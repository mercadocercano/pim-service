-- Migration: 007_create_category_attributes_table.sql
-- Description: Crear tabla category_attributes para relacionar categorías con atributos
-- Date: 2025-06-01

DROP TABLE IF EXISTS category_attributes;

CREATE TABLE category_attributes (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    tenant_id VARCHAR(36) NOT NULL,
    category_id VARCHAR(36) NOT NULL,
    attribute_id VARCHAR(36) NOT NULL,
    allowed_values TEXT[] DEFAULT '{}',
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Constraints
    CONSTRAINT category_attributes_status_check CHECK (status IN ('active', 'inactive', 'deleted')),
    CONSTRAINT unique_category_attribute_tenant UNIQUE (tenant_id, category_id, attribute_id)
);

-- Índices para optimizar las consultas
CREATE INDEX IF NOT EXISTS idx_category_attributes_tenant_id ON category_attributes(tenant_id);
CREATE INDEX IF NOT EXISTS idx_category_attributes_category_id ON category_attributes(category_id);
CREATE INDEX IF NOT EXISTS idx_category_attributes_attribute_id ON category_attributes(attribute_id);
CREATE INDEX IF NOT EXISTS idx_category_attributes_tenant_category ON category_attributes(tenant_id, category_id);
CREATE INDEX IF NOT EXISTS idx_category_attributes_tenant_attribute ON category_attributes(tenant_id, attribute_id);
CREATE INDEX IF NOT EXISTS idx_category_attributes_status ON category_attributes(status);

-- Comentarios para documentación
COMMENT ON TABLE category_attributes IS 'Tabla para relacionar categorías con atributos en el PIM';
COMMENT ON COLUMN category_attributes.id IS 'Identificador único de la relación';
COMMENT ON COLUMN category_attributes.tenant_id IS 'ID del tenant propietario de la relación';
COMMENT ON COLUMN category_attributes.category_id IS 'ID de la categoría';
COMMENT ON COLUMN category_attributes.attribute_id IS 'ID del atributo';
COMMENT ON COLUMN category_attributes.allowed_values IS 'Valores permitidos específicos para esta relación';
COMMENT ON COLUMN category_attributes.status IS 'Estado de la relación: active, inactive, deleted';
COMMENT ON COLUMN category_attributes.created_at IS 'Fecha y hora de creación';
COMMENT ON COLUMN category_attributes.updated_at IS 'Fecha y hora de última actualización'; 