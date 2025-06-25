-- Seed: 008_marketplace_brands_argentina_seed.sql
-- Purpose: Poblar catálogo inicial de marcas populares en Argentina
-- Date: 2025-01-20
-- Beneficio: Tener marcas comunes pre-cargadas para facilitar curación y onboarding

-- Limpiar marcas existentes
DELETE FROM marketplace_brands;

-- Marcas de Alimentos y Bebidas
INSERT INTO marketplace_brands (name, slug, normalized_name, description, website, category_tags, verification_status, quality_score, aliases) VALUES
('Coca-Cola', 'coca-cola', 'COCA-COLA', 'Marca líder mundial de bebidas gaseosas y refrescos', 'https://www.coca-cola.com.ar', ARRAY['bebidas', 'gaseosas'], 'verified', 0.95, ARRAY['Coca Cola', 'CocaCola', 'Coke']),
('La Serenísima', 'la-serenisima', 'LA SERENISIMA', 'Marca argentina líder en productos lácteos', 'https://www.laserenisima.com.ar', ARRAY['lacteos', 'alimentos'], 'verified', 0.90, ARRAY['Serenisima', 'Mastellone']),
('Arcor', 'arcor', 'ARCOR', 'Grupo argentino líder en golosinas y alimentos', 'https://www.arcor.com.ar', ARRAY['golosinas', 'alimentos'], 'verified', 0.92, ARRAY['Grupo Arcor']),
('Molinos Río de la Plata', 'molinos-rio-de-la-plata', 'MOLINOS RIO DE LA PLATA', 'Empresa argentina de alimentos procesados', 'https://www.molinos.com.ar', ARRAY['alimentos', 'harinas'], 'verified', 0.88, ARRAY['Molinos', 'Molinos Ríos de la Plata']),
('Nestlé', 'nestle', 'NESTLE', 'Multinacional de alimentos y bebidas', 'https://www.nestle.com.ar', ARRAY['alimentos', 'bebidas', 'lacteos'], 'verified', 0.94, ARRAY['Nestle Argentina']),
('Unilever', 'unilever', 'UNILEVER', 'Multinacional de productos de consumo', 'https://www.unilever.com.ar', ARRAY['alimentos', 'cuidado-personal'], 'verified', 0.91, ARRAY['Unilever Argentina']),
('Quilmes', 'quilmes', 'QUILMES', 'Cervecería argentina tradicional', 'https://www.quilmes.com.ar', ARRAY['bebidas', 'cerveza'], 'verified', 0.89, ARRAY['Cervecería Quilmes']),
('Havanna', 'havanna', 'HAVANNA', 'Marca argentina de alfajores y dulces', 'https://www.havanna.com.ar', ARRAY['golosinas', 'alfajores'], 'verified', 0.87, ARRAY['Alfajores Havanna']),
('Bagley', 'bagley', 'BAGLEY', 'Marca argentina de galletitas y productos horneados', 'https://www.bagley.com.ar', ARRAY['galletitas', 'alimentos'], 'verified', 0.85, ARRAY['Bagley S.A.']),
('Terrabusi', 'terrabusi', 'TERRABUSI', 'Marca argentina de galletitas', NULL, ARRAY['galletitas', 'alimentos'], 'verified', 0.83, ARRAY['Terrabusi S.A.']),

-- Marcas de Limpieza y Cuidado del Hogar
('Ayudín', 'ayudin', 'AYUDIN', 'Marca argentina de productos de limpieza', NULL, ARRAY['limpieza', 'hogar'], 'verified', 0.82, ARRAY['Ayudin']),
('Magistral', 'magistral', 'MAGISTRAL', 'Marca de productos de limpieza', NULL, ARRAY['limpieza', 'hogar'], 'verified', 0.80, NULL),
('Skip', 'skip', 'SKIP', 'Marca de detergentes y productos de limpieza', NULL, ARRAY['limpieza', 'detergente'], 'verified', 0.84, NULL),
('Procenex', 'procenex', 'PROCENEX', 'Marca argentina de productos de limpieza profesional', NULL, ARRAY['limpieza', 'profesional'], 'verified', 0.78, NULL),

-- Marcas de Cuidado Personal
('Sedal', 'sedal', 'SEDAL', 'Marca de productos para el cuidado del cabello', NULL, ARRAY['cuidado-personal', 'cabello'], 'verified', 0.86, NULL),
('Dove', 'dove', 'DOVE', 'Marca internacional de cuidado personal', 'https://www.dove.com', ARRAY['cuidado-personal', 'belleza'], 'verified', 0.93, NULL),
('Rexona', 'rexona', 'REXONA', 'Marca de desodorantes y antitranspirantes', NULL, ARRAY['cuidado-personal', 'desodorante'], 'verified', 0.88, NULL),
('Colgate', 'colgate', 'COLGATE', 'Marca líder en cuidado bucal', 'https://www.colgate.com.ar', ARRAY['cuidado-personal', 'higiene-bucal'], 'verified', 0.92, NULL),

-- Marcas de Tecnología
('Samsung', 'samsung', 'SAMSUNG', 'Multinacional surcoreana de tecnología', 'https://www.samsung.com/ar', ARRAY['tecnologia', 'electronica'], 'verified', 0.96, ARRAY['Samsung Electronics', 'Samsung Galaxy']),
('Apple', 'apple', 'APPLE', 'Empresa tecnológica estadounidense', 'https://www.apple.com', ARRAY['tecnologia', 'electronica'], 'verified', 0.97, ARRAY['Apple Inc.', 'iPhone']),
('Sony', 'sony', 'SONY', 'Multinacional japonesa de electrónicos', 'https://www.sony.com.ar', ARRAY['tecnologia', 'electronica'], 'verified', 0.94, ARRAY['Sony Corporation']),
('LG', 'lg', 'LG', 'Multinacional surcoreana de electrónicos', 'https://www.lg.com/ar', ARRAY['tecnologia', 'electrodomesticos'], 'verified', 0.91, ARRAY['LG Electronics']),
('Philips', 'philips', 'PHILIPS', 'Multinacional holandesa de tecnología', 'https://www.philips.com.ar', ARRAY['tecnologia', 'electrodomesticos', 'salud'], 'verified', 0.90, NULL),

-- Marcas de Indumentaria
('Adidas', 'adidas', 'ADIDAS', 'Marca alemana de ropa y calzado deportivo', 'https://www.adidas.com.ar', ARRAY['deportes', 'indumentaria'], 'verified', 0.94, NULL),
('Nike', 'nike', 'NIKE', 'Marca estadounidense de artículos deportivos', 'https://www.nike.com', ARRAY['deportes', 'indumentaria'], 'verified', 0.95, NULL),
('Puma', 'puma', 'PUMA', 'Marca alemana de artículos deportivos', 'https://www.puma.com', ARRAY['deportes', 'indumentaria'], 'verified', 0.89, NULL),
('Topper', 'topper', 'TOPPER', 'Marca argentina de calzado deportivo', 'https://www.topper.com.ar', ARRAY['deportes', 'calzado'], 'verified', 0.85, NULL),

-- Marcas Farmacéuticas y de Salud
('Bayer', 'bayer', 'BAYER', 'Multinacional alemana farmacéutica', 'https://www.bayer.com.ar', ARRAY['salud', 'farmacia'], 'verified', 0.93, NULL),
('Roemmers', 'roemmers', 'ROEMMERS', 'Laboratorio farmacéutico argentino', 'https://www.roemmers.com.ar', ARRAY['salud', 'farmacia'], 'verified', 0.87, NULL),
('Bagó', 'bago', 'BAGO', 'Laboratorio farmacéutico argentino', 'https://www.bago.com.ar', ARRAY['salud', 'farmacia'], 'verified', 0.86, ARRAY['Laboratorios Bagó']),

-- Marcas Genéricas Comunes (para identificar y mejorar)
('Sin Marca', 'sin-marca', 'SIN MARCA', 'Productos sin marca específica identificada', NULL, ARRAY['generico'], 'unverified', 0.10, ARRAY['Sin marca', 'Genérica', 'No brand']),
('Marca Propia', 'marca-propia', 'MARCA PROPIA', 'Productos de marca propia de comercios', NULL, ARRAY['generico'], 'unverified', 0.15, ARRAY['Propia', 'Local']),
('Almacén Local', 'almacen-local', 'ALMACEN LOCAL', 'Productos de almacenes locales sin marca específica', NULL, ARRAY['generico', 'local'], 'unverified', 0.05, ARRAY['Local', 'Almacén']);

-- Comentario sobre la tabla
COMMENT ON TABLE marketplace_brands IS 'Catálogo inicial de marcas populares en Argentina con datos curados manualmente';

-- Mostrar estadísticas de la carga
DO $$
DECLARE 
    total_brands INTEGER;
    verified_brands INTEGER;
    categories_covered TEXT[];
BEGIN
    SELECT COUNT(*) INTO total_brands FROM marketplace_brands;
    SELECT COUNT(*) INTO verified_brands FROM marketplace_brands WHERE verification_status = 'verified';
    SELECT ARRAY_AGG(DISTINCT unnest) INTO categories_covered FROM (
        SELECT unnest(category_tags) FROM marketplace_brands WHERE verification_status = 'verified'
    ) t;
    
    RAISE NOTICE '✅ Seed completado exitosamente:';
    RAISE NOTICE '   📊 Total marcas cargadas: %', total_brands;
    RAISE NOTICE '   ✅ Marcas verificadas: %', verified_brands;
    RAISE NOTICE '   🏷️  Categorías cubiertas: %', array_length(categories_covered, 1);
    RAISE NOTICE '   📈 Puntuación promedio: %.2f', (SELECT AVG(quality_score) FROM marketplace_brands WHERE verification_status = 'verified');
END;
$$; 