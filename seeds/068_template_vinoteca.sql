-- Seed 068: Template curado para Vinoteca (business_type code='vinoteca')
-- CICLO: cycle-004-brands-catalog-expansion
-- PROPÓSITO: Categorías, marcas y productos ancla para vinotecas argentinas (foco NEA/Posadas)
-- IDEMPOTENTE: INSERT business_type si no existe; INSERT/UPDATE template
-- REQUIERE: 001 (business_types), 004 (business_type_quickstart_templates), 063 (vinos), 062 (cervezas aperitivos)
-- VERSION: 1.0.0-curated | generated_by: catalog-researcher | FECHA: 2026-04-21
--
-- NOTA NEA: Las vinotecas en Posadas combinan vinos, cervezas artesanales, aperitivos y
--           en algunos casos espumantes y sidras. La venta por copa no aplica en este template.

-- Paso 1: Asegurar que existe business_type vinoteca
INSERT INTO business_types (code, name, description, icon, color, sort_order, is_active)
VALUES (
  'vinoteca',
  'Vinoteca',
  'Venta especializada de vinos, espumantes, cervezas artesanales y aperitivos',
  'wine',
  '#7B1414',
  75,
  true
)
ON CONFLICT (code) DO NOTHING;

-- Paso 2: INSERT o UPDATE del template default
INSERT INTO business_type_templates (
  business_type_id, name, description, region, categories, brands, products, attributes,
  is_default, is_active, generated_by, version
)
SELECT
  bt.id,
  'Vinoteca',
  'Template curado para vinotecas — vinos, espumantes, aperitivos y cervezas artesanales',
  'AR',

  -- CATEGORÍAS
  '[
    {"slug": "vinos-tintos",            "name": "Vinos Tintos"},
    {"slug": "vinos-blancos",           "name": "Vinos Blancos"},
    {"slug": "vinos-rosados",           "name": "Vinos Rosados"},
    {"slug": "espumantes",              "name": "Espumantes y Champagnes"},
    {"slug": "vermouths-aperitivos",    "name": "Vermouths y Aperitivos"},
    {"slug": "cervezas-artesanales",    "name": "Cervezas Artesanales y Premium"}
  ]'::jsonb,

  -- MARCAS (agrupadas por segmento)
  '[
    {"name": "Trapiche",        "segment": "premium",   "suggested_for_categories": ["vinos-tintos","vinos-blancos","vinos-rosados","espumantes"]},
    {"name": "Norton",          "segment": "premium",   "suggested_for_categories": ["vinos-tintos","vinos-blancos"]},
    {"name": "Luigi Bosca",     "segment": "premium",   "suggested_for_categories": ["vinos-tintos","vinos-blancos","espumantes"]},
    {"name": "Rutini",          "segment": "premium",   "suggested_for_categories": ["vinos-tintos","vinos-blancos"]},
    {"name": "Catena",          "segment": "premium",   "suggested_for_categories": ["vinos-tintos","vinos-blancos"]},
    {"name": "Finca Las Moras", "segment": "medio",     "suggested_for_categories": ["vinos-tintos","vinos-blancos","vinos-rosados"]},
    {"name": "Alamos",          "segment": "medio",     "suggested_for_categories": ["vinos-tintos","vinos-blancos"]},
    {"name": "Bianchi",         "segment": "medio",     "suggested_for_categories": ["vinos-tintos","espumantes"]},
    {"name": "Santa Julia",     "segment": "medio",     "suggested_for_categories": ["vinos-tintos","vinos-blancos","vinos-rosados"]},
    {"name": "Gato Negro",      "segment": "medio",     "suggested_for_categories": ["vinos-tintos","vinos-blancos","vinos-rosados"]},
    {"name": "Callia",          "segment": "economico", "suggested_for_categories": ["vinos-tintos","vinos-blancos"]},
    {"name": "La Linda",        "segment": "economico", "suggested_for_categories": ["vinos-tintos","vinos-blancos"]},
    {"name": "Vasco Viejo",     "segment": "economico", "suggested_for_categories": ["vinos-tintos","vinos-blancos"]},
    {"name": "Don Valentín Lacrado", "segment": "economico", "suggested_for_categories": ["vinos-tintos","vinos-blancos"]},
    {"name": "Quara",           "segment": "economico", "suggested_for_categories": ["vinos-tintos","vinos-blancos"]},
    {"name": "Gancia",          "segment": "aperitivos","suggested_for_categories": ["vermouths-aperitivos","espumantes"]},
    {"name": "Campari",         "segment": "aperitivos","suggested_for_categories": ["vermouths-aperitivos"]},
    {"name": "Cinzano",         "segment": "aperitivos","suggested_for_categories": ["vermouths-aperitivos"]},
    {"name": "Fernet Branca",   "segment": "aperitivos","suggested_for_categories": ["vermouths-aperitivos"]},
    {"name": "Heineken",        "segment": "cervezas",  "suggested_for_categories": ["cervezas-artesanales"]},
    {"name": "Patagonia",       "segment": "cervezas",  "suggested_for_categories": ["cervezas-artesanales"]}
  ]'::jsonb,

  -- PRODUCTOS (≥20, representativos para carga rápida en quickstart)
  '[
    {"name": "Trapiche Malbec 750ml",             "category_slug": "vinos-tintos",         "brand": "Trapiche",     "price_reference": 7500,  "unit": "botella"},
    {"name": "Trapiche Malbec Reserva 750ml",      "category_slug": "vinos-tintos",         "brand": "Trapiche",     "price_reference": 12000, "unit": "botella"},
    {"name": "Norton Malbec 750ml",                "category_slug": "vinos-tintos",         "brand": "Norton",       "price_reference": 9000,  "unit": "botella"},
    {"name": "Finca Las Moras Malbec 750ml",       "category_slug": "vinos-tintos",         "brand": "Finca Las Moras","price_reference": 7000, "unit": "botella"},
    {"name": "Gato Negro Cabernet 750ml",          "category_slug": "vinos-tintos",         "brand": "Gato Negro",   "price_reference": 5500,  "unit": "botella"},
    {"name": "Callia Malbec 750ml",                "category_slug": "vinos-tintos",         "brand": "Callia",        "price_reference": 4500,  "unit": "botella"},
    {"name": "La Linda Malbec 750ml",              "category_slug": "vinos-tintos",         "brand": "La Linda",      "price_reference": 4200,  "unit": "botella"},
    {"name": "Don Valentín Lacrado Tinto 750ml",   "category_slug": "vinos-tintos",         "brand": "Don Valentín Lacrado","price_reference": 3500,"unit": "botella"},
    {"name": "Vasco Viejo Tinto 750ml",            "category_slug": "vinos-tintos",         "brand": "Vasco Viejo",   "price_reference": 3200,  "unit": "botella"},
    {"name": "Trapiche Malbec caja x6",            "category_slug": "vinos-tintos",         "brand": "Trapiche",     "price_reference": 42000, "unit": "caja x6"},
    {"name": "Trapiche Chardonnay 750ml",          "category_slug": "vinos-blancos",        "brand": "Trapiche",     "price_reference": 7000,  "unit": "botella"},
    {"name": "Finca Las Moras Torrontés 750ml",    "category_slug": "vinos-blancos",        "brand": "Finca Las Moras","price_reference": 6500, "unit": "botella"},
    {"name": "Santa Julia Sauvignon Blanc 750ml",  "category_slug": "vinos-blancos",        "brand": "Santa Julia",  "price_reference": 6000,  "unit": "botella"},
    {"name": "Gato Negro Chardonnay 750ml",        "category_slug": "vinos-blancos",        "brand": "Gato Negro",   "price_reference": 5200,  "unit": "botella"},
    {"name": "Callia Blanco Suave 750ml",          "category_slug": "vinos-blancos",        "brand": "Callia",        "price_reference": 4200,  "unit": "botella"},
    {"name": "Gato Negro Rosé 750ml",              "category_slug": "vinos-rosados",        "brand": "Gato Negro",   "price_reference": 5500,  "unit": "botella"},
    {"name": "Santa Julia Rosé 750ml",             "category_slug": "vinos-rosados",        "brand": "Santa Julia",  "price_reference": 6000,  "unit": "botella"},
    {"name": "Trapiche Extra Brut 750ml",          "category_slug": "espumantes",           "brand": "Trapiche",     "price_reference": 10000, "unit": "botella"},
    {"name": "Luigi Bosca Brut Nature 750ml",      "category_slug": "espumantes",           "brand": "Luigi Bosca",  "price_reference": 14000, "unit": "botella"},
    {"name": "Bianchi Extra Brut 750ml",           "category_slug": "espumantes",           "brand": "Bianchi",       "price_reference": 9000,  "unit": "botella"},
    {"name": "Gancia Clasico 750ml",               "category_slug": "espumantes",           "brand": "Gancia",       "price_reference": 7500,  "unit": "botella"},
    {"name": "Fernet Branca 750ml",                "category_slug": "vermouths-aperitivos", "brand": "Fernet Branca","price_reference": 22000, "unit": "botella"},
    {"name": "Fernet Branca 1L",                   "category_slug": "vermouths-aperitivos", "brand": "Fernet Branca","price_reference": 28000, "unit": "botella"},
    {"name": "Campari 750ml",                      "category_slug": "vermouths-aperitivos", "brand": "Campari",      "price_reference": 18000, "unit": "botella"},
    {"name": "Cinzano Bianco 750ml",               "category_slug": "vermouths-aperitivos", "brand": "Cinzano",      "price_reference": 9000,  "unit": "botella"},
    {"name": "Gancia Vermouth 750ml",              "category_slug": "vermouths-aperitivos", "brand": "Gancia",       "price_reference": 8000,  "unit": "botella"},
    {"name": "Heineken 330ml",                     "category_slug": "cervezas-artesanales", "brand": "Heineken",     "price_reference": 2200,  "unit": "botella"},
    {"name": "Heineken 650ml",                     "category_slug": "cervezas-artesanales", "brand": "Heineken",     "price_reference": 3500,  "unit": "botella"},
    {"name": "Patagonia Amber Lager 730ml",        "category_slug": "cervezas-artesanales", "brand": "Patagonia",    "price_reference": 4500,  "unit": "botella"},
    {"name": "Patagonia Bohemian Pilsener 730ml",  "category_slug": "cervezas-artesanales", "brand": "Patagonia",    "price_reference": 4200,  "unit": "botella"}
  ]'::jsonb,

  -- ATRIBUTOS
  '[
    {"name": "Varietal",       "slug": "varietal-vino",    "values": ["Malbec","Cabernet Sauvignon","Syrah","Torrontés","Chardonnay","Sauvignon Blanc","Pinot Grigio","Blend","Rosé"], "applies_to_categories": ["vinos-tintos","vinos-blancos","vinos-rosados"]},
    {"name": "Presentación",   "slug": "presentacion-vino","values": ["187ml","375ml","750ml","1.125L","Caja x6","Damajuana 5L"],                                                      "applies_to_categories": ["vinos-tintos","vinos-blancos","vinos-rosados","espumantes"]},
    {"name": "Método",         "slug": "metodo-espumante", "values": ["Charmat","Champenoise","Extra Brut","Brut Nature","Demi Sec"],                                                  "applies_to_categories": ["espumantes"]},
    {"name": "Presentación",   "slug": "presentacion-licor","values": ["375ml","700ml","750ml","1L"],                                                                                  "applies_to_categories": ["vermouths-aperitivos"]},
    {"name": "Presentación cerveza","slug": "presentacion-cerveza","values": ["330ml","473ml","650ml","730ml","1L"],                                                                   "applies_to_categories": ["cervezas-artesanales"]}
  ]'::jsonb,

  true, true, 'catalog-researcher', '1.0.0-curated'

FROM business_types bt
WHERE bt.code = 'vinoteca'

ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories   = EXCLUDED.categories,
  brands       = EXCLUDED.brands,
  products     = EXCLUDED.products,
  attributes   = EXCLUDED.attributes,
  version      = EXCLUDED.version,
  generated_by = EXCLUDED.generated_by,
  updated_at   = CURRENT_TIMESTAMP;
