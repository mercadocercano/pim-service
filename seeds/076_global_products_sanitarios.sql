-- Seed 076: Sanitarios y griferías — 30 productos
-- CICLO: cycle-004-brands-catalog-expansion
-- FECHA: 2026-04-22
-- FUENTE: catálogos corralones NEA (Posadas) + distribuidores Ferrum/FV 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- NOTA: Numeración saltó de 073 a 076 porque 074 (business_type_piletas) y
--       075 (business_type_limpieza) fueron creados en la misma iteración del ciclo.
-- REQUIERE: seed 065 ya ejecutado (marcas sanitarios: ferrum, roca-sanitarios, deca, piazza,
--           vasser, peirano). seed 055 ya ejecutado (fv).

-- Categorías requeridas (existentes en marketplace_categories):
--   sanitarios, griferias (o griferia-banos, griferia-cocina)
--   Si las categorías exactas difieren, ajustar el code al valor en BD.

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- INODOROS
-- ============================================================
('Inodoro económico con mochila',            'Piazza',  'sanitarios',  42000.00, 'seed', 0.5, 72, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Inodoro Ferrum línea Iceberg con mochila', 'Ferrum',  'sanitarios',  68000.00, 'seed', 0.6, 78, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Inodoro Ferrum línea Méndez con mochila', 'Ferrum',  'sanitarios',  58000.00, 'seed', 0.6, 76, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Inodoro Roca línea Debba con mochila',    'Roca',    'sanitarios', 105000.00, 'seed', 0.6, 82, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Inodoro Deca modelo Carrara con mochila', 'Deca',    'sanitarios',  72000.00, 'seed', 0.5, 75, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),

-- ============================================================
-- BIDETS
-- ============================================================
('Bidet económico',                          'Piazza',  'sanitarios',  32000.00, 'seed', 0.5, 70, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Bidet Ferrum línea Iceberg',               'Ferrum',  'sanitarios',  52000.00, 'seed', 0.6, 76, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Bidet Roca línea Debba',                   'Roca',    'sanitarios',  78000.00, 'seed', 0.6, 80, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Bidet Deca modelo Carrara',                'Deca',    'sanitarios',  55000.00, 'seed', 0.5, 73, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),

-- ============================================================
-- LAVABOS / PILETAS DE BAÑO
-- ============================================================
('Lavabo colgante 50cm blanco',              'Piazza',  'sanitarios',  28000.00, 'seed', 0.5, 70, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Lavabo Ferrum 55cm línea Iceberg',         'Ferrum',  'sanitarios',  38000.00, 'seed', 0.6, 76, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Lavabo Roca 60cm línea Debba',             'Roca',    'sanitarios',  58000.00, 'seed', 0.6, 80, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Lavabo bajo mesada 60cm',                  'Deca',    'sanitarios',  45000.00, 'seed', 0.5, 73, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),

-- ============================================================
-- BAÑERAS
-- ============================================================
('Bañera acrílica 1.50m x 0.70m',          'Ferrum',  'sanitarios', 195000.00, 'seed', 0.5, 74, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Bañera acrílica 1.70m x 0.75m',          'Roca',    'sanitarios', 270000.00, 'seed', 0.5, 78, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),

-- ============================================================
-- GRIFERÍAS DE BAÑO
-- ============================================================
('Grifería lavabo monocomando FV Zen',       'FV',      'sanitarios',  38000.00, 'seed', 0.6, 78, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "GRI"}'),
('Grifería lavabo de dos llaves FV',         'FV',      'sanitarios',  24000.00, 'seed', 0.6, 75, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "GRI"}'),
('Grifería lavabo monocomando Vasser',       'Vasser',  'sanitarios',  22000.00, 'seed', 0.5, 72, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "GRI"}'),
('Grifería lavabo de dos llaves Peirano',    'Peirano', 'sanitarios',  18000.00, 'seed', 0.5, 70, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "GRI"}'),
('Columna de ducha FV Soho',                 'FV',      'sanitarios',  65000.00, 'seed', 0.6, 78, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "GRI"}'),
('Duchador con flexible 1.5m FV',            'FV',      'sanitarios',  18000.00, 'seed', 0.6, 75, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "GRI"}'),
('Mezclador termostático ducha FV',          'FV',      'sanitarios',  95000.00, 'seed', 0.5, 76, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "GRI"}'),
('Llave de ducha FV económica',              'FV',      'sanitarios',  28000.00, 'seed', 0.6, 74, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "GRI"}'),

-- ============================================================
-- GRIFERÍAS DE COCINA
-- ============================================================
('Grifería pileta cocina cuello alto FV',    'FV',      'sanitarios',  35000.00, 'seed', 0.6, 77, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "GRI"}'),
('Grifería pileta cocina FV modelo Lena',    'FV',      'sanitarios',  28000.00, 'seed', 0.6, 76, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "GRI"}'),
('Grifería cocina monocomando Vasser',       'Vasser',  'sanitarios',  22000.00, 'seed', 0.5, 71, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "GRI"}'),
('Grifería cocina dos llaves Peirano',       'Peirano', 'sanitarios',  16000.00, 'seed', 0.5, 68, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "GRI"}'),

-- ============================================================
-- ACCESORIOS DE BAÑO
-- ============================================================
('Sifón botella PVC 1¼',                    NULL,      'sanitarios',   3500.00, 'seed', 0.5, 65, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Sifón botella PVC 1½',                    NULL,      'sanitarios',   4200.00, 'seed', 0.5, 65, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Flotante de mochila universal',            NULL,      'sanitarios',   4800.00, 'seed', 0.5, 65, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}'),
('Asiento de inodoro blanco universal',      NULL,      'sanitarios',   8500.00, 'seed', 0.5, 65, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "SAN"}')

ON CONFLICT DO NOTHING;

-- Total: 30 productos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
-- Ferrum domina segmento medio-alto en NEA (grupo Saint-Gobain, distribución nacional).
-- FV es grifería masiva líder (slug 'fv' en marketplace_brands desde seed 055).
-- Roca tiene presencia en corralones especializados y proyectos de obra.
-- Piazza cubre el segmento económico con amplia distribución NEA.
