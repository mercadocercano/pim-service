package usecase

import (
	"context"
	"database/sql"
	"fmt"
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
	db *sql.DB
}

// NewListTemplatesUseCase crea una nueva instancia del caso de uso
func NewListTemplatesUseCase(db *sql.DB) *ListTemplatesUseCase {
	return &ListTemplatesUseCase{
		db: db,
	}
}

// Execute ejecuta el caso de uso para listar templates
func (uc *ListTemplatesUseCase) Execute(ctx context.Context) (*ListTemplatesResponse, error) {
	// Por ahora retornamos solo el template de ferretería hardcodeado
	// En el futuro esto se puede extender para leer desde BD
	templates := []Template{
		{
			ID:          "ferreteria-corralon",
			Name:        "Ferretería / Corralón",
			Slug:        "ferreteria-corralon",
			Description: "Template completo para negocios de ferretería y corralón con 6 categorías principales",
			Categories: []string{
				"Tornillería",
				"Herramientas Manuales",
				"Herramientas Eléctricas",
				"Materiales de Construcción",
				"Pinturas",
				"Plomería y Sanitarios",
			},
			IsActive: true,
		},
	}

	// Opcional: Verificar que las categorías existan en la BD
	for i := range templates {
		if err := uc.verifyTemplateCategories(ctx, &templates[i]); err != nil {
			return nil, fmt.Errorf("error verifying template %s: %w", templates[i].ID, err)
		}
	}

	return &ListTemplatesResponse{
		Templates: templates,
		Total:     len(templates),
	}, nil
}

// verifyTemplateCategories verifica que las categorías del template existan en marketplace_categories
func (uc *ListTemplatesUseCase) verifyTemplateCategories(ctx context.Context, template *Template) error {
	// Query para verificar que existan las categorías del template
	query := `
		SELECT COUNT(*) as count
		FROM marketplace_categories
		WHERE parent_id = 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f00'
		AND is_active = true
	`

	var count int
	err := uc.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return fmt.Errorf("error querying categories: %w", err)
	}

	// Verificar que haya al menos 6 categorías (las del template)
	if count < 6 {
		return fmt.Errorf("template categories not found in database (found %d, expected 6)", count)
	}

	return nil
}

