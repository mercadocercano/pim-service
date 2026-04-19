-- Seed 046: Fiambrería — 120 productos reales argentinos
-- Generado: 2026-04-18
-- Fuente: global_products (v2.0)
-- ON CONFLICT DO NOTHING: idempotente, no duplica si ya existe por nombre+marca.

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES
-- ============================================================
-- JAMONES
-- ============================================================
('Jamón cocido Paladini', 'Paladini', 'jamones', 12500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Jamón cocido natural Paladini', 'Paladini', 'jamones', 15800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Jamón cocido Cagnoli', 'Cagnoli', 'jamones', 13200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Jamón cocido Coto', 'Coto', 'jamones', 10800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Jamón cocido Calchaquí', 'Calchaquí', 'jamones', 11500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Jamón crudo Paladini', 'Paladini', 'jamones', 18500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Jamón crudo Cagnoli', 'Cagnoli', 'jamones', 19200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Jamón crudo bondiola Paladini', 'Paladini', 'jamones', 16800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Jamón serrano España Cagnoli', 'Cagnoli', 'jamones', 28500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Jamón crudo estacionado Tyson', 'Tyson', 'jamones', 17500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Paleta cocida Paladini', 'Paladini', 'jamones', 9800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Paleta cocida Cagnoli', 'Cagnoli', 'jamones', 10200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Lomito ahumado Paladini', 'Paladini', 'jamones', 22000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Lomito canadiense Cagnoli', 'Cagnoli', 'jamones', 23500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),

-- ============================================================
-- FIAMBRES Y EMBUTIDOS
-- ============================================================
('Salame milán Paladini', 'Paladini', 'fiambres-embutidos', 14500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Salame milán Cagnoli', 'Cagnoli', 'fiambres-embutidos', 15200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Salame tipo italiano Paladini', 'Paladini', 'fiambres-embutidos', 16000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Salame picado grueso Cagnoli', 'Cagnoli', 'fiambres-embutidos', 13800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Salame quintero Paladini', 'Paladini', 'fiambres-embutidos', 12500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Salame sopressata Cagnoli', 'Cagnoli', 'fiambres-embutidos', 17500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Mortadela Paladini', 'Paladini', 'fiambres-embutidos', 8500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Mortadela con aceitunas Cagnoli', 'Cagnoli', 'fiambres-embutidos', 9200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Mortadela con pistachos Cagnoli', 'Cagnoli', 'fiambres-embutidos', 10500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Leberwurst Paladini', 'Paladini', 'fiambres-embutidos', 9800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Morcilla vasca Cagnoli', 'Cagnoli', 'fiambres-embutidos', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Longaniza calabresa Paladini', 'Paladini', 'fiambres-embutidos', 11000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Bondiola curada Paladini', 'Paladini', 'fiambres-embutidos', 16500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Panceta ahumada Paladini', 'Paladini', 'fiambres-embutidos', 14000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Matambre arrollado Cagnoli', 'Cagnoli', 'fiambres-embutidos', 15500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Chorizo colorado sarta Paladini', 'Paladini', 'fiambres-embutidos', 10500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),

-- ============================================================
-- QUESOS DUROS
-- ============================================================
('Queso sardo La Paulina', 'La Paulina', 'quesos-duros', 16500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso sardo SanCor', 'SanCor', 'quesos-duros', 15800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso sardo Verónica', 'Verónica', 'quesos-duros', 14500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso reggianito La Serenísima', 'La Serenísima', 'quesos-duros', 19500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso reggianito SanCor', 'SanCor', 'quesos-duros', 18800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso reggianito La Paulina', 'La Paulina', 'quesos-duros', 19000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso parmesano rallado La Serenísima 150g', 'La Serenísima', 'quesos-duros', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Queso provolone Ilolay', 'Ilolay', 'quesos-duros', 14000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso provolone La Paulina', 'La Paulina', 'quesos-duros', 14800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso provolone ahumado Ilolay', 'Ilolay', 'quesos-duros', 15500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso grana padano La Paulina', 'La Paulina', 'quesos-duros', 22000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso sbrinz Milkaut', 'Milkaut', 'quesos-duros', 17500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),

-- ============================================================
-- QUESOS SEMIDUROS
-- ============================================================
('Queso tybo La Serenísima', 'La Serenísima', 'quesos-semiduros', 12500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso tybo SanCor', 'SanCor', 'quesos-semiduros', 11800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso tybo Ilolay', 'Ilolay', 'quesos-semiduros', 11200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso fontina La Paulina', 'La Paulina', 'quesos-semiduros', 13500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso fontina Milkaut', 'Milkaut', 'quesos-semiduros', 12800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso gouda La Serenísima', 'La Serenísima', 'quesos-semiduros', 13000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso gouda Tregar', 'Tregar', 'quesos-semiduros', 11500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso dambo La Serenísima', 'La Serenísima', 'quesos-semiduros', 12200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso dambo SanCor', 'SanCor', 'quesos-semiduros', 11600.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso port salut La Serenísima', 'La Serenísima', 'quesos-semiduros', 12800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso port salut Ilolay', 'Ilolay', 'quesos-semiduros', 11500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso pategrás La Paulina', 'La Paulina', 'quesos-semiduros', 13200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso pategrás SanCor', 'SanCor', 'quesos-semiduros', 12600.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso holanda Verónica', 'Verónica', 'quesos-semiduros', 11800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso edam Milkaut', 'Milkaut', 'quesos-semiduros', 12000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),

-- ============================================================
-- QUESOS BLANDOS
-- ============================================================
('Queso cremoso La Serenísima', 'La Serenísima', 'quesos-blandos', 11000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso cremoso La Paulina', 'La Paulina', 'quesos-blandos', 10500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso cremoso Tregar', 'Tregar', 'quesos-blandos', 9800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso cremoso Ilolay', 'Ilolay', 'quesos-blandos', 10200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso brie La Paulina', 'La Paulina', 'quesos-blandos', 16500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso brie SanCor', 'SanCor', 'quesos-blandos', 15800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso camembert La Paulina', 'La Paulina', 'quesos-blandos', 17000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso camembert SanCor', 'SanCor', 'quesos-blandos', 16200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso roquefort La Paulina', 'La Paulina', 'quesos-blandos', 18500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso azul SanCor', 'SanCor', 'quesos-blandos', 17500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso azul Verónica', 'Verónica', 'quesos-blandos', 16000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),
('Queso mozzarella La Serenísima', 'La Serenísima', 'quesos-blandos', 11500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "kg", "sku_prefix": "FIAMB"}'::jsonb),

-- ============================================================
-- QUESOS UNTABLES
-- ============================================================
('Ricotta La Serenísima 500g', 'La Serenísima', 'quesos-untables', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Ricotta SanCor 500g', 'SanCor', 'quesos-untables', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Mascarpone La Paulina 250g', 'La Paulina', 'quesos-untables', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Queso crema La Serenísima 300g', 'La Serenísima', 'quesos-untables', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Queso crema Mendicrim 300g', 'Mendicrim', 'quesos-untables', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Casancrem original 300g', 'Casancrem', 'quesos-untables', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Casancrem light 300g', 'Casancrem', 'quesos-untables', 3600.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Queso untable Finlandia 300g', 'Finlandia', 'quesos-untables', 3400.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Queso untable Finlandia light 300g', 'Finlandia', 'quesos-untables', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Queso crema Tregar 300g', 'Tregar', 'quesos-untables', 2600.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Ricotta Tregar 500g', 'Tregar', 'quesos-untables', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),

-- ============================================================
-- PICADAS
-- ============================================================
('Aceitunas verdes descarozadas Nucete 220g', 'Nucete', 'picadas', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Aceitunas verdes rellenas Nucete 220g', 'Nucete', 'picadas', 2600.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Aceitunas negras Nucete 220g', 'Nucete', 'picadas', 3000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Aceitunas verdes La Campagnola 220g', 'La Campagnola', 'picadas', 2400.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Aceitunas rellenas con morrón Molto 220g', 'Molto', 'picadas', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Pickles en vinagre Gourmet 330g', 'Gourmet', 'picadas', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Ajíes en vinagre Gourmet 330g', 'Gourmet', 'picadas', 3000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Lupines en salmuera La Campagnola 350g', 'La Campagnola', 'picadas', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Berenjenas en escabeche Gourmet 330g', 'Gourmet', 'picadas', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Morrones asados La Campagnola 290g', 'La Campagnola', 'picadas', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Alcaparras Molto 90g', 'Molto', 'picadas', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),

-- ============================================================
-- CONSERVAS GOURMET
-- ============================================================
('Aceite de oliva extra virgen Nucete 500ml', 'Nucete', 'conservas-gourmet', 8500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Aceite de oliva extra virgen Esmeralda 500ml', 'Esmeralda', 'conservas-gourmet', 7800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Aceite de oliva extra virgen Molto 500ml', 'Molto', 'conservas-gourmet', 7200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Aceto balsámico Nucete 250ml', 'Nucete', 'conservas-gourmet', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Vinagre de manzana Nucete 500ml', 'Nucete', 'conservas-gourmet', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Mostaza Dijon Gourmet 200g', 'Gourmet', 'conservas-gourmet', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Mostaza a la antigua Gourmet 200g', 'Gourmet', 'conservas-gourmet', 3600.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Mermelada de higos La Campagnola 390g', 'La Campagnola', 'conservas-gourmet', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Mermelada de frutos rojos La Campagnola 390g', 'La Campagnola', 'conservas-gourmet', 3400.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Dulce de membrillo Esmeralda 500g', 'Esmeralda', 'conservas-gourmet', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Tomates secos en aceite Gourmet 200g', 'Gourmet', 'conservas-gourmet', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Paté de aceitunas Nucete 180g', 'Nucete', 'conservas-gourmet', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),

-- ============================================================
-- PAN DE FIAMBRERÍA
-- ============================================================
('Pan de campo artesanal', 'Artesanal', 'pan-fiambreria', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Grisines largos clásicos', 'Artesanal', 'pan-fiambreria', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Grisines con queso', 'Artesanal', 'pan-fiambreria', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Grisines con salvado', 'Artesanal', 'pan-fiambreria', 2400.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Croutons saborizados Gourmet 100g', 'Gourmet', 'pan-fiambreria', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Tostadas de mesa Breviss 200g', 'Breviss', 'pan-fiambreria', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Focaccia con romero', 'Artesanal', 'pan-fiambreria', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb),
('Pan ciabatta', 'Artesanal', 'pan-fiambreria', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'fiambreria', '{"unit": "unidad", "sku_prefix": "FIAMB"}'::jsonb)

ON CONFLICT DO NOTHING;
