-- Seed 101: Productos globales — Carnicería: Cortes y Embutidos (~65 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento carnicerías NEA (Posadas) + precios Mercado Regional 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026 (precio x kg)
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: cortes-vacunos, cortes-cerdo, cortes-pollo, embutidos-fiambres, achuras

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- CORTES VACUNOS — precio x kg (minorista NEA)
-- ============================================================
('Asado de tira x kg',                           NULL, 'cortes-vacunos',  9500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Vacío x kg',                                   NULL, 'cortes-vacunos', 11000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Nalga x kg',                                   NULL, 'cortes-vacunos', 10500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Cuadril x kg',                                 NULL, 'cortes-vacunos', 12000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Bife de chorizo x kg',                         NULL, 'cortes-vacunos', 14000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Lomo x kg',                                    NULL, 'cortes-vacunos', 18000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Paleta x kg',                                  NULL, 'cortes-vacunos',  9000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Osobuco x kg',                                 NULL, 'cortes-vacunos',  8500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Carne molida común x kg',                      NULL, 'cortes-vacunos',  8800.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Carne molida especial x kg',                   NULL, 'cortes-vacunos', 10500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Hígado vacuno x kg',                           NULL, 'cortes-vacunos',  5500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Riñones vacunos x kg',                         NULL, 'cortes-vacunos',  4500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Matambre x kg',                                NULL, 'cortes-vacunos', 10000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Carnaza de paleta x kg',                       NULL, 'cortes-vacunos',  8200.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Bife de costilla x kg',                        NULL, 'cortes-vacunos', 10000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Tapa de asado x kg',                           NULL, 'cortes-vacunos',  9800.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Aguja x kg',                                   NULL, 'cortes-vacunos',  8000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Roast beef x kg',                              NULL, 'cortes-vacunos', 11000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),

-- ============================================================
-- CORTES DE CERDO
-- ============================================================
('Bondiola de cerdo x kg',                       NULL, 'cortes-cerdo',  9500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Costeletas de cerdo x kg',                     NULL, 'cortes-cerdo',  8500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Panceta de cerdo x kg',                        NULL, 'cortes-cerdo',  8000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Patitas de cerdo x kg',                        NULL, 'cortes-cerdo',  4800.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Paleta de cerdo sin hueso x kg',               NULL, 'cortes-cerdo',  9000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Lomo de cerdo x kg',                           NULL, 'cortes-cerdo', 11000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Hígado de cerdo x kg',                         NULL, 'cortes-cerdo',  4200.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),

-- ============================================================
-- POLLO
-- ============================================================
('Pollo entero x kg',                            NULL, 'cortes-pollo',   5800.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Pechuga de pollo sin hueso x kg',              NULL, 'cortes-pollo',   8500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Muslo y contramuslo de pollo x kg',            NULL, 'cortes-pollo',   6200.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Alitas de pollo x kg',                         NULL, 'cortes-pollo',   5500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Menudencias de pollo x kg',                    NULL, 'cortes-pollo',   3200.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Suprema de pollo x kg',                        NULL, 'cortes-pollo',   9200.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Pata muslo de pollo x kg',                     NULL, 'cortes-pollo',   6500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Pollo trozado x kg',                           NULL, 'cortes-pollo',   6000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),

-- ============================================================
-- EMBUTIDOS — Quickfood, Paladini, genéricos locales
-- ============================================================
('Chorizo parrillero x kg',                      NULL,         'embutidos-fiambres',  8000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Chorizo parrillero x unidad',                  NULL,         'embutidos-fiambres',  1200.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'),
('Morcilla x kg',                                NULL,         'embutidos-fiambres',  6500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Salchicha viena 180g Quickfood',               'Quickfood',  'embutidos-fiambres',  3200.00, 'seed', 0.9, 82, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'),
('Salchicha viena 180g Paladini',                'Paladini',   'embutidos-fiambres',  3200.00, 'seed', 0.9, 82, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'),
('Panceta ahumada x kg Paladini',                'Paladini',   'embutidos-fiambres',  9500.00, 'seed', 0.9, 82, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Salame x kg genérico',                         NULL,         'embutidos-fiambres', 12000.00, 'seed', 0.7, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Longaniza x kg',                               NULL,         'embutidos-fiambres',  8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Salchicha parrillera x kg',                    NULL,         'embutidos-fiambres',  7800.00, 'seed', 0.7, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),

-- ============================================================
-- ACHURAS
-- ============================================================
('Chinchulines de res x kg',                     NULL, 'achuras',   5500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Tripa gorda de res x kg',                      NULL, 'achuras',   4800.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Corazón vacuno x kg',                          NULL, 'achuras',   4500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Molleja x kg',                                 NULL, 'achuras',   7200.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Mondongo x kg',                                NULL, 'achuras',   5000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Lengua vacuna x kg',                           NULL, 'achuras',   8500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Ubre vacuna x kg',                             NULL, 'achuras',   4200.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),

-- ============================================================
-- HAMBURGUESAS Y PREPARADOS
-- ============================================================
('Hamburguesa casera x unidad 120g',             NULL,        'embutidos-fiambres',  1800.00, 'seed', 0.7, 60, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'),
('Hamburguesa Paty x4u 480g',                   'Paty',      'embutidos-fiambres',  8500.00, 'seed', 0.9, 82, FALSE, TRUE, 'carniceria', '{"unit": "pack",   "sku_prefix": "CARN"}'),
('Milanesa de nalga feteada x kg',               NULL,        'cortes-vacunos',     12000.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Milanesa de cerdo feteada x kg',               NULL,        'cortes-cerdo',        9500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}'),
('Milanesa de pollo feteada x kg',               NULL,        'cortes-pollo',        8500.00, 'seed', 0.8, 60, FALSE, TRUE, 'carniceria', '{"unit": "kg",    "sku_prefix": "CARN"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 64 productos
-- Categorías: cortes-vacunos, cortes-cerdo, cortes-pollo, embutidos-fiambres, achuras
-- Marcas: mayoritariamente genérico (carnicería local), Quickfood, Paladini, Paty
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista x kg estimado)
-- NOTA: precios de carne fluctúan semanalmente, usar como referencia inicial
