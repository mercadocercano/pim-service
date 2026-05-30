-- Migration 045: Add color column to brands table
-- Agrega soporte de color hex (#RRGGBB) para personalización del badge de marca en POS
ALTER TABLE brands ADD COLUMN IF NOT EXISTS color VARCHAR(7) NULL;

COMMENT ON COLUMN brands.color IS 'Color hex #RRGGBB opcional para personalizar el badge de la marca en POS. NULL = fallback al design system.';
