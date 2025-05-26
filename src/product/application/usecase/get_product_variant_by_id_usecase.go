package usecase

import (
	"context"
	"fmt"

	"pim/src/product/application/mapper"
	"pim/src/product/application/response"
	"pim/src/product/domain/port"

	"github.com/google/uuid"
)

// GetProductVariantByIDUseCase maneja la obtención de una variante por ID a través del agregado Product
type GetProductVariantByIDUseCase struct {
	productRepo port.ProductCriteriaRepository
	mapper      *mapper.ProductVariantMapper
}

// NewGetProductVariantByIDUseCase crea una nueva instancia del use case
func NewGetProductVariantByIDUseCase(
	productRepo port.ProductCriteriaRepository,
	mapper *mapper.ProductVariantMapper,
) *GetProductVariantByIDUseCase {
	return &GetProductVariantByIDUseCase{
		productRepo: productRepo,
		mapper:      mapper,
	}
}

// Execute ejecuta la obtención de una variante por ID a través del producto
func (uc *GetProductVariantByIDUseCase) Execute(
	ctx context.Context,
	productID uuid.UUID,
	variantID uuid.UUID,
	tenantID string,
) (*response.ProductVariantResponse, error) {
	// Obtener el producto con sus variantes
	product, err := uc.productRepo.FindByIDWithVariants(ctx, productID, tenantID)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo producto: %w", err)
	}

	if product == nil {
		return nil, fmt.Errorf("producto no encontrado")
	}

	// Obtener la variante específica del agregado
	variant := product.GetVariantByID(variantID)
	if variant == nil {
		return nil, fmt.Errorf("variante no encontrada")
	}

	// Convertir a response y retornar
	return uc.mapper.ToResponse(variant), nil
}
