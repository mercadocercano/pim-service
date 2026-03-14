-- Seed 026: Template curado para Indumentaria (business_type code='ropa')
-- PROPÓSITO: Productos ancla, marcas y atributos curados para tiendas de ropa argentinas
-- IDEMPOTENTE: UPDATE solo afecta template existente
-- REQUIERE: 001 (business_types), 004 (business_type_quickstart_templates)
-- VERSION: 3.0.0-curated | generated_by: manual-curation

UPDATE business_type_templates
SET
  categories = '[
    {"slug": "remeras", "name": "Remeras"},
    {"slug": "pantalones", "name": "Pantalones"},
    {"slug": "camperas-abrigos", "name": "Camperas/Abrigos"},
    {"slug": "buzos-sweaters", "name": "Buzos/Sweaters"},
    {"slug": "vestidos-faldas", "name": "Vestidos/Faldas"},
    {"slug": "accesorios", "name": "Accesorios"}
  ]'::jsonb,
  brands = '[
    {"name": "Nike", "suggested_for_categories": ["remeras","pantalones","buzos-sweaters"]},
    {"name": "Adidas", "suggested_for_categories": ["remeras","pantalones","buzos-sweaters"]},
    {"name": "Topper", "suggested_for_categories": ["remeras","pantalones","accesorios"]},
    {"name": "Puma", "suggested_for_categories": ["remeras","pantalones","buzos-sweaters"]},
    {"name": "Levi''s", "suggested_for_categories": ["pantalones","camperas-abrigos"]},
    {"name": "Wrangler", "suggested_for_categories": ["pantalones"]},
    {"name": "Narrow", "suggested_for_categories": ["remeras","pantalones","camperas-abrigos"]},
    {"name": "Kevingston", "suggested_for_categories": ["remeras","pantalones","vestidos-faldas"]}
  ]'::jsonb,
  products = '[
    {"name": "Remera básica algodón", "category_slug": "remeras", "brand": null, "price_reference": 15000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Remera estampada", "category_slug": "remeras", "brand": null, "price_reference": 18000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Remera deportiva", "category_slug": "remeras", "brand": "Nike", "price_reference": 22000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Jean clásico recto", "category_slug": "pantalones", "brand": "Levi''s", "price_reference": 35000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Jean slim fit", "category_slug": "pantalones", "brand": "Wrangler", "price_reference": 38000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Pantalón cargo", "category_slug": "pantalones", "brand": null, "price_reference": 30000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Pantalón jogger", "category_slug": "pantalones", "brand": null, "price_reference": 25000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Campera rompeviento", "category_slug": "camperas-abrigos", "brand": null, "price_reference": 42000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Campera de abrigo", "category_slug": "camperas-abrigos", "brand": null, "price_reference": 55000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Campera de jean", "category_slug": "camperas-abrigos", "brand": "Levi''s", "price_reference": 48000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Buzo canguro frisa", "category_slug": "buzos-sweaters", "brand": null, "price_reference": 28000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Sweater cuello redondo", "category_slug": "buzos-sweaters", "brand": null, "price_reference": 25000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Hoodie con capucha", "category_slug": "buzos-sweaters", "brand": "Adidas", "price_reference": 32000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Vestido casual", "category_slug": "vestidos-faldas", "brand": null, "price_reference": 30000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Falda midi", "category_slug": "vestidos-faldas", "brand": null, "price_reference": 22000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Bermuda/Short", "category_slug": "pantalones", "brand": null, "price_reference": 18000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Camisa manga larga", "category_slug": "remeras", "brand": null, "price_reference": 28000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Camisa manga corta", "category_slug": "remeras", "brand": null, "price_reference": 24000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Zapatilla urbana", "category_slug": "accesorios", "brand": "Nike", "price_reference": 35000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Zapatilla deportiva", "category_slug": "accesorios", "brand": "Adidas", "price_reference": 45000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Cinturón", "category_slug": "accesorios", "brand": null, "price_reference": 8000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Gorra", "category_slug": "accesorios", "brand": null, "price_reference": 6000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Bufanda", "category_slug": "accesorios", "brand": null, "price_reference": 10000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Medias pack x3", "category_slug": "accesorios", "brand": null, "price_reference": 5000, "unit": "unidad", "sku_prefix": "INDUM"}
  ]'::jsonb,
  attributes = '[
    {"name": "Talle", "slug": "talle-ropa", "values": ["XS","S","M","L","XL","XXL"], "applies_to_categories": ["remeras","pantalones","camperas-abrigos","buzos-sweaters","vestidos-faldas"]},
    {"name": "Color", "slug": "color", "values": ["Negro","Blanco","Azul","Gris","Rojo","Verde","Beige","Rosa","Celeste"], "applies_to_categories": ["remeras","pantalones","camperas-abrigos","buzos-sweaters","vestidos-faldas","accesorios"]},
    {"name": "Material", "slug": "material-textil", "values": ["Algodón","Poliéster","Jean","Frisa","Lana","Sintético"], "applies_to_categories": ["remeras","pantalones","camperas-abrigos","buzos-sweaters","vestidos-faldas"]},
    {"name": "Género", "slug": "genero", "values": ["Hombre","Mujer","Unisex","Niño","Niña"], "applies_to_categories": ["remeras","pantalones","camperas-abrigos","buzos-sweaters","vestidos-faldas"]}
  ]'::jsonb,
  version = '3.0.0-curated',
  generated_by = 'manual-curation',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'ropa')
  AND is_default = true;
