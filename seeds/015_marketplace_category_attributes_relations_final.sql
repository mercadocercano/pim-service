-- Seeder 015: Marketplace Category-Attribute Relations (Final - Real IDs)
-- PROPÓSITO: Mapear atributos específicos a categorías para filtros contextuales
-- BENEFICIO: Usuarios ven solo filtros relevantes por categoría

-- ==============================================
-- RELACIONES MODA E INDUMENTARIA
-- ==============================================

-- MODA E INDUMENTARIA (9c7dd83f-9a5e-4885-b105-81df963a52fb)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('9c7dd83f-9a5e-4885-b105-81df963a52fb', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('9c7dd83f-9a5e-4885-b105-81df963a52fb', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('9c7dd83f-9a5e-4885-b105-81df963a52fb', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', false, 3), -- Talle
('9c7dd83f-9a5e-4885-b105-81df963a52fb', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', false, 4), -- Talle Calzado
('9c7dd83f-9a5e-4885-b105-81df963a52fb', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', false, 5), -- Material Textil
('9c7dd83f-9a5e-4885-b105-81df963a52fb', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', false, 6)  -- Género
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES TECNOLOGÍA Y ELECTRÓNICOS
-- ==============================================

-- TECNOLOGÍA Y ELECTRÓNICOS (5bc8c462-3ce2-46f8-b454-55c7c3d5c924)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('5bc8c462-3ce2-46f8-b454-55c7c3d5c924', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', false, 1), -- Marca Tech
('5bc8c462-3ce2-46f8-b454-55c7c3d5c924', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('5bc8c462-3ce2-46f8-b454-55c7c3d5c924', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', false, 3), -- Almacenamiento
('5bc8c462-3ce2-46f8-b454-55c7c3d5c924', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', false, 4)  -- RAM
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES HOGAR Y DECORACIÓN
-- ==============================================

-- HOGAR Y DECORACIÓN (bd1ca977-db94-4e89-b699-b71c5fbfd17d)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('bd1ca977-db94-4e89-b699-b71c5fbfd17d', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('bd1ca977-db94-4e89-b699-b71c5fbfd17d', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('bd1ca977-db94-4e89-b699-b71c5fbfd17d', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', false, 3), -- Material Hogar
('bd1ca977-db94-4e89-b699-b71c5fbfd17d', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', false, 4), -- Ambiente
('bd1ca977-db94-4e89-b699-b71c5fbfd17d', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 5)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES ELECTRODOMÉSTICOS
-- ==============================================

-- ELECTRODOMÉSTICOS (05e06d40-6d4e-4f50-9822-bc813b3b8e61)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('05e06d40-6d4e-4f50-9822-bc813b3b8e61', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('05e06d40-6d4e-4f50-9822-bc813b3b8e61', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('05e06d40-6d4e-4f50-9822-bc813b3b8e61', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', false, 3), -- Ambiente
('05e06d40-6d4e-4f50-9822-bc813b3b8e61', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 4)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES DEPORTES Y FITNESS
-- ==============================================

-- DEPORTES Y FITNESS (4431e331-6bca-4ec0-80b7-4655243d31e7)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('4431e331-6bca-4ec0-80b7-4655243d31e7', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('4431e331-6bca-4ec0-80b7-4655243d31e7', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('4431e331-6bca-4ec0-80b7-4655243d31e7', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', false, 3), -- Talle
('4431e331-6bca-4ec0-80b7-4655243d31e7', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', false, 4), -- Talle Calzado
('4431e331-6bca-4ec0-80b7-4655243d31e7', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', false, 5), -- Material
('4431e331-6bca-4ec0-80b7-4655243d31e7', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', false, 6), -- Género
('4431e331-6bca-4ec0-80b7-4655243d31e7', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 7)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES AUTOMOTRIZ
-- ==============================================

-- AUTOMOTRIZ (e6b75c80-70a5-4bb6-ad77-38154847609e)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('e6b75c80-70a5-4bb6-ad77-38154847609e', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('e6b75c80-70a5-4bb6-ad77-38154847609e', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('e6b75c80-70a5-4bb6-ad77-38154847609e', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', false, 3)  -- Material
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES BELLEZA Y CUIDADO PERSONAL
-- ==============================================

-- BELLEZA Y CUIDADO PERSONAL (a6fed950-b19f-476c-b191-9a166c75786c)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('a6fed950-b19f-476c-b191-9a166c75786c', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('a6fed950-b19f-476c-b191-9a166c75786c', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('a6fed950-b19f-476c-b191-9a166c75786c', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', false, 3), -- Género
('a6fed950-b19f-476c-b191-9a166c75786c', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 4)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES MASCOTAS
-- ==============================================

-- MASCOTAS (0be2ec6e-3c2a-4e44-a347-c513e6770f60)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('0be2ec6e-3c2a-4e44-a347-c513e6770f60', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('0be2ec6e-3c2a-4e44-a347-c513e6770f60', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('0be2ec6e-3c2a-4e44-a347-c513e6770f60', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 3)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES HERRAMIENTAS Y CONSTRUCCIÓN
-- ==============================================

-- HERRAMIENTAS Y CONSTRUCCIÓN (afe9c235-39e7-4ed3-baeb-407793923b89)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('afe9c235-39e7-4ed3-baeb-407793923b89', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('afe9c235-39e7-4ed3-baeb-407793923b89', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('afe9c235-39e7-4ed3-baeb-407793923b89', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', false, 3), -- Material
('afe9c235-39e7-4ed3-baeb-407793923b89', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 4)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Comentario final
COMMENT ON TABLE marketplace_category_attributes IS 'Seeder 015: Relaciones categoría-atributos contextuales usando IDs reales de la base de datos'; 