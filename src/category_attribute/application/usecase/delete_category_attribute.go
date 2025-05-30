package usecase

import (
	"context"

	"pim/src/category_attribute/domain/port"
)

// DeleteCategoryAttributeUseCase implementa el caso de uso para eliminar un atributo de categoría
type DeleteCategoryAttributeUseCase struct {
	categoryAttrRepo port.CategoryAttributeRepository
}

// NewDeleteCategoryAttributeUseCase crea una nueva instancia del caso de uso
func NewDeleteCategoryAttributeUseCase(repo port.CategoryAttributeRepository) *DeleteCategoryAttributeUseCase {
	return &DeleteCategoryAttributeUseCase{
		categoryAttrRepo: repo,
	}
}

// Execute ejecuta el caso de uso para eliminar un atributo de categoría
func (uc *DeleteCategoryAttributeUseCase) Execute(ctx context.Context, id, tenantID string) error {
	return uc.categoryAttrRepo.Delete(ctx, id, tenantID)
}
