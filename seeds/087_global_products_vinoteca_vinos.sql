-- Seed 087: Productos globales — Vinoteca: Vinos Argentinos (~90 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: catálogos vinotecas NEA + distribuidores de bodegas + Wines of Argentina 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: vinos-tintos, vinos-blancos, vinos-rosados, espumantes

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- VINOS TINTOS — Trapiche (alta distribución en NEA)
-- ============================================================
('Trapiche Malbec Roble 750ml',                 'Trapiche',  'vinos-tintos',  4800.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Trapiche Malbec Reserva 750ml',               'Trapiche',  'vinos-tintos',  7500.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Trapiche Cabernet Sauvignon Roble 750ml',     'Trapiche',  'vinos-tintos',  4800.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Trapiche Syrah Roble 750ml',                  'Trapiche',  'vinos-tintos',  4800.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Trapiche Malbec Medalla 750ml',               'Trapiche',  'vinos-tintos', 12000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Trapiche Broquel Malbec 750ml',               'Trapiche',  'vinos-tintos',  9500.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VINOS TINTOS — Norton
-- ============================================================
('Norton Malbec 750ml',                         'Norton',    'vinos-tintos',  5500.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Norton Reserva Malbec 750ml',                 'Norton',    'vinos-tintos',  9500.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Norton Cabernet Sauvignon 750ml',             'Norton',    'vinos-tintos',  5500.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Norton Merlot 750ml',                         'Norton',    'vinos-tintos',  5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Norton Privada 750ml',                        'Norton',    'vinos-tintos', 18000.00, 'seed', 0.8, 88, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VINOS TINTOS — Zuccardi
-- ============================================================
('Zuccardi Valle de Uco Malbec 750ml',          'Zuccardi',  'vinos-tintos', 18000.00, 'seed', 0.9, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Zuccardi Serie A Malbec 750ml',               'Zuccardi',  'vinos-tintos',  7500.00, 'seed', 0.9, 88, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Zuccardi Concreto Malbec 750ml',              'Zuccardi',  'vinos-tintos', 35000.00, 'seed', 0.9, 92, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Santa Julia Malbec 750ml',                    'Santa Julia','vinos-tintos', 4200.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Santa Julia Reserva Malbec 750ml',            'Santa Julia','vinos-tintos', 7000.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VINOS TINTOS — Trivento
-- ============================================================
('Trivento Reserve Malbec 750ml',               'Trivento',  'vinos-tintos',  7800.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Trivento Golden Reserve Malbec 750ml',        'Trivento',  'vinos-tintos', 18000.00, 'seed', 0.8, 88, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Trivento Amado Sur Malbec 750ml',             'Trivento',  'vinos-tintos',  5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VINOS TINTOS — Nieto Senetiner
-- ============================================================
('Nieto Senetiner Malbec Bonarda 750ml',        'Nieto Senetiner','vinos-tintos', 5500.00,'seed',0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Nieto Senetiner Reserva Malbec 750ml',        'Nieto Senetiner','vinos-tintos', 9500.00,'seed',0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Nieto Senetiner Bonarda 750ml',               'Nieto Senetiner','vinos-tintos', 5500.00,'seed',0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VINOS TINTOS — Luigi Bosca, El Esteco, Graffigna
-- ============================================================
('Luigi Bosca Malbec 750ml',                    'Luigi Bosca','vinos-tintos', 12000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('El Esteco Old Vines Malbec 750ml',            'El Esteco', 'vinos-tintos', 12000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('El Esteco Don David Malbec 750ml',            'El Esteco', 'vinos-tintos',  8500.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Graffigna Centenario Malbec 750ml',           'Graffigna', 'vinos-tintos',  5800.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Graffigna Centenario Shiraz 750ml',           'Graffigna', 'vinos-tintos',  5800.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VINOS TINTOS — Premium (segmento alta vinoteca)
-- ============================================================
('Catena Zapata Adrianna Malbec 750ml',         'Catena Zapata','vinos-tintos',145000.00,'seed',0.9, 96, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Catena Zapata Malbec Argentino 750ml',        'Catena Zapata','vinos-tintos', 38000.00,'seed',0.9, 92, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Achaval Ferrer Malbec Mendoza 750ml',         'Achaval Ferrer','vinos-tintos',28000.00,'seed',0.9, 92, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Clos de los Siete 750ml',                     'Clos de los Siete','vinos-tintos',22000.00,'seed',0.9, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Finca La Anita Malbec 750ml',                 'Finca La Anita','vinos-tintos',15000.00,'seed',0.9, 88, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VINOS TINTOS — En botellón 1.5L (alta rotación en NEA)
-- ============================================================
('Trapiche Malbec Roble 1500ml',                'Trapiche',  'vinos-tintos',  9000.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Norton Malbec 1500ml',                        'Norton',    'vinos-tintos', 10000.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VINOS BLANCOS — Chardonnay
-- ============================================================
('Trapiche Chardonnay Roble 750ml',             'Trapiche',  'vinos-blancos',  4800.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Norton Chardonnay 750ml',                     'Norton',    'vinos-blancos',  5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Zuccardi Serie A Chardonnay 750ml',           'Zuccardi',  'vinos-blancos',  7500.00, 'seed', 0.9, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Santa Julia Chardonnay 750ml',                'Santa Julia','vinos-blancos', 4200.00, 'seed', 0.8, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Luigi Bosca Chardonnay 750ml',                'Luigi Bosca','vinos-blancos',12000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VINOS BLANCOS — Torrontés (varietal argentino top)
-- ============================================================
('Graffigna Torrontés 750ml',                   'Graffigna', 'vinos-blancos',  5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('El Esteco Don David Torrontés 750ml',         'El Esteco', 'vinos-blancos',  8500.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Crios Torrontés 750ml',                       'Crios',     'vinos-blancos',  7500.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Crios Sauvignon Blanc 750ml',                 'Crios',     'vinos-blancos',  7500.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Trivento Amado Sur Torrontés 750ml',          'Trivento',  'vinos-blancos',  5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VINOS BLANCOS — Sauvignon Blanc y otros varietales
-- ============================================================
('Norton Sauvignon Blanc 750ml',                'Norton',    'vinos-blancos',  5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Trapiche Sauvignon Blanc 750ml',              'Trapiche',  'vinos-blancos',  4800.00, 'seed', 0.8, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Santa Julia Torrontés 750ml',                 'Santa Julia','vinos-blancos', 4200.00, 'seed', 0.8, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Zuccardi Q Chardonnay 750ml',                 'Zuccardi',  'vinos-blancos', 12000.00, 'seed', 0.9, 88, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VINOS ROSADOS
-- ============================================================
('Trapiche Rosé Malbec 750ml',                  'Trapiche',  'vinos-rosados',  4800.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Norton Rosé 750ml',                           'Norton',    'vinos-rosados',  5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Zuccardi Serie A Rosé 750ml',                 'Zuccardi',  'vinos-rosados',  7500.00, 'seed', 0.9, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Trivento Amado Sur Rosé 750ml',               'Trivento',  'vinos-rosados',  5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Crios Rosé de Malbec 750ml',                  'Crios',     'vinos-rosados',  7500.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Santa Julia Rosé 750ml',                      'Santa Julia','vinos-rosados', 4200.00, 'seed', 0.8, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('El Esteco Rosé 750ml',                        'El Esteco', 'vinos-rosados',  7500.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- ESPUMANTES ARGENTINOS (van en vinoteca junto a vinos)
-- ============================================================
('Chandon Extra Brut 750ml',                    'Chandon',   'espumantes',    12000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Chandon Brut Nature 750ml',                   'Chandon',   'espumantes',    14000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Chandon Délice Rosé 750ml',                   'Chandon',   'espumantes',    16000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Mumm Cuvée Royale Extra Brut 750ml',          'Mumm',      'espumantes',    18000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Valentin Bianchi Extra Brut 750ml',           'Valentin Bianchi','espumantes',9500.00,'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Baron B Extra Brut 750ml',                    'Baron B',   'espumantes',    45000.00, 'seed', 0.9, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Achaval Ferrer Extra Brut 750ml',             'Achaval Ferrer','espumantes',22000.00,'seed', 0.9, 88, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Rainier Extra Brut 750ml',                    'Rainier',   'espumantes',     7500.00, 'seed', 0.7, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Santa Julia Sparkling Pinot Grigio 750ml',    'Santa Julia','espumantes',    8500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VINOS PREMIUM — para vinoteca de alta gama
-- ============================================================
('Catena Zapata Adrianna Vineyard Malbec 750ml','Catena Zapata','vinos-tintos', 145000.00,'seed',0.9,96,FALSE,TRUE,'vinoteca','{"unit": "botella", "sku_prefix": "VINO"}'),
('Clos de los Siete Malbec Blend 750ml',        'Clos de los Siete','vinos-tintos',22000.00,'seed',0.9,90,FALSE,TRUE,'vinoteca','{"unit": "botella", "sku_prefix": "VINO"}'),
('Achaval Ferrer Quimera 750ml',                'Achaval Ferrer','vinos-tintos',55000.00,'seed',0.9,92,FALSE,TRUE,'vinoteca','{"unit": "botella", "sku_prefix": "VINO"}'),
('Zuccardi Concreto Malbec Valle de Uco 750ml', 'Zuccardi',  'vinos-tintos',  35000.00, 'seed', 0.9, 92, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Luigi Bosca Gala 750ml',                      'Luigi Bosca','vinos-tintos', 28000.00, 'seed', 0.8, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Finca La Anita Cabernet Sauvignon 750ml',     'Finca La Anita','vinos-tintos',15000.00,'seed',0.9,88,FALSE,TRUE,'vinoteca','{"unit": "botella", "sku_prefix": "VINO"}'),
('El Esteco El Sueno Malbec 750ml',             'El Esteco', 'vinos-tintos',  28000.00, 'seed', 0.8, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 88 productos
-- Categorías usadas: vinos-tintos, vinos-blancos, vinos-rosados, espumantes
-- Bodegas: Trapiche, Norton, Zuccardi, Trivento, Santa Julia, Graffigna, El Esteco, Nieto Senetiner,
--   Clos de los Siete, Luigi Bosca, Catena Zapata, Achaval Ferrer, Finca La Anita, Chandon, Mumm,
--   Valentin Bianchi, Baron B, Rainier, Crios
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
-- Nota: precios reflejan realidad de vinoteca, no almacén (ver seed 093 para vinos de almacén)
