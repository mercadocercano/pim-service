-- Migration 009: Create marketplace_attributes tables
-- PROPÓSITO: Atributos globales para filtros consistentes cross-tenant
-- BENEFICIO: Filtros estándares que funcionan en todo el marketplace

CREATE TABLE marketplace_attributes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    type VARCHAR(50) NOT NULL, -- text, number, boolean, select, multi_select
    is_filterable BOOLEAN DEFAULT FALSE,
    is_searchable BOOLEAN DEFAULT FALSE,
    is_required_for_listing BOOLEAN DEFAULT FALSE,
    validation_rules JSONB DEFAULT '{}', -- {"min": 1, "max": 100, "regex": "..."}
    sort_order INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Constraints
    CONSTRAINT marketplace_attributes_type_check CHECK (type IN ('text', 'number', 'boolean', 'select', 'multi_select')),
    CONSTRAINT marketplace_attributes_name_not_empty CHECK (LENGTH(TRIM(name)) > 0),
    CONSTRAINT marketplace_attributes_slug_not_empty CHECK (LENGTH(TRIM(slug)) > 0)
);

-- Valores predefinidos para atributos tipo select y multi_select
CREATE TABLE marketplace_attribute_values (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    attribute_id UUID NOT NULL REFERENCES marketplace_attributes(id) ON DELETE CASCADE,
    value VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    sort_order INTEGER DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Constraints
    CONSTRAINT marketplace_attribute_values_value_not_empty CHECK (LENGTH(TRIM(value)) > 0),
    CONSTRAINT marketplace_attribute_values_slug_not_empty CHECK (LENGTH(TRIM(slug)) > 0),
    UNIQUE(attribute_id, slug)
);

-- Relación entre categorías marketplace y atributos
CREATE TABLE marketplace_category_attributes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    category_id UUID NOT NULL REFERENCES marketplace_categories(id) ON DELETE CASCADE,
    attribute_id UUID NOT NULL REFERENCES marketplace_attributes(id) ON DELETE CASCADE,
    is_required BOOLEAN DEFAULT FALSE,
    sort_order INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(category_id, attribute_id)
);

-- Mapeo de categorías tenant a categorías marketplace
CREATE TABLE tenant_category_mappings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE, -- categoría tenant existente
    marketplace_category_id UUID NOT NULL REFERENCES marketplace_categories(id) ON DELETE CASCADE,
    custom_name VARCHAR(255), -- Nombre personalizado que ve el tenant
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(tenant_id, category_id),
    UNIQUE(tenant_id, marketplace_category_id)
);

-- Extensiones de atributos por tenant
CREATE TABLE tenant_attribute_extensions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    marketplace_attribute_id UUID NOT NULL REFERENCES marketplace_attributes(id) ON DELETE CASCADE,
    custom_name VARCHAR(255), -- Nombre personalizado del atributo
    additional_values JSONB DEFAULT '[]', -- Valores adicionales específicos del tenant
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(tenant_id, marketplace_attribute_id)
);

-- Atributos completamente custom por tenant
CREATE TABLE tenant_custom_attributes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
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

-- NUEVA TABLA: variant_marketplace_attributes (CORRIGIENDO EL DISEÑO)
-- JUSTIFICACIÓN: Los atributos van en variantes, no en productos directamente
CREATE TABLE variant_marketplace_attributes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    variant_id UUID NOT NULL REFERENCES product_variants(id) ON DELETE CASCADE,
    marketplace_attribute_id UUID REFERENCES marketplace_attributes(id) ON DELETE CASCADE,
    tenant_custom_attribute_id UUID REFERENCES tenant_custom_attributes(id) ON DELETE CASCADE,
    value TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Solo uno de los dos tipos de atributo puede estar presente
    CONSTRAINT variant_marketplace_attributes_single_attribute CHECK (
        (marketplace_attribute_id IS NOT NULL AND tenant_custom_attribute_id IS NULL) OR
        (marketplace_attribute_id IS NULL AND tenant_custom_attribute_id IS NOT NULL)
    ),
    -- Un variant no puede tener el mismo atributo duplicado
    UNIQUE(variant_id, marketplace_attribute_id),
    UNIQUE(variant_id, tenant_custom_attribute_id)
);

-- Índices para performance
CREATE INDEX IF NOT EXISTS idx_marketplace_attributes_type ON marketplace_attributes(type);
CREATE INDEX IF NOT EXISTS idx_marketplace_attributes_filterable ON marketplace_attributes(is_filterable);
CREATE INDEX IF NOT EXISTS idx_marketplace_attributes_searchable ON marketplace_attributes(is_searchable);
CREATE INDEX IF NOT EXISTS idx_marketplace_attributes_slug ON marketplace_attributes(slug);

CREATE INDEX IF NOT EXISTS idx_marketplace_attribute_values_attribute_id ON marketplace_attribute_values(attribute_id);
CREATE INDEX IF NOT EXISTS idx_marketplace_attribute_values_active ON marketplace_attribute_values(is_active);

CREATE INDEX IF NOT EXISTS idx_marketplace_category_attributes_category ON marketplace_category_attributes(category_id);
CREATE INDEX IF NOT EXISTS idx_marketplace_category_attributes_attribute ON marketplace_category_attributes(attribute_id);

CREATE INDEX IF NOT EXISTS idx_tenant_category_mappings_tenant ON tenant_category_mappings(tenant_id);
CREATE INDEX IF NOT EXISTS idx_tenant_category_mappings_category ON tenant_category_mappings(category_id);
CREATE INDEX IF NOT EXISTS idx_tenant_category_mappings_marketplace ON tenant_category_mappings(marketplace_category_id);

CREATE INDEX IF NOT EXISTS idx_tenant_attribute_extensions_tenant ON tenant_attribute_extensions(tenant_id);
CREATE INDEX IF NOT EXISTS idx_tenant_attribute_extensions_marketplace ON tenant_attribute_extensions(marketplace_attribute_id);

CREATE INDEX IF NOT EXISTS idx_tenant_custom_attributes_tenant ON tenant_custom_attributes(tenant_id);
CREATE INDEX IF NOT EXISTS idx_tenant_custom_attributes_category ON tenant_custom_attributes(marketplace_category_id);

CREATE INDEX IF NOT EXISTS idx_variant_marketplace_attributes_variant ON variant_marketplace_attributes(variant_id);
CREATE INDEX IF NOT EXISTS idx_variant_marketplace_attributes_marketplace ON variant_marketplace_attributes(marketplace_attribute_id);
CREATE INDEX IF NOT EXISTS idx_variant_marketplace_attributes_custom ON variant_marketplace_attributes(tenant_custom_attribute_id);
CREATE INDEX IF NOT EXISTS idx_variant_marketplace_attributes_value ON variant_marketplace_attributes(value);

-- Comentarios para documentación
COMMENT ON TABLE marketplace_attributes IS 'Atributos globales del marketplace para filtros consistentes';
COMMENT ON TABLE marketplace_attribute_values IS 'Valores predefinidos para atributos tipo select/multi_select';
COMMENT ON TABLE marketplace_category_attributes IS 'Relación entre categorías marketplace y sus atributos aplicables';
COMMENT ON TABLE tenant_category_mappings IS 'Mapeo de categorías tenant existentes a categorías marketplace';
COMMENT ON TABLE tenant_attribute_extensions IS 'Extensiones de atributos marketplace por tenant';
COMMENT ON TABLE tenant_custom_attributes IS 'Atributos completamente personalizados por tenant';
COMMENT ON TABLE variant_marketplace_attributes IS 'Atributos marketplace asignados a variantes específicas';

-- Triggers para updated_at
CREATE OR REPLACE FUNCTION update_marketplace_attributes_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_marketplace_attributes_updated_at
    BEFORE UPDATE ON marketplace_attributes
    FOR EACH ROW
    EXECUTE FUNCTION update_marketplace_attributes_updated_at();

CREATE TRIGGER trigger_tenant_category_mappings_updated_at
    BEFORE UPDATE ON tenant_category_mappings
    FOR EACH ROW
    EXECUTE FUNCTION update_marketplace_attributes_updated_at();

CREATE TRIGGER trigger_tenant_attribute_extensions_updated_at
    BEFORE UPDATE ON tenant_attribute_extensions
    FOR EACH ROW
    EXECUTE FUNCTION update_marketplace_attributes_updated_at();

CREATE TRIGGER trigger_tenant_custom_attributes_updated_at
    BEFORE UPDATE ON tenant_custom_attributes
    FOR EACH ROW
    EXECUTE FUNCTION update_marketplace_attributes_updated_at();

CREATE TRIGGER trigger_variant_marketplace_attributes_updated_at
    BEFORE UPDATE ON variant_marketplace_attributes
    FOR EACH ROW
    EXECUTE FUNCTION update_marketplace_attributes_updated_at(); 