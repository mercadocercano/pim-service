-- Eliminar trigger y función
DROP TRIGGER IF EXISTS trigger_products_updated_at ON products;
DROP FUNCTION IF EXISTS update_products_updated_at();

-- Eliminar índices (se eliminan automáticamente con la tabla, pero por claridad)
DROP INDEX IF EXISTS idx_products_tenant_brand;
DROP INDEX IF EXISTS idx_products_tenant_category;
DROP INDEX IF EXISTS idx_products_tenant_status;
DROP INDEX IF EXISTS idx_products_tenant_sku_unique;
DROP INDEX IF EXISTS idx_products_tenant_name_unique;
DROP INDEX IF EXISTS idx_products_name;
DROP INDEX IF EXISTS idx_products_created_at;
DROP INDEX IF EXISTS idx_products_brand_id;
DROP INDEX IF EXISTS idx_products_category_id;
DROP INDEX IF EXISTS idx_products_status;
DROP INDEX IF EXISTS idx_products_tenant_id;

-- Eliminar tabla
DROP TABLE IF EXISTS products; 