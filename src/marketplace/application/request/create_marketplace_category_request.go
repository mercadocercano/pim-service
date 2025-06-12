package request

import (
	"errors"
	"regexp"
	"strings"
)

// CreateMarketplaceCategoryRequest representa la petición para crear una categoría marketplace
type CreateMarketplaceCategoryRequest struct {
	Name        string  `json:"name" binding:"required"`
	Slug        string  `json:"slug" binding:"required"`
	Description *string `json:"description"`
	ParentID    *string `json:"parent_id"`
	SortOrder   int     `json:"sort_order"`
}

// Validate valida los datos de la petición
func (r *CreateMarketplaceCategoryRequest) Validate() error {
	// Validar nombre
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required")
	}

	if len(r.Name) > 255 {
		return errors.New("name cannot exceed 255 characters")
	}

	// Validar slug
	if strings.TrimSpace(r.Slug) == "" {
		return errors.New("slug is required")
	}

	if len(r.Slug) > 255 {
		return errors.New("slug cannot exceed 255 characters")
	}

	// Validar formato del slug (solo letras, números, guiones)
	slugPattern := regexp.MustCompile(`^[a-z0-9-]+$`)
	if !slugPattern.MatchString(r.Slug) {
		return errors.New("slug must contain only lowercase letters, numbers, and hyphens")
	}

	// Validar descripción si se proporciona
	if r.Description != nil && len(*r.Description) > 1000 {
		return errors.New("description cannot exceed 1000 characters")
	}

	// Validar parent_id si se proporciona
	if r.ParentID != nil && strings.TrimSpace(*r.ParentID) == "" {
		return errors.New("parent_id cannot be empty string")
	}

	// Validar sort_order
	if r.SortOrder < 0 {
		return errors.New("sort_order cannot be negative")
	}

	return nil
}
