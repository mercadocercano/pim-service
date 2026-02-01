-- Seed v2 (idempotente): Extiende marketplace_categories con slugs provenientes de curación (Mongo)
-- NO borra datos existentes. Reutiliza slugs ya presentes; crea los faltantes y ajusta name/description.
--
-- Fuente de slugs: scraper_db.scraped_products.curated_data.category (sincronizado a Postgres en global_products.category)
-- Mapping documental: seeds/curation_category_slug_map.json
--
-- Reglas:
-- - Si el slug ya existe: UPDATE name/description/is_active/sort_order (no toca parent_id para no romper jerarquías ya definidas)
-- - Si no existe: INSERT con parent_id resuelto por parent_slug (si aplica)

DO $$
DECLARE
  -- Parents comunes (pueden no existir en todas las BD; se validan antes de insertar)
  p_alimentos_bebidas UUID;
  p_alimentos_frescos UUID;
  p_productos_secos UUID;
  p_bebidas UUID;
  p_lacteos UUID;
  p_belleza UUID;
  p_bebes_ninos UUID;
  p_snacks_dulces UUID;
  p_pastas_cereales UUID;
BEGIN
  -- Resolver parents por slug (si no existen quedan NULL y el INSERT se hace como raíz)
  SELECT id INTO p_alimentos_bebidas FROM marketplace_categories WHERE slug = 'alimentos-bebidas';
  SELECT id INTO p_alimentos_frescos FROM marketplace_categories WHERE slug = 'alimentos-frescos';
  SELECT id INTO p_productos_secos FROM marketplace_categories WHERE slug = 'productos-secos';
  SELECT id INTO p_bebidas FROM marketplace_categories WHERE slug = 'bebidas';
  SELECT id INTO p_lacteos FROM marketplace_categories WHERE slug = 'lacteos';
  SELECT id INTO p_belleza FROM marketplace_categories WHERE slug = 'belleza-cuidado-personal';
  SELECT id INTO p_bebes_ninos FROM marketplace_categories WHERE slug = 'bebes-ninos';
  SELECT id INTO p_snacks_dulces FROM marketplace_categories WHERE slug = 'snacks-dulces';
  SELECT id INTO p_pastas_cereales FROM marketplace_categories WHERE slug = 'pastas-cereales';

  -- UPSERT helper: insert if missing with parent_id; always update non-hierarchy fields

  -- Root/Top-level (si existen, solo update)
  INSERT INTO marketplace_categories (name, slug, description, parent_id, sort_order, is_active)
  VALUES
    ('Hogar y Jardín', 'hogar-jardin', 'Categoría curada: hogar y jardín', NULL, 2, true),
    ('Tecnología', 'tecnologia', 'Categoría curada: tecnología', NULL, 3, true),
    ('Salud y Belleza', 'salud-belleza', 'Categoría curada: salud y belleza', NULL, 5, true),
    ('Moda y Accesorios', 'moda-accesorios', 'Categoría curada: moda y accesorios', NULL, 1, true),
    ('Bebés y Niños', 'bebes-ninos', 'Categoría curada: bebés y niños', NULL, 8, true)
  ON CONFLICT (slug) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    sort_order = EXCLUDED.sort_order,
    is_active = EXCLUDED.is_active;

  -- Food/CPG
  INSERT INTO marketplace_categories (name, slug, description, parent_id, sort_order, is_active)
  VALUES
    ('Alimentos Envasados', 'alimentos-envasados', 'Productos industrializados/envasados', COALESCE(p_alimentos_bebidas, NULL), 3, true),
    ('Frutas y Verduras', 'frutas-verduras', 'Frutas y verduras frescas', COALESCE(p_alimentos_frescos, p_alimentos_bebidas, NULL), 1, true),
    ('Carnes y Pescados', 'carnes-pescados', 'Carnes, pescados y productos cárnicos', COALESCE(p_alimentos_frescos, p_alimentos_bebidas, NULL), 2, true),
    ('Panificados', 'panificados', 'Pan, facturas y productos de panadería', COALESCE(p_alimentos_frescos, p_alimentos_bebidas, NULL), 4, true),
    ('Pastas y Cereales', 'pastas-cereales', 'Pastas secas, cereales y similares', COALESCE(p_productos_secos, p_alimentos_bebidas, NULL), 5, true),
    ('Conservas', 'conservas', 'Conservas y enlatados', COALESCE(p_productos_secos, p_alimentos_bebidas, NULL), 6, true),
    ('Snacks y Dulces', 'snacks-dulces', 'Snacks y dulces', COALESCE(p_alimentos_bebidas, NULL), 7, true),
    ('Snacks Salados', 'snacks-salados', 'Snacks salados', COALESCE(p_snacks_dulces, p_alimentos_bebidas, NULL), 1, true),
    ('Galletas', 'galletas', 'Galletas y galletitas', COALESCE(p_snacks_dulces, p_alimentos_bebidas, NULL), 2, true),
    ('Golosinas', 'golosinas', 'Golosinas y chocolates', COALESCE(p_snacks_dulces, p_alimentos_bebidas, NULL), 3, true),
    ('Bebidas sin Alcohol', 'bebidas-sin-alcohol', 'Gaseosas, aguas, jugos', COALESCE(p_bebidas, p_alimentos_bebidas, NULL), 1, true),
    ('Bebidas con Alcohol', 'bebidas-con-alcohol', 'Cervezas, vinos, espirituosas', COALESCE(p_bebidas, p_alimentos_bebidas, NULL), 2, true),
    ('Lácteos Líquidos', 'lacteos-liquidos', 'Leches, yogures bebibles', COALESCE(p_lacteos, p_alimentos_bebidas, NULL), 1, true),
    ('Lácteos y Fiambres', 'lacteos-fiambres', 'Quesos, fiambres y afines', COALESCE(p_lacteos, p_alimentos_bebidas, NULL), 2, true),
    ('Cereales de Desayuno', 'cereales-desayuno', 'Cereales, granolas y barras', COALESCE(p_pastas_cereales, p_productos_secos, p_alimentos_bebidas, NULL), 1, true)
  ON CONFLICT (slug) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    sort_order = EXCLUDED.sort_order,
    is_active = EXCLUDED.is_active;

  -- Beauty/Personal care
  INSERT INTO marketplace_categories (name, slug, description, parent_id, sort_order, is_active)
  VALUES
    ('Higiene Personal', 'higiene-personal', 'Jabones, desodorantes, higiene', COALESCE(p_belleza, NULL), 1, true),
    ('Cuidado del Cabello', 'cuidado-cabello', 'Shampoos, acondicionadores, tratamientos', COALESCE(p_belleza, NULL), 2, true),
    ('Cuidado Facial', 'cuidado-facial', 'Cremas y limpieza facial', COALESCE(p_belleza, NULL), 3, true),
    ('Cuidado Corporal', 'cuidado-corporal', 'Cremas corporales y cuidado', COALESCE(p_belleza, NULL), 4, true),
    ('Cuidado Bucal', 'cuidado-bucal', 'Pasta dental, cepillos y cuidado bucal', COALESCE(p_belleza, NULL), 5, true)
  ON CONFLICT (slug) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    sort_order = EXCLUDED.sort_order,
    is_active = EXCLUDED.is_active;

  -- Babies
  INSERT INTO marketplace_categories (name, slug, description, parent_id, sort_order, is_active)
  VALUES
    ('Bebés (0-2 años)', 'bebes-0-2', 'Productos para bebés de 0 a 2 años', COALESCE(p_bebes_ninos, NULL), 1, true),
    ('Alimentación Bebé', 'alimentacion-bebe', 'Fórmulas, papillas, alimentos para bebé', COALESCE(p_bebes_ninos, NULL), 2, true)
  ON CONFLICT (slug) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    sort_order = EXCLUDED.sort_order,
    is_active = EXCLUDED.is_active;

END $$;


