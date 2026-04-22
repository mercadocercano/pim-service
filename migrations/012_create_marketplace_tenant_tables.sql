-- Migration 012: Create marketplace tenant tables with correct data types
-- PROPÓSITO: Crear tablas de mapeo y atributos tenant con tipos de datos correctos
-- CORRECCIÓN: Usar VARCHAR(36) para category_id para coincidir con la tabla categories existente

-- Reemplaza versión incorrecta de 009 (category_id era UUID, debe ser VARCHAR(36))
DROP TABLE IF EXISTS tenant_custom_attributes CASCADE;
DROP TABLE IF EXISTS tenant_category_mappings CASCADE;

-- Mapeo de categorías tenant a categorías marketplace
CREATE TABLE tenant_category_mappings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id VARCHAR(36) NOT NULL,
    category_id VARCHAR(36) NOT NULL REFERENCES categories(id) ON DELETE CASCADE, -- categoría tenant existente
    marketplace_category_id UUID NOT NULL REFERENCES marketplace_categories(id) ON DELETE CASCADE,
    custom_name VARCHAR(255), -- Nombre personalizado que ve el tenant
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(tenant_id, category_id),
    UNIQUE(tenant_id, marketplace_category_id)
);

-- Atributos completamente custom por tenant
CREATE TABLE tenant_custom_attributes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id VARCHAR(36) NOT NULL,
    marketplace_category_id UUID REFERENCES marketplace_categories(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    is_filterable BOOLEAN DEFAULT FALSE,
    is_searchable BOOLEAN DEFAULT FALSE,
    validation_rules JSONB DEFAULT '{}',
    sort_order INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Constraints
    CONSTRAINT tenant_custom_attributes_type_check CHECK (type IN ('text', 'number', 'boolean', 'select', 'multi_select')),
    CONSTRAINT tenant_custom_attributes_name_not_empty CHECK (LENGTH(TRIM(name)) > 0),
    CONSTRAINT tenant_custom_attributes_slug_not_empty CHECK (LENGTH(TRIM(slug)) > 0),
    UNIQUE(tenant_id, marketplace_category_id, slug)
);

-- Índices para performance
CREATE INDEX IF NOT EXISTS idx_tenant_category_mappings_tenant ON tenant_category_mappings(tenant_id);
CREATE INDEX IF NOT EXISTS idx_tenant_category_mappings_category ON tenant_category_mappings(category_id);
CREATE INDEX IF NOT EXISTS idx_tenant_category_mappings_marketplace ON tenant_category_mappings(marketplace_category_id);

CREATE INDEX IF NOT EXISTS idx_tenant_custom_attributes_tenant ON tenant_custom_attributes(tenant_id);
CREATE INDEX IF NOT EXISTS idx_tenant_custom_attributes_category ON tenant_custom_attributes(marketplace_category_id);
CREATE INDEX IF NOT EXISTS idx_tenant_custom_attributes_slug ON tenant_custom_attributes(slug);

-- Comentarios para documentación
COMMENT ON TABLE tenant_category_mappings IS 'Mapeo de categorías tenant existentes a categorías marketplace';
COMMENT ON COLUMN tenant_category_mappings.tenant_id IS 'ID del tenant propietario del mapeo';
COMMENT ON COLUMN tenant_category_mappings.category_id IS 'ID de la categoría tenant existente';
COMMENT ON COLUMN tenant_category_mappings.marketplace_category_id IS 'ID de la categoría marketplace asociada';
COMMENT ON COLUMN tenant_category_mappings.custom_name IS 'Nombre personalizado que ve el tenant';

COMMENT ON TABLE tenant_custom_attributes IS 'Atributos completamente personalizados por tenant';
COMMENT ON COLUMN tenant_custom_attributes.tenant_id IS 'ID del tenant propietario del atributo';
COMMENT ON COLUMN tenant_custom_attributes.marketplace_category_id IS 'Categoría marketplace asociada (opcional)';
COMMENT ON COLUMN tenant_custom_attributes.slug IS 'Identificador único del atributo dentro del tenant';

-- Triggers para updated_at
CREATE OR REPLACE FUNCTION update_tenant_tables_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_tenant_category_mappings_updated_at
    BEFORE UPDATE ON tenant_category_mappings
    FOR EACH ROW
    EXECUTE FUNCTION update_tenant_tables_updated_at();

CREATE TRIGGER trigger_tenant_custom_attributes_updated_at
    BEFORE UPDATE ON tenant_custom_attributes
    FOR EACH ROW
    EXECUTE FUNCTION update_tenant_tables_updated_at(); 