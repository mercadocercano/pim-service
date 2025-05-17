package response

import (
	"pim/src/category/domain/entity"
)

// CategoryTreeResponse representa una categoría en formato de árbol con hijos
type CategoryTreeResponse struct {
	ID          string                  `json:"id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	ParentID    string                  `json:"parent_id,omitempty"`
	Active      bool                    `json:"active"`
	CreatedAt   string                  `json:"created_at"`
	UpdatedAt   string                  `json:"updated_at"`
	Children    []*CategoryTreeResponse `json:"children,omitempty"`
}

// BuildCategoryTree convierte una lista plana de categorías en una estructura jerárquica de árbol
func BuildCategoryTree(categories []*entity.Category) []*CategoryTreeResponse {
	// Mapa para acceder rápidamente a las categorías por ID
	categoryMap := make(map[string]*CategoryTreeResponse)

	// Primero, convertimos todas las categorías a la estructura de respuesta de árbol
	for _, category := range categories {
		var parentID string
		if category.ParentID != nil {
			parentID = *category.ParentID
		}

		categoryMap[category.ID] = &CategoryTreeResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			ParentID:    parentID,
			Active:      category.IsActive(),
			CreatedAt:   category.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   category.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			Children:    []*CategoryTreeResponse{},
		}
	}

	// Lista para almacenar solo las categorías raíz (sin padre)
	var rootCategories []*CategoryTreeResponse

	// Ahora organizamos las categorías en estructura de árbol
	for _, category := range categories {
		categoryResponse := categoryMap[category.ID]

		if category.ParentID == nil || *category.ParentID == "" {
			// Es una categoría raíz (sin padre)
			rootCategories = append(rootCategories, categoryResponse)
		} else {
			// Es una categoría con padre, la agregamos como hijo
			if parent, exists := categoryMap[*category.ParentID]; exists {
				parent.Children = append(parent.Children, categoryResponse)
			}
		}
	}

	return rootCategories
}
