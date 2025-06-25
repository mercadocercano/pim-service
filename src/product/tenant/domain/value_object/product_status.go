package value_object

import "errors"

// ProductStatus representa el estado de un producto
type ProductStatus struct {
	value string
}

// Constantes para los estados de producto
const (
	ProductStatusDraftValue        = "draft"        // Importado desde template, pendiente de configuración
	ProductStatusPendingValue      = "pending"      // Configurado pero pendiente de precios/stock
	ProductStatusActiveValue       = "active"       // Listo para vender
	ProductStatusInactiveValue     = "inactive"     // Temporalmente deshabilitado
	ProductStatusDiscontinuedValue = "discontinued" // Descontinuado
	ProductStatusDeletedValue      = "deleted"      // Eliminado (soft delete)
)

// NewProductStatus crea un nuevo ProductStatus
func NewProductStatus(value string) (ProductStatus, error) {
	if !isValidProductStatus(value) {
		return ProductStatus{}, errors.New("estado de producto inválido")
	}
	return ProductStatus{value: value}, nil
}

// ProductStatusDraft retorna un estado borrador
func ProductStatusDraft() ProductStatus {
	return ProductStatus{value: ProductStatusDraftValue}
}

// ProductStatusPending retorna un estado pendiente
func ProductStatusPending() ProductStatus {
	return ProductStatus{value: ProductStatusPendingValue}
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

// IsDraft verifica si el estado es borrador
func (ps ProductStatus) IsDraft() bool {
	return ps.value == ProductStatusDraftValue
}

// IsPending verifica si el estado es pendiente
func (ps ProductStatus) IsPending() bool {
	return ps.value == ProductStatusPendingValue
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

// CanBeActivated verifica si el producto puede ser activado
func (ps ProductStatus) CanBeActivated() bool {
	return ps.value == ProductStatusDraftValue ||
		ps.value == ProductStatusPendingValue ||
		ps.value == ProductStatusInactiveValue
}

// CanBeEdited verifica si el producto puede ser editado
func (ps ProductStatus) CanBeEdited() bool {
	return ps.value != ProductStatusDeletedValue
}

// IsReadyForMarketplace verifica si está listo para aparecer en el marketplace
func (ps ProductStatus) IsReadyForMarketplace() bool {
	return ps.value == ProductStatusActiveValue
}

// RequiresConfiguration verifica si necesita configuración adicional
func (ps ProductStatus) RequiresConfiguration() bool {
	return ps.value == ProductStatusDraftValue || ps.value == ProductStatusPendingValue
}

// IsValid verifica si el estado es válido
func (ps ProductStatus) IsValid() bool {
	return isValidProductStatus(ps.value)
}

// GetNextValidStates retorna los estados válidos a los que puede transicionar
func (ps ProductStatus) GetNextValidStates() []string {
	switch ps.value {
	case ProductStatusDraftValue:
		return []string{ProductStatusPendingValue, ProductStatusInactiveValue, ProductStatusDeletedValue}
	case ProductStatusPendingValue:
		return []string{ProductStatusActiveValue, ProductStatusInactiveValue, ProductStatusDeletedValue}
	case ProductStatusActiveValue:
		return []string{ProductStatusInactiveValue, ProductStatusDiscontinuedValue}
	case ProductStatusInactiveValue:
		return []string{ProductStatusActiveValue, ProductStatusDiscontinuedValue, ProductStatusDeletedValue}
	case ProductStatusDiscontinuedValue:
		return []string{ProductStatusDeletedValue}
	case ProductStatusDeletedValue:
		return []string{} // No puede transicionar desde eliminado
	default:
		return []string{}
	}
}

// CanTransitionTo verifica si puede transicionar a un estado específico
func (ps ProductStatus) CanTransitionTo(newStatus string) bool {
	validStates := ps.GetNextValidStates()
	for _, state := range validStates {
		if state == newStatus {
			return true
		}
	}
	return false
}

// isValidProductStatus verifica si un valor es un estado válido
func isValidProductStatus(value string) bool {
	validStatuses := []string{
		ProductStatusDraftValue,
		ProductStatusPendingValue,
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
