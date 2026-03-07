package usecase

import (
	"context"
	"strings"

	"saas-mt-pim-service/src/quickstart/domain/port"
)

// Template representa un template de quickstart disponible
type Template struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	IsActive    bool     `json:"is_active"`
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

// legacyTemplates: templates estáticos que siempre se muestran si no vienen del DB
var legacyTemplates = []port.ListTemplate{
	{ID: "ferreteria-corralon", Name: "Ferretería / Corralón", Slug: "ferreteria-corralon", Categories: []string{"Tornilleria", "Herramientas Manuales", "Herramientas Electricas", "Materiales de Construccion", "Pinturas", "Plomeria y Sanitarios"}, IsActive: true},
	{ID: "bazar", Name: "Bazar", Slug: "bazar", Categories: []string{"Cocina", "Bazar", "Decoracion", "Organizacion"}, IsActive: true},
	{ID: "jugueteria", Name: "Juguetería", Slug: "jugueteria", Categories: []string{"Bebes y Ninos", "Juguetes"}, IsActive: true},
	{ID: "ropa", Name: "Indumentaria", Slug: "ropa", Categories: []string{"Remeras", "Pantalones", "Buzos", "Camperas", "Zapatillas"}, IsActive: true},
	{ID: "electricidad", Name: "Materiales Eléctricos / Iluminación", Slug: "electricidad", Categories: []string{"Electricidad", "Iluminacion"}, IsActive: true},
	{ID: "zapateria", Name: "Zapatería", Slug: "zapateria", Categories: []string{"Calzado", "Escarpines", "Calzado Deportivo"}, IsActive: true},
	{ID: "deportes", Name: "Indumentaria y Accesorios Deportivos", Slug: "deportes", Categories: []string{"Ropa Deportiva", "Indumentaria Deportiva", "Deportes", "Accesorios Deportivos"}, IsActive: true},
}

// NewListTemplatesUseCase crea una nueva instancia del caso de uso
func NewListTemplatesUseCase(repo port.ListTemplatesRepository) *ListTemplatesUseCase {
	return &ListTemplatesUseCase{
		repo: repo,
	}
}

// Execute ejecuta el caso de uso para listar templates
func (uc *ListTemplatesUseCase) Execute(ctx context.Context) (*ListTemplatesResponse, error) {
	templates, err := uc.repo.LoadTemplatesFromBusinessTypeTemplates(ctx)
	if err != nil {
		return nil, err
	}

	// Merge con legacy: agregar los que no están (por slug)
	hasSlug := make(map[string]bool)
	for _, t := range templates {
		if t.Slug != "" {
			hasSlug[t.Slug] = true
		}
		if t.Name != "" && strings.Contains(strings.ToLower(t.Name), "ferreter") {
			hasSlug["ferreteria-corralon"] = true
		}
	}

	for _, lt := range legacyTemplates {
		if !hasSlug[lt.Slug] {
			templates = append(templates, lt)
			hasSlug[lt.Slug] = true
		}
	}

	if len(templates) == 0 {
		templates = append([]port.ListTemplate{}, legacyTemplates...)
	}

	// Mapear port.ListTemplate a usecase.Template
	result := make([]Template, len(templates))
	for i, t := range templates {
		result[i] = Template{
			ID:          t.ID,
			Name:        t.Name,
			Slug:        t.Slug,
			Description: t.Description,
			Categories:  t.Categories,
			IsActive:    t.IsActive,
		}
	}

	return &ListTemplatesResponse{
		Templates: result,
		Total:     len(result),
	}, nil
}
