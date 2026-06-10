package port

import (
	"context"
	cr "github.com/mercadocercano/criteria"
	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
)

// ProductsourceRepository define los métodos para persistir Productsource
type ProductsourceRepository interface {
	Create(ctx context.Context, product_source *entity.Productsource) error
	Update(ctx context.Context, product_source *entity.Productsource) error
	FindByID(ctx context.Context, id string, tenantID string) (*entity.Productsource, error)
	FindByTenant(ctx context.Context, tenantID string) ([]*entity.Productsource, error)
	Delete(ctx context.Context, id string, tenantID string) error
	SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.Productsource, error)
	CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error)
}
