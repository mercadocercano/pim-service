package port

import (
	"context"

	"pim/src/attribute/domain/entity"
	"pim/src/shared/domain/criteria"
)

// TenantCustomAttributeRepository define las operaciones de persistencia para atributos custom de tenant
type TenantCustomAttributeRepository interface {
	// Save guarda un atributo custom de tenant
	Save(ctx context.Context, attribute *entity.TenantCustomAttribute) error

	// GetByID obtiene un atributo por su ID
	GetByID(ctx context.Context, id string) (*entity.TenantCustomAttribute, error)

	// GetByTenantAndSlug obtiene un atributo por tenant y slug
	GetByTenantAndSlug(ctx context.Context, tenantID, slug string) (*entity.TenantCustomAttribute, error)

	// GetByTenantID obtiene todos los atributos de un tenant
	GetByTenantID(ctx context.Context, tenantID string) ([]*entity.TenantCustomAttribute, error)

	// GetByTenantAndCategory obtiene atributos de un tenant para una categoría específica
	GetByTenantAndCategory(ctx context.Context, tenantID, marketplaceCategoryID string) ([]*entity.TenantCustomAttribute, error)

	// GetGlobalByTenant obtiene atributos globales de un tenant (sin categoría específica)
	GetGlobalByTenant(ctx context.Context, tenantID string) ([]*entity.TenantCustomAttribute, error)

	// FindByCriteria busca atributos según criterios
	FindByCriteria(ctx context.Context, criteria criteria.Criteria) ([]*entity.TenantCustomAttribute, error)

	// CountByCriteria cuenta atributos según criterios
	CountByCriteria(ctx context.Context, criteria criteria.Criteria) (int, error)

	// Update actualiza un atributo
	Update(ctx context.Context, attribute *entity.TenantCustomAttribute) error

	// Delete elimina un atributo (soft delete)
	Delete(ctx context.Context, id string) error

	// ExistsByTenantAndSlug verifica si ya existe un atributo con el slug en el tenant
	ExistsByTenantAndSlug(ctx context.Context, tenantID, slug string, marketplaceCategoryID *string) (bool, error)

	// GetTenantAttributesForCategory obtiene todos los atributos aplicables a una categoría
	// (incluye atributos globales del tenant + atributos específicos de la categoría)
	GetTenantAttributesForCategory(ctx context.Context, tenantID, marketplaceCategoryID string) ([]*entity.TenantCustomAttribute, error)

	// GetFilterableByTenant obtiene todos los atributos filtrables de un tenant
	GetFilterableByTenant(ctx context.Context, tenantID string) ([]*entity.TenantCustomAttribute, error)

	// GetSearchableByTenant obtiene todos los atributos buscables de un tenant
	GetSearchableByTenant(ctx context.Context, tenantID string) ([]*entity.TenantCustomAttribute, error)
}
