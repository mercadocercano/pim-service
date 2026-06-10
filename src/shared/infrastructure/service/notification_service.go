package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"saas-mt-pim-service/src/shared/domain/entity"
	"saas-mt-pim-service/src/shared/domain/port"
)

// HTTPNotificationService implementa el servicio de notificaciones usando HTTP
type HTTPNotificationService struct {
	httpClient   *http.Client
	emailService EmailService
}

// EmailService define la interfaz para enviar emails
type EmailService interface {
	Send(to, subject, body string) error
}

// NewHTTPNotificationService crea una nueva instancia del servicio
func NewHTTPNotificationService(emailService EmailService) port.NotificationService {
	return &HTTPNotificationService{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		emailService: emailService,
	}
}

// SendWebhook envía una notificación webhook
func (s *HTTPNotificationService) SendWebhook(ctx context.Context, webhookURL string, payload interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling webhook payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating webhook request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "PIM-Import-Service/1.0")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending webhook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("webhook returned non-2xx status: %d", resp.StatusCode)
	}

	return nil
}

// SendEmail envía una notificación por email
func (s *HTTPNotificationService) SendEmail(ctx context.Context, to string, subject string, body string) error {
	if s.emailService == nil {
		return fmt.Errorf("email service not configured")
	}

	return s.emailService.Send(to, subject, body)
}

// NotifyImportJobComplete notifica la finalización de un trabajo de importación
func (s *HTTPNotificationService) NotifyImportJobComplete(ctx context.Context, job *entity.ImportJob) error {
	// Preparar payload de notificación
	notification := port.ImportJobNotification{
		JobID:         job.ID.String(),
		TenantID:      job.TenantID,
		Type:          string(job.Type),
		Status:        string(job.Status),
		FileName:      job.FileName,
		TotalRecords:  job.TotalRecords,
		SuccessCount:  job.SuccessCount,
		FailureCount:  job.FailureCount,
		Duration:      job.Duration().String(),
		ResultFileURL: job.ResultFileURL,
		ErrorMessage:  job.ErrorMessage,
		CompletedAt:   job.CompletedAt.Format(time.RFC3339),
	}

	var errors []error

	// Enviar webhook si está configurado
	if job.WebhookURL != nil && *job.WebhookURL != "" {
		if err := s.SendWebhook(ctx, *job.WebhookURL, notification); err != nil {
			errors = append(errors, fmt.Errorf("webhook notification failed: %w", err))
		}
	}

	// Enviar email si está configurado
	if job.EmailNotify != nil && *job.EmailNotify != "" {
		subject := fmt.Sprintf("Import Job %s - %s", job.ID, job.Status)
		body := s.buildEmailBody(job, notification)

		if err := s.SendEmail(ctx, *job.EmailNotify, subject, body); err != nil {
			errors = append(errors, fmt.Errorf("email notification failed: %w", err))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("notification errors: %v", errors)
	}

	return nil
}

// buildEmailBody construye el cuerpo del email
func (s *HTTPNotificationService) buildEmailBody(job *entity.ImportJob, notification port.ImportJobNotification) string {
	status := "✅ Completed"
	if job.Status == entity.ImportJobStatusFailed {
		status = "❌ Failed"
	} else if job.Status == entity.ImportJobStatusCancelled {
		status = "⚠️ Cancelled"
	}

	body := fmt.Sprintf(`
Import Job Notification

Job ID: %s
Status: %s
File: %s
Type: %s

Results:
- Total Records: %d
- Successfully Processed: %d
- Failed: %d
- Duration: %s

`, job.ID, status, job.FileName, job.Type,
		job.TotalRecords, job.SuccessCount, job.FailureCount,
		notification.Duration)

	if job.ErrorMessage != nil {
		body += fmt.Sprintf("Error: %s\n\n", *job.ErrorMessage)
	}

	if job.ResultFileURL != nil {
		body += fmt.Sprintf("Download results: %s\n", *job.ResultFileURL)
	}

	body += "\n---\nThis is an automated notification from PIM Import Service"

	return body
}

// MockEmailService es una implementación mock para testing
type MockEmailService struct{}

func (m *MockEmailService) Send(to, subject, body string) error {
	// Log para desarrollo, en producción esto enviaría emails reales
	fmt.Printf("Mock Email - To: %s, Subject: %s\n", to, subject)
	return nil
}
