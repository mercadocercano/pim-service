-- Seed 080: Productos globales — Ferretería (~23 productos)
-- CICLO: cycle-005
-- FECHA: 2026-04-24
-- FUENTE: catálogos ferreterías NEA (Posadas) + distribuidores Stanley/Bosch/Makita 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- REQUIERE:
--   seed 077 ejecutado (marca black-decker)
--   seeds 055/065 ejecutados (marcas stanley, tramontina, bosch, makita, tigre, fv, tres-m)
-- CATEGORÍAS REQUERIDAS (slugs confirmados):
--   herramientas-manuales, taladros-percutores, sierras-electricas,
--   canos-pvc, latex-ferret, membranas-aislantes

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- HERRAMIENTAS MANUALES — Stanley
-- ============================================================
('Martillo carpintero 16oz Stanley',                    'Stanley',     'herramientas-manuales',  12500.00, 'seed', 0.7, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Set destornilladores Phillips/plano 6u Stanley',      'Stanley',     'herramientas-manuales',  14800.00, 'seed', 0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),
('Alicate multigrip 8" Stanley',                        'Stanley',     'herramientas-manuales',   9200.00, 'seed', 0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cinta métrica 5m Stanley FatMax',                     'Stanley',     'herramientas-manuales',   8500.00, 'seed', 0.7, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- HERRAMIENTAS MANUALES — Tramontina
-- ============================================================
('Llave inglesa 12" Tramontina',                        'Tramontina',  'herramientas-manuales',  11000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Juego llaves combinadas 8u Tramontina',               'Tramontina',  'herramientas-manuales',  22000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),

-- ============================================================
-- HERRAMIENTAS ELÉCTRICAS — Bosch
-- ============================================================
('Taladro percutor GSB 13 RE 650W Bosch',               'Bosch',       'taladros-percutores',    68000.00, 'seed', 0.8, 88, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Amoladora GWS 700 4.5" Bosch',                        'Bosch',       'sierras-electricas',     52000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Sierra caladora GST 700 500W Bosch',                  'Bosch',       'sierras-electricas',     78000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- HERRAMIENTAS ELÉCTRICAS — Makita
-- ============================================================
('Taladro inalámbrico DDF453 18V Makita',               'Makita',      'taladros-percutores',   145000.00, 'seed', 0.8, 88, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Amoladora GA4530 4.5" 720W Makita',                   'Makita',      'sierras-electricas',     58000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- HERRAMIENTAS ELÉCTRICAS — Black+Decker
-- ============================================================
('Taladro percutor CD714CRES 500W Black+Decker',        'Black+Decker','taladros-percutores',    42000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- PLOMERÍA — Tigre
-- ============================================================
('Caño PVC 4" x 3m Tigre',                             'Tigre',       'canos-pvc',             14500.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Codo 90° 4" PVC Tigre',                              'Tigre',       'canos-pvc',              2200.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Unión simple 4" PVC Tigre',                          'Tigre',       'canos-pvc',              1800.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- PLOMERÍA — FV
-- ============================================================
('Llave de paso para jardín 3/4" FV',                  'FV',          'canos-pvc',              5800.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Ducha teléfono cromada FV',                          'FV',          'canos-pvc',             12500.00, 'seed', 0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- FIJACIONES Y CINTAS — 3M
-- ============================================================
('Cinta doble faz 18mm x 10m 3M',                     '3M',          'membranas-aislantes',    4200.00, 'seed', 0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cinta aisladora negra 18mm x 20m 3M',               '3M',          'membranas-aislantes',    1850.00, 'seed', 0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Masilla poliuretano espuma 600ml 3M',                '3M',          'latex-ferret',           8500.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "cartucho", "sku_prefix": "FERRET"}'),

-- ============================================================
-- FIJACIONES — Genérico / Distribuidor local
-- ============================================================
('Tarugos nylon 6mm x100u',                            NULL,          'herramientas-manuales',  1200.00, 'seed', 0.5, 68, FALSE, TRUE, 'ferreteria', '{"unit": "caja",   "sku_prefix": "FERRET"}'),
('Tornillos cabeza phillips 5x60mm caja 25u',          NULL,          'herramientas-manuales',  1800.00, 'seed', 0.5, 68, FALSE, TRUE, 'ferreteria', '{"unit": "caja",   "sku_prefix": "FERRET"}'),
('Nivel de burbuja 40cm aluminio',                     NULL,          'herramientas-manuales',  4500.00, 'seed', 0.5, 68, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 23 productos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
-- Marcas requeridas: stanley, tramontina, bosch, makita (seed 055),
--   black-decker (seed 077), tigre, fv, tres-m (seed 055 — slug 'tres-m')
-- Items sin marca (NULL): fijaciones genéricas sin proveedor identificado
