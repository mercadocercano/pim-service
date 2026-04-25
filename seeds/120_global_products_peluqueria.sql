-- =============================================================
-- SEED 120: Peluquería y Estética (business_type = 'peluqueria')
-- Research date: 2026-04-25
-- Source: Catálogo market research NEA/Posadas 2026
--   - Distribuidoras profesionales (Roca Distribuciones, La Paleta)
--   - Precios de reventa minorista y mayorista abril 2026
-- Zona: Posadas, Misiones (NEA)
-- Objetivo: >= 100 productos verificados
-- Categorías: tinturas-capilares, cuidado-capilar,
--             tratamientos-capilares, accesorios-belleza
-- =============================================================

-- ============================================================
-- CATEGORÍA: tinturas-capilares — Tinturas, oxidantes, decolorantes
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Koleston (Wella) — referente profesional NEA
('Tintura capilar profesional Koleston 2/0 negro 60g Wella', 'Wella', 'peluqueria', 'tinturas-capilares', 2200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "2/0 negro", "linea": "Koleston", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Koleston 3/0 castaño oscuro 60g Wella', 'Wella', 'peluqueria', 'tinturas-capilares', 2200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "3/0 castaño oscuro", "linea": "Koleston", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Koleston 4/0 castaño medio 60g Wella', 'Wella', 'peluqueria', 'tinturas-capilares', 2200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "4/0 castaño medio", "linea": "Koleston", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Koleston 5/0 castaño claro 60g Wella', 'Wella', 'peluqueria', 'tinturas-capilares', 2200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "5/0 castaño claro", "linea": "Koleston", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Koleston 6/0 rubio oscuro 60g Wella', 'Wella', 'peluqueria', 'tinturas-capilares', 2200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "6/0 rubio oscuro", "linea": "Koleston", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Koleston 7/0 rubio medio 60g Wella', 'Wella', 'peluqueria', 'tinturas-capilares', 2200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "7/0 rubio medio", "linea": "Koleston", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Koleston 8/0 rubio claro 60g Wella', 'Wella', 'peluqueria', 'tinturas-capilares', 2200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "8/0 rubio claro", "linea": "Koleston", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Koleston 9/0 rubio muy claro 60g Wella', 'Wella', 'peluqueria', 'tinturas-capilares', 2200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "9/0 rubio muy claro", "linea": "Koleston", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Koleston 4/65 castaño caoba 60g Wella', 'Wella', 'peluqueria', 'tinturas-capilares', 2200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "4/65 castaño caoba", "linea": "Koleston", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Koleston 6/45 rubio oscuro cobrizo 60g Wella', 'Wella', 'peluqueria', 'tinturas-capilares', 2200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "6/45 rubio oscuro cobrizo", "linea": "Koleston", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Igora Royal (Schwarzkopf)
('Tintura capilar profesional Igora Royal 3-0 negro oscuro 60g Schwarzkopf', 'Schwarzkopf', 'peluqueria', 'tinturas-capilares', 2400.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "3-0 negro oscuro", "linea": "Igora Royal", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Igora Royal 4-0 castaño medio 60g Schwarzkopf', 'Schwarzkopf', 'peluqueria', 'tinturas-capilares', 2400.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "4-0 castaño medio", "linea": "Igora Royal", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Igora Royal 5-0 castaño claro 60g Schwarzkopf', 'Schwarzkopf', 'peluqueria', 'tinturas-capilares', 2400.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "5-0 castaño claro", "linea": "Igora Royal", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Igora Royal 6-0 rubio oscuro 60g Schwarzkopf', 'Schwarzkopf', 'peluqueria', 'tinturas-capilares', 2400.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "6-0 rubio oscuro", "linea": "Igora Royal", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Igora Royal 7-0 rubio medio 60g Schwarzkopf', 'Schwarzkopf', 'peluqueria', 'tinturas-capilares', 2400.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "7-0 rubio medio", "linea": "Igora Royal", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar profesional Igora Royal 5-65 castaño cobrizo 60g Schwarzkopf', 'Schwarzkopf', 'peluqueria', 'tinturas-capilares', 2400.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 60g", "tono": "5-65 castaño cobrizo", "linea": "Igora Royal", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Nutrisse (Garnier) — accesible, alta rotación
('Tintura capilar Nutrisse 10 negro azulado Garnier', 'Garnier', 'peluqueria', 'tinturas-capilares', 1800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "tono": "10 negro azulado", "linea": "Nutrisse", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar Nutrisse 30 castaño oscuro Garnier', 'Garnier', 'peluqueria', 'tinturas-capilares', 1800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "tono": "30 castaño oscuro", "linea": "Nutrisse", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar Nutrisse 50 castaño claro Garnier', 'Garnier', 'peluqueria', 'tinturas-capilares', 1800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "tono": "50 castaño claro", "linea": "Nutrisse", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar Nutrisse 60 rubio oscuro Garnier', 'Garnier', 'peluqueria', 'tinturas-capilares', 1800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "tono": "60 rubio oscuro", "linea": "Nutrisse", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar Nutrisse 70 rubio Garnier', 'Garnier', 'peluqueria', 'tinturas-capilares', 1800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "tono": "70 rubio", "linea": "Nutrisse", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar Nutrisse 43 caoba Garnier', 'Garnier', 'peluqueria', 'tinturas-capilares', 1800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "tono": "43 caoba", "linea": "Nutrisse", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tintura capilar Nutrisse 66 rojo intenso Garnier', 'Garnier', 'peluqueria', 'tinturas-capilares', 1800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "tono": "66 rojo intenso", "linea": "Nutrisse", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Oxidantes
('Oxidante 10 volúmenes 90ml Wella Welloxon', 'Wella', 'peluqueria', 'tinturas-capilares', 800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 90ml", "volumen": "10vol", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Oxidante 20 volúmenes 90ml Wella Welloxon', 'Wella', 'peluqueria', 'tinturas-capilares', 800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 90ml", "volumen": "20vol", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Oxidante 30 volúmenes 90ml Wella Welloxon', 'Wella', 'peluqueria', 'tinturas-capilares', 800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 90ml", "volumen": "30vol", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Oxidante 40 volúmenes 90ml Wella Welloxon', 'Wella', 'peluqueria', 'tinturas-capilares', 800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 90ml", "volumen": "40vol", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Oxidante 20 volúmenes 1L Schwarzkopf Igora', 'Schwarzkopf', 'peluqueria', 'tinturas-capilares', 4500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 1L", "volumen": "20vol", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Oxidante 30 volúmenes 1L Schwarzkopf Igora', 'Schwarzkopf', 'peluqueria', 'tinturas-capilares', 4500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 1L", "volumen": "30vol", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Decolorante
('Polvo decolorante 500g Schwarzkopf Blondme', 'Schwarzkopf', 'peluqueria', 'tinturas-capilares', 5500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad 500g", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Polvo decolorante 50g sobre Wella Blonderma', 'Wella', 'peluqueria', 'tinturas-capilares', 1200.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "sobre 50g", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Polvo decolorante 500g Bleach L Oréal', 'L''Oréal', 'peluqueria', 'tinturas-capilares', 5000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad 500g", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============================================================
-- CATEGORÍA: cuidado-capilar — Shampoos y acondicionadores profesionales
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Shampoos profesionales
('Shampoo profesional Wella Invigo Color Brilliance 1L', 'Wella', 'peluqueria', 'cuidado-capilar', 8500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 1L", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Shampoo profesional Wella Invigo Moisture Moi 1L', 'Wella', 'peluqueria', 'cuidado-capilar', 8500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 1L", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Shampoo profesional Wella Invigo Nutri-Enrich 1L', 'Wella', 'peluqueria', 'cuidado-capilar', 8500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 1L", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Shampoo profesional L Oréal Série Expert Vitamino 1.5L', 'L''Oréal', 'peluqueria', 'cuidado-capilar', 11000.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 1.5L", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Shampoo profesional L Oréal Série Expert Absolut Repair 1.5L', 'L''Oréal', 'peluqueria', 'cuidado-capilar', 11000.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 1.5L", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Shampoo profesional Joico Color Endure 1L', 'Joico', 'peluqueria', 'cuidado-capilar', 9000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad 1L", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Shampoo profesional Schwarzkopf BC Color Freeze 1L', 'Schwarzkopf', 'peluqueria', 'cuidado-capilar', 7500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad 1L", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Shampoo matizador violeta sin amarillo 300ml', NULL, 'peluqueria', 'cuidado-capilar', 3500.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad 300ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Acondicionadores profesionales
('Acondicionador profesional Wella Invigo Color Brilliance 1L', 'Wella', 'peluqueria', 'cuidado-capilar', 8500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 1L", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Acondicionador profesional Wella Invigo Moisture Moi 1L', 'Wella', 'peluqueria', 'cuidado-capilar', 8500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 1L", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Acondicionador profesional L Oréal Série Expert Vitamino 1.5L', 'L''Oréal', 'peluqueria', 'cuidado-capilar', 11000.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad 1.5L", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Acondicionador profesional Joico Color Endure 1L', 'Joico', 'peluqueria', 'cuidado-capilar', 9000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad 1L", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Geles, ceras y fijadores
('Cera modeladora mate 100g Schwarzkopf OSiS', 'Schwarzkopf', 'peluqueria', 'cuidado-capilar', 4500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 100g", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cera modeladora brillo 100g Wella Shockwaves', 'Wella', 'peluqueria', 'cuidado-capilar', 4000.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 100g", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Gel fijador fuerte 250ml Wella Shockwaves', 'Wella', 'peluqueria', 'cuidado-capilar', 3500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 250ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Gel fijador mega fuerte 300ml Schwarzkopf Taft', 'Schwarzkopf', 'peluqueria', 'cuidado-capilar', 3200.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 300ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Laca capilar spray fuerte 400ml Wella Shockwaves', 'Wella', 'peluqueria', 'cuidado-capilar', 3800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 400ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Laca capilar spray extra fuerte 400ml Schwarzkopf Taft', 'Schwarzkopf', 'peluqueria', 'cuidado-capilar', 3500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 400ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tónico capilar anticaída 200ml Yaneke', 'Yaneke', 'peluqueria', 'cuidado-capilar', 3000.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad 200ml", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============================================================
-- CATEGORÍA: tratamientos-capilares — Máscaras, keratinas, ampollas
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Mascarillas y tratamientos intensivos
('Mascarilla capilar nutrición 200ml Wella Invigo', 'Wella', 'peluqueria', 'tratamientos-capilares', 6500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad 200ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Mascarilla capilar nutrición 500ml Wella Invigo', 'Wella', 'peluqueria', 'tratamientos-capilares', 12000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad 500ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Mascarilla capilar reparación 200ml L Oréal Absolut Repair', 'L''Oréal', 'peluqueria', 'tratamientos-capilares', 7000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad 200ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Mascarilla capilar reparación 500ml L Oréal Absolut Repair', 'L''Oréal', 'peluqueria', 'tratamientos-capilares', 14000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad 500ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Mascarilla capilar hidratación 300ml Schwarzkopf BC', 'Schwarzkopf', 'peluqueria', 'tratamientos-capilares', 6000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad 300ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Mascarilla capilar hidratación 200ml Joico', 'Joico', 'peluqueria', 'tratamientos-capilares', 7500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad 200ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Keratina
('Keratina alisante 300ml sin formol Brazilian', NULL, 'peluqueria', 'tratamientos-capilares', 12000.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad 300ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Keratina alisante 1L sin formol profesional', NULL, 'peluqueria', 'tratamientos-capilares', 35000.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad 1L", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Progresiva capilar 300ml sin formol', NULL, 'peluqueria', 'tratamientos-capilares', 10000.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad 300ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Ampollas de tratamiento
('Ampolla capilar hidratante x12 unidades Joico', 'Joico', 'peluqueria', 'tratamientos-capilares', 8000.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "caja 12u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Ampolla capilar reparadora x12 unidades Wella', 'Wella', 'peluqueria', 'tratamientos-capilares', 7500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "caja 12u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Ampolla anti-caída x12 unidades genérico', NULL, 'peluqueria', 'tratamientos-capilares', 3500.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "caja 12u", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Protectores térmicos
('Protector térmico spray 250ml Wella EIMI', 'Wella', 'peluqueria', 'tratamientos-capilares', 5500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 250ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Protector térmico spray 200ml Schwarzkopf OSiS', 'Schwarzkopf', 'peluqueria', 'tratamientos-capilares', 5000.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 200ml", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Sérum de brillo capilar 50ml L Oréal', 'L''Oréal', 'peluqueria', 'tratamientos-capilares', 4500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad 50ml", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============================================================
-- CATEGORÍA: accesorios-belleza — Herramientas y accesorios profesionales
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Tijeras profesionales
('Tijera profesional pelo 5.5 pulgadas Jaguar', 'Jaguar', 'peluqueria', 'accesorios-belleza', 18000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "largo": "5.5 pulgadas", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tijera profesional pelo 6 pulgadas Jaguar', 'Jaguar', 'peluqueria', 'accesorios-belleza', 20000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "largo": "6 pulgadas", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tijera profesional pelo 6.5 pulgadas Jaguar', 'Jaguar', 'peluqueria', 'accesorios-belleza', 22000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "largo": "6.5 pulgadas", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tijera entresacar fileteadora 6 pulgadas', NULL, 'peluqueria', 'accesorios-belleza', 12000.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad", "largo": "6 pulgadas", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tijera peluquería acero inox 6 pulgadas económica', NULL, 'peluqueria', 'accesorios-belleza', 4500.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "unidad", "largo": "6 pulgadas", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Peines y cepillos
('Peine de corte profesional cola fina', NULL, 'peluqueria', 'accesorios-belleza', 500.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Peine de cola largo profesional', NULL, 'peluqueria', 'accesorios-belleza', 600.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Peine de dientes anchos antistático', NULL, 'peluqueria', 'accesorios-belleza', 400.00, 'catalog-researcher', 0.80, 52, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cepillo paleta grande tipo Denman', NULL, 'peluqueria', 'accesorios-belleza', 2500.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cepillo redondo térmico 25mm cerda jabalí', NULL, 'peluqueria', 'accesorios-belleza', 3000.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "diametro": "25mm", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cepillo redondo térmico 40mm cerda mixta', NULL, 'peluqueria', 'accesorios-belleza', 3500.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "diametro": "40mm", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Secadores y planchas
('Secador de cabello profesional 2000W Gamma Più', 'Gamma Più', 'peluqueria', 'accesorios-belleza', 35000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "potencia": "2000W", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Secador de cabello profesional 2200W Parlux', 'Parlux', 'peluqueria', 'accesorios-belleza', 65000.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "potencia": "2200W", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Secador doméstico 1800W Philips', 'Philips', 'peluqueria', 'accesorios-belleza', 18000.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "potencia": "1800W", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Plancha alisadora cerámica 230°C Gamma Più', 'Gamma Più', 'peluqueria', 'accesorios-belleza', 22000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "temperatura_max": "230°C", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Plancha alisadora cerámica 230°C Remington', 'Remington', 'peluqueria', 'accesorios-belleza', 20000.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "temperatura_max": "230°C", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Rizador barril 25mm 200°C Gamma Più', 'Gamma Più', 'peluqueria', 'accesorios-belleza', 18000.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "barril": "25mm", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Accesorios varios
('Gorros de tintura descartables x50 unidades', NULL, 'peluqueria', 'accesorios-belleza', 800.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "paquete 50u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Guantes de vinilo talle M caja x100 unidades', NULL, 'peluqueria', 'accesorios-belleza', 2500.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "caja 100u", "talle": "M", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Guantes de vinilo talle G caja x100 unidades', NULL, 'peluqueria', 'accesorios-belleza', 2500.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "caja 100u", "talle": "G", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Papel aluminio para mechones x500m rollo', NULL, 'peluqueria', 'accesorios-belleza', 5500.00, 'catalog-researcher', 0.75, 55, TRUE, TRUE, '{"unit": "rollo 500m", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pincel para tintura plano N4', NULL, 'peluqueria', 'accesorios-belleza', 400.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Bowl para mezcla tintura plástico', NULL, 'peluqueria', 'accesorios-belleza', 350.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cuello descartable x100 unidades tira', NULL, 'peluqueria', 'accesorios-belleza', 900.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "tira 100u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Capa de corte impermeable negra', NULL, 'peluqueria', 'accesorios-belleza', 3500.00, 'catalog-researcher', 0.75, 55, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Clips para mechas ondas x12 unidades', NULL, 'peluqueria', 'accesorios-belleza', 600.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "pack 12u", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;
