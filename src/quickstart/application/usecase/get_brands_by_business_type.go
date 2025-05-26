package usecase

import (
	"context"

	"pim/src/quickstart/domain/service"
)

// GetBrandsByBusinessTypeUseCase implementa el caso de uso para obtener marcas por tipo de negocio
type GetBrandsByBusinessTypeUseCase struct {
	quickstartService *service.QuickstartService
}

// NewGetBrandsByBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewGetBrandsByBusinessTypeUseCase(quickstartService *service.QuickstartService) *GetBrandsByBusinessTypeUseCase {
	return &GetBrandsByBusinessTypeUseCase{
		quickstartService: quickstartService,
	}
}

// Execute ejecuta el caso de uso para obtener marcas por tipo de negocio
func (uc *GetBrandsByBusinessTypeUseCase) Execute(ctx context.Context, businessType string) (interface{}, error) {
	return uc.quickstartService.GetBrandsByBusinessType(ctx, businessType)
}
