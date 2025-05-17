package request

// CreateCategoryRequest representa la solicitud para crear una nueva categoría
type CreateCategoryRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	ParentID    *string `json:"parent_id,omitempty"`
}
