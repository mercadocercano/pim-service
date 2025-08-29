package usecase

import (
	"context"

	"saas-mt-pim-service/src/product/tenant/application/request"
	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/domain/port"
	"saas-mt-pim-service/src/shared/domain/criteria"
)

// ValidateSKUsUseCase valida qué SKUs ya existen en el sistema
type ValidateSKUsUseCase struct {
	productRepo port.ProductCriteriaRepository
}

// NewValidateSKUsUseCase crea una nueva instancia del caso de uso
func NewValidateSKUsUseCase(productRepo port.ProductCriteriaRepository) *ValidateSKUsUseCase {
	return &ValidateSKUsUseCase{
		productRepo: productRepo,
	}
}

// Execute valida los SKUs proporcionados
func (uc *ValidateSKUsUseCase) Execute(ctx context.Context, req *request.ValidateSKUsRequest, tenantID string) (*response.ValidateSKUsResponse, error) {
	existingSKUs := []string{}
	availableSKUs := []string{}

	// Verificar cada SKU
	for _, sku := range req.SKUs {
		// Crear criterio para buscar por SKU
		filters := criteria.NewFilters(
			criteria.NewFilter("sku", criteria.OpEqual, sku),
			criteria.NewFilter("tenant_id", criteria.OpEqual, tenantID),
		)
		order := criteria.NewOrder("created_at", "DESC")
		pagination := criteria.NewPagination(1, 1) // Solo necesitamos saber si existe
		
		searchCriteria := criteria.NewCriteria(filters, order, pagination)

		// Buscar productos con ese SKU
		products, err := uc.productRepo.SearchByCriteria(ctx, searchCriteria)
		if err != nil {
			// En caso de error, asumir que el SKU está disponible
			availableSKUs = append(availableSKUs, sku)
			continue
		}

		if len(products) > 0 {
			existingSKUs = append(existingSKUs, sku)
		} else {
			availableSKUs = append(availableSKUs, sku)
		}
	}

	return &response.ValidateSKUsResponse{
		ExistingSKUs:  existingSKUs,
		AvailableSKUs: availableSKUs,
	}, nil
}