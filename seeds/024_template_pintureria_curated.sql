-- Seed 024: Template curado para Pinturería (business_type code='pintureria')
-- PROPÓSITO: Productos ancla, marcas y atributos curados para pinturerías argentinas
-- IDEMPOTENTE: INSERT business_type si no existe; INSERT/UPDATE template
-- REQUIERE: 001 (business_types), 004 (business_type_quickstart_templates)
-- VERSION: 3.0.0-curated | generated_by: manual-curation

-- Paso 1: Asegurar que existe business_type pintureria
INSERT INTO business_types (code, name, description, icon, color, sort_order, is_active)
VALUES ('pintureria', 'Pinturería', 'Venta de pinturas, esmaltes, accesorios y materiales para pintar', 'paint-bucket', '#FF6B35', 45, true)
ON CONFLICT (code) DO NOTHING;

-- Paso 2: INSERT o UPDATE del template default
INSERT INTO business_type_templates (
  business_type_id, name, description, region, categories, brands, products, attributes,
  is_default, is_active, generated_by, version
)
SELECT bt.id, 'Pinturería', 'Template curado para pinturerías', 'AR',
  '[
    {"slug": "latex-interior", "name": "Látex Interior"},
    {"slug": "latex-exterior", "name": "Látex Exterior"},
    {"slug": "esmaltes-sinteticos", "name": "Esmaltes Sintéticos"},
    {"slug": "enduidos-masillas", "name": "Enduidos/Masillas"},
    {"slug": "accesorios-pintura", "name": "Accesorios de Pintura"}
  ]'::jsonb,
  '[
    {"name": "Alba", "suggested_for_categories": ["latex-interior","latex-exterior","esmaltes-sinteticos","enduidos-masillas"]},
    {"name": "Sinteplast", "suggested_for_categories": ["latex-interior","latex-exterior","enduidos-masillas"]},
    {"name": "Sherwin Williams", "suggested_for_categories": ["latex-interior","latex-exterior","esmaltes-sinteticos"]},
    {"name": "Tersuave", "suggested_for_categories": ["latex-interior","latex-exterior"]},
    {"name": "Colorín", "suggested_for_categories": ["esmaltes-sinteticos","accesorios-pintura"]},
    {"name": "Riopint", "suggested_for_categories": ["esmaltes-sinteticos","latex-exterior"]}
  ]'::jsonb,
  '[
    {"name": "Alba Albalatex Interior Mate 4L", "category_slug": "latex-interior", "brand": "Alba", "price_reference": 15000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Alba Albalatex Interior Mate 20L", "category_slug": "latex-interior", "brand": "Alba", "price_reference": 55000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Sinteplast Recuplast Interior 4L", "category_slug": "latex-interior", "brand": "Sinteplast", "price_reference": 12000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Sinteplast Recuplast Interior 20L", "category_slug": "latex-interior", "brand": "Sinteplast", "price_reference": 48000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Sherwin Williams Z10 4L", "category_slug": "latex-interior", "brand": "Sherwin Williams", "price_reference": 22000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Tersuave Látex Exterior 4L", "category_slug": "latex-exterior", "brand": "Tersuave", "price_reference": 14000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Albalux Esmalte Sintético Blanco 1L", "category_slug": "esmaltes-sinteticos", "brand": "Alba", "price_reference": 8000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Albalux Esmalte Sintético Blanco 4L", "category_slug": "esmaltes-sinteticos", "brand": "Alba", "price_reference": 25000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Riopint Riolux Esmalte Negro 4L", "category_slug": "esmaltes-sinteticos", "brand": "Riopint", "price_reference": 18000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Colorín Barniz 1L", "category_slug": "esmaltes-sinteticos", "brand": "Colorín", "price_reference": 7000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Albaplast Enduido Interior 4L", "category_slug": "enduidos-masillas", "brand": "Alba", "price_reference": 8000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Enduido Sinteplast 4L", "category_slug": "enduidos-masillas", "brand": "Sinteplast", "price_reference": 6000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Membrana Líquida Sika 20L", "category_slug": "enduidos-masillas", "brand": null, "price_reference": 45000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Fijador/Sellador 4L", "category_slug": "enduidos-masillas", "brand": null, "price_reference": 6000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Pincel 2\"", "category_slug": "accesorios-pintura", "brand": null, "price_reference": 1500, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Rodillo lana 23cm", "category_slug": "accesorios-pintura", "brand": null, "price_reference": 3500, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Bandeja para rodillo", "category_slug": "accesorios-pintura", "brand": null, "price_reference": 1200, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Cinta de enmascarar 18mm", "category_slug": "accesorios-pintura", "brand": null, "price_reference": 800, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Lija al agua 220", "category_slug": "accesorios-pintura", "brand": null, "price_reference": 300, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Diluyente 1L", "category_slug": "accesorios-pintura", "brand": null, "price_reference": 3000, "unit": "unidad", "sku_prefix": "PINTUR"}
  ]'::jsonb,
  '[
    {"name": "Presentación Pintura", "slug": "presentacion-pintura", "values": ["1L","4L","10L","20L"], "applies_to_categories": ["latex-interior","latex-exterior","esmaltes-sinteticos"]},
    {"name": "Acabado", "slug": "acabado", "values": ["Mate","Satinado","Brillante"], "applies_to_categories": ["latex-interior","latex-exterior","esmaltes-sinteticos"]},
    {"name": "Tipo Pintura", "slug": "tipo-pintura", "values": ["Látex","Esmalte","Barniz","Enduido"], "applies_to_categories": ["latex-interior","latex-exterior","esmaltes-sinteticos","enduidos-masillas"]},
    {"name": "Color", "slug": "color", "values": ["Blanco","Negro","Gris","Beige","Celeste","Verde","Rojo","Amarillo"], "applies_to_categories": ["latex-interior","latex-exterior","esmaltes-sinteticos"]}
  ]'::jsonb,
  true, true, 'manual-curation', '3.0.0-curated'
FROM business_types bt WHERE bt.code = 'pintureria'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories,
  brands = EXCLUDED.brands,
  products = EXCLUDED.products,
  attributes = EXCLUDED.attributes,
  version = EXCLUDED.version,
  generated_by = EXCLUDED.generated_by,
  updated_at = CURRENT_TIMESTAMP;
