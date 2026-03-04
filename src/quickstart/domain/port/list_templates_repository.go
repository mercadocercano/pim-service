package port

import (
	"context"
)

// ListTemplate representa un template de quickstart disponible
type ListTemplate struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	IsActive    bool     `json:"is_active"`
}

// ListTemplatesRepository define las operaciones para listar templates
type ListTemplatesRepository interface {
	LoadTemplatesFromBusinessTypeTemplates(ctx context.Context) ([]ListTemplate, error)
}
