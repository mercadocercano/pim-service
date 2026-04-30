-- Seed 094: Productos globales — Almacén: Perfumería Básica (~40 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento supermercados y almacenes NEA (Posadas) + distribuidores 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: detergentes-jabones, higiene-personal, cuidado-bebe

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- SHAMPOO
-- ============================================================
('Shampoo hidratación intensa 400ml Sedal',     'Sedal',      'higiene-personal',  3200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Shampoo caída defensa 400ml Sedal',           'Sedal',      'higiene-personal',  3200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Shampoo reparación total 400ml Pantene',      'Pantene',    'higiene-personal',  3800.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Shampoo caspa 400ml Head & Shoulders',        'Head & Shoulders','higiene-personal',3800.00,'seed',0.8,82,FALSE,TRUE,'almacen','{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Shampoo nutritivo 400ml Dove',                'Dove',        'higiene-personal',  3800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Shampoo nutritivo 400ml Elvive L''Oreal',     'L''Oreal',   'higiene-personal',  4200.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Acondicionador reparación total 400ml Pantene','Pantene',   'higiene-personal',  3800.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- JABÓN DE TOCADOR
-- ============================================================
('Jabón en barra 90g Dove',                     'Dove',       'higiene-personal',   950.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Jabón en barra x3u 90g Dove',                 'Dove',       'higiene-personal',  2500.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Jabón en barra 100g Lux',                     'Lux',        'higiene-personal',   850.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Jabón en barra 90g Rexona',                   'Rexona',     'higiene-personal',   920.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Jabón líquido tocador 225ml Dove',            'Dove',       'higiene-personal',  2800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- DESODORANTE
-- ============================================================
('Desodorante aerosol 150ml Rexona feminino',   'Rexona',     'higiene-personal',  2800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Desodorante aerosol 150ml Rexona masculino',  'Rexona',     'higiene-personal',  2800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Desodorante roll-on 50ml Rexona feminino',    'Rexona',     'higiene-personal',  2400.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Desodorante aerosol 150ml Dove',              'Dove',       'higiene-personal',  3200.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Desodorante aerosol 150ml Nivea Fresh',       'Nivea',      'higiene-personal',  2800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Desodorante roll-on 50ml Nivea',              'Nivea',      'higiene-personal',  2200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- CREMA CORPORAL
-- ============================================================
('Crema de manos humectante 200ml Nivea',       'Nivea',      'higiene-personal',  2800.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Crema de manos 75ml Hinds',                   'Hinds',      'higiene-personal',  1200.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Crema corporal humectante 400ml Nivea',       'Nivea',      'higiene-personal',  5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Vaselina pura 100g Nivea',                    'Nivea',      'higiene-personal',  1800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- HIGIENE BUCAL
-- ============================================================
('Pasta dental triple acción 90g Colgate',      'Colgate',    'higiene-personal',  1850.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Pasta dental 130g Colgate luminous white',    'Colgate',    'higiene-personal',  2400.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Pasta dental 100g Oral-B',                    'Oral-B',     'higiene-personal',  2200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Cepillo dental suave Colgate',                'Colgate',    'higiene-personal',  1200.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Enjuague bucal 500ml Listerine menta',        'Listerine',  'higiene-personal',  3800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- PAÑALES — alta rotación en almacén NEA
-- ============================================================
('Pañales Huggies Active Sec talle S x24u',     'Huggies',    'cuidado-bebe',     8500.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Pañales Huggies Active Sec talle M x24u',     'Huggies',    'cuidado-bebe',     8500.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Pañales Huggies Active Sec talle G x20u',     'Huggies',    'cuidado-bebe',     8500.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Pañales Huggies Active Sec talle XG x16u',    'Huggies',    'cuidado-bebe',     8500.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Pañales Pampers talle M x20u',                'Pampers',    'cuidado-bebe',     9200.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Pañales Pampers talle G x16u',                'Pampers',    'cuidado-bebe',     9200.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),

-- ============================================================
-- TOALLAS FEMENINAS
-- ============================================================
('Toallas femeninas Always tela x8u',           'Always',     'higiene-personal',  2400.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Toallas femeninas Always noches x8u',         'Always',     'higiene-personal',  2600.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Toallas femeninas Kotex x8u',                 'Kotex',      'higiene-personal',  2200.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),

-- ============================================================
-- AFEITADO
-- ============================================================
('Gillette Prestobarba 3 hojas x3u',            'Gillette',   'higiene-personal',  3200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Gillette Venus desechable x3u',               'Gillette',   'higiene-personal',  3800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Crema de afeitado 200ml Gillete',             'Gillette',   'higiene-personal',  2800.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 41 productos
-- Categorías usadas: higiene-personal, cuidado-bebe, detergentes-jabones
-- Marcas: Sedal, Pantene, Head & Shoulders, Dove, L'Oreal, Lux, Rexona, Nivea, Hinds,
--   Colgate, Oral-B, Listerine, Huggies, Pampers, Always, Kotex, Gillette
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
