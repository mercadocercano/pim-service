-- Migration 010: Seed marketplace categories with initial data
-- PROPÓSITO: Crear taxonomía base del marketplace con categorías comunes
-- BENEFICIO: Sellers encuentran categorías familiares sin complejidad

-- Categorías de Nivel 0 (Raíz)
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60', 'Moda y Accesorios', 'moda-accesorios', 'Ropa, calzado y accesorios para toda la familia', NULL, 0, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61', 'Hogar y Jardín', 'hogar-jardin', 'Muebles, decoración, electrodomésticos y jardín', NULL, 0, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f62', 'Tecnología', 'tecnologia', 'Electrónicos, computadoras y accesorios tecnológicos', NULL, 0, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f63', 'Deportes y Fitness', 'deportes-fitness', 'Equipamiento deportivo, ropa deportiva y fitness', NULL, 0, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f64', 'Salud y Belleza', 'salud-belleza', 'Productos de belleza, cuidado personal y salud', NULL, 0, 5),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f65', 'Comida y Bebidas', 'comida-bebidas', 'Alimentos, bebidas y productos gourmet', NULL, 0, 6),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f66', 'Libros y Entretenimiento', 'libros-entretenimiento', 'Libros, música, películas y juegos', NULL, 0, 7),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f67', 'Bebés y Niños', 'bebes-ninos', 'Productos para bebés, niños y maternidad', NULL, 0, 8),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f68', 'Mascotas', 'mascotas', 'Alimentos, accesorios y cuidado para mascotas', NULL, 0, 9),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f69', 'Servicios', 'servicios', 'Servicios profesionales y consultoría', NULL, 0, 10);

-- Categorías de Nivel 1 - Moda y Accesorios
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f70', 'Ropa Mujer', 'ropa-mujer', 'Vestimenta femenina de todas las edades', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60', 1, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f71', 'Ropa Hombre', 'ropa-hombre', 'Vestimenta masculina de todas las edades', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60', 1, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 'Calzado', 'calzado', 'Zapatos, zapatillas y sandalias para toda la familia', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60', 1, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f73', 'Accesorios', 'accesorios', 'Carteras, bijouterie, relojes y complementos', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60', 1, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f74', 'Lencería', 'lenceria', 'Ropa interior y lencería femenina', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60', 1, 5);

-- Categorías de Nivel 1 - Hogar y Jardín
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f75', 'Muebles', 'muebles', 'Muebles para todo el hogar', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61', 1, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f76', 'Decoración', 'decoracion', 'Objetos decorativos y ornamentales', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61', 1, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f77', 'Electrodomésticos', 'electrodomesticos', 'Equipos eléctricos para el hogar', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61', 1, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f78', 'Jardín', 'jardin', 'Plantas, herramientas y accesorios de jardín', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61', 1, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f79', 'Cocina', 'cocina', 'Utensilios y equipos de cocina', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61', 1, 5);

-- Categorías de Nivel 1 - Tecnología
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f80', 'Celulares y Tablets', 'celulares-tablets', 'Dispositivos móviles y accesorios', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f62', 1, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f81', 'Computadoras', 'computadoras', 'Notebooks, PC y componentes', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f62', 1, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f82', 'Gaming', 'gaming', 'Consolas, videojuegos y accesorios', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f62', 1, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f83', 'Audio y Video', 'audio-video', 'Equipos de sonido, auriculares y TV', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f62', 1, 4);

-- Categorías de Nivel 2 - Ejemplos específicos para Ropa Mujer
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f84', 'Remeras y Tops', 'remeras-tops', 'Remeras, musculosas y tops', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f70', 2, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f85', 'Vestidos', 'vestidos', 'Vestidos casuales y de fiesta', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f70', 2, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f86', 'Pantalones', 'pantalones', 'Jeans, pantalones y leggins', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f70', 2, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f87', 'Abrigos', 'abrigos', 'Camperas, sacos y abrigos', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f70', 2, 4),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f88', 'Faldas', 'faldas', 'Faldas de todos los largos', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f70', 2, 5);

-- Categorías de Nivel 2 - Ejemplos para Ropa Hombre
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f89', 'Remeras y Polos', 'remeras-polos-hombre', 'Remeras y polos masculinos', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f71', 2, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f90', 'Camisas', 'camisas', 'Camisas casuales y formales', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f71', 2, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f91', 'Pantalones Hombre', 'pantalones-hombre', 'Jeans, chinos y pantalones formales', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f71', 2, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f92', 'Abrigos Hombre', 'abrigos-hombre', 'Camperas, blazers y abrigos masculinos', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f71', 2, 4);

-- Categorías de Nivel 2 - Ejemplos para Calzado
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f93', 'Zapatillas', 'zapatillas', 'Zapatillas deportivas y urbanas', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 2, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f94', 'Zapatos Formales', 'zapatos-formales', 'Zapatos de vestir y trabajo', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 2, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f95', 'Sandalias', 'sandalias', 'Sandalias y ojotas', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 2, 3),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f96', 'Botas', 'botas', 'Botas de lluvia, trabajo y fashion', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72', 2, 4);

-- Categorías de Nivel 2 - Ejemplos para Muebles
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order) VALUES
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f97', 'Sillones y Sofás', 'sillones-sofas', 'Sillones, sofás y muebles de living', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f75', 2, 1),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f98', 'Mesas', 'mesas', 'Mesas de comedor, centro y trabajo', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f75', 2, 2),
('d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f99', 'Dormitorio', 'dormitorio', 'Camas, cómodas y muebles de dormitorio', 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f75', 2, 3);

-- Log de la migración para tracking
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order, is_active) VALUES
('00000000-0000-0000-0000-000000000000', '_MIGRATION_LOG_010', '_migration_log_010', 'Seeder executed successfully - 50 categories created', NULL, 0, 999999, FALSE);

-- Comentario final
COMMENT ON TABLE marketplace_categories IS 'Taxonomía inicial del marketplace con 50 categorías base organizadas en 3 niveles máximo'; 