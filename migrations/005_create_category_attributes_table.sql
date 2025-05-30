-- Migración para crear la tabla category_attributes
-- +goose Up
CREATE TABLE IF NOT EXISTS category_attributes (
    id VARCHAR(36) PRIMARY KEY,
    tenant_id VARCHAR(36) NOT NULL,
    category_id VARCHAR(36) NOT NULL,
    attribute_id VARCHAR(36) NOT NULL,
    allowed_values TEXT[] DEFAULT '{}',
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Índices para mejorar el rendimiento
    CONSTRAINT unique_category_attribute_tenant UNIQUE (tenant_id, category_id, attribute_id)
);

-- Crear índices para optimizar las consultas
CREATE INDEX IF NOT EXISTS idx_category_attributes_tenant_id ON category_attributes(tenant_id);
CREATE INDEX IF NOT EXISTS idx_category_attributes_category_id ON category_attributes(category_id);
CREATE INDEX IF NOT EXISTS idx_category_attributes_attribute_id ON category_attributes(attribute_id);
CREATE INDEX IF NOT EXISTS idx_category_attributes_tenant_category ON category_attributes(tenant_id, category_id);

-- +goose Down
DROP TABLE IF EXISTS category_attributes; 