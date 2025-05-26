-- Crear tabla products
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    sku VARCHAR(50),
    category_id UUID,
    category_name VARCHAR(255),
    brand_id UUID,
    brand_name VARCHAR(255),
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    -- Constraints
    CONSTRAINT products_status_check CHECK (status IN ('active', 'inactive', 'discontinued', 'deleted')),
    CONSTRAINT products_name_length_check CHECK (LENGTH(name) >= 2 AND LENGTH(name) <= 255),
    CONSTRAINT products_description_length_check CHECK (description IS NULL OR LENGTH(description) <= 1000),
    CONSTRAINT products_sku_length_check CHECK (sku IS NULL OR (LENGTH(sku) >= 3 AND LENGTH(sku) <= 50))
);

-- Índices para optimizar consultas
CREATE INDEX IF NOT EXISTS idx_products_tenant_id ON products(tenant_id);
CREATE INDEX IF NOT EXISTS idx_products_status ON products(status);
CREATE INDEX IF NOT EXISTS idx_products_category_id ON products(category_id);
CREATE INDEX IF NOT EXISTS idx_products_brand_id ON products(brand_id);
CREATE INDEX IF NOT EXISTS idx_products_created_at ON products(created_at);
CREATE INDEX IF NOT EXISTS idx_products_name ON products(name);

-- Índices únicos para evitar duplicados
CREATE UNIQUE INDEX IF NOT EXISTS idx_products_tenant_name_unique 
    ON products(tenant_id, name) 
    WHERE status != 'deleted';

CREATE UNIQUE INDEX IF NOT EXISTS idx_products_tenant_sku_unique 
    ON products(tenant_id, sku) 
    WHERE sku IS NOT NULL AND status != 'deleted';

-- Índices compuestos para consultas comunes
CREATE INDEX IF NOT EXISTS idx_products_tenant_status ON products(tenant_id, status);
CREATE INDEX IF NOT EXISTS idx_products_tenant_category ON products(tenant_id, category_id);
CREATE INDEX IF NOT EXISTS idx_products_tenant_brand ON products(tenant_id, brand_id);

-- Trigger para actualizar updated_at automáticamente
CREATE OR REPLACE FUNCTION update_products_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_products_updated_at
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE FUNCTION update_products_updated_at();

-- Comentarios para documentación
COMMENT ON TABLE products IS 'Tabla de productos del sistema PIM';
COMMENT ON COLUMN products.id IS 'Identificador único del producto';
COMMENT ON COLUMN products.tenant_id IS 'ID del tenant al que pertenece el producto';
COMMENT ON COLUMN products.name IS 'Nombre del producto (2-255 caracteres)';
COMMENT ON COLUMN products.description IS 'Descripción del producto (máximo 1000 caracteres)';
COMMENT ON COLUMN products.sku IS 'Código SKU del producto (3-50 caracteres, único por tenant)';
COMMENT ON COLUMN products.category_id IS 'ID de la categoría (referencia desacoplada)';
COMMENT ON COLUMN products.category_name IS 'Nombre de la categoría (desnormalizado para performance)';
COMMENT ON COLUMN products.brand_id IS 'ID de la marca (referencia desacoplada)';
COMMENT ON COLUMN products.brand_name IS 'Nombre de la marca (desnormalizado para performance)';
COMMENT ON COLUMN products.status IS 'Estado del producto: active, inactive, discontinued, deleted';
COMMENT ON COLUMN products.created_at IS 'Fecha y hora de creación';
COMMENT ON COLUMN products.updated_at IS 'Fecha y hora de última actualización'; 