-- Seeder 016: Marketplace Attributes para Alimentos (Almacén / Kiosco)
-- PROPÓSITO: Atributos de Contenido Neto y Tipo Envase para productos alimenticios
-- IDEMPOTENTE: Usa ON CONFLICT DO NOTHING

-- ==============================================
-- CONTENIDO NETO
-- ==============================================
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', 'Contenido Neto', 'contenido-neto', 'select', true, true, false, 40)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '50g', '50g', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '100g', '100g', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '150g', '150g', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '200g', '200g', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '250g', '250g', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '300g', '300g', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '500g', '500g', 7),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '1kg', '1kg', 8),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '2.5kg', '2-5kg', 9),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '5kg', '5kg', 10),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '200ml', '200ml', 11),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '250ml', '250ml', 12),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '330ml', '330ml', 13),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '473ml', '473ml', 14),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '500ml', '500ml', 15),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '750ml', '750ml', 16),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '1L', '1l', 17),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '1.5L', '1-5l', 18),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '2.25L', '2-25l', 19),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', '3L', '3l', 20)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- TIPO ENVASE
-- ==============================================
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f21', 'Tipo Envase', 'tipo-envase', 'select', true, false, false, 41)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f21', 'Botella', 'botella', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f21', 'Lata', 'lata', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f21', 'Tetra Brik', 'tetra-brik', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f21', 'Sachet', 'sachet', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f21', 'Caja', 'caja', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f21', 'Bolsa', 'bolsa', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f21', 'Frasco', 'frasco', 7),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f21', 'Paquete', 'paquete', 8),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f21', 'Bandeja', 'bandeja', 9),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f21', 'Bidón', 'bidon', 10)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- RELACIONES: Vincular a categorías de Alimentación
-- ==============================================
DO $$
DECLARE
  cat_id UUID;
  cat_slugs TEXT[] := ARRAY[
    'alimentos-envasados', 'bebidas-sin-alcohol', 'bebidas-con-alcohol',
    'lacteos-liquidos', 'lacteos-fiambres', 'conservas', 'galletas',
    'golosinas', 'snacks-salados', 'pastas-cereales', 'panificados',
    'cereales-desayuno'
  ];
  s TEXT;
BEGIN
  FOREACH s IN ARRAY cat_slugs LOOP
    SELECT id INTO cat_id FROM marketplace_categories WHERE slug = s LIMIT 1;
    IF cat_id IS NOT NULL THEN
      INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
        (cat_id, 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f20', false, 10),
        (cat_id, 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f21', false, 11)
      ON CONFLICT (category_id, attribute_id) DO NOTHING;
    END IF;
  END LOOP;
END $$;
