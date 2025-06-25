package entity

import (
	"strings"
	"time"

	"pim/src/attribute/domain/entity"

	"github.com/google/uuid"
)

// MarketplaceAttributeMother implementa el patrón Object Mother para crear entities MarketplaceAttribute de prueba
type MarketplaceAttributeMother struct{}

// WithDefaults crea un atributo marketplace con valores por defecto
func (MarketplaceAttributeMother) WithDefaults() *entity.MarketplaceAttribute {
	now := time.Now()
	description := "Descripción del atributo marketplace"
	return &entity.MarketplaceAttribute{
		ID:            uuid.New().String(),
		Name:          "Atributo Marketplace",
		Type:          "text",
		Description:   &description,
		IsRequired:    false,
		IsFilterable:  true,
		IsSearchable:  true,
		AllowedValues: []string{},
		IsActive:      true,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

// WithID crea un atributo marketplace con un ID específico
func (m MarketplaceAttributeMother) WithID(id string) *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.ID = id
	return attribute
}

// WithName crea un atributo marketplace con un nombre específico
func (m MarketplaceAttributeMother) WithName(name string) *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Name = name
	return attribute
}

// WithType crea un atributo marketplace con un tipo específico
func (m MarketplaceAttributeMother) WithType(attributeType string) *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = attributeType
	return attribute
}

// WithAllowedValues crea un atributo marketplace con valores permitidos específicos
func (m MarketplaceAttributeMother) WithAllowedValues(allowedValues []string) *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.AllowedValues = allowedValues
	return attribute
}

// Required crea un atributo marketplace requerido
func (m MarketplaceAttributeMother) Required() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.IsRequired = true
	return attribute
}

// NotFilterable crea un atributo marketplace no filtrable
func (m MarketplaceAttributeMother) NotFilterable() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.IsFilterable = false
	return attribute
}

// NotSearchable crea un atributo marketplace no buscable
func (m MarketplaceAttributeMother) NotSearchable() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.IsSearchable = false
	return attribute
}

// Inactive crea un atributo marketplace inactivo
func (m MarketplaceAttributeMother) Inactive() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.IsActive = false
	return attribute
}

// ColorAttribute crea un atributo de color para marketplace
func (m MarketplaceAttributeMother) ColorAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Name = "Color"
	attribute.Type = "enum"
	attribute.AllowedValues = []string{"Rojo", "Azul", "Verde", "Negro", "Blanco"}
	attribute.IsFilterable = true
	return attribute
}

// SizeAttribute crea un atributo de talla para marketplace
func (m MarketplaceAttributeMother) SizeAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Name = "Talla"
	attribute.Type = "enum"
	attribute.AllowedValues = []string{"XS", "S", "M", "L", "XL", "XXL"}
	attribute.IsFilterable = true
	return attribute
}

// BrandAttribute crea un atributo de marca para marketplace
func (m MarketplaceAttributeMother) BrandAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Name = "Marca"
	attribute.Type = "string"
	attribute.IsRequired = true
	attribute.IsFilterable = true
	attribute.IsSearchable = true
	return attribute
}

// PriceAttribute crea un atributo de precio para marketplace
func (m MarketplaceAttributeMother) PriceAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Name = "Precio"
	attribute.Type = "number"
	attribute.IsRequired = true
	attribute.IsFilterable = true
	attribute.IsSearchable = false
	return attribute
}

// TextAttribute crea un atributo de texto para marketplace
func (m MarketplaceAttributeMother) TextAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = "text"
	attribute.AllowedValues = []string{}
	return attribute
}

// NumberAttribute crea un atributo numérico para marketplace
func (m MarketplaceAttributeMother) NumberAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = "number"
	attribute.AllowedValues = []string{}
	return attribute
}

// BooleanAttribute crea un atributo booleano para marketplace
func (m MarketplaceAttributeMother) BooleanAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = "boolean"
	attribute.AllowedValues = []string{}
	return attribute
}

// EnumAttribute crea un atributo de enumeración para marketplace
func (m MarketplaceAttributeMother) EnumAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = "enum"
	attribute.AllowedValues = []string{"Opción 1", "Opción 2", "Opción 3"}
	return attribute
}

// DateAttribute crea un atributo de fecha para marketplace
func (m MarketplaceAttributeMother) DateAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = "date"
	attribute.AllowedValues = []string{}
	return attribute
}

// StringAttribute crea un atributo de cadena para marketplace
func (m MarketplaceAttributeMother) StringAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = "string"
	attribute.AllowedValues = []string{}
	return attribute
}

// Complete crea un atributo marketplace con todos los parámetros especificados
func (MarketplaceAttributeMother) Complete(
	id, name, attributeType string,
	description *string,
	isRequired, isFilterable, isSearchable, isActive bool,
	allowedValues []string,
) *entity.MarketplaceAttribute {
	now := time.Now()
	return &entity.MarketplaceAttribute{
		ID:            id,
		Name:          name,
		Type:          attributeType,
		Description:   description,
		IsRequired:    isRequired,
		IsFilterable:  isFilterable,
		IsSearchable:  isSearchable,
		AllowedValues: allowedValues,
		IsActive:      isActive,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

// Helper function to generate slug from name (not used but kept for compatibility)
func generateSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
