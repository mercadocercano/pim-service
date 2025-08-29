package usecase

import (
	"context"
	"errors"

	"saas-mt-pim-service/src/category_attribute/domain/entity"
	"saas-mt-pim-service/src/category_attribute/domain/port"
)

// GetCategoryAttributesUseCase implementa el caso de uso para obtener atributos de categoría
type GetCategoryAttributesUseCase struct {
	categoryAttrRepo port.CategoryAttributeRepository
}

// GetDetailedCategoryAttributesUseCase implementa el caso de uso para obtener atributos detallados de categoría
type GetDetailedCategoryAttributesUseCase struct {
	categoryAttrRepo port.CategoryAttributeRepository
}

// NewGetCategoryAttributesUseCase crea una nueva instancia del caso de uso
func NewGetCategoryAttributesUseCase(repo port.CategoryAttributeRepository) *GetCategoryAttributesUseCase {
	return &GetCategoryAttributesUseCase{
		categoryAttrRepo: repo,
	}
}

// NewGetDetailedCategoryAttributesUseCase crea una nueva instancia del caso de uso detallado
func NewGetDetailedCategoryAttributesUseCase(repo port.CategoryAttributeRepository) *GetDetailedCategoryAttributesUseCase {
	return &GetDetailedCategoryAttributesUseCase{
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

// Execute ejecuta el caso de uso para obtener atributos detallados de categoría
func (uc *GetDetailedCategoryAttributesUseCase) Execute(ctx context.Context, tenantID, categoryID string) ([]*port.DetailedCategoryAttribute, error) {
	if categoryID == "" {
		return nil, errors.New("category_id es obligatorio para obtener atributos detallados")
	}

	return uc.categoryAttrRepo.FindDetailedByCategoryAndTenant(ctx, categoryID, tenantID)
}
