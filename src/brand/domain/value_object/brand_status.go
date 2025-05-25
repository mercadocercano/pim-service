package value_object

import "errors"

// BrandStatus representa los estados válidos de una marca
type BrandStatus string

const (
	BrandStatusActive   BrandStatus = "active"
	BrandStatusInactive BrandStatus = "inactive"
	BrandStatusDeleted  BrandStatus = "deleted"
)

// ValidBrandStatuses contiene todos los estados válidos
var ValidBrandStatuses = []BrandStatus{
	BrandStatusActive,
	BrandStatusInactive,
	BrandStatusDeleted,
}

// NewBrandStatus crea un nuevo BrandStatus validando el valor
func NewBrandStatus(status string) (BrandStatus, error) {
	brandStatus := BrandStatus(status)
	if !brandStatus.IsValid() {
		return "", errors.New("estado de marca inválido")
	}
	return brandStatus, nil
}

// IsValid verifica si el estado es válido
func (bs BrandStatus) IsValid() bool {
	for _, validStatus := range ValidBrandStatuses {
		if bs == validStatus {
			return true
		}
	}
	return false
}

// String devuelve la representación en string del estado
func (bs BrandStatus) String() string {
	return string(bs)
}

// IsActive verifica si la marca está activa
func (bs BrandStatus) IsActive() bool {
	return bs == BrandStatusActive
}

// IsDeleted verifica si la marca está eliminada
func (bs BrandStatus) IsDeleted() bool {
	return bs == BrandStatusDeleted
}
