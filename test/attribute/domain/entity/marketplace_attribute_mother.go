package entity

import (
	"strings"
	"time"

	"saas-mt-pim-service/src/attribute/domain/entity"

	"github.com/google/uuid"
)

// MarketplaceAttributeMother implementa el patrón Object Mother para crear entities MarketplaceAttribute de prueba
type MarketplaceAttributeMother struct{}

// WithDefaults crea un atributo marketplace con valores por defecto
func (MarketplaceAttributeMother) WithDefaults() *entity.MarketplaceAttribute {
	now := time.Now()
	return &entity.MarketplaceAttribute{
		ID:                   uuid.New().String(),
		Name:                 "Atributo Marketplace",
		Slug:                 "atributo-marketplace",
		Type:                 "text",
		IsFilterable:         true,
		IsSearchable:         true,
		IsRequiredForListing: false,
		ValidationRules:      make(map[string]interface{}),
		SortOrder:            0,
		CreatedAt:            now,
		UpdatedAt:            now,
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
	attribute.Slug = generateSlug(name)
	return attribute
}

// WithSlug crea un atributo marketplace con un slug específico
func (m MarketplaceAttributeMother) WithSlug(slug string) *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Slug = slug
	return attribute
}

// WithType crea un atributo marketplace con un tipo específico
func (m MarketplaceAttributeMother) WithType(attributeType string) *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = attributeType
	return attribute
}

// WithValidationRules crea un atributo marketplace con reglas de validación específicas
func (m MarketplaceAttributeMother) WithValidationRules(rules map[string]interface{}) *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.ValidationRules = rules
	return attribute
}

// WithSortOrder crea un atributo marketplace con un orden específico
func (m MarketplaceAttributeMother) WithSortOrder(sortOrder int) *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.SortOrder = sortOrder
	return attribute
}

// RequiredForListing crea un atributo marketplace requerido para listar
func (m MarketplaceAttributeMother) RequiredForListing() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.IsRequiredForListing = true
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

// ColorAttribute crea un atributo de color para marketplace
func (m MarketplaceAttributeMother) ColorAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Name = "Color"
	attribute.Slug = "color"
	attribute.Type = "select"
	attribute.IsFilterable = true
	attribute.ValidationRules = map[string]interface{}{
		"options": []string{"Rojo", "Azul", "Verde", "Negro", "Blanco"},
	}
	return attribute
}

// SizeAttribute crea un atributo de talla para marketplace
func (m MarketplaceAttributeMother) SizeAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Name = "Talla"
	attribute.Slug = "talla"
	attribute.Type = "select"
	attribute.IsFilterable = true
	attribute.ValidationRules = map[string]interface{}{
		"options": []string{"XS", "S", "M", "L", "XL", "XXL"},
	}
	return attribute
}

// BrandAttribute crea un atributo de marca para marketplace
func (m MarketplaceAttributeMother) BrandAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Name = "Marca"
	attribute.Slug = "marca"
	attribute.Type = "text"
	attribute.IsRequiredForListing = true
	attribute.IsFilterable = true
	attribute.IsSearchable = true
	return attribute
}

// PriceAttribute crea un atributo de precio para marketplace
func (m MarketplaceAttributeMother) PriceAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Name = "Precio"
	attribute.Slug = "precio"
	attribute.Type = "number"
	attribute.IsRequiredForListing = true
	attribute.IsFilterable = true
	attribute.IsSearchable = false
	attribute.ValidationRules = map[string]interface{}{
		"min": 0,
		"max": 999999,
	}
	return attribute
}

// TextAttribute crea un atributo de texto para marketplace
func (m MarketplaceAttributeMother) TextAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = "text"
	return attribute
}

// NumberAttribute crea un atributo numérico para marketplace
func (m MarketplaceAttributeMother) NumberAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = "number"
	return attribute
}

// BooleanAttribute crea un atributo booleano para marketplace
func (m MarketplaceAttributeMother) BooleanAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = "boolean"
	return attribute
}

// SelectAttribute crea un atributo de selección para marketplace
func (m MarketplaceAttributeMother) SelectAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = "select"
	attribute.ValidationRules = map[string]interface{}{
		"options": []string{"Opción 1", "Opción 2", "Opción 3"},
	}
	return attribute
}

// MultiSelectAttribute crea un atributo de selección múltiple para marketplace
func (m MarketplaceAttributeMother) MultiSelectAttribute() *entity.MarketplaceAttribute {
	attribute := m.WithDefaults()
	attribute.Type = "multi_select"
	attribute.ValidationRules = map[string]interface{}{
		"options": []string{"Opción A", "Opción B", "Opción C"},
	}
	return attribute
}

// Complete crea un atributo marketplace con todos los parámetros especificados
func (MarketplaceAttributeMother) Complete(
	id, name, slug, attributeType string,
	isFilterable, isSearchable, isRequiredForListing bool,
	validationRules map[string]interface{},
	sortOrder int,
) *entity.MarketplaceAttribute {
	now := time.Now()
	return &entity.MarketplaceAttribute{
		ID:                   id,
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
	}
}

// generateSlug genera un slug a partir de un nombre
func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "á", "a")
	slug = strings.ReplaceAll(slug, "é", "e")
	slug = strings.ReplaceAll(slug, "í", "i")
	slug = strings.ReplaceAll(slug, "ó", "o")
	slug = strings.ReplaceAll(slug, "ú", "u")
	slug = strings.ReplaceAll(slug, "ñ", "n")
	return slug
}
