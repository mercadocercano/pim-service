package value_object

import (
	"errors"
	"strings"
)

// VariantAttribute representa un atributo de una variante de producto
type VariantAttribute struct {
	name  string
	value string
}

// NewVariantAttribute crea un nuevo atributo de variante
func NewVariantAttribute(name, value string) (*VariantAttribute, error) {
	if err := validateAttributeName(name); err != nil {
		return nil, err
	}

	if err := validateAttributeValue(value); err != nil {
		return nil, err
	}

	return &VariantAttribute{
		name:  strings.TrimSpace(name),
		value: strings.TrimSpace(value),
	}, nil
}

// validateAttributeName valida el nombre del atributo
func validateAttributeName(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return errors.New("el nombre del atributo es requerido")
	}

	if len(name) < 2 {
		return errors.New("el nombre del atributo debe tener al menos 2 caracteres")
	}

	if len(name) > 100 {
		return errors.New("el nombre del atributo no puede exceder 100 caracteres")
	}

	return nil
}

// validateAttributeValue valida el valor del atributo
func validateAttributeValue(value string) error {
	value = strings.TrimSpace(value)

	if value == "" {
		return errors.New("el valor del atributo es requerido")
	}

	if len(value) > 255 {
		return errors.New("el valor del atributo no puede exceder 255 caracteres")
	}

	return nil
}

// Name retorna el nombre del atributo
func (va *VariantAttribute) Name() string {
	return va.name
}

// Value retorna el valor del atributo
func (va *VariantAttribute) Value() string {
	return va.value
}

// UpdateValue actualiza el valor del atributo
func (va *VariantAttribute) UpdateValue(newValue string) error {
	if err := validateAttributeValue(newValue); err != nil {
		return err
	}

	va.value = strings.TrimSpace(newValue)
	return nil
}

// Equals compara dos atributos de variante
func (va *VariantAttribute) Equals(other *VariantAttribute) bool {
	if other == nil {
		return false
	}
	return va.name == other.name && va.value == other.value
}

// HasSameName verifica si dos atributos tienen el mismo nombre
func (va *VariantAttribute) HasSameName(other *VariantAttribute) bool {
	if other == nil {
		return false
	}
	return va.name == other.name
}

// String retorna la representación en string del atributo
func (va *VariantAttribute) String() string {
	return va.name + ": " + va.value
}

// VariantAttributeCollection representa una colección de atributos de variante
type VariantAttributeCollection struct {
	attributes []*VariantAttribute
}

// NewVariantAttributeCollection crea una nueva colección de atributos
func NewVariantAttributeCollection(attributes []*VariantAttribute) (*VariantAttributeCollection, error) {
	if err := validateAttributeCollection(attributes); err != nil {
		return nil, err
	}

	return &VariantAttributeCollection{
		attributes: attributes,
	}, nil
}

// validateAttributeCollection valida que no haya nombres duplicados
func validateAttributeCollection(attributes []*VariantAttribute) error {
	nameMap := make(map[string]bool)

	for _, attr := range attributes {
		if attr == nil {
			return errors.New("atributo nulo encontrado en la colección")
		}

		if nameMap[attr.Name()] {
			return errors.New("nombres de atributos duplicados: " + attr.Name())
		}

		nameMap[attr.Name()] = true
	}

	return nil
}

// Attributes retorna todos los atributos
func (vac *VariantAttributeCollection) Attributes() []*VariantAttribute {
	return vac.attributes
}

// Count retorna el número de atributos
func (vac *VariantAttributeCollection) Count() int {
	return len(vac.attributes)
}

// IsEmpty verifica si la colección está vacía
func (vac *VariantAttributeCollection) IsEmpty() bool {
	return len(vac.attributes) == 0
}

// GetByName obtiene un atributo por nombre
func (vac *VariantAttributeCollection) GetByName(name string) *VariantAttribute {
	for _, attr := range vac.attributes {
		if attr.Name() == name {
			return attr
		}
	}
	return nil
}

// HasAttribute verifica si existe un atributo con el nombre dado
func (vac *VariantAttributeCollection) HasAttribute(name string) bool {
	return vac.GetByName(name) != nil
}

// AddAttribute agrega un nuevo atributo
func (vac *VariantAttributeCollection) AddAttribute(attribute *VariantAttribute) error {
	if attribute == nil {
		return errors.New("no se puede agregar un atributo nulo")
	}

	if vac.HasAttribute(attribute.Name()) {
		return errors.New("ya existe un atributo con el nombre: " + attribute.Name())
	}

	vac.attributes = append(vac.attributes, attribute)
	return nil
}

// UpdateAttribute actualiza un atributo existente
func (vac *VariantAttributeCollection) UpdateAttribute(name, newValue string) error {
	attr := vac.GetByName(name)
	if attr == nil {
		return errors.New("atributo no encontrado: " + name)
	}

	return attr.UpdateValue(newValue)
}

// RemoveAttribute elimina un atributo por nombre
func (vac *VariantAttributeCollection) RemoveAttribute(name string) error {
	for i, attr := range vac.attributes {
		if attr.Name() == name {
			// Eliminar el elemento del slice
			vac.attributes = append(vac.attributes[:i], vac.attributes[i+1:]...)
			return nil
		}
	}
	return errors.New("atributo no encontrado: " + name)
}

// ReplaceAll reemplaza todos los atributos con una nueva colección
func (vac *VariantAttributeCollection) ReplaceAll(newAttributes []*VariantAttribute) error {
	if err := validateAttributeCollection(newAttributes); err != nil {
		return err
	}

	vac.attributes = newAttributes
	return nil
}

// ToMap convierte la colección a un mapa nombre->valor
func (vac *VariantAttributeCollection) ToMap() map[string]string {
	result := make(map[string]string)
	for _, attr := range vac.attributes {
		result[attr.Name()] = attr.Value()
	}
	return result
}

// Equals compara dos colecciones de atributos
func (vac *VariantAttributeCollection) Equals(other *VariantAttributeCollection) bool {
	if other == nil {
		return false
	}

	if vac.Count() != other.Count() {
		return false
	}

	for _, attr := range vac.attributes {
		otherAttr := other.GetByName(attr.Name())
		if otherAttr == nil || !attr.Equals(otherAttr) {
			return false
		}
	}

	return true
}
