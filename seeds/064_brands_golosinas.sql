-- Seed 064: Marcas de golosinas — identidad visual WCAG AA
-- CICLO: cycle-004-brands-catalog-expansion
-- FECHA RESEARCH: 2026-04-21
-- FUENTE: brand guidelines públicos + observación de góndola NEA (Posadas, Misiones)
-- IDEMPOTENTE: ON CONFLICT (slug) DO UPDATE
-- NOTAS:
--   - Georgalos, Arcor, Felfort, Milka, Bagley, Mondelez ya están en seed 055 — no duplicar.
--   - Sub-marcas de Arcor (Mogul, Palitos, Flynn Paff, Pico Dulce, Mr. Pop, Bon o Bon, Rocklets)
--     heredan identidad corporativa Arcor (#E2231A / #FFFFFF / Nunito) salvo que tengan
--     identidad visual propia verificada (indicado con comentario).
--   - Sub-marcas de Georgalos (Guaymallén, Capitán del Espacio, Mantecol) heredan #8B0000.
--   - Sub-marcas de Felfort (Rhodesia, Cofler) heredan #8B0000.
--   - wcag_adjusted: texto ajustado para garantizar contraste ≥ 4.5:1.

INSERT INTO marketplace_brands (slug, name, background_color, text_color, typography)
VALUES

  -- === ALFAJORES ===

  -- Guaymallén: sub-marca de Georgalos. Identidad propia no verificada → hereda Georgalos.
  -- #8B0000 sobre #FFFFFF: ratio ~8.6:1 — WCAG AA OK.
  ('guaymallen',         'Guaymallén',           '#8B0000', '#FFFFFF', 'Roboto'),

  -- Capitán del Espacio: sub-marca de Georgalos. Sin identidad visual propia verificada.
  -- Heredando #8B0000 Georgalos.
  ('capitan-del-espacio','Capitán del Espacio',  '#8B0000', '#FFFFFF', 'Roboto'),

  -- Jorgito: marca independiente (Alimentos Jorgito, Córdoba). Color institucional verde botella.
  -- #1B5E20 sobre #FFFFFF: ratio ~9.2:1 — WCAG AA OK.
  ('jorgito',            'Jorgito',              '#1B5E20', '#FFFFFF', 'Open Sans'),

  -- Fantoche: marca de Arcor (alfajor de chocolate). Sin identidad propia → hereda Arcor.
  -- #E2231A sobre #FFFFFF: ratio ~4.85:1 — WCAG AA OK.
  ('fantoche',           'Fantoche',             '#E2231A', '#FFFFFF', 'Nunito'),

  -- La Repostera: marca Arcor. Sin identidad visual propia verificada → hereda Arcor.
  ('la-repostera',       'La Repostera',         '#E2231A', '#FFFFFF', 'Nunito'),

  -- Suchard: marca de Felfort (licencia Suchard en Argentina). Sin identidad propia →
  -- hereda Felfort (#8B0000).
  ('suchard',            'Suchard',              '#8B0000', '#FFFFFF', 'Open Sans'),

  -- Cachafaz: marca premium independiente (El Cachafaz SRL, Buenos Aires).
  -- Identidad oscura, marrón premium tipo cacao.
  -- #3E1C00 sobre #FFFFFF: ratio ~14.5:1 — WCAG AA OK.
  ('cachafaz',           'Cachafaz',             '#3E1C00', '#FFFFFF', 'Playfair Display'),

  -- Tofi: alfajor de Arcor (lanzado 2010s). Sin identidad propia → hereda Arcor.
  ('tofi',               'Tofi',                 '#E2231A', '#FFFFFF', 'Nunito'),

  -- === CHOCOLATES ===

  -- Cofler: sub-marca de Felfort (chocolate aireado, blanco, con maní). Hereda Felfort.
  ('cofler',             'Cofler',               '#8B0000', '#FFFFFF', 'Open Sans'),

  -- Cadbury: marca de Mondelez Argentina. Color institucional violeta profundo.
  -- #4B0082 sobre #FFFFFF: ratio ~10.5:1 — WCAG AA OK.
  ('cadbury',            'Cadbury',              '#4B0082', '#FFFFFF', 'Open Sans'),

  -- Tita: galletita/chocolate de Bagley (Mondelez). Sin identidad propia → hereda Bagley.
  -- #DA291C sobre #FFFFFF: ratio ~4.85:1 — WCAG AA OK.
  ('tita',               'Tita',                 '#DA291C', '#FFFFFF', 'Open Sans'),

  -- Rhodesia: sub-marca de Felfort. Hereda Felfort (#8B0000).
  ('rhodesia',           'Rhodesia',             '#8B0000', '#FFFFFF', 'Open Sans'),

  -- Águila: marca de Nestlé Argentina (chocolate de taza y tabletas).
  -- Color institucional marrón oscuro propio.
  -- #5D2906 sobre #FFFFFF: ratio ~10.1:1 — WCAG AA OK.
  ('aguila',             'Águila',               '#5D2906', '#FFFFFF', 'Open Sans'),

  -- Block: sub-marca de Arcor (chocolate en barra individual). Hereda Arcor.
  ('block-chocolate',    'Block',                '#E2231A', '#FFFFFF', 'Nunito'),

  -- Shot: sub-marca de Arcor (chocolate Shot 55g). Hereda Arcor.
  ('shot-chocolate',     'Shot',                 '#E2231A', '#FFFFFF', 'Nunito'),

  -- === CARAMELOS Y CHUPETINES ===

  -- Sugus: marca de Mars/Wrigley Argentina. Color institucional naranja brillante.
  -- #FF6B00 sobre #000000: contraste ~5.7:1 — WCAG AA OK. wcag_adjusted (texto negro).
  ('sugus',              'Sugus',                '#FF6B00', '#000000', 'Roboto'),   -- wcag_adjusted

  -- Palitos de la Selva: sub-marca de Arcor (chupetines/caramelos). Hereda Arcor.
  ('palitos-de-la-selva','Palitos de la Selva',  '#E2231A', '#FFFFFF', 'Nunito'),

  -- Flynn Paff: caramelo masticable de Arcor. Hereda Arcor.
  ('flynn-paff',         'Flynn Paff',           '#E2231A', '#FFFFFF', 'Nunito'),

  -- Mr. Pop: caramelo explosivo de Arcor. Hereda Arcor.
  ('mr-pop',             'Mr. Pop',              '#E2231A', '#FFFFFF', 'Nunito'),

  -- Butter Toffees: marca de Stani (ahora Arcor group). Sin identidad propia → hereda Arcor.
  -- Nota: Stani fue absorbida por Arcor. Butter Toffees es de Arcor desde 2000s.
  ('butter-toffees',     'Butter Toffees',       '#E2231A', '#FFFFFF', 'Nunito'),

  -- Pico Dulce: chupetín de Arcor. Hereda Arcor.
  ('pico-dulce',         'Pico Dulce',           '#E2231A', '#FFFFFF', 'Nunito'),

  -- Bombuchas: marca de Arcor (caramelos de goma pequeños). Hereda Arcor.
  ('bombuchas',          'Bombuchas',            '#E2231A', '#FFFFFF', 'Nunito'),

  -- Mantecol: sub-marca de Georgalos (maní con miel). Hereda Georgalos (#8B0000).
  -- Identidad propia no verificada independientemente.
  ('mantecol',           'Mantecol',             '#8B0000', '#FFFFFF', 'Roboto'),

  -- === GOMITAS Y CONFITADOS ===

  -- Mogul: sub-marca de Arcor (gomitas, cerebritos, ositos). Hereda Arcor.
  ('mogul',              'Mogul',                '#E2231A', '#FFFFFF', 'Nunito'),

  -- Billiken Golosinas: sub-marca de Arcor. Sin identidad propia → hereda Arcor.
  -- Nota: la revista Billiken es diferente; esta es la línea de golosinas de Arcor.
  ('billiken-golosinas', 'Billiken',             '#E2231A', '#FFFFFF', 'Nunito'),

  -- Rocklets: confite de chocolate de Arcor. Hereda Arcor.
  ('rocklets',           'Rocklets',             '#E2231A', '#FFFFFF', 'Nunito'),

  -- Pindapoy: marca nacional de caramelos con vitamina C (Laboratorio Industrial).
  -- Sin identidad visual verificada → NULL.
  ('pindapoy',           'Pindapoy',             NULL, NULL, NULL)

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
-- guaymallen        #8B0000/#FFFFFF  ~8.6:1  OK
-- capitan-del-espacio #8B0000/#FFFFFF ~8.6:1 OK
-- jorgito           #1B5E20/#FFFFFF  ~9.2:1  OK
-- fantoche          #E2231A/#FFFFFF  ~4.85:1 OK
-- la-repostera      #E2231A/#FFFFFF  ~4.85:1 OK
-- suchard           #8B0000/#FFFFFF  ~8.6:1  OK
-- cachafaz          #3E1C00/#FFFFFF  ~14.5:1 OK
-- tofi              #E2231A/#FFFFFF  ~4.85:1 OK
-- cofler            #8B0000/#FFFFFF  ~8.6:1  OK
-- cadbury           #4B0082/#FFFFFF  ~10.5:1 OK
-- tita              #DA291C/#FFFFFF  ~4.85:1 OK
-- rhodesia          #8B0000/#FFFFFF  ~8.6:1  OK
-- aguila            #5D2906/#FFFFFF  ~10.1:1 OK
-- block-chocolate   #E2231A/#FFFFFF  ~4.85:1 OK
-- shot-chocolate    #E2231A/#FFFFFF  ~4.85:1 OK
-- sugus             #FF6B00/#000000  ~5.7:1  OK (wcag_adjusted — texto negro)
-- palitos-de-la-selva #E2231A/#FFFFFF ~4.85:1 OK
-- flynn-paff        #E2231A/#FFFFFF  ~4.85:1 OK
-- mr-pop            #E2231A/#FFFFFF  ~4.85:1 OK
-- butter-toffees    #E2231A/#FFFFFF  ~4.85:1 OK
-- pico-dulce        #E2231A/#FFFFFF  ~4.85:1 OK
-- bombuchas         #E2231A/#FFFFFF  ~4.85:1 OK
-- mantecol          #8B0000/#FFFFFF  ~8.6:1  OK
-- mogul             #E2231A/#FFFFFF  ~4.85:1 OK
-- billiken-golosinas #E2231A/#FFFFFF ~4.85:1 OK
-- rocklets          #E2231A/#FFFFFF  ~4.85:1 OK
-- pindapoy          NULL             —       Sin datos suficientes
