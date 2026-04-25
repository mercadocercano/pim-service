-- Seed 102: Productos globales — Verdulería: Frutas y Verduras (~65 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: relevamiento verdulerías NEA (Posadas) + Mercado Regional Misiones 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: frutas-frescas, verduras-hoja, verduras-raiz, verduras-fruto, hierbas-aromáticas
-- NOTA: productos estacionales — disponibilidad varía según temporada en NEA (clima subtropical)

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- FRUTAS — precio x kg salvo indicación
-- ============================================================
('Manzana roja x kg',                            NULL, 'frutas-frescas',  2800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Manzana verde Granny Smith x kg',              NULL, 'frutas-frescas',  3200.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Naranja x kg',                                 NULL, 'frutas-frescas',  1800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Banana x kg',                                  NULL, 'frutas-frescas',  2200.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Mandarina x kg',                               NULL, 'frutas-frescas',  2000.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Pera x kg',                                    NULL, 'frutas-frescas',  3500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Uva negra sin semilla x kg',                   NULL, 'frutas-frescas',  4500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Uva verde x kg',                               NULL, 'frutas-frescas',  4200.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Durazno x kg',                                 NULL, 'frutas-frescas',  3800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Sandía x unidad',                              NULL, 'frutas-frescas', 12000.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),
('Melón x unidad',                               NULL, 'frutas-frescas',  8500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),
('Limón x kg',                                   NULL, 'frutas-frescas',  2500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Pomelo rosa x kg',                             NULL, 'frutas-frescas',  2200.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Kiwi x unidad',                                NULL, 'frutas-frescas',   800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),
('Ananá x unidad',                               NULL, 'frutas-frescas',  6500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),
('Frutilla x bandeja 250g',                      NULL, 'frutas-frescas',  3500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "bandeja", "sku_prefix": "VERD"}'),
('Ciruela negra x kg',                           NULL, 'frutas-frescas',  3800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Mango x unidad',                               NULL, 'frutas-frescas',  2500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),
('Maracuyá x unidad',                            NULL, 'frutas-frescas',   800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),

-- ============================================================
-- VERDURAS DE HOJA
-- ============================================================
('Lechuga mantecosa x unidad',                   NULL, 'verduras-hoja',   1800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),
('Lechuga rizada x unidad',                      NULL, 'verduras-hoja',   2000.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),
('Acelga x atado',                               NULL, 'verduras-hoja',   1500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "atado",   "sku_prefix": "VERD"}'),
('Espinaca x atado',                             NULL, 'verduras-hoja',   1500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "atado",   "sku_prefix": "VERD"}'),
('Rúcula x atado',                               NULL, 'verduras-hoja',   2200.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "atado",   "sku_prefix": "VERD"}'),
('Repollo blanco x kg',                          NULL, 'verduras-hoja',   1800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Repollo colorado x kg',                        NULL, 'verduras-hoja',   2000.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Brócoli x unidad',                             NULL, 'verduras-hoja',   3500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),
('Coliflor x unidad',                            NULL, 'verduras-hoja',   3800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),
('Kale x atado',                                 NULL, 'verduras-hoja',   2800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "atado",   "sku_prefix": "VERD"}'),

-- ============================================================
-- VERDURAS DE RAÍZ Y TUBÉRCULO
-- ============================================================
('Papa blanca x kg',                             NULL, 'verduras-raiz',   2200.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Papa negra x kg',                              NULL, 'verduras-raiz',   2000.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Cebolla blanca x kg',                          NULL, 'verduras-raiz',   2000.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Cebolla morada x kg',                          NULL, 'verduras-raiz',   2500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Zanahoria x kg',                               NULL, 'verduras-raiz',   2200.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Remolacha x kg',                               NULL, 'verduras-raiz',   1800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Batata x kg',                                  NULL, 'verduras-raiz',   2500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Ajo x cabeza',                                 NULL, 'verduras-raiz',    500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),
('Ajo x kg',                                     NULL, 'verduras-raiz',   4500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Nabo x kg',                                    NULL, 'verduras-raiz',   1800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),

-- ============================================================
-- VERDURAS DE FRUTO
-- ============================================================
('Tomate perita x kg',                           NULL, 'verduras-fruto',  3200.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Tomate redondo x kg',                          NULL, 'verduras-fruto',  2800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Zapallo anco x kg',                            NULL, 'verduras-fruto',  1800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Zapallo criollo x kg',                         NULL, 'verduras-fruto',  1500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Pimiento rojo x kg',                           NULL, 'verduras-fruto',  3800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Pimiento verde x kg',                          NULL, 'verduras-fruto',  3200.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Berenjena x kg',                               NULL, 'verduras-fruto',  3000.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Zapallito verde x kg',                         NULL, 'verduras-fruto',  2500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Choclo x unidad',                              NULL, 'verduras-fruto',  1500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),
('Pepino x kg',                                  NULL, 'verduras-fruto',  2500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),
('Chaucha x kg',                                 NULL, 'verduras-fruto',  4500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "kg",      "sku_prefix": "VERD"}'),

-- ============================================================
-- HIERBAS AROMÁTICAS Y CONDIMENTOS FRESCOS
-- ============================================================
('Perejil x atado',                              NULL, 'hierbas-aromaticas',  800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "atado",   "sku_prefix": "VERD"}'),
('Ciboulette x atado',                           NULL, 'hierbas-aromaticas', 1000.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "atado",   "sku_prefix": "VERD"}'),
('Albahaca x atado',                             NULL, 'hierbas-aromaticas',  900.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "atado",   "sku_prefix": "VERD"}'),
('Menta x atado',                                NULL, 'hierbas-aromaticas',  800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "atado",   "sku_prefix": "VERD"}'),
('Cilantro x atado',                             NULL, 'hierbas-aromaticas',  900.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "atado",   "sku_prefix": "VERD"}'),
('Apio x atado',                                 NULL, 'hierbas-aromaticas', 1800.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "atado",   "sku_prefix": "VERD"}'),
('Puerro x unidad',                              NULL, 'hierbas-aromaticas', 1500.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "unidad",  "sku_prefix": "VERD"}'),
('Cebolla de verdeo x atado',                    NULL, 'hierbas-aromaticas', 1200.00, 'seed', 0.8, 60, FALSE, TRUE, 'verduleria', '{"unit": "atado",   "sku_prefix": "VERD"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 65 productos
-- Categorías: frutas-frescas, verduras-hoja, verduras-raiz, verduras-fruto, hierbas-aromaticas
-- Marcas: todas genérico (sin marca en verdulería)
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
-- NOTA: frutas tropicales (mango, maracuyá, ananá) son abundantes y de menor costo en NEA vs resto del país
