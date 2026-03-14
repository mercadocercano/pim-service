-- Seed 023: Template curado para Ferretería (business_type code='ferreteria')
-- PROPÓSITO: UPSERT del template default con categorías, marcas, productos y atributos curados para ferreterías argentinas
-- IDEMPOTENTE: Usa UPDATE con WHERE; no afecta si el template no existe (se ejecuta después de 004)
-- REQUIERE: 001 (business_types), 004 (business_type_quickstart_templates), 020 (ferreteria marketplace_categories)
-- VERSION: 3.0.0-curated | generated_by: manual-curation

UPDATE business_type_templates
SET
  categories = '[
    {"slug": "materiales-construccion", "name": "Materiales de Construcción"},
    {"slug": "herramientas-manuales", "name": "Herramientas Manuales"},
    {"slug": "herramientas-electricas", "name": "Herramientas Eléctricas"},
    {"slug": "tornilleria", "name": "Tornillería y Fijaciones"},
    {"slug": "plomeria", "name": "Plomería y Sanitarios"},
    {"slug": "electricidad", "name": "Electricidad"},
    {"slug": "pinturas", "name": "Pinturas y Accesorios"},
    {"slug": "hogar-jardin", "name": "Ferretería General"}
  ]'::jsonb,
  brands = '[
    {"name": "Loma Negra", "suggested_for_categories": ["materiales-construccion"]},
    {"name": "Avellaneda", "suggested_for_categories": ["materiales-construccion"]},
    {"name": "Stanley", "suggested_for_categories": ["herramientas-manuales"]},
    {"name": "Bosch", "suggested_for_categories": ["herramientas-electricas"]},
    {"name": "Black+Decker", "suggested_for_categories": ["herramientas-electricas"]},
    {"name": "Acindar", "suggested_for_categories": ["materiales-construccion"]},
    {"name": "Tigre", "suggested_for_categories": ["plomeria"]},
    {"name": "FV", "suggested_for_categories": ["plomeria"]},
    {"name": "Alba", "suggested_for_categories": ["pinturas"]},
    {"name": "Sinteplast", "suggested_for_categories": ["pinturas"]},
    {"name": "Makita", "suggested_for_categories": ["herramientas-electricas"]},
    {"name": "Tramontina", "suggested_for_categories": ["herramientas-manuales"]}
  ]'::jsonb,
  products = '[
    {"name": "Cemento Portland Loma Negra 50kg", "category_slug": "materiales-construccion", "brand": "Loma Negra", "price_reference": 9500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Cemento Avellaneda 25kg", "category_slug": "materiales-construccion", "brand": "Avellaneda", "price_reference": 5600, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Cal hidráulica 25kg", "category_slug": "materiales-construccion", "brand": "Loma Negra", "price_reference": 3500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Arena bolsa 25kg", "category_slug": "materiales-construccion", "brand": null, "price_reference": 1500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Ladrillo hueco 12x18x33", "category_slug": "materiales-construccion", "brand": null, "price_reference": 180, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Bloque de hormigón", "category_slug": "materiales-construccion", "brand": null, "price_reference": 250, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Hierro ø8mm x 12m Acindar", "category_slug": "materiales-construccion", "brand": "Acindar", "price_reference": 8500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Hierro ø10mm x 12m", "category_slug": "materiales-construccion", "brand": "Acindar", "price_reference": 13000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Membrana asfáltica 4mm x 10m²", "category_slug": "materiales-construccion", "brand": null, "price_reference": 35000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Malla SIMA 15x15", "category_slug": "materiales-construccion", "brand": null, "price_reference": 15000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Yeso fino 25kg", "category_slug": "materiales-construccion", "brand": "Loma Negra", "price_reference": 4200, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Varilla acero ø6mm x 12m", "category_slug": "materiales-construccion", "brand": "Acindar", "price_reference": 4500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Martillo carpintero Stanley 500g", "category_slug": "herramientas-manuales", "brand": "Stanley", "price_reference": 12000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Destornillador Phillips Stanley", "category_slug": "herramientas-manuales", "brand": "Stanley", "price_reference": 3500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Llave francesa 10\"", "category_slug": "herramientas-manuales", "brand": "Stanley", "price_reference": 8000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Pinza universal", "category_slug": "herramientas-manuales", "brand": "Tramontina", "price_reference": 4500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Sierra manual", "category_slug": "herramientas-manuales", "brand": "Tramontina", "price_reference": 6000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Cinta métrica 5m", "category_slug": "herramientas-manuales", "brand": "Stanley", "price_reference": 3000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Nivel 40cm", "category_slug": "herramientas-manuales", "brand": "Stanley", "price_reference": 5000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Escuadra carpintera 30cm", "category_slug": "herramientas-manuales", "brand": "Stanley", "price_reference": 3500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Serrucho de mano", "category_slug": "herramientas-manuales", "brand": "Tramontina", "price_reference": 4500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Taladro percutor Bosch 700W", "category_slug": "herramientas-electricas", "brand": "Bosch", "price_reference": 45000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Amoladora angular 4½\" Bosch", "category_slug": "herramientas-electricas", "brand": "Bosch", "price_reference": 38000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Sierra caladora Black+Decker", "category_slug": "herramientas-electricas", "brand": "Black+Decker", "price_reference": 32000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Atornillador inalámbrico Makita", "category_slug": "herramientas-electricas", "brand": "Makita", "price_reference": 55000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Mecha para cemento 10mm", "category_slug": "herramientas-electricas", "brand": "Bosch", "price_reference": 2500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Disco abrasivo 4½\"", "category_slug": "herramientas-electricas", "brand": "Bosch", "price_reference": 1500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Tornillo autoperforante 8x1 (x100)", "category_slug": "tornilleria", "brand": null, "price_reference": 1500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Clavo 2\" (x1kg)", "category_slug": "tornilleria", "brand": null, "price_reference": 2000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Tarugo 8mm (x100)", "category_slug": "tornilleria", "brand": null, "price_reference": 800, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Tirafondo 6x50 (x50)", "category_slug": "tornilleria", "brand": null, "price_reference": 1200, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Caño PVC 110mm x 4m", "category_slug": "plomeria", "brand": "Tigre", "price_reference": 8000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Codo PVC 110mm", "category_slug": "plomeria", "brand": "Tigre", "price_reference": 600, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Grifería cocina FV", "category_slug": "plomeria", "brand": "FV", "price_reference": 25000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Flotante para tanque", "category_slug": "plomeria", "brand": "FV", "price_reference": 2500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Rosca galvanizada 1\" x 6m", "category_slug": "plomeria", "brand": "Tigre", "price_reference": 5500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Caño PBC ¾\" x 3m", "category_slug": "plomeria", "brand": "Tigre", "price_reference": 2500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Cable 2.5mm² x 100m", "category_slug": "electricidad", "brand": null, "price_reference": 18000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Llave térmica 20A", "category_slug": "electricidad", "brand": null, "price_reference": 5000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Tomacorriente doble", "category_slug": "electricidad", "brand": null, "price_reference": 1200, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Cinta aisladora", "category_slug": "electricidad", "brand": null, "price_reference": 500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Interruptor simples", "category_slug": "electricidad", "brand": null, "price_reference": 800, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Zapatilla 3 bocas", "category_slug": "electricidad", "brand": null, "price_reference": 3500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Pintura látex interior blanco 4L Alba", "category_slug": "pinturas", "brand": "Alba", "price_reference": 15000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Esmalte sintético 1L", "category_slug": "pinturas", "brand": "Sinteplast", "price_reference": 8000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Enduido plástico 4kg", "category_slug": "pinturas", "brand": "Alba", "price_reference": 4500, "unit": "unidad", "sku_prefix": "FERRET"}
  ]'::jsonb,
  attributes = '[
    {"name": "Presentación", "slug": "presentacion-pintura", "values": ["1L","4L","10L","20L"], "applies_to_categories": ["pinturas"]},
    {"name": "Material", "slug": "material-hogar", "values": ["Acero","PVC","Bronce","Aluminio","Hierro","Madera"], "applies_to_categories": ["tornilleria","plomeria","herramientas-manuales"]},
    {"name": "Tamaño", "slug": "tamano", "values": ["Pequeño","Mediano","Grande","Extra Grande"], "applies_to_categories": ["materiales-construccion","hogar-jardin"]}
  ]'::jsonb,
  version = '3.0.0-curated',
  generated_by = 'manual-curation',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'ferreteria')
  AND is_default = true;
