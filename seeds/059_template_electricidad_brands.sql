-- Seed 059: Template de Electricidad con brands curadas
-- CICLO: cycle-002-global-brands-colors / ADR-001
-- FUENTE: curación manual — materiales eléctricos mercado argentino
-- IDEMPOTENTE: INSERT ... ON CONFLICT DO UPDATE — puede correrse N veces
-- REQUIERE: 001 (business_types), 004 (quickstart templates), 056 (marcas adicionales)
-- VERSION: 1.0.0
-- NOTA: Si el template no existe, lo crea. Si existe, actualiza solo el campo brands.

-- Paso 1: Asegurar que existe business_type electricidad
INSERT INTO business_types (code, name, description, icon, color, sort_order, is_active)
VALUES ('electricidad', 'Electricidad', 'Venta de materiales eléctricos, iluminación y automatización', 'zap', '#FACC15', 50, true)
ON CONFLICT (code) DO NOTHING;

-- Paso 2: UPSERT del template de electricidad
INSERT INTO business_type_templates (
  business_type_id, name, description, region, categories, brands, products, attributes,
  is_default, is_active, generated_by, version
)
SELECT
  bt.id,
  'Electricidad',
  'Template curado para locales de materiales eléctricos argentinos',
  'AR',
  '[
    {"slug": "tableros",              "name": "Tableros y Protección"},
    {"slug": "automatizacion",        "name": "Automatización"},
    {"slug": "cables-conductores",    "name": "Cables y Conductores"},
    {"slug": "enchufes-interruptores","name": "Enchufes e Interruptores"},
    {"slug": "iluminacion",           "name": "Iluminación"}
  ]'::jsonb,
  '[
    {"name": "Schneider Electric", "suggested_for_categories": ["tableros","automatizacion"]},
    {"name": "ABB",                "suggested_for_categories": ["tableros","automatizacion"]},
    {"name": "Siemens",            "suggested_for_categories": ["tableros","automatizacion"]},
    {"name": "Cambre",             "suggested_for_categories": ["enchufes-interruptores"]},
    {"name": "Sica",               "suggested_for_categories": ["cables-conductores"]},
    {"name": "Trefilcon",          "suggested_for_categories": ["cables-conductores"]},
    {"name": "Genrod",             "suggested_for_categories": ["cables-conductores"]},
    {"name": "Roker",              "suggested_for_categories": ["iluminacion"]},
    {"name": "Macroled",           "suggested_for_categories": ["iluminacion"]},
    {"name": "Jeluz",              "suggested_for_categories": ["iluminacion"]},
    {"name": "BAW",                "suggested_for_categories": ["iluminacion"]},
    {"name": "Kalop",              "suggested_for_categories": ["iluminacion"]},
    {"name": "Lexo",               "suggested_for_categories": ["iluminacion"]},
    {"name": "EQ",                 "suggested_for_categories": ["iluminacion"]},
    {"name": "Philips",            "suggested_for_categories": ["iluminacion"]},
    {"name": "3M",                 "suggested_for_categories": ["cables-conductores","herramientas"]}
  ]'::jsonb,
  '[]'::jsonb,
  '[]'::jsonb,
  true, true, 'manual-curation', '1.0.0'
FROM business_types bt
WHERE bt.code = 'electricidad'
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
  WHERE bt.code = 'electricidad' AND btt.is_default = true;
  RAISE NOTICE 'Seed 059 — template electricidad: % registro(s) activo(s)', v_count;
END $$;
