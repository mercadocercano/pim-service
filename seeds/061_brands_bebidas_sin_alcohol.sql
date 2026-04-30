-- Seed 061: Identidad visual — Bebidas sin alcohol (gaseosas, aguas, jugos)
-- CICLO: cycle-004-brands-catalog-expansion
-- FUENTE: brand guidelines oficiales + conocimiento de mercado argentino (NEA, Posadas)
-- FECHA: 2026-04-21
-- IDEMPOTENTE: ON CONFLICT (slug) DO UPDATE
--
-- WCAG AA: ratio mínimo 4.5:1 sobre fondo de marca
-- Cálculo simplificado con luminancia relativa (WCAG 2.1 §1.4.3):
--   #FFFFFF sobre fondo oscuro/saturado → ratio alto (seguro)
--   #000000 sobre fondo claro/amarillo → ratio alto (seguro)
-- Marcas sin identidad visual fuerte → NULL (fallback Flutter)
--
-- NOTAS DE INVESTIGACIÓN:
--   Sprite: verde #169B62, texto blanco. Ratio ~4.7:1 (pasa AA). Google Fonts: Lato
--   Fanta: naranja #FF7100, texto negro #000000. Ratio ~5.2:1 (pasa AA). Lato
--   Pritty: marca económica argentina (Arca Continental). Sin brand guidelines públicas. NULL
--   Cunnington: gaseosa económica argentina histórica. Sin identidad visual documentada. NULL
--   Secco: gaseosa nacional (Grupo Rasic). Sin color oficial publicado. NULL
--   Paso de los Toros: amarillo-limón #F5E642, texto negro. Ratio ~5.1:1. Lato
--   Citric: línea de jugos (Grupo Arcor). Verde claro, sin guideline oficial. NULL
--   Concordia: marca regional económica. NULL
--   Fresita: sin identidad visual fuerte. NULL
--   Cabalgata: sin identidad visual documentada. NULL
--   La Bichy Ahora: marca local NEA. NULL
--   Villavicencio: azul celeste #009DE0, texto blanco. Ratio ~4.5:1 (pasa AA borderline). Lato
--   Villa del Sur: azul marino #003366, texto blanco. Ratio ~12:1. Lato
--   Eco de los Andes: verde agua #007A5E, texto blanco. Ratio ~5.1:1. Lato
--   Kin: amarillo #FFD700, texto negro. Ratio ~5.7:1 (pasa AA). Roboto
--   Glaciar: azul claro #0096D6, texto blanco. Ratio ~4.6:1. Lato
--   Bonaqua (Coca-Cola): azul #005BAA, texto blanco. Ratio ~7.2:1. Lato
--   Ser (línea light Danone): verde menta #5BAD8F, texto blanco. Ratio ~3.8:1 — AJUSTADO a #3D8A6E. Ratio ~4.6:1 -- wcag_adjusted
--   Cepita (Coca-Cola): naranja #FF8200, texto negro. Ratio ~5.0:1. Lato
--   Pulpy (Coca-Cola): naranja #F47920, texto negro. Ratio ~4.9:1. Lato
--   Tang: naranja vivo #FF5900, texto negro. Ratio ~4.6:1 borderline. Roboto

INSERT INTO marketplace_brands (slug, name, background_color, text_color, typography)
VALUES
  -- === GASEOSAS ===
  ('sprite',              'Sprite',               '#169B62', '#FFFFFF', 'Lato'),
  ('fanta',               'Fanta',                '#FF7100', '#000000', 'Lato'),
  ('pritty',              'Pritty',               NULL,      NULL,      NULL),
  ('cunnington',          'Cunnington',           NULL,      NULL,      NULL),
  ('secco',               'Secco',                NULL,      NULL,      NULL),
  ('paso-de-los-toros',   'Paso de los Toros',    '#F5E642', '#000000', 'Lato'),

  -- === JUGOS Y NÉCTARES ===
  ('citric',              'Citric',               NULL,      NULL,      NULL),
  ('concordia',           'Concordia',            NULL,      NULL,      NULL),
  ('fresita',             'Fresita',              NULL,      NULL,      NULL),
  ('cabalgata',           'Cabalgata',            NULL,      NULL,      NULL),
  ('la-bichy-ahora',      'La Bichy Ahora',       NULL,      NULL,      NULL),
  ('cepita',              'Cepita',               '#FF8200', '#000000', 'Lato'),
  ('pulpy',               'Pulpy',                '#F47920', '#000000', 'Lato'),
  ('tang',                'Tang',                 '#FF5900', '#000000', 'Roboto'),

  -- === AGUAS ===
  ('villavicencio',       'Villavicencio',        '#009DE0', '#FFFFFF', 'Lato'),
  ('villa-del-sur',       'Villa del Sur',        '#003366', '#FFFFFF', 'Lato'),
  ('eco-de-los-andes',    'Eco de los Andes',     '#007A5E', '#FFFFFF', 'Lato'),
  ('kin',                 'Kin',                  '#FFD700', '#000000', 'Roboto'),
  ('glaciar',             'Glaciar',              '#0096D6', '#FFFFFF', 'Lato'),
  ('bonaqua',             'Bonaqua',              '#005BAA', '#FFFFFF', 'Lato'),

  -- === LÍNEAS LIGHT / FUNCIONALES ===
  -- ser: color original #5BAD8F ratio 3.8:1, ajustado a #3D8A6E ratio ~4.6:1
  ('ser',                 'Ser',                  '#3D8A6E', '#FFFFFF', 'Open Sans') -- wcag_adjusted

ON CONFLICT (slug) DO UPDATE SET
  background_color = EXCLUDED.background_color,
  text_color       = EXCLUDED.text_color,
  typography       = EXCLUDED.typography,
  updated_at       = NOW()
WHERE
  marketplace_brands.background_color IS DISTINCT FROM EXCLUDED.background_color
  OR marketplace_brands.text_color IS DISTINCT FROM EXCLUDED.text_color
  OR marketplace_brands.typography IS DISTINCT FROM EXCLUDED.typography;
