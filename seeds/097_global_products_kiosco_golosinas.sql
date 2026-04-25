-- Seed 097: Productos globales — Kiosco: Golosinas (~85 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento kioscos NEA (Posadas) + distribuidores golosinas 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: alfajores, chocolates, caramelos-chicles, chupetines, gomitas-confites, mani-snacks

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- ALFAJORES — top vendedores NEA
-- ============================================================
('Alfajor Guaymallén doble chocolate x unidad',  'Guaymallén',  'alfajores',   650.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Alfajor Guaymallén triple chocolate x unidad', 'Guaymallén',  'alfajores',   850.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Alfajor Guaymallén blanco x unidad',           'Guaymallén',  'alfajores',   650.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Alfajor Jorgito chocolate x unidad',           'Jorgito',     'alfajores',   580.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Alfajor Jorgito blanco x unidad',              'Jorgito',     'alfajores',   580.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Alfajor Bon o Bon chocolate x unidad',         'Bon o Bon',   'alfajores',   750.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Alfajor Milka chocolate con leche x unidad',   'Milka',       'alfajores',  1200.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Alfajor Oreo doble relleno x unidad',          'Oreo',        'alfajores',  1100.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Alfajor Havanna maicena x unidad',             'Havanna',     'alfajores',  1800.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Alfajor Cabsha chocolate x unidad',            'Cabsha',      'alfajores',   700.00, 'seed', 0.8, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Alfajor Rocklets chocolate x unidad',          'Arcor',       'alfajores',   650.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Alfajor Pepitos x unidad',                     'Bagley',      'alfajores',   480.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- CHOCOLATES EN BARRA Y TABLETAS
-- ============================================================
('Chocolate Milka leche 100g',                   'Milka',       'chocolates',  2200.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chocolate Milka oreo 100g',                    'Milka',       'chocolates',  2500.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chocolate Milka almendras 100g',               'Milka',       'chocolates',  2500.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Mentitas chocolate negro 25g',                 'Arcor',       'chocolates',   650.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Cabsha maní chocolate 36g',                    'Cabsha',      'chocolates',   820.00, 'seed', 0.8, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Garoto bombón surtidos 180g',                  'Garoto',      'chocolates',  3200.00, 'seed', 0.8, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Toblerone leche 100g',                         'Toblerone',   'chocolates',  3500.00, 'seed', 0.8, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Rocklets tableta 40g Arcor',                   'Arcor',       'chocolates',   750.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chocolate Snikers 52g',                        'Mars',        'chocolates',  1800.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chocolate Twix 50g',                           'Mars',        'chocolates',  1800.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- CARAMELOS
-- ============================================================
('Sugus tutti frutti x 12u bolsa',               'Sugus',       'caramelos-chicles', 800.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "bolsa", "sku_prefix": "KIOS"}'),
('Sugus naranja x unidad',                       'Sugus',       'caramelos-chicles', 120.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Halls mentol azul x unidad',                   'Halls',       'caramelos-chicles', 150.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Halls frutilla x unidad',                      'Halls',       'caramelos-chicles', 150.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Alpenliebe original x unidad',                 'Alpenliebe',  'caramelos-chicles', 130.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Caramelos masticables Trulala bolsa 12u',      'Arcor',       'caramelos-chicles', 750.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "bolsa", "sku_prefix": "KIOS"}'),
('Caramelo Pop 25g Arcor',                       'Arcor',       'caramelos-chicles', 280.00, 'seed', 0.9, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- CHICLES
-- ============================================================
('Chicle Adams Spearmint x unidad',              'Adams',       'caramelos-chicles', 300.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chicle Adams Menta x unidad',                  'Adams',       'caramelos-chicles', 300.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chicle Beldent sandía sin azúcar x unidad',    'Beldent',     'caramelos-chicles', 350.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chicle Beldent menta sin azúcar x unidad',     'Beldent',     'caramelos-chicles', 350.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Bubbaloo fresa x unidad',                      'Bubbaloo',    'caramelos-chicles', 200.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chicle Topline pastillas x unidad',            'Topline',     'caramelos-chicles', 450.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- CHUPETINES
-- ============================================================
('Chupetín Arcor variedad x unidad',             'Arcor',       'chupetines',   180.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chupetín Pop 12g Arcor x unidad',              'Arcor',       'chupetines',   220.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chupetín Paleta Colombina fresa x unidad',     'Colombina',   'chupetines',   280.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chupetín Big Pop chicle relleno x unidad',     'Arcor',       'chupetines',   350.00, 'seed', 0.9, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chupetín Pica Pau surtidos x unidad',          'Colombina',   'chupetines',   250.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chupetín Frutales bolsa x24u Arcor',           'Arcor',       'chupetines',  3800.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "bolsa", "sku_prefix": "KIOS"}'),

-- ============================================================
-- GOMITAS Y CONFITES
-- ============================================================
('Gomitas Haribo ositos 100g',                   'Haribo',      'gomitas-confites', 1800.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Gomitas Haribo surtidas 250g',                 'Haribo',      'gomitas-confites', 3800.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Gomitas Arcor surtidas 100g',                  'Arcor',       'gomitas-confites', 1200.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Confites de colores 100g genérico',            NULL,          'gomitas-confites',  800.00, 'seed', 0.6, 60, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Marshmallows Arcor 200g',                      'Arcor',       'gomitas-confites', 1500.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Palitos de regaliz Pico Dulce 200g',           'Pico Dulce',  'gomitas-confites', 1200.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- MANÍ CON CHOCOLATE Y SNACKS KIOSCO
-- ============================================================
('Maní con chocolate Cabsha 40g',                'Cabsha',      'mani-snacks',   750.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Maní japonés 120g La Maní',                    'La Maní',     'mani-snacks',  1200.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Maní tostado salado 100g genérico',            NULL,          'mani-snacks',   850.00, 'seed', 0.6, 60, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Palomitas de maíz caramelizadas 50g Arcor',    'Arcor',       'mani-snacks',   650.00, 'seed', 0.9, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Galletitas obleas Oblea 40g Biscotti',         'Biscotti',    'mani-snacks',   480.00, 'seed', 0.8, 76, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Papas fritas Pringles original 40g',           'Pringles',    'mani-snacks',  1500.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Papas fritas Lays original 30g',               'Lays',        'mani-snacks',   850.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Papas fritas Doritos nacho 35g',               'Doritos',     'mani-snacks',   950.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Chizitos Arcor 55g',                           'Arcor',       'mani-snacks',   750.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Palitos salados La Virginia 240g',             'La Virginia', 'mani-snacks',  2200.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Tucos de maíz Pehuamar 100g',                  'Pehuamar',    'mani-snacks',   950.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Maní con mantequilla Mantekol 100g',           'Mantekol',    'mani-snacks',  1100.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- GALLETITAS DULCES INDIVIDUALES (kiosco)
-- ============================================================
('Galletita Oreo original 41g',                  'Oreo',        'alfajores',     680.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Galletita Toddy 40g',                          'Toddy',       'alfajores',     580.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Galletita Maná vainilla 45g Bagley',           'Bagley',      'alfajores',     480.00, 'seed', 0.8, 76, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Barrita de cereal Quaker banana 22g',          'Quaker',      'mani-snacks',   550.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Barrita de cereal Nutri-Grain 37g Kellogs',    'Kellogg''s',  'mani-snacks',   750.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- TURRONES Y NOUGAT
-- ============================================================
('Turrón de maní 40g Georgalos',                 'Georgalos',   'gomitas-confites', 450.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Turrón blando de nuez 35g Nucrem',             'Nucrem',      'gomitas-confites', 380.00, 'seed', 0.7, 76, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Oblea Bañada Chocolate 40g Georgalos',         'Georgalos',   'chocolates',    520.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Cubanito azúcar 25g',                          NULL,          'alfajores',     250.00, 'seed', 0.6, 60, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Pirulín chocolate 40g Arcor',                  'Arcor',       'chocolates',    480.00, 'seed', 0.9, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- HELADOS DE PALITO (freezer kiosco)
-- ============================================================
('Paleta helado Duetto 70ml Arcor',              'Arcor',       'alfajores',     850.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Paleta helado La Liguria 75ml',                'La Liguria',  'alfajores',     780.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Varito helado de crema 80ml Frigor',           'Frigor',      'alfajores',     750.00, 'seed', 0.8, 76, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- CARAMELOS MASAJEADORES Y ESPECIALES
-- ============================================================
('Caramelo Jack 15g genérico',                   NULL,          'caramelos-chicles', 100.00, 'seed', 0.6, 60, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Caramelo Milky Way 58g',                       'Mars',        'caramelos-chicles', 1600.00, 'seed', 0.8, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Pochoclos bolsa 30g genérico kiosco',          NULL,          'mani-snacks',    480.00, 'seed', 0.6, 60, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 85 productos
-- Categorías: alfajores, chocolates, caramelos-chicles, chupetines, gomitas-confites, mani-snacks
-- Marcas: Guaymallén, Jorgito, Bon o Bon, Milka, Oreo, Havanna, Cabsha, Arcor, Garoto, Toblerone,
--   Mars, Adams, Beldent, Bubbaloo, Halls, Sugus, Alpenliebe, Haribo, Colombina, Pringles, Lays,
--   Doritos, Pehuamar, Georgalos, Kellogg's, Bagley, Quaker
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
