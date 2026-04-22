-- Migration 011: Seed marketplace attributes
-- Reescrito para usar UUIDs generados (no string IDs) + subqueries por slug

INSERT INTO marketplace_attributes (name, slug, type, is_filterable, is_searchable, is_required_for_listing, sort_order) VALUES
('Color',                'color',              'select', true,  true,  false, 1),
('Talle',                'size',               'select', true,  true,  false, 2),
('Marca',                'brand',              'text',   true,  true,  false, 3),
('Material',             'material',           'select', true,  false, false, 4),
('Género',               'gender',             'select', true,  true,  false, 5),
('Talle de Calzado',     'shoe-size',          'select', true,  true,  false, 10),
('Tipo de Prenda',       'clothing-type',      'select', true,  true,  false, 11),
('Tamaño de Pantalla',   'screen-size',        'select', true,  false, false, 20),
('Almacenamiento',       'storage',            'select', true,  false, false, 21),
('Memoria RAM',          'ram',                'select', true,  false, false, 22),
('Procesador',           'processor',          'text',   false, true,  false, 23),
('Ambiente',             'room',               'select', true,  false, false, 30),
('Material de Mueble',   'furniture-material', 'select', true,  false, false, 31),
('Potencia',             'power',              'text',   false, false, false, 32),
('Tipo de Deporte',      'sport-type',         'select', true,  true,  false, 40),
('Grupo de Edad',        'age-group',          'select', true,  false, false, 41),
('Tipo de Piel',         'skin-type',          'select', true,  false, false, 50),
('Tipo de Cabello',      'hair-type',          'select', true,  false, false, 51),
('Peso',                 'weight',             'text',   false, false, false, 60),
('Volumen',              'volume',             'text',   false, false, false, 61),
('Fecha de Vencimiento', 'expiry',             'text',   false, false, false, 62),
('Condición',            'condition',          'select', true,  false, false, 90),
('Origen',               'origin',             'select', false, false, false, 91),
('Garantía',             'warranty',           'text',   false, false, false, 92)
ON CONFLICT (slug) DO NOTHING;

-- Valores para Color
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order)
SELECT a.id, v.value, v.slug, v.sort_order
FROM marketplace_attributes a
CROSS JOIN (VALUES
  ('Negro','negro',1),('Blanco','blanco',2),('Gris','gris',3),
  ('Gris Claro','gris-claro',4),('Gris Oscuro','gris-oscuro',5),
  ('Rojo','rojo',6),('Azul','azul',7),('Azul Marino','azul-marino',8),
  ('Azul Claro','azul-claro',9),('Verde','verde',10),
  ('Verde Militar','verde-militar',11),('Amarillo','amarillo',12),
  ('Naranja','naranja',13),('Rosa','rosa',14),('Rosa Palo','rosa-palo',15),
  ('Morado','morado',16),('Marrón','marron',17),('Beige','beige',18),
  ('Dorado','dorado',19),('Plateado','plateado',20),
  ('Multicolor','multicolor',21),('Estampado','estampado',22)
) AS v(value, slug, sort_order)
WHERE a.slug = 'color'
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Valores para Talle
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order)
SELECT a.id, v.value, v.slug, v.sort_order
FROM marketplace_attributes a
CROSS JOIN (VALUES
  ('XS','xs',1),('S','s',2),('M','m',3),('L','l',4),('XL','xl',5),
  ('XXL','xxl',6),('XXXL','xxxl',7),('Único','unico',8),
  ('1','1',10),('2','2',11),('3','3',12),('4','4',13),
  ('6','6',14),('8','8',15),('10','10',16),('12','12',17),('14','14',18)
) AS v(value, slug, sort_order)
WHERE a.slug = 'size'
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Valores para Talle de Calzado
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order)
SELECT a.id, v.value, v.slug, v.sort_order
FROM marketplace_attributes a
CROSS JOIN (VALUES
  ('35','35',35),('36','36',36),('37','37',37),('38','38',38),
  ('39','39',39),('40','40',40),('41','41',41),('42','42',42),
  ('43','43',43),('44','44',44),('45','45',45),('46','46',46)
) AS v(value, slug, sort_order)
WHERE a.slug = 'shoe-size'
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Valores para Material
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order)
SELECT a.id, v.value, v.slug, v.sort_order
FROM marketplace_attributes a
CROSS JOIN (VALUES
  ('Algodón','algodon',1),('Algodón Orgánico','algodon-organico',2),
  ('Poliéster','poliester',3),('Lycra','lycra',4),('Lana','lana',5),
  ('Cuero','cuero',6),('Cuero Sintético','cuero-sintetico',7),
  ('Denim','denim',8),('Seda','seda',9),('Lino','lino',10)
) AS v(value, slug, sort_order)
WHERE a.slug = 'material'
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Valores para Género
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order)
SELECT a.id, v.value, v.slug, v.sort_order
FROM marketplace_attributes a
CROSS JOIN (VALUES
  ('Mujer','mujer',1),('Hombre','hombre',2),('Unisex','unisex',3),
  ('Niña','nina',4),('Niño','nino',5),('Bebé','bebe',6)
) AS v(value, slug, sort_order)
WHERE a.slug = 'gender'
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Valores para Almacenamiento
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order)
SELECT a.id, v.value, v.slug, v.sort_order
FROM marketplace_attributes a
CROSS JOIN (VALUES
  ('16GB','16gb',1),('32GB','32gb',2),('64GB','64gb',3),('128GB','128gb',4),
  ('256GB','256gb',5),('512GB','512gb',6),('1TB','1tb',7),('2TB','2tb',8)
) AS v(value, slug, sort_order)
WHERE a.slug = 'storage'
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Valores para RAM
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order)
SELECT a.id, v.value, v.slug, v.sort_order
FROM marketplace_attributes a
CROSS JOIN (VALUES
  ('2GB','2gb',1),('4GB','4gb',2),('6GB','6gb',3),('8GB','8gb',4),
  ('12GB','12gb',5),('16GB','16gb',6),('32GB','32gb',7)
) AS v(value, slug, sort_order)
WHERE a.slug = 'ram'
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Valores para Condición
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order)
SELECT a.id, v.value, v.slug, v.sort_order
FROM marketplace_attributes a
CROSS JOIN (VALUES
  ('Nuevo','nuevo',1),('Usado - Excelente','usado-excelente',2),
  ('Usado - Muy Bueno','usado-muy-bueno',3),('Usado - Bueno','usado-bueno',4),
  ('Usado - Aceptable','usado-aceptable',5)
) AS v(value, slug, sort_order)
WHERE a.slug = 'condition'
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Valores para Ambiente
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order)
SELECT a.id, v.value, v.slug, v.sort_order
FROM marketplace_attributes a
CROSS JOIN (VALUES
  ('Living','living',1),('Comedor','comedor',2),('Dormitorio','dormitorio',3),
  ('Cocina','cocina',4),('Baño','bano',5),('Jardín','jardin',6),
  ('Oficina','oficina',7),('Balcón','balcon',8)
) AS v(value, slug, sort_order)
WHERE a.slug = 'room'
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Valores para Grupo de Edad
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order)
SELECT a.id, v.value, v.slug, v.sort_order
FROM marketplace_attributes a
CROSS JOIN (VALUES
  ('Bebé (0-2 años)','bebe',1),('Niño (3-12 años)','nino',2),
  ('Adolescente (13-17 años)','adolescente',3),('Adulto (18-64 años)','adulto',4),
  ('Adulto Mayor (65+ años)','adulto-mayor',5)
) AS v(value, slug, sort_order)
WHERE a.slug = 'age-group'
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Relaciones categorías-atributos (usando subquery por slug)
-- Moda y Accesorios (d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order)
SELECT 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f60'::uuid, a.id, false, v.sort_order
FROM marketplace_attributes a
JOIN (VALUES ('color',1),('size',2),('brand',3),('material',4),('gender',5)) AS v(slug, sort_order)
  ON a.slug = v.slug
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Calzado (d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order)
SELECT 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f72'::uuid, a.id, false, v.sort_order
FROM marketplace_attributes a
JOIN (VALUES ('shoe-size',1),('color',2),('brand',3),('material',4)) AS v(slug, sort_order)
  ON a.slug = v.slug
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Tecnología (d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f62)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order)
SELECT 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f62'::uuid, a.id, false, v.sort_order
FROM marketplace_attributes a
JOIN (VALUES ('brand',1),('condition',2)) AS v(slug, sort_order)
  ON a.slug = v.slug
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Celulares y Tablets (d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f80)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order)
SELECT 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f80'::uuid, a.id, false, v.sort_order
FROM marketplace_attributes a
JOIN (VALUES ('storage',1),('ram',2),('screen-size',3),('color',4)) AS v(slug, sort_order)
  ON a.slug = v.slug
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Hogar y Jardín (d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order)
SELECT 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f61'::uuid, a.id, false, v.sort_order
FROM marketplace_attributes a
JOIN (VALUES ('room',1),('brand',2),('condition',3)) AS v(slug, sort_order)
  ON a.slug = v.slug
ON CONFLICT (category_id, attribute_id) DO NOTHING;

-- Muebles (d1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f75)
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order)
SELECT 'd1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f75'::uuid, a.id, false, v.sort_order
FROM marketplace_attributes a
JOIN (VALUES ('furniture-material',1),('color',2)) AS v(slug, sort_order)
  ON a.slug = v.slug
ON CONFLICT (category_id, attribute_id) DO NOTHING;
