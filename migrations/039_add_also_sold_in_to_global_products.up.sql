-- Migración idempotente: agrega campo also_sold_in a global_products
-- Usado para cross-rubro sharing en el quickstart computado
ALTER TABLE global_products
  ADD COLUMN IF NOT EXISTS also_sold_in JSONB DEFAULT '[]';

CREATE INDEX IF NOT EXISTS idx_global_products_also_sold_in
  ON global_products USING GIN (also_sold_in);

