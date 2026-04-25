-- =============================================================
-- SEED 118: Electricidad (business_type = 'electricidad')
-- Research date: 2026-04-25
-- Source: Catálogo market research NEA/Posadas 2026
--   - Distribuidoras eléctricas Posadas (Electro NEA, Casanova)
--   - Precios mayoristas y minoristas relevados abril 2026
-- Zona: Posadas, Misiones (NEA)
-- Objetivo: >= 150 productos verificados
-- Categorías: cables-electricidad, proteccion-electrica,
--             iluminacion-led, herramientas-electricas
-- Nota: precios de cables son por metro o por rollo según variante
-- =============================================================

-- ============================================================
-- CATEGORÍA: cables-electricidad — Cables por sección y presentación
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Cable unipolar 1.5mm² (precio por metro y rollo)
('Cable unipolar 1.5mm² blanco x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 320.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "1.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable unipolar 1.5mm² negro x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 320.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "1.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable unipolar 1.5mm² verde x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 320.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "1.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable unipolar 1.5mm² rojo x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 320.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "1.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable unipolar 1.5mm² rollo x100m Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 28000.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "rollo 100m", "seccion": "1.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable unipolar 1.5mm² rollo x100m Centelsa', 'Centelsa', 'electricidad', 'cables-electricidad', 26000.00, 'catalog-researcher', 0.80, 65, TRUE, TRUE, '{"unit": "rollo 100m", "seccion": "1.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Cable unipolar 2.5mm²
('Cable unipolar 2.5mm² blanco x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 520.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "2.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable unipolar 2.5mm² negro x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 520.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "2.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable unipolar 2.5mm² verde x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 520.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "2.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable unipolar 2.5mm² rollo x100m Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 46000.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "rollo 100m", "seccion": "2.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable unipolar 2.5mm² rollo x100m Centelsa', 'Centelsa', 'electricidad', 'cables-electricidad', 42000.00, 'catalog-researcher', 0.80, 65, TRUE, TRUE, '{"unit": "rollo 100m", "seccion": "2.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Cable unipolar 4mm²
('Cable unipolar 4mm² negro x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 820.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "4mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable unipolar 4mm² blanco x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 820.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "4mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable unipolar 4mm² rollo x100m Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 72000.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "rollo 100m", "seccion": "4mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Cable unipolar 6mm²
('Cable unipolar 6mm² negro x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "6mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable unipolar 6mm² rollo x50m Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 55000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "rollo 50m", "seccion": "6mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Cable unipolar 10mm²
('Cable unipolar 10mm² negro x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 1900.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "metro", "seccion": "10mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Cables bipolares y tripolares
('Cable bipolar 2x1.5mm² x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 650.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "2x1.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable bipolar 2x2.5mm² x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 1050.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "2x2.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable tripolar 3x2.5mm² x metro Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 1400.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "metro", "seccion": "3x2.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable bipolar 2x1.5mm² rollo x100m Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 58000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "rollo 100m", "seccion": "2x1.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cable bipolar 2x2.5mm² rollo x100m Prysmian', 'Prysmian', 'electricidad', 'cables-electricidad', 92000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "rollo 100m", "seccion": "2x2.5mm2", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Caño corrugado
('Caño corrugado PVC 20mm x metro', NULL, 'electricidad', 'cables-electricidad', 180.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "metro", "diametro": "20mm", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Caño corrugado PVC 25mm x metro', NULL, 'electricidad', 'cables-electricidad', 220.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "metro", "diametro": "25mm", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Caño corrugado PVC 32mm x metro', NULL, 'electricidad', 'cables-electricidad', 280.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "metro", "diametro": "32mm", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Caño corrugado PVC 20mm rollo x50m', NULL, 'electricidad', 'cables-electricidad', 7500.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "rollo 50m", "diametro": "20mm", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Caño corrugado PVC 25mm rollo x50m', NULL, 'electricidad', 'cables-electricidad', 9500.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "rollo 50m", "diametro": "25mm", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Caño rígido PVC 25mm barra x3m', NULL, 'electricidad', 'cables-electricidad', 1200.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "barra 3m", "diametro": "25mm", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Caño rígido PVC 32mm barra x3m', NULL, 'electricidad', 'cables-electricidad', 1600.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "barra 3m", "diametro": "32mm", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Prensacables y cajas
('Prensacable PG7 nylon negro x10 unidades', NULL, 'electricidad', 'cables-electricidad', 500.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "pack 10u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Prensacable PG9 nylon negro x10 unidades', NULL, 'electricidad', 'cables-electricidad', 550.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "pack 10u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Prensacable PG11 nylon negro x10 unidades', NULL, 'electricidad', 'cables-electricidad', 600.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "pack 10u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Caja de paso plástica 10x10cm tapa ciega', NULL, 'electricidad', 'cables-electricidad', 350.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Caja octogonal PVC embutir', NULL, 'electricidad', 'cables-electricidad', 180.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Caja rectangular PVC embutir tomacorriente', NULL, 'electricidad', 'cables-electricidad', 150.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cinta aisladora negra 20m 3M', '3M', 'electricidad', 'cables-electricidad', 400.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "rollo 20m", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cinta aisladora negra 10m genérico', NULL, 'electricidad', 'cables-electricidad', 200.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "rollo 10m", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cintillo plástico 100mm x2.5mm x100 unidades', NULL, 'electricidad', 'cables-electricidad', 350.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "paquete 100u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cintillo plástico 200mm x3.6mm x100 unidades', NULL, 'electricidad', 'cables-electricidad', 500.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "paquete 100u", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============================================================
-- CATEGORÍA: proteccion-electrica — Tableros, termomagnéticas, diferenciales
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Llaves termomagnéticas
('Llave termomagnética 1x10A Schneider Electric', 'Schneider Electric', 'electricidad', 'proteccion-electrica', 4500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "amperaje": "10A", "polos": 1, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave termomagnética 1x16A Schneider Electric', 'Schneider Electric', 'electricidad', 'proteccion-electrica', 4700.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "amperaje": "16A", "polos": 1, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave termomagnética 1x20A Schneider Electric', 'Schneider Electric', 'electricidad', 'proteccion-electrica', 4900.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "amperaje": "20A", "polos": 1, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave termomagnética 1x25A Schneider Electric', 'Schneider Electric', 'electricidad', 'proteccion-electrica', 5200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "amperaje": "25A", "polos": 1, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave termomagnética 1x32A Schneider Electric', 'Schneider Electric', 'electricidad', 'proteccion-electrica', 5800.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "amperaje": "32A", "polos": 1, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave termomagnética 2x20A Schneider Electric', 'Schneider Electric', 'electricidad', 'proteccion-electrica', 9500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "amperaje": "20A", "polos": 2, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave termomagnética 2x25A Schneider Electric', 'Schneider Electric', 'electricidad', 'proteccion-electrica', 10200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "amperaje": "25A", "polos": 2, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave termomagnética 2x32A Schneider Electric', 'Schneider Electric', 'electricidad', 'proteccion-electrica', 11500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "amperaje": "32A", "polos": 2, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave termomagnética 2x40A Schneider Electric', 'Schneider Electric', 'electricidad', 'proteccion-electrica', 13000.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "amperaje": "40A", "polos": 2, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave termomagnética 1x10A BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 4200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "amperaje": "10A", "polos": 1, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave termomagnética 1x20A BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 4600.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "amperaje": "20A", "polos": 1, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave termomagnética 2x20A BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 9000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "amperaje": "20A", "polos": 2, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave termomagnética 2x32A BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 11000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "amperaje": "32A", "polos": 2, "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Interruptores diferenciales
('Interruptor diferencial 2x25A 30mA Schneider Electric', 'Schneider Electric', 'electricidad', 'proteccion-electrica', 18000.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "amperaje": "25A", "sensibilidad": "30mA", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Interruptor diferencial 2x32A 30mA Schneider Electric', 'Schneider Electric', 'electricidad', 'proteccion-electrica', 19500.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "amperaje": "32A", "sensibilidad": "30mA", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Interruptor diferencial 2x40A 30mA Schneider Electric', 'Schneider Electric', 'electricidad', 'proteccion-electrica', 22000.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "amperaje": "40A", "sensibilidad": "30mA", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Interruptor diferencial 2x25A 30mA BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 17000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "amperaje": "25A", "sensibilidad": "30mA", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Interruptor diferencial 2x40A 30mA BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 21000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "amperaje": "40A", "sensibilidad": "30mA", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Tableros
('Tablero monofásico 4 módulos embutir Conextube', 'Conextube', 'electricidad', 'proteccion-electrica', 4500.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad", "modulos": 4, "tipo": "monofasico", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tablero monofásico 6 módulos embutir Conextube', 'Conextube', 'electricidad', 'proteccion-electrica', 5500.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad", "modulos": 6, "tipo": "monofasico", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tablero monofásico 8 módulos embutir Conextube', 'Conextube', 'electricidad', 'proteccion-electrica', 6500.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad", "modulos": 8, "tipo": "monofasico", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tablero monofásico 12 módulos embutir Conextube', 'Conextube', 'electricidad', 'proteccion-electrica', 8500.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad", "modulos": 12, "tipo": "monofasico", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tablero monofásico 4 módulos superficie Ferretería', NULL, 'electricidad', 'proteccion-electrica', 4000.00, 'catalog-researcher', 0.75, 55, TRUE, TRUE, '{"unit": "unidad", "modulos": 4, "tipo": "monofasico", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tablero trifásico 12 módulos embutir Conextube', 'Conextube', 'electricidad', 'proteccion-electrica', 12000.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad", "modulos": 12, "tipo": "trifasico", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Tomacorrientes y llaves de luz
('Tomacorriente doble 2P+T Ornamental blanco BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 1800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "tipo": "doble", "sistema": "Ornamental", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tomacorriente simple 2P+T Ornamental blanco BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "tipo": "simple", "sistema": "Ornamental", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tomacorriente doble 2P+T Magic blanco BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 2200.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "tipo": "doble", "sistema": "Magic", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tomacorriente triple 2P+T Cambre blanco', 'Cambre', 'electricidad', 'proteccion-electrica', 2500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "tipo": "triple", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave de luz simple Ornamental BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 1000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "sistema": "Ornamental", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave de luz doble Ornamental BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 1500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "sistema": "Ornamental", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Llave llave simple Magic BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 1300.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "sistema": "Magic", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Placa simple ciega Ornamental BTicino', 'BTicino', 'electricidad', 'proteccion-electrica', 600.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Caja de 2x4 PVC embutir para mecanismos', NULL, 'electricidad', 'proteccion-electrica', 200.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Enchufes y fichas industriales
('Enchufe industrial 2P+T 16A 220V Conextube', 'Conextube', 'electricidad', 'proteccion-electrica', 1200.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "amperaje": "16A", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Enchufe industrial 2P+T 32A 220V Conextube', 'Conextube', 'electricidad', 'proteccion-electrica', 2500.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "amperaje": "32A", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Ficha macho con tierra 16A Conextube', 'Conextube', 'electricidad', 'proteccion-electrica', 800.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Ficha hembra con tierra 16A Conextube', 'Conextube', 'electricidad', 'proteccion-electrica', 900.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Prolongador 3 tomas con tierra 3m', NULL, 'electricidad', 'proteccion-electrica', 2200.00, 'catalog-researcher', 0.75, 55, TRUE, TRUE, '{"unit": "unidad", "largo": "3m", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Prolongador 3 tomas con tierra 5m', NULL, 'electricidad', 'proteccion-electrica', 3000.00, 'catalog-researcher', 0.75, 55, TRUE, TRUE, '{"unit": "unidad", "largo": "5m", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Zapatilla múltiple 6 tomas con interruptor', NULL, 'electricidad', 'proteccion-electrica', 3500.00, 'catalog-researcher', 0.75, 55, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============================================================
-- CATEGORÍA: iluminacion-led — Focos, tiras, tubos LED
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Focos LED E27
('Foco LED 9W luz cálida E27 Osram', 'Osram', 'electricidad', 'iluminacion-led', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "potencia": "9W", "temperatura": "3000K", "base": "E27", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Foco LED 9W luz fría E27 Osram', 'Osram', 'electricidad', 'iluminacion-led', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "potencia": "9W", "temperatura": "6500K", "base": "E27", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Foco LED 13W luz cálida E27 Osram', 'Osram', 'electricidad', 'iluminacion-led', 1500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "potencia": "13W", "temperatura": "3000K", "base": "E27", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Foco LED 13W luz fría E27 Osram', 'Osram', 'electricidad', 'iluminacion-led', 1500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "potencia": "13W", "temperatura": "6500K", "base": "E27", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Foco LED 18W luz fría E27 Osram', 'Osram', 'electricidad', 'iluminacion-led', 1900.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "potencia": "18W", "temperatura": "6500K", "base": "E27", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Foco LED 9W luz cálida E27 Philips', 'Philips', 'electricidad', 'iluminacion-led', 1300.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "potencia": "9W", "temperatura": "3000K", "base": "E27", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Foco LED 13W luz fría E27 Philips', 'Philips', 'electricidad', 'iluminacion-led', 1600.00, 'catalog-researcher', 0.85, 72, TRUE, TRUE, '{"unit": "unidad", "potencia": "13W", "temperatura": "6500K", "base": "E27", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Foco LED 9W luz fría E27 genérico', NULL, 'electricidad', 'iluminacion-led', 600.00, 'catalog-researcher', 0.70, 45, TRUE, TRUE, '{"unit": "unidad", "potencia": "9W", "temperatura": "6500K", "base": "E27", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Foco LED 12W luz fría E27 genérico', NULL, 'electricidad', 'iluminacion-led', 750.00, 'catalog-researcher', 0.70, 45, TRUE, TRUE, '{"unit": "unidad", "potencia": "12W", "temperatura": "6500K", "base": "E27", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Focos dicroicos y especiales
('Foco LED dicroico 7W GU10 luz fría Osram', 'Osram', 'electricidad', 'iluminacion-led', 1400.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "potencia": "7W", "base": "GU10", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Foco LED dicroico 5W GU5.3 luz fría Osram', 'Osram', 'electricidad', 'iluminacion-led', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "potencia": "5W", "base": "GU5.3", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Foco LED vela 6W E14 luz cálida Osram', 'Osram', 'electricidad', 'iluminacion-led', 1100.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "potencia": "6W", "base": "E14", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Tubos LED T8
('Tubo LED T8 18W 120cm luz fría Osram', 'Osram', 'electricidad', 'iluminacion-led', 2200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "potencia": "18W", "largo": "120cm", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tubo LED T8 9W 60cm luz fría Osram', 'Osram', 'electricidad', 'iluminacion-led', 1500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "potencia": "9W", "largo": "60cm", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tubo LED T8 18W 120cm luz fría genérico', NULL, 'electricidad', 'iluminacion-led', 900.00, 'catalog-researcher', 0.70, 45, TRUE, TRUE, '{"unit": "unidad", "potencia": "18W", "largo": "120cm", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Tiras LED
('Tira LED 5050 RGB 5m con control 12V', NULL, 'electricidad', 'iluminacion-led', 4500.00, 'catalog-researcher', 0.75, 55, TRUE, TRUE, '{"unit": "rollo 5m", "tipo": "RGB", "voltaje": "12V", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tira LED 3528 blanco frío 5m autoadhesiva', NULL, 'electricidad', 'iluminacion-led', 2800.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "rollo 5m", "color": "blanco-frio", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tira LED 3528 blanco cálido 5m autoadhesiva', NULL, 'electricidad', 'iluminacion-led', 2800.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "rollo 5m", "color": "blanco-calido", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Fuente 12V 2A para tira LED', NULL, 'electricidad', 'iluminacion-led', 1500.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Fuente 12V 5A para tira LED', NULL, 'electricidad', 'iluminacion-led', 2500.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Reflectores LED
('Reflector LED 10W luz fría exterior', NULL, 'electricidad', 'iluminacion-led', 2200.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "unidad", "potencia": "10W", "uso": "exterior", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Reflector LED 20W luz fría exterior', NULL, 'electricidad', 'iluminacion-led', 3200.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "unidad", "potencia": "20W", "uso": "exterior", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Reflector LED 30W luz fría exterior', NULL, 'electricidad', 'iluminacion-led', 4500.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "unidad", "potencia": "30W", "uso": "exterior", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Reflector LED 50W luz fría exterior', NULL, 'electricidad', 'iluminacion-led', 6500.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "unidad", "potencia": "50W", "uso": "exterior", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Luminarias de emergencia
('Luminaria de emergencia LED 2hs autonomía', NULL, 'electricidad', 'iluminacion-led', 5500.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "unidad", "autonomia": "2hs", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Señal de emergencia Salida LED', NULL, 'electricidad', 'iluminacion-led', 3500.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============================================================
-- CATEGORÍA: herramientas-electricas — Herramientas para electricista
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

('Tester digital digital Multímetro Uni-T UT33B', 'Uni-T', 'electricidad', 'herramientas-electricas', 8000.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Tester de fase y neutro de enchufe', NULL, 'electricidad', 'herramientas-electricas', 1200.00, 'catalog-researcher', 0.75, 55, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pinza amperométrica digital Uni-T', 'Uni-T', 'electricidad', 'herramientas-electricas', 12000.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Destornillador de prueba 100-500V', NULL, 'electricidad', 'herramientas-electricas', 350.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pelacable automático AWG10-20', NULL, 'electricidad', 'herramientas-electricas', 3500.00, 'catalog-researcher', 0.80, 58, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pelacable manual 1.5-6mm²', NULL, 'electricidad', 'herramientas-electricas', 1500.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Alicate electricista 200mm aislado 1000V Stanley', 'Stanley', 'electricidad', 'herramientas-electricas', 6500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pinza corte lateral 160mm aislada 1000V Stanley', 'Stanley', 'electricidad', 'herramientas-electricas', 5500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Escalera de aluminio 6 escalones', NULL, 'electricidad', 'herramientas-electricas', 18000.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "escalones": 6, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Terminal tipo ojo 6mm² pack x50 unidades', NULL, 'electricidad', 'herramientas-electricas', 1200.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "pack 50u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Terminal tipo punta 1.5-2.5mm² pack x100 unidades', NULL, 'electricidad', 'herramientas-electricas', 800.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "pack 100u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Conector Wago 3 polos pack x5 unidades', 'Wago', 'electricidad', 'herramientas-electricas', 1500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "pack 5u", "polos": 3, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Conector Wago 2 polos pack x5 unidades', 'Wago', 'electricidad', 'herramientas-electricas', 1200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "pack 5u", "polos": 2, "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cinta doble faz 19mm x5m 3M', '3M', 'electricidad', 'herramientas-electricas', 800.00, 'catalog-researcher', 0.85, 62, TRUE, TRUE, '{"unit": "rollo 5m", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Funda termocontráctil 3mm x1m surtida x5', NULL, 'electricidad', 'herramientas-electricas', 600.00, 'catalog-researcher', 0.75, 52, TRUE, TRUE, '{"unit": "pack 5u", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;
