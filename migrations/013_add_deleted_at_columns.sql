-- Migration 013: Add deleted_at columns for soft delete functionality
-- PROPÓSITO: Agregar soporte para soft delete en tablas marketplace
-- BENEFICIO: Mantener historial de datos eliminados para auditoría

-- Agregar deleted_at a marketplace_categories
ALTER TABLE marketplace_categories 
ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL;

-- Agregar deleted_at a tenant_category_mappings
ALTER TABLE tenant_category_mappings 
ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL;

-- Agregar deleted_at a tenant_custom_attributes
ALTER TABLE tenant_custom_attributes 
ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL;

-- Agregar índices para performance en consultas con deleted_at
CREATE INDEX IF NOT EXISTS idx_marketplace_categories_deleted_at ON marketplace_categories(deleted_at);
CREATE INDEX IF NOT EXISTS idx_tenant_category_mappings_deleted_at ON tenant_category_mappings(deleted_at);
CREATE INDEX IF NOT EXISTS idx_tenant_custom_attributes_deleted_at ON tenant_custom_attributes(deleted_at);

-- Comentarios para documentación
COMMENT ON COLUMN marketplace_categories.deleted_at IS 'Timestamp de eliminación lógica, NULL si no está eliminado';
COMMENT ON COLUMN tenant_category_mappings.deleted_at IS 'Timestamp de eliminación lógica, NULL si no está eliminado';
COMMENT ON COLUMN tenant_custom_attributes.deleted_at IS 'Timestamp de eliminación lógica, NULL si no está eliminado'; 