package usecase

import (
	"context"

	"pim/src/category_attribute/domain/entity"
	"pim/src/category_attribute/domain/port"
)

// GetCategoryAttributesUseCase implementa el caso de uso para obtener atributos de categoría
type GetCategoryAttributesUseCase struct {
	categoryAttrRepo port.CategoryAttributeRepository
}

// NewGetCategoryAttributesUseCase crea una nueva instancia del caso de uso
func NewGetCategoryAttributesUseCase(repo port.CategoryAttributeRepository) *GetCategoryAttributesUseCase {
	return &GetCategoryAttributesUseCase{
		categoryAttrRepo: repo,
	}
}

// Execute ejecuta el caso de uso para obtener atributos de categoría por tenant y category
func (uc *GetCategoryAttributesUseCase) Execute(ctx context.Context, tenantID, categoryID string) ([]*entity.CategoryAttribute, error) {
	if categoryID != "" {
		return uc.categoryAttrRepo.FindByCategoryAndTenant(ctx, categoryID, tenantID)
	}

	return uc.categoryAttrRepo.FindByTenant(ctx, tenantID)
}
