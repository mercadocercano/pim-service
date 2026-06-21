-- Alta del rubro "congelados" en el catálogo de business_types.
-- Motivo (E17): los productos scrapeados bajo /Congelados/... necesitan un rubro
-- propio en vez de caer en almacen. Idempotente (migrate.sh re-corre todo).

INSERT INTO business_types (id, code, name, description, icon, color, is_active, sort_order, created_at, updated_at)
VALUES
  (gen_random_uuid(), 'congelados', 'Congelados', 'Productos congelados y ultracongelados (helados, hamburguesas, vegetales, etc.)', 'ac_unit', '#0EA5E9', true, 18, NOW(), NOW())
ON CONFLICT (code) DO NOTHING;

