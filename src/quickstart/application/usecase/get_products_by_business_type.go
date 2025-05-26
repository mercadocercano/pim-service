package usecase

import (
	"context"

	"pim/src/quickstart/domain/service"
)

// GetProductsByBusinessTypeUseCase implementa el caso de uso para obtener productos por tipo de negocio
type GetProductsByBusinessTypeUseCase struct {
	quickstartService *service.QuickstartService
}

// NewGetProductsByBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewGetProductsByBusinessTypeUseCase(quickstartService *service.QuickstartService) *GetProductsByBusinessTypeUseCase {
	return &GetProductsByBusinessTypeUseCase{
		quickstartService: quickstartService,
	}
}

// Execute ejecuta el caso de uso para obtener productos por tipo de negocio
func (uc *GetProductsByBusinessTypeUseCase) Execute(ctx context.Context, businessType string) (interface{}, error) {
	return uc.quickstartService.GetProductsByBusinessType(ctx, businessType)
}
