-- Seed 095: Productos globales — Corralón: Materiales de Construcción (~90 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento corralones NEA (Posadas) + distribuidores materiales 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: cementos-cales, revoques-yesos, ladrillos-bloques, membranas-chapas, ceramicos-porcelanato, madera-molduras

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- CEMENTO — Loma Negra, Cementos Avellaneda
-- ============================================================
('Cemento Portland 50kg Loma Negra',             'Loma Negra',         'cementos-cales',   18500.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Cemento Portland 25kg Loma Negra',             'Loma Negra',         'cementos-cales',    9800.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Cemento de albañilería 40kg Loma Negra',       'Loma Negra',         'cementos-cales',   14000.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Cemento Portland 50kg Cementos Avellaneda',    'Cementos Avellaneda','cementos-cales',   17800.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Cemento de albañilería 40kg Cerro Negro',      'Cerro Negro',        'cementos-cales',   13500.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Cemento rápido 25kg Loma Negra',               'Loma Negra',         'cementos-cales',   12000.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),

-- ============================================================
-- CAL
-- ============================================================
('Cal hidratada 25kg Minetti',                   'Minetti',            'cementos-cales',    6500.00, 'seed', 0.8, 80, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Cal hidratada 30kg genérica',                  NULL,                 'cementos-cales',    5800.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Cal viva 25kg genérica',                       NULL,                 'cementos-cales',    4800.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),

-- ============================================================
-- ARENAS Y ÁRIDOS
-- ============================================================
('Arena fina bolsa 25kg',                        NULL,                 'cementos-cales',    3500.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Arena gruesa bolsa 25kg',                      NULL,                 'cementos-cales',    3200.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Canto rodado bolsa 25kg',                      NULL,                 'cementos-cales',    3800.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Pedregullo 6/12 bolsa 25kg',                   NULL,                 'cementos-cales',    4000.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),

-- ============================================================
-- REVOQUES Y YESOS
-- ============================================================
('Revoque fino 25kg Loma Negra',                 'Loma Negra',         'revoques-yesos',    7500.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Revoque grueso 40kg Loma Negra',               'Loma Negra',         'revoques-yesos',    9000.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Revoque interior 25kg Parex',                  'Parex',              'revoques-yesos',    8200.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Yeso 25kg Knauf',                              'Knauf',              'revoques-yesos',    6800.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Yeso 40kg Knauf',                              'Knauf',              'revoques-yesos',   10500.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Yeso proyectable 25kg genérico',               NULL,                 'revoques-yesos',    6000.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Enduido exterior 25kg Sinteplast',             'Sinteplast',         'revoques-yesos',   14000.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Enduido interior 20kg Plavicon',               'Plavicon',           'revoques-yesos',   12500.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),

-- ============================================================
-- LADRILLOS Y BLOQUES
-- ============================================================
('Ladrillo común 8x18x38cm x unidad',            NULL,                 'ladrillos-bloques',   180.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Ladrillo común x1000 unidades',                NULL,                 'ladrillos-bloques',160000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "millar", "sku_prefix": "CORR"}'),
('Ladrillo hueco 8x18x33cm x unidad',            NULL,                 'ladrillos-bloques',   220.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Ladrillo hueco 12x18x33cm x unidad',           NULL,                 'ladrillos-bloques',   280.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Bloque de hormigón 15x20x40cm x unidad',       NULL,                 'ladrillos-bloques',   650.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Bloque de hormigón 20x20x40cm x unidad',       NULL,                 'ladrillos-bloques',   850.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Ladrillo refractario 25x12x6cm x unidad',      NULL,                 'ladrillos-bloques',   480.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- MEMBRANAS E IMPERMEABILIZANTES
-- ============================================================
('Membrana asfáltica 3mm 10m2 Impetek',          'Impetek',            'membranas-chapas',  28000.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "rollo",  "sku_prefix": "CORR"}'),
('Membrana asfáltica 4mm 10m2 Impetek',          'Impetek',            'membranas-chapas',  36000.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "rollo",  "sku_prefix": "CORR"}'),
('Membrana líquida 20kg Iggam',                  'Iggam',              'membranas-chapas',  18000.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "balde",  "sku_prefix": "CORR"}'),
('Membrana líquida 4kg Sika',                    'Sika',               'membranas-chapas',   8500.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "balde",  "sku_prefix": "CORR"}'),
('Hidrofugo líquido 1L Sika',                    'Sika',               'membranas-chapas',   6200.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Hidrofugo líquido 5L Sika',                    'Sika',               'membranas-chapas',  26000.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "bidon",  "sku_prefix": "CORR"}'),

-- ============================================================
-- CHAPAS Y TEJAS
-- ============================================================
('Chapa acanalada 2m calibre 27 Eternit',        'Eternit',            'membranas-chapas',  12500.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa acanalada 3m calibre 27 Eternit',        'Eternit',            'membranas-chapas',  18000.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa acanalada 4m calibre 25 Eternit',        'Eternit',            'membranas-chapas',  26000.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Chapa trapezoidal 3m calibre 25',              NULL,                 'membranas-chapas',  19500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Teja cerámica española x unidad',              NULL,                 'membranas-chapas',    950.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Teja de hormigón colonial x unidad',           NULL,                 'membranas-chapas',   1100.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Teja pizarreña fibrocemento x unidad',         'Eternit',            'membranas-chapas',   1450.00, 'seed', 0.9, 82, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- CERÁMICOS Y PORCELANATO
-- ============================================================
('Cerámico piso 33x33cm liso blanco x m2',       NULL,                 'ceramicos-porcelanato', 4800.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "m2",    "sku_prefix": "CORR"}'),
('Cerámico piso 45x45cm beige x m2',             NULL,                 'ceramicos-porcelanato', 6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "m2",    "sku_prefix": "CORR"}'),
('Cerámico pared 20x30cm blanco x m2',           NULL,                 'ceramicos-porcelanato', 5200.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "m2",    "sku_prefix": "CORR"}'),
('Porcelanato rectificado 60x60cm gris x m2',    NULL,                 'ceramicos-porcelanato',12000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "m2",    "sku_prefix": "CORR"}'),
('Porcelanato madera simil 20x80cm x m2',        NULL,                 'ceramicos-porcelanato',14500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "m2",    "sku_prefix": "CORR"}'),
('Pastina gris 1kg Klaukol',                     'Klaukol',            'ceramicos-porcelanato', 2200.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Pastina blanca 1kg Klaukol',                   'Klaukol',            'ceramicos-porcelanato', 2200.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Adhesivo gris 25kg Klaukol',                   'Klaukol',            'ceramicos-porcelanato', 8500.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Adhesivo flexible 25kg Klaukol',               'Klaukol',            'ceramicos-porcelanato',11000.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Rejilla plástica desagüe 10x10cm',             NULL,                 'ceramicos-porcelanato', 1200.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- MADERA Y TIRANTES
-- ============================================================
('Tirante de pino 1"x4" 3m',                     NULL,                 'madera-molduras',   4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Tirante de pino 2"x4" 3m',                     NULL,                 'madera-molduras',   8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Tirante de pino 2"x6" 3m',                     NULL,                 'madera-molduras',  12000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Tirante de pino 2"x4" 6m',                     NULL,                 'madera-molduras',  16000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Tabla de pino 1"x6" 3m',                       NULL,                 'madera-molduras',   6800.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Fenólico 18mm 2.44x1.22m',                     NULL,                 'madera-molduras',  45000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "plancha","sku_prefix": "CORR"}'),
('MDF 15mm 2.44x1.83m',                          NULL,                 'madera-molduras',  38000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "plancha","sku_prefix": "CORR"}'),
('OSB 11mm 2.44x1.22m',                          NULL,                 'madera-molduras',  28000.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "plancha","sku_prefix": "CORR"}'),
('Moldura cielorraso de madera 3m',              NULL,                 'madera-molduras',   3200.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Zócalo madera pino 1"x2" 3m',                  NULL,                 'madera-molduras',   2800.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- CAÑOS Y CONEXIONES PVC CONSTRUCTIVO
-- ============================================================
('Caño PVC sanitario 110mm x 3m',               NULL,                 'cementos-cales',   12500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Caño PVC sanitario 75mm x 3m',                NULL,                 'cementos-cales',    8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Caño PVC presión 1/2" x 3m',                  NULL,                 'cementos-cales',    4800.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Caño PVC presión 3/4" x 3m',                  NULL,                 'cementos-cales',    6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Codo PVC 110mm x 90° sanitario',              NULL,                 'cementos-cales',    2800.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Unión PVC 110mm',                              NULL,                 'cementos-cales',    1800.00, 'seed', 0.7, 60, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- ADITIVOS Y PEGAMENTOS PARA CONSTRUCCIÓN
-- ============================================================
('Hidrófugo concentrado 1L Sika',               'Sika',               'revoques-yesos',    5800.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Aditivo para hormigón 1L Sika',               'Sika',               'revoques-yesos',    6200.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Pegamento para cerámica 5kg Sika',            'Sika',               'ceramicos-porcelanato',9500.00,'seed',0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),
('Alisador de pisos autonivelante 25kg Knauf',  'Knauf',              'revoques-yesos',   18000.00, 'seed', 0.9, 85, FALSE, TRUE, 'corralon', '{"unit": "bolsa", "sku_prefix": "CORR"}'),

-- ============================================================
-- MESADAS Y ACCESORIOS DE BAÑO
-- ============================================================
('Bacha de acero inoxidable 45x35cm',           NULL,                 'ceramicos-porcelanato',18000.00,'seed',0.7,60,FALSE,TRUE,'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Inodoro común blanco FERRUM',                 'Ferrum',             'ceramicos-porcelanato',55000.00,'seed',0.8,82,FALSE,TRUE,'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),
('Lavatorio blanco 60cm FERRUM',                'Ferrum',             'ceramicos-porcelanato',42000.00,'seed',0.8,82,FALSE,TRUE,'corralon', '{"unit": "unidad", "sku_prefix": "CORR"}'),

-- ============================================================
-- PINTURA LÁTEX CONSTRUCTIVA
-- ============================================================
('Pintura látex interior 4L Alba',              'Alba',               'revoques-yesos',    18000.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "lata",  "sku_prefix": "CORR"}'),
('Pintura látex exterior 4L Alba',              'Alba',               'revoques-yesos',    22000.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "lata",  "sku_prefix": "CORR"}'),
('Pintura látex interior 20L Plavicon',         'Plavicon',           'revoques-yesos',    75000.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "balde", "sku_prefix": "CORR"}'),
('Pintura látex exterior 20L Sinteplast',       'Sinteplast',         'revoques-yesos',    85000.00, 'seed', 0.8, 82, FALSE, TRUE, 'corralon', '{"unit": "balde", "sku_prefix": "CORR"}'),
('Sellador fijador 4L genérico',                NULL,                 'revoques-yesos',    12000.00, 'seed', 0.6, 60, FALSE, TRUE, 'corralon', '{"unit": "lata",  "sku_prefix": "CORR"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 81 productos
-- Categorías: cementos-cales, revoques-yesos, ladrillos-bloques, membranas-chapas, ceramicos-porcelanato, madera-molduras
-- Marcas: Loma Negra, Cementos Avellaneda, Cerro Negro, Minetti, Knauf, Parex, Impetek, Sika, Iggam,
--   Eternit, Klaukol, Sinteplast, Plavicon, Alba, Ferrum
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
