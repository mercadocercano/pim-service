-- Migration 040: Migrar productos JSONB de business_type_templates → global_products
-- Los seeds anteriores guardaban productos como JSONB en business_type_templates.products.
-- Esta migración los extrae e inserta en global_products para conectarlos al pipeline
-- de scraping y enrichment. El JSONB queda vacío: el refresh nocturno toma el control.

-- Insertar productos del JSONB que no existan ya en global_products (por nombre+marca)
INSERT INTO global_products (
    name,
    brand,
    category,
    price,
    source,
    source_reliability,
    quality_score,
    is_verified,
    is_active,
    business_type,
    metadata,
    created_at,
    updated_at
)
SELECT
    p->>'name'                                          AS name,
    NULLIF(TRIM(p->>'brand'), '')                       AS brand,
    NULLIF(TRIM(p->>'category_slug'), '')               AS category,
    CASE
        WHEN p->>'price_reference' IS NOT NULL
             AND p->>'price_reference' != ''
        THEN (p->>'price_reference')::NUMERIC
        ELSE NULL
    END                                                 AS price,
    'seed'                                              AS source,
    0.5                                                 AS source_reliability,
    -- quality_score: 25 pts por cada campo presente (name siempre, brand, price, image)
    (
        25
        + CASE WHEN p->>'brand'           IS NOT NULL AND p->>'brand' != ''           THEN 25 ELSE 0 END
        + CASE WHEN p->>'price_reference' IS NOT NULL AND p->>'price_reference' != '' THEN 25 ELSE 0 END
        -- image_url: los seeds no la tienen → +0; llegará via webdata enrichment
    )                                                   AS quality_score,
    FALSE                                               AS is_verified,
    TRUE                                                AS is_active,
    bt.code                                             AS business_type,
    jsonb_build_object(
        'unit',       COALESCE(p->>'unit', 'unidad'),
        'sku_prefix', COALESCE(p->>'sku_prefix', ''),
        'migrated_from', 'business_type_templates'
    )                                                   AS metadata,
    CURRENT_TIMESTAMP                                   AS created_at,
    CURRENT_TIMESTAMP                                   AS updated_at
FROM
    business_type_templates btt
    JOIN business_types bt ON bt.id = btt.business_type_id
    CROSS JOIN LATERAL jsonb_array_elements(
        CASE
            WHEN jsonb_typeof(btt.products) = 'array' THEN btt.products
            ELSE '[]'::jsonb
        END
    ) AS p
WHERE
    btt.products IS NOT NULL
    AND jsonb_array_length(
        CASE WHEN jsonb_typeof(btt.products) = 'array' THEN btt.products ELSE '[]'::jsonb END
    ) > 0
    AND p->>'name' IS NOT NULL
    AND TRIM(p->>'name') != ''
    -- Evitar duplicados: no insertar si ya existe con mismo nombre y misma marca
    AND NOT EXISTS (
        SELECT 1
        FROM global_products gp
        WHERE gp.name = TRIM(p->>'name')
          AND COALESCE(gp.brand, '') = COALESCE(NULLIF(TRIM(p->>'brand'), ''), '')
    );

-- Reportar cuántos productos se migraron por rubro
DO $$
DECLARE
    migrated_count INTEGER;
BEGIN
    SELECT COUNT(*) INTO migrated_count
    FROM global_products
    WHERE source = 'seed'
      AND created_at >= NOW() - INTERVAL '5 seconds';

    RAISE NOTICE 'Migración completada: % productos insertados en global_products desde templates', migrated_count;
END $$;

-- Vaciar el JSONB de productos en los templates que ya fueron migrados.
-- El refresh nocturno (RefreshTemplateProductsUseCase) reconstruirá
-- business_type_product_templates.suggested_products desde global_products.
UPDATE business_type_templates
SET products    = '[]'::jsonb,
    updated_at  = CURRENT_TIMESTAMP
WHERE products IS NOT NULL
  AND jsonb_typeof(products) = 'array'
  AND jsonb_array_length(products) > 0;
