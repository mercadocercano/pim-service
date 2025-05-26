-- Seeds para variantes de productos
-- Nota: Estos seeds asumen que ya existen productos en la tabla products
-- Las variantes por defecto se crean automáticamente desde el código del agregado Product

-- Obtener el ID del iPhone 15 Pro Max para crear variantes adicionales
DO $$
DECLARE
    iphone_product_id UUID;
    tenant_id_val UUID := '9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8';
    default_variant_id UUID;
    variant_256_nat_id UUID;
    variant_512_blue_id UUID;
    variant_1tb_white_id UUID;
    variant_black_id UUID;
BEGIN
    -- Obtener el ID del iPhone 15 Pro Max
    SELECT id INTO iphone_product_id 
    FROM products 
    WHERE name = 'iPhone 15 Pro Max' 
    AND tenant_id = tenant_id_val
    LIMIT 1;
    
    -- Solo proceder si encontramos el producto
    IF iphone_product_id IS NOT NULL THEN
        -- Actualizar la variante por defecto que ya fue creada automáticamente
        UPDATE product_variants 
        SET name = 'iPhone 15 Pro Max - 128GB - Natural Titanium',
            sku = 'IPHONE-15-PRO-MAX-128-NAT'
        WHERE product_id = iphone_product_id 
        AND is_default = true;
        
        -- Obtener el ID de la variante por defecto
        SELECT id INTO default_variant_id 
        FROM product_variants 
        WHERE product_id = iphone_product_id 
        AND is_default = true
        LIMIT 1;
        
        -- Agregar atributos a la variante por defecto
        INSERT INTO variant_attributes (tenant_id, variant_id, attribute_name, attribute_value) VALUES
        (tenant_id_val, default_variant_id, 'Storage', '128GB'),
        (tenant_id_val, default_variant_id, 'Color', 'Natural Titanium'),
        (tenant_id_val, default_variant_id, 'Material', 'Titanium');
        
        -- Variante 256GB - Natural Titanium
        INSERT INTO product_variants (tenant_id, product_id, name, sku, status, is_default, sort_order)
        VALUES (tenant_id_val, iphone_product_id, 'iPhone 15 Pro Max - 256GB - Natural Titanium', 'IPHONE-15-PRO-MAX-256-NAT', 'active', false, 2)
        RETURNING id INTO variant_256_nat_id;
        
        INSERT INTO variant_attributes (tenant_id, variant_id, attribute_name, attribute_value) VALUES
        (tenant_id_val, variant_256_nat_id, 'Storage', '256GB'),
        (tenant_id_val, variant_256_nat_id, 'Color', 'Natural Titanium'),
        (tenant_id_val, variant_256_nat_id, 'Material', 'Titanium');
        
        -- Variante 512GB - Blue Titanium
        INSERT INTO product_variants (tenant_id, product_id, name, sku, status, is_default, sort_order)
        VALUES (tenant_id_val, iphone_product_id, 'iPhone 15 Pro Max - 512GB - Blue Titanium', 'IPHONE-15-PRO-MAX-512-BLUE', 'active', false, 3)
        RETURNING id INTO variant_512_blue_id;
        
        INSERT INTO variant_attributes (tenant_id, variant_id, attribute_name, attribute_value) VALUES
        (tenant_id_val, variant_512_blue_id, 'Storage', '512GB'),
        (tenant_id_val, variant_512_blue_id, 'Color', 'Blue Titanium'),
        (tenant_id_val, variant_512_blue_id, 'Material', 'Titanium');
        
        -- Variante 1TB - White Titanium
        INSERT INTO product_variants (tenant_id, product_id, name, sku, status, is_default, sort_order)
        VALUES (tenant_id_val, iphone_product_id, 'iPhone 15 Pro Max - 1TB - White Titanium', 'IPHONE-15-PRO-MAX-1TB-WHITE', 'active', false, 4)
        RETURNING id INTO variant_1tb_white_id;
        
        INSERT INTO variant_attributes (tenant_id, variant_id, attribute_name, attribute_value) VALUES
        (tenant_id_val, variant_1tb_white_id, 'Storage', '1TB'),
        (tenant_id_val, variant_1tb_white_id, 'Color', 'White Titanium'),
        (tenant_id_val, variant_1tb_white_id, 'Material', 'Titanium');
        
        -- Variante descontinuada (ejemplo)
        INSERT INTO product_variants (tenant_id, product_id, name, sku, status, is_default, sort_order)
        VALUES (tenant_id_val, iphone_product_id, 'iPhone 15 Pro Max - 128GB - Black Titanium', 'IPHONE-15-PRO-MAX-128-BLACK', 'discontinued', false, 5)
        RETURNING id INTO variant_black_id;
        
        INSERT INTO variant_attributes (tenant_id, variant_id, attribute_name, attribute_value) VALUES
        (tenant_id_val, variant_black_id, 'Storage', '128GB'),
        (tenant_id_val, variant_black_id, 'Color', 'Black Titanium'),
        (tenant_id_val, variant_black_id, 'Material', 'Titanium');
        
        RAISE NOTICE 'Se actualizaron/crearon % variantes para el producto iPhone 15 Pro Max', 
            (SELECT COUNT(*) FROM product_variants WHERE product_id = iphone_product_id);
    ELSE
        RAISE NOTICE 'No se encontró el producto iPhone 15 Pro Max para crear variantes';
    END IF;
END $$;

-- Nota: Los demás productos ya tienen sus variantes por defecto creadas automáticamente
-- desde el código del agregado Product cuando se crearon los productos

-- Crear algunas variantes para otros productos si existen
DO $$
DECLARE
    product_record RECORD;
    variant_id UUID;
    tenant_id_val UUID := '9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8';
BEGIN
    -- Buscar productos que no sean el iPhone y que no tengan variantes
    FOR product_record IN 
        SELECT p.id, p.name, p.sku 
        FROM products p 
        LEFT JOIN product_variants pv ON p.id = pv.product_id
        WHERE p.tenant_id = tenant_id_val 
        AND p.status != 'deleted'
        AND p.name != 'iPhone 15 Pro Max'
        AND pv.id IS NULL
        LIMIT 3
    LOOP
        -- Crear variante por defecto para cada producto
        INSERT INTO product_variants (tenant_id, product_id, name, sku, status, is_default, sort_order)
        VALUES (tenant_id_val, product_record.id, product_record.name || ' - Default', product_record.sku || '-DEFAULT', 'active', true, 1)
        RETURNING id INTO variant_id;
        
        -- Agregar un atributo básico
        INSERT INTO variant_attributes (tenant_id, variant_id, attribute_name, attribute_value)
        VALUES (tenant_id_val, variant_id, 'Type', 'Standard');
        
        RAISE NOTICE 'Creada variante por defecto para producto: %', product_record.name;
    END LOOP;
END $$; 