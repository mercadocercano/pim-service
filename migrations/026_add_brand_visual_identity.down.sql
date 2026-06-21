-- Down Migration 043: Remove brand visual identity columns
ALTER TABLE marketplace_brands
    DROP CONSTRAINT IF EXISTS marketplace_brands_background_color_format,
    DROP CONSTRAINT IF EXISTS marketplace_brands_text_color_format,
    DROP CONSTRAINT IF EXISTS marketplace_brands_typography_format;

ALTER TABLE marketplace_brands
    DROP COLUMN IF EXISTS background_color,
    DROP COLUMN IF EXISTS text_color,
    DROP COLUMN IF EXISTS typography;
