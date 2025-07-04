-- Seeder 015: Marketplace Category-Attribute Relations (Fixed with UUIDs)
-- PROPÓSITO: Mapear atributos específicos a categorías para filtros contextuales
-- BENEFICIO: Usuarios ven solo filtros relevantes por categoría

-- ==============================================
-- RELACIONES MODA Y ACCESORIOS
-- ==============================================

-- ROPA DE MUJER
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f01', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f01', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f01', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', false, 3), -- Talle
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f01', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', false, 4), -- Material Textil
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f01', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', false, 5)  -- Género
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ROPA DE HOMBRE
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f02', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f02', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f02', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', false, 3), -- Talle
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f02', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', false, 4), -- Material Textil
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f02', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', false, 5)  -- Género
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- CALZADO DE MUJER
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f03', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f03', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f03', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', false, 3), -- Talle Calzado
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f03', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', false, 4), -- Material
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f03', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', false, 5)  -- Género
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- CALZADO DE HOMBRE
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f04', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f04', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f04', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', false, 3), -- Talle Calzado
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f04', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', false, 4), -- Material
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f04', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', false, 5)  -- Género
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ACCESORIOS
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f05', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f05', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f05', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', false, 3), -- Material
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f05', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', false, 4)  -- Género
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES TECNOLOGÍA
-- ==============================================

-- CELULARES Y SMARTPHONES
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f10', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', false, 1), -- Marca Tech
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f10', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f10', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', false, 3), -- Almacenamiento
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f10', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', false, 4)  -- RAM
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- COMPUTADORAS Y LAPTOPS
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f11', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', false, 1), -- Marca Tech
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f11', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f11', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', false, 3), -- Almacenamiento
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f11', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', false, 4)  -- RAM
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- TABLETS
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f12', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', false, 1), -- Marca Tech
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f12', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f12', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', false, 3), -- Almacenamiento
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f12', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', false, 4)  -- RAM
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- AUDIO Y VIDEO
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f13', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', false, 1), -- Marca Tech
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f13', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2)  -- Color
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ACCESORIOS TECNOLOGÍA
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f14', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', false, 1), -- Marca Tech
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f14', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2)  -- Color
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES HOGAR Y DECORACIÓN
-- ==============================================

-- MUEBLES DE LIVING
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f20', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f20', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f20', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', false, 3), -- Material Hogar
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f20', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', false, 4), -- Ambiente
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f20', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 5)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- MUEBLES DE DORMITORIO
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f21', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f21', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f21', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', false, 3), -- Material Hogar
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f21', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', false, 4), -- Ambiente
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f21', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 5)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- MUEBLES DE COCINA
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f22', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f22', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f22', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', false, 3), -- Material Hogar
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f22', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', false, 4), -- Ambiente
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f22', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 5)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ELECTRODOMÉSTICOS GRANDES
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f23', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f23', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f23', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', false, 3), -- Ambiente
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f23', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 4)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ELECTRODOMÉSTICOS PEQUEÑOS
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f24', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f24', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f24', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', false, 3), -- Ambiente
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f24', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 4)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- DECORACIÓN Y TEXTILES
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f25', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f25', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f25', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', false, 3), -- Material
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f25', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', false, 4), -- Ambiente
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f25', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 5)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES DEPORTES Y FITNESS
-- ==============================================

-- ROPA DEPORTIVA
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f30', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f30', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f30', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', false, 3), -- Talle
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f30', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', false, 4), -- Material
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f30', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', false, 5)  -- Género
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- CALZADO DEPORTIVO
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f31', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f31', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f31', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', false, 3), -- Talle Calzado
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f31', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', false, 4), -- Material
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f31', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', false, 5)  -- Género
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- SUPLEMENTOS Y NUTRICIÓN
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f32', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1) -- Marca
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- EQUIPAMIENTO DEPORTIVO
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f33', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f33', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f33', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 3)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES AUTOS Y VEHÍCULOS
-- ==============================================

-- ACCESORIOS PARA AUTOS
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f40', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f40', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f40', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', false, 3)  -- Material
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- REPUESTOS
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f41', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1) -- Marca
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- RELACIONES BEBÉS Y NIÑOS
-- ==============================================

-- ROPA DE BEBÉ
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f50', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f50', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f50', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', false, 3), -- Talle
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f50', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', false, 4), -- Material
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f50', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', false, 5)  -- Género
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- JUGUETES
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f51', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f51', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f51', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', false, 3), -- Material
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f51', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', false, 4)  -- Género
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- CUIDADO DEL BEBÉ
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f52', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', false, 1), -- Marca
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f52', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', false, 2), -- Color
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f52', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', false, 3), -- Material
('44e6f2a1-2b3c-4d5e-8f9a-1b2c3d4e5f52', 'a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', false, 4)  -- Tamaño
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Comentario final
COMMENT ON TABLE marketplace_category_attributes IS 'Seeder 015: Relaciones categoría-atributos contextuales para filtros específicos por categoría (UUIDs corregidos)'; 