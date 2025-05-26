package value_object

import (
	"errors"

	"github.com/google/uuid"
)

// BrandReference representa una referencia a una marca
type BrandReference struct {
	id   string
	name string
}

// NewBrandReference crea una nueva referencia de marca
func NewBrandReference(id, name string) (*BrandReference, error) {
	if id == "" {
		return nil, errors.New("el ID de marca es obligatorio")
	}

	if name == "" {
		return nil, errors.New("el nombre de marca es obligatorio")
	}

	// Validar que el ID sea un UUID válido
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("el ID de marca debe ser un UUID válido")
	}

	return &BrandReference{
		id:   id,
		name: name,
	}, nil
}

// ID retorna el ID de la marca
func (br *BrandReference) ID() string {
	return br.id
}

// Name retorna el nombre de la marca
func (br *BrandReference) Name() string {
	return br.name
}

// IsEmpty verifica si la referencia está vacía
func (br *BrandReference) IsEmpty() bool {
	return br.id == "" || br.name == ""
}

// Equals compara dos referencias de marca
func (br *BrandReference) Equals(other *BrandReference) bool {
	if other == nil {
		return false
	}
	return br.id == other.id
}

// String retorna una representación en string
func (br *BrandReference) String() string {
	return br.name
}
