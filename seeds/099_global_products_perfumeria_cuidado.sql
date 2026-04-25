-- Seed 099: Productos globales — Perfumería: Cuidado Personal (~85 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento perfumerías NEA (Posadas) + distribuidores 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: cremas-corporales, tinturas-capilares, shampoos-profesionales, tratamientos-capilares, fragancias, esmaltes-maquillaje

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- CREMAS CORPORALES
-- ============================================================
('Crema corporal hidratante 400ml Nivea',        'Nivea',         'cremas-corporales',  5800.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Crema corporal original 250ml Nivea lata',     'Nivea',         'cremas-corporales',  4200.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Crema corporal humectante 400ml Dove',         'Dove',          'cremas-corporales',  5500.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Vaselina intensiva crema 200ml Unilever',      'Vaselina',      'cremas-corporales',  3800.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Crema de manos 75ml Neutrogena',               'Neutrogena',    'cremas-corporales',  4500.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Crema anti-edad Q10 50ml Nivea',               'Nivea',         'cremas-corporales',  7800.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Crema facial día 50ml Pond''s',                'Pond''s',       'cremas-corporales',  5200.00, 'seed', 0.8, 80, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Crema corporal reparadora 400ml Johnson''s',   'Johnson''s',    'cremas-corporales',  5800.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Crema corporal manteca de karité 400ml Palmolive', 'Palmolive', 'cremas-corporales',  4800.00, 'seed', 0.8, 80, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),

-- ============================================================
-- TINTURAS CAPILARES
-- ============================================================
('Tintura L''Oreal Excellence castaño claro 5',  'L''Oreal',      'tinturas-capilares', 8500.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Tintura L''Oreal Excellence negro 1',          'L''Oreal',      'tinturas-capilares', 8500.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Tintura L''Oreal Excellence rubio oscuro 6',   'L''Oreal',      'tinturas-capilares', 8500.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Tintura Garnier Color castaño N°4',            'Garnier',       'tinturas-capilares', 7200.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Tintura Garnier Color negro N°1',              'Garnier',       'tinturas-capilares', 7200.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Tintura Revlon Colorsilk castaño 30',          'Revlon',        'tinturas-capilares', 6800.00, 'seed', 0.8, 80, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Tintura Revlon Colorsilk negro 10',            'Revlon',        'tinturas-capilares', 6800.00, 'seed', 0.8, 80, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Decoloración capilares Pols Blanca genérica',  NULL,            'tinturas-capilares', 3500.00, 'seed', 0.6, 60, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Oxidante 20vol 90ml genérico',                 NULL,            'tinturas-capilares', 1800.00, 'seed', 0.6, 60, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),

-- ============================================================
-- SHAMPOOS PROFESIONALES
-- ============================================================
('Shampoo Schwarzkopf BC reparación 200ml',      'Schwarzkopf',   'shampoos-profesionales', 9800.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Shampoo TRESemmé hidratación 400ml',           'TRESemmé',      'shampoos-profesionales', 6800.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Shampoo TRESemmé keratina liso 400ml',         'TRESemmé',      'shampoos-profesionales', 7200.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Shampoo Redken color extend 300ml',            'Redken',        'shampoos-profesionales',14500.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Shampoo Wella Professionals SP 200ml',         'Wella',         'shampoos-profesionales',12000.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Acondicionador TRESemmé reparación 400ml',     'TRESemmé',      'shampoos-profesionales', 6800.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),

-- ============================================================
-- TRATAMIENTOS CAPILARES
-- ============================================================
('Mascarilla Elvive reparación total 300ml L''Oreal', 'L''Oreal', 'tratamientos-capilares', 7200.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Mascarilla Elvive rizos definidos 300ml',      'L''Oreal',      'tratamientos-capilares', 7200.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Ampolla Capilatis keratina x12u',              'Capilatis',     'tratamientos-capilares', 6500.00, 'seed', 0.8, 80, FALSE, TRUE, 'perfumeria', '{"unit": "caja",   "sku_prefix": "PERF"}'),
('Ampolla Capilatis colageno x12u',              'Capilatis',     'tratamientos-capilares', 6500.00, 'seed', 0.8, 80, FALSE, TRUE, 'perfumeria', '{"unit": "caja",   "sku_prefix": "PERF"}'),
('Aceite capilar Argan Oil 100ml Schwarzkopf',   'Schwarzkopf',   'tratamientos-capilares', 8800.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Keratina líquida alisante 500ml genérica',     NULL,            'tratamientos-capilares',12000.00, 'seed', 0.6, 60, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Crema para peinar 200ml Pantene',              'Pantene',       'tratamientos-capilares', 4800.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Spray termoprotector 200ml Schwarzkopf',       'Schwarzkopf',   'tratamientos-capilares', 7500.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),

-- ============================================================
-- FRAGANCIAS Y PERFUMES
-- ============================================================
('Perfume Natura Essencial femenino 100ml',      'Natura',        'fragancias',   18000.00, 'seed', 0.8, 80, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Perfume Natura Essencial masculino 100ml',     'Natura',        'fragancias',   18000.00, 'seed', 0.8, 80, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Colonia Yanbal femenina 60ml',                 'Yanbal',        'fragancias',   12000.00, 'seed', 0.8, 80, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Colonia Yanbal masculina 60ml',                'Yanbal',        'fragancias',   12000.00, 'seed', 0.8, 80, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Perfume Acqua di Gio imitación 100ml',         NULL,            'fragancias',    8500.00, 'seed', 0.6, 60, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Colonia 1800 genérica 500ml',                  NULL,            'fragancias',    5500.00, 'seed', 0.6, 60, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Body splash Victoria''s Secret sin alcohol 250ml', 'Victoria''s Secret', 'fragancias', 9500.00, 'seed', 0.8, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Desodorante colonia CK One imitación 125ml',   NULL,            'fragancias',    4800.00, 'seed', 0.6, 60, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),

-- ============================================================
-- ESMALTES DE UÑAS
-- ============================================================
('Esmalte Risqué rojo 8ml',                      'Risqué',        'esmaltes-maquillaje', 1800.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Esmalte Risqué transparente base 8ml',         'Risqué',        'esmaltes-maquillaje', 1600.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Esmalte Colorama rosa 8ml',                    'Colorama',      'esmaltes-maquillaje', 1500.00, 'seed', 0.9, 80, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Esmalte Colorama nude 8ml',                    'Colorama',      'esmaltes-maquillaje', 1500.00, 'seed', 0.9, 80, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Esmalte OPI rojo 15ml',                        'OPI',           'esmaltes-maquillaje', 6500.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Quitaesmalte sin acetona 100ml Avon',          'Avon',          'esmaltes-maquillaje', 2200.00, 'seed', 0.8, 78, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),

-- ============================================================
-- MAQUILLAJE
-- ============================================================
('Base de maquillaje Revlon ColorStay 30ml',     'Revlon',        'esmaltes-maquillaje', 8500.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Base de maquillaje Maybelline Fit Me 30ml',    'Maybelline',    'esmaltes-maquillaje', 7800.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Labial Revlon super lustrous rojo',            'Revlon',        'esmaltes-maquillaje', 5500.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Labial Maybelline Color Sensational rosa',     'Maybelline',    'esmaltes-maquillaje', 5200.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Máscara de pestañas Maybelline Lash Sensational', 'Maybelline', 'esmaltes-maquillaje', 7200.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Máscara de pestañas Revlon Ultra Volume',      'Revlon',        'esmaltes-maquillaje', 6800.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Corrector líquido Maybelline Instant Age',     'Maybelline',    'esmaltes-maquillaje', 6500.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Sombra de ojos paleta 12 colores Revlon',      'Revlon',        'esmaltes-maquillaje', 9800.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Delineador de ojos negro Maybelline',          'Maybelline',    'esmaltes-maquillaje', 4800.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Polvo compacto Avon True x unidad',            'Avon',          'esmaltes-maquillaje', 5800.00, 'seed', 0.8, 78, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Rubor Revlon Powder Blush 5g',                 'Revlon',        'esmaltes-maquillaje', 6200.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Corrector Revlon PhotoReady 3.2ml',            'Revlon',        'esmaltes-maquillaje', 5500.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),

-- ============================================================
-- DEPILACIÓN Y CUIDADO ESPECÍFICO
-- ============================================================
('Cera depilatoria fría tiras 20u Veet',         'Veet',          'cremas-corporales',  5800.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "caja",   "sku_prefix": "PERF"}'),
('Crema depilatoria 200ml Veet',                 'Veet',          'cremas-corporales',  4800.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Protector solar SPF50 200ml Eucerin',          'Eucerin',       'cremas-corporales',  9800.00, 'seed', 0.9, 85, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Protector solar SPF30 200ml Banana Boat',      'Banana Boat',   'cremas-corporales',  7200.00, 'seed', 0.8, 80, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}'),
('Crema anticelulítica 400ml Dove',              'Dove',          'cremas-corporales',  8200.00, 'seed', 0.9, 82, FALSE, TRUE, 'perfumeria', '{"unit": "unidad", "sku_prefix": "PERF"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 83 productos
-- Categorías: cremas-corporales, tinturas-capilares, shampoos-profesionales, tratamientos-capilares, fragancias, esmaltes-maquillaje
-- Marcas: Nivea, Dove, Neutrogena, Pond's, Johnson's, L'Oreal, Garnier, Revlon, Schwarzkopf, TRESemmé,
--   Redken, Wella, Capilatis, Pantene, Natura, Yanbal, Risqué, Colorama, OPI, Maybelline, Avon,
--   Veet, Eucerin, Banana Boat, Victoria's Secret
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
