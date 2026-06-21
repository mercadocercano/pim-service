-- Migration 038: Make EAN nullable in global_products
-- Products synced from webdata-service often lack EAN codes.
-- EAN uniqueness is preserved for non-null, non-empty values.

ALTER TABLE global_products ALTER COLUMN ean DROP NOT NULL;

ALTER TABLE global_products DROP CONSTRAINT IF EXISTS global_products_ean_key;

CREATE UNIQUE INDEX IF NOT EXISTS global_products_ean_unique
    ON global_products(ean)
    WHERE ean IS NOT NULL AND ean != '';
