package port

import (
	"context"

	"pim/src/brand/domain/entity"
	"pim/src/shared/domain/criteria"
)

// BrandRepository define el contrato para acceder a los datos de marcas
type BrandRepository interface {
	// Create guarda una nueva marca
	Create(ctx context.Context, brand *entity.Brand) error

	// FindByID busca una marca por su ID y tenantID
	FindByID(ctx context.Context, id string, tenantID string) (*entity.Brand, error)

	// FindByName busca una marca por su nombre y tenantID
	FindByName(ctx context.Context, name string, tenantID string) (*entity.Brand, error)

	// FindAll recupera todas las marcas de un tenant
	FindAll(ctx context.Context, tenantID string) ([]*entity.Brand, error)

	// Update actualiza una marca existente
	Update(ctx context.Context, brand *entity.Brand) error

	// Delete elimina una marca por su ID y tenantID
	Delete(ctx context.Context, id string, tenantID string) error

	// ExistsByName verifica si existe una marca con el nombre dado
	ExistsByName(ctx context.Context, name string, tenantID string, excludeID *string) (bool, error)
}

// BrandCriteriaRepository extiende BrandRepository con soporte para criteria
type BrandCriteriaRepository interface {
	BrandRepository
	criteria.CriteriaRepository[entity.Brand]
	criteria.ListRepository[entity.Brand]
}
