-- Seed 078: Productos globales — Almacén (~55 productos)
-- CICLO: cycle-005
-- FECHA: 2026-04-24
-- FUENTE: relevamiento supermercados y almacenes NEA (Posadas) + listas de distribuidores 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- REQUIERE: marcas en marketplace_brands (seeds 055, 060, 061, etc. ya ejecutados)
-- CATEGORÍAS REQUERIDAS (slugs confirmados de seeds existentes):
--   quesos-manteca, conservas-enlatados, arroz-legumbres, pastas-secas,
--   harinas-premezclas, aceites-vinagres, detergentes-jabones, lavandina-desinfectantes

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- LÁCTEOS — La Serenísima
-- ============================================================
('Leche entera 1L La Serenísima',            'La Serenísima', 'quesos-manteca',       1350.00, 'seed', 0.7, 82, FALSE, TRUE, 'almacen', '{"unit": "litro",   "sku_prefix": "ALMAC"}'),
('Leche descremada 1L La Serenísima',        'La Serenísima', 'quesos-manteca',       1380.00, 'seed', 0.7, 82, FALSE, TRUE, 'almacen', '{"unit": "litro",   "sku_prefix": "ALMAC"}'),
('Leche larga vida entera 1L La Serenísima', 'La Serenísima', 'quesos-manteca',       1420.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "litro",   "sku_prefix": "ALMAC"}'),
('Yogur entero frutado 190g La Serenísima',  'La Serenísima', 'quesos-manteca',        780.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Queso crema 190g La Serenísima',           'La Serenísima', 'quesos-manteca',       1650.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Manteca 200g La Serenísima',               'La Serenísima', 'quesos-manteca',       2200.00, 'seed', 0.7, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Crema de leche 200ml La Serenísima',       'La Serenísima', 'quesos-manteca',       1450.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- LÁCTEOS — Sancor
-- ============================================================
('Leche entera 1L Sancor',                   'Sancor',        'quesos-manteca',       1320.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "litro",   "sku_prefix": "ALMAC"}'),
('Queso cuartirolo 250g Sancor',             'Sancor',        'quesos-manteca',       2800.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Yogur bebible frutado 900ml Sancor',       'Sancor',        'quesos-manteca',       2100.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- LÁCTEOS — Danone
-- ============================================================
('Yogur Activia natural 190g Danone',        'Danone',        'quesos-manteca',        820.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Yogur Danet vainilla 100g Danone',         'Danone',        'quesos-manteca',        550.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Actimel líquido frutado 100ml Danone',     'Danone',        'quesos-manteca',        650.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- LÁCTEOS — Tregar
-- ============================================================
('Queso por salut 350g Tregar',              'Tregar',        'quesos-manteca',       4200.00, 'seed', 0.6, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Queso pategrás 350g Tregar',               'Tregar',        'quesos-manteca',       4500.00, 'seed', 0.6, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- CONSERVAS Y ENLATADOS — La Campagnola
-- ============================================================
('Tomates perita lata 400g La Campagnola',   'La Campagnola', 'conservas-enlatados',   980.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Choclo en lata 300g La Campagnola',        'La Campagnola', 'conservas-enlatados',   920.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Arvejas en lata 300g La Campagnola',       'La Campagnola', 'conservas-enlatados',   880.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Palmitos en lata 400g La Campagnola',      'La Campagnola', 'conservas-enlatados',  2100.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Duraznos en almíbar lata 825g La Campagnola', 'La Campagnola', 'conservas-enlatados', 2400.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Pechuga de pollo en lata 354g La Campagnola', 'La Campagnola', 'conservas-enlatados', 3200.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- ACEITES Y CONDIMENTOS — Marolio
-- ============================================================
('Aceite de girasol 900ml Marolio',          'Marolio',       'aceites-vinagres',     2100.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Aceite de oliva extra virgen 500ml Marolio', 'Marolio',     'aceites-vinagres',     4800.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Tomates triturados 520g Marolio',          'Marolio',       'conservas-enlatados',   750.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Sal fina 500g Marolio',                    'Marolio',       'conservas-enlatados',   480.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- PASTAS, ARROCES Y HARINAS — Molinos Río de la Plata
-- ============================================================
('Arroz largo fino 1kg Molinos',             'Molinos Río de la Plata', 'arroz-legumbres',    1400.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Fideos tallarín 500g Molinos',             'Molinos Río de la Plata', 'pastas-secas',        820.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Polenta 500g Molinos',                     'Molinos Río de la Plata', 'harinas-premezclas',  750.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Harina 0000 1kg Molinos',                  'Molinos Río de la Plata', 'harinas-premezclas',  980.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),

-- ============================================================
-- AZÚCAR — Ledesma
-- ============================================================
('Azúcar blanca 1kg Ledesma',                'Ledesma',       'conservas-enlatados',  1200.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "kg",      "sku_prefix": "ALMAC"}'),
('Azúcar impalpable 500g Ledesma',           'Ledesma',       'harinas-premezclas',    850.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- LIMPIEZA ROPA — Skip
-- ============================================================
('Jabón en polvo 800g Skip',                 'Skip',          'detergentes-jabones',  3200.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Jabón líquido ropa 900ml Skip',            'Skip',          'detergentes-jabones',  3800.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- LIMPIEZA ROPA — Ariel
-- ============================================================
('Jabón en polvo 800g Ariel',                'Ariel',         'detergentes-jabones',  3500.00, 'seed', 0.7, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Jabón líquido ropa 800ml Ariel',           'Ariel',         'detergentes-jabones',  4200.00, 'seed', 0.7, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Pods detergente 12u Ariel',                'Ariel',         'detergentes-jabones',  4800.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- LIMPIEZA HOGAR — Cif
-- ============================================================
('Crema limpiadora multiuso 750ml Cif',      'Cif',           'detergentes-jabones',  2100.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Crema limpiadora baño 500ml Cif',          'Cif',           'detergentes-jabones',  1850.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- HIGIENE PERSONAL — Dove
-- ============================================================
('Jabón en barra 90g Dove',                  'Dove',          'detergentes-jabones',   950.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Shampoo nutritivo 400ml Dove',             'Dove',          'detergentes-jabones',  3800.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Acondicionador nutritivo 400ml Dove',      'Dove',          'detergentes-jabones',  3800.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- DESINFECCIÓN — Lavandina
-- ============================================================
('Lavandina concentrada 1L Ayudín',          'Ayudín',        'lavandina-desinfectantes', 1450.00, 'seed', 0.6, 76, FALSE, TRUE, 'almacen', '{"unit": "litro",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- DETERGENTE VAJILLA
-- ============================================================
('Detergente limón 750ml Magistral',         'Magistral',     'detergentes-jabones',  1200.00, 'seed', 0.6, 74, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- HIGIENE BUCAL — Colgate
-- ============================================================
('Pasta dental triple acción 90g Colgate',   'Colgate',       'detergentes-jabones',  1850.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Cepillo dental suave Colgate',             'Colgate',       'detergentes-jabones',  1200.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- DESODORANTE — Rexona
-- ============================================================
('Desodorante aerosol 150ml Rexona',         'Rexona',        'detergentes-jabones',  2800.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Desodorante roll-on 50ml Rexona',          'Rexona',        'detergentes-jabones',  2400.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- EXTRAS ALMACÉN (para llegar a ~55 productos)
-- ============================================================
('Fideos moño 500g Molinos',                 'Molinos Río de la Plata', 'pastas-secas',        820.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Fideos spaghetti 500g Molinos',            'Molinos Río de la Plata', 'pastas-secas',        820.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Aceite de girasol 2L Marolio',             'Marolio',       'aceites-vinagres',     4200.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Vinagre de alcohol 500ml Marolio',         'Marolio',       'aceites-vinagres',      680.00, 'seed', 0.7, 74, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Arroz largo fino 500g Molinos',            'Molinos Río de la Plata', 'arroz-legumbres',     720.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Leche entera en polvo 800g La Serenísima', 'La Serenísima', 'quesos-manteca',       6500.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}'),
('Dulce de leche repostero 400g La Serenísima', 'La Serenísima', 'conservas-enlatados', 2600.00, 'seed', 0.7, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad",  "sku_prefix": "ALMAC"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 54 productos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
-- Marcas requeridas: la-serenisima, sancor, danone, tregar, la-campagnola, marolio,
--   molinos-rio-de-la-plata, ledesma, skip, ariel, cif, dove, colgate, rexona
-- Ayudín y Magistral: marcas nacionales de limpieza sin seed propio, se insertan
--   con la información disponible (sin marca en marketplace_brands — revisar si se requiere).
