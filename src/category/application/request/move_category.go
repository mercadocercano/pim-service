package request

// MoveCategoryRequest representa la solicitud para mover una categoría a un nuevo padre
type MoveCategoryRequest struct {
	ParentID string `json:"parent_id,omitempty"`
}
