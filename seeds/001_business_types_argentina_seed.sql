-- Seed: 001_business_types_argentina_seed.sql
-- Purpose: Limpiar y cargar tipos de comercio minorista argentinos reales
-- Date: 2025-01-27

-- Primero limpiamos los datos existentes
DELETE FROM tenant_business_type_setup;
DELETE FROM business_type_templates;
DELETE FROM business_types;

-- Insertamos tipos de comercios físicos argentinos reales
INSERT INTO business_types (code, name, description, icon, color, sort_order) VALUES

-- Comercios alimentarios
('almacen', 'Almacén de Barrio', 'Almacén tradicional con productos básicos, bebidas, conservas y artículos de primera necesidad', 'shopping-basket', '#4F46E5', 1),
('supermercado', 'Supermercado', 'Supermercado con amplia variedad de productos alimentarios, bebidas, limpieza y hogar', 'shopping-cart', '#10B981', 2),
('carniceria', 'Carnicería', 'Carnes frescas, embutidos, fiambres y productos cárnicos elaborados', 'chef-hat', '#EF4444', 3),
('panaderia', 'Panadería', 'Productos de panadería, facturas, tortas y productos de confitería', 'croissant', '#F59E0B', 4),
('verduleria', 'Verdulería', 'Frutas y verduras frescas, productos estacionales y orgánicos', 'carrot', '#22C55E', 5),
('fiambreria', 'Fiambrería', 'Fiambres, quesos, productos gourmet y especialidades delicatessen', 'cheese', '#FF6B35', 6),
('heladeria', 'Heladería', 'Helados artesanales, postres helados y productos de temporada', 'ice-cream', '#8B5CF6', 7),

-- Farmacias y salud
('farmacia', 'Farmacia', 'Medicamentos, productos de cuidado personal, cosméticos y artículos de higiene', 'heart-pulse', '#EC4899', 8),
('perfumeria', 'Perfumería', 'Perfumes, cosméticos, productos de belleza y cuidado personal', 'sparkles', '#E879F9', 9),

-- Indumentaria y calzado
('ropa', 'Tienda de Ropa', 'Indumentaria para hombre, mujer y niños, ropa casual y formal', 'shirt', '#6366F1', 10),
('zapateria', 'Zapatería', 'Calzado para toda la familia, zapatillas, zapatos formales y botas', 'footprints', '#7C3AED', 11),
('deportes', 'Artículos Deportivos', 'Ropa deportiva, calzado deportivo, equipamiento y accesorios para deportes', 'dumbbell', '#059669', 12),

-- Hogar y construcción
('ferreteria', 'Ferretería', 'Herramientas, materiales de construcción, pintura, electricidad y plomería', 'hammer', '#DC2626', 13),
('electricidad', 'Materiales Eléctricos / Iluminación', 'Cables, iluminación, interruptores, tomas y materiales eléctricos', 'zap', '#F59E0B', 14),
('muebleria', 'Mueblería', 'Muebles para el hogar, decoración, colchones y electrodomésticos', 'armchair', '#B45309', 15),
('bazar', 'Bazar', 'Artículos para el hogar, utensilios de cocina, decoración y regalos', 'gift', '#DB2777', 16),

-- Tecnología y comunicaciones
('electronica', 'Casa de Electrodomésticos', 'Electrodomésticos grandes y pequeños, equipos de audio y video', 'tv', '#1D4ED8', 17),
('celulares', 'Casa de Celulares', 'Teléfonos celulares, accesorios, fundas y servicios técnicos', 'smartphone', '#2563EB', 18),
('computacion', 'Computación', 'Computadoras, notebooks, periféricos, software y accesorios informáticos', 'monitor', '#1E40AF', 19),

-- Automotriz
('repuestos', 'Casa de Repuestos', 'Repuestos automotrices, neumáticos, baterías y accesorios para vehículos', 'car', '#374151', 20),
('lubricentro', 'Lubricentro', 'Cambio de aceite, filtros, service automotriz y productos de mantenimiento', 'fuel', '#6B7280', 21),

-- Servicios especializados
('optica', 'Óptica', 'Anteojos, lentes de contacto, anteojos de sol y servicios ópticos', 'glasses', '#4338CA', 22),
('relojeria', 'Relojería', 'Relojes, joyería, reparación de relojes y accesorios', 'clock', '#7C2D12', 23),
('libreria', 'Librería', 'Libros, útiles escolares, material didáctico y papelería', 'book', '#0F766E', 24),
('jugueteria', 'Juguetería', 'Juguetes para todas las edades, juegos didácticos y artículos para bebés', 'gamepad-2', '#C2410C', 25),

-- Mascotas
('veterinaria', 'Veterinaria', 'Alimentos para mascotas, productos veterinarios, accesorios y medicamentos', 'dog', '#15803D', 26),

-- Servicios diversos
('kiosco', 'Kiosco', 'Golosinas, cigarrillos, bebidas, diarios, revistas y productos de impulso', 'newspaper', '#92400E', 27),
('floreria', 'Florería', 'Flores frescas, plantas, arreglos florales y accesorios para jardín', 'flower', '#059669', 28),
('polirubro', 'Polirubro', 'Comercio multirrubro con productos diversos: limpieza, bazar, ferretería básica', 'store', '#6B21A8', 29),
('delivery', 'Delivery de Comidas', 'Entrega de comidas y productos a domicilio', 'truck', '#10B981', 30);

-- Comentarios sobre la tabla
COMMENT ON TABLE business_types IS 'Tipos de comercios físicos argentinos para marketplace SaaS multitenant'; 