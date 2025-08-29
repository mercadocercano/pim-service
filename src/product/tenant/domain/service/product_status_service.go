package service

import (
	"errors"
	"fmt"

	"saas-mt-pim-service/src/product/tenant/domain/entity"
	"saas-mt-pim-service/src/product/tenant/domain/value_object"
)

// ProductStatusService maneja las transiciones de estado de productos
type ProductStatusService struct{}

// NewProductStatusService crea una nueva instancia del servicio
func NewProductStatusService() *ProductStatusService {
	return &ProductStatusService{}
}

// TransitionTo valida y ejecuta una transición de estado
func (s *ProductStatusService) TransitionTo(product *entity.Product, newStatus string) error {
	currentStatus := product.Status().Value()

	// Validar que la transición sea permitida
	if !s.isTransitionAllowed(currentStatus, newStatus) {
		return fmt.Errorf("transición no permitida: %s -> %s", currentStatus, newStatus)
	}

	// Ejecutar validaciones específicas del estado destino
	if err := s.validateStateRequirements(product, newStatus); err != nil {
		return fmt.Errorf("validación de estado fallida: %w", err)
	}

	// Ejecutar la transición usando los métodos existentes
	switch newStatus {
	case value_object.ProductStatusDraftValue:
		return product.SetToDraft()
	case value_object.ProductStatusPendingValue:
		return product.SetToPending()
	case value_object.ProductStatusActiveValue:
		return product.ActivateWithValidation()
	case value_object.ProductStatusInactiveValue:
		return product.DeactivateWithValidation()
	case value_object.ProductStatusDiscontinuedValue:
		product.Discontinue()
	case value_object.ProductStatusDeletedValue:
		product.Delete()
	default:
		return fmt.Errorf("estado no válido: %s", newStatus)
	}

	return nil
}

// isTransitionAllowed verifica si una transición de estado es permitida
func (s *ProductStatusService) isTransitionAllowed(from, to string) bool {
	// Matriz de transiciones permitidas
	allowedTransitions := map[string][]string{
		value_object.ProductStatusDraftValue: {
			value_object.ProductStatusPendingValue,
			value_object.ProductStatusActiveValue,
			value_object.ProductStatusDeletedValue,
		},
		value_object.ProductStatusPendingValue: {
			value_object.ProductStatusDraftValue,
			value_object.ProductStatusActiveValue,
			value_object.ProductStatusDeletedValue,
		},
		value_object.ProductStatusActiveValue: {
			value_object.ProductStatusInactiveValue,
			value_object.ProductStatusDiscontinuedValue,
			value_object.ProductStatusDeletedValue,
		},
		value_object.ProductStatusInactiveValue: {
			value_object.ProductStatusActiveValue,
			value_object.ProductStatusDiscontinuedValue,
			value_object.ProductStatusDeletedValue,
		},
		value_object.ProductStatusDiscontinuedValue: {
			value_object.ProductStatusDeletedValue,
		},
		// Los productos eliminados no pueden cambiar de estado
		value_object.ProductStatusDeletedValue: {},
	}

	allowedStates, exists := allowedTransitions[from]
	if !exists {
		return false
	}

	for _, allowedState := range allowedStates {
		if allowedState == to {
			return true
		}
	}

	return false
}

// validateStateRequirements valida los requisitos específicos para cada estado
func (s *ProductStatusService) validateStateRequirements(product *entity.Product, newStatus string) error {
	switch newStatus {
	case value_object.ProductStatusActiveValue:
		return s.validateActiveRequirements(product)
	case value_object.ProductStatusPendingValue:
		return s.validatePendingRequirements(product)
	case value_object.ProductStatusDraftValue:
		return s.validateDraftRequirements(product)
	default:
		// Otros estados no requieren validaciones especiales
		return nil
	}
}

// validateActiveRequirements valida que el producto puede activarse
func (s *ProductStatusService) validateActiveRequirements(product *entity.Product) error {
	var validationErrors []string

	// Validar que tenga nombre
	if product.Name() == "" {
		validationErrors = append(validationErrors, "el producto debe tener un nombre")
	}

	// Validar que tenga descripción
	if product.Description() == nil || *product.Description() == "" {
		validationErrors = append(validationErrors, "el producto debe tener una descripción")
	}

	// Validar que tenga categoría
	if product.CategoryReference() == nil {
		validationErrors = append(validationErrors, "el producto debe estar asignado a una categoría")
	}

	// Validar que tenga al menos una variante activa
	activeVariants := 0
	for _, variant := range product.GetVariants() {
		if variant.Status().IsActive() {
			activeVariants++
		}
	}
	if activeVariants == 0 {
		validationErrors = append(validationErrors, "el producto debe tener al menos una variante activa")
	}

	if len(validationErrors) > 0 {
		return errors.New("requisitos para activación no cumplidos: " + fmt.Sprintf("%v", validationErrors))
	}

	return nil
}

// validatePendingRequirements valida que el producto puede marcarse como pendiente
func (s *ProductStatusService) validatePendingRequirements(product *entity.Product) error {
	var validationErrors []string

	// Validar que tenga nombre
	if product.Name() == "" {
		validationErrors = append(validationErrors, "el producto debe tener un nombre")
	}

	// Validar que tenga descripción
	if product.Description() == nil || *product.Description() == "" {
		validationErrors = append(validationErrors, "el producto debe tener una descripción")
	}

	// Validar que tenga categoría
	if product.CategoryReference() == nil {
		validationErrors = append(validationErrors, "el producto debe estar asignado a una categoría")
	}

	if len(validationErrors) > 0 {
		return errors.New("requisitos para estado pendiente no cumplidos: " + fmt.Sprintf("%v", validationErrors))
	}

	return nil
}

// validateDraftRequirements valida que el producto puede marcarse como borrador
func (s *ProductStatusService) validateDraftRequirements(product *entity.Product) error {
	// Los productos draft tienen requisitos mínimos
	if product.Name() == "" {
		return errors.New("el producto debe tener al menos un nombre")
	}

	return nil
}

// GetAvailableTransitions retorna las transiciones disponibles para un producto
func (s *ProductStatusService) GetAvailableTransitions(product *entity.Product) []string {
	currentStatus := product.Status().Value()

	allowedTransitions := map[string][]string{
		value_object.ProductStatusDraftValue: {
			value_object.ProductStatusPendingValue,
			value_object.ProductStatusActiveValue,
			value_object.ProductStatusDeletedValue,
		},
		value_object.ProductStatusPendingValue: {
			value_object.ProductStatusDraftValue,
			value_object.ProductStatusActiveValue,
			value_object.ProductStatusDeletedValue,
		},
		value_object.ProductStatusActiveValue: {
			value_object.ProductStatusInactiveValue,
			value_object.ProductStatusDiscontinuedValue,
			value_object.ProductStatusDeletedValue,
		},
		value_object.ProductStatusInactiveValue: {
			value_object.ProductStatusActiveValue,
			value_object.ProductStatusDiscontinuedValue,
			value_object.ProductStatusDeletedValue,
		},
		value_object.ProductStatusDiscontinuedValue: {
			value_object.ProductStatusDeletedValue,
		},
		value_object.ProductStatusDeletedValue: {},
	}

	transitions, exists := allowedTransitions[currentStatus]
	if !exists {
		return []string{}
	}

	// Filtrar solo las transiciones que cumplen los requisitos
	var availableTransitions []string
	for _, transition := range transitions {
		if s.validateStateRequirements(product, transition) == nil {
			availableTransitions = append(availableTransitions, transition)
		}
	}

	return availableTransitions
}
