-- Seed 029: Enriquecimiento v4 — lote 2 de templates con categorías jerárquicas
-- PROPÓSITO: Actualizar templates de 8 rubros adicionales con árboles de categorías (parent_slug + level),
--            marcas enriquecidas del mercado argentino. Productos vacíos (se cargan aparte).
-- IDEMPOTENTE: Solo UPDATE con WHERE; no inserta ni elimina registros.
-- REQUIERE: 001 (business_types), 004 (business_type_quickstart_templates)
-- VERSION: 4.0.0-enriched | generated_by: manual-curation-v4
-- RUBROS: perfumeria, muebleria, celulares, computacion, repuestos, lubricentro, jugueteria, veterinaria

-- =====================================================
-- 1. PERFUMERÍA — 6 padres + 17 hijas = 23 categorías, 10 marcas
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "perfumes", "name": "Perfumes", "level": 0},
    {"slug": "perfumes-mujer", "name": "Perfumes Mujer", "parent_slug": "perfumes", "level": 1},
    {"slug": "perfumes-hombre", "name": "Perfumes Hombre", "parent_slug": "perfumes", "level": 1},
    {"slug": "perfumes-unisex", "name": "Perfumes Unisex", "parent_slug": "perfumes", "level": 1},

    {"slug": "maquillaje", "name": "Maquillaje", "level": 0},
    {"slug": "base-correctores", "name": "Base y Correctores", "parent_slug": "maquillaje", "level": 1},
    {"slug": "labiales", "name": "Labiales", "parent_slug": "maquillaje", "level": 1},
    {"slug": "ojos", "name": "Ojos", "parent_slug": "maquillaje", "level": 1},

    {"slug": "cuidado-capilar", "name": "Cuidado Capilar", "level": 0},
    {"slug": "shampoo", "name": "Shampoo", "parent_slug": "cuidado-capilar", "level": 1},
    {"slug": "acondicionador", "name": "Acondicionador", "parent_slug": "cuidado-capilar", "level": 1},
    {"slug": "tratamiento-capilar", "name": "Tratamiento Capilar", "parent_slug": "cuidado-capilar", "level": 1},

    {"slug": "cuidado-corporal", "name": "Cuidado Corporal", "level": 0},
    {"slug": "cremas-corporales", "name": "Cremas Corporales", "parent_slug": "cuidado-corporal", "level": 1},
    {"slug": "desodorantes", "name": "Desodorantes", "parent_slug": "cuidado-corporal", "level": 1},
    {"slug": "jabones", "name": "Jabones", "parent_slug": "cuidado-corporal", "level": 1},

    {"slug": "cuidado-facial", "name": "Cuidado Facial", "level": 0},
    {"slug": "limpieza-facial", "name": "Limpieza Facial", "parent_slug": "cuidado-facial", "level": 1},
    {"slug": "hidratantes-faciales", "name": "Hidratantes Faciales", "parent_slug": "cuidado-facial", "level": 1},

    {"slug": "accesorios-belleza", "name": "Accesorios de Belleza", "level": 0},
    {"slug": "brochas-pinceles", "name": "Brochas y Pinceles", "parent_slug": "accesorios-belleza", "level": 1},
    {"slug": "espejos-organizadores", "name": "Espejos y Organizadores", "parent_slug": "accesorios-belleza", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Natura", "suggested_for_categories": ["perfumes-mujer", "cremas-corporales"]},
    {"name": "Avon", "suggested_for_categories": ["perfumes-mujer", "labiales"]},
    {"name": "L''Oréal", "suggested_for_categories": ["base-correctores", "tratamiento-capilar"]},
    {"name": "Nivea", "suggested_for_categories": ["cremas-corporales", "hidratantes-faciales"]},
    {"name": "Dove", "suggested_for_categories": ["jabones", "desodorantes"]},
    {"name": "TRESemmé", "suggested_for_categories": ["shampoo", "acondicionador"]},
    {"name": "Revlon", "suggested_for_categories": ["labiales", "ojos"]},
    {"name": "Maybelline", "suggested_for_categories": ["base-correctores", "ojos"]},
    {"name": "Sedal", "suggested_for_categories": ["shampoo", "acondicionador"]},
    {"name": "Pantene", "suggested_for_categories": ["shampoo", "tratamiento-capilar"]}
  ]'::jsonb,
  products = '[]'::jsonb,
  version = '4.0.0-enriched',
  generated_by = 'manual-curation-v4',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'perfumeria') AND is_default = true;

-- =====================================================
-- 2. MUEBLERÍA — 7 padres + 17 hijas = 24 categorías, 7 marcas
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "living", "name": "Living", "level": 0},
    {"slug": "sillones", "name": "Sillones", "parent_slug": "living", "level": 1},
    {"slug": "mesas-ratonas", "name": "Mesas Ratonas", "parent_slug": "living", "level": 1},
    {"slug": "racks-tv", "name": "Racks y Muebles de TV", "parent_slug": "living", "level": 1},

    {"slug": "dormitorio", "name": "Dormitorio", "level": 0},
    {"slug": "camas", "name": "Camas", "parent_slug": "dormitorio", "level": 1},
    {"slug": "colchones", "name": "Colchones", "parent_slug": "dormitorio", "level": 1},
    {"slug": "placares-roperos", "name": "Placares y Roperos", "parent_slug": "dormitorio", "level": 1},

    {"slug": "cocina-comedor", "name": "Cocina y Comedor", "level": 0},
    {"slug": "mesas-comedor", "name": "Mesas", "parent_slug": "cocina-comedor", "level": 1},
    {"slug": "sillas-comedor", "name": "Sillas", "parent_slug": "cocina-comedor", "level": 1},

    {"slug": "oficina", "name": "Oficina", "level": 0},
    {"slug": "escritorios", "name": "Escritorios", "parent_slug": "oficina", "level": 1},
    {"slug": "sillas-oficina", "name": "Sillas de Oficina", "parent_slug": "oficina", "level": 1},
    {"slug": "bibliotecas-estantes", "name": "Bibliotecas y Estantes", "parent_slug": "oficina", "level": 1},

    {"slug": "exterior", "name": "Exterior", "level": 0},
    {"slug": "juegos-jardin", "name": "Juegos de Jardín", "parent_slug": "exterior", "level": 1},
    {"slug": "reposeras-sillones-ext", "name": "Reposeras y Sillones", "parent_slug": "exterior", "level": 1},

    {"slug": "bano", "name": "Baño", "level": 0},
    {"slug": "vanitorios", "name": "Vanitorios", "parent_slug": "bano", "level": 1},
    {"slug": "botiquines", "name": "Botiquines", "parent_slug": "bano", "level": 1},

    {"slug": "infantil", "name": "Infantil", "level": 0},
    {"slug": "cunas-camas-infantiles", "name": "Cunas y Camas Infantiles", "parent_slug": "infantil", "level": 1},
    {"slug": "escritorios-infantiles", "name": "Escritorios Infantiles", "parent_slug": "infantil", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Inval", "suggested_for_categories": ["escritorios", "bibliotecas-estantes"]},
    {"name": "Platinum", "suggested_for_categories": ["colchones", "camas"]},
    {"name": "Fiplasto", "suggested_for_categories": ["placares-roperos", "vanitorios"]},
    {"name": "Mosconi", "suggested_for_categories": ["sillones", "sillas-comedor"]},
    {"name": "BRN", "suggested_for_categories": ["mesas-comedor", "mesas-ratonas"]},
    {"name": "Tables", "suggested_for_categories": ["mesas-comedor", "escritorios"]},
    {"name": "Forbidan", "suggested_for_categories": ["racks-tv", "bibliotecas-estantes"]}
  ]'::jsonb,
  products = '[]'::jsonb,
  version = '4.0.0-enriched',
  generated_by = 'manual-curation-v4',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'muebleria') AND is_default = true;

-- =====================================================
-- 3. CASA DE CELULARES — 6 padres + 15 hijas = 21 categorías, 7 marcas
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "smartphones", "name": "Smartphones", "level": 0},
    {"slug": "android", "name": "Android", "parent_slug": "smartphones", "level": 1},
    {"slug": "iphone", "name": "iPhone", "parent_slug": "smartphones", "level": 1},
    {"slug": "celulares-basicos", "name": "Celulares Básicos", "parent_slug": "smartphones", "level": 1},

    {"slug": "accesorios-celulares", "name": "Accesorios", "level": 0},
    {"slug": "fundas-carcasas", "name": "Fundas y Carcasas", "parent_slug": "accesorios-celulares", "level": 1},
    {"slug": "vidrios-templados", "name": "Vidrios Templados", "parent_slug": "accesorios-celulares", "level": 1},
    {"slug": "soportes-celular", "name": "Soportes", "parent_slug": "accesorios-celulares", "level": 1},

    {"slug": "audio-celulares", "name": "Audio", "level": 0},
    {"slug": "auriculares-bluetooth", "name": "Auriculares Bluetooth", "parent_slug": "audio-celulares", "level": 1},
    {"slug": "auriculares-cable", "name": "Auriculares con Cable", "parent_slug": "audio-celulares", "level": 1},
    {"slug": "parlantes-portatiles", "name": "Parlantes Portátiles", "parent_slug": "audio-celulares", "level": 1},

    {"slug": "cargadores-cables", "name": "Cargadores y Cables", "level": 0},
    {"slug": "cargadores", "name": "Cargadores", "parent_slug": "cargadores-cables", "level": 1},
    {"slug": "cables-datos", "name": "Cables de Datos", "parent_slug": "cargadores-cables", "level": 1},

    {"slug": "smartwatch-wearables", "name": "Smartwatch y Wearables", "level": 0},
    {"slug": "smartwatch", "name": "Smartwatch", "parent_slug": "smartwatch-wearables", "level": 1},
    {"slug": "smartbands", "name": "Smartbands", "parent_slug": "smartwatch-wearables", "level": 1},

    {"slug": "reparacion-celulares", "name": "Reparación", "level": 0},
    {"slug": "pantallas-repuesto", "name": "Pantallas de Repuesto", "parent_slug": "reparacion-celulares", "level": 1},
    {"slug": "baterias-repuesto", "name": "Baterías de Repuesto", "parent_slug": "reparacion-celulares", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Samsung", "suggested_for_categories": ["android", "auriculares-bluetooth"]},
    {"name": "Apple", "suggested_for_categories": ["iphone", "smartwatch"]},
    {"name": "Motorola", "suggested_for_categories": ["android", "celulares-basicos"]},
    {"name": "Xiaomi", "suggested_for_categories": ["android", "smartbands"]},
    {"name": "JBL", "suggested_for_categories": ["auriculares-bluetooth", "parlantes-portatiles"]},
    {"name": "Anker", "suggested_for_categories": ["cargadores", "cables-datos"]},
    {"name": "Spigen", "suggested_for_categories": ["fundas-carcasas", "vidrios-templados"]}
  ]'::jsonb,
  products = '[]'::jsonb,
  version = '4.0.0-enriched',
  generated_by = 'manual-curation-v4',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'celulares') AND is_default = true;

-- =====================================================
-- 4. COMPUTACIÓN — 7 padres + 18 hijas = 25 categorías, 9 marcas
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "notebooks", "name": "Notebooks", "level": 0},
    {"slug": "notebooks-gamer", "name": "Notebooks Gamer", "parent_slug": "notebooks", "level": 1},
    {"slug": "notebooks-oficina", "name": "Notebooks Oficina", "parent_slug": "notebooks", "level": 1},
    {"slug": "notebooks-ultralivianas", "name": "Notebooks Ultralivianas", "parent_slug": "notebooks", "level": 1},

    {"slug": "pcs-escritorio", "name": "PCs de Escritorio", "level": 0},
    {"slug": "pc-armada", "name": "PC Armada", "parent_slug": "pcs-escritorio", "level": 1},
    {"slug": "all-in-one", "name": "All in One", "parent_slug": "pcs-escritorio", "level": 1},

    {"slug": "perifericos", "name": "Periféricos", "level": 0},
    {"slug": "teclados", "name": "Teclados", "parent_slug": "perifericos", "level": 1},
    {"slug": "mouse", "name": "Mouse", "parent_slug": "perifericos", "level": 1},
    {"slug": "monitores", "name": "Monitores", "parent_slug": "perifericos", "level": 1},

    {"slug": "componentes", "name": "Componentes", "level": 0},
    {"slug": "placas-video", "name": "Placas de Video", "parent_slug": "componentes", "level": 1},
    {"slug": "memorias-ram", "name": "Memorias RAM", "parent_slug": "componentes", "level": 1},
    {"slug": "procesadores", "name": "Procesadores", "parent_slug": "componentes", "level": 1},

    {"slug": "redes", "name": "Redes", "level": 0},
    {"slug": "routers", "name": "Routers", "parent_slug": "redes", "level": 1},
    {"slug": "switches-access-points", "name": "Switches y Access Points", "parent_slug": "redes", "level": 1},

    {"slug": "almacenamiento", "name": "Almacenamiento", "level": 0},
    {"slug": "discos-ssd", "name": "Discos SSD", "parent_slug": "almacenamiento", "level": 1},
    {"slug": "discos-hdd", "name": "Discos HDD", "parent_slug": "almacenamiento", "level": 1},
    {"slug": "pendrives", "name": "Pendrives", "parent_slug": "almacenamiento", "level": 1},

    {"slug": "impresoras", "name": "Impresoras", "level": 0},
    {"slug": "impresoras-tinta", "name": "Impresoras de Tinta", "parent_slug": "impresoras", "level": 1},
    {"slug": "impresoras-laser", "name": "Impresoras Láser", "parent_slug": "impresoras", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Lenovo", "suggested_for_categories": ["notebooks-oficina", "all-in-one"]},
    {"name": "HP", "suggested_for_categories": ["notebooks-oficina", "impresoras-tinta"]},
    {"name": "Dell", "suggested_for_categories": ["notebooks-ultralivianas", "monitores"]},
    {"name": "Logitech", "suggested_for_categories": ["teclados", "mouse"]},
    {"name": "Kingston", "suggested_for_categories": ["memorias-ram", "pendrives"]},
    {"name": "Samsung", "suggested_for_categories": ["discos-ssd", "monitores"]},
    {"name": "WD", "suggested_for_categories": ["discos-hdd", "discos-ssd"]},
    {"name": "Epson", "suggested_for_categories": ["impresoras-tinta", "impresoras-laser"]},
    {"name": "TP-Link", "suggested_for_categories": ["routers", "switches-access-points"]}
  ]'::jsonb,
  products = '[]'::jsonb,
  version = '4.0.0-enriched',
  generated_by = 'manual-curation-v4',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'computacion') AND is_default = true;

-- =====================================================
-- 5. CASA DE REPUESTOS — 7 padres + 18 hijas = 25 categorías, 8 marcas
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "motor", "name": "Motor", "level": 0},
    {"slug": "bujias", "name": "Bujías", "parent_slug": "motor", "level": 1},
    {"slug": "bomba-agua", "name": "Bomba de Agua", "parent_slug": "motor", "level": 1},
    {"slug": "termostato", "name": "Termostato", "parent_slug": "motor", "level": 1},

    {"slug": "suspension", "name": "Suspensión", "level": 0},
    {"slug": "amortiguadores", "name": "Amortiguadores", "parent_slug": "suspension", "level": 1},
    {"slug": "rotulas-extremos", "name": "Rótulas y Extremos", "parent_slug": "suspension", "level": 1},
    {"slug": "bujes-silent-blocks", "name": "Bujes y Silent Blocks", "parent_slug": "suspension", "level": 1},

    {"slug": "frenos", "name": "Frenos", "level": 0},
    {"slug": "pastillas-freno", "name": "Pastillas de Freno", "parent_slug": "frenos", "level": 1},
    {"slug": "discos-freno", "name": "Discos de Freno", "parent_slug": "frenos", "level": 1},
    {"slug": "cilindros-freno", "name": "Cilindros de Freno", "parent_slug": "frenos", "level": 1},

    {"slug": "electricidad-automotor", "name": "Electricidad Automotor", "level": 0},
    {"slug": "alternador", "name": "Alternador", "parent_slug": "electricidad-automotor", "level": 1},
    {"slug": "burro-arranque", "name": "Burro de Arranque", "parent_slug": "electricidad-automotor", "level": 1},
    {"slug": "baterias-auto", "name": "Baterías", "parent_slug": "electricidad-automotor", "level": 1},

    {"slug": "filtros-repuestos", "name": "Filtros", "level": 0},
    {"slug": "filtro-aceite", "name": "Filtro de Aceite", "parent_slug": "filtros-repuestos", "level": 1},
    {"slug": "filtro-aire", "name": "Filtro de Aire", "parent_slug": "filtros-repuestos", "level": 1},
    {"slug": "filtro-combustible", "name": "Filtro de Combustible", "parent_slug": "filtros-repuestos", "level": 1},

    {"slug": "correas-distribucion", "name": "Correas y Distribución", "level": 0},
    {"slug": "kit-distribucion", "name": "Kit de Distribución", "parent_slug": "correas-distribucion", "level": 1},
    {"slug": "correas-accesorios", "name": "Correas de Accesorios", "parent_slug": "correas-distribucion", "level": 1},

    {"slug": "carroceria", "name": "Carrocería", "level": 0},
    {"slug": "opticas-faros", "name": "Ópticas y Faros", "parent_slug": "carroceria", "level": 1},
    {"slug": "espejos-retrovisores", "name": "Espejos Retrovisores", "parent_slug": "carroceria", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Bosch", "suggested_for_categories": ["bujias", "alternador"]},
    {"name": "NGK", "suggested_for_categories": ["bujias", "cables-bujia"]},
    {"name": "Monroe", "suggested_for_categories": ["amortiguadores", "rotulas-extremos"]},
    {"name": "Ferodo", "suggested_for_categories": ["pastillas-freno", "discos-freno"]},
    {"name": "Mann Filter", "suggested_for_categories": ["filtro-aceite", "filtro-aire"]},
    {"name": "Gates", "suggested_for_categories": ["kit-distribucion", "correas-accesorios"]},
    {"name": "Sachs", "suggested_for_categories": ["amortiguadores", "bujes-silent-blocks"]},
    {"name": "Mahle", "suggested_for_categories": ["filtro-combustible", "termostato"]}
  ]'::jsonb,
  products = '[]'::jsonb,
  version = '4.0.0-enriched',
  generated_by = 'manual-curation-v4',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'repuestos') AND is_default = true;

-- =====================================================
-- 6. LUBRICENTRO — 6 padres + 16 hijas = 22 categorías, 7 marcas
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "aceites-motor", "name": "Aceites Motor", "level": 0},
    {"slug": "aceite-sintetico", "name": "Aceite Sintético", "parent_slug": "aceites-motor", "level": 1},
    {"slug": "aceite-semi-sintetico", "name": "Aceite Semi-Sintético", "parent_slug": "aceites-motor", "level": 1},
    {"slug": "aceite-mineral", "name": "Aceite Mineral", "parent_slug": "aceites-motor", "level": 1},

    {"slug": "aceites-transmision", "name": "Aceites Transmisión", "level": 0},
    {"slug": "aceite-caja-manual", "name": "Aceite Caja Manual", "parent_slug": "aceites-transmision", "level": 1},
    {"slug": "aceite-caja-automatica", "name": "Aceite Caja Automática", "parent_slug": "aceites-transmision", "level": 1},
    {"slug": "aceite-diferencial", "name": "Aceite Diferencial", "parent_slug": "aceites-transmision", "level": 1},

    {"slug": "filtros-lubricentro", "name": "Filtros", "level": 0},
    {"slug": "filtro-aceite-lub", "name": "Filtro de Aceite", "parent_slug": "filtros-lubricentro", "level": 1},
    {"slug": "filtro-aire-lub", "name": "Filtro de Aire", "parent_slug": "filtros-lubricentro", "level": 1},
    {"slug": "filtro-habitaculo", "name": "Filtro de Habitáculo", "parent_slug": "filtros-lubricentro", "level": 1},

    {"slug": "refrigerantes", "name": "Refrigerantes", "level": 0},
    {"slug": "anticongelante", "name": "Anticongelante", "parent_slug": "refrigerantes", "level": 1},
    {"slug": "agua-destilada", "name": "Agua Destilada", "parent_slug": "refrigerantes", "level": 1},

    {"slug": "limpieza-automotor", "name": "Limpieza Automotor", "level": 0},
    {"slug": "lavado-motor", "name": "Lavado de Motor", "parent_slug": "limpieza-automotor", "level": 1},
    {"slug": "liquido-limpiaparabrisas", "name": "Líquido Limpiaparabrisas", "parent_slug": "limpieza-automotor", "level": 1},
    {"slug": "aditivos", "name": "Aditivos", "parent_slug": "limpieza-automotor", "level": 1},

    {"slug": "servicios-lubricentro", "name": "Servicios", "level": 0},
    {"slug": "cambio-aceite", "name": "Cambio de Aceite", "parent_slug": "servicios-lubricentro", "level": 1},
    {"slug": "cambio-filtros", "name": "Cambio de Filtros", "parent_slug": "servicios-lubricentro", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Shell", "suggested_for_categories": ["aceite-sintetico", "aceite-semi-sintetico"]},
    {"name": "YPF", "suggested_for_categories": ["aceite-sintetico", "aceite-mineral"]},
    {"name": "Mobil", "suggested_for_categories": ["aceite-sintetico", "aceite-caja-automatica"]},
    {"name": "Castrol", "suggested_for_categories": ["aceite-sintetico", "aceite-semi-sintetico"]},
    {"name": "Total", "suggested_for_categories": ["aceite-mineral", "aceite-caja-manual"]},
    {"name": "Liqui Moly", "suggested_for_categories": ["aditivos", "aceite-sintetico"]},
    {"name": "Motul", "suggested_for_categories": ["aceite-sintetico", "aceite-diferencial"]}
  ]'::jsonb,
  products = '[]'::jsonb,
  version = '4.0.0-enriched',
  generated_by = 'manual-curation-v4',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'lubricentro') AND is_default = true;

-- =====================================================
-- 7. JUGUETERÍA — 7 padres + 17 hijas = 24 categorías, 7 marcas
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "bebes-primera-infancia", "name": "Bebés y Primera Infancia", "level": 0},
    {"slug": "sonajeros-mordillos", "name": "Sonajeros y Mordillos", "parent_slug": "bebes-primera-infancia", "level": 1},
    {"slug": "andadores-caminadores", "name": "Andadores y Caminadores", "parent_slug": "bebes-primera-infancia", "level": 1},

    {"slug": "juegos-mesa", "name": "Juegos de Mesa", "level": 0},
    {"slug": "rompecabezas", "name": "Rompecabezas", "parent_slug": "juegos-mesa", "level": 1},
    {"slug": "juegos-cartas", "name": "Juegos de Cartas", "parent_slug": "juegos-mesa", "level": 1},
    {"slug": "juegos-tablero", "name": "Juegos de Tablero", "parent_slug": "juegos-mesa", "level": 1},

    {"slug": "aire-libre", "name": "Aire Libre", "level": 0},
    {"slug": "pelotas", "name": "Pelotas", "parent_slug": "aire-libre", "level": 1},
    {"slug": "bicicletas-triciclos", "name": "Bicicletas y Triciclos", "parent_slug": "aire-libre", "level": 1},
    {"slug": "pistolas-agua", "name": "Pistolas de Agua", "parent_slug": "aire-libre", "level": 1},

    {"slug": "electronicos-juguetes", "name": "Electrónicos", "level": 0},
    {"slug": "consolas-videojuegos", "name": "Consolas y Videojuegos", "parent_slug": "electronicos-juguetes", "level": 1},
    {"slug": "drones-rc", "name": "Drones y Radio Control", "parent_slug": "electronicos-juguetes", "level": 1},

    {"slug": "didacticos", "name": "Didácticos", "level": 0},
    {"slug": "bloques-encastre", "name": "Bloques y Encastre", "parent_slug": "didacticos", "level": 1},
    {"slug": "arte-manualidades", "name": "Arte y Manualidades", "parent_slug": "didacticos", "level": 1},
    {"slug": "ciencia-experimentos", "name": "Ciencia y Experimentos", "parent_slug": "didacticos", "level": 1},

    {"slug": "munecos-figuras", "name": "Muñecos y Figuras", "level": 0},
    {"slug": "munecas", "name": "Muñecas", "parent_slug": "munecos-figuras", "level": 1},
    {"slug": "figuras-accion", "name": "Figuras de Acción", "parent_slug": "munecos-figuras", "level": 1},
    {"slug": "peluches", "name": "Peluches", "parent_slug": "munecos-figuras", "level": 1},

    {"slug": "vehiculos-pistas", "name": "Vehículos y Pistas", "level": 0},
    {"slug": "autos-camiones", "name": "Autos y Camiones", "parent_slug": "vehiculos-pistas", "level": 1},
    {"slug": "pistas-circuitos", "name": "Pistas y Circuitos", "parent_slug": "vehiculos-pistas", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Mattel", "suggested_for_categories": ["munecas", "autos-camiones"]},
    {"name": "Hasbro", "suggested_for_categories": ["figuras-accion", "juegos-tablero"]},
    {"name": "Fisher-Price", "suggested_for_categories": ["sonajeros-mordillos", "andadores-caminadores"]},
    {"name": "LEGO", "suggested_for_categories": ["bloques-encastre", "figuras-accion"]},
    {"name": "Playmobil", "suggested_for_categories": ["figuras-accion", "vehiculos-pistas"]},
    {"name": "Rasti", "suggested_for_categories": ["bloques-encastre", "didacticos"]},
    {"name": "Duravit", "suggested_for_categories": ["pelotas", "bicicletas-triciclos"]}
  ]'::jsonb,
  products = '[]'::jsonb,
  version = '4.0.0-enriched',
  generated_by = 'manual-curation-v4',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'jugueteria') AND is_default = true;

-- =====================================================
-- 8. VETERINARIA — 5 padres + 15 hijas = 20 categorías, 7 marcas
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "alimentos-mascotas", "name": "Alimentos", "level": 0},
    {"slug": "alimento-perro", "name": "Alimento Perro", "parent_slug": "alimentos-mascotas", "level": 1},
    {"slug": "alimento-gato", "name": "Alimento Gato", "parent_slug": "alimentos-mascotas", "level": 1},
    {"slug": "snacks-premios", "name": "Snacks y Premios", "parent_slug": "alimentos-mascotas", "level": 1},

    {"slug": "accesorios-mascotas", "name": "Accesorios", "level": 0},
    {"slug": "correas-collares", "name": "Correas y Collares", "parent_slug": "accesorios-mascotas", "level": 1},
    {"slug": "cuchas-camas", "name": "Cuchas y Camas", "parent_slug": "accesorios-mascotas", "level": 1},
    {"slug": "comederos-bebederos", "name": "Comederos y Bebederos", "parent_slug": "accesorios-mascotas", "level": 1},
    {"slug": "juguetes-mascotas", "name": "Juguetes", "parent_slug": "accesorios-mascotas", "level": 1},

    {"slug": "higiene-cuidado-mascotas", "name": "Higiene y Cuidado", "level": 0},
    {"slug": "shampoo-mascotas", "name": "Shampoo", "parent_slug": "higiene-cuidado-mascotas", "level": 1},
    {"slug": "antipulgas-garrapatas", "name": "Antipulgas y Garrapatas", "parent_slug": "higiene-cuidado-mascotas", "level": 1},
    {"slug": "piedras-sanitarias", "name": "Piedras Sanitarias", "parent_slug": "higiene-cuidado-mascotas", "level": 1},

    {"slug": "farmacia-veterinaria", "name": "Farmacia Veterinaria", "level": 0},
    {"slug": "antiparasitarios", "name": "Antiparasitarios", "parent_slug": "farmacia-veterinaria", "level": 1},
    {"slug": "vitaminas-suplementos", "name": "Vitaminas y Suplementos", "parent_slug": "farmacia-veterinaria", "level": 1},

    {"slug": "servicios-veterinaria", "name": "Servicios", "level": 0},
    {"slug": "consulta-veterinaria", "name": "Consulta Veterinaria", "parent_slug": "servicios-veterinaria", "level": 1},
    {"slug": "vacunacion", "name": "Vacunación", "parent_slug": "servicios-veterinaria", "level": 1},
    {"slug": "peluqueria-canina", "name": "Peluquería Canina", "parent_slug": "servicios-veterinaria", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Royal Canin", "suggested_for_categories": ["alimento-perro", "alimento-gato"]},
    {"name": "Purina", "suggested_for_categories": ["alimento-perro", "alimento-gato"]},
    {"name": "Eukanuba", "suggested_for_categories": ["alimento-perro", "snacks-premios"]},
    {"name": "Pedigree", "suggested_for_categories": ["alimento-perro", "snacks-premios"]},
    {"name": "Whiskas", "suggested_for_categories": ["alimento-gato", "snacks-premios"]},
    {"name": "Old Prince", "suggested_for_categories": ["alimento-perro", "alimento-gato"]},
    {"name": "Bayer Vet", "suggested_for_categories": ["antipulgas-garrapatas", "antiparasitarios"]}
  ]'::jsonb,
  products = '[]'::jsonb,
  version = '4.0.0-enriched',
  generated_by = 'manual-curation-v4',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'veterinaria') AND is_default = true;
