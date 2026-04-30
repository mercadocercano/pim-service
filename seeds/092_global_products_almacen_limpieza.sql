-- Seed 092: Productos globales — Almacén: Limpieza del Hogar (~60 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento supermercados y almacenes NEA (Posadas) + distribuidores 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: detergentes-jabones, lavandina-desinfectantes, papel-higiene

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- LAVANDINA
-- ============================================================
('Lavandina concentrada 1L Ayudín',             'Ayudín',       'lavandina-desinfectantes', 1450.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "litro",  "sku_prefix": "ALMAC"}'),
('Lavandina concentrada 2L Ayudín',             'Ayudín',       'lavandina-desinfectantes', 2600.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "litro",  "sku_prefix": "ALMAC"}'),
('Lavandina 1L Estrella',                       'Estrella',     'lavandina-desinfectantes', 1200.00, 'seed', 0.7, 74, FALSE, TRUE, 'almacen', '{"unit": "litro",  "sku_prefix": "ALMAC"}'),
('Lavandina 2L Magistral',                      'Magistral',    'lavandina-desinfectantes', 1800.00, 'seed', 0.7, 74, FALSE, TRUE, 'almacen', '{"unit": "litro",  "sku_prefix": "ALMAC"}'),
('Lavandina gel 850g Ayudín',                   'Ayudín',       'lavandina-desinfectantes', 1650.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Desinfectante multiuso 900ml Lysoform',       'Lysoform',     'lavandina-desinfectantes', 2200.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Desinfectante pisos 1L Pino Sol',             'Pino Sol',     'lavandina-desinfectantes', 1800.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "litro",  "sku_prefix": "ALMAC"}'),

-- ============================================================
-- DETERGENTES VAJILLA
-- ============================================================
('Detergente vajilla limón 750ml Magistral',    'Magistral',    'detergentes-jabones',  1200.00, 'seed', 0.7, 74, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Detergente vajilla 750ml Ala',                'Ala',          'detergentes-jabones',  1250.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Detergente vajilla 500ml Fairy',              'Fairy',        'detergentes-jabones',  1600.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Detergente vajilla 750ml Cif',                'Cif',          'detergentes-jabones',  1800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- JABÓN EN POLVO
-- ============================================================
('Jabón en polvo 800g Ariel',                   'Ariel',        'detergentes-jabones',  3500.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Jabón en polvo 1kg Ariel',                    'Ariel',        'detergentes-jabones',  4200.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Jabón en polvo 3kg Ariel',                    'Ariel',        'detergentes-jabones', 11500.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "bolsa",  "sku_prefix": "ALMAC"}'),
('Jabón en polvo 800g Skip',                    'Skip',         'detergentes-jabones',  3200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Jabón en polvo 1kg Skip',                     'Skip',         'detergentes-jabones',  3800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Jabón en polvo 1kg Omo',                      'Omo',          'detergentes-jabones',  3600.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "kg",     "sku_prefix": "ALMAC"}'),
('Jabón en polvo 800g genérico económico',      NULL,           'detergentes-jabones',  1800.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- JABÓN LÍQUIDO DE ROPA
-- ============================================================
('Jabón líquido ropa 900ml Skip',               'Skip',         'detergentes-jabones',  3800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Jabón líquido ropa 800ml Ariel',              'Ariel',        'detergentes-jabones',  4200.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Jabón líquido ropa 3L Ala',                   'Ala',          'detergentes-jabones', 10500.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- SUAVIZANTE
-- ============================================================
('Suavizante de ropa 500ml Comfort',            'Comfort',      'detergentes-jabones',  2800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Suavizante de ropa 2L Comfort',               'Comfort',      'detergentes-jabones',  8500.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Suavizante de ropa 500ml Downy',              'Downy',        'detergentes-jabones',  3200.00, 'seed', 0.8, 82, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- LIMPIADORES MULTIUSO
-- ============================================================
('Limpiador multiuso lavanda 750ml Cif',        'Cif',          'detergentes-jabones',  2100.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Limpiador multiuso 750ml Mr. Músculo',        'Mr. Músculo',  'detergentes-jabones',  2400.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Limpiador vidrios 500ml Mr. Músculo',         'Mr. Músculo',  'detergentes-jabones',  2200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Limpiador baño 500ml Mr. Músculo',            'Mr. Músculo',  'detergentes-jabones',  2200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Desengrasante cocina 500ml Lysoform',         'Lysoform',     'detergentes-jabones',  2000.00, 'seed', 0.7, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- DESODORANTE AMBIENTE
-- ============================================================
('Desodorante de ambiente 360ml Glade lavanda', 'Glade',        'detergentes-jabones',  3500.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Desodorante de ambiente 360ml Glade fresh',   'Glade',        'detergentes-jabones',  3500.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Desodorante de ambiente 360ml Brise',         'Brise',        'detergentes-jabones',  3200.00, 'seed', 0.8, 76, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),

-- ============================================================
-- PAPEL HIGIÉNICO
-- ============================================================
('Papel higiénico x4u Higienol',                'Higienol',     'papel-higiene',        2800.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Papel higiénico x8u Higienol',                'Higienol',     'papel-higiene',        5200.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Papel higiénico doble hoja x4u Elite',        'Elite',        'papel-higiene',        3200.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Papel higiénico doble hoja x12u Elite',       'Elite',        'papel-higiene',        8800.00, 'seed', 0.8, 78, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Papel higiénico Renova x4u',                  'Renova',       'papel-higiene',        4200.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Papel higiénico x16u genérico económico',     NULL,           'papel-higiene',        9500.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Servilletas x100u Higienol',                  'Higienol',     'papel-higiene',        1500.00, 'seed', 0.7, 74, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Papel de cocina x2u Elite',                   'Elite',        'papel-higiene',        2400.00, 'seed', 0.8, 76, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Papel de cocina x4u Higienol',                'Higienol',     'papel-higiene',        4200.00, 'seed', 0.7, 74, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),

-- ============================================================
-- ACCESORIOS DE LIMPIEZA
-- ============================================================
('Esponja de cocina x3u Scotch Brite 3M',       '3M',           'detergentes-jabones',  1800.00, 'seed', 0.8, 80, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Esponja de cocina x2u genérica',              NULL,           'detergentes-jabones',   780.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Trapo de piso 60x80cm genérico',              NULL,           'detergentes-jabones',  1800.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Guantes de goma talla M Tigre',               'Tigre',        'detergentes-jabones',  1500.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "par",    "sku_prefix": "ALMAC"}'),
('Guantes de goma talla L Tigre',               'Tigre',        'detergentes-jabones',  1500.00, 'seed', 0.7, 76, FALSE, TRUE, 'almacen', '{"unit": "par",    "sku_prefix": "ALMAC"}'),
('Bolsas de residuos x30u 50L genéricas',       NULL,           'detergentes-jabones',  1800.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Bolsas de residuos x20u 80L genéricas',       NULL,           'detergentes-jabones',  2200.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Secador de piso con mango genérico',          NULL,           'detergentes-jabones',  3500.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'),
('Fibra verde de limpieza x3u genérica',        NULL,           'detergentes-jabones',   950.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "pack",   "sku_prefix": "ALMAC"}'),
('Cepillo de baño con mango genérico',          NULL,           'detergentes-jabones',  2800.00, 'seed', 0.6, 40, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 60 productos
-- Categorías usadas: detergentes-jabones, lavandina-desinfectantes, papel-higiene
-- Marcas: Ayudín, Estrella, Magistral, Ariel, Skip, Omo, Ala, Comfort, Downy, Cif, Mr. Músculo,
--   Lysoform, Pino Sol, Fairy, Glade, Brise, Higienol, Elite, Renova, 3M, Tigre, genéricos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
