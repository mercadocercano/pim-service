package usecase

import (
	"context"

	"pim/src/quickstart/domain/service"
)

// GetCategoriesByBusinessTypeUseCase implementa el caso de uso para obtener categorías por tipo de negocio
type GetCategoriesByBusinessTypeUseCase struct {
	quickstartService *service.QuickstartService
}

// NewGetCategoriesByBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewGetCategoriesByBusinessTypeUseCase(quickstartService *service.QuickstartService) *GetCategoriesByBusinessTypeUseCase {
	return &GetCategoriesByBusinessTypeUseCase{
		quickstartService: quickstartService,
	}
}

// Execute ejecuta el caso de uso para obtener categorías por tipo de negocio
func (uc *GetCategoriesByBusinessTypeUseCase) Execute(ctx context.Context, businessType string) (interface{}, error) {
	return uc.quickstartService.GetCategoriesByBusinessType(ctx, businessType)
}
