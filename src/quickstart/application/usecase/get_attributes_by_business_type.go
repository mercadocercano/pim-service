package usecase

import (
	"context"
)

// GetAttributesByBusinessTypeUseCase obtiene atributos por tipo de negocio
type GetAttributesByBusinessTypeUseCase struct {
	// Por ahora retorna datos mock
}

// NewGetAttributesByBusinessTypeUseCase crea una nueva instancia
func NewGetAttributesByBusinessTypeUseCase() *GetAttributesByBusinessTypeUseCase {
	return &GetAttributesByBusinessTypeUseCase{}
}

// Execute ejecuta el caso de uso
func (uc *GetAttributesByBusinessTypeUseCase) Execute(ctx context.Context, businessType string) (interface{}, error) {
	// Mock data temporal
	attributes := []map[string]interface{}{
		{"id": 1, "name": "Precio", "type": "number"},
		{"id": 2, "name": "Stock", "type": "number"},
		{"id": 3, "name": "Código de barras", "type": "text"},
	}
	
	return map[string]interface{}{
		"attributes": attributes,
	}, nil
}