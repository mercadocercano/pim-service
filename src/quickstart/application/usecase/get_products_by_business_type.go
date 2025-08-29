package usecase

import (
	"context"
)

// GetProductsByBusinessTypeUseCase obtiene productos por tipo de negocio
type GetProductsByBusinessTypeUseCase struct {
}

// NewGetProductsByBusinessTypeUseCase crea una nueva instancia
func NewGetProductsByBusinessTypeUseCase() *GetProductsByBusinessTypeUseCase {
	return &GetProductsByBusinessTypeUseCase{}
}

// Execute ejecuta el caso de uso
func (uc *GetProductsByBusinessTypeUseCase) Execute(ctx context.Context, businessType string) (interface{}, error) {
	// Mock data temporal
	products := []map[string]interface{}{
		{"id": 1, "name": "Coca-Cola 2.5L", "category": "beverages"},
		{"id": 2, "name": "Leche Entera 1L", "category": "dairy"},
		{"id": 3, "name": "Pan Blanco", "category": "bakery"},
	}
	
	return map[string]interface{}{
		"products": products,
	}, nil
}