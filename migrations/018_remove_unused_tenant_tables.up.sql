-- Migration 018: Remove unused tenant_* tables
-- PROPÓSITO: Limpiar arquitectura eliminando tablas tenant_* que no se usan
-- MANTENER: variant_attributes (en uso), tenant_brand_mappings, tenant_category_mappings

-- 1. Eliminar tabla variant_marketplace_attributes primero (tiene FK a tenant_custom_attributes)
DROP TRIGGER IF EXISTS trigger_variant_marketplace_attributes_updated_at ON variant_marketplace_attributes;

-- Eliminar índices de variant_marketplace_attributes
DROP INDEX IF EXISTS idx_variant_marketplace_attributes_value;
DROP INDEX IF EXISTS idx_variant_marketplace_attributes_custom;
DROP INDEX IF EXISTS idx_variant_marketplace_attributes_marketplace;
DROP INDEX IF EXISTS idx_variant_marketplace_attributes_variant;

-- Eliminar la tabla
DROP TABLE IF EXISTS variant_marketplace_attributes;

-- 2. Ahora eliminar tenant_custom_attributes y dependencias
DROP TRIGGER IF EXISTS trigger_tenant_custom_attributes_updated_at ON tenant_custom_attributes;

-- Eliminar índices de tenant_custom_attributes
DROP INDEX IF EXISTS idx_tenant_custom_attributes_slug;
DROP INDEX IF EXISTS idx_tenant_custom_attributes_category;
DROP INDEX IF EXISTS idx_tenant_custom_attributes_tenant;
DROP INDEX IF EXISTS idx_tenant_custom_attributes_deleted_at;

-- Eliminar la tabla
DROP TABLE IF EXISTS tenant_custom_attributes;

-- 3. Eliminar tabla tenant_business_type_setup
DROP TRIGGER IF EXISTS update_tenant_business_type_setup_updated_at ON tenant_business_type_setup;

-- Eliminar índices de tenant_business_type_setup
DROP INDEX IF EXISTS idx_tenant_business_type_setup_business_type_id;
DROP INDEX IF EXISTS idx_tenant_business_type_setup_tenant_id;

-- Eliminar la tabla
DROP TABLE IF EXISTS tenant_business_type_setup;

-- 4. Eliminar tabla tenant_attribute_extensions
DROP TRIGGER IF EXISTS trigger_tenant_attribute_extensions_updated_at ON tenant_attribute_extensions;

-- Eliminar índices de tenant_attribute_extensions
DROP INDEX IF EXISTS idx_tenant_attribute_extensions_marketplace;
DROP INDEX IF EXISTS idx_tenant_attribute_extensions_tenant;

-- Eliminar la tabla
DROP TABLE IF EXISTS tenant_attribute_extensions;

-- 5. Eliminar función update_tenant_tables_updated_at si ya no se usa
DROP FUNCTION IF EXISTS update_tenant_tables_updated_at() CASCADE;

-- Comentarios para documentación
COMMENT ON SCHEMA public IS 'Limpieza completada: eliminadas tablas tenant_* no utilizadas. Mantenidas: variant_attributes, tenant_brand_mappings, tenant_category_mappings';