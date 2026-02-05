package mapper

import (
	"math"

	"saas-mt-pim-service/src/product/tenant/application/request"
	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/domain/entity"
	"saas-mt-pim-service/src/product/tenant/domain/value_object"
)

// ProductVariantMapper maneja la conversión entre entidades de dominio y DTOs
type ProductVariantMapper struct{}

// NewProductVariantMapper crea una nueva instancia del mapper
func NewProductVariantMapper() *ProductVariantMapper {
	return &ProductVariantMapper{}
}

// ToResponse convierte una entidad ProductVariant a ProductVariantResponse
func (m *ProductVariantMapper) ToResponse(variant *entity.ProductVariant) *response.ProductVariantResponse {
	if variant == nil {
		return nil
	}

	var sku *string
	if variant.HasSKU() {
		skuValue := variant.SKU().Value()
		sku = &skuValue
	}

	attributes := m.attributesToResponse(variant.Attributes())

	return &response.ProductVariantResponse{
		ID:         variant.ID(),
		ProductID:  variant.ProductID(),
		Name:       variant.Name(),
		SKU:        sku,
		Status:     variant.Status().Value(),
		IsDefault:  variant.IsDefault(),
		SortOrder:  variant.SortOrder(),
		Price:      variant.Price(),
		Stock:      variant.Stock(),
		Attributes: attributes,
		CreatedAt:  variant.CreatedAt(),
		UpdatedAt:  variant.UpdatedAt(),
	}
}

// ToResponseList convierte una lista de entidades a lista de responses
func (m *ProductVariantMapper) ToResponseList(variants []*entity.ProductVariant) []response.ProductVariantResponse {
	if variants == nil {
		return []response.ProductVariantResponse{}
	}

	responses := make([]response.ProductVariantResponse, len(variants))
	for i, variant := range variants {
		if response := m.ToResponse(variant); response != nil {
			responses[i] = *response
		}
	}

	return responses
}

// ToListResponse convierte una lista de entidades a ProductVariantListResponse con paginación
func (m *ProductVariantMapper) ToListResponse(
	variants []*entity.ProductVariant,
	page, pageSize, totalItems int,
) *response.ProductVariantListResponse {
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))

	return &response.ProductVariantListResponse{
		Variants: m.ToResponseList(variants),
		Pagination: response.PaginationResponse{
			Page:       page,
			PageSize:   pageSize,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
	}
}

// ToStatusChangeResponse convierte una entidad a VariantStatusChangeResponse
func (m *ProductVariantMapper) ToStatusChangeResponse(variant *entity.ProductVariant, message string) *response.VariantStatusChangeResponse {
	if variant == nil {
		return nil
	}

	return &response.VariantStatusChangeResponse{
		ID:        variant.ID(),
		Status:    variant.Status().Value(),
		UpdatedAt: variant.UpdatedAt(),
		Message:   message,
	}
}

// attributesToResponse convierte atributos de dominio a response
func (m *ProductVariantMapper) attributesToResponse(attributes *value_object.VariantAttributeCollection) []response.VariantAttributeResponse {
	if attributes == nil || attributes.IsEmpty() {
		return []response.VariantAttributeResponse{}
	}

	responses := make([]response.VariantAttributeResponse, 0, attributes.Count())
	for _, attr := range attributes.Attributes() {
		responses = append(responses, response.VariantAttributeResponse{
			Name:  attr.Name(),
			Value: attr.Value(),
		})
	}

	return responses
}

// RequestAttributesToDomain convierte atributos de request a value objects de dominio
func (m *ProductVariantMapper) RequestAttributesToDomain(reqAttributes []request.CreateVariantAttributeRequest) (*value_object.VariantAttributeCollection, error) {
	if len(reqAttributes) == 0 {
		return value_object.NewVariantAttributeCollection([]*value_object.VariantAttribute{})
	}

	attributes := make([]*value_object.VariantAttribute, 0, len(reqAttributes))
	for _, reqAttr := range reqAttributes {
		attr, err := value_object.NewVariantAttribute(reqAttr.Name, reqAttr.Value)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, attr)
	}

	return value_object.NewVariantAttributeCollection(attributes)
}

// UpdateRequestAttributesToDomain convierte atributos de update request a value objects de dominio
func (m *ProductVariantMapper) UpdateRequestAttributesToDomain(reqAttributes []request.UpdateVariantAttributeRequest) (*value_object.VariantAttributeCollection, error) {
	if len(reqAttributes) == 0 {
		return value_object.NewVariantAttributeCollection([]*value_object.VariantAttribute{})
	}

	attributes := make([]*value_object.VariantAttribute, 0, len(reqAttributes))
	for _, reqAttr := range reqAttributes {
		attr, err := value_object.NewVariantAttribute(reqAttr.Name, reqAttr.Value)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, attr)
	}

	return value_object.NewVariantAttributeCollection(attributes)
}
