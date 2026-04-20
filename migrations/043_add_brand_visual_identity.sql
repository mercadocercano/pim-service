-- Migration 043: Add brand visual identity columns
-- CONTEXTO: cycle-002-global-brands-colors / ADR-001
ALTER TABLE marketplace_brands
    ADD COLUMN IF NOT EXISTS background_color VARCHAR(7)   NULL,
    ADD COLUMN IF NOT EXISTS text_color       VARCHAR(7)   NULL,
    ADD COLUMN IF NOT EXISTS typography       VARCHAR(100) NULL;

ALTER TABLE marketplace_brands
    DROP CONSTRAINT IF EXISTS marketplace_brands_background_color_format;
ALTER TABLE marketplace_brands ADD CONSTRAINT marketplace_brands_background_color_format
    CHECK (background_color IS NULL OR background_color ~ '^#[0-9A-Fa-f]{6}$');

ALTER TABLE marketplace_brands
    DROP CONSTRAINT IF EXISTS marketplace_brands_text_color_format;
ALTER TABLE marketplace_brands ADD CONSTRAINT marketplace_brands_text_color_format
    CHECK (text_color IS NULL OR text_color ~ '^#[0-9A-Fa-f]{6}$');

ALTER TABLE marketplace_brands
    DROP CONSTRAINT IF EXISTS marketplace_brands_typography_format;
ALTER TABLE marketplace_brands ADD CONSTRAINT marketplace_brands_typography_format
    CHECK (typography IS NULL OR (LENGTH(TRIM(typography)) BETWEEN 1 AND 100));

COMMENT ON COLUMN marketplace_brands.background_color IS 'Hex #RRGGBB color primario de fondo. NULL = fallback design system.';
COMMENT ON COLUMN marketplace_brands.text_color       IS 'Hex #RRGGBB color de texto sobre background. NULL = fallback.';
COMMENT ON COLUMN marketplace_brands.typography       IS 'Familia Google Fonts (ej. Oswald). NULL = hereda theme.';
