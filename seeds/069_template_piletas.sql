-- Seed 069: Template piletas/piscinas (business_type code='piletas')
-- CICLO: cycle-004-brands-catalog-expansion
-- FECHA: 2026-04-21
-- PROPÓSITO: Crear template default para el rubro piletas si existe el business_type.
--            Si el business_type 'piletas' no existe, loguea aviso y no inserta.
-- IDEMPOTENTE: INSERT... ON CONFLICT DO UPDATE
-- REQUIERE: seed 065 ya ejecutado (marcas piletas en marketplace_brands).
-- REQUIERE: que exista business_type con code='piletas' O code='corralon' (como extensión).
-- NOTA: En Posadas las piletas se venden en ferreterías y corralones.
--       Este template se aplica a comercios especializados en piletas/natación.
--       Si el rubro 'piletas' no está en el catálogo, este seed es un no-op hasta que se cree.

-- =====================================================
-- PASO 1: Asegurar categorías de piletas
-- =====================================================
INSERT INTO marketplace_categories (name, slug, description, parent_id, sort_order, is_active)
VALUES
  ('Cloro y Químicos',         'cloro-quimicos',        'Cloro, algicidas, pH, floculantes para piletas', NULL, 1, true),
  ('Bombas y Filtros',         'bombas-filtros-pileta', 'Bombas de recirculación y filtros de arena',     NULL, 2, true),
  ('Accesorios de Pileta',     'accesorios-pileta',     'Cepillos, mangas, skimmers, mangueras',          NULL, 3, true),
  ('Piletas Estructurales',    'piletas-estructurales', 'Piletas desmontables, inflables, de armazón',    NULL, 4, true),
  ('Mantenimiento de Pileta',  'mantenimiento-pileta',  'Robots limpiafondos, cobertores, kits de test',  NULL, 5, true)
ON CONFLICT (slug) DO UPDATE SET
  name        = EXCLUDED.name,
  description = EXCLUDED.description,
  is_active   = EXCLUDED.is_active;

-- =====================================================
-- PASO 2: Construir e insertar/actualizar template
-- =====================================================
DO $$
DECLARE
  v_bt_id      UUID;
  v_categories JSONB := '[]'::jsonb;
  v_cat_slugs  TEXT[] := ARRAY[
    'cloro-quimicos', 'bombas-filtros-pileta', 'accesorios-pileta',
    'piletas-estructurales', 'mantenimiento-pileta'
  ];
  v_cat_names  TEXT[] := ARRAY[
    'Cloro y Químicos', 'Bombas y Filtros', 'Accesorios',
    'Piletas', 'Mantenimiento'
  ];
  v_cat_id   UUID;
  v_cat_slug TEXT;
  v_cat_name TEXT;
  i          INT;
BEGIN
  -- Buscar business_type piletas
  SELECT id INTO v_bt_id FROM business_types WHERE code = 'piletas' LIMIT 1;

  IF v_bt_id IS NULL THEN
    RAISE NOTICE 'business_type ''piletas'' no existe. Seed 069 es no-op hasta que se cree el rubro.';
    RETURN;
  END IF;

  -- Construir array de categorías con IDs reales
  FOR i IN 1..array_length(v_cat_slugs, 1) LOOP
    v_cat_slug := v_cat_slugs[i];
    v_cat_name := v_cat_names[i];
    SELECT id INTO v_cat_id FROM marketplace_categories WHERE slug = v_cat_slug AND is_active = true LIMIT 1;
    IF v_cat_id IS NOT NULL THEN
      v_categories := v_categories || jsonb_build_object(
        'id',   v_cat_id,
        'slug', v_cat_slug,
        'name', v_cat_name
      );
    END IF;
  END LOOP;

  -- Insertar o actualizar template default
  INSERT INTO business_type_templates (
    business_type_id,
    name,
    is_default,
    categories,
    brands,
    products,
    attributes,
    version,
    generated_by,
    created_at,
    updated_at
  )
  VALUES (
    v_bt_id,
    'Template Piletas y Jardín',
    true,
    v_categories,
    '[
      {"name": "Clorotec",   "suggested_for_categories": ["cloro-quimicos"]},
      {"name": "Nataclor",   "suggested_for_categories": ["cloro-quimicos"]},
      {"name": "Freshclor",  "suggested_for_categories": ["cloro-quimicos"]},
      {"name": "Deep Blue",  "suggested_for_categories": ["cloro-quimicos"]},
      {"name": "Bestway",    "suggested_for_categories": ["piletas-estructurales", "accesorios-pileta"]},
      {"name": "Maytronics", "suggested_for_categories": ["mantenimiento-pileta"]},
      {"name": "Kokido",     "suggested_for_categories": ["accesorios-pileta"]},
      {"name": "Difran",     "suggested_for_categories": ["bombas-filtros-pileta"]},
      {"name": "Pool Xpert", "suggested_for_categories": ["cloro-quimicos", "accesorios-pileta"]}
    ]'::jsonb,
    '[
      {"name": "Cloro granulado 10kg",         "category_slug": "cloro-quimicos",        "brand": "Nataclor",   "price_reference": 55000,  "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "Cloro tabletas 1kg",           "category_slug": "cloro-quimicos",        "brand": "Clorotec",   "price_reference": 18000,  "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "Cloro líquido 5L",             "category_slug": "cloro-quimicos",        "brand": "Clorotec",   "price_reference": 12000,  "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "Algicida 1L",                  "category_slug": "cloro-quimicos",        "brand": "Nataclor",   "price_reference": 14000,  "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "pH menos 1kg",                 "category_slug": "cloro-quimicos",        "brand": "Nataclor",   "price_reference": 12000,  "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "pH más 1kg",                   "category_slug": "cloro-quimicos",        "brand": "Nataclor",   "price_reference": 11000,  "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "Kit test análisis agua",       "category_slug": "mantenimiento-pileta",  "brand": "Nataclor",   "price_reference": 8000,   "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "Bomba de recirculación 1/2HP", "category_slug": "bombas-filtros-pileta", "brand": "Difran",     "price_reference": 120000, "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "Filtro de arena 12 pulgadas",  "category_slug": "bombas-filtros-pileta", "brand": "Difran",     "price_reference": 180000, "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "Cepillo de fondo 45cm",        "category_slug": "accesorios-pileta",     "brand": "Kokido",     "price_reference": 6000,   "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "Manga aspiradora telescópica", "category_slug": "accesorios-pileta",     "brand": "Kokido",     "price_reference": 22000,  "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "Pileta estructural 3.66m",     "category_slug": "piletas-estructurales", "brand": "Bestway",    "price_reference": 280000, "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "Pileta inflable familiar 3m",  "category_slug": "piletas-estructurales", "brand": "Bestway",    "price_reference": 85000,  "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "Cobertor solar 4x8m",          "category_slug": "mantenimiento-pileta",  "brand": "Deep Blue",  "price_reference": 45000,  "unit": "unidad", "sku_prefix": "PIL"},
      {"name": "Robot limpiafondos Dolphin",   "category_slug": "mantenimiento-pileta",  "brand": "Maytronics", "price_reference": 500000, "unit": "unidad", "sku_prefix": "PIL"}
    ]'::jsonb,
    '[
      {"name": "Presentación",  "slug": "presentacion-pileta",  "values": ["1kg","5kg","10kg","1L","5L","10L","unidad"], "applies_to_categories": ["cloro-quimicos"]},
      {"name": "Potencia",      "slug": "potencia-bomba",       "values": ["1/3HP","1/2HP","3/4HP","1HP","1.5HP"],       "applies_to_categories": ["bombas-filtros-pileta"]},
      {"name": "Diámetro",      "slug": "diametro-filtro",      "values": ["10\"","12\"","16\"","20\"","24\""],           "applies_to_categories": ["bombas-filtros-pileta"]},
      {"name": "Medida pileta", "slug": "medida-pileta",        "values": ["2m","2.44m","3m","3.66m","4.57m","5.49m"],   "applies_to_categories": ["piletas-estructurales"]}
    ]'::jsonb,
    '1.0.0',
    'cycle-004-catalog-researcher',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  )
  ON CONFLICT (business_type_id, region) WHERE is_default = true
  DO UPDATE SET
    categories   = EXCLUDED.categories,
    brands       = EXCLUDED.brands,
    products     = EXCLUDED.products,
    attributes   = EXCLUDED.attributes,
    version      = EXCLUDED.version,
    generated_by = EXCLUDED.generated_by,
    updated_at   = CURRENT_TIMESTAMP
  WHERE
    business_type_templates.brands IS DISTINCT FROM EXCLUDED.brands;

  RAISE NOTICE 'Template piletas insertado/actualizado: 5 categorías, 9 marcas, 15 productos, 4 atributos.';
END $$;
