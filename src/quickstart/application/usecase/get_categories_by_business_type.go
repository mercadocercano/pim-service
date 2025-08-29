package usecase

import (
	"context"
)

// GetCategoriesByBusinessTypeUseCase obtiene categorías por tipo de negocio
type GetCategoriesByBusinessTypeUseCase struct {
	// Por ahora retorna datos mock, después se conectará con business_type_templates
}

// NewGetCategoriesByBusinessTypeUseCase crea una nueva instancia
func NewGetCategoriesByBusinessTypeUseCase() *GetCategoriesByBusinessTypeUseCase {
	return &GetCategoriesByBusinessTypeUseCase{}
}

// Execute ejecuta el caso de uso
func (uc *GetCategoriesByBusinessTypeUseCase) Execute(ctx context.Context, businessType string) (interface{}, error) {
	// Mock data temporal - después vendrá de business_type_templates
	categories := []map[string]interface{}{
		{"id": 1, "name": "Bebidas", "code": "beverages"},
		{"id": 2, "name": "Lácteos", "code": "dairy"},
		{"id": 3, "name": "Panadería", "code": "bakery"},
	}
	
	return map[string]interface{}{
		"categories": categories,
	}, nil
}