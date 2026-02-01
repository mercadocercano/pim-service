package model

import "time"

// Category representa el modelo de base de datos para una categoría
type Category struct {
	ID          string
	TenantID    string
	Name        string
	Slug        string  // HITO 2: Para búsqueda por slug
	Description string
	ParentID    *string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
