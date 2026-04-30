-- Seed 070: Template corralón v2 — agrega marcas de sanitarios y griferías
-- CICLO: cycle-004-brands-catalog-expansion
-- FECHA: 2026-04-21
-- PROPÓSITO: UPDATE del template corralón (is_default=true) incorporando marcas
--            de sanitarios (Ferrum, Roca, Deca, Piazza) y griferías (FV, Peirano,
--            Vasser, Hidromet). No modifica categorías ni productos del seed base.
-- IDEMPOTENTE: UPDATE con WHERE; sin efecto si el template no existe.
-- REQUIERE: seed 052 ya ejecutado (global_products corralón con sanitarios).
-- REQUIERE: seed 055 (FV, Poxipol) y seed 065 (Ferrum, Roca, Deca, Piazza, Vasser, etc.).
-- NOTA: FV ya está en marketplace_brands (seed 055, slug 'fv').
--       Poxipol ya está en marketplace_brands (seed 055, slug 'poxipol').

DO $$
BEGIN
  UPDATE business_type_templates
  SET
    brands = '[
      {"name": "Loma Negra",    "suggested_for_categories": ["cemento-cal"]},
      {"name": "Acindar",       "suggested_for_categories": ["hierro-acero"]},
      {"name": "Amanco",        "suggested_for_categories": ["canierias"]},
      {"name": "Tigre",         "suggested_for_categories": ["canierias"]},
      {"name": "Durlock",       "suggested_for_categories": ["maderas", "adhesivos-mezclas"]},
      {"name": "Weber",         "suggested_for_categories": ["adhesivos-mezclas", "cemento-cal"]},
      {"name": "Klaukol",       "suggested_for_categories": ["adhesivos-mezclas"]},
      {"name": "Alba",          "suggested_for_categories": ["pintura"]},
      {"name": "Sherwin Williams", "suggested_for_categories": ["pintura"]},
      {"name": "Sinteplast",    "suggested_for_categories": ["pintura"]},
      {"name": "Sika",          "suggested_for_categories": ["adhesivos-mezclas", "aislantes"]},
      {"name": "Megaflex",      "suggested_for_categories": ["techos"]},
      {"name": "Cincalum",      "suggested_for_categories": ["techos"]},
      {"name": "Isover",        "suggested_for_categories": ["aislantes"]},
      {"name": "Masisa",        "suggested_for_categories": ["maderas"]},
      {"name": "Poxipol",       "suggested_for_categories": ["adhesivos-mezclas"]},
      {"name": "Ferrum",        "suggested_for_categories": ["sanitarios"]},
      {"name": "Roca",          "suggested_for_categories": ["sanitarios"]},
      {"name": "Deca",          "suggested_for_categories": ["sanitarios"]},
      {"name": "Piazza",        "suggested_for_categories": ["sanitarios"]},
      {"name": "FV",            "suggested_for_categories": ["sanitarios"]},
      {"name": "Peirano",       "suggested_for_categories": ["sanitarios"]},
      {"name": "Vasser",        "suggested_for_categories": ["sanitarios"]},
      {"name": "Hidromet",      "suggested_for_categories": ["sanitarios"]}
    ]'::jsonb,
    version      = '2.0.0-sanitarios-griferia',
    generated_by = 'cycle-004-catalog-researcher',
    updated_at   = CURRENT_TIMESTAMP
  WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'corralon')
    AND is_default = true;

  IF NOT FOUND THEN
    RAISE NOTICE 'Template corralón (is_default=true) no encontrado. Verificar que exista el business_type ''corralon'' y su template default.';
  ELSE
    RAISE NOTICE 'Template corralón actualizado: 24 marcas (construcción + sanitarios + griferías) — v2.0.0.';
  END IF;
END $$;
