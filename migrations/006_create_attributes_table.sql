-- Migration: 006_create_attributes_table.sql
-- Description: Crear tabla attributes para el módulo quickstart
-- Date: 2025-06-01

DROP TABLE IF EXISTS attributes;

CREATE TABLE attributes (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    tenant_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    type VARCHAR(50) NOT NULL DEFAULT 'text', -- text, number, select, boolean, multi-select
    required BOOLEAN DEFAULT FALSE,
    options TEXT[] DEFAULT '{}', -- Para tipos select y multi-select
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Constraints
    CONSTRAINT attributes_status_check CHECK (status IN ('active', 'inactive', 'deleted')),
    CONSTRAINT attributes_type_check CHECK (type IN ('text', 'number', 'select', 'boolean', 'multi-select')),
    CONSTRAINT attributes_name_not_empty CHECK (LENGTH(TRIM(name)) > 0),
    CONSTRAINT unique_attribute_name_tenant UNIQUE (tenant_id, name)
);

-- Índices para optimizar las consultas
CREATE INDEX IF NOT EXISTS idx_attributes_tenant_id ON attributes(tenant_id);
CREATE INDEX IF NOT EXISTS idx_attributes_name ON attributes(name);
CREATE INDEX IF NOT EXISTS idx_attributes_type ON attributes(type);
CREATE INDEX IF NOT EXISTS idx_attributes_status ON attributes(status);
CREATE INDEX IF NOT EXISTS idx_attributes_tenant_status ON attributes(tenant_id, status);
CREATE INDEX IF NOT EXISTS idx_attributes_name_search ON attributes USING gin(to_tsvector('spanish', name));

-- Comentarios para documentación
COMMENT ON TABLE attributes IS 'Tabla para almacenar atributos de productos en el PIM';
COMMENT ON COLUMN attributes.id IS 'Identificador único del atributo';
COMMENT ON COLUMN attributes.tenant_id IS 'ID del tenant propietario del atributo';
COMMENT ON COLUMN attributes.name IS 'Nombre del atributo (único por tenant)';
COMMENT ON COLUMN attributes.description IS 'Descripción del atributo';
COMMENT ON COLUMN attributes.type IS 'Tipo de atributo: text, number, select, boolean, multi-select';
COMMENT ON COLUMN attributes.required IS 'Indica si el atributo es obligatorio';
COMMENT ON COLUMN attributes.options IS 'Opciones disponibles para atributos de tipo select';
COMMENT ON COLUMN attributes.status IS 'Estado del atributo: active, inactive, deleted';
COMMENT ON COLUMN attributes.created_at IS 'Fecha y hora de creación';
COMMENT ON COLUMN attributes.updated_at IS 'Fecha y hora de última actualización'; 