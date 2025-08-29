-- Seed: 009_marketplace_brands_regional_argentina.sql
-- Purpose: Agregar marcas regionales específicas de Argentina y marcas propias de cadenas
-- Date: 2025-01-20
-- Beneficio: Mejorar detección de marcas locales y manejo de marcas propias

-- ============================================
-- MARCAS REGIONALES POR PROVINCIA/REGIÓN
-- ============================================

-- Mendoza
INSERT INTO marketplace_brands (name, slug, normalized_name, description, category_tags, verification_status, quality_score, aliases, country_code) VALUES
('La Campagnola', 'la-campagnola', 'LA CAMPAGNOLA', 'Marca de conservas, dulces y mermeladas originaria de Mendoza', ARRAY['alimentos', 'conservas', 'dulces'], 'verified', 0.78, ARRAY['Campagnola', 'Molto'], 'AR'),
('Canale', 'canale', 'CANALE', 'Marca regional de arroz de Mendoza', ARRAY['alimentos', 'arroz'], 'verified', 0.72, ARRAY['Arroz Canale'], 'AR'),
('Baggio', 'baggio', 'BAGGIO', 'Jugos y bebidas de Mendoza', ARRAY['bebidas', 'jugos'], 'verified', 0.76, ARRAY['Jugos Baggio'], 'AR'),
('Matas', 'matas', 'MATAS', 'Vinos y aceites de Mendoza', ARRAY['bebidas', 'vinos', 'aceites'], 'verified', 0.73, NULL, 'AR')
ON CONFLICT (slug) DO UPDATE SET
    aliases = EXCLUDED.aliases,
    quality_score = EXCLUDED.quality_score,
    category_tags = EXCLUDED.category_tags;

-- Córdoba
INSERT INTO marketplace_brands (name, slug, normalized_name, description, category_tags, verification_status, quality_score, aliases, country_code) VALUES
('Georgalos', 'georgalos', 'GEORGALOS', 'Fábrica de golosinas y caramelos de Córdoba', ARRAY['golosinas', 'caramelos'], 'verified', 0.74, ARRAY['Georgalos Hnos'], 'AR'),
('Fantoche', 'fantoche', 'FANTOCHE', 'Helados artesanales de Córdoba', ARRAY['helados', 'postres'], 'verified', 0.71, NULL, 'AR'),
('Los Cortaderos', 'los-cortaderos', 'LOS CORTADEROS', 'Productos lácteos artesanales de Córdoba', ARRAY['lacteos', 'quesos'], 'verified', 0.69, NULL, 'AR')
ON CONFLICT (slug) DO UPDATE SET
    aliases = EXCLUDED.aliases,
    quality_score = EXCLUDED.quality_score;

-- Buenos Aires Interior
INSERT INTO marketplace_brands (name, slug, normalized_name, description, category_tags, verification_status, quality_score, aliases, country_code) VALUES
('Vauquita', 'vauquita', 'VAUQUITA', 'Productos lácteos del interior de Buenos Aires', ARRAY['lacteos'], 'verified', 0.70, ARRAY['La Vauquita', 'Lácteos Vauquita'], 'AR'),
('Don Satur', 'don-satur', 'DON SATUR', 'Bizcochos tradicionales argentinos', ARRAY['galletitas', 'panificados'], 'verified', 0.79, ARRAY['Bizcochos Don Satur'], 'AR'),
('Manaos', 'manaos', 'MANAOS', 'Bebidas gaseosas económicas', ARRAY['bebidas', 'gaseosas'], 'verified', 0.65, NULL, 'AR')
ON CONFLICT (slug) DO UPDATE SET
    aliases = EXCLUDED.aliases,
    quality_score = EXCLUDED.quality_score;

-- NOA (Salta, Jujuy, Tucumán)
INSERT INTO marketplace_brands (name, slug, normalized_name, description, category_tags, verification_status, quality_score, aliases, country_code) VALUES
('La Morenita', 'la-morenita', 'LA MORENITA', 'Yerba mate del norte argentino', ARRAY['alimentos', 'yerba'], 'verified', 0.75, ARRAY['Yerba La Morenita'], 'AR'),
('Ledesma', 'ledesma', 'LEDESMA', 'Azúcar y papel de Jujuy', ARRAY['alimentos', 'azucar'], 'verified', 0.77, ARRAY['Azúcar Ledesma'], 'AR'),
('Arcor NOA', 'arcor-noa', 'ARCOR NOA', 'Productos Arcor fabricados en el NOA', ARRAY['golosinas', 'alimentos'], 'verified', 0.80, ARRAY['Arcor Tucumán'], 'AR'),
('Los Nietitos', 'los-nietitos', 'LOS NIETITOS', 'Dulces regionales y alfajores del NOA', ARRAY['dulces', 'golosinas'], 'verified', 0.68, NULL, 'AR')
ON CONFLICT (slug) DO UPDATE SET
    aliases = EXCLUDED.aliases,
    quality_score = EXCLUDED.quality_score;

-- NEA (Misiones, Corrientes, Chaco, Formosa)
INSERT INTO marketplace_brands (name, slug, normalized_name, description, category_tags, verification_status, quality_score, aliases, country_code) VALUES
('Las Marías', 'las-marias', 'LAS MARIAS', 'Yerba mate de Corrientes', ARRAY['alimentos', 'yerba'], 'verified', 0.82, ARRAY['Taragui', 'Union', 'La Merced'], 'AR'),
('Rosamonte', 'rosamonte', 'ROSAMONTE', 'Yerba mate de Misiones', ARRAY['alimentos', 'yerba'], 'verified', 0.78, NULL, 'AR'),
('Amanda', 'amanda', 'AMANDA', 'Yerba mate de Misiones', ARRAY['alimentos', 'yerba'], 'verified', 0.81, ARRAY['Yerba Amanda'], 'AR')
ON CONFLICT (slug) DO UPDATE SET
    aliases = EXCLUDED.aliases,
    quality_score = EXCLUDED.quality_score;

-- Patagonia
INSERT INTO marketplace_brands (name, slug, normalized_name, description, category_tags, verification_status, quality_score, aliases, country_code) VALUES
('La Anónima', 'la-anonima-marca', 'LA ANONIMA', 'Marca propia de Supermercados La Anónima', ARRAY['generico', 'marca-propia'], 'verified', 0.60, ARRAY['Supermercados La Anónima', 'Anónima'], 'AR'),
('Patagonia', 'patagonia-cerveza', 'PATAGONIA', 'Cerveza artesanal patagónica', ARRAY['bebidas', 'cerveza'], 'verified', 0.83, ARRAY['Cerveza Patagonia'], 'AR'),
('Mamusia', 'mamusia', 'MAMUSIA', 'Chocolates artesanales de Bariloche', ARRAY['golosinas', 'chocolates'], 'verified', 0.77, NULL, 'AR'),
('Rapa Nui', 'rapa-nui', 'RAPA NUI', 'Chocolates de Bariloche', ARRAY['golosinas', 'chocolates'], 'verified', 0.76, NULL, 'AR')
ON CONFLICT (slug) DO UPDATE SET
    aliases = EXCLUDED.aliases,
    quality_score = EXCLUDED.quality_score;

-- Cuyo (San Juan, La Rioja)
INSERT INTO marketplace_brands (name, slug, normalized_name, description, category_tags, verification_status, quality_score, aliases, country_code) VALUES
('Alco', 'alco', 'ALCO', 'Alimentos en conserva de San Juan', ARRAY['alimentos', 'conservas'], 'verified', 0.71, ARRAY['Alco Canale'], 'AR'),
('La Riojana', 'la-riojana', 'LA RIOJANA', 'Cooperativa vitivinícola de La Rioja', ARRAY['bebidas', 'vinos'], 'verified', 0.74, NULL, 'AR')
ON CONFLICT (slug) DO UPDATE SET
    aliases = EXCLUDED.aliases,
    quality_score = EXCLUDED.quality_score;

-- ============================================
-- MARCAS PROPIAS DE CADENAS DE SUPERMERCADOS
-- ============================================
-- Estas marcas deben ser identificadas como "Sin Marca" en el proceso de curación

INSERT INTO marketplace_brands (name, slug, normalized_name, description, category_tags, verification_status, quality_score, aliases, country_code) VALUES
-- Carrefour
('Carrefour', 'carrefour-marca-propia', 'CARREFOUR', 'Marca propia de Carrefour', ARRAY['marca-propia'], 'verified', 0.20, ARRAY['Marca Carrefour'], 'AR'),
('TEX', 'tex', 'TEX', 'Marca propia de ropa de Carrefour', ARRAY['marca-propia', 'indumentaria'], 'verified', 0.20, NULL, 'AR'),
('First Line', 'first-line', 'FIRST LINE', 'Marca propia económica de Carrefour', ARRAY['marca-propia'], 'verified', 0.15, NULL, 'AR'),

-- DIA
('DIA', 'dia-marca-propia', 'DIA', 'Marca propia de supermercados DIA', ARRAY['marca-propia'], 'verified', 0.20, ARRAY['Marca DIA', 'Día'], 'AR'),
('Bell''s', 'bells', 'BELLS', 'Marca propia de DIA', ARRAY['marca-propia', 'alimentos'], 'verified', 0.18, ARRAY['Bells'], 'AR'),
('Bonté', 'bonte', 'BONTE', 'Marca propia de cosmética de DIA', ARRAY['marca-propia', 'cuidado-personal'], 'verified', 0.18, NULL, 'AR'),
('Delicious', 'delicious-dia', 'DELICIOUS', 'Marca propia premium de DIA', ARRAY['marca-propia', 'alimentos'], 'verified', 0.20, NULL, 'AR'),

-- Coto
('Coto', 'coto-marca-propia', 'COTO', 'Marca propia de supermercados Coto', ARRAY['marca-propia'], 'verified', 0.20, ARRAY['Marca Coto'], 'AR'),
('Ciudad del Lago', 'ciudad-del-lago', 'CIUDAD DEL LAGO', 'Marca propia premium de Coto', ARRAY['marca-propia', 'alimentos'], 'verified', 0.22, NULL, 'AR'),

-- Walmart/Chango Más
('Great Value', 'great-value', 'GREAT VALUE', 'Marca propia de Walmart', ARRAY['marca-propia'], 'verified', 0.20, ARRAY['GreatValue'], 'AR'),
('Mainstays', 'mainstays', 'MAINSTAYS', 'Marca propia de hogar de Walmart', ARRAY['marca-propia', 'hogar'], 'verified', 0.18, NULL, 'AR'),

-- Jumbo/Disco/Vea
('Jumbo', 'jumbo-marca-propia', 'JUMBO', 'Marca propia de hipermercados Jumbo', ARRAY['marca-propia'], 'verified', 0.20, ARRAY['Marca Jumbo'], 'AR'),
('Econo', 'econo', 'ECONO', 'Marca propia económica de Disco', ARRAY['marca-propia'], 'verified', 0.15, NULL, 'AR'),
('Qualitá', 'qualita', 'QUALITA', 'Marca propia del grupo Jumbo/Disco', ARRAY['marca-propia'], 'verified', 0.18, NULL, 'AR'),

-- Farmacity
('Farmacity', 'farmacity-marca-propia', 'FARMACITY', 'Marca propia de Farmacity', ARRAY['marca-propia', 'farmacia'], 'verified', 0.20, ARRAY['Marca Farmacity'], 'AR'),
('Get the Look', 'get-the-look', 'GET THE LOOK', 'Marca propia de cosmética de Farmacity', ARRAY['marca-propia', 'cosmetica'], 'verified', 0.20, NULL, 'AR'),

-- Más Online (ex Libertad)
('Más', 'mas-marca-propia', 'MAS', 'Marca propia de Más Online', ARRAY['marca-propia'], 'verified', 0.18, ARRAY['Mas', 'Más Online'], 'AR')
ON CONFLICT (slug) DO UPDATE SET
    aliases = EXCLUDED.aliases,
    quality_score = EXCLUDED.quality_score,
    category_tags = EXCLUDED.category_tags;

-- ============================================
-- ACTUALIZAR MARCA GENÉRICA
-- ============================================
UPDATE marketplace_brands 
SET 
    aliases = array_cat(
        COALESCE(aliases, '{}'),
        ARRAY[
            'S/Marca', 
            'Sin marca',
            'Genérico',
            'Producto genérico',
            'Marca blanca',
            'Private label'
        ]
    ),
    description = 'Productos sin marca específica o marcas propias de cadenas identificadas'
WHERE slug = 'sin-marca';

-- ============================================
-- ESTADÍSTICAS DE LA CARGA
-- ============================================
DO $$
DECLARE 
    total_regional INTEGER;
    total_private_label INTEGER;
    total_brands INTEGER;
BEGIN
    SELECT COUNT(*) INTO total_regional 
    FROM marketplace_brands 
    WHERE category_tags && ARRAY['alimentos', 'bebidas', 'lacteos'] 
    AND verification_status = 'verified'
    AND quality_score > 0.60;
    
    SELECT COUNT(*) INTO total_private_label 
    FROM marketplace_brands 
    WHERE 'marca-propia' = ANY(category_tags);
    
    SELECT COUNT(*) INTO total_brands 
    FROM marketplace_brands;
    
    RAISE NOTICE '✅ Seeds de marcas regionales completados:';
    RAISE NOTICE '   📍 Marcas regionales argentinas: %', total_regional;
    RAISE NOTICE '   🏪 Marcas propias de cadenas: %', total_private_label;
    RAISE NOTICE '   📊 Total marcas en sistema: %', total_brands;
    RAISE NOTICE '';
    RAISE NOTICE '   Regiones cubiertas:';
    RAISE NOTICE '   - Mendoza (Canale, La Campagnola, Baggio)';
    RAISE NOTICE '   - Córdoba (Georgalos, Fantoche)';
    RAISE NOTICE '   - NOA (La Morenita, Ledesma)';
    RAISE NOTICE '   - NEA (Las Marías, Rosamonte, Amanda)';
    RAISE NOTICE '   - Patagonia (La Anónima, Patagonia, Mamusia)';
    RAISE NOTICE '   - Buenos Aires (Vauquita, Don Satur, Manaos)';
    RAISE NOTICE '';
    RAISE NOTICE '   Cadenas con marcas propias identificadas:';
    RAISE NOTICE '   - Carrefour (TEX, First Line)';
    RAISE NOTICE '   - DIA (Bell''s, Bonté, Delicious)';
    RAISE NOTICE '   - Coto (Ciudad del Lago)';
    RAISE NOTICE '   - Walmart (Great Value, Mainstays)';
    RAISE NOTICE '   - Jumbo/Disco (Econo, Qualitá)';
    RAISE NOTICE '   - Farmacity (Get the Look)';
    RAISE NOTICE '   - La Anónima';
    RAISE NOTICE '   - Más Online';
END;
$$;