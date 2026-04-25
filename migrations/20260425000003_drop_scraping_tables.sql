-- Migration: Drop scraping tables from PIM
-- Scraping responsibility belongs to webdata-service, not PIM.
-- global_product_variations and tenant_global_product_links are kept (planned features).

DROP TABLE IF EXISTS scraping_logs;
DROP TABLE IF EXISTS scraping_targets;
