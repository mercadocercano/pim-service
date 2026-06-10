package port

import (
	"context"
	"saas-mt-pim-service/src/quickstart/domain/entity"
)

// QuickstartHistoryRepository define los métodos para persistir el historial del wizard
type QuickstartHistoryRepository interface {
	// FindByTenantID busca un wizard activo (no completado) para un tenant
	FindActiveByTenantID(ctx context.Context, tenantID string) (*entity.TenantQuickstartHistory, error)

	// FindByID busca un registro por su ID
	FindByID(ctx context.Context, id string) (*entity.TenantQuickstartHistory, error)

	// Create crea un nuevo registro de historial
	Create(ctx context.Context, history *entity.TenantQuickstartHistory) error

	// Update actualiza un registro existente
	Update(ctx context.Context, history *entity.TenantQuickstartHistory) error

	// MarkAsCompleted marca un wizard como completado
	MarkAsCompleted(ctx context.Context, id string) error

	// FindAllByTenantID obtiene todo el historial de un tenant
	FindAllByTenantID(ctx context.Context, tenantID string) ([]*entity.TenantQuickstartHistory, error)
}
