package request

import (
	"errors"
	"strings"
)

// ValidateCategoryHierarchyRequest representa la petición para validar jerarquía de categorías
type ValidateCategoryHierarchyRequest struct {
	CategoryID       string  `json:"category_id" binding:"required"`
	NewParentID      *string `json:"new_parent_id,omitempty"` // null para mover a raíz
	MaxDepth         int     `json:"max_depth,omitempty"`     // default 3
	ValidateChildren bool    `json:"validate_children"`       // validar que los hijos no excedan profundidad
}

// Validate valida los datos de la petición
func (r *ValidateCategoryHierarchyRequest) Validate() error {
	// Validar category_id
	if strings.TrimSpace(r.CategoryID) == "" {
		return errors.New("category_id is required")
	}

	// Validar new_parent_id si se proporciona
	if r.NewParentID != nil && strings.TrimSpace(*r.NewParentID) == "" {
		return errors.New("new_parent_id cannot be empty string")
	}

	// Validar que no sea auto-referencia
	if r.NewParentID != nil && *r.NewParentID == r.CategoryID {
		return errors.New("category cannot be its own parent")
	}

	// Validar max_depth
	if r.MaxDepth <= 0 {
		r.MaxDepth = 3 // Default
	}

	if r.MaxDepth > 10 {
		return errors.New("max_depth cannot exceed 10 levels")
	}

	return nil
}
