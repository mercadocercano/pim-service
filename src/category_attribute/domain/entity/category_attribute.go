package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// ErrInvalidCategoryAttribute representa errores relacionados con la validación de atributos de categoría
var ErrInvalidCategoryAttribute = errors.New("atributo de categoría inválido")

// CategoryAttribute representa la entidad que relaciona categorías con atributos y sus valores permitidos
type CategoryAttribute struct {
	ID            string
	TenantID      string
	CategoryID    string
	AttributeID   string
	AllowedValues []string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// NewCategoryAttribute crea una nueva instancia de CategoryAttribute con validaciones
func NewCategoryAttribute(tenantID, categoryID, attributeID string, allowedValues []string) (*CategoryAttribute, error) {
	if tenantID == "" {
		return nil, errors.New("el tenant ID es obligatorio")
	}

	if categoryID == "" {
		return nil, errors.New("el category ID es obligatorio")
	}

	if attributeID == "" {
		return nil, errors.New("el attribute ID es obligatorio")
	}

	now := time.Now()
	return &CategoryAttribute{
		ID:            uuid.New().String(),
		TenantID:      tenantID,
		CategoryID:    categoryID,
		AttributeID:   attributeID,
		AllowedValues: allowedValues,
		Status:        "active",
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}

// UpdateAllowedValues actualiza los valores permitidos para el atributo
func (ca *CategoryAttribute) UpdateAllowedValues(allowedValues []string) {
	ca.AllowedValues = allowedValues
	ca.UpdatedAt = time.Now()
}

// AddAllowedValue agrega un nuevo valor permitido
func (ca *CategoryAttribute) AddAllowedValue(value string) error {
	if value == "" {
		return errors.New("el valor no puede estar vacío")
	}

	// Verificar si el valor ya existe
	for _, existing := range ca.AllowedValues {
		if existing == value {
			return errors.New("el valor ya existe en la lista de valores permitidos")
		}
	}

	ca.AllowedValues = append(ca.AllowedValues, value)
	ca.UpdatedAt = time.Now()
	return nil
}

// RemoveAllowedValue elimina un valor permitido
func (ca *CategoryAttribute) RemoveAllowedValue(value string) error {
	for i, existing := range ca.AllowedValues {
		if existing == value {
			ca.AllowedValues = append(ca.AllowedValues[:i], ca.AllowedValues[i+1:]...)
			ca.UpdatedAt = time.Now()
			return nil
		}
	}
	return errors.New("el valor no existe en la lista de valores permitidos")
}

// Activate establece el atributo de categoría como activo
func (ca *CategoryAttribute) Activate() {
	ca.Status = "active"
	ca.UpdatedAt = time.Now()
}

// Deactivate establece el atributo de categoría como inactivo
func (ca *CategoryAttribute) Deactivate() {
	ca.Status = "inactive"
	ca.UpdatedAt = time.Now()
}

// IsActive verifica si el atributo de categoría está activo
func (ca *CategoryAttribute) IsActive() bool {
	return ca.Status == "active"
}

// HasAllowedValues verifica si el atributo tiene valores permitidos
func (ca *CategoryAttribute) HasAllowedValues() bool {
	return len(ca.AllowedValues) > 0
}
