-- Migration 014: Add detailed food subcategories for better product categorization
-- Date: 2025-11-21
-- Purpose: Create granular subcategories under "alimentacion" for AI Gateway product curation

-- Add level 3 subcategories under "Snacks y Dulces"
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order, is_active) 
VALUES
(
    'a1b2c3d4-e5f6-4a5b-8c9d-1e2f3a4b5c6d',
    'Snacks Salados',
    'snacks-salados',
    'Papas fritas, palitos, nachos, productos salados para picar',
    (SELECT id FROM marketplace_categories WHERE slug = 'snacks-dulces'),
    3,
    1,
    TRUE
),
(
    'b2c3d4e5-f6a7-4b5c-8d9e-2f3a4b5c6d7e',
    'Galletas',
    'galletas',
    'Galletas dulces, galletitas, cookies, obleas, crackers',
    (SELECT id FROM marketplace_categories WHERE slug = 'snacks-dulces'),
    3,
    2,
    TRUE
),
(
    'c3d4e5f6-a7b8-4c5d-8e9f-3a4b5c6d7e8f',
    'Golosinas',
    'golosinas',
    'Chocolates, caramelos, alfajores, chicles, gomitas, bombones',
    (SELECT id FROM marketplace_categories WHERE slug = 'snacks-dulces'),
    3,
    3,
    TRUE
);

-- Add level 3 subcategory under "Pastas y Cereales"
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order, is_active)
VALUES
(
    'd4e5f6a7-b8c9-4d5e-8f9a-4b5c6d7e8f9a',
    'Cereales de Desayuno',
    'cereales-desayuno',
    'Cereales, granolas, barras de cereal, avena',
    (SELECT id FROM marketplace_categories WHERE slug = 'pastas-cereales'),
    3,
    1,
    TRUE
);

-- Add level 2 subcategory under "Alimentos Frescos"
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order, is_active)
VALUES
(
    'e5f6a7b8-c9d0-4e5f-8a9b-5c6d7e8f9a0b',
    'Panificados',
    'panificados',
    'Pan, facturas, productos de panadería, tostadas',
    (SELECT id FROM marketplace_categories WHERE slug = 'alimentos-frescos'),
    2,
    4,
    TRUE
);

-- Create indices for better query performance
CREATE INDEX IF NOT EXISTS idx_marketplace_categories_slug_active 
ON marketplace_categories(slug, is_active) 
WHERE is_active = TRUE;

-- Log migration completion
COMMENT ON TABLE marketplace_categories IS 'Updated 2025-11-21: Added 5 detailed food subcategories for AI curation (snacks-salados, galletas, golosinas, cereales-desayuno, panificados)';

