package request

import (
	"errors"
	"regexp"
	"strings"
)

// UpdateMarketplaceCategoryRequest representa la petición para actualizar una categoría marketplace
type UpdateMarketplaceCategoryRequest struct {
	Name        *string `json:"name"`
	Slug        *string `json:"slug"`
	Description *string `json:"description"`
	IsActive    *bool   `json:"is_active"`
	SortOrder   *int    `json:"sort_order"`
}

// Validate valida los datos de la petición
func (r *UpdateMarketplaceCategoryRequest) Validate() error {
	// Validar nombre si se proporciona
	if r.Name != nil {
		if strings.TrimSpace(*r.Name) == "" {
			return errors.New("name cannot be empty")
		}

		if len(*r.Name) > 255 {
			return errors.New("name cannot exceed 255 characters")
		}
	}

	// Validar slug si se proporciona
	if r.Slug != nil {
		if strings.TrimSpace(*r.Slug) == "" {
			return errors.New("slug cannot be empty")
		}

		if len(*r.Slug) > 255 {
			return errors.New("slug cannot exceed 255 characters")
		}

		// Validar formato del slug (solo letras, números, guiones)
		slugPattern := regexp.MustCompile(`^[a-z0-9-]+$`)
		if !slugPattern.MatchString(*r.Slug) {
			return errors.New("slug must contain only lowercase letters, numbers, and hyphens")
		}
	}

	// Validar descripción si se proporciona
	if r.Description != nil && len(*r.Description) > 1000 {
		return errors.New("description cannot exceed 1000 characters")
	}

	// Validar sort_order si se proporciona
	if r.SortOrder != nil && *r.SortOrder < 0 {
		return errors.New("sort_order cannot be negative")
	}

	return nil
}
