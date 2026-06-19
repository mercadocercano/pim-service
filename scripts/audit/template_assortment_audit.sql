-- template_assortment_audit.sql
-- Auditoría READ-ONLY del surtido de los templates del Quickstart (business_type_templates).
-- Objetivo: dimensionar la calidad del surtido curado y producir el insumo del mapa EN->ES.
-- No modifica datos. Salida ordenada/determinista.
--
-- Uso:
--   docker exec -e PGPASSWORD=postgres lab-postgres \
--     psql -U postgres -d pim_db -f /ruta/template_assortment_audit.sql
--
-- Dimensiones:
--   A1 Resumen por template (huérfanos + cobertura de marcas)
--   A2 Slugs huérfanos consolidados (INSUMO del mapa EN->ES)
--   A3 Marcas del surtido real NO representadas en la tira curada
--   A4 Productos sospechosos (nombre vs marca incoherente)

\pset pager off
\pset footer off
\timing off
\set ON_ERROR_STOP off

-- Slugs declarados por template (categories[].slug)
DROP TABLE IF EXISTS _decl;
CREATE TEMP TABLE _decl AS
SELECT btt.id AS tid, c->>'slug' AS slug
FROM business_type_templates btt, jsonb_array_elements(btt.categories) c
WHERE btt.is_active;

-- Productos del template con flag de match contra los slugs declarados
DROP TABLE IF EXISTS _prod;
CREATE TEMP TABLE _prod AS
SELECT btt.id                       AS tid,
       btt.name                     AS tname,
       coalesce(bt.code, '')        AS slug_bt,
       p->>'name'                   AS pname,
       nullif(p->>'category_slug','') AS slug,
       nullif(p->>'brand','')       AS brand,
       EXISTS (
         SELECT 1 FROM _decl d
         WHERE d.tid = btt.id AND d.slug = nullif(p->>'category_slug','')
       )                            AS matched
FROM business_type_templates btt
LEFT JOIN business_types bt ON bt.id = btt.business_type_id,
     jsonb_array_elements(btt.products) p
WHERE btt.is_active;

-- Marcas curadas (tira highlight) por template
DROP TABLE IF EXISTS _brands_cur;
CREATE TEMP TABLE _brands_cur AS
SELECT btt.id AS tid, b->>'name' AS brand
FROM business_type_templates btt, jsonb_array_elements(btt.brands) b
WHERE btt.is_active;

\echo '================================================================'
\echo 'A1  RESUMEN POR TEMPLATE  (huerfanos + cobertura de marcas)'
\echo '================================================================'
\echo '  matched   = productos cuyo category_slug esta en las categorias declaradas'
\echo '  huerfanos = productos cuyo category_slug NO matchea (taxonomia rota)'
\echo '  m_real    = marcas distintas en el surtido (products[].brand)'
\echo '  m_tira    = marcas en la tira curada (brands[])'
SELECT pr.slug_bt                                            AS bt,
       count(*)                                              AS total,
       count(*) FILTER (WHERE pr.matched)                    AS matched,
       count(*) FILTER (WHERE NOT pr.matched)                AS huerfanos,
       round(100.0 * count(*) FILTER (WHERE NOT pr.matched)
             / nullif(count(*),0), 1)                        AS pct_huerf,
       count(DISTINCT pr.brand)                              AS m_real,
       (SELECT count(DISTINCT bc.brand) FROM _brands_cur bc WHERE bc.tid = pr.tid) AS m_tira
FROM _prod pr
GROUP BY pr.tid, pr.slug_bt
ORDER BY pct_huerf DESC, total DESC;

\echo ''
\echo '================================================================'
\echo 'A2  SLUGS HUERFANOS CONSOLIDADOS  (INSUMO del mapa EN->ES)'
\echo '================================================================'
\echo '  en_const = slug que parece constante de MercadoLibre (^[A-Z_]+$)'
SELECT pr.slug                                               AS slug_huerfano,
       (pr.slug ~ '^[A-Z][A-Z0-9_]*$')                       AS en_const,
       count(*)                                              AS prods,
       count(DISTINCT pr.tid)                                AS en_n_templates
FROM _prod pr
WHERE NOT pr.matched AND pr.slug IS NOT NULL
GROUP BY pr.slug
ORDER BY en_const DESC, prods DESC, slug_huerfano;

\echo ''
\echo '-- Total de constantes EN distintas (universo a mapear) --'
SELECT count(*) AS slugs_en_distintos,
       sum(prods) AS productos_afectados
FROM (
  SELECT pr.slug, count(*) AS prods
  FROM _prod pr
  WHERE NOT pr.matched AND pr.slug ~ '^[A-Z][A-Z0-9_]*$'
  GROUP BY pr.slug
) s;

\echo ''
\echo '================================================================'
\echo 'A3  MARCAS REALES NO REPRESENTADAS EN LA TIRA CURADA'
\echo '================================================================'
SELECT pr.slug_bt AS bt, pr.brand AS marca_en_surtido, count(*) AS prods
FROM _prod pr
WHERE pr.brand IS NOT NULL
  AND NOT EXISTS (SELECT 1 FROM _brands_cur bc WHERE bc.tid = pr.tid AND bc.brand = pr.brand)
GROUP BY pr.slug_bt, pr.brand
ORDER BY pr.slug_bt, prods DESC, marca_en_surtido;

\echo ''
\echo '================================================================'
\echo 'A4  PRODUCTOS SOSPECHOSOS  (nombre menciona una marca != brand)'
\echo '================================================================'
WITH tokens(tok) AS (
  VALUES ('quilmes'),('brahma'),('stella'),('andes'),('coca'),('pepsi'),
         ('arcor'),('serenisima'),('sancor'),('marolio'),('terrabusi'),
         ('bagley'),('molinos'),('rosamonte'),('dove'),('natura')
)
SELECT pr.slug_bt AS bt, pr.pname AS producto, pr.brand AS marca, pr.slug AS category_slug
FROM _prod pr
JOIN tokens t ON lower(pr.pname) LIKE '%'||t.tok||'%'
WHERE pr.brand IS NULL OR lower(pr.brand) NOT LIKE '%'||t.tok||'%'
ORDER BY pr.slug_bt, producto;
