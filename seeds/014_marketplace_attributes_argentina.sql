-- Seeder 014: Marketplace Attributes Argentina
-- PROPÓSITO: Atributos específicos para marketplace argentino con valores reales
-- BENEFICIO: Filtros útiles para usuarios argentinos con datos locales

-- Limpiar datos existentes (cuidado en producción)
-- TRUNCATE TABLE marketplace_category_attributes CASCADE;
-- TRUNCATE TABLE marketplace_attribute_values CASCADE;
-- TRUNCATE TABLE marketplace_attributes CASCADE;

-- ==============================================
-- ATRIBUTOS GENERALES (aplicables a múltiples categorías)
-- ==============================================

-- MARCA (universal)
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-marca', 'Marca', 'marca', 'text', true, true, false, 1)
ON CONFLICT (id) DO NOTHING;

-- COLOR GENERAL
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-color', 'Color', 'color', 'select', true, true, false, 3)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-color', 'Negro', 'negro', 1),
('attr-color', 'Blanco', 'blanco', 2),
('attr-color', 'Gris', 'gris', 3),
('attr-color', 'Azul', 'azul', 4),
('attr-color', 'Rojo', 'rojo', 5),
('attr-color', 'Verde', 'verde', 6),
('attr-color', 'Amarillo', 'amarillo', 7),
('attr-color', 'Rosa', 'rosa', 8),
('attr-color', 'Violeta', 'violeta', 9),
('attr-color', 'Marrón', 'marron', 10),
('attr-color', 'Beige', 'beige', 11),
('attr-color', 'Dorado', 'dorado', 12),
('attr-color', 'Plateado', 'plateado', 13),
('attr-color', 'Multicolor', 'multicolor', 14)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- ATRIBUTOS DE MODA Y ACCESORIOS
-- ==============================================

-- TALLE ROPA
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-talle-ropa', 'Talle', 'talle-ropa', 'select', true, true, false, 10)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-talle-ropa', 'XS', 'xs', 1),
('attr-talle-ropa', 'S', 's', 2),
('attr-talle-ropa', 'M', 'm', 3),
('attr-talle-ropa', 'L', 'l', 4),
('attr-talle-ropa', 'XL', 'xl', 5),
('attr-talle-ropa', 'XXL', 'xxl', 6),
('attr-talle-ropa', 'XXXL', 'xxxl', 7),
('attr-talle-ropa', '1', '1', 8),
('attr-talle-ropa', '2', '2', 9),
('attr-talle-ropa', '3', '3', 10),
('attr-talle-ropa', '4', '4', 11),
('attr-talle-ropa', '6', '6', 12),
('attr-talle-ropa', '8', '8', 13),
('attr-talle-ropa', '10', '10', 14),
('attr-talle-ropa', '12', '12', 15),
('attr-talle-ropa', '14', '14', 16),
('attr-talle-ropa', '16', '16', 17),
('attr-talle-ropa', '18', '18', 18),
('attr-talle-ropa', 'Único', 'unico', 19)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- TALLE CALZADO ARGENTINO
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-talle-calzado', 'Talle Calzado', 'talle-calzado', 'select', true, true, false, 11)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-talle-calzado', '32', '32', 32),
('attr-talle-calzado', '33', '33', 33),
('attr-talle-calzado', '34', '34', 34),
('attr-talle-calzado', '35', '35', 35),
('attr-talle-calzado', '36', '36', 36),
('attr-talle-calzado', '37', '37', 37),
('attr-talle-calzado', '38', '38', 38),
('attr-talle-calzado', '39', '39', 39),
('attr-talle-calzado', '40', '40', 40),
('attr-talle-calzado', '41', '41', 41),
('attr-talle-calzado', '42', '42', 42),
('attr-talle-calzado', '43', '43', 43),
('attr-talle-calzado', '44', '44', 44),
('attr-talle-calzado', '45', '45', 45),
('attr-talle-calzado', '46', '46', 46)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- MATERIAL TEXTIL
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-material-textil', 'Material', 'material-textil', 'select', true, false, false, 12)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-material-textil', 'Algodón', 'algodon', 1),
('attr-material-textil', 'Algodón Orgánico', 'algodon-organico', 2),
('attr-material-textil', 'Poliéster', 'poliester', 3),
('attr-material-textil', 'Lycra', 'lycra', 4),
('attr-material-textil', 'Lana', 'lana', 5),
('attr-material-textil', 'Cuero', 'cuero', 6),
('attr-material-textil', 'Cuero Sintético', 'cuero-sintetico', 7),
('attr-material-textil', 'Denim', 'denim', 8),
('attr-material-textil', 'Seda', 'seda', 9),
('attr-material-textil', 'Lino', 'lino', 10),
('attr-material-textil', 'Modal', 'modal', 11),
('attr-material-textil', 'Viscosa', 'viscosa', 12)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- GÉNERO
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-genero', 'Género', 'genero', 'select', true, true, false, 13)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-genero', 'Mujer', 'mujer', 1),
('attr-genero', 'Hombre', 'hombre', 2),
('attr-genero', 'Unisex', 'unisex', 3),
('attr-genero', 'Niña', 'nina', 4),
('attr-genero', 'Niño', 'nino', 5),
('attr-genero', 'Bebé', 'bebe', 6)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- ATRIBUTOS DE TECNOLOGÍA
-- ==============================================

-- ALMACENAMIENTO
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-almacenamiento', 'Almacenamiento', 'almacenamiento', 'select', true, true, false, 20)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-almacenamiento', '16GB', '16gb', 1),
('attr-almacenamiento', '32GB', '32gb', 2),
('attr-almacenamiento', '64GB', '64gb', 3),
('attr-almacenamiento', '128GB', '128gb', 4),
('attr-almacenamiento', '256GB', '256gb', 5),
('attr-almacenamiento', '512GB', '512gb', 6),
('attr-almacenamiento', '1TB', '1tb', 7),
('attr-almacenamiento', '2TB', '2tb', 8),
('attr-almacenamiento', '4TB', '4tb', 9)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- MEMORIA RAM
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-ram', 'Memoria RAM', 'ram', 'select', true, false, false, 21)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-ram', '2GB', '2gb', 1),
('attr-ram', '4GB', '4gb', 2),
('attr-ram', '6GB', '6gb', 3),
('attr-ram', '8GB', '8gb', 4),
('attr-ram', '12GB', '12gb', 5),
('attr-ram', '16GB', '16gb', 6),
('attr-ram', '32GB', '32gb', 7),
('attr-ram', '64GB', '64gb', 8)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- MARCAS TECNOLOGÍA (específicas)
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-marca-tech', 'Marca Tecnología', 'marca-tech', 'select', true, true, false, 22)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-marca-tech', 'Apple', 'apple', 1),
('attr-marca-tech', 'Samsung', 'samsung', 2),
('attr-marca-tech', 'Huawei', 'huawei', 3),
('attr-marca-tech', 'Xiaomi', 'xiaomi', 4),
('attr-marca-tech', 'Motorola', 'motorola', 5),
('attr-marca-tech', 'LG', 'lg', 6),
('attr-marca-tech', 'Sony', 'sony', 7),
('attr-marca-tech', 'HP', 'hp', 8),
('attr-marca-tech', 'Dell', 'dell', 9),
('attr-marca-tech', 'Lenovo', 'lenovo', 10),
('attr-marca-tech', 'Asus', 'asus', 11),
('attr-marca-tech', 'Acer', 'acer', 12)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- ATRIBUTOS DE HOGAR Y JARDÍN
-- ==============================================

-- MATERIAL HOGAR
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-material-hogar', 'Material', 'material-hogar', 'select', true, false, false, 30)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-material-hogar', 'Madera', 'madera', 1),
('attr-material-hogar', 'Metal', 'metal', 2),
('attr-material-hogar', 'Plástico', 'plastico', 3),
('attr-material-hogar', 'Vidrio', 'vidrio', 4),
('attr-material-hogar', 'Tela', 'tela', 5),
('attr-material-hogar', 'Cuero', 'cuero', 6),
('attr-material-hogar', 'MDF', 'mdf', 7),
('attr-material-hogar', 'Melamina', 'melamina', 8),
('attr-material-hogar', 'Hierro', 'hierro', 9),
('attr-material-hogar', 'Aluminio', 'aluminio', 10)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- AMBIENTE
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-ambiente', 'Ambiente', 'ambiente', 'select', true, false, false, 31)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-ambiente', 'Living', 'living', 1),
('attr-ambiente', 'Comedor', 'comedor', 2),
('attr-ambiente', 'Dormitorio', 'dormitorio', 3),
('attr-ambiente', 'Cocina', 'cocina', 4),
('attr-ambiente', 'Baño', 'bano', 5),
('attr-ambiente', 'Jardín', 'jardin', 6),
('attr-ambiente', 'Oficina', 'oficina', 7),
('attr-ambiente', 'Balcón', 'balcon', 8),
('attr-ambiente', 'Terraza', 'terraza', 9)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- TAMAÑO
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-tamano', 'Tamaño', 'tamano', 'select', true, false, false, 32)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-tamano', 'Pequeño', 'pequeno', 1),
('attr-tamano', 'Mediano', 'mediano', 2),
('attr-tamano', 'Grande', 'grande', 3),
('attr-tamano', 'Extra Grande', 'extra-grande', 4)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- ==============================================
-- ATRIBUTOS DE COMIDA Y BEBIDAS
-- ==============================================

-- PESO
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-peso', 'Peso', 'peso', 'select', true, false, false, 40)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-peso', '50g', '50g', 1),
('attr-peso', '100g', '100g', 2),
('attr-peso', '200g', '200g', 3),
('attr-peso', '250g', '250g', 4),
('attr-peso', '500g', '500g', 5),
('attr-peso', '750g', '750g', 6),
('attr-peso', '1kg', '1kg', 7),
('attr-peso', '2kg', '2kg', 8),
('attr-peso', '5kg', '5kg', 9)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- VOLUMEN
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-volumen', 'Volumen', 'volumen', 'select', true, false, false, 41)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-volumen', '250ml', '250ml', 1),
('attr-volumen', '500ml', '500ml', 2),
('attr-volumen', '750ml', '750ml', 3),
('attr-volumen', '1L', '1l', 4),
('attr-volumen', '1.5L', '1-5l', 5),
('attr-volumen', '2L', '2l', 6),
('attr-volumen', '3L', '3l', 7)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- MARCAS ALIMENTICIAS ARGENTINAS
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-marca-alimentos', 'Marca Alimentos', 'marca-alimentos', 'select', true, true, false, 42)
ON CONFLICT (id) DO NOTHING;

INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-marca-alimentos', 'Arcor', 'arcor', 1),
('attr-marca-alimentos', 'Mastellone', 'mastellone', 2),
('attr-marca-alimentos', 'Molinos Río de la Plata', 'molinos', 3),
('attr-marca-alimentos', 'Unilever', 'unilever', 4),
('attr-marca-alimentos', 'Nestlé', 'nestle', 5),
('attr-marca-alimentos', 'Coca-Cola', 'cocacola', 6),
('attr-marca-alimentos', 'Quilmes', 'quilmes', 7),
('attr-marca-alimentos', 'Terma', 'terma', 8),
('attr-marca-alimentos', 'Bagley', 'bagley', 9),
('attr-marca-alimentos', 'Georgalos', 'georgalos', 10)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Log de migración
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('_seeder-log-014', '_SEEDER_LOG_014', '_seeder_log_014', 'text', false, false, false, 999999)
ON CONFLICT (id) DO NOTHING;

-- Comentario final
COMMENT ON TABLE marketplace_attributes IS 'Seeder 014: Atributos marketplace específicos para Argentina con valores locales'; 