package mapper

import (
	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/domain/entity"
)

// ProductMapper maneja la conversión entre entidades de dominio y DTOs
type ProductMapper struct {
	variantMapper *ProductVariantMapper
}

// NewProductMapper crea una nueva instancia del mapper
func NewProductMapper() *ProductMapper {
	return &ProductMapper{
		variantMapper: NewProductVariantMapper(),
	}
}

// ToResponse convierte una entidad Product a ProductResponse (sin variantes)
func (m *ProductMapper) ToResponse(product *entity.Product) *response.ProductResponse {
	if product == nil {
		return nil
	}

	resp := &response.ProductResponse{
		ID:        product.IDString(),
		Name:      product.Name(),
		Status:    product.Status().Value(),
		CreatedAt: product.CreatedAt(),
		UpdatedAt: product.UpdatedAt(),
	}

	// Descripción opcional
	if product.Description() != nil {
		resp.Description = product.Description()
	}

	// SKU opcional
	if product.HasSKU() {
		sku := product.SKU().Value()
		resp.SKU = &sku
	}

	// Categoría opcional
	if product.HasCategory() {
		resp.Category = &response.CategoryReferenceResponse{
			ID:   product.CategoryReference().ID(),
			Name: product.CategoryReference().Name(),
		}
	}

	// Marca opcional
	if product.HasBrand() {
		resp.Brand = &response.BrandReferenceResponse{
			ID:   product.BrandReference().ID(),
			Name: product.BrandReference().Name(),
		}
	}

	// Price/stock de la variante default (si las variantes están cargadas)
	if defaultVariant := product.GetDefaultVariant(); defaultVariant != nil {
		price := defaultVariant.Price()
		stock := defaultVariant.Stock()
		resp.Price = &price
		resp.Stock = &stock
	}

	return resp
}

// ToResponseWithVariants convierte una entidad Product a ProductResponse incluyendo variantes
func (m *ProductMapper) ToResponseWithVariants(product *entity.Product) *response.ProductResponse {
	if product == nil {
		return nil
	}

	resp := m.ToResponse(product)

	// Agregar variantes activas
	if product.HasVariants() {
		variants := product.GetVariants()
		resp.Variants = m.variantMapper.ToResponseList(variants)
	}

	return resp
}

// ToProductWithVariantsResponse convierte una entidad Product a ProductWithVariantsResponse
func (m *ProductMapper) ToProductWithVariantsResponse(product *entity.Product) *response.ProductWithVariantsResponse {
	if product == nil {
		return nil
	}

	productResp := m.ToResponse(product)
	variants := product.GetVariants()
	defaultVariant := product.GetDefaultVariant()

	var defaultVariantResp *response.ProductVariantResponse
	if defaultVariant != nil {
		defaultVariantResp = m.variantMapper.ToResponse(defaultVariant)
	}

	return &response.ProductWithVariantsResponse{
		Product:        *productResp,
		DefaultVariant: defaultVariantResp,
		Variants:       m.variantMapper.ToResponseList(variants),
		TotalVariants:  len(variants),
	}
}

// ToResponseList convierte una lista de entidades Product a lista de ProductResponse
func (m *ProductMapper) ToResponseList(products []*entity.Product) []*response.ProductResponse {
	if products == nil {
		return []*response.ProductResponse{}
	}

	responses := make([]*response.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = m.ToResponse(product)
	}

	return responses
}

// ToResponseListWithVariants convierte una lista de entidades Product a lista de ProductResponse con variantes
func (m *ProductMapper) ToResponseListWithVariants(products []*entity.Product) []*response.ProductResponse {
	if products == nil {
		return []*response.ProductResponse{}
	}

	responses := make([]*response.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = m.ToResponseWithVariants(product)
	}

	return responses
}

// ToListResponse convierte una lista de productos con metadatos de paginación
func (m *ProductMapper) ToListResponse(
	products []*entity.Product,
	page, pageSize, totalItems int,
) *response.ProductListResponse {
	totalPages := 0
	if pageSize > 0 {
		totalPages = (totalItems + pageSize - 1) / pageSize
	}

	return &response.ProductListResponse{
		Products: m.ToResponseList(products),
		Pagination: response.PaginationResponse{
			Page:       page,
			PageSize:   pageSize,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
	}
}
