package request

import (
	"errors"
	"strings"
)

// UpdateTenantCustomAttributeRequest representa la petición para actualizar un atributo personalizado de un tenant
type UpdateTenantCustomAttributeRequest struct {
	Name            *string                `json:"name,omitempty"`             // nuevo nombre del atributo
	IsFilterable    *bool                  `json:"is_filterable,omitempty"`    // si el atributo es filtrable
	IsSearchable    *bool                  `json:"is_searchable,omitempty"`    // si el atributo es buscable
	SortOrder       *int                   `json:"sort_order,omitempty"`       // orden de clasificación
	ValidationRules map[string]interface{} `json:"validation_rules,omitempty"` // reglas de validación actualizadas
}

// Validate valida los datos de la petición
func (r *UpdateTenantCustomAttributeRequest) Validate() error {
	// Validar nombre si se proporciona
	if r.Name != nil && strings.TrimSpace(*r.Name) == "" {
		return errors.New("name cannot be empty")
	}

	// Validar sort_order si se proporciona
	if r.SortOrder != nil && *r.SortOrder < 0 {
		return errors.New("sort_order must be non-negative")
	}

	// Validar que al menos un campo esté presente para actualizar
	if r.Name == nil && r.IsFilterable == nil && r.IsSearchable == nil &&
		r.SortOrder == nil && r.ValidationRules == nil {
		return errors.New("at least one field must be provided for update")
	}

	return nil
}
