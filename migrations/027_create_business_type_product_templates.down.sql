-- Migration 044 rollback: Drop business_type_product_templates table

DROP INDEX IF EXISTS idx_btpt_updated_at;
DROP INDEX IF EXISTS idx_btpt_template_id;
DROP TABLE IF EXISTS business_type_product_templates;
