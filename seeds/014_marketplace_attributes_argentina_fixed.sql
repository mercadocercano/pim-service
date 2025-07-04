-- Seeder 014: Marketplace Attributes Argentina (Fixed with UUIDs)
-- PROPÓSITO: Atributos específicos para marketplace argentino con valores reales
-- BENEFICIO: Filtros útiles para usuarios argentinos con datos locales

-- ==============================================
-- ATRIBUTOS GENERALES (aplicables a múltiples categorías)
-- ==============================================

-- MARCA (universal)
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'Marca', 'marca', 'text', true, true, false, 1)
ON CONFLICT (id) DO NOTHING;

-- COLOR GENERAL
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Color', 'color', 'select', true, true, false, 3)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Negro', 'negro', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Blanco', 'blanco', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Gris', 'gris', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Azul', 'azul', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Rojo', 'rojo', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Verde', 'verde', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Amarillo', 'amarillo', 7),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Rosa', 'rosa', 8),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Violeta', 'violeta', 9),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Marrón', 'marron', 10),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Beige', 'beige', 11),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Dorado', 'dorado', 12),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Plateado', 'plateado', 13),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Multicolor', 'multicolor', 14)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- ATRIBUTOS DE MODA Y ACCESORIOS
-- ==============================================

-- TALLE ROPA
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'Talle', 'talle-ropa', 'select', true, true, false, 10)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'XS', 'xs', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'S', 's', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'M', 'm', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'L', 'l', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'XL', 'xl', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'XXL', 'xxl', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'XXXL', 'xxxl', 7),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', '1', '1', 8),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', '2', '2', 9),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', '3', '3', 10),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', '4', '4', 11),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', '6', '6', 12),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', '8', '8', 13),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', '10', '10', 14),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', '12', '12', 15),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', '14', '14', 16),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', '16', '16', 17),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', '18', '18', 18),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'Único', 'unico', 19)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- TALLE CALZADO ARGENTINO
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', 'Talle Calzado', 'talle-calzado', 'select', true, true, false, 11)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '32', '32', 32),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '33', '33', 33),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '34', '34', 34),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '35', '35', 35),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '36', '36', 36),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '37', '37', 37),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '38', '38', 38),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '39', '39', 39),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '40', '40', 40),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '41', '41', 41),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '42', '42', 42),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '43', '43', 43),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '44', '44', 44),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '45', '45', 45),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', '46', '46', 46)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- MATERIAL TEXTIL
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Material', 'material-textil', 'select', true, false, false, 12)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Algodón', 'algodon', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Algodón Orgánico', 'algodon-organico', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Poliéster', 'poliester', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Lycra', 'lycra', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Lana', 'lana', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Cuero', 'cuero', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Cuero Sintético', 'cuero-sintetico', 7),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Denim', 'denim', 8),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Seda', 'seda', 9),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Lino', 'lino', 10),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Modal', 'modal', 11),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Viscosa', 'viscosa', 12)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- GÉNERO
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'Género', 'genero', 'select', true, true, false, 13)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'Mujer', 'mujer', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'Hombre', 'hombre', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'Unisex', 'unisex', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'Niña', 'nina', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'Niño', 'nino', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'Bebé', 'bebe', 6)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- ATRIBUTOS DE TECNOLOGÍA
-- ==============================================

-- ALMACENAMIENTO
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', 'Almacenamiento', 'almacenamiento', 'select', true, true, false, 20)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', '16GB', '16gb', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', '32GB', '32gb', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', '64GB', '64gb', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', '128GB', '128gb', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', '256GB', '256gb', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', '512GB', '512gb', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', '1TB', '1tb', 7),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', '2TB', '2tb', 8),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f07', '4TB', '4tb', 9)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- MEMORIA RAM
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', 'Memoria RAM', 'ram', 'select', true, false, false, 21)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', '2GB', '2gb', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', '4GB', '4gb', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', '6GB', '6gb', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', '8GB', '8gb', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', '12GB', '12gb', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', '16GB', '16gb', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', '32GB', '32gb', 7),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f08', '64GB', '64gb', 8)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- MARCAS TECNOLOGÍA (específicas)
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'Marca Tecnología', 'marca-tech', 'select', true, true, false, 22)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'Apple', 'apple', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'Samsung', 'samsung', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'Huawei', 'huawei', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'Xiaomi', 'xiaomi', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'Motorola', 'motorola', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'LG', 'lg', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'Sony', 'sony', 7),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'HP', 'hp', 8),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'Dell', 'dell', 9),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'Lenovo', 'lenovo', 10),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'Asus', 'asus', 11),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f09', 'Acer', 'acer', 12)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- ATRIBUTOS DE HOGAR Y JARDÍN
-- ==============================================

-- MATERIAL HOGAR
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', 'Material', 'material-hogar', 'select', true, false, false, 30)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', 'Madera', 'madera', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', 'Metal', 'metal', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', 'Plástico', 'plastico', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', 'Vidrio', 'vidrio', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', 'Tela', 'tela', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', 'Cuero', 'cuero', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', 'MDF', 'mdf', 7),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', 'Melamina', 'melamina', 8),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', 'Hierro', 'hierro', 9),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f10', 'Aluminio', 'aluminio', 10)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- AMBIENTE
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', 'Ambiente', 'ambiente', 'select', true, false, false, 31)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', 'Living', 'living', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', 'Comedor', 'comedor', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', 'Dormitorio', 'dormitorio', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', 'Cocina', 'cocina', 4),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', 'Baño', 'bano', 5),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', 'Jardín', 'jardin', 6),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', 'Oficina', 'oficina', 7),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', 'Balcón', 'balcon', 8),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f11', 'Terraza', 'terraza', 9)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- TAMAÑO
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', 'Tamaño', 'tamano', 'select', true, false, false, 32)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', 'Pequeño', 'pequeno', 1),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', 'Mediano', 'mediano', 2),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', 'Grande', 'grande', 3),
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f12', 'Extra Grande', 'extra-grande', 4)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Log del seeder
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('a1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f99', '_SEEDER_LOG_014', '_seeder_log_014', 'text', false, false, false, 999999)
ON CONFLICT (id) DO NOTHING;

-- Comentario final
COMMENT ON TABLE marketplace_attributes IS 'Seeder 014: Atributos marketplace específicos para Argentina con valores locales (UUIDs corregidos)'; 