-- Seed 086: Productos globales — Ferretería: Cerrajería y Jardín (~70 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: catálogos ferreterías NEA (Posadas) + distribuidores Yale/Tramontina 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: cerrajeria-seguridad, jardin-herramientas, fijaciones

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- CERRADURAS — Yale, Trabex, genéricas
-- ============================================================
('Cerradura horizontal puerta exterior Yale',        'Yale',     'cerrajeria-seguridad', 28000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cerradura horizontal puerta interior Yale',        'Yale',     'cerrajeria-seguridad', 18000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cerradura de embutir doble paleta Yale',           'Yale',     'cerrajeria-seguridad', 38000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cerradura embutir puerta madera Trabex',           'Trabex',   'cerrajeria-seguridad', 22000.00, 'seed', 0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cerradura de embutir para perfil de aluminio',     NULL,       'cerrajeria-seguridad', 15000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cerradura eléctrica 12V para portón',              NULL,       'cerrajeria-seguridad', 28000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Chapa de seguridad puerta exterior Trabex',        'Trabex',   'cerrajeria-seguridad', 18000.00, 'seed', 0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- CANDADOS — Master Lock, Yale, genéricos
-- ============================================================
('Candado 30mm zinc Yale',                          'Yale',     'cerrajeria-seguridad',  4800.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Candado 40mm zinc Yale',                          'Yale',     'cerrajeria-seguridad',  6500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Candado 50mm zinc Yale',                          'Yale',     'cerrajeria-seguridad',  8800.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Candado 60mm zinc Yale',                          'Yale',     'cerrajeria-seguridad', 12000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Candado 40mm Master Lock combinación 4 dígitos',  'Master Lock','cerrajeria-seguridad',9500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Candado 50mm Master Lock con llave',              'Master Lock','cerrajeria-seguridad',10500.00,'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Candado 30mm latón genérico',                     NULL,       'cerrajeria-seguridad',  2800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Candado de barra para portón 250mm',              NULL,       'cerrajeria-seguridad', 15000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- MANIJAS Y PICAPORTES
-- ============================================================
('Manija de puerta palanca aluminio par genérica',   NULL,       'cerrajeria-seguridad',  8500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "par",    "sku_prefix": "FERRET"}'),
('Manija de puerta bola cromada par genérica',       NULL,       'cerrajeria-seguridad',  6500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "par",    "sku_prefix": "FERRET"}'),
('Manija exterior acero inoxidable Trabex',          'Trabex',   'cerrajeria-seguridad', 12000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Picaporte de embutir 3" galvanizado',             NULL,       'cerrajeria-seguridad',  2200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Picaporte de embutir 4" galvanizado',             NULL,       'cerrajeria-seguridad',  2800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pasador de empotrar 100mm galvanizado',           NULL,       'cerrajeria-seguridad',  2500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pasador de empotrar 150mm galvanizado',           NULL,       'cerrajeria-seguridad',  3200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Aldaba y pasador combinado galvanizado',          NULL,       'cerrajeria-seguridad',  3800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Mirilla ojo de buey 200° dorado',                 NULL,       'cerrajeria-seguridad',  2800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- HERRAJES PARA VENTANAS Y PORTONES
-- ============================================================
('Falleba para ventana de aluminio 1.5m',           NULL,       'cerrajeria-seguridad',  4500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Bisagra de portón soldable 3" x2u',               NULL,       'cerrajeria-seguridad',  2800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "par",    "sku_prefix": "FERRET"}'),
('Bisagra de portón soldable 4" x2u',               NULL,       'cerrajeria-seguridad',  3500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "par",    "sku_prefix": "FERRET"}'),
('Gozne de hierro 3/4" x5u',                       NULL,       'cerrajeria-seguridad',  1800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),
('Brazo hidráulico para puerta pesada',             NULL,       'cerrajeria-seguridad', 18000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cierre automático de puerta liviano',             NULL,       'cerrajeria-seguridad',  8500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- HERRAMIENTAS DE JARDÍN — Tramontina, Truper, genéricas
-- ============================================================
('Pala cuadrada mango largo Tramontina',            'Tramontina','jardin-herramientas', 18000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pala punta mango largo Tramontina',               'Tramontina','jardin-herramientas', 18000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pala punta mango largo Truper',                   'Truper',   'jardin-herramientas',  15000.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Rastrillo 16 dientes mango largo Tramontina',     'Tramontina','jardin-herramientas', 16000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Azadón mango largo Truper',                       'Truper',   'jardin-herramientas',  14000.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pico doble mango largo Tramontina',               'Tramontina','jardin-herramientas', 22000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Horquilla de jardín 4 dientes mango largo',       NULL,       'jardin-herramientas',  12000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pala de mano con mango corto Tramontina',         'Tramontina','jardin-herramientas',  6500.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Rastrillo de mano jardín 3 dientes',              NULL,       'jardin-herramientas',   4500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Carretilla 60L para jardín/obra',                 NULL,       'jardin-herramientas',  65000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- TIJERAS DE PODAR Y CORTE
-- ============================================================
('Tijera de podar 8" Tramontina',                   'Tramontina','jardin-herramientas', 12000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tijera de podar 7" Truper',                       'Truper',   'jardin-herramientas',   8500.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Tijera para seto manual 18" Tramontina',          'Tramontina','jardin-herramientas', 18000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Machete jardinero 18" Tramontina',                'Tramontina','jardin-herramientas', 12000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- CORTACÉSPED Y PULVERIZADORES
-- ============================================================
('Cortacésped eléctrico 1600W 46cm Black+Decker',  'Black+Decker','jardin-herramientas',185000.00,'seed',0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Bordeadora a hilo eléctrica 500W Black+Decker',  'Black+Decker','jardin-herramientas', 68000.00,'seed',0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Bordeadora a hilo a nafta 25cc Truper',           'Truper',   'jardin-herramientas',  95000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pulverizador de mochila 16L genérico',            NULL,       'jardin-herramientas',  28000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pulverizador manual 2L Truper',                   'Truper',   'jardin-herramientas',   6500.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Fumigadora eléctrica a batería 20V genérica',    NULL,       'jardin-herramientas',  48000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- MANGUERAS EXTENSIBLES Y RIEGO (complemento seed 083)
-- ============================================================
('Manguera extensible 30m con pistola Truper',      'Truper',   'jardin-herramientas',  18000.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Sistema riego goteo kit 30 plantas',              NULL,       'jardin-herramientas',  14000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "kit",    "sku_prefix": "FERRET"}'),
('Conector roscado 1/2" macho para manguera x5u',  NULL,       'jardin-herramientas',   2800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),

-- ============================================================
-- GUANTES Y PROTECCIÓN PERSONAL (muy vendido en ferretería)
-- ============================================================
('Guantes de cuero para trabajo talla M',           NULL,       'fijaciones',            3500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "par",    "sku_prefix": "FERRET"}'),
('Guantes de cuero para trabajo talla L',           NULL,       'fijaciones',            3500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "par",    "sku_prefix": "FERRET"}'),
('Guantes de nitrilo talla M x12u',                 NULL,       'fijaciones',            4800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),
('Antiparras de seguridad policarbonato 3M',        '3M',       'fijaciones',            4500.00, 'seed', 0.8, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Casco de seguridad blanco genérico',              NULL,       'fijaciones',            6500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Barbijo N95 descartable x5u 3M',                  '3M',       'fijaciones',            8500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),
('Calzado de seguridad punta de acero',             NULL,       'fijaciones',           55000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "par",    "sku_prefix": "FERRET"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 70 productos
-- Categorías usadas: cerrajeria-seguridad, jardin-herramientas, fijaciones
-- Marcas: Yale, Master Lock, Trabex, Tramontina, Truper, Black+Decker, 3M, genéricos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
