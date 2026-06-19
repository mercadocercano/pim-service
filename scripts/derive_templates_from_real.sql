-- ⚠️ DESCARTADO / NO EJECUTAR — SUPERSEDED BY ADR-007 (2026-06-18).
-- Este script fue un atajo (patchear el JSONB editorial de business_type_templates) que
-- el diseño firmado descartó. El camino de referencia es ADR-007
-- (services/pim-service/docs/adr/ADR-007-normalizacion-category-slug-y-regeneracion-templates.md):
-- normalizar global_products.category_slug (resolver en go-shared + tabla overrides),
-- backfill como operación S2S, y switchear el read-path del quickstart de editorial→computado
-- (business_type_product_templates: suggested_products + priority_brands). Se conserva solo como
-- referencia histórica del intento. Los datos que tocó YA fueron revertidos desde el backup.
--
-- derive_templates_from_real.sql
--
-- Deriva las categorías y marcas de los templates por defecto (business_type_templates)
-- DESDE el catálogo real (global_products), reemplazando los valores curados/aspiracionales.
-- Mantiene lo curado como backup en business_type_templates_curated_backup.
--
-- Decisión owner (2026-06-18): "derivar templates de lo real, y tener lo curado como
-- summary/backup". Reencuadra E31 (en vez de mapear real→curado, derivamos curado desde real).
--
-- Normalización (lo real está sucio: paths VTEX, MAYÚSCULAS, duplicados de formato):
--   - Categorías: se toma la HOJA del path (/Kiosco/Chocolates/Chocolates/ → Chocolates),
--     Title Case, dedupe (CHOCOLATES y /.../Chocolates/ colapsan a "Chocolates"). Top 15 por volumen.
--   - Marcas: dedupe case-insensitive (ARCOR + Arcor → la grafía más frecuente). Top 12 por volumen.
--   - products NO se toca → total_products sigue saliendo de su JSONB.
--
-- Idempotente: el backup solo se toma la primera vez (ON CONFLICT DO NOTHING); la derivación
-- es re-ejecutable (recalcula desde el estado actual de global_products).

BEGIN;

-- 1. Tabla de backup de lo curado (una sola vez por template).
CREATE TABLE IF NOT EXISTS business_type_templates_curated_backup (
  template_id  uuid PRIMARY KEY REFERENCES business_type_templates(id) ON DELETE CASCADE,
  categories   jsonb,
  brands       jsonb,
  backed_up_at timestamptz NOT NULL DEFAULT now()
);

-- 2. Snapshot del estado curado actual (solo si no se backupeó antes).
INSERT INTO business_type_templates_curated_backup (template_id, categories, brands)
SELECT id, categories, brands
FROM business_type_templates
WHERE is_default = true
ON CONFLICT (template_id) DO NOTHING;

-- 3. Derivar categorías (top 15) y marcas (top 12) desde global_products y escribirlas.
WITH cat_norm AS (
  SELECT
    bt.id AS bt_id,
    initcap(lower(replace(
      (regexp_split_to_array(btrim(gp.category, '/'), '/'))[
        cardinality(regexp_split_to_array(btrim(gp.category, '/'), '/'))
      ]
    , '_', ' '))) AS cat_name,
    COUNT(*) AS n
  FROM global_products gp
  JOIN business_types bt ON bt.code = gp.business_type
  WHERE gp.is_active AND gp.category IS NOT NULL AND btrim(gp.category) <> ''
  GROUP BY 1, 2
),
cat_ranked AS (
  SELECT bt_id, cat_name, n,
         ROW_NUMBER() OVER (PARTITION BY bt_id ORDER BY n DESC, cat_name) AS rn
  FROM cat_norm
  WHERE cat_name IS NOT NULL AND cat_name <> ''
),
cat_agg AS (
  SELECT bt_id,
    jsonb_agg(
      jsonb_build_object(
        'name', cat_name,
        'slug', NULLIF(regexp_replace(lower(translate(cat_name, 'áéíóúñ', 'aeioun')), '[^a-z0-9]+', '-', 'g'), ''),
        'level', 0
      ) ORDER BY n DESC
    ) AS categories
  FROM cat_ranked
  WHERE rn <= 15
  GROUP BY bt_id
),
brand_norm AS (
  SELECT
    bt.id AS bt_id,
    mode() WITHIN GROUP (ORDER BY gp.brand) AS brand_name,  -- grafía más frecuente
    COUNT(*) AS n
  FROM global_products gp
  JOIN business_types bt ON bt.code = gp.business_type
  WHERE gp.is_active AND gp.brand IS NOT NULL AND btrim(gp.brand) <> ''
  GROUP BY bt.id, lower(btrim(gp.brand))
),
brand_ranked AS (
  SELECT bt_id, brand_name, n,
         ROW_NUMBER() OVER (PARTITION BY bt_id ORDER BY n DESC, brand_name) AS rn
  FROM brand_norm
),
brand_agg AS (
  SELECT bt_id,
    jsonb_agg(
      jsonb_build_object(
        -- Si la grafía elegida es TODO MAYÚSCULAS, sin puntos y >3 chars (ej ARCOR, COFLER),
        -- pasarla a Title Case. Se preservan acrónimos con puntos (F.V.) y siglas cortas.
        'name', CASE
          WHEN brand_name = upper(brand_name)
               AND position('.' in brand_name) = 0
               AND length(brand_name) > 3
          THEN initcap(lower(brand_name))
          ELSE brand_name
        END,
        'suggested_for_categories', '[]'::jsonb
      ) ORDER BY n DESC
    ) AS brands
  FROM brand_ranked
  WHERE rn <= 12
  GROUP BY bt_id
)
UPDATE business_type_templates btt
SET categories = COALESCE(ca.categories, btt.categories),
    brands     = COALESCE(ba.brands, btt.brands),
    generated_by = 'derived_from_real',
    updated_at = now()
FROM business_types bt
LEFT JOIN cat_agg   ca ON ca.bt_id = bt.id
LEFT JOIN brand_agg ba ON ba.bt_id = bt.id
WHERE btt.business_type_id = bt.id
  AND btt.is_default = true
  AND (ca.categories IS NOT NULL OR ba.brands IS NOT NULL);

COMMIT;
