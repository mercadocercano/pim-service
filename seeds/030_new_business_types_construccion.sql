-- Seed 030: Nuevos business types — Corralón, Construcción en Seco
-- VERSION: 4.0.0-enriched | IDEMPOTENTE: UPSERT

-- =====================================================
-- PASO 0: Crear business_types nuevos
-- =====================================================
INSERT INTO business_types (code, name, description, icon, color, sort_order, is_active)
VALUES
  ('corralon', 'Corralón', 'Materiales de construcción, hierros, cementos, áridos, maderas, aberturas y techos', 'building-2', '#78716C', 42, true),
  ('construccion-seco', 'Construcción en Seco', 'Placas de yeso, perfilería, masillas, aislantes térmicos y acústicos, cielorrasos', 'layers', '#A3A3A3', 43, true)
ON CONFLICT (code) DO UPDATE SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  icon = EXCLUDED.icon,
  color = EXCLUDED.color,
  is_active = EXCLUDED.is_active;

-- =====================================================
-- 1. CORRALÓN — 8 padres + 24 hijas = 32 categorías
-- =====================================================
INSERT INTO business_type_templates (business_type_id, name, description, version, region, is_default, is_active, generated_by, categories, brands, products)
SELECT bt.id, 'Corralón', 'Template de Corralón', '4.0.0-enriched', 'AR', true, true, 'manual-curation-v4',
  '[
    {"slug": "cementos-cales", "name": "Cementos y Cales", "level": 0},
    {"slug": "cemento-portland", "name": "Cemento Portland", "parent_slug": "cementos-cales", "level": 1},
    {"slug": "cal-hidraulica", "name": "Cal Hidráulica", "parent_slug": "cementos-cales", "level": 1},
    {"slug": "morteros-premezclados", "name": "Morteros Premezclados", "parent_slug": "cementos-cales", "level": 1},

    {"slug": "aridos-agregados", "name": "Áridos y Agregados", "level": 0},
    {"slug": "arena", "name": "Arena", "parent_slug": "aridos-agregados", "level": 1},
    {"slug": "piedra-granza", "name": "Piedra y Granza", "parent_slug": "aridos-agregados", "level": 1},
    {"slug": "tosca-relleno", "name": "Tosca y Relleno", "parent_slug": "aridos-agregados", "level": 1},

    {"slug": "ladrillos-bloques", "name": "Ladrillos y Bloques", "level": 0},
    {"slug": "ladrillo-comun", "name": "Ladrillo Común", "parent_slug": "ladrillos-bloques", "level": 1},
    {"slug": "ladrillo-hueco", "name": "Ladrillo Hueco", "parent_slug": "ladrillos-bloques", "level": 1},
    {"slug": "bloques-cemento", "name": "Bloques de Cemento", "parent_slug": "ladrillos-bloques", "level": 1},

    {"slug": "hierros-acero", "name": "Hierros y Acero", "level": 0},
    {"slug": "hierro-construccion", "name": "Hierro de Construcción", "parent_slug": "hierros-acero", "level": 1},
    {"slug": "mallas-electrosoldadas", "name": "Mallas Electrosoldadas", "parent_slug": "hierros-acero", "level": 1},
    {"slug": "perfiles-metalicos", "name": "Perfiles Metálicos", "parent_slug": "hierros-acero", "level": 1},

    {"slug": "maderas", "name": "Maderas", "level": 0},
    {"slug": "tirantes-tablas", "name": "Tirantes y Tablas", "parent_slug": "maderas", "level": 1},
    {"slug": "fenolicos-aglomerados", "name": "Fenólicos y Aglomerados", "parent_slug": "maderas", "level": 1},
    {"slug": "machimbre", "name": "Machimbre", "parent_slug": "maderas", "level": 1},

    {"slug": "techos-cubiertas", "name": "Techos y Cubiertas", "level": 0},
    {"slug": "chapas", "name": "Chapas", "parent_slug": "techos-cubiertas", "level": 1},
    {"slug": "tejas", "name": "Tejas", "parent_slug": "techos-cubiertas", "level": 1},
    {"slug": "membranas-aislantes", "name": "Membranas y Aislantes", "parent_slug": "techos-cubiertas", "level": 1},

    {"slug": "aberturas", "name": "Aberturas", "level": 0},
    {"slug": "puertas", "name": "Puertas", "parent_slug": "aberturas", "level": 1},
    {"slug": "ventanas", "name": "Ventanas", "parent_slug": "aberturas", "level": 1},
    {"slug": "portones", "name": "Portones", "parent_slug": "aberturas", "level": 1},

    {"slug": "pisos-revestimientos", "name": "Pisos y Revestimientos", "level": 0},
    {"slug": "ceramicos", "name": "Cerámicos", "parent_slug": "pisos-revestimientos", "level": 1},
    {"slug": "porcelanatos", "name": "Porcelanatos", "parent_slug": "pisos-revestimientos", "level": 1},
    {"slug": "adhesivos-pastina", "name": "Adhesivos y Pastina", "parent_slug": "pisos-revestimientos", "level": 1}
  ]'::jsonb,
  '[
    {"name": "Loma Negra", "suggested_for_categories": ["cemento-portland", "morteros-premezclados"]},
    {"name": "Avellaneda", "suggested_for_categories": ["cemento-portland", "cal-hidraulica"]},
    {"name": "Holcim", "suggested_for_categories": ["cemento-portland"]},
    {"name": "Weber", "suggested_for_categories": ["morteros-premezclados", "adhesivos-pastina"]},
    {"name": "Klaukol", "suggested_for_categories": ["adhesivos-pastina"]},
    {"name": "Acindar", "suggested_for_categories": ["hierro-construccion", "mallas-electrosoldadas"]},
    {"name": "Cintoplom", "suggested_for_categories": ["membranas-aislantes"]},
    {"name": "Aluar", "suggested_for_categories": ["ventanas", "perfiles-metalicos"]}
  ]'::jsonb,
  '[]'::jsonb
FROM business_types bt WHERE bt.code = 'corralon'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories,
  brands = EXCLUDED.brands,
  version = EXCLUDED.version,
  generated_by = EXCLUDED.generated_by,
  updated_at = CURRENT_TIMESTAMP;

-- =====================================================
-- 2. CONSTRUCCIÓN EN SECO — 6 padres + 18 hijas = 24 categorías
-- =====================================================
INSERT INTO business_type_templates (business_type_id, name, description, version, region, is_default, is_active, generated_by, categories, brands, products)
SELECT bt.id, 'Construcción en Seco', 'Template de Construcción en Seco', '4.0.0-enriched', 'AR', true, true, 'manual-curation-v4',
  '[
    {"slug": "placas", "name": "Placas", "level": 0},
    {"slug": "placas-yeso", "name": "Placas de Yeso (Durlock)", "parent_slug": "placas", "level": 1},
    {"slug": "placas-cementicia", "name": "Placas Cementicias", "parent_slug": "placas", "level": 1},
    {"slug": "placas-especiales", "name": "Placas Especiales (RF, RH)", "parent_slug": "placas", "level": 1},

    {"slug": "perfileria", "name": "Perfilería", "level": 0},
    {"slug": "montantes", "name": "Montantes", "parent_slug": "perfileria", "level": 1},
    {"slug": "soleras", "name": "Soleras", "parent_slug": "perfileria", "level": 1},
    {"slug": "omega-furring", "name": "Omega y Furring", "parent_slug": "perfileria", "level": 1},

    {"slug": "fijaciones-tornilleria", "name": "Fijaciones y Tornillería", "level": 0},
    {"slug": "tornillos-drywall", "name": "Tornillos Drywall", "parent_slug": "fijaciones-tornilleria", "level": 1},
    {"slug": "tarugos-anclajes", "name": "Tarugos y Anclajes", "parent_slug": "fijaciones-tornilleria", "level": 1},
    {"slug": "cintas-angulos", "name": "Cintas y Ángulos", "parent_slug": "fijaciones-tornilleria", "level": 1},

    {"slug": "masillas-terminaciones", "name": "Masillas y Terminaciones", "level": 0},
    {"slug": "masilla-secado-rapido", "name": "Masilla Secado Rápido", "parent_slug": "masillas-terminaciones", "level": 1},
    {"slug": "masilla-multiuso", "name": "Masilla Multiuso", "parent_slug": "masillas-terminaciones", "level": 1},
    {"slug": "selladores", "name": "Selladores", "parent_slug": "masillas-terminaciones", "level": 1},

    {"slug": "aislacion", "name": "Aislación", "level": 0},
    {"slug": "lana-vidrio", "name": "Lana de Vidrio", "parent_slug": "aislacion", "level": 1},
    {"slug": "poliestireno", "name": "Poliestireno Expandido", "parent_slug": "aislacion", "level": 1},
    {"slug": "barreras-vapor", "name": "Barreras de Vapor", "parent_slug": "aislacion", "level": 1},

    {"slug": "cielorrasos", "name": "Cielorrasos", "level": 0},
    {"slug": "cielorrasos-desmontables", "name": "Cielorrasos Desmontables", "parent_slug": "cielorrasos", "level": 1},
    {"slug": "cielorrasos-aplicados", "name": "Cielorrasos Aplicados", "parent_slug": "cielorrasos", "level": 1}
  ]'::jsonb,
  '[
    {"name": "Durlock", "suggested_for_categories": ["placas-yeso", "masilla-multiuso"]},
    {"name": "Knauf", "suggested_for_categories": ["placros-yeso", "masilla-secado-rapido"]},
    {"name": "Isover", "suggested_for_categories": ["lana-vidrio", "poliestireno"]},
    {"name": "Barbieri", "suggested_for_categories": ["montantes", "soleras"]},
    {"name": "Superboard", "suggested_for_categories": ["placas-cementicia"]},
    {"name": "Hunter Douglas", "suggested_for_categories": ["cielorrasos-desmontables"]}
  ]'::jsonb,
  '[]'::jsonb
FROM business_types bt WHERE bt.code = 'construccion-seco'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories,
  brands = EXCLUDED.brands,
  version = EXCLUDED.version,
  generated_by = EXCLUDED.generated_by,
  updated_at = CURRENT_TIMESTAMP;
