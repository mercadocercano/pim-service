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

const legacyTemplateID = "ferreteria-corralon"

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

	hasFerreteria := false
	for _, template := range templates {
		if template.Name != "" && strings.Contains(strings.ToLower(template.Name), "ferreter") {
			hasFerreteria = true
			break
		}
	}

	if !hasFerreteria {
		templates = append(templates, port.ListTemplate{
			ID:          legacyTemplateID,
			Name:        "Ferreteria / Corralon",
			Slug:        "ferreteria-corralon",
			Description: "Template completo para negocios de ferreteria y corralon con 6 categorias principales",
			Categories: []string{
				"Tornilleria",
				"Herramientas Manuales",
				"Herramientas Electricas",
				"Materiales de Construccion",
				"Pinturas",
				"Plomeria y Sanitarios",
			},
			IsActive: true,
		})
	}

	if len(templates) == 0 {
		templates = []port.ListTemplate{
			{
				ID:          legacyTemplateID,
				Name:        "Ferreteria / Corralon",
				Slug:        "ferreteria-corralon",
				Description: "Template completo para negocios de ferreteria y corralon con 6 categorias principales",
				Categories: []string{
					"Tornilleria",
					"Herramientas Manuales",
					"Herramientas Electricas",
					"Materiales de Construccion",
					"Pinturas",
					"Plomeria y Sanitarios",
				},
				IsActive: true,
			},
		}
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
