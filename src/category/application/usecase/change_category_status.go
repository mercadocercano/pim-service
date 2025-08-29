package usecase

import (
	"context"

	"saas-mt-pim-service/src/category/domain/entity"
	"saas-mt-pim-service/src/category/domain/port"
)

// ChangeCategoryStatusUseCase implementa el caso de uso para cambiar el estado de una categoría
type ChangeCategoryStatusUseCase struct {
	categoryRepo port.CategoryRepository
}

// NewChangeCategoryStatusUseCase crea una nueva instancia del caso de uso
func NewChangeCategoryStatusUseCase(repo port.CategoryRepository) *ChangeCategoryStatusUseCase {
	return &ChangeCategoryStatusUseCase{
		categoryRepo: repo,
	}
}

// Activate activa una categoría
func (uc *ChangeCategoryStatusUseCase) Activate(ctx context.Context, id string, tenantID string) (*entity.Category, error) {
	category, err := uc.categoryRepo.FindByID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	category.Activate()
	err = uc.categoryRepo.Update(ctx, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// Deactivate desactiva una categoría
func (uc *ChangeCategoryStatusUseCase) Deactivate(ctx context.Context, id string, tenantID string) (*entity.Category, error) {
	category, err := uc.categoryRepo.FindByID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	category.Deactivate()
	err = uc.categoryRepo.Update(ctx, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}
