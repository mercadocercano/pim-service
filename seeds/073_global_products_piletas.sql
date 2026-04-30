-- Seed 073: Piletas y piscinas — 28 productos
-- CICLO: cycle-004-brands-catalog-expansion
-- FECHA: 2026-04-21
-- FUENTE: catálogos distribuidores NEA (Posadas) + relevamiento de ferreterías 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- NOTA: Marcas deben estar en marketplace_brands (seed 065 primero).
--       business_type='piletas' si existe el rubro; sino puede usarse 'corralon' como
--       alternativa para corralones que venden insumos de pileta.
--       Se usa 'piletas' para separar el catálogo correctamente.

-- Categorías requeridas (creadas en seed 069):
--   cloro-quimicos, bombas-filtros-pileta, accesorios-pileta,
--   piletas-estructurales, mantenimiento-pileta

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- CLORO Y QUÍMICOS
-- ============================================================
('Cloro granulado 10kg',                     'Nataclor',  'cloro-quimicos',        55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Cloro granulado 10kg',                     'Clorotec',  'cloro-quimicos',        48000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Cloro granulado 1kg',                      'Nataclor',  'cloro-quimicos',         8000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Cloro tabletas tricloro 200g x5u',         'Nataclor',  'cloro-quimicos',        18000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Cloro tabletas 1kg',                       'Clorotec',  'cloro-quimicos',        15000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Cloro líquido 5L',                         'Clorotec',  'cloro-quimicos',        12000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Algicida antiverdin 1L',                   'Nataclor',  'cloro-quimicos',        14000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Algicida concentrado 1L',                  'Freshclor', 'cloro-quimicos',        12000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('pH menos 1kg',                             'Nataclor',  'cloro-quimicos',        12000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('pH más 1kg',                               'Nataclor',  'cloro-quimicos',        11000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Floculante clarificador 1L',               'Nataclor',  'cloro-quimicos',        11500.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Kit test análisis agua (50 test)',         'Nataclor',  'mantenimiento-pileta',   8000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),

-- ============================================================
-- BOMBAS Y FILTROS
-- ============================================================
('Bomba de recirculación 1/2HP',             'Difran',    'bombas-filtros-pileta', 120000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Bomba de recirculación 3/4HP',             'Difran',    'bombas-filtros-pileta', 165000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Filtro de arena 12 pulgadas',              'Difran',    'bombas-filtros-pileta', 180000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Filtro de arena 16 pulgadas',              'Difran',    'bombas-filtros-pileta', 240000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Arena filtrante sílice 25kg',              NULL,        'bombas-filtros-pileta',  18000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),

-- ============================================================
-- ACCESORIOS
-- ============================================================
('Manguera flexible 9m 38mm',                'Kokido',    'accesorios-pileta',     22000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Cepillo de fondo 45cm',                    'Kokido',    'accesorios-pileta',      6000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Manga aspiradora manual telescópica',      'Kokido',    'accesorios-pileta',     22000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Skimmer flotante con malla',               'Kokido',    'accesorios-pileta',      9500.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Termómetro de pileta flotante',            NULL,        'accesorios-pileta',      3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),

-- ============================================================
-- PILETAS ESTRUCTURALES
-- ============================================================
('Pileta estructural Bestway 3.05m x 76cm',  'Bestway',   'piletas-estructurales', 180000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Pileta estructural Bestway 3.66m x 91cm',  'Bestway',   'piletas-estructurales', 280000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Pileta inflable familiar 3m Bestway',      'Bestway',   'piletas-estructurales',  85000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),

-- ============================================================
-- MANTENIMIENTO Y COBERTORES
-- ============================================================
('Cobertor solar burbujas 4x8m',             'Deep Blue', 'mantenimiento-pileta',  45000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Cobertor solar burbujas 6x10m',            'Deep Blue', 'mantenimiento-pileta',  78000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Robot limpiafondos Dolphin E10',           'Maytronics','mantenimiento-pileta', 500000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb),
('Robot limpiafondos Dolphin E20',           'Maytronics','mantenimiento-pileta', 700000.00, 'seed', 0.5, 75, FALSE, TRUE, 'piletas', '{"unit": "unidad", "sku_prefix": "PIL"}'::jsonb)

ON CONFLICT DO NOTHING;

-- Total: 28 productos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
-- Variación cloro granulado: entre $4.800 y $5.500/kg (al 04/2026)
-- Piletas estructurales Bestway: presentes en ferreterías y supermercados de Posadas
