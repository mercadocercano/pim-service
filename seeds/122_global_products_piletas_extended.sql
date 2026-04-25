-- Seed 122: Productos globales — Piletas extendida (~35 productos)
-- CICLO: cycle-009-catalog-volume-expansion (T013)
-- FECHA: 2026-04-25
-- FUENTE: relevamiento locales piletas NEA + distribuidores AstralPool/CTX/Genco/Biopool 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- OBJETIVO: llevar piletas a >= 113 verificados (hoy: 83 → +30)
-- MARCAS: Biopool, CTX, AstralPool, Genco, HTH, Nataclor
-- NOTA: temporada extendida NEA oct-abr; clima subtropical, alta demanda mantenimiento

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- pH REDUCTORES/ELEVADORES — presentaciones faltantes
-- ============================================================
('pH menos granulado 5kg CTX',                 'CTX',       'cloro-quimicos',      38000.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "bolsa", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('pH menos granulado 1kg CTX',                 'CTX',       'cloro-quimicos',       9500.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "bolsa", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('pH más granulado 5kg CTX',                   'CTX',       'cloro-quimicos',      36000.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "bolsa", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('pH más granulado 1kg CTX',                   'CTX',       'cloro-quimicos',       9000.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "bolsa", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('pH menos 5kg Biopool',                       'Biopool',   'cloro-quimicos',      35000.00, 'seed', 0.85, 80, TRUE, TRUE, 'piletas', '{"unit": "bolsa", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('pH más 5kg Biopool',                         'Biopool',   'cloro-quimicos',      34000.00, 'seed', 0.85, 80, TRUE, TRUE, 'piletas', '{"unit": "bolsa", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- ALGICIDAS — presentaciones 5L faltantes
-- ============================================================
('Algicida concentrado 5L CTX',                'CTX',       'quimicos-pileta',     28000.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "bidon", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Algicida concentrado 1L CTX',                'CTX',       'quimicos-pileta',      7500.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "bidon", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Algicida concentrado 5L Biopool',            'Biopool',   'quimicos-pileta',     26000.00, 'seed', 0.85, 80, TRUE, TRUE, 'piletas', '{"unit": "bidon", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Algicida concentrado 1L Biopool',            'Biopool',   'quimicos-pileta',      7000.00, 'seed', 0.85, 80, TRUE, TRUE, 'piletas', '{"unit": "bidon", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Algicida 1L AstralPool',                     'AstralPool','quimicos-pileta',      8000.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "bidon", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Algicida 5L AstralPool',                     'AstralPool','quimicos-pileta',     30000.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "bidon", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- CLARIFICADORES Y FLOCULANTES
-- ============================================================
('Clarificador de agua 1L CTX',                'CTX',       'quimicos-pileta',      8500.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "bidon", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Clarificador líquido 1L AstralPool',         'AstralPool','quimicos-pileta',      9000.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "bidon", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Floculante en pastillas 1kg HTH',            'HTH',       'quimicos-pileta',     11000.00, 'seed', 0.90, 85, TRUE, TRUE, 'piletas', '{"unit": "bolsa", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Floculante en pastillas 1kg AstralPool',     'AstralPool','quimicos-pileta',     12000.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "bolsa", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- ARENA DE FILTRO
-- ============================================================
('Arena de filtro sílice 25kg AstralPool',     'AstralPool','equipos-filtracion',  18000.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "bolsa", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Arena de filtro sílice 25kg Genco',          'Genco',     'equipos-filtracion',  15000.00, 'seed', 0.85, 80, TRUE, TRUE, 'piletas', '{"unit": "bolsa", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- BOMBAS DE CIRCULACIÓN — marcas adicionales
-- ============================================================
('Bomba de circulación 1/2HP AstralPool',      'AstralPool','bombas-filtros-pileta',85000.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Bomba de circulación 3/4HP AstralPool',      'AstralPool','bombas-filtros-pileta',115000.00,'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Bomba de circulación 1HP Sodramar',          'Sodramar',  'bombas-filtros-pileta',130000.00,'seed', 0.85, 80, TRUE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- VÁLVULA SELECTORA Y ACCESORIOS FILTRO
-- ============================================================
('Válvula selectora 6 vías 1.5" Sodramar',     'Sodramar',  'equipos-filtracion',  22000.00, 'seed', 0.85, 80, TRUE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Válvula selectora 6 vías 2" Sodramar',       'Sodramar',  'equipos-filtracion',  28000.00, 'seed', 0.85, 80, TRUE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Válvula selectora 6 vías 1.5" AstralPool',   'AstralPool','equipos-filtracion',  24000.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Manómetro de presión para filtro 1/4"',      NULL,        'equipos-filtracion',   4500.00, 'seed', 0.78, 70, TRUE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- LIMPIEZA Y ACCESORIOS — ítems faltantes
-- ============================================================
('Termómetro flotante para pileta genérico',   NULL,        'accesorios-pileta',    2800.00, 'seed', 0.78, 70, TRUE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Aspiradora manual de pileta telescópica',    NULL,        'limpieza-pileta',      8500.00, 'seed', 0.80, 75, TRUE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Aspiradora manual de pileta Kokido',         'Kokido',    'limpieza-pileta',     10500.00, 'seed', 0.85, 82, TRUE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Red de recolección hojas mango 4m',          NULL,        'limpieza-pileta',       4200.00, 'seed', 0.80, 75, TRUE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),
('Cepillo limpia paredes pileta Kokido',       'Kokido',    'limpieza-pileta',       3800.00, 'seed', 0.85, 82, TRUE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA"}'),

-- ============================================================
-- KIT MANTENIMIENTO
-- ============================================================
('Kit mantenimiento pileta temporada Genco',   'Genco',     'mantenimiento-pileta', 45000.00, 'seed', 0.85, 82, TRUE, TRUE, 'piletas', '{"unit": "kit", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA", "contents": "cloro + algicida + pH"}'),
('Kit mantenimiento pileta temporada HTH',     'HTH',       'mantenimiento-pileta', 55000.00, 'seed', 0.90, 85, TRUE, TRUE, 'piletas', '{"unit": "kit", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA", "contents": "cloro + algicida + pH + clarificador"}'),
('Kit apertura de pileta CTX',                 'CTX',       'mantenimiento-pileta', 62000.00, 'seed', 0.88, 83, TRUE, TRUE, 'piletas', '{"unit": "kit", "sku_prefix": "PILE", "research_date": "2026-04-25", "zone": "Posadas NEA", "contents": "alguicida + cloro + reductor pH + clarificador"}')

ON CONFLICT (name, business_type) DO NOTHING;
