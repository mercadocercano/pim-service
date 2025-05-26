package value_object

import (
	"errors"
	"regexp"
	"strings"
)

// ProductSKU representa el código único de un producto
type ProductSKU struct {
	value string
}

// NewProductSKU crea un nuevo ProductSKU
func NewProductSKU(value string) (*ProductSKU, error) {
	if value == "" {
		return nil, errors.New("el SKU no puede estar vacío")
	}

	// Normalizar: convertir a mayúsculas y quitar espacios
	normalizedValue := strings.ToUpper(strings.TrimSpace(value))

	if !isValidSKU(normalizedValue) {
		return nil, errors.New("formato de SKU inválido")
	}

	return &ProductSKU{value: normalizedValue}, nil
}

// Value retorna el valor del SKU
func (ps *ProductSKU) Value() string {
	return ps.value
}

// String implementa la interfaz Stringer
func (ps *ProductSKU) String() string {
	return ps.value
}

// Equals compara dos SKUs
func (ps *ProductSKU) Equals(other *ProductSKU) bool {
	if other == nil {
		return false
	}
	return ps.value == other.value
}

// IsEmpty verifica si el SKU está vacío
func (ps *ProductSKU) IsEmpty() bool {
	return ps.value == ""
}

// isValidSKU valida el formato del SKU
// Permite letras, números, guiones y guiones bajos
// Longitud entre 3 y 50 caracteres
func isValidSKU(value string) bool {
	if len(value) < 3 || len(value) > 50 {
		return false
	}

	// Patrón: solo letras, números, guiones y guiones bajos
	pattern := `^[A-Z0-9_-]+$`
	matched, _ := regexp.MatchString(pattern, value)
	return matched
}
