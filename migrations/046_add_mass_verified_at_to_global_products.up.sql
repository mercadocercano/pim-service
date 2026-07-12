-- Migration: add mass_verified_at to global_products
-- Epic: MC-E34 (Pantalla de gestión y verificación de global_products en mc_admin)
-- Date: 2026-07-11
-- Reason: el stopgap de 2026-06-19 marcó ~58.921 productos como verificados masivamente
--   para destrabar el surtido. Se necesita un flag derivado reversible para aislar ese
--   lote en la UI de revisión sin depender de mantener la tabla backup en el hot path.

-- UP 1: columna derivada (nullable) que indica cuándo el producto fue verificado masivamente
ALTER TABLE global_products
    ADD COLUMN IF NOT EXISTS mass_verified_at TIMESTAMPTZ;

-- UP 2: backfill desde el backup del mass-verify stopgap.
-- La tabla backup contiene los IDs exactos que se flipearon el 2026-06-19.
-- Si la tabla backup no existe, el UPDATE no afecta ninguna fila (no falla).
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM pg_tables
        WHERE tablename = 'global_products_isverified_bkp_20260619_010824'
    ) THEN
        UPDATE global_products
        SET mass_verified_at = '2026-06-19T01:08:24Z'
        WHERE id IN (
            SELECT id FROM global_products_isverified_bkp_20260619_010824
        )
          AND mass_verified_at IS NULL;
    END IF;
END $$;

-- Índice parcial para filtrar el lote stopgap de forma eficiente
CREATE INDEX IF NOT EXISTS idx_global_products_mass_verified_at
    ON global_products (mass_verified_at)
    WHERE mass_verified_at IS NOT NULL;

COMMENT ON COLUMN global_products.mass_verified_at IS
    'Timestamp de verificación masiva stopgap (2026-06-19). NULL para productos verificados '
    'individualmente. Permite filtrar el lote MC-E34 "verificados sin revisión humana". '
    'Reversible: UPDATE global_products SET is_verified=false, mass_verified_at=NULL WHERE id IN '
    '(SELECT id FROM global_products_isverified_bkp_20260619_010824);';

-- DOWN (comentado, nunca destructivo según política de migraciones)
-- ALTER TABLE global_products DROP COLUMN IF EXISTS mass_verified_at;
