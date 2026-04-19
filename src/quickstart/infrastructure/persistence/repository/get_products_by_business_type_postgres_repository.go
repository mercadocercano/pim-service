package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"saas-mt-pim-service/src/quickstart/domain/port"
)

// GetProductsByBusinessTypePostgresRepository implementa GetProductsByBusinessTypeRepository para PostgreSQL
type GetProductsByBusinessTypePostgresRepository struct {
	db *sql.DB
}

// NewGetProductsByBusinessTypePostgresRepository crea una nueva instancia del repositorio
func NewGetProductsByBusinessTypePostgresRepository(db *sql.DB) port.GetProductsByBusinessTypeRepository {
	return &GetProductsByBusinessTypePostgresRepository{db: db}
}

// GetProductsByBusinessType obtiene los productos sugeridos del template por slug de tipo de negocio
func (r *GetProductsByBusinessTypePostgresRepository) GetProductsByBusinessType(ctx context.Context, businessTypeSlug string) ([]port.TemplateProduct, error) {
	query := `
		SELECT COALESCE(btt.products, '[]'::jsonb)
		FROM business_type_templates btt
		JOIN business_types bt ON bt.id = btt.business_type_id
		WHERE bt.code = $1
		  AND btt.is_active = true
		ORDER BY btt.is_default DESC,
		         CASE WHEN btt.region = 'AR' THEN 0 ELSE 1 END
		LIMIT 1
	`

	var productsRaw []byte
	err := r.db.QueryRowContext(ctx, query, businessTypeSlug).Scan(&productsRaw)
	if err == sql.ErrNoRows {
		return []port.TemplateProduct{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying products for business type %q: %w", businessTypeSlug, err)
	}

	var products []port.TemplateProduct
	if err := json.Unmarshal(productsRaw, &products); err != nil {
		return nil, fmt.Errorf("error parsing products json for business type %q: %w", businessTypeSlug, err)
	}

	return products, nil
}
