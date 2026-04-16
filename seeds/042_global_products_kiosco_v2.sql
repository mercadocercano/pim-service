-- Seed 042: Kiosco — productos complementarios (v2.0 → global_products)
-- Generado: 2026-04-16
-- Complementa los 227 productos migrados desde templates via migración 040.
-- ON CONFLICT DO NOTHING: idempotente.

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES
-- ============================================================
-- PALITOS Y CHIZITOS (refuerzo: 6 existentes)
-- ============================================================
('Palitos salados Pehuamar 120g', 'Pehuamar', 'palitos-chizitos', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Chizitos Cheetos 75g', 'Cheetos', 'palitos-chizitos', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Conitos 3D 40g', 'Lays', 'palitos-chizitos', 1600.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Tutucas 75g', 'Arcor', 'palitos-chizitos', 1000.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- PILAS Y BATERIAS (refuerzo: 7 existentes)
-- ============================================================
('Pilas AAA Energizer x2', 'Energizer', 'pilas-baterias', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Pilas AA Eveready x4', 'Eveready', 'pilas-baterias', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Pila botón CR2032 Duracell', 'Duracell', 'pilas-baterias', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- ENCENDEDORES (refuerzo: 7 existentes)
-- ============================================================
('Encendedor Cricket', 'Cricket', 'encendedores', 900.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Fósforos Tres Patitos', 'Tres Patitos', 'encendedores', 300.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- FRUTOS SECOS (refuerzo: 7 existentes)
-- ============================================================
('Maní pelado salado Pehuamar 120g', 'Pehuamar', 'frutos-secos', 1600.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Mix de frutos secos Georgalos 100g', 'Georgalos', 'frutos-secos', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Almendras Georgalos 100g', 'Georgalos', 'frutos-secos', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- ENERGIZANTES (refuerzo: 7 existentes)
-- ============================================================
('Monster Energy 473ml', 'Monster', 'energizantes', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Speed Unlimited 473ml', 'Speed', 'energizantes', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Gatorade 500ml', 'Gatorade', 'energizantes', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Powerade 500ml', 'Powerade', 'energizantes', 2400.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- GOMITAS Y MALVAVISCOS (refuerzo: 8 existentes)
-- ============================================================
('Gomitas Mogul cerebritos 55g', 'Arcor', 'gomitas-malvaviscos', 1000.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Malvaviscos Guolis 100g', 'Guolis', 'gomitas-malvaviscos', 800.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Gomitas Trolli 100g', 'Trolli', 'gomitas-malvaviscos', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Gomitas Haribo ositos 80g', 'Haribo', 'gomitas-malvaviscos', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- GALLETITAS SALADAS (refuerzo: 8 existentes)
-- ============================================================
('Crackers Traviata 300g', 'Bagley', 'galletitas-saladas', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Club Social original 144g', 'Mondelez', 'galletitas-saladas', 1400.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Rex 75g', 'Bagley', 'galletitas-saladas', 600.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Pepitos salados 120g', 'Bagley', 'galletitas-saladas', 1100.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- AGUAS Y SABORIZADAS (refuerzo: 9 existentes)
-- ============================================================
('Agua Glaciar 600ml', 'Glaciar', 'aguas-saborizadas', 900.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Agua saborizada Villa del Sur 500ml', 'Villa del Sur', 'aguas-saborizadas', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Levité Manzana 500ml', 'Levité', 'aguas-saborizadas', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- GALLETITAS DULCES (refuerzo: 9 existentes)
-- ============================================================
('Pepitos 118g', 'Bagley', 'galletitas-dulces', 1400.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Toddy 126g', 'Toddy', 'galletitas-dulces', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Melba 120g', 'Bagley', 'galletitas-dulces', 1000.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Rumba 112g', 'Bagley', 'galletitas-dulces', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- ALFAJORES (refuerzo: 20 existentes, agregar marcas faltantes)
-- ============================================================
('Alfajor Terrabusi triple', 'Terrabusi', 'alfajores', 1000.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Cachafaz triple', 'Cachafaz', 'alfajores', 1100.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Alfajor Jorgelin', 'Jorgelin', 'alfajores', 600.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- CHOCOLATES (refuerzo: 12 existentes)
-- ============================================================
('Chocolate Cofler aireado 55g', 'Arcor', 'chocolates', 1600.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Chocolate Toblerone 100g', 'Toblerone', 'chocolates', 4000.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Rocklets Arcor 40g', 'Arcor', 'chocolates', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Kinder Bueno 43g', 'Kinder', 'chocolates', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- CARAMELOS Y CHICLES (refuerzo: 14 existentes)
-- ============================================================
('Caramelos Butter Toffees 150g', 'Arcor', 'caramelos-chicles', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Chicle Topline 7 unidades', 'Topline', 'caramelos-chicles', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Pastillas Menthoplus 30g', 'Arcor', 'caramelos-chicles', 800.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- GASEOSAS (refuerzo: 19 existentes, agregar presentaciones)
-- ============================================================
('Coca-Cola 2.25L', 'Coca-Cola', 'gaseosas', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Coca-Cola Zero 500ml', 'Coca-Cola', 'gaseosas', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('7Up 500ml', 'Pepsi', 'gaseosas', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Manaos Cola 2.25L', 'Manaos', 'gaseosas', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Schweppes Pomelo 500ml', 'Schweppes', 'gaseosas', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- CERVEZAS (refuerzo: 10 existentes)
-- ============================================================
('Cerveza Stella Artois 473ml lata', 'Stella Artois', 'cervezas', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Cerveza Andes Origen 473ml', 'Andes', 'cervezas', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Cerveza Imperial 473ml lata', 'Imperial', 'cervezas', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Cerveza Patagonia Amber Lager 473ml', 'Patagonia', 'cervezas', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- HIGIENE PERSONAL (refuerzo: 9 existentes)
-- ============================================================
('Preservativos Tulipán x3', 'Tulipán', 'higiene-personal', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Pañuelos descartables Day x10', 'Day', 'higiene-personal', 600.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Toallitas femeninas Kotex x8', 'Kotex', 'higiene-personal', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- PAPAS FRITAS (refuerzo: 9 existentes)
-- ============================================================
('Papas Pringles original 40g', 'Pringles', 'papas-fritas', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Doritos Dippas 50g', 'Lays', 'papas-fritas', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Papas Lays corte americano 47g', 'Lays', 'papas-fritas', 2100.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),

-- ============================================================
-- JUGOS (refuerzo: 11 existentes)
-- ============================================================
('Jugo Ades soja 200ml', 'Ades', 'jugos', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb),
('Jugo Baggio Pronto 1L', 'Baggio', 'jugos', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'kiosco', '{"unit": "unidad", "sku_prefix": "KIOSK"}'::jsonb)

ON CONFLICT DO NOTHING;
