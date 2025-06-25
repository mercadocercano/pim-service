package value_object

import (
	"errors"

	"github.com/google/uuid"
)

// CategoryReference representa una referencia a una categoría
type CategoryReference struct {
	id   string
	name string
}

// NewCategoryReference crea una nueva referencia de categoría
func NewCategoryReference(id, name string) (*CategoryReference, error) {
	if id == "" {
		return nil, errors.New("el ID de categoría es obligatorio")
	}

	if name == "" {
		return nil, errors.New("el nombre de categoría es obligatorio")
	}

	// Validar que el ID sea un UUID válido
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("el ID de categoría debe ser un UUID válido")
	}

	return &CategoryReference{
		id:   id,
		name: name,
	}, nil
}

// ID retorna el ID de la categoría
func (cr *CategoryReference) ID() string {
	return cr.id
}

// Name retorna el nombre de la categoría
func (cr *CategoryReference) Name() string {
	return cr.name
}

// IsEmpty verifica si la referencia está vacía
func (cr *CategoryReference) IsEmpty() bool {
	return cr.id == "" || cr.name == ""
}

// Equals compara dos referencias de categoría
func (cr *CategoryReference) Equals(other *CategoryReference) bool {
	if other == nil {
		return false
	}
	return cr.id == other.id
}

// String retorna una representación en string
func (cr *CategoryReference) String() string {
	return cr.name
}
