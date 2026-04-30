-- Seed 082: Productos globales — Ferretería: Fijaciones y Sujeción (~110 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: catálogos ferreterías NEA (Posadas) + Fischer Argentina + distribuidores 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: tornillos-clavos, fijaciones

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- TORNILLOS PARA MADERA — medidas habituales de ferretería NEA
-- ============================================================
('Tornillo madera 3x20mm galvanizado x100u',    NULL,  'tornillos-clavos',   850.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo madera 3.5x25mm galvanizado x100u',  NULL,  'tornillos-clavos',   980.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo madera 3.5x35mm galvanizado x100u',  NULL,  'tornillos-clavos',  1100.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo madera 4x40mm galvanizado x100u',    NULL,  'tornillos-clavos',  1250.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo madera 4x50mm galvanizado x100u',    NULL,  'tornillos-clavos',  1400.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo madera 4.5x60mm galvanizado x50u',   NULL,  'tornillos-clavos',  1100.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo madera 5x80mm galvanizado x50u',     NULL,  'tornillos-clavos',  1350.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo madera 5x100mm galvanizado x25u',    NULL,  'tornillos-clavos',   980.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo madera 6x120mm galvanizado x10u',    NULL,  'tornillos-clavos',   750.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo madera autoperforante 4x25mm x100u', NULL,  'tornillos-clavos',  1100.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),

-- ============================================================
-- TORNILLOS PARA CHAPA — usados en construcción liviana
-- ============================================================
('Tornillo chapa 4.2x13mm autoperforante x100u',NULL,  'tornillos-clavos',   950.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo chapa 4.2x19mm autoperforante x100u',NULL,  'tornillos-clavos',  1050.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo chapa 4.8x25mm autoperforante x100u',NULL,  'tornillos-clavos',  1150.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo chapa 4.8x38mm autoperforante x100u',NULL,  'tornillos-clavos',  1350.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo chapa hexagonal 5.5x25mm x100u',     NULL,  'tornillos-clavos',  1200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),

-- ============================================================
-- TORNILLOS DE MÁQUINA (métricos) — bulonería básica
-- ============================================================
('Bulón M6x30mm c/tuerca galvanizado x20u',     NULL,  'tornillos-clavos',   680.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Bulón M8x40mm c/tuerca galvanizado x20u',     NULL,  'tornillos-clavos',   850.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Bulón M8x60mm c/tuerca galvanizado x10u',     NULL,  'tornillos-clavos',   720.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Bulón M10x50mm c/tuerca galvanizado x10u',    NULL,  'tornillos-clavos',   850.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Bulón M10x100mm c/tuerca galvanizado x5u',    NULL,  'tornillos-clavos',   680.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Bulón carruaje M8x80mm galvanizado x10u',     NULL,  'tornillos-clavos',   780.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),

-- ============================================================
-- TUERCAS Y ARANDELAS
-- ============================================================
('Tuerca hexagonal M6 galvanizada x50u',        NULL,  'tornillos-clavos',   480.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tuerca hexagonal M8 galvanizada x50u',        NULL,  'tornillos-clavos',   620.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tuerca hexagonal M10 galvanizada x25u',       NULL,  'tornillos-clavos',   520.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Arandela plana M6 x50u',                      NULL,  'tornillos-clavos',   320.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Arandela plana M8 x50u',                      NULL,  'tornillos-clavos',   420.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Arandela presión M8 x50u',                    NULL,  'tornillos-clavos',   380.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Arandela plana M10 x25u',                     NULL,  'tornillos-clavos',   350.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),

-- ============================================================
-- CLAVOS — varios tamaños y tipos
-- ============================================================
('Clavo de obra 1.5" (40x1.8mm) x1kg',         NULL,  'tornillos-clavos',  1800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "kg",   "sku_prefix": "FERRET"}'),
('Clavo de obra 2" (50x2.0mm) x1kg',            NULL,  'tornillos-clavos',  1900.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "kg",   "sku_prefix": "FERRET"}'),
('Clavo de obra 2.5" (65x2.5mm) x1kg',          NULL,  'tornillos-clavos',  2000.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "kg",   "sku_prefix": "FERRET"}'),
('Clavo de obra 3" (75x3.0mm) x1kg',            NULL,  'tornillos-clavos',  2100.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "kg",   "sku_prefix": "FERRET"}'),
('Clavo de obra 4" (100x4.0mm) x1kg',           NULL,  'tornillos-clavos',  2200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "kg",   "sku_prefix": "FERRET"}'),
('Clavo sin cabeza 1" (25x1.4mm) x100u',        NULL,  'tornillos-clavos',   480.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Clavo ardox 2" para block x1kg',              NULL,  'tornillos-clavos',  2400.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "kg",   "sku_prefix": "FERRET"}'),
('Clavo para amure 1/4" x3/4" x50u',            NULL,  'tornillos-clavos',   650.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Grampa para cerco 1" x1kg',                   NULL,  'tornillos-clavos',  2200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "kg",   "sku_prefix": "FERRET"}'),

-- ============================================================
-- TARUGOS — Fischer y genéricos, múltiples medidas
-- ============================================================
('Tarugo Fischer S6 x100u',                     'Fischer',  'fijaciones',    4500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tarugo Fischer S8 x100u',                     'Fischer',  'fijaciones',    5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tarugo Fischer S10 x50u',                     'Fischer',  'fijaciones',    4200.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tarugo Fischer SX6 c/tornillo x100u',         'Fischer',  'fijaciones',    6800.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tarugo Fischer SX8 c/tornillo x50u',          'Fischer',  'fijaciones',    5200.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tarugo Fischer FUR 10 para hormigón x10u',    'Fischer',  'fijaciones',    5800.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tarugo plástico 6mm x50u genérico',           NULL,       'fijaciones',    1200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tarugo plástico 8mm x50u genérico',           NULL,       'fijaciones',    1500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tarugo plástico 10mm x25u genérico',          NULL,       'fijaciones',    1200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tarugo de expansión metálico 8mm x10u',       NULL,       'fijaciones',    2400.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tarugo de expansión metálico 10mm x10u',      NULL,       'fijaciones',    2800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Taco mariposa 5mm x20u genérico',             NULL,       'fijaciones',    1800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),

-- ============================================================
-- ESCUADRAS Y ÁNGULOS METÁLICOS
-- ============================================================
('Escuadra metálica 40x40x2mm x5u',             NULL,  'fijaciones',         950.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Escuadra metálica 60x60x3mm x5u',             NULL,  'fijaciones',        1350.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Escuadra metálica 80x80x3mm x5u',             NULL,  'fijaciones',        1650.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Placa de unión para madera 80x60mm x5u',      NULL,  'fijaciones',         850.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Sujetaviga metálica 48mm x2u',                NULL,  'fijaciones',         780.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),

-- ============================================================
-- ALAMBRE Y TENSORES
-- ============================================================
('Alambre de hierro recocido N°17 x1kg',        NULL,  'fijaciones',        2200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "kg",   "sku_prefix": "FERRET"}'),
('Alambre de hierro recocido N°18 x1kg',        NULL,  'fijaciones',        2400.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "kg",   "sku_prefix": "FERRET"}'),
('Alambre de fardo x10kg',                      NULL,  'fijaciones',       22000.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "rollo","sku_prefix": "FERRET"}'),
('Alambre galvanizado cerco N°16 x5kg',         NULL,  'fijaciones',       12000.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "rollo","sku_prefix": "FERRET"}'),
('Tensor de alambre largo galvanizado x5u',     NULL,  'fijaciones',        1800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tensor de alambre corto galvanizado x5u',     NULL,  'fijaciones',        1400.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),

-- ============================================================
-- ABRAZADERAS Y GRAPAS
-- ============================================================
('Abrazadera metálica 1/2" (12-20mm) x5u',     NULL,  'fijaciones',         680.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Abrazadera metálica 3/4" (18-32mm) x5u',     NULL,  'fijaciones',         780.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Abrazadera metálica 1" (25-40mm) x5u',        NULL,  'fijaciones',         880.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Abrazadera metálica 2" (50-70mm) x5u',        NULL,  'fijaciones',        1100.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Grapa para manguera 1/2" x20u',              NULL,  'fijaciones',          650.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Grapa para cable 6mm x50u',                  NULL,  'fijaciones',          580.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Grapa para caño EMT 3/4" x20u',              NULL,  'fijaciones',          650.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),

-- ============================================================
-- VARILLAS ROSCADAS Y ESPÁRRAGOS
-- ============================================================
('Varilla roscada M8 x1m',                      NULL,  'fijaciones',        3200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad","sku_prefix": "FERRET"}'),
('Varilla roscada M10 x1m',                     NULL,  'fijaciones',        4200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad","sku_prefix": "FERRET"}'),
('Varilla roscada M12 x1m',                     NULL,  'fijaciones',        5500.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad","sku_prefix": "FERRET"}'),

-- ============================================================
-- CADENAS Y GRILLETES
-- ============================================================
('Cadena galvanizada 3mm x10m',                 NULL,  'fijaciones',        6500.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "rollo","sku_prefix": "FERRET"}'),
('Cadena galvanizada 5mm x5m',                  NULL,  'fijaciones',        7500.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "rollo","sku_prefix": "FERRET"}'),
('Grillete galvanizado 1/4" x2u',               NULL,  'fijaciones',         850.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Grillete galvanizado 3/8" x2u',               NULL,  'fijaciones',        1200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Gancho de acero con seguro 50mm x2u',         NULL,  'fijaciones',        1400.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),

-- ============================================================
-- REMACHES Y ROBLONES
-- ============================================================
('Remache pop 3.2x6mm x100u',                   NULL,  'fijaciones',         680.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Remache pop 4x10mm x100u',                    NULL,  'fijaciones',         820.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Remache pop 4.8x12mm x50u',                   NULL,  'fijaciones',         780.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Remachadora manual para pop Truper',           'Truper', 'fijaciones',     4500.00, 'seed', 0.7, 76, FALSE, TRUE, 'ferreteria', '{"unit": "unidad","sku_prefix": "FERRET"}'),

-- ============================================================
-- PASADORES Y TRABAS
-- ============================================================
('Pasador de puerta 2" galvanizado x5u',        NULL,  'fijaciones',         820.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Pasador de puerta 4" galvanizado x5u',        NULL,  'fijaciones',        1100.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Tornillo de seguridad 5/16" con arandela x10u',NULL, 'fijaciones',         950.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),

-- ============================================================
-- BISAGRAS (también útil en cerrajería)
-- ============================================================
('Bisagra de puerta 3" galvanizada x2u',        NULL,  'fijaciones',        1200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Bisagra de puerta 4" galvanizada x2u',        NULL,  'fijaciones',        1600.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Bisagra de piano 50mm x30cm acero',           NULL,  'fijaciones',        2800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad","sku_prefix": "FERRET"}'),
('Bisagra de llave piano de piano 40mm x1m',    NULL,  'fijaciones',        5500.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad","sku_prefix": "FERRET"}'),

-- ============================================================
-- TACO QUÍMICO Y ANCLAJES ESPECIALES — Fischer
-- ============================================================
('Taco químico Fischer FIS V 360S',             'Fischer', 'fijaciones',    18000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad","sku_prefix": "FERRET"}'),
('Taco químico Fischer FIS EМ plus 390T',       'Fischer', 'fijaciones',    25000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad","sku_prefix": "FERRET"}'),
('Barra de anclaje M8x80mm Fischer x10u',       'Fischer', 'fijaciones',     8500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}'),
('Barra de anclaje M10x100mm Fischer x10u',     'Fischer', 'fijaciones',    11000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "pack", "sku_prefix": "FERRET"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 101 productos
-- Categorías usadas: tornillos-clavos, fijaciones
-- Marcas: Fischer, Truper, genéricos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
-- Nota: tornillos/clavos se venden mayoritariamente a granel o en pack en ferreterías NEA
