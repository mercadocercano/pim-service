package port

import (
	"context"

	"saas-mt-pim-service/src/category/domain/entity"
	cr "github.com/mercadocercano/criteria"
)

// CategoryRepository define el contrato para acceder a los datos de categorías
type CategoryRepository interface {
	// Create guarda una nueva categoría
	Create(ctx context.Context, category *entity.Category) error

	// FindByID busca una categoría por su ID y tenantID
	FindByID(ctx context.Context, id string, tenantID string) (*entity.Category, error)

	// FindAll recupera todas las categorías de un tenant
	FindAll(ctx context.Context, tenantID string) ([]*entity.Category, error)

	// Update actualiza una categoría existente
	Update(ctx context.Context, category *entity.Category) error

	// Delete elimina una categoría por su ID y tenantID
	Delete(ctx context.Context, id string, tenantID string) error
}

// CategoryCriteriaRepository extiende CategoryRepository con soporte para criteria
type CategoryCriteriaRepository interface {
	CategoryRepository
	cr.CriteriaRepository[entity.Category]
	cr.ListRepository[entity.Category]
}
