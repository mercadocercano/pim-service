package port

import (
	"context"

	"saas-mt-pim-service/src/shared/domain/entity"
)

// NotificationService define las operaciones para enviar notificaciones
type NotificationService interface {
	// SendWebhook envía una notificación webhook
	SendWebhook(ctx context.Context, webhookURL string, payload interface{}) error

	// SendEmail envía una notificación por email
	SendEmail(ctx context.Context, to string, subject string, body string) error

	// NotifyImportJobComplete notifica la finalización de un trabajo de importación
	NotifyImportJobComplete(ctx context.Context, job *entity.ImportJob) error
}

// ImportJobNotification representa la estructura de notificación para trabajos de importación
type ImportJobNotification struct {
	JobID         string  `json:"job_id"`
	TenantID      string  `json:"tenant_id"`
	Type          string  `json:"type"`
	Status        string  `json:"status"`
	FileName      string  `json:"file_name"`
	TotalRecords  int     `json:"total_records"`
	SuccessCount  int     `json:"success_count"`
	FailureCount  int     `json:"failure_count"`
	Duration      string  `json:"duration"`
	ResultFileURL *string `json:"result_file_url,omitempty"`
	ErrorMessage  *string `json:"error_message,omitempty"`
	CompletedAt   string  `json:"completed_at"`
}
