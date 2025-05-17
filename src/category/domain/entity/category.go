package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// ErrInvalidCategory representa errores relacionados con la validación de categorías
var ErrInvalidCategory = errors.New("categoría inválida")

// Category representa la entidad principal del dominio de categorías
type Category struct {
	ID          string
	TenantID    string
	Name        string
	Description string
	ParentID    *string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewCategory crea una nueva instancia de Category con validaciones
func NewCategory(tenantID, name, description string, parentID *string) (*Category, error) {
	if name == "" {
		return nil, errors.New("el nombre de la categoría es obligatorio")
	}

	if tenantID == "" {
		return nil, errors.New("el tenant ID es obligatorio")
	}

	now := time.Now()
	return &Category{
		ID:          uuid.New().String(),
		TenantID:    tenantID,
		Name:        name,
		Description: description,
		ParentID:    parentID,
		Status:      "active",
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// Update actualiza los campos de la categoría
func (c *Category) Update(name, description string, parentID *string) error {
	if name == "" {
		return errors.New("el nombre de la categoría es obligatorio")
	}

	c.Name = name
	c.Description = description
	c.ParentID = parentID
	c.UpdatedAt = time.Now()
	return nil
}

// Activate establece la categoría como activa
func (c *Category) Activate() {
	c.Status = "active"
	c.UpdatedAt = time.Now()
}

// Deactivate establece la categoría como inactiva
func (c *Category) Deactivate() {
	c.Status = "inactive"
	c.UpdatedAt = time.Now()
}

// IsActive verifica si la categoría está activa
func (c *Category) IsActive() bool {
	return c.Status == "active"
}

// HasParent verifica si la categoría tiene un padre
func (c *Category) HasParent() bool {
	return c.ParentID != nil
}
