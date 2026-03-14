-- Seeder 018: Marketplace Attributes para Librería / Papelería
-- PROPÓSITO: Atributos de Formato, Cantidad de Hojas y Tipo de Rayado
-- IDEMPOTENTE: Usa ON CONFLICT DO NOTHING

-- ==============================================
-- FORMATO CUADERNO
-- ==============================================
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f25', 'Formato', 'formato-cuaderno', 'select', true, true, false, 60)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f25', 'Tapa Dura', 'tapa-dura', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f25', 'Tapa Blanda', 'tapa-blanda', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f25', 'Espiralado', 'espiralado', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f25', 'A4', 'a4', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f25', 'A5', 'a5', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f25', 'Oficio', 'oficio', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f25', 'Universitario', 'universitario', 7)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- CANTIDAD DE HOJAS
-- ==============================================
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f26', 'Cantidad de Hojas', 'cantidad-hojas', 'select', true, false, false, 61)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f26', '24 hojas', '24-hojas', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f26', '48 hojas', '48-hojas', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f26', '84 hojas', '84-hojas', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f26', '96 hojas', '96-hojas', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f26', '150 hojas', '150-hojas', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f26', '200 hojas', '200-hojas', 6)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- TIPO DE RAYADO
-- ==============================================
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f27', 'Tipo de Rayado', 'tipo-rayado', 'select', true, false, false, 62)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f27', 'Rayado', 'rayado', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f27', 'Cuadriculado', 'cuadriculado', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f27', 'Liso', 'liso', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f27', 'Pautado', 'pautado', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f27', 'Puntos (Dot Grid)', 'dot-grid', 5)
ON CONFLICT (attribute_id, slug) DO NOTHING;
