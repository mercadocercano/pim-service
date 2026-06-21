-- Migration: create global_product_reclassification_audit table
-- Epic: E24 — Re-clasificación de business_type como endpoint S2S
-- ADR: ADR-005 §4 — Audit log persistente separado del canonical log ADR-001
-- Date: 2026-06-17
-- Reason: rastro forense con retención propia, queryable, sobrevive rotación de logs.
--   Distinto del canonical log (Loki) que es observabilidad operacional.

-- UP: crear tabla de auditoría
CREATE TABLE IF NOT EXISTS global_product_reclassification_audit (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    operator_id     TEXT NOT NULL,
    executed_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    mode            TEXT NOT NULL CHECK (mode IN ('dry_run', 'applied')),
    scope           JSONB NOT NULL,
    snapshot_ref    TEXT,
    summary         JSONB NOT NULL,
    affected_count  INT NOT NULL DEFAULT 0
);

-- Índices para consultas operacionales frecuentes
CREATE INDEX IF NOT EXISTS idx_reclassification_audit_executed_at
    ON global_product_reclassification_audit (executed_at DESC);

CREATE INDEX IF NOT EXISTS idx_reclassification_audit_operator_id
    ON global_product_reclassification_audit (operator_id);

CREATE INDEX IF NOT EXISTS idx_reclassification_audit_mode
    ON global_product_reclassification_audit (mode);

COMMENT ON TABLE global_product_reclassification_audit IS
    'Rastro forense de operaciones de re-clasificación masiva de business_type. '
    'NO debe tener endpoint DELETE ni UPDATE expuesto — inmutable por diseño (ADR-005 §4, E24-L4). '
    'El detalle antes/después por fila vive en el snapshot tabular global_products_bkp_<ts>.';

COMMENT ON COLUMN global_product_reclassification_audit.operator_id IS
    'ID del operador humano propagado por mc_admin via header X-Operator-Id (ADR-005 §4).';

COMMENT ON COLUMN global_product_reclassification_audit.mode IS
    'dry_run: simulación sin mutación. applied: cambios reales aplicados a global_products.';

COMMENT ON COLUMN global_product_reclassification_audit.scope IS
    'Criterio de selección JSONB: {"source_prefix":"scraper","max_rows":50000}.';

COMMENT ON COLUMN global_product_reclassification_audit.snapshot_ref IS
    'Nombre de la tabla de backup creada: global_products_bkp_<timestamp>. NULL en dry_run.';

COMMENT ON COLUMN global_product_reclassification_audit.summary IS
    'Resumen estructurado JSONB con total_evaluados, candidatos, updates_por_rubro, '
    'colisiones_skipeadas, skips. Mismo formato que la response del endpoint.';

COMMENT ON COLUMN global_product_reclassification_audit.affected_count IS
    'Cantidad de rows efectivamente actualizados en global_products. 0 en dry_run.';

-- DOWN (comentado, nunca destructivo según política de migraciones)
-- DROP TABLE IF EXISTS global_product_reclassification_audit;
