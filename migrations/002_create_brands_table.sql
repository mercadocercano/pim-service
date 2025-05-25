-- Migration: 002_create_brands_table.sql
-- Description: Crear tabla brands para el módulo Brand
-- Date: 2024-01-15

CREATE TABLE IF NOT EXISTS brands (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    logo_url VARCHAR(500),
    website VARCHAR(200),
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Constraints
    CONSTRAINT brands_status_check CHECK (status IN ('active', 'inactive', 'deleted')),
    CONSTRAINT brands_name_not_empty CHECK (LENGTH(TRIM(name)) > 0),
    CONSTRAINT brands_logo_url_format CHECK (logo_url IS NULL OR logo_url ~ '^https?://'),
    CONSTRAINT brands_website_format CHECK (website IS NULL OR website ~ '^https?://'),
    
    -- Unique constraint para evitar marcas duplicadas por tenant
    UNIQUE(tenant_id, name)
);

-- Índices para optimizar consultas
CREATE INDEX IF NOT EXISTS idx_brands_tenant_id ON brands(tenant_id);
CREATE INDEX IF NOT EXISTS idx_brands_status ON brands(status);
CREATE INDEX IF NOT EXISTS idx_brands_tenant_status ON brands(tenant_id, status);
CREATE INDEX IF NOT EXISTS idx_brands_name_search ON brands USING gin(to_tsvector('spanish', name));
CREATE INDEX IF NOT EXISTS idx_brands_created_at ON brands(created_at);

-- Comentarios para documentación
COMMENT ON TABLE brands IS 'Tabla para almacenar marcas de productos en el PIM';
COMMENT ON COLUMN brands.id IS 'Identificador único de la marca';
COMMENT ON COLUMN brands.tenant_id IS 'ID del tenant propietario de la marca';
COMMENT ON COLUMN brands.name IS 'Nombre de la marca (único por tenant)';
COMMENT ON COLUMN brands.description IS 'Descripción de la marca';
COMMENT ON COLUMN brands.logo_url IS 'URL del logo de la marca';
COMMENT ON COLUMN brands.website IS 'Sitio web oficial de la marca';
COMMENT ON COLUMN brands.status IS 'Estado de la marca: active, inactive, deleted';
COMMENT ON COLUMN brands.created_at IS 'Fecha y hora de creación';
COMMENT ON COLUMN brands.updated_at IS 'Fecha y hora de última actualización'; 