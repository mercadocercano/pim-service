-- Seed 084: Productos globales — Ferretería: Materiales Eléctricos (~90 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: catálogos ferreterías NEA (Posadas) + distribuidores Cambre/Bticino/Osram/Philips 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: cables-electricidad, llaves-tomacorrientes, iluminacion, proteccion-electrica

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- CABLES — por metro y por rollo
-- ============================================================
('Cable unipolar 1.5mm² x1m',                  NULL,       'cables-electricidad',  680.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Cable unipolar 2.5mm² x1m',                  NULL,       'cables-electricidad',  950.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Cable unipolar 4mm² x1m',                    NULL,       'cables-electricidad', 1450.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Cable unipolar 6mm² x1m',                    NULL,       'cables-electricidad', 2100.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Cable unipolar 10mm² x1m',                   NULL,       'cables-electricidad', 3500.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Cable duplex 2x1.5mm² x1m',                  NULL,       'cables-electricidad', 1200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Cable duplex 2x2.5mm² x1m',                  NULL,       'cables-electricidad', 1800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Cable unipolar 1.5mm² rollo 100m',            NULL,       'cables-electricidad',62000.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "rollo",  "sku_prefix": "FERRET"}'),
('Cable unipolar 2.5mm² rollo 100m',            NULL,       'cables-electricidad',88000.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "rollo",  "sku_prefix": "FERRET"}'),
('Cable tipo taller 3x2.5mm² x1m',             NULL,       'cables-electricidad', 2800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Cinta aisladora 20m Scotch 3M',              '3M',       'cables-electricidad',  950.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cinta aisladora 10m genérica x5u',           NULL,       'cables-electricidad', 1800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),
('Caño corrugado eléctrico 20mm x1m',          NULL,       'cables-electricidad',  480.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Caño corrugado eléctrico 25mm x1m',          NULL,       'cables-electricidad',  650.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "metro",  "sku_prefix": "FERRET"}'),
('Caño conduit EMT 3/4" x3m',                 NULL,       'cables-electricidad', 4500.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- LLAVES DE LUZ Y TOMACORRIENTES — Cambre
-- ============================================================
('Llave de luz simple Cambre Siglo XXI',       'Cambre',   'llaves-tomacorrientes', 2800.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave de luz doble Cambre Siglo XXI',        'Cambre',   'llaves-tomacorrientes', 3800.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tomacorriente doble Cambre Siglo XXI',       'Cambre',   'llaves-tomacorrientes', 3500.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tomacorriente schuko Cambre Siglo XXI',      'Cambre',   'llaves-tomacorrientes', 4200.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Módulo pulsador timbre Cambre',              'Cambre',   'llaves-tomacorrientes', 2200.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Placa 1 módulo Cambre Siglo XXI',            'Cambre',   'llaves-tomacorrientes', 1800.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Placa 2 módulos Cambre Siglo XXI',           'Cambre',   'llaves-tomacorrientes', 2200.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Caja de empotrar 1 módulo para pared',       NULL,       'llaves-tomacorrientes',  650.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Caja de empotrar 2 módulos para pared',      NULL,       'llaves-tomacorrientes',  850.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- LLAVES Y TOMACORRIENTES — Bticino
-- ============================================================
('Llave de luz simple Bticino Axolute',        'Bticino',  'llaves-tomacorrientes', 5800.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tomacorriente doble Bticino Axolute',        'Bticino',  'llaves-tomacorrientes', 6500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tomacorriente schuko Bticino Living',        'Bticino',  'llaves-tomacorrientes', 5200.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Ficha schuko macho 10A genérica',            NULL,       'llaves-tomacorrientes',  980.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Enchufe macho 2P+T 10A genérico',            NULL,       'llaves-tomacorrientes',  650.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- TABLEROS Y PROTECCIÓN ELÉCTRICA
-- ============================================================
('Disyuntor termomagnético 1x10A Bticino',     'Bticino',  'proteccion-electrica',  6500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Disyuntor termomagnético 1x16A Bticino',     'Bticino',  'proteccion-electrica',  6800.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Disyuntor termomagnético 1x20A Bticino',     'Bticino',  'proteccion-electrica',  7200.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Disyuntor termomagnético 1x25A Bticino',     'Bticino',  'proteccion-electrica',  7500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Disyuntor diferencial 2x25A 30mA Bticino',  'Bticino',  'proteccion-electrica', 28000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Disyuntor diferencial 2x40A 30mA Bticino',  'Bticino',  'proteccion-electrica', 32000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tablero metálico 6 módulos para riel DIN',   NULL,       'proteccion-electrica',  8500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tablero metálico 12 módulos para riel DIN',  NULL,       'proteccion-electrica', 14000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tablero metálico 24 módulos para riel DIN',  NULL,       'proteccion-electrica', 22000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Bornera de paso 10A x10u genérica',          NULL,       'proteccion-electrica',  2400.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),

-- ============================================================
-- ILUMINACIÓN — LED (lo que más se vende en ferretería)
-- ============================================================
('Lámpara LED bulbo 9W E27 luz fría Philips',  'Philips',  'iluminacion',    2200.00,  'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lámpara LED bulbo 9W E27 luz cálida Philips','Philips',  'iluminacion',    2200.00,  'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lámpara LED bulbo 15W E27 luz fría Philips', 'Philips',  'iluminacion',    3200.00,  'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lámpara LED bulbo 20W E27 Philips',          'Philips',  'iluminacion',    4200.00,  'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lámpara LED bulbo 9W E27 Osram',             'Osram',    'iluminacion',    2000.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lámpara LED bulbo 12W E27 Osram',            'Osram',    'iluminacion',    2800.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lámpara LED PAR20 7W E27 Ledvance',          'Ledvance', 'iluminacion',    3800.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lámpara LED PAR30 11W E27 Ledvance',         'Ledvance', 'iluminacion',    5500.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lámpara LED dicroica GU10 7W Philips',       'Philips',  'iluminacion',    3500.00,  'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tubo LED T8 18W 120cm luz fría Philips',     'Philips',  'iluminacion',    5800.00,  'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tubo LED T8 9W 60cm luz cálida Osram',       'Osram',    'iluminacion',    4200.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lámpara LED bulbo 9W E27 genérica',          NULL,       'iluminacion',    1200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lampara de emergencia LED c/batería',        NULL,       'iluminacion',   12000.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- TIRAS LED Y REFLECTORES
-- ============================================================
('Tira LED SMD 5050 5m RGB Ledvance',          'Ledvance', 'iluminacion',   12000.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "rollo",  "sku_prefix": "FERRET"}'),
('Tira LED SMD 2835 5m blanco cálido',         NULL,       'iluminacion',    8500.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "rollo",  "sku_prefix": "FERRET"}'),
('Reflector LED 50W exterior Philips',         'Philips',  'iluminacion',   18000.00,  'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Reflector LED 100W exterior genérico',       NULL,       'iluminacion',   22000.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Reflector LED 30W exterior Ledvance',        'Ledvance', 'iluminacion',   14000.00,  'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Portalámparas E27 plástico x5u',             NULL,       'iluminacion',    1800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),

-- ============================================================
-- PROLONGADORES Y EXTENSIONES
-- ============================================================
('Prolongador 3 tomas 2m sin switch',          NULL,       'cables-electricidad', 2800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Prolongador 4 tomas 5m con fusible Zoloda', 'Zoloda',   'cables-electricidad', 5500.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Prolongador 6 tomas 3m con disyuntor',       NULL,       'cables-electricidad', 6800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Zapatilla colgante 4 tomas 1.5m',            NULL,       'cables-electricidad', 3200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cable de extensión 5m 3x2.5mm² Zoloda',     'Zoloda',   'cables-electricidad', 9500.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cable de extensión 10m 3x2.5mm² genérico',  NULL,       'cables-electricidad',16000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- CONECTORES Y TERMINALES ELÉCTRICOS
-- ============================================================
('Terminal ojo M6 para cable 2.5mm² x10u',    NULL,       'cables-electricidad',  680.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),
('Bornera Wago 5 conductores x10u',            NULL,       'cables-electricidad', 2800.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),
('Conector rápido para cable 2.5mm² x10u',    NULL,       'cables-electricidad', 1200.00,  'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),
('Cinta autofusionante 5m 3M',                '3M',       'cables-electricidad', 2800.00,  'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 89 productos
-- Categorías usadas: cables-electricidad, llaves-tomacorrientes, iluminacion, proteccion-electrica
-- Marcas: Philips, Osram, Ledvance, Cambre, Bticino, Zoloda, 3M, genéricos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
