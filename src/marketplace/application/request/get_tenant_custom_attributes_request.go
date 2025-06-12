package request

import (
	"errors"
	"strings"
)

// GetTenantCustomAttributesRequest representa la petición para obtener atributos personalizados de un tenant
type GetTenantCustomAttributesRequest struct {
	TenantID              string  `json:"tenant_id" binding:"required"`
	MarketplaceCategoryID *string `json:"marketplace_category_id,omitempty"` // filtrar por categoría específica
	AttributeType         *string `json:"attribute_type,omitempty"`          // filtrar por tipo (text, number, select, etc.)
	IsFilterable          *bool   `json:"is_filterable,omitempty"`           // filtrar solo atributos filtrables
	IsSearchable          *bool   `json:"is_searchable,omitempty"`           // filtrar solo atributos buscables
	IncludeInactive       bool    `json:"include_inactive"`                  // incluir atributos inactivos
}

// Validate valida los datos de la petición
func (r *GetTenantCustomAttributesRequest) Validate() error {
	// Validar tenant_id
	if strings.TrimSpace(r.TenantID) == "" {
		return errors.New("tenant_id is required")
	}

	// Validar marketplace_category_id si se proporciona
	if r.MarketplaceCategoryID != nil && strings.TrimSpace(*r.MarketplaceCategoryID) == "" {
		return errors.New("marketplace_category_id cannot be empty")
	}

	// Validar attribute_type si se proporciona
	if r.AttributeType != nil {
		validTypes := map[string]bool{
			"text":         true,
			"number":       true,
			"boolean":      true,
			"select":       true,
			"multi_select": true,
			"date":         true,
			"datetime":     true,
			"url":          true,
			"email":        true,
		}

		if !validTypes[*r.AttributeType] {
			return errors.New("invalid attribute_type")
		}
	}

	return nil
}
