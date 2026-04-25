-- Migración idempotente: agrega campo also_sold_in a global_products
-- Usado para cross-rubro sharing en el quickstart computado
ALTER TABLE global_products
  ADD COLUMN IF NOT EXISTS also_sold_in JSONB DEFAULT '[]';

CREATE INDEX IF NOT EXISTS idx_global_products_also_sold_in
  ON global_products USING GIN (also_sold_in);

-- Registrar en schema_migrations (tabla tiene columna 'filename' no 'version')
INSERT INTO schema_migrations (filename, applied_at)
VALUES ('20260425000004_add_also_sold_in_to_global_products.sql', NOW())
ON CONFLICT (filename) DO NOTHING;
