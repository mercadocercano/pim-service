-- Seed 027: Template curado para Deportes (business_type code='deportes')
-- PROPÓSITO: Productos ancla, marcas y atributos curados para tiendas deportivas argentinas
-- IDEMPOTENTE: UPDATE solo afecta template existente
-- REQUIERE: 001 (business_types), 004 (business_type_quickstart_templates)
-- VERSION: 3.0.0-curated | generated_by: manual-curation

UPDATE business_type_templates
SET
  categories = '[
    {"slug": "ropa-deportiva", "name": "Ropa Deportiva"},
    {"slug": "calzado-deportivo", "name": "Calzado Deportivo"},
    {"slug": "accesorios-deportivos", "name": "Accesorios Deportivos"},
    {"slug": "equipamiento", "name": "Equipamiento"},
    {"slug": "indumentaria-fitness", "name": "Indumentaria Fitness"}
  ]'::jsonb,
  brands = '[
    {"name": "Nike", "suggested_for_categories": ["ropa-deportiva","calzado-deportivo","indumentaria-fitness"]},
    {"name": "Adidas", "suggested_for_categories": ["ropa-deportiva","calzado-deportivo","indumentaria-fitness"]},
    {"name": "Topper", "suggested_for_categories": ["ropa-deportiva","calzado-deportivo"]},
    {"name": "Puma", "suggested_for_categories": ["ropa-deportiva","calzado-deportivo"]},
    {"name": "Under Armour", "suggested_for_categories": ["ropa-deportiva","indumentaria-fitness"]},
    {"name": "DRB", "suggested_for_categories": ["equipamiento"]},
    {"name": "Nassau", "suggested_for_categories": ["equipamiento"]}
  ]'::jsonb,
  products = '[
    {"name": "Camiseta deportiva dry-fit", "category_slug": "ropa-deportiva", "brand": "Nike", "price_reference": 18000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Short deportivo", "category_slug": "ropa-deportiva", "brand": null, "price_reference": 12000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Calza running", "category_slug": "indumentaria-fitness", "brand": "Adidas", "price_reference": 22000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Zapatilla running", "category_slug": "calzado-deportivo", "brand": "Nike", "price_reference": 55000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Zapatilla fútbol/botín", "category_slug": "calzado-deportivo", "brand": "Adidas", "price_reference": 45000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Zapatilla training", "category_slug": "calzado-deportivo", "brand": "Topper", "price_reference": 40000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Pelota fútbol Nassau", "category_slug": "equipamiento", "brand": "Nassau", "price_reference": 15000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Pelota basketball", "category_slug": "equipamiento", "brand": null, "price_reference": 12000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Bolso deportivo", "category_slug": "accesorios-deportivos", "brand": null, "price_reference": 20000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Botella deportiva 750ml", "category_slug": "accesorios-deportivos", "brand": null, "price_reference": 5000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Toalla deportiva", "category_slug": "accesorios-deportivos", "brand": null, "price_reference": 8000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Remera térmica manga larga", "category_slug": "indumentaria-fitness", "brand": "Under Armour", "price_reference": 25000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Campera deportiva", "category_slug": "ropa-deportiva", "brand": "Puma", "price_reference": 48000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Conjunto deportivo", "category_slug": "indumentaria-fitness", "brand": null, "price_reference": 50000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Tobillera/Rodillera", "category_slug": "accesorios-deportivos", "brand": null, "price_reference": 6000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Soga para saltar", "category_slug": "equipamiento", "brand": null, "price_reference": 4000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Banda elástica", "category_slug": "equipamiento", "brand": null, "price_reference": 3000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Guantes de arquero", "category_slug": "equipamiento", "brand": null, "price_reference": 12000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Canillera", "category_slug": "accesorios-deportivos", "brand": null, "price_reference": 5000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Medias deportivas x3", "category_slug": "accesorios-deportivos", "brand": null, "price_reference": 6000, "unit": "unidad", "sku_prefix": "DEPORT"}
  ]'::jsonb,
  attributes = '[
    {"name": "Talle", "slug": "talle-ropa", "values": ["XS","S","M","L","XL","XXL"], "applies_to_categories": ["ropa-deportiva","indumentaria-fitness"]},
    {"name": "Color", "slug": "color", "values": ["Negro","Blanco","Azul","Rojo","Verde","Gris","Amarillo","Rosa"], "applies_to_categories": ["ropa-deportiva","calzado-deportivo","accesorios-deportivos","indumentaria-fitness"]},
    {"name": "Disciplina", "slug": "disciplina-deportiva", "values": ["Fútbol","Running","Basket","Gym","Tenis","Vóley","Natación"], "applies_to_categories": ["ropa-deportiva","calzado-deportivo","equipamiento","indumentaria-fitness"]},
    {"name": "Talle Calzado", "slug": "talle-calzado", "values": ["36","37","38","39","40","41","42","43","44","45"], "applies_to_categories": ["calzado-deportivo"]}
  ]'::jsonb,
  version = '3.0.0-curated',
  generated_by = 'manual-curation',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'deportes')
  AND is_default = true;
