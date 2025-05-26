-- Crear tabla de variantes de productos
CREATE TABLE IF NOT EXISTS product_variants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    sku VARCHAR(100) UNIQUE,
    status VARCHAR(50) NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'discontinued', 'deleted')),
    is_default BOOLEAN NOT NULL DEFAULT false,
    sort_order INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    -- Constraints
    CONSTRAINT unique_product_variant_name_per_tenant UNIQUE (tenant_id, product_id, name),
    CONSTRAINT unique_default_variant_per_product EXCLUDE (product_id WITH =) WHERE (is_default = true)
);

-- Crear tabla de atributos de variantes
CREATE TABLE IF NOT EXISTS variant_attributes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    variant_id UUID NOT NULL REFERENCES product_variants(id) ON DELETE CASCADE,
    attribute_name VARCHAR(100) NOT NULL,
    attribute_value VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    -- Constraints
    CONSTRAINT unique_variant_attribute UNIQUE (variant_id, attribute_name)
);

-- Índices para product_variants
CREATE INDEX IF NOT EXISTS idx_product_variants_tenant_id ON product_variants(tenant_id);
CREATE INDEX IF NOT EXISTS idx_product_variants_product_id ON product_variants(product_id);
CREATE INDEX IF NOT EXISTS idx_product_variants_sku ON product_variants(sku) WHERE sku IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_product_variants_status ON product_variants(status);
CREATE INDEX IF NOT EXISTS idx_product_variants_is_default ON product_variants(is_default) WHERE is_default = true;
CREATE INDEX IF NOT EXISTS idx_product_variants_sort_order ON product_variants(sort_order);
CREATE INDEX IF NOT EXISTS idx_product_variants_created_at ON product_variants(created_at);

-- Índices para variant_attributes
CREATE INDEX IF NOT EXISTS idx_variant_attributes_tenant_id ON variant_attributes(tenant_id);
CREATE INDEX IF NOT EXISTS idx_variant_attributes_variant_id ON variant_attributes(variant_id);
CREATE INDEX IF NOT EXISTS idx_variant_attributes_name ON variant_attributes(attribute_name);
CREATE INDEX IF NOT EXISTS idx_variant_attributes_value ON variant_attributes(attribute_value);

-- Trigger técnico para actualizar updated_at en product_variants
CREATE OR REPLACE FUNCTION update_product_variants_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_product_variants_updated_at
    BEFORE UPDATE ON product_variants
    FOR EACH ROW
    EXECUTE FUNCTION update_product_variants_updated_at();

-- Trigger técnico para actualizar updated_at en variant_attributes
CREATE OR REPLACE FUNCTION update_variant_attributes_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_variant_attributes_updated_at
    BEFORE UPDATE ON variant_attributes
    FOR EACH ROW
    EXECUTE FUNCTION update_variant_attributes_updated_at();

-- NOTA: La lógica de creación de variante por defecto se maneja desde el código
-- en el agregado Product, siguiendo principios de DDD y arquitectura hexagonal 