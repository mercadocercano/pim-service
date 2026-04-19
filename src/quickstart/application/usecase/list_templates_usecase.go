package usecase

import (
	"context"

	"saas-mt-pim-service/src/quickstart/domain/port"
)

// Template representa un template de quickstart disponible
type TemplateBrand struct {
	Name    string `json:"name"`
	LogoURL string `json:"logo_url,omitempty"`
}

type Template struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Slug            string          `json:"slug"`
	Description     string          `json:"description"`
	Icon            string          `json:"icon"`
	Categories      []string        `json:"categories"`
	Brands          []TemplateBrand `json:"brands,omitempty"`
	TotalCategories int             `json:"total_categories"`
	TotalProducts   int             `json:"total_products"`
	IsActive        bool            `json:"is_active"`
}

// ListTemplatesResponse es la respuesta del caso de uso
type ListTemplatesResponse struct {
	Templates []Template `json:"templates"`
	Total     int        `json:"total"`
}

// ListTemplatesUseCase obtiene la lista de templates disponibles
type ListTemplatesUseCase struct {
	repo port.ListTemplatesRepository
}

// NewListTemplatesUseCase crea una nueva instancia del caso de uso
func NewListTemplatesUseCase(repo port.ListTemplatesRepository) *ListTemplatesUseCase {
	return &ListTemplatesUseCase{
		repo: repo,
	}
}

// Execute lista templates desde business_type_templates en la DB.
// Si la DB está vacía (seeds no ejecutados), retorna lista vacía y el frontend
// muestra su propio fallback estático.
func (uc *ListTemplatesUseCase) Execute(ctx context.Context) (*ListTemplatesResponse, error) {
	templates, err := uc.repo.LoadTemplatesFromBusinessTypeTemplates(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]Template, len(templates))
	for i, t := range templates {
		var brands []TemplateBrand
		for _, b := range t.Brands {
			brands = append(brands, TemplateBrand{Name: b.Name, LogoURL: b.LogoURL})
		}

		result[i] = Template{
			ID:              t.ID,
			Name:            t.Name,
			Slug:            t.Slug,
			Description:     t.Description,
			Icon:            t.Icon,
			Categories:      t.Categories,
			Brands:          brands,
			TotalCategories: t.TotalCategories,
			TotalProducts:   t.TotalProducts,
			IsActive:        t.IsActive,
		}
	}

	return &ListTemplatesResponse{
		Templates: result,
		Total:     len(result),
	}, nil
}
