-- photoless_segmentation.sql
-- Auditoría READ-ONLY que clasifica los productos del catálogo global SIN foto real
-- (image_url null o placeholder wikimedia) en buckets accionables para E16.
-- No modifica datos. Salida determinista (ordenada) para diff limpio.
--
-- Uso:
--   docker exec -e PGPASSWORD=postgres lab-postgres \
--     psql -U postgres -d pim_db -f /ruta/photoless_segmentation.sql
--   (o: psql -h <host> -p <port> -U postgres -d pim_db -f photoless_segmentation.sql)
--
-- Buckets (precedencia M > B2 > B1 > A):
--   M  — Merge:      duplicado exacto por nombre de un producto activo CON foto real.
--   B1 — Marca:      Grupo B, con marca, no a granel → resoluble por nombre (ML/OFF/Firecrawl/Perplexity).
--   B2 — Granel:     Grupo B, sin marca o vendido por peso → fuentes reales inciertas; curaduría manual si no resuelve.
--   A  — Saturado:   rubros donde la expansión VTEX ya dejó miles de fotos reales (prioridad baja).
--
-- NOTA QA (T2, 2026-06-17): NO existe bucket "inválido/no-producto". Se inspeccionaron los
-- candidatos a basura (nombres cortos sin marca, 'test', numéricos): todos resultaron productos
-- legítimos (cortes de carne, verduras a granel, tests de embarazo, kits de pileta). Los 2264 son
-- reales; el eje real es resoluble-por-nombre (B1) vs genérico-a-granel sin fuente de foto (B2).
--
-- Grupo A (saturados) = farmacia, perfumeria, fiambreria, limpieza, bazar, vinoteca, veterinaria, almacen.
-- Grupo B (gap real)  = el resto (ropa, libreria, verduleria, piletas, peluqueria, carniceria, jugueteria, electrodomesticos, ...).

\pset pager off
\pset footer off
\timing off
\set ON_ERROR_STOP off

-- ──────────────────────────────────────────────────────────────────────────
-- CTEs base
-- ──────────────────────────────────────────────────────────────────────────

-- Universo problema: activos sin foto real.
DROP VIEW IF EXISTS _prob CASCADE;
CREATE TEMP VIEW _prob AS
SELECT id, name, brand, business_type, quality_score, source, is_verified, description,
       lower(trim(name)) AS nm,
       (image_url IS NULL OR image_url = '') AS img_null
FROM global_products
WHERE is_active
  AND (image_url IS NULL OR image_url = '' OR image_url ILIKE '%wikimedia%');

-- Nombres (case-insensitive) que YA tienen un producto activo con foto real.
DROP VIEW IF EXISTS _real_names CASCADE;
CREATE TEMP VIEW _real_names AS
SELECT DISTINCT lower(trim(name)) AS nm
FROM global_products
WHERE is_active
  AND image_url IS NOT NULL AND image_url <> ''
  AND image_url NOT ILIKE '%wikimedia%';

-- Clasificación de cada producto del universo problema.
DROP VIEW IF EXISTS _classified CASCADE;
CREATE TEMP VIEW _classified AS
SELECT
  p.*,
  (p.business_type IN ('farmacia','perfumeria','fiambreria','limpieza',
                       'bazar','vinoteca','veterinaria','almacen')) AS is_grupo_a,
  (rn.nm IS NOT NULL)                                               AS dup_of_real,
  (coalesce(trim(p.brand),'') = '')                                 AS sin_marca,
  -- a granel / por peso: típico de perecederos sin packaging
  (p.name ILIKE '%x kg%' OR p.name ILIKE '% por kg%' OR p.name ILIKE '% kg'
   OR p.name ILIKE '%x unidad%' OR p.name ILIKE '% docena%')        AS by_weight
FROM _prob p
LEFT JOIN _real_names rn ON rn.nm = p.nm;

-- Asignación de bucket con la precedencia definida.
DROP VIEW IF EXISTS _bucketed CASCADE;
CREATE TEMP VIEW _bucketed AS
SELECT *,
  CASE
    WHEN dup_of_real                              THEN 'M'
    WHEN NOT is_grupo_a AND NOT sin_marca
         AND NOT by_weight                        THEN 'B1'
    WHEN NOT is_grupo_a                           THEN 'B2'
    ELSE                                               'A'
  END AS bucket
FROM _classified;

-- ──────────────────────────────────────────────────────────────────────────
\echo '================================================================'
\echo 'S0  TOTAL del universo problema (debe == 2264 aprox)'
\echo '================================================================'
SELECT count(*) AS total_sin_foto,
       count(*) FILTER (WHERE img_null)       AS null_img,
       count(*) FILTER (WHERE NOT img_null)   AS wikimedia
FROM _prob;

\echo ''
\echo '================================================================'
\echo 'S1  CONTEO POR BUCKET'
\echo '================================================================'
SELECT bucket, count(*) AS n,
       round(100.0*count(*)/sum(count(*)) over (),1) AS pct
FROM _bucketed GROUP BY bucket ORDER BY bucket;

\echo ''
\echo '================================================================'
\echo 'S2  BUCKET x RUBRO'
\echo '================================================================'
SELECT coalesce(business_type,'(null)') AS rubro,
       count(*) FILTER (WHERE bucket='M')  AS m,
       count(*) FILTER (WHERE bucket='B1') AS b1,
       count(*) FILTER (WHERE bucket='B2') AS b2,
       count(*) FILTER (WHERE bucket='A')  AS a,
       count(*)                            AS total
FROM _bucketed GROUP BY rubro ORDER BY total DESC;

\echo ''
\echo '================================================================'
\echo 'S3  IDs del bucket M (merge — desactivar placeholder en T3)'
\echo '================================================================'
SELECT id, business_type, left(name,50) AS name, brand
FROM _bucketed WHERE bucket='M' ORDER BY business_type, name;

\echo ''
\echo '================================================================'
\echo 'S4  MUESTRA bucket B1 (marca, resoluble por nombre)'
\echo '================================================================'
SELECT business_type, left(name,50) AS name, brand
FROM _bucketed WHERE bucket='B1' ORDER BY business_type, name LIMIT 30;

\echo ''
\echo '================================================================'
\echo 'S5  MUESTRA bucket B2 (genérico a granel)'
\echo '================================================================'
SELECT business_type, left(name,50) AS name, coalesce(brand,'(sin marca)') AS brand
FROM _bucketed WHERE bucket='B2' ORDER BY business_type, name LIMIT 30;
