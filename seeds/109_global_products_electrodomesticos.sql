-- Seed 109: Productos globales — Electrodomésticos (~65 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento locales de electro NEA (Posadas) + distribuidores 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: linea-blanca-pequena, audio-video, climatizacion, iluminacion-led

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- LÍNEA BLANCA PEQUEÑA
-- ============================================================
('Licuadora Atma 500W 3 velocidades',            'Atma',       'linea-blanca-pequena', 32000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Licuadora Oster 500W 12 velocidades',          'Oster',      'linea-blanca-pequena', 45000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Minipimer Oster 350W',                         'Oster',      'linea-blanca-pequena', 28000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Minipimer Atma 350W con accesorios',           'Atma',       'linea-blanca-pequena', 25000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Tostadora 2 ranuras Atma',                     'Atma',       'linea-blanca-pequena', 22000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Sandwichera / waflera Atma 750W',              'Atma',       'linea-blanca-pequena', 28000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Freidora de aire Atma 3.5L 1500W',             'Atma',       'linea-blanca-pequena', 65000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Freidora de aire Atma 5L 1700W',               'Atma',       'linea-blanca-pequena', 85000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Cafetera goteo Oster 12 tazas 900W',           'Oster',      'linea-blanca-pequena', 35000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Cafetera goteo Atma 12 tazas 800W',            'Atma',       'linea-blanca-pequena', 28000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Cafetera Nespresso Essenza Mini cápsulas',     'Nespresso',  'linea-blanca-pequena',120000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Plancha de ropa vapor Philips 2400W',          'Philips',    'linea-blanca-pequena', 48000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Plancha de ropa vapor Oster 2000W',            'Oster',      'linea-blanca-pequena', 38000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Aspiradora ciclónica Atma 1200W',              'Atma',       'linea-blanca-pequena', 65000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Aspiradora sin cable Atma 21.6V',              'Atma',       'linea-blanca-pequena', 85000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Procesadora de alimentos Atma 900W',           'Atma',       'linea-blanca-pequena', 42000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Microondas Atma 20L 700W',                     'Atma',       'linea-blanca-pequena', 78000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),

-- ============================================================
-- AUDIO Y VIDEO
-- ============================================================
('Parlante Bluetooth JBL Flip Essential 20W',   'JBL',        'audio-video',       55000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Parlante Bluetooth JBL Charge 5 30W',         'JBL',        'audio-video',       95000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Parlante Bluetooth Sony SRS-XB23',            'Sony',       'audio-video',       65000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Auriculares Bluetooth Sony WH-CH520',         'Sony',       'audio-video',       55000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Auriculares Bluetooth JBL Tune 520BT',        'JBL',        'audio-video',       48000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Auriculares in-ear JBL Tune 130NC TWS',       'JBL',        'audio-video',       45000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Smart TV 32" HD Samsung genérico',            'Samsung',    'audio-video',      280000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Smart TV 43" FHD LG genérico',               'LG',         'audio-video',      420000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Decodificador TDA Zinex HD',                  'Zinex',      'audio-video',       22000.00, 'seed', 0.8, 80, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Antena digital interior UHF/VHF',             NULL,         'audio-video',        8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Teléfono fijo inalámbrico Motorola',          'Motorola',   'audio-video',       35000.00, 'seed', 0.8, 80, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),

-- ============================================================
-- CLIMATIZACIÓN — importante en NEA (clima subtropical, veranos muy calurosos)
-- ============================================================
('Ventilador de pie Liliana 16" 3 velocidades', 'Liliana',    'climatizacion',     28000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Ventilador de pie Liliana 18" 3 velocidades', 'Liliana',    'climatizacion',     35000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Ventilador de mesa 12" Liliana',              'Liliana',    'climatizacion',     18000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Ventilador de techo Peabody 52" 5 aspas',     'Peabody',    'climatizacion',     55000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Ventilador de torre Liliana 45cm oscilante',  'Liliana',    'climatizacion',     42000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Caloventor eléctrico 2000W Liliana',          'Liliana',    'climatizacion',     32000.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Estufa halógena cuarzo 1500W',               NULL,          'climatizacion',     22000.00, 'seed', 0.7, 60, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Estufa a gas infrarrojo 3500 cal Eskabe',     'Eskabe',     'climatizacion',     85000.00, 'seed', 0.8, 80, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Split frío/calor Gree 2200fg 220V',           'Gree',       'climatizacion',    380000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Split frío/calor Gree 3200fg 220V',           'Gree',       'climatizacion',    480000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Split frío/calor BGH 2200fg inverter',        'BGH',        'climatizacion',    420000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Split frío/calor BGH 3200fg inverter',        'BGH',        'climatizacion',    520000.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),

-- ============================================================
-- ILUMINACIÓN LED
-- ============================================================
('Lamparita LED 9W E27 luz fría Osram',         'Osram',      'iluminacion-led',    3200.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Lamparita LED 12W E27 luz fría Osram',        'Osram',      'iluminacion-led',    3800.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Lamparita LED 9W E27 luz cálida Philips',     'Philips',    'iluminacion-led',    3500.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Lamparita LED 12W E27 luz cálida Philips',    'Philips',    'iluminacion-led',    4200.00, 'seed', 0.9, 85, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Lamparita LED dicroica 7W GU10 Osram',        'Osram',      'iluminacion-led',    4500.00, 'seed', 0.9, 82, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Tira LED 5m 12V RGB multicolor',             NULL,          'iluminacion-led',   18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'electrodomesticos', '{"unit": "rollo",  "sku_prefix": "ELEC"}'),
('Tira LED 5m 220V blanca cálida',             NULL,          'iluminacion-led',   14000.00, 'seed', 0.7, 60, FALSE, TRUE, 'electrodomesticos', '{"unit": "rollo",  "sku_prefix": "ELEC"}'),
('Reflector LED exterior 20W frío',            NULL,          'iluminacion-led',   14000.00, 'seed', 0.7, 60, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'),
('Reflector LED exterior 50W frío',            NULL,          'iluminacion-led',   24000.00, 'seed', 0.7, 60, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 64 productos
-- Categorías: linea-blanca-pequena, audio-video, climatizacion, iluminacion-led
-- Marcas: Atma, Oster, Philips, Nespresso, JBL, Sony, Samsung, LG, Zinex, Liliana, Peabody,
--   Eskabe, Gree, BGH, Osram, Motorola
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
-- NOTA: climatización tiene altísima rotación en NEA — veranos extremos (40°C+), temporada oct-mar
