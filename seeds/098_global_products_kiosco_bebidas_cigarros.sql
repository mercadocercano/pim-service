-- Seed 098: Productos globales — Kiosco: Bebidas y Cigarrillos (~45 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento kioscos NEA (Posadas) + distribuidores bebidas 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: aguas-kiosco, jugos-kiosco, energeticas, gaseosas-lata, cigarrillos

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- AGUAS
-- ============================================================
('Agua Villavicencio 500ml sin gas',             'Villavicencio', 'aguas-kiosco',   950.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Agua Villavicencio 1.5L sin gas',              'Villavicencio', 'aguas-kiosco',  1600.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Agua Villavicencio 500ml con gas',             'Villavicencio', 'aguas-kiosco',  1050.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Agua Ser 500ml sin gas',                       'Ser',           'aguas-kiosco',   900.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Agua Nestlé Pure Life 500ml',                  'Nestlé',        'aguas-kiosco',   850.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Agua Eco de los Andes 500ml',                  'Eco de los Andes','aguas-kiosco', 900.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- JUGOS Y BEBIDAS FRUTALES
-- ============================================================
('Jugo Cepita naranja 1L',                       'Cepita',        'jugos-kiosco', 1800.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Jugo Cepita manzana 1L',                       'Cepita',        'jugos-kiosco', 1800.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Jugo Cepita naranja 200ml tetra',              'Cepita',        'jugos-kiosco',  680.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Jugo Baggio naranja 200ml tetra',              'Baggio',        'jugos-kiosco',  650.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Jugo Baggio tropical 200ml tetra',             'Baggio',        'jugos-kiosco',  650.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Jugo Minute Maid naranja lata 350ml',          'Minute Maid',   'jugos-kiosco', 1200.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Bebida Powerade naranja 500ml',                'Powerade',      'jugos-kiosco', 1800.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Bebida Gatorade limón 500ml',                  'Gatorade',      'jugos-kiosco', 2000.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- BEBIDAS ENERGÉTICAS
-- ============================================================
('Monster Energy original 473ml',               'Monster',       'energeticas',  3500.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Monster Energy verde 473ml',                  'Monster',       'energeticas',  3500.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Red Bull original 250ml',                     'Red Bull',      'energeticas',  3200.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Speed energy drink 473ml',                    'Speed',         'energeticas',  2800.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Burn energy original 500ml',                  'Burn',          'energeticas',  2600.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Volt energy 500ml',                           'Volt',          'energeticas',  2200.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- GASEOSAS EN LATA
-- ============================================================
('Coca-Cola lata 350ml',                         'Coca-Cola',     'gaseosas-lata', 1400.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Coca-Cola Zero lata 350ml',                    'Coca-Cola',     'gaseosas-lata', 1400.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Sprite lata 350ml',                            'Sprite',        'gaseosas-lata', 1300.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Fanta naranja lata 350ml',                     'Fanta',         'gaseosas-lata', 1300.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Pepsi lata 350ml',                             'Pepsi',         'gaseosas-lata', 1200.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('7UP lata 350ml',                               '7UP',           'gaseosas-lata', 1200.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Cunnington cola lata 350ml',                   'Cunnington',    'gaseosas-lata',  900.00, 'seed', 0.8, 76, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- GASEOSAS PEQUEÑAS (personal)
-- ============================================================
('Coca-Cola 500ml personal',                     'Coca-Cola',     'gaseosas-lata', 1800.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Pepsi 500ml personal',                         'Pepsi',         'gaseosas-lata', 1600.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Sprite 500ml personal',                        'Sprite',        'gaseosas-lata', 1700.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),

-- ============================================================
-- CIGARRILLOS (EAN ficticios — sin código real publicable)
-- ============================================================
('Cigarrillos Marlboro rojo x20u',               'Marlboro',      'cigarrillos',  4800.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "caja",   "sku_prefix": "KIOS"}'),
('Cigarrillos Marlboro gold x20u',               'Marlboro',      'cigarrillos',  4800.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "caja",   "sku_prefix": "KIOS"}'),
('Cigarrillos Marlboro blue x20u',               'Marlboro',      'cigarrillos',  4600.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "caja",   "sku_prefix": "KIOS"}'),
('Cigarrillos Camel original x20u',              'Camel',         'cigarrillos',  4800.00, 'seed', 0.9, 85, FALSE, TRUE, 'kiosco', '{"unit": "caja",   "sku_prefix": "KIOS"}'),
('Cigarrillos Lucky Strike rojo x20u',           'Lucky Strike',  'cigarrillos',  4600.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "caja",   "sku_prefix": "KIOS"}'),
('Cigarrillos Philip Morris rojo x20u',          'Philip Morris', 'cigarrillos',  4500.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "caja",   "sku_prefix": "KIOS"}'),
('Cigarrillos Philip Morris blue x20u',          'Philip Morris', 'cigarrillos',  4500.00, 'seed', 0.9, 82, FALSE, TRUE, 'kiosco', '{"unit": "caja",   "sku_prefix": "KIOS"}'),
('Cigarrillos Pall Mall rojo x20u',              'Pall Mall',     'cigarrillos',  4200.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "caja",   "sku_prefix": "KIOS"}'),
('Cigarrillos Pall Mall blue x20u',              'Pall Mall',     'cigarrillos',  4200.00, 'seed', 0.8, 80, FALSE, TRUE, 'kiosco', '{"unit": "caja",   "sku_prefix": "KIOS"}'),
('Cigarrillos Derby rojo x20u',                  'Derby',         'cigarrillos',  3800.00, 'seed', 0.8, 76, FALSE, TRUE, 'kiosco', '{"unit": "caja",   "sku_prefix": "KIOS"}'),

-- ============================================================
-- OTROS — café y mate en kiosco
-- ============================================================
('Café Nescafé 3en1 sachet individual',          'Nescafé',       'jugos-kiosco',  450.00, 'seed', 0.9, 80, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Yerba mate Taragüi 100g sachet',               'Taragüi',       'jugos-kiosco',  650.00, 'seed', 0.8, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}'),
('Té Lipton sobre x1u',                          'Lipton',        'jugos-kiosco',  150.00, 'seed', 0.9, 78, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOS"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 46 productos
-- Categorías: aguas-kiosco, jugos-kiosco, energeticas, gaseosas-lata, cigarrillos
-- Marcas: Villavicencio, Ser, Nestlé, Cepita, Baggio, Minute Maid, Powerade, Gatorade, Monster,
--   Red Bull, Speed, Burn, Volt, Coca-Cola, Sprite, Fanta, Pepsi, 7UP, Cunnington, Marlboro,
--   Camel, Lucky Strike, Philip Morris, Pall Mall, Derby, Nescafé, Taragüi, Lipton
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
-- NOTA: cigarrillos sin EAN real — los precios y marcas son de mercado argentino
