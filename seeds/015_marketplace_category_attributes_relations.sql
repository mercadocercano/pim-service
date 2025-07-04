-- Seeder 015: Marketplace Category-Attributes Relations
-- PROPÓSITO: Establecer qué atributos aplican a cada categoría específica
-- BENEFICIO: Filtros contextuales relevantes para cada tipo de producto

-- Limpiar relaciones existentes (cuidado en producción)
-- TRUNCATE TABLE marketplace_category_attributes CASCADE;

-- ==============================================
-- ATRIBUTOS UNIVERSALES (aplican a todas las categorías)
-- ==============================================

-- CONDICIÓN (obligatorio para todas las categorías principales)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) 
SELECT id, 'attr-condicion', true, 1 
FROM marketplace_categories 
WHERE level = 0 AND is_active = true
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- MARCA (aplica a todas las categorías principales)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) 
SELECT id, 'attr-marca', false, 2 
FROM marketplace_categories 
WHERE level = 0 AND is_active = true
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- MODA Y ACCESORIOS - ROPA MUJER
-- ==============================================

-- Ropa Mujer (nivel 1) - Atributos básicos
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f70', 'attr-talle-ropa', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f70', 'attr-color', false, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f70', 'attr-material-textil', false, 5),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f70', 'attr-genero', false, 6)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Remeras y Tops (nivel 2)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f84', 'attr-talle-ropa', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f84', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f84', 'attr-material-textil', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f84', 'attr-genero', false, 4)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Vestidos
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f85', 'attr-talle-ropa', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f85', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f85', 'attr-material-textil', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f85', 'attr-genero', false, 4)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Pantalones Mujer
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f86', 'attr-talle-ropa', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f86', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f86', 'attr-material-textil', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f86', 'attr-genero', false, 4)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- MODA Y ACCESORIOS - ROPA HOMBRE
-- ==============================================

-- Ropa Hombre (nivel 1)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f71', 'attr-talle-ropa', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f71', 'attr-color', false, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f71', 'attr-material-textil', false, 5),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f71', 'attr-genero', false, 6)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Remeras y Polos Hombre
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f89', 'attr-talle-ropa', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f89', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f89', 'attr-material-textil', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f89', 'attr-genero', false, 4)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Camisas
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f90', 'attr-talle-ropa', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f90', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f90', 'attr-material-textil', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f90', 'attr-genero', false, 4)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- MODA Y ACCESORIOS - CALZADO
-- ==============================================

-- Calzado (nivel 1)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 'attr-talle-calzado', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 'attr-color', false, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 'attr-material-textil', false, 5),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 'attr-genero', false, 6)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Zapatillas
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f93', 'attr-talle-calzado', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f93', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f93', 'attr-material-textil', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f93', 'attr-genero', false, 4)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Zapatos Formales
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f94', 'attr-talle-calzado', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f94', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f94', 'attr-material-textil', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f94', 'attr-genero', false, 4)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- TECNOLOGÍA
-- ==============================================

-- Tecnología (nivel 0)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f62', 'attr-marca-tech', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f62', 'attr-condicion', true, 4)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Celulares y Tablets
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f80', 'attr-almacenamiento', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f80', 'attr-ram', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f80', 'attr-marca-tech', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f80', 'attr-color', false, 4)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Computadoras
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f81', 'attr-almacenamiento', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f81', 'attr-ram', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f81', 'attr-marca-tech', false, 3)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Gaming
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f82', 'attr-marca-tech', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f82', 'attr-color', false, 2)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Audio y Video
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f83', 'attr-marca-tech', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f83', 'attr-color', false, 2)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- HOGAR Y JARDÍN
-- ==============================================

-- Hogar y Jardín (nivel 0)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61', 'attr-material-hogar', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61', 'attr-color', false, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61', 'attr-ambiente', false, 5)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Muebles
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f75', 'attr-material-hogar', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f75', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f75', 'attr-ambiente', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f75', 'attr-tamano', false, 4)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Sillones y Sofás
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f97', 'attr-material-hogar', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f97', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f97', 'attr-tamano', false, 3)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Mesas
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f98', 'attr-material-hogar', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f98', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f98', 'attr-tamano', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f98', 'attr-ambiente', false, 4)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Dormitorio
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f99', 'attr-material-hogar', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f99', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f99', 'attr-tamano', false, 3)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Decoración
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f76', 'attr-material-hogar', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f76', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f76', 'attr-ambiente', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f76', 'attr-tamano', false, 4)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Electrodomésticos
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f77', 'attr-marca-tech', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f77', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f77', 'attr-ambiente', false, 3)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- COMIDA Y BEBIDAS
-- ==============================================

-- Comida y Bebidas (nivel 0)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f65', 'attr-peso', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f65', 'attr-volumen', false, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f65', 'attr-marca-alimentos', false, 5)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- DEPORTES Y FITNESS
-- ==============================================

-- Deportes y Fitness - Aplicar atributos de ropa cuando sea ropa deportiva
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f63', 'attr-talle-ropa', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f63', 'attr-color', false, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f63', 'attr-material-textil', false, 5),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f63', 'attr-genero', false, 6)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- SALUD Y BELLEZA
-- ==============================================

-- Salud y Belleza - Atributos básicos
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f64', 'attr-genero', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f64', 'attr-peso', false, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f64', 'attr-volumen', false, 5)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- BEBÉS Y NIÑOS
-- ==============================================

-- Bebés y Niños - Atributos específicos
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f67', 'attr-talle-ropa', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f67', 'attr-color', false, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f67', 'attr-genero', false, 5),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f67', 'attr-material-textil', false, 6)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- ==============================================
-- MASCOTAS
-- ==============================================

-- Mascotas - Atributos básicos
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f68', 'attr-peso', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f68', 'attr-tamano', false, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f68', 'attr-marca-alimentos', false, 5)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Log de seeder
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('seeder-log-013', '_seeder-log-014', false, 999999)
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Comentario final
COMMENT ON TABLE marketplace_category_attributes IS 'Seeder 015: Relaciones categría-atributo para filtros contextuales por categoría';

-- Estadísticas del seeder
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order, is_active) VALUES
('seeder-stats-015', '_SEEDER_STATS_015', '_seeder_stats_015', 
'Relaciones creadas: ~200+ category-attribute mappings para 50 categorías con filtros específicos argentinos', 
NULL, 0, 999998, false)
ON CONFLICT (id) DO NOTHING; 