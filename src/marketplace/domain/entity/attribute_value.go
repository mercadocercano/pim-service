package entity

import (
	"errors"
	"sort"
	"strings"
)

// ErrInvalidAttributeValue representa errores relacionados con valores de atributos
var ErrInvalidAttributeValue = errors.New("valor de atributo inválido")

// AttributeValueSet representa un conjunto de valores para un atributo específico
type AttributeValueSet struct {
	AttributeID            string
	AttributeName          string
	AttributeType          AttributeType
	MarketplaceValues      []AttributeValueItem // Valores globales del marketplace
	TenantAdditionalValues []AttributeValueItem // Valores adicionales del tenant
	AllValues              []AttributeValueItem // Todos los valores combinados
	IsFilterable           bool
	IsSearchable           bool
	ValidationRules        map[string]interface{}
}

// AttributeValueItem representa un valor individual de un atributo
type AttributeValueItem struct {
	ID        string
	Value     string
	Slug      string
	Source    ValueSource // Marketplace o Tenant
	SortOrder int
	IsActive  bool
	Count     int // Número de productos que usan este valor (para filtros)
}

// ValueSource indica el origen del valor
type ValueSource string

const (
	ValueSourceMarketplace ValueSource = "marketplace"
	ValueSourceTenant      ValueSource = "tenant"
)

// NewAttributeValueSet crea un nuevo conjunto de valores para un atributo
func NewAttributeValueSet(attributeID, attributeName string, attributeType AttributeType) *AttributeValueSet {
	return &AttributeValueSet{
		AttributeID:            attributeID,
		AttributeName:          attributeName,
		AttributeType:          attributeType,
		MarketplaceValues:      []AttributeValueItem{},
		TenantAdditionalValues: []AttributeValueItem{},
		AllValues:              []AttributeValueItem{},
		ValidationRules:        make(map[string]interface{}),
	}
}

// AddMarketplaceValue agrega un valor del marketplace
func (avs *AttributeValueSet) AddMarketplaceValue(id, value, slug string, sortOrder int) error {
	if value == "" {
		return errors.New("el valor del atributo no puede estar vacío")
	}

	if slug == "" {
		return errors.New("el slug del valor no puede estar vacío")
	}

	// Verificar que no exista ya
	for _, existing := range avs.MarketplaceValues {
		if existing.Value == value || existing.Slug == slug {
			return errors.New("el valor ya existe en el marketplace")
		}
	}

	valueItem := AttributeValueItem{
		ID:        id,
		Value:     value,
		Slug:      slug,
		Source:    ValueSourceMarketplace,
		SortOrder: sortOrder,
		IsActive:  true,
		Count:     0,
	}

	avs.MarketplaceValues = append(avs.MarketplaceValues, valueItem)
	avs.rebuildAllValues()
	return nil
}

// AddTenantValue agrega un valor adicional del tenant
func (avs *AttributeValueSet) AddTenantValue(value string, sortOrder int) error {
	if value == "" {
		return errors.New("el valor del tenant no puede estar vacío")
	}

	// Verificar que no exista ya en valores del tenant
	for _, existing := range avs.TenantAdditionalValues {
		if existing.Value == value {
			return errors.New("el valor ya existe en los valores del tenant")
		}
	}

	// Verificar que no exista en valores del marketplace
	for _, existing := range avs.MarketplaceValues {
		if existing.Value == value {
			return errors.New("el valor ya existe en el marketplace")
		}
	}

	slug := strings.ToLower(strings.ReplaceAll(value, " ", "-"))
	valueItem := AttributeValueItem{
		ID:        "", // Se asignará cuando se persista
		Value:     value,
		Slug:      slug,
		Source:    ValueSourceTenant,
		SortOrder: sortOrder,
		IsActive:  true,
		Count:     0,
	}

	avs.TenantAdditionalValues = append(avs.TenantAdditionalValues, valueItem)
	avs.rebuildAllValues()
	return nil
}

// RemoveMarketplaceValue elimina un valor del marketplace
func (avs *AttributeValueSet) RemoveMarketplaceValue(valueID string) {
	for i, value := range avs.MarketplaceValues {
		if value.ID == valueID {
			avs.MarketplaceValues = append(avs.MarketplaceValues[:i], avs.MarketplaceValues[i+1:]...)
			avs.rebuildAllValues()
			break
		}
	}
}

// RemoveTenantValue elimina un valor del tenant
func (avs *AttributeValueSet) RemoveTenantValue(value string) {
	for i, val := range avs.TenantAdditionalValues {
		if val.Value == value {
			avs.TenantAdditionalValues = append(avs.TenantAdditionalValues[:i], avs.TenantAdditionalValues[i+1:]...)
			avs.rebuildAllValues()
			break
		}
	}
}

// UpdateValueCount actualiza el conteo de uso de un valor
func (avs *AttributeValueSet) UpdateValueCount(value string, count int) {
	// Actualizar en todos los arrays
	for i := range avs.MarketplaceValues {
		if avs.MarketplaceValues[i].Value == value {
			avs.MarketplaceValues[i].Count = count
		}
	}

	for i := range avs.TenantAdditionalValues {
		if avs.TenantAdditionalValues[i].Value == value {
			avs.TenantAdditionalValues[i].Count = count
		}
	}

	for i := range avs.AllValues {
		if avs.AllValues[i].Value == value {
			avs.AllValues[i].Count = count
		}
	}
}

// GetValueBySlug retorna un valor por su slug
func (avs *AttributeValueSet) GetValueBySlug(slug string) (*AttributeValueItem, error) {
	for _, value := range avs.AllValues {
		if value.Slug == slug {
			return &value, nil
		}
	}
	return nil, errors.New("valor no encontrado con slug: " + slug)
}

// GetActiveValues retorna solo los valores activos
func (avs *AttributeValueSet) GetActiveValues() []AttributeValueItem {
	var activeValues []AttributeValueItem
	for _, value := range avs.AllValues {
		if value.IsActive {
			activeValues = append(activeValues, value)
		}
	}
	return activeValues
}

// GetValuesWithCount retorna valores que tienen conteo > 0 (hay productos)
func (avs *AttributeValueSet) GetValuesWithCount() []AttributeValueItem {
	var valuesWithCount []AttributeValueItem
	for _, value := range avs.AllValues {
		if value.IsActive && value.Count > 0 {
			valuesWithCount = append(valuesWithCount, value)
		}
	}
	return valuesWithCount
}

// GetMarketplaceOnlyValues retorna solo valores del marketplace
func (avs *AttributeValueSet) GetMarketplaceOnlyValues() []AttributeValueItem {
	return avs.MarketplaceValues
}

// GetTenantOnlyValues retorna solo valores adicionales del tenant
func (avs *AttributeValueSet) GetTenantOnlyValues() []AttributeValueItem {
	return avs.TenantAdditionalValues
}

// ValidateValue valida un valor contra las reglas del atributo
func (avs *AttributeValueSet) ValidateValue(value interface{}) error {
	if value == nil || value == "" {
		return errors.New("el valor no puede estar vacío")
	}

	valueStr := value.(string)

	// Para atributos tipo select, verificar que el valor existe
	if avs.AttributeType == AttributeTypeSelect || avs.AttributeType == AttributeTypeMultiSelect {
		found := false
		for _, availableValue := range avs.AllValues {
			if availableValue.Value == valueStr {
				found = true
				break
			}
		}
		if !found {
			return errors.New("el valor no está en la lista de valores permitidos")
		}
	}

	// Aplicar reglas de validación específicas
	return avs.applyValidationRules(valueStr)
}

// ValidateMultipleValues valida múltiples valores para atributos multi_select
func (avs *AttributeValueSet) ValidateMultipleValues(values []string) error {
	if avs.AttributeType != AttributeTypeMultiSelect {
		return errors.New("múltiples valores solo son permitidos para atributos multi_select")
	}

	for _, value := range values {
		if err := avs.ValidateValue(value); err != nil {
			return err
		}
	}

	return nil
}

// SortValues ordena todos los valores por SortOrder
func (avs *AttributeValueSet) SortValues() {
	sort.Slice(avs.MarketplaceValues, func(i, j int) bool {
		return avs.MarketplaceValues[i].SortOrder < avs.MarketplaceValues[j].SortOrder
	})

	sort.Slice(avs.TenantAdditionalValues, func(i, j int) bool {
		return avs.TenantAdditionalValues[i].SortOrder < avs.TenantAdditionalValues[j].SortOrder
	})

	avs.rebuildAllValues()
}

// GetFilterOptions retorna opciones para filtros con conteos
func (avs *AttributeValueSet) GetFilterOptions() []FilterOption {
	if !avs.IsFilterable {
		return []FilterOption{}
	}

	var options []FilterOption
	for _, value := range avs.GetValuesWithCount() {
		options = append(options, FilterOption{
			Label: value.Value,
			Value: value.Slug,
			Count: value.Count,
		})
	}

	// Ordenar por conteo descendente
	sort.Slice(options, func(i, j int) bool {
		return options[i].Count > options[j].Count
	})

	return options
}

// FilterOption representa una opción de filtro
type FilterOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Count int    `json:"count"`
}

// Métodos privados

// rebuildAllValues reconstruye el array de todos los valores combinados
func (avs *AttributeValueSet) rebuildAllValues() {
	avs.AllValues = []AttributeValueItem{}

	// Agregar valores del marketplace primero
	avs.AllValues = append(avs.AllValues, avs.MarketplaceValues...)

	// Agregar valores adicionales del tenant
	avs.AllValues = append(avs.AllValues, avs.TenantAdditionalValues...)

	// Ordenar por SortOrder
	sort.Slice(avs.AllValues, func(i, j int) bool {
		// Valores del marketplace van primero, luego del tenant
		if avs.AllValues[i].Source != avs.AllValues[j].Source {
			return avs.AllValues[i].Source == ValueSourceMarketplace
		}
		return avs.AllValues[i].SortOrder < avs.AllValues[j].SortOrder
	})
}

// applyValidationRules aplica las reglas de validación específicas
func (avs *AttributeValueSet) applyValidationRules(value string) error {
	// Verificar longitud mínima
	if minLength, exists := avs.ValidationRules["min_length"]; exists {
		if minLen, ok := minLength.(int); ok && len(value) < minLen {
			return errors.New("el valor es demasiado corto")
		}
	}

	// Verificar longitud máxima
	if maxLength, exists := avs.ValidationRules["max_length"]; exists {
		if maxLen, ok := maxLength.(int); ok && len(value) > maxLen {
			return errors.New("el valor es demasiado largo")
		}
	}

	// Verificar pattern regex
	if pattern, exists := avs.ValidationRules["pattern"]; exists {
		if patternStr, ok := pattern.(string); ok {
			// Aquí se podría implementar validación con regex
			_ = patternStr // Por ahora solo lo asignamos para evitar warning
		}
	}

	return nil
}
