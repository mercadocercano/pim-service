package usecase

import (
	"context"

	"saas-mt-pim-service/src/quickstart/domain/port"
)

// GetCategoriesByBusinessTypeUseCase obtiene categorías por tipo de negocio desde la DB
type GetCategoriesByBusinessTypeUseCase struct {
	repo port.GetCategoriesByBusinessTypeRepository
}

// NewGetCategoriesByBusinessTypeUseCase crea una nueva instancia con el repositorio inyectado
func NewGetCategoriesByBusinessTypeUseCase(repo port.GetCategoriesByBusinessTypeRepository) *GetCategoriesByBusinessTypeUseCase {
	return &GetCategoriesByBusinessTypeUseCase{repo: repo}
}

// Execute obtiene las categorías asociadas al tipo de negocio dado su slug
func (uc *GetCategoriesByBusinessTypeUseCase) Execute(ctx context.Context, businessTypeSlug string) ([]port.CategoryByBusinessType, error) {
	return uc.repo.GetCategoriesByBusinessType(ctx, businessTypeSlug)
}
