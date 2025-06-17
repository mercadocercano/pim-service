package port

import (
	"context"
	"pim/src/global_catalog/domain/entity"
	"pim/src/shared/domain/criteria"
)

// EanregistryRepository define los métodos para persistir Eanregistry
type EanregistryRepository interface {
	Create(ctx context.Context, ean_registry *entity.Eanregistry) error
	Update(ctx context.Context, ean_registry *entity.Eanregistry) error
	FindByID(ctx context.Context, id string, tenantID string) (*entity.Eanregistry, error)
	FindByTenant(ctx context.Context, tenantID string) ([]*entity.Eanregistry, error)
	Delete(ctx context.Context, id string, tenantID string) error
	SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.Eanregistry, error)
	CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error)
}
