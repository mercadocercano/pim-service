-- Seed 044: Verdulería — 120 productos reales argentinos
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
-- VERDURAS DE HOJA (verduras-hoja)
-- ============================================================
('Lechuga criolla', NULL, 'verduras-hoja', 1800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Lechuga mantecosa', NULL, 'verduras-hoja', 2200.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Lechuga morada', NULL, 'verduras-hoja', 2500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Espinaca', NULL, 'verduras-hoja', 3200.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Acelga', NULL, 'verduras-hoja', 2000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Rúcula', NULL, 'verduras-hoja', 2800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Radicheta', NULL, 'verduras-hoja', 2200.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Berro', NULL, 'verduras-hoja', 2500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- VERDURAS DE FRUTO (verduras-fruto)
-- ============================================================
('Tomate redondo', NULL, 'verduras-fruto', 3500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Tomate perita', NULL, 'verduras-fruto', 3200.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Tomate cherry', NULL, 'verduras-fruto', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Morrón rojo', NULL, 'verduras-fruto', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Morrón verde', NULL, 'verduras-fruto', 3800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Morrón amarillo', NULL, 'verduras-fruto', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Zapallito redondo', NULL, 'verduras-fruto', 2800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Zapallito largo', NULL, 'verduras-fruto', 2500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Berenjena', NULL, 'verduras-fruto', 3500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Pepino', NULL, 'verduras-fruto', 3000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Zapallo anco', NULL, 'verduras-fruto', 1800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Zapallo plomo', NULL, 'verduras-fruto', 1500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Calabaza', NULL, 'verduras-fruto', 1800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Chaucha', NULL, 'verduras-fruto', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- TUBÉRCULOS (tuberculos)
-- ============================================================
('Papa negra', NULL, 'tuberculos', 1800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Papa blanca', NULL, 'tuberculos', 1600.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Papa rosada', NULL, 'tuberculos', 2000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Batata', NULL, 'tuberculos', 2200.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Mandioca', NULL, 'tuberculos', 2500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- RAÍCES (raices)
-- ============================================================
('Zanahoria', NULL, 'raices', 1800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Remolacha', NULL, 'raices', 2200.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Nabo', NULL, 'raices', 2000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Rabanito', NULL, 'raices', 2500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Jengibre', NULL, 'raices', 8000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- BULBOS (bulbos)
-- ============================================================
('Cebolla blanca', NULL, 'bulbos', 1500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Cebolla morada', NULL, 'bulbos', 2200.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Cebolla de verdeo', NULL, 'bulbos', 1800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Ajo', NULL, 'bulbos', 6000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Ajo colorado', NULL, 'bulbos', 7000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Puerro', NULL, 'bulbos', 3000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Echalote', NULL, 'bulbos', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- CRUCÍFERAS (cruciferas)
-- ============================================================
('Brócoli', NULL, 'cruciferas', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Coliflor', NULL, 'cruciferas', 4000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Repollo blanco', NULL, 'cruciferas', 2000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Repollo colorado', NULL, 'cruciferas', 2500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Repollito de Bruselas', NULL, 'cruciferas', 6000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- LEGUMBRES FRESCAS (legumbres-frescas)
-- ============================================================
('Choclo', NULL, 'legumbres-frescas', 1500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Arvejas frescas', NULL, 'legumbres-frescas', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Habas frescas', NULL, 'legumbres-frescas', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Choclo desgranado', NULL, 'legumbres-frescas', 4000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- FRUTAS CÍTRICOS (frutas-citricos)
-- ============================================================
('Naranja de jugo', NULL, 'frutas-citricos', 2000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Naranja de mesa', NULL, 'frutas-citricos', 2500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Mandarina', NULL, 'frutas-citricos', 2800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Limón', NULL, 'frutas-citricos', 2200.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Pomelo rosado', NULL, 'frutas-citricos', 2500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Pomelo blanco', NULL, 'frutas-citricos', 2200.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Lima', NULL, 'frutas-citricos', 3500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- FRUTAS DE CAROZO (frutas-carozo)
-- ============================================================
('Durazno', NULL, 'frutas-carozo', 4000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Ciruela roja', NULL, 'frutas-carozo', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Ciruela negra', NULL, 'frutas-carozo', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Damasco', NULL, 'frutas-carozo', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Pelón', NULL, 'frutas-carozo', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Cereza', NULL, 'frutas-carozo', 12000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- FRUTAS DE PEPITA (frutas-pepita)
-- ============================================================
('Manzana roja deliciosa', NULL, 'frutas-pepita', 3000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Manzana verde Granny Smith', NULL, 'frutas-pepita', 3200.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Manzana Gala', NULL, 'frutas-pepita', 3500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Pera Williams', NULL, 'frutas-pepita', 3800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Pera Packham', NULL, 'frutas-pepita', 3500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Membrillo', NULL, 'frutas-pepita', 3000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Uva negra', NULL, 'frutas-pepita', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Uva blanca', NULL, 'frutas-pepita', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Uva rosada', NULL, 'frutas-pepita', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- FRUTAS TROPICALES (frutas-tropicales)
-- ============================================================
('Banana Ecuador', NULL, 'frutas-tropicales', 2500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Banana de Formosa', NULL, 'frutas-tropicales', 2800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Ananá', NULL, 'frutas-tropicales', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Mango', NULL, 'frutas-tropicales', 6000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Kiwi', NULL, 'frutas-tropicales', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Palta Hass', NULL, 'frutas-tropicales', 7000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Maracuyá', NULL, 'frutas-tropicales', 8000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Papaya', NULL, 'frutas-tropicales', 6500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Frutilla', NULL, 'frutas-tropicales', 6000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Arándano', NULL, 'frutas-tropicales', 10000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Frambuesa', NULL, 'frutas-tropicales', 12000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Sandía', NULL, 'frutas-tropicales', 1500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Melón', NULL, 'frutas-tropicales', 2500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- FRUTAS SECAS (frutas-secas)
-- ============================================================
('Nuez con cáscara', NULL, 'frutas-secas', 8000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Nuez pelada', NULL, 'frutas-secas', 18000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Almendra con cáscara', NULL, 'frutas-secas', 10000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Almendra pelada', NULL, 'frutas-secas', 20000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Maní tostado con cáscara', NULL, 'frutas-secas', 4000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Maní pelado salado', NULL, 'frutas-secas', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Castañas de cajú', NULL, 'frutas-secas', 22000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Mix de frutos secos Georgalos 250g', 'Georgalos', 'frutas-secas', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Nueces Georgalos 250g', 'Georgalos', 'frutas-secas', 6000.00, 'seed', 0.5, 75, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Almendras Georgalos 250g', 'Georgalos', 'frutas-secas', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Maní tostado Marolio 400g', 'Marolio', 'frutas-secas', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Pasas de uva', NULL, 'frutas-secas', 6000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),
('Ciruelas desecadas', NULL, 'frutas-secas', 7000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "kg", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- AROMÁTICAS (aromaticas)
-- ============================================================
('Perejil', NULL, 'aromaticas', 1200.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Cilantro', NULL, 'aromaticas', 1500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Albahaca', NULL, 'aromaticas', 1800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Orégano fresco', NULL, 'aromaticas', 1500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Romero fresco', NULL, 'aromaticas', 1500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Ciboulette', NULL, 'aromaticas', 1800.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Menta fresca', NULL, 'aromaticas', 1500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Tomillo fresco', NULL, 'aromaticas', 1500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),

-- ============================================================
-- HUEVOS (huevos)
-- ============================================================
('Huevos blancos x12', NULL, 'huevos', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "docena", "sku_prefix": "VERD"}'::jsonb),
('Huevos colorados x12', NULL, 'huevos', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "docena", "sku_prefix": "VERD"}'::jsonb),
('Huevos de campo x6', NULL, 'huevos', 4000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Huevos de campo x12', NULL, 'huevos', 7500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "docena", "sku_prefix": "VERD"}'::jsonb),
('Huevos orgánicos x6', NULL, 'huevos', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Huevos Granja Iris x12', 'Granja Iris', 'huevos', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'verduleria', '{"unit": "docena", "sku_prefix": "VERD"}'::jsonb),
('Huevos Las Dinas x12', 'Las Dinas', 'huevos', 5200.00, 'seed', 0.5, 75, FALSE, TRUE, 'verduleria', '{"unit": "docena", "sku_prefix": "VERD"}'::jsonb),
('Huevos Oro x12', 'Huevos Oro', 'huevos', 5000.00, 'seed', 0.5, 75, FALSE, TRUE, 'verduleria', '{"unit": "docena", "sku_prefix": "VERD"}'::jsonb),
('Huevos Granja Iris x30', 'Granja Iris', 'huevos', 12000.00, 'seed', 0.5, 75, FALSE, TRUE, 'verduleria', '{"unit": "unidad", "sku_prefix": "VERD"}'::jsonb),
('Huevos de codorniz x12', NULL, 'huevos', 3000.00, 'seed', 0.5, 50, FALSE, TRUE, 'verduleria', '{"unit": "docena", "sku_prefix": "VERD"}'::jsonb)

ON CONFLICT DO NOTHING;
