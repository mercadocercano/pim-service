-- Seeder 017: Marketplace Attributes para Pinturería
-- PROPÓSITO: Atributos de Presentación, Acabado y Tipo de Pintura
-- IDEMPOTENTE: Usa ON CONFLICT DO NOTHING

-- ==============================================
-- PRESENTACIÓN PINTURA
-- ==============================================
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('fa1e8f2a-0005-0000-0000-000000000002', 'Presentación Pintura', 'presentacion-pintura', 'select', true, true, false, 50)
ON CONFLICT (slug) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('fa1e8f2a-0005-0000-0000-000000000002', '0.25L', '0-25l', 1),
('fa1e8f2a-0005-0000-0000-000000000002', '0.5L', '0-5l', 2),
('fa1e8f2a-0005-0000-0000-000000000002', '1L', '1l', 3),
('fa1e8f2a-0005-0000-0000-000000000002', '4L', '4l', 4),
('fa1e8f2a-0005-0000-0000-000000000002', '10L', '10l', 5),
('fa1e8f2a-0005-0000-0000-000000000002', '20L', '20l', 6)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- ACABADO
-- ==============================================
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f23', 'Acabado', 'acabado-pintura', 'select', true, false, false, 51)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f23', 'Mate', 'mate', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f23', 'Satinado', 'satinado', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f23', 'Semi-brillante', 'semi-brillante', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f23', 'Brillante', 'brillante', 4)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- TIPO PINTURA
-- ==============================================
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f24', 'Tipo Pintura', 'tipo-pintura', 'select', true, true, false, 52)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f24', 'Látex Interior', 'latex-interior', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f24', 'Látex Exterior', 'latex-exterior', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f24', 'Esmalte Sintético', 'esmalte-sintetico', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f24', 'Enduido', 'enduido', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f24', 'Membrana Líquida', 'membrana-liquida', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f24', 'Impermeabilizante', 'impermeabilizante', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f24', 'Fijador/Sellador', 'fijador-sellador', 7),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f24', 'Barniz', 'barniz', 8)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- RELACIONES: Vincular a categoría Ferretería y Construcción (que incluye pinturas)
-- ==============================================
DO $$
DECLARE
  cat_id UUID;
BEGIN
  SELECT id INTO cat_id FROM marketplace_categories WHERE slug = 'ferreteria-construccion' LIMIT 1;
  IF cat_id IS NULL THEN
    SELECT id INTO cat_id FROM marketplace_categories WHERE name ILIKE '%Ferretería%' AND level = 0 LIMIT 1;
  END IF;
  IF cat_id IS NOT NULL THEN
    INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
      (cat_id, 'fa1e8f2a-0005-0000-0000-000000000002', false, 10),
      (cat_id, 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f23', false, 11),
      (cat_id, 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f24', false, 12)
    ON CONFLICT (category_id, attribute_id) DO NOTHING;
  END IF;

  SELECT id INTO cat_id FROM marketplace_categories WHERE slug = 'pinturas' LIMIT 1;
  IF cat_id IS NOT NULL THEN
    INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
      (cat_id, 'fa1e8f2a-0005-0000-0000-000000000002', false, 10),
      (cat_id, 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f23', false, 11),
      (cat_id, 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f24', false, 12)
    ON CONFLICT (category_id, attribute_id) DO NOTHING;
  END IF;
END $$;
