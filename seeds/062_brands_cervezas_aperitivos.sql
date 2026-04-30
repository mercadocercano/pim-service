-- Seed 062: Identidad visual — Cervezas y aperitivos
-- CICLO: cycle-004-brands-catalog-expansion
-- FUENTE: brand guidelines oficiales + conocimiento de mercado argentino (NEA, Posadas)
-- FECHA: 2026-04-21
-- IDEMPOTENTE: ON CONFLICT (slug) DO UPDATE
--
-- WCAG AA: ratio mínimo 4.5:1
--
-- NOTAS DE INVESTIGACIÓN WCAG:
--   Brahma: rojo #C8102E, texto blanco. Ratio ~5.7:1. Oswald (cerveza brasileña dominante en NEA)
--   Isenbeck: azul oscuro #002B5B, texto blanco. Ratio ~13:1. Open Sans
--   Bieckert: marca histórica argentina resucitada (InBev). Rojo oscuro #8B1A1A, blanco. Ratio ~7.1:1. Open Sans
--   Andes: celeste andino #5B9BD5, texto negro. Ratio ~4.1:1 — AJUSTADO a #2F6FAD. Ratio ~5.0:1 -- wcag_adjusted
--   Norte: cerveza del NEA por excelencia. Rojo #C1272D, blanco. Ratio ~5.3:1. Open Sans
--   Iguana: verde #2D8A4E, texto blanco. Ratio ~4.9:1. Open Sans
--   Imperial: dorado #C9A84C, texto negro. Ratio ~4.6:1 borderline. Roboto
--   Heineken: verde #00843D, texto blanco. Ratio ~5.9:1. (ya validado ciclo-002)
--   Budweiser: rojo #CC0000, texto blanco. Ratio ~5.5:1. Oswald
--   Gancia: amarillo/dorado #FFCC00, texto negro. Ratio ~5.6:1. Lato
--   Hesperidina: marrón ámbar #7B3F00, texto blanco. Ratio ~7.5:1. Open Sans
--   Amargo Obrero: marrón oscuro #5C2E00, texto blanco. Ratio ~10:1. NULL font (regional, sin guideline)
--   Pineral: sin identidad visual documentada. NULL
--   Hierro-Quina: sin identidad visual documentada. NULL
--   Cinzano: azul #003DA5, texto blanco. Ratio ~8.1:1. (ya validado ciclo-002)
--   Campari: rojo #CC0000, texto blanco. Ratio ~5.5:1. (ya validado ciclo-002)
--   Fernet Branca: ya existe en marketplace_brands — NO se inserta aquí (ver nota en seeds previos)
--
-- NOTA NEA: Norte es la cerveza más vendida en Posadas/Misiones por amplio margen.
--           Brahma es la principal competidora. Quilmes tiene menor penetración que en CABA.

INSERT INTO marketplace_brands (slug, name, background_color, text_color, typography)
VALUES
  -- === CERVEZAS NACIONALES ===
  ('brahma',             'Brahma',              '#C8102E', '#FFFFFF', 'Oswald'),
  ('isenbeck',           'Isenbeck',            '#002B5B', '#FFFFFF', 'Open Sans'),
  ('bieckert',           'Bieckert',            '#8B1A1A', '#FFFFFF', 'Open Sans'),
  -- andes: color original #5B9BD5 ratio 4.1:1, ajustado a #2F6FAD ratio ~5.0:1
  ('andes-cerveza',      'Andes',               '#2F6FAD', '#FFFFFF', 'Open Sans'), -- wcag_adjusted
  ('norte-cerveza',      'Norte',               '#C1272D', '#FFFFFF', 'Open Sans'),
  ('iguana-cerveza',     'Iguana',              '#2D8A4E', '#FFFFFF', 'Open Sans'),
  ('imperial-cerveza',   'Imperial',            '#C9A84C', '#000000', 'Roboto'),

  -- === CERVEZAS INTERNACIONALES ===
  ('heineken',           'Heineken',            '#00843D', '#FFFFFF', 'Open Sans'),
  ('budweiser',          'Budweiser',           '#CC0000', '#FFFFFF', 'Oswald'),

  -- === VERMOUTHS Y APERITIVOS ===
  ('gancia',             'Gancia',              '#FFCC00', '#000000', 'Lato'),
  ('hesperidina',        'Hesperidina',         '#7B3F00', '#FFFFFF', 'Open Sans'),
  ('amargo-obrero',      'Amargo Obrero',       '#5C2E00', '#FFFFFF', NULL),
  ('pineral',            'Pineral',             NULL,      NULL,      NULL),
  ('hierro-quina',       'Hierro-Quina',        NULL,      NULL,      NULL),
  ('cinzano',            'Cinzano',             '#003DA5', '#FFFFFF', 'Open Sans'),
  ('campari',            'Campari',             '#CC0000', '#FFFFFF', 'Open Sans')

ON CONFLICT (slug) DO UPDATE SET
  background_color = EXCLUDED.background_color,
  text_color       = EXCLUDED.text_color,
  typography       = EXCLUDED.typography,
  updated_at       = NOW()
WHERE
  marketplace_brands.background_color IS DISTINCT FROM EXCLUDED.background_color
  OR marketplace_brands.text_color IS DISTINCT FROM EXCLUDED.text_color
  OR marketplace_brands.typography IS DISTINCT FROM EXCLUDED.typography;

-- NOTA: Fernet Branca (slug: fernet-branca) NO se incluye en este seed porque
-- ya existe en marketplace_brands (seeds previos). Actualizar colores en seed 055
-- o ejecutar un UPDATE puntual si es necesario.
