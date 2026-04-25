-- Seed 104: Productos globales — Piletas: Químicos y Equipos (~85 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento locales de piletas NEA (Posadas) + distribuidores 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: quimicos-pileta, tabletas-cloro, elevadores-ph, equipos-filtracion, accesorios-pileta, limpieza-pileta
-- NOTA: temporada de piletas NEA: oct-abr (clima subtropical, uso extendido)

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- CLORO GRANULADO
-- ============================================================
('Cloro granulado estabilizado 1kg HTH',         'HTH',   'quimicos-pileta', 12000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Cloro granulado estabilizado 5kg HTH',         'HTH',   'quimicos-pileta', 52000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Cloro granulado estabilizado 10kg HTH',        'HTH',   'quimicos-pileta', 98000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Cloro granulado 1kg Genco',                    'Genco', 'quimicos-pileta', 10000.00, 'seed', 0.8, 80, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Cloro granulado 5kg Genco',                    'Genco', 'quimicos-pileta', 46000.00, 'seed', 0.8, 80, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Cloro líquido 10% 5L genérico',                NULL,    'quimicos-pileta',  6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "bidon",  "sku_prefix": "PILE"}'),
('Cloro líquido 10% 20L genérico',               NULL,    'quimicos-pileta', 22000.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "bidon",  "sku_prefix": "PILE"}'),

-- ============================================================
-- TABLETAS DE CLORO
-- ============================================================
('Tabletas de cloro 1" x1kg HTH',                'HTH',   'tabletas-cloro', 14000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Tabletas de cloro 3" x1kg HTH',                'HTH',   'tabletas-cloro', 16000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Tabletas de cloro 3" x4kg HTH',                'HTH',   'tabletas-cloro', 58000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Tabletas de cloro 1" x1kg Genco',              'Genco', 'tabletas-cloro', 12000.00, 'seed', 0.8, 80, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Tabletas de cloro 3" x1kg Genco',              'Genco', 'tabletas-cloro', 14000.00, 'seed', 0.8, 80, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Cloro multiacción tabletas 200g x5u HTH',      'HTH',   'tabletas-cloro', 12000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "pack",   "sku_prefix": "PILE"}'),

-- ============================================================
-- ALGICIDAS Y FLOCULANTES
-- ============================================================
('Algicida líquido 1L HTH',                      'HTH',   'quimicos-pileta',  9500.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Algicida líquido 5L HTH',                      'HTH',   'quimicos-pileta', 42000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "bidon",  "sku_prefix": "PILE"}'),
('Algicida concentrado 1L Genco',                'Genco', 'quimicos-pileta',  8500.00, 'seed', 0.8, 80, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Floculante líquido 1L HTH',                    'HTH',   'quimicos-pileta',  9000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Clarificante líquido 1L genérico',             NULL,    'quimicos-pileta',  6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Floculante tabletas 1kg HTH',                  'HTH',   'quimicos-pileta', 12000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),

-- ============================================================
-- ELEVADORES Y REDUCTORES DE pH
-- ============================================================
('Elevador de pH 1kg HTH',                       'HTH',   'elevadores-ph',   8500.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Elevador de pH 5kg HTH',                       'HTH',   'elevadores-ph',  36000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Reductor de pH ácido 1kg HTH',                 'HTH',   'elevadores-ph',   9500.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Reductor de pH ácido 5kg HTH',                 'HTH',   'elevadores-ph',  40000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Estabilizador ácido cianúrico 1kg genérico',   NULL,    'elevadores-ph',  14000.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Kit análisis agua tiras reactivas 50u',        NULL,    'elevadores-ph',   5500.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "pack",   "sku_prefix": "PILE"}'),
('Kit análisis agua reactivos líquidos',         'HTH',   'elevadores-ph',   8000.00, 'seed', 0.9, 82, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),

-- ============================================================
-- EQUIPOS DE FILTRACIÓN
-- ============================================================
('Bomba de filtración 1/3 HP Sodramar',          'Sodramar', 'equipos-filtracion', 85000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Bomba de filtración 1/2 HP Sodramar',          'Sodramar', 'equipos-filtracion',110000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Bomba de filtración 1 HP Sodramar',            'Sodramar', 'equipos-filtracion',145000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Bomba de filtración 1/3 HP Dancas',            'Dancas',   'equipos-filtracion', 80000.00, 'seed', 0.8, 82, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Bomba de filtración 1/2 HP Dancas',            'Dancas',   'equipos-filtracion',105000.00, 'seed', 0.8, 82, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Filtro de arena completo SF244 Sodramar',      'Sodramar', 'equipos-filtracion',185000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Filtro de arena completo SF310 Sodramar',      'Sodramar', 'equipos-filtracion',250000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Filtro de cartucho para piscina pequeña',      NULL,       'equipos-filtracion', 42000.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Arena sílice filtrante 25kg',                  NULL,       'equipos-filtracion', 12000.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "bolsa",  "sku_prefix": "PILE"}'),
('Cartucho de filtro repuesto 10" genérico',     NULL,       'equipos-filtracion',  8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),

-- ============================================================
-- ACCESORIOS DE PILETA
-- ============================================================
('Manguera flexible 38mm x metro',               NULL,    'accesorios-pileta',  1800.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "metro",  "sku_prefix": "PILE"}'),
('Manguera flexible 38mm x 10m',                 NULL,    'accesorios-pileta', 16000.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "rollo",  "sku_prefix": "PILE"}'),
('Acople rápido manguera 38mm',                  NULL,    'accesorios-pileta',  2500.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Skimmer de superficie plástico',               NULL,    'accesorios-pileta',  8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Boquilla de retorno inundación 1.5"',          NULL,    'accesorios-pileta',  2800.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Escalera de acero inoxidable 2 peldaños',      NULL,    'accesorios-pileta', 45000.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Escalera de acero inoxidable 3 peldaños',      NULL,    'accesorios-pileta', 58000.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Cubre pileta lona 5x10m',                      NULL,    'accesorios-pileta', 35000.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Cubre pileta lona 7x12m',                      NULL,    'accesorios-pileta', 52000.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Termómetro de pileta flotante',                NULL,    'accesorios-pileta',  4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Reflector LED sumergible 12W',                 NULL,    'accesorios-pileta', 18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Difusor de cloro flotante automático',         'HTH',   'accesorios-pileta',  6500.00, 'seed', 0.9, 82, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),

-- ============================================================
-- LIMPIEZA DE PILETA
-- ============================================================
('Limpiafondo manual con mango telescópico',     NULL,    'limpieza-pileta', 18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Red cazahojas para pileta mango 3m',           NULL,    'limpieza-pileta',  8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Cepillo limpia paredes pileta',                NULL,    'limpieza-pileta',  6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Mango telescópico aluminio 2-4m',              NULL,    'limpieza-pileta', 12000.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Limpiafondo automático Aqua Products basic',   'Aqua',  'limpieza-pileta', 85000.00, 'seed', 0.8, 82, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Robot limpiafondo Zodiac OT15',                'Zodiac', 'limpieza-pileta',180000.00, 'seed', 0.9, 85, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Antiincrustante para pileta 1L genérico',      NULL,    'limpieza-pileta',  7500.00, 'seed', 0.7, 60, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}'),
('Desincrustante de azulejos 1L HTH',            'HTH',   'limpieza-pileta', 10000.00, 'seed', 0.9, 82, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 82 productos
-- Categorías: quimicos-pileta, tabletas-cloro, elevadores-ph, equipos-filtracion, accesorios-pileta, limpieza-pileta
-- Marcas: HTH, Genco, Sodramar, Dancas, Aqua, Zodiac
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
-- NOTA: temporada alta piletas NEA es oct-abr (clima tropical/subtropical)
