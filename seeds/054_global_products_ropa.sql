-- Seed 054: Ropa — 95 productos reales argentinos
-- Generado: 2026-04-18
-- Fuente: global_products (v2.0)
-- ON CONFLICT DO NOTHING: idempotente, no duplica si ya existe por nombre+marca.

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES
-- ============================================================
-- REMERAS
-- ============================================================
('Remera manga corta algodón hombre', 'Kevingston', 'remeras', 28500.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Remera manga corta logo bordado hombre', 'Lacoste', 'remeras', 89000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Remera manga corta Dry-Fit hombre', 'Adidas', 'remeras', 45000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Remera manga corta básica mujer', 'Kosiuko', 'remeras', 32000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Remera manga larga rayada hombre', 'Tascani', 'remeras', 38500.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Remera manga larga térmica mujer', 'Mistral', 'remeras', 29500.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Musculosa deportiva mujer', 'Nike', 'remeras', 42000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Musculosa algodón hombre', 'Topper', 'remeras', 18500.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Remera polo clásica hombre', 'Lacoste', 'remeras', 115000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Remera polo piqué hombre', 'Kevingston', 'remeras', 42000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),

-- ============================================================
-- PANTALONES
-- ============================================================
('Jean recto clásico hombre', 'Levi''s', 'pantalones', 89000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Jean skinny mujer', 'Levi''s', 'pantalones', 85000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Jean clásico hombre', 'Wrangler', 'pantalones', 72000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Jean mom fit mujer', 'Kosiuko', 'pantalones', 68000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Jogger algodón frizado hombre', 'Puma', 'pantalones', 52000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Jogger deportivo mujer', 'Adidas', 'pantalones', 58000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Pantalón cargo hombre', 'Rusty', 'pantalones', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Pantalón de vestir hombre', 'Tascani', 'pantalones', 62000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Bermuda gabardina hombre', 'Kevingston', 'pantalones', 38000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Bermuda jean mujer', 'Wrangler', 'pantalones', 45000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),

-- ============================================================
-- CAMPERAS
-- ============================================================
('Campera rompeviento hombre', 'Adidas', 'camperas', 85000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Campera rompeviento liviana mujer', 'Nike', 'camperas', 92000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Campera inflable hombre', 'Puma', 'camperas', 125000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Campera inflable mujer', 'Mistral', 'camperas', 98000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Campera jean clásica hombre', 'Levi''s', 'camperas', 135000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Campera jean oversize mujer', 'Wrangler', 'camperas', 110000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Campera polar full zip hombre', 'Topper', 'camperas', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Parka impermeable hombre', 'Mistral', 'camperas', 145000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),

-- ============================================================
-- BUZOS
-- ============================================================
('Hoodie algodón frizado hombre', 'Nike', 'buzos', 78000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Hoodie oversize mujer', 'Adidas', 'buzos', 72000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Buzo canguro hombre', 'Puma', 'buzos', 58000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Buzo canguro mujer', 'Rusty', 'buzos', 52000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Buzo polar medio cierre hombre', 'Topper', 'buzos', 45000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Buzo polar mujer', 'Mistral', 'buzos', 42000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Sweater cuello redondo hombre', 'Kevingston', 'buzos', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Sweater cuello V mujer', 'Portsaid', 'buzos', 48000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Sweater lana hombre', 'Tascani', 'buzos', 68000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),

-- ============================================================
-- CAMISAS
-- ============================================================
('Camisa manga corta lino hombre', 'Kevingston', 'camisas', 52000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Camisa manga corta estampada hombre', 'Bensimon', 'camisas', 45000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Camisa manga larga lisa hombre', 'Tascani', 'camisas', 58000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Camisa manga larga mujer', 'Portsaid', 'camisas', 48000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Camisa formal slim fit hombre', 'Lacoste', 'camisas', 95000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Camisa casual cuadros hombre', 'Kevingston', 'camisas', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Camisa jean mujer', 'Kosiuko', 'camisas', 52000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),

-- ============================================================
-- ROPA INTERIOR
-- ============================================================
('Boxer algodón hombre', 'Eyelit', 'ropa-interior', 12500.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Boxer pack x3 hombre', 'Cocot', 'ropa-interior', 28000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Slip algodón hombre', 'Eyelit', 'ropa-interior', 8500.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Medias deportivas pack x3 hombre', 'Topper', 'ropa-interior', 15000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Medias invisibles pack x3 mujer', 'Cocot', 'ropa-interior', 12000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Calza térmica mujer', 'Caro Cuore', 'ropa-interior', 18500.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Corpiño deportivo mujer', 'Cocot', 'ropa-interior', 22000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Corpiño sin costura mujer', 'Luz de Mar', 'ropa-interior', 25000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Bombacha algodón pack x3 mujer', 'Caro Cuore', 'ropa-interior', 15000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),

-- ============================================================
-- DEPORTIVA
-- ============================================================
('Conjunto deportivo hombre', 'Adidas', 'deportiva', 115000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Conjunto deportivo mujer', 'Nike', 'deportiva', 125000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Calza deportiva larga mujer', 'Puma', 'deportiva', 48000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Calza deportiva capri mujer', 'Topper', 'deportiva', 35000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Short deportivo hombre', 'Adidas', 'deportiva', 38000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Short running mujer', 'Nike', 'deportiva', 42000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Top deportivo mujer', 'Puma', 'deportiva', 32000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Top deportivo tirantes mujer', 'Topper', 'deportiva', 25000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Camiseta fútbol entrenamiento hombre', 'Adidas', 'deportiva', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),

-- ============================================================
-- CALZADO
-- ============================================================
('Zapatilla urbana hombre', 'Topper', 'calzado', 65000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Zapatilla running hombre', 'Adidas', 'calzado', 125000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Zapatilla running mujer', 'Nike', 'calzado', 135000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Zapatilla skate hombre', 'DC Shoes', 'calzado', 95000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Zapatilla lona unisex', 'Topper', 'calzado', 42000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Zapatilla urbana mujer', 'Gaelle', 'calzado', 72000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Bota cuero hombre', 'Jaguar', 'calzado', 115000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Bota mujer caña media', 'Kosiuko', 'calzado', 125000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Sandalia tiras mujer', 'Reef', 'calzado', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Ojota hombre', 'Reef', 'calzado', 35000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Ojota mujer', 'Rip Curl', 'calzado', 32000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Zapato vestir hombre', 'Jaguar', 'calzado', 135000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),

-- ============================================================
-- ACCESORIOS
-- ============================================================
('Cinturón cuero hombre', 'Kevingston', 'accesorios', 32000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Cinturón trenzado hombre', 'Tascani', 'accesorios', 28000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Billetera cuero hombre', 'Rusty', 'accesorios', 38000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Billetera mujer', 'Portsaid', 'accesorios', 35000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Gorra trucker unisex', 'Rip Curl', 'accesorios', 25000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Gorra deportiva unisex', 'Adidas', 'accesorios', 28000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Bufanda lana unisex', 'Mistral', 'accesorios', 22000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Guantes polar unisex', 'Topper', 'accesorios', 18000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Mochila urbana', 'Rip Curl', 'accesorios', 65000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),

-- ============================================================
-- NIÑOS
-- ============================================================
('Remera algodón niño', 'Mimo', 'ninos', 18000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Remera estampada niña', 'Cheeky', 'ninos', 19500.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Remera manga larga niño', '47 Street', 'ninos', 22000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Pantalón jean niño', 'Mimo', 'ninos', 32000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Pantalón jogger niña', 'Cheeky', 'ninos', 28000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Campera polar niño', 'Mimo', 'ninos', 42000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Campera inflable niña', 'Cheeky', 'ninos', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Guardapolvo escolar blanco', 'Mimo', 'ninos', 35000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Guardapolvo escolar tableado', 'Cheeky', 'ninos', 38000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb),
('Zapatilla niño urbana', 'Topper', 'ninos', 45000.00, 'seed', 0.5, 75, FALSE, TRUE, 'ropa', '{"unit": "unidad", "sku_prefix": "ROPA"}'::jsonb)
ON CONFLICT DO NOTHING;
