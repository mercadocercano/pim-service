-- Seed 066: Template kiosco v2 — amplía marcas a ≥30 (golosinas completas)
-- CICLO: cycle-004-brands-catalog-expansion
-- FECHA: 2026-04-21
-- PROPÓSITO: UPDATE del template kiosco (is_default=true) ampliando el array de brands
--            de 10 a 31 marcas. No modifica categorías ni productos — solo brands.
-- IDEMPOTENTE: UPDATE con WHERE; sin efecto si el template no existe.
-- REQUIERE: seed 021 ya ejecutado (template base kiosco).
-- REQUIERE: seeds 055 y 064 ya ejecutados (marcas con colores en marketplace_brands).
-- NOTA: Las marcas referenciadas aquí deben existir en marketplace_brands.
--       Marcas que ya estaban en v1: Coca-Cola, Pepsi, Arcor, Lay's, Milka, Beldent, Quilmes,
--       Fanta, Speed, Bagley.
--       Marcas que NO estaban y se agregan: resto listado abajo.

DO $$
BEGIN
  UPDATE business_type_templates
  SET
    brands = '[
      {"name": "Coca-Cola",           "suggested_for_categories": ["bebidas-sin-alcohol", "gaseosas"]},
      {"name": "Pepsi",               "suggested_for_categories": ["bebidas-sin-alcohol", "gaseosas"]},
      {"name": "Manaos",              "suggested_for_categories": ["bebidas-sin-alcohol", "gaseosas"]},
      {"name": "Fanta",               "suggested_for_categories": ["bebidas-sin-alcohol", "gaseosas"]},
      {"name": "Speed",               "suggested_for_categories": ["bebidas-sin-alcohol", "energizantes"]},
      {"name": "Quilmes",             "suggested_for_categories": ["bebidas-con-alcohol", "cervezas"]},
      {"name": "Brahma",              "suggested_for_categories": ["bebidas-con-alcohol", "cervezas"]},
      {"name": "Heineken",            "suggested_for_categories": ["bebidas-con-alcohol", "cervezas"]},
      {"name": "Arcor",               "suggested_for_categories": ["golosinas", "caramelos-chicles", "gomitas-malvaviscos"]},
      {"name": "Georgalos",           "suggested_for_categories": ["golosinas", "frutos-secos"]},
      {"name": "Felfort",             "suggested_for_categories": ["golosinas", "chocolates"]},
      {"name": "Milka",               "suggested_for_categories": ["golosinas", "chocolates", "alfajores"]},
      {"name": "Bagley",              "suggested_for_categories": ["galletas", "galletitas-dulces", "galletitas-saladas"]},
      {"name": "Mondelez",            "suggested_for_categories": ["galletas", "galletitas-dulces", "galletitas-saladas"]},
      {"name": "Lay''s",              "suggested_for_categories": ["snacks-salados", "papas-fritas"]},
      {"name": "Beldent",             "suggested_for_categories": ["golosinas", "caramelos-chicles"]},
      {"name": "Guaymallén",          "suggested_for_categories": ["golosinas", "alfajores"]},
      {"name": "Jorgito",             "suggested_for_categories": ["golosinas", "alfajores"]},
      {"name": "Cachafaz",            "suggested_for_categories": ["golosinas", "alfajores"]},
      {"name": "Mantecol",            "suggested_for_categories": ["golosinas"]},
      {"name": "Sugus",               "suggested_for_categories": ["golosinas", "caramelos-chicles"]},
      {"name": "Cofler",              "suggested_for_categories": ["golosinas", "chocolates"]},
      {"name": "Tita",                "suggested_for_categories": ["golosinas", "chocolates"]},
      {"name": "Rhodesia",            "suggested_for_categories": ["golosinas", "chocolates"]},
      {"name": "Mogul",               "suggested_for_categories": ["golosinas", "gomitas-malvaviscos"]},
      {"name": "Rocklets",            "suggested_for_categories": ["golosinas", "chocolates"]},
      {"name": "Billiken",            "suggested_for_categories": ["golosinas", "gomitas-malvaviscos"]},
      {"name": "Palitos de la Selva", "suggested_for_categories": ["golosinas", "caramelos-chicles"]},
      {"name": "Flynn Paff",          "suggested_for_categories": ["golosinas", "caramelos-chicles"]},
      {"name": "Butter Toffees",      "suggested_for_categories": ["golosinas", "caramelos-chicles"]},
      {"name": "Havanna",             "suggested_for_categories": ["golosinas", "alfajores"]},
      {"name": "Bonafide",            "suggested_for_categories": ["golosinas", "chocolates", "alfajores"]},
      {"name": "Pico Dulce",          "suggested_for_categories": ["golosinas", "caramelos-chicles"]}
    ]'::jsonb,
    version      = '4.0.0-golosinas-completo',
    generated_by = 'cycle-004-catalog-researcher',
    updated_at   = CURRENT_TIMESTAMP
  WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'kiosco')
    AND is_default = true;

  IF NOT FOUND THEN
    RAISE NOTICE 'Template kiosco (is_default=true) no encontrado. Ejecutar seed 021 primero.';
  ELSE
    RAISE NOTICE 'Template kiosco actualizado: 33 marcas incluyendo golosinas completas (v4.0.0).';
  END IF;
END $$;
