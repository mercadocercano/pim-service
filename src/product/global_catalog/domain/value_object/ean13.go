package value_object

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// EAN13 representa un código de barras EAN-13 válido
type EAN13 struct {
	value string
}

// NewEAN13 crea un nuevo EAN13 validando el formato y el checksum
func NewEAN13(value string) (*EAN13, error) {
	if value == "" {
		return nil, errors.New("EAN-13 es obligatorio")
	}

	// Validar formato (13 dígitos)
	if !isValidEAN13Format(value) {
		return nil, errors.New("el código EAN-13 debe contener exactamente 13 dígitos")
	}

	// Validar checksum
	if !isValidEAN13Checksum(value) {
		return nil, errors.New("el código EAN-13 tiene un checksum inválido")
	}

	return &EAN13{value: value}, nil
}

// Value retorna el valor del EAN-13
func (e *EAN13) Value() string {
	return e.value
}

// String implementa la interfaz Stringer
func (e *EAN13) String() string {
	return e.value
}

// Formatted retorna el EAN-13 con formato legible (xxx-x-xxxxx-xxxxx-x)
func (e *EAN13) Formatted() string {
	if len(e.value) != 13 {
		return e.value
	}
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		e.value[0:3],   // País
		e.value[3:4],   // Grupo
		e.value[4:9],   // Fabricante
		e.value[9:12],  // Producto
		e.value[12:13], // Checksum
	)
}

// CountryCode retorna el código de país del EAN-13
func (e *EAN13) CountryCode() string {
	if len(e.value) >= 3 {
		return e.value[0:3]
	}
	return ""
}

// IsArgentineProduct verifica si el producto es de origen argentino
func (e *EAN13) IsArgentineProduct() bool {
	countryCode := e.CountryCode()
	// Códigos de Argentina: 779, 780-799
	return countryCode == "779" || (countryCode >= "780" && countryCode <= "799")
}

// isValidEAN13Format valida que el string contenga exactamente 13 dígitos
func isValidEAN13Format(value string) bool {
	matched, _ := regexp.MatchString(`^\d{13}$`, value)
	return matched
}

// isValidEAN13Checksum valida el checksum del EAN-13 usando el algoritmo estándar
func isValidEAN13Checksum(value string) bool {
	if len(value) != 13 {
		return false
	}

	sum := 0
	for i, char := range value[:12] {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}

		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}

	checksum := (10 - (sum % 10)) % 10
	expectedChecksum, err := strconv.Atoi(string(value[12]))
	if err != nil {
		return false
	}

	return checksum == expectedChecksum
}

// Equals compara dos EAN13
func (e *EAN13) Equals(other *EAN13) bool {
	if other == nil {
		return false
	}
	return e.value == other.value
}
