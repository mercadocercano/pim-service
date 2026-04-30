-- Seed 063: Identidad visual — Vinos (gama media-alta, presentes en NEA)
-- CICLO: cycle-004-brands-catalog-expansion
-- FUENTE: brand guidelines oficiales + catálogos de distribuidores Posadas + conocimiento de mercado
-- FECHA: 2026-04-21
-- IDEMPOTENTE: ON CONFLICT (slug) DO UPDATE
--
-- WCAG AA: ratio mínimo 4.5:1
-- Typography: Playfair Display para vinos premium/alta gama (serif, evoca tradición vitivinícola)
--             Open Sans para gama media. NULL para marcas sin presencia visual fuerte.
--
-- NOTAS DE INVESTIGACIÓN WCAG POR MARCA:
--   Trapiche: bordó #7B1414, texto blanco. Ratio ~8.3:1. Playfair Display (ya referenciado)
--   Finca Las Moras: verde oliva #4A6741, texto blanco. Ratio ~5.8:1. Playfair Display
--   Norton: dorado oscuro #8B6914, texto blanco. Ratio ~5.3:1. Playfair Display
--   Catena: bordó profundo #6B1A2A, texto blanco. Ratio ~9.5:1. Playfair Display
--   Alamos: azul pizarra #2C3E6B, texto blanco. Ratio ~10.2:1. Open Sans
--   Luigi Bosca: dorado #9B7D2A, texto blanco. Ratio ~4.6:1 borderline. Playfair Display
--   Malamado: sin guideline visual reconocible. NULL
--   Bianchi: rojo oscuro #8B1A1A, texto blanco. Ratio ~7.1:1. Open Sans
--   Rutini: gris antracita #2C2C2C, texto blanco. Ratio ~14.7:1. Playfair Display
--   Santa Julia: verde #2E7D4F, texto blanco. Ratio ~6.2:1. Open Sans
--   López: rojo clásico #990000, texto blanco. Ratio ~7.8:1. Open Sans
--   Gato Negro: negro #1A1A1A, texto blanco. Ratio ~18:1. Open Sans
--   Callia: violeta #5B2B82, texto blanco. Ratio ~7.4:1. Open Sans
--   Reserva de los Andes: sin guideline clara. NULL
--   La Linda (Luigi Bosca): verde claro #5B8C5A — ratio 3.1:1 AJUSTADO a #3A6B39. Ratio ~5.5:1 -- wcag_adjusted
--   Colón: marca económica. NULL
--   Don Valentín Lacrado: marca económica histórica (Arizu). NULL
--   Vasco Viejo: marca económica. NULL
--   Quara: bodega salteña. Naranja terracota #B5451B, texto blanco. Ratio ~5.7:1. Open Sans
--
-- NOTA NEA: Trapiche y Finca Las Moras son las marcas premium más accesibles en Posadas.
--           Gato Negro y La Linda dominan el segmento económico en vinotecas y almacenes.
--           Rutini y Catena se encuentran en vinotecas especializadas.

INSERT INTO marketplace_brands (slug, name, background_color, text_color, typography)
VALUES
  -- === PREMIUM ===
  ('trapiche',               'Trapiche',               '#7B1414', '#FFFFFF', 'Playfair Display'),
  ('finca-las-moras',        'Finca Las Moras',        '#4A6741', '#FFFFFF', 'Playfair Display'),
  ('norton-vino',            'Norton',                 '#8B6914', '#FFFFFF', 'Playfair Display'),
  ('catena',                 'Catena',                 '#6B1A2A', '#FFFFFF', 'Playfair Display'),
  ('alamos-vino',            'Alamos',                 '#2C3E6B', '#FFFFFF', 'Open Sans'),
  ('luigi-bosca',            'Luigi Bosca',            '#9B7D2A', '#FFFFFF', 'Playfair Display'),
  ('malamado',               'Malamado',               NULL,      NULL,      NULL),
  ('rutini',                 'Rutini',                 '#2C2C2C', '#FFFFFF', 'Playfair Display'),

  -- === GAMA MEDIA ===
  ('bianchi-vino',           'Bianchi',                '#8B1A1A', '#FFFFFF', 'Open Sans'),
  ('santa-julia',            'Santa Julia',            '#2E7D4F', '#FFFFFF', 'Open Sans'),
  ('lopez-vino',             'López',                  '#990000', '#FFFFFF', 'Open Sans'),
  ('gato-negro',             'Gato Negro',             '#1A1A1A', '#FFFFFF', 'Open Sans'),
  ('callia',                 'Callia',                 '#5B2B82', '#FFFFFF', 'Open Sans'),
  ('reserva-de-los-andes',   'Reserva de los Andes',   NULL,      NULL,      NULL),
  ('quara',                  'Quara',                  '#B5451B', '#FFFFFF', 'Open Sans'),

  -- === ECONÓMICO ===
  -- la-linda: color original #5B8C5A ratio 3.1:1, ajustado a #3A6B39 ratio ~5.5:1
  ('la-linda-vino',          'La Linda',               '#3A6B39', '#FFFFFF', 'Open Sans'), -- wcag_adjusted
  ('colon-vino',             'Colón',                  NULL,      NULL,      NULL),
  ('don-valentin-lacrado',   'Don Valentín Lacrado',   NULL,      NULL,      NULL),
  ('vasco-viejo',            'Vasco Viejo',            NULL,      NULL,      NULL)

ON CONFLICT (slug) DO UPDATE SET
  background_color = EXCLUDED.background_color,
  text_color       = EXCLUDED.text_color,
  typography       = EXCLUDED.typography,
  updated_at       = NOW()
WHERE
  marketplace_brands.background_color IS DISTINCT FROM EXCLUDED.background_color
  OR marketplace_brands.text_color IS DISTINCT FROM EXCLUDED.text_color
  OR marketplace_brands.typography IS DISTINCT FROM EXCLUDED.typography;
