package usecase

import (
	"context"
)

// GetBrandsByBusinessTypeUseCase obtiene marcas por tipo de negocio
type GetBrandsByBusinessTypeUseCase struct {
}

// NewGetBrandsByBusinessTypeUseCase crea una nueva instancia
func NewGetBrandsByBusinessTypeUseCase() *GetBrandsByBusinessTypeUseCase {
	return &GetBrandsByBusinessTypeUseCase{}
}

// Execute ejecuta el caso de uso
func (uc *GetBrandsByBusinessTypeUseCase) Execute(ctx context.Context, businessType string) (interface{}, error) {
	// Mock data temporal
	brands := []string{"Coca-Cola", "La Serenísima", "Bimbo", "Skip", "Arcor"}
	
	return map[string]interface{}{
		"brands": brands,
	}, nil
}