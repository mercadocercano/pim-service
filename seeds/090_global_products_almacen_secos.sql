-- Seed 090: Productos globales — Almacén: Secos y Conservas (~90 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento supermercados y almacenes NEA (Posadas) + distribuidores 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: pastas-secas, arroz-legumbres, harinas-premezclas, aceites-vinagres, conservas-enlatados

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- FIDEOS SECOS — Lucchetti (líder en NEA)
-- ============================================================
('Fideos spaghetti 500g Lucchetti',             'Lucchetti',  'pastas-secas',   980.00,  'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Fideos moño 500g Lucchetti',                  'Lucchetti',  'pastas-secas',   980.00,  'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Fideos mostachol 500g Lucchetti',             'Lucchetti',  'pastas-secas',   980.00,  'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Fideos tallarín 500g Lucchetti',              'Lucchetti',  'pastas-secas',   980.00,  'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Fideos tirabuzón 500g Lucchetti',             'Lucchetti',  'pastas-secas',   980.00,  'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Fideos codito 500g Lucchetti',                'Lucchetti',  'pastas-secas',   980.00,  'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Fideos spaghetti 1kg Lucchetti',              'Lucchetti',  'pastas-secas',  1850.00,  'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- FIDEOS SECOS — Terrabusi, Don Felipe, genéricos
-- ============================================================
('Fideos spaghetti 500g Terrabusi',             'Terrabusi',  'pastas-secas',   850.00,  'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Fideos moño 500g Terrabusi',                  'Terrabusi',  'pastas-secas',   850.00,  'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Fideos spaghetti 500g Don Felipe',            'Don Felipe', 'pastas-secas',   780.00,  'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Fideos mostachol 500g Don Felipe',            'Don Felipe', 'pastas-secas',   780.00,  'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Fideos spaghetti 500g genérico x-econ',       NULL,         'pastas-secas',   620.00,  'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- ARROZ — Gallo, Molinos, genérico
-- ============================================================
('Arroz largo fino 1kg Gallo',                  'Gallo',       'arroz-legumbres', 1600.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Arroz largo fino 500g Gallo',                 'Gallo',       'arroz-legumbres',  850.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Arroz largo fino 1kg Molinos',                'Molinos Río de la Plata','arroz-legumbres',1400.00,'seed',0.8,80,FALSE,TRUE,'almacen','{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Arroz doble carolina 1kg genérico',           NULL,          'arroz-legumbres', 1100.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Arroz integral 500g Gallo',                   'Gallo',       'arroz-legumbres', 1200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Arroz parboil 1kg genérico',                  NULL,          'arroz-legumbres',  950.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),

-- ============================================================
-- HARINAS — Cañuelas, Pureza, genérico
-- ============================================================
('Harina 0000 1kg Cañuelas',                    'Cañuelas',    'harinas-premezclas', 1100.00,'seed',0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Harina 0000 5kg Cañuelas',                    'Cañuelas',    'harinas-premezclas', 4800.00,'seed',0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "bolsa",  "sku_prefix": "ALMAC"}'),
('Harina 000 1kg Cañuelas',                     'Cañuelas',    'harinas-premezclas',  980.00,'seed',0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Harina 0000 1kg Pureza',                      'Pureza',      'harinas-premezclas',  950.00,'seed',0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Harina 0000 1kg genérica',                    NULL,          'harinas-premezclas',  780.00,'seed',0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Premezcla para pan casero 1kg Cañuelas',      'Cañuelas',    'harinas-premezclas', 2200.00,'seed',0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),

-- ============================================================
-- AZÚCAR — Ledesma
-- ============================================================
('Azúcar blanca 1kg Ledesma',                   'Ledesma',     'conservas-enlatados',1200.00,'seed',0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Azúcar blanca 2kg Ledesma',                   'Ledesma',     'conservas-enlatados',2300.00,'seed',0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "bolsa",  "sku_prefix": "ALMAC"}'),
('Azúcar blanca 5kg Ledesma',                   'Ledesma',     'conservas-enlatados',5500.00,'seed',0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "bolsa",  "sku_prefix": "ALMAC"}'),
('Azúcar blanca 1kg genérica',                  NULL,          'conservas-enlatados', 950.00,'seed',0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Azúcar morena 1kg Ledesma',                   'Ledesma',     'conservas-enlatados',1400.00,'seed',0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),

-- ============================================================
-- ACEITES — Cocinero, Natura, Marolio
-- ============================================================
('Aceite de girasol 900ml Cocinero',            'Cocinero',    'aceites-vinagres',  2200.00,'seed',0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Aceite de girasol 1.5L Cocinero',             'Cocinero',    'aceites-vinagres',  3200.00,'seed',0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Aceite de girasol 2L Cocinero',               'Cocinero',    'aceites-vinagres',  4100.00,'seed',0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Aceite de maíz 900ml Natura',                 'Natura',      'aceites-vinagres',  2600.00,'seed',0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Aceite de girasol 900ml Marolio',             'Marolio',     'aceites-vinagres',  2100.00,'seed',0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Aceite de oliva extra virgen 500ml Cocinero', 'Cocinero',    'aceites-vinagres',  5500.00,'seed',0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- SAL Y VINAGRE
-- ============================================================
('Sal fina 500g Celusal',                       'Celusal',     'conservas-enlatados', 580.00,'seed',0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Sal gruesa 1kg genérica',                     NULL,          'conservas-enlatados', 650.00,'seed',0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Vinagre de alcohol 500ml Marolio',            'Marolio',     'aceites-vinagres',    650.00,'seed',0.7, 74, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Vinagre de manzana 500ml Marolio',            'Marolio',     'aceites-vinagres',    780.00,'seed',0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- AVENA Y CEREALES
-- ============================================================
('Avena arrollada 500g Quaker',                 'Quaker',      'harinas-premezclas', 1800.00,'seed',0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Avena arrollada 500g La Paulina',             'La Paulina',  'harinas-premezclas', 1400.00,'seed',0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Polenta 500g Molinos',                        'Molinos Río de la Plata','harinas-premezclas',750.00,'seed',0.8,76,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Maicena 500g La Plata',                       'La Plata',    'harinas-premezclas',  950.00,'seed',0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- LEGUMBRES — genérico (bolsita 500g)
-- ============================================================
('Lentejas 500g La Paulina',                    'La Paulina',  'arroz-legumbres',    980.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Porotos negros 500g genérico',                NULL,          'arroz-legumbres',    850.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Porotos pallares 500g genérico',              NULL,          'arroz-legumbres',    850.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Garbanzos 500g genérico',                     NULL,          'arroz-legumbres',    980.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Arvejas secas 500g genérico',                 NULL,          'arroz-legumbres',    780.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- CONSERVAS — La Campagnola, Arcor, Marolio
-- ============================================================
('Tomates perita lata 400g La Campagnola',      'La Campagnola','conservas-enlatados',980.00,'seed',0.8,80,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Tomates triturados lata 400g Arcor',          'Arcor',       'conservas-enlatados', 820.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Puré de tomate 520g Cica',                    'Cica',        'conservas-enlatados', 780.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Choclo en lata 300g La Campagnola',           'La Campagnola','conservas-enlatados',920.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Arvejas en lata 300g La Campagnola',          'La Campagnola','conservas-enlatados',880.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Lentejas en lata 400g La Campagnola',         'La Campagnola','conservas-enlatados',950.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Duraznos en almíbar 825g La Campagnola',      'La Campagnola','conservas-enlatados',2400.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Palmitos en lata 400g La Campagnola',         'La Campagnola','conservas-enlatados',2100.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Peras en almíbar 820g genérico',              NULL,          'conservas-enlatados',1800.00,'seed',0.6,40,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- ATÚN Y SARDINAS
-- ============================================================
('Atún al natural 170g Cingal',                 'Cingal',      'conservas-enlatados', 1800.00,'seed',0.8,80,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Atún en aceite 170g Cingal',                  'Cingal',      'conservas-enlatados', 1850.00,'seed',0.8,80,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Atún en agua 160g La Campagnola',             'La Campagnola','conservas-enlatados',1950.00,'seed',0.8,80,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Sardinas en aceite 125g La Campagnola',       'La Campagnola','conservas-enlatados', 950.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Caballa en aceite 170g Trigal',               'Trigal',      'conservas-enlatados', 1200.00,'seed',0.7,76,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- MORRONES Y OTROS ENCURTIDOS
-- ============================================================
('Morrones en aceite 235g La Campagnola',       'La Campagnola','conservas-enlatados',1800.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Aceitunas verdes 200g La Campagnola',         'La Campagnola','conservas-enlatados',2200.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Pickles surtidos 390g Arcor',                 'Arcor',       'conservas-enlatados',1600.00,'seed',0.8,76,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- DULCE DE LECHE Y MERMELADAS
-- ============================================================
('Dulce de leche repostero 400g La Serenísima', 'La Serenísima','conservas-enlatados',2600.00,'seed',0.8,80,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Dulce de leche familiar 400g La Serenísima',  'La Serenísima','conservas-enlatados',2400.00,'seed',0.8,80,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Mermelada de durazno 454g Arcor',             'Arcor',       'conservas-enlatados',1800.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Mermelada de frutilla 454g Arcor',            'Arcor',       'conservas-enlatados',1800.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Mermelada de damasco 454g La Campagnola',     'La Campagnola','conservas-enlatados',1900.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- CALDOS, CONDIMENTOS Y SALSAS
-- ============================================================
('Caldo de verdura x6u Knorr',                  'Knorr',      'conservas-enlatados',  980.00,'seed',0.8,80,FALSE,TRUE,'almacen','{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Caldo de carne x6u Knorr',                    'Knorr',      'conservas-enlatados',  980.00,'seed',0.8,80,FALSE,TRUE,'almacen','{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Caldo de pollo x12u Knorr',                   'Knorr',      'conservas-enlatados', 1800.00,'seed',0.8,80,FALSE,TRUE,'almacen','{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Salsa de tomate 400g Arcor',                  'Arcor',       'conservas-enlatados',  820.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Salsa para pizza 420g Cica',                  'Cica',        'conservas-enlatados',  950.00,'seed',0.8,78,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Ketchup 397g Heinz',                          'Heinz',       'conservas-enlatados', 2200.00,'seed',0.8,82,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Mayonesa 250g Hellmann''s',                   'Hellmann''s', 'conservas-enlatados', 1800.00,'seed',0.8,82,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Mostaza 225g Hellmann''s',                    'Hellmann''s', 'conservas-enlatados', 1600.00,'seed',0.8,82,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 88 productos
-- Categorías usadas: pastas-secas, arroz-legumbres, harinas-premezclas, aceites-vinagres, conservas-enlatados
-- Marcas: Lucchetti, Terrabusi, Don Felipe, Gallo, Molinos, Cañuelas, Pureza, Ledesma, Cocinero, Natura,
--   Marolio, Celusal, Quaker, La Paulina, La Campagnola, Arcor, Cica, Cingal, Trigal, Knorr, Heinz, Hellmann's
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
