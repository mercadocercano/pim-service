package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"saas-mt-pim-service/src/quickstart/domain/port"
)

// GetCategoriesByBusinessTypePostgresRepository implementa GetCategoriesByBusinessTypeRepository para PostgreSQL
type GetCategoriesByBusinessTypePostgresRepository struct {
	db *sql.DB
}

// NewGetCategoriesByBusinessTypePostgresRepository crea una nueva instancia del repositorio
func NewGetCategoriesByBusinessTypePostgresRepository(db *sql.DB) port.GetCategoriesByBusinessTypeRepository {
	return &GetCategoriesByBusinessTypePostgresRepository{db: db}
}

// GetCategoriesByBusinessType obtiene las categorías del template activo para el tipo de negocio dado
func (r *GetCategoriesByBusinessTypePostgresRepository) GetCategoriesByBusinessType(ctx context.Context, businessTypeSlug string) ([]port.CategoryByBusinessType, error) {
	query := `
		SELECT btt.categories
		FROM business_type_templates btt
		JOIN business_types bt ON bt.id = btt.business_type_id
		WHERE bt.code = $1 AND btt.is_active = true
		ORDER BY btt.is_default DESC
		LIMIT 1
	`

	var categoriesRaw []byte
	err := r.db.QueryRowContext(ctx, query, businessTypeSlug).Scan(&categoriesRaw)
	if err == sql.ErrNoRows {
		return []port.CategoryByBusinessType{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying categories for business type %q: %w", businessTypeSlug, err)
	}

	return parseCategoryObjects(categoriesRaw)
}

func parseCategoryObjects(categoriesRaw []byte) ([]port.CategoryByBusinessType, error) {
	if len(categoriesRaw) == 0 {
		return []port.CategoryByBusinessType{}, nil
	}

	var payload []templateCategoryPayload
	if err := json.Unmarshal(categoriesRaw, &payload); err != nil {
		return nil, fmt.Errorf("error parsing categories JSON: %w", err)
	}

	categories := make([]port.CategoryByBusinessType, 0, len(payload))
	for _, p := range payload {
		if p.ParentSlug != "" {
			continue
		}
		categories = append(categories, port.CategoryByBusinessType{
			ID:   p.ID,
			Name: p.Name,
			Slug: p.Slug,
		})
	}

	return categories, nil
}
