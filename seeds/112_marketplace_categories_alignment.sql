-- =============================================================================
-- SEED 112: Alineación marketplace_categories ↔ global_products
-- =============================================================================
-- PROPÓSITO:
--   1. Insertar categorías raíz por business_type (level=0, parent_id NULL)
--   2. Insertar subcategorías faltantes (level=1) con slugs reales de global_products
--   3. Normalizar slugs problemáticos en global_products
--   4. Corregir mismatches de slugs en business_type_templates.categories (JSONB)
--
-- IDEMPOTENTE:
--   - INSERT usa ON CONFLICT (slug) DO NOTHING
--   - UPDATE de global_products son seguros (condicionales con WHERE exacto)
--   - UPDATE de templates usan jsonb_agg con CASE, idempotentes por naturaleza
--
-- REQUIERE: 001 (business_types), 003 (marketplace_categories), 004 (templates)
-- FECHA: 2026-04-25 | ZONA: Posadas, Misiones — NEA
-- =============================================================================


-- =============================================================================
-- PARTE 1 — Categorías raíz por rubro (level=0, parent_id NULL)
-- UUIDs fijos para poder referenciarlos en Parte 2
-- =============================================================================

INSERT INTO marketplace_categories (id, name, slug, level, sort_order, is_active) VALUES
  ('aa000001-0000-4000-8000-000000000001', 'Almacén',          'almacen-rubro',          0, 10,  true),
  ('aa000001-0000-4000-8000-000000000002', 'Ferretería',        'ferreteria-rubro',        0, 20,  true),
  ('aa000001-0000-4000-8000-000000000003', 'Vinoteca',          'vinoteca-rubro',          0, 30,  true),
  ('aa000001-0000-4000-8000-000000000004', 'Kiosco',            'kiosco-rubro',            0, 40,  true),
  ('aa000001-0000-4000-8000-000000000005', 'Perfumería',        'perfumeria-rubro',        0, 50,  true),
  ('aa000001-0000-4000-8000-000000000006', 'Panadería',         'panaderia-rubro',         0, 60,  true),
  ('aa000001-0000-4000-8000-000000000007', 'Carnicería',        'carniceria-rubro',        0, 70,  true),
  ('aa000001-0000-4000-8000-000000000008', 'Verdulería',        'verduleria-rubro',        0, 80,  true),
  ('aa000001-0000-4000-8000-000000000009', 'Fiambrería',        'fiambreria-rubro',        0, 90,  true),
  ('aa000001-0000-4000-8000-000000000010', 'Piletas',           'piletas-rubro',           0, 100, true),
  ('aa000001-0000-4000-8000-000000000011', 'Bazar',             'bazar-rubro',             0, 110, true),
  ('aa000001-0000-4000-8000-000000000012', 'Juguetería',        'jugueteria-rubro',        0, 120, true),
  ('aa000001-0000-4000-8000-000000000013', 'Librería',          'libreria-rubro',          0, 130, true),
  ('aa000001-0000-4000-8000-000000000014', 'Ropa y Calzado',   'ropa-rubro',              0, 140, true),
  ('aa000001-0000-4000-8000-000000000015', 'Electrodomésticos', 'electrodomesticos-rubro', 0, 150, true),
  ('aa000001-0000-4000-8000-000000000016', 'Corralón',          'corralon-rubro',          0, 160, true),
  ('aa000001-0000-4000-8000-000000000017', 'Veterinaria',       'veterinaria-rubro',       0, 170, true)
ON CONFLICT (slug) DO NOTHING;


-- =============================================================================
-- PARTE 2 — Subcategorías faltantes (level=1)
-- UUIDs generados con gen_random_uuid() — no fijos
-- Nombre: versión legible del slug (guiones → espacios, primera letra mayúscula)
-- =============================================================================

-- -----------------------------------------------------------
-- ALMACÉN (parent: aa000001-0000-4000-8000-000000000001)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Aceites vinagres',        'aceites-vinagres',        'aa000001-0000-4000-8000-000000000001', 1, 10,  true),
  (gen_random_uuid(), 'Alfajores',               'alfajores',               'aa000001-0000-4000-8000-000000000001', 1, 20,  true),
  (gen_random_uuid(), 'Almacen seco',            'almacen-seco',            'aa000001-0000-4000-8000-000000000001', 1, 30,  true),
  (gen_random_uuid(), 'Aperitivos licores',       'aperitivos-licores',      'aa000001-0000-4000-8000-000000000001', 1, 40,  true),
  (gen_random_uuid(), 'Arroz legumbres',          'arroz-legumbres',         'aa000001-0000-4000-8000-000000000001', 1, 50,  true),
  (gen_random_uuid(), 'Cervezas vinos',           'cervezas-vinos',          'aa000001-0000-4000-8000-000000000001', 1, 60,  true),
  (gen_random_uuid(), 'Chocolates',              'chocolates',              'aa000001-0000-4000-8000-000000000001', 1, 70,  true),
  (gen_random_uuid(), 'Conservas enlatados',      'conservas-enlatados',     'aa000001-0000-4000-8000-000000000001', 1, 80,  true),
  (gen_random_uuid(), 'Cuidado bebe',             'cuidado-bebe',            'aa000001-0000-4000-8000-000000000001', 1, 90,  true),
  (gen_random_uuid(), 'Detergentes jabones',      'detergentes-jabones',     'aa000001-0000-4000-8000-000000000001', 1, 100, true),
  (gen_random_uuid(), 'Espumantes',              'espumantes',              'aa000001-0000-4000-8000-000000000001', 1, 110, true),
  (gen_random_uuid(), 'Fiambres embutidos',       'fiambres-embutidos',      'aa000001-0000-4000-8000-000000000001', 1, 120, true),
  (gen_random_uuid(), 'Galletitas almacen',       'galletitas-almacen',      'aa000001-0000-4000-8000-000000000001', 1, 130, true),
  (gen_random_uuid(), 'Galletitas dulces',        'galletitas-dulces',       'aa000001-0000-4000-8000-000000000001', 1, 140, true),
  (gen_random_uuid(), 'Galletitas saladas',       'galletitas-saladas',      'aa000001-0000-4000-8000-000000000001', 1, 150, true),
  (gen_random_uuid(), 'Gaseosas aguas',           'gaseosas-aguas',          'aa000001-0000-4000-8000-000000000001', 1, 160, true),
  (gen_random_uuid(), 'Harinas premezclas',       'harinas-premezclas',      'aa000001-0000-4000-8000-000000000001', 1, 170, true),
  (gen_random_uuid(), 'Higiene personal',         'higiene-personal',        'aa000001-0000-4000-8000-000000000001', 1, 180, true),
  (gen_random_uuid(), 'Jabones desodorantes',     'jabones-desodorantes',    'aa000001-0000-4000-8000-000000000001', 1, 190, true),
  (gen_random_uuid(), 'Jugos polvos',             'jugos-polvos',            'aa000001-0000-4000-8000-000000000001', 1, 200, true),
  (gen_random_uuid(), 'Lavandina desinfectantes', 'lavandina-desinfectantes','aa000001-0000-4000-8000-000000000001', 1, 210, true),
  (gen_random_uuid(), 'Panales',                 'panales',                 'aa000001-0000-4000-8000-000000000001', 1, 220, true),
  (gen_random_uuid(), 'Pan envasado',             'pan-envasado',            'aa000001-0000-4000-8000-000000000001', 1, 230, true),
  (gen_random_uuid(), 'Papel higiene',            'papel-higiene',           'aa000001-0000-4000-8000-000000000001', 1, 240, true),
  (gen_random_uuid(), 'Papel higienico',          'papel-higienico',         'aa000001-0000-4000-8000-000000000001', 1, 250, true),
  (gen_random_uuid(), 'Pastas secas',             'pastas-secas',            'aa000001-0000-4000-8000-000000000001', 1, 260, true),
  (gen_random_uuid(), 'Quesos manteca',           'quesos-manteca',          'aa000001-0000-4000-8000-000000000001', 1, 270, true),
  (gen_random_uuid(), 'Shampoo acondicionador',   'shampoo-acondicionador',  'aa000001-0000-4000-8000-000000000001', 1, 280, true),
  (gen_random_uuid(), 'Snacks salados alm',       'snacks-salados-alm',      'aa000001-0000-4000-8000-000000000001', 1, 290, true),
  (gen_random_uuid(), 'Vinos blancos',            'vinos-blancos',           'aa000001-0000-4000-8000-000000000001', 1, 300, true),
  (gen_random_uuid(), 'Vinos rosados',            'vinos-rosados',           'aa000001-0000-4000-8000-000000000001', 1, 310, true),
  (gen_random_uuid(), 'Vinos tintos',             'vinos-tintos',            'aa000001-0000-4000-8000-000000000001', 1, 320, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- FERRETERÍA (parent: aa000001-0000-4000-8000-000000000002)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Accesorios pintura',    'accesorios-pintura',    'aa000001-0000-4000-8000-000000000002', 1, 10,  true),
  (gen_random_uuid(), 'Adhesivos selladores',  'adhesivos-selladores',  'aa000001-0000-4000-8000-000000000002', 1, 20,  true),
  (gen_random_uuid(), 'Cables electricidad',   'cables-electricidad',   'aa000001-0000-4000-8000-000000000002', 1, 30,  true),
  (gen_random_uuid(), 'Canos pvc',             'canos-pvc',             'aa000001-0000-4000-8000-000000000002', 1, 40,  true),
  (gen_random_uuid(), 'Cerrajeria seguridad',  'cerrajeria-seguridad',  'aa000001-0000-4000-8000-000000000002', 1, 50,  true),
  (gen_random_uuid(), 'Esmaltes barnices',     'esmaltes-barnices',     'aa000001-0000-4000-8000-000000000002', 1, 60,  true),
  (gen_random_uuid(), 'Fijaciones',            'fijaciones',            'aa000001-0000-4000-8000-000000000002', 1, 70,  true),
  (gen_random_uuid(), 'Herramientas electricas','herramientas-electricas','aa000001-0000-4000-8000-000000000002', 1, 80,  true),
  (gen_random_uuid(), 'Herramientas manuales', 'herramientas-manuales', 'aa000001-0000-4000-8000-000000000002', 1, 90,  true),
  (gen_random_uuid(), 'Jardin herramientas',   'jardin-herramientas',   'aa000001-0000-4000-8000-000000000002', 1, 100, true),
  (gen_random_uuid(), 'Latex ferret',          'latex-ferret',          'aa000001-0000-4000-8000-000000000002', 1, 110, true),
  (gen_random_uuid(), 'Llaves tomacorrientes', 'llaves-tomacorrientes', 'aa000001-0000-4000-8000-000000000002', 1, 120, true),
  (gen_random_uuid(), 'Llaves valvulas',       'llaves-valvulas',       'aa000001-0000-4000-8000-000000000002', 1, 130, true),
  (gen_random_uuid(), 'Mangueras riego',       'mangueras-riego',       'aa000001-0000-4000-8000-000000000002', 1, 140, true),
  (gen_random_uuid(), 'Medicion',              'medicion',              'aa000001-0000-4000-8000-000000000002', 1, 150, true),
  (gen_random_uuid(), 'Membranas aislantes',   'membranas-aislantes',   'aa000001-0000-4000-8000-000000000002', 1, 160, true),
  (gen_random_uuid(), 'Pinturas latex',        'pinturas-latex',        'aa000001-0000-4000-8000-000000000002', 1, 170, true),
  (gen_random_uuid(), 'Proteccion electrica',  'proteccion-electrica',  'aa000001-0000-4000-8000-000000000002', 1, 180, true),
  (gen_random_uuid(), 'Sanitarios plomeria',   'sanitarios-plomeria',   'aa000001-0000-4000-8000-000000000002', 1, 190, true),
  (gen_random_uuid(), 'Sierras electricas',    'sierras-electricas',    'aa000001-0000-4000-8000-000000000002', 1, 200, true),
  (gen_random_uuid(), 'Taladros percutores',   'taladros-percutores',   'aa000001-0000-4000-8000-000000000002', 1, 210, true),
  (gen_random_uuid(), 'Tornillos clavos',      'tornillos-clavos',      'aa000001-0000-4000-8000-000000000002', 1, 220, true),
  (gen_random_uuid(), 'Iluminacion',           'iluminacion',           'aa000001-0000-4000-8000-000000000002', 1, 230, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- VINOTECA (parent: aa000001-0000-4000-8000-000000000003)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Accesorios vino',        'accesorios-vino',        'aa000001-0000-4000-8000-000000000003', 1, 10,  true),
  (gen_random_uuid(), 'Aperitivos licores',      'aperitivos-licores',     'aa000001-0000-4000-8000-000000000003', 1, 20,  true),
  (gen_random_uuid(), 'Cervezas artesanales',   'cervezas-artesanales',   'aa000001-0000-4000-8000-000000000003', 1, 30,  true),
  (gen_random_uuid(), 'Cervezas importadas',    'cervezas-importadas',    'aa000001-0000-4000-8000-000000000003', 1, 40,  true),
  (gen_random_uuid(), 'Espumantes',             'espumantes',             'aa000001-0000-4000-8000-000000000003', 1, 50,  true),
  (gen_random_uuid(), 'Gin',                   'gin',                   'aa000001-0000-4000-8000-000000000003', 1, 60,  true),
  (gen_random_uuid(), 'Rum brandy',             'rum-brandy',             'aa000001-0000-4000-8000-000000000003', 1, 70,  true),
  (gen_random_uuid(), 'Vermouths aperitivos',   'vermouths-aperitivos',   'aa000001-0000-4000-8000-000000000003', 1, 80,  true),
  (gen_random_uuid(), 'Vinos blancos',          'vinos-blancos',          'aa000001-0000-4000-8000-000000000003', 1, 90,  true),
  (gen_random_uuid(), 'Vinos rosados',          'vinos-rosados',          'aa000001-0000-4000-8000-000000000003', 1, 100, true),
  (gen_random_uuid(), 'Vinos tintos',           'vinos-tintos',           'aa000001-0000-4000-8000-000000000003', 1, 110, true),
  (gen_random_uuid(), 'Vodka',                 'vodka',                 'aa000001-0000-4000-8000-000000000003', 1, 120, true),
  (gen_random_uuid(), 'Whisky',                'whisky',                'aa000001-0000-4000-8000-000000000003', 1, 130, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- KIOSCO (parent: aa000001-0000-4000-8000-000000000004)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Aguas kiosco',          'aguas-kiosco',          'aa000001-0000-4000-8000-000000000004', 1, 10,  true),
  (gen_random_uuid(), 'Aguas saborizadas',      'aguas-saborizadas',      'aa000001-0000-4000-8000-000000000004', 1, 20,  true),
  (gen_random_uuid(), 'Alfajores',             'alfajores',             'aa000001-0000-4000-8000-000000000004', 1, 30,  true),
  (gen_random_uuid(), 'Caramelos chicles',      'caramelos-chicles',      'aa000001-0000-4000-8000-000000000004', 1, 40,  true),
  (gen_random_uuid(), 'Cervezas',             'cervezas',             'aa000001-0000-4000-8000-000000000004', 1, 50,  true),
  (gen_random_uuid(), 'Chocolates',           'chocolates',           'aa000001-0000-4000-8000-000000000004', 1, 60,  true),
  (gen_random_uuid(), 'Chupetines',           'chupetines',           'aa000001-0000-4000-8000-000000000004', 1, 70,  true),
  (gen_random_uuid(), 'Cigarrillos',          'cigarrillos',          'aa000001-0000-4000-8000-000000000004', 1, 80,  true),
  (gen_random_uuid(), 'Encendedores',         'encendedores',         'aa000001-0000-4000-8000-000000000004', 1, 90,  true),
  (gen_random_uuid(), 'Energeticas',          'energeticas',          'aa000001-0000-4000-8000-000000000004', 1, 100, true),
  (gen_random_uuid(), 'Energizantes',         'energizantes',         'aa000001-0000-4000-8000-000000000004', 1, 110, true),
  (gen_random_uuid(), 'Frutos secos',          'frutos-secos',          'aa000001-0000-4000-8000-000000000004', 1, 120, true),
  (gen_random_uuid(), 'Galletitas dulces',     'galletitas-dulces',     'aa000001-0000-4000-8000-000000000004', 1, 130, true),
  (gen_random_uuid(), 'Galletitas saladas',    'galletitas-saladas',    'aa000001-0000-4000-8000-000000000004', 1, 140, true),
  (gen_random_uuid(), 'Gaseosas',             'gaseosas',             'aa000001-0000-4000-8000-000000000004', 1, 150, true),
  (gen_random_uuid(), 'Gaseosas lata',         'gaseosas-lata',         'aa000001-0000-4000-8000-000000000004', 1, 160, true),
  (gen_random_uuid(), 'Gomitas confites',      'gomitas-confites',      'aa000001-0000-4000-8000-000000000004', 1, 170, true),
  (gen_random_uuid(), 'Gomitas malvaviscos',   'gomitas-malvaviscos',   'aa000001-0000-4000-8000-000000000004', 1, 180, true),
  (gen_random_uuid(), 'Higiene personal',      'higiene-personal',      'aa000001-0000-4000-8000-000000000004', 1, 190, true),
  (gen_random_uuid(), 'Jugos',               'jugos',               'aa000001-0000-4000-8000-000000000004', 1, 200, true),
  (gen_random_uuid(), 'Jugos kiosco',          'jugos-kiosco',          'aa000001-0000-4000-8000-000000000004', 1, 210, true),
  (gen_random_uuid(), 'Mani snacks',           'mani-snacks',           'aa000001-0000-4000-8000-000000000004', 1, 220, true),
  (gen_random_uuid(), 'Palitos chizitos',      'palitos-chizitos',      'aa000001-0000-4000-8000-000000000004', 1, 230, true),
  (gen_random_uuid(), 'Papas fritas',          'papas-fritas',          'aa000001-0000-4000-8000-000000000004', 1, 240, true),
  (gen_random_uuid(), 'Pilas baterias',        'pilas-baterias',        'aa000001-0000-4000-8000-000000000004', 1, 250, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- PERFUMERÍA (parent: aa000001-0000-4000-8000-000000000005)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Accesorios belleza',     'accesorios-belleza',     'aa000001-0000-4000-8000-000000000005', 1, 10,  true),
  (gen_random_uuid(), 'Cremas corporales',      'cremas-corporales',      'aa000001-0000-4000-8000-000000000005', 1, 20,  true),
  (gen_random_uuid(), 'Cuidado capilar',        'cuidado-capilar',        'aa000001-0000-4000-8000-000000000005', 1, 30,  true),
  (gen_random_uuid(), 'Cuidado corporal',       'cuidado-corporal',       'aa000001-0000-4000-8000-000000000005', 1, 40,  true),
  (gen_random_uuid(), 'Cuidado facial',         'cuidado-facial',         'aa000001-0000-4000-8000-000000000005', 1, 50,  true),
  (gen_random_uuid(), 'Esmaltes maquillaje',    'esmaltes-maquillaje',    'aa000001-0000-4000-8000-000000000005', 1, 60,  true),
  (gen_random_uuid(), 'Esmaltes unas',          'esmaltes-unas',          'aa000001-0000-4000-8000-000000000005', 1, 70,  true),
  (gen_random_uuid(), 'Fragancias',            'fragancias',            'aa000001-0000-4000-8000-000000000005', 1, 80,  true),
  (gen_random_uuid(), 'Fragancias hombre',      'fragancias-hombre',      'aa000001-0000-4000-8000-000000000005', 1, 90,  true),
  (gen_random_uuid(), 'Fragancias mujer',       'fragancias-mujer',       'aa000001-0000-4000-8000-000000000005', 1, 100, true),
  (gen_random_uuid(), 'Maquillaje labios',      'maquillaje-labios',      'aa000001-0000-4000-8000-000000000005', 1, 110, true),
  (gen_random_uuid(), 'Maquillaje ojos',        'maquillaje-ojos',        'aa000001-0000-4000-8000-000000000005', 1, 120, true),
  (gen_random_uuid(), 'Maquillaje rostro',      'maquillaje-rostro',      'aa000001-0000-4000-8000-000000000005', 1, 130, true),
  (gen_random_uuid(), 'Shampoos profesionales', 'shampoos-profesionales', 'aa000001-0000-4000-8000-000000000005', 1, 140, true),
  (gen_random_uuid(), 'Tinturas capilares',     'tinturas-capilares',     'aa000001-0000-4000-8000-000000000005', 1, 150, true),
  (gen_random_uuid(), 'Tratamientos capilares', 'tratamientos-capilares', 'aa000001-0000-4000-8000-000000000005', 1, 160, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- PANADERÍA (parent: aa000001-0000-4000-8000-000000000006)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Bizcochos',               'bizcochos',               'aa000001-0000-4000-8000-000000000006', 1, 10,  true),
  (gen_random_uuid(), 'Budines',                'budines',                'aa000001-0000-4000-8000-000000000006', 1, 20,  true),
  (gen_random_uuid(), 'Confiteria',             'confiteria',             'aa000001-0000-4000-8000-000000000006', 1, 30,  true),
  (gen_random_uuid(), 'Dulce leche reposteria',  'dulce-leche-reposteria',  'aa000001-0000-4000-8000-000000000006', 1, 40,  true),
  (gen_random_uuid(), 'Esencias aditivos',       'esencias-aditivos',       'aa000001-0000-4000-8000-000000000006', 1, 50,  true),
  (gen_random_uuid(), 'Facturas',               'facturas',               'aa000001-0000-4000-8000-000000000006', 1, 60,  true),
  (gen_random_uuid(), 'Galletitas saladas',      'galletitas-saladas',      'aa000001-0000-4000-8000-000000000006', 1, 70,  true),
  (gen_random_uuid(), 'Galletitas secas',        'galletitas-secas',        'aa000001-0000-4000-8000-000000000006', 1, 80,  true),
  (gen_random_uuid(), 'Harinas panaderia',       'harinas-panaderia',       'aa000001-0000-4000-8000-000000000006', 1, 90,  true),
  (gen_random_uuid(), 'Insumos panaderia',       'insumos-panaderia',       'aa000001-0000-4000-8000-000000000006', 1, 100, true),
  (gen_random_uuid(), 'Levaduras',              'levaduras',              'aa000001-0000-4000-8000-000000000006', 1, 110, true),
  (gen_random_uuid(), 'Margarinas industriales', 'margarinas-industriales', 'aa000001-0000-4000-8000-000000000006', 1, 120, true),
  (gen_random_uuid(), 'Pan fresco',              'pan-fresco',              'aa000001-0000-4000-8000-000000000006', 1, 130, true),
  (gen_random_uuid(), 'Papeles moldes',          'papeles-moldes',          'aa000001-0000-4000-8000-000000000006', 1, 140, true),
  (gen_random_uuid(), 'Productos integrales',    'productos-integrales',    'aa000001-0000-4000-8000-000000000006', 1, 150, true),
  (gen_random_uuid(), 'Rellenos coberturas',     'rellenos-coberturas',     'aa000001-0000-4000-8000-000000000006', 1, 160, true),
  (gen_random_uuid(), 'Sandwiches',             'sandwiches',             'aa000001-0000-4000-8000-000000000006', 1, 170, true),
  (gen_random_uuid(), 'Tortas',                'tortas',                'aa000001-0000-4000-8000-000000000006', 1, 180, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- CARNICERÍA (parent: aa000001-0000-4000-8000-000000000007)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Achuras',              'achuras',              'aa000001-0000-4000-8000-000000000007', 1, 10,  true),
  (gen_random_uuid(), 'Carne molida',         'carne-molida',         'aa000001-0000-4000-8000-000000000007', 1, 20,  true),
  (gen_random_uuid(), 'Cerdo',               'cerdo',               'aa000001-0000-4000-8000-000000000007', 1, 30,  true),
  (gen_random_uuid(), 'Chacinados',           'chacinados',           'aa000001-0000-4000-8000-000000000007', 1, 40,  true),
  (gen_random_uuid(), 'Cortes cerdo',         'cortes-cerdo',         'aa000001-0000-4000-8000-000000000007', 1, 50,  true),
  (gen_random_uuid(), 'Cortes pollo',         'cortes-pollo',         'aa000001-0000-4000-8000-000000000007', 1, 60,  true),
  (gen_random_uuid(), 'Cortes vacunos',       'cortes-vacunos',       'aa000001-0000-4000-8000-000000000007', 1, 70,  true),
  (gen_random_uuid(), 'Embutidos',           'embutidos',           'aa000001-0000-4000-8000-000000000007', 1, 80,  true),
  (gen_random_uuid(), 'Embutidos fiambres',   'embutidos-fiambres',   'aa000001-0000-4000-8000-000000000007', 1, 90,  true),
  (gen_random_uuid(), 'Fiambres',            'fiambres',            'aa000001-0000-4000-8000-000000000007', 1, 100, true),
  (gen_random_uuid(), 'Hamburguesas',        'hamburguesas',        'aa000001-0000-4000-8000-000000000007', 1, 110, true),
  (gen_random_uuid(), 'Milanesas',           'milanesas',           'aa000001-0000-4000-8000-000000000007', 1, 120, true),
  (gen_random_uuid(), 'Pollo',              'pollo',              'aa000001-0000-4000-8000-000000000007', 1, 130, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- VERDULERÍA (parent: aa000001-0000-4000-8000-000000000008)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Aromaticas',         'aromaticas',         'aa000001-0000-4000-8000-000000000008', 1, 10,  true),
  (gen_random_uuid(), 'Bulbos',            'bulbos',            'aa000001-0000-4000-8000-000000000008', 1, 20,  true),
  (gen_random_uuid(), 'Cruciferas',         'cruciferas',         'aa000001-0000-4000-8000-000000000008', 1, 30,  true),
  (gen_random_uuid(), 'Frutas carozo',      'frutas-carozo',      'aa000001-0000-4000-8000-000000000008', 1, 40,  true),
  (gen_random_uuid(), 'Frutas citricos',    'frutas-citricos',    'aa000001-0000-4000-8000-000000000008', 1, 50,  true),
  (gen_random_uuid(), 'Frutas frescas',     'frutas-frescas',     'aa000001-0000-4000-8000-000000000008', 1, 60,  true),
  (gen_random_uuid(), 'Frutas pepita',      'frutas-pepita',      'aa000001-0000-4000-8000-000000000008', 1, 70,  true),
  (gen_random_uuid(), 'Frutas secas',       'frutas-secas',       'aa000001-0000-4000-8000-000000000008', 1, 80,  true),
  (gen_random_uuid(), 'Frutas tropicales',  'frutas-tropicales',  'aa000001-0000-4000-8000-000000000008', 1, 90,  true),
  (gen_random_uuid(), 'Hierbas aromaticas', 'hierbas-aromaticas', 'aa000001-0000-4000-8000-000000000008', 1, 100, true),
  (gen_random_uuid(), 'Huevos',            'huevos',            'aa000001-0000-4000-8000-000000000008', 1, 110, true),
  (gen_random_uuid(), 'Legumbres frescas',  'legumbres-frescas',  'aa000001-0000-4000-8000-000000000008', 1, 120, true),
  (gen_random_uuid(), 'Raices',            'raices',            'aa000001-0000-4000-8000-000000000008', 1, 130, true),
  (gen_random_uuid(), 'Tuberculos',         'tuberculos',         'aa000001-0000-4000-8000-000000000008', 1, 140, true),
  (gen_random_uuid(), 'Verduras fruto',     'verduras-fruto',     'aa000001-0000-4000-8000-000000000008', 1, 150, true),
  (gen_random_uuid(), 'Verduras hoja',      'verduras-hoja',      'aa000001-0000-4000-8000-000000000008', 1, 160, true),
  (gen_random_uuid(), 'Verduras raiz',      'verduras-raiz',      'aa000001-0000-4000-8000-000000000008', 1, 170, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- FIAMBRERÍA (parent: aa000001-0000-4000-8000-000000000009)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Conservas gourmet',   'conservas-gourmet',   'aa000001-0000-4000-8000-000000000009', 1, 10,  true),
  (gen_random_uuid(), 'Fiambres cocidos',    'fiambres-cocidos',    'aa000001-0000-4000-8000-000000000009', 1, 20,  true),
  (gen_random_uuid(), 'Fiambres curados',    'fiambres-curados',    'aa000001-0000-4000-8000-000000000009', 1, 30,  true),
  (gen_random_uuid(), 'Fiambres embutidos',  'fiambres-embutidos',  'aa000001-0000-4000-8000-000000000009', 1, 40,  true),
  (gen_random_uuid(), 'Jamones',            'jamones',            'aa000001-0000-4000-8000-000000000009', 1, 50,  true),
  (gen_random_uuid(), 'Lacteos frescos',     'lacteos-frescos',     'aa000001-0000-4000-8000-000000000009', 1, 60,  true),
  (gen_random_uuid(), 'Pan fiambreria',      'pan-fiambreria',      'aa000001-0000-4000-8000-000000000009', 1, 70,  true),
  (gen_random_uuid(), 'Picadas',            'picadas',            'aa000001-0000-4000-8000-000000000009', 1, 80,  true),
  (gen_random_uuid(), 'Quesos blandos',      'quesos-blandos',      'aa000001-0000-4000-8000-000000000009', 1, 90,  true),
  (gen_random_uuid(), 'Quesos duros',        'quesos-duros',        'aa000001-0000-4000-8000-000000000009', 1, 100, true),
  (gen_random_uuid(), 'Quesos frescos',      'quesos-frescos',      'aa000001-0000-4000-8000-000000000009', 1, 110, true),
  (gen_random_uuid(), 'Quesos semiduros',    'quesos-semiduros',    'aa000001-0000-4000-8000-000000000009', 1, 120, true),
  (gen_random_uuid(), 'Quesos untables',     'quesos-untables',     'aa000001-0000-4000-8000-000000000009', 1, 130, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- PILETAS (parent: aa000001-0000-4000-8000-000000000010)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Elevadores ph',      'elevadores-ph',      'aa000001-0000-4000-8000-000000000010', 1, 10, true),
  (gen_random_uuid(), 'Equipos filtracion', 'equipos-filtracion', 'aa000001-0000-4000-8000-000000000010', 1, 20, true),
  (gen_random_uuid(), 'Limpieza pileta',    'limpieza-pileta',    'aa000001-0000-4000-8000-000000000010', 1, 30, true),
  (gen_random_uuid(), 'Quimicos pileta',    'quimicos-pileta',    'aa000001-0000-4000-8000-000000000010', 1, 40, true),
  (gen_random_uuid(), 'Tabletas cloro',     'tabletas-cloro',     'aa000001-0000-4000-8000-000000000010', 1, 50, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- BAZAR (parent: aa000001-0000-4000-8000-000000000011)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Almacenamiento hogar', 'almacenamiento-hogar', 'aa000001-0000-4000-8000-000000000011', 1, 10,  true),
  (gen_random_uuid(), 'Cubiertos',           'cubiertos',           'aa000001-0000-4000-8000-000000000011', 1, 20,  true),
  (gen_random_uuid(), 'Iluminacion',         'iluminacion',         'aa000001-0000-4000-8000-000000000011', 1, 30,  true),
  (gen_random_uuid(), 'Limpieza',            'limpieza',            'aa000001-0000-4000-8000-000000000011', 1, 40,  true),
  (gen_random_uuid(), 'Mesa',               'mesa',               'aa000001-0000-4000-8000-000000000011', 1, 50,  true),
  (gen_random_uuid(), 'Organizacion',        'organizacion',        'aa000001-0000-4000-8000-000000000011', 1, 60,  true),
  (gen_random_uuid(), 'Organizacion hogar',  'organizacion-hogar',  'aa000001-0000-4000-8000-000000000011', 1, 70,  true),
  (gen_random_uuid(), 'Textiles hogar',      'textiles-hogar',      'aa000001-0000-4000-8000-000000000011', 1, 80,  true),
  (gen_random_uuid(), 'Textil hogar',        'textil-hogar',        'aa000001-0000-4000-8000-000000000011', 1, 90,  true),
  (gen_random_uuid(), 'Utensilios cocina',   'utensilios-cocina',   'aa000001-0000-4000-8000-000000000011', 1, 100, true),
  (gen_random_uuid(), 'Vajilla',            'vajilla',            'aa000001-0000-4000-8000-000000000011', 1, 110, true),
  (gen_random_uuid(), 'Vajilla vidrieria',   'vajilla-vidrieria',   'aa000001-0000-4000-8000-000000000011', 1, 120, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- JUGUETERÍA (parent: aa000001-0000-4000-8000-000000000012)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Adolescentes',     'adolescentes',     'aa000001-0000-4000-8000-000000000012', 1, 10,  true),
  (gen_random_uuid(), 'Aire libre',       'aire-libre',       'aa000001-0000-4000-8000-000000000012', 1, 20,  true),
  (gen_random_uuid(), 'Bebes',           'bebes',           'aa000001-0000-4000-8000-000000000012', 1, 30,  true),
  (gen_random_uuid(), 'Bebes 0 2',       'bebes-0-2',       'aa000001-0000-4000-8000-000000000012', 1, 40,  true),
  (gen_random_uuid(), 'Construccion',     'construccion',     'aa000001-0000-4000-8000-000000000012', 1, 50,  true),
  (gen_random_uuid(), 'Cotillon',        'cotillon',        'aa000001-0000-4000-8000-000000000012', 1, 60,  true),
  (gen_random_uuid(), 'Creativos',       'creativos',       'aa000001-0000-4000-8000-000000000012', 1, 70,  true),
  (gen_random_uuid(), 'Didacticos',      'didacticos',      'aa000001-0000-4000-8000-000000000012', 1, 80,  true),
  (gen_random_uuid(), 'Disfraces',       'disfraces',       'aa000001-0000-4000-8000-000000000012', 1, 90,  true),
  (gen_random_uuid(), 'Electronico',     'electronico',     'aa000001-0000-4000-8000-000000000012', 1, 100, true),
  (gen_random_uuid(), 'Juegos mesa',     'juegos-mesa',     'aa000001-0000-4000-8000-000000000012', 1, 110, true),
  (gen_random_uuid(), 'Munecos figuras', 'munecos-figuras', 'aa000001-0000-4000-8000-000000000012', 1, 120, true),
  (gen_random_uuid(), 'Ninos 3 5',      'ninos-3-5',      'aa000001-0000-4000-8000-000000000012', 1, 130, true),
  (gen_random_uuid(), 'Ninos 6 10',     'ninos-6-10',     'aa000001-0000-4000-8000-000000000012', 1, 140, true),
  (gen_random_uuid(), 'Vehiculos',       'vehiculos',       'aa000001-0000-4000-8000-000000000012', 1, 150, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- LIBRERÍA (parent: aa000001-0000-4000-8000-000000000013)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Adhesivos',           'adhesivos',           'aa000001-0000-4000-8000-000000000013', 1, 10,  true),
  (gen_random_uuid(), 'Arte',               'arte',               'aa000001-0000-4000-8000-000000000013', 1, 20,  true),
  (gen_random_uuid(), 'Arte pintura',        'arte-pintura',        'aa000001-0000-4000-8000-000000000013', 1, 30,  true),
  (gen_random_uuid(), 'Carpetas',           'carpetas',           'aa000001-0000-4000-8000-000000000013', 1, 40,  true),
  (gen_random_uuid(), 'Cuadernos',          'cuadernos',          'aa000001-0000-4000-8000-000000000013', 1, 50,  true),
  (gen_random_uuid(), 'Cuadernos carpetas', 'cuadernos-carpetas', 'aa000001-0000-4000-8000-000000000013', 1, 60,  true),
  (gen_random_uuid(), 'Escritura',          'escritura',          'aa000001-0000-4000-8000-000000000013', 1, 70,  true),
  (gen_random_uuid(), 'Juegos didacticos',  'juegos-didacticos',  'aa000001-0000-4000-8000-000000000013', 1, 80,  true),
  (gen_random_uuid(), 'Lapices lapiceras',  'lapices-lapiceras',  'aa000001-0000-4000-8000-000000000013', 1, 90,  true),
  (gen_random_uuid(), 'Mochilas',           'mochilas',           'aa000001-0000-4000-8000-000000000013', 1, 100, true),
  (gen_random_uuid(), 'Organizacion',       'organizacion',       'aa000001-0000-4000-8000-000000000013', 1, 110, true),
  (gen_random_uuid(), 'Papeleria',          'papeleria',          'aa000001-0000-4000-8000-000000000013', 1, 120, true),
  (gen_random_uuid(), 'Papeleria insumos',  'papeleria-insumos',  'aa000001-0000-4000-8000-000000000013', 1, 130, true),
  (gen_random_uuid(), 'Tecnologia basica',  'tecnologia-basica',  'aa000001-0000-4000-8000-000000000013', 1, 140, true),
  (gen_random_uuid(), 'Utiles escolares',   'utiles-escolares',   'aa000001-0000-4000-8000-000000000013', 1, 150, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- ROPA (parent: aa000001-0000-4000-8000-000000000014)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Buzos',               'buzos',               'aa000001-0000-4000-8000-000000000014', 1, 10,  true),
  (gen_random_uuid(), 'Calzado basico',      'calzado-basico',      'aa000001-0000-4000-8000-000000000014', 1, 20,  true),
  (gen_random_uuid(), 'Camperas',           'camperas',           'aa000001-0000-4000-8000-000000000014', 1, 30,  true),
  (gen_random_uuid(), 'Deportiva',          'deportiva',          'aa000001-0000-4000-8000-000000000014', 1, 40,  true),
  (gen_random_uuid(), 'Medias',             'medias',             'aa000001-0000-4000-8000-000000000014', 1, 50,  true),
  (gen_random_uuid(), 'Ninos',             'ninos',             'aa000001-0000-4000-8000-000000000014', 1, 60,  true),
  (gen_random_uuid(), 'Remeras',            'remeras',            'aa000001-0000-4000-8000-000000000014', 1, 70,  true),
  (gen_random_uuid(), 'Remeras basicas',    'remeras-basicas',    'aa000001-0000-4000-8000-000000000014', 1, 80,  true),
  (gen_random_uuid(), 'Ropa interior',      'ropa-interior',      'aa000001-0000-4000-8000-000000000014', 1, 90,  true),
  (gen_random_uuid(), 'Ropa interior hombre','ropa-interior-hombre','aa000001-0000-4000-8000-000000000014', 1, 100, true),
  (gen_random_uuid(), 'Ropa interior mujer', 'ropa-interior-mujer', 'aa000001-0000-4000-8000-000000000014', 1, 110, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- ELECTRODOMÉSTICOS (parent: aa000001-0000-4000-8000-000000000015)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Aires acondicionados', 'aires-acondicionados', 'aa000001-0000-4000-8000-000000000015', 1, 10,  true),
  (gen_random_uuid(), 'Calefaccion',          'calefaccion',          'aa000001-0000-4000-8000-000000000015', 1, 20,  true),
  (gen_random_uuid(), 'Climatizacion',        'climatizacion',        'aa000001-0000-4000-8000-000000000015', 1, 30,  true),
  (gen_random_uuid(), 'Cocinas',             'cocinas',             'aa000001-0000-4000-8000-000000000015', 1, 40,  true),
  (gen_random_uuid(), 'Heladeras',           'heladeras',           'aa000001-0000-4000-8000-000000000015', 1, 50,  true),
  (gen_random_uuid(), 'Iluminacion led',     'iluminacion-led',     'aa000001-0000-4000-8000-000000000015', 1, 60,  true),
  (gen_random_uuid(), 'Lavarropas',          'lavarropas',          'aa000001-0000-4000-8000-000000000015', 1, 70,  true),
  (gen_random_uuid(), 'Linea blanca pequena','linea-blanca-pequena','aa000001-0000-4000-8000-000000000015', 1, 80,  true),
  (gen_random_uuid(), 'Pequenos cocina',     'pequenos-cocina',     'aa000001-0000-4000-8000-000000000015', 1, 90,  true),
  (gen_random_uuid(), 'Pequenos limpieza',   'pequenos-limpieza',   'aa000001-0000-4000-8000-000000000015', 1, 100, true),
  (gen_random_uuid(), 'Tv audio',            'tv-audio',            'aa000001-0000-4000-8000-000000000015', 1, 110, true),
  (gen_random_uuid(), 'Ventilacion',         'ventilacion',         'aa000001-0000-4000-8000-000000000015', 1, 120, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- CORRALÓN (parent: aa000001-0000-4000-8000-000000000016)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Aberturas',            'aberturas',            'aa000001-0000-4000-8000-000000000016', 1, 10,  true),
  (gen_random_uuid(), 'Adhesivos mezclas',    'adhesivos-mezclas',    'aa000001-0000-4000-8000-000000000016', 1, 20,  true),
  (gen_random_uuid(), 'Aislantes',           'aislantes',           'aa000001-0000-4000-8000-000000000016', 1, 30,  true),
  (gen_random_uuid(), 'Canierias',           'canierias',           'aa000001-0000-4000-8000-000000000016', 1, 40,  true),
  (gen_random_uuid(), 'Canos acero',          'canos-acero',          'aa000001-0000-4000-8000-000000000016', 1, 50,  true),
  (gen_random_uuid(), 'Cemento cal',          'cemento-cal',          'aa000001-0000-4000-8000-000000000016', 1, 60,  true),
  (gen_random_uuid(), 'Cementos cales',       'cementos-cales',       'aa000001-0000-4000-8000-000000000016', 1, 70,  true),
  (gen_random_uuid(), 'Ceramicos porcelanato','ceramicos-porcelanato','aa000001-0000-4000-8000-000000000016', 1, 80,  true),
  (gen_random_uuid(), 'Chapas acanaladas',    'chapas-acanaladas',    'aa000001-0000-4000-8000-000000000016', 1, 90,  true),
  (gen_random_uuid(), 'Hierro acero',         'hierro-acero',         'aa000001-0000-4000-8000-000000000016', 1, 100, true),
  (gen_random_uuid(), 'Hierro redondo',       'hierro-redondo',       'aa000001-0000-4000-8000-000000000016', 1, 110, true),
  (gen_random_uuid(), 'Ladrillos',           'ladrillos',           'aa000001-0000-4000-8000-000000000016', 1, 120, true),
  (gen_random_uuid(), 'Ladrillos bloques',    'ladrillos-bloques',    'aa000001-0000-4000-8000-000000000016', 1, 130, true),
  (gen_random_uuid(), 'Madera molduras',      'madera-molduras',      'aa000001-0000-4000-8000-000000000016', 1, 140, true),
  (gen_random_uuid(), 'Maderas',             'maderas',             'aa000001-0000-4000-8000-000000000016', 1, 150, true),
  (gen_random_uuid(), 'Mallas electrosoldadas','mallas-electrosoldadas','aa000001-0000-4000-8000-000000000016', 1, 160, true),
  (gen_random_uuid(), 'Membranas chapas',     'membranas-chapas',     'aa000001-0000-4000-8000-000000000016', 1, 170, true),
  (gen_random_uuid(), 'Perfiles aluminio',    'perfiles-aluminio',    'aa000001-0000-4000-8000-000000000016', 1, 180, true),
  (gen_random_uuid(), 'Pintura',             'pintura',             'aa000001-0000-4000-8000-000000000016', 1, 190, true),
  (gen_random_uuid(), 'Pisos revestimientos', 'pisos-revestimientos', 'aa000001-0000-4000-8000-000000000016', 1, 200, true),
  (gen_random_uuid(), 'Revoques yesos',       'revoques-yesos',       'aa000001-0000-4000-8000-000000000016', 1, 210, true),
  (gen_random_uuid(), 'Sanitarios',          'sanitarios',          'aa000001-0000-4000-8000-000000000016', 1, 220, true),
  (gen_random_uuid(), 'Techos',             'techos',             'aa000001-0000-4000-8000-000000000016', 1, 230, true),
  (gen_random_uuid(), 'Tirantes madera',     'tirantes-madera',     'aa000001-0000-4000-8000-000000000016', 1, 240, true)
ON CONFLICT (slug) DO NOTHING;

-- -----------------------------------------------------------
-- VETERINARIA (parent: aa000001-0000-4000-8000-000000000017)
-- -----------------------------------------------------------
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active) VALUES
  (gen_random_uuid(), 'Accesorios gato',  'accesorios-gato',  'aa000001-0000-4000-8000-000000000017', 1, 10,  true),
  (gen_random_uuid(), 'Accesorios perro', 'accesorios-perro', 'aa000001-0000-4000-8000-000000000017', 1, 20,  true),
  (gen_random_uuid(), 'Alimento gato',    'alimento-gato',    'aa000001-0000-4000-8000-000000000017', 1, 30,  true),
  (gen_random_uuid(), 'Alimento humedo',  'alimento-humedo',  'aa000001-0000-4000-8000-000000000017', 1, 40,  true),
  (gen_random_uuid(), 'Alimento perro',   'alimento-perro',   'aa000001-0000-4000-8000-000000000017', 1, 50,  true),
  (gen_random_uuid(), 'Higiene mascotas', 'higiene-mascotas', 'aa000001-0000-4000-8000-000000000017', 1, 60,  true),
  (gen_random_uuid(), 'Medicamentos otc', 'medicamentos-otc', 'aa000001-0000-4000-8000-000000000017', 1, 70,  true),
  (gen_random_uuid(), 'Otros animales',   'otros-animales',   'aa000001-0000-4000-8000-000000000017', 1, 80,  true),
  (gen_random_uuid(), 'Ropa mascotas',    'ropa-mascotas',    'aa000001-0000-4000-8000-000000000017', 1, 90,  true),
  (gen_random_uuid(), 'Snacks mascotas',  'snacks-mascotas',  'aa000001-0000-4000-8000-000000000017', 1, 100, true)
ON CONFLICT (slug) DO NOTHING;


-- =============================================================================
-- PARTE 2 (extra) — Slugs normalizados que necesitan categoría propia
-- =============================================================================

-- lacteos → almacén
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active)
VALUES (gen_random_uuid(), 'Lacteos', 'lacteos', 'aa000001-0000-4000-8000-000000000001', 1, 330, true)
ON CONFLICT (slug) DO NOTHING;

-- bebidas → kiosco
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active)
VALUES (gen_random_uuid(), 'Bebidas', 'bebidas', 'aa000001-0000-4000-8000-000000000004', 1, 260, true)
ON CONFLICT (slug) DO NOTHING;

-- golosinas → kiosco
INSERT INTO marketplace_categories (id, name, slug, parent_id, level, sort_order, is_active)
VALUES (gen_random_uuid(), 'Golosinas', 'golosinas', 'aa000001-0000-4000-8000-000000000004', 1, 270, true)
ON CONFLICT (slug) DO NOTHING;


-- =============================================================================
-- PARTE 3 — Normalizar slugs problemáticos en global_products
-- Solo corrige casos donde la categoria existe en global_products con valor
-- incorrecto y no tiene match en marketplace_categories por ese slug.
-- =============================================================================

-- Corregir slugs con mayúsculas / acentos / valores no normalizados
UPDATE global_products
SET category = 'lacteos'
WHERE category = 'Lácteos'
  AND business_type = 'almacen';

UPDATE global_products
SET category = 'harinas-panaderia'
WHERE category = 'Panadería'
  AND business_type = 'almacen';

UPDATE global_products
SET category = 'bebidas'
WHERE category = 'Bebidas'
  AND business_type = 'kiosco';

UPDATE global_products
SET category = 'gaseosas'
WHERE category = 'beverages-and-beverages-preparations'
  AND business_type = 'almacen';


-- =============================================================================
-- PARTE 4 — Corregir mismatches de slugs en business_type_templates.categories
-- Los UPDATE usan jsonb_agg + CASE — idempotentes (aplicar el mismo cambio dos
-- veces produce el mismo resultado porque el slug ya fue corregido).
-- =============================================================================

-- ------------------------------------------------------------
-- ALMACÉN (b2000001-0000-4000-8000-000000000001)
-- Corrige 7 slugs que no coinciden con global_products
-- ------------------------------------------------------------
UPDATE business_type_templates
SET categories = (
  SELECT jsonb_agg(
    CASE
      WHEN cat->>'slug' = 'galletitas'       THEN jsonb_set(cat, '{slug}', '"galletitas-dulces"')
      WHEN cat->>'slug' = 'bebidas'          THEN jsonb_set(cat, '{slug}', '"gaseosas-aguas"')
      WHEN cat->>'slug' = 'conservas'        THEN jsonb_set(cat, '{slug}', '"conservas-enlatados"')
      WHEN cat->>'slug' = 'lacteos'          THEN jsonb_set(cat, '{slug}', '"quesos-manteca"')
      WHEN cat->>'slug' = 'limpieza'         THEN jsonb_set(cat, '{slug}', '"lavandina-desinfectantes"')
      WHEN cat->>'slug' = 'pastas-harinas'   THEN jsonb_set(cat, '{slug}', '"pastas-secas"')
      WHEN cat->>'slug' = 'perfumeria-basica'THEN jsonb_set(cat, '{slug}', '"higiene-personal"')
      ELSE cat
    END
  )
  FROM jsonb_array_elements(categories) AS cat
)
WHERE id = 'b2000001-0000-4000-8000-000000000001';

-- ------------------------------------------------------------
-- CARNICERÍA (b2000001-0000-4000-8000-000000000006)
-- vacuno → cortes-vacunos
-- ------------------------------------------------------------
UPDATE business_type_templates
SET categories = (
  SELECT jsonb_agg(
    CASE
      WHEN cat->>'slug' = 'vacuno' THEN jsonb_set(cat, '{slug}', '"cortes-vacunos"')
      ELSE cat
    END
  )
  FROM jsonb_array_elements(categories) AS cat
)
WHERE id = 'b2000001-0000-4000-8000-000000000006';

-- ------------------------------------------------------------
-- FERRETERÍA (b2000001-0000-4000-8000-000000000003)
-- Corrige slugs genéricos que no matchean global_products
-- ------------------------------------------------------------
UPDATE business_type_templates
SET categories = (
  SELECT jsonb_agg(
    CASE
      WHEN cat->>'slug' = 'pinturas'         THEN jsonb_set(cat, '{slug}', '"pinturas-latex"')
      WHEN cat->>'slug' = 'herramientas'     THEN jsonb_set(cat, '{slug}', '"herramientas-manuales"')
      WHEN cat->>'slug' = 'plomeria'         THEN jsonb_set(cat, '{slug}', '"sanitarios-plomeria"')
      WHEN cat->>'slug' = 'electrica'        THEN jsonb_set(cat, '{slug}', '"cables-electricidad"')
      WHEN cat->>'slug' = 'adhesivos'        THEN jsonb_set(cat, '{slug}', '"adhesivos-selladores"')
      ELSE cat
    END
  )
  FROM jsonb_array_elements(categories) AS cat
)
WHERE id = 'b2000001-0000-4000-8000-000000000003';

-- ------------------------------------------------------------
-- KIOSCO (b2000001-0000-4000-8000-000000000002)
-- Corrige slugs de bebidas y golosinas
-- ------------------------------------------------------------
UPDATE business_type_templates
SET categories = (
  SELECT jsonb_agg(
    CASE
      WHEN cat->>'slug' = 'bebidas-frias'    THEN jsonb_set(cat, '{slug}', '"gaseosas"')
      WHEN cat->>'slug' = 'golosinas-varias' THEN jsonb_set(cat, '{slug}', '"golosinas"')
      WHEN cat->>'slug' = 'snacks'           THEN jsonb_set(cat, '{slug}', '"palitos-chizitos"')
      ELSE cat
    END
  )
  FROM jsonb_array_elements(categories) AS cat
)
WHERE id = 'b2000001-0000-4000-8000-000000000002';

-- ------------------------------------------------------------
-- PANADERÍA (b2000001-0000-4000-8000-000000000004)
-- harinas → harinas-panaderia, insumos → insumos-panaderia
-- ------------------------------------------------------------
UPDATE business_type_templates
SET categories = (
  SELECT jsonb_agg(
    CASE
      WHEN cat->>'slug' = 'harinas'  THEN jsonb_set(cat, '{slug}', '"harinas-panaderia"')
      WHEN cat->>'slug' = 'insumos'  THEN jsonb_set(cat, '{slug}', '"insumos-panaderia"')
      ELSE cat
    END
  )
  FROM jsonb_array_elements(categories) AS cat
)
WHERE id = 'b2000001-0000-4000-8000-000000000004';

-- ------------------------------------------------------------
-- VERDULERÍA (b2000001-0000-4000-8000-000000000005)
-- frutas → frutas-frescas, verduras → verduras-hoja
-- ------------------------------------------------------------
UPDATE business_type_templates
SET categories = (
  SELECT jsonb_agg(
    CASE
      WHEN cat->>'slug' = 'frutas'   THEN jsonb_set(cat, '{slug}', '"frutas-frescas"')
      WHEN cat->>'slug' = 'verduras' THEN jsonb_set(cat, '{slug}', '"verduras-hoja"')
      ELSE cat
    END
  )
  FROM jsonb_array_elements(categories) AS cat
)
WHERE id = 'b2000001-0000-4000-8000-000000000005';


-- =============================================================================
-- VERIFICACIÓN POST-SEED
-- Resultado esperado: 0 o muy cercano a 0
-- =============================================================================
SELECT COUNT(DISTINCT gp.category) AS gp_cats_sin_mc
FROM global_products gp
LEFT JOIN marketplace_categories mc ON mc.slug = gp.category
WHERE gp.is_active = true
  AND mc.id IS NULL;
