-- Seed: 003_marketplace_categories_argentina_seed.sql
-- Purpose: Categorías globales del marketplace basadas en investigación de mercado argentino 2024-2025
-- Date: 2025-01-27
-- Fuente: Estudios CACE, Kantar TNS, PCMI, Americas Market Intelligence

-- Limpiar datos existentes
DELETE FROM marketplace_categories;

-- NIVEL 0: CATEGORÍAS PRINCIPALES (basadas en los datos de facturación más alta)
INSERT INTO marketplace_categories (name, slug, description, parent_id, level, sort_order, is_active) VALUES

-- 1. ALIMENTOS Y BEBIDAS (Top 1 en múltiples estudios, crecimiento 111% y 20.7%)
('Alimentos y Bebidas', 'alimentos-bebidas', 'Productos alimentarios, bebidas, comestibles y artículos de consumo diario', NULL, 0, 1, true),

-- 2. MODA E INDUMENTARIA (Top en ventas, $162.9 mil millones en EE.UU.)
('Moda e Indumentaria', 'moda-indumentaria', 'Ropa, calzado, accesorios de moda para toda la familia', NULL, 0, 2, true),

-- 3. TECNOLOGÍA Y ELECTRÓNICOS (Top 3 consistente, $120.1 mil millones en EE.UU.)
('Tecnología y Electrónicos', 'tecnologia-electronicos', 'Dispositivos electrónicos, tecnología informática, equipos de audio e imagen', NULL, 0, 3, true),

-- 4. HOGAR Y DECORACIÓN (Top 4 consistente, $74.5 mil millones en muebles)
('Hogar y Decoración', 'hogar-decoracion', 'Muebles, decoración, artículos para el hogar y jardín', NULL, 0, 4, true),

-- 5. ELECTRODOMÉSTICOS (Top 5 consistente, línea blanca y marrón)
('Electrodomésticos', 'electrodomesticos', 'Electrodomésticos grandes y pequeños, línea blanca y marrón', NULL, 0, 5, true),

-- 6. HERRAMIENTAS Y CONSTRUCCIÓN (Top en Argentina, $108.6 mil millones bricolaje)
('Herramientas y Construcción', 'herramientas-construccion', 'Herramientas, materiales de construcción, ferretería y bricolaje', NULL, 0, 6, true),

-- 7. AUTOMOTRIZ (Top en Argentina, crecimiento 30.1%)
('Automotriz', 'automotriz', 'Repuestos, accesorios automotrices y productos para vehículos', NULL, 0, 7, true),

-- 8. BELLEZA Y CUIDADO PERSONAL (Top consistente, $28.4 mil millones)
('Belleza y Cuidado Personal', 'belleza-cuidado-personal', 'Cosméticos, perfumes, productos de cuidado e higiene personal', NULL, 0, 8, true),

-- 9. DEPORTES Y FITNESS (Presente en múltiples estudios)
('Deportes y Fitness', 'deportes-fitness', 'Artículos deportivos, ropa deportiva, equipamiento y suplementos', NULL, 0, 9, true),

-- 10. MASCOTAS (Nicho en crecimiento según Mailchimp)
('Mascotas', 'mascotas', 'Alimentos, accesorios, productos veterinarios y cuidado de mascotas', NULL, 0, 10, true),

-- 11. SALUD Y FARMACIA (Sector esencial)
('Salud y Farmacia', 'salud-farmacia', 'Medicamentos, productos de salud, equipos médicos y bienestar', NULL, 0, 11, true),

-- 12. LIBROS Y ENTRETENIMIENTO (Multimedia $37.7 mil millones)
('Libros y Entretenimiento', 'libros-entretenimiento', 'Libros, medios audiovisuales, juegos, entradas a espectáculos', NULL, 0, 12, true),

-- 13. SERVICIOS DIGITALES (Contenidos audiovisuales, software)
('Servicios Digitales', 'servicios-digitales', 'Software, aplicaciones, contenido digital, cursos online', NULL, 0, 13, true),

-- 14. TURISMO Y VIAJES (22% de e-commerce en 2019)
('Turismo y Viajes', 'turismo-viajes', 'Pasajes, paquetes turísticos, hospedaje y servicios de viaje', NULL, 0, 14, true);

-- Obtener IDs de categorías principales para subcategorías
DO $$
DECLARE
    cat_alimentos_id UUID;
    cat_moda_id UUID;
    cat_tecnologia_id UUID;
    cat_hogar_id UUID;
    cat_electrodomesticos_id UUID;
    cat_herramientas_id UUID;
    cat_automotriz_id UUID;
    cat_belleza_id UUID;
    cat_deportes_id UUID;
    cat_mascotas_id UUID;
    cat_salud_id UUID;
    cat_libros_id UUID;
BEGIN
    -- Obtener IDs de categorías principales
    SELECT id INTO cat_alimentos_id FROM marketplace_categories WHERE slug = 'alimentos-bebidas';
    SELECT id INTO cat_moda_id FROM marketplace_categories WHERE slug = 'moda-indumentaria';
    SELECT id INTO cat_tecnologia_id FROM marketplace_categories WHERE slug = 'tecnologia-electronicos';
    SELECT id INTO cat_hogar_id FROM marketplace_categories WHERE slug = 'hogar-decoracion';
    SELECT id INTO cat_electrodomesticos_id FROM marketplace_categories WHERE slug = 'electrodomesticos';
    SELECT id INTO cat_herramientas_id FROM marketplace_categories WHERE slug = 'herramientas-construccion';
    SELECT id INTO cat_automotriz_id FROM marketplace_categories WHERE slug = 'automotriz';
    SELECT id INTO cat_belleza_id FROM marketplace_categories WHERE slug = 'belleza-cuidado-personal';
    SELECT id INTO cat_deportes_id FROM marketplace_categories WHERE slug = 'deportes-fitness';
    SELECT id INTO cat_mascotas_id FROM marketplace_categories WHERE slug = 'mascotas';
    SELECT id INTO cat_salud_id FROM marketplace_categories WHERE slug = 'salud-farmacia';
    SELECT id INTO cat_libros_id FROM marketplace_categories WHERE slug = 'libros-entretenimiento';

    -- NIVEL 1: SUBCATEGORÍAS PRINCIPALES
    
    -- ALIMENTOS Y BEBIDAS
    INSERT INTO marketplace_categories (name, slug, description, parent_id, level, sort_order, is_active) VALUES
    ('Alimentos Frescos', 'alimentos-frescos', 'Carnes, verduras, frutas, productos frescos', cat_alimentos_id, 1, 1, true),
    ('Bebidas', 'bebidas', 'Bebidas alcohólicas, gaseosas, jugos, aguas', cat_alimentos_id, 1, 2, true),
    ('Productos Secos', 'productos-secos', 'Cereales, legumbres, conservas, productos enlatados', cat_alimentos_id, 1, 3, true),
    ('Panadería y Confitería', 'panaderia-confiteria', 'Pan, facturas, tortas, productos de panadería', cat_alimentos_id, 1, 4, true),
    ('Lácteos', 'lacteos', 'Leche, quesos, yogures, manteca, fiambres', cat_alimentos_id, 1, 5, true),
    ('Limpieza', 'limpieza', 'Productos de limpieza del hogar, detergentes, desinfectantes', cat_alimentos_id, 1, 6, true);

    -- MODA E INDUMENTARIA
    INSERT INTO marketplace_categories (name, slug, description, parent_id, level, sort_order, is_active) VALUES
    ('Ropa Mujer', 'ropa-mujer', 'Indumentaria femenina, vestidos, remeras, pantalones', cat_moda_id, 1, 1, true),
    ('Ropa Hombre', 'ropa-hombre', 'Indumentaria masculina, camisas, pantalones, trajes', cat_moda_id, 1, 2, true),
    ('Ropa Infantil', 'ropa-infantil', 'Ropa para bebés, niños y adolescentes', cat_moda_id, 1, 3, true),
    ('Calzado', 'calzado', 'Zapatos, zapatillas, botas para toda la familia', cat_moda_id, 1, 4, true),
    ('Accesorios', 'accesorios-moda', 'Carteras, cinturones, sombreros, bufandas', cat_moda_id, 1, 5, true),
    ('Joyería y Relojes', 'joyeria-relojes', 'Joyas, relojes, bisutería, accesorios personales', cat_moda_id, 1, 6, true);

    -- TECNOLOGÍA Y ELECTRÓNICOS
    INSERT INTO marketplace_categories (name, slug, description, parent_id, level, sort_order, is_active) VALUES
    ('Celulares y Telefonía', 'celulares-telefonia', 'Smartphones, teléfonos, accesorios móviles', cat_tecnologia_id, 1, 1, true),
    ('Computación', 'computacion', 'Computadoras, notebooks, tablets, periféricos', cat_tecnologia_id, 1, 2, true),
    ('Audio e Imagen', 'audio-imagen', 'Equipos de sonido, televisores, cámaras, parlantes', cat_tecnologia_id, 1, 3, true),
    ('Consolas y Videojuegos', 'consolas-videojuegos', 'Consolas, videojuegos, accesorios gaming', cat_tecnologia_id, 1, 4, true),
    ('Accesorios Tecnológicos', 'accesorios-tecnologicos', 'Cables, cargadores, fundas, soportes', cat_tecnologia_id, 1, 5, true);

    -- HOGAR Y DECORACIÓN
    INSERT INTO marketplace_categories (name, slug, description, parent_id, level, sort_order, is_active) VALUES
    ('Muebles', 'muebles', 'Muebles para living, dormitorio, cocina, oficina', cat_hogar_id, 1, 1, true),
    ('Decoración', 'decoracion', 'Cuadros, plantas, velas, objetos decorativos', cat_hogar_id, 1, 2, true),
    ('Textiles del Hogar', 'textiles-hogar', 'Sábanas, toallas, cortinas, alfombras', cat_hogar_id, 1, 3, true),
    ('Cocina y Comedor', 'cocina-comedor', 'Utensilios, vajilla, electrodomésticos pequeños', cat_hogar_id, 1, 4, true),
    ('Jardín y Exterior', 'jardin-exterior', 'Plantas, macetas, muebles de jardín, parrillas', cat_hogar_id, 1, 5, true),
    ('Organización', 'organizacion', 'Contenedores, organizadores, sistemas de almacenamiento', cat_hogar_id, 1, 6, true);

    -- HERRAMIENTAS Y CONSTRUCCIÓN
    INSERT INTO marketplace_categories (name, slug, description, parent_id, level, sort_order, is_active) VALUES
    ('Herramientas Manuales', 'herramientas-manuales', 'Destornilladores, martillos, llaves, herramientas básicas', cat_herramientas_id, 1, 1, true),
    ('Herramientas Eléctricas', 'herramientas-electricas', 'Taladros, amoladoras, sierras eléctricas', cat_herramientas_id, 1, 2, true),
    ('Materiales de Construcción', 'materiales-construccion', 'Cemento, ladrillos, caños, materiales básicos', cat_herramientas_id, 1, 3, true),
    ('Pintura', 'pintura', 'Pinturas, barnices, pinceles, rodillos', cat_herramientas_id, 1, 4, true),
    ('Electricidad', 'electricidad', 'Cables, enchufes, llaves, materiales eléctricos', cat_herramientas_id, 1, 5, true),
    ('Plomería', 'plomeria', 'Caños, grifos, accesorios de plomería', cat_herramientas_id, 1, 6, true);

    -- AUTOMOTRIZ
    INSERT INTO marketplace_categories (name, slug, description, parent_id, level, sort_order, is_active) VALUES
    ('Repuestos', 'repuestos', 'Repuestos originales y alternativos para vehículos', cat_automotriz_id, 1, 1, true),
    ('Neumáticos', 'neumaticos', 'Neumáticos, llantas, válvulas, cámaras', cat_automotriz_id, 1, 2, true),
    ('Aceites y Lubricantes', 'aceites-lubricantes', 'Aceites de motor, lubricantes, fluidos', cat_automotriz_id, 1, 3, true),
    ('Accesorios Exterior', 'accesorios-exterior', 'Espejos, luces, parrillas, emblemas', cat_automotriz_id, 1, 4, true),
    ('Accesorios Interior', 'accesorios-interior', 'Fundas, alfombras, volantes, stereos', cat_automotriz_id, 1, 5, true),
    ('Motos', 'motos', 'Repuestos y accesorios para motocicletas', cat_automotriz_id, 1, 6, true);

    -- BELLEZA Y CUIDADO PERSONAL
    INSERT INTO marketplace_categories (name, slug, description, parent_id, level, sort_order, is_active) VALUES
    ('Cosméticos', 'cosmeticos', 'Maquillaje, bases, labiales, sombras', cat_belleza_id, 1, 1, true),
    ('Cuidado Facial', 'cuidado-facial', 'Cremas, limpiadores, tratamientos faciales', cat_belleza_id, 1, 2, true),
    ('Cuidado Corporal', 'cuidado-corporal', 'Cremas corporales, lociones, aceites', cat_belleza_id, 1, 3, true),
    ('Perfumes', 'perfumes', 'Perfumes, colonias, fragancias', cat_belleza_id, 1, 4, true),
    ('Cuidado Capilar', 'cuidado-capilar', 'Shampoos, acondicionadores, tratamientos', cat_belleza_id, 1, 5, true),
    ('Higiene Personal', 'higiene-personal', 'Jabones, desodorantes, productos de higiene', cat_belleza_id, 1, 6, true);

    -- DEPORTES Y FITNESS
    INSERT INTO marketplace_categories (name, slug, description, parent_id, level, sort_order, is_active) VALUES
    ('Ropa Deportiva', 'ropa-deportiva', 'Indumentaria deportiva, calzado deportivo', cat_deportes_id, 1, 1, true),
    ('Fitness', 'fitness', 'Equipos de gimnasio, pesas, máquinas', cat_deportes_id, 1, 2, true),
    ('Fútbol', 'futbol', 'Pelotas, botines, camisetas, equipamiento', cat_deportes_id, 1, 3, true),
    ('Deportes de Agua', 'deportes-agua', 'Natación, surf, deportes acuáticos', cat_deportes_id, 1, 4, true),
    ('Suplementos', 'suplementos', 'Proteínas, vitaminas, suplementos deportivos', cat_deportes_id, 1, 5, true);

    -- MASCOTAS
    INSERT INTO marketplace_categories (name, slug, description, parent_id, level, sort_order, is_active) VALUES
    ('Alimento para Mascotas', 'alimento-mascotas', 'Alimento balanceado, snacks, premios', cat_mascotas_id, 1, 1, true),
    ('Accesorios para Perros', 'accesorios-perros', 'Collares, correas, camas, juguetes', cat_mascotas_id, 1, 2, true),
    ('Accesorios para Gatos', 'accesorios-gatos', 'Rascadores, juguetes, piedras sanitarias', cat_mascotas_id, 1, 3, true),
    ('Salud Veterinaria', 'salud-veterinaria', 'Medicamentos, productos veterinarios', cat_mascotas_id, 1, 4, true);

    -- SALUD Y FARMACIA
    INSERT INTO marketplace_categories (name, slug, description, parent_id, level, sort_order, is_active) VALUES
    ('Medicamentos', 'medicamentos', 'Medicamentos de venta libre, analgésicos', cat_salud_id, 1, 1, true),
    ('Equipos Médicos', 'equipos-medicos', 'Tensiómetros, termómetros, nebulizadores', cat_salud_id, 1, 2, true),
    ('Primeros Auxilios', 'primeros-auxilios', 'Vendas, gasas, antisépticos, botiquines', cat_salud_id, 1, 3, true),
    ('Ortopedia', 'ortopedia', 'Fajas, plantillas, productos ortopédicos', cat_salud_id, 1, 4, true);

    -- LIBROS Y ENTRETENIMIENTO
    INSERT INTO marketplace_categories (name, slug, description, parent_id, level, sort_order, is_active) VALUES
    ('Libros', 'libros', 'Literatura, educación, técnicos, infantiles', cat_libros_id, 1, 1, true),
    ('Juguetes', 'juguetes', 'Juguetes para todas las edades, didácticos', cat_libros_id, 1, 2, true),
    ('Juegos de Mesa', 'juegos-mesa', 'Juegos familiares, cartas, puzzles', cat_libros_id, 1, 3, true),
    ('Instrumentos Musicales', 'instrumentos-musicales', 'Guitarras, teclados, accesorios musicales', cat_libros_id, 1, 4, true),
    ('Papelería', 'papeleria', 'Útiles escolares, material de oficina', cat_libros_id, 1, 5, true);

END $$;

-- Comentarios
COMMENT ON TABLE marketplace_categories IS 'Categorías globales del marketplace basadas en investigación de mercado argentino 2024-2025 (CACE, Kantar TNS, PCMI)'; 