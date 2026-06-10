package port

import (
	"context"

	"saas-mt-pim-service/src/shared/domain/entity"

	"github.com/google/uuid"
)

// ImportJobRepository define las operaciones de persistencia para trabajos de importación
type ImportJobRepository interface {
	// Create crea un nuevo trabajo de importación
	Create(ctx context.Context, job *entity.ImportJob) error

	// Update actualiza un trabajo existente
	Update(ctx context.Context, job *entity.ImportJob) error

	// FindByID busca un trabajo por su ID
	FindByID(ctx context.Context, id uuid.UUID) (*entity.ImportJob, error)

	// FindByTenantID busca trabajos por tenant
	FindByTenantID(ctx context.Context, tenantID string, limit int) ([]*entity.ImportJob, error)

	// FindPendingJobs busca trabajos pendientes de procesar
	FindPendingJobs(ctx context.Context, limit int) ([]*entity.ImportJob, error)

	// FindJobsNeedingNotification busca trabajos que necesitan enviar notificación
	FindJobsNeedingNotification(ctx context.Context, limit int) ([]*entity.ImportJob, error)

	// FindActiveJobsByTenant busca trabajos activos de un tenant
	FindActiveJobsByTenant(ctx context.Context, tenantID string) ([]*entity.ImportJob, error)
}
