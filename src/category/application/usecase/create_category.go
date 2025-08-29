package usecase

import (
	"context"

	"saas-mt-pim-service/src/category/domain/entity"
	"saas-mt-pim-service/src/category/domain/port"
	"saas-mt-pim-service/src/category/domain/value_object"
)

// CreateCategoryUseCase implementa el caso de uso para crear una nueva categoría
type CreateCategoryUseCase struct {
	categoryRepo port.CategoryRepository
}

// NewCreateCategoryUseCase crea una nueva instancia del caso de uso
func NewCreateCategoryUseCase(repo port.CategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		categoryRepo: repo,
	}
}

// Execute ejecuta el caso de uso para crear una categoría
func (uc *CreateCategoryUseCase) Execute(ctx context.Context, tenantID, nameStr, description string, parentID *string) (*entity.Category, error) {
	// Validar el nombre usando el objeto de valor
	name, err := value_object.NewName(nameStr)
	if err != nil {
		return nil, err
	}

	// Crear la entidad con los valores validados
	category, err := entity.NewCategory(tenantID, name.Value(), description, parentID)
	if err != nil {
		return nil, err
	}

	// Persistir la entidad
	err = uc.categoryRepo.Create(ctx, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}
