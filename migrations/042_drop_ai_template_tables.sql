-- Migration: Drop deprecated AI template tables and columns
-- Description: Removes AI template capabilities added in migration 034 (modules batch and template_ai are deprecated)

-- 1. Drop triggers
DROP TRIGGER IF EXISTS update_template_global_products_updated_at ON template_global_products;

-- 2. Drop tables in reverse order of dependencies
DROP TABLE IF EXISTS ai_product_feedback;
DROP TABLE IF EXISTS ai_template_performance_metrics;
DROP TABLE IF EXISTS ai_template_generation_history;
DROP TABLE IF EXISTS template_global_products;

-- 3. Remove AI columns from business_type_templates
ALTER TABLE business_type_templates
DROP COLUMN IF EXISTS generated_by,
DROP COLUMN IF EXISTS ai_generation_params,
DROP COLUMN IF EXISTS performance_metrics,
DROP COLUMN IF EXISTS last_ai_update,
DROP COLUMN IF EXISTS product_selection_rules,
DROP COLUMN IF EXISTS category_distribution,
DROP COLUMN IF EXISTS brand_preferences,
DROP COLUMN IF EXISTS regional_adaptations;

-- 4. Drop indexes
DROP INDEX IF EXISTS idx_business_type_templates_generated_by;
