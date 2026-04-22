-- Migration to update business_type_templates data format
-- Updates categories, attributes, and products to include all required fields

-- First, let's create a temporary function to help with the migration
CREATE OR REPLACE FUNCTION migrate_template_categories(categories_json jsonb)
RETURNS jsonb AS $$
DECLARE
    result jsonb := '[]'::jsonb;
    category jsonb;
    marketplace_cat record;
BEGIN
    -- Loop through each category in the array
    FOR category IN SELECT * FROM jsonb_array_elements(categories_json)
    LOOP
        -- Look up the category in marketplace_categories
        SELECT id, name, slug, description, parent_id, level 
        INTO marketplace_cat
        FROM marketplace_categories 
        WHERE id = (category->>'id')::uuid;
        
        IF FOUND THEN
            -- Build the new category format
            result := result || jsonb_build_object(
                'id', marketplace_cat.id::text,
                'name', marketplace_cat.name,
                'slug', marketplace_cat.slug,
                'description', COALESCE(marketplace_cat.description, ''),
                'parent_id', COALESCE(marketplace_cat.parent_id::text, ''),
                'level', marketplace_cat.level
            );
        ELSE
            -- If category not found, try to use slug as name
            result := result || jsonb_build_object(
                'id', category->>'id',
                'name', REPLACE(category->>'slug', '-', ' '),
                'slug', category->>'slug',
                'description', '',
                'parent_id', '',
                'level', 0
            );
        END IF;
    END LOOP;
    
    RETURN result;
END;
$$ LANGUAGE plpgsql;

-- Update all templates with the new category format
UPDATE business_type_templates 
SET categories = migrate_template_categories(categories)
WHERE categories IS NOT NULL 
  AND jsonb_typeof(categories) = 'array'
  AND categories::text LIKE '%"slug"%';

-- Clean up the temporary function
DROP FUNCTION IF EXISTS migrate_template_categories(jsonb);

-- Add any missing fields to templates that might not have them
UPDATE business_type_templates 
SET 
    is_active = COALESCE(is_active, true),
    is_default = COALESCE(is_default, false)
WHERE is_active IS NULL OR is_default IS NULL;