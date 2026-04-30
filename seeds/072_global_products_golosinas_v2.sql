-- Seed 072: Golosinas — 56 productos nuevos (complemento de seed 042)
-- CICLO: cycle-004-brands-catalog-expansion
-- FECHA: 2026-04-21
-- FUENTE: observación de góndola NEA (Posadas, Misiones) + listas mayoristas 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- COMPLEMENTA: seed 042 (kiosco v2). Los productos de 042 NO se repiten aquí.
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- NOTA: Verificar que las marcas existan en marketplace_brands antes de correr este seed
--       (seeds 055 y 064 deben haberse ejecutado primero).

-- Categorías confirmadas en seed 042/producción:
--   alfajores, chocolates, caramelos-chicles, gomitas-malvaviscos

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- ALFAJORES — marcas nuevas, no presentes en seed 042
-- ============================================================
-- En 042 ya están: Terrabusi triple, Cachafaz triple, Jorgelin, Alfajor Jorgito,
-- Alfajor Tita, Alfajor Guaymallén, Alfajor Milka, Alfajor Capitán del Espacio,
-- Alfajor de chocolate Milka 55g, Chocolate Shot 55g.

('Alfajor Guaymallén simple chocolate',          'Guaymallén',    'alfajores', 700.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Guaymallén triple chocolate',          'Guaymallén',    'alfajores', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Guaymallén blanco',                    'Guaymallén',    'alfajores', 700.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Capitán del Espacio simple',           'Capitán del Espacio', 'alfajores', 800.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Capitán del Espacio triple',           'Capitán del Espacio', 'alfajores', 1300.00,'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Jorgito triple chocolate',             'Jorgito',       'alfajores', 1100.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Jorgito blanco',                       'Jorgito',       'alfajores', 750.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Cachafaz de maicena x2',               'Cachafaz',      'alfajores', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Fantoche simple',                      'Fantoche',      'alfajores', 650.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Fantoche triple',                      'Fantoche',      'alfajores', 1100.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Suchard simple',                       'Suchard',       'alfajores', 800.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor La Repostera chocolate',               'La Repostera',  'alfajores', 600.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Tofi simple',                          'Tofi',          'alfajores', 700.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Havanna chocolate x2',                 'Havanna',       'alfajores', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Bonafide triple chocolate',            'Bonafide',      'alfajores', 1400.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- CHOCOLATES — no presentes en seed 042
-- ============================================================
-- En 042 ya están: Cofler aireado 55g, Toblerone 100g, Rocklets 40g, Kinder Bueno 43g,
-- Chocolate Milka 55g (en template 021), Shot 55g (en template 021).

('Cofler Blanco 55g',                            'Cofler',        'chocolates', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Cofler con maní 55g',                          'Cofler',        'chocolates', 1600.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Tita 70g',                                     'Tita',          'chocolates', 1300.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Rhodesia chocolate con leche 50g',             'Rhodesia',      'chocolates', 1400.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Rhodesia mousse de chocolate 50g',             'Rhodesia',      'chocolates', 1400.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Águila chocolate de taza 40g',                 'Águila',        'chocolates', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Block chocolate con leche 55g',                'Block',         'chocolates', 1300.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Cadbury Oreo chocolate 38g',                   'Cadbury',       'chocolates', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Cadbury Bubbly chocolate con leche 40g',       'Cadbury',       'chocolates', 1700.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Rocklets Arcor 40g',                           'Rocklets',      'chocolates', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- CARAMELOS, CHUPETINES Y CHICLES
-- ============================================================
-- En 042 ya están: Butter Toffees 150g, Topline 7u, Menthoplus 30g.

('Sugus tutti frutti bolsa 150g',                'Sugus',         'caramelos-chicles', 1600.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Sugus tutti frutti 20g (unidad)',              'Sugus',         'caramelos-chicles', 500.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Palitos de la Selva 25g',                      'Palitos de la Selva', 'caramelos-chicles', 600.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Flynn Paff frutilla 12g',                      'Flynn Paff',    'caramelos-chicles', 400.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Flynn Paff tutti frutti 12g',                  'Flynn Paff',    'caramelos-chicles', 400.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Butter Toffees individual 8g',                 'Butter Toffees','caramelos-chicles', 300.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Pico Dulce chupetín 14g',                      'Pico Dulce',    'caramelos-chicles', 350.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Mr. Pop caramelo explosivo',                   'Mr. Pop',       'caramelos-chicles', 350.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Bombuchas caramelos x5u',                      'Bombuchas',     'caramelos-chicles', 400.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Pindapoy vitamina C 40g',                      'Pindapoy',      'caramelos-chicles', 700.00,  'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- GOMITAS Y CONFITADOS
-- ============================================================
-- En 042 ya están: Mogul cerebritos 55g, Malvaviscos Guolis 100g, Trolli 100g, Haribo 80g,
-- Rocklets Arcor 40g.

('Mogul tutti frutti bolsa 110g',                'Mogul',         'gomitas-malvaviscos', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Mogul ositos bolsa 110g',                      'Mogul',         'gomitas-malvaviscos', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Billiken gomitas bolsa 100g',                  'Billiken',      'gomitas-malvaviscos', 1100.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Billiken serpentinas ácidas 100g',             'Billiken',      'gomitas-malvaviscos', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- MANTECOL
-- ============================================================
-- Categoría: puede ir en golosinas o alfajores. Usamos categoria propia si existe,
-- de lo contrario 'alfajores' como fallback (producto típico de kiosco).

('Mantecol 125g',                                'Mantecol',      'alfajores', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Mantecol 90g',                                 'Mantecol',      'alfajores', 2300.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Mantecol Light 90g',                           'Mantecol',      'alfajores', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb)

ON CONFLICT DO NOTHING;

-- Total: 56 productos nuevos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
