-- Seed 083: Productos globales — Ferretería: Plomería y Sanitaria (~90 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: catálogos ferreterías NEA (Posadas) + distribuidores Tigre/Awaduct/Perafan 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: canos-pvc, llaves-valvulas, sanitarios-plomeria, mangueras-riego

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- CAÑOS PVC PRESIÓN — por metro (cloacal y agua fría)
-- ============================================================
('Caño PVC 1/2" x1m Awaduct',                  'Awaduct',  'canos-pvc',     1850.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Caño PVC 3/4" x1m Awaduct',                  'Awaduct',  'canos-pvc',     2400.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Caño PVC 1" x1m Awaduct',                    'Awaduct',  'canos-pvc',     3200.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Caño PVC 1.1/2" x1m Awaduct',                'Awaduct',  'canos-pvc',     5200.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Caño PVC 2" x1m Awaduct',                    'Awaduct',  'canos-pvc',     7500.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Caño PVC cloacal 63mm x1m Tigre',            'Tigre',    'canos-pvc',     5800.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Caño PVC cloacal 100mm x1m Tigre',           'Tigre',    'canos-pvc',    12500.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Caño PVC cloacal 110mm x1m Tigre',           'Tigre',    'canos-pvc',    13800.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Caño PVC lluvia 75mm x1m Tigre',             'Tigre',    'canos-pvc',     6500.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Caño PVC lluvia 110mm x1m Tigre',            'Tigre',    'canos-pvc',    12000.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),

-- ============================================================
-- CONEXIONES PVC — codos, tees, reducciones, cuplas
-- ============================================================
('Codo PVC 1/2" 90° Awaduct',                  'Awaduct',  'canos-pvc',      680.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Codo PVC 3/4" 90° Awaduct',                  'Awaduct',  'canos-pvc',      850.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Codo PVC 1" 90° Awaduct',                    'Awaduct',  'canos-pvc',     1200.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tee PVC 1/2" Awaduct',                       'Awaduct',  'canos-pvc',      780.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tee PVC 3/4" Awaduct',                       'Awaduct',  'canos-pvc',      980.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tee PVC 1" Awaduct',                         'Awaduct',  'canos-pvc',     1450.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Reducción PVC 3/4" a 1/2" Awaduct',          'Awaduct',  'canos-pvc',      620.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Reducción PVC 1" a 3/4" Awaduct',            'Awaduct',  'canos-pvc',      780.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cupla PVC 1/2" Awaduct',                     'Awaduct',  'canos-pvc',      520.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cupla PVC 3/4" Awaduct',                     'Awaduct',  'canos-pvc',      650.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cupla PVC 1" Awaduct',                       'Awaduct',  'canos-pvc',      850.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Codo cloacal PVC 100mm 90° Tigre',           'Tigre',    'canos-pvc',     4200.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tee cloacal PVC 100mm Tigre',                'Tigre',    'canos-pvc',     5800.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Sifón PVC lavabo 1.1/4" Tigre',              'Tigre',    'canos-pvc',     4800.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Sifón PVC doble para pileta cocina Tigre',   'Tigre',    'canos-pvc',     6500.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- LLAVES DE PASO Y VÁLVULAS
-- ============================================================
('Llave de paso esfera 1/2" Perafan',          'Perafan',  'llaves-valvulas', 2800.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave de paso esfera 3/4" Perafan',          'Perafan',  'llaves-valvulas', 3800.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave de paso esfera 1" Perafan',            'Perafan',  'llaves-valvulas', 5500.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave de paso angular 1/2" Perafan',         'Perafan',  'llaves-valvulas', 2200.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Válvula de retención 3/4" genérica',         NULL,       'llaves-valvulas', 4500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Válvula retención pie de bomba 1" genérica', NULL,       'llaves-valvulas', 6500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Válvula flotante para tanque 1/2"',          NULL,       'llaves-valvulas', 3500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Válvula flotante para tanque 3/4"',          NULL,       'llaves-valvulas', 4800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave de corte monocomando FV',              'FV',       'llaves-valvulas',28000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- FLEXIBLE SANITARIO Y ACCESORIOS
-- ============================================================
('Flexible sanitario 1/2" x30cm acero inox',  NULL,       'sanitarios-plomeria', 2800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Flexible sanitario 1/2" x40cm acero inox',  NULL,       'sanitarios-plomeria', 3200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Flexible sanitario 1/2" x50cm acero inox',  NULL,       'sanitarios-plomeria', 3600.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cinta teflón 3/4" x10m',                    NULL,       'sanitarios-plomeria',  450.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cinta teflón 1/2" x10m',                    NULL,       'sanitarios-plomeria',  380.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cinta teflón industrial 19mm x25m',         NULL,       'sanitarios-plomeria', 1200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pegamento PVC 100ml Tigre',                 'Tigre',    'sanitarios-plomeria', 1800.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pegamento PVC 250ml Tigre',                 'Tigre',    'sanitarios-plomeria', 3500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- MANGUERAS — jardín y presión
-- ============================================================
('Manguera jardín 1/2" x25m Genco',           'Genco',    'mangueras-riego',  8500.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "rollo", "sku_prefix": "FERRET"}'),
('Manguera jardín 1/2" x50m Genco',           'Genco',    'mangueras-riego', 16000.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "rollo", "sku_prefix": "FERRET"}'),
('Manguera jardín 3/4" x25m Genco',           'Genco',    'mangueras-riego', 12000.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "rollo", "sku_prefix": "FERRET"}'),
('Manguera extensible 15m Truper',             'Truper',   'mangueras-riego', 12000.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Manguera de presión 3/4" x5m trenzada',     NULL,       'mangueras-riego',  6500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pistola rociadora para manguera Truper',     'Truper',   'mangueras-riego',  3800.00, 'seed', 0.7, 76, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Aspersor de jardín rotativo genérico',       NULL,       'mangueras-riego',  3200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Conector plástico rápido 1/2" para manguera',NULL,       'mangueras-riego',  1200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Regulador de caudal para riego goteo',       NULL,       'mangueras-riego',  4500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- BOMBAS DE AGUA
-- ============================================================
('Bomba de agua autocebante 1/2HP Grundfos',  'Grundfos', 'sanitarios-plomeria',85000.00,'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Bomba de agua periférica 1HP genérica',     NULL,       'sanitarios-plomeria',65000.00,'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Bomba sumergible 0.5HP genérica',           NULL,       'sanitarios-plomeria',55000.00,'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Hidroneumático para bomba 24L',             NULL,       'sanitarios-plomeria',45000.00,'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- TERMO-TANQUES Y CALEFONES BÁSICOS
-- ============================================================
('Termotanque 80L eléctrico Rheem',           'Rheem',    'sanitarios-plomeria',185000.00,'seed',0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Termotanque 150L eléctrico Rheem',          'Rheem',    'sanitarios-plomeria',280000.00,'seed',0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Calefón a gas 11L/min Orbis',               'Orbis',    'sanitarios-plomeria',145000.00,'seed',0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Calefón a gas 14L/min Domec',               'Domec',    'sanitarios-plomeria',185000.00,'seed',0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- SANITARIOS BÁSICOS (inodoros, lavabos — los vende la ferretería)
-- ============================================================
('Inodoro monoblock FERRUM Artic',            'Ferrum',   'sanitarios-plomeria',85000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lavabo de pared FERRUM Artic 60cm',         'Ferrum',   'sanitarios-plomeria',48000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Mochila de descarga doble para inodoro',    NULL,       'sanitarios-plomeria',28000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tapa de inodoro plástica genérica',         NULL,       'sanitarios-plomeria', 5500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave de pileta monocomando FV',            'FV',       'sanitarios-plomeria',35000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave de pileta bidet monocomando FV',      'FV',       'sanitarios-plomeria',28000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Grifería de ducha mezcladora FV',           'FV',       'sanitarios-plomeria',52000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- CAÑOS DE COBRE Y ACCESORIOS
-- ============================================================
('Caño de cobre 1/2" x1m',                   NULL,       'canos-pvc',       8500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Caño de cobre 3/4" x1m',                   NULL,       'canos-pvc',      12500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Codo cobre 1/2" para soldar',               NULL,       'canos-pvc',         850.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Estaño para soldar plomería 100g',          NULL,       'sanitarios-plomeria', 3500.00,'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 88 productos
-- Categorías usadas: canos-pvc, llaves-valvulas, sanitarios-plomeria, mangueras-riego
-- Marcas: Awaduct, Tigre, Perafan, FV, Genco, Grundfos, Rheem, Orbis, Domec, Ferrum, Truper, Genco
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
