package usecase

import (
	"context"

	cr "github.com/mercadocercano/criteria"
	"saas-mt-pim-service/src/product/tenant/application/request"
	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/domain/port"
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
		filters := cr.NewFilters(
			cr.NewFilter("sku", cr.OpEqual, sku),
			cr.NewFilter("tenant_id", cr.OpEqual, tenantID),
		)
		order := cr.NewOrder("created_at", cr.OrderDesc)
		pagination := cr.NewPagination(1, 1) // Solo necesitamos saber si existe

		searchCriteria := cr.NewCriteria(filters, []cr.Order{order}, pagination)

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
