-- Seed 077: Marcas faltantes para cycle-005 — Klaukol, Megaflex, Black+Decker
-- CICLO: cycle-005
-- FECHA: 2026-04-24
-- FUENTE: sitios corporativos + catálogos distribuidores NEA 2026
-- IDEMPOTENTE: ON CONFLICT (slug) DO NOTHING
-- REQUIERE: tabla marketplace_brands existente (seeds anteriores)
-- NOTAS WCAG AA:
--   Klaukol: rojo institucional #E30613 sobre blanco → ratio ~4.6:1 — WCAG AA OK (mínimo).
--   Megaflex: azul institucional #005A9C sobre blanco → ratio ~6.8:1 — WCAG AA OK.
--   Black+Decker: amarillo corporativo #FFC600 sobre negro → ratio ~9.3:1 — WCAG AA OK.

INSERT INTO marketplace_brands (slug, name, background_color, text_color, typography)
VALUES

  -- Klaukol: empresa argentina líder en adhesivos para revestimientos cerámicos y porcelanato.
  -- Muy presente en corralones y ferreterías de todo el país, incluido NEA.
  -- Rojo institucional #E30613. Ratio #E30613/#FFFFFF: ~4.6:1 — WCAG AA OK (mínimo).
  ('klaukol',       'Klaukol',       '#E30613', '#FFFFFF', 'Open Sans'),

  -- Megaflex: empresa argentina especializada en membranas impermeabilizantes y techos.
  -- Líder nacional en el segmento. Muy distribuida en corralones NEA.
  -- Azul institucional #005A9C. Ratio #005A9C/#FFFFFF: ~6.8:1 — WCAG AA OK.
  ('megaflex',      'Megaflex',      '#005A9C', '#FFFFFF', 'Open Sans'),

  -- Black+Decker: marca estadounidense (Stanley Black & Decker) de herramientas eléctricas
  -- y manuales. Amplia distribución en ferreterías argentinas.
  -- Amarillo corporativo #FFC600 sobre negro. Ratio ~9.3:1 — WCAG AA OK.
  ('black-decker',  'Black+Decker',  '#FFC600', '#000000', 'Open Sans')

ON CONFLICT (slug) DO NOTHING;

-- WCAG AA log:
-- klaukol      #E30613/#FFFFFF  ~4.6:1   OK (mínimo)
-- megaflex     #005A9C/#FFFFFF  ~6.8:1   OK
-- black-decker #FFC600/#000000  ~9.3:1   OK
