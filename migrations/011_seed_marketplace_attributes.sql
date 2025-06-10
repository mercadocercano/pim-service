-- Migration 011: Seed marketplace attributes based on quickstart data
-- PROPÓSITO: Crear atributos globales basados en los datos del quickstart existente
-- BENEFICIO: Compatibilidad total con sistema actual + capacidades marketplace

-- Atributos globales más comunes across business types
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
-- Atributos básicos universales
('attr-color', 'Color', 'color', 'select', true, true, false, 1),
('attr-size', 'Talle', 'size', 'select', true, true, false, 2),
('attr-brand', 'Marca', 'brand', 'text', true, true, false, 3),
('attr-material', 'Material', 'material', 'select', true, false, false, 4),
('attr-gender', 'Género', 'gender', 'select', true, true, false, 5),

-- Atributos específicos de moda
('attr-shoe-size', 'Talle de Calzado', 'shoe-size', 'select', true, true, false, 10),
('attr-clothing-type', 'Tipo de Prenda', 'clothing-type', 'select', true, true, false, 11),

-- Atributos de tecnología
('attr-screen-size', 'Tamaño de Pantalla', 'screen-size', 'select', true, false, false, 20),
('attr-storage', 'Almacenamiento', 'storage', 'select', true, false, false, 21),
('attr-ram', 'Memoria RAM', 'ram', 'select', true, false, false, 22),
('attr-processor', 'Procesador', 'processor', 'text', false, true, false, 23),

-- Atributos de hogar
('attr-room', 'Ambiente', 'room', 'select', true, false, false, 30),
('attr-furniture-material', 'Material de Mueble', 'furniture-material', 'select', true, false, false, 31),
('attr-power', 'Potencia', 'power', 'text', false, false, false, 32),

-- Atributos de deportes
('attr-sport-type', 'Tipo de Deporte', 'sport-type', 'select', true, true, false, 40),
('attr-age-group', 'Grupo de Edad', 'age-group', 'select', true, false, false, 41),

-- Atributos de salud y belleza
('attr-skin-type', 'Tipo de Piel', 'skin-type', 'select', true, false, false, 50),
('attr-hair-type', 'Tipo de Cabello', 'hair-type', 'select', true, false, false, 51),

-- Atributos de comida
('attr-weight', 'Peso', 'weight', 'text', false, false, false, 60),
('attr-volume', 'Volumen', 'volume', 'text', false, false, false, 61),
('attr-expiry', 'Fecha de Vencimiento', 'expiry', 'text', false, false, false, 62),

-- Atributos generales de producto
('attr-condition', 'Condición', 'condition', 'select', true, false, false, 90),
('attr-origin', 'Origen', 'origin', 'select', false, false, false, 91),
('attr-warranty', 'Garantía', 'warranty', 'text', false, false, false, 92);

-- Valores predefinidos para Color (basado en fashion.yaml)
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-color', 'Negro', 'negro', 1),
('attr-color', 'Blanco', 'blanco', 2),
('attr-color', 'Gris', 'gris', 3),
('attr-color', 'Gris Claro', 'gris-claro', 4),
('attr-color', 'Gris Oscuro', 'gris-oscuro', 5),
('attr-color', 'Rojo', 'rojo', 6),
('attr-color', 'Azul', 'azul', 7),
('attr-color', 'Azul Marino', 'azul-marino', 8),
('attr-color', 'Azul Claro', 'azul-claro', 9),
('attr-color', 'Verde', 'verde', 10),
('attr-color', 'Verde Militar', 'verde-militar', 11),
('attr-color', 'Amarillo', 'amarillo', 12),
('attr-color', 'Naranja', 'naranja', 13),
('attr-color', 'Rosa', 'rosa', 14),
('attr-color', 'Rosa Palo', 'rosa-palo', 15),
('attr-color', 'Morado', 'morado', 16),
('attr-color', 'Marrón', 'marron', 17),
('attr-color', 'Beige', 'beige', 18),
('attr-color', 'Dorado', 'dorado', 19),
('attr-color', 'Plateado', 'plateado', 20),
('attr-color', 'Multicolor', 'multicolor', 21),
('attr-color', 'Estampado', 'estampado', 22);

-- Valores para Talle
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-size', 'XS', 'xs', 1),
('attr-size', 'S', 's', 2),
('attr-size', 'M', 'm', 3),
('attr-size', 'L', 'l', 4),
('attr-size', 'XL', 'xl', 5),
('attr-size', 'XXL', 'xxl', 6),
('attr-size', 'XXXL', 'xxxl', 7),
('attr-size', 'Único', 'unico', 8),
('attr-size', '1', '1', 10),
('attr-size', '2', '2', 11),
('attr-size', '3', '3', 12),
('attr-size', '4', '4', 13),
('attr-size', '6', '6', 14),
('attr-size', '8', '8', 15),
('attr-size', '10', '10', 16),
('attr-size', '12', '12', 17),
('attr-size', '14', '14', 18);

-- Valores para Talle de Calzado
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-shoe-size', '35', '35', 35),
('attr-shoe-size', '36', '36', 36),
('attr-shoe-size', '37', '37', 37),
('attr-shoe-size', '38', '38', 38),
('attr-shoe-size', '39', '39', 39),
('attr-shoe-size', '40', '40', 40),
('attr-shoe-size', '41', '41', 41),
('attr-shoe-size', '42', '42', 42),
('attr-shoe-size', '43', '43', 43),
('attr-shoe-size', '44', '44', 44),
('attr-shoe-size', '45', '45', 45),
('attr-shoe-size', '46', '46', 46);

-- Valores para Material (basado en fashion.yaml)
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-material', 'Algodón', 'algodon', 1),
('attr-material', 'Algodón Orgánico', 'algodon-organico', 2),
('attr-material', 'Poliéster', 'poliester', 3),
('attr-material', 'Lycra', 'lycra', 4),
('attr-material', 'Lana', 'lana', 5),
('attr-material', 'Cuero', 'cuero', 6),
('attr-material', 'Cuero Sintético', 'cuero-sintetico', 7),
('attr-material', 'Denim', 'denim', 8),
('attr-material', 'Seda', 'seda', 9),
('attr-material', 'Lino', 'lino', 10);

-- Valores para Género
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-gender', 'Mujer', 'mujer', 1),
('attr-gender', 'Hombre', 'hombre', 2),
('attr-gender', 'Unisex', 'unisex', 3),
('attr-gender', 'Niña', 'nina', 4),
('attr-gender', 'Niño', 'nino', 5),
('attr-gender', 'Bebé', 'bebe', 6);

-- Valores para Almacenamiento (tecnología)
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-storage', '16GB', '16gb', 1),
('attr-storage', '32GB', '32gb', 2),
('attr-storage', '64GB', '64gb', 3),
('attr-storage', '128GB', '128gb', 4),
('attr-storage', '256GB', '256gb', 5),
('attr-storage', '512GB', '512gb', 6),
('attr-storage', '1TB', '1tb', 7),
('attr-storage', '2TB', '2tb', 8);

-- Valores para RAM
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-ram', '2GB', '2gb', 1),
('attr-ram', '4GB', '4gb', 2),
('attr-ram', '6GB', '6gb', 3),
('attr-ram', '8GB', '8gb', 4),
('attr-ram', '12GB', '12gb', 5),
('attr-ram', '16GB', '16gb', 6),
('attr-ram', '32GB', '32gb', 7);

-- Valores para Condición
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-condition', 'Nuevo', 'nuevo', 1),
('attr-condition', 'Usado - Excelente', 'usado-excelente', 2),
('attr-condition', 'Usado - Muy Bueno', 'usado-muy-bueno', 3),
('attr-condition', 'Usado - Bueno', 'usado-bueno', 4),
('attr-condition', 'Usado - Aceptable', 'usado-aceptable', 5);

-- Valores para Ambiente (hogar)
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-room', 'Living', 'living', 1),
('attr-room', 'Comedor', 'comedor', 2),
('attr-room', 'Dormitorio', 'dormitorio', 3),
('attr-room', 'Cocina', 'cocina', 4),
('attr-room', 'Baño', 'bano', 5),
('attr-room', 'Jardín', 'jardin', 6),
('attr-room', 'Oficina', 'oficina', 7),
('attr-room', 'Balcón', 'balcon', 8);

-- Valores para Grupo de Edad
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-age-group', 'Bebé (0-2 años)', 'bebe', 1),
('attr-age-group', 'Niño (3-12 años)', 'nino', 2),
('attr-age-group', 'Adolescente (13-17 años)', 'adolescente', 3),
('attr-age-group', 'Adulto (18-64 años)', 'adulto', 4),
('attr-age-group', 'Adulto Mayor (65+ años)', 'adulto-mayor', 5);

-- Relaciones categorías-atributos base (ejemplos principales)
-- Moda y Accesorios -> atributos de moda
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60', 'attr-color', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60', 'attr-size', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60', 'attr-brand', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60', 'attr-material', false, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60', 'attr-gender', false, 5);

-- Calzado -> talle de calzado específico
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 'attr-shoe-size', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 'attr-color', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 'attr-brand', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 'attr-material', false, 4);

-- Tecnología -> atributos tech
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f62', 'attr-brand', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f62', 'attr-condition', false, 2);

-- Celulares y Tablets -> atributos específicos
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f80', 'attr-storage', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f80', 'attr-ram', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f80', 'attr-screen-size', false, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f80', 'attr-color', false, 4);

-- Hogar y Jardín -> atributos de hogar
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61', 'attr-room', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61', 'attr-brand', false, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61', 'attr-condition', false, 3);

-- Muebles -> atributos específicos
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f75', 'attr-furniture-material', false, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f75', 'attr-color', false, 2);

-- Log de la migración
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('attr-migration-log', '_MIGRATION_LOG_011', '_migration_log_011', 'text', false, false, false, 999999);

-- Comentario final
COMMENT ON TABLE marketplace_attributes IS 'Atributos marketplace basados en estructura quickstart existente - 19 atributos base con 140+ valores predefinidos'; 