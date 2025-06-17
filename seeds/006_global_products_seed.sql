-- ================================================
-- GLOBAL PRODUCTS INITIAL SEED
-- Seed 006: Productos populares argentinos para catálogo global
-- ================================================

-- Primero, obtener IDs de categorías para vincular productos
-- Estas categorías ya fueron creadas en el seed 003

DO $$
DECLARE
    cat_bebidas UUID;
    cat_lacteos UUID;
    cat_alimentos UUID;
    cat_limpieza UUID;
    cat_electronica UUID;
    cat_celulares UUID;
    cat_snacks UUID;
    cat_cuidado_personal UUID;
BEGIN
    -- Obtener IDs de categorías
    SELECT id INTO cat_bebidas FROM marketplace_categories WHERE slug = 'bebidas' AND parent_id IS NOT NULL LIMIT 1;
    SELECT id INTO cat_lacteos FROM marketplace_categories WHERE slug = 'lacteos' AND parent_id IS NOT NULL LIMIT 1;
    SELECT id INTO cat_alimentos FROM marketplace_categories WHERE slug = 'alimentos' AND parent_id IS NOT NULL LIMIT 1;
    SELECT id INTO cat_limpieza FROM marketplace_categories WHERE slug = 'limpieza' AND parent_id IS NOT NULL LIMIT 1;
    SELECT id INTO cat_electronica FROM marketplace_categories WHERE slug = 'electronica' AND parent_id IS NOT NULL LIMIT 1;
    SELECT id INTO cat_celulares FROM marketplace_categories WHERE slug = 'celulares' AND parent_id IS NOT NULL LIMIT 1;
    SELECT id INTO cat_snacks FROM marketplace_categories WHERE slug = 'snacks' AND parent_id IS NOT NULL LIMIT 1;
    SELECT id INTO cat_cuidado_personal FROM marketplace_categories WHERE slug = 'cuidado-personal' AND parent_id IS NOT NULL LIMIT 1;

    -- PRODUCTOS FMCG - BEBIDAS
    -- ======================
    
    INSERT INTO global_products (
        ean, name, brand, manufacturer, marketplace_category_id, product_type,
        specifications, source_type, source_confidence, verification_status,
        quality_score, popularity_rank, is_active
    ) VALUES
    
    -- Coca Cola
    ('7790895000126', 'Coca Cola 600ml', 'Coca Cola', 'Coca Cola FEMSA', cat_bebidas, 'fmcg', 
     '{"volume": "600ml", "container": "botella plástica", "calories_per_100ml": 42, "sugar_per_100ml": "10.6g"}', 
     'manual', 0.95, 'verified', 0.98, 1, true),
    
    ('7790895001239', 'Coca Cola 1.5L', 'Coca Cola', 'Coca Cola FEMSA', cat_bebidas, 'fmcg',
     '{"volume": "1.5L", "container": "botella plástica", "calories_per_100ml": 42, "sugar_per_100ml": "10.6g"}',
     'manual', 0.95, 'verified', 0.97, 2, true),
    
    ('7790895000567', 'Coca Cola Zero 600ml', 'Coca Cola', 'Coca Cola FEMSA', cat_bebidas, 'fmcg',
     '{"volume": "600ml", "container": "botella plástica", "calories_per_100ml": 0, "sweetener": "aspartame"}',
     'manual', 0.95, 'verified', 0.96, 3, true),
    
    -- Pepsi
    ('7790742000101', 'Pepsi 600ml', 'Pepsi', 'PepsiCo', cat_bebidas, 'fmcg',
     '{"volume": "600ml", "container": "botella plástica", "calories_per_100ml": 43}',
     'manual', 0.92, 'verified', 0.94, 4, true),
    
    -- Sprite
    ('7790895000890', 'Sprite 600ml', 'Sprite', 'Coca Cola FEMSA', cat_bebidas, 'fmcg',
     '{"volume": "600ml", "container": "botella plástica", "flavor": "lima-limón"}',
     'manual', 0.93, 'verified', 0.93, 5, true),
    
    -- Fanta
    ('7790895001456', 'Fanta Naranja 600ml', 'Fanta', 'Coca Cola FEMSA', cat_bebidas, 'fmcg',
     '{"volume": "600ml", "container": "botella plástica", "flavor": "naranja"}',
     'manual', 0.91, 'verified', 0.92, 6, true),
    
    -- Agua
    ('7790375000234', 'Agua Villa del Sur 500ml', 'Villa del Sur', 'Danone', cat_bebidas, 'fmcg',
     '{"volume": "500ml", "container": "botella plástica", "type": "agua mineral"}',
     'manual', 0.90, 'verified', 0.89, 10, true),
    
    -- PRODUCTOS FMCG - LÁCTEOS
    -- ========================
    
    -- La Serenísima
    ('7790315000123', 'Leche La Serenísima Entera 1L', 'La Serenísima', 'Mastellone', cat_lacteos, 'fmcg',
     '{"volume": "1L", "type": "leche entera", "fat_content": "3%", "protein_per_100ml": "3.2g"}',
     'manual', 0.96, 'verified', 0.95, 7, true),
    
    ('7790315001456', 'Yogur La Serenísima Natural 190g', 'La Serenísima', 'Mastellone', cat_lacteos, 'fmcg',
     '{"weight": "190g", "type": "yogur natural", "fat_content": "2.5%"}',
     'manual', 0.94, 'verified', 0.91, 12, true),
    
    -- Sancor
    ('7790742123456', 'Leche Sancor Descremada 1L', 'Sancor', 'SanCor', cat_lacteos, 'fmcg',
     '{"volume": "1L", "type": "leche descremada", "fat_content": "0.5%"}',
     'manual', 0.93, 'verified', 0.90, 13, true),
    
    -- Ilolay
    ('7790070000789', 'Manteca Ilolay 200g', 'Ilolay', 'Ilolay', cat_lacteos, 'fmcg',
     '{"weight": "200g", "type": "manteca", "fat_content": "82%"}',
     'manual', 0.92, 'verified', 0.88, 15, true),
    
    -- PRODUCTOS FMCG - SNACKS Y GOLOSINAS
    -- ===================================
    
    -- Arcor
    ('7790350000456', 'Bon o Bon Arcor', 'Bon o Bon', 'Arcor', cat_snacks, 'fmcg',
     '{"weight": "15g", "type": "chocolate", "ingredients": ["chocolate", "maní", "oblea"]}',
     'manual', 0.94, 'verified', 0.93, 8, true),
    
    ('7790350001789', 'Caramelos Menthoplus', 'Menthoplus', 'Arcor', cat_snacks, 'fmcg',
     '{"weight": "50g", "type": "caramelos", "flavor": "eucalipto"}',
     'manual', 0.89, 'verified', 0.85, 20, true),
    
    -- Bagley
    ('7790070123456', 'Galletitas Tita', 'Tita', 'Bagley', cat_snacks, 'fmcg',
     '{"weight": "168g", "type": "galletitas dulces"}',
     'manual', 0.91, 'verified', 0.87, 18, true),
    
    ('7790070456789', 'Galletitas Opera', 'Opera', 'Bagley', cat_snacks, 'fmcg',
     '{"weight": "55g", "type": "galletitas rellenas", "flavor": "chocolate"}',
     'manual', 0.90, 'verified', 0.86, 19, true),
    
    -- PRODUCTOS FMCG - ALIMENTOS BÁSICOS
    -- ==================================
    
    -- Molinos
    ('7790742000890', 'Aceite Cocinero 900ml', 'Cocinero', 'Molinos', cat_alimentos, 'fmcg',
     '{"volume": "900ml", "type": "aceite girasol"}',
     'manual', 0.92, 'verified', 0.89, 14, true),
    
    -- Marolio
    ('7790070000123', 'Fideos Marolio Tallarines 500g', 'Marolio', 'Marolio', cat_alimentos, 'fmcg',
     '{"weight": "500g", "type": "pasta", "shape": "tallarines"}',
     'manual', 0.88, 'verified', 0.84, 22, true),
    
    -- PRODUCTOS FMCG - LIMPIEZA
    -- =========================
    
    -- Unilever
    ('7791293000234', 'Detergente Ala 800g', 'Ala', 'Unilever', cat_limpieza, 'fmcg',
     '{"weight": "800g", "type": "detergente en polvo", "fragrance": "limón"}',
     'manual', 0.91, 'verified', 0.88, 16, true),
    
    ('7791293001567', 'Suavizante Comfort 500ml', 'Comfort', 'Unilever', cat_limpieza, 'fmcg',
     '{"volume": "500ml", "type": "suavizante", "fragrance": "caricias"}',
     'manual', 0.89, 'verified', 0.85, 21, true),
    
    -- P&G
    ('7791234000456', 'Shampoo Head & Shoulders 400ml', 'Head & Shoulders', 'P&G', cat_cuidado_personal, 'fmcg',
     '{"volume": "400ml", "type": "shampoo anticaspa"}',
     'manual', 0.93, 'verified', 0.90, 17, true),
    
    -- PRODUCTOS INDUSTRIALIZADOS - ELECTRÓNICOS
    -- =========================================
    
    -- Samsung
    ('8801643000123', 'Samsung Galaxy A54 128GB', 'Samsung', 'Samsung Electronics', cat_celulares, 'industrialized',
     '{"storage": "128GB", "ram": "6GB", "screen_size": "6.4\"", "camera": "50MP", "battery": "5000mAh", "os": "Android 13"}',
     'manual', 0.97, 'verified', 0.99, 1, true),
    
    ('8801643001456', 'Samsung Galaxy A34 128GB', 'Samsung', 'Samsung Electronics', cat_celulares, 'industrialized',
     '{"storage": "128GB", "ram": "6GB", "screen_size": "6.6\"", "camera": "48MP", "battery": "5000mAh", "os": "Android 13"}',
     'manual', 0.96, 'verified', 0.95, 5, true),
    
    -- Apple (distribuidor oficial Argentina)
    ('0194252000123', 'iPhone 14 128GB', 'iPhone', 'Apple', cat_celulares, 'industrialized',
     '{"storage": "128GB", "ram": "6GB", "screen_size": "6.1\"", "camera": "12MP dual", "battery": "3279mAh", "os": "iOS 16"}',
     'manual', 0.98, 'verified', 0.97, 2, true),
    
    -- Xiaomi
    ('6934177000456', 'Xiaomi Redmi Note 12 128GB', 'Redmi', 'Xiaomi', cat_celulares, 'industrialized',
     '{"storage": "128GB", "ram": "4GB", "screen_size": "6.67\"", "camera": "50MP", "battery": "5000mAh", "os": "MIUI 14"}',
     'manual', 0.94, 'verified', 0.92, 8, true),
    
    -- Motorola
    ('0723755000789', 'Motorola Moto G73 128GB', 'Moto G', 'Motorola', cat_celulares, 'industrialized',
     '{"storage": "128GB", "ram": "8GB", "screen_size": "6.5\"", "camera": "50MP", "battery": "5000mAh", "os": "Android 13"}',
     'manual', 0.93, 'verified', 0.91, 11, true);

    -- PRODUCTOS LOCALES CON EAN
    -- =========================
    
    INSERT INTO global_products (
        ean, name, brand, manufacturer, marketplace_category_id, product_type,
        specifications, source_type, source_confidence, verification_status,
        quality_score, popularity_rank, is_active
    ) VALUES
    
    -- Dulce de Leche La Serenísima
    ('7790315002345', 'Dulce de Leche La Serenísima 400g', 'La Serenísima', 'Mastellone', cat_alimentos, 'local_brand',
     '{"weight": "400g", "type": "dulce de leche tradicional", "origin": "Argentina"}',
     'manual', 0.94, 'verified', 0.89, 25, true),
    
    -- Alfajores Havanna
    ('7790070789123', 'Alfajor Havanna Mixto', 'Havanna', 'Havanna', cat_snacks, 'local_brand',
     '{"weight": "70g", "type": "alfajor", "filling": "dulce de leche", "coating": "chocolate"}',
     'manual', 0.95, 'verified', 0.93, 9, true),
    
    -- Yerba Mate
    ('7790742345678', 'Yerba Mate La Merced 1kg', 'La Merced', 'Las Marías', cat_alimentos, 'local_brand',
     '{"weight": "1kg", "type": "yerba mate", "blend": "suave", "origin": "Misiones"}',
     'manual', 0.91, 'verified', 0.87, 23, true),
    
    ('7790123456789', 'Yerba Mate Taragui 1kg', 'Taragui', 'Las Marías', cat_alimentos, 'local_brand',
     '{"weight": "1kg", "type": "yerba mate", "blend": "tradicional", "origin": "Corrientes"}',
     'manual', 0.90, 'verified', 0.86, 24, true);

    RAISE NOTICE 'Productos globales insertados correctamente: % productos', 
        (SELECT COUNT(*) FROM global_products WHERE source_type = 'manual');

END $$;

-- Crear algunas variaciones de productos
INSERT INTO global_product_variations (global_product_id, variation_type, variation_value, ean, specifications_diff)
SELECT 
    gp.id,
    'size',
    '2L',
    '7790895002345',
    '{"volume": "2L", "price_difference": 200}'
FROM global_products gp 
WHERE gp.ean = '7790895000126' -- Coca Cola 600ml
LIMIT 1;

INSERT INTO global_product_variations (global_product_id, variation_type, variation_value, ean, specifications_diff)
SELECT 
    gp.id,
    'flavor',
    'Lima-Limón',
    '7790895003456',
    '{"flavor": "lima-limón", "price_difference": 0}'
FROM global_products gp 
WHERE gp.ean = '7790895000126' -- Coca Cola 600ml para variación de sabor
LIMIT 1;

-- Agregar productos sugeridos a los templates de tipos de negocio
-- Esto vincula productos específicos con cada tipo de negocio

DO $$
DECLARE
    template_id UUID;
    almacen_products JSONB;
    farmacia_products JSONB;
    kiosco_products JSONB;
    super_products JSONB;
    electronica_products JSONB;
BEGIN
    -- ALMACÉN - Productos básicos de alimentación
    SELECT id INTO template_id FROM business_type_templates btt
    JOIN business_types bt ON btt.business_type_id = bt.id
    WHERE bt.code = 'almacen' LIMIT 1;
    
    IF template_id IS NOT NULL THEN
        SELECT jsonb_agg(gp.id) INTO almacen_products
        FROM global_products gp
        WHERE gp.ean IN (
            '7790895000126', -- Coca Cola 600ml
            '7790315000123', -- Leche La Serenísima
            '7790742000890', -- Aceite Cocinero
            '7790070000123', -- Fideos Marolio
            '7790742345678', -- Yerba Mate La Merced
            '7790315002345'  -- Dulce de Leche
        );
        
        INSERT INTO business_type_product_templates (
            business_type_template_id, 
            suggested_products,
            priority_brands,
            max_products_per_category
        ) VALUES (
            template_id,
            almacen_products,
            '["Coca Cola", "La Serenísima", "Marolio", "Molinos"]',
            30
        ) ON CONFLICT (business_type_template_id) DO UPDATE SET
            suggested_products = EXCLUDED.suggested_products,
            priority_brands = EXCLUDED.priority_brands;
    END IF;
    
    -- FARMACIA - Productos de cuidado personal y salud
    SELECT id INTO template_id FROM business_type_templates btt
    JOIN business_types bt ON btt.business_type_id = bt.id
    WHERE bt.code = 'farmacia' LIMIT 1;
    
    IF template_id IS NOT NULL THEN
        SELECT jsonb_agg(gp.id) INTO farmacia_products
        FROM global_products gp
        WHERE gp.ean IN (
            '7791234000456', -- Shampoo Head & Shoulders
            '7791293001567', -- Suavizante Comfort
            '7790375000234'  -- Agua Villa del Sur
        );
        
        INSERT INTO business_type_product_templates (
            business_type_template_id, 
            suggested_products,
            priority_brands,
            max_products_per_category
        ) VALUES (
            template_id,
            farmacia_products,
            '["Head & Shoulders", "P&G", "Unilever"]',
            25
        ) ON CONFLICT (business_type_template_id) DO UPDATE SET
            suggested_products = EXCLUDED.suggested_products,
            priority_brands = EXCLUDED.priority_brands;
    END IF;
    
    -- KIOSCO - Bebidas, snacks y golosinas
    SELECT id INTO template_id FROM business_type_templates btt
    JOIN business_types bt ON btt.business_type_id = bt.id
    WHERE bt.code = 'kiosco' LIMIT 1;
    
    IF template_id IS NOT NULL THEN
        SELECT jsonb_agg(gp.id) INTO kiosco_products
        FROM global_products gp
        WHERE gp.ean IN (
            '7790895000126', -- Coca Cola 600ml
            '7790742000101', -- Pepsi 600ml
            '7790895000890', -- Sprite 600ml
            '7790350000456', -- Bon o Bon
            '7790070123456', -- Galletitas Tita
            '7790070789123'  -- Alfajor Havanna
        );
        
        INSERT INTO business_type_product_templates (
            business_type_template_id, 
            suggested_products,
            priority_brands,
            max_products_per_category
        ) VALUES (
            template_id,
            kiosco_products,
            '["Coca Cola", "Pepsi", "Arcor", "Bagley", "Havanna"]',
            40
        ) ON CONFLICT (business_type_template_id) DO UPDATE SET
            suggested_products = EXCLUDED.suggested_products,
            priority_brands = EXCLUDED.priority_brands;
    END IF;
    
    -- SUPERMERCADO - Amplia variedad de productos
    SELECT id INTO template_id FROM business_type_templates btt
    JOIN business_types bt ON btt.business_type_id = bt.id
    WHERE bt.code = 'supermercado' LIMIT 1;
    
    IF template_id IS NOT NULL THEN
        SELECT jsonb_agg(gp.id) INTO super_products
        FROM global_products gp
        WHERE gp.product_type IN ('fmcg', 'local_brand')
        LIMIT 20;
        
        INSERT INTO business_type_product_templates (
            business_type_template_id, 
            suggested_products,
            priority_brands,
            max_products_per_category
        ) VALUES (
            template_id,
            super_products,
            '["Coca Cola", "La Serenísima", "Arcor", "Unilever", "P&G"]',
            100
        ) ON CONFLICT (business_type_template_id) DO UPDATE SET
            suggested_products = EXCLUDED.suggested_products,
            priority_brands = EXCLUDED.priority_brands;
    END IF;
    
    -- ELECTRÓNICA - Productos tecnológicos
    SELECT id INTO template_id FROM business_type_templates btt
    JOIN business_types bt ON btt.business_type_id = bt.id
    WHERE bt.code = 'electronica' LIMIT 1;
    
    IF template_id IS NOT NULL THEN
        SELECT jsonb_agg(gp.id) INTO electronica_products
        FROM global_products gp
        WHERE gp.product_type = 'industrialized';
        
        INSERT INTO business_type_product_templates (
            business_type_template_id, 
            suggested_products,
            priority_brands,
            max_products_per_category
        ) VALUES (
            template_id,
            electronica_products,
            '["Samsung", "Apple", "Xiaomi", "Motorola"]',
            50
        ) ON CONFLICT (business_type_template_id) DO UPDATE SET
            suggested_products = EXCLUDED.suggested_products,
            priority_brands = EXCLUDED.priority_brands;
    END IF;
    
    RAISE NOTICE 'Templates de productos configurados para tipos de negocio';
    
END $$;

-- Comentarios para documentación
COMMENT ON TABLE global_products IS 'Catálogo global con productos populares argentinos iniciales';

-- Estadísticas finales
DO $$
DECLARE
    total_products INTEGER;
    fmcg_products INTEGER;
    industrialized_products INTEGER;
    local_products INTEGER;
    templates_configured INTEGER;
BEGIN
    SELECT COUNT(*) INTO total_products FROM global_products;
    SELECT COUNT(*) INTO fmcg_products FROM global_products WHERE product_type = 'fmcg';
    SELECT COUNT(*) INTO industrialized_products FROM global_products WHERE product_type = 'industrialized';
    SELECT COUNT(*) INTO local_products FROM global_products WHERE product_type = 'local_brand';
    SELECT COUNT(*) INTO templates_configured FROM business_type_product_templates;
    
    RAISE NOTICE '=== ESTADÍSTICAS DEL CATÁLOGO GLOBAL ===';
    RAISE NOTICE 'Total productos: %', total_products;
    RAISE NOTICE 'Productos FMCG: %', fmcg_products;
    RAISE NOTICE 'Productos industrializados: %', industrialized_products;
    RAISE NOTICE 'Productos locales: %', local_products;
    RAISE NOTICE 'Templates configurados: %', templates_configured;
    RAISE NOTICE '========================================';
END $$; 