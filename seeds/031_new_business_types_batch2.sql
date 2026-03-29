-- Seed 031: Nuevos business types — Batch 2 (Agencia de Viajes, Electrodomésticos, Hogar, Jardinería, Ropa de Hogar, Regalería)
-- VERSION: 4.0.0-enriched | IDEMPOTENTE: UPSERT

-- =====================================================
-- PASO 0: Crear business_types nuevos
-- =====================================================
INSERT INTO business_types (code, name, description, icon, color, sort_order, is_active)
VALUES
  ('agencia-viajes', 'Agencia de Viajes', 'Vuelos, paquetes turísticos, alquiler de autos, hoteles, asistencia al viajero y excursiones', 'plane', '#3B82F6', 44, true),
  ('electrodomesticos', 'Electrodomésticos', 'Línea blanca, refrigeración, cocción, TV, audio, calefacción y climatización', 'refrigerator', '#6366F1', 45, true),
  ('hogar', 'Hogar', 'Muebles, electrodomésticos pequeños, jardín, decoración y organización del hogar', 'home', '#8B5CF6', 46, true),
  ('jardineria', 'Jardinería y Máquinas', 'Cortadoras, herramientas de jardín, riego, repuestos, motores e insumos', 'trees', '#22C55E', 47, true),
  ('ropa-hogar', 'Ropa de Hogar', 'Ropa de cama, baño, cortinas, mesa y cocina, mantas y abrigo', 'bed-double', '#EC4899', 48, true),
  ('regaleria', 'Regalería', 'Regalos, bazar, cotillón, papelería y decoración', 'gift', '#F59E0B', 49, true)
ON CONFLICT (code) DO UPDATE SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  icon = EXCLUDED.icon,
  color = EXCLUDED.color,
  is_active = EXCLUDED.is_active;

-- =====================================================
-- 1. AGENCIA DE VIAJES — 6 padres + 17 hijas = 23 categorías
-- =====================================================
INSERT INTO business_type_templates (business_type_id, name, description, version, region, is_default, is_active, generated_by, categories, brands, products)
SELECT bt.id, 'Agencia de Viajes', 'Template de Agencia de Viajes', '4.0.0-enriched', 'AR', true, true, 'manual-curation-v4',
  '[
    {"slug": "aereos", "name": "Aéreos", "level": 0},
    {"slug": "vuelos-nacionales", "name": "Vuelos Nacionales", "parent_slug": "aereos", "level": 1},
    {"slug": "vuelos-internacionales", "name": "Vuelos Internacionales", "parent_slug": "aereos", "level": 1},
    {"slug": "low-cost", "name": "Low Cost", "parent_slug": "aereos", "level": 1},

    {"slug": "paquetes-turisticos", "name": "Paquetes Turísticos", "level": 0},
    {"slug": "paquetes-nacionales", "name": "Nacionales", "parent_slug": "paquetes-turisticos", "level": 1},
    {"slug": "paquetes-internacionales", "name": "Internacionales", "parent_slug": "paquetes-turisticos", "level": 1},
    {"slug": "escapadas", "name": "Escapadas", "parent_slug": "paquetes-turisticos", "level": 1},

    {"slug": "alquiler-autos", "name": "Alquiler de Autos", "level": 0},
    {"slug": "autos-economicos", "name": "Autos Económicos", "parent_slug": "alquiler-autos", "level": 1},
    {"slug": "autos-premium", "name": "Autos Premium", "parent_slug": "alquiler-autos", "level": 1},
    {"slug": "camionetas", "name": "Camionetas", "parent_slug": "alquiler-autos", "level": 1},

    {"slug": "hoteles-alojamiento", "name": "Hoteles y Alojamiento", "level": 0},
    {"slug": "hoteles", "name": "Hoteles", "parent_slug": "hoteles-alojamiento", "level": 1},
    {"slug": "cabanas", "name": "Cabañas", "parent_slug": "hoteles-alojamiento", "level": 1},
    {"slug": "apart-hotels", "name": "Apart Hotels", "parent_slug": "hoteles-alojamiento", "level": 1},

    {"slug": "asistencia-viajero", "name": "Asistencia al Viajero", "level": 0},
    {"slug": "seguros-viaje", "name": "Seguros de Viaje", "parent_slug": "asistencia-viajero", "level": 1},
    {"slug": "visas-tramites", "name": "Visas y Trámites", "parent_slug": "asistencia-viajero", "level": 1},

    {"slug": "cruceros-excursiones", "name": "Cruceros y Excursiones", "level": 0},
    {"slug": "cruceros", "name": "Cruceros", "parent_slug": "cruceros-excursiones", "level": 1},
    {"slug": "excursiones-terrestres", "name": "Excursiones Terrestres", "parent_slug": "cruceros-excursiones", "level": 1},
    {"slug": "aventura", "name": "Aventura", "parent_slug": "cruceros-excursiones", "level": 1}
  ]'::jsonb,
  '[
    {"name": "Aerolíneas Argentinas", "suggested_for_categories": ["vuelos-nacionales", "vuelos-internacionales"]},
    {"name": "LATAM", "suggested_for_categories": ["vuelos-nacionales", "vuelos-internacionales"]},
    {"name": "Flybondi", "suggested_for_categories": ["low-cost", "vuelos-nacionales"]},
    {"name": "JetSMART", "suggested_for_categories": ["low-cost", "vuelos-nacionales"]},
    {"name": "Booking", "suggested_for_categories": ["hoteles", "apart-hotels", "cabanas"]},
    {"name": "Despegar", "suggested_for_categories": ["paquetes-nacionales", "paquetes-internacionales", "hoteles"]},
    {"name": "Assist Card", "suggested_for_categories": ["seguros-viaje"]}
  ]'::jsonb,
  '[]'::jsonb
FROM business_types bt WHERE bt.code = 'agencia-viajes'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories,
  brands = EXCLUDED.brands,
  version = EXCLUDED.version,
  generated_by = EXCLUDED.generated_by,
  updated_at = CURRENT_TIMESTAMP;

-- =====================================================
-- 2. ELECTRODOMÉSTICOS — 7 padres + 22 hijas = 29 categorías
-- =====================================================
INSERT INTO business_type_templates (business_type_id, name, description, version, region, is_default, is_active, generated_by, categories, brands, products)
SELECT bt.id, 'Electrodomésticos', 'Template de Electrodomésticos', '4.0.0-enriched', 'AR', true, true, 'manual-curation-v4',
  '[
    {"slug": "linea-blanca", "name": "Línea Blanca", "level": 0},
    {"slug": "lavarropas", "name": "Lavarropas", "parent_slug": "linea-blanca", "level": 1},
    {"slug": "secarropas", "name": "Secarropas", "parent_slug": "linea-blanca", "level": 1},
    {"slug": "lavavajillas", "name": "Lavavajillas", "parent_slug": "linea-blanca", "level": 1},

    {"slug": "refrigeracion", "name": "Refrigeración", "level": 0},
    {"slug": "heladeras", "name": "Heladeras", "parent_slug": "refrigeracion", "level": 1},
    {"slug": "freezers", "name": "Freezers", "parent_slug": "refrigeracion", "level": 1},
    {"slug": "cavas-vino", "name": "Cavas de Vino", "parent_slug": "refrigeracion", "level": 1},

    {"slug": "coccion", "name": "Cocción", "level": 0},
    {"slug": "cocinas", "name": "Cocinas", "parent_slug": "coccion", "level": 1},
    {"slug": "hornos", "name": "Hornos", "parent_slug": "coccion", "level": 1},
    {"slug": "anafes", "name": "Anafes", "parent_slug": "coccion", "level": 1},
    {"slug": "microondas", "name": "Microondas", "parent_slug": "coccion", "level": 1},

    {"slug": "tv-video", "name": "TV y Video", "level": 0},
    {"slug": "televisores-led-smart", "name": "Televisores LED/Smart", "parent_slug": "tv-video", "level": 1},
    {"slug": "soportes-accesorios", "name": "Soportes y Accesorios", "parent_slug": "tv-video", "level": 1},

    {"slug": "audio", "name": "Audio", "level": 0},
    {"slug": "parlantes-bluetooth", "name": "Parlantes Bluetooth", "parent_slug": "audio", "level": 1},
    {"slug": "barras-sonido", "name": "Barras de Sonido", "parent_slug": "audio", "level": 1},
    {"slug": "home-theater", "name": "Home Theater", "parent_slug": "audio", "level": 1},

    {"slug": "calefaccion", "name": "Calefacción", "level": 0},
    {"slug": "estufas", "name": "Estufas", "parent_slug": "calefaccion", "level": 1},
    {"slug": "calefactores", "name": "Calefactores", "parent_slug": "calefaccion", "level": 1},
    {"slug": "caloventores", "name": "Caloventores", "parent_slug": "calefaccion", "level": 1},

    {"slug": "climatizacion", "name": "Climatización", "level": 0},
    {"slug": "aires-acondicionados", "name": "Aires Acondicionados", "parent_slug": "climatizacion", "level": 1},
    {"slug": "ventiladores", "name": "Ventiladores", "parent_slug": "climatizacion", "level": 1},
    {"slug": "purificadores", "name": "Purificadores", "parent_slug": "climatizacion", "level": 1}
  ]'::jsonb,
  '[
    {"name": "Samsung", "suggested_for_categories": ["televisores-led-smart", "heladeras", "lavarropas", "microondas"]},
    {"name": "LG", "suggested_for_categories": ["televisores-led-smart", "heladeras", "lavarropas", "aires-acondicionados"]},
    {"name": "Whirlpool", "suggested_for_categories": ["heladeras", "lavarropas", "lavavajillas", "cocinas"]},
    {"name": "Drean", "suggested_for_categories": ["lavarropas", "secarropas", "cocinas"]},
    {"name": "BGH", "suggested_for_categories": ["aires-acondicionados", "heladeras", "microondas"]},
    {"name": "Philco", "suggested_for_categories": ["televisores-led-smart", "heladeras", "cocinas"]},
    {"name": "Electrolux", "suggested_for_categories": ["lavarropas", "heladeras", "lavavajillas"]},
    {"name": "Atma", "suggested_for_categories": ["ventiladores", "calefactores", "caloventores"]}
  ]'::jsonb,
  '[]'::jsonb
FROM business_types bt WHERE bt.code = 'electrodomesticos'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories,
  brands = EXCLUDED.brands,
  version = EXCLUDED.version,
  generated_by = EXCLUDED.generated_by,
  updated_at = CURRENT_TIMESTAMP;

-- =====================================================
-- 3. HOGAR — 5 padres + 16 hijas = 21 categorías
-- =====================================================
INSERT INTO business_type_templates (business_type_id, name, description, version, region, is_default, is_active, generated_by, categories, brands, products)
SELECT bt.id, 'Hogar', 'Template de Hogar', '4.0.0-enriched', 'AR', true, true, 'manual-curation-v4',
  '[
    {"slug": "muebles", "name": "Muebles", "level": 0},
    {"slug": "living", "name": "Living", "parent_slug": "muebles", "level": 1},
    {"slug": "dormitorio", "name": "Dormitorio", "parent_slug": "muebles", "level": 1},
    {"slug": "cocina-comedor", "name": "Cocina y Comedor", "parent_slug": "muebles", "level": 1},
    {"slug": "infantil", "name": "Infantil", "parent_slug": "muebles", "level": 1},

    {"slug": "electrodomesticos-hogar", "name": "Electrodomésticos", "level": 0},
    {"slug": "pequenos-electrodomesticos", "name": "Pequeños Electrodomésticos", "parent_slug": "electrodomesticos-hogar", "level": 1},
    {"slug": "linea-blanca-hogar", "name": "Línea Blanca", "parent_slug": "electrodomesticos-hogar", "level": 1},

    {"slug": "jardin", "name": "Jardín", "level": 0},
    {"slug": "muebles-jardin", "name": "Muebles de Jardín", "parent_slug": "jardin", "level": 1},
    {"slug": "macetas", "name": "Macetas", "parent_slug": "jardin", "level": 1},
    {"slug": "riego-jardin", "name": "Riego", "parent_slug": "jardin", "level": 1},

    {"slug": "decoracion", "name": "Decoración", "level": 0},
    {"slug": "cuadros", "name": "Cuadros", "parent_slug": "decoracion", "level": 1},
    {"slug": "espejos", "name": "Espejos", "parent_slug": "decoracion", "level": 1},
    {"slug": "iluminacion-deco", "name": "Iluminación Deco", "parent_slug": "decoracion", "level": 1},

    {"slug": "organizacion", "name": "Organización", "level": 0},
    {"slug": "cajas", "name": "Cajas", "parent_slug": "organizacion", "level": 1},
    {"slug": "estanterias", "name": "Estanterías", "parent_slug": "organizacion", "level": 1},
    {"slug": "closets", "name": "Closets", "parent_slug": "organizacion", "level": 1}
  ]'::jsonb,
  '[
    {"name": "Botánica Urbana", "suggested_for_categories": ["macetas", "muebles-jardin"]},
    {"name": "TST", "suggested_for_categories": ["pequenos-electrodomesticos"]},
    {"name": "Philips", "suggested_for_categories": ["iluminacion-deco", "pequenos-electrodomesticos"]},
    {"name": "Tramontina", "suggested_for_categories": ["muebles-jardin", "cocina-comedor"]},
    {"name": "Krea", "suggested_for_categories": ["living", "dormitorio", "organizacion"]}
  ]'::jsonb,
  '[]'::jsonb
FROM business_types bt WHERE bt.code = 'hogar'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories,
  brands = EXCLUDED.brands,
  version = EXCLUDED.version,
  generated_by = EXCLUDED.generated_by,
  updated_at = CURRENT_TIMESTAMP;

-- =====================================================
-- 4. JARDINERÍA Y MÁQUINAS — 6 padres + 20 hijas = 26 categorías
-- =====================================================
INSERT INTO business_type_templates (business_type_id, name, description, version, region, is_default, is_active, generated_by, categories, brands, products)
SELECT bt.id, 'Jardinería y Máquinas', 'Template de Jardinería y Máquinas', '4.0.0-enriched', 'AR', true, true, 'manual-curation-v4',
  '[
    {"slug": "cortadoras-podadoras", "name": "Cortadoras y Podadoras", "level": 0},
    {"slug": "cortadoras-cesped", "name": "Cortadoras de Césped", "parent_slug": "cortadoras-podadoras", "level": 1},
    {"slug": "desmalezadoras", "name": "Desmalezadoras", "parent_slug": "cortadoras-podadoras", "level": 1},
    {"slug": "podadoras", "name": "Podadoras", "parent_slug": "cortadoras-podadoras", "level": 1},

    {"slug": "herramientas-jardin", "name": "Herramientas de Jardín", "level": 0},
    {"slug": "palas", "name": "Palas", "parent_slug": "herramientas-jardin", "level": 1},
    {"slug": "rastrillos", "name": "Rastrillos", "parent_slug": "herramientas-jardin", "level": 1},
    {"slug": "tijeras-podar", "name": "Tijeras de Podar", "parent_slug": "herramientas-jardin", "level": 1},
    {"slug": "mangueras", "name": "Mangueras", "parent_slug": "herramientas-jardin", "level": 1},

    {"slug": "riego", "name": "Riego", "level": 0},
    {"slug": "aspersores", "name": "Aspersores", "parent_slug": "riego", "level": 1},
    {"slug": "goteo", "name": "Goteo", "parent_slug": "riego", "level": 1},
    {"slug": "programadores", "name": "Programadores", "parent_slug": "riego", "level": 1},

    {"slug": "repuestos", "name": "Repuestos", "level": 0},
    {"slug": "cuchillas", "name": "Cuchillas", "parent_slug": "repuestos", "level": 1},
    {"slug": "filtros", "name": "Filtros", "parent_slug": "repuestos", "level": 1},
    {"slug": "bujias", "name": "Bujías", "parent_slug": "repuestos", "level": 1},
    {"slug": "correas", "name": "Correas", "parent_slug": "repuestos", "level": 1},

    {"slug": "motores-bombas", "name": "Motores y Bombas", "level": 0},
    {"slug": "motobombas", "name": "Motobombas", "parent_slug": "motores-bombas", "level": 1},
    {"slug": "grupos-electrogenos", "name": "Grupos Electrógenos", "parent_slug": "motores-bombas", "level": 1},
    {"slug": "motosierras", "name": "Motosierras", "parent_slug": "motores-bombas", "level": 1},

    {"slug": "insumos", "name": "Insumos", "level": 0},
    {"slug": "fertilizantes", "name": "Fertilizantes", "parent_slug": "insumos", "level": 1},
    {"slug": "tierra", "name": "Tierra", "parent_slug": "insumos", "level": 1},
    {"slug": "sustratos", "name": "Sustratos", "parent_slug": "insumos", "level": 1},
    {"slug": "semillas", "name": "Semillas", "parent_slug": "insumos", "level": 1}
  ]'::jsonb,
  '[
    {"name": "STIHL", "suggested_for_categories": ["cortadoras-cesped", "desmalezadoras", "motosierras"]},
    {"name": "Honda", "suggested_for_categories": ["cortadoras-cesped", "motobombas", "grupos-electrogenos"]},
    {"name": "Husqvarna", "suggested_for_categories": ["cortadoras-cesped", "desmalezadoras", "motosierras"]},
    {"name": "Niwa", "suggested_for_categories": ["cortadoras-cesped", "desmalezadoras", "grupos-electrogenos"]},
    {"name": "Gamma", "suggested_for_categories": ["cortadoras-cesped", "desmalezadoras", "motobombas"]},
    {"name": "FMT", "suggested_for_categories": ["cortadoras-cesped", "motobombas", "grupos-electrogenos"]}
  ]'::jsonb,
  '[]'::jsonb
FROM business_types bt WHERE bt.code = 'jardineria'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories,
  brands = EXCLUDED.brands,
  version = EXCLUDED.version,
  generated_by = EXCLUDED.generated_by,
  updated_at = CURRENT_TIMESTAMP;

-- =====================================================
-- 5. ROPA DE HOGAR — 5 padres + 17 hijas = 22 categorías
-- =====================================================
INSERT INTO business_type_templates (business_type_id, name, description, version, region, is_default, is_active, generated_by, categories, brands, products)
SELECT bt.id, 'Ropa de Hogar', 'Template de Ropa de Hogar', '4.0.0-enriched', 'AR', true, true, 'manual-curation-v4',
  '[
    {"slug": "ropa-cama", "name": "Ropa de Cama", "level": 0},
    {"slug": "sabanas", "name": "Sábanas", "parent_slug": "ropa-cama", "level": 1},
    {"slug": "acolchados", "name": "Acolchados", "parent_slug": "ropa-cama", "level": 1},
    {"slug": "cubrecamas", "name": "Cubrecamas", "parent_slug": "ropa-cama", "level": 1},
    {"slug": "almohadas", "name": "Almohadas", "parent_slug": "ropa-cama", "level": 1},

    {"slug": "bano", "name": "Baño", "level": 0},
    {"slug": "toallas", "name": "Toallas", "parent_slug": "bano", "level": 1},
    {"slug": "toallones", "name": "Toallones", "parent_slug": "bano", "level": 1},
    {"slug": "cortinas-bano", "name": "Cortinas de Baño", "parent_slug": "bano", "level": 1},
    {"slug": "alfombras-bano", "name": "Alfombras de Baño", "parent_slug": "bano", "level": 1},

    {"slug": "cortinas", "name": "Cortinas", "level": 0},
    {"slug": "cortinas-roller", "name": "Cortinas Roller", "parent_slug": "cortinas", "level": 1},
    {"slug": "cortinas-tela", "name": "Cortinas de Tela", "parent_slug": "cortinas", "level": 1},
    {"slug": "barras-accesorios", "name": "Barras y Accesorios", "parent_slug": "cortinas", "level": 1},

    {"slug": "mesa-cocina", "name": "Mesa y Cocina", "level": 0},
    {"slug": "manteles", "name": "Manteles", "parent_slug": "mesa-cocina", "level": 1},
    {"slug": "individuales", "name": "Individuales", "parent_slug": "mesa-cocina", "level": 1},
    {"slug": "repasadores", "name": "Repasadores", "parent_slug": "mesa-cocina", "level": 1},

    {"slug": "abrigo", "name": "Abrigo", "level": 0},
    {"slug": "mantas", "name": "Mantas", "parent_slug": "abrigo", "level": 1},
    {"slug": "frazadas", "name": "Frazadas", "parent_slug": "abrigo", "level": 1},
    {"slug": "edredones", "name": "Edredones", "parent_slug": "abrigo", "level": 1}
  ]'::jsonb,
  '[
    {"name": "Arredo", "suggested_for_categories": ["sabanas", "acolchados", "toallas", "toallones"]},
    {"name": "Cannon", "suggested_for_categories": ["toallas", "toallones", "sabanas"]},
    {"name": "Danubio", "suggested_for_categories": ["sabanas", "acolchados", "frazadas"]},
    {"name": "Palette", "suggested_for_categories": ["sabanas", "almohadas", "acolchados"]},
    {"name": "Pierre Cardin", "suggested_for_categories": ["acolchados", "cubrecamas", "edredones"]},
    {"name": "Espalma", "suggested_for_categories": ["toallas", "toallones", "alfombras-bano"]}
  ]'::jsonb,
  '[]'::jsonb
FROM business_types bt WHERE bt.code = 'ropa-hogar'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories,
  brands = EXCLUDED.brands,
  version = EXCLUDED.version,
  generated_by = EXCLUDED.generated_by,
  updated_at = CURRENT_TIMESTAMP;

-- =====================================================
-- 6. REGALERÍA — 5 padres + 17 hijas = 22 categorías
-- =====================================================
INSERT INTO business_type_templates (business_type_id, name, description, version, region, is_default, is_active, generated_by, categories, brands, products)
SELECT bt.id, 'Regalería', 'Template de Regalería', '4.0.0-enriched', 'AR', true, true, 'manual-curation-v4',
  '[
    {"slug": "regalos", "name": "Regalos", "level": 0},
    {"slug": "souvenirs", "name": "Souvenirs", "parent_slug": "regalos", "level": 1},
    {"slug": "portarretratos", "name": "Portarretratos", "parent_slug": "regalos", "level": 1},
    {"slug": "velas-aromaticas", "name": "Velas Aromáticas", "parent_slug": "regalos", "level": 1},
    {"slug": "difusores", "name": "Difusores", "parent_slug": "regalos", "level": 1},

    {"slug": "bazar", "name": "Bazar", "level": 0},
    {"slug": "vajilla", "name": "Vajilla", "parent_slug": "bazar", "level": 1},
    {"slug": "vasos", "name": "Vasos", "parent_slug": "bazar", "level": 1},
    {"slug": "tazas", "name": "Tazas", "parent_slug": "bazar", "level": 1},
    {"slug": "termos", "name": "Termos", "parent_slug": "bazar", "level": 1},

    {"slug": "cotillon", "name": "Cotillón", "level": 0},
    {"slug": "globos", "name": "Globos", "parent_slug": "cotillon", "level": 1},
    {"slug": "guirnaldas", "name": "Guirnaldas", "parent_slug": "cotillon", "level": 1},
    {"slug": "descartables", "name": "Descartables", "parent_slug": "cotillon", "level": 1},
    {"slug": "pinatas", "name": "Piñatas", "parent_slug": "cotillon", "level": 1},

    {"slug": "papeleria", "name": "Papelería", "level": 0},
    {"slug": "tarjetas", "name": "Tarjetas", "parent_slug": "papeleria", "level": 1},
    {"slug": "bolsas-regalo", "name": "Bolsas de Regalo", "parent_slug": "papeleria", "level": 1},
    {"slug": "papel-envolver", "name": "Papel de Envolver", "parent_slug": "papeleria", "level": 1},

    {"slug": "deco", "name": "Deco", "level": 0},
    {"slug": "figuras-decorativas", "name": "Figuras Decorativas", "parent_slug": "deco", "level": 1},
    {"slug": "porta-velas", "name": "Porta Velas", "parent_slug": "deco", "level": 1},
    {"slug": "cuadros-pequenos", "name": "Cuadros Pequeños", "parent_slug": "deco", "level": 1}
  ]'::jsonb,
  '[
    {"name": "Royal Norfolk", "suggested_for_categories": ["vajilla", "vasos", "tazas"]},
    {"name": "Tramontina", "suggested_for_categories": ["vajilla", "termos"]},
    {"name": "Stanley", "suggested_for_categories": ["termos", "vasos"]}
  ]'::jsonb,
  '[]'::jsonb
FROM business_types bt WHERE bt.code = 'regaleria'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories,
  brands = EXCLUDED.brands,
  version = EXCLUDED.version,
  generated_by = EXCLUDED.generated_by,
  updated_at = CURRENT_TIMESTAMP;
