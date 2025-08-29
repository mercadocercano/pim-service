package usecase

import (
	"context"
	"errors"

	"saas-mt-pim-service/src/category_attribute/domain/entity"
	"saas-mt-pim-service/src/category_attribute/domain/port"
)

// CreateCategoryAttributeUseCase implementa el caso de uso para crear un nuevo atributo de categoría
type CreateCategoryAttributeUseCase struct {
	categoryAttrRepo port.CategoryAttributeRepository
}

// NewCreateCategoryAttributeUseCase crea una nueva instancia del caso de uso
func NewCreateCategoryAttributeUseCase(repo port.CategoryAttributeRepository) *CreateCategoryAttributeUseCase {
	return &CreateCategoryAttributeUseCase{
		categoryAttrRepo: repo,
	}
}

// Execute ejecuta el caso de uso para crear un atributo de categoría
func (uc *CreateCategoryAttributeUseCase) Execute(ctx context.Context, tenantID, categoryID, attributeID string, allowedValues []string) (*entity.CategoryAttribute, error) {
	// Verificar si ya existe la relación
	existing, err := uc.categoryAttrRepo.FindByAttributeAndCategory(ctx, attributeID, categoryID, tenantID)
	if err == nil && existing != nil {
		return nil, errors.New("la relación atributo-categoría ya existe para este tenant")
	}

	// Crear la entidad
	categoryAttribute, err := entity.NewCategoryAttribute(tenantID, categoryID, attributeID, allowedValues)
	if err != nil {
		return nil, err
	}

	// Persistir la entidad
	err = uc.categoryAttrRepo.Create(ctx, categoryAttribute)
	if err != nil {
		return nil, err
	}

	return categoryAttribute, nil
}
