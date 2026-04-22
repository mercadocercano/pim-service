-- Migration: Add is_active column to marketplace_brands table
-- Description: Permite desactivar marcas globales sin eliminarlas
-- Date: 2024-12-18

-- Add is_active column with default true
ALTER TABLE marketplace_brands
ADD COLUMN IF NOT EXISTS is_active boolean NOT NULL DEFAULT true;

-- Add comment for documentation
COMMENT ON COLUMN marketplace_brands.is_active IS 'Indica si la marca está activa y disponible para uso';

-- Create index for performance on active brands queries
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_is_active ON marketplace_brands(is_active);

-- Update existing records to be active (redundant but explicit)
UPDATE marketplace_brands SET is_active = true WHERE is_active IS NULL; 