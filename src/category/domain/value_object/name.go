package value_object

import (
	"errors"
	"strings"
)

// ErrEmptyName error cuando el nombre está vacío
var ErrEmptyName = errors.New("el nombre de la categoría es obligatorio")

// ErrNameTooLong error cuando el nombre es demasiado largo
var ErrNameTooLong = errors.New("el nombre de la categoría es demasiado largo")

// Name representa el nombre de una categoría como objeto de valor
type Name struct {
	value string
}

// NewName crea un nuevo objeto de valor Name con validaciones
func NewName(value string) (*Name, error) {
	// Eliminar espacios en blanco al inicio y final
	value = strings.TrimSpace(value)

	if value == "" {
		return nil, ErrEmptyName
	}

	if len(value) > 100 {
		return nil, ErrNameTooLong
	}

	return &Name{value: value}, nil
}

// Value devuelve el valor del nombre
func (n *Name) Value() string {
	return n.value
}

// Equals compara dos nombres
func (n *Name) Equals(other *Name) bool {
	if other == nil {
		return false
	}
	return n.value == other.value
}
