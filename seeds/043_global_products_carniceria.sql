-- Seed 043: Carnicería — 120 productos reales argentinos
-- Generado: 2026-04-18
-- Fuente: global_products (v2.0)
-- ON CONFLICT DO NOTHING: idempotente.

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES
-- ============================================================
-- CORTES VACUNOS (25 productos)
-- ============================================================
('Asado de tira', NULL, 'cortes-vacunos', 8500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Asado de tira angosto', NULL, 'cortes-vacunos', 9200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Vacío', NULL, 'cortes-vacunos', 9800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Bife ancho', NULL, 'cortes-vacunos', 10500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Bife angosto', NULL, 'cortes-vacunos', 11000.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Bife de chorizo', NULL, 'cortes-vacunos', 11500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Lomo', NULL, 'cortes-vacunos', 14500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Cuadril', NULL, 'cortes-vacunos', 9500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Nalga', NULL, 'cortes-vacunos', 8800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Paleta', NULL, 'cortes-vacunos', 7500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Matambre', NULL, 'cortes-vacunos', 10000.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Entraña', NULL, 'cortes-vacunos', 12000.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Tapa de asado', NULL, 'cortes-vacunos', 7800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Falda', NULL, 'cortes-vacunos', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Peceto', NULL, 'cortes-vacunos', 9800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Colita de cuadril', NULL, 'cortes-vacunos', 10200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Roast beef', NULL, 'cortes-vacunos', 7200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Osobuco', NULL, 'cortes-vacunos', 6500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Aguja', NULL, 'cortes-vacunos', 6800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Carnaza', NULL, 'cortes-vacunos', 7000.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Tapa de nalga', NULL, 'cortes-vacunos', 8500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Bola de lomo', NULL, 'cortes-vacunos', 8200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Tortuguita', NULL, 'cortes-vacunos', 7800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Vacío Cabaña Argentina', 'Cabaña Argentina', 'cortes-vacunos', 13500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Bife ancho Swift', 'Swift', 'cortes-vacunos', 14000.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),

-- ============================================================
-- CARNE MOLIDA (6 productos)
-- ============================================================
('Carne picada común', NULL, 'carne-molida', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Carne picada especial', NULL, 'carne-molida', 7200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Carne picada premium', NULL, 'carne-molida', 8800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Carne picada especial Swift 500g', 'Swift', 'carne-molida', 4800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Carne picada premium Swift 500g', 'Swift', 'carne-molida', 5900.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Carne picada común Good Mark 500g', 'Good Mark', 'carne-molida', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),

-- ============================================================
-- MILANESAS (10 productos)
-- ============================================================
('Milanesas de nalga', NULL, 'milanesas', 10500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Milanesas de cuadril', NULL, 'milanesas', 11500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Milanesas de peceto', NULL, 'milanesas', 12000.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Milanesas de bola de lomo', NULL, 'milanesas', 10800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Milanesas de pollo', NULL, 'milanesas', 8500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Milanesas de cerdo', NULL, 'milanesas', 8000.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Milanesas de suprema rebozadas Granja Tres Arroyos x4', 'Granja Tres Arroyos', 'milanesas', 6500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Milanesas de carne rebozadas Swift x4', 'Swift', 'milanesas', 5800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Milanesas de soja Good Mark x4', 'Good Mark', 'milanesas', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Supremas napolitanas Granja del Sol x2', 'Granja del Sol', 'milanesas', 5200.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),

-- ============================================================
-- POLLO (14 productos)
-- ============================================================
('Pollo entero', NULL, 'pollo', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Pata muslo de pollo', NULL, 'pollo', 4800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Suprema de pollo', NULL, 'pollo', 8500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Alitas de pollo', NULL, 'pollo', 4200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Muslo de pollo', NULL, 'pollo', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Pechuga de pollo con hueso', NULL, 'pollo', 6800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Pollo entero Granja Tres Arroyos', 'Granja Tres Arroyos', 'pollo', 6200.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Pata muslo Cresta Roja x2kg', 'Cresta Roja', 'pollo', 9500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Suprema Granja Tres Arroyos x1kg', 'Granja Tres Arroyos', 'pollo', 9200.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Pollo entero Cresta Roja', 'Cresta Roja', 'pollo', 5800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Alitas de pollo Campo Austral x1kg', 'Campo Austral', 'pollo', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Medallones de pollo Granja del Sol x4', 'Granja del Sol', 'pollo', 4800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Nuggets de pollo Granja del Sol x12', 'Granja del Sol', 'pollo', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Patitas de pollo Granja Tres Arroyos x12', 'Granja Tres Arroyos', 'pollo', 5800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),

-- ============================================================
-- CERDO (10 productos)
-- ============================================================
('Bondiola de cerdo', NULL, 'cerdo', 8500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Costilla de cerdo', NULL, 'cerdo', 7200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Carré de cerdo', NULL, 'cerdo', 8800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Matambre de cerdo', NULL, 'cerdo', 9200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Pechito de cerdo', NULL, 'cerdo', 6800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Solomillo de cerdo', NULL, 'cerdo', 10500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Pata de cerdo', NULL, 'cerdo', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Paleta de cerdo', NULL, 'cerdo', 6200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Bondiola de cerdo Campo Austral x1kg', 'Campo Austral', 'cerdo', 9800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Carré de cerdo Cabaña Argentina x1kg', 'Cabaña Argentina', 'cerdo', 10500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),

-- ============================================================
-- ACHURAS (10 productos)
-- ============================================================
('Chinchulines de vaca', NULL, 'achuras', 6500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Riñón de vaca', NULL, 'achuras', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Mollejas', NULL, 'achuras', 12000.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Tripa gorda', NULL, 'achuras', 5800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Hígado de vaca', NULL, 'achuras', 3800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Corazón de vaca', NULL, 'achuras', 4200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Lengua de vaca', NULL, 'achuras', 7500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Ubre', NULL, 'achuras', 4000.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Mondongo', NULL, 'achuras', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Chinchulín trenzado', NULL, 'achuras', 7200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),

-- ============================================================
-- EMBUTIDOS (12 productos)
-- ============================================================
('Chorizo parrillero', NULL, 'embutidos', 6500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Chorizo bombón', NULL, 'embutidos', 7200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Morcilla', NULL, 'embutidos', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Salchicha parrillera', NULL, 'embutidos', 5800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Chorizo colorado', NULL, 'embutidos', 6800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Chorizo parrillero Paladini x6', 'Paladini', 'embutidos', 5800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Morcilla vasca Paladini x3', 'Paladini', 'embutidos', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Salchicha parrillera Swift x6', 'Swift', 'embutidos', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Chorizo ahumado Paladini x3', 'Paladini', 'embutidos', 4800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Salchicha tipo viena Paladini x6', 'Paladini', 'embutidos', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Chorizo parrillero Campo Austral x4', 'Campo Austral', 'embutidos', 5200.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Morcilla criolla Swift x3', 'Swift', 'embutidos', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),

-- ============================================================
-- CHACINADOS (10 productos)
-- ============================================================
('Salame quintero Paladini 250g', 'Paladini', 'chacinados', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Salame milán Paladini 150g', 'Paladini', 'chacinados', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Longaniza Paladini 250g', 'Paladini', 'chacinados', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Bondiola curada Paladini 200g', 'Paladini', 'chacinados', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Salame tandilero Estancias del Sur 250g', 'Estancias del Sur', 'chacinados', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Sopressata Paladini 200g', 'Paladini', 'chacinados', 4800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Salame picado grueso Swift 300g', 'Swift', 'chacinados', 4000.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Longaniza seca campo 200g', NULL, 'chacinados', 3200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Chorizo seco Paladini 250g', 'Paladini', 'chacinados', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Salame tipo italiano campo 300g', NULL, 'chacinados', 3800.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),

-- ============================================================
-- FIAMBRES (11 productos)
-- ============================================================
('Jamón cocido natural Paladini 150g', 'Paladini', 'fiambres', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Jamón cocido Paladini x kg', 'Paladini', 'fiambres', 12500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Jamón crudo Paladini 150g', 'Paladini', 'fiambres', 5200.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Paleta cocida Paladini 150g', 'Paladini', 'fiambres', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Lomito ahumado Paladini 150g', 'Paladini', 'fiambres', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Jamón cocido Swift 200g', 'Swift', 'fiambres', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Paleta cocida Swift 200g', 'Swift', 'fiambres', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Jamón crudo x kg', NULL, 'fiambres', 18000.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Mortadela Paladini 150g', 'Paladini', 'fiambres', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Fiambre de cerdo cocido Super Valu 200g', 'Super Valu', 'fiambres', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Matambre de cerdo cocido x kg', NULL, 'fiambres', 14000.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),

-- ============================================================
-- HAMBURGUESAS (12 productos)
-- ============================================================
('Hamburguesa casera de carne', NULL, 'hamburguesas', 7500.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Hamburguesa casera de carne y cerdo', NULL, 'hamburguesas', 7000.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Hamburguesa Paty clásica x4', 'Paty', 'hamburguesas', 5200.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Hamburguesa Paty clásica x8', 'Paty', 'hamburguesas', 9500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Hamburguesa Good Mark clásica x4', 'Good Mark', 'hamburguesas', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Hamburguesa Good Mark clásica x8', 'Good Mark', 'hamburguesas', 6800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Hamburguesa Swift premium x4', 'Swift', 'hamburguesas', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Hamburguesa Paty doble x4', 'Paty', 'hamburguesas', 6200.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Hamburguesa casera de pollo', NULL, 'hamburguesas', 7200.00, 'seed', 0.5, 50, FALSE, TRUE, 'carniceria', '{"unit": "kg", "sku_prefix": "CARN"}'::jsonb),
('Hamburguesa de carne Super Valu x4', 'Super Valu', 'hamburguesas', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Hamburguesa de garbanzo Good Mark x4', 'Good Mark', 'hamburguesas', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb),
('Medallones de carne Swift x4', 'Swift', 'hamburguesas', 4800.00, 'seed', 0.5, 75, FALSE, TRUE, 'carniceria', '{"unit": "unidad", "sku_prefix": "CARN"}'::jsonb)

ON CONFLICT DO NOTHING;
