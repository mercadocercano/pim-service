-- Seed 121: Productos globales — Carnicería extendida (~40 productos)
-- CICLO: cycle-009-catalog-volume-expansion (T013)
-- FECHA: 2026-04-25
-- FUENTE: relevamiento carnicerías NEA + distribuidores frigoríficos Posadas 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- OBJETIVO: llevar carnicería a >= 120 verificados (hoy: 174 — ya cumple; ampliamos cobertura)
-- NOTA: Priorizamos cortes con marca, variantes por peso y cortes de cerdo/pollo premium

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- CORTES VACUNOS — variantes con marca y peso específico
-- ============================================================
('Osobuco x kg Swift',                         'Swift',            'cortes-vacunos',  4200.00, 'seed', 0.85, 80, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Osobuco x kg Cabaña Argentina',              'Cabaña Argentina', 'cortes-vacunos',  4500.00, 'seed', 0.85, 80, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Cuadrada x kg',                              NULL,               'cortes-vacunos',  3900.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Tapa de cuadril x kg',                       NULL,               'cortes-vacunos',  5500.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Bola de lomo x kg',                          NULL,               'cortes-vacunos',  4800.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Peceto x kg',                                NULL,               'cortes-vacunos',  5200.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Bife de chorizo Swift Angus x kg',           'Swift',            'cortes-vacunos',  8500.00, 'seed', 0.85, 82, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Entraña fina x kg',                          NULL,               'cortes-vacunos',  7200.00, 'seed', 0.80, 78, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Costillar de res entero x kg',               NULL,               'cortes-vacunos',  3600.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Colita de cuadril x kg',                     NULL,               'cortes-vacunos',  5800.00, 'seed', 0.80, 78, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Colita de cuadril Cabaña Argentina x kg',    'Cabaña Argentina', 'cortes-vacunos',  6200.00, 'seed', 0.85, 82, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- CORTES DE CERDO — ampliación con presentaciones
-- ============================================================
('Chuleta de cerdo x kg',                      NULL,               'cortes-cerdo',    3200.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Panceta de cerdo fresca x kg',               NULL,               'cortes-cerdo',    3500.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Bondiola de cerdo entera x kg Campo Austral','Campo Austral',    'cortes-cerdo',    4200.00, 'seed', 0.85, 82, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Paleta de cerdo con hueso x kg',             NULL,               'cortes-cerdo',    2900.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Lomo de cerdo con hueso x kg',               NULL,               'cortes-cerdo',    4000.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Pechito de cerdo x kg',                      NULL,               'cortes-cerdo',    2800.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- EMBUTIDOS FRESCOS — chorizo parrillero y morcilla artesanal
-- ============================================================
('Chorizo parrillero fresco artesanal x kg',   NULL,               'embutidos',       3800.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Chorizo parrillero fresco Paladini x kg',    'Paladini',         'embutidos',       4200.00, 'seed', 0.85, 82, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Morcilla criolla artesanal x kg',            NULL,               'embutidos',       3500.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Morcilla dulce x kg',                        NULL,               'embutidos',       3600.00, 'seed', 0.78, 72, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Salchicha parrillera Campo Austral x kg',    'Campo Austral',    'embutidos',       4000.00, 'seed', 0.85, 82, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- POLLO — cortes con marca premium y variantes de peso
-- ============================================================
('Cuarto trasero de pollo x kg',               NULL,               'cortes-pollo',    2400.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Cuarto trasero Cresta Roja x kg',            'Cresta Roja',      'cortes-pollo',    2600.00, 'seed', 0.85, 82, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Pechuga de pollo con hueso x kg',            NULL,               'cortes-pollo',    2800.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Alas de pollo Granja Tres Arroyos x kg',     'Granja Tres Arroyos', 'cortes-pollo', 2200.00, 'seed', 0.85, 82, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Pollo entero Granja del Sol x kg',           'Granja del Sol',   'cortes-pollo',    2500.00, 'seed', 0.85, 82, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Muslo de pollo sin hueso x kg',              NULL,               'cortes-pollo',    3200.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- ACHURAS — variantes adicionales
-- ============================================================
('Seso de vaca x kg',                          NULL,               'achuras',         2500.00, 'seed', 0.78, 70, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Bofe de vaca x kg',                          NULL,               'achuras',         1800.00, 'seed', 0.75, 65, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Panza de vaca (librillo) x kg',              NULL,               'achuras',         2200.00, 'seed', 0.78, 70, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Tripa fina x kg',                            NULL,               'achuras',         2000.00, 'seed', 0.78, 70, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- MILANESAS — variantes con marca
-- ============================================================
('Milanesas de bola de lomo x kg',             NULL,               'milanesas',       6500.00, 'seed', 0.80, 78, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Milanesas de cuadril Swift x4 500g',         'Swift',            'milanesas',      3800.00, 'seed', 0.85, 82, TRUE, TRUE, 'carniceria', '{"unit": "bandeja", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Milanesas de nalga Swift x4 500g',           'Swift',            'milanesas',      3600.00, 'seed', 0.85, 82, TRUE, TRUE, 'carniceria', '{"unit": "bandeja", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Milanesas de pollo artesanales x kg',        NULL,               'milanesas',       4200.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- HAMBURGUESAS — variantes artesanales
-- ============================================================
('Hamburguesa artesanal vacuna x kg',          NULL,               'hamburguesas',    5500.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Hamburguesa artesanal vacuna x4 480g',       NULL,               'hamburguesas',    2800.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "bandeja", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Hamburguesa Paty triple x2 300g',            'Paty',             'hamburguesas',    2200.00, 'seed', 0.85, 82, TRUE, TRUE, 'carniceria', '{"unit": "pack", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- CARNE MOLIDA — variantes adicionales
-- ============================================================
('Carne picada mixta (vaca/cerdo) x kg',       NULL,               'carne-molida',    3400.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Carne picada de pollo x kg',                 NULL,               'carne-molida',    2600.00, 'seed', 0.80, 75, TRUE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN", "research_date": "2026-04-25", "zone": "Posadas NEA"}')

ON CONFLICT (name, business_type) DO NOTHING;
