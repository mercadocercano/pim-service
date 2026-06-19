-- pim_parity_audit.sql
-- Auditoría READ-ONLY de paridad de pim_db (local vs producción).
-- Se corre idéntico contra ambas bases; la salida es determinista (ordenada)
-- para que `diff` sea limpio. No modifica datos.
--
-- Uso:
--   psql -h <host> -p <port> -U postgres -d pim_db -f pim_parity_audit.sql
--
-- Dimensiones:
--   D1 Business types
--   D2 Templates por business type (+ business_type_product_templates)
--   D3 Marcas renderizables (identidad visual)
--   D4 Productos con imagen en CDN de Digital Ocean
--   D5 Readiness UX por business type ("listo para vender en 2 min")

\pset pager off
\pset footer off
\timing off
\set ON_ERROR_STOP off

\echo '================================================================'
\echo 'D1  BUSINESS TYPES'
\echo '================================================================'
SELECT count(*) AS total,
       count(*) FILTER (WHERE is_active) AS activos
FROM business_types;

SELECT code, name, is_active, sort_order
FROM business_types
ORDER BY code;

\echo ''
\echo '================================================================'
\echo 'D2  TEMPLATES POR BUSINESS TYPE'
\echo '================================================================'
SELECT bt.code,
       count(t.*)                                          AS templates,
       count(*) FILTER (WHERE t.is_default)                AS defaults,
       coalesce(max(jsonb_array_length(t.categories)), 0)  AS max_cats,
       coalesce(max(jsonb_array_length(t.products)), 0)    AS max_prods,
       coalesce(max(jsonb_array_length(t.brands)), 0)      AS max_brands
FROM business_types bt
LEFT JOIN business_type_templates t ON t.business_type_id = bt.id
GROUP BY bt.code
ORDER BY bt.code;

\echo '-- business_type_product_templates: cobertura de suggested_products --'
SELECT count(*)                                                        AS template_products_rows,
       count(*) FILTER (WHERE jsonb_array_length(suggested_products) > 0) AS con_suggested,
       coalesce(max(jsonb_array_length(suggested_products)), 0)        AS max_suggested,
       coalesce(max(jsonb_array_length(priority_brands)), 0)          AS max_priority_brands
FROM business_type_product_templates;

\echo ''
\echo '================================================================'
\echo 'D3  MARCAS RENDERIZABLES (marketplace_brands)'
\echo '================================================================'
SELECT count(*) AS total,
       count(*) FILTER (WHERE deleted_at IS NULL AND is_active)        AS activas,
       count(*) FILTER (WHERE background_color IS NOT NULL
                          AND text_color IS NOT NULL
                          AND typography IS NOT NULL)                  AS renderizables,
       count(*) FILTER (WHERE logo_url IS NOT NULL AND logo_url <> '') AS con_logo,
       count(*) FILTER (WHERE verification_status = 'verified')        AS verificadas
FROM marketplace_brands;

\echo '-- marcas activas SIN identidad visual completa (orden por name) --'
SELECT name,
       (background_color IS NOT NULL) AS bg,
       (text_color IS NOT NULL)       AS txt,
       (typography IS NOT NULL)       AS typo,
       (logo_url IS NOT NULL AND logo_url <> '') AS logo
FROM marketplace_brands
WHERE deleted_at IS NULL AND is_active
  AND NOT (background_color IS NOT NULL
           AND text_color IS NOT NULL
           AND typography IS NOT NULL)
ORDER BY name;

\echo ''
\echo '================================================================'
\echo 'D4  PRODUCTOS CON IMAGEN EN CDN DE DIGITAL OCEAN'
\echo '================================================================'
\echo '-- global_products (catálogo global) --'
SELECT count(*) AS total,
       count(*) FILTER (WHERE image_url ILIKE '%digitaloceanspaces.com%')      AS en_cdn_do,
       count(*) FILTER (WHERE image_url IS NOT NULL AND image_url <> ''
                          AND image_url NOT ILIKE '%digitaloceanspaces.com%')  AS otra_url,
       count(*) FILTER (WHERE image_url IS NULL OR image_url = '')             AS sin_imagen
FROM global_products
WHERE is_active;

\echo '-- products (tenant) --'
SELECT count(*) AS total,
       count(*) FILTER (WHERE image_url ILIKE '%digitaloceanspaces.com%')      AS en_cdn_do,
       count(*) FILTER (WHERE image_url IS NOT NULL AND image_url <> ''
                          AND image_url NOT ILIKE '%digitaloceanspaces.com%')  AS otra_url,
       count(*) FILTER (WHERE image_url IS NULL OR image_url = '')             AS sin_imagen
FROM products
WHERE status <> 'deleted';

\echo ''
\echo '================================================================'
\echo 'D5  READINESS UX POR BUSINESS TYPE (global_products)'
\echo '================================================================'
SELECT business_type,
       count(*)                                                          AS prods_catalogo,
       count(*) FILTER (WHERE image_url ILIKE '%digitaloceanspaces.com%') AS con_imagen_cdn,
       round(100.0 * count(*) FILTER (WHERE image_url ILIKE '%digitaloceanspaces.com%')
             / nullif(count(*), 0), 1)                                   AS pct_con_imagen
FROM global_products
WHERE is_active AND business_type IS NOT NULL
GROUP BY business_type
ORDER BY pct_con_imagen ASC, business_type;
