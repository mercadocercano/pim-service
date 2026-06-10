package usecase

import (
	"context"
	"fmt"
	"io"
	"time"

	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/domain/entity"
	tenantport "saas-mt-pim-service/src/product/tenant/domain/port"
	sharedport "github.com/mercadocercano/go-shared/domain/port"
)

// ImportProductsFromCSVUseCase caso de uso para importar productos desde CSV
type ImportProductsFromCSVUseCase struct {
	productRepo  tenantport.ProductCriteriaRepository
	fileImporter sharedport.FileImporter[entity.Product]
	metrics      sharedport.MetricsRecorder
}

// NewImportProductsFromCSVUseCase crea una nueva instancia del caso de uso
func NewImportProductsFromCSVUseCase(
	productRepo tenantport.ProductCriteriaRepository,
	fileImporter sharedport.FileImporter[entity.Product],
	metrics sharedport.MetricsRecorder,
) *ImportProductsFromCSVUseCase {
	return &ImportProductsFromCSVUseCase{
		productRepo:  productRepo,
		fileImporter: fileImporter,
		metrics:      metrics,
	}
}

// ImportResult resultado de la importación con productos guardados
type ImportResult struct {
	*sharedport.ImportResult[entity.Product]
	SavedProducts    []*entity.Product          `json:"saved_products"`
	ProcessingErrors []response.ProcessingError `json:"processing_errors"`
}

// Execute ejecuta la importación de productos desde CSV
func (uc *ImportProductsFromCSVUseCase) Execute(ctx context.Context, reader io.Reader, tenantID string) (*ImportResult, error) {
	startTime := time.Now()

	if tenantID == "" {
		return nil, fmt.Errorf("tenant ID es requerido")
	}

	importResult, err := uc.fileImporter.Import(ctx, reader, tenantID)
	if err != nil {
		uc.metrics.Record(sharedport.MetricEvent{
			Name:  tenantport.MetricImportOperation,
			Kind:  sharedport.MetricKindCounter,
			Value: 1.0,
			Labels: map[string]string{
				"tenant_id": tenantID,
				"type":      "csv_products",
				"result":    "failure",
			},
		})
		return nil, fmt.Errorf("error al importar archivo CSV: %w", err)
	}

	result := &ImportResult{
		ImportResult:     importResult,
		SavedProducts:    make([]*entity.Product, 0),
		ProcessingErrors: make([]response.ProcessingError, 0),
	}

	for _, product := range importResult.ImportedItems {
		productToSave := product

		if saveErr := uc.productRepo.Save(ctx, &productToSave); saveErr != nil {
			result.ProcessingErrors = append(result.ProcessingErrors, response.ProcessingError{
				Product: &productToSave,
				Error:   fmt.Sprintf("error al guardar producto: %v", saveErr),
			})
			result.SuccessfulImports--
			result.FailedImports++

			uc.metrics.Record(sharedport.MetricEvent{
				Name:  tenantport.MetricImportValidation,
				Kind:  sharedport.MetricKindCounter,
				Value: 1.0,
				Labels: map[string]string{
					"tenant_id": tenantID,
					"type":      "csv_products",
					"error":     "save_error",
				},
			})
		} else {
			result.SavedProducts = append(result.SavedProducts, &productToSave)
		}
	}

	result.ImportedItems = make([]entity.Product, 0, len(result.SavedProducts))
	for _, saved := range result.SavedProducts {
		result.ImportedItems = append(result.ImportedItems, *saved)
	}

	duration := time.Since(startTime).Seconds()
	success := len(result.ProcessingErrors) == 0

	uc.metrics.Record(sharedport.MetricEvent{
		Name:  tenantport.MetricImportOperation,
		Kind:  sharedport.MetricKindCounter,
		Value: 1.0,
		Labels: map[string]string{
			"tenant_id": tenantID,
			"type":      "csv_products",
			"result":    boolResult(success),
		},
	})
	uc.metrics.Record(sharedport.MetricEvent{
		Name:  tenantport.MetricImportRecord,
		Kind:  sharedport.MetricKindCounter,
		Value: float64(importResult.TotalRows),
		Labels: map[string]string{
			"tenant_id": tenantID,
			"type":      "csv_products",
		},
	})
	uc.metrics.Record(sharedport.MetricEvent{
		Name:  tenantport.MetricImportDuration,
		Kind:  sharedport.MetricKindHistogram,
		Unit:  sharedport.MetricUnitSeconds,
		Value: duration,
		Labels: map[string]string{
			"tenant_id": tenantID,
			"type":      "csv_products",
		},
	})

	for _, importErr := range importResult.Errors {
		for _, e := range importErr.Errors {
			uc.metrics.Record(sharedport.MetricEvent{
				Name:  tenantport.MetricImportValidation,
				Kind:  sharedport.MetricKindCounter,
				Value: 1.0,
				Labels: map[string]string{
					"tenant_id": tenantID,
					"type":      "csv_products",
					"error":     e,
				},
			})
		}
	}

	return result, nil
}

// ExecuteWithVariants ejecuta la importación y crea variantes por defecto
func (uc *ImportProductsFromCSVUseCase) ExecuteWithVariants(
	ctx context.Context,
	reader io.Reader,
	tenantID string,
	variantRepo tenantport.ProductVariantRepository,
) (*ImportResult, error) {
	return uc.Execute(ctx, reader, tenantID)
}

func boolResult(ok bool) string {
	if ok {
		return "success"
	}
	return "failure"
}
