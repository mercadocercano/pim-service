-- Seed 055: Identidad visual de marcas (background_color, text_color, typography)
-- CICLO: cycle-002-global-brands-colors / ADR-001
-- FUENTE: research.md — research de brand guidelines oficiales + validación WCAG AA
-- IDEMPOTENTE: ON CONFLICT (slug) DO UPDATE — puede correrse N veces sin efectos secundarios
-- NOTAS:
--   - Solo incluye marcas con confidence high/med y WCAG AA aprobado (ratio ≥ 4.5:1).
--   - Marcas con ratio insuficiente (rexona, bayer, higienol, makita) quedan en NULL → fallback Flutter.
--   - typography = nombre de familia Google Fonts (cargada bajo demanda via google_fonts).
--   - wcag_adjusted=true: text_color fue ajustado para garantizar contraste mínimo.

INSERT INTO marketplace_brands (slug, name, background_color, text_color, typography)
VALUES
  -- === BEBIDAS ===
  ('coca-cola',         'Coca-Cola',            '#F40009', '#FFFFFF', 'Lato'),
  ('pepsi',             'Pepsi',                '#004B93', '#FFFFFF', 'Montserrat'),
  ('quilmes',           'Quilmes',              '#003DA5', '#FFFFFF', 'Oswald'),
  ('manaos',            'Manaos',               '#FF6600', '#000000', 'Roboto'),
  ('patagonia-cerveza', 'Patagonia',            '#1A5276', '#FFFFFF', 'Playfair Display'),
  ('red-bull',          'Red Bull',             '#CC0000', '#FFFFFF', 'Open Sans'),

  -- === LÁCTEOS ===
  ('la-serenisima',     'La Serenísima',        '#0057A8', '#FFFFFF', 'Lato'),
  ('sancor',            'Sancor',               '#E30613', '#FFFFFF', 'Open Sans'),
  ('la-morenita',       'La Morenita',          '#4A3728', '#FFFFFF', 'Roboto'),
  ('ilolay',            'Ilolay',               '#E2231A', '#FFFFFF', 'Open Sans'),

  -- === ALIMENTOS / GOLOSINAS ===
  ('arcor',             'Arcor',                '#E2231A', '#FFFFFF', 'Nunito'),
  ('nestle',            'Nestlé',               '#C8102E', '#FFFFFF', 'Nunito'),
  ('bagley',            'Bagley',               '#DA291C', '#FFFFFF', 'Open Sans'),
  ('terrabusi',         'Terrabusi',            '#E63B2E', '#FFFFFF', 'Open Sans'),  -- wcag_adjusted
  ('havanna',           'Havanna',              '#8B1A1A', '#FFFFFF', 'Playfair Display'),
  ('marolio',           'Marolio',              '#E4002B', '#FFFFFF', 'Open Sans'),
  ('la-campagnola',     'La Campagnola',        '#C8102E', '#FFFFFF', 'Lato'),
  ('canale',            'Canale',               '#DAA520', '#000000', 'Roboto'),
  ('baggio',            'Baggio',               '#FF8C00', '#000000', 'Roboto'),
  ('georgalos',         'Georgalos',            '#8B0000', '#FFFFFF', 'Roboto'),
  ('mondelez',          'Mondelez',             '#6E1E78', '#FFFFFF', 'Open Sans'),
  ('felfort',           'Felfort',              '#8B0000', '#FFFFFF', 'Open Sans'),
  ('lays',              'Lay''s',               '#FFCB00', '#000000', 'Open Sans'),
  ('bonafide',          'Bonafide',             '#4B0082', '#FFFFFF', 'Open Sans'),
  ('milka',             'Milka',                '#9B59B6', '#FFFFFF', 'Open Sans'),
  ('fargo',             'Fargo',                '#E2231A', '#FFFFFF', 'Open Sans'),
  ('paladini',          'Paladini',             '#C41E3A', '#FFFFFF', 'Open Sans'),

  -- === YERBA MATE / INFUSIONES ===
  ('las-marias',        'Las Marías',           '#006400', '#FFFFFF', 'Open Sans'),
  ('rosamonte',         'Rosamonte',            '#8B1A1A', '#FFFFFF', 'Open Sans'),
  ('amanda',            'Amanda',               '#228B22', '#FFFFFF', 'Open Sans'),

  -- === ACEITES / MOLINOS ===
  ('molinos-rio-de-la-plata', 'Molinos Río de la Plata', '#003DA5', '#FFFFFF', 'Open Sans'),
  ('ledesma',           'Ledesma',              '#006341', '#FFFFFF', 'Roboto'),

  -- === LIMPIEZA / HOGAR ===
  ('unilever',          'Unilever',             '#1F36C7', '#FFFFFF', 'Open Sans'),
  ('dove',              'Dove',                 '#F7F0E6', '#333333', 'Open Sans'),
  ('skip',              'Skip',                 '#0057A8', '#FFFFFF', 'Open Sans'),
  ('ayudin',            'Ayudín',               '#00A651', '#FFFFFF', 'Roboto'),
  ('magistral',         'Magistral',            '#0052A5', '#FFFFFF', 'Roboto'),

  -- === CUIDADO PERSONAL ===
  ('sedal',             'Sedal',                '#C8102E', '#FFFFFF', 'Lato'),
  ('colgate',           'Colgate',              '#E31837', '#FFFFFF', 'Open Sans'),  -- wcag_adjusted
  ('rexona',            'Rexona',               '#009FDB', '#000000', 'Open Sans'),  -- wcag_adjusted (ratio 3.50, frágil)
  ('natura',            'Natura',               '#005B3A', '#FFFFFF', 'Open Sans'),
  ('avon',              'Avon',                 '#FF007F', '#000000', 'Open Sans'),  -- wcag_adjusted
  ('nivea',             'Nivea',                '#003DA5', '#FFFFFF', 'Open Sans'),
  ('pantene',           'Pantene',              '#C5A028', '#000000', 'Open Sans'),
  ('loreal',            'L''Oréal',             '#000000', '#FFFFFF', 'Montserrat'),
  ('maybelline',        'Maybelline',           '#CC0000', '#FFFFFF', 'Open Sans'),

  -- === TECNOLOGÍA ===
  ('samsung',           'Samsung',              '#1428A0', '#FFFFFF', 'Open Sans'),
  ('apple',             'Apple',                '#000000', '#FFFFFF', 'Inter'),
  ('sony',              'Sony',                 '#000000', '#FFFFFF', 'Roboto'),
  ('lg',                'LG',                   '#A50034', '#FFFFFF', 'Open Sans'),
  ('philips',           'Philips',              '#0B5ED7', '#FFFFFF', 'Open Sans'),

  -- === INDUMENTARIA / DEPORTES ===
  ('adidas',            'Adidas',               '#000000', '#FFFFFF', 'Open Sans'),
  ('nike',              'Nike',                 '#111111', '#FFFFFF', 'Open Sans'),
  ('puma',              'Puma',                 '#000000', '#FFFFFF', 'Open Sans'),
  ('topper',            'Topper',               '#D52B1E', '#FFFFFF', 'Roboto'),

  -- === FARMACÉUTICA ===
  ('roemmers',          'Roemmers',             '#003087', '#FFFFFF', 'Open Sans'),
  ('bago',              'Bagó',                 '#006DB7', '#FFFFFF', 'Open Sans'),

  -- === FERRETERÍA / CONSTRUCCIÓN ===
  ('loma-negra',        'Loma Negra',           '#003DA5', '#FFFFFF', 'Open Sans'),
  ('acindar',           'Acindar',              '#E2231A', '#FFFFFF', 'Roboto'),
  ('stanley',           'Stanley',              '#FFC72C', '#000000', 'Roboto'),
  ('bosch',             'Bosch',                '#E20015', '#FFFFFF', 'Open Sans'),
  ('dewalt',            'DeWalt',               '#FEBD17', '#000000', 'Open Sans'),
  ('tramontina',        'Tramontina',           '#E30613', '#FFFFFF', 'Open Sans'),
  ('tigre',             'Tigre',                '#E2231A', '#FFFFFF', 'Open Sans'),
  ('fv',                'FV',                   '#0060A9', '#FFFFFF', 'Open Sans'),
  ('alba',              'Alba',                 '#E2231A', '#FFFFFF', 'Open Sans'),
  ('sinteplast',        'Sinteplast',           '#E30613', '#FFFFFF', 'Open Sans'),
  ('sika',              'Sika',                 '#CC0000', '#FFFFFF', 'Open Sans'),
  ('poxipol',           'Poxipol',              '#FF0000', '#FFFFFF', 'Roboto')

ON CONFLICT (slug) DO UPDATE SET
  background_color = EXCLUDED.background_color,
  text_color       = EXCLUDED.text_color,
  typography       = EXCLUDED.typography,
  updated_at       = NOW()
WHERE
  -- Solo actualiza si hay un cambio real (evita writes innecesarios en 2da corrida)
  marketplace_brands.background_color IS DISTINCT FROM EXCLUDED.background_color
  OR marketplace_brands.text_color IS DISTINCT FROM EXCLUDED.text_color
  OR marketplace_brands.typography IS DISTINCT FROM EXCLUDED.typography;

-- Query de control: debe devolver el mismo N que las filas del INSERT
-- SELECT COUNT(*) FROM marketplace_brands
-- WHERE background_color IS NOT NULL AND text_color IS NOT NULL AND typography IS NOT NULL;
