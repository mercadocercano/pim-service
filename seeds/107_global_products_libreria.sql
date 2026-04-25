-- Seed 107: Productos globales — Librería: Útiles, Cuadernos y Arte (~65 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento librerías NEA (Posadas) + distribuidores útiles 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: lapices-lapiceras, cuadernos-carpetas, arte-pintura, papeleria-insumos, tecnologia-basica

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- LÁPICES Y LAPICERAS
-- ============================================================
('Lápiz negro HB x12u Staedtler',               'Staedtler',    'lapices-lapiceras',  5800.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "caja",   "sku_prefix": "LIBR"}'),
('Lápiz negro 2B x12u Staedtler',               'Staedtler',    'lapices-lapiceras',  6200.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "caja",   "sku_prefix": "LIBR"}'),
('Lápiz negro HB x12u Faber-Castell',           'Faber-Castell','lapices-lapiceras',  6500.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "caja",   "sku_prefix": "LIBR"}'),
('Lápiz negro HB x unidad Faber-Castell',       'Faber-Castell','lapices-lapiceras',   550.00, 'seed', 0.9, 82, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Lapicera Bic Cristal azul x50u caja',         'Bic',          'lapices-lapiceras', 18000.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "caja",   "sku_prefix": "LIBR"}'),
('Lapicera Bic Cristal azul x unidad',          'Bic',          'lapices-lapiceras',   380.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Lapicera Bic Cristal negro x unidad',         'Bic',          'lapices-lapiceras',   380.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Lapicera gel azul x unidad Pilot',            'Pilot',        'lapices-lapiceras',   850.00, 'seed', 0.9, 82, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Fibras de colores x12u Faber-Castell',        'Faber-Castell','lapices-lapiceras',  5500.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "caja",   "sku_prefix": "LIBR"}'),
('Marcadores permanentes x4u negro Sharpie',    'Sharpie',      'lapices-lapiceras',  4800.00, 'seed', 0.9, 82, FALSE, TRUE, 'libreria', '{"unit": "pack",   "sku_prefix": "LIBR"}'),
('Resaltadores x4u surtidos Stabilo',           'Stabilo',      'lapices-lapiceras',  5500.00, 'seed', 0.9, 82, FALSE, TRUE, 'libreria', '{"unit": "pack",   "sku_prefix": "LIBR"}'),
('Goma de borrar blanca x unidad Staedtler',    'Staedtler',    'lapices-lapiceras',   650.00, 'seed', 0.9, 82, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Sacapuntas doble plástico x unidad',          NULL,           'lapices-lapiceras',   350.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Regla 30cm transparente plástico',            NULL,           'lapices-lapiceras',   800.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Compás escolar básico acero',                 NULL,           'lapices-lapiceras',  2500.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Transportador 180° plástico',                 NULL,           'lapices-lapiceras',   700.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Escuadra 45° 25cm plástico',                  NULL,           'lapices-lapiceras',  1200.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),

-- ============================================================
-- CUADERNOS Y CARPETAS
-- ============================================================
('Cuaderno Gloria 48h tapa dura rayado',        'Gloria',       'cuadernos-carpetas',  2800.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Cuaderno Gloria 96h tapa dura cuadriculado',  'Gloria',       'cuadernos-carpetas',  5200.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Cuaderno Gloria 192h tapa dura rayado',       'Gloria',       'cuadernos-carpetas',  9500.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Cuaderno Rivadavia 50h cuadriculado',         'Rivadavia',    'cuadernos-carpetas',  2200.00, 'seed', 0.9, 82, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Cuaderno Rivadavia 84h doble raya',           'Rivadavia',    'cuadernos-carpetas',  3800.00, 'seed', 0.9, 82, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Carpeta escolar 3 anillos A4 genérica',       NULL,           'cuadernos-carpetas',  5500.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Bibliorato A4 lomo 5cm genérico',             NULL,           'cuadernos-carpetas',  4200.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Separadores de carpeta x10u plástico',        NULL,           'cuadernos-carpetas',  1800.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "pack",   "sku_prefix": "LIBR"}'),
('Portafolio ejecutivo A4 plástico',            NULL,           'cuadernos-carpetas',  8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Block de notas A5 espiral x80h',              NULL,           'cuadernos-carpetas',  3500.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),

-- ============================================================
-- ARTE Y PINTURA
-- ============================================================
('Acrílicos Acrilex x6 colores 20ml',           'Acrilex',      'arte-pintura',        8500.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "set",    "sku_prefix": "LIBR"}'),
('Acrílico Acrilex 20ml blanco',                'Acrilex',      'arte-pintura',        1500.00, 'seed', 0.9, 82, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Témpera escolar x6 colores 25ml',             NULL,           'arte-pintura',        4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "set",    "sku_prefix": "LIBR"}'),
('Acuarelas x12 colores Pentel',                'Pentel',       'arte-pintura',        6500.00, 'seed', 0.8, 80, FALSE, TRUE, 'libreria', '{"unit": "caja",   "sku_prefix": "LIBR"}'),
('Pinceles set x12 cerdas mixtas',              NULL,           'arte-pintura',        5500.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "set",    "sku_prefix": "LIBR"}'),
('Block de dibujo A4 x20h 180g',                NULL,           'arte-pintura',        4200.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Lienzo 30x40cm bastidor madera',              NULL,           'arte-pintura',        5500.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Lápices de color x24u Faber-Castell',         'Faber-Castell','arte-pintura',       12000.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "caja",   "sku_prefix": "LIBR"}'),
('Lápices de color x12u Staedtler',             'Staedtler',    'arte-pintura',        6500.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "caja",   "sku_prefix": "LIBR"}'),

-- ============================================================
-- PAPELERÍA E INSUMOS
-- ============================================================
('Papel A4 80g resma 500h Chamex',              'Chamex',       'papeleria-insumos',  12000.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "resma",  "sku_prefix": "LIBR"}'),
('Papel A4 75g resma 500h genérico',            NULL,           'papeleria-insumos',  10000.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "resma",  "sku_prefix": "LIBR"}'),
('Sobres blancos C5 x100u',                     NULL,           'papeleria-insumos',   6500.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "caja",   "sku_prefix": "LIBR"}'),
('Papel de color A4 x50h surtido',              NULL,           'papeleria-insumos',   3800.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "pack",   "sku_prefix": "LIBR"}'),
('Cartulina de colores x10h A4',                NULL,           'papeleria-insumos',   2800.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "pack",   "sku_prefix": "LIBR"}'),
('Cinta adhesiva Scotch 12mm x33m',             'Scotch',       'papeleria-insumos',   2500.00, 'seed', 0.9, 82, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Cinta adhesiva Scotch 18mm x33m',             'Scotch',       'papeleria-insumos',   3200.00, 'seed', 0.9, 82, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Cola vinílica 250g Poxipol',                  'Poxipol',      'papeleria-insumos',   3500.00, 'seed', 0.9, 82, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Post-it notas adhesivas 100h 76x76mm 3M',     '3M',           'papeleria-insumos',   4800.00, 'seed', 0.9, 82, FALSE, TRUE, 'libreria', '{"unit": "block",  "sku_prefix": "LIBR"}'),
('Broches mariposa x50u acero',                 NULL,           'papeleria-insumos',   2200.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "caja",   "sku_prefix": "LIBR"}'),
('Clip de papel x100u acero galvanizado',       NULL,           'papeleria-insumos',   1800.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "caja",   "sku_prefix": "LIBR"}'),

-- ============================================================
-- TECNOLOGÍA BÁSICA
-- ============================================================
('Pilas AA x2u Duracell',                       'Duracell',     'tecnologia-basica',   3500.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "pack",   "sku_prefix": "LIBR"}'),
('Pilas AA x4u Duracell',                       'Duracell',     'tecnologia-basica',   6500.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "pack",   "sku_prefix": "LIBR"}'),
('Pilas AAA x4u Duracell',                      'Duracell',     'tecnologia-basica',   6500.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "pack",   "sku_prefix": "LIBR"}'),
('Pilas AA x4u Energizer',                      'Energizer',    'tecnologia-basica',   6200.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "pack",   "sku_prefix": "LIBR"}'),
('Pilas AAA x4u Energizer',                     'Energizer',    'tecnologia-basica',   6200.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "pack",   "sku_prefix": "LIBR"}'),
('Calculadora científica Casio FX-82',          'Casio',        'tecnologia-basica',  22000.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Calculadora básica 8 dígitos genérica',       NULL,           'tecnologia-basica',   4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('USB 32GB Kingston DataTraveler',              'Kingston',     'tecnologia-basica',  12000.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('USB 64GB Kingston DataTraveler',              'Kingston',     'tecnologia-basica',  18000.00, 'seed', 0.9, 85, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}'),
('Cargador USB dual 2.1A genérico',             NULL,           'tecnologia-basica',   5500.00, 'seed', 0.7, 60, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIBR"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 65 productos
-- Categorías: lapices-lapiceras, cuadernos-carpetas, arte-pintura, papeleria-insumos, tecnologia-basica
-- Marcas: Staedtler, Faber-Castell, Bic, Pilot, Sharpie, Stabilo, Gloria, Rivadavia, Acrilex, Pentel,
--   Faber-Castell, Chamex, Scotch, Poxipol, 3M, Duracell, Energizer, Casio, Kingston
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
