-- Seed 067: Template almacén v2 — agrega cervezas y vinos básicos
-- CICLO: cycle-004-brands-catalog-expansion
-- FECHA: 2026-04-22
-- PROPÓSITO: UPDATE del template almacén (is_default=true) incorporando marcas de
--            bebidas alcohólicas (cervezas + vinos de rotación básica).
--            Conserva todas las marcas de seed 057 y suma bebidas alcohólicas.
-- IDEMPOTENTE: UPDATE con WHERE; sin efecto si el template no existe.
-- REQUIERE: seed 057 ya ejecutado (template almacén v1 con 28 marcas).
-- REQUIERE: seeds 062 (cervezas/aperitivos) y 063 (vinos) ya ejecutados.
-- VERSION: 2.0.0-cervezas-vinos

DO $$
BEGIN
  UPDATE business_type_templates
  SET
    brands = '[
      {"name": "La Serenísima",          "suggested_for_categories": ["lacteos-liquidos", "lacteos-fiambres"]},
      {"name": "Sancor",                 "suggested_for_categories": ["lacteos-liquidos"]},
      {"name": "La Morenita",            "suggested_for_categories": ["lacteos-liquidos"]},
      {"name": "Danone",                 "suggested_for_categories": ["lacteos-liquidos"]},
      {"name": "Tregar",                 "suggested_for_categories": ["lacteos-fiambres"]},
      {"name": "Coca-Cola",              "suggested_for_categories": ["bebidas-sin-alcohol", "gaseosas"]},
      {"name": "Pepsi",                  "suggested_for_categories": ["bebidas-sin-alcohol", "gaseosas"]},
      {"name": "Manaos",                 "suggested_for_categories": ["bebidas-sin-alcohol", "gaseosas"]},
      {"name": "Fanta",                  "suggested_for_categories": ["bebidas-sin-alcohol", "gaseosas"]},
      {"name": "Arcor",                  "suggested_for_categories": ["alimentos-envasados", "conservas"]},
      {"name": "Bagley",                 "suggested_for_categories": ["snacks-salados", "alimentos-envasados"]},
      {"name": "Fargo",                  "suggested_for_categories": ["alimentos-envasados"]},
      {"name": "Paladini",               "suggested_for_categories": ["lacteos-fiambres"]},
      {"name": "Marolio",                "suggested_for_categories": ["alimentos-envasados", "conservas"]},
      {"name": "La Campagnola",          "suggested_for_categories": ["conservas"]},
      {"name": "Molinos Río de la Plata","suggested_for_categories": ["pastas-cereales", "alimentos-envasados"]},
      {"name": "Molino Cañuelas",        "suggested_for_categories": ["pastas-cereales", "alimentos-envasados"]},
      {"name": "Ledesma",                "suggested_for_categories": ["alimentos-envasados"]},
      {"name": "Nescafé",                "suggested_for_categories": ["alimentos-envasados"]},
      {"name": "Rosamonte",              "suggested_for_categories": ["alimentos-envasados"]},
      {"name": "Las Marías",             "suggested_for_categories": ["alimentos-envasados"]},
      {"name": "Skip",                   "suggested_for_categories": ["general"]},
      {"name": "Ariel",                  "suggested_for_categories": ["general"]},
      {"name": "Dove",                   "suggested_for_categories": ["general"]},
      {"name": "Colgate",                "suggested_for_categories": ["general"]},
      {"name": "Rexona",                 "suggested_for_categories": ["general"]},
      {"name": "Sedal",                  "suggested_for_categories": ["general"]},
      {"name": "Cif",                    "suggested_for_categories": ["general"]},
      {"name": "Quilmes",                "suggested_for_categories": ["bebidas-con-alcohol", "cervezas"]},
      {"name": "Brahma",                 "suggested_for_categories": ["bebidas-con-alcohol", "cervezas"]},
      {"name": "Norte",                  "suggested_for_categories": ["bebidas-con-alcohol", "cervezas"]},
      {"name": "Heineken",               "suggested_for_categories": ["bebidas-con-alcohol", "cervezas"]},
      {"name": "Trapiche",               "suggested_for_categories": ["bebidas-con-alcohol", "vinos-tintos", "vinos-blancos"]},
      {"name": "Gato Negro",             "suggested_for_categories": ["bebidas-con-alcohol", "vinos-tintos", "vinos-blancos"]},
      {"name": "Callia",                 "suggested_for_categories": ["bebidas-con-alcohol", "vinos-tintos", "vinos-blancos"]},
      {"name": "Fernet Branca",          "suggested_for_categories": ["bebidas-con-alcohol", "aperitivos"]}
    ]'::jsonb,
    version      = '2.0.0-cervezas-vinos',
    generated_by = 'cycle-004-catalog-researcher',
    updated_at   = CURRENT_TIMESTAMP
  WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'almacen')
    AND is_default = true;

  IF NOT FOUND THEN
    RAISE NOTICE 'Template almacén (is_default=true) no encontrado. Ejecutar seed 022 primero.';
  ELSE
    RAISE NOTICE 'Template almacén actualizado: 36 marcas (v2.0.0 +cervezas +vinos básicos).';
  END IF;
END $$;
