-- Migration: 020_add_slug_to_categories.sql
-- Description: Agregar columna slug a tabla categories para búsqueda eficiente (HITO 2)
-- Date: 2026-02-01

-- Agregar columna slug si no existe
ALTER TABLE categories ADD COLUMN IF NOT EXISTS slug VARCHAR(255);

-- Crear índice para búsqueda por slug
CREATE INDEX IF NOT EXISTS idx_categories_slug ON categories(slug);
CREATE INDEX IF NOT EXISTS idx_categories_tenant_slug ON categories(tenant_id, slug);

-- Generar slugs para categorías existentes que no tengan
UPDATE categories
SET slug = LOWER(REPLACE(REPLACE(REPLACE(name, ' ', '-'), 'á', 'a'), 'é', 'e'))
WHERE slug IS NULL OR slug = '';

-- Comentarios
COMMENT ON COLUMN categories.slug IS 'Slug único de la categoría para búsquedas y URLs';

