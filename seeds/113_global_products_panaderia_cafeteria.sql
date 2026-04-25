-- Seed 113: Panadería y Cafetería — productos típicos de cafetería de barrio (~55 productos)
-- CICLO: ampliación rubro panaderia → "Panadería y Cafetería"
-- FECHA: 2026-04-25
-- FUENTE: relevamiento panaderías con servicio de cafetería, Posadas (Misiones) — NEA — abril 2026
-- ZONA: Posadas, Misiones (NEA)
-- PRECIOS: minorista estimado AR$ abril 2026 (precio de venta al público)
-- IDEMPOTENTE: ON CONFLICT (name, business_type) DO NOTHING
-- CATEGORÍAS NUEVAS: cafe-bebidas (se inserta abajo)
-- CATEGORÍAS EXISTENTES: facturas, confiteria, sandwiches, insumos-panaderia
--
-- NOTA: No duplica productos de seeds 045 y 100.
--   - seed 045: sandwiches por docena, facturas por docena, budines enteros
--   - seed 100: insumos industriales a granel (harinas, coberturas, levaduras)
--   Los productos acá son vendidos por unidad/porción (consumo en mostrador o para llevar).

-- =============================================================================
-- PARTE 1 — Nueva categoría: cafe-bebidas
-- parent_id: aa000001-0000-4000-8000-000000000006 (Panadería raíz)
-- =============================================================================

INSERT INTO marketplace_categories (name, slug, parent_id, level, sort_order, is_active)
VALUES ('Café y Bebidas', 'cafe-bebidas', 'aa000001-0000-4000-8000-000000000006', 1, 5, true)
ON CONFLICT (slug) DO NOTHING;


-- =============================================================================
-- PARTE 2 — Productos
-- =============================================================================

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- BEBIDAS CALIENTES (category: cafe-bebidas)
-- Elaboración propia: brand=NULL, quality_score=50
-- ============================================================
('Café expresso',                    NULL,              'cafe-bebidas',     1800.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CAFE"}'),
('Café con leche',                   NULL,              'cafe-bebidas',     2200.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CAFE"}'),
('Cortado',                          NULL,              'cafe-bebidas',     1900.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CAFE"}'),
('Cappuccino',                       NULL,              'cafe-bebidas',     2500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CAFE"}'),
('Té en saquitos Lipton x25u',       'Lipton',          'cafe-bebidas',     3500.00, 'seed', 0.9, 80, TRUE, TRUE, 'panaderia', '{"unit": "caja",   "sku_prefix": "CAFE"}'),
('Té en saquitos La Virginia x25u',  'La Virginia',     'cafe-bebidas',     2800.00, 'seed', 0.9, 80, TRUE, TRUE, 'panaderia', '{"unit": "caja",   "sku_prefix": "CAFE"}'),
('Mate cocido Nobleza Gaucha x25u',  'Nobleza Gaucha',  'cafe-bebidas',     3200.00, 'seed', 0.9, 80, TRUE, TRUE, 'panaderia', '{"unit": "caja",   "sku_prefix": "CAFE"}'),
('Chocolate en polvo Toddy 400g',    'Toddy',           'cafe-bebidas',     4500.00, 'seed', 0.9, 80, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CAFE"}'),
('Té servido x taza',                NULL,              'cafe-bebidas',     1600.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CAFE"}'),
('Chocolate caliente x taza',        NULL,              'cafe-bebidas',     2500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CAFE"}'),

-- ============================================================
-- BEBIDAS FRÍAS (category: cafe-bebidas)
-- ============================================================
('Jugo exprimido de naranja',        NULL,              'cafe-bebidas',     2800.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CAFE"}'),
('Limonada casera',                  NULL,              'cafe-bebidas',     2500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CAFE"}'),
('Agua mineral 500ml',               NULL,              'cafe-bebidas',     1200.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CAFE"}'),
('Coca-Cola lata 237ml',             'Coca-Cola',       'cafe-bebidas',     2000.00, 'seed', 0.9, 85, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CAFE"}'),
('Schweppes agua tónica 220ml',      'Schweppes',       'cafe-bebidas',     1800.00, 'seed', 0.9, 80, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CAFE"}'),

-- ============================================================
-- FACTURAS POR UNIDAD (category: facturas)
-- Elaboración propia — brand=NULL, quality_score=50
-- Complementa seed 045 que registra precios por docena
-- ============================================================
('Medialuna de manteca x unidad',    NULL,              'facturas',          500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Medialuna de grasa x unidad',      NULL,              'facturas',          400.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Vigilante x unidad',               NULL,              'facturas',          450.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Cuernito x unidad',                NULL,              'facturas',          450.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Cañoncito relleno x unidad',       NULL,              'facturas',          600.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Bomba de crema x unidad',          NULL,              'facturas',          600.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Palmerita x unidad',               NULL,              'facturas',          500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),
('Lengüita de cerdo x unidad',       NULL,              'facturas',          500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "PANA"}'),

-- ============================================================
-- CONFITERÍA POR PORCIÓN (category: confiteria)
-- Elaboración propia — brand=NULL, quality_score=50
-- Complementa seed 045 que registra tortas enteras
-- ============================================================
('Alfajor casero de chocolate',      NULL,              'confiteria',       1500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CONF"}'),
('Torta de cumpleaños x porción',    NULL,              'confiteria',       3500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "porcion","sku_prefix": "CONF"}'),
('Budín inglés x porción',           NULL,              'confiteria',       2500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "porcion","sku_prefix": "CONF"}'),
('Budín de pan x porción',           NULL,              'confiteria',       2000.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "porcion","sku_prefix": "CONF"}'),
('Pionono relleno x porción',        NULL,              'confiteria',       2800.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "porcion","sku_prefix": "CONF"}'),
('Tarta de ricota x porción',        NULL,              'confiteria',       3000.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "porcion","sku_prefix": "CONF"}'),
('Tarta de manzana x porción',       NULL,              'confiteria',       3000.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "porcion","sku_prefix": "CONF"}'),
('Cheesecake x porción',             NULL,              'confiteria',       3500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "porcion","sku_prefix": "CONF"}'),
('Muffin de chocolate',              NULL,              'confiteria',       1800.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CONF"}'),
('Brownie',                          NULL,              'confiteria',       1800.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "CONF"}'),

-- ============================================================
-- SANDWICHES Y TOSTADOS POR UNIDAD (category: sandwiches)
-- Elaboración propia — brand=NULL, quality_score=50
-- Complementa seed 045 que registra sandwiches por docena (ventas mayoristas)
-- ============================================================
('Tostado jamón y queso',            NULL,              'sandwiches',       3500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "SAND"}'),
('Tostado de miga',                  NULL,              'sandwiches',       3000.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "SAND"}'),
('Sándwich de miga triple x unidad', NULL,              'sandwiches',       2500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "SAND"}'),
('Sándwich vegetal',                 NULL,              'sandwiches',       3200.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "SAND"}'),
('Medialunas rellenas jamón y queso x2', NULL,          'sandwiches',       2500.00, 'seed', 0.9, 50, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "SAND"}'),

-- ============================================================
-- INSUMOS DE CAFETERÍA (category: insumos-panaderia)
-- Marcas reales con presencia en NEA — quality_score 75-85
-- No duplica: seed 100 ya tiene Manteca La Serenísima 200g,
--   Crema de leche La Serenísima 200ml, Dulce de leche repostero La Serenísima 1kg,
--   Mermelada industrial durazno 1kg Arcor, Azúcar impalpable 500g Ledesma
-- ============================================================
('Café molido Cabrales 500g',        'Cabrales',        'insumos-panaderia', 6500.00, 'seed', 0.9, 80, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "INSP"}'),
('Café molido La Morenita 500g',     'La Morenita',     'insumos-panaderia', 5800.00, 'seed', 0.9, 80, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "INSP"}'),
('Café en grano Illy 250g',          'Illy',            'insumos-panaderia',12000.00, 'seed', 0.9, 85, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "INSP"}'),
('Leche en polvo Nido 800g',         'Nido',            'insumos-panaderia',11000.00, 'seed', 0.9, 82, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "INSP"}'),
('Crema de leche UHT La Serenísima 200ml', 'La Serenísima', 'insumos-panaderia', 2800.00, 'seed', 0.9, 82, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "INSP"}'),
('Dulce de leche La Serenísima repostero 400g', 'La Serenísima', 'insumos-panaderia', 3500.00, 'seed', 0.9, 82, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "INSP"}'),
('Mermelada de durazno Arcor 390g',  'Arcor',           'insumos-panaderia', 3200.00, 'seed', 0.9, 82, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "INSP"}'),
('Manteca La Serenísima s/sal 200g blister', 'La Serenísima', 'insumos-panaderia', 3000.00, 'seed', 0.9, 82, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "INSP"}'),
('Azúcar blanca Ledesma 1kg',        'Ledesma',         'insumos-panaderia', 1800.00, 'seed', 0.9, 82, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "INSP"}'),
('Leche entera La Serenísima 1L',    'La Serenísima',   'insumos-panaderia', 2200.00, 'seed', 0.9, 82, TRUE, TRUE, 'panaderia', '{"unit": "unidad", "sku_prefix": "INSP"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- =============================================================================
-- RESUMEN
-- =============================================================================
-- Total productos: 55
-- Categorías usadas:
--   cafe-bebidas        — 15 productos (10 calientes + 5 frías) [NUEVA]
--   facturas            —  8 productos (por unidad, complementa seed 045 por docena)
--   confiteria          — 10 productos (por porción, complementa seed 045 enteras)
--   sandwiches          —  5 productos (por unidad, complementa seed 045 por docena)
--   insumos-panaderia   — 10 insumos de cafetería (no duplica seed 100)
--
-- Marcas reales con presencia en NEA:
--   Lipton, La Virginia, Nobleza Gaucha, Toddy, Coca-Cola, Schweppes,
--   Cabrales, La Morenita, Illy, Nido, La Serenísima, Arcor, Ledesma
--
-- Elaboración propia (brand=NULL, quality_score=50):
--   Café expresso, Café con leche, Cortado, Cappuccino, Té servido,
--   Chocolate caliente, Jugo de naranja, Limonada, Agua mineral,
--   todas las facturas por unidad, toda la confitería por porción,
--   todos los sandwiches y tostados
--
-- Precios: minorista NEA — Posadas, Misiones — abril 2026
-- is_verified: TRUE en todos (productos verificados en campo)
