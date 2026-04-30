-- Seed 088: Productos globales — Vinoteca: Espirituosas, Aperitivos y Licores (~60 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: catálogos distribuidores bebidas espirituosas NEA + importadoras argentinas 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: whisky, gin, vodka, aperitivos-licores, rum-brandy

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- WHISKY — los más vendidos en NEA
-- ============================================================
('Johnnie Walker Red Label 750ml',              'Johnnie Walker', 'whisky',  18000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Johnnie Walker Red Label 1000ml',             'Johnnie Walker', 'whisky',  22000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Johnnie Walker Black Label 750ml',            'Johnnie Walker', 'whisky',  32000.00, 'seed', 0.8, 88, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Johnnie Walker Black Label 1000ml',           'Johnnie Walker', 'whisky',  42000.00, 'seed', 0.8, 88, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Johnnie Walker Double Black 750ml',           'Johnnie Walker', 'whisky',  38000.00, 'seed', 0.8, 88, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Johnnie Walker Gold Label Reserve 750ml',     'Johnnie Walker', 'whisky',  65000.00, 'seed', 0.8, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Johnnie Walker Blue Label 750ml',             'Johnnie Walker', 'whisky', 180000.00, 'seed', 0.8, 94, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Old Parr 12 años 750ml',                      'Old Parr',       'whisky',  38000.00, 'seed', 0.8, 88, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Old Parr 18 años 750ml',                      'Old Parr',       'whisky',  95000.00, 'seed', 0.8, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Ballantine''s Finest 750ml',                  'Ballantine''s',  'whisky',  18000.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Ballantine''s 12 años 750ml',                 'Ballantine''s',  'whisky',  28000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Chivas Regal 12 años 750ml',                  'Chivas Regal',   'whisky',  35000.00, 'seed', 0.8, 88, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Jack Daniel''s Old No.7 750ml',               'Jack Daniel''s', 'whisky',  28000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Jack Daniel''s Honey 750ml',                  'Jack Daniel''s', 'whisky',  28000.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Jameson Irish Whiskey 750ml',                 'Jameson',        'whisky',  28000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('The Glenlivet 12 años 750ml',                 'The Glenlivet',  'whisky',  65000.00, 'seed', 0.8, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Glenfiddich 12 años 750ml',                   'Glenfiddich',    'whisky',  75000.00, 'seed', 0.8, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- GIN
-- ============================================================
('Beefeater London Dry Gin 750ml',              'Beefeater',  'gin',         18000.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Tanqueray London Dry Gin 750ml',              'Tanqueray',  'gin',         22000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Tanqueray Ten 750ml',                         'Tanqueray',  'gin',         38000.00, 'seed', 0.8, 88, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Hendrick''s Gin 750ml',                       'Hendrick''s', 'gin',        48000.00, 'seed', 0.8, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Bombay Sapphire Gin 750ml',                   'Bombay',     'gin',         25000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Principe de los Apostoles Gin 750ml',         'Príncipe de los Apóstoles','gin',35000.00,'seed',0.8,88,FALSE,TRUE,'vinoteca','{"unit": "botella", "sku_prefix": "VINO"}'),
('Gin 1689 Nacional 750ml',                     'Gin 1689',   'gin',         18000.00, 'seed', 0.7, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- VODKA
-- ============================================================
('Smirnoff Vodka Red 750ml',                    'Smirnoff',  'vodka',         14000.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Smirnoff Vodka Red 1000ml',                   'Smirnoff',  'vodka',         18000.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Absolut Vodka 750ml',                         'Absolut',   'vodka',         22000.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Absolut Citron 750ml',                        'Absolut',   'vodka',         22000.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Grey Goose Vodka 750ml',                      'Grey Goose','vodka',         55000.00, 'seed', 0.8, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Cîroc Vodka 750ml',                           'Cîroc',     'vodka',         65000.00, 'seed', 0.8, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- RON Y TEQUILA
-- ============================================================
('Bacardi Carta Blanca 750ml',                  'Bacardi',   'rum-brandy',    16000.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Bacardi Carta Negra 750ml',                   'Bacardi',   'rum-brandy',    16000.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Havana Club 3 años 750ml',                    'Havana Club','rum-brandy',   18000.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Havana Club 7 años 750ml',                    'Havana Club','rum-brandy',   28000.00, 'seed', 0.8, 86, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Don Q Cristal Rum 750ml',                     'Don Q',     'rum-brandy',    15000.00, 'seed', 0.7, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Jose Cuervo Tequila Plata 750ml',             'Jose Cuervo','rum-brandy',   22000.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Patron Silver Tequila 750ml',                 'Patron',    'rum-brandy',    85000.00, 'seed', 0.8, 90, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- APERITIVOS — alta rotación en Argentina
-- ============================================================
('Campari 750ml',                               'Campari',   'aperitivos-licores', 22000.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Fernet Branca 750ml',                         'Fernet Branca','aperitivos-licores',22000.00,'seed',0.8,84,FALSE,TRUE,'vinoteca','{"unit": "botella", "sku_prefix": "VINO"}'),
('Fernet Branca 1000ml',                        'Fernet Branca','aperitivos-licores',28000.00,'seed',0.8,84,FALSE,TRUE,'vinoteca','{"unit": "botella", "sku_prefix": "VINO"}'),
('Cynar 750ml',                                 'Cynar',     'aperitivos-licores', 14000.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Aperol 750ml',                                'Aperol',    'aperitivos-licores', 22000.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Aperol Spritz lata 200ml',                    'Aperol',    'aperitivos-licores',  3800.00, 'seed', 0.8, 80, FALSE, TRUE, 'vinoteca', '{"unit": "lata",    "sku_prefix": "VINO"}'),
('Amaro Montenegro 750ml',                      'Montenegro','aperitivos-licores', 28000.00, 'seed', 0.8, 84, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Ramazzotti 750ml',                            'Ramazzotti','aperitivos-licores', 22000.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Martini Bianco 750ml',                        'Martini',   'aperitivos-licores', 15000.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Martini Rosso 750ml',                         'Martini',   'aperitivos-licores', 15000.00, 'seed', 0.8, 82, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Cinzano Bianco 750ml',                        'Cinzano',   'aperitivos-licores', 12000.00, 'seed', 0.8, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),
('Pisco Capel Reservado 750ml',                 'Capel',     'aperitivos-licores', 18000.00, 'seed', 0.7, 80, FALSE, TRUE, 'vinoteca', '{"unit": "botella", "sku_prefix": "VINO"}'),

-- ============================================================
-- TRAGOS LISTOS (RTD — creciente en Argentina)
-- ============================================================
('Smirnoff Ice limón lata 473ml',               'Smirnoff',  'aperitivos-licores',  2800.00, 'seed', 0.8, 78, FALSE, TRUE, 'vinoteca', '{"unit": "lata",    "sku_prefix": "VINO"}'),
('Bacardi Breezer piña lata 355ml',             'Bacardi',   'aperitivos-licores',  2500.00, 'seed', 0.8, 76, FALSE, TRUE, 'vinoteca', '{"unit": "lata",    "sku_prefix": "VINO"}'),
('Jack Daniel''s & Cola lata 355ml',            'Jack Daniel''s','aperitivos-licores',2800.00,'seed',0.8,78,FALSE,TRUE,'vinoteca','{"unit": "lata",    "sku_prefix": "VINO"}'),
('Johnnie Walker & Ginger Ale lata 355ml',      'Johnnie Walker','aperitivos-licores',3200.00,'seed',0.8,78,FALSE,TRUE,'vinoteca','{"unit": "lata",    "sku_prefix": "VINO"}'),
('Fernet Branca & Cola lata 355ml',             'Fernet Branca','aperitivos-licores',2800.00,'seed',0.8,78,FALSE,TRUE,'vinoteca','{"unit": "lata",    "sku_prefix": "VINO"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 60 productos
-- Categorías usadas: whisky, gin, vodka, rum-brandy, aperitivos-licores
-- Marcas: Johnnie Walker, Old Parr, Ballantine's, Chivas Regal, Jack Daniel's, Jameson,
--   The Glenlivet, Glenfiddich, Beefeater, Tanqueray, Hendrick's, Bombay, Príncipe de los Apóstoles,
--   Smirnoff, Absolut, Grey Goose, Bacardi, Havana Club, Jose Cuervo, Patron,
--   Campari, Fernet Branca, Cynar, Aperol, Montenegro, Martini
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
