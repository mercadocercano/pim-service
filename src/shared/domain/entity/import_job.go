package entity

import (
	"time"

	"github.com/google/uuid"
)

// ImportJobStatus representa el estado de un trabajo de importación
type ImportJobStatus string

const (
	ImportJobStatusPending    ImportJobStatus = "pending"
	ImportJobStatusProcessing ImportJobStatus = "processing"
	ImportJobStatusCompleted  ImportJobStatus = "completed"
	ImportJobStatusFailed     ImportJobStatus = "failed"
	ImportJobStatusCancelled  ImportJobStatus = "cancelled"
)

// ImportJobType representa el tipo de importación
type ImportJobType string

const (
	ImportJobTypeCSVProducts  ImportJobType = "csv_products"
	ImportJobTypeJSONProducts ImportJobType = "json_products"
	ImportJobTypeBatchCreate  ImportJobType = "batch_create"
)

// ImportJob representa un trabajo de importación asíncrono
type ImportJob struct {
	ID               uuid.UUID       `json:"id" db:"id"`
	TenantID         string          `json:"tenant_id" db:"tenant_id"`
	Type             ImportJobType   `json:"type" db:"type"`
	Status           ImportJobStatus `json:"status" db:"status"`
	FileName         string          `json:"file_name" db:"file_name"`
	FileSizeBytes    int64           `json:"file_size_bytes" db:"file_size_bytes"`
	TotalRecords     int             `json:"total_records" db:"total_records"`
	ProcessedRecords int             `json:"processed_records" db:"processed_records"`
	SuccessCount     int             `json:"success_count" db:"success_count"`
	FailureCount     int             `json:"failure_count" db:"failure_count"`
	Progress         float64         `json:"progress" db:"progress"` // 0-100
	StartedAt        *time.Time      `json:"started_at" db:"started_at"`
	CompletedAt      *time.Time      `json:"completed_at" db:"completed_at"`
	CreatedAt        time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" db:"updated_at"`
	CreatedBy        string          `json:"created_by" db:"created_by"`
	ErrorMessage     *string         `json:"error_message,omitempty" db:"error_message"`
	ResultFileURL    *string         `json:"result_file_url,omitempty" db:"result_file_url"`
	NotificationSent bool            `json:"notification_sent" db:"notification_sent"`
	WebhookURL       *string         `json:"webhook_url,omitempty" db:"webhook_url"`
	EmailNotify      *string         `json:"email_notify,omitempty" db:"email_notify"`
}

// NewImportJob crea un nuevo trabajo de importación
func NewImportJob(tenantID, fileName, createdBy string, jobType ImportJobType, fileSizeBytes int64) *ImportJob {
	now := time.Now()
	return &ImportJob{
		ID:               uuid.New(),
		TenantID:         tenantID,
		Type:             jobType,
		Status:           ImportJobStatusPending,
		FileName:         fileName,
		FileSizeBytes:    fileSizeBytes,
		TotalRecords:     0,
		ProcessedRecords: 0,
		SuccessCount:     0,
		FailureCount:     0,
		Progress:         0,
		CreatedAt:        now,
		UpdatedAt:        now,
		CreatedBy:        createdBy,
		NotificationSent: false,
	}
}

// Start marca el trabajo como iniciado
func (j *ImportJob) Start() {
	now := time.Now()
	j.Status = ImportJobStatusProcessing
	j.StartedAt = &now
	j.UpdatedAt = now
}

// UpdateProgress actualiza el progreso del trabajo
func (j *ImportJob) UpdateProgress(processed, success, failure int) {
	j.ProcessedRecords = processed
	j.SuccessCount = success
	j.FailureCount = failure
	if j.TotalRecords > 0 {
		j.Progress = float64(processed) / float64(j.TotalRecords) * 100
	}
	j.UpdatedAt = time.Now()
}

// Complete marca el trabajo como completado
func (j *ImportJob) Complete(resultFileURL string) {
	now := time.Now()
	j.Status = ImportJobStatusCompleted
	j.CompletedAt = &now
	j.UpdatedAt = now
	j.Progress = 100
	if resultFileURL != "" {
		j.ResultFileURL = &resultFileURL
	}
}

// Fail marca el trabajo como fallido
func (j *ImportJob) Fail(errorMessage string) {
	now := time.Now()
	j.Status = ImportJobStatusFailed
	j.CompletedAt = &now
	j.UpdatedAt = now
	j.ErrorMessage = &errorMessage
}

// Cancel cancela el trabajo
func (j *ImportJob) Cancel() {
	now := time.Now()
	j.Status = ImportJobStatusCancelled
	j.CompletedAt = &now
	j.UpdatedAt = now
}

// IsFinished indica si el trabajo ha terminado
func (j *ImportJob) IsFinished() bool {
	return j.Status == ImportJobStatusCompleted ||
		j.Status == ImportJobStatusFailed ||
		j.Status == ImportJobStatusCancelled
}

// NeedsNotification indica si el trabajo necesita enviar notificación
func (j *ImportJob) NeedsNotification() bool {
	return j.IsFinished() && !j.NotificationSent && (j.WebhookURL != nil || j.EmailNotify != nil)
}

// MarkNotificationSent marca la notificación como enviada
func (j *ImportJob) MarkNotificationSent() {
	j.NotificationSent = true
	j.UpdatedAt = time.Now()
}

// Duration calcula la duración del trabajo
func (j *ImportJob) Duration() time.Duration {
	if j.StartedAt == nil {
		return 0
	}
	endTime := time.Now()
	if j.CompletedAt != nil {
		endTime = *j.CompletedAt
	}
	return endTime.Sub(*j.StartedAt)
}