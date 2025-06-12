package request

import (
	"errors"
	"strings"
)

// ExtendTenantAttributesRequest representa la petición para extender atributos de tenant
type ExtendTenantAttributesRequest struct {
	MarketplaceCategoryID string                   `json:"marketplace_category_id" binding:"required"`
	CustomAttributes      []CustomAttributeRequest `json:"custom_attributes" binding:"required"`
}

// CustomAttributeRequest representa un atributo personalizado del tenant
type CustomAttributeRequest struct {
	Name         string   `json:"name" binding:"required"`
	Type         string   `json:"type" binding:"required"` // text, number, boolean, select, multi_select
	IsRequired   bool     `json:"is_required"`
	IsFilterable bool     `json:"is_filterable"`
	Options      []string `json:"options,omitempty"` // Para tipos select/multi_select
	DefaultValue *string  `json:"default_value,omitempty"`
}

// Validate valida los datos de la petición
func (r *ExtendTenantAttributesRequest) Validate() error {
	// Validar marketplace_category_id
	if strings.TrimSpace(r.MarketplaceCategoryID) == "" {
		return errors.New("marketplace_category_id is required")
	}

	// Validar que hay al menos un atributo
	if len(r.CustomAttributes) == 0 {
		return errors.New("at least one custom attribute is required")
	}

	// Validar cada atributo personalizado
	attributeNames := make(map[string]bool)
	for i, attr := range r.CustomAttributes {
		if err := attr.Validate(); err != nil {
			return errors.New("custom_attributes[" + string(rune(i)) + "]: " + err.Error())
		}

		// Verificar nombres únicos
		if attributeNames[attr.Name] {
			return errors.New("duplicate attribute name: " + attr.Name)
		}
		attributeNames[attr.Name] = true
	}

	return nil
}

// Validate valida un atributo personalizado individual
func (c *CustomAttributeRequest) Validate() error {
	// Validar nombre
	if strings.TrimSpace(c.Name) == "" {
		return errors.New("name is required")
	}

	if len(c.Name) > 255 {
		return errors.New("name cannot exceed 255 characters")
	}

	// Validar tipo
	validTypes := map[string]bool{
		"text":         true,
		"number":       true,
		"boolean":      true,
		"select":       true,
		"multi_select": true,
	}

	if !validTypes[c.Type] {
		return errors.New("invalid type. Must be one of: text, number, boolean, select, multi_select")
	}

	// Validar opciones para tipos select
	if (c.Type == "select" || c.Type == "multi_select") && len(c.Options) == 0 {
		return errors.New("options are required for select and multi_select types")
	}

	// Validar que no hay opciones para otros tipos
	if c.Type != "select" && c.Type != "multi_select" && len(c.Options) > 0 {
		return errors.New("options are only allowed for select and multi_select types")
	}

	return nil
}
