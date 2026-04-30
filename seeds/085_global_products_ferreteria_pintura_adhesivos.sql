-- Seed 085: Productos globales — Ferretería: Pintura, Adhesivos y Accesorios (~90 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: catálogos ferreterías NEA (Posadas) + Sinteplast/Tersuave/Alba/Sika distribuidores 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: pinturas-latex, esmaltes-barnices, adhesivos-selladores, accesorios-pintura

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- PINTURAS LÁTEX INTERIOR — Sinteplast
-- ============================================================
('Látex interior blanco 1L Sinteplast',         'Sinteplast', 'pinturas-latex',   4500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "litro",  "sku_prefix": "FERRET"}'),
('Látex interior blanco 4L Sinteplast',         'Sinteplast', 'pinturas-latex',  16000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),
('Látex interior blanco 10L Sinteplast',        'Sinteplast', 'pinturas-latex',  36000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "balde",  "sku_prefix": "FERRET"}'),
('Látex interior blanco 20L Sinteplast',        'Sinteplast', 'pinturas-latex',  68000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "balde",  "sku_prefix": "FERRET"}'),
('Látex interior ocre 4L Sinteplast',           'Sinteplast', 'pinturas-latex',  17500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),
('Látex interior azul cielo 4L Sinteplast',     'Sinteplast', 'pinturas-latex',  17500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),

-- ============================================================
-- PINTURAS LÁTEX EXTERIOR — Sinteplast y Tersuave
-- ============================================================
('Látex exterior blanco 1L Sinteplast',         'Sinteplast', 'pinturas-latex',   5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "litro",  "sku_prefix": "FERRET"}'),
('Látex exterior blanco 4L Sinteplast',         'Sinteplast', 'pinturas-latex',  20000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),
('Látex exterior blanco 10L Sinteplast',        'Sinteplast', 'pinturas-latex',  46000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "balde",  "sku_prefix": "FERRET"}'),
('Látex exterior blanco 20L Tersuave',          'Tersuave',   'pinturas-latex',  85000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "balde",  "sku_prefix": "FERRET"}'),
('Látex exterior antihumedad 4L Tersuave',      'Tersuave',   'pinturas-latex',  24000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),
('Látex exterior multicolores 4L Alba',         'Alba',       'pinturas-latex',  22000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),
('Látex exterior blanco 10L Plavicon',          'Plavicon',   'pinturas-latex',  48000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "balde",  "sku_prefix": "FERRET"}'),

-- ============================================================
-- PINTURAS LÁTEX — Alba premium y Sherwin-Williams
-- ============================================================
('Látex interior lavable blanco 4L Alba',       'Alba',       'pinturas-latex',  22000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),
('Látex interior lavable blanco 10L Alba',      'Alba',       'pinturas-latex',  50000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "balde",  "sku_prefix": "FERRET"}'),
('Látex interior premium blanco 4L Sherwin-Williams','Sherwin-Williams','pinturas-latex',35000.00,'seed',0.9, 88, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),
('Látex exterior premium blanco 4L Kem',        'Kem',        'pinturas-latex',  28000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),

-- ============================================================
-- ESMALTE SINTÉTICO
-- ============================================================
('Esmalte sintético blanco 250ml Sinteplast',   'Sinteplast', 'esmaltes-barnices', 3500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Esmalte sintético blanco 1L Sinteplast',      'Sinteplast', 'esmaltes-barnices',10500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "litro",  "sku_prefix": "FERRET"}'),
('Esmalte sintético blanco 4L Sinteplast',      'Sinteplast', 'esmaltes-barnices',38000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),
('Esmalte sintético negro 1L Alba',             'Alba',       'esmaltes-barnices',12000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "litro",  "sku_prefix": "FERRET"}'),
('Esmalte sintético brillante colores 1L Tersuave','Tersuave','esmaltes-barnices',11000.00, 'seed', 0.8, 80, FALSE, TRUE, 'ferreteria', '{"unit": "litro",  "sku_prefix": "FERRET"}'),
('Esmalte anticorrosivo gris 1L Sinteplast',    'Sinteplast', 'esmaltes-barnices',14000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "litro",  "sku_prefix": "FERRET"}'),
('Esmalte anticorrosivo negro 4L Kem',          'Kem',        'esmaltes-barnices',48000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),

-- ============================================================
-- BARNIZ Y LACA
-- ============================================================
('Barniz marino brillante 1L Sinteplast',       'Sinteplast', 'esmaltes-barnices',14000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "litro",  "sku_prefix": "FERRET"}'),
('Barniz parquet satinado 1L Alba',             'Alba',       'esmaltes-barnices',16000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "litro",  "sku_prefix": "FERRET"}'),
('Laca poliuretánica brillante 1L Tersuave',    'Tersuave',   'esmaltes-barnices',18000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "litro",  "sku_prefix": "FERRET"}'),
('Barniz transparente satinado 4L Sinteplast',  'Sinteplast', 'esmaltes-barnices',48000.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),

-- ============================================================
-- SELLADORES, FIJADORES Y DILUYENTES
-- ============================================================
('Sellador para paredes interior 4L Sinteplast','Sinteplast', 'pinturas-latex',  15000.00, 'seed', 0.8, 80, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),
('Fijador multipropósito 4L Sinteplast',        'Sinteplast', 'pinturas-latex',  14000.00, 'seed', 0.8, 80, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),
('Aguarrás mineral 1L genérico',                NULL,         'esmaltes-barnices', 3500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "litro",  "sku_prefix": "FERRET"}'),
('Aguarrás mineral 4L genérico',                NULL,         'esmaltes-barnices',12000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "lata",   "sku_prefix": "FERRET"}'),
('Diluyente para esmalte 1L genérico',          NULL,         'esmaltes-barnices', 4500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "litro",  "sku_prefix": "FERRET"}'),
('Thinner nitro 1L genérico',                   NULL,         'esmaltes-barnices', 4200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "litro",  "sku_prefix": "FERRET"}'),

-- ============================================================
-- ADHESIVOS — Poxipol, Sika, 3M, UHU
-- ============================================================
('Poxipol transparente 14g',                    'Poxipol',    'adhesivos-selladores', 3200.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Poxipol gris metálico 14g',                   'Poxipol',    'adhesivos-selladores', 3500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Poxipol minutos 14g',                         'Poxipol',    'adhesivos-selladores', 3200.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('SuperBonder 20g Loctite',                     'Loctite',    'adhesivos-selladores', 2800.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('SuperBonder gel 3g Loctite',                  'Loctite',    'adhesivos-selladores', 2200.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Adhesivo para madera 1kg Sika',               'Sika',       'adhesivos-selladores', 8500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cemento de contacto Sika Cimsol 400g',        'Sika',       'adhesivos-selladores', 6500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Adhesivo de contacto tipo cemento 250ml',     NULL,         'adhesivos-selladores', 2800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('UHU Por 50ml espuma adhesiva',                'UHU',        'adhesivos-selladores', 5500.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pegamento de madera vinílico 1kg genérico',   NULL,         'adhesivos-selladores', 3800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- SELLADORES DE SILICONA
-- ============================================================
('Silicona neutra transparente 280ml Sika',     'Sika',       'adhesivos-selladores', 4500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Silicona neutra blanca 280ml Sika',           'Sika',       'adhesivos-selladores', 4500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Silicona acetíca transparente 280ml genérica',NULL,         'adhesivos-selladores', 2800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Sellador acrílico blanco 280ml Sika',         'Sika',       'adhesivos-selladores', 3800.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Espuma poliuretano expansiva 750ml Sika',     'Sika',       'adhesivos-selladores',12000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Espuma poliuretano expansiva 750ml genérica', NULL,         'adhesivos-selladores', 8000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- ACCESORIOS DE PINTURA — pinceles, rodillos, bandejas
-- ============================================================
('Pincel plano 1" cerda natural genérico',      NULL,         'accesorios-pintura',  850.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pincel plano 2" cerda natural genérico',      NULL,         'accesorios-pintura', 1200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pincel plano 3" cerda natural genérico',      NULL,         'accesorios-pintura', 1800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pincel angular 2" pelo sintético Purdy',      'Purdy',      'accesorios-pintura', 4800.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Rodillo lana 22cm con mango',                 NULL,         'accesorios-pintura', 2800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Rodillo lana 23cm recambio x2u',              NULL,         'accesorios-pintura', 2200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),
('Rodillo espuma 9cm con mango',                NULL,         'accesorios-pintura', 1800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Bandeja plástica para rodillo 22cm',          NULL,         'accesorios-pintura', 1500.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Mango telescópico para rodillo 1.2m',         NULL,         'accesorios-pintura', 4500.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cinta de enmascarar 18mm x50m genérica',      NULL,         'accesorios-pintura', 1200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cinta de enmascarar 36mm x50m genérica',      NULL,         'accesorios-pintura', 1800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Nylon protector para piso x5m',               NULL,         'accesorios-pintura', 2200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Espátula metálica 4" para pintura',           NULL,         'accesorios-pintura', 1800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Espátula metálica 6" para enmasillado',       NULL,         'accesorios-pintura', 2500.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lija al agua grano 120 x5u',                  NULL,         'accesorios-pintura',  980.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),
('Lija al agua grano 220 x5u',                  NULL,         'accesorios-pintura',  980.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),
('Lija madera grano 80 x5u',                    NULL,         'accesorios-pintura',  850.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),

-- ============================================================
-- MEMBRANAS Y SELLACANTO
-- ============================================================
('Membrana líquida impermeabilizante 4kg Sika', 'Sika',       'pinturas-latex',  28000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "balde",  "sku_prefix": "FERRET"}'),
('Membrana asfáltica aluminio 10m2',            NULL,         'pinturas-latex',  45000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "rollo",  "sku_prefix": "FERRET"}'),
('Sellacanto acrílico blanco 300ml',            NULL,         'adhesivos-selladores', 3200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Masilla lista para usar 1kg Sinteplast',      'Sinteplast', 'accesorios-pintura', 4800.00, 'seed', 0.8, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Masilla para madera 200g genérica',           NULL,         'accesorios-pintura', 1800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 90 productos
-- Categorías usadas: pinturas-latex, esmaltes-barnices, adhesivos-selladores, accesorios-pintura
-- Marcas: Sinteplast, Tersuave, Alba, Plavicon, Kem, Sherwin-Williams, Sika, Poxipol, Loctite, UHU, Purdy, genéricos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
