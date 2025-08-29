package response

import (
	"time"

	"saas-mt-pim-service/src/category/domain/entity"
)

// CategoryResponse representa la respuesta con datos de una categoría
type CategoryResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentID    string    `json:"parent_id,omitempty"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// FromEntity convierte una entidad de categoría en una respuesta
func FromEntity(category *entity.Category) *CategoryResponse {
	var parentID string
	if category.ParentID != nil {
		parentID = *category.ParentID
	}

	return &CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ParentID:    parentID,
		Active:      category.IsActive(),
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}

// FromEntityList convierte una lista de entidades en una lista de respuestas
func FromEntityList(categories []*entity.Category) []*CategoryResponse {
	var response []*CategoryResponse
	for _, category := range categories {
		response = append(response, FromEntity(category))
	}
	return response
}
