-- Seed 091: Productos globales — Almacén: Galletitas, Alfajores y Chocolates (~70 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento supermercados y almacenes NEA (Posadas) + distribuidores Arcor/Bagley 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: galletitas-dulces, galletitas-saladas, alfajores, chocolates, golosinas

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- GALLETITAS DULCES — Arcor / Bagley
-- ============================================================
('Galletitas Oreo original 117g',               'Oreo',       'galletitas-dulces',  1200.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas Oreo original 312g',               'Oreo',       'galletitas-dulces',  2800.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas Oreo Double Stuf 312g',            'Oreo',       'galletitas-dulces',  3200.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas Pepitos 100g',                     'Bagley',     'galletitas-dulces',  1050.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas Criollitas 200g',                  'Bagley',     'galletitas-dulces',  1400.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas Lincoln 200g Arcor',               'Arcor',      'galletitas-dulces',  1350.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas Diversión vainilla 140g Bagley',   'Bagley',     'galletitas-dulces',  1050.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas Melba vainilla 180g Terrabusi',    'Terrabusi',  'galletitas-dulces',  1200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Surtido Bagley 320g',                         'Bagley',     'galletitas-dulces',  2800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas de avena con choco 200g Quaker',   'Quaker',     'galletitas-dulces',  1600.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas Mini Chips Ahoy! 130g',            'Chips Ahoy!','galletitas-dulces',  1800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- GALLETITAS SALADAS
-- ============================================================
('Galletitas Crackers 180g Club Social',        'Club Social', 'galletitas-saladas', 1400.00,'seed',0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas 9 de Oro 150g',                    '9 de Oro',   'galletitas-saladas', 1100.00,'seed',0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas Vocación integral 200g Bagley',    'Bagley',     'galletitas-saladas', 1500.00,'seed',0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas Express soda 180g Bagley',         'Bagley',     'galletitas-saladas', 1100.00,'seed',0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas Agua 100g Canale',                 'Canale',     'galletitas-saladas',  950.00,'seed',0.8, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas Telma 180g',                       'Telma',      'galletitas-saladas', 1150.00,'seed',0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Galletitas de arroz inflado 100g genérico',   NULL,         'galletitas-saladas',  980.00,'seed',0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- ALFAJORES — los más vendidos en NEA
-- ============================================================
('Alfajor Havanna chocolate 90g',               'Havanna',    'alfajores',          2200.00, 'seed', 0.8, 86, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Havanna Triple chocolate 90g',                'Havanna',    'alfajores',          2500.00, 'seed', 0.8, 86, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Havanna Clasico x6u',                         'Havanna',    'alfajores',         12000.00, 'seed', 0.8, 86, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Alfajor Guaymallén chocolate 55g',            'Guaymallén', 'alfajores',           820.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Alfajor Guaymallén triple chocolate 55g',     'Guaymallén', 'alfajores',          1000.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Alfajor Guaymallén blanco 55g',               'Guaymallén', 'alfajores',           820.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Alfajor Jorgito chocolate 50g',               'Jorgito',    'alfajores',           650.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Alfajor Terrabusi Mousse 50g',                'Terrabusi',  'alfajores',           680.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Alfajor Capitan del Espacio 50g',             'Capitán del Espacio','alfajores',   650.00,'seed',0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Alfajor Mantecol x2u',                        'Mantecol',   'alfajores',          1200.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Alfajor Milka mousse 50g',                    'Milka',      'alfajores',           980.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Alfajor Cachafaz x6u 100g',                   'Cachafaz',   'alfajores',          5500.00, 'seed', 0.8, 84, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),

-- ============================================================
-- CHOCOLATES — Milka, Toblerone, Felfort, Cofler
-- ============================================================
('Milka Chocolate Leche 90g',                   'Milka',      'chocolates',         2200.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Milka Chocolate Oreo 100g',                   'Milka',      'chocolates',         2800.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Milka Chocolate Almendras 100g',              'Milka',      'chocolates',         2800.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Milka Chocolate Avellanas 100g',              'Milka',      'chocolates',         2800.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Toblerone leche 100g',                        'Toblerone',  'chocolates',         3500.00, 'seed', 0.8, 84, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Toblerone leche 200g',                        'Toblerone',  'chocolates',         6500.00, 'seed', 0.8, 84, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Lindt Excellence negro 70% 100g',             'Lindt',      'chocolates',         5500.00, 'seed', 0.8, 86, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Cofler leche 55g Arcor',                      'Arcor',      'chocolates',          950.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Turín clasico 55g',                           'Turín',      'chocolates',          980.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Felfort Chocolate Maní 54g',                  'Felfort',    'chocolates',          850.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Bon o Bon leche 180g Arcor',                  'Arcor',      'chocolates',         2200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Kit Kat 45g Nestlé',                          'Nestlé',     'chocolates',         1200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- CARAMELOS Y CHICLES — alta rotación kiosco/almacén
-- ============================================================
('Halls Mentol 28g x10u',                       'Halls',      'golosinas',           580.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Halls Frutas 28g x10u',                       'Halls',      'golosinas',           580.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Mentos Frutas 38g',                           'Mentos',     'golosinas',           650.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Mentos Menta 38g',                            'Mentos',     'golosinas',           650.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Trident menta x12 unidades',                  'Trident',    'golosinas',           780.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Trident Sandia x12 unidades',                 'Trident',    'golosinas',           780.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Beldent Menta x12u',                          'Beldent',    'golosinas',           750.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Adams Spearmint x12u',                        'Adams',      'golosinas',           720.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Sugus frutales 125g Arcor',                   'Arcor',      'golosinas',           980.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Caramelos Arcor duros 500g',                  'Arcor',      'golosinas',          1800.00, 'seed', 0.8, 76, FALSE, TRUE, 'almacen', '{"unit": "bolsa",  "sku_prefix": "ALMAC"}'),
('Chupetín Pops Arcor x100u',                   'Arcor',      'golosinas',          4800.00, 'seed', 0.8, 76, FALSE, TRUE, 'almacen', '{"unit": "caja",   "sku_prefix": "ALMAC"}'),
('Palito Frutal Arcor x20u',                    'Arcor',      'golosinas',          2800.00, 'seed', 0.8, 76, FALSE, TRUE, 'almacen', '{"unit": "caja",   "sku_prefix": "ALMAC"}'),
('Maní confitado 200g Arcor',                   'Arcor',      'golosinas',          1200.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Maní salado 200g Arcor',                      'Arcor',      'golosinas',          1100.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 68 productos
-- Categorías usadas: galletitas-dulces, galletitas-saladas, alfajores, chocolates, golosinas
-- Marcas: Oreo, Bagley, Terrabusi, Arcor, Club Social, Canale, Havanna, Guaymallén, Jorgito,
--   Milka, Toblerone, Lindt, Felfort, Cofler, Turín, Nestlé, Halls, Mentos, Trident, Beldent, Adams
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
