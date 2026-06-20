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
	repo        port.ListTemplatesRepository
	useComputed bool
}

// NewListTemplatesUseCase crea una nueva instancia del caso de uso.
// useComputed activa el read-path COMPUTADO (surtido derivado de global_products,
// ADR-007 Fase 2). Con fallback automático al editorial ante error.
func NewListTemplatesUseCase(repo port.ListTemplatesRepository, useComputed bool) *ListTemplatesUseCase {
	return &ListTemplatesUseCase{
		repo:        repo,
		useComputed: useComputed,
	}
}

// Execute lista templates desde business_type_templates en la DB.
// Si la DB está vacía (seeds no ejecutados), retorna lista vacía y el frontend
// muestra su propio fallback estático.
func (uc *ListTemplatesUseCase) Execute(ctx context.Context) (*ListTemplatesResponse, error) {
	templates, err := uc.loadTemplates(ctx)
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

// loadTemplates elige el read-path: computado (ADR-007 Fase 2) cuando el flag
// está activo, con fallback al editorial si el computado falla; editorial por
// defecto.
func (uc *ListTemplatesUseCase) loadTemplates(ctx context.Context) ([]port.ListTemplate, error) {
	if uc.useComputed {
		if templates, err := uc.repo.LoadTemplatesComputed(ctx); err == nil {
			return templates, nil
		}
		// Fallback: si el computado falla, servir editorial en vez de romper el onboarding.
	}
	return uc.repo.LoadTemplatesFromBusinessTypeTemplates(ctx)
}
