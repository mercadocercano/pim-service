-- Seed 075: Template curado para Casa de Limpieza (business_type code='limpieza')
-- PROPÓSITO: Productos ancla, marcas y atributos curados para casas de limpieza argentinas
-- IDEMPOTENTE: INSERT/UPDATE business_type; INSERT/UPDATE template
-- REQUIERE: 001 (business_types), schema business_type_templates
-- CICLO: cycle-004-brands-catalog-expansion
-- FECHA: 2026-04-22
-- VERSION: 1.0.0-curated | generated_by: manual-curation

-- Paso 1: Asegurar que existe business_type limpieza
INSERT INTO business_types (code, name, description, icon, color, sort_order, is_active)
VALUES (
  'limpieza',
  'Casa de Limpieza',
  'Artículos de limpieza del hogar, productos a granel (cloro, suavizante, detergente), insumos para pileta y accesorios de jardín',
  'sparkles',
  '#00838F',
  47,
  true
)
ON CONFLICT (code) DO UPDATE SET
  name        = EXCLUDED.name,
  description = EXCLUDED.description,
  icon        = EXCLUDED.icon,
  color       = EXCLUDED.color,
  is_active   = EXCLUDED.is_active;

-- Paso 2: INSERT o UPDATE del template default
INSERT INTO business_type_templates (
  business_type_id, name, description, region, categories, brands, products, attributes,
  is_default, is_active, generated_by, version
)
SELECT bt.id, 'Casa de Limpieza', 'Template curado para casas de limpieza y artículos del hogar', 'AR',
  '[
    {"slug": "limpieza-hogar",         "name": "Limpieza del Hogar"},
    {"slug": "liquidos-granel",        "name": "Líquidos a Granel"},
    {"slug": "desodorantes-ambientales","name": "Desodorantes y Ambientales"},
    {"slug": "insumos-pileta",         "name": "Insumos para Pileta"},
    {"slug": "accesorios-limpieza",    "name": "Accesorios y Utensilios"},
    {"slug": "jardin-exterior",        "name": "Jardín y Exterior"}
  ]'::jsonb,
  '[
    {"name": "Skip",        "suggested_for_categories": ["limpieza-hogar","liquidos-granel"]},
    {"name": "Ariel",       "suggested_for_categories": ["limpieza-hogar","liquidos-granel"]},
    {"name": "Cif",         "suggested_for_categories": ["limpieza-hogar"]},
    {"name": "Mr. Músculo", "suggested_for_categories": ["limpieza-hogar"]},
    {"name": "Lysoform",    "suggested_for_categories": ["limpieza-hogar"]},
    {"name": "Ayudín",      "suggested_for_categories": ["limpieza-hogar"]},
    {"name": "Magistral",   "suggested_for_categories": ["limpieza-hogar"]},
    {"name": "Poett",       "suggested_for_categories": ["desodorantes-ambientales"]},
    {"name": "Glade",       "suggested_for_categories": ["desodorantes-ambientales"]},
    {"name": "Clorotec",    "suggested_for_categories": ["insumos-pileta","liquidos-granel"]},
    {"name": "Nataclor",    "suggested_for_categories": ["insumos-pileta"]},
    {"name": "Dove",        "suggested_for_categories": ["limpieza-hogar"]},
    {"name": "Rexona",      "suggested_for_categories": ["limpieza-hogar"]},
    {"name": "Colgate",     "suggested_for_categories": ["limpieza-hogar"]},
    {"name": "Sedal",       "suggested_for_categories": ["limpieza-hogar"]},
    {"name": "Tramontina",  "suggested_for_categories": ["accesorios-limpieza","jardin-exterior"]}
  ]'::jsonb,
  '[
    {"name": "Detergente lavavajillas Skip 750ml",      "category_slug": "limpieza-hogar",          "brand": "Skip",        "price_reference": 2800,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Detergente lavavajillas Skip 1.5L",       "category_slug": "limpieza-hogar",          "brand": "Skip",        "price_reference": 5200,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Jabón en polvo Ariel 800g",               "category_slug": "limpieza-hogar",          "brand": "Ariel",       "price_reference": 4500,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Jabón en polvo Skip 500g",                "category_slug": "limpieza-hogar",          "brand": "Skip",        "price_reference": 2800,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Lavandina Ayudín 1L",                     "category_slug": "limpieza-hogar",          "brand": "Ayudín",      "price_reference": 1800,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Lavandina Magistral 2L",                  "category_slug": "limpieza-hogar",          "brand": "Magistral",   "price_reference": 3200,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Limpiador multiuso Cif 500ml",            "category_slug": "limpieza-hogar",          "brand": "Cif",         "price_reference": 2500,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Mr. Músculo cocina 500ml",                "category_slug": "limpieza-hogar",          "brand": "Mr. Músculo", "price_reference": 3200,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Mr. Músculo baño 500ml",                  "category_slug": "limpieza-hogar",          "brand": "Mr. Músculo", "price_reference": 3200,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Lysoform desinfectante 500ml",            "category_slug": "limpieza-hogar",          "brand": "Lysoform",    "price_reference": 3500,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Suavizante a granel 1L",                  "category_slug": "liquidos-granel",         "brand": null,          "price_reference": 2200,  "unit": "litro",  "sku_prefix": "LIMPIEZA"},
    {"name": "Suavizante a granel 5L",                  "category_slug": "liquidos-granel",         "brand": null,          "price_reference": 9500,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Detergente a granel 1L",                  "category_slug": "liquidos-granel",         "brand": null,          "price_reference": 1800,  "unit": "litro",  "sku_prefix": "LIMPIEZA"},
    {"name": "Detergente a granel 5L",                  "category_slug": "liquidos-granel",         "brand": null,          "price_reference": 7500,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Lavandina a granel 1L",                   "category_slug": "liquidos-granel",         "brand": null,          "price_reference": 1200,  "unit": "litro",  "sku_prefix": "LIMPIEZA"},
    {"name": "Lavandina a granel 5L",                   "category_slug": "liquidos-granel",         "brand": null,          "price_reference": 5000,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Cloro pileta granel 1L",                  "category_slug": "liquidos-granel",         "brand": "Clorotec",    "price_reference": 2500,  "unit": "litro",  "sku_prefix": "LIMPIEZA"},
    {"name": "Desodorante ambiente Poett 360ml",        "category_slug": "desodorantes-ambientales","brand": "Poett",        "price_reference": 3500,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Glade aerosol 360ml",                     "category_slug": "desodorantes-ambientales","brand": "Glade",        "price_reference": 4200,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Cloro tabletas pileta 1kg",               "category_slug": "insumos-pileta",          "brand": "Clorotec",    "price_reference": 12000, "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Algicida 1L",                             "category_slug": "insumos-pileta",          "brand": "Nataclor",    "price_reference": 8500,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Escoba plástica",                         "category_slug": "accesorios-limpieza",     "brand": null,          "price_reference": 4500,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Trapo de piso",                           "category_slug": "accesorios-limpieza",     "brand": null,          "price_reference": 2800,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Balde 10L",                               "category_slug": "accesorios-limpieza",     "brand": null,          "price_reference": 3500,  "unit": "unidad", "sku_prefix": "LIMPIEZA"},
    {"name": "Guantes de látex par",                    "category_slug": "accesorios-limpieza",     "brand": null,          "price_reference": 1800,  "unit": "par",    "sku_prefix": "LIMPIEZA"},
    {"name": "Esponja rejilla x3",                      "category_slug": "accesorios-limpieza",     "brand": null,          "price_reference": 1500,  "unit": "pack",   "sku_prefix": "LIMPIEZA"},
    {"name": "Manguera jardín 1/2\" 10m",               "category_slug": "jardin-exterior",         "brand": null,          "price_reference": 8500,  "unit": "unidad", "sku_prefix": "LIMPIEZA"}
  ]'::jsonb,
  '[
    {"name": "Presentación Granel","slug": "presentacion-granel","values": ["500ml","1L","2L","5L","10L"],                      "applies_to_categories": ["liquidos-granel"]},
    {"name": "Tipo Lavado",        "slug": "tipo-lavado",        "values": ["Ropa blanca","Ropa color","Universal"],            "applies_to_categories": ["limpieza-hogar"]},
    {"name": "Fragancia",          "slug": "fragancia",          "values": ["Sin fragancia","Lavanda","Limón","Pino","Floral"], "applies_to_categories": ["limpieza-hogar","desodorantes-ambientales"]}
  ]'::jsonb,
  true, true, 'manual-curation', '1.0.0-curated'
FROM business_types bt WHERE bt.code = 'limpieza'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories,
  brands     = EXCLUDED.brands,
  products   = EXCLUDED.products,
  attributes = EXCLUDED.attributes,
  version    = EXCLUDED.version,
  generated_by = EXCLUDED.generated_by,
  updated_at = CURRENT_TIMESTAMP;

-- RESUMEN: 27 productos en template limpieza
--   limpieza-hogar:          10 productos
--   liquidos-granel:          7 productos
--   desodorantes-ambientales: 2 productos
--   insumos-pileta:           2 productos
--   accesorios-limpieza:      5 productos
--   jardin-exterior:          1 producto
