-- Migration 044: Create business_type_product_templates table
-- Output computado de RefreshTemplateProductsUseCase
-- Idempotente: CREATE TABLE IF NOT EXISTS, CREATE INDEX IF NOT EXISTS

CREATE TABLE IF NOT EXISTS business_type_product_templates (
    id                          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_type_template_id   UUID NOT NULL,
    suggested_products          JSONB NOT NULL DEFAULT '[]'::jsonb,
    priority_brands             JSONB NOT NULL DEFAULT '[]'::jsonb,
    max_products_per_category   INTEGER NOT NULL DEFAULT 30,
    created_at                  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at                  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_btpt_template
        FOREIGN KEY (business_type_template_id)
        REFERENCES business_type_templates(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_btpt_template_id
        UNIQUE (business_type_template_id)
);

CREATE INDEX IF NOT EXISTS idx_btpt_template_id
    ON business_type_product_templates (business_type_template_id);

CREATE INDEX IF NOT EXISTS idx_btpt_updated_at
    ON business_type_product_templates (updated_at DESC);

COMMENT ON TABLE business_type_product_templates IS
    'Output computado de RefreshTemplateProductsUseCase. No editar manualmente.';
