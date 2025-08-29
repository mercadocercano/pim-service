package response

import (
	"saas-mt-pim-service/src/attribute/domain/entity"
	"time"
)

// MarketplaceAttributeResponse representa la respuesta de un atributo marketplace
type MarketplaceAttributeResponse struct {
	ID                   string                 `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name                 string                 `json:"name" example:"Color"`
	Slug                 string                 `json:"slug" example:"color"`
	Type                 string                 `json:"type" example:"select"`
	IsFilterable         bool                   `json:"is_filterable" example:"true"`
	IsSearchable         bool                   `json:"is_searchable" example:"true"`
	IsRequiredForListing bool                   `json:"is_required_for_listing" example:"false"`
	ValidationRules      map[string]interface{} `json:"validation_rules" example:"{\"required\":true}"`
	SortOrder            int                    `json:"sort_order" example:"10"`
	CreatedAt            time.Time              `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt            time.Time              `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// FromMarketplaceEntity convierte una entidad MarketplaceAttribute a MarketplaceAttributeResponse
func FromMarketplaceEntity(attribute *entity.MarketplaceAttribute) *MarketplaceAttributeResponse {
	return &MarketplaceAttributeResponse{
		ID:                   attribute.ID,
		Name:                 attribute.Name,
		Slug:                 attribute.Slug,
		Type:                 attribute.Type,
		IsFilterable:         attribute.IsFilterable,
		IsSearchable:         attribute.IsSearchable,
		IsRequiredForListing: attribute.IsRequiredForListing,
		ValidationRules:      attribute.ValidationRules,
		SortOrder:            attribute.SortOrder,
		CreatedAt:            attribute.CreatedAt,
		UpdatedAt:            attribute.UpdatedAt,
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
