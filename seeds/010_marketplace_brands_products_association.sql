-- Seed: 010_marketplace_brands_products_association.sql
-- Purpose: Asociar marcas con productos típicos y categorías por tipo de negocio
-- Date: 2025-01-20
-- Beneficio: Mejorar la curación identificando productos típicos por marca y categoría

-- ============================================
-- MARCAS Y PRODUCTOS TÍPICOS POR CATEGORÍA
-- ============================================

-- Tabla temporal para asociaciones marca-producto-categoría
CREATE TEMP TABLE IF NOT EXISTS brand_product_associations (
    brand_name VARCHAR(255),
    category VARCHAR(100),
    typical_products TEXT[],
    business_types TEXT[]
);

-- ============================================
-- ALMACÉN DE BARRIO - Productos y Marcas
-- ============================================
INSERT INTO brand_product_associations VALUES
-- Yerba Mate
('La Merced', 'yerba', ARRAY['Yerba Mate La Merced 500g', 'Yerba Mate La Merced 1kg', 'Yerba Mate La Merced Barbacuá'], ARRAY['almacen', 'supermercado']),
('Amanda', 'yerba', ARRAY['Yerba Amanda Tradicional 500g', 'Yerba Amanda Tradicional 1kg', 'Yerba Amanda Despalada'], ARRAY['almacen', 'supermercado']),
('Taragui', 'yerba', ARRAY['Yerba Taragui 500g', 'Yerba Taragui 1kg', 'Yerba Taragui Sin Palo'], ARRAY['almacen', 'supermercado']),
('Rosamonte', 'yerba', ARRAY['Yerba Rosamonte 500g', 'Yerba Rosamonte 1kg', 'Yerba Rosamonte Especial'], ARRAY['almacen', 'supermercado']),
('La Morenita', 'yerba', ARRAY['Yerba La Morenita 500g', 'Yerba La Morenita 1kg'], ARRAY['almacen', 'supermercado']),
('Cruz de Malta', 'yerba', ARRAY['Yerba Cruz de Malta 500g', 'Yerba Cruz de Malta 1kg'], ARRAY['almacen', 'supermercado']),

-- Aceites
('Cocinero', 'aceites', ARRAY['Aceite Cocinero 900ml', 'Aceite Cocinero 1.5L', 'Aceite Cocinero Girasol'], ARRAY['almacen', 'supermercado']),
('Natura', 'aceites', ARRAY['Aceite Natura 900ml', 'Aceite Natura 1.5L', 'Aceite Natura Girasol Alto Oleico'], ARRAY['almacen', 'supermercado']),
('Cada Día', 'aceites', ARRAY['Aceite Cada Día 900ml', 'Aceite Cada Día 1.5L'], ARRAY['almacen', 'supermercado']),
('Marolio', 'aceites', ARRAY['Aceite Marolio 900ml', 'Aceite Marolio 1.5L', 'Aceite Marolio Mezcla'], ARRAY['almacen', 'supermercado']),

-- Harinas
('Blancaflor', 'harinas', ARRAY['Harina Blancaflor 000 1kg', 'Harina Blancaflor 0000 1kg', 'Harina Blancaflor Leudante'], ARRAY['almacen', 'supermercado']),
('Pureza', 'harinas', ARRAY['Harina Pureza 000 1kg', 'Harina Pureza 0000 1kg', 'Harina Pureza Integral'], ARRAY['almacen', 'supermercado']),
('Favorita', 'harinas', ARRAY['Harina Favorita 000 1kg', 'Harina Favorita 0000 1kg'], ARRAY['almacen', 'supermercado']),

-- Fideos
('Matarazzo', 'fideos', ARRAY['Fideos Matarazzo Spaghetti 500g', 'Fideos Matarazzo Tallarines', 'Fideos Matarazzo Mostacholes'], ARRAY['almacen', 'supermercado']),
('Lucchetti', 'fideos', ARRAY['Fideos Lucchetti Spaghetti 500g', 'Fideos Lucchetti Tallarines', 'Fideos Lucchetti Tirabuzón'], ARRAY['almacen', 'supermercado']),
('Don Vicente', 'fideos', ARRAY['Fideos Don Vicente Spaghetti 500g', 'Fideos Don Vicente Tallarines'], ARRAY['almacen', 'supermercado']),
('Knorr', 'fideos', ARRAY['Fideos Knorr Quick', 'Fideos Knorr Vitina'], ARRAY['almacen', 'supermercado']),

-- Arroz
('Gallo', 'arroz', ARRAY['Arroz Gallo Oro 1kg', 'Arroz Gallo Doble Carolina', 'Arroz Gallo Integral'], ARRAY['almacen', 'supermercado']),
('Molinos Ala', 'arroz', ARRAY['Arroz Molinos Ala 1kg', 'Arroz Molinos Ala Parboiled'], ARRAY['almacen', 'supermercado']),
('Lucchetti', 'arroz', ARRAY['Arroz Lucchetti 1kg', 'Arroz Lucchetti Largo Fino'], ARRAY['almacen', 'supermercado']),
('Canale', 'arroz', ARRAY['Arroz Canale 1kg', 'Arroz Canale Doble Carolina'], ARRAY['almacen', 'supermercado']),

-- Azúcar
('Ledesma', 'azucar', ARRAY['Azúcar Ledesma 1kg', 'Azúcar Ledesma Light', 'Azúcar Ledesma Rubia'], ARRAY['almacen', 'supermercado']),
('Chango', 'azucar', ARRAY['Azúcar Chango 1kg', 'Azúcar Chango Impalpable'], ARRAY['almacen', 'supermercado']),
('Domino', 'azucar', ARRAY['Azúcar Domino 1kg', 'Azúcar Domino Rubia'], ARRAY['almacen', 'supermercado']),

-- Lácteos
('La Serenísima', 'lacteos', ARRAY['Leche La Serenísima 1L', 'Yogur La Serenísima', 'Queso La Serenísima Cremoso', 'Manteca La Serenísima'], ARRAY['almacen', 'supermercado']),
('SanCor', 'lacteos', ARRAY['Leche SanCor 1L', 'Yogur SanCor', 'Queso SanCor Por Salut', 'Dulce de Leche SanCor'], ARRAY['almacen', 'supermercado']),
('Ilolay', 'lacteos', ARRAY['Leche Ilolay 1L', 'Queso Ilolay Untable', 'Manteca Ilolay'], ARRAY['almacen', 'supermercado']),
('Milkaut', 'lacteos', ARRAY['Leche Milkaut 1L', 'Yogur Milkaut', 'Dulce de Leche Milkaut'], ARRAY['almacen', 'supermercado']),
('Tregar', 'lacteos', ARRAY['Leche Tregar 1L', 'Queso Tregar Rallado', 'Crema Tregar'], ARRAY['almacen', 'supermercado']),
('Vauquita', 'lacteos', ARRAY['Dulce de Leche Vauquita', 'Leche Vauquita 1L'], ARRAY['almacen', 'supermercado']),

-- Galletitas
('Bagley', 'galletitas', ARRAY['Galletitas Criollitas', 'Galletitas Sonrisas', 'Galletitas Rumba', 'Galletitas Rex'], ARRAY['almacen', 'supermercado', 'kiosco']),
('Terrabusi', 'galletitas', ARRAY['Galletitas Lincoln', 'Galletitas Melba', 'Galletitas Express', 'Galletitas Variedad'], ARRAY['almacen', 'supermercado', 'kiosco']),
('Arcor', 'galletitas', ARRAY['Galletitas Surtido Arcor', 'Galletitas Merengadas', 'Galletitas Saladix'], ARRAY['almacen', 'supermercado', 'kiosco']),
('Don Satur', 'galletitas', ARRAY['Bizcochos Don Satur Salados', 'Bizcochos Don Satur Dulces', 'Bizcochos Don Satur Agridulces'], ARRAY['almacen', 'supermercado']),

-- Golosinas
('Arcor', 'golosinas', ARRAY['Bon o Bon', 'Rocklets', 'Butter Toffees', 'Mogul', 'Topline'], ARRAY['kiosco', 'almacen', 'supermercado']),
('Georgalos', 'golosinas', ARRAY['Pico Dulce', 'Flynn Paff', 'Palitos de la Selva', 'Chupetines Mr. Pop'], ARRAY['kiosco', 'almacen']),
('Felfort', 'golosinas', ARRAY['Chocolates Felfort', 'Dos Corazones', 'Marroc'], ARRAY['kiosco', 'almacen', 'supermercado']),
('Havanna', 'golosinas', ARRAY['Alfajores Havanna', 'Havanna 70% Cacao', 'Havanna Nuez'], ARRAY['kiosco', 'almacen']),

-- Bebidas Gaseosas
('Coca-Cola', 'gaseosas', ARRAY['Coca-Cola 500ml', 'Coca-Cola 1.5L', 'Coca-Cola 2.25L', 'Coca-Cola Zero'], ARRAY['almacen', 'supermercado', 'kiosco']),
('Pepsi', 'gaseosas', ARRAY['Pepsi 500ml', 'Pepsi 1.5L', 'Pepsi 2L', 'Pepsi Black'], ARRAY['almacen', 'supermercado', 'kiosco']),
('Sprite', 'gaseosas', ARRAY['Sprite 500ml', 'Sprite 1.5L', 'Sprite 2.25L', 'Sprite Zero'], ARRAY['almacen', 'supermercado', 'kiosco']),
('Manaos', 'gaseosas', ARRAY['Manaos Cola 2.25L', 'Manaos Naranja 2.25L', 'Manaos Lima Limón 2.25L', 'Manaos Uva'], ARRAY['almacen', 'supermercado']),
('Cunnington', 'gaseosas', ARRAY['Cunnington Cola 2.25L', 'Cunnington Naranja', 'Cunnington Pomelo'], ARRAY['almacen', 'supermercado']),

-- Cervezas
('Quilmes', 'cerveza', ARRAY['Quilmes Clásica 1L', 'Quilmes 473ml', 'Quilmes Stout', 'Quilmes Red Lager'], ARRAY['almacen', 'supermercado', 'kiosco']),
('Brahma', 'cerveza', ARRAY['Brahma 1L', 'Brahma 473ml', 'Brahma Dorada'], ARRAY['almacen', 'supermercado', 'kiosco']),
('Stella Artois', 'cerveza', ARRAY['Stella Artois 1L', 'Stella Artois 473ml', 'Stella Artois Noire'], ARRAY['almacen', 'supermercado']),
('Patagonia', 'cerveza', ARRAY['Patagonia Amber Lager', 'Patagonia Weisse', 'Patagonia IPA'], ARRAY['almacen', 'supermercado']),
('Schneider', 'cerveza', ARRAY['Schneider 1L', 'Schneider 473ml', 'Schneider Rubia'], ARRAY['almacen', 'supermercado']),

-- Limpieza
('Ayudín', 'limpieza', ARRAY['Detergente Ayudín 300ml', 'Detergente Ayudín 500ml', 'Detergente Ayudín Limón'], ARRAY['almacen', 'supermercado']),
('Magistral', 'limpieza', ARRAY['Detergente Magistral 300ml', 'Detergente Magistral 500ml', 'Detergente Magistral Ultra'], ARRAY['almacen', 'supermercado']),
('Skip', 'limpieza', ARRAY['Jabón en Polvo Skip 800g', 'Skip Líquido 3L', 'Skip Intelligence'], ARRAY['almacen', 'supermercado']),
('Procenex', 'limpieza', ARRAY['Lavandina Procenex 1L', 'Lavandina Procenex 2L', 'Desinfectante Procenex'], ARRAY['almacen', 'supermercado']),
('Vim', 'limpieza', ARRAY['Vim Crema 250ml', 'Vim Líquido 500ml', 'Vim Antigrasa'], ARRAY['almacen', 'supermercado']),
('Mr. Musculo', 'limpieza', ARRAY['Mr. Musculo Cocina', 'Mr. Musculo Baño', 'Mr. Musculo Vidrios'], ARRAY['almacen', 'supermercado']);

-- ============================================
-- CARNICERÍA - Productos típicos sin marca
-- ============================================
INSERT INTO brand_product_associations VALUES
('Sin Marca', 'carnes', ARRAY['Asado', 'Vacío', 'Matambre', 'Bife de Chorizo', 'Milanesas', 'Carne Picada', 'Pechuga de Pollo', 'Chorizo', 'Morcilla'], ARRAY['carniceria']),
('Swift', 'carnes', ARRAY['Hamburguesas Swift', 'Milanesas Swift', 'Medallones Swift'], ARRAY['supermercado']),
('Paty', 'carnes', ARRAY['Hamburguesas Paty', 'Paty Finitas', 'Paty XL'], ARRAY['supermercado']),
('Granja del Sol', 'carnes', ARRAY['Pollo Entero', 'Pechuga de Pollo', 'Pata Muslo'], ARRAY['carniceria', 'supermercado']);

-- ============================================
-- VERDULERÍA - Productos frescos
-- ============================================
INSERT INTO brand_product_associations VALUES
('Sin Marca', 'verduras', ARRAY['Tomate', 'Lechuga', 'Papa', 'Cebolla', 'Zanahoria', 'Zapallo', 'Morron', 'Ajo', 'Perejil'], ARRAY['verduleria']),
('Sin Marca', 'frutas', ARRAY['Manzana', 'Banana', 'Naranja', 'Mandarina', 'Pera', 'Durazno', 'Frutilla', 'Uva', 'Limón'], ARRAY['verduleria']);

-- ============================================
-- FARMACIA - Medicamentos y cuidado personal
-- ============================================
INSERT INTO brand_product_associations VALUES
('Bayer', 'medicamentos', ARRAY['Aspirina', 'Bayaspirina', 'Redoxon', 'Supradyn'], ARRAY['farmacia']),
('Roemmers', 'medicamentos', ARRAY['Amoxidal', 'Lotrial', 'Ibupirac', 'Tafirol'], ARRAY['farmacia']),
('Bagó', 'medicamentos', ARRAY['Buscapina', 'Sertal', 'Cafiaspirina'], ARRAY['farmacia']),
('Andrómaco', 'medicamentos', ARRAY['Flexicamin B12', 'Hepatalgina', 'Qura Plus'], ARRAY['farmacia']),

-- Cuidado Personal
('Dove', 'cuidado-personal', ARRAY['Jabón Dove', 'Shampoo Dove', 'Acondicionador Dove', 'Desodorante Dove'], ARRAY['farmacia', 'perfumeria', 'supermercado']),
('Nivea', 'cuidado-personal', ARRAY['Crema Nivea', 'Desodorante Nivea', 'Protector Solar Nivea'], ARRAY['farmacia', 'perfumeria', 'supermercado']),
('Sedal', 'cuidado-personal', ARRAY['Shampoo Sedal', 'Acondicionador Sedal', 'Crema de Peinar Sedal'], ARRAY['farmacia', 'perfumeria', 'supermercado']),
('Rexona', 'cuidado-personal', ARRAY['Desodorante Rexona Men', 'Desodorante Rexona Women', 'Rexona Clinical'], ARRAY['farmacia', 'perfumeria', 'supermercado']),
('Colgate', 'cuidado-personal', ARRAY['Pasta Dental Colgate Total', 'Cepillo Colgate', 'Enjuague Bucal Colgate'], ARRAY['farmacia', 'supermercado']);

-- ============================================
-- FERRETERÍA - Herramientas y materiales
-- ============================================
INSERT INTO brand_product_associations VALUES
('Alba', 'pintura', ARRAY['Látex Alba', 'Alba Antihongo', 'Alba Techos'], ARRAY['ferreteria']),
('Sinteplast', 'pintura', ARRAY['Látex Sinteplast', 'Sinteplast Impermeabilizante', 'Enduido Sinteplast'], ARRAY['ferreteria']),
('Tersuave', 'pintura', ARRAY['Látex Tersuave', 'Tersuave Satinado', 'Tersuave Techos'], ARRAY['ferreteria']),
('Bahco', 'herramientas', ARRAY['Destornilladores Bahco', 'Llaves Bahco', 'Alicates Bahco'], ARRAY['ferreteria']),
('Stanley', 'herramientas', ARRAY['Martillo Stanley', 'Cinta Métrica Stanley', 'Destornilladores Stanley'], ARRAY['ferreteria']),
('Black & Decker', 'herramientas', ARRAY['Taladro Black & Decker', 'Amoladora Black & Decker', 'Lijadora Black & Decker'], ARRAY['ferreteria']);

-- ============================================
-- KIOSCO - Golosinas y cigarrillos
-- ============================================
INSERT INTO brand_product_associations VALUES
('Marlboro', 'cigarrillos', ARRAY['Marlboro Box', 'Marlboro Gold', 'Marlboro Red'], ARRAY['kiosco']),
('Philip Morris', 'cigarrillos', ARRAY['Philip Morris Box', 'Philip Morris Blue'], ARRAY['kiosco']),
('Lucky Strike', 'cigarrillos', ARRAY['Lucky Strike Box', 'Lucky Strike Blue'], ARRAY['kiosco']),
('7up', 'gaseosas', ARRAY['7up 500ml', '7up 1.5L'], ARRAY['kiosco', 'almacen']),
('Bimbo', 'panificados', ARRAY['Pan Bimbo', 'Pan Lactal Bimbo', 'Rapiditas Bimbo'], ARRAY['kiosco', 'almacen', 'supermercado']);

-- ============================================
-- ACTUALIZAR marketplace_brands CON INFORMACIÓN ADICIONAL
-- ============================================

-- Actualizar las marcas con información de productos típicos
UPDATE marketplace_brands mb
SET web_data = jsonb_set(
    COALESCE(web_data, '{}'::jsonb),
    '{typical_products}',
    to_jsonb(bpa.typical_products)
)
FROM brand_product_associations bpa
WHERE mb.name = bpa.brand_name
AND bpa.typical_products IS NOT NULL;

-- Actualizar category_tags basado en las asociaciones
UPDATE marketplace_brands mb
SET category_tags = array_cat(
    COALESCE(category_tags, '{}'),
    ARRAY(SELECT DISTINCT unnest(string_to_array(bpa.category, ','))
          FROM brand_product_associations bpa
          WHERE mb.name = bpa.brand_name)
)
WHERE EXISTS (
    SELECT 1 FROM brand_product_associations bpa
    WHERE mb.name = bpa.brand_name
);

-- ============================================
-- MARCAS ADICIONALES DETECTADAS EN SCRAPERS
-- ============================================
INSERT INTO marketplace_brands (name, slug, normalized_name, description, category_tags, verification_status, quality_score, aliases, country_code) VALUES
-- Marcas de productos frescos/locales
('Granja del Sol', 'granja-del-sol', 'GRANJA DEL SOL', 'Productos avícolas frescos', ARRAY['carnes', 'alimentos-frescos'], 'verified', 0.72, NULL, 'AR'),
('Campo Austral', 'campo-austral', 'CAMPO AUSTRAL', 'Productos orgánicos y naturales', ARRAY['alimentos-frescos', 'organicos'], 'verified', 0.68, NULL, 'AR'),
('Don Atilio', 'don-atilio', 'DON ATILIO', 'Embutidos y chacinados artesanales', ARRAY['carnes', 'embutidos'], 'verified', 0.70, NULL, 'AR'),

-- Marcas económicas/segundas marcas
('Molto', 'molto', 'MOLTO', 'Segunda marca de La Campagnola', ARRAY['alimentos', 'conservas'], 'verified', 0.65, ARRAY['Campagnola Molto'], 'AR'),
('Marolio', 'marolio', 'MAROLIO', 'Marca económica multi-categoría', ARRAY['alimentos', 'limpieza', 'bebidas'], 'verified', 0.63, ARRAY['Maxiconsumo'], 'AR'),
('Cada Día', 'cada-dia', 'CADA DIA', 'Marca económica de alimentos', ARRAY['alimentos', 'aceites'], 'verified', 0.60, NULL, 'AR'),

-- Marcas de productos importados comunes
('Hellmanns', 'hellmanns', 'HELLMANNS', 'Mayonesas y aderezos', ARRAY['alimentos', 'aderezos'], 'verified', 0.85, ARRAY['Hellmann''s'], 'AR'),
('Heinz', 'heinz', 'HEINZ', 'Ketchup y salsas', ARRAY['alimentos', 'aderezos'], 'verified', 0.83, NULL, 'AR'),
('Pringles', 'pringles', 'PRINGLES', 'Papas en tubo', ARRAY['snacks', 'alimentos'], 'verified', 0.88, NULL, 'AR')
ON CONFLICT (slug) DO UPDATE SET
    aliases = EXCLUDED.aliases,
    quality_score = EXCLUDED.quality_score,
    category_tags = array_cat(marketplace_brands.category_tags, EXCLUDED.category_tags);

-- ============================================
-- ESTADÍSTICAS FINALES
-- ============================================
DO $$
DECLARE 
    total_associations INTEGER;
    brands_with_products INTEGER;
    total_typical_products INTEGER;
BEGIN
    SELECT COUNT(DISTINCT brand_name) INTO total_associations FROM brand_product_associations;
    
    SELECT COUNT(*) INTO brands_with_products 
    FROM marketplace_brands 
    WHERE web_data->>'typical_products' IS NOT NULL;
    
    SELECT SUM(array_length(typical_products, 1)) INTO total_typical_products
    FROM brand_product_associations;
    
    RAISE NOTICE '✅ Asociaciones de marcas y productos completadas:';
    RAISE NOTICE '   📊 Marcas con productos típicos: %', total_associations;
    RAISE NOTICE '   🛍️ Total de productos típicos identificados: %', total_typical_products;
    RAISE NOTICE '   🏷️ Marcas actualizadas en BD: %', brands_with_products;
    RAISE NOTICE '';
    RAISE NOTICE '   Categorías principales cubiertas:';
    RAISE NOTICE '   - Almacén: yerba, aceites, harinas, fideos, lácteos';
    RAISE NOTICE '   - Carnicería: carnes frescas y procesadas';
    RAISE NOTICE '   - Verdulería: frutas y verduras frescas';
    RAISE NOTICE '   - Farmacia: medicamentos y cuidado personal';
    RAISE NOTICE '   - Ferretería: herramientas y pintura';
    RAISE NOTICE '   - Kiosco: golosinas, cigarrillos, bebidas';
END;
$$;

-- Limpiar tabla temporal
DROP TABLE IF EXISTS brand_product_associations;