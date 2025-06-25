package usecase

import (
	"context"
	"fmt"

	"pim/src/product/tenant/application/mapper"
	"pim/src/product/tenant/application/request"
	"pim/src/product/tenant/application/response"
	"pim/src/product/tenant/domain/port"
	"pim/src/product/tenant/domain/value_object"
)

// CreateProductVariantUseCase maneja la creación de variantes de productos a través del agregado Product
type CreateProductVariantUseCase struct {
	productRepo port.ProductCriteriaRepository
	mapper      *mapper.ProductVariantMapper
}

// NewCreateProductVariantUseCase crea una nueva instancia del use case
func NewCreateProductVariantUseCase(
	productRepo port.ProductCriteriaRepository,
	mapper *mapper.ProductVariantMapper,
) *CreateProductVariantUseCase {
	return &CreateProductVariantUseCase{
		productRepo: productRepo,
		mapper:      mapper,
	}
}

// Execute ejecuta la creación de una variante de producto a través del agregado
func (uc *CreateProductVariantUseCase) Execute(
	ctx context.Context,
	req *request.CreateProductVariantRequest,
	tenantID string,
) (*response.ProductVariantResponse, error) {
	// Obtener el producto (agregado raíz) con sus variantes
	product, err := uc.productRepo.FindByIDWithVariants(ctx, req.ProductID, tenantID)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo producto: %w", err)
	}
	if product == nil {
		return nil, fmt.Errorf("producto no encontrado")
	}

	// Convertir atributos de request a value objects de dominio
	attributes, err := uc.mapper.RequestAttributesToDomain(req.Attributes)
	if err != nil {
		return nil, fmt.Errorf("error procesando atributos: %w", err)
	}

	// Crear SKU si se proporcionó
	var sku *value_object.ProductSKU
	if req.HasSKU() {
		sku, err = value_object.NewProductSKU(req.GetSKU())
		if err != nil {
			return nil, fmt.Errorf("error creando SKU: %w", err)
		}
	}

	// Agregar la variante al producto (el agregado maneja la lógica de negocio)
	variant, err := product.AddVariant(
		req.Name,
		sku,
		req.GetIsDefault(),
		req.GetSortOrder(),
		attributes,
	)
	if err != nil {
		return nil, fmt.Errorf("error agregando variante al producto: %w", err)
	}

	// Guardar la variante a través del repositorio del producto
	err = uc.productRepo.SaveVariant(ctx, product.ID(), variant)
	if err != nil {
		return nil, fmt.Errorf("error guardando variante: %w", err)
	}

	// Actualizar el producto para reflejar los cambios
	err = uc.productRepo.Update(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("error actualizando producto: %w", err)
	}

	// Convertir a response y retornar
	return uc.mapper.ToResponse(variant), nil
}
