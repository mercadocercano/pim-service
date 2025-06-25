package value_object

import "errors"

// VariantStatus representa el estado de una variante de producto
type VariantStatus struct {
	value string
}

// Constantes para los estados de variante
const (
	VariantStatusActive       = "active"
	VariantStatusInactive     = "inactive"
	VariantStatusDiscontinued = "discontinued"
	VariantStatusDeleted      = "deleted"
)

// NewVariantStatus crea un nuevo VariantStatus
func NewVariantStatus(value string) (*VariantStatus, error) {
	if !isValidVariantStatus(value) {
		return nil, errors.New("estado de variante inválido")
	}
	return &VariantStatus{value: value}, nil
}

// NewActiveVariantStatus retorna un estado activo
func NewActiveVariantStatus() *VariantStatus {
	return &VariantStatus{value: VariantStatusActive}
}

// NewInactiveVariantStatus retorna un estado inactivo
func NewInactiveVariantStatus() *VariantStatus {
	return &VariantStatus{value: VariantStatusInactive}
}

// NewDiscontinuedVariantStatus retorna un estado descontinuado
func NewDiscontinuedVariantStatus() *VariantStatus {
	return &VariantStatus{value: VariantStatusDiscontinued}
}

// NewDeletedVariantStatus retorna un estado eliminado
func NewDeletedVariantStatus() *VariantStatus {
	return &VariantStatus{value: VariantStatusDeleted}
}

// isValidVariantStatus valida si el estado es válido
func isValidVariantStatus(value string) bool {
	switch value {
	case VariantStatusActive, VariantStatusInactive, VariantStatusDiscontinued, VariantStatusDeleted:
		return true
	default:
		return false
	}
}

// Value retorna el valor del estado
func (vs *VariantStatus) Value() string {
	return vs.value
}

// IsActive verifica si el estado es activo
func (vs *VariantStatus) IsActive() bool {
	return vs.value == VariantStatusActive
}

// IsInactive verifica si el estado es inactivo
func (vs *VariantStatus) IsInactive() bool {
	return vs.value == VariantStatusInactive
}

// IsDiscontinued verifica si el estado es descontinuado
func (vs *VariantStatus) IsDiscontinued() bool {
	return vs.value == VariantStatusDiscontinued
}

// IsDeleted verifica si el estado es eliminado
func (vs *VariantStatus) IsDeleted() bool {
	return vs.value == VariantStatusDeleted
}

// CanBeActivated verifica si puede ser activado desde el estado actual
func (vs *VariantStatus) CanBeActivated() bool {
	switch vs.value {
	case VariantStatusInactive, VariantStatusDiscontinued:
		return true
	case VariantStatusDeleted, VariantStatusActive:
		return false
	default:
		return false
	}
}

// CanBeDeactivated verifica si puede ser desactivado desde el estado actual
func (vs *VariantStatus) CanBeDeactivated() bool {
	switch vs.value {
	case VariantStatusActive:
		return true
	case VariantStatusInactive, VariantStatusDiscontinued, VariantStatusDeleted:
		return false
	default:
		return false
	}
}

// CanBeDiscontinued verifica si puede ser descontinuado desde el estado actual
func (vs *VariantStatus) CanBeDiscontinued() bool {
	switch vs.value {
	case VariantStatusActive, VariantStatusInactive:
		return true
	case VariantStatusDiscontinued, VariantStatusDeleted:
		return false
	default:
		return false
	}
}

// CanBeDeleted verifica si puede ser eliminado desde el estado actual
func (vs *VariantStatus) CanBeDeleted() bool {
	switch vs.value {
	case VariantStatusActive, VariantStatusInactive, VariantStatusDiscontinued:
		return true
	case VariantStatusDeleted:
		return false
	default:
		return false
	}
}

// Equals compara dos estados de variante
func (vs *VariantStatus) Equals(other *VariantStatus) bool {
	if other == nil {
		return false
	}
	return vs.value == other.value
}

// String retorna la representación en string del estado
func (vs *VariantStatus) String() string {
	return vs.value
}
