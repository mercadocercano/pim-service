package entity

import (
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"
)

// MarketplaceValidator contiene todas las validaciones de negocio del dominio marketplace
type MarketplaceValidator struct{}

// NewMarketplaceValidator crea una nueva instancia del validador
func NewMarketplaceValidator() *MarketplaceValidator {
	return &MarketplaceValidator{}
}

// ValidateCategoryHierarchy valida la jerarquía completa de categorías
func (mv *MarketplaceValidator) ValidateCategoryHierarchy(categories []MarketplaceCategory) error {
	tree, err := NewCategoryTree(categories)
	if err != nil {
		return err
	}

	return tree.ValidateHierarchy()
}

// ValidateCategoryName valida el nombre de una categoría
func (mv *MarketplaceValidator) ValidateCategoryName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("el nombre de la categoría no puede estar vacío")
	}

	if utf8.RuneCountInString(name) < 2 {
		return errors.New("el nombre de la categoría debe tener al menos 2 caracteres")
	}

	if utf8.RuneCountInString(name) > 100 {
		return errors.New("el nombre de la categoría no puede exceder 100 caracteres")
	}

	// Verificar que no contenga caracteres especiales problemáticos
	if strings.ContainsAny(name, "<>\"'&") {
		return errors.New("el nombre de la categoría contiene caracteres no permitidos")
	}

	return nil
}

// ValidateCategorySlug valida el slug de una categoría
func (mv *MarketplaceValidator) ValidateCategorySlug(slug string) error {
	if strings.TrimSpace(slug) == "" {
		return errors.New("el slug de la categoría no puede estar vacío")
	}

	// Slug debe ser lowercase, guiones, números, sin espacios
	slugPattern := regexp.MustCompile(`^[a-z0-9-]+$`)
	if !slugPattern.MatchString(slug) {
		return errors.New("el slug debe contener solo letras minúsculas, números y guiones")
	}

	if len(slug) < 2 {
		return errors.New("el slug debe tener al menos 2 caracteres")
	}

	if len(slug) > 50 {
		return errors.New("el slug no puede exceder 50 caracteres")
	}

	// No puede empezar o terminar con guión
	if strings.HasPrefix(slug, "-") || strings.HasSuffix(slug, "-") {
		return errors.New("el slug no puede empezar o terminar con guión")
	}

	// No puede tener guiones consecutivos
	if strings.Contains(slug, "--") {
		return errors.New("el slug no puede tener guiones consecutivos")
	}

	return nil
}

// ValidateAttributeName valida el nombre de un atributo
func (mv *MarketplaceValidator) ValidateAttributeName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("el nombre del atributo no puede estar vacío")
	}

	if utf8.RuneCountInString(name) < 2 {
		return errors.New("el nombre del atributo debe tener al menos 2 caracteres")
	}

	if utf8.RuneCountInString(name) > 50 {
		return errors.New("el nombre del atributo no puede exceder 50 caracteres")
	}

	return nil
}

// ValidateAttributeSlug valida el slug de un atributo
func (mv *MarketplaceValidator) ValidateAttributeSlug(slug string) error {
	return mv.ValidateCategorySlug(slug) // Mismas reglas que categorías
}

// ValidateAttributeType valida el tipo de un atributo
func (mv *MarketplaceValidator) ValidateAttributeType(attrType AttributeType) error {
	switch attrType {
	case AttributeTypeText, AttributeTypeNumber, AttributeTypeBoolean, AttributeTypeSelect, AttributeTypeMultiSelect:
		return nil
	default:
		return errors.New("tipo de atributo no válido")
	}
}

// ValidateAttributeValue valida un valor de atributo
func (mv *MarketplaceValidator) ValidateAttributeValue(value string, attrType AttributeType) error {
	if strings.TrimSpace(value) == "" {
		return errors.New("el valor del atributo no puede estar vacío")
	}

	switch attrType {
	case AttributeTypeText:
		return mv.validateTextValue(value)
	case AttributeTypeNumber:
		return mv.validateNumberValue(value)
	case AttributeTypeBoolean:
		return mv.validateBooleanValue(value)
	case AttributeTypeSelect, AttributeTypeMultiSelect:
		return mv.validateSelectValue(value)
	default:
		return errors.New("tipo de atributo no reconocido")
	}
}

// ValidateTenantCategoryMapping valida un mapeo de categoría tenant
func (mv *MarketplaceValidator) ValidateTenantCategoryMapping(mapping *TenantCategoryMapping, existingMappings []TenantCategoryMapping) error {
	// Validar IDs requeridos
	if mapping.TenantID == "" {
		return errors.New("el ID del tenant es obligatorio")
	}

	if mapping.CategoryID == "" {
		return errors.New("el ID de la categoría tenant es obligatorio")
	}

	if mapping.MarketplaceCategoryID == "" {
		return errors.New("el ID de la categoría marketplace es obligatorio")
	}

	// Validar nombre personalizado si está presente
	if mapping.CustomName != nil && strings.TrimSpace(*mapping.CustomName) != "" {
		if err := mv.ValidateCategoryName(*mapping.CustomName); err != nil {
			return errors.New("nombre personalizado inválido: " + err.Error())
		}
	}

	// Validar unicidad
	return mapping.ValidateUniqueMapping(existingMappings)
}

// ValidateTenantAttributeExtension valida una extensión de atributo tenant
func (mv *MarketplaceValidator) ValidateTenantAttributeExtension(extension *TenantAttributeExtension, existingExtensions []TenantAttributeExtension) error {
	// Validar IDs requeridos
	if extension.TenantID == "" {
		return errors.New("el ID del tenant es obligatorio")
	}

	if extension.MarketplaceAttributeID == "" {
		return errors.New("el ID del atributo marketplace es obligatorio")
	}

	// Validar nombre personalizado si está presente
	if extension.CustomName != nil && strings.TrimSpace(*extension.CustomName) != "" {
		if err := mv.ValidateAttributeName(*extension.CustomName); err != nil {
			return errors.New("nombre personalizado inválido: " + err.Error())
		}
	}

	// Validar valores adicionales
	for _, value := range extension.AdditionalValues {
		if strings.TrimSpace(value) == "" {
			return errors.New("los valores adicionales no pueden estar vacíos")
		}
		if utf8.RuneCountInString(value) > 100 {
			return errors.New("los valores adicionales no pueden exceder 100 caracteres")
		}
	}

	// Validar unicidad
	return extension.ValidateUniqueExtension(existingExtensions)
}

// ValidateBusinessRules valida reglas de negocio específicas del marketplace
func (mv *MarketplaceValidator) ValidateBusinessRules(categories []MarketplaceCategory, attributes []MarketplaceAttribute) error {
	// Validar que hay al menos una categoría raíz
	hasRootCategory := false
	for _, category := range categories {
		if category.ParentID == nil && category.IsActive {
			hasRootCategory = true
			break
		}
	}
	if !hasRootCategory {
		return errors.New("debe existir al menos una categoría raíz activa")
	}

	// Validar que no hay categorías huérfanas
	categoryMap := make(map[string]bool)
	for _, category := range categories {
		categoryMap[category.ID] = true
	}

	for _, category := range categories {
		if category.ParentID != nil {
			if !categoryMap[*category.ParentID] {
				return errors.New("existe una categoría con padre inexistente: " + category.ID)
			}
		}
	}

	// Validar que hay atributos básicos requeridos
	hasColorAttribute := false
	hasSizeAttribute := false
	for _, attribute := range attributes {
		if attribute.Slug == "color" && attribute.IsFilterable {
			hasColorAttribute = true
		}
		if attribute.Slug == "size" && attribute.IsFilterable {
			hasSizeAttribute = true
		}
	}

	if !hasColorAttribute {
		return errors.New("debe existir un atributo 'color' filterable")
	}

	if !hasSizeAttribute {
		return errors.New("debe existir un atributo 'size' filterable")
	}

	return nil
}

// ValidateMarketplaceConsistency valida la consistencia general del marketplace
func (mv *MarketplaceValidator) ValidateMarketplaceConsistency(
	categories []MarketplaceCategory,
	attributes []MarketplaceAttribute,
	categoryMappings []TenantCategoryMapping,
	attributeExtensions []TenantAttributeExtension,
) error {
	// Validar jerarquía de categorías
	if err := mv.ValidateCategoryHierarchy(categories); err != nil {
		return errors.New("error en jerarquía de categorías: " + err.Error())
	}

	// Validar reglas de negocio
	if err := mv.ValidateBusinessRules(categories, attributes); err != nil {
		return errors.New("error en reglas de negocio: " + err.Error())
	}

	// Validar que todos los mapeos de categorías referencian categorías existentes
	categoryMap := make(map[string]bool)
	for _, category := range categories {
		categoryMap[category.ID] = true
	}

	for _, mapping := range categoryMappings {
		if !categoryMap[mapping.MarketplaceCategoryID] {
			return errors.New("mapeo de categoría referencia categoría marketplace inexistente: " + mapping.MarketplaceCategoryID)
		}
	}

	// Validar que todas las extensiones de atributos referencian atributos existentes
	attributeMap := make(map[string]bool)
	for _, attribute := range attributes {
		attributeMap[attribute.ID] = true
	}

	for _, extension := range attributeExtensions {
		if !attributeMap[extension.MarketplaceAttributeID] {
			return errors.New("extensión de atributo referencia atributo marketplace inexistente: " + extension.MarketplaceAttributeID)
		}
	}

	return nil
}

// Métodos privados para validaciones específicas por tipo

func (mv *MarketplaceValidator) validateTextValue(value string) error {
	if utf8.RuneCountInString(value) > 500 {
		return errors.New("el valor de texto no puede exceder 500 caracteres")
	}
	return nil
}

func (mv *MarketplaceValidator) validateNumberValue(value string) error {
	// Verificar que sea un número válido
	numberPattern := regexp.MustCompile(`^-?\d+(\.\d+)?$`)
	if !numberPattern.MatchString(value) {
		return errors.New("el valor debe ser un número válido")
	}
	return nil
}

func (mv *MarketplaceValidator) validateBooleanValue(value string) error {
	lowerValue := strings.ToLower(value)
	if lowerValue != "true" && lowerValue != "false" && lowerValue != "1" && lowerValue != "0" {
		return errors.New("el valor booleano debe ser true, false, 1 o 0")
	}
	return nil
}

func (mv *MarketplaceValidator) validateSelectValue(value string) error {
	if utf8.RuneCountInString(value) > 100 {
		return errors.New("el valor de selección no puede exceder 100 caracteres")
	}

	// No debe contener caracteres de control
	if strings.ContainsAny(value, "\n\r\t") {
		return errors.New("el valor de selección no puede contener saltos de línea o tabulaciones")
	}

	return nil
}
