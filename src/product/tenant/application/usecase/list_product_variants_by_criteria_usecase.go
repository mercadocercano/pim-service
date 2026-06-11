package usecase

import (
	"context"
	"fmt"
	"net/url"

	cr "github.com/hornosg/go-shared/criteria"
	"saas-mt-pim-service/src/product/tenant/application/mapper"
	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/domain/entity"
	"saas-mt-pim-service/src/product/tenant/domain/port"

	"github.com/google/uuid"
)

// ListProductVariantsByCriteriaUseCase maneja la búsqueda de variantes a través del agregado Product
type ListProductVariantsByCriteriaUseCase struct {
	productRepo port.ProductCriteriaRepository
	mapper      *mapper.ProductVariantMapper
}

// NewListProductVariantsByCriteriaUseCase crea una nueva instancia del use case
func NewListProductVariantsByCriteriaUseCase(
	productRepo port.ProductCriteriaRepository,
	variantMapper *mapper.ProductVariantMapper,
) *ListProductVariantsByCriteriaUseCase {
	return &ListProductVariantsByCriteriaUseCase{
		productRepo: productRepo,
		mapper:      variantMapper,
	}
}

// Execute ejecuta la búsqueda de variantes con criterios a través del repositorio de productos
func (uc *ListProductVariantsByCriteriaUseCase) Execute(
	ctx context.Context,
	urlValues url.Values,
	tenantID string,
) (*response.ProductVariantListResponse, error) {
	// Crear builder de criterios desde URL values
	builder := cr.NewCriteriaBuilder().FromURLValues(urlValues)

	// Agregar filtro de tenant automáticamente
	builder.AddEqualFilter("tenant_id", tenantID)

	// Excluir variantes eliminadas por defecto (a menos que se especifique lo contrario)
	if urlValues.Get("include_deleted") != "true" {
		builder.AddNotEqualFilter("status", "deleted")
	}

	// Agregar filtros específicos de variantes desde URL
	if productID := urlValues.Get("product_id"); productID != "" {
		builder.AddUUIDFilter("product_id", productID)
	}

	if name := urlValues.Get("name"); name != "" {
		builder.AddLikeFilter("name", name)
	}

	if sku := urlValues.Get("sku"); sku != "" {
		builder.AddLikeFilter("sku", sku)
	}

	if status := urlValues.Get("status"); status != "" {
		builder.AddEqualFilter("status", status)
	}

	if isDefault := urlValues.Get("is_default"); isDefault != "" {
		builder.AddBoolFilter("is_default", isDefault)
	}

	// Construir criterios
	domainCriteria := builder.Build()

	// Obtener variantes a través del repositorio de productos
	variants, err := uc.productRepo.FindVariantsByCriteria(ctx, &domainCriteria)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo variantes: %w", err)
	}

	// Obtener total de elementos
	totalItems, err := uc.productRepo.CountVariantsByCriteria(ctx, &domainCriteria)
	if err != nil {
		return nil, fmt.Errorf("error contando variantes: %w", err)
	}

	// Obtener información de paginación
	pagination := domainCriteria.Pagination

	// Convertir a response con paginación
	return uc.mapper.ToListResponse(
		variants,
		pagination.Page,
		pagination.PageSize,
		totalItems,
	), nil
}

// GetVariantsByProductID obtiene todas las variantes de un producto específico
func (uc *ListProductVariantsByCriteriaUseCase) GetVariantsByProductID(
	ctx context.Context,
	productID uuid.UUID,
	tenantID string,
	includeDeleted bool,
) ([]response.ProductVariantResponse, error) {
	// Obtener el producto con sus variantes
	product, err := uc.productRepo.FindByIDWithVariants(ctx, productID, tenantID)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo producto: %w", err)
	}

	if product == nil {
		return nil, fmt.Errorf("producto no encontrado")
	}

	// Obtener variantes del agregado
	var variants []*entity.ProductVariant
	if includeDeleted {
		variants = product.Variants() // Todas las variantes
	} else {
		variants = product.GetVariants() // Solo las activas
	}

	// Convertir a response
	return uc.mapper.ToResponseList(variants), nil
}
