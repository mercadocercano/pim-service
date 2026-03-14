-- Seeder 019: Marketplace Attributes adicionales para Deportes y Zapatería
-- PROPÓSITO: Atributo de Disciplina Deportiva (Talle/Calzado/Material ya existen en seed 014)
-- IDEMPOTENTE: Usa ON CONFLICT DO NOTHING

-- ==============================================
-- DISCIPLINA DEPORTIVA
-- ==============================================
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f28', 'Disciplina', 'disciplina-deportiva', 'select', true, true, false, 70)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f28', 'Fútbol', 'futbol', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f28', 'Running', 'running', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f28', 'Fitness/Gym', 'fitness', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f28', 'Natación', 'natacion', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f28', 'Tenis/Padel', 'tenis-padel', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f28', 'Basketball', 'basketball', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f28', 'Ciclismo', 'ciclismo', 7),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f28', 'Outdoor/Trekking', 'outdoor', 8)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- TIPO CALZADO (complementa talle-calzado existente en seed 014)
-- ==============================================
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f29', 'Tipo Calzado', 'tipo-calzado', 'select', true, false, false, 71)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f29', 'Urbano', 'urbano', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f29', 'Deportivo', 'deportivo', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f29', 'Formal', 'formal', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f29', 'Sandalia', 'sandalia', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f29', 'Bota', 'bota', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f29', 'Ojotas', 'ojotas', 6)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- RELACIONES: Vincular Disciplina a categoría Deportes y Fitness
-- ==============================================
DO $$
DECLARE
  cat_id UUID;
BEGIN
  SELECT id INTO cat_id FROM marketplace_categories WHERE slug = 'deportes-fitness' LIMIT 1;
  IF cat_id IS NULL THEN
    SELECT id INTO cat_id FROM marketplace_categories WHERE name ILIKE '%Deportes%' AND level = 0 LIMIT 1;
  END IF;
  IF cat_id IS NOT NULL THEN
    INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
      (cat_id, 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f28', false, 8),
      (cat_id, 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f29', false, 9)
    ON CONFLICT (category_id, attribute_id) DO NOTHING;
  END IF;

  SELECT id INTO cat_id FROM marketplace_categories WHERE slug = 'moda-accesorios' LIMIT 1;
  IF cat_id IS NULL THEN
    SELECT id INTO cat_id FROM marketplace_categories WHERE name ILIKE '%Moda%' AND level = 0 LIMIT 1;
  END IF;
  IF cat_id IS NOT NULL THEN
    INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
      (cat_id, 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f29', false, 7)
    ON CONFLICT (category_id, attribute_id) DO NOTHING;
  END IF;
END $$;
