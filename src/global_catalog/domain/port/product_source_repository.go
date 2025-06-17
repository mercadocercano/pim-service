package port

import (
	"context"
	"pim/src/global_catalog/domain/entity"
	"pim/src/shared/domain/criteria"
)

// ProductsourceRepository define los métodos para persistir Productsource
type ProductsourceRepository interface {
	Create(ctx context.Context, product_source *entity.Productsource) error
	Update(ctx context.Context, product_source *entity.Productsource) error
	FindByID(ctx context.Context, id string, tenantID string) (*entity.Productsource, error)
	FindByTenant(ctx context.Context, tenantID string) ([]*entity.Productsource, error)
	Delete(ctx context.Context, id string, tenantID string) error
	SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.Productsource, error)
	CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error)
}
