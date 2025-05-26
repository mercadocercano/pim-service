package usecase

import (
	"context"
	"fmt"

	"pim/src/product/application/mapper"
	"pim/src/product/application/request"
	"pim/src/product/application/response"
	"pim/src/product/domain/port"
	"pim/src/product/domain/value_object"

	"github.com/google/uuid"
)

// UpdateProductVariantUseCase maneja la actualización de variantes a través del agregado Product
type UpdateProductVariantUseCase struct {
	productRepo port.ProductCriteriaRepository
	mapper      *mapper.ProductVariantMapper
}

// NewUpdateProductVariantUseCase crea una nueva instancia del use case
func NewUpdateProductVariantUseCase(
	productRepo port.ProductCriteriaRepository,
	mapper *mapper.ProductVariantMapper,
) *UpdateProductVariantUseCase {
	return &UpdateProductVariantUseCase{
		productRepo: productRepo,
		mapper:      mapper,
	}
}

// Execute ejecuta la actualización de una variante a través del agregado Product
func (uc *UpdateProductVariantUseCase) Execute(
	ctx context.Context,
	productID uuid.UUID,
	variantID uuid.UUID,
	req *request.UpdateProductVariantRequest,
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

	// Verificar que la variante existe en el agregado
	variant := product.GetVariantByID(variantID)
	if variant == nil {
		return nil, fmt.Errorf("variante no encontrada")
	}

	// Preparar los campos para actualizar
	var name *string
	if req.HasName() {
		nameValue := req.GetName()
		name = &nameValue
	}

	var sku *value_object.ProductSKU
	if req.HasSKU() {
		if req.GetSKU() != "" {
			sku, err = value_object.NewProductSKU(req.GetSKU())
			if err != nil {
				return nil, fmt.Errorf("error creando SKU: %w", err)
			}
		}
	}

	var sortOrder *int
	if req.HasSortOrder() {
		sortOrderValue := req.GetSortOrder()
		sortOrder = &sortOrderValue
	}

	var attributes *value_object.VariantAttributeCollection
	if req.HasAttributes() {
		attributes, err = uc.mapper.UpdateRequestAttributesToDomain(req.Attributes)
		if err != nil {
			return nil, fmt.Errorf("error procesando atributos: %w", err)
		}
	}

	// Actualizar la variante a través del agregado
	err = product.UpdateVariant(variantID, name, sku, sortOrder, attributes)
	if err != nil {
		return nil, fmt.Errorf("error actualizando variante: %w", err)
	}

	// Guardar los cambios a través del repositorio
	err = uc.productRepo.UpdateVariant(ctx, variant)
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
