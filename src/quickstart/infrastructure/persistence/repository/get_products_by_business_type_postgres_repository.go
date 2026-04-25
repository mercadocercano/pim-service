package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/lib/pq"
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

// GetProductsByBusinessType retorna productos del template computado si están disponibles,
// o del JSONB editorial como fallback. No rompe el flujo existente.
func (r *GetProductsByBusinessTypePostgresRepository) GetProductsByBusinessType(ctx context.Context, businessTypeSlug string) ([]port.TemplateProduct, error) {
	productIDs, err := r.fetchComputedTemplateIDs(ctx, businessTypeSlug)
	if err != nil {
		return nil, err
	}

	if len(productIDs) > 0 {
		products, err := r.resolveGlobalProducts(ctx, businessTypeSlug, productIDs)
		if err != nil {
			return nil, err
		}
		if len(products) > 0 {
			return products, nil
		}
	}

	return r.fetchEditorialProducts(ctx, businessTypeSlug)
}

// fetchComputedTemplateIDs obtiene los UUIDs de suggested_products del template computado activo.
// Retorna slice vacío si no hay registro o si suggested_products es un array vacío.
func (r *GetProductsByBusinessTypePostgresRepository) fetchComputedTemplateIDs(ctx context.Context, slug string) ([]string, error) {
	query := `
		SELECT btpt.suggested_products
		FROM business_type_product_templates btpt
		JOIN business_type_templates btt ON btt.id = btpt.business_type_template_id
		JOIN business_types bt ON bt.id = btt.business_type_id
		WHERE bt.code = $1
		  AND btt.is_active = true
		  AND btt.is_default = true
		ORDER BY btt.is_default DESC
		LIMIT 1
	`

	var rawJSON []byte
	err := r.db.QueryRowContext(ctx, query, slug).Scan(&rawJSON)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("fetchComputedTemplateIDs %q: %w", slug, err)
	}

	var ids []string
	if err := json.Unmarshal(rawJSON, &ids); err != nil {
		return nil, fmt.Errorf("fetchComputedTemplateIDs parse json %q: %w", slug, err)
	}

	return ids, nil
}

// resolveGlobalProducts convierte UUIDs de suggested_products a TemplateProduct consultando global_products.
func (r *GetProductsByBusinessTypePostgresRepository) resolveGlobalProducts(ctx context.Context, slug string, ids []string) ([]port.TemplateProduct, error) {
	query := `
		SELECT gp.name, COALESCE(gp.brand, ''), COALESCE(gp.category, ''), gp.quality_score
		FROM global_products gp
		WHERE gp.id = ANY($1)
		  AND gp.is_active = true
		ORDER BY gp.quality_score DESC
		LIMIT 30
	`

	rows, err := r.db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, fmt.Errorf("resolveGlobalProducts %q: %w", slug, err)
	}
	defer rows.Close()

	var products []port.TemplateProduct
	for rows.Next() {
		var p port.TemplateProduct
		var qualityScore float64
		if err := rows.Scan(&p.Name, &p.Brand, &p.CategorySlug, &qualityScore); err != nil {
			return nil, fmt.Errorf("resolveGlobalProducts scan %q: %w", slug, err)
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("resolveGlobalProducts rows %q: %w", slug, err)
	}

	log.Printf("[quickstart-source] business_type=%s source=computed products=%d", slug, len(products))
	return products, nil
}

// fetchEditorialProducts obtiene productos del JSONB editorial en business_type_templates (fallback).
func (r *GetProductsByBusinessTypePostgresRepository) fetchEditorialProducts(ctx context.Context, slug string) ([]port.TemplateProduct, error) {
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
	err := r.db.QueryRowContext(ctx, query, slug).Scan(&productsRaw)
	if err == sql.ErrNoRows {
		return []port.TemplateProduct{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("fetchEditorialProducts %q: %w", slug, err)
	}

	var products []port.TemplateProduct
	if err := json.Unmarshal(productsRaw, &products); err != nil {
		return nil, fmt.Errorf("fetchEditorialProducts parse json %q: %w", slug, err)
	}

	log.Printf("[quickstart-source] business_type=%s source=editorial products=%d", slug, len(products))
	return products, nil
}
