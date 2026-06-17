package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"

	pimport "saas-mt-pim-service/src/pim/domain/port"
	"saas-mt-pim-service/src/product/tenant/domain/entity"
	tenantport "saas-mt-pim-service/src/product/tenant/domain/port"
	sharedEntity "saas-mt-pim-service/src/shared/domain/entity"
	localport "saas-mt-pim-service/src/shared/domain/port"
	sharedport "github.com/hornosg/go-shared/domain/port"
)

// ImportProductsAsyncUseCase maneja importaciones asíncronas de productos
type ImportProductsAsyncUseCase struct {
	productRepo     tenantport.ProductCriteriaRepository
	fileImporter    sharedport.FileImporter[entity.Product]
	importJobRepo   localport.ImportJobRepository
	notificationSvc localport.NotificationService
	fileStorage     FileStorageService
	metrics         sharedport.MetricsRecorder
	logger          pimport.PIMEventLogger
}

// FileStorageService define operaciones de almacenamiento de archivos
type FileStorageService interface {
	Store(ctx context.Context, fileName string, reader io.Reader) (string, error)
	StoreResults(ctx context.Context, jobID string, results interface{}) (string, error)
	GetReader(ctx context.Context, fileURL string) (io.ReadCloser, error)
}

// NewImportProductsAsyncUseCase crea una nueva instancia
func NewImportProductsAsyncUseCase(
	productRepo tenantport.ProductCriteriaRepository,
	fileImporter sharedport.FileImporter[entity.Product],
	importJobRepo localport.ImportJobRepository,
	notificationSvc localport.NotificationService,
	fileStorage FileStorageService,
	metrics sharedport.MetricsRecorder,
) *ImportProductsAsyncUseCase {
	return &ImportProductsAsyncUseCase{
		productRepo:     productRepo,
		fileImporter:    fileImporter,
		importJobRepo:   importJobRepo,
		notificationSvc: notificationSvc,
		fileStorage:     fileStorage,
		metrics:         metrics,
	}
}

// NewImportProductsAsyncUseCaseWithLogger crea el use case con logger canónico inyectado.
func NewImportProductsAsyncUseCaseWithLogger(
	productRepo tenantport.ProductCriteriaRepository,
	fileImporter sharedport.FileImporter[entity.Product],
	importJobRepo localport.ImportJobRepository,
	notificationSvc localport.NotificationService,
	fileStorage FileStorageService,
	metrics sharedport.MetricsRecorder,
	logger pimport.PIMEventLogger,
) *ImportProductsAsyncUseCase {
	return &ImportProductsAsyncUseCase{
		productRepo:     productRepo,
		fileImporter:    fileImporter,
		importJobRepo:   importJobRepo,
		notificationSvc: notificationSvc,
		fileStorage:     fileStorage,
		metrics:         metrics,
		logger:          logger,
	}
}

// logEvent emite un evento canónico si hay logger inyectado (nil-safe).
func (uc *ImportProductsAsyncUseCase) logEvent(e pimport.PIMEvent) {
	if uc.logger != nil {
		uc.logger.Log(e)
	}
}

// ImportAsyncRequest representa la solicitud de importación asíncrona
type ImportAsyncRequest struct {
	FileName    string
	FileReader  io.Reader
	FileSize    int64
	TenantID    string
	UserID      string
	WebhookURL  *string
	EmailNotify *string
}

// ImportAsyncResponse representa la respuesta de importación asíncrona
type ImportAsyncResponse struct {
	JobID    string `json:"job_id"`
	Status   string `json:"status"`
	Message  string `json:"message"`
	TrackURL string `json:"track_url"`
}

// StartImport inicia una importación asíncrona
func (uc *ImportProductsAsyncUseCase) StartImport(ctx context.Context, req *ImportAsyncRequest) (*ImportAsyncResponse, error) {
	job := sharedEntity.NewImportJob(
		req.TenantID,
		req.FileName,
		req.UserID,
		sharedEntity.ImportJobTypeCSVProducts,
		req.FileSize,
	)

	if req.WebhookURL != nil {
		job.WebhookURL = req.WebhookURL
	}
	if req.EmailNotify != nil {
		job.EmailNotify = req.EmailNotify
	}

	fileURL, err := uc.fileStorage.Store(ctx, req.FileName, req.FileReader)
	if err != nil {
		return nil, fmt.Errorf("error storing file: %w", err)
	}

	if err := uc.importJobRepo.Create(ctx, job); err != nil {
		return nil, fmt.Errorf("error creating import job: %w", err)
	}

	go uc.processImportJob(context.Background(), job.ID, fileURL)

	return &ImportAsyncResponse{
		JobID:    job.ID.String(),
		Status:   string(job.Status),
		Message:  "Import job created successfully. You will be notified when it completes.",
		TrackURL: fmt.Sprintf("/api/v1/import-jobs/%s", job.ID),
	}, nil
}

// processImportJob procesa el trabajo de importación de forma asíncrona
func (uc *ImportProductsAsyncUseCase) processImportJob(ctx context.Context, jobID uuid.UUID, fileURL string) {
	startTime := time.Now()

	job, err := uc.importJobRepo.FindByID(ctx, jobID)
	if err != nil {
		return
	}

	job.Start()
	if err := uc.importJobRepo.Update(ctx, job); err != nil {
		return
	}

	fileReader, err := uc.fileStorage.GetReader(ctx, fileURL)
	if err != nil {
		uc.handleJobError(ctx, job, fmt.Errorf("error reading file: %w", err))
		return
	}
	defer fileReader.Close()

	importResult, err := uc.fileImporter.Import(ctx, fileReader, job.TenantID)
	if err != nil {
		uc.handleJobError(ctx, job, fmt.Errorf("error importing file: %w", err))
		return
	}

	job.TotalRecords = importResult.TotalRows
	job.UpdateProgress(len(importResult.ImportedItems), 0, 0)
	if err := uc.importJobRepo.Update(ctx, job); err != nil {
		return
	}

	savedProducts := []interface{}{}
	processingErrors := []interface{}{}
	successCount := 0
	failureCount := 0

	for i, product := range importResult.ImportedItems {
		productToSave := product

		if saveErr := uc.productRepo.Save(ctx, &productToSave); saveErr != nil {
			failureCount++
			processingErrors = append(processingErrors, map[string]interface{}{
				"row":     i + 1,
				"product": productToSave,
				"error":   saveErr.Error(),
			})
			uc.metrics.Record(sharedport.MetricEvent{
				Name:  tenantport.MetricImportValidation,
				Kind:  sharedport.MetricKindCounter,
				Value: 1.0,
				Labels: map[string]string{
					"tenant_id": job.TenantID,
					"type":      "csv_products_async",
					"error":     "save_error",
				},
			})
		} else {
			successCount++
			savedProducts = append(savedProducts, productToSave)
		}

		if (i+1)%100 == 0 || i == len(importResult.ImportedItems)-1 {
			job.UpdateProgress(i+1, successCount, failureCount)
			uc.importJobRepo.Update(ctx, job)
		}
	}

	for _, importErr := range importResult.Errors {
		failureCount++
		processingErrors = append(processingErrors, importErr)
	}

	results := map[string]interface{}{
		"summary": map[string]interface{}{
			"total_processed": job.TotalRecords,
			"success_count":   successCount,
			"failure_count":   failureCount,
			"duration":        time.Since(startTime).String(),
		},
		"saved_products": savedProducts,
		"errors":         processingErrors,
	}

	resultURL, err := uc.fileStorage.StoreResults(ctx, jobID.String(), results)
	if err != nil {
		uc.logEvent(pimport.PIMEvent{
			Event:  "pim.import_failed",
			JobID:  jobID.String(),
			Reason: fmt.Sprintf("storing results: %v", err),
		})
	}

	job.UpdateProgress(job.TotalRecords, successCount, failureCount)
	job.Complete(resultURL)
	if err := uc.importJobRepo.Update(ctx, job); err != nil {
		return
	}

	duration := time.Since(startTime).Seconds()
	uc.metrics.Record(sharedport.MetricEvent{
		Name:  tenantport.MetricImportOperation,
		Kind:  sharedport.MetricKindCounter,
		Value: 1.0,
		Labels: map[string]string{
			"tenant_id": job.TenantID,
			"type":      "csv_products_async",
			"result":    boolResult(failureCount == 0),
		},
	})
	uc.metrics.Record(sharedport.MetricEvent{
		Name:  tenantport.MetricImportDuration,
		Kind:  sharedport.MetricKindHistogram,
		Unit:  sharedport.MetricUnitSeconds,
		Value: duration,
		Labels: map[string]string{
			"tenant_id": job.TenantID,
			"type":      "csv_products_async",
		},
	})

	if job.NeedsNotification() {
		if notifyErr := uc.notificationSvc.NotifyImportJobComplete(ctx, job); notifyErr != nil {
			uc.logEvent(pimport.PIMEvent{
				Event:    "pim.import_failed",
				TenantID: job.TenantID,
				JobID:    jobID.String(),
				Reason:   fmt.Sprintf("notification failed: %v", notifyErr),
			})
		} else {
			job.MarkNotificationSent()
			uc.importJobRepo.Update(ctx, job)
		}
	}

	if failureCount == 0 {
		uc.logEvent(pimport.PIMEvent{
			Event:    "pim.import_from_global_catalog_completed",
			TenantID: job.TenantID,
			JobID:    jobID.String(),
			Count:    successCount,
		})
	} else {
		uc.logEvent(pimport.PIMEvent{
			Event:    "pim.import_failed",
			TenantID: job.TenantID,
			JobID:    jobID.String(),
			Count:    failureCount,
			Reason:   fmt.Sprintf("%d/%d records failed", failureCount, job.TotalRecords),
		})
	}
}

// handleJobError maneja errores durante el procesamiento
func (uc *ImportProductsAsyncUseCase) handleJobError(ctx context.Context, job *sharedEntity.ImportJob, err error) {
	job.Fail(err.Error())
	uc.importJobRepo.Update(ctx, job)

	uc.metrics.Record(sharedport.MetricEvent{
		Name:  tenantport.MetricImportOperation,
		Kind:  sharedport.MetricKindCounter,
		Value: 1.0,
		Labels: map[string]string{
			"tenant_id": job.TenantID,
			"type":      "csv_products_async",
			"result":    "failure",
		},
	})

	if job.NeedsNotification() {
		if notifyErr := uc.notificationSvc.NotifyImportJobComplete(ctx, job); notifyErr == nil {
			job.MarkNotificationSent()
			uc.importJobRepo.Update(ctx, job)
		}
	}
}

// GetImportJobStatus obtiene el estado de un trabajo de importación
func (uc *ImportProductsAsyncUseCase) GetImportJobStatus(ctx context.Context, jobID uuid.UUID) (*sharedEntity.ImportJob, error) {
	return uc.importJobRepo.FindByID(ctx, jobID)
}

// CancelImportJob cancela un trabajo de importación
func (uc *ImportProductsAsyncUseCase) CancelImportJob(ctx context.Context, jobID uuid.UUID) error {
	job, err := uc.importJobRepo.FindByID(ctx, jobID)
	if err != nil {
		return err
	}

	if job.IsFinished() {
		return fmt.Errorf("job already finished with status: %s", job.Status)
	}

	job.Cancel()
	return uc.importJobRepo.Update(ctx, job)
}

// LocalFileStorageService implementación local para desarrollo
type LocalFileStorageService struct {
	basePath string
}

func NewLocalFileStorageService(basePath string) FileStorageService {
	return &LocalFileStorageService{basePath: basePath}
}

func (s *LocalFileStorageService) Store(ctx context.Context, fileName string, reader io.Reader) (string, error) {
	uploadPath := filepath.Join(s.basePath, "uploads")
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		return "", err
	}

	uniqueName := fmt.Sprintf("%s_%s", uuid.New().String(), fileName)
	filePath := filepath.Join(uploadPath, uniqueName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(file, reader); err != nil {
		return "", err
	}

	return filePath, nil
}

func (s *LocalFileStorageService) StoreResults(ctx context.Context, jobID string, results interface{}) (string, error) {
	resultsPath := filepath.Join(s.basePath, "results")
	if err := os.MkdirAll(resultsPath, 0755); err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%s_results.json", jobID)
	filePath := filepath.Join(resultsPath, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		return "", err
	}

	return filePath, nil
}

func (s *LocalFileStorageService) GetReader(ctx context.Context, fileURL string) (io.ReadCloser, error) {
	return os.Open(fileURL)
}
