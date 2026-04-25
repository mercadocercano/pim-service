-- Seed 100: Productos globales — Panadería: Insumos (~65 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento panaderías NEA (Posadas) + distribuidores insumos 2026
-- ZONA: Posadas, Misiones (NEA) — precios mayoristas estimados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: harinas-panaderia, levaduras, margarinas-industriales, dulce-leche-reposteria, rellenos-coberturas, esencias-aditivos, papeles-moldes

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- HARINAS — uso panadería
-- ============================================================
('Harina 0000 25kg Cañuelas',                    'Cañuelas',      'harinas-panaderia', 22000.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "bolsa", "sku_prefix": "PANA"}'),
('Harina 0000 50kg Cañuelas',                    'Cañuelas',      'harinas-panaderia', 42000.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "bolsa", "sku_prefix": "PANA"}'),
('Harina 000 25kg Cañuelas',                     'Cañuelas',      'harinas-panaderia', 20000.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "bolsa", "sku_prefix": "PANA"}'),
('Harina 0000 25kg Pureza',                      'Pureza',        'harinas-panaderia', 21000.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "bolsa", "sku_prefix": "PANA"}'),
('Harina 0000 25kg Morixe',                      'Morixe',        'harinas-panaderia', 20500.00, 'seed', 0.8, 82, FALSE, TRUE, 'panaderia', '{"unit": "bolsa", "sku_prefix": "PANA"}'),
('Harina integral 25kg Cañuelas',                'Cañuelas',      'harinas-panaderia', 24000.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "bolsa", "sku_prefix": "PANA"}'),
('Harina de maíz amarilla 25kg genérica',        NULL,            'harinas-panaderia', 14000.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "bolsa", "sku_prefix": "PANA"}'),
('Sémola fina 25kg genérica',                    NULL,            'harinas-panaderia', 16000.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "bolsa", "sku_prefix": "PANA"}'),
('Salvado de trigo fino 10kg genérico',          NULL,            'harinas-panaderia',  8000.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "bolsa", "sku_prefix": "PANA"}'),

-- ============================================================
-- LEVADURAS
-- ============================================================
('Levadura seca activa 500g Fleischmann',        'Fleischmann',   'levaduras',      8500.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Levadura seca activa 100g Fleischmann',        'Fleischmann',   'levaduras',      2200.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Levadura fresca 500g Fleischmann',             'Fleischmann',   'levaduras',      5500.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Levadura fresca 500g Calsa',                   'Calsa',         'levaduras',      5200.00, 'seed', 0.8, 82, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Levadura instantánea 100g Dr. Oetker',         'Dr. Oetker',    'levaduras',      2800.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Polvo de hornear 100g Royal',                  'Royal',         'levaduras',      2500.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Polvo de hornear 100g Dr. Oetker',             'Dr. Oetker',    'levaduras',      2800.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Bicarbonato de sodio 500g La Serenísima',      'La Serenísima', 'levaduras',      1800.00, 'seed', 0.9, 80, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),

-- ============================================================
-- MARGARINAS INDUSTRIALES
-- ============================================================
('Margarina industrial Primavera 1kg bloque',    'Primavera',     'margarinas-industriales', 4800.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Margarina industrial Primavera 5kg bloque',    'Primavera',     'margarinas-industriales',22000.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "caja",  "sku_prefix": "PANA"}'),
('Margarina Cocinero bloque 1kg',                'Cocinero',      'margarinas-industriales', 5200.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Margarina Mambotá blanda 1kg balde',           'Mambotá',       'margarinas-industriales', 4500.00, 'seed', 0.8, 80, FALSE, TRUE, 'panaderia', '{"unit": "balde", "sku_prefix": "PANA"}'),
('Grasa vegetal 1kg genérica',                   NULL,            'margarinas-industriales', 3800.00, 'seed', 0.6, 60, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Manteca sin sal 200g La Serenísima',           'La Serenísima', 'margarinas-industriales', 2800.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),

-- ============================================================
-- DULCE DE LECHE INDUSTRIAL Y REPOSTERÍA
-- ============================================================
('Dulce de leche repostero 1kg La Serenísima',   'La Serenísima', 'dulce-leche-reposteria', 6500.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Dulce de leche repostero 5kg La Serenísima',   'La Serenísima', 'dulce-leche-reposteria',30000.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "balde", "sku_prefix": "PANA"}'),
('Dulce de leche repostero 1kg SanCor',          'SanCor',        'dulce-leche-reposteria', 6200.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Dulce de leche repostero 1kg Mastellone',      'Mastellone',    'dulce-leche-reposteria', 5800.00, 'seed', 0.8, 82, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Crema pastelera lista 1kg Puratos',            'Puratos',       'dulce-leche-reposteria', 5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Mermelada industrial frutilla 1kg Arcor',      'Arcor',         'dulce-leche-reposteria', 4800.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Mermelada industrial durazno 1kg Arcor',       'Arcor',         'dulce-leche-reposteria', 4800.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Mermelada industrial damasco 3kg La Campagnola','La Campagnola', 'dulce-leche-reposteria',12000.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "lata",  "sku_prefix": "PANA"}'),

-- ============================================================
-- COBERTURAS DE CHOCOLATE
-- ============================================================
('Cobertura chocolate semiamargo 1kg Fenix',     'Fenix',         'rellenos-coberturas',  8500.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Cobertura chocolate con leche 1kg Fenix',      'Fenix',         'rellenos-coberturas',  8800.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Cobertura chocolate blanco 1kg Fenix',         'Fenix',         'rellenos-coberturas',  9200.00, 'seed', 0.9, 85, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Cobertura chocolate 1kg Molsa',                'Molsa',         'rellenos-coberturas',  7800.00, 'seed', 0.8, 82, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Cobertura chocolate 1kg Famosa',               'Famosa',        'rellenos-coberturas',  7500.00, 'seed', 0.8, 80, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Baño de repostería blanco 1kg genérico',       NULL,            'rellenos-coberturas',  5500.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Chips de chocolate 500g Fenix',                'Fenix',         'rellenos-coberturas',  5200.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Cacao amargo en polvo 200g genérico',          NULL,            'rellenos-coberturas',  2800.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),

-- ============================================================
-- ESENCIAS Y ADITIVOS
-- ============================================================
('Esencia de vainilla 60ml Georgalos',           'Georgalos',     'esencias-aditivos',  1800.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Esencia de vainilla 60ml Fleischmann',         'Fleischmann',   'esencias-aditivos',  1950.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Esencia de limón 60ml Georgalos',              'Georgalos',     'esencias-aditivos',  1800.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Esencia de naranja 60ml Georgalos',            'Georgalos',     'esencias-aditivos',  1800.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Colorante alimentario rojo 30ml genérico',     NULL,            'esencias-aditivos',  1200.00, 'seed', 0.6, 60, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Colorante alimentario amarillo 30ml genérico', NULL,            'esencias-aditivos',  1200.00, 'seed', 0.6, 60, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Cremor tártaro 100g genérico',                 NULL,            'esencias-aditivos',  2200.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Azúcar impalpable 500g Ledesma',               'Ledesma',       'esencias-aditivos',  2200.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Azúcar rubia 1kg Ledesma',                     'Ledesma',       'esencias-aditivos',  1800.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Glucosa 1kg genérica',                         NULL,            'esencias-aditivos',  3800.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "kg",    "sku_prefix": "PANA"}'),
('Gelatina sin sabor 7g Royal x caja 12u',       'Royal',         'esencias-aditivos',  3600.00, 'seed', 0.9, 82, FALSE, TRUE, 'panaderia', '{"unit": "caja",  "sku_prefix": "PANA"}'),

-- ============================================================
-- PAPELES Y MOLDES
-- ============================================================
('Papel manteca 25m rollo',                      NULL,            'papeles-moldes',     3200.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "rollo", "sku_prefix": "PANA"}'),
('Papel film transparente 30cmx50m',             NULL,            'papeles-moldes',     3500.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "rollo", "sku_prefix": "PANA"}'),
('Papel de aluminio 30cmx25m',                   NULL,            'papeles-moldes',     4200.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "rollo", "sku_prefix": "PANA"}'),
('Molde rectangular para budín 25cm genérico',   NULL,            'papeles-moldes',     2200.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Molde de torta redondo 26cm aluminio',         NULL,            'papeles-moldes',     3800.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Pirotines de papel n°5 x100u',                 NULL,            'papeles-moldes',     2500.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "pack",   "sku_prefix": "PANA"}'),
('Pirotines de papel n°8 x100u',                 NULL,            'papeles-moldes',     3200.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "pack",   "sku_prefix": "PANA"}'),
('Manga de repostería desechable x10u',          NULL,            'papeles-moldes',     2800.00, 'seed', 0.7, 60, FALSE, TRUE, 'panaderia', '{"unit": "pack",   "sku_prefix": "PANA"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 64 productos
-- Categorías: harinas-panaderia, levaduras, margarinas-industriales, dulce-leche-reposteria, rellenos-coberturas, esencias-aditivos, papeles-moldes
-- Marcas: Cañuelas, Pureza, Morixe, Fleischmann, Calsa, Dr. Oetker, Royal, Primavera, Cocinero, Mambotá,
--   La Serenísima, SanCor, Mastellone, Puratos, Arcor, La Campagnola, Fenix, Molsa, Famosa, Georgalos, Ledesma
-- Precios referencia zona Posadas, NEA — abril 2026 (mayorista/industrial estimado)
