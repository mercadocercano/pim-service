-- Seed 079: Productos globales — Corralón: adhesivos, membranas, pinturas (~20 productos)
-- CICLO: cycle-005
-- FECHA: 2026-04-24
-- FUENTE: catálogos distribuidores NEA (Posadas) + relevamiento corralones 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- REQUIERE:
--   seed 077 ejecutado (marcas klaukol, megaflex)
--   seeds 055/065 ejecutados (marcas poxipol, sika, sinteplast, alba, sherwin-williams)
-- CATEGORÍAS REQUERIDAS (slugs confirmados):
--   adhesivos-mezclas, techos, aislantes, pintura, membranas-aislantes (ferretería también)

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- ADHESIVOS — Poxipol
-- ============================================================
('Pegamento epóxico transparente 14ml Poxipol',  'Poxipol', 'adhesivos-mezclas',   2800.00, 'seed', 0.7, 82, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Pegamento epóxico gris 14ml Poxipol',          'Poxipol', 'adhesivos-mezclas',   2800.00, 'seed', 0.7, 82, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Pegamento para metales 14ml Poxipol',          'Poxipol', 'adhesivos-mezclas',   3200.00, 'seed', 0.7, 82, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- ADHESIVOS CERÁMICOS Y PASTINAS — Klaukol
-- ============================================================
('Adhesivo para cerámicos 5kg Klaukol',          'Klaukol', 'adhesivos-mezclas',   5800.00, 'seed', 0.7, 82, FALSE, TRUE, 'corralon', '{"unit": "bolsa",  "sku_prefix": "CORR"}'),
('Adhesivo para porcelanato 25kg Klaukol',       'Klaukol', 'adhesivos-mezclas',  18500.00, 'seed', 0.7, 82, FALSE, TRUE, 'corralon', '{"unit": "bolsa",  "sku_prefix": "CORR"}'),
('Pastina blanca 1kg Klaukol',                   'Klaukol', 'adhesivos-mezclas',   2200.00, 'seed', 0.7, 80, FALSE, TRUE, 'corralon', '{"unit": "kg",     "sku_prefix": "CORR"}'),
('Pastina gris 1kg Klaukol',                     'Klaukol', 'adhesivos-mezclas',   2200.00, 'seed', 0.7, 80, FALSE, TRUE, 'corralon', '{"unit": "kg",     "sku_prefix": "CORR"}'),

-- ============================================================
-- SELLADORES Y ADHESIVOS — Sika
-- ============================================================
('Sikaflex 11FC sellador gris 600ml',            'Sika',    'adhesivos-mezclas',  14500.00, 'seed', 0.7, 84, FALSE, TRUE, 'corralon', '{"unit": "cartucho", "sku_prefix": "CORR"}'),
('SikaTop Armatec 110 puente adherente 1kg',     'Sika',    'adhesivos-mezclas',   8200.00, 'seed', 0.7, 82, FALSE, TRUE, 'corralon', '{"unit": "kg",     "sku_prefix": "CORR"}'),
('Sika Látex aditivo para mezclas 1L',           'Sika',    'adhesivos-mezclas',   5500.00, 'seed', 0.7, 82, FALSE, TRUE, 'corralon', '{"unit": "litro",  "sku_prefix": "CORR"}'),

-- ============================================================
-- MEMBRANAS E IMPERMEABILIZANTES — Megaflex
-- ============================================================
('Membrana líquida blanca 20kg Megaflex',        'Megaflex', 'techos',            32000.00, 'seed', 0.7, 84, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Membrana asfáltica con aluminio 10kg Megaflex','Megaflex', 'techos',            24000.00, 'seed', 0.7, 82, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- MEMBRANAS — Sinteplast
-- ============================================================
('Membrana acrílica impermeabilizante 20L Sinteplast', 'Sinteplast', 'techos',   42000.00, 'seed', 0.7, 80, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- PINTURAS COMPLEMENTARIAS — Alba
-- (complementa seed 052, solo esmalte no cubierto)
-- ============================================================
('Esmalte sintético blanco brillante 1L Alba',   'Alba',    'pintura',           12500.00, 'seed', 0.7, 80, FALSE, TRUE, 'corralon', '{"unit": "litro",  "sku_prefix": "CORR"}'),
('Esmalte sintético negro mate 1L Alba',         'Alba',    'pintura',           12500.00, 'seed', 0.7, 78, FALSE, TRUE, 'corralon', '{"unit": "litro",  "sku_prefix": "CORR"}'),

-- ============================================================
-- PINTURAS COMPLEMENTARIAS — Sherwin-Williams
-- ============================================================
('Látex interior blanco 4L Sherwin-Williams',    'Sherwin-Williams', 'pintura', 28000.00, 'seed', 0.7, 82, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Látex exterior blanco 4L Sherwin-Williams',    'Sherwin-Williams', 'pintura', 32000.00, 'seed', 0.7, 82, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- AISLANTES COMPLEMENTARIOS
-- ============================================================
('Sikaflex 221 sellador multiuso negro 300ml',   'Sika',    'aislantes',         12000.00, 'seed', 0.7, 82, FALSE, TRUE, 'corralon', '{"unit": "cartucho", "sku_prefix": "CORR"}'),
('Adhesivo para porcelanato grande 5kg Klaukol', 'Klaukol', 'adhesivos-mezclas',  5800.00, 'seed', 0.7, 80, FALSE, TRUE, 'corralon', '{"unit": "bolsa",  "sku_prefix": "CORR"}'),
('Pastina porcelanato gris perla 1kg Klaukol',   'Klaukol', 'adhesivos-mezclas',  2500.00, 'seed', 0.7, 80, FALSE, TRUE, 'corralon', '{"unit": "kg",     "sku_prefix": "CORR"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 20 productos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
-- Marcas requeridas: poxipol (seed 055), klaukol (seed 077), sika (seed 055),
--   megaflex (seed 077), sinteplast (seed 055), alba (seed 055), sherwin-williams (seed 055)
