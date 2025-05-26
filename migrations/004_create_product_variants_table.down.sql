-- Eliminar triggers y funciones de variant_attributes
DROP TRIGGER IF EXISTS trigger_update_variant_attributes_updated_at ON variant_attributes;
DROP FUNCTION IF EXISTS update_variant_attributes_updated_at();

-- Eliminar triggers y funciones de product_variants
DROP TRIGGER IF EXISTS trigger_update_product_variants_updated_at ON product_variants;
DROP FUNCTION IF EXISTS update_product_variants_updated_at();

-- Eliminar índices de variant_attributes
DROP INDEX IF EXISTS idx_variant_attributes_value;
DROP INDEX IF EXISTS idx_variant_attributes_name;
DROP INDEX IF EXISTS idx_variant_attributes_variant_id;
DROP INDEX IF EXISTS idx_variant_attributes_tenant_id;

-- Eliminar índices de product_variants
DROP INDEX IF EXISTS idx_product_variants_created_at;
DROP INDEX IF EXISTS idx_product_variants_sort_order;
DROP INDEX IF EXISTS idx_product_variants_is_default;
DROP INDEX IF EXISTS idx_product_variants_status;
DROP INDEX IF EXISTS idx_product_variants_sku;
DROP INDEX IF EXISTS idx_product_variants_product_id;
DROP INDEX IF EXISTS idx_product_variants_tenant_id;

-- Eliminar tablas
DROP TABLE IF EXISTS variant_attributes;
DROP TABLE IF EXISTS product_variants; 