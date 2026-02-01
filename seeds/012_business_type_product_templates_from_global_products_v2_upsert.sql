-- Seed v2 (idempotente): Poblar productos sugeridos por template default desde global_products verificados
-- Usa business_type_product_templates (JSONB) por ser el modelo más simple para el quickstart.
--
-- Selección:
-- - Solo productos: global_products.is_active=true AND is_verified=true
-- - Categorías: se toman desde business_type_templates.categories (campo JSONB [{id, slug}, ...])
-- - Límite: max_products_per_category (default 30 por template en este seeder)
-- - Orden: quality_score DESC, source_reliability DESC, created_at DESC
--
-- Resultado:
-- - business_type_product_templates.suggested_products: JSONB array de UUIDs (global_products.id)
-- - priority_brands: top brands presentes en el set sugerido

WITH template_categories AS (
  SELECT
    btt.id AS template_id,
    (cat_obj->>'slug') AS category_slug
  FROM business_type_templates btt
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
  updated_at = CURRENT_TIMESTAMP;


