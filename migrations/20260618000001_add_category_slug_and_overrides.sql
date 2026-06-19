-- Migration: add category_slug to global_products + category_slug_overrides table
-- Epic: Saneamiento de surtido de templates del Quickstart
-- ADR: ADR-007 §1, §2 — eje category_slug normalizado (preserva category raw)
-- Date: 2026-06-18
-- Reason: global_products.category viene en 3 formatos sucios (paths VTEX, MAYÚSCULAS/EN,
--   kebab). El join de RefreshProductTemplates es exact-match contra el slug declarado y por
--   eso ignora ~2.000 categorías. Se agrega una columna NORMALIZADA category_slug (derivada,
--   recalculable) preservando category raw para auditoría, y una tabla de overrides curables
--   para el long tail que el resolver determinístico no cubre limpio.

-- UP 1: columna normalizada (nullable; se puebla por el resolver en ingestión + backfill S2S)
ALTER TABLE global_products
    ADD COLUMN IF NOT EXISTS category_slug VARCHAR(200);

-- Índice parcial: solo filas ya normalizadas (el join del refresh filtra por category_slug)
CREATE INDEX IF NOT EXISTS idx_global_products_category_slug
    ON global_products (category_slug)
    WHERE category_slug IS NOT NULL;

COMMENT ON COLUMN global_products.category_slug IS
    'Slug de categoría NORMALIZADO (ADR-007). Derivado mecánico de category raw vía el resolver '
    'de go-shared (domain/category) + tabla category_slug_overrides. Se recalcula siempre en '
    'ingestión; category raw se preserva como rastro de la fuente. NULL hasta el primer backfill.';

-- UP 2: tabla de overrides curables (long tail que el resolver no resuelve limpio)
-- Precedencia (ADR-007 §2): override (esta tabla) > resolver (código go-shared) > 'sin-clasificar'.
CREATE TABLE IF NOT EXISTS category_slug_overrides (
    raw_category   TEXT PRIMARY KEY,            -- valor crudo exacto de global_products.category
    category_slug  VARCHAR(200) NOT NULL,       -- slug declarado destino
    note           TEXT,                        -- por qué este override (curación manual)
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE category_slug_overrides IS
    'Overrides curados raw_category -> category_slug para el long tail de categorías que el '
    'resolver determinístico de go-shared no mapea limpio (ADR-007 §2). Seedeable por migración '
    'y editable sin redeploy. Gana sobre el resolver. Sin fuzzy: match exacto por raw_category.';

COMMENT ON COLUMN category_slug_overrides.raw_category IS
    'Valor crudo EXACTO de global_products.category (path VTEX, constante EN o kebab).';

COMMENT ON COLUMN category_slug_overrides.category_slug IS
    'Slug declarado destino (ej. cervezas-vinos, vinos-tintos). Debe existir en algún '
    'business_type_templates.categories[].slug para que el refresh lo aproveche.';

-- DOWN (comentado, nunca destructivo según política de migraciones)
-- ALTER TABLE global_products DROP COLUMN IF EXISTS category_slug;
-- DROP TABLE IF EXISTS category_slug_overrides;
