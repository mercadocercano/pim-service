package entity

import (
	"errors"
	"time"
)

// ErrInvalidCategory representa errores relacionados con la validación de categorías
var ErrInvalidCategory = errors.New("categoría inválida")

// Category representa la entidad principal del dominio de categorías
// Test comment for git hook validation
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
