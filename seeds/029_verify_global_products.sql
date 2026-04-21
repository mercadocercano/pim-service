-- Seed 029: Auto-verificación de global_products
-- Umbral aprobado: quality_score >= 60 (D1 del ciclo-003)
-- IDEMPOTENTE: solo actualiza los que aún están en false

UPDATE global_products
SET
    is_verified = true,
    updated_at  = CURRENT_TIMESTAMP
WHERE is_active    = true
  AND is_verified  = false
  AND quality_score >= 60;

-- Reporte post-update
SELECT
    business_type,
    COUNT(*)                                              AS total,
    COUNT(*) FILTER (WHERE is_verified = true)           AS verificados,
    ROUND(AVG(quality_score), 1)                         AS avg_quality
FROM global_products
WHERE is_active = true
GROUP BY business_type
ORDER BY total DESC;
