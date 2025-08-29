package usecase

import (
	"context"
	"errors"

	"saas-mt-pim-service/src/category/domain/entity"
	"saas-mt-pim-service/src/category/domain/port"
	"saas-mt-pim-service/src/category/domain/value_object"
)

// ErrCategoryNotFound error cuando la categoría no se encuentra
var ErrCategoryNotFound = errors.New("categoría no encontrada")

// UpdateCategoryUseCase implementa el caso de uso para actualizar una categoría
type UpdateCategoryUseCase struct {
	categoryRepo port.CategoryRepository
}

// NewUpdateCategoryUseCase crea una nueva instancia del caso de uso
func NewUpdateCategoryUseCase(repo port.CategoryRepository) *UpdateCategoryUseCase {
	return &UpdateCategoryUseCase{
		categoryRepo: repo,
	}
}

// Execute ejecuta el caso de uso para actualizar una categoría
func (uc *UpdateCategoryUseCase) Execute(ctx context.Context, id string, tenantID string, nameStr, description string, parentID *string) (*entity.Category, error) {
	// Validar el nombre usando el objeto de valor
	name, err := value_object.NewName(nameStr)
	if err != nil {
		return nil, err
	}

	// Buscar la categoría existente
	category, err := uc.categoryRepo.FindByID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	// Actualizar los campos
	err = category.Update(name.Value(), description, parentID)
	if err != nil {
		return nil, err
	}

	// Persistir los cambios
	err = uc.categoryRepo.Update(ctx, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}
