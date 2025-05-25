package value_object

import "errors"

// BrandReference representa una referencia a una marca para usar en otros contextos
type BrandReference struct {
	ID          string
	Name        string
	Description string
}

// NewBrandReference crea una nueva referencia de marca con validaciones
func NewBrandReference(id, name, description string) (*BrandReference, error) {
	if id == "" {
		return nil, errors.New("el ID de la marca es obligatorio")
	}

	if name == "" {
		return nil, errors.New("el nombre de la marca es obligatorio")
	}

	return &BrandReference{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

// IsEmpty verifica si la referencia está vacía
func (br *BrandReference) IsEmpty() bool {
	return br == nil || br.ID == ""
}

// Equals compara dos referencias de marca
func (br *BrandReference) Equals(other *BrandReference) bool {
	if br == nil && other == nil {
		return true
	}
	if br == nil || other == nil {
		return false
	}
	return br.ID == other.ID
}
