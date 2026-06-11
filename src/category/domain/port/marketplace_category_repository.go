package port

import (
	"context"

	cr "github.com/hornosg/go-shared/criteria"
	"saas-mt-pim-service/src/category/domain/entity"
)

// MarketplaceCategoryRepository define las operaciones de persistencia para categorías marketplace
type MarketplaceCategoryRepository interface {
	// Save guarda una categoría marketplace
	Save(ctx context.Context, category *entity.MarketplaceCategory) error

	// GetByID obtiene una categoría por su ID
	GetByID(ctx context.Context, id string) (*entity.MarketplaceCategory, error)

	// GetBySlug obtiene una categoría por su slug
	GetBySlug(ctx context.Context, slug string) (*entity.MarketplaceCategory, error)

	// GetByParentID obtiene las categorías hijas de un parent
	GetByParentID(ctx context.Context, parentID *string) ([]*entity.MarketplaceCategory, error)

	// GetRootCategories obtiene las categorías raíz (sin parent)
	GetRootCategories(ctx context.Context) ([]*entity.MarketplaceCategory, error)

	// GetTree obtiene el árbol completo de categorías
	GetTree(ctx context.Context) ([]*entity.MarketplaceCategory, error)

	// FindByCriteria busca categorías según criterios
	FindByCriteria(ctx context.Context, criteria cr.Criteria) ([]*entity.MarketplaceCategory, error)

	// CountByCriteria cuenta categorías según criterios
	CountByCriteria(ctx context.Context, criteria cr.Criteria) (int, error)

	// Update actualiza una categoría
	Update(ctx context.Context, category *entity.MarketplaceCategory) error

	// Delete elimina una categoría (soft delete)
	Delete(ctx context.Context, id string) error

	// ExistsBySlug verifica si ya existe una categoría con el slug
	ExistsBySlug(ctx context.Context, slug string) (bool, error)

	// GetCategoryPath obtiene el path completo de una categoría
	GetCategoryPath(ctx context.Context, categoryID string) ([]*entity.MarketplaceCategory, error)
}
