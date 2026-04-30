-- Seed 074: Template curado para Piletas y Jardín (business_type code='piletas')
-- PROPÓSITO: Productos ancla, marcas y atributos curados para comercios de piletas y jardín argentinos
-- IDEMPOTENTE: INSERT/UPDATE business_type; INSERT/UPDATE template
-- REQUIERE: 001 (business_types), schema business_type_templates
-- CICLO: cycle-004-brands-catalog-expansion
-- FECHA: 2026-04-22
-- VERSION: 1.0.0-curated | generated_by: manual-curation

-- Paso 1: Asegurar que existe business_type piletas
INSERT INTO business_types (code, name, description, icon, color, sort_order, is_active)
VALUES (
  'piletas',
  'Piletas y Jardín',
  'Mantenimiento de piletas, productos químicos, bombas, filtros, accesorios de jardín y artículos de exterior',
  'waves',
  '#0288D1',
  46,
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
SELECT bt.id, 'Piletas y Jardín', 'Template curado para comercios de piletas y jardín', 'AR',
  '[
    {"slug": "cloro-quimicos",          "name": "Cloro y Químicos"},
    {"slug": "bombas-filtros-pileta",   "name": "Bombas y Filtros"},
    {"slug": "accesorios-pileta",       "name": "Accesorios de Pileta"},
    {"slug": "piletas-estructurales",   "name": "Piletas"},
    {"slug": "jardin-exterior",         "name": "Jardín y Exterior"},
    {"slug": "limpieza-hogar",          "name": "Limpieza del Hogar"}
  ]'::jsonb,
  '[
    {"name": "Clorotec",    "suggested_for_categories": ["cloro-quimicos"]},
    {"name": "Nataclor",    "suggested_for_categories": ["cloro-quimicos"]},
    {"name": "Deep Blue",   "suggested_for_categories": ["cloro-quimicos","accesorios-pileta"]},
    {"name": "Bestway",     "suggested_for_categories": ["piletas-estructurales","accesorios-pileta"]},
    {"name": "Maytronics",  "suggested_for_categories": ["bombas-filtros-pileta"]},
    {"name": "Freshclor",   "suggested_for_categories": ["cloro-quimicos"]},
    {"name": "Kokido",      "suggested_for_categories": ["accesorios-pileta"]},
    {"name": "Pool Xpert",  "suggested_for_categories": ["cloro-quimicos","accesorios-pileta"]},
    {"name": "Tramontina",  "suggested_for_categories": ["jardin-exterior"]},
    {"name": "3M",          "suggested_for_categories": ["limpieza-hogar"]}
  ]'::jsonb,
  '[
    {"name": "Cloro granulado Clorotec 10kg",         "category_slug": "cloro-quimicos",        "brand": "Clorotec",    "price_reference": 65000,  "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Cloro granulado Nataclor 10kg",          "category_slug": "cloro-quimicos",        "brand": "Nataclor",    "price_reference": 68000,  "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Cloro tabletas Clorotec 1kg",            "category_slug": "cloro-quimicos",        "brand": "Clorotec",    "price_reference": 12000,  "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Algicida Nataclor 1L",                   "category_slug": "cloro-quimicos",        "brand": "Nataclor",    "price_reference": 8500,   "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "pH Minus 1kg",                           "category_slug": "cloro-quimicos",        "brand": null,          "price_reference": 7500,   "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "pH Plus 1kg",                            "category_slug": "cloro-quimicos",        "brand": null,          "price_reference": 7500,   "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Kit test agua",                          "category_slug": "cloro-quimicos",        "brand": null,          "price_reference": 3500,   "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Cloro líquido 5L",                       "category_slug": "cloro-quimicos",        "brand": null,          "price_reference": 4500,   "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Bomba recirculación 1/2HP",              "category_slug": "bombas-filtros-pileta", "brand": null,          "price_reference": 95000,  "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Bomba recirculación 3/4HP",              "category_slug": "bombas-filtros-pileta", "brand": null,          "price_reference": 140000, "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Filtro de arena 11\"",                   "category_slug": "bombas-filtros-pileta", "brand": null,          "price_reference": 75000,  "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Filtro de arena 16\"",                   "category_slug": "bombas-filtros-pileta", "brand": null,          "price_reference": 120000, "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Robot limpiafondos Maytronics",          "category_slug": "bombas-filtros-pileta", "brand": "Maytronics",  "price_reference": 450000, "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Cepillo de fondo",                       "category_slug": "accesorios-pileta",     "brand": null,          "price_reference": 3500,   "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Manga aspiradora 7m",                    "category_slug": "accesorios-pileta",     "brand": null,          "price_reference": 8000,   "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Skimmer flotante",                       "category_slug": "accesorios-pileta",     "brand": null,          "price_reference": 12000,  "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Termómetro digital pileta",              "category_slug": "accesorios-pileta",     "brand": null,          "price_reference": 2500,   "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Cubierta solar 4x3m",                    "category_slug": "accesorios-pileta",     "brand": null,          "price_reference": 25000,  "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Pileta estructural Bestway 3.66m",       "category_slug": "piletas-estructurales", "brand": "Bestway",     "price_reference": 195000, "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Pileta estructural Bestway 4.88m",       "category_slug": "piletas-estructurales", "brand": "Bestway",     "price_reference": 320000, "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Pileta inflable familiar 2.44m",         "category_slug": "piletas-estructurales", "brand": null,          "price_reference": 45000,  "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Manguera jardín 1/2\" 25m",              "category_slug": "jardin-exterior",       "brand": null,          "price_reference": 18000,  "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Aspersor jardín",                        "category_slug": "jardin-exterior",       "brand": null,          "price_reference": 4500,   "unit": "unidad", "sku_prefix": "PILETA"},
    {"name": "Maceta 40cm",                            "category_slug": "jardin-exterior",       "brand": null,          "price_reference": 3500,   "unit": "unidad", "sku_prefix": "PILETA"}
  ]'::jsonb,
  '[
    {"name": "Tipo Químico",       "slug": "tipo-quimico",       "values": ["Cloro granulado","Cloro tabletas","Cloro líquido","Algicida","pH Minus","pH Plus","Floculante"],  "applies_to_categories": ["cloro-quimicos"]},
    {"name": "Presentación Pileta","slug": "presentacion-pileta","values": ["Redonda","Rectangular","Ovalada"],                                                                 "applies_to_categories": ["piletas-estructurales"]},
    {"name": "Potencia Bomba",     "slug": "potencia-bomba",     "values": ["1/3 HP","1/2 HP","3/4 HP","1 HP"],                                                                "applies_to_categories": ["bombas-filtros-pileta"]}
  ]'::jsonb,
  true, true, 'manual-curation', '1.0.0-curated'
FROM business_types bt WHERE bt.code = 'piletas'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories,
  brands     = EXCLUDED.brands,
  products   = EXCLUDED.products,
  attributes = EXCLUDED.attributes,
  version    = EXCLUDED.version,
  generated_by = EXCLUDED.generated_by,
  updated_at = CURRENT_TIMESTAMP;

-- RESUMEN: 24 productos en template piletas
--   cloro-quimicos:        8 productos
--   bombas-filtros-pileta: 5 productos
--   accesorios-pileta:     5 productos
--   piletas-estructurales: 3 productos
--   jardin-exterior:       3 productos
--   limpieza-hogar:        0 productos (categoría disponible para el comercio)
