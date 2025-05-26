package usecase

import (
	"context"

	"pim/src/quickstart/domain/entity"
	"pim/src/quickstart/domain/service"
)

// GetBusinessTypesUseCase implementa el caso de uso para obtener tipos de negocio
type GetBusinessTypesUseCase struct {
	quickstartService *service.QuickstartService
}

// NewGetBusinessTypesUseCase crea una nueva instancia del caso de uso
func NewGetBusinessTypesUseCase(quickstartService *service.QuickstartService) *GetBusinessTypesUseCase {
	return &GetBusinessTypesUseCase{
		quickstartService: quickstartService,
	}
}

// Execute ejecuta el caso de uso para obtener todos los tipos de negocio disponibles
func (uc *GetBusinessTypesUseCase) Execute(ctx context.Context) ([]*entity.BusinessType, error) {
	return uc.quickstartService.GetBusinessTypes(ctx)
}
