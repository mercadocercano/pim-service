package usecase

import (
	"context"

	"pim/src/quickstart/domain/service"
)

// GetVariantsByBusinessTypeUseCase implementa el caso de uso para obtener variantes por tipo de negocio
type GetVariantsByBusinessTypeUseCase struct {
	quickstartService *service.QuickstartService
}

// NewGetVariantsByBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewGetVariantsByBusinessTypeUseCase(quickstartService *service.QuickstartService) *GetVariantsByBusinessTypeUseCase {
	return &GetVariantsByBusinessTypeUseCase{
		quickstartService: quickstartService,
	}
}

// Execute ejecuta el caso de uso para obtener variantes por tipo de negocio
func (uc *GetVariantsByBusinessTypeUseCase) Execute(ctx context.Context, businessType string) (interface{}, error) {
	return uc.quickstartService.GetVariantsByBusinessType(ctx, businessType)
}
