-- Seed 093: Productos globales — Almacén: Vinos Populares de Almacén (~30 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento almacenes y súper NEA (Posadas) + distribuidores vinos masivos 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- NOTA: Estos son vinos de alta rotación en almacén (no vinoteca). Precios de góndola masivo.
-- CATEGORÍAS: vinos-tintos, vinos-blancos, vinos-rosados, espumantes

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- VINOS GATO NEGRO — el más vendido en almacenes NEA
-- ============================================================
('Gato Negro Tinto 750ml',                      'Gato Negro',  'vinos-tintos',    3200.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Gato Negro Blanco 750ml',                     'Gato Negro',  'vinos-blancos',   3200.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Gato Negro Rosado 750ml',                     'Gato Negro',  'vinos-rosados',   3200.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Gato Negro Malbec 750ml',                     'Gato Negro',  'vinos-tintos',    3500.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Gato Negro Cabernet Sauvignon 750ml',         'Gato Negro',  'vinos-tintos',    3500.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Gato Negro Tetra 1L Tinto',                   'Gato Negro',  'vinos-tintos',    2800.00, 'seed', 0.8, 74, FALSE, TRUE, 'almacen', '{"unit": "tetra",   "sku_prefix": "ALMAC"}'),
('Gato Negro Tetra 1L Blanco',                  'Gato Negro',  'vinos-blancos',   2800.00, 'seed', 0.8, 74, FALSE, TRUE, 'almacen', '{"unit": "tetra",   "sku_prefix": "ALMAC"}'),

-- ============================================================
-- TRAPICHE — versiones económicas para almacén
-- ============================================================
('Trapiche Malbec Roble 750ml',                 'Trapiche',    'vinos-tintos',    4800.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Trapiche Vineyards Malbec 750ml',             'Trapiche',    'vinos-tintos',    3500.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Trapiche Chardonnay Vineyards 750ml',         'Trapiche',    'vinos-blancos',   3500.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- NORTON Y SANTA JULIA — populares en almacén
-- ============================================================
('Norton Malbec 750ml',                         'Norton',      'vinos-tintos',    5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Santa Julia Malbec 750ml',                    'Santa Julia', 'vinos-tintos',    4200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Santa Julia Torrontés 750ml',                 'Santa Julia', 'vinos-blancos',   4200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- CRIOS — populares en almacén NEA
-- ============================================================
('Crios Torrontés 750ml',                       'Crios',       'vinos-blancos',   7500.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- VINO EN CAJA — alta rotación en almacén
-- ============================================================
('Vino tinto caja 3L genérico',                 NULL,          'vinos-tintos',    7500.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "caja",    "sku_prefix": "ALMAC"}'),
('Vino blanco caja 3L genérico',                NULL,          'vinos-blancos',   7500.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "caja",    "sku_prefix": "ALMAC"}'),
('Toro Viejo tinto 3L',                         'Toro Viejo',  'vinos-tintos',    6500.00, 'seed', 0.7, 70, FALSE, TRUE, 'almacen', '{"unit": "caja",    "sku_prefix": "ALMAC"}'),
('Tocornal tinto 3L',                           'Tocornal',    'vinos-tintos',    5500.00, 'seed', 0.7, 70, FALSE, TRUE, 'almacen', '{"unit": "caja",    "sku_prefix": "ALMAC"}'),

-- ============================================================
-- ESPUMANTES BÁSICOS — para almacén
-- ============================================================
('Chandon Extra Brut 750ml',                    'Chandon',     'espumantes',     12000.00, 'seed', 0.8, 84, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Rainier Extra Brut 750ml',                    'Rainier',     'espumantes',      7500.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Valentin Bianchi Extra Brut 750ml',           'Valentin Bianchi','espumantes',  9500.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "botella", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- FERNET Y APERITIVOS BÁSICOS (de compra frecuente en almacén)
-- ============================================================
('Fernet Branca 750ml',                         'Fernet Branca','aperitivos-licores',22000.00,'seed',0.8,82,FALSE,TRUE,'almacen','{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Campari 750ml',                               'Campari',      'aperitivos-licores',22000.00,'seed',0.8,82,FALSE,TRUE,'almacen','{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Cynar 750ml',                                 'Cynar',        'aperitivos-licores',14000.00,'seed',0.8,80,FALSE,TRUE,'almacen','{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Aperol 750ml',                                'Aperol',       'aperitivos-licores',22000.00,'seed',0.8,82,FALSE,TRUE,'almacen','{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Gancia Quinado 950ml',                        'Gancia',       'aperitivos-licores', 8500.00,'seed',0.7,78,FALSE,TRUE,'almacen','{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Smirnoff Vodka Red 750ml',                    'Smirnoff',     'aperitivos-licores',14000.00,'seed',0.8,80,FALSE,TRUE,'almacen','{"unit": "botella", "sku_prefix": "ALMAC"}'),
('Johnnie Walker Red Label 750ml',              'Johnnie Walker','aperitivos-licores',18000.00,'seed',0.8,84,FALSE,TRUE,'almacen','{"unit": "botella", "sku_prefix": "ALMAC"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 29 productos
-- Categorías usadas: vinos-tintos, vinos-blancos, vinos-rosados, espumantes, aperitivos-licores
-- Nota: business_type = 'almacen' — diferente al seed 087-088 que son 'vinoteca'
-- Marcas: Gato Negro, Trapiche, Norton, Santa Julia, Crios, Chandon, Rainier,
--   Fernet Branca, Campari, Cynar, Aperol, Gancia, Smirnoff, Johnnie Walker
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
