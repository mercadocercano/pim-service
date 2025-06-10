package entity

import (
	"errors"
	"time"
)

// AttributeType define los tipos de atributos disponibles
type AttributeType string

const (
	AttributeTypeText        AttributeType = "text"
	AttributeTypeNumber      AttributeType = "number"
	AttributeTypeBoolean     AttributeType = "boolean"
	AttributeTypeSelect      AttributeType = "select"
	AttributeTypeMultiSelect AttributeType = "multi_select"
)

// ErrInvalidMarketplaceAttribute representa errores relacionados con la validación de atributos marketplace
var ErrInvalidMarketplaceAttribute = errors.New("atributo marketplace inválido")

// MarketplaceAttribute representa un atributo global del marketplace
type MarketplaceAttribute struct {
	ID                   string
	Name                 string
	Slug                 string
	Type                 AttributeType
	IsFilterable         bool
	IsSearchable         bool
	IsRequiredForListing bool
	ValidationRules      map[string]interface{}
	SortOrder            int
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Values               []MarketplaceAttributeValue
}

// MarketplaceAttributeValue representa un valor predefinido para atributos tipo select
type MarketplaceAttributeValue struct {
	ID          string
	AttributeID string
	Value       string
	Slug        string
	SortOrder   int
	IsActive    bool
	CreatedAt   time.Time
}

// NewMarketplaceAttribute crea una nueva instancia de MarketplaceAttribute con validaciones
func NewMarketplaceAttribute(name, slug string, attrType AttributeType) (*MarketplaceAttribute, error) {
	if name == "" {
		return nil, errors.New("el nombre del atributo marketplace es obligatorio")
	}

	if slug == "" {
		return nil, errors.New("el slug del atributo marketplace es obligatorio")
	}

	if !isValidAttributeType(attrType) {
		return nil, errors.New("tipo de atributo marketplace inválido")
	}

	now := time.Now()
	return &MarketplaceAttribute{
		Name:                 name,
		Slug:                 slug,
		Type:                 attrType,
		IsFilterable:         false,
		IsSearchable:         false,
		IsRequiredForListing: false,
		ValidationRules:      make(map[string]interface{}),
		SortOrder:            0,
		CreatedAt:            now,
		UpdatedAt:            now,
		Values:               []MarketplaceAttributeValue{},
	}, nil
}

// Update actualiza los campos del atributo marketplace
func (ma *MarketplaceAttribute) Update(name, slug string, attrType AttributeType) error {
	if name == "" {
		return errors.New("el nombre del atributo marketplace es obligatorio")
	}

	if slug == "" {
		return errors.New("el slug del atributo marketplace es obligatorio")
	}

	if !isValidAttributeType(attrType) {
		return errors.New("tipo de atributo marketplace inválido")
	}

	ma.Name = name
	ma.Slug = slug
	ma.Type = attrType
	ma.UpdatedAt = time.Now()
	return nil
}

// SetFilterable establece si el atributo es filterable
func (ma *MarketplaceAttribute) SetFilterable(filterable bool) {
	ma.IsFilterable = filterable
	ma.UpdatedAt = time.Now()
}

// SetSearchable establece si el atributo es buscable
func (ma *MarketplaceAttribute) SetSearchable(searchable bool) {
	ma.IsSearchable = searchable
	ma.UpdatedAt = time.Now()
}

// SetRequired establece si el atributo es requerido para listar
func (ma *MarketplaceAttribute) SetRequired(required bool) {
	ma.IsRequiredForListing = required
	ma.UpdatedAt = time.Time{}
}

// SetValidationRules establece las reglas de validación
func (ma *MarketplaceAttribute) SetValidationRules(rules map[string]interface{}) {
	ma.ValidationRules = rules
	ma.UpdatedAt = time.Now()
}

// UpdateSortOrder actualiza el orden del atributo
func (ma *MarketplaceAttribute) UpdateSortOrder(sortOrder int) {
	ma.SortOrder = sortOrder
	ma.UpdatedAt = time.Now()
}

// AddValue agrega un valor predefinido (para atributos tipo select)
func (ma *MarketplaceAttribute) AddValue(value, slug string, sortOrder int) error {
	if ma.Type != AttributeTypeSelect && ma.Type != AttributeTypeMultiSelect {
		return errors.New("solo los atributos tipo select pueden tener valores predefinidos")
	}

	if value == "" {
		return errors.New("el valor del atributo es obligatorio")
	}

	if slug == "" {
		return errors.New("el slug del valor es obligatorio")
	}

	attributeValue := MarketplaceAttributeValue{
		AttributeID: ma.ID,
		Value:       value,
		Slug:        slug,
		SortOrder:   sortOrder,
		IsActive:    true,
		CreatedAt:   time.Now(),
	}

	ma.Values = append(ma.Values, attributeValue)
	ma.UpdatedAt = time.Now()
	return nil
}

// RemoveValue remueve un valor predefinido
func (ma *MarketplaceAttribute) RemoveValue(valueID string) {
	for i, val := range ma.Values {
		if val.ID == valueID {
			ma.Values = append(ma.Values[:i], ma.Values[i+1:]...)
			break
		}
	}
	ma.UpdatedAt = time.Now()
}

// HasValues verifica si el atributo tiene valores predefinidos
func (ma *MarketplaceAttribute) HasValues() bool {
	return len(ma.Values) > 0
}

// RequiresValues verifica si el atributo debe tener valores predefinidos
func (ma *MarketplaceAttribute) RequiresValues() bool {
	return ma.Type == AttributeTypeSelect || ma.Type == AttributeTypeMultiSelect
}

// ValidateValue valida un valor contra las reglas del atributo
func (ma *MarketplaceAttribute) ValidateValue(value interface{}) error {
	// Implementar validaciones específicas según el tipo y reglas
	if ma.IsRequiredForListing && (value == nil || value == "") {
		return errors.New("el atributo es requerido")
	}

	// Agregar más validaciones según ValidationRules
	return nil
}

// isValidAttributeType verifica si el tipo de atributo es válido
func isValidAttributeType(attrType AttributeType) bool {
	switch attrType {
	case AttributeTypeText, AttributeTypeNumber, AttributeTypeBoolean, AttributeTypeSelect, AttributeTypeMultiSelect:
		return true
	default:
		return false
	}
}
