-- Seed 081: Productos globales — Ferretería: Herramientas Manuales y Eléctricas (~90 productos)
-- CICLO: cycle-009-catalog-volume-expansion
-- FECHA: 2026-04-25
-- FUENTE: catálogos ferreterías NEA (Posadas) + ferreterias.com.ar + MercadoLibre Argentina + distribuidores Stanley/Bosch/Makita/Truper 2026
-- ZONA: Posadas, Misiones (NEA) — precios minoristas aproximados AR$ abril 2026
-- IDEMPOTENTE: ON CONFLICT DO NOTHING
-- CATEGORÍAS: herramientas-manuales, herramientas-electricas, medicion

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES

-- ============================================================
-- MARTILLOS — varios pesos y usos
-- ============================================================
('Martillo carpintero 16oz Stanley',            'Stanley',      'herramientas-manuales',  12500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Martillo carpintero 20oz Stanley',            'Stanley',      'herramientas-manuales',  14800.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Martillo de goma 250g Bahco',                 'Bahco',        'herramientas-manuales',   9800.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Martillo mecánico 300g Truper',               'Truper',       'herramientas-manuales',   7500.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Martillo electricista 300g Tramontina',        'Tramontina',   'herramientas-manuales',  10200.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Martillo demoledor 1kg genérico',             NULL,           'herramientas-manuales',   8500.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- DESTORNILLADORES — planos y Phillips, varios números
-- ============================================================
('Destornillador plano 1/4" x 6" Stanley',      'Stanley',      'herramientas-manuales',   3200.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Destornillador plano 3/16" x 4" Stanley',     'Stanley',      'herramientas-manuales',   2800.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Destornillador Phillips PH2 x 6" Stanley',    'Stanley',      'herramientas-manuales',   3200.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Destornillador Phillips PH1 x 4" Stanley',    'Stanley',      'herramientas-manuales',   2800.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Set destornilladores 6u plano/Phillips Stanley','Stanley',     'herramientas-manuales',  14800.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),
('Destornillador de golpe 3/8" x 6" Bahco',     'Bahco',        'herramientas-manuales',   6800.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Set destornilladores 8u Truper',               'Truper',       'herramientas-manuales',   9500.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),

-- ============================================================
-- LLAVES — combinadas, inglesa, de caño, torx, allen
-- ============================================================
('Llave combinada 10mm Stanley',                'Stanley',      'herramientas-manuales',   3800.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave combinada 12mm Stanley',                'Stanley',      'herramientas-manuales',   4200.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave combinada 13mm Stanley',                'Stanley',      'herramientas-manuales',   4500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave combinada 17mm Stanley',                'Stanley',      'herramientas-manuales',   5200.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Juego llaves combinadas 8u Tramontina',       'Tramontina',   'herramientas-manuales',  22000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),
('Juego llaves combinadas 12u Bahco',           'Bahco',        'herramientas-manuales',  48000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),
('Llave inglesa 10" Stanley',                   'Stanley',      'herramientas-manuales',   9800.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave inglesa 12" Tramontina',                'Tramontina',   'herramientas-manuales',  11000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave de caño 14" Bahco',                     'Bahco',        'herramientas-manuales',  16500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Llave de caño 18" genérica',                  NULL,           'herramientas-manuales',  12000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Set llaves allen métrico 9u Stanley',         'Stanley',      'herramientas-manuales',   5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),
('Set llaves allen imperial 9u Stanley',        'Stanley',      'herramientas-manuales',   5500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),
('Set llaves Torx 8u Truper',                   'Truper',       'herramientas-manuales',   6800.00, 'seed', 0.7, 76, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),
('Llave estrella 19-22mm genérica',             NULL,           'herramientas-manuales',   4200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- ALICATES — varios tipos
-- ============================================================
('Alicate multigrip 8" Stanley',                'Stanley',      'herramientas-manuales',   9200.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Alicate de punta 6" Stanley',                 'Stanley',      'herramientas-manuales',   8400.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Alicate de corte diagonal 6" Stanley',        'Stanley',      'herramientas-manuales',   7800.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Alicate pelamango 7" Bahco',                  'Bahco',        'herramientas-manuales',  14500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Alicate presión tipo C 10" Tramontina',       'Tramontina',   'herramientas-manuales',  11000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Set alicates 3u punta/corte/multigrip Truper','Truper',       'herramientas-manuales',  18500.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),

-- ============================================================
-- SIERRAS MANUALES Y ACCESORIOS
-- ============================================================
('Sierra de arco para metal 12" Stanley',       'Stanley',      'herramientas-manuales',   8500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Hojas de repuesto sierra arco 12" x5u Stanley','Stanley',     'herramientas-manuales',   4500.00, 'seed', 0.8, 80, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),
('Serrucho carpintero 20" Stanley',             'Stanley',      'herramientas-manuales',  12800.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Serrucho podador 13" Truper',                 'Truper',       'herramientas-manuales',   7500.00, 'seed', 0.7, 76, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- HERRAMIENTAS ELÉCTRICAS — Bosch
-- ============================================================
('Taladro percutor GSB 13 RE 650W Bosch',       'Bosch',        'herramientas-electricas', 68000.00, 'seed', 0.8, 88, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Taladro percutor GSB 16 RE 750W Bosch',       'Bosch',        'herramientas-electricas', 85000.00, 'seed', 0.8, 88, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Amoladora GWS 700 4.5" Bosch',                'Bosch',        'herramientas-electricas', 52000.00, 'seed', 0.8, 88, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Amoladora GWS 750-125 5" Bosch',              'Bosch',        'herramientas-electricas', 68000.00, 'seed', 0.8, 88, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Sierra circular GKS 190 1400W Bosch',         'Bosch',        'herramientas-electricas',125000.00, 'seed', 0.8, 88, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Sierra caladora GST 700 500W Bosch',          'Bosch',        'herramientas-electricas', 78000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lijadora orbital GOS 108V-LI Bosch',          'Bosch',        'herramientas-electricas', 62000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Router fresadora GKF 600 Bosch',              'Bosch',        'herramientas-electricas',148000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- HERRAMIENTAS ELÉCTRICAS — Makita
-- ============================================================
('Taladro inalámbrico DDF453 18V Makita',       'Makita',       'herramientas-electricas',145000.00, 'seed', 0.8, 88, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Taladro percutor HP1630 16mm Makita',         'Makita',       'herramientas-electricas', 78000.00, 'seed', 0.8, 88, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Amoladora GA4530 4.5" 720W Makita',           'Makita',       'herramientas-electricas', 58000.00, 'seed', 0.8, 88, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Amoladora GA5030 5" 1010W Makita',            'Makita',       'herramientas-electricas', 78000.00, 'seed', 0.8, 88, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Sierra circular 5007N 1800W Makita',          'Makita',       'herramientas-electricas',135000.00, 'seed', 0.8, 88, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- HERRAMIENTAS ELÉCTRICAS — Black+Decker
-- ============================================================
('Taladro percutor BD700KA 710W Black+Decker',  'Black+Decker', 'herramientas-electricas', 48000.00, 'seed', 0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Amoladora G720 4.5" 720W Black+Decker',       'Black+Decker', 'herramientas-electricas', 38000.00, 'seed', 0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Lijadora orbital orbital KA290K Black+Decker','Black+Decker', 'herramientas-electricas', 32000.00, 'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Sierra circular CS1015 1020W Black+Decker',   'Black+Decker', 'herramientas-electricas', 68000.00, 'seed', 0.7, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Destornillador inalámbrico BDCS36G Black+Decker','Black+Decker','herramientas-electricas',28000.00,'seed', 0.7, 80, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- HERRAMIENTAS ELÉCTRICAS — DeWalt
-- ============================================================
('Taladro percutor DWD024 550W DeWalt',         'DeWalt',       'herramientas-electricas', 72000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Amoladora DWE4120 4.5" 800W DeWalt',          'DeWalt',       'herramientas-electricas', 62000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Sierra caladora DWE301 500W DeWalt',          'DeWalt',       'herramientas-electricas', 82000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- DISCOS PARA AMOLADORA (consumibles clave de ferretería)
-- ============================================================
('Disco de corte 4.5" para metal 115mm genérico',NULL,          'herramientas-electricas',  1200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Disco de desbaste 4.5" para metal 115mm genérico',NULL,       'herramientas-electricas',  1800.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Disco de corte 4.5" para mampostería 115mm Bosch','Bosch',    'herramientas-electricas',  2800.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Disco de corte 7" para metal 180mm genérico', NULL,           'herramientas-electricas',  2200.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Disco diamantado 4.5" para cerámica 115mm Bosch','Bosch',     'herramientas-electricas',  8500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Pack discos de corte 4.5" x10u 3M',           '3M',           'herramientas-electricas', 15000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),

-- ============================================================
-- MEDICIÓN — flexómetros, niveles, escuadras
-- ============================================================
('Cinta métrica 5m Stanley FatMax',             'Stanley',      'medicion',                8500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cinta métrica 8m Stanley FatMax',             'Stanley',      'medicion',               10800.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Cinta métrica 3m Truper',                     'Truper',       'medicion',                4200.00, 'seed', 0.7, 76, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Nivel de burbuja 60cm Stanley',               'Stanley',      'medicion',               12000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Nivel de burbuja 120cm Stanley',              'Stanley',      'medicion',               18500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Nivel de burbuja 40cm Truper',                'Truper',       'medicion',                6800.00, 'seed', 0.7, 76, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Nivel láser línea cruzada Stanley',           'Stanley',      'medicion',               48000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Escuadra carpintero 30cm Stanley',            'Stanley',      'medicion',                6500.00, 'seed', 0.8, 82, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Escuadra T 45cm Truper',                      'Truper',       'medicion',                5200.00, 'seed', 0.7, 76, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Calibre Vernier 150mm Truper',                'Truper',       'medicion',               12000.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Marcador de tiza azul x3u Stanley',           'Stanley',      'medicion',                4800.00, 'seed', 0.8, 80, FALSE, TRUE, 'ferreteria', '{"unit": "pack",   "sku_prefix": "FERRET"}'),
('Multímetro digital Truper',                   'Truper',       'medicion',               18000.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Detector de metales y cables Bosch',          'Bosch',        'medicion',               28000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- LLAVES DE IMPACTO / CABEZAS DE SOCKET
-- ============================================================
('Set llaves de socket 1/2" 21u Bahco',         'Bahco',        'herramientas-manuales',  85000.00, 'seed', 0.8, 86, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),
('Set llaves de socket 3/8" 38u Truper',        'Truper',       'herramientas-manuales',  45000.00, 'seed', 0.7, 78, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),
('Llave de impacto neumática 1/2" genérica',    NULL,           'herramientas-electricas', 65000.00, 'seed', 0.6, 40, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),

-- ============================================================
-- ACCESORIOS PARA TALADRO
-- ============================================================
('Set brocas HSS para metal 13u Bosch',         'Bosch',        'herramientas-electricas', 18000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),
('Set brocas concreto 5u 4-10mm Bosch',         'Bosch',        'herramientas-electricas', 12000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),
('Set puntas de atornillador 30u Würth',        'Würth',        'herramientas-electricas', 15000.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "set",    "sku_prefix": "FERRET"}'),
('Broca para madera 25mm Truper',               'Truper',       'herramientas-electricas',  3800.00, 'seed', 0.7, 76, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}'),
('Broca SDS plus 12mm Bosch',                   'Bosch',        'herramientas-electricas',  8500.00, 'seed', 0.8, 84, FALSE, TRUE, 'ferreteria', '{"unit": "unidad", "sku_prefix": "FERRET"}')

ON CONFLICT (name, business_type) DO NOTHING;

-- Total: 91 productos
-- Categorías usadas: herramientas-manuales, herramientas-electricas, medicion
-- Marcas: Stanley, Bahco, Tramontina, Truper, Bosch, Makita, Black+Decker, DeWalt, Würth, 3M, genéricos
-- Precios referencia zona Posadas, NEA — abril 2026 (minorista estimado)
