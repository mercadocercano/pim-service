-- Seed 071: Global products — Bebidas alcohólicas (vinos, cervezas, aperitivos)
-- CICLO: cycle-004-brands-catalog-expansion
-- FUENTE: precios de góndola observados en Posadas/NEA + distribuidores mayoristas (2026-04)
-- IDEMPOTENTE: ON CONFLICT DO NOTHING (por nombre+marca, sin duplicar)
-- REQUIERE: 062 (marcas cervezas/aperitivos), 063 (marcas vinos)
--
-- PRECIOS DE REFERENCIA (zona Posadas, abril 2026):
--   Vino económico 750ml:    AR$ 3.000 - 6.000
--   Vino gama media 750ml:   AR$ 6.000 - 15.000
--   Vino premium 750ml:      AR$ 15.000 - 30.000
--   Cerveza lata 473ml:      AR$ 1.500 - 2.500
--   Cerveza botella 330ml:   AR$ 1.800 - 2.500
--   Cerveza botella 1L:      AR$ 3.000 - 4.500
--   Fernet 750ml:            AR$ 18.000 - 25.000
--   Fernet 1L:               AR$ 24.000 - 30.000
--   Campari 750ml:           AR$ 16.000 - 20.000
--   Gancia vermouth 750ml:   AR$ 7.000 - 10.000
--   Gancia espumante 750ml:  AR$ 7.000 - 9.000

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- VINOS TINTOS
-- ============================================================
('Trapiche Malbec 750ml',              'Trapiche',       'vinos-tintos', 7500.00,  'seed', 0.7, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Trapiche Malbec Reserva 750ml',      'Trapiche',       'vinos-tintos', 12000.00, 'seed', 0.7, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Trapiche Malbec caja x6 750ml',      'Trapiche',       'vinos-tintos', 42000.00, 'seed', 0.7, 78, FALSE, TRUE, 'vinoteca', '{"unit": "caja x6", "sku_prefix": "VINOT"}'),
('Norton Malbec 750ml',                'Norton',         'vinos-tintos', 9000.00,  'seed', 0.7, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Finca Las Moras Malbec 750ml',       'Finca Las Moras','vinos-tintos', 7000.00,  'seed', 0.7, 78, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Catena Malbec 750ml',                'Catena',         'vinos-tintos', 18000.00, 'seed', 0.7, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Alamos Malbec 750ml',                'Alamos',         'vinos-tintos', 8500.00,  'seed', 0.7, 78, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Luigi Bosca Malbec 750ml',           'Luigi Bosca',    'vinos-tintos', 14000.00, 'seed', 0.7, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Santa Julia Malbec 750ml',           'Santa Julia',    'vinos-tintos', 6500.00,  'seed', 0.6, 75, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Gato Negro Cabernet Sauvignon 750ml','Gato Negro',     'vinos-tintos', 5500.00,  'seed', 0.6, 75, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Callia Malbec 750ml',                'Callia',         'vinos-tintos', 4500.00,  'seed', 0.6, 72, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('La Linda Malbec 750ml',              'La Linda',       'vinos-tintos', 4200.00,  'seed', 0.6, 72, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Don Valentín Lacrado Tinto 750ml',   'Don Valentín Lacrado','vinos-tintos',3500.00,'seed',0.5, 68, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Vasco Viejo Tinto 750ml',            'Vasco Viejo',    'vinos-tintos', 3200.00,  'seed', 0.5, 68, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Quara Malbec 750ml',                 'Quara',          'vinos-tintos', 5000.00,  'seed', 0.6, 73, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),

-- ============================================================
-- VINOS BLANCOS
-- ============================================================
('Trapiche Chardonnay 750ml',          'Trapiche',       'vinos-blancos', 7000.00, 'seed', 0.7, 78, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Finca Las Moras Torrontés 750ml',    'Finca Las Moras','vinos-blancos', 6500.00, 'seed', 0.7, 77, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Santa Julia Sauvignon Blanc 750ml',  'Santa Julia',    'vinos-blancos', 6000.00, 'seed', 0.6, 75, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Gato Negro Chardonnay 750ml',        'Gato Negro',     'vinos-blancos', 5200.00, 'seed', 0.6, 74, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Callia Blanco Suave 750ml',          'Callia',         'vinos-blancos', 4200.00, 'seed', 0.6, 70, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Alamos Torrontés 750ml',             'Alamos',         'vinos-blancos', 8000.00, 'seed', 0.7, 77, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),

-- ============================================================
-- VINOS ROSADOS
-- ============================================================
('Gato Negro Rosé 750ml',              'Gato Negro',     'vinos-rosados', 5500.00, 'seed', 0.6, 74, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Santa Julia Rosé 750ml',             'Santa Julia',    'vinos-rosados', 6000.00, 'seed', 0.6, 74, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Finca Las Moras Rosé 750ml',         'Finca Las Moras','vinos-rosados', 6000.00, 'seed', 0.6, 74, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),

-- ============================================================
-- ESPUMANTES
-- ============================================================
('Trapiche Extra Brut 750ml',          'Trapiche',       'espumantes', 10000.00,  'seed', 0.7, 78, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Luigi Bosca Brut Nature 750ml',      'Luigi Bosca',    'espumantes', 14000.00,  'seed', 0.7, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Bianchi Extra Brut 750ml',           'Bianchi',        'espumantes',  9000.00,  'seed', 0.6, 76, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Gancia Clasico Espumante 750ml',     'Gancia',         'espumantes',  7500.00,  'seed', 0.6, 74, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),

-- ============================================================
-- VERMOUTHS Y APERITIVOS
-- ============================================================
('Fernet Branca 750ml',                'Fernet Branca',  'vermouths-aperitivos', 22000.00, 'seed', 0.8, 85, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Fernet Branca 1L',                   'Fernet Branca',  'vermouths-aperitivos', 28000.00, 'seed', 0.8, 85, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Campari 750ml',                      'Campari',        'vermouths-aperitivos', 18000.00, 'seed', 0.7, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Cinzano Bianco 750ml',               'Cinzano',        'vermouths-aperitivos',  9000.00, 'seed', 0.7, 78, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Cinzano Rosso 750ml',                'Cinzano',        'vermouths-aperitivos',  9000.00, 'seed', 0.7, 78, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Gancia Vermouth 750ml',              'Gancia',         'vermouths-aperitivos',  8000.00, 'seed', 0.7, 77, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Hesperidina 500ml',                  'Hesperidina',    'vermouths-aperitivos',  6500.00, 'seed', 0.6, 70, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Amargo Obrero 500ml',                'Amargo Obrero',  'vermouths-aperitivos',  4500.00, 'seed', 0.5, 65, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),

-- ============================================================
-- CERVEZAS (segmento premium en vinoteca)
-- ============================================================
('Heineken 330ml',                     'Heineken',       'cervezas-artesanales', 2200.00, 'seed', 0.7, 78, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Heineken 650ml',                     'Heineken',       'cervezas-artesanales', 3500.00, 'seed', 0.7, 78, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Patagonia Amber Lager 730ml',        'Patagonia',      'cervezas-artesanales', 4500.00, 'seed', 0.7, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Patagonia Bohemian Pilsener 730ml',  'Patagonia',      'cervezas-artesanales', 4200.00, 'seed', 0.7, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Patagonia Weisse 730ml',             'Patagonia',      'cervezas-artesanales', 4500.00, 'seed', 0.7, 78, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Budweiser 473ml lata',               'Budweiser',      'cervezas-artesanales', 2000.00, 'seed', 0.6, 72, FALSE, TRUE, 'vinoteca', '{"unit": "lata",    "sku_prefix": "VINOT"}'),
-- Cervezas corrientes también vendidas en vinotecas de NEA
('Norte 1L',                           'Norte',          'cervezas-artesanales', 4000.00, 'seed', 0.7, 75, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINOT"}'),
('Brahma 473ml lata',                  'Brahma',         'cervezas-artesanales', 1800.00, 'seed', 0.7, 74, FALSE, TRUE, 'vinoteca', '{"unit": "lata",    "sku_prefix": "VINOT"}')

ON CONFLICT DO NOTHING;
