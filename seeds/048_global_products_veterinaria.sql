-- Seed 048: Veterinaria — 120 productos reales argentinos
-- Generado: 2026-04-18
-- Fuente: global_products (v2.0)
-- ON CONFLICT DO NOTHING: idempotente.

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES
-- ============================================================
-- ALIMENTO PERRO
-- ============================================================
('Dog Chow Cachorro Razas Medianas y Grandes 21kg', 'Purina', 'alimento-perro', 62000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Dog Chow Adulto Razas Medianas y Grandes 21kg', 'Purina', 'alimento-perro', 58000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Dog Chow Adulto Razas Pequeñas 21kg', 'Purina', 'alimento-perro', 59000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Pro Plan Cachorro Razas Medianas 15kg', 'Purina', 'alimento-perro', 85000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Pro Plan Adulto Razas Grandes 15kg', 'Purina', 'alimento-perro', 82000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Pro Plan Senior +7 Razas Medianas 15kg', 'Purina', 'alimento-perro', 84000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Royal Canin Medium Adult 15kg', 'Royal Canin', 'alimento-perro', 95000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Royal Canin Medium Puppy 15kg', 'Royal Canin', 'alimento-perro', 98000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Royal Canin Mini Adult 7.5kg', 'Royal Canin', 'alimento-perro', 62000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Royal Canin Maxi Adult 15kg', 'Royal Canin', 'alimento-perro', 92000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Eukanuba Cachorro Razas Medianas 15kg', 'Eukanuba', 'alimento-perro', 78000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Eukanuba Adulto Razas Grandes 15kg', 'Eukanuba', 'alimento-perro', 75000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Pedigree Adulto Carne y Vegetales 21kg', 'Pedigree', 'alimento-perro', 42000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Pedigree Cachorro 21kg', 'Pedigree', 'alimento-perro', 44000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Sabrositos Adulto Carne 20kg', 'Sabrositos', 'alimento-perro', 28000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Ken-L Adulto 22kg', 'Ken-L', 'alimento-perro', 30000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Cooperación Adulto 25kg', 'Cooperación', 'alimento-perro', 22000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Old Prince Cachorro 15kg', 'Old Prince', 'alimento-perro', 48000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Vital Can Complete Adulto 20kg', 'Vital Can', 'alimento-perro', 35000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Excellent Adulto Razas Medianas y Grandes 20kg', 'Excellent', 'alimento-perro', 52000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Sieger Adulto Premium 20kg', 'Sieger', 'alimento-perro', 45000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Unik Adulto 22kg', 'Unik', 'alimento-perro', 32000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Nutrience Cachorro Razas Pequeñas 7.5kg', 'Nutrience', 'alimento-perro', 38000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Tiernitos Cachorro 15kg', 'Tiernitos', 'alimento-perro', 26000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),

-- ============================================================
-- ALIMENTO GATO
-- ============================================================
('Cat Chow Adulto Pollo 15kg', 'Purina', 'alimento-gato', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Cat Chow Gatito 15kg', 'Purina', 'alimento-gato', 57000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Pro Plan Gato Adulto Pollo 15kg', 'Purina', 'alimento-gato', 82000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Pro Plan Gato Castrado 15kg', 'Purina', 'alimento-gato', 84000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Royal Canin Indoor 27 7.5kg', 'Royal Canin', 'alimento-gato', 68000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Royal Canin Kitten 7.5kg', 'Royal Canin', 'alimento-gato', 72000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Royal Canin Sterilised 37 7.5kg', 'Royal Canin', 'alimento-gato', 70000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Whiskas Adulto Pescado 10kg', 'Whiskas', 'alimento-gato', 32000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Whiskas Gatito Pollo y Leche 10kg', 'Whiskas', 'alimento-gato', 34000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Old Prince Gato Adulto 7.5kg', 'Old Prince', 'alimento-gato', 42000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Excellent Gato Adulto 7.5kg', 'Excellent', 'alimento-gato', 45000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Vital Can Gato Adulto 7.5kg', 'Vital Can', 'alimento-gato', 28000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Cooperación Gato Adulto 10kg', 'Cooperación', 'alimento-gato', 18000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),

-- ============================================================
-- ALIMENTO HÚMEDO
-- ============================================================
('Pedigree Lata Adulto Carne 340g', 'Pedigree', 'alimento-humedo', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Pedigree Sobre Adulto Carne 100g', 'Pedigree', 'alimento-humedo', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Whiskas Sobre Adulto Atún 85g', 'Whiskas', 'alimento-humedo', 1600.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Whiskas Lata Paté Pollo 340g', 'Whiskas', 'alimento-humedo', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Pro Plan Sobre Gato Adulto Pollo 85g', 'Purina', 'alimento-humedo', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Pro Plan Sobre Perro Adulto Carne 100g', 'Purina', 'alimento-humedo', 2400.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Royal Canin Lata Recovery 195g', 'Royal Canin', 'alimento-humedo', 6500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Dog Selection Lata Carne 340g', 'Dog Selection', 'alimento-humedo', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),

-- ============================================================
-- SNACKS MASCOTAS
-- ============================================================
('Pedigree Dentastix Razas Medianas x7', 'Pedigree', 'snacks-mascotas', 6500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Pedigree Dentastix Razas Pequeñas x7', 'Pedigree', 'snacks-mascotas', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Whiskas Temptaciones Pollo 75g', 'Whiskas', 'snacks-mascotas', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Hueso Prensado Natural 15cm', 'Full Health', 'snacks-mascotas', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Galletas de Arroz para Perro 200g', 'Osspret', 'snacks-mascotas', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Palitos Saborizados Carne x10', 'Tiernitos', 'snacks-mascotas', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Oreja de Cerdo Deshidratada', 'Full Health', 'snacks-mascotas', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),

-- ============================================================
-- HIGIENE MASCOTAS
-- ============================================================
('Shampoo Osspret Pelo Largo 250ml', 'Osspret', 'higiene-mascotas', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Shampoo Osspret Pelo Corto 250ml', 'Osspret', 'higiene-mascotas', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Shampoo Osspret Cachorro 250ml', 'Osspret', 'higiene-mascotas', 5200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Shampoo Pet Society Pelo Blanco 300ml', 'Pet Society', 'higiene-mascotas', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Shampoo Full Health Avena 500ml', 'Full Health', 'higiene-mascotas', 6000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Colonia Osspret Macho 250ml', 'Osspret', 'higiene-mascotas', 4800.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Colonia Osspret Hembra 250ml', 'Osspret', 'higiene-mascotas', 4800.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Desodorante Ambiental Mascotas Labyes 350ml', 'Labyes', 'higiene-mascotas', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),

-- ============================================================
-- MEDICAMENTOS OTC (antiparasitarios, vitaminas, suplementos)
-- ============================================================
('Frontline Spray Antipulgas 250ml', 'Frontline', 'medicamentos-otc', 38000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Frontline Plus Perro 10-20kg Pipeta', 'Frontline', 'medicamentos-otc', 18000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Frontline Plus Perro 20-40kg Pipeta', 'Frontline', 'medicamentos-otc', 20000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Frontline Plus Gato Pipeta', 'Frontline', 'medicamentos-otc', 16000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('NexGard Perro 4-10kg Comprimido', 'NexGard', 'medicamentos-otc', 22000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('NexGard Perro 10-25kg Comprimido', 'NexGard', 'medicamentos-otc', 25000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('NexGard Perro 25-50kg Comprimido', 'NexGard', 'medicamentos-otc', 28000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Bravecto Perro 4.5-10kg Comprimido', 'Bravecto', 'medicamentos-otc', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Bravecto Perro 10-20kg Comprimido', 'Bravecto', 'medicamentos-otc', 60000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Bravecto Perro 20-40kg Comprimido', 'Bravecto', 'medicamentos-otc', 65000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Pipeta Labyes Fipronil Perro Grande', 'Labyes', 'medicamentos-otc', 8500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Pipeta Labyes Fipronil Gato', 'Labyes', 'medicamentos-otc', 6500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Collar Antipulgas Perro Osspret', 'Osspret', 'medicamentos-otc', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Collar Antipulgas Gato Osspret', 'Osspret', 'medicamentos-otc', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Vitaminas Labyes Vitaminthe Pasta 20g', 'Labyes', 'medicamentos-otc', 9500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Suplemento Articular Condrovet 60 comprimidos', 'Labyes', 'medicamentos-otc', 35000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),

-- ============================================================
-- ACCESORIOS PERRO
-- ============================================================
('Correa de Paseo Nylon 1.5m', 'Osspret', 'accesorios-perro', 6500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Correa Retráctil 5m hasta 20kg', 'Pet Society', 'accesorios-perro', 15000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Collar de Nylon Regulable Mediano', 'Osspret', 'accesorios-perro', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Collar de Nylon Regulable Grande', 'Osspret', 'accesorios-perro', 5200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Pretal Pechera Acolchada Mediano', 'Pet Society', 'accesorios-perro', 9500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Plato Acero Inoxidable 21cm', 'Osspret', 'accesorios-perro', 4800.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Plato Doble Acero Inoxidable', 'Full Health', 'accesorios-perro', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Cucha Plástica Grande', 'Osspret', 'accesorios-perro', 45000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Cucha Plástica Mediana', 'Osspret', 'accesorios-perro', 32000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Juguete Mordillo Hueso Goma', 'Full Health', 'accesorios-perro', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Juguete Pelota Tenis Perro x3', 'Pet Society', 'accesorios-perro', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Juguete Soga Resistente con Nudos', 'Full Health', 'accesorios-perro', 5800.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),

-- ============================================================
-- ACCESORIOS GATO
-- ============================================================
('Rascador Poste con Base 50cm', 'Pet Society', 'accesorios-gato', 18000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Rascador Torre 3 Niveles 80cm', 'Pet Society', 'accesorios-gato', 42000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Arenero Sanitario con Borde Alto', 'Osspret', 'accesorios-gato', 15000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Arenero Cerrado con Puerta', 'Pet Society', 'accesorios-gato', 28000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Arena Sanitaria Absorsol 3.6kg', 'Absorsol', 'accesorios-gato', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Arena Sanitaria Piedras Sanitarias 4kg', 'Piedras Sanitarias', 'accesorios-gato', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Arena Aglomerante Premium Cat 10kg', 'Premium Cat', 'accesorios-gato', 9500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Juguete Ratón con Plumas para Gato', 'Full Health', 'accesorios-gato', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Juguete Caña con Plumas para Gato', 'Pet Society', 'accesorios-gato', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Transportadora Plástica Mediana', 'Osspret', 'accesorios-gato', 25000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Transportadora Plástica Grande', 'Osspret', 'accesorios-gato', 32000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),

-- ============================================================
-- ROPA MASCOTAS
-- ============================================================
('Abrigo Polar Perro Talle 4', 'Osspret', 'ropa-mascotas', 12000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Abrigo Polar Perro Talle 6', 'Osspret', 'ropa-mascotas', 14000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Abrigo Polar Perro Talle 8', 'Osspret', 'ropa-mascotas', 16000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Piloto Impermeable Perro Talle 4', 'Pet Society', 'ropa-mascotas', 10000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Piloto Impermeable Perro Talle 6', 'Pet Society', 'ropa-mascotas', 12000.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Camiseta de Algodón Perro Talle 3', 'Full Health', 'ropa-mascotas', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Camiseta de Algodón Perro Talle 5', 'Full Health', 'ropa-mascotas', 8500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),

-- ============================================================
-- OTROS ANIMALES (peces, aves, roedores)
-- ============================================================
('Alimento Peces Tropicales TetraMin 52g', 'Tetra', 'otros-animales', 8500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Alimento Peces Goldfish TetraFin 62g', 'Tetra', 'otros-animales', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Alimento Peces Betta TetraBetta 27g', 'Tetra', 'otros-animales', 6500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Alimento Aves Mixtura Canarios 500g', 'Vita Force', 'otros-animales', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Alimento Aves Mixtura Loros 1kg', 'Vita Force', 'otros-animales', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Alimento Aves Alpiste Premium 500g', 'Vita Force', 'otros-animales', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Alimento Hamster y Roedores 750g', 'Vita Force', 'otros-animales', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Alimento Conejos 1kg', 'Vita Force', 'otros-animales', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Heno Natural para Roedores 500g', 'Full Health', 'otros-animales', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Piedra Mineral para Aves', 'Vita Force', 'otros-animales', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Viruta Prensada para Roedores 1kg', 'Full Health', 'otros-animales', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Alimento Tortugas Acuáticas Tetra ReptoMin 55g', 'Tetra', 'otros-animales', 7200.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb),
('Alimento Cobayo 1kg', 'Vita Force', 'otros-animales', 4800.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "kg", "sku_prefix": "VET"}'::jsonb),
('Acondicionador de Agua AquaSafe 100ml', 'Tetra', 'otros-animales', 6800.00, 'seed', 0.5, 75, FALSE, TRUE, 'veterinaria', '{"unit": "unidad", "sku_prefix": "VET"}'::jsonb)

ON CONFLICT DO NOTHING;
