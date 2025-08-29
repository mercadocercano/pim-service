package response

import (
	"time"

	"saas-mt-pim-service/src/category_attribute/domain/entity"
	"saas-mt-pim-service/src/category_attribute/domain/port"
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

// DetailedCategoryAttributeResponse representa la respuesta detallada con datos completos del atributo
type DetailedCategoryAttributeResponse struct {
	ID               string    `json:"id"`
	CategoryID       string    `json:"category_id"`
	AttributeID      string    `json:"attribute_id"`
	AttributeName    string    `json:"attribute_name"`
	AttributeType    string    `json:"attribute_type"`
	Description      string    `json:"description"`
	Required         bool      `json:"required"`
	AttributeOptions []string  `json:"attribute_options"` // Opciones del atributo base
	AllowedValues    []string  `json:"allowed_values"`    // Valores específicos para esta categoría
	FinalValues      []string  `json:"final_values"`      // Valores finales combinados
	Active           bool      `json:"active"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
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

// FromDetailedEntity convierte un DetailedCategoryAttribute en una respuesta detallada
func FromDetailedEntity(detailed *port.DetailedCategoryAttribute) *DetailedCategoryAttributeResponse {
	if detailed == nil {
		return nil
	}

	// Combinar attribute_options con allowed_values
	finalValues := detailed.AttributeOptions
	if len(detailed.AllowedValues) > 0 {
		// Si hay allowed_values específicos, usar esos en lugar de las opciones generales
		finalValues = detailed.AllowedValues
	}

	return &DetailedCategoryAttributeResponse{
		ID:               detailed.ID,
		CategoryID:       detailed.CategoryID,
		AttributeID:      detailed.AttributeID,
		AttributeName:    detailed.AttributeName,
		AttributeType:    detailed.AttributeType,
		Description:      detailed.Description,
		Required:         detailed.Required,
		AttributeOptions: detailed.AttributeOptions,
		AllowedValues:    detailed.AllowedValues,
		FinalValues:      finalValues,
		Active:           detailed.Active,
		CreatedAt:        detailed.CreatedAt,
		UpdatedAt:        detailed.UpdatedAt,
	}
}

// FromDetailedEntityList convierte una lista de DetailedCategoryAttribute en una lista de respuestas detalladas
func FromDetailedEntityList(detailedAttributes []*port.DetailedCategoryAttribute) []*DetailedCategoryAttributeResponse {
	var response []*DetailedCategoryAttributeResponse
	for _, detailed := range detailedAttributes {
		response = append(response, FromDetailedEntity(detailed))
	}
	return response
}
