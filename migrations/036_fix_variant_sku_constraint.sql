-- Migración 036: Corregir constraint de SKU en product_variants
-- Problema: SKU es UNIQUE global, debe ser UNIQUE por tenant
-- Fecha: 2026-02-04

BEGIN;

-- 1. Eliminar constraint incorrecta (UNIQUE global)
ALTER TABLE product_variants
DROP CONSTRAINT IF EXISTS product_variants_sku_key;

-- 2. Crear constraint correcta (UNIQUE por tenant)
ALTER TABLE product_variants
ADD CONSTRAINT product_variants_tenant_sku_key
UNIQUE (tenant_id, sku);

-- 3. Actualizar índice para optimizar búsquedas por tenant+sku
DROP INDEX IF EXISTS idx_product_variants_sku;
CREATE INDEX IF NOT EXISTS idx_product_variants_tenant_sku 
ON product_variants(tenant_id, sku) WHERE sku IS NOT NULL;

COMMIT;
