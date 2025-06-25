package request

import (
	"errors"
	"strings"
)

// GetTenantTaxonomyRequest representa la petición para obtener la taxonomía de un tenant
type GetTenantTaxonomyRequest struct {
	TenantID                string   `json:"tenant_id" binding:"required"`
	IncludeCustomAttributes bool     `json:"include_custom_attributes"` // incluir atributos personalizados
	IncludeMarketplaceData  bool     `json:"include_marketplace_data"`  // incluir datos de categorías marketplace
	CategoryIDs             []string `json:"category_ids,omitempty"`    // filtrar por categorías específicas
	MaxDepth                *int     `json:"max_depth,omitempty"`       // limitar profundidad del árbol
	IncludeInactive         bool     `json:"include_inactive"`          // incluir categorías inactivas
	Format                  string   `json:"format"`                    // tree, flat, hierarchical
}

// Validate valida los datos de la petición
func (r *GetTenantTaxonomyRequest) Validate() error {
	// Validar tenant_id
	if strings.TrimSpace(r.TenantID) == "" {
		return errors.New("tenant_id is required")
	}

	// Validar category_ids si se proporcionan
	for i, categoryID := range r.CategoryIDs {
		if strings.TrimSpace(categoryID) == "" {
			return errors.New("category_ids cannot contain empty strings")
		}
		r.CategoryIDs[i] = strings.TrimSpace(categoryID)
	}

	// Validar max_depth
	if r.MaxDepth != nil && *r.MaxDepth <= 0 {
		return errors.New("max_depth must be greater than 0")
	}

	// Validar format
	validFormats := map[string]bool{
		"tree":         true,
		"flat":         true,
		"hierarchical": true,
	}

	if r.Format == "" {
		r.Format = "tree" // Default
	}

	if !validFormats[r.Format] {
		return errors.New("format must be one of: tree, flat, hierarchical")
	}

	return nil
}
