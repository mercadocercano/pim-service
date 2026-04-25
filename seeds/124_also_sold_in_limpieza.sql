-- Seed 124: also_sold_in — Productos de limpieza → también en almacén (T021)
-- CICLO: cycle-009-catalog-volume-expansion (T021)
-- FECHA: 2026-04-25
-- FUENTE: observación comercial NEA — lavandina, detergente y jabón en polvo son ítem
--   habitual en almacenes de barrio de Posadas que no son exclusivamente limpieza.
-- ZONA: Posadas, Misiones (NEA)
-- IDEMPOTENTE: UPDATE idempotente (SET siempre el mismo valor JSONB)
-- CRITERIO: Productos de higiene del hogar de venta masiva que el almacenero stockea.
--   Excluir accesorios especializados (escobas Vileda premium, robots, etc.).
--   Incluir: lavandina, detergente, jabón en polvo, suavizante, desinfectante básico,
--   papel higiénico, servilletas, bolsas de basura básicas.
-- NOTA: brand 'Ayudín' en la DB lleva tilde — se usa ILIKE para cubrir variantes.

-- ============================================================
-- LIMPIEZA BÁSICA — limpieza → también en almacén
-- ============================================================
UPDATE global_products
SET also_sold_in = '["almacen"]'
WHERE business_type = 'limpieza'
  AND category IN ('limpieza-hogar', 'liquidos-granel', 'papel-higiene', 'accesorios-limpieza')
  AND is_verified = true
  AND (
    -- Por marca de consumo masivo
    brand IN ('Skip', 'Drive', 'Ala', 'Magistral', 'Cif', 'Comfort')
    OR brand ILIKE '%ayud%'
    -- Por tipo de producto (nombres genéricos frecuentes en almacenes)
    OR name ILIKE '%lavandina%'
    OR name ILIKE '%detergente lavavajillas%'
    OR name ILIKE '%jabón en polvo%'
    OR name ILIKE '%suavizante%'
    OR name ILIKE '%papel higiénico%'
    OR name ILIKE '%papel higienico%'
    OR name ILIKE '%servilletas%'
    OR name ILIKE '%bolsas de basura%'
    OR name ILIKE '%desinfectante%'
  );
