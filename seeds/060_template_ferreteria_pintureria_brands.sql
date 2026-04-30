-- Seed 060: Actualización de brands en templates de Ferretería y Pinturería
-- CICLO: cycle-002-global-brands-colors / ADR-001
-- FUENTE: curación manual — ampliación con marcas del seed 056
-- IDEMPOTENTE: UPDATE con WHERE is_default = true — puede correrse N veces
-- REQUIERE: 023 (template ferreteria), 024 (template pintureria), 056 (marcas adicionales)
-- VERSION: 1.0.0

-- =====================================================
-- PARTE 1: Ferretería — brands completas (existentes + nuevas)
-- =====================================================
DO $$
DECLARE
  v_updated INT;
BEGIN
  UPDATE business_type_templates
  SET
    brands = '[
      {"name": "Loma Negra",      "suggested_for_categories": ["materiales-construccion"]},
      {"name": "Avellaneda",      "suggested_for_categories": ["materiales-construccion"]},
      {"name": "Acindar",         "suggested_for_categories": ["materiales-construccion"]},
      {"name": "Tigre",           "suggested_for_categories": ["plomeria"]},
      {"name": "FV",              "suggested_for_categories": ["plomeria"]},
      {"name": "Stanley",         "suggested_for_categories": ["herramientas-manuales"]},
      {"name": "Tramontina",      "suggested_for_categories": ["herramientas-manuales"]},
      {"name": "Bosch",           "suggested_for_categories": ["herramientas-electricas"]},
      {"name": "Black+Decker",    "suggested_for_categories": ["herramientas-electricas"]},
      {"name": "Makita",          "suggested_for_categories": ["herramientas-electricas"]},
      {"name": "DeWalt",          "suggested_for_categories": ["herramientas-electricas"]},
      {"name": "3M",              "suggested_for_categories": ["herramientas-manuales","materiales-construccion"]},
      {"name": "Sherwin-Williams","suggested_for_categories": ["pinturas"]},
      {"name": "Alba",            "suggested_for_categories": ["pinturas"]},
      {"name": "Sinteplast",      "suggested_for_categories": ["pinturas"]},
      {"name": "Sika",            "suggested_for_categories": ["materiales-construccion","pinturas"]},
      {"name": "Poxipol",         "suggested_for_categories": ["materiales-construccion"]},
      {"name": "Workpro",         "suggested_for_categories": ["herramientas-manuales","herramientas-electricas"]},
      {"name": "Sekur",           "suggested_for_categories": ["herramientas-manuales"]}
    ]'::jsonb,
    updated_at = NOW()
  WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'ferreteria')
    AND is_default = true;

  GET DIAGNOSTICS v_updated = ROW_COUNT;
  RAISE NOTICE 'Seed 060 — template ferreteria brands actualizado: % fila(s)', v_updated;
END $$;

-- =====================================================
-- PARTE 2: Pinturería — brands completas (existentes + nuevas)
-- =====================================================
DO $$
DECLARE
  v_updated INT;
BEGIN
  UPDATE business_type_templates
  SET
    brands = '[
      {"name": "Alba",             "suggested_for_categories": ["latex-interior","latex-exterior","esmaltes-sinteticos","enduidos-masillas"]},
      {"name": "Sinteplast",       "suggested_for_categories": ["latex-interior","latex-exterior","enduidos-masillas"]},
      {"name": "Sherwin-Williams", "suggested_for_categories": ["latex-interior","latex-exterior","esmaltes-sinteticos"]},
      {"name": "Tersuave",         "suggested_for_categories": ["latex-interior","latex-exterior"]},
      {"name": "Colorín",          "suggested_for_categories": ["esmaltes-sinteticos","accesorios-pintura"]},
      {"name": "Riopint",          "suggested_for_categories": ["esmaltes-sinteticos","latex-exterior"]},
      {"name": "Cetol",            "suggested_for_categories": ["esmaltes-sinteticos","accesorios-pintura"]},
      {"name": "Plavicon",         "suggested_for_categories": ["latex-interior","latex-exterior","enduidos-masillas"]},
      {"name": "Dux",              "suggested_for_categories": ["latex-interior","latex-exterior"]},
      {"name": "Sika",             "suggested_for_categories": ["enduidos-masillas"]}
    ]'::jsonb,
    updated_at = NOW()
  WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'pintureria')
    AND is_default = true;

  GET DIAGNOSTICS v_updated = ROW_COUNT;
  RAISE NOTICE 'Seed 060 — template pintureria brands actualizado: % fila(s)', v_updated;
END $$;
