package entity

import (
	"errors"
	"time"
)

// ErrInvalidTenantCustomAttribute representa errores relacionados con atributos custom de tenant
var ErrInvalidTenantCustomAttribute = errors.New("atributo custom de tenant inválido")

// TenantCustomAttribute representa un atributo completamente personalizado por un tenant
type TenantCustomAttribute struct {
	ID                    string
	TenantID              string
	MarketplaceCategoryID *string // Puede ser null para atributos globales del tenant
	Name                  string
	Slug                  string
	Type                  string // text, number, boolean, select, multi_select
	IsFilterable          bool
	IsSearchable          bool
	ValidationRules       map[string]interface{} // JSON con reglas de validación
	SortOrder             int
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

// NewTenantCustomAttribute crea una nueva instancia de TenantCustomAttribute
func NewTenantCustomAttribute(
	tenantID string,
	marketplaceCategoryID *string,
	name, slug, attributeType string,
) (*TenantCustomAttribute, error) {
	if tenantID == "" {
		return nil, errors.New("el ID del tenant es obligatorio")
	}

	if name == "" {
		return nil, errors.New("el nombre del atributo es obligatorio")
	}

	if slug == "" {
		return nil, errors.New("el slug del atributo es obligatorio")
	}

	// Validar tipo
	validTypes := map[string]bool{
		"text":         true,
		"number":       true,
		"boolean":      true,
		"select":       true,
		"multi_select": true,
	}

	if !validTypes[attributeType] {
		return nil, errors.New("tipo de atributo inválido")
	}

	now := time.Now()
	return &TenantCustomAttribute{
		TenantID:              tenantID,
		MarketplaceCategoryID: marketplaceCategoryID,
		Name:                  name,
		Slug:                  slug,
		Type:                  attributeType,
		IsFilterable:          false,
		IsSearchable:          false,
		ValidationRules:       make(map[string]interface{}),
		SortOrder:             0,
		CreatedAt:             now,
		UpdatedAt:             now,
	}, nil
}

// Update actualiza los campos del atributo custom
func (tca *TenantCustomAttribute) Update(name, slug string) error {
	if name == "" {
		return errors.New("el nombre del atributo es obligatorio")
	}

	if slug == "" {
		return errors.New("el slug del atributo es obligatorio")
	}

	tca.Name = name
	tca.Slug = slug
	tca.UpdatedAt = time.Now()
	return nil
}

// SetFilterable establece si el atributo es filterable
func (tca *TenantCustomAttribute) SetFilterable(filterable bool) {
	tca.IsFilterable = filterable
	tca.UpdatedAt = time.Now()
}

// SetSearchable establece si el atributo es buscable
func (tca *TenantCustomAttribute) SetSearchable(searchable bool) {
	tca.IsSearchable = searchable
	tca.UpdatedAt = time.Now()
}

// UpdateSortOrder actualiza el orden del atributo
func (tca *TenantCustomAttribute) UpdateSortOrder(sortOrder int) {
	tca.SortOrder = sortOrder
	tca.UpdatedAt = time.Now()
}

// SetValidationRules establece las reglas de validación
func (tca *TenantCustomAttribute) SetValidationRules(rules map[string]interface{}) {
	if rules == nil {
		tca.ValidationRules = make(map[string]interface{})
	} else {
		tca.ValidationRules = rules
	}
	tca.UpdatedAt = time.Now()
}

// IsSelectType verifica si el atributo es de tipo select o multi_select
func (tca *TenantCustomAttribute) IsSelectType() bool {
	return tca.Type == "select" || tca.Type == "multi_select"
}

// IsGlobal verifica si el atributo es global del tenant (no específico de categoría)
func (tca *TenantCustomAttribute) IsGlobal() bool {
	return tca.MarketplaceCategoryID == nil
}

// ValidateUniqueSlug verifica que el slug sea único para el tenant y categoría
func (tca *TenantCustomAttribute) ValidateUniqueSlug(existingAttributes []TenantCustomAttribute) error {
	for _, attr := range existingAttributes {
		if attr.ID != tca.ID &&
			attr.TenantID == tca.TenantID &&
			attr.Slug == tca.Slug {
			// Verificar si están en la misma categoría o ambos son globales
			if (attr.MarketplaceCategoryID == nil && tca.MarketplaceCategoryID == nil) ||
				(attr.MarketplaceCategoryID != nil && tca.MarketplaceCategoryID != nil &&
					*attr.MarketplaceCategoryID == *tca.MarketplaceCategoryID) {
				return errors.New("ya existe un atributo con este slug en el tenant")
			}
		}
	}
	return nil
}
