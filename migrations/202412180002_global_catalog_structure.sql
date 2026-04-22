-- Migration 023: Global catalog structure
-- Crea tablas auxiliares del catálogo global (variaciones, links tenant, scraping)
-- global_products se crea en 202412180001
-- business_type_product_templates se crea en 044

-- Variaciones de productos globales
CREATE TABLE IF NOT EXISTS global_product_variations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    global_product_id UUID NOT NULL REFERENCES global_products(id) ON DELETE CASCADE,
    variation_type VARCHAR(50) NOT NULL,
    variation_value VARCHAR(100) NOT NULL,
    ean VARCHAR(13),
    sku_variation VARCHAR(100),
    specifications_diff JSONB DEFAULT '{}',
    price_difference DECIMAL(10,2),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(global_product_id, variation_type, variation_value)
);

CREATE INDEX IF NOT EXISTS idx_global_product_variations_product ON global_product_variations(global_product_id);
CREATE INDEX IF NOT EXISTS idx_global_product_variations_ean ON global_product_variations(ean);

-- Vinculación entre productos de tenant y catálogo global
CREATE TABLE IF NOT EXISTS tenant_global_product_links (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    tenant_product_id UUID NOT NULL,
    global_product_id UUID NOT NULL REFERENCES global_products(id),
    price DECIMAL(12,2),
    stock_quantity INTEGER DEFAULT 0,
    is_available BOOLEAN DEFAULT true,
    custom_name VARCHAR(500),
    custom_description TEXT,
    custom_images JSONB DEFAULT '[]',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(tenant_id, tenant_product_id),
    UNIQUE(tenant_id, global_product_id)
);

CREATE INDEX IF NOT EXISTS idx_tenant_global_links_tenant ON tenant_global_product_links(tenant_id);
CREATE INDEX IF NOT EXISTS idx_tenant_global_links_global ON tenant_global_product_links(global_product_id);
CREATE INDEX IF NOT EXISTS idx_tenant_global_links_available ON tenant_global_product_links(is_available);

-- Log de scraping
CREATE TABLE IF NOT EXISTS scraping_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    job_id VARCHAR(100),
    target_site VARCHAR(100) NOT NULL,
    target_url TEXT NOT NULL,
    status VARCHAR(20) CHECK (status IN ('success', 'failed', 'partial', 'skipped')),
    products_found INTEGER DEFAULT 0,
    products_processed INTEGER DEFAULT 0,
    duration_ms INTEGER,
    error_message TEXT,
    user_agent TEXT,
    started_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_scraping_logs_site ON scraping_logs(target_site);
CREATE INDEX IF NOT EXISTS idx_scraping_logs_status ON scraping_logs(status);
CREATE INDEX IF NOT EXISTS idx_scraping_logs_started ON scraping_logs(started_at);

-- Configuración de scraping por sitio
CREATE TABLE IF NOT EXISTS scraping_targets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    base_url TEXT NOT NULL,
    is_active BOOLEAN DEFAULT true,
    rate_limit_rps DECIMAL(4,2) DEFAULT 1.0,
    delay_between_requests_ms INTEGER DEFAULT 1000,
    headers JSONB DEFAULT '{}',
    selectors JSONB DEFAULT '{}',
    categories JSONB DEFAULT '[]',
    respect_robots_txt BOOLEAN DEFAULT true,
    robots_txt_url TEXT,
    last_robots_check TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO scraping_targets (name, base_url, rate_limit_rps, categories, selectors) VALUES
('disco', 'https://www.disco.com.ar', 2.0,
 '["bebidas", "lacteos", "limpieza", "snacks"]',
 '{"title": ".product-title", "price": ".price", "ean": ".product-ean", "brand": ".brand", "image": ".product-image img"}'),
('carrefour', 'https://www.carrefour.com.ar', 1.5,
 '["electronica", "hogar", "bebidas", "alimentos"]',
 '{"title": "h1.product-name", "price": ".price-current", "ean": ".product-code", "brand": ".brand-name", "image": ".product-image img"}'),
('fravega', 'https://www.fravega.com', 1.0,
 '["celulares", "computacion", "electrodomesticos"]',
 '{"title": ".product-title", "price": ".price-box .price", "ean": ".sku", "brand": ".brand", "image": ".product-image img"}')
ON CONFLICT (name) DO NOTHING;
