-- ================================================
-- GLOBAL CATALOG STRUCTURE
-- Migración 005: Tablas para catálogo global de productos argentinos
-- ================================================

-- Tabla principal de productos globales
CREATE TABLE global_products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Identificadores únicos
    ean VARCHAR(13) UNIQUE,                     -- Código EAN-13 (obligatorio)
    gtin VARCHAR(14),                           -- GTIN alternativo
    sku_global VARCHAR(100) UNIQUE,             -- SKU interno global
    
    -- Información básica
    name VARCHAR(500) NOT NULL,                 -- "Coca Cola 600ml"
    brand VARCHAR(200) NOT NULL,                -- "Coca Cola"
    manufacturer VARCHAR(200),                  -- "Coca Cola FEMSA"
    description TEXT,                           -- Descripción detallada
    
    -- Categorización
    marketplace_category_id UUID REFERENCES marketplace_categories(id),
    product_type VARCHAR(20) CHECK (product_type IN ('industrialized', 'fmcg', 'local_brand')),
    
    -- Especificaciones técnicas (JSONB para flexibilidad)
    specifications JSONB DEFAULT '{}',          -- Peso, dimensiones, ingredientes, etc.
    technical_specs JSONB DEFAULT '{}',         -- Para electrónicos: CPU, RAM, etc.
    nutritional_info JSONB DEFAULT '{}',        -- Para alimentos: calorías, etc.
    
    -- Imágenes y multimedia
    images JSONB DEFAULT '[]',                  -- URLs de imágenes
    primary_image_url TEXT,                     -- Imagen principal
    video_urls JSONB DEFAULT '[]',              -- Videos del producto
    
    -- Metadatos de origen
    source_type VARCHAR(20) CHECK (source_type IN ('scraping', 'api', 'manual', 'partnership')),
    source_url TEXT,                            -- URL original
    source_confidence DECIMAL(3,2) CHECK (source_confidence >= 0 AND source_confidence <= 1),
    
    -- Regulatorio argentino
    anmat_code VARCHAR(50),                     -- Código ANMAT para alimentos/farmacos
    senasa_code VARCHAR(50),                    -- Código SENASA para productos agropecuarios
    
    -- Estado y calidad
    verification_status VARCHAR(10) CHECK (verification_status IN ('pending', 'verified', 'rejected')) DEFAULT 'pending',
    quality_score DECIMAL(3,2) CHECK (quality_score >= 0 AND quality_score <= 1),
    popularity_rank INTEGER DEFAULT 0,          -- Ranking de popularidad
    is_active BOOLEAN DEFAULT true,
    
    -- Timestamps
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_scraped_at TIMESTAMP WITH TIME ZONE
);

-- Índices para performance
CREATE INDEX idx_global_products_ean ON global_products(ean);
CREATE INDEX idx_global_products_brand ON global_products(brand);
CREATE INDEX idx_global_products_category ON global_products(marketplace_category_id);
CREATE INDEX idx_global_products_type ON global_products(product_type);
CREATE INDEX idx_global_products_verification ON global_products(verification_status);
CREATE INDEX idx_global_products_quality ON global_products(quality_score);
CREATE INDEX idx_global_products_popularity ON global_products(popularity_rank);

-- Índice de texto completo para búsqueda
CREATE INDEX idx_global_products_search ON global_products 
USING gin(to_tsvector('spanish', name || ' ' || brand));

-- Variaciones de productos globales
CREATE TABLE global_product_variations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    global_product_id UUID NOT NULL REFERENCES global_products(id) ON DELETE CASCADE,
    
    -- Variaciones (color, tamaño, sabor, etc.)
    variation_type VARCHAR(50) NOT NULL,        -- "color", "size", "flavor"
    variation_value VARCHAR(100) NOT NULL,      -- "rojo", "500ml", "frutilla"
    
    -- Identificadores específicos de la variación
    ean VARCHAR(13),                            -- EAN específico de esta variación
    sku_variation VARCHAR(100),                 -- SKU específico
    
    -- Diferencias de specs
    specifications_diff JSONB DEFAULT '{}',     -- Solo campos que difieren
    price_difference DECIMAL(10,2),             -- Diferencia de precio base
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(global_product_id, variation_type, variation_value)
);

-- Índices para variaciones
CREATE INDEX idx_global_product_variations_product ON global_product_variations(global_product_id);
CREATE INDEX idx_global_product_variations_ean ON global_product_variations(ean);

-- Templates de productos por tipo de negocio (extensión de templates existentes)
CREATE TABLE business_type_product_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_type_template_id UUID NOT NULL REFERENCES business_type_templates(id) ON DELETE CASCADE,
    
    -- Productos sugeridos por tipo de negocio
    suggested_products JSONB DEFAULT '[]',      -- Array de global_product IDs
    
    -- Configuración de sugerencias
    max_products_per_category INTEGER DEFAULT 50,
    priority_brands JSONB DEFAULT '[]',         -- Marcas prioritarias para este tipo
    exclude_brands JSONB DEFAULT '[]',          -- Marcas a excluir
    
    -- Metadatos
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(business_type_template_id)
);

-- Vinculación entre productos de tenant y catálogo global
CREATE TABLE tenant_global_product_links (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    tenant_product_id UUID NOT NULL,           -- ID del producto en el catálogo del tenant
    global_product_id UUID NOT NULL REFERENCES global_products(id),
    
    -- Metadatos de la vinculación
    price DECIMAL(12,2),                       -- Precio específico del tenant
    stock_quantity INTEGER DEFAULT 0,          -- Stock actual
    is_available BOOLEAN DEFAULT true,         -- Disponibilidad
    
    -- Customizaciones del tenant
    custom_name VARCHAR(500),                  -- Nombre personalizado (opcional)
    custom_description TEXT,                   -- Descripción personalizada (opcional)
    custom_images JSONB DEFAULT '[]',          -- Imágenes adicionales del tenant
    
    -- Timestamps
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(tenant_id, tenant_product_id),
    UNIQUE(tenant_id, global_product_id)       -- Un tenant no puede linkear el mismo producto global twice
);

-- Índices para links
CREATE INDEX idx_tenant_global_links_tenant ON tenant_global_product_links(tenant_id);
CREATE INDEX idx_tenant_global_links_global ON tenant_global_product_links(global_product_id);
CREATE INDEX idx_tenant_global_links_available ON tenant_global_product_links(is_available);

-- Log de scraping para tracking y debugging
CREATE TABLE scraping_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Información del job de scraping
    job_id VARCHAR(100),                       -- ID único del job
    target_site VARCHAR(100) NOT NULL,        -- "disco", "carrefour", etc.
    target_url TEXT NOT NULL,                  -- URL específica scrapeada
    
    -- Resultado
    status VARCHAR(20) CHECK (status IN ('success', 'failed', 'partial', 'skipped')),
    products_found INTEGER DEFAULT 0,          -- Productos encontrados
    products_processed INTEGER DEFAULT 0,      -- Productos procesados exitosamente
    
    -- Metadatos
    duration_ms INTEGER,                       -- Duración en milisegundos
    error_message TEXT,                        -- Error si falló
    user_agent TEXT,                          -- User-Agent usado
    
    -- Timestamps
    started_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP WITH TIME ZONE
);

-- Índices para scraping logs
CREATE INDEX idx_scraping_logs_site ON scraping_logs(target_site);
CREATE INDEX idx_scraping_logs_status ON scraping_logs(status);
CREATE INDEX idx_scraping_logs_started ON scraping_logs(started_at);

-- Configuración de scraping por sitio
CREATE TABLE scraping_targets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Configuración del target
    name VARCHAR(100) UNIQUE NOT NULL,         -- "disco", "carrefour"
    base_url TEXT NOT NULL,                    -- "https://www.disco.com.ar"
    is_active BOOLEAN DEFAULT true,
    
    -- Rate limiting
    rate_limit_rps DECIMAL(4,2) DEFAULT 1.0,  -- Requests per second
    delay_between_requests_ms INTEGER DEFAULT 1000,
    
    -- Headers y configuración
    headers JSONB DEFAULT '{}',                -- Headers HTTP personalizados
    selectors JSONB DEFAULT '{}',              -- Selectores CSS/XPath
    categories JSONB DEFAULT '[]',             -- Categorías a scrapear
    
    -- Configuración de robots.txt
    respect_robots_txt BOOLEAN DEFAULT true,
    robots_txt_url TEXT,
    last_robots_check TIMESTAMP WITH TIME ZONE,
    
    -- Metadatos
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Targets iniciales para Argentina
INSERT INTO scraping_targets (name, base_url, rate_limit_rps, categories, selectors) VALUES
('disco', 'https://www.disco.com.ar', 2.0, 
 '["bebidas", "lacteos", "limpieza", "snacks"]',
 '{"title": ".product-title", "price": ".price", "ean": ".product-ean", "brand": ".brand", "image": ".product-image img"}'),

('carrefour', 'https://www.carrefour.com.ar', 1.5, 
 '["electronica", "hogar", "bebidas", "alimentos"]',
 '{"title": "h1.product-name", "price": ".price-current", "ean": ".product-code", "brand": ".brand-name", "image": ".product-image img"}'),

('fravega', 'https://www.fravega.com', 1.0, 
 '["celulares", "computacion", "electrodomesticos"]',
 '{"title": ".product-title", "price": ".price-box .price", "ean": ".sku", "brand": ".brand", "image": ".product-image img"}');

-- Función para actualizar updated_at automáticamente
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Triggers para updated_at
CREATE TRIGGER update_global_products_updated_at BEFORE UPDATE ON global_products 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_business_type_product_templates_updated_at BEFORE UPDATE ON business_type_product_templates 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_tenant_global_product_links_updated_at BEFORE UPDATE ON tenant_global_product_links 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_scraping_targets_updated_at BEFORE UPDATE ON scraping_targets 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Comentarios para documentación
COMMENT ON TABLE global_products IS 'Catálogo global de productos argentinos con códigos EAN';
COMMENT ON TABLE global_product_variations IS 'Variaciones de productos globales (tamaños, colores, sabores)';
COMMENT ON TABLE business_type_product_templates IS 'Productos sugeridos por tipo de negocio para quickstart';
COMMENT ON TABLE tenant_global_product_links IS 'Vinculación entre productos específicos de tenant y catálogo global';
COMMENT ON TABLE scraping_logs IS 'Log de actividades de scraping para monitoreo';
COMMENT ON TABLE scraping_targets IS 'Configuración de sitios web para scraping de productos';

COMMENT ON COLUMN global_products.ean IS 'Código EAN-13 único del producto';
COMMENT ON COLUMN global_products.quality_score IS 'Score 0-1 de calidad de datos del producto';
COMMENT ON COLUMN global_products.popularity_rank IS 'Ranking de popularidad basado en ventas/búsquedas';
COMMENT ON COLUMN tenant_global_product_links.price IS 'Precio específico que maneja este tenant';
COMMENT ON COLUMN scraping_targets.rate_limit_rps IS 'Límite de requests por segundo para scraping ético'; 