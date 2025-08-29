package response

import (
	"saas-mt-pim-service/src/attribute/domain/entity"
	"time"
)

// AttributeResponse representa la respuesta de un atributo básico
type AttributeResponse struct {
	ID        string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	TenantID  string    `json:"tenant_id" example:"tenant-123"`
	Name      string    `json:"name" example:"color"`
	Active    bool      `json:"active" example:"true"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// FromEntity convierte una entidad Attribute a AttributeResponse
func FromEntity(attribute *entity.Attribute) *AttributeResponse {
	return &AttributeResponse{
		ID:        attribute.ID,
		TenantID:  attribute.TenantID,
		Name:      attribute.Name,
		Active:    attribute.Active,
		CreatedAt: attribute.CreatedAt,
		UpdatedAt: attribute.UpdatedAt,
	}
}

// FromEntities convierte una lista de entidades a una lista de responses
func FromEntities(attributes []*entity.Attribute) []*AttributeResponse {
	responses := make([]*AttributeResponse, len(attributes))
	for i, attribute := range attributes {
		responses[i] = FromEntity(attribute)
	}
	return responses
}
