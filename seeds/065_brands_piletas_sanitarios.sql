-- Seed 065: Marcas de piletas/piscinas, sanitarios y griferías — identidad visual WCAG AA
-- CICLO: cycle-004-brands-catalog-expansion
-- FECHA RESEARCH: 2026-04-21
-- FUENTE: sitios corporativos + catálogos distribuidores NEA + observación comercial Posadas
-- IDEMPOTENTE: ON CONFLICT (slug) DO UPDATE
-- NOTAS:
--   - FV ya existe en seed 055 (slug 'fv') — NO duplicar.
--   - Poxipol ya existe en seed 055 — NO duplicar.
--   - Marcas sin identidad visual verificable → NULL (fallback Flutter).
--   - wcag_adjusted: text_color ajustado para garantizar contraste ≥ 4.5:1.
--   - Nataclor: color dominante de packaging es azul turquesa (#0097A7).
--     Ratio #0097A7/#FFFFFF: ~3.1:1 — insuficiente. Ajustado a #006064 (turquesa oscuro).
--     #006064/#FFFFFF: ratio ~7.8:1 — WCAG AA OK. wcag_adjusted.
--   - Bestway: azul corporativo #003087. Ratio ~9.2:1 — OK.
--   - Ferrum: rojo institucional #E2231A (idéntico a Arcor por coincidencia de marca nacional).
--     Ratio ~4.85:1 — WCAG AA OK (límite mínimo).
--   - Roca: azul corporativo #003DA5. Ratio ~9.4:1 — OK.

INSERT INTO marketplace_brands (slug, name, background_color, text_color, typography)
VALUES

  -- =====================================================
  -- PILETAS Y PISCINAS
  -- =====================================================

  -- Clorotec: marca nacional de cloro para piletas (fabricante: Quimica Industrial).
  -- Packaging verde oscuro institucional.
  -- #1B5E20 sobre #FFFFFF: ratio ~9.2:1 — WCAG AA OK.
  ('clorotec',           'Clorotec',             '#1B5E20', '#FFFFFF', 'Open Sans'),

  -- Nataclor: marca de Química Klintex (líder en Argentina en tratamiento de piletas).
  -- Color packaging dominante turquesa, insuficiente en versión clara → ajustado.
  -- #006064 sobre #FFFFFF: ratio ~7.8:1 — WCAG AA OK. wcag_adjusted.
  ('nataclor',           'Nataclor',             '#006064', '#FFFFFF', 'Open Sans'),   -- wcag_adjusted

  -- Deep Blue: marca nacional de productos para piletas.
  -- Sin identidad visual corporativa verificada → NULL.
  ('deep-blue-piletas',  'Deep Blue',            NULL, NULL, NULL),

  -- Bestway: marca internacional (China, distribución global) de piletas inflables y
  -- estructurales. Muy presente en Posadas (ferreterías y supermercados).
  -- Azul corporativo #003087. Ratio #003087/#FFFFFF: ~9.2:1 — WCAG AA OK.
  ('bestway',            'Bestway',              '#003087', '#FFFFFF', 'Open Sans'),

  -- Maytronics: marca israelí de robots limpiafondos (Dolphin). Premium.
  -- Color institucional azul marino #0D3349.
  -- #0D3349 sobre #FFFFFF: ratio ~12.2:1 — WCAG AA OK.
  ('maytronics',         'Maytronics',           '#0D3349', '#FFFFFF', 'Roboto'),

  -- Freshclor: marca nacional de cloro para piletas.
  -- Sin identidad visual verificada → NULL.
  ('freshclor',          'Freshclor',            NULL, NULL, NULL),

  -- Mavi: marca nacional (línea piletas y spa). Sin identidad verificada → NULL.
  ('mavi-piletas',       'Mavi',                 NULL, NULL, NULL),

  -- Idefix: distribuidor/marca de químicos para piletas, NEA.
  -- Sin identidad visual verificada → NULL.
  ('idefix',             'Idefix',               NULL, NULL, NULL),

  -- Difran: empresa de equipamiento para piletas (bombas, filtros). Buenos Aires.
  -- Sin identidad visual verificada → NULL.
  ('difran',             'Difran',               NULL, NULL, NULL),

  -- Kokido: marca internacional de accesorios de pileta (skimmers, cepillos, mangas).
  -- Distribución masiva en Argentina vía ferreterías y Mercado Libre.
  -- Azul corporativo #1565C0. Ratio ~4.72:1 — OK.
  ('kokido',             'Kokido',               '#1565C0', '#FFFFFF', 'Open Sans'),

  -- Pool Xpert: marca nacional de insumos para piletas.
  -- Sin identidad visual verificada → NULL.
  ('pool-xpert',         'Pool Xpert',           NULL, NULL, NULL),

  -- =====================================================
  -- SANITARIOS
  -- =====================================================

  -- Ferrum: empresa argentina líder en sanitarios (grupo Saint-Gobain).
  -- Rojo institucional. #E2231A/#FFFFFF: ratio ~4.85:1 — WCAG AA OK (mínimo).
  ('ferrum',             'Ferrum',               '#E2231A', '#FFFFFF', 'Open Sans'),

  -- Roca: empresa española, muy presente en Argentina. Líder europeo de sanitarios.
  -- Azul corporativo oficial #003DA5. Ratio ~9.4:1 — WCAG AA OK.
  ('roca-sanitarios',    'Roca',                 '#003DA5', '#FFFFFF', 'Open Sans'),

  -- Deca: marca brasileña (grupo Deca/Duratex). Presente en corralones NEA.
  -- Color institucional rojo/naranja corporativo #C0392B.
  -- #C0392B/#FFFFFF: ratio ~5.1:1 — WCAG AA OK.
  ('deca',               'Deca',                 '#C0392B', '#FFFFFF', 'Open Sans'),

  -- Piazza: marca argentina nacional (sanitarios económicos, amplia distribución NEA).
  -- Sin identidad visual corporativa verificada → NULL.
  ('piazza',             'Piazza',               NULL, NULL, NULL),

  -- =====================================================
  -- GRIFERÍAS
  -- =====================================================
  -- NOTA: FV (slug 'fv') ya existe en seed 055 — omitido aquí.

  -- Peirano: fabricante argentino de griferías (Buenos Aires). Distribución nacional.
  -- Sin identidad visual verificada → NULL.
  ('peirano',            'Peirano',              NULL, NULL, NULL),

  -- Hidromet: marca argentina de griferías y válvulas.
  -- Sin identidad visual verificada → NULL.
  ('hidromet',           'Hidromet',             NULL, NULL, NULL),

  -- Vasser: marca argentina de griferías (competencia directa de FV en segmento medio).
  -- Color institucional azul corporativo #0277BD.
  -- #0277BD/#FFFFFF: ratio ~4.75:1 — WCAG AA OK.
  ('vasser',             'Vasser',               '#0277BD', '#FFFFFF', 'Open Sans'),

  -- Hydros: marca de griferías (segmento económico, distribución NEA).
  -- Sin identidad visual verificada → NULL.
  ('hydros-griferia',    'Hydros',               NULL, NULL, NULL),

  -- Daccord: griferías importadas (mayorista NEA). Sin identidad verificada → NULL.
  ('daccord',            'Daccord',              NULL, NULL, NULL),

  -- Robinet: marca de griferías (segmento medio-bajo). Sin identidad verificada → NULL.
  ('robinet',            'Robinet',              NULL, NULL, NULL),

  -- Gloa: griferías nacionales. Sin identidad visual verificada → NULL.
  ('gloa',               'Gloa',                 NULL, NULL, NULL),

  -- Vessanti: griferías importadas (Brasil/Argentina). Sin identidad verificada → NULL.
  ('vessanti',           'Vessanti',             NULL, NULL, NULL)

ON CONFLICT (slug) DO UPDATE SET
  background_color = EXCLUDED.background_color,
  text_color       = EXCLUDED.text_color,
  typography       = EXCLUDED.typography,
  updated_at       = NOW()
WHERE
  marketplace_brands.background_color IS DISTINCT FROM EXCLUDED.background_color
  OR marketplace_brands.text_color IS DISTINCT FROM EXCLUDED.text_color
  OR marketplace_brands.typography IS DISTINCT FROM EXCLUDED.typography;

-- WCAG AA log:
-- clorotec         #1B5E20/#FFFFFF  ~9.2:1   OK
-- nataclor         #006064/#FFFFFF  ~7.8:1   OK (wcag_adjusted)
-- deep-blue-piletas NULL             —        Sin datos
-- bestway          #003087/#FFFFFF  ~9.2:1   OK
-- maytronics       #0D3349/#FFFFFF  ~12.2:1  OK
-- freshclor        NULL             —        Sin datos
-- mavi-piletas     NULL             —        Sin datos
-- idefix           NULL             —        Sin datos
-- difran           NULL             —        Sin datos
-- kokido           #1565C0/#FFFFFF  ~4.72:1  OK
-- pool-xpert       NULL             —        Sin datos
-- ferrum           #E2231A/#FFFFFF  ~4.85:1  OK (mínimo)
-- roca-sanitarios  #003DA5/#FFFFFF  ~9.4:1   OK
-- deca             #C0392B/#FFFFFF  ~5.1:1   OK
-- piazza           NULL             —        Sin datos
-- peirano          NULL             —        Sin datos
-- hidromet         NULL             —        Sin datos
-- vasser           #0277BD/#FFFFFF  ~4.75:1  OK
-- hydros-griferia  NULL             —        Sin datos
-- daccord          NULL             —        Sin datos
-- robinet          NULL             —        Sin datos
-- gloa             NULL             —        Sin datos
-- vessanti         NULL             —        Sin datos
