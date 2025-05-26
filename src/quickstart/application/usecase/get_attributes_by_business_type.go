package usecase

import (
	"context"

	"pim/src/quickstart/domain/service"
)

// GetAttributesByBusinessTypeUseCase implementa el caso de uso para obtener atributos por tipo de negocio
type GetAttributesByBusinessTypeUseCase struct {
	quickstartService *service.QuickstartService
}

// NewGetAttributesByBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewGetAttributesByBusinessTypeUseCase(quickstartService *service.QuickstartService) *GetAttributesByBusinessTypeUseCase {
	return &GetAttributesByBusinessTypeUseCase{
		quickstartService: quickstartService,
	}
}

// Execute ejecuta el caso de uso para obtener atributos por tipo de negocio
func (uc *GetAttributesByBusinessTypeUseCase) Execute(ctx context.Context, businessType string) (interface{}, error) {
	return uc.quickstartService.GetAttributesByBusinessType(ctx, businessType)
}
