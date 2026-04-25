-- Seed 105: Productos globales — Bazar: Vajilla, Utensilios y Hogar (~65 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento bazares NEA (Posadas) + distribuidores 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: vajilla-vidrieria, utensilios-cocina, almacenamiento-hogar, textiles-hogar, organizacion-hogar

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- VAJILLA Y VIDRIERÍA
-- ============================================================
('Juego de platos hondo x6u loza blanca',        NULL,          'vajilla-vidrieria',  18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "juego",  "sku_prefix": "BAZA"}'),
('Juego de platos plano x6u loza blanca',        NULL,          'vajilla-vidrieria',  18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "juego",  "sku_prefix": "BAZA"}'),
('Juego vajilla 18 piezas loza blanca',          NULL,          'vajilla-vidrieria',  48000.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "juego",  "sku_prefix": "BAZA"}'),
('Taza con plato café porcelana blanca',         NULL,          'vajilla-vidrieria',   3500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Set tazas x4u 350ml OLA',                      'OLA',         'vajilla-vidrieria',  18000.00, 'seed', 0.8, 80, FALSE, TRUE, 'bazar', '{"unit": "set",    "sku_prefix": "BAZA"}'),
('Vaso alto vidrio 400ml x6u',                   NULL,          'vajilla-vidrieria',  12000.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "set",    "sku_prefix": "BAZA"}'),
('Vaso plástico irrompible 300ml x6u',           NULL,          'vajilla-vidrieria',   6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "set",    "sku_prefix": "BAZA"}'),
('Ensaladera vidrio templado 3L',                NULL,          'vajilla-vidrieria',  14000.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Fuente oval vidrio refractario 2L Marinex',    'Marinex',     'vajilla-vidrieria',  18000.00, 'seed', 0.8, 80, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Set cubiertos x24u acero inoxidable',          NULL,          'vajilla-vidrieria',  28000.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "set",    "sku_prefix": "BAZA"}'),
('Cubiertos x12u acero Tramontina',              'Tramontina',  'vajilla-vidrieria',  22000.00, 'seed', 0.9, 85, FALSE, TRUE, 'bazar', '{"unit": "set",    "sku_prefix": "BAZA"}'),

-- ============================================================
-- UTENSILIOS DE COCINA
-- ============================================================
('Cuchillo chef 20cm Tramontina',                'Tramontina',  'utensilios-cocina',   8500.00, 'seed', 0.9, 85, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Cuchillo multiuso 15cm Tramontina',            'Tramontina',  'utensilios-cocina',   6500.00, 'seed', 0.9, 85, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Cuchillo pan 20cm Tramontina',                 'Tramontina',  'utensilios-cocina',   7800.00, 'seed', 0.9, 85, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Cuchillo set x5u con taco Mundial',            'Mundial',     'utensilios-cocina',  45000.00, 'seed', 0.9, 85, FALSE, TRUE, 'bazar', '{"unit": "set",    "sku_prefix": "BAZA"}'),
('Tabla de picar plástico 30x20cm',              NULL,          'utensilios-cocina',   5500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Tabla de picar madera bambú 35x25cm',          NULL,          'utensilios-cocina',  12000.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Olla enlozada 20cm con tapa',                  NULL,          'utensilios-cocina',  18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Olla acero inoxidable 22cm Imusa',             'Imusa',       'utensilios-cocina',  32000.00, 'seed', 0.8, 80, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Sartén teflón 24cm con mango',                 NULL,          'utensilios-cocina',  18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Sartén acero 26cm OLA',                        'OLA',         'utensilios-cocina',  28000.00, 'seed', 0.8, 80, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Espátula silicona',                            NULL,          'utensilios-cocina',   3500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Cucharón acero inoxidable',                    'Tramontina',  'utensilios-cocina',   4500.00, 'seed', 0.9, 82, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Rallador 4 caras acero',                       NULL,          'utensilios-cocina',   6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Colador acero inoxidable 20cm',                NULL,          'utensilios-cocina',   7500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Abrelatas manual acero',                       'Tramontina',  'utensilios-cocina',   4800.00, 'seed', 0.9, 80, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Pelador de papas acero',                       NULL,          'utensilios-cocina',   2500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),

-- ============================================================
-- ALMACENAMIENTO Y HERMÉTICOS
-- ============================================================
('Contenedor hermético 1L plástico',             NULL,          'almacenamiento-hogar', 3200.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Contenedor hermético 2L plástico',             NULL,          'almacenamiento-hogar', 4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Set contenedores herméticos x3u plástico',     NULL,          'almacenamiento-hogar', 9500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "set",    "sku_prefix": "BAZA"}'),
('Frasco de vidrio 500ml con tapa',              NULL,          'almacenamiento-hogar', 3800.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Frasco de vidrio 1L con tapa',                 NULL,          'almacenamiento-hogar', 5200.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Bolsas zip reutilizables x20u 1L',             NULL,          'almacenamiento-hogar', 3500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "pack",   "sku_prefix": "BAZA"}'),
('Tarro con tapa hermética 3L plástico',         NULL,          'almacenamiento-hogar', 5800.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Canasto plástico multiuso 10L',                NULL,          'almacenamiento-hogar', 4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),

-- ============================================================
-- TEXTILES Y HOGAR
-- ============================================================
('Vela aromática soja 200g',                     NULL,          'textiles-hogar',  4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Vela de emergencia blanca 8h x6u',             NULL,          'textiles-hogar',  2800.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "pack",   "sku_prefix": "BAZA"}'),
('Mantel hule 1.4x1.8m estampado',              NULL,          'textiles-hogar',  8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Repasador algodón x3u',                        NULL,          'textiles-hogar',  4800.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "pack",   "sku_prefix": "BAZA"}'),
('Delantal de cocina tela',                      NULL,          'textiles-hogar',  5500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),

-- ============================================================
-- ORGANIZACIÓN Y HOGAR
-- ============================================================
('Organizador de cajón plástico 30x20cm',        NULL,          'organizacion-hogar', 3500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Gancho adhesivo de pared 5kg x4u',             NULL,          'organizacion-hogar', 3200.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "pack",   "sku_prefix": "BAZA"}'),
('Colgador de ropa pared 4 ganchos',             NULL,          'organizacion-hogar', 5500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Portarrollo papel higiénico baño',             NULL,          'organizacion-hogar', 4200.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Jabonera para baño plástico',                  NULL,          'organizacion-hogar', 2800.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Papelera de baño plástico 10L',                NULL,          'organizacion-hogar', 3500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Perchero con base 6 ganchos',                  NULL,          'organizacion-hogar',12000.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Cesto de ropa sucia 30L plástico',             NULL,          'organizacion-hogar', 8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Tendedero de ropa plegable 12m',               NULL,          'organizacion-hogar',12000.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Pinzas de ropa plástico x40u',                 NULL,          'organizacion-hogar', 2800.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "pack",   "sku_prefix": "BAZA"}'),
('Escobillón para baño con soporte',             NULL,          'organizacion-hogar', 4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Palangana plástico 10L',                       NULL,          'organizacion-hogar', 3200.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Balde plástico con mango 10L',                 NULL,          'organizacion-hogar', 4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}'),
('Felpudo antideslizante 60x40cm',               NULL,          'textiles-hogar',  5800.00, 'seed', 0.7, 60, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZA"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 62 productos
-- Categorías: vajilla-vidrieria, utensilios-cocina, almacenamiento-hogar, textiles-hogar, organizacion-hogar
-- Marcas: Tramontina, Imusa, OLA, Mundial, Marinex — mayoría genérico
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
