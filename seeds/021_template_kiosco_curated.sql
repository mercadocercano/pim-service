-- Seed 021: Template curado para Kiosco (business_type code='kiosco')
-- PROPÓSITO: UPSERT del template default con categorías, marcas, productos y atributos curados para kioscos argentinos
-- IDEMPOTENTE: Usa UPDATE con WHERE; no afecta si el template no existe (se ejecuta después de 004)
-- REQUIERE: 001 (business_types), 011 (marketplace_categories), 016 (marketplace_attributes alimentos)
-- VERSION: 3.0.0-curated | generated_by: manual-curation

-- =====================================================
-- PASO 1: Asegurar que existe categoría 'general' (si no existe)
-- =====================================================
INSERT INTO marketplace_categories (name, slug, description, parent_id, sort_order, is_active)
VALUES ('Varios / General', 'general', 'Productos varios y miscelánea', NULL, 99, true)
ON CONFLICT (slug) DO UPDATE SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  is_active = EXCLUDED.is_active;

-- =====================================================
-- PASO 2: Construir categorías con IDs desde marketplace_categories
-- =====================================================
DO $$
DECLARE
  v_categories JSONB := '[]'::jsonb;
  v_category_slugs TEXT[] := ARRAY[
    'bebidas-sin-alcohol', 'golosinas', 'snacks-salados', 'galletas',
    'bebidas-con-alcohol', 'general'
  ];
  v_category_names TEXT[] := ARRAY[
    'Bebidas', 'Golosinas', 'Snacks', 'Galletitas',
    'Bebidas con Alcohol', 'Varios Kiosco'
  ];
  v_cat_id UUID;
  v_cat_slug TEXT;
  v_cat_name TEXT;
  i INT;
BEGIN
  FOR i IN 1..array_length(v_category_slugs, 1) LOOP
    v_cat_slug := v_category_slugs[i];
    v_cat_name := v_category_names[i];
    SELECT id INTO v_cat_id FROM marketplace_categories WHERE slug = v_cat_slug AND is_active = true LIMIT 1;
    IF v_cat_id IS NOT NULL THEN
      v_categories := v_categories || jsonb_build_object(
        'id', v_cat_id,
        'slug', v_cat_slug,
        'name', v_cat_name
      );
    END IF;
  END LOOP;

  -- PASO 3: UPDATE del template kiosco
  UPDATE business_type_templates
  SET
    categories = v_categories,
    brands = '[
      {"name": "Coca-Cola", "suggested_for_categories": ["bebidas-sin-alcohol"]},
      {"name": "Pepsi", "suggested_for_categories": ["bebidas-sin-alcohol"]},
      {"name": "Arcor", "suggested_for_categories": ["golosinas", "galletas"]},
      {"name": "Lay''s", "suggested_for_categories": ["snacks-salados"]},
      {"name": "Milka", "suggested_for_categories": ["golosinas"]},
      {"name": "Beldent", "suggested_for_categories": ["golosinas"]},
      {"name": "Quilmes", "suggested_for_categories": ["bebidas-con-alcohol"]},
      {"name": "Fanta", "suggested_for_categories": ["bebidas-sin-alcohol"]},
      {"name": "Speed", "suggested_for_categories": ["bebidas-sin-alcohol"]},
      {"name": "Bagley", "suggested_for_categories": ["galletas", "golosinas"]}
    ]'::jsonb,
    products = '[
      {"name": "Coca-Cola 500ml", "category_slug": "bebidas-sin-alcohol", "brand": "Coca-Cola", "price_reference": 2000, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Sprite 500ml", "category_slug": "bebidas-sin-alcohol", "brand": "Coca-Cola", "price_reference": 1900, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Fanta 500ml", "category_slug": "bebidas-sin-alcohol", "brand": "Fanta", "price_reference": 1900, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Pepsi 500ml", "category_slug": "bebidas-sin-alcohol", "brand": "Pepsi", "price_reference": 1800, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Agua Villavicencio 500ml", "category_slug": "bebidas-sin-alcohol", "brand": "Villavicencio", "price_reference": 1300, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Levité 500ml", "category_slug": "bebidas-sin-alcohol", "brand": "Levité", "price_reference": 1500, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Speed 473ml", "category_slug": "bebidas-sin-alcohol", "brand": "Speed", "price_reference": 2500, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Red Bull 473ml", "category_slug": "bebidas-sin-alcohol", "brand": "Red Bull", "price_reference": 3000, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Cepita 200ml", "category_slug": "bebidas-sin-alcohol", "brand": "Coca-Cola", "price_reference": 900, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Tang 1L", "category_slug": "bebidas-sin-alcohol", "brand": "Tang", "price_reference": 2200, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Alfajor Jorgito", "category_slug": "golosinas", "brand": "Jorgito", "price_reference": 700, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Alfajor Tita", "category_slug": "golosinas", "brand": "Arcor", "price_reference": 600, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Alfajor Guaymallén", "category_slug": "golosinas", "brand": "Jorgito", "price_reference": 400, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Alfajor Milka", "category_slug": "golosinas", "brand": "Milka", "price_reference": 900, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Alfajor Capitán del Espacio", "category_slug": "golosinas", "brand": "Jorgito", "price_reference": 900, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Chocolate Milka 55g", "category_slug": "golosinas", "brand": "Milka", "price_reference": 1500, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Chocolate Shot 55g", "category_slug": "golosinas", "brand": "Arcor", "price_reference": 1400, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Bon o Bon unidad", "category_slug": "golosinas", "brand": "Arcor", "price_reference": 350, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Gomitas Mogul", "category_slug": "golosinas", "brand": "Arcor", "price_reference": 1200, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Chicle Beldent", "category_slug": "golosinas", "brand": "Beldent", "price_reference": 1000, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Papas Lay''s 47g", "category_slug": "snacks-salados", "brand": "Lay''s", "price_reference": 2000, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Papas Pehuamar 45g", "category_slug": "snacks-salados", "brand": "Pehuamar", "price_reference": 1700, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Cheetos", "category_slug": "snacks-salados", "brand": "Lay''s", "price_reference": 1600, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Doritos", "category_slug": "snacks-salados", "brand": "Lay''s", "price_reference": 1800, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Maní Georgalos", "category_slug": "snacks-salados", "brand": "Georgalos", "price_reference": 1400, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Oreo 118g", "category_slug": "galletas", "brand": "Oreo", "price_reference": 1500, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Criollitas 100g", "category_slug": "galletas", "brand": "Bagley", "price_reference": 1300, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Galletitas Terrabusi", "category_slug": "galletas", "brand": "Bagley", "price_reference": 1300, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Cerveza Quilmes 500ml", "category_slug": "bebidas-con-alcohol", "brand": "Quilmes", "price_reference": 1800, "unit": "unidad", "sku_prefix": "KIOSCO"},
      {"name": "Fernet Branca 750ml", "category_slug": "bebidas-con-alcohol", "brand": "Branca", "price_reference": 4500, "unit": "unidad", "sku_prefix": "KIOSCO"}
    ]'::jsonb,
    attributes = '[
      {"name": "Contenido Neto", "slug": "contenido-neto", "values": ["200ml","330ml","473ml","500ml","1L","1.5L","2.25L","50g","100g","200g"], "applies_to_categories": ["bebidas-sin-alcohol","bebidas-con-alcohol","golosinas","snacks-salados","galletas"]},
      {"name": "Tipo Envase", "slug": "tipo-envase", "values": ["Botella","Lata","Paquete","Caja","Bolsa"], "applies_to_categories": ["bebidas-sin-alcohol","bebidas-con-alcohol","golosinas","snacks-salados","galletas"]}
    ]'::jsonb,
    version = '3.0.0-curated',
    generated_by = 'manual-curation',
    updated_at = CURRENT_TIMESTAMP
  WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'kiosco') AND is_default = true;

  IF NOT FOUND THEN
    RAISE NOTICE 'Template kiosco (is_default=true) no encontrado. Ejecutar seed 004 antes de este seed.';
  ELSE
    RAISE NOTICE 'Template kiosco actualizado correctamente: 6 categorías, 10 marcas, 30 productos, 2 atributos.';
  END IF;
END $$;
