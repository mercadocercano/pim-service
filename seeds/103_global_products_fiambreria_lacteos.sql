-- Seed 103: Productos globales — Fiambrería: Quesos, Fiambres y Lácteos (~45 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento fiambrerías NEA (Posadas) + distribuidores lácteos 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: quesos-frescos, quesos-duros, fiambres-cocidos, fiambres-curados, lacteos-frescos

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- QUESOS FRESCOS — precio x kg o x unidad
-- ============================================================
('Queso cremoso x kg La Serenísima',             'La Serenísima', 'quesos-frescos', 12500.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Queso cuartirolo x kg SanCor',                 'SanCor',        'quesos-frescos', 12000.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Queso mozzarella x kg La Serenísima',          'La Serenísima', 'quesos-frescos', 13500.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Queso de máquina x kg genérico',               NULL,            'quesos-frescos', 11000.00, 'seed', 0.7, 60, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Queso brie Bel Paese x unidad 100g',           'Bel Paese',     'quesos-frescos',  5800.00, 'seed', 0.8, 80, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),
('Queso crema vasito 200g La Serenísima',        'La Serenísima', 'quesos-frescos',  4200.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),
('Queso crema Philadelphia 165g Kraft',          'Kraft',         'quesos-frescos',  6500.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),

-- ============================================================
-- QUESOS DUROS Y SEMIBLANDOS
-- ============================================================
('Queso reggianito x kg Tregar',                 'Tregar',        'quesos-duros', 18000.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Queso provolone x kg La Serenísima',           'La Serenísima', 'quesos-duros', 16000.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Queso gruyère x kg genérico',                  NULL,            'quesos-duros', 18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Queso pategras x kg SanCor',                   'SanCor',        'quesos-duros', 15000.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Queso sardo x kg genérico',                    NULL,            'quesos-duros', 14500.00, 'seed', 0.7, 60, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Queso tybo x kg La Serenísima',                'La Serenísima', 'quesos-duros', 14000.00, 'seed', 0.9, 82, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Queso rallado 100g La Serenísima',             'La Serenísima', 'quesos-duros',  3500.00, 'seed', 0.9, 82, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),

-- ============================================================
-- FIAMBRES COCIDOS
-- ============================================================
('Jamón cocido x kg Paladini',                   'Paladini',      'fiambres-cocidos', 14000.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Paleta cocida x kg Paladini',                  'Paladini',      'fiambres-cocidos', 12500.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Mortadela x kg Cagnoli',                       'Cagnoli',       'fiambres-cocidos', 10000.00, 'seed', 0.9, 82, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Mortadela con aceitunas x kg Cagnoli',         'Cagnoli',       'fiambres-cocidos', 10500.00, 'seed', 0.9, 82, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Queso de cerdo x kg genérico',                 NULL,            'fiambres-cocidos',  9000.00, 'seed', 0.7, 60, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Lomito cocido x kg Quickfood',                 'Quickfood',     'fiambres-cocidos', 18000.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Pollo en gelatina x kg genérico',              NULL,            'fiambres-cocidos',  8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),

-- ============================================================
-- FIAMBRES CURADOS
-- ============================================================
('Salame x kg Cagnoli',                          'Cagnoli',       'fiambres-curados', 20000.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Salame tipo Milano x kg Paladini',             'Paladini',      'fiambres-curados', 22000.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Bresaola x kg genérico',                       NULL,            'fiambres-curados', 25000.00, 'seed', 0.7, 60, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Panceta ahumada x kg Cagnoli',                 'Cagnoli',       'fiambres-curados', 14000.00, 'seed', 0.9, 82, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Prosciutto x kg genérico',                     NULL,            'fiambres-curados', 32000.00, 'seed', 0.7, 60, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),
('Chorizo colorado x kg La Salamandra',          'La Salamandra', 'fiambres-curados', 15000.00, 'seed', 0.8, 80, FALSE, TRUE, 'fiambreria', '{"unit": "kg",    "sku_prefix": "FIAM"}'),

-- ============================================================
-- LÁCTEOS FRESCOS
-- ============================================================
('Manteca sin sal 200g La Serenísima',           'La Serenísima', 'lacteos-frescos',  3200.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),
('Manteca con sal 200g La Serenísima',           'La Serenísima', 'lacteos-frescos',  3200.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),
('Crema de leche 200ml SanCor',                  'SanCor',        'lacteos-frescos',  3800.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),
('Crema de leche 200ml La Serenísima',           'La Serenísima', 'lacteos-frescos',  3800.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),
('Ricota 300g La Serenísima',                    'La Serenísima', 'lacteos-frescos',  3500.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),
('Ricota 300g SanCor',                           'SanCor',        'lacteos-frescos',  3500.00, 'seed', 0.9, 85, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),
('Yogur entero natural 190g vasito La Serenísima','La Serenísima','lacteos-frescos',  1800.00, 'seed', 0.9, 82, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),
('Yogur entero frutilla 190g vasito SanCor',     'SanCor',        'lacteos-frescos',  1800.00, 'seed', 0.9, 82, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),
('Yogur bebible 200ml La Serenísima',            'La Serenísima', 'lacteos-frescos',  2200.00, 'seed', 0.9, 82, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}'),
('Leche cultivada 950ml La Serenísima',          'La Serenísima', 'lacteos-frescos',  4800.00, 'seed', 0.9, 82, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAM"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 44 productos
-- Categorías: quesos-frescos, quesos-duros, fiambres-cocidos, fiambres-curados, lacteos-frescos
-- Marcas: La Serenísima, SanCor, Tregar, Kraft, Paladini, Cagnoli, Quickfood, La Salamandra, Bel Paese
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
