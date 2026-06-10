package usecase

import (
	"context"
)

// GetVariantsByBusinessTypeUseCase obtiene variantes por tipo de negocio
type GetVariantsByBusinessTypeUseCase struct {
}

// NewGetVariantsByBusinessTypeUseCase crea una nueva instancia
func NewGetVariantsByBusinessTypeUseCase() *GetVariantsByBusinessTypeUseCase {
	return &GetVariantsByBusinessTypeUseCase{}
}

// Execute ejecuta el caso de uso
func (uc *GetVariantsByBusinessTypeUseCase) Execute(ctx context.Context, businessType string) (interface{}, error) {
	// Mock data temporal
	return map[string]interface{}{
		"variants": []interface{}{},
	}, nil
}
