-- Migración: Crear tabla para catálogo global de productos argentinos
-- Fecha: 2024-12-18
-- Descripción: Tabla principal para almacenar productos del catálogo global con EAN-13

-- Crear extensión UUID si no existe
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Crear tabla global_products
CREATE TABLE IF NOT EXISTS global_products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ean VARCHAR(13) NOT NULL UNIQUE,
    name VARCHAR(500) NOT NULL,
    description TEXT,
    brand VARCHAR(200),
    category VARCHAR(200),
    price DECIMAL(10,2),
    image_url TEXT,
    image_urls TEXT[], -- Array de URLs de imágenes adicionales
    
    -- Campos de fuente y calidad
    source VARCHAR(50) NOT NULL,
    source_url TEXT,
    source_reliability DECIMAL(3,2) DEFAULT 0.5,
    quality_score INTEGER NOT NULL DEFAULT 0 CHECK (quality_score >= 0 AND quality_score <= 100),
    
    -- Campos de estado
    is_verified BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    business_type VARCHAR(100),
    
    -- Campos de metadatos
    tags TEXT[], -- Array de etiquetas
    metadata JSONB DEFAULT '{}',
    
    -- Campos de auditoría
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    scraped_at TIMESTAMP WITH TIME ZONE,
    last_scraped_at TIMESTAMP WITH TIME ZONE
);

-- Crear índices para optimizar búsquedas
CREATE INDEX IF NOT EXISTS idx_global_products_ean ON global_products(ean);
CREATE INDEX IF NOT EXISTS idx_global_products_name ON global_products USING gin(to_tsvector('spanish', name));
CREATE INDEX IF NOT EXISTS idx_global_products_brand ON global_products(brand);
CREATE INDEX IF NOT EXISTS idx_global_products_category ON global_products(category);
CREATE INDEX IF NOT EXISTS idx_global_products_source ON global_products(source);
CREATE INDEX IF NOT EXISTS idx_global_products_business_type ON global_products(business_type);
CREATE INDEX IF NOT EXISTS idx_global_products_quality_score ON global_products(quality_score);
CREATE INDEX IF NOT EXISTS idx_global_products_is_active ON global_products(is_active);
CREATE INDEX IF NOT EXISTS idx_global_products_is_verified ON global_products(is_verified);
CREATE INDEX IF NOT EXISTS idx_global_products_tags ON global_products USING gin(tags);
CREATE INDEX IF NOT EXISTS idx_global_products_metadata ON global_products USING gin(metadata);
CREATE INDEX IF NOT EXISTS idx_global_products_created_at ON global_products(created_at);
CREATE INDEX IF NOT EXISTS idx_global_products_last_scraped_at ON global_products(last_scraped_at);

-- Índice compuesto para búsquedas por tipo de negocio y calidad
CREATE INDEX IF NOT EXISTS idx_global_products_business_quality ON global_products(business_type, quality_score DESC, is_active);

-- Índice para productos argentinos (EAN que empiecen con 779, 780-799)
CREATE INDEX IF NOT EXISTS idx_global_products_argentine ON global_products(ean) 
WHERE ean LIKE '779%' OR (ean >= '7800000000000' AND ean <= '7999999999999');

-- Función para actualizar updated_at automáticamente
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Trigger para actualizar updated_at automáticamente
CREATE TRIGGER update_global_products_updated_at 
    BEFORE UPDATE ON global_products 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Comentarios para documentar la tabla
COMMENT ON TABLE global_products IS 'Catálogo global de productos argentinos con información agregada de múltiples fuentes';
COMMENT ON COLUMN global_products.ean IS 'Código de barras EAN-13 único del producto';
COMMENT ON COLUMN global_products.quality_score IS 'Puntuación de calidad de 0-100 basada en completitud de datos';
COMMENT ON COLUMN global_products.source IS 'Fuente del producto (disco, carrefour, fravega, manual, etc)';
COMMENT ON COLUMN global_products.source_reliability IS 'Confiabilidad de la fuente (0.0-1.0)';
COMMENT ON COLUMN global_products.business_type IS 'Tipo de negocio recomendado para el producto';
COMMENT ON COLUMN global_products.metadata IS 'Metadatos adicionales en formato JSON';
COMMENT ON COLUMN global_products.tags IS 'Etiquetas del producto para clasificación adicional';

-- Insertar algunos productos de ejemplo para testing
INSERT INTO global_products (ean, name, description, brand, category, price, source, quality_score, business_type, tags) VALUES
('7791234567890', 'Coca Cola 500ml', 'Bebida cola sabor original 500ml', 'Coca Cola', 'Bebidas', 150.00, 'manual', 85, 'kiosco', ARRAY['bebida', 'cola', 'gaseosa']),
('7791111222333', 'Pan Lactal Bimbo', 'Pan de molde lactal 450g', 'Bimbo', 'Panadería', 280.50, 'manual', 75, 'almacen', ARRAY['pan', 'lactal', 'panaderia']),
('7792345678901', 'Leche La Serenísima 1L', 'Leche entera pasteurizada 1 litro', 'La Serenísima', 'Lácteos', 320.00, 'manual', 90, 'almacen', ARRAY['leche', 'lacteos', 'serenisima'])
ON CONFLICT (ean) DO NOTHING; 