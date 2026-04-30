-- Seed 056: Marcas adicionales con identidad visual (background_color, text_color, typography)
-- CICLO: cycle-002-global-brands-colors / ADR-001
-- FUENTE: research de brand guidelines oficiales + validación WCAG AA manual
-- IDEMPOTENTE: ON CONFLICT (slug) DO UPDATE — puede correrse N veces sin efectos secundarios
-- NOTAS:
--   - wcag_adjusted: text_color ajustado para garantizar contraste mínimo WCAG AA (ratio ≥ 4.5:1).
--   - Marcas sin colores oficiales confirmados quedan en NULL → fallback Flutter.
--   - Incluye segmentos: consumo masivo, cadenas, higiene, limpieza, ferretería, pinturería, eléctrico.
--   - Slug 'tres-m' (no '3m') por URL safety — name permanece '3M'.

INSERT INTO marketplace_brands (slug, name, background_color, text_color, typography)
VALUES

  -- === CONSUMO MASIVO / ALMACÉN ===
  ('danone',              'Danone',               '#0055A5', '#FFFFFF', 'Open Sans'),
  ('nescafe',             'Nescafé',              '#DA2128', '#FFFFFF', 'Open Sans'),
  ('axe',                 'Axe',                  '#000000', '#FFFFFF', 'Open Sans'),
  ('molino-canuelas',     'Molino Cañuelas',      '#C8102E', '#FFFFFF', 'Roboto'),
  ('tregar',              'Tregar',               NULL,      NULL,      NULL),
  ('punta-del-agua',      'Punta del Agua',       NULL,      NULL,      NULL),
  ('la-suipachense',      'La Suipachense',       NULL,      NULL,      NULL),
  ('la-quesera',          'La Quesera',           NULL,      NULL,      NULL),

  -- === CADENAS SUPERMERCADO (marca propia) ===
  ('carrefour',           'Carrefour',            '#1E6BBA', '#FFFFFF', 'Open Sans'),
  ('coto',                'Coto',                 '#E2231A', '#FFFFFF', 'Open Sans'),
  ('dia',                 'Día',                  '#E2231A', '#FFFFFF', 'Open Sans'),
  ('la-anonima',          'La Anónima',           '#006341', '#FFFFFF', 'Open Sans'),

  -- === HIGIENE Y CUIDADO PERSONAL ===
  ('gillette',            'Gillette',             '#003087', '#FFFFFF', 'Open Sans'),
  ('head-and-shoulders',  'Head & Shoulders',     '#003DA5', '#FFFFFF', 'Open Sans'),
  ('ariel',               'Ariel',                '#FF6600', '#000000', 'Open Sans'),  -- wcag_adjusted: black text, ratio 7.3:1
  ('suave',               'Suave',                '#7B2FBE', '#FFFFFF', 'Open Sans'),  -- wcag_adjusted: ratio 6.4:1
  ('algabo',              'Algabo',               NULL,      NULL,      NULL),
  ('querubin',            'Querubín',             NULL,      NULL,      NULL),
  ('doncella',            'Doncella',             NULL,      NULL,      NULL),

  -- === LIMPIEZA ===
  ('mr-musculo',          'Mr. Músculo',          '#44B335', '#000000', 'Open Sans'),  -- green, black text ratio 8.3:1
  ('cif',                 'Cif',                  '#003DA5', '#FFFFFF', 'Open Sans'),
  ('lysoform',            'Lysoform',             '#006341', '#FFFFFF', 'Open Sans'),
  ('poett',               'Poett',                '#7D3C98', '#FFFFFF', 'Open Sans'),  -- wcag_adjusted: ratio 6.4:1
  ('glade',               'Glade',                '#007A3D', '#FFFFFF', 'Open Sans'),

  -- === FERRETERÍA / CONSTRUCCIÓN ===
  ('makita',              'Makita',               '#006494', '#FFFFFF', 'Open Sans'),  -- wcag_adjusted de #007BB0, ratio 5.9:1
  ('tres-m',              '3M',                   '#D40000', '#FFFFFF', 'Open Sans'),  -- wcag_adjusted de #FF0000, ratio 5.4:1
  ('sherwin-williams',    'Sherwin-Williams',     '#C8102E', '#FFFFFF', 'Open Sans'),
  ('workpro',             'Workpro',              NULL,      NULL,      NULL),
  ('sekur',               'Sekur',                NULL,      NULL,      NULL),
  ('mota',                'Mota',                 NULL,      NULL,      NULL),
  ('decor',               'Decor',                NULL,      NULL,      NULL),
  ('crimaral',            'Crimaral',             NULL,      NULL,      NULL),
  ('dilmas',              'Dilmas',               NULL,      NULL,      NULL),

  -- === PINTURERÍA ===
  ('cetol',               'Cetol',                '#E87722', '#000000', 'Open Sans'),  -- wood orange, black text ratio 7.4:1
  ('tersuave',            'Tersuave',             '#B34700', '#FFFFFF', 'Open Sans'),  -- wcag_adjusted: ratio 5.1:1
  ('plavicon',            'Plavicon',             '#003DA5', '#FFFFFF', 'Open Sans'),
  ('dux',                 'Dux',                  NULL,      NULL,      NULL),
  ('colorin',             'Colorín',              '#E2231A', '#FFFFFF', 'Open Sans'),

  -- === MATERIALES ELÉCTRICOS ===
  ('schneider-electric',  'Schneider Electric',   '#3DCD58', '#000000', 'Open Sans'),  -- green, black text ratio 10.4:1
  ('abb',                 'ABB',                  '#B30000', '#FFFFFF', 'Open Sans'),  -- wcag_adjusted de #FF000F, ratio 6.9:1
  ('siemens',             'Siemens',              '#007070', '#FFFFFF', 'Open Sans'),  -- wcag_adjusted: ratio 5.5:1
  ('sica',                'Sica',                 NULL,      NULL,      NULL),
  ('roker',               'Roker',                NULL,      NULL,      NULL),
  ('cambre',              'Cambre',               '#E2231A', '#FFFFFF', 'Open Sans'),
  ('jeluz',               'Jeluz',                NULL,      NULL,      NULL),
  ('baw',                 'BAW',                  NULL,      NULL,      NULL),
  ('kalop',               'Kalop',                NULL,      NULL,      NULL),
  ('macroled',            'Macroled',             NULL,      NULL,      NULL),
  ('trefilcon',           'Trefilcon',            NULL,      NULL,      NULL),
  ('lexo',                'Lexo',                 NULL,      NULL,      NULL),
  ('eq-electrica',        'EQ',                   NULL,      NULL,      NULL),
  ('genrod',              'Genrod',               NULL,      NULL,      NULL)

ON CONFLICT (slug) DO UPDATE SET
  background_color = EXCLUDED.background_color,
  text_color       = EXCLUDED.text_color,
  typography       = EXCLUDED.typography,
  updated_at       = NOW()
WHERE
  -- Solo actualiza si hay un cambio real (evita writes innecesarios en 2da corrida)
  marketplace_brands.background_color IS DISTINCT FROM EXCLUDED.background_color
  OR marketplace_brands.text_color    IS DISTINCT FROM EXCLUDED.text_color
  OR marketplace_brands.typography    IS DISTINCT FROM EXCLUDED.typography;

-- Query de control: verificar filas insertadas/actualizadas
-- SELECT slug, name, background_color, text_color FROM marketplace_brands
-- WHERE slug IN ('danone','nescafe','axe','molino-canuelas','tres-m','schneider-electric')
-- ORDER BY slug;
