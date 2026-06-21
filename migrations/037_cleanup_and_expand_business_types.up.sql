-- Migration: Cleanup generic business types, add NEA-specific rubros
-- Removes 8 generic non-NEA types (no templates, no real data)
-- Adds 7 missing types that have global_products seeds

-- ─── DELETE generic types (no templates, not NEA-relevant) ───────────────────
DELETE FROM business_types WHERE code IN (
  'retail', 'restaurant', 'services', 'wholesale',
  'manufacturing', 'healthcare', 'education', 'technology'
);

-- ─── ADD missing NEA rubros from seeds ───────────────────────────────────────
INSERT INTO business_types (id, code, name, description, icon, color, is_active, sort_order, created_at, updated_at)
VALUES
  (gen_random_uuid(), 'bazar',           'Bazar',                     'Artículos para el hogar, decoración y regalos',      'shopping_basket', '#F97316', true, 130, NOW(), NOW()),
  (gen_random_uuid(), 'fiambreria',      'Fiambrería y Rotisería',    'Fiambres, quesos, platos preparados y rotisería',    'set_meal',        '#BE185D', true, 140, NOW(), NOW()),
  (gen_random_uuid(), 'jugueteria',      'Juguetería',                'Juguetes, juegos y artículos infantiles',             'toys',            '#7C3AED', true, 150, NOW(), NOW()),
  (gen_random_uuid(), 'libreria',        'Librería y Papelería',      'Útiles escolares, libros y artículos de oficina',    'menu_book',       '#1D4ED8', true, 160, NOW(), NOW()),
  (gen_random_uuid(), 'ropa',            'Ropa y Calzado',            'Indumentaria, calzado y accesorios de moda',          'checkroom',       '#6D28D9', true, 170, NOW(), NOW()),
  (gen_random_uuid(), 'veterinaria',     'Veterinaria y Mascotas',    'Alimentos y accesorios para mascotas, veterinaria',  'pets',            '#0F766E', true, 180, NOW(), NOW()),
  (gen_random_uuid(), 'electrodomesticos','Electrodomésticos',        'Pequeños y grandes electrodomésticos para el hogar', 'electrical_services', '#374151', true, 190, NOW(), NOW())
ON CONFLICT (code) DO NOTHING;

-- ─── FIX sort_orders: templates-first order ──────────────────────────────────
UPDATE business_types SET sort_order = 10  WHERE code = 'almacen';
UPDATE business_types SET sort_order = 20  WHERE code = 'kiosco';
UPDATE business_types SET sort_order = 30  WHERE code = 'ferreteria';
UPDATE business_types SET sort_order = 40  WHERE code = 'panaderia';
UPDATE business_types SET sort_order = 50  WHERE code = 'verduleria';
UPDATE business_types SET sort_order = 60  WHERE code = 'carniceria';
UPDATE business_types SET sort_order = 70  WHERE code = 'farmacia';
UPDATE business_types SET sort_order = 80  WHERE code = 'vinoteca';
UPDATE business_types SET sort_order = 90  WHERE code = 'peluqueria';
UPDATE business_types SET sort_order = 100 WHERE code = 'limpieza';
UPDATE business_types SET sort_order = 110 WHERE code = 'piletas';
UPDATE business_types SET sort_order = 120 WHERE code = 'perfumeria';
UPDATE business_types SET sort_order = 125 WHERE code = 'electricidad';
UPDATE business_types SET sort_order = 130 WHERE code = 'bazar';
UPDATE business_types SET sort_order = 140 WHERE code = 'fiambreria';
UPDATE business_types SET sort_order = 150 WHERE code = 'jugueteria';
UPDATE business_types SET sort_order = 160 WHERE code = 'libreria';
UPDATE business_types SET sort_order = 170 WHERE code = 'ropa';
UPDATE business_types SET sort_order = 180 WHERE code = 'veterinaria';
UPDATE business_types SET sort_order = 190 WHERE code = 'electrodomesticos';

-- Track migration
