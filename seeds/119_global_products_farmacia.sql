-- =============================================================
-- SEED 119: Farmacia (business_type = 'farmacia')
-- Research date: 2026-04-25
-- Source: Catálogo market research NEA/Posadas 2026
--   - Farmacias de la zona (Farmacity online, Dr. Ahorro, Del Pueblo)
--   - Precios de góndola/venta libre relevados abril 2026
-- Zona: Posadas, Misiones (NEA)
-- Objetivo: >= 100 productos verificados
-- IMPORTANTE: Solo productos sin receta / venta libre
-- Categorías: higiene-personal, papel-higiene, cuidado-bebe,
--             medicamentos-otc, cuidado-corporal
-- =============================================================

-- ============================================================
-- CATEGORÍA: medicamentos-otc — Venta libre, sin receta
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Antisépticos y primeros auxilios
('Alcohol etílico 70% 250ml Bioxcin', 'Bioxcin', 'farmacia', 'medicamentos-otc', 600.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "concentracion": "70%", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Alcohol etílico 70% 500ml Bioxcin', 'Bioxcin', 'farmacia', 'medicamentos-otc', 1000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "concentracion": "70%", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Alcohol etílico 96% 250ml Bioxcin', 'Bioxcin', 'farmacia', 'medicamentos-otc', 700.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "concentracion": "96%", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Alcohol etílico 96% 500ml Bioxcin', 'Bioxcin', 'farmacia', 'medicamentos-otc', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "concentracion": "96%", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Agua oxigenada 10 volúmenes 250ml Bioxcin', 'Bioxcin', 'farmacia', 'medicamentos-otc', 500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "volumen": "10vol", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Agua oxigenada 20 volúmenes 250ml Bioxcin', 'Bioxcin', 'farmacia', 'medicamentos-otc', 550.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "volumen": "20vol", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Agua oxigenada 30 volúmenes 250ml Bioxcin', 'Bioxcin', 'farmacia', 'medicamentos-otc', 600.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "volumen": "30vol", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Yodo povidona solución 50ml Isodine', 'Isodine', 'farmacia', 'medicamentos-otc', 900.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Algodón hidrófilo 100g', NULL, 'farmacia', 'medicamentos-otc', 600.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad 100g", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Algodón hidrófilo 500g', NULL, 'farmacia', 'medicamentos-otc', 2200.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad 500g", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Gasa estéril 10x10cm x10 unidades', NULL, 'farmacia', 'medicamentos-otc', 450.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "pack 10u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Venda elástica tubular 5cm x4.5m', NULL, 'farmacia', 'medicamentos-otc', 700.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad", "ancho": "5cm", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Venda elástica tubular 10cm x4.5m', NULL, 'farmacia', 'medicamentos-otc', 950.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad", "ancho": "10cm", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Venda elástica autoadhesiva 7.5cm Nexcare', 'Nexcare', 'farmacia', 'medicamentos-otc', 1200.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "ancho": "7.5cm", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Curita curaciones surtidas x40 Band-Aid', 'Band-Aid', 'farmacia', 'medicamentos-otc', 1100.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "caja 40u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Curita curaciones surtidas x20 genérico', NULL, 'farmacia', 'medicamentos-otc', 400.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "caja 20u", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Instrumentos médicos básicos
('Termómetro digital axilar con memoria', NULL, 'farmacia', 'medicamentos-otc', 3500.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Termómetro digital infrarrojo frente y oído', NULL, 'farmacia', 'medicamentos-otc', 9500.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tensiómetro digital automático muñeca', NULL, 'farmacia', 'medicamentos-otc', 18000.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Nebulizador de pistón básico 0.3ml/min', NULL, 'farmacia', 'medicamentos-otc', 18000.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Glucómetro básico con tiras x10', NULL, 'farmacia', 'medicamentos-otc', 12000.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "kit", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tiras reactivas glucosa x50 unidades', NULL, 'farmacia', 'medicamentos-otc', 8000.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "caja 50u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Lancetas descartables x100 unidades', NULL, 'farmacia', 'medicamentos-otc', 1500.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "caja 100u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Jeringa descartable 5ml sin aguja x10', NULL, 'farmacia', 'medicamentos-otc', 800.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "pack 10u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Jeringa descartable 10ml sin aguja x10', NULL, 'farmacia', 'medicamentos-otc', 950.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "pack 10u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Jeringa descartable 20ml sin aguja x5', NULL, 'farmacia', 'medicamentos-otc', 900.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "pack 5u", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Suplementos y vitaminas (venta libre)
('Vitamina C 1g efervescente x10 comprimidos Redoxon', 'Redoxon', 'farmacia', 'medicamentos-otc', 1800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "caja 10u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Vitamina C 1g efervescente x20 comprimidos Redoxon', 'Redoxon', 'farmacia', 'medicamentos-otc', 3200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "caja 20u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Complejo B x30 comprimidos Laboratorio Elea', 'Elea', 'farmacia', 'medicamentos-otc', 2500.00, 'catalog-researcher', 0.80, 65, TRUE, TRUE, '{"unit": "caja 30u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Omega 3 1000mg x30 cápsulas', NULL, 'farmacia', 'medicamentos-otc', 3000.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "caja 30u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Calcio + Vitamina D x30 comprimidos', NULL, 'farmacia', 'medicamentos-otc', 2800.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "caja 30u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Magnesio 250mg x30 comprimidos', NULL, 'farmacia', 'medicamentos-otc', 2500.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "caja 30u", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Tests y preservativos
('Test de embarazo x1 unidad Confirm', 'Confirm', 'farmacia', 'medicamentos-otc', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Test de embarazo x2 unidades', NULL, 'farmacia', 'medicamentos-otc', 1800.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "pack 2u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Preservativos x3 unidades Tulipán', 'Tulipán', 'farmacia', 'medicamentos-otc', 1000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "pack 3u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Preservativos x6 unidades Tulipán', 'Tulipán', 'farmacia', 'medicamentos-otc', 1800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "pack 6u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Preservativos x12 unidades Tulipán', 'Tulipán', 'farmacia', 'medicamentos-otc', 3200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "pack 12u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Preservativos x3 unidades Prime', 'Prime', 'farmacia', 'medicamentos-otc', 900.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "pack 3u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Preservativos x6 unidades Prime', 'Prime', 'farmacia', 'medicamentos-otc', 1600.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "pack 6u", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============================================================
-- CATEGORÍA: higiene-personal — Higiene básica sin receta
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Protectores solares (alta rotación en NEA)
('Protector solar FPS 30 200ml Sundown', 'Sundown', 'farmacia', 'higiene-personal', 4500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "fps": 30, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Protector solar FPS 50 200ml Sundown', 'Sundown', 'farmacia', 'higiene-personal', 5500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "fps": 50, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Protector solar FPS 50+ 200ml Sundown kids', 'Sundown', 'farmacia', 'higiene-personal', 6000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "fps": 50, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Protector solar FPS 30 150ml Photoderm', 'Photoderm', 'farmacia', 'higiene-personal', 5500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "fps": 30, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Protector solar FPS 50 150ml Photoderm', 'Photoderm', 'farmacia', 'higiene-personal', 6500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "fps": 50, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Protector solar FPS 50+ Rostro 50ml Anasol', 'Anasol', 'farmacia', 'higiene-personal', 4000.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "fps": 50, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Repelente de insectos spray 200ml Off!', 'Off!', 'farmacia', 'higiene-personal', 2800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Repelente de insectos crema 100ml Off!', 'Off!', 'farmacia', 'higiene-personal', 2200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Higiene bucal
('Cepillo de dientes adulto blando Oral-B', 'Oral-B', 'farmacia', 'higiene-personal', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cepillo de dientes adulto medio Oral-B', 'Oral-B', 'farmacia', 'higiene-personal', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pasta dental 90ml Colgate Triple Acción', 'Colgate', 'farmacia', 'higiene-personal', 1500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pasta dental 90ml Colgate Blanqueadora', 'Colgate', 'farmacia', 'higiene-personal', 1600.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Hilo dental 50m Oral-B', 'Oral-B', 'farmacia', 'higiene-personal', 1000.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Enjuague bucal 500ml Listerine Cool Mint', 'Listerine', 'farmacia', 'higiene-personal', 2800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Desodorantes
('Desodorante aerosol 150ml Rexona Women', 'Rexona', 'farmacia', 'higiene-personal', 1800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "genero": "mujer", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Desodorante aerosol 150ml Rexona Men', 'Rexona', 'farmacia', 'higiene-personal', 1800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "genero": "hombre", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Desodorante roll-on 50ml Dove Original', 'Dove', 'farmacia', 'higiene-personal', 1400.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Desodorante roll-on 50ml Nivea Pearl Beauty', 'Nivea', 'farmacia', 'higiene-personal', 1500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Shampoo y acondicionador
('Shampoo 400ml Pantene Pro-V', 'Pantene', 'farmacia', 'higiene-personal', 2500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Shampoo 400ml Head & Shoulders', 'Head & Shoulders', 'farmacia', 'higiene-personal', 2800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Acondicionador 400ml Pantene Pro-V', 'Pantene', 'farmacia', 'higiene-personal', 2600.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============================================================
-- CATEGORÍA: cuidado-bebe — Productos para bebés
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Pañales
('Pañales descartables recién nacido x40 Pampers', 'Pampers', 'farmacia', 'cuidado-bebe', 5500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "paquete 40u", "talle": "RN", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pañales descartables talle G x36 Pampers', 'Pampers', 'farmacia', 'cuidado-bebe', 6200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "paquete 36u", "talle": "G", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pañales descartables talle XG x30 Pampers', 'Pampers', 'farmacia', 'cuidado-bebe', 6500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "paquete 30u", "talle": "XG", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pañales descartables talle M x40 Huggies', 'Huggies', 'farmacia', 'cuidado-bebe', 5800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "paquete 40u", "talle": "M", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pañales descartables talle G x36 Huggies', 'Huggies', 'farmacia', 'cuidado-bebe', 6000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "paquete 36u", "talle": "G", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pañales adulto talle M x8 Tena', 'Tena', 'farmacia', 'cuidado-bebe', 4500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "paquete 8u", "talle": "M", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pañales adulto talle G x8 Tena', 'Tena', 'farmacia', 'cuidado-bebe', 4800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "paquete 8u", "talle": "G", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Toallitas húmedas y cremas
('Toallitas húmedas bebé x80 unidades Pampers', 'Pampers', 'farmacia', 'cuidado-bebe', 2200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "paquete 80u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Toallitas húmedas bebé x80 unidades Huggies', 'Huggies', 'farmacia', 'cuidado-bebe', 2100.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "paquete 80u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Toallitas húmedas desmaquillantes x20 unidades', NULL, 'farmacia', 'cuidado-bebe', 800.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "paquete 20u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Crema para bebé con óxido de zinc 100g Bübchen', 'Bübchen', 'farmacia', 'cuidado-bebe', 2200.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 100g", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Talco para bebé 200g Johnson', 'Johnson', 'farmacia', 'cuidado-bebe', 1500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 200g", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Shampoo bebé 200ml Johnson', 'Johnson', 'farmacia', 'cuidado-bebe', 1800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Colonia bebé 200ml Johnson', 'Johnson', 'farmacia', 'cuidado-bebe', 2000.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============================================================
-- CATEGORÍA: cuidado-corporal — Cremas, hidratantes, higiene general
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Cremas e hidratantes
('Crema hidratante corporal 400ml Nivea Milk', 'Nivea', 'farmacia', 'cuidado-corporal', 3500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Crema hidratante corporal 400ml Dove', 'Dove', 'farmacia', 'cuidado-corporal', 3800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Crema hidratante manos 75ml Neutrogena', 'Neutrogena', 'farmacia', 'cuidado-corporal', 2200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Vaselina pura 100g', NULL, 'farmacia', 'cuidado-corporal', 800.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad 100g", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Gel aloe vera 250ml', NULL, 'farmacia', 'cuidado-corporal', 2000.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Jabones
('Jabón en barra Dove Original x3 unidades', 'Dove', 'farmacia', 'cuidado-corporal', 1800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "pack 3u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Jabón líquido antibacterial 250ml Protex', 'Protex', 'farmacia', 'cuidado-corporal', 1500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Jabón líquido antibacterial 500ml Protex', 'Protex', 'farmacia', 'cuidado-corporal', 2500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Jabón de glicerina 100g', NULL, 'farmacia', 'cuidado-corporal', 600.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "unidad 100g", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Higiene femenina
('Toallas femeninas x8 compresa normal Always', 'Always', 'farmacia', 'cuidado-corporal', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "paquete 8u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Toallas femeninas x16 compresa normal Always', 'Always', 'farmacia', 'cuidado-corporal', 2100.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "paquete 16u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tampones x8 unidades Tampax regular', 'Tampax', 'farmacia', 'cuidado-corporal', 1400.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "paquete 8u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Protegeslip x20 unidades Always', 'Always', 'farmacia', 'cuidado-corporal', 900.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "paquete 20u", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Accesorios de higiene
('Máquina de afeitar descartable x5 Gillette', 'Gillette', 'farmacia', 'cuidado-corporal', 2200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "pack 5u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Espuma de afeitar 200ml Gillette', 'Gillette', 'farmacia', 'cuidado-corporal', 1800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cotonetes x100 unidades', NULL, 'farmacia', 'cuidado-corporal', 500.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "caja 100u", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;
