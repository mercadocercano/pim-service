package port

import (
	"context"

	cr "github.com/mercadocercano/criteria"
	"saas-mt-pim-service/src/product/tenant/domain/entity"

	"github.com/google/uuid"
)

// ProductVariantRepository define las operaciones básicas para variantes de productos
type ProductVariantRepository interface {
	Create(ctx context.Context, variant *entity.ProductVariant) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.ProductVariant, error)
	Update(ctx context.Context, variant *entity.ProductVariant) error
	Delete(ctx context.Context, id uuid.UUID) error

	// Operaciones específicas de variantes
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]*entity.ProductVariant, error)
	GetDefaultByProductID(ctx context.Context, productID uuid.UUID) (*entity.ProductVariant, error)
	GetBySKU(ctx context.Context, sku string, tenantID uuid.UUID) (*entity.ProductVariant, error)

	// Validaciones
	ExistsByName(ctx context.Context, name string, productID uuid.UUID, tenantID uuid.UUID) (bool, error)
	ExistsBySKU(ctx context.Context, sku string, tenantID uuid.UUID) (bool, error)
	ExistsByNameExcludingID(ctx context.Context, name string, productID uuid.UUID, tenantID uuid.UUID, excludeID uuid.UUID) (bool, error)
	ExistsBySKUExcludingID(ctx context.Context, sku string, tenantID uuid.UUID, excludeID uuid.UUID) (bool, error)

	// Operaciones de estado
	CountActiveByProductID(ctx context.Context, productID uuid.UUID) (int, error)
	HasDefaultVariant(ctx context.Context, productID uuid.UUID) (bool, error)
}

// ProductVariantCriteriaRepository extiende el repositorio básico con capacidades de búsqueda
type ProductVariantCriteriaRepository interface {
	ProductVariantRepository
	FindByCriteria(ctx context.Context, criteria *cr.Criteria) ([]*entity.ProductVariant, error)
	CountByCriteria(ctx context.Context, criteria *cr.Criteria) (int, error)
}
