-- Seed 022: Template curado para Almacén (business_type code='almacen')
-- PROPÓSITO: UPSERT del template default con categorías, marcas, productos y atributos curados para almacenes argentinos
-- DIFERENCIA vs KIOSCO: Almacén enfoca en productos para el hogar, formatos familiares (1L, 1kg), NO porciones individuales
-- IDEMPOTENTE: Usa UPDATE con WHERE; no afecta si el template no existe (se ejecuta después de 004)
-- REQUIERE: 001 (business_types), 011 (marketplace_categories)
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
    'bebidas-sin-alcohol', 'lacteos-liquidos', 'lacteos-fiambres', 'alimentos-envasados',
    'conservas', 'pastas-cereales', 'snacks-salados', 'general'
  ];
  v_category_names TEXT[] := ARRAY[
    'Bebidas', 'Lácteos', 'Fiambres y Quesos', 'Almacén Seco',
    'Conservas y Enlatados', 'Harinas y Pastas', 'Snacks', 'Limpieza y Hogar'
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

  -- PASO 3: UPDATE del template almacen
  UPDATE business_type_templates
  SET
    categories = v_categories,
    brands = '[
      {"name": "Marolio", "suggested_for_categories": ["alimentos-envasados","conservas"]},
      {"name": "La Serenísima", "suggested_for_categories": ["lacteos-liquidos","lacteos-fiambres"]},
      {"name": "Arcor", "suggested_for_categories": ["alimentos-envasados","conservas"]},
      {"name": "Molinos Río de la Plata", "suggested_for_categories": ["pastas-cereales","alimentos-envasados"]},
      {"name": "Unilever", "suggested_for_categories": ["general"]},
      {"name": "Coca-Cola", "suggested_for_categories": ["bebidas-sin-alcohol"]},
      {"name": "Quilmes", "suggested_for_categories": ["bebidas-sin-alcohol"]},
      {"name": "Bagley", "suggested_for_categories": ["snacks-salados","alimentos-envasados"]},
      {"name": "Fargo", "suggested_for_categories": ["alimentos-envasados"]},
      {"name": "Skip", "suggested_for_categories": ["general"]},
      {"name": "Cif", "suggested_for_categories": ["general"]},
      {"name": "Sancor", "suggested_for_categories": ["lacteos-liquidos"]}
    ]'::jsonb,
    products = '[
      {"name": "Coca-Cola 2.25L", "category_slug": "bebidas-sin-alcohol", "brand": "Coca-Cola", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Agua mineral 2L", "category_slug": "bebidas-sin-alcohol", "brand": "Villavicencio", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Jugo Tang polvo", "category_slug": "bebidas-sin-alcohol", "brand": "Tang", "price_reference": 600, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Leche La Serenísima 1L", "category_slug": "lacteos-liquidos", "brand": "La Serenísima", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Yogur La Serenísima 1kg", "category_slug": "lacteos-liquidos", "brand": "La Serenísima", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Manteca 200g", "category_slug": "lacteos-liquidos", "brand": "La Serenísima", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Crema de leche 200ml", "category_slug": "lacteos-liquidos", "brand": "La Serenísima", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Queso cremoso kg", "category_slug": "lacteos-fiambres", "brand": "La Serenísima", "price_reference": 8000, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Jamón cocido kg", "category_slug": "lacteos-fiambres", "brand": "Paladini", "price_reference": 12000, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Salame kg", "category_slug": "lacteos-fiambres", "brand": "Paladini", "price_reference": 10000, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Queso de máquina kg", "category_slug": "lacteos-fiambres", "brand": "Sancor", "price_reference": 7000, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Mortadela kg", "category_slug": "lacteos-fiambres", "brand": "Paladini", "price_reference": 8500, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Leche Sancor 1L", "category_slug": "lacteos-liquidos", "brand": "Sancor", "price_reference": 1700, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Aceite girasol 1.5L", "category_slug": "alimentos-envasados", "brand": "Marolio", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Yerba mate 1kg", "category_slug": "alimentos-envasados", "brand": "Rosamonte", "price_reference": 4200, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Azúcar 1kg", "category_slug": "alimentos-envasados", "brand": "Ledesma", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Arroz 1kg", "category_slug": "alimentos-envasados", "brand": "Molinos", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Harina 1kg", "category_slug": "alimentos-envasados", "brand": "Blancaflor", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Fideos Matarazzo 500g", "category_slug": "pastas-cereales", "brand": "Molinos Río de la Plata", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Pan lactal Fargo", "category_slug": "alimentos-envasados", "brand": "Fargo", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Tomate triturado 520g", "category_slug": "conservas", "brand": "Marolio", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Arvejas lata", "category_slug": "conservas", "brand": "Arcor", "price_reference": 800, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Atún lata", "category_slug": "conservas", "brand": "La Campagnola", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Duraznos en almíbar", "category_slug": "conservas", "brand": "Arcor", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Fideos Lucchetti 500g", "category_slug": "pastas-cereales", "brand": "Lucchetti", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Polenta 500g", "category_slug": "pastas-cereales", "brand": "Marolio", "price_reference": 1000, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Puré de tomate", "category_slug": "conservas", "brand": "Marolio", "price_reference": 900, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Sprite 2.25L", "category_slug": "bebidas-sin-alcohol", "brand": "Coca-Cola", "price_reference": 2700, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Gatorade 1L", "category_slug": "bebidas-sin-alcohol", "brand": "PepsiCo", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Dulce de leche 200g", "category_slug": "alimentos-envasados", "brand": "La Serenísima", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Café torrado 500g", "category_slug": "alimentos-envasados", "brand": "Nescafé", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Té negro 25 saquitos", "category_slug": "alimentos-envasados", "brand": "Taragüí", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Choclo lata", "category_slug": "conservas", "brand": "La Campagnola", "price_reference": 1100, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Maíz en conserva", "category_slug": "conservas", "brand": "Arcor", "price_reference": 900, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Arroz integral 1kg", "category_slug": "alimentos-envasados", "brand": "Molinos", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Avena 500g", "category_slug": "pastas-cereales", "brand": "Quaker", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Papas fritas Bagley 170g", "category_slug": "snacks-salados", "brand": "Bagley", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Maní con chocolate", "category_slug": "snacks-salados", "brand": "Georgalos", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Criollitas 100g", "category_slug": "snacks-salados", "brand": "Bagley", "price_reference": 1300, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Galletitas Oreo 118g", "category_slug": "snacks-salados", "brand": "Oreo", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Lavandina 2L", "category_slug": "general", "brand": "Ayudín", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Detergente Magistral 750ml", "category_slug": "general", "brand": "Magistral", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Jabón en polvo Skip 800g", "category_slug": "general", "brand": "Skip", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Limpiador Cif 750ml", "category_slug": "general", "brand": "Cif", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Papel higiénico x4", "category_slug": "general", "brand": "Higienol", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Esponja de cocina", "category_slug": "general", "brand": "Scotch-Brite", "price_reference": 800, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Suavizante 1L", "category_slug": "general", "brand": "Skip", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMACEN"},
      {"name": "Jabón líquido 500ml", "category_slug": "general", "brand": "Ala", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMACEN"}
    ]'::jsonb,
    attributes = '[
      {"name": "Contenido Neto", "slug": "contenido-neto", "values": ["200g","500g","1kg","2.5kg","200ml","500ml","1L","1.5L","2L","2.25L"], "applies_to_categories": ["bebidas-sin-alcohol","lacteos-liquidos","alimentos-envasados","conservas","pastas-cereales"]},
      {"name": "Tipo Envase", "slug": "tipo-envase", "values": ["Botella","Sachet","Caja","Bolsa","Lata","Frasco","Tetra Brik","Bidón"], "applies_to_categories": ["bebidas-sin-alcohol","lacteos-liquidos","alimentos-envasados","conservas","pastas-cereales","general"]}
    ]'::jsonb,
    version = '3.0.0-curated',
    generated_by = 'manual-curation',
    updated_at = CURRENT_TIMESTAMP
  WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'almacen') AND is_default = true;

  IF NOT FOUND THEN
    RAISE NOTICE 'Template almacen (is_default=true) no encontrado. Ejecutar seed 004 antes de este seed.';
  ELSE
    RAISE NOTICE 'Template almacen actualizado correctamente: 8 categorías, 13 marcas, 45 productos, 2 atributos.';
  END IF;
END $$;
