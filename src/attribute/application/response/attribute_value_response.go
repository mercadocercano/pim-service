package response

import (
	"time"

	"saas-mt-pim-service/src/attribute/domain/entity"
)

// AttributeValueResponse representa la respuesta de un valor de atributo
type AttributeValueResponse struct {
	ID          string    `json:"id"`
	AttributeID string    `json:"attribute_id"`
	Value       string    `json:"value"`
	Slug        string    `json:"slug"`
	SortOrder   int       `json:"sort_order"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

// FromAttributeValueEntity convierte una entidad a response
func FromAttributeValueEntity(v *entity.AttributeValue) *AttributeValueResponse {
	return &AttributeValueResponse{
		ID:          v.ID,
		AttributeID: v.AttributeID,
		Value:       v.Value,
		Slug:        v.Slug,
		SortOrder:   v.SortOrder,
		IsActive:    v.IsActive,
		CreatedAt:   v.CreatedAt,
	}
}

// FromAttributeValueEntities convierte una lista de entidades a responses
func FromAttributeValueEntities(values []*entity.AttributeValue) []*AttributeValueResponse {
	out := make([]*AttributeValueResponse, len(values))
	for i, v := range values {
		out[i] = FromAttributeValueEntity(v)
	}
	return out
}
