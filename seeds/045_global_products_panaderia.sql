-- Seed: Panadería — 93 productos reales argentinos
-- Generado: 2026-04-18
-- Fuente: global_products (v2.0)

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES
-- ============================================================
-- PAN FRESCO (elaboración propia)
-- ============================================================
('Pan francés', NULL, 'pan-fresco', 2500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Pan Felipe', NULL, 'pan-fresco', 2800.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Pan lactal', NULL, 'pan-fresco', 3200.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Pan integral', NULL, 'pan-fresco', 3500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Pan de centeno', NULL, 'pan-fresco', 4000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Pan de salvado', NULL, 'pan-fresco', 3800.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Pan de campo', NULL, 'pan-fresco', 3000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Pan árabe', NULL, 'pan-fresco', 3200.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Pan de viena', NULL, 'pan-fresco', 3500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Flautitas', NULL, 'pan-fresco', 2800.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),

-- ============================================================
-- FACTURAS
-- ============================================================
('Medialunas de manteca', NULL, 'facturas', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Medialunas de grasa', NULL, 'facturas', 3800.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Vigilantes', NULL, 'facturas', 4000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Cañoncitos de dulce de leche', NULL, 'facturas', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Bolas de fraile', NULL, 'facturas', 4200.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Sacramentos', NULL, 'facturas', 4000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Tortitas negras', NULL, 'facturas', 3800.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Churros rellenos de dulce de leche', NULL, 'facturas', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Berlinesas', NULL, 'facturas', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),

-- ============================================================
-- GALLETITAS SECAS (elaboración propia)
-- ============================================================
('Galletitas de manteca', NULL, 'galletitas-secas', 3500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Galletitas de limón', NULL, 'galletitas-secas', 3500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Pepas de membrillo', NULL, 'galletitas-secas', 4000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Pepas de batata', NULL, 'galletitas-secas', 4000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Galletitas surtidas', NULL, 'galletitas-secas', 3800.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Galletitas de coco', NULL, 'galletitas-secas', 3800.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Palmeritas', NULL, 'galletitas-secas', 4200.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),

-- ============================================================
-- GALLETITAS SALADAS
-- ============================================================
('Grisines', NULL, 'galletitas-saladas', 3000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Bay biscuit', NULL, 'galletitas-saladas', 3200.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Criollitas', NULL, 'galletitas-saladas', 3000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Galletas de agua', NULL, 'galletitas-saladas', 2800.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Bizcochos de grasa', NULL, 'galletitas-saladas', 3500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),

-- ============================================================
-- TORTAS
-- ============================================================
('Torta de cumpleaños clásica', NULL, 'tortas', 18000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Rogel', NULL, 'tortas', 22000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Selva negra', NULL, 'tortas', 20000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Lemon pie', NULL, 'tortas', 16000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Chocotorta', NULL, 'tortas', 15000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Torta de ricota', NULL, 'tortas', 14000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Torta de mousse de chocolate', NULL, 'tortas', 19000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Tarta de frutillas con crema', NULL, 'tortas', 16000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),

-- ============================================================
-- BUDINES
-- ============================================================
('Budín inglés', NULL, 'budines', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Budín marmolado', NULL, 'budines', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Budín de naranja', NULL, 'budines', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Budín de limón', NULL, 'budines', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Budín de banana', NULL, 'budines', 5200.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Budín de zanahoria', NULL, 'budines', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),

-- ============================================================
-- BIZCOCHOS
-- ============================================================
('Bizcochuelo', NULL, 'bizcochos', 6000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Vainillas', NULL, 'bizcochos', 3500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Pionono', NULL, 'bizcochos', 4000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Bizcochuelo de chocolate', NULL, 'bizcochos', 6500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),

-- ============================================================
-- SANDWICHES
-- ============================================================
('Sándwiches de miga surtidos x12', NULL, 'sandwiches', 8000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Sándwiches de miga de jamón y queso x12', NULL, 'sandwiches', 7500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Sándwiches triples x12', NULL, 'sandwiches', 9000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Sándwiches de miga de palmitos x12', NULL, 'sandwiches', 9500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Sándwiches de miga de atún x12', NULL, 'sandwiches', 9000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),

-- ============================================================
-- CONFITERÍA
-- ============================================================
('Masas finas surtidas', NULL, 'confiteria', 12000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Alfajores de maicena', NULL, 'confiteria', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Alfajores de chocolate', NULL, 'confiteria', 6000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Bombones surtidos', NULL, 'confiteria', 10000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Cañoncitos de crema', NULL, 'confiteria', 5500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Conitos de dulce de leche', NULL, 'confiteria', 5000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "docena", "sku_prefix": "PANA"}'::jsonb),
('Masitas secas surtidas', NULL, 'confiteria', 8000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Pastafrola de membrillo', NULL, 'confiteria', 6000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Pastafrola de batata', NULL, 'confiteria', 6000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),

-- ============================================================
-- PRODUCTOS INTEGRALES
-- ============================================================
('Pan integral con semillas', NULL, 'productos-integrales', 4200.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Galletitas integrales de avena', NULL, 'productos-integrales', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Budín integral de manzana', NULL, 'productos-integrales', 5800.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Galletas de arroz integrales', NULL, 'productos-integrales', 3500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),
('Pan multicereal', NULL, 'productos-integrales', 4500.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "kg", "sku_prefix": "PANA"}'::jsonb),

-- ============================================================
-- INSUMOS PANADERÍA (marcas reales)
-- ============================================================
('Harina 000 Pureza 25kg', 'Pureza', 'insumos-panaderia', 18000.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Harina 0000 Pureza 25kg', 'Pureza', 'insumos-panaderia', 19000.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Harina 000 Blancaflor 1kg', 'Blancaflor', 'insumos-panaderia', 1400.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Harina 000 Cañuelas 25kg', 'Cañuelas', 'insumos-panaderia', 16000.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Harina 000 Favorita 25kg', 'Favorita', 'insumos-panaderia', 14000.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Levadura fresca Calsa 500g', 'Calsa', 'insumos-panaderia', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Levadura seca Fleischmann 500g', 'Fleischmann', 'insumos-panaderia', 6000.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Levadura seca Calsa 500g', 'Calsa', 'insumos-panaderia', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Manteca La Serenísima 200g', 'La Serenísima', 'insumos-panaderia', 3000.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Manteca Ilolay 200g', 'Ilolay', 'insumos-panaderia', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Azúcar Ledesma 1kg', 'Ledesma', 'insumos-panaderia', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Azúcar Ledesma 50kg', 'Ledesma', 'insumos-panaderia', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Dulce de leche La Serenísima 1kg', 'La Serenísima', 'insumos-panaderia', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Dulce de leche Ilolay 1kg', 'Ilolay', 'insumos-panaderia', 5000.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Dulce de leche repostero La Serenísima 1kg', 'La Serenísima', 'insumos-panaderia', 5800.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Grasa vacuna Primer Precio 500g', 'Primer Precio', 'insumos-panaderia', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Esencia de vainilla Marolio 100ml', 'Marolio', 'insumos-panaderia', 800.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Cacao amargo Aguila 150g', 'Aguila', 'insumos-panaderia', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Chocolate cobertura Aguila 150g', 'Aguila', 'insumos-panaderia', 4000.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Crema de leche La Serenísima 200ml', 'La Serenísima', 'insumos-panaderia', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Mermelada de damasco Arcor 454g', 'Arcor', 'insumos-panaderia', 2800.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Membrillo Arcor 500g', 'Arcor', 'insumos-panaderia', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Huevos frescos x30', NULL, 'insumos-panaderia', 8000.00, 'seed', 0.5, 50, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Sal fina Celusal 500g', 'Celusal', 'insumos-panaderia', 600.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb),
('Mejorador de pan Fleischmann 1kg', 'Fleischmann', 'insumos-panaderia', 7000.00, 'seed', 0.5, 75, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'::jsonb)

ON CONFLICT DO NOTHING;
