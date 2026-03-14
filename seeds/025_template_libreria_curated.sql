-- Seed 025: Template curado para Librería (business_type code='libreria')
-- PROPÓSITO: Productos ancla, marcas y atributos curados para librerías argentinas
-- IDEMPOTENTE: UPDATE solo afecta template existente
-- REQUIERE: 001 (business_types), 004 (business_type_quickstart_templates)
-- VERSION: 3.0.0-curated | generated_by: manual-curation

UPDATE business_type_templates
SET
  categories = '[
    {"slug": "cuadernos", "name": "Cuadernos"},
    {"slug": "escritura", "name": "Escritura"},
    {"slug": "utiles-escolares", "name": "Útiles Escolares"},
    {"slug": "papeleria-oficina", "name": "Papelería/Oficina"},
    {"slug": "mochilas-cartucheras", "name": "Mochilas/Cartucheras"},
    {"slug": "arte-dibujo", "name": "Arte/Dibujo"}
  ]'::jsonb,
  brands = '[
    {"name": "Rivadavia", "suggested_for_categories": ["cuadernos"]},
    {"name": "Faber Castell", "suggested_for_categories": ["escritura","utiles-escolares","arte-dibujo"]},
    {"name": "Maped", "suggested_for_categories": ["utiles-escolares","arte-dibujo"]},
    {"name": "BIC", "suggested_for_categories": ["escritura"]},
    {"name": "Filgo", "suggested_for_categories": ["escritura"]},
    {"name": "Éxito", "suggested_for_categories": ["utiles-escolares"]},
    {"name": "Ledesma", "suggested_for_categories": ["papeleria-oficina"]},
    {"name": "Mooving", "suggested_for_categories": ["mochilas-cartucheras"]}
  ]'::jsonb,
  products = '[
    {"name": "Cuaderno Rivadavia tapa dura 48 hojas rayado", "category_slug": "cuadernos", "brand": "Rivadavia", "price_reference": 7600, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Cuaderno Rivadavia tapa dura 98 hojas", "category_slug": "cuadernos", "brand": "Rivadavia", "price_reference": 9800, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Cuaderno espiralado A4 84 hojas", "category_slug": "cuadernos", "brand": null, "price_reference": 5500, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Lápiz grafito Faber Castell HB", "category_slug": "escritura", "brand": "Faber Castell", "price_reference": 500, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Lápices de colores Faber Castell x12", "category_slug": "escritura", "brand": "Faber Castell", "price_reference": 5000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Lápices de colores Faber Castell x24", "category_slug": "escritura", "brand": "Faber Castell", "price_reference": 9000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Marcadores Maped Color Peps x12", "category_slug": "escritura", "brand": "Maped", "price_reference": 6000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Resaltador Filgo x4", "category_slug": "escritura", "brand": "Filgo", "price_reference": 3500, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Bolígrafo BIC Cristal azul", "category_slug": "escritura", "brand": "BIC", "price_reference": 400, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Bolígrafo BIC Cristal rojo", "category_slug": "escritura", "brand": "BIC", "price_reference": 400, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Roller borrable Frixion", "category_slug": "escritura", "brand": "Pilot", "price_reference": 2500, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Goma de borrar Maped", "category_slug": "utiles-escolares", "brand": "Maped", "price_reference": 500, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Sacapuntas metálico", "category_slug": "utiles-escolares", "brand": null, "price_reference": 300, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Regla 30cm Maped", "category_slug": "utiles-escolares", "brand": "Maped", "price_reference": 800, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Tijera escolar Maped 13cm", "category_slug": "utiles-escolares", "brand": "Maped", "price_reference": 2000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Plasticola 40g", "category_slug": "utiles-escolares", "brand": null, "price_reference": 800, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Cinta adhesiva", "category_slug": "utiles-escolares", "brand": null, "price_reference": 600, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Cartuchera simple", "category_slug": "mochilas-cartucheras", "brand": null, "price_reference": 4000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Mochila escolar", "category_slug": "mochilas-cartucheras", "brand": null, "price_reference": 25000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Resma A4 500 hojas Ledesma", "category_slug": "papeleria-oficina", "brand": "Ledesma", "price_reference": 8000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Carpeta 3 anillos", "category_slug": "papeleria-oficina", "brand": null, "price_reference": 4000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Repuesto hojas A4 x480", "category_slug": "papeleria-oficina", "brand": null, "price_reference": 5000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Pinturitas Faber Castell x12", "category_slug": "arte-dibujo", "brand": "Faber Castell", "price_reference": 4000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Block de dibujo A4", "category_slug": "arte-dibujo", "brand": null, "price_reference": 2000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Témperas x6", "category_slug": "arte-dibujo", "brand": null, "price_reference": 3500, "unit": "unidad", "sku_prefix": "LIBRE"}
  ]'::jsonb,
  attributes = '[
    {"name": "Formato Cuaderno", "slug": "formato-cuaderno", "values": ["A4","A5","Letter","Oficio"], "applies_to_categories": ["cuadernos"]},
    {"name": "Cantidad de Hojas", "slug": "cantidad-hojas", "values": ["48","72","84","98","120","200"], "applies_to_categories": ["cuadernos","papeleria-oficina"]},
    {"name": "Tipo de Rayado", "slug": "tipo-rayado", "values": ["Rayado","Cuadriculado","Liso","Pautado"], "applies_to_categories": ["cuadernos"]},
    {"name": "Color", "slug": "color", "values": ["Negro","Azul","Rojo","Verde","Rosa","Amarillo"], "applies_to_categories": ["cuadernos","escritura","utiles-escolares","mochilas-cartucheras"]}
  ]'::jsonb,
  version = '3.0.0-curated',
  generated_by = 'manual-curation',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'libreria')
  AND is_default = true;
