-- Migration: add unique constraint (name, business_type) to global_products
-- Cycle: cycle-005-global-products-by-business-type
-- Date: 2026-04-24
-- Reason: seeds use ON CONFLICT (name, business_type) DO NOTHING for idempotency

-- UP
ALTER TABLE global_products
    ADD CONSTRAINT uq_global_products_name_business_type
    UNIQUE (name, business_type);

-- DOWN
-- ALTER TABLE global_products DROP CONSTRAINT uq_global_products_name_business_type;
