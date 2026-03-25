-- Seed v2 (idempotente): Tipos de comercio minorista argentinos (NO borra)
-- Reemplaza el comportamiento destructivo de seeds/001_business_types_argentina_seed.sql
-- Upsert por (code) usando la constraint UNIQUE business_types_code_key.

INSERT INTO business_types (code, name, description, icon, color, sort_order, is_active)
VALUES
  -- Comercios alimentarios
  ('almacen', 'Almacén de Barrio', 'Almacén tradicional con productos básicos, bebidas, conservas y artículos de primera necesidad', 'shopping-basket', '#4F46E5', 1, true),
  ('supermercado', 'Supermercado', 'Supermercado con amplia variedad de productos alimentarios, bebidas, limpieza y hogar', 'shopping-cart', '#10B981', 2, true),
  ('carniceria', 'Carnicería', 'Carnes frescas, embutidos, fiambres y productos cárnicos elaborados', 'chef-hat', '#EF4444', 3, true),
  ('panaderia', 'Panadería', 'Productos de panadería, facturas, tortas y productos de confitería', 'croissant', '#F59E0B', 4, true),
  ('verduleria', 'Verdulería', 'Frutas y verduras frescas, productos estacionales y orgánicos', 'carrot', '#22C55E', 5, true),
  ('fiambreria', 'Fiambrería', 'Fiambres, quesos, productos gourmet y especialidades delicatessen', 'cheese', '#FF6B35', 6, true),
  ('heladeria', 'Heladería', 'Helados artesanales, postres helados y productos de temporada', 'ice-cream', '#8B5CF6', 7, true),

  -- Farmacias y salud
  ('farmacia', 'Farmacia', 'Medicamentos, productos de cuidado personal, cosméticos y artículos de higiene', 'heart-pulse', '#EC4899', 8, true),
  ('perfumeria', 'Perfumería', 'Perfumes, cosméticos, productos de belleza y cuidado personal', 'sparkles', '#E879F9', 9, true),

  -- Indumentaria y calzado
  ('ropa', 'Tienda de Ropa', 'Indumentaria para hombre, mujer y niños, ropa casual y formal', 'shirt', '#6366F1', 10, true),
  ('zapateria', 'Zapatería', 'Calzado para toda la familia, zapatillas, zapatos formales y botas', 'footprints', '#7C3AED', 11, true),
  ('deportes', 'Artículos Deportivos', 'Ropa deportiva, calzado deportivo, equipamiento y accesorios para deportes', 'medal', '#059669', 12, true),

  -- Hogar y construcción
  ('ferreteria', 'Ferretería', 'Herramientas, materiales de construcción, pintura, electricidad y plomería', 'hammer', '#DC2626', 13, true),
  ('muebleria', 'Mueblería', 'Muebles para el hogar, decoración, colchones y electrodomésticos', 'armchair', '#B45309', 14, true),
  ('bazar', 'Bazar', 'Artículos para el hogar, utensilios de cocina, decoración y regalos', 'gift', '#DB2777', 15, true),

  -- Tecnología y comunicaciones
  ('electronica', 'Casa de Electrodomésticos', 'Electrodomésticos grandes y pequeños, equipos de audio y video', 'tv', '#1D4ED8', 16, true),
  ('celulares', 'Casa de Celulares', 'Teléfonos celulares, accesorios, fundas y servicios técnicos', 'smartphone', '#2563EB', 17, true),
  ('computacion', 'Computación', 'Computadoras, notebooks, periféricos, software y accesorios informáticos', 'monitor', '#1E40AF', 18, true),

  -- Automotriz
  ('repuestos', 'Casa de Repuestos', 'Repuestos automotrices, neumáticos, baterías y accesorios para vehículos', 'car', '#374151', 19, true),
  ('lubricentro', 'Lubricentro', 'Cambio de aceite, filtros, service automotriz y productos de mantenimiento', 'fuel', '#6B7280', 20, true),

  -- Servicios especializados
  ('optica', 'Óptica', 'Anteojos, lentes de contacto, anteojos de sol y servicios ópticos', 'glasses', '#4338CA', 21, true),
  ('relojeria', 'Relojería', 'Relojes, joyería, reparación de relojes y accesorios', 'clock', '#7C2D12', 22, true),
  ('libreria', 'Librería', 'Libros, útiles escolares, material didáctico y papelería', 'book', '#0F766E', 23, true),
  ('jugueteria', 'Juguetería', 'Juguetes para todas las edades, juegos didácticos y artículos para bebés', 'toy-brick', '#C2410C', 24, true),

  -- Mascotas
  ('veterinaria', 'Veterinaria', 'Alimentos para mascotas, productos veterinarios, accesorios y medicamentos', 'dog', '#15803D', 25, true),

  -- Servicios diversos
  ('kiosco', 'Kiosco', 'Golosinas, cigarrillos, bebidas, diarios, revistas y productos de impulso', 'lollipop', '#92400E', 26, true),
  ('floreria', 'Florería', 'Flores frescas, plantas, arreglos florales y accesorios para jardín', 'flower', '#059669', 27, true),
  ('polirubro', 'Polirubro', 'Comercio multirrubro con productos diversos: limpieza, bazar, ferretería básica', 'store', '#6B21A8', 28, true)
ON CONFLICT (code) DO UPDATE SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  icon = EXCLUDED.icon,
  color = EXCLUDED.color,
  sort_order = EXCLUDED.sort_order,
  is_active = EXCLUDED.is_active;


