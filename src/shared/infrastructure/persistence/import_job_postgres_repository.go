package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"saas-mt-pim-service/src/shared/domain/entity"
	"saas-mt-pim-service/src/shared/domain/port"

	"github.com/google/uuid"
)

// ImportJobPostgresRepository implementa el repositorio de trabajos de importación usando PostgreSQL
type ImportJobPostgresRepository struct {
	db *sql.DB
}

// NewImportJobPostgresRepository crea una nueva instancia del repositorio
func NewImportJobPostgresRepository(db *sql.DB) port.ImportJobRepository {
	return &ImportJobPostgresRepository{db: db}
}

// Create crea un nuevo trabajo de importación
func (r *ImportJobPostgresRepository) Create(ctx context.Context, job *entity.ImportJob) error {
	query := `
		INSERT INTO import_jobs (
			id, tenant_id, type, status, file_name, file_size_bytes,
			total_records, processed_records, success_count, failure_count,
			progress, started_at, completed_at, created_at, updated_at,
			created_by, error_message, result_file_url, notification_sent,
			webhook_url, email_notify
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21
		)
	`

	_, err := r.db.ExecContext(ctx, query,
		job.ID, job.TenantID, job.Type, job.Status, job.FileName, job.FileSizeBytes,
		job.TotalRecords, job.ProcessedRecords, job.SuccessCount, job.FailureCount,
		job.Progress, job.StartedAt, job.CompletedAt, job.CreatedAt, job.UpdatedAt,
		job.CreatedBy, job.ErrorMessage, job.ResultFileURL, job.NotificationSent,
		job.WebhookURL, job.EmailNotify,
	)

	return err
}

// Update actualiza un trabajo existente
func (r *ImportJobPostgresRepository) Update(ctx context.Context, job *entity.ImportJob) error {
	query := `
		UPDATE import_jobs SET
			status = $2, total_records = $3, processed_records = $4,
			success_count = $5, failure_count = $6, progress = $7,
			started_at = $8, completed_at = $9, updated_at = $10,
			error_message = $11, result_file_url = $12, notification_sent = $13
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		job.ID, job.Status, job.TotalRecords, job.ProcessedRecords,
		job.SuccessCount, job.FailureCount, job.Progress,
		job.StartedAt, job.CompletedAt, job.UpdatedAt,
		job.ErrorMessage, job.ResultFileURL, job.NotificationSent,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("import job with id %s not found", job.ID)
	}

	return nil
}

// FindByID busca un trabajo por su ID
func (r *ImportJobPostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.ImportJob, error) {
	query := `
		SELECT 
			id, tenant_id, type, status, file_name, file_size_bytes,
			total_records, processed_records, success_count, failure_count,
			progress, started_at, completed_at, created_at, updated_at,
			created_by, error_message, result_file_url, notification_sent,
			webhook_url, email_notify
		FROM import_jobs
		WHERE id = $1
	`

	job := &entity.ImportJob{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&job.ID, &job.TenantID, &job.Type, &job.Status, &job.FileName, &job.FileSizeBytes,
		&job.TotalRecords, &job.ProcessedRecords, &job.SuccessCount, &job.FailureCount,
		&job.Progress, &job.StartedAt, &job.CompletedAt, &job.CreatedAt, &job.UpdatedAt,
		&job.CreatedBy, &job.ErrorMessage, &job.ResultFileURL, &job.NotificationSent,
		&job.WebhookURL, &job.EmailNotify,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("import job with id %s not found", id)
	}

	return job, err
}

// FindByTenantID busca trabajos por tenant
func (r *ImportJobPostgresRepository) FindByTenantID(ctx context.Context, tenantID string, limit int) ([]*entity.ImportJob, error) {
	query := `
		SELECT 
			id, tenant_id, type, status, file_name, file_size_bytes,
			total_records, processed_records, success_count, failure_count,
			progress, started_at, completed_at, created_at, updated_at,
			created_by, error_message, result_file_url, notification_sent,
			webhook_url, email_notify
		FROM import_jobs
		WHERE tenant_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []*entity.ImportJob
	for rows.Next() {
		job := &entity.ImportJob{}
		err := rows.Scan(
			&job.ID, &job.TenantID, &job.Type, &job.Status, &job.FileName, &job.FileSizeBytes,
			&job.TotalRecords, &job.ProcessedRecords, &job.SuccessCount, &job.FailureCount,
			&job.Progress, &job.StartedAt, &job.CompletedAt, &job.CreatedAt, &job.UpdatedAt,
			&job.CreatedBy, &job.ErrorMessage, &job.ResultFileURL, &job.NotificationSent,
			&job.WebhookURL, &job.EmailNotify,
		)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

// FindPendingJobs busca trabajos pendientes de procesar
func (r *ImportJobPostgresRepository) FindPendingJobs(ctx context.Context, limit int) ([]*entity.ImportJob, error) {
	query := `
		SELECT 
			id, tenant_id, type, status, file_name, file_size_bytes,
			total_records, processed_records, success_count, failure_count,
			progress, started_at, completed_at, created_at, updated_at,
			created_by, error_message, result_file_url, notification_sent,
			webhook_url, email_notify
		FROM import_jobs
		WHERE status = $1
		ORDER BY created_at ASC
		LIMIT $2
	`

	rows, err := r.db.QueryContext(ctx, query, entity.ImportJobStatusPending, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []*entity.ImportJob
	for rows.Next() {
		job := &entity.ImportJob{}
		err := rows.Scan(
			&job.ID, &job.TenantID, &job.Type, &job.Status, &job.FileName, &job.FileSizeBytes,
			&job.TotalRecords, &job.ProcessedRecords, &job.SuccessCount, &job.FailureCount,
			&job.Progress, &job.StartedAt, &job.CompletedAt, &job.CreatedAt, &job.UpdatedAt,
			&job.CreatedBy, &job.ErrorMessage, &job.ResultFileURL, &job.NotificationSent,
			&job.WebhookURL, &job.EmailNotify,
		)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

// FindJobsNeedingNotification busca trabajos que necesitan enviar notificación
func (r *ImportJobPostgresRepository) FindJobsNeedingNotification(ctx context.Context, limit int) ([]*entity.ImportJob, error) {
	query := `
		SELECT 
			id, tenant_id, type, status, file_name, file_size_bytes,
			total_records, processed_records, success_count, failure_count,
			progress, started_at, completed_at, created_at, updated_at,
			created_by, error_message, result_file_url, notification_sent,
			webhook_url, email_notify
		FROM import_jobs
		WHERE status IN ($1, $2, $3)
			AND notification_sent = false
			AND (webhook_url IS NOT NULL OR email_notify IS NOT NULL)
		ORDER BY completed_at ASC
		LIMIT $4
	`

	rows, err := r.db.QueryContext(ctx, query, 
		entity.ImportJobStatusCompleted,
		entity.ImportJobStatusFailed,
		entity.ImportJobStatusCancelled,
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []*entity.ImportJob
	for rows.Next() {
		job := &entity.ImportJob{}
		err := rows.Scan(
			&job.ID, &job.TenantID, &job.Type, &job.Status, &job.FileName, &job.FileSizeBytes,
			&job.TotalRecords, &job.ProcessedRecords, &job.SuccessCount, &job.FailureCount,
			&job.Progress, &job.StartedAt, &job.CompletedAt, &job.CreatedAt, &job.UpdatedAt,
			&job.CreatedBy, &job.ErrorMessage, &job.ResultFileURL, &job.NotificationSent,
			&job.WebhookURL, &job.EmailNotify,
		)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

// FindActiveJobsByTenant busca trabajos activos de un tenant
func (r *ImportJobPostgresRepository) FindActiveJobsByTenant(ctx context.Context, tenantID string) ([]*entity.ImportJob, error) {
	query := `
		SELECT 
			id, tenant_id, type, status, file_name, file_size_bytes,
			total_records, processed_records, success_count, failure_count,
			progress, started_at, completed_at, created_at, updated_at,
			created_by, error_message, result_file_url, notification_sent,
			webhook_url, email_notify
		FROM import_jobs
		WHERE tenant_id = $1
			AND status IN ($2, $3)
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID,
		entity.ImportJobStatusPending,
		entity.ImportJobStatusProcessing,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []*entity.ImportJob
	for rows.Next() {
		job := &entity.ImportJob{}
		err := rows.Scan(
			&job.ID, &job.TenantID, &job.Type, &job.Status, &job.FileName, &job.FileSizeBytes,
			&job.TotalRecords, &job.ProcessedRecords, &job.SuccessCount, &job.FailureCount,
			&job.Progress, &job.StartedAt, &job.CompletedAt, &job.CreatedAt, &job.UpdatedAt,
			&job.CreatedBy, &job.ErrorMessage, &job.ResultFileURL, &job.NotificationSent,
			&job.WebhookURL, &job.EmailNotify,
		)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}