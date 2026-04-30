-- Seed 058: Actualización de brands en template de Perfumería
-- CICLO: cycle-002-global-brands-colors / ADR-001
-- FUENTE: curación manual — higiene personal, limpieza y cosmética argentina
-- IDEMPOTENTE: INSERT ... ON CONFLICT DO UPDATE — puede correrse N veces
-- REQUIERE: 001 (business_types), 004 (quickstart templates), 056 (marcas adicionales)
-- VERSION: 1.0.0
-- NOTA: Si el template no existe, lo crea. Si existe, actualiza solo el campo brands.

-- Paso 1: Asegurar que existe business_type perfumeria
INSERT INTO business_types (code, name, description, icon, color, sort_order, is_active)
VALUES ('perfumeria', 'Perfumería', 'Venta de higiene personal, cosméticos y productos de limpieza', 'sparkles', '#A855F7', 40, true)
ON CONFLICT (code) DO NOTHING;

-- Paso 2: UPSERT del template de perfumería
INSERT INTO business_type_templates (
  business_type_id, name, description, region, categories, brands, products, attributes,
  is_default, is_active, generated_by, version
)
SELECT
  bt.id,
  'Perfumería',
  'Template curado para perfumerías argentinas',
  'AR',
  '[
    {"slug": "cuidado-personal",  "name": "Cuidado Personal"},
    {"slug": "higiene-bucal",     "name": "Higiene Bucal"},
    {"slug": "limpieza",          "name": "Limpieza del Hogar"},
    {"slug": "cosmeticos",        "name": "Cosméticos y Maquillaje"}
  ]'::jsonb,
  '[
    {"name": "Colgate",          "suggested_for_categories": ["higiene-bucal"]},
    {"name": "Gillette",         "suggested_for_categories": ["cuidado-personal"]},
    {"name": "Dove",             "suggested_for_categories": ["cuidado-personal"]},
    {"name": "Head & Shoulders", "suggested_for_categories": ["cuidado-personal"]},
    {"name": "Pantene",          "suggested_for_categories": ["cuidado-personal"]},
    {"name": "Sedal",            "suggested_for_categories": ["cuidado-personal"]},
    {"name": "Rexona",           "suggested_for_categories": ["cuidado-personal"]},
    {"name": "Axe",              "suggested_for_categories": ["cuidado-personal"]},
    {"name": "Suave",            "suggested_for_categories": ["cuidado-personal"]},
    {"name": "Nivea",            "suggested_for_categories": ["cuidado-personal"]},
    {"name": "Ariel",            "suggested_for_categories": ["limpieza"]},
    {"name": "Skip",             "suggested_for_categories": ["limpieza"]},
    {"name": "Cif",              "suggested_for_categories": ["limpieza"]},
    {"name": "Mr. Músculo",      "suggested_for_categories": ["limpieza"]},
    {"name": "Lysoform",         "suggested_for_categories": ["limpieza"]},
    {"name": "Poett",            "suggested_for_categories": ["limpieza"]},
    {"name": "Glade",            "suggested_for_categories": ["limpieza"]},
    {"name": "L''Oréal",         "suggested_for_categories": ["cosmeticos"]},
    {"name": "Maybelline",       "suggested_for_categories": ["cosmeticos"]},
    {"name": "Avon",             "suggested_for_categories": ["cosmeticos"]},
    {"name": "Natura",           "suggested_for_categories": ["cosmeticos"]},
    {"name": "Algabo",           "suggested_for_categories": ["cuidado-personal"]},
    {"name": "Querubín",         "suggested_for_categories": ["cuidado-personal"]},
    {"name": "Doncella",         "suggested_for_categories": ["cuidado-personal"]}
  ]'::jsonb,
  '[]'::jsonb,
  '[]'::jsonb,
  true, true, 'manual-curation', '1.0.0'
FROM business_types bt
WHERE bt.code = 'perfumeria'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  brands     = EXCLUDED.brands,
  version    = EXCLUDED.version,
  updated_at = CURRENT_TIMESTAMP;

DO $$
DECLARE
  v_count INT;
BEGIN
  SELECT COUNT(*) INTO v_count
  FROM business_type_templates btt
  JOIN business_types bt ON bt.id = btt.business_type_id
  WHERE bt.code = 'perfumeria' AND btt.is_default = true;
  RAISE NOTICE 'Seed 058 — template perfumeria: % registro(s) activo(s)', v_count;
END $$;
