package port

import (
	"context"
)

// ListTemplate representa un template de quickstart disponible
type ListTemplateBrand struct {
	Name    string `json:"name"`
	LogoURL string `json:"logo_url,omitempty"`
}

type ListTemplate struct {
	ID              string              `json:"id"`
	Name            string              `json:"name"`
	Slug            string              `json:"slug"`
	Description     string              `json:"description"`
	Icon            string              `json:"icon"`
	Categories      []string            `json:"categories"`
	Brands          []ListTemplateBrand  `json:"brands,omitempty"`
	TotalCategories int                 `json:"total_categories"`
	TotalProducts   int                 `json:"total_products"`
	IsActive        bool                `json:"is_active"`
}

// ListTemplatesRepository define las operaciones para listar templates
type ListTemplatesRepository interface {
	LoadTemplatesFromBusinessTypeTemplates(ctx context.Context) ([]ListTemplate, error)
}
