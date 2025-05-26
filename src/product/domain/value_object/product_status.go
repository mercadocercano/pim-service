package value_object

import "errors"

// ProductStatus representa el estado de un producto
type ProductStatus struct {
	value string
}

// Constantes para los estados de producto
const (
	ProductStatusActiveValue       = "active"
	ProductStatusInactiveValue     = "inactive"
	ProductStatusDiscontinuedValue = "discontinued"
	ProductStatusDeletedValue      = "deleted"
)

// NewProductStatus crea un nuevo ProductStatus
func NewProductStatus(value string) (ProductStatus, error) {
	if !isValidProductStatus(value) {
		return ProductStatus{}, errors.New("estado de producto inválido")
	}
	return ProductStatus{value: value}, nil
}

// ProductStatusActive retorna un estado activo
func ProductStatusActive() ProductStatus {
	return ProductStatus{value: ProductStatusActiveValue}
}

// ProductStatusInactive retorna un estado inactivo
func ProductStatusInactive() ProductStatus {
	return ProductStatus{value: ProductStatusInactiveValue}
}

// ProductStatusDiscontinued retorna un estado descontinuado
func ProductStatusDiscontinued() ProductStatus {
	return ProductStatus{value: ProductStatusDiscontinuedValue}
}

// ProductStatusDeleted retorna un estado eliminado
func ProductStatusDeleted() ProductStatus {
	return ProductStatus{value: ProductStatusDeletedValue}
}

// Value retorna el valor del estado
func (ps ProductStatus) Value() string {
	return ps.value
}

// IsActive verifica si el estado es activo
func (ps ProductStatus) IsActive() bool {
	return ps.value == ProductStatusActiveValue
}

// IsInactive verifica si el estado es inactivo
func (ps ProductStatus) IsInactive() bool {
	return ps.value == ProductStatusInactiveValue
}

// IsDiscontinued verifica si el estado es descontinuado
func (ps ProductStatus) IsDiscontinued() bool {
	return ps.value == ProductStatusDiscontinuedValue
}

// IsDeleted verifica si el estado es eliminado
func (ps ProductStatus) IsDeleted() bool {
	return ps.value == ProductStatusDeletedValue
}

// IsValid verifica si el estado es válido
func (ps ProductStatus) IsValid() bool {
	return isValidProductStatus(ps.value)
}

// isValidProductStatus verifica si un valor es un estado válido
func isValidProductStatus(value string) bool {
	validStatuses := []string{
		ProductStatusActiveValue,
		ProductStatusInactiveValue,
		ProductStatusDiscontinuedValue,
		ProductStatusDeletedValue,
	}

	for _, status := range validStatuses {
		if value == status {
			return true
		}
	}
	return false
}
