ALTER TABLE global_products ADD COLUMN IF NOT EXISTS gtin VARCHAR(14);
CREATE INDEX IF NOT EXISTS idx_global_products_gtin ON global_products(gtin) WHERE gtin IS NOT NULL;
