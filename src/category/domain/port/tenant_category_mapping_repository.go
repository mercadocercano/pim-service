package port

import (
	"context"

	"saas-mt-pim-service/src/category/domain/entity"
	"saas-mt-pim-service/src/shared/domain/criteria"
)

// TenantCategoryMappingRepository define las operaciones de persistencia para mapeos de categorías tenant
type TenantCategoryMappingRepository interface {
	// Save guarda un mapeo de categoría tenant
	Save(ctx context.Context, mapping *entity.TenantCategoryMapping) error

	// GetByID obtiene un mapeo por su ID
	GetByID(ctx context.Context, id string) (*entity.TenantCategoryMapping, error)

	// GetByTenantAndMarketplaceCategory obtiene un mapeo específico
	GetByTenantAndMarketplaceCategory(ctx context.Context, tenantID, marketplaceCategoryID string) (*entity.TenantCategoryMapping, error)

	// GetByTenantID obtiene todos los mapeos de un tenant
	GetByTenantID(ctx context.Context, tenantID string) ([]*entity.TenantCategoryMapping, error)

	// GetByMarketplaceCategoryID obtiene todos los mapeos de una categoría marketplace
	GetByMarketplaceCategoryID(ctx context.Context, marketplaceCategoryID string) ([]*entity.TenantCategoryMapping, error)

	// FindByCriteria busca mapeos según criterios
	FindByCriteria(ctx context.Context, criteria criteria.Criteria) ([]*entity.TenantCategoryMapping, error)

	// CountByCriteria cuenta mapeos según criterios
	CountByCriteria(ctx context.Context, criteria criteria.Criteria) (int, error)

	// Update actualiza un mapeo
	Update(ctx context.Context, mapping *entity.TenantCategoryMapping) error

	// Delete elimina un mapeo (soft delete)
	Delete(ctx context.Context, id string) error

	// GetTenantTaxonomy obtiene la taxonomía completa de un tenant (categorías + mapeos)
	GetTenantTaxonomy(ctx context.Context, tenantID string) ([]*entity.TenantCategoryMapping, error)
}
