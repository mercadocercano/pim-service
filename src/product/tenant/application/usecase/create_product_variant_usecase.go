package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/product/tenant/application/mapper"
	"saas-mt-pim-service/src/product/tenant/application/request"
	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/domain/port"
	"saas-mt-pim-service/src/product/tenant/domain/value_object"
)

// CreateProductVariantUseCase maneja la creación de variantes de productos a través del agregado Product
type CreateProductVariantUseCase struct {
	productRepo port.ProductCriteriaRepository
	mapper      *mapper.ProductVariantMapper
}

// NewCreateProductVariantUseCase crea una nueva instancia del use case
func NewCreateProductVariantUseCase(
	productRepo port.ProductCriteriaRepository,
	variantMapper *mapper.ProductVariantMapper,
) *CreateProductVariantUseCase {
	return &CreateProductVariantUseCase{
		productRepo: productRepo,
		mapper:      variantMapper,
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

		// VALIDACIÓN: Verificar que el SKU no exista en el tenant
		if sku != nil {
			existing, err := uc.productRepo.FindBySKUs(ctx, tenantID, []string{sku.Value()})
			if err != nil {
				return nil, fmt.Errorf("error validando SKU: %w", err)
			}
			if len(existing) > 0 {
				return nil, fmt.Errorf("ya existe una variante con el SKU '%s' en este tenant", sku.Value())
			}
		}
	} else {
		return nil, fmt.Errorf("el SKU es obligatorio para crear una variante")
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
