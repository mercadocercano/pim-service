-- Seed de categorías principales
INSERT INTO categories (id, tenant_id, name, description, parent_id, created_at, updated_at)
VALUES 
    (gen_random_uuid(), 'default', 'Hogar y Jardín', 'Muebles, decoración, herramientas de jardín y más', NULL, NOW(), NOW()),
    (gen_random_uuid(), 'default', 'Ropa y Accesorios', 'Moda para hombres, mujeres y niños', NULL, NOW(), NOW()),
    (gen_random_uuid(), 'default', 'Salud y Belleza', 'Productos para el cuidado personal, cosméticos y suplementos', NULL, NOW(), NOW()),
    (gen_random_uuid(), 'default', 'Deportes y Aire Libre', 'Artículos deportivos, camping y actividades al aire libre', NULL, NOW(), NOW()),
    (gen_random_uuid(), 'default', 'Automotriz', 'Accesorios para autos, repuestos y herramientas', NULL, NOW(), NOW()),
    (gen_random_uuid(), 'default', 'Mascotas', 'Alimentos, juguetes y accesorios para mascotas', NULL, NOW(), NOW()),
    (gen_random_uuid(), 'default', 'Juguetes y Juegos', 'Juguetes para todas las edades, juegos de mesa y videojuegos', NULL, NOW(), NOW()),
    (gen_random_uuid(), 'default', 'Alimentos y Bebidas', 'Productos comestibles, gourmet, bebidas y más', NULL, NOW(), NOW()),
    (gen_random_uuid(), 'default', 'Libros y Educación', 'Libros físicos, ebooks, material educativo y cursos', NULL, NOW(), NOW()),
    (gen_random_uuid(), 'default', 'Servicios', 'Servicios profesionales, técnicos y personales', NULL, NOW(), NOW()); 