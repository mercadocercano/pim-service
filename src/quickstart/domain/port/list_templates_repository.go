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
	Brands          []ListTemplateBrand `json:"brands,omitempty"`
	TotalCategories int                 `json:"total_categories"`
	TotalProducts   int                 `json:"total_products"`
	IsActive        bool                `json:"is_active"`
}

// ListTemplatesRepository define las operaciones para listar templates
type ListTemplatesRepository interface {
	// LoadTemplatesFromBusinessTypeTemplates carga el surtido editorial (JSONB
	// curado de business_type_templates: brands + products).
	LoadTemplatesFromBusinessTypeTemplates(ctx context.Context) ([]ListTemplate, error)

	// LoadTemplatesComputed carga el surtido computado desde global_products
	// (business_type_product_templates: priority_brands + suggested_products).
	// Cae a editorial por-template cuando un template no tiene surtido computado.
	LoadTemplatesComputed(ctx context.Context) ([]ListTemplate, error)
}
