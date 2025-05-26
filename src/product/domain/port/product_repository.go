package port

import (
	"context"

	"pim/src/product/domain/entity"
	"pim/src/shared/domain/criteria"

	"github.com/google/uuid"
)

// ProductRepository define las operaciones de persistencia para productos
type ProductRepository interface {
	// Operaciones básicas CRUD
	Save(ctx context.Context, product *entity.Product) error
	FindByID(ctx context.Context, id uuid.UUID, tenantID string) (*entity.Product, error)
	FindByIDWithVariants(ctx context.Context, id uuid.UUID, tenantID string) (*entity.Product, error)
	FindBySKU(ctx context.Context, sku, tenantID string) (*entity.Product, error)
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, id uuid.UUID, tenantID string) error

	// Operaciones de variantes (a través del agregado)
	SaveVariant(ctx context.Context, productID uuid.UUID, variant *entity.ProductVariant) error
	UpdateVariant(ctx context.Context, variant *entity.ProductVariant) error
	DeleteVariant(ctx context.Context, variantID uuid.UUID) error
	LoadVariantsForProduct(ctx context.Context, productID uuid.UUID) ([]*entity.ProductVariant, error)

	// Verificaciones de existencia
	ExistsByID(ctx context.Context, id uuid.UUID, tenantID string) (bool, error)
	ExistsBySKU(ctx context.Context, sku, tenantID string) (bool, error)
	ExistsByName(ctx context.Context, name, tenantID string) (bool, error)
	ExistsByNameExcludingID(ctx context.Context, name, tenantID string, excludeID uuid.UUID) (bool, error)
	ExistsBySKUExcludingID(ctx context.Context, sku, tenantID string, excludeID uuid.UUID) (bool, error)

	// Verificaciones de variantes
	VariantExistsByName(ctx context.Context, name string, productID uuid.UUID, tenantID string) (bool, error)
	VariantExistsBySKU(ctx context.Context, sku, tenantID string) (bool, error)
	VariantExistsByNameExcludingID(ctx context.Context, name string, productID uuid.UUID, tenantID string, excludeID uuid.UUID) (bool, error)
	VariantExistsBySKUExcludingID(ctx context.Context, sku, tenantID string, excludeID uuid.UUID) (bool, error)
}

// ProductCriteriaRepository extiende ProductRepository con capacidades de búsqueda por criterios
type ProductCriteriaRepository interface {
	ProductRepository
	criteria.CriteriaRepository[entity.Product]

	// Búsqueda de variantes por criterios (a través del producto)
	FindVariantsByCriteria(ctx context.Context, criteria *criteria.Criteria) ([]*entity.ProductVariant, error)
	CountVariantsByCriteria(ctx context.Context, criteria *criteria.Criteria) (int, error)
}
