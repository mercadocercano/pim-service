package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// MarketplaceAttribute representa un atributo global del marketplace
type MarketplaceAttribute struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"` // string, number, boolean, date, enum, text
	Description   *string   `json:"description"`
	IsRequired    bool      `json:"is_required"`
	IsSearchable  bool      `json:"is_searchable"`
	IsFilterable  bool      `json:"is_filterable"`
	AllowedValues []string  `json:"allowed_values"` // Para tipo enum
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// NewMarketplaceAttribute crea una nueva instancia de MarketplaceAttribute
func NewMarketplaceAttribute(
	name string,
	attributeType string,
	description *string,
	isRequired bool,
	isSearchable bool,
	isFilterable bool,
	allowedValues []string,
) (*MarketplaceAttribute, error) {
	if name == "" {
		return nil, fmt.Errorf("name es requerido")
	}

	if attributeType == "" {
		return nil, fmt.Errorf("type es requerido")
	}

	// Validar tipos permitidos
	validTypes := []string{"string", "number", "boolean", "date", "enum", "text"}
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

	// Validar que tipo enum tenga valores permitidos
	if attributeType == "enum" && len(allowedValues) == 0 {
		return nil, fmt.Errorf("atributos de tipo enum deben tener valores permitidos")
	}

	now := time.Now()
	return &MarketplaceAttribute{
		ID:            uuid.New().String(),
		Name:          name,
		Type:          attributeType,
		Description:   description,
		IsRequired:    isRequired,
		IsSearchable:  isSearchable,
		IsFilterable:  isFilterable,
		AllowedValues: allowedValues,
		IsActive:      true,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}

// Update actualiza los campos de la entidad
func (e *MarketplaceAttribute) Update() {
	e.UpdatedAt = time.Now()
}

// Activate activa el atributo
func (e *MarketplaceAttribute) Activate() {
	e.IsActive = true
	e.Update()
}

// Deactivate desactiva el atributo
func (e *MarketplaceAttribute) Deactivate() {
	e.IsActive = false
	e.Update()
}
