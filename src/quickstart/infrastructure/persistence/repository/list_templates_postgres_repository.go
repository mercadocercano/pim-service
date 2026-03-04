package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"saas-mt-pim-service/src/quickstart/domain/port"
)

// ListTemplatesPostgresRepository implementa ListTemplatesRepository para PostgreSQL
type ListTemplatesPostgresRepository struct {
	db *sql.DB
}

// NewListTemplatesPostgresRepository crea una nueva instancia del repositorio
func NewListTemplatesPostgresRepository(db *sql.DB) port.ListTemplatesRepository {
	return &ListTemplatesPostgresRepository{db: db}
}

type templateCategoryPayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// LoadTemplatesFromBusinessTypeTemplates carga templates desde business_type_templates
func (r *ListTemplatesPostgresRepository) LoadTemplatesFromBusinessTypeTemplates(ctx context.Context) ([]port.ListTemplate, error) {
	query := `
		SELECT id, name, description, categories, is_default, region
		FROM business_type_templates
		WHERE is_active = true
		  AND region IN ('AR', 'GLOBAL')
		  AND COALESCE(jsonb_array_length(categories), 0) >= 3
		ORDER BY is_default DESC,
		         CASE WHEN region = 'AR' THEN 0 ELSE 1 END,
		         name
		LIMIT 8
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying business_type_templates: %w", err)
	}
	defer rows.Close()

	templates := make([]port.ListTemplate, 0)
	for rows.Next() {
		var id, name, description sql.NullString
		var categoriesRaw []byte
		var isDefault sql.NullBool
		var region sql.NullString
		if err := rows.Scan(&id, &name, &description, &categoriesRaw, &isDefault, &region); err != nil {
			return nil, fmt.Errorf("error scanning business_type_template: %w", err)
		}
		if !id.Valid || id.String == "" {
			continue
		}

		categories, err := parseTemplateCategoryNames(categoriesRaw)
		if err != nil {
			return nil, err
		}

		templates = append(templates, port.ListTemplate{
			ID:          id.String,
			Name:        name.String,
			Slug:        id.String,
			Description: description.String,
			Categories:  categories,
			IsActive:    true,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating business_type_templates: %w", err)
	}

	return templates, nil
}

func parseTemplateCategoryNames(categoriesRaw []byte) ([]string, error) {
	if len(categoriesRaw) == 0 {
		return nil, nil
	}

	var payload []templateCategoryPayload
	if err := json.Unmarshal(categoriesRaw, &payload); err != nil {
		return nil, fmt.Errorf("error parsing template categories: %w", err)
	}

	categories := make([]string, 0, len(payload))
	for _, category := range payload {
		if category.Name != "" {
			categories = append(categories, category.Name)
			continue
		}
		if category.Slug != "" {
			categories = append(categories, category.Slug)
		}
	}

	return categories, nil
}
