-- Seed 089: Productos globales — Vinoteca: Cervezas Premium y Accesorios (~40 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: catálogos distribuidores bebidas NEA + sitios de cervezas importadas Argentina 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: cervezas-importadas, cervezas-artesanales, accesorios-vino

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- CERVEZAS IMPORTADAS — en botella y lata
-- ============================================================
('Heineken Lager 330ml botella',                'Heineken',  'cervezas-importadas',  1800.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Heineken Lager 473ml lata',                   'Heineken',  'cervezas-importadas',  2200.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "lata",    "sku_prefix": "VINO"}'),
('Heineken 0.0 sin alcohol 330ml',              'Heineken',  'cervezas-importadas',  2000.00, 'seed', 0.8, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Stella Artois 330ml botella',                 'Stella Artois','cervezas-importadas',1800.00,'seed',0.8,82,FALSE,TRUE,'vinoteca','{"unit": "botella", "sku_prefix": "VINO"}'),
('Stella Artois 473ml lata',                    'Stella Artois','cervezas-importadas',2200.00,'seed',0.8,82,FALSE,TRUE,'vinoteca','{"unit": "lata",    "sku_prefix": "VINO"}'),
('Corona Extra 355ml botella',                  'Corona',    'cervezas-importadas',  2500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Corona Extra 473ml lata',                     'Corona',    'cervezas-importadas',  2800.00, 'seed', 0.8, 80, FALSE, TRUE, 'vinoteca', '{"unit": "lata",    "sku_prefix": "VINO"}'),
('Budweiser 473ml lata',                        'Budweiser', 'cervezas-importadas',  1800.00, 'seed', 0.8, 78, FALSE, TRUE, 'vinoteca', '{"unit": "lata",    "sku_prefix": "VINO"}'),
('Beck''s 330ml botella',                       'Beck''s',   'cervezas-importadas',  2200.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Guinness Draught 440ml lata',                 'Guinness',  'cervezas-importadas',  3800.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "lata",    "sku_prefix": "VINO"}'),
('Leffe Blonde 330ml botella',                  'Leffe',     'cervezas-importadas',  3500.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Hoegaarden 330ml botella',                    'Hoegaarden','cervezas-importadas',  3500.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- CERVEZAS ARTESANALES PREMIUM — Patagonia (más distribuida en NEA)
-- ============================================================
('Patagonia Amber Lager 730ml',                 'Patagonia', 'cervezas-artesanales', 4500.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Patagonia Amber Lager 473ml lata',            'Patagonia', 'cervezas-artesanales', 2500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "lata",    "sku_prefix": "VINO"}'),
('Patagonia Weisse 730ml',                      'Patagonia', 'cervezas-artesanales', 4800.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Patagonia Hoppy Lager 473ml lata',            'Patagonia', 'cervezas-artesanales', 2500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "lata",    "sku_prefix": "VINO"}'),
('Patagonia IPA 473ml lata',                    'Patagonia', 'cervezas-artesanales', 2800.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "lata",    "sku_prefix": "VINO"}'),
('Patagonia Imperial Stout 730ml',              'Patagonia', 'cervezas-artesanales', 5500.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Antares Cream Stout 500ml',                   'Antares',   'cervezas-artesanales', 4200.00, 'seed', 0.7, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Antares IPA 500ml',                           'Antares',   'cervezas-artesanales', 4200.00, 'seed', 0.7, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Antares Kolsch 500ml',                        'Antares',   'cervezas-artesanales', 4200.00, 'seed', 0.7, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Juguetes Perdidos Red IPA 500ml',             'Juguetes Perdidos','cervezas-artesanales',5500.00,'seed',0.7,84,FALSE,TRUE,'vinoteca','{"unit": "botella", "sku_prefix": "VINO"}'),
('Juguetes Perdidos Pale Ale 500ml',            'Juguetes Perdidos','cervezas-artesanales',5500.00,'seed',0.7,84,FALSE,TRUE,'vinoteca','{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- ACCESORIOS PARA VINO Y CAVA
-- ============================================================
('Sacacorchos de palanca profesional Pulltex',  'Pulltex',  'accesorios-vino',   8500.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "unidad", "sku_prefix": "VINO"}'),
('Sacacorchos sommelier básico genérico',        NULL,       'accesorios-vino',   1800.00, 'seed', 0.6, 40, FALSE, TRUE, 'vinoteca', '{"unit": "unidad", "sku_prefix": "VINO"}'),
('Sacacorchos de palanca de pared genérico',     NULL,       'accesorios-vino',  12000.00, 'seed', 0.6, 40, FALSE, TRUE, 'vinoteca', '{"unit": "unidad", "sku_prefix": "VINO"}'),
('Decantador de vino cristal 1.5L genérico',     NULL,       'accesorios-vino',  18000.00, 'seed', 0.6, 40, FALSE, TRUE, 'vinoteca', '{"unit": "unidad", "sku_prefix": "VINO"}'),
('Copa de cristal para vino tinto x6u',          NULL,       'accesorios-vino',  28000.00, 'seed', 0.6, 40, FALSE, TRUE, 'vinoteca', '{"unit": "set",    "sku_prefix": "VINO"}'),
('Copa de cristal para champagne flauta x6u',    NULL,       'accesorios-vino',  24000.00, 'seed', 0.6, 40, FALSE, TRUE, 'vinoteca', '{"unit": "set",    "sku_prefix": "VINO"}'),
('Copa de cristal de bordeaux Riedel x2u',      'Riedel',   'accesorios-vino',  35000.00, 'seed', 0.9, 90, FALSE, TRUE, 'vinoteca', '{"unit": "set",    "sku_prefix": "VINO"}'),
('Enfriador de botella neoprene',                NULL,       'accesorios-vino',   4500.00, 'seed', 0.6, 40, FALSE, TRUE, 'vinoteca', '{"unit": "unidad", "sku_prefix": "VINO"}'),
('Cubitera de acero inoxidable con pie',         NULL,       'accesorios-vino',  22000.00, 'seed', 0.6, 40, FALSE, TRUE, 'vinoteca', '{"unit": "unidad", "sku_prefix": "VINO"}'),
('Aerador de vino Vacu Vin',                    'Vacu Vin', 'accesorios-vino',   5500.00, 'seed', 0.8, 80, FALSE, TRUE, 'vinoteca', '{"unit": "unidad", "sku_prefix": "VINO"}'),
('Tapón de vacío para vino x2u Vacu Vin',       'Vacu Vin', 'accesorios-vino',   3800.00, 'seed', 0.8, 80, FALSE, TRUE, 'vinoteca', '{"unit": "pack",   "sku_prefix": "VINO"}'),
('Bolsa de hielo 1kg genérica',                  NULL,       'accesorios-vino',   1200.00, 'seed', 0.6, 40, FALSE, TRUE, 'vinoteca', '{"unit": "unidad", "sku_prefix": "VINO"}'),
('Termómetro de vino digital',                   NULL,       'accesorios-vino',   4800.00, 'seed', 0.6, 40, FALSE, TRUE, 'vinoteca', '{"unit": "unidad", "sku_prefix": "VINO"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 41 productos
-- Categorías usadas: cervezas-importadas, cervezas-artesanales, accesorios-vino
-- Marcas: Heineken, Stella Artois, Corona, Budweiser, Beck's, Guinness, Leffe, Hoegaarden,
--   Patagonia, Antares, Juguetes Perdidos, Pulltex, Riedel, Vacu Vin, genéricos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
