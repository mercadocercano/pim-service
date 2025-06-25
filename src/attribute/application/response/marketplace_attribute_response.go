package response

import (
	"pim/src/attribute/domain/entity"
	"time"
)

// MarketplaceAttributeResponse representa la respuesta de un atributo marketplace
type MarketplaceAttributeResponse struct {
	ID            string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name          string    `json:"name" example:"color"`
	Type          string    `json:"type" example:"enum"`
	Description   *string   `json:"description" example:"Color del producto"`
	IsRequired    bool      `json:"is_required" example:"false"`
	IsSearchable  bool      `json:"is_searchable" example:"true"`
	IsFilterable  bool      `json:"is_filterable" example:"true"`
	AllowedValues []string  `json:"allowed_values" example:"[\"rojo\", \"azul\", \"verde\"]"`
	IsActive      bool      `json:"is_active" example:"true"`
	CreatedAt     time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt     time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// FromMarketplaceEntity convierte una entidad MarketplaceAttribute a MarketplaceAttributeResponse
func FromMarketplaceEntity(attribute *entity.MarketplaceAttribute) *MarketplaceAttributeResponse {
	return &MarketplaceAttributeResponse{
		ID:            attribute.ID,
		Name:          attribute.Name,
		Type:          attribute.Type,
		Description:   attribute.Description,
		IsRequired:    attribute.IsRequired,
		IsSearchable:  attribute.IsSearchable,
		IsFilterable:  attribute.IsFilterable,
		AllowedValues: attribute.AllowedValues,
		IsActive:      attribute.IsActive,
		CreatedAt:     attribute.CreatedAt,
		UpdatedAt:     attribute.UpdatedAt,
	}
}

// FromMarketplaceEntities convierte una lista de entidades a una lista de responses
func FromMarketplaceEntities(attributes []*entity.MarketplaceAttribute) []*MarketplaceAttributeResponse {
	responses := make([]*MarketplaceAttributeResponse, len(attributes))
	for i, attribute := range attributes {
		responses[i] = FromMarketplaceEntity(attribute)
	}
	return responses
}
