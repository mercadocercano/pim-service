package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"pim/src/category/domain/entity"
	"pim/src/category/domain/port"
)

// ErrInvalidUUID representa un error cuando el ID no es un UUID válido
var ErrInvalidUUID = fmt.Errorf("el ID proporcionado no es un UUID válido")

// GetCategoryByIDUseCase implementa el caso de uso para obtener una categoría por su ID
type GetCategoryByIDUseCase struct {
	categoryRepo port.CategoryRepository
}

// NewGetCategoryByIDUseCase crea una nueva instancia del caso de uso
func NewGetCategoryByIDUseCase(repo port.CategoryRepository) *GetCategoryByIDUseCase {
	return &GetCategoryByIDUseCase{
		categoryRepo: repo,
	}
}

// Execute ejecuta el caso de uso para obtener una categoría por su ID
func (uc *GetCategoryByIDUseCase) Execute(ctx context.Context, id string, tenantID string) (*entity.Category, error) {
	// Validar que el ID sea un UUID válido
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrInvalidUUID
	}

	// Buscar la categoría en el repositorio
	category, err := uc.categoryRepo.FindByID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}
	return category, nil
}
