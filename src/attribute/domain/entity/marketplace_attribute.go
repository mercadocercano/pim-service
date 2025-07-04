package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// MarketplaceAttribute representa un atributo global del marketplace
type MarketplaceAttribute struct {
	ID                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	Slug                 string                 `json:"slug"`
	Type                 string                 `json:"type"` // text, number, boolean, select, multi_select
	IsFilterable         bool                   `json:"is_filterable"`
	IsSearchable         bool                   `json:"is_searchable"`
	IsRequiredForListing bool                   `json:"is_required_for_listing"`
	ValidationRules      map[string]interface{} `json:"validation_rules"`
	SortOrder            int                    `json:"sort_order"`
	CreatedAt            time.Time              `json:"created_at"`
	UpdatedAt            time.Time              `json:"updated_at"`
}

// NewMarketplaceAttribute crea una nueva instancia de MarketplaceAttribute
func NewMarketplaceAttribute(
	name string,
	slug string,
	attributeType string,
	isFilterable bool,
	isSearchable bool,
	isRequiredForListing bool,
	validationRules map[string]interface{},
	sortOrder int,
) (*MarketplaceAttribute, error) {
	if name == "" {
		return nil, fmt.Errorf("name es requerido")
	}

	if slug == "" {
		return nil, fmt.Errorf("slug es requerido")
	}

	if attributeType == "" {
		return nil, fmt.Errorf("type es requerido")
	}

	// Validar tipos permitidos según la constraint de la tabla
	validTypes := []string{"text", "number", "boolean", "select", "multi_select"}
	isValidType := false
	for _, validType := range validTypes {
		if attributeType == validType {
			isValidType = true
			break
		}
	}
	if !isValidType {
		return nil, fmt.Errorf("tipo de atributo inválido: %s", attributeType)
	}

	if validationRules == nil {
		validationRules = make(map[string]interface{})
	}

	now := time.Now()
	return &MarketplaceAttribute{
		ID:                   uuid.New().String(),
		Name:                 name,
		Slug:                 slug,
		Type:                 attributeType,
		IsFilterable:         isFilterable,
		IsSearchable:         isSearchable,
		IsRequiredForListing: isRequiredForListing,
		ValidationRules:      validationRules,
		SortOrder:            sortOrder,
		CreatedAt:            now,
		UpdatedAt:            now,
	}, nil
}

// Update actualiza el timestamp de modificación
func (e *MarketplaceAttribute) Update() {
	e.UpdatedAt = time.Now()
}

// SetFilterable establece si el atributo es filtrable
func (e *MarketplaceAttribute) SetFilterable(filterable bool) {
	e.IsFilterable = filterable
	e.Update()
}

// SetSearchable establece si el atributo es buscable
func (e *MarketplaceAttribute) SetSearchable(searchable bool) {
	e.IsSearchable = searchable
	e.Update()
}

// SetRequiredForListing establece si el atributo es requerido para listar productos
func (e *MarketplaceAttribute) SetRequiredForListing(required bool) {
	e.IsRequiredForListing = required
	e.Update()
}
