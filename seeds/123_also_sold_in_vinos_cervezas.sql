-- Seed 123: also_sold_in — Vinos y cervezas de vinoteca → también en almacén/kiosco (T020)
-- CICLO: cycle-009-catalog-volume-expansion (T020)
-- FECHA: 2026-04-25
-- FUENTE: observación comercial NEA — vinos de entrada/media gama se venden en almacenes
-- ZONA: Posadas, Misiones (NEA)
-- IDEMPOTENTE: UPDATE idempotente (SET siempre el mismo valor JSONB)
-- CRITERIO: Vinos y espumantes masivos/accesibles que el almacenero también stockea.
--   No incluir vinos premium (Catena Zapata, Achaval Ferrer, Clos de los Siete) — esos
--   solo vinoteca. Sí incluir: Alamos, Gato Negro, Norton, Trapiche, Trivento, Finca Las Moras,
--   Graffigna, Callia, La Linda, Chandon, Bianchi, Baron B, Nieto Senetiner.
-- NOTA: also_sold_in no duplica el producto; indica en qué rubros puede aparecer vía cross-sell.

-- ============================================================
-- VINOS Y ESPUMANTES MASIVOS — vinoteca → también en almacén/kiosco
-- ============================================================
UPDATE global_products
SET also_sold_in = '["almacen", "kiosco"]'
WHERE business_type = 'vinoteca'
  AND category IN ('vinos-tintos', 'vinos-blancos', 'vinos-rosados', 'espumantes')
  AND is_verified = true
  AND (
    name ILIKE '%alamos%'
    OR name ILIKE '%gato negro%'
    OR name ILIKE '%norton%'
    OR name ILIKE '%trapiche%'
    OR name ILIKE '%trivento%'
    OR name ILIKE '%finca las moras%'
    OR name ILIKE '%graffigna%'
    OR name ILIKE '%callia%'
    OR name ILIKE '%la linda%'
    OR name ILIKE '%chandon%'
    OR name ILIKE '%bianchi%'
    OR name ILIKE '%baron b%'
    OR name ILIKE '%nieto senetiner%'
    OR name ILIKE '%don valentín%'
    OR name ILIKE '%gancia%'
  );

-- ============================================================
-- CERVEZAS — almacén → también en kiosco
-- (las cervezas de almacén son las masivas: Quilmes, Brahma, Stella, Heineken, etc.)
-- ============================================================
UPDATE global_products
SET also_sold_in = '["kiosco"]'
WHERE business_type = 'almacen'
  AND category IN ('cervezas-vinos', 'cervezas')
  AND is_verified = true
  AND (
    name ILIKE '%quilmes%'
    OR name ILIKE '%brahma%'
    OR name ILIKE '%stella artois%'
    OR name ILIKE '%heineken%'
    OR name ILIKE '%corona%'
    OR name ILIKE '%norte%'
    OR name ILIKE '%imperial%'
    OR name ILIKE '%schneider%'
    OR name ILIKE '%andes%'
  );
