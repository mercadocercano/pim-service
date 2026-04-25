-- Seed 106: Productos globales — Juguetería (~65 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento jugueterías NEA (Posadas) + distribuidores 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: bebes-0-2, ninos-3-5, ninos-6-10, adolescentes, deportes-exterior, cotillon

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- BEBÉS 0-2 AÑOS
-- ============================================================
('Sonajero multicolor plástico',                 NULL,      'bebes-0-2',   2500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Mordillo silicona BPA free',                   NULL,      'bebes-0-2',   3200.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Peluche osito 30cm suave',                     NULL,      'bebes-0-2',   5500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Alfombra de juego plegable 150x200cm',         NULL,      'bebes-0-2',  22000.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Andador bebé plástico reforzado',              NULL,      'bebes-0-2',  28000.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Balde playero set 6 piezas',                   NULL,      'bebes-0-2',   3800.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "set",    "sku_prefix": "JUGG"}'),
('Juguete centro de actividades bebé',           NULL,      'bebes-0-2',  18000.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),

-- ============================================================
-- NIÑOS 3-5 AÑOS
-- ============================================================
('Rasti básico x100 piezas',                     'Rasti',   'ninos-3-5',  12000.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "caja",   "sku_prefix": "JUGG"}'),
('Plastilina Play-Doh colores básicos x4u',      'Play-Doh','ninos-3-5',   8500.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "set",    "sku_prefix": "JUGG"}'),
('Puzzle 24 piezas madera animales',             NULL,      'ninos-3-5',   5500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Muñeca Barbie básica Fashionistas Mattel',     'Mattel',  'ninos-3-5',  12000.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Auto Hot Wheels básico surtido',               'Hot Wheels','ninos-3-5', 1800.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Pistola de burbujas eléctrica',                NULL,      'ninos-3-5',   4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Kit cocina juguete 20 piezas',                 NULL,      'ninos-3-5',   9500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "set",    "sku_prefix": "JUGG"}'),
('Kit médico juguete maletín',                   NULL,      'ninos-3-5',   7500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Lápices de cera x16u Fisher-Price',            'Fisher-Price','ninos-3-5',2800.00,'seed', 0.8, 78, FALSE, TRUE, 'jugueteria', '{"unit": "caja",   "sku_prefix": "JUGG"}'),

-- ============================================================
-- NIÑOS 6-10 AÑOS
-- ============================================================
('Lego Creator 3en1 básico 200 piezas',          'Lego',    'ninos-6-10',  22000.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "caja",   "sku_prefix": "JUGG"}'),
('Rasti x300 piezas avanzado',                   'Rasti',   'ninos-6-10',  22000.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "caja",   "sku_prefix": "JUGG"}'),
('Monopoly clásico Hasbro',                      'Hasbro',  'ninos-6-10',  15000.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "caja",   "sku_prefix": "JUGG"}'),
('Batalla Naval Hasbro',                         'Hasbro',  'ninos-6-10',  12000.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "caja",   "sku_prefix": "JUGG"}'),
('Jenga clásico Hasbro',                         'Hasbro',  'ninos-6-10',   9500.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "caja",   "sku_prefix": "JUGG"}'),
('Uno cartas Mattel',                            'Mattel',  'ninos-6-10',   5500.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "caja",   "sku_prefix": "JUGG"}'),
('Bicicleta 16" niño acero',                     NULL,      'ninos-6-10', 85000.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Bicicleta 20" niño acero',                     NULL,      'ninos-6-10',105000.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Rollers ajustables talle 30-37',               NULL,      'ninos-6-10',  22000.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "par",    "sku_prefix": "JUGG"}'),
('Pelota de fútbol N°3 cuero sintético',         NULL,      'ninos-6-10',   8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Pelota de fútbol N°4 cuero sintético',         NULL,      'ninos-6-10',  10500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Dardo volador Nerf básico',                    'Hasbro',  'ninos-6-10',   8500.00, 'seed', 0.9, 82, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Muñeca Barbie con accesorios Mattel',          'Mattel',  'ninos-6-10',  18000.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),

-- ============================================================
-- ADOLESCENTES 10+ AÑOS
-- ============================================================
('Cluedo misterio Hasbro',                       'Hasbro',  'adolescentes', 18000.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "caja",   "sku_prefix": "JUGG"}'),
('Scrabble español Mattel',                      'Mattel',  'adolescentes', 15000.00, 'seed', 0.9, 85, FALSE, TRUE, 'jugueteria', '{"unit": "caja",   "sku_prefix": "JUGG"}'),
('Ajedrez plástico talle 3 genérico',            NULL,      'adolescentes',  8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Dardos acero con tablero 45cm',                NULL,      'adolescentes', 22000.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "set",    "sku_prefix": "JUGG"}'),
('Frisbee standard 27cm',                        NULL,      'adolescentes',  4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Raquetas de playa set x2 + pelota',            NULL,      'adolescentes',  8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "set",    "sku_prefix": "JUGG"}'),
('Mazo de cartas español 40 naipes',             NULL,      'adolescentes',  2800.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Puzzle adulto 1000 piezas paisaje',            NULL,      'adolescentes', 12000.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "caja",   "sku_prefix": "JUGG"}'),
('Trivia Perseguido juego preguntas',             NULL,      'adolescentes', 14000.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "caja",   "sku_prefix": "JUGG"}'),

-- ============================================================
-- COTILLÓN
-- ============================================================
('Globos látex colores x50u',                    NULL,      'cotillon',   3500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "bolsa",  "sku_prefix": "JUGG"}'),
('Gorros de cumpleaños x10u',                    NULL,      'cotillon',   2800.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "bolsa",  "sku_prefix": "JUGG"}'),
('Souvenir recuerdos x10u surtidos',             NULL,      'cotillon',   4500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "bolsa",  "sku_prefix": "JUGG"}'),
('Piñata de cartón decorada',                    NULL,      'cotillon',   8500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Confeti multicolor 100g',                      NULL,      'cotillon',   1800.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}'),
('Serpentinas papel x10u',                       NULL,      'cotillon',   2500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "pack",   "sku_prefix": "JUGG"}'),
('Vela cumpleaños x10u colores',                 NULL,      'cotillon',   1500.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "pack",   "sku_prefix": "JUGG"}'),
('Guirnalda "Feliz Cumpleaños" letras',          NULL,      'cotillon',   3800.00, 'seed', 0.7, 60, FALSE, TRUE, 'jugueteria', '{"unit": "unidad", "sku_prefix": "JUGG"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 63 productos
-- Categorías: bebes-0-2, ninos-3-5, ninos-6-10, adolescentes, deportes-exterior, cotillon
-- Marcas: Mattel, Hasbro, Rasti, Lego, Hot Wheels, Play-Doh, Fisher-Price — mayoría genérico
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
