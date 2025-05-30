package response

import (
	"time"

	"pim/src/category_attribute/domain/entity"
)

// CategoryAttributeResponse representa la respuesta con datos de un atributo de categoría
type CategoryAttributeResponse struct {
	ID            string    `json:"id"`
	CategoryID    string    `json:"category_id"`
	AttributeID   string    `json:"attribute_id"`
	AllowedValues []string  `json:"allowed_values"`
	Active        bool      `json:"active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// FromEntity convierte una entidad de atributo de categoría en una respuesta
func FromEntity(categoryAttribute *entity.CategoryAttribute) *CategoryAttributeResponse {
	return &CategoryAttributeResponse{
		ID:            categoryAttribute.ID,
		CategoryID:    categoryAttribute.CategoryID,
		AttributeID:   categoryAttribute.AttributeID,
		AllowedValues: categoryAttribute.AllowedValues,
		Active:        categoryAttribute.IsActive(),
		CreatedAt:     categoryAttribute.CreatedAt,
		UpdatedAt:     categoryAttribute.UpdatedAt,
	}
}

// FromEntityList convierte una lista de entidades en una lista de respuestas
func FromEntityList(categoryAttributes []*entity.CategoryAttribute) []*CategoryAttributeResponse {
	var response []*CategoryAttributeResponse
	for _, categoryAttribute := range categoryAttributes {
		response = append(response, FromEntity(categoryAttribute))
	}
	return response
}
