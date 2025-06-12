package request

import (
	"errors"
	"strings"
)

// MapTenantCategoryRequest representa la petición para mapear una categoría tenant a marketplace
type MapTenantCategoryRequest struct {
	CategoryID            string `json:"category_id" binding:"required"`
	MarketplaceCategoryID string `json:"marketplace_category_id" binding:"required"`
	CustomName            string `json:"custom_name" binding:"required"`
}

// Validate valida los datos de la petición
func (r *MapTenantCategoryRequest) Validate() error {
	// Validar category_id
	if strings.TrimSpace(r.CategoryID) == "" {
		return errors.New("category_id is required")
	}

	// Validar marketplace_category_id
	if strings.TrimSpace(r.MarketplaceCategoryID) == "" {
		return errors.New("marketplace_category_id is required")
	}

	// Validar custom_name
	if strings.TrimSpace(r.CustomName) == "" {
		return errors.New("custom_name is required")
	}

	if len(r.CustomName) > 255 {
		return errors.New("custom_name cannot exceed 255 characters")
	}

	return nil
}
