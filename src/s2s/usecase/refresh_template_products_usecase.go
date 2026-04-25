package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

// RefreshTemplateProductsUseCase re-calculates business_type_product_templates
// from verified global_products. Idempotent — safe to call repeatedly.
type RefreshTemplateProductsUseCase struct {
	db *sql.DB
}

func NewRefreshTemplateProductsUseCase(db *sql.DB) *RefreshTemplateProductsUseCase {
	return &RefreshTemplateProductsUseCase{db: db}
}

type RefreshResult struct {
	TemplatesUpdated int `json:"templates_updated"`
}

func (uc *RefreshTemplateProductsUseCase) Execute(ctx context.Context) (*RefreshResult, error) {
	log.Println("[refresh-template-products] Starting refresh from global_products...")

	query := `
	WITH template_categories AS (
	  SELECT
	    btt.id AS template_id,
	    bt.code AS business_type_code,
	    (cat_obj->>'slug') AS category_slug
	  FROM business_type_templates btt
	  JOIN business_types bt ON bt.id = btt.business_type_id
	  CROSS JOIN LATERAL jsonb_array_elements(btt.categories) AS cat_obj
	  WHERE btt.is_default = true
	    AND btt.is_active = true
	),
	ranked_products AS (
	  SELECT
	    tc.template_id,
	    gp.id AS global_product_id,
	    gp.brand,
	    tc.category_slug,
	    ROW_NUMBER() OVER (
	      PARTITION BY tc.template_id, tc.category_slug
	      ORDER BY gp.quality_score DESC, gp.source_reliability DESC, gp.created_at DESC
	    ) AS rn
	  FROM template_categories tc
	  JOIN global_products gp
	    ON gp.category = tc.category_slug
	   AND (
	     gp.business_type = tc.business_type_code
	     OR tc.business_type_code = ANY(
	       SELECT jsonb_array_elements_text(gp.also_sold_in)
	     )
	   )
	  WHERE gp.is_active = true
	    AND gp.is_verified = true
	),
	limited_products AS (
	  SELECT
	    template_id,
	    global_product_id,
	    brand
	  FROM ranked_products
	  WHERE rn <= 30
	),
	agg AS (
	  SELECT
	    template_id,
	    jsonb_agg(global_product_id ORDER BY global_product_id) AS suggested_products,
	    (
	      SELECT jsonb_agg(brand ORDER BY cnt DESC, brand)
	      FROM (
	        SELECT brand, COUNT(*) AS cnt
	        FROM limited_products lp2
	        WHERE lp2.template_id = lp.template_id
	          AND lp2.brand IS NOT NULL
	          AND trim(lp2.brand) <> ''
	        GROUP BY brand
	        ORDER BY cnt DESC, brand
	        LIMIT 10
	      ) b
	    ) AS priority_brands
	  FROM limited_products lp
	  GROUP BY template_id
	)
	INSERT INTO business_type_product_templates (
	  business_type_template_id,
	  suggested_products,
	  priority_brands,
	  max_products_per_category
	)
	SELECT
	  a.template_id,
	  COALESCE(a.suggested_products, '[]'::jsonb),
	  COALESCE(a.priority_brands, '[]'::jsonb),
	  30
	FROM agg a
	ON CONFLICT (business_type_template_id) DO UPDATE SET
	  suggested_products = EXCLUDED.suggested_products,
	  priority_brands = EXCLUDED.priority_brands,
	  max_products_per_category = EXCLUDED.max_products_per_category,
	  updated_at = CURRENT_TIMESTAMP`

	result, err := uc.db.ExecContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("refresh template products: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	log.Printf("[refresh-template-products] Done: %d templates updated", rowsAffected)

	return &RefreshResult{
		TemplatesUpdated: int(rowsAffected),
	}, nil
}
