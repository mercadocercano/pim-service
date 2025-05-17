package usecase

import (
	"context"
	"errors"

	"pim/src/category/domain/entity"
	"pim/src/category/domain/port"
)

// ErrInvalidMove error cuando el movimiento de la categoría no es válido
var ErrInvalidMove = errors.New("movimiento de categoría inválido")

// MoveCategoryUseCase implementa el caso de uso para mover una categoría a otro padre
type MoveCategoryUseCase struct {
	categoryRepo port.CategoryRepository
}

// NewMoveCategoryUseCase crea una nueva instancia del caso de uso
func NewMoveCategoryUseCase(repo port.CategoryRepository) *MoveCategoryUseCase {
	return &MoveCategoryUseCase{
		categoryRepo: repo,
	}
}

// Execute ejecuta el caso de uso para mover una categoría
func (uc *MoveCategoryUseCase) Execute(ctx context.Context, id string, tenantID string, newParentID *string) (*entity.Category, error) {
	// Buscar la categoría a mover
	category, err := uc.categoryRepo.FindByID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	// Si se proporciona un nuevo padre, verificar que exista
	if newParentID != nil {
		parent, err := uc.categoryRepo.FindByID(ctx, *newParentID, tenantID)
		if err != nil {
			return nil, err
		}

		// Verificar que no se intente mover una categoría a sí misma o a sus descendientes
		if parent.ID == category.ID {
			return nil, ErrInvalidMove
		}

		// Verificar que el padre no sea un descendiente de la categoría
		currentParent := parent
		for currentParent.ParentID != nil {
			if *currentParent.ParentID == category.ID {
				return nil, ErrInvalidMove
			}
			currentParent, err = uc.categoryRepo.FindByID(ctx, *currentParent.ParentID, tenantID)
			if err != nil {
				return nil, err
			}
		}
	}

	// Actualizar el parentID
	err = category.Update(category.Name, category.Description, newParentID)
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
