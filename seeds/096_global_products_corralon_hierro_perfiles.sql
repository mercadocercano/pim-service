-- Seed 096: Productos globales — Corralón: Hierro y Perfiles (~65 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento corralones NEA (Posadas) + distribuidores acero 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: hierro-redondo, mallas-electrosoldadas, canos-acero, perfiles-aluminio, tirantes-madera, chapas-acanaladas

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- HIERRO REDONDO — barras lisas y nervuradas
-- ============================================================
('Hierro redondo liso 6mm x 6m',                 NULL, 'hierro-redondo',     2800.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "barra",  "sku_prefix": "CORR"}'),
('Hierro redondo liso 8mm x 6m',                 NULL, 'hierro-redondo',     4800.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "barra",  "sku_prefix": "CORR"}'),
('Hierro redondo liso 10mm x 6m',                NULL, 'hierro-redondo',     7200.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "barra",  "sku_prefix": "CORR"}'),
('Hierro redondo liso 12mm x 6m',                NULL, 'hierro-redondo',    10500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "barra",  "sku_prefix": "CORR"}'),
('Hierro redondo liso 16mm x 6m',                NULL, 'hierro-redondo',    18500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "barra",  "sku_prefix": "CORR"}'),
('Hierro redondo liso 6mm x 12m',                NULL, 'hierro-redondo',     5400.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "barra",  "sku_prefix": "CORR"}'),
('Hierro redondo liso 8mm x 12m',                NULL, 'hierro-redondo',     9200.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "barra",  "sku_prefix": "CORR"}'),
('Hierro redondo liso 10mm x 12m',               NULL, 'hierro-redondo',    13800.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "barra",  "sku_prefix": "CORR"}'),
('Hierro redondo liso 12mm x 12m',               NULL, 'hierro-redondo',    20000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "barra",  "sku_prefix": "CORR"}'),
('Hierro nervurado 10mm x 6m ADN 420',           NULL, 'hierro-redondo',     8000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "barra",  "sku_prefix": "CORR"}'),
('Hierro nervurado 12mm x 6m ADN 420',           NULL, 'hierro-redondo',    11500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "barra",  "sku_prefix": "CORR"}'),
('Hierro nervurado 16mm x 6m ADN 420',           NULL, 'hierro-redondo',    20000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "barra",  "sku_prefix": "CORR"}'),

-- ============================================================
-- MALLAS ELECTROSOLDADAS
-- ============================================================
('Malla electrosoldada 15x15cm 5mm 2.44x6m',    NULL, 'mallas-electrosoldadas', 28000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "plancha","sku_prefix": "CORR"}'),
('Malla electrosoldada 15x15cm 6mm 2.44x6m',    NULL, 'mallas-electrosoldadas', 38000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "plancha","sku_prefix": "CORR"}'),
('Malla electrosoldada 20x20cm 5mm 2.44x6m',    NULL, 'mallas-electrosoldadas', 22000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "plancha","sku_prefix": "CORR"}'),
('Malla electrosoldada 20x20cm 6mm 2.44x6m',    NULL, 'mallas-electrosoldadas', 30000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "plancha","sku_prefix": "CORR"}'),
('Malla electrosoldada 10x10cm 4.2mm 2.44x6m',  NULL, 'mallas-electrosoldadas', 35000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "plancha","sku_prefix": "CORR"}'),

-- ============================================================
-- CAÑOS DE ACERO
-- ============================================================
('Caño de acero negro 1" x 6m SCH40',           NULL, 'canos-acero',       18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Caño de acero negro 1.5" x 6m SCH40',         NULL, 'canos-acero',       26000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Caño de acero negro 2" x 6m SCH40',           NULL, 'canos-acero',       38000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Caño de acero negro 3/4" x 6m SCH40',         NULL, 'canos-acero',       13500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Caño galvanizado 1" x 6m',                    NULL, 'canos-acero',       22000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Caño galvanizado 2" x 6m',                    NULL, 'canos-acero',       45000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Caño cuadrado 25x25mm x 6m',                  NULL, 'canos-acero',       14000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Caño cuadrado 40x40mm x 6m',                  NULL, 'canos-acero',       22000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Caño rectangular 40x20mm x 6m',               NULL, 'canos-acero',       13000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Caño rectangular 60x40mm x 6m',               NULL, 'canos-acero',       21000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- PERFILES DE ACERO
-- ============================================================
('Perfil C 100x50mm x 6m galvanizado',           NULL, 'canos-acero',       28000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Perfil U 100x50mm x 6m galvanizado',           NULL, 'canos-acero',       26000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Perfil L 50x50mm x 6m acero',                 NULL, 'canos-acero',       16000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Viga doble T 100mm x 6m',                     NULL, 'canos-acero',       85000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- PERFILES DE ALUMINIO
-- ============================================================
('Perfil L aluminio 1" x 6m',                   NULL, 'perfiles-aluminio',   8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Perfil L aluminio 2" x 6m',                   NULL, 'perfiles-aluminio',  14000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Perfil T aluminio 1" x 6m',                   NULL, 'perfiles-aluminio',   9000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Perfil T aluminio 2" x 6m',                   NULL, 'perfiles-aluminio',  15500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Perfil U aluminio 1" x 6m',                   NULL, 'perfiles-aluminio',   8000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Perfil U aluminio 2" x 6m',                   NULL, 'perfiles-aluminio',  13500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Perfil cuadrado aluminio 1"x1" x 6m',         NULL, 'perfiles-aluminio',  11000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Perfil rectangular aluminio 1"x2" x 6m',      NULL, 'perfiles-aluminio',  14000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa lisa aluminio 1mm 1.22x2.44m',          NULL, 'perfiles-aluminio',  48000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "plancha","sku_prefix": "CORR"}'),

-- ============================================================
-- TIRANTES DE MADERA (referencia para corralón con aserradero)
-- ============================================================
('Tirante pino 1"x4" 3m cepillado',             NULL, 'tirantes-madera',    4800.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Tirante pino 1"x4" 6m cepillado',             NULL, 'tirantes-madera',    9200.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Tirante pino 2"x4" 3m cepillado',             NULL, 'tirantes-madera',    9000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Tirante pino 2"x4" 6m cepillado',             NULL, 'tirantes-madera',   17000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Tirante eucaliptus rollizo 8cm x 3m',          NULL, 'tirantes-madera',    6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Tirante eucaliptus rollizo 10cm x 3m',         NULL, 'tirantes-madera',    9000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Tirante eucaliptus rollizo 8cm x 6m',          NULL, 'tirantes-madera',   12500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Tirante eucaliptus rollizo 10cm x 6m',         NULL, 'tirantes-madera',   17000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- CHAPAS ACANALADAS
-- ============================================================
('Chapa acanalada galvanizada cal.27 2m',        NULL, 'chapas-acanaladas',  9800.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa acanalada galvanizada cal.27 3m',        NULL, 'chapas-acanaladas', 14500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa acanalada galvanizada cal.27 4m',        NULL, 'chapas-acanaladas', 19000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa acanalada galvanizada cal.25 2m',        NULL, 'chapas-acanaladas', 12500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa acanalada galvanizada cal.25 3m',        NULL, 'chapas-acanaladas', 18500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa acanalada galvanizada cal.25 4m',        NULL, 'chapas-acanaladas', 24500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa acanalada color cal.27 3m',              NULL, 'chapas-acanaladas', 18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa acanalada color cal.27 4m',              NULL, 'chapas-acanaladas', 23500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa acanalada traslúcida policarbonato 3m',  NULL, 'chapas-acanaladas', 32000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa acanalada traslúcida fibra vidrio 3m',   NULL, 'chapas-acanaladas', 22000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 62 productos
-- Categorías: hierro-redondo, mallas-electrosoldadas, canos-acero, perfiles-aluminio, tirantes-madera, chapas-acanaladas
-- Marcas: mayoritariamente genérico (comercio de materiales)
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
