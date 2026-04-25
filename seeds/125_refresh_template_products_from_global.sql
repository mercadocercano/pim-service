WITH top_products AS (
  SELECT
    gp.name,
    COALESCE(gp.brand, '') as brand,
    gp.category as category_slug,
    COALESCE(gp.price, 0) as price,
    COALESCE(gp.image_url, '') as image_url,
    gp.business_type,
    gp.quality_score,
    ROW_NUMBER() OVER (
      PARTITION BY gp.business_type, gp.category
      ORDER BY gp.quality_score DESC
    ) as rn
  FROM global_products gp
  WHERE gp.is_verified = true
    AND gp.is_active = true
),
limited AS (
  SELECT * FROM top_products WHERE rn <= 8
),
aggregated AS (
  SELECT
    l.business_type,
    jsonb_agg(
      jsonb_build_object(
        'name', l.name,
        'brand', l.brand,
        'category_slug', l.category_slug,
        'price_reference', l.price,
        'image_url', l.image_url,
        'unit', 'unidad',
        'sku_prefix', UPPER(REPLACE(l.business_type, '-', '_'))
      ) ORDER BY l.quality_score DESC
    ) as products_json
  FROM limited l
  GROUP BY l.business_type
)
UPDATE business_type_templates btt
SET products = a.products_json
FROM aggregated a
JOIN business_types bt ON bt.code = a.business_type
WHERE btt.business_type_id = bt.id
  AND btt.is_active = true;
