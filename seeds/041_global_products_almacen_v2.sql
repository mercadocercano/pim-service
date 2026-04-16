-- Seed 041: Almacén — productos complementarios (v2.0 → global_products)
-- Generado: 2026-04-15
-- Complementa los 188 productos migrados desde templates via migración 040.
-- Refuerza categorías sub-representadas: harinas, detergentes, pan, shampoo, etc.
-- ON CONFLICT DO NOTHING: idempotente, no duplica si ya existe por nombre+marca.

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES
-- ============================================================
-- HARINAS Y PREMEZCLAS (refuerzo: solo 1 existente)
-- ============================================================
('Harina 000 Cañuelas 1kg', 'Cañuelas', 'harinas-premezclas', 1100.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Harina 0000 Pureza 1kg', 'Pureza', 'harinas-premezclas', 1250.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Harina leudante Blancaflor 1kg', 'Blancaflor', 'harinas-premezclas', 1400.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Premezcla para pizza Pureza 1kg', 'Pureza', 'harinas-premezclas', 1600.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Harina de maíz Maizena 500g', 'Maizena', 'harinas-premezclas', 1300.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Almidón de maíz Maizena 500g', 'Maizena', 'harinas-premezclas', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Azúcar Ledesma 1kg', 'Ledesma', 'harinas-premezclas', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Azúcar Chango 1kg', 'Chango', 'harinas-premezclas', 1450.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Polenta Presto Pronta 500g', 'Presto Pronta', 'harinas-premezclas', 900.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Levadura Levex seca 2 sobres', 'Levex', 'harinas-premezclas', 600.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- DETERGENTES Y JABONES (refuerzo: solo 1 existente)
-- ============================================================
('Detergente Ala 750ml', 'Ala', 'detergentes-jabones', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Detergente Magistral 500ml', 'Magistral', 'detergentes-jabones', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Detergente Cif Active Gel 750ml', 'Cif', 'detergentes-jabones', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Jabón en polvo Skip 800g', 'Skip', 'detergentes-jabones', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Jabón en polvo Ala 800g', 'Ala', 'detergentes-jabones', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Jabón en polvo Ace 800g', 'Ace', 'detergentes-jabones', 2600.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Suavizante Vivere 900ml', 'Vivere', 'detergentes-jabones', 2400.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Suavizante Comfort 900ml', 'Comfort', 'detergentes-jabones', 2600.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Jabón de tocador Lux 3x125g', 'Lux', 'detergentes-jabones', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Jabón blanco Federal 200g', 'Federal', 'detergentes-jabones', 500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- PAN ENVASADO (refuerzo: solo 2 existentes)
-- ============================================================
('Pan lactal Bimbo 450g', 'Bimbo', 'pan-envasado', 2400.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Pan lactal integral Fargo 350g', 'Fargo', 'pan-envasado', 2600.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Pan de hamburguesa Fargo x4', 'Fargo', 'pan-envasado', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Pan de pancho Bimbo x6', 'Bimbo', 'pan-envasado', 1600.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Tostadas Breviss light 200g', 'Breviss', 'pan-envasado', 1400.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Grisines Valente 200g', 'Valente', 'pan-envasado', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- SHAMPOO Y ACONDICIONADOR (refuerzo: solo 3 existentes)
-- ============================================================
('Shampoo Sedal reconstrucción 340ml', 'Sedal', 'shampoo-acondicionador', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Shampoo Head & Shoulders 375ml', 'Head & Shoulders', 'shampoo-acondicionador', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Shampoo Pantene 400ml', 'Pantene', 'shampoo-acondicionador', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Acondicionador Dove 400ml', 'Dove', 'shampoo-acondicionador', 4000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Shampoo Suave 930ml', 'Suave', 'shampoo-acondicionador', 3000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Shampoo Plusbelle 1L', 'Plusbelle', 'shampoo-acondicionador', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- ALMACEN SECO — yerba, café, té, dulces (refuerzo: 4 existentes)
-- ============================================================
('Yerba Taragüi 1kg', 'Taragüi', 'almacen-seco', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Yerba Playadito 1kg', 'Playadito', 'almacen-seco', 4800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Yerba Cruz de Malta 1kg', 'Cruz de Malta', 'almacen-seco', 4000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Yerba Cbsé Silueta 1kg', 'Cbsé', 'almacen-seco', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Café Cabrales molido 500g', 'Cabrales', 'almacen-seco', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Café La Virginia molido 500g', 'La Virginia', 'almacen-seco', 5000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Té La Virginia 25 saquitos', 'La Virginia', 'almacen-seco', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Té Taragüi 25 saquitos', 'Taragüi', 'almacen-seco', 1100.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Dulce de batata Arcor 500g', 'Arcor', 'almacen-seco', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Dulce de membrillo Arcor 500g', 'Arcor', 'almacen-seco', 1700.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Mermelada de frutilla BC La Campagnola 454g', 'La Campagnola', 'almacen-seco', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Mermelada de durazno BC La Campagnola 454g', 'La Campagnola', 'almacen-seco', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Sal fina Dos Anclas 500g', 'Dos Anclas', 'almacen-seco', 600.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Sal gruesa Dos Anclas 1kg', 'Dos Anclas', 'almacen-seco', 500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- SNACKS SALADOS (refuerzo: 4 existentes)
-- ============================================================
('Papas fritas Lays clásicas 270g', 'Lays', 'snacks-salados-alm', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Papas fritas Pringles original 124g', 'Pringles', 'snacks-salados-alm', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Chizitos Cheetos 200g', 'Cheetos', 'snacks-salados-alm', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Maní salado Marolio 120g', 'Marolio', 'snacks-salados-alm', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Palitos salados Pehuamar 120g', 'Pehuamar', 'snacks-salados-alm', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Tutucas 75g', 'Arcor', 'snacks-salados-alm', 1000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- CERVEZAS Y VINOS (refuerzo: 4 existentes)
-- ============================================================
('Cerveza Quilmes 1L retornable', 'Quilmes', 'cervezas-vinos', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Cerveza Brahma 1L retornable', 'Brahma', 'cervezas-vinos', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Cerveza Stella Artois 1L', 'Stella Artois', 'cervezas-vinos', 3000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Cerveza Andes Origen 473ml lata', 'Andes', 'cervezas-vinos', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Vino tinto Estancia Mendoza 750ml', 'Estancia Mendoza', 'cervezas-vinos', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Vino Termidor tinto 1L', 'Termidor', 'cervezas-vinos', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Fernet Branca 750ml', 'Branca', 'cervezas-vinos', 8500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- LAVANDINA Y DESINFECTANTES (refuerzo: 5 existentes)
-- ============================================================
('Lavandina Ayudín 1L', 'Ayudín', 'lavandina-desinfectantes', 800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Lavandina Querubin 2L', 'Querubín', 'lavandina-desinfectantes', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Desinfectante Lysoform 900ml', 'Lysoform', 'lavandina-desinfectantes', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Limpiador Cif crema 750ml', 'Cif', 'lavandina-desinfectantes', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Limpiador Procenex pisos 1.8L', 'Procenex', 'lavandina-desinfectantes', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- ARROZ Y LEGUMBRES (refuerzo: 6 existentes)
-- ============================================================
('Arroz Gallo Oro 1kg', 'Gallo Oro', 'arroz-legumbres', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Arroz Dos Hermanos 1kg', 'Dos Hermanos', 'arroz-legumbres', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Arroz Lucchetti parboil 1kg', 'Lucchetti', 'arroz-legumbres', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Porotos pallares Marolio 500g', 'Marolio', 'arroz-legumbres', 1300.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Garbanzos Marolio 500g', 'Marolio', 'arroz-legumbres', 1400.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- PASTAS SECAS (refuerzo: 5 existentes)
-- ============================================================
('Fideos tirabuzón Matarazzo 500g', 'Matarazzo', 'pastas-secas', 1600.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Fideos mostachol Don Vicente 500g', 'Don Vicente', 'pastas-secas', 1400.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Fideos coditos Marolio 500g', 'Marolio', 'pastas-secas', 1000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Fideos spaghetti Lucchetti 500g', 'Lucchetti', 'pastas-secas', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Tapas de empanada La Salteña 12u', 'La Salteña', 'pastas-secas', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Tapas de tarta La Salteña', 'La Salteña', 'pastas-secas', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- ACEITES Y VINAGRES (refuerzo: 7 existentes)
-- ============================================================
('Aceite girasol Cada Día 1.5L', 'Cada Día', 'aceites-vinagres', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Aceite mezcla Cañuelas 1.5L', 'Cañuelas', 'aceites-vinagres', 3000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Vinagre de manzana Menoyo 1L', 'Menoyo', 'aceites-vinagres', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Aceite de oliva Cocinero 500ml', 'Cocinero', 'aceites-vinagres', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- CONSERVAS Y ENLATADOS (refuerzo: 13 existentes, agregar pocos)
-- ============================================================
('Puré de tomate Arcor 520g', 'Arcor', 'conservas-enlatados', 1100.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Choclo cremoso Marolio 340g', 'Marolio', 'conservas-enlatados', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Sardinas en aceite Gomes da Costa', 'Gomes da Costa', 'conservas-enlatados', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Mostaza Savora 250g', 'Savora', 'conservas-enlatados', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Ketchup Hellmanns 250g', 'Hellmanns', 'conservas-enlatados', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- JABONES Y DESODORANTES (refuerzo: 4 existentes)
-- ============================================================
('Desodorante Rexona 150ml', 'Rexona', 'jabones-desodorantes', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Desodorante Axe 150ml', 'Axe', 'jabones-desodorantes', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Desodorante Dove 150ml', 'Dove', 'jabones-desodorantes', 4000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Jabón líquido Espadol 250ml', 'Espadol', 'jabones-desodorantes', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Pasta dental Colgate 90g', 'Colgate', 'jabones-desodorantes', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Cepillo dental Colgate', 'Colgate', 'jabones-desodorantes', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- FIAMBRES Y EMBUTIDOS (refuerzo: 7 existentes)
-- ============================================================
('Salchicha Vienissima 6 unidades', 'Vienissima', 'fiambres-embutidos', 2400.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Salchicha Paladini 6 unidades', 'Paladini', 'fiambres-embutidos', 2700.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Jamón crudo Cagnoli feteado 120g', 'Cagnoli', 'fiambres-embutidos', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Salame tandilero Cagnoli 130g', 'Cagnoli', 'fiambres-embutidos', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Mortadela Paladini 200g', 'Paladini', 'fiambres-embutidos', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- QUESOS Y MANTECA (refuerzo: 6 existentes)
-- ============================================================
('Queso rallado La Serenísima 150g', 'La Serenísima', 'quesos-manteca', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Queso untable Finlandia 200g', 'Finlandia', 'quesos-manteca', 3000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Queso crema Mendicrim 300g', 'Mendicrim', 'quesos-manteca', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Manteca La Serenísima 200g', 'La Serenísima', 'quesos-manteca', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- JUGOS Y POLVOS (refuerzo: 7 existentes)
-- ============================================================
('Jugo Cepita del Valle 1L', 'Cepita', 'jugos-polvos', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Jugo en polvo Zuko limón', 'Zuko', 'jugos-polvos', 500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Cacao Toddy 360g', 'Toddy', 'jugos-polvos', 3000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Chocolatada Nesquik 1L', 'Nesquik', 'jugos-polvos', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- PAPEL HIGIENICO Y SERVILLETAS (refuerzo: 9 existentes, pocos más)
-- ============================================================
('Papel higiénico Higienol Exportación 4x30m', 'Higienol', 'papel-higienico', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Rollo de cocina Sussex 3x50', 'Sussex', 'papel-higienico', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Servilletas Elite 50u', 'Elite', 'papel-higienico', 1000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- GASEOSAS Y AGUAS (refuerzo: 8 existentes)
-- ============================================================
('Sprite 2.25L', 'Sprite', 'gaseosas-aguas', 2600.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Fanta naranja 2.25L', 'Fanta', 'gaseosas-aguas', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Pepsi 2.25L', 'Pepsi', 'gaseosas-aguas', 2400.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Agua mineral Glaciar 2L', 'Glaciar', 'gaseosas-aguas', 1000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Soda Ivess 2L', 'Ivess', 'gaseosas-aguas', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- GALLETITAS (refuerzo: 9 existentes)
-- ============================================================
('Galletitas Criollitas 100g', 'Bagley', 'galletitas-almacen', 800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Galletitas Traviata 300g', 'Bagley', 'galletitas-almacen', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Galletitas Diversión Variedad 400g', 'Arcor', 'galletitas-almacen', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Galletitas Rex 75g', 'Bagley', 'galletitas-almacen', 600.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Galletitas Oreo 117g', 'Oreo', 'galletitas-almacen', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),

-- ============================================================
-- PAÑALES (refuerzo: 8 existentes, agregar pocos)
-- ============================================================
('Pañales Pampers Premium Care M x40', 'Pampers', 'panales', 12000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb),
('Toallitas húmedas Huggies x48', 'Huggies', 'panales', 3000.00, 'seed', 0.5, 75, FALSE, TRUE, 'almacen', '{"unit": "unidad", "sku_prefix": "ALMAC"}'::jsonb)

ON CONFLICT DO NOTHING;
