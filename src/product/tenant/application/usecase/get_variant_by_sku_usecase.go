package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"saas-mt-pim-service/src/product/tenant/application/mapper"
	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/domain/port"
)

// GetVariantBySKUUseCase caso de uso para obtener una variante por SKU
type GetVariantBySKUUseCase struct {
	productRepo port.ProductCriteriaRepository
	mapper      *mapper.ProductVariantMapper
}

// NewGetVariantBySKUUseCase crea una nueva instancia del caso de uso
func NewGetVariantBySKUUseCase(productRepo port.ProductCriteriaRepository, variantMapper *mapper.ProductVariantMapper) *GetVariantBySKUUseCase {
	return &GetVariantBySKUUseCase{
		productRepo: productRepo,
		mapper:      variantMapper,
	}
}

// Execute ejecuta el caso de uso
func (uc *GetVariantBySKUUseCase) Execute(ctx context.Context, sku string, tenantID string) (*response.ProductVariantResponse, error) {
	// Validar tenant ID
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		return nil, fmt.Errorf("invalid tenant_id: %w", err)
	}

	// Validar SKU
	if sku == "" {
		return nil, fmt.Errorf("sku is required")
	}

	// Buscar variante por SKU
	variant, err := uc.productRepo.GetBySKU(ctx, sku, tenantUUID)
	if err != nil {
		return nil, fmt.Errorf("variant not found: %s", sku)
	}

	// Convertir a response usando mapper
	return uc.mapper.ToResponse(variant), nil
}
