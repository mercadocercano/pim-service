package usecase

import (
	"context"
	"errors"

	"saas-mt-pim-service/src/category_attribute/domain/port"
)

// UpdateCategoryAttributeUseCase implementa el caso de uso para actualizar un atributo de categoría
type UpdateCategoryAttributeUseCase struct {
	categoryAttrRepo port.CategoryAttributeRepository
}

// NewUpdateCategoryAttributeUseCase crea una nueva instancia del caso de uso
func NewUpdateCategoryAttributeUseCase(repo port.CategoryAttributeRepository) *UpdateCategoryAttributeUseCase {
	return &UpdateCategoryAttributeUseCase{
		categoryAttrRepo: repo,
	}
}

// Execute ejecuta el caso de uso para actualizar los valores permitidos de un atributo de categoría
func (uc *UpdateCategoryAttributeUseCase) Execute(ctx context.Context, id, tenantID string, allowedValues []string) error {
	// Buscar la entidad existente
	categoryAttribute, err := uc.categoryAttrRepo.FindByID(ctx, id, tenantID)
	if err != nil {
		return err
	}

	if categoryAttribute == nil {
		return errors.New("atributo de categoría no encontrado")
	}

	// Actualizar los valores permitidos
	categoryAttribute.UpdateAllowedValues(allowedValues)

	// Persistir los cambios
	return uc.categoryAttrRepo.Update(ctx, categoryAttribute)
}
