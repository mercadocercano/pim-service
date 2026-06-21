-- Migration: make ean nullable in global_products
-- Reason: seed products don't have EAN codes; the NOT NULL constraint
-- prevented seeding the catalog with reference data from seeds 041+.
-- PostgreSQL UNIQUE allows multiple NULL values, so data integrity is preserved.
ALTER TABLE global_products ALTER COLUMN ean DROP NOT NULL;
