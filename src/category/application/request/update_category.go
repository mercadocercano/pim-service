package request

// UpdateCategoryRequest representa la solicitud para actualizar una categoría
type UpdateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ParentID    string `json:"parent_id,omitempty"`
}
