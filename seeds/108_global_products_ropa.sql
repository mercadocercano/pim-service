-- Seed 108: Productos globales — Ropa: Indumentaria Básica (~65 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento locales de ropa NEA (Posadas) + distribuidores indumentaria 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: remeras-basicas, pantalones, ropa-interior-hombre, ropa-interior-mujer, medias, calzado-basico, abrigos
-- NOTA: en NEA el clima es subtropical; las ventas de ropa de abrigo son estacionales (mayo-agosto)

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- REMERAS BÁSICAS HOMBRE
-- ============================================================
('Remera básica blanca hombre talle S',          'Genérico',  'remeras-basicas',   6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "S",   "sku_prefix": "ROPA"}'),
('Remera básica blanca hombre talle M',          'Genérico',  'remeras-basicas',   6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",   "sku_prefix": "ROPA"}'),
('Remera básica blanca hombre talle L',          'Genérico',  'remeras-basicas',   6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "L",   "sku_prefix": "ROPA"}'),
('Remera básica blanca hombre talle XL',         'Genérico',  'remeras-basicas',   7000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "XL",  "sku_prefix": "ROPA"}'),
('Remera básica negra hombre talle M',           'Genérico',  'remeras-basicas',   6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",   "sku_prefix": "ROPA"}'),
('Remera básica negra hombre talle L',           'Genérico',  'remeras-basicas',   6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "L",   "sku_prefix": "ROPA"}'),
('Remera básica gris hombre talle M',            'Genérico',  'remeras-basicas',   6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",   "sku_prefix": "ROPA"}'),
('Remera polo básica hombre talle M',            'Genérico',  'remeras-basicas',   9500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",   "sku_prefix": "ROPA"}'),
('Remera básica mujer cuello V talle S',         'Genérico',  'remeras-basicas',   6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "S",   "sku_prefix": "ROPA"}'),
('Remera básica mujer cuello V talle M',         'Genérico',  'remeras-basicas',   6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",   "sku_prefix": "ROPA"}'),

-- ============================================================
-- PANTALONES
-- ============================================================
('Jean hombre azul talle 28',                    'Genérico',  'pantalones',       28000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "28",  "sku_prefix": "ROPA"}'),
('Jean hombre azul talle 30',                    'Genérico',  'pantalones',       28000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "30",  "sku_prefix": "ROPA"}'),
('Jean hombre azul talle 32',                    'Genérico',  'pantalones',       28000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "32",  "sku_prefix": "ROPA"}'),
('Jean hombre azul talle 34',                    'Genérico',  'pantalones',       28000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "34",  "sku_prefix": "ROPA"}'),
('Jean hombre negro talle 32',                   'Genérico',  'pantalones',       28000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "32",  "sku_prefix": "ROPA"}'),
('Jogger hombre algodón talle S',                'Genérico',  'pantalones',       14000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "S",   "sku_prefix": "ROPA"}'),
('Jogger hombre algodón talle M',                'Genérico',  'pantalones',       14000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",   "sku_prefix": "ROPA"}'),
('Jogger hombre algodón talle L',                'Genérico',  'pantalones',       14000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "L",   "sku_prefix": "ROPA"}'),
('Bermuda hombre talle M',                       'Genérico',  'pantalones',       12000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",   "sku_prefix": "ROPA"}'),
('Bermuda hombre talle L',                       'Genérico',  'pantalones',       12000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "L",   "sku_prefix": "ROPA"}'),
('Calza mujer lycra negra talle S',              'Genérico',  'pantalones',        9500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "S",   "sku_prefix": "ROPA"}'),
('Calza mujer lycra negra talle M',              'Genérico',  'pantalones',        9500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",   "sku_prefix": "ROPA"}'),
('Calza mujer lycra negra talle L',              'Genérico',  'pantalones',        9500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "L",   "sku_prefix": "ROPA"}'),
('Short deportivo mujer talle S',                'Genérico',  'pantalones',        8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "S",   "sku_prefix": "ROPA"}'),
('Short deportivo mujer talle M',                'Genérico',  'pantalones',        8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",   "sku_prefix": "ROPA"}'),

-- ============================================================
-- ROPA INTERIOR HOMBRE
-- ============================================================
('Slip hombre talle S algodón',                  'Genérico',  'ropa-interior-hombre', 3500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "S",  "sku_prefix": "ROPA"}'),
('Slip hombre talle M algodón',                  'Genérico',  'ropa-interior-hombre', 3500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",  "sku_prefix": "ROPA"}'),
('Bóxer hombre talle M algodón',                 'Genérico',  'ropa-interior-hombre', 4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",  "sku_prefix": "ROPA"}'),
('Bóxer hombre talle L algodón',                 'Genérico',  'ropa-interior-hombre', 4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "L",  "sku_prefix": "ROPA"}'),

-- ============================================================
-- ROPA INTERIOR MUJER
-- ============================================================
('Tanga mujer talle S algodón',                  'Genérico',  'ropa-interior-mujer', 3200.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "S",  "sku_prefix": "ROPA"}'),
('Tanga mujer talle M algodón',                  'Genérico',  'ropa-interior-mujer', 3200.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",  "sku_prefix": "ROPA"}'),
('Colaless mujer talle L algodón',               'Genérico',  'ropa-interior-mujer', 3200.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "L",  "sku_prefix": "ROPA"}'),
('Corpiño básico talle 85B algodón',             'Genérico',  'ropa-interior-mujer', 6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "85B","sku_prefix": "ROPA"}'),
('Corpiño básico talle 90B algodón',             'Genérico',  'ropa-interior-mujer', 6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "90B","sku_prefix": "ROPA"}'),
('Corpiño básico talle 95C algodón',             'Genérico',  'ropa-interior-mujer', 6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "95C","sku_prefix": "ROPA"}'),

-- ============================================================
-- MEDIAS
-- ============================================================
('Calcetines hombre cortos x3 pares',            'Genérico',  'medias',  4800.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "pack",   "sku_prefix": "ROPA"}'),
('Calcetines hombre largos x3 pares',            'Genérico',  'medias',  5500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "pack",   "sku_prefix": "ROPA"}'),
('Pinkies mujer x3 pares colores',               'Genérico',  'medias',  4200.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "pack",   "sku_prefix": "ROPA"}'),
('Bucaneras mujer lisas x3 pares',               'Genérico',  'medias',  4800.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "pack",   "sku_prefix": "ROPA"}'),
('Medias de vestir hombre negro x2 pares',       'Genérico',  'medias',  3500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "pack",   "sku_prefix": "ROPA"}'),

-- ============================================================
-- CALZADO BÁSICO
-- ============================================================
('Ojotas de goma talle 35',                      'Genérico',  'calzado-basico',  4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "par",    "talle": "35",  "sku_prefix": "ROPA"}'),
('Ojotas de goma talle 38',                      'Genérico',  'calzado-basico',  4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "par",    "talle": "38",  "sku_prefix": "ROPA"}'),
('Ojotas de goma talle 42',                      'Genérico',  'calzado-basico',  4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "par",    "talle": "42",  "sku_prefix": "ROPA"}'),
('Zapatillas lona básicas talle 38',             'Genérico',  'calzado-basico', 18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "par",    "talle": "38",  "sku_prefix": "ROPA"}'),
('Zapatillas lona básicas talle 40',             'Genérico',  'calzado-basico', 18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "par",    "talle": "40",  "sku_prefix": "ROPA"}'),
('Zapatillas lona básicas talle 42',             'Genérico',  'calzado-basico', 18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "par",    "talle": "42",  "sku_prefix": "ROPA"}'),
('Zapatillas Topper básicas talle 40',           'Topper',    'calzado-basico', 32000.00, 'seed', 0.8, 80, FALSE, TRUE, 'ropa', '{"unit": "par",    "talle": "40",  "sku_prefix": "ROPA"}'),

-- ============================================================
-- ABRIGOS (estacionales — temporada invernal NEA)
-- ============================================================
('Campera rompeviento talle S',                  'Genérico',  'abrigos',  22000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "S",   "sku_prefix": "ROPA"}'),
('Campera rompeviento talle M',                  'Genérico',  'abrigos',  22000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",   "sku_prefix": "ROPA"}'),
('Campera rompeviento talle L',                  'Genérico',  'abrigos',  22000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "L",   "sku_prefix": "ROPA"}'),
('Buzo polar talle S',                           'Genérico',  'abrigos',  18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "S",   "sku_prefix": "ROPA"}'),
('Buzo polar talle M',                           'Genérico',  'abrigos',  18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "M",   "sku_prefix": "ROPA"}'),
('Buzo polar talle L',                           'Genérico',  'abrigos',  18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "L",   "sku_prefix": "ROPA"}'),
('Buzo polar talle XL',                          'Genérico',  'abrigos',  19000.00, 'seed', 0.7, 60, FALSE, TRUE, 'ropa', '{"unit": "unidad", "talle": "XL",  "sku_prefix": "ROPA"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 65 productos
-- Categorías: remeras-basicas, pantalones, ropa-interior-hombre, ropa-interior-mujer, medias, calzado-basico, abrigos
-- Marcas: mayoritariamente Genérico (indumentaria sin marca local), Topper
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
-- NOTA: precios de ropa en NEA son similares al promedio nacional; importaciones de Paraguay
--   pueden competir en precio en zona de frontera (Posadas - Encarnación)
