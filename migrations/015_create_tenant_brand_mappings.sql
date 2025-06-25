-- Migration 015: Create tenant_brand_mappings table
-- PROPÓSITO: Crear tabla de mapeo entre marcas tenant y marcas marketplace globales
-- BENEFICIO: Permite personalización de marcas por tenant manteniendo consistencia global

CREATE TABLE tenant_brand_mappings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id VARCHAR(36) NOT NULL,
    brand_id UUID NOT NULL REFERENCES brands(id) ON DELETE CASCADE, -- marca tenant existente
    marketplace_brand_id UUID NOT NULL REFERENCES marketplace_brands(id) ON DELETE CASCADE,
    custom_name VARCHAR(255), -- Nombre personalizado que ve el tenant
    custom_description TEXT, -- Descripción personalizada
    custom_logo_url VARCHAR(500), -- Logo personalizado por el tenant
    mapping_confidence DECIMAL(3,2) DEFAULT 1.0, -- Confianza del mapping (0.0-1.0)
    mapping_source VARCHAR(50) DEFAULT 'manual', -- manual, auto_curated, suggested
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Constraints
    CONSTRAINT tenant_brand_mappings_confidence_valid CHECK (mapping_confidence >= 0.0 AND mapping_confidence <= 1.0),
    CONSTRAINT tenant_brand_mappings_source_valid CHECK (mapping_source IN ('manual', 'auto_curated', 'suggested', 'imported')),
    CONSTRAINT tenant_brand_mappings_custom_logo_format CHECK (custom_logo_url IS NULL OR custom_logo_url ~ '^https?://'),
    
    -- Unique constraints
    UNIQUE(tenant_id, brand_id),
    UNIQUE(tenant_id, marketplace_brand_id)
);

-- Índices para performance
CREATE INDEX IF NOT EXISTS idx_tenant_brand_mappings_tenant ON tenant_brand_mappings(tenant_id);
CREATE INDEX IF NOT EXISTS idx_tenant_brand_mappings_brand ON tenant_brand_mappings(brand_id);
CREATE INDEX IF NOT EXISTS idx_tenant_brand_mappings_marketplace_brand ON tenant_brand_mappings(marketplace_brand_id);
CREATE INDEX IF NOT EXISTS idx_tenant_brand_mappings_source ON tenant_brand_mappings(mapping_source);
CREATE INDEX IF NOT EXISTS idx_tenant_brand_mappings_confidence ON tenant_brand_mappings(mapping_confidence);
CREATE INDEX IF NOT EXISTS idx_tenant_brand_mappings_active ON tenant_brand_mappings(is_active);

-- Índice compuesto para consultas frecuentes
CREATE INDEX IF NOT EXISTS idx_tenant_brand_mappings_tenant_active ON tenant_brand_mappings(tenant_id, is_active);

-- Comentarios para documentación
COMMENT ON TABLE tenant_brand_mappings IS 'Mapeo de marcas tenant existentes a marcas marketplace globales';
COMMENT ON COLUMN tenant_brand_mappings.tenant_id IS 'ID del tenant propietario del mapeo';
COMMENT ON COLUMN tenant_brand_mappings.brand_id IS 'ID de la marca tenant existente';
COMMENT ON COLUMN tenant_brand_mappings.marketplace_brand_id IS 'ID de la marca marketplace asociada';
COMMENT ON COLUMN tenant_brand_mappings.custom_name IS 'Nombre personalizado que ve el tenant';
COMMENT ON COLUMN tenant_brand_mappings.custom_description IS 'Descripción personalizada por el tenant';
COMMENT ON COLUMN tenant_brand_mappings.custom_logo_url IS 'URL del logo personalizado por el tenant';
COMMENT ON COLUMN tenant_brand_mappings.mapping_confidence IS 'Confianza del mapping automático (0.0-1.0)';
COMMENT ON COLUMN tenant_brand_mappings.mapping_source IS 'Origen del mapping: manual, auto_curated, suggested, imported';

-- Trigger para updated_at
CREATE OR REPLACE FUNCTION update_tenant_brand_mappings_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_tenant_brand_mappings_updated_at
    BEFORE UPDATE ON tenant_brand_mappings
    FOR EACH ROW
    EXECUTE FUNCTION update_tenant_brand_mappings_updated_at();

-- Función para actualizar contadores de productos en marketplace_brands
CREATE OR REPLACE FUNCTION update_marketplace_brand_product_count()
RETURNS TRIGGER AS $$
DECLARE
    brand_count INTEGER;
BEGIN
    -- Si se está insertando o actualizando un mapping activo
    IF TG_OP = 'INSERT' OR (TG_OP = 'UPDATE' AND NEW.is_active = TRUE) THEN
        -- Contar productos que usan marcas mapeadas a esta marketplace_brand
        SELECT COUNT(DISTINCT p.id) INTO brand_count
        FROM products p
        JOIN brands b ON p.brand_id = b.id
        JOIN tenant_brand_mappings tbm ON tbm.brand_id = b.id
        WHERE tbm.marketplace_brand_id = COALESCE(NEW.marketplace_brand_id, OLD.marketplace_brand_id)
        AND tbm.is_active = TRUE;
        
        -- Actualizar el contador en marketplace_brands
        UPDATE marketplace_brands 
        SET product_count = brand_count,
            updated_at = NOW()
        WHERE id = COALESCE(NEW.marketplace_brand_id, OLD.marketplace_brand_id);
    END IF;
    
    -- Si se está eliminando o desactivando un mapping
    IF TG_OP = 'DELETE' OR (TG_OP = 'UPDATE' AND OLD.is_active = TRUE AND NEW.is_active = FALSE) THEN
        -- Recalcular contador para la marketplace_brand antigua
        SELECT COUNT(DISTINCT p.id) INTO brand_count
        FROM products p
        JOIN brands b ON p.brand_id = b.id
        JOIN tenant_brand_mappings tbm ON tbm.brand_id = b.id
        WHERE tbm.marketplace_brand_id = OLD.marketplace_brand_id
        AND tbm.is_active = TRUE;
        
        UPDATE marketplace_brands 
        SET product_count = brand_count,
            updated_at = NOW()
        WHERE id = OLD.marketplace_brand_id;
    END IF;
    
    RETURN COALESCE(NEW, OLD);
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_update_marketplace_brand_product_count
    AFTER INSERT OR UPDATE OR DELETE ON tenant_brand_mappings
    FOR EACH ROW
    EXECUTE FUNCTION update_marketplace_brand_product_count(); 