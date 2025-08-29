package usecase

import (
	"context"

	"saas-mt-pim-service/src/category/domain/entity"
	"saas-mt-pim-service/src/category/domain/port"
)

// GetCategoriesUseCase implementa el caso de uso para obtener todas las categorías
type GetCategoriesUseCase struct {
	categoryRepo port.CategoryRepository
}

// NewGetCategoriesUseCase crea una nueva instancia del caso de uso
func NewGetCategoriesUseCase(repo port.CategoryRepository) *GetCategoriesUseCase {
	return &GetCategoriesUseCase{
		categoryRepo: repo,
	}
}

// Execute ejecuta el caso de uso para obtener todas las categorías
func (uc *GetCategoriesUseCase) Execute(ctx context.Context, tenantID string) ([]*entity.Category, error) {
	return uc.categoryRepo.FindAll(ctx, tenantID)
}
