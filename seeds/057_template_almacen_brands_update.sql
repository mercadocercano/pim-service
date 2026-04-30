-- Seed 057: Actualización de brands en template de Almacén
-- CICLO: cycle-002-global-brands-colors / ADR-001
-- FUENTE: curación manual — ampliación del array brands con marcas del seed 056
-- IDEMPOTENTE: UPDATE con WHERE is_default = true — puede correrse N veces
-- REQUIERE: 022 (template almacen curado), 056 (marcas adicionales)
-- VERSION: 1.0.0

DO $$
DECLARE
  v_updated INT;
BEGIN
  UPDATE business_type_templates
  SET
    brands = '[
      {"name": "La Serenísima",          "suggested_for_categories": ["lacteos-liquidos","lacteos-fiambres"]},
      {"name": "Sancor",                 "suggested_for_categories": ["lacteos-liquidos"]},
      {"name": "La Morenita",            "suggested_for_categories": ["lacteos-liquidos"]},
      {"name": "Danone",                 "suggested_for_categories": ["lacteos-liquidos"]},
      {"name": "Tregar",                 "suggested_for_categories": ["lacteos-fiambres"]},
      {"name": "Coca-Cola",              "suggested_for_categories": ["bebidas-sin-alcohol"]},
      {"name": "Pepsi",                  "suggested_for_categories": ["bebidas-sin-alcohol"]},
      {"name": "Manaos",                 "suggested_for_categories": ["bebidas-sin-alcohol"]},
      {"name": "Quilmes",                "suggested_for_categories": ["bebidas-sin-alcohol"]},
      {"name": "Arcor",                  "suggested_for_categories": ["alimentos-envasados","conservas"]},
      {"name": "Bagley",                 "suggested_for_categories": ["snacks-salados","alimentos-envasados"]},
      {"name": "Fargo",                  "suggested_for_categories": ["alimentos-envasados"]},
      {"name": "Paladini",               "suggested_for_categories": ["lacteos-fiambres"]},
      {"name": "Marolio",                "suggested_for_categories": ["alimentos-envasados","conservas"]},
      {"name": "La Campagnola",          "suggested_for_categories": ["conservas"]},
      {"name": "Molinos Río de la Plata","suggested_for_categories": ["pastas-cereales","alimentos-envasados"]},
      {"name": "Molino Cañuelas",        "suggested_for_categories": ["pastas-cereales","alimentos-envasados"]},
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
      {"name": "Cif",                    "suggested_for_categories": ["general"]}
    ]'::jsonb,
    updated_at = NOW()
  WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'almacen')
    AND is_default = true;

  GET DIAGNOSTICS v_updated = ROW_COUNT;
  RAISE NOTICE 'Seed 057 — template almacen brands actualizado: % fila(s)', v_updated;
END $$;
