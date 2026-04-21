-- Seed 028: Enriquecimiento v4 de TODOS los templates con categorías jerárquicas
-- PROPÓSITO: Actualizar templates existentes con árboles de categorías (parent_slug + level),
--            marcas enriquecidas del mercado argentino, y productos representativos por subcategoría.
-- IDEMPOTENTE: UPDATE con WHERE; INSERT para nuevos business_types y templates.
-- REQUIERE: 001 (business_types), 004 (business_type_quickstart_templates)
-- VERSION: 4.0.0-enriched | generated_by: manual-curation + web-research
-- NOTA: Las categorías ahora incluyen "parent_slug" y "level" para soporte jerárquico.
--       level=0 son raíz, level=1 son subcategorías. El backend resuelve parent_id al aplicar.

-- =====================================================
-- PASO 0: Asegurar business_types nuevos
-- =====================================================
INSERT INTO business_types (code, name, description, icon, color, sort_order, is_active)
VALUES
  ('electricidad', 'Materiales Eléctricos', 'Cables, iluminación, tableros, protección eléctrica y automatización', 'zap', '#F59E0B', 40, true),
  ('sanitarios', 'Sanitarios y Griferías', 'Inodoros, griferías, bachas, duchas, muebles y accesorios de baño', 'bath', '#0EA5E9', 41, true)
ON CONFLICT (code) DO UPDATE SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  icon = EXCLUDED.icon,
  color = EXCLUDED.color,
  is_active = EXCLUDED.is_active;

-- =====================================================
-- 1. KIOSCO — 6 padres + 18 hijas, 15 marcas, 35 productos
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "bebidas", "name": "Bebidas", "level": 0},
    {"slug": "gaseosas", "name": "Gaseosas", "parent_slug": "bebidas", "level": 1},
    {"slug": "aguas-saborizadas", "name": "Aguas y Saborizadas", "parent_slug": "bebidas", "level": 1},
    {"slug": "jugos", "name": "Jugos", "parent_slug": "bebidas", "level": 1},
    {"slug": "energizantes", "name": "Energizantes", "parent_slug": "bebidas", "level": 1},
    {"slug": "cervezas", "name": "Cervezas", "parent_slug": "bebidas", "level": 1},

    {"slug": "golosinas", "name": "Golosinas", "level": 0},
    {"slug": "alfajores", "name": "Alfajores", "parent_slug": "golosinas", "level": 1},
    {"slug": "chocolates", "name": "Chocolates", "parent_slug": "golosinas", "level": 1},
    {"slug": "caramelos-chicles", "name": "Caramelos y Chicles", "parent_slug": "golosinas", "level": 1},
    {"slug": "gomitas-malvaviscos", "name": "Gomitas y Malvaviscos", "parent_slug": "golosinas", "level": 1},

    {"slug": "snacks", "name": "Snacks", "level": 0},
    {"slug": "papas-fritas", "name": "Papas Fritas", "parent_slug": "snacks", "level": 1},
    {"slug": "palitos-chizitos", "name": "Palitos y Chizitos", "parent_slug": "snacks", "level": 1},
    {"slug": "frutos-secos", "name": "Maníes y Frutos Secos", "parent_slug": "snacks", "level": 1},

    {"slug": "galletitas", "name": "Galletitas", "level": 0},
    {"slug": "galletitas-dulces", "name": "Dulces", "parent_slug": "galletitas", "level": 1},
    {"slug": "galletitas-saladas", "name": "Saladas", "parent_slug": "galletitas", "level": 1},

    {"slug": "cigarrillos-tabaco", "name": "Cigarrillos y Tabaco", "level": 0},
    {"slug": "encendedores", "name": "Encendedores", "parent_slug": "cigarrillos-tabaco", "level": 1},

    {"slug": "varios-kiosco", "name": "Varios Kiosco", "level": 0},
    {"slug": "pilas-baterias", "name": "Pilas y Baterías", "parent_slug": "varios-kiosco", "level": 1},
    {"slug": "higiene-personal", "name": "Higiene Personal", "parent_slug": "varios-kiosco", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Coca-Cola", "suggested_for_categories": ["gaseosas","jugos"]},
    {"name": "Pepsi", "suggested_for_categories": ["gaseosas"]},
    {"name": "Arcor", "suggested_for_categories": ["caramelos-chicles","gomitas-malvaviscos","galletitas-dulces"]},
    {"name": "Mondelez", "suggested_for_categories": ["chocolates","galletitas-dulces"]},
    {"name": "Bagley", "suggested_for_categories": ["galletitas-dulces","galletitas-saladas"]},
    {"name": "Lay''s", "suggested_for_categories": ["papas-fritas","palitos-chizitos"]},
    {"name": "Quilmes", "suggested_for_categories": ["cervezas"]},
    {"name": "Beldent", "suggested_for_categories": ["caramelos-chicles"]},
    {"name": "Georgalos", "suggested_for_categories": ["frutos-secos"]},
    {"name": "Felfort", "suggested_for_categories": ["chocolates"]},
    {"name": "Guaymallén", "suggested_for_categories": ["alfajores"]},
    {"name": "Red Bull", "suggested_for_categories": ["energizantes"]},
    {"name": "Bonafide", "suggested_for_categories": ["chocolates","alfajores"]},
    {"name": "Villavicencio", "suggested_for_categories": ["aguas-saborizadas"]},
    {"name": "Levité", "suggested_for_categories": ["aguas-saborizadas"]}
  ]'::jsonb,
  products = '[
    {"name": "Coca-Cola 500ml", "category_slug": "gaseosas", "brand": "Coca-Cola", "price_reference": 2000, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Sprite 500ml", "category_slug": "gaseosas", "brand": "Coca-Cola", "price_reference": 1900, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Pepsi 500ml", "category_slug": "gaseosas", "brand": "Pepsi", "price_reference": 1800, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Fanta 500ml", "category_slug": "gaseosas", "brand": "Coca-Cola", "price_reference": 1900, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Agua Villavicencio 500ml", "category_slug": "aguas-saborizadas", "brand": "Villavicencio", "price_reference": 1300, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Levité Pomelo 500ml", "category_slug": "aguas-saborizadas", "brand": "Levité", "price_reference": 1500, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Cepita 200ml", "category_slug": "jugos", "brand": "Coca-Cola", "price_reference": 900, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Baggio Multifruta 200ml", "category_slug": "jugos", "brand": "Baggio", "price_reference": 800, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Speed Max 473ml", "category_slug": "energizantes", "brand": "Speed", "price_reference": 2500, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Red Bull 250ml", "category_slug": "energizantes", "brand": "Red Bull", "price_reference": 3000, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Cerveza Quilmes 473ml", "category_slug": "cervezas", "brand": "Quilmes", "price_reference": 1800, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Cerveza Brahma 473ml", "category_slug": "cervezas", "brand": "Quilmes", "price_reference": 1500, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Alfajor Jorgito", "category_slug": "alfajores", "brand": "Jorgito", "price_reference": 700, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Alfajor Guaymallén triple", "category_slug": "alfajores", "brand": "Guaymallén", "price_reference": 500, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Alfajor Havanna", "category_slug": "alfajores", "brand": "Havanna", "price_reference": 1200, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Alfajor Capitán del Espacio", "category_slug": "alfajores", "brand": "Capitán del Espacio", "price_reference": 900, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Chocolate Milka 55g", "category_slug": "chocolates", "brand": "Mondelez", "price_reference": 1500, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Chocolate Shot 55g", "category_slug": "chocolates", "brand": "Felfort", "price_reference": 1400, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Bon o Bon", "category_slug": "chocolates", "brand": "Arcor", "price_reference": 350, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Chicle Beldent", "category_slug": "caramelos-chicles", "brand": "Beldent", "price_reference": 1000, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Caramelos Flynn Paff", "category_slug": "caramelos-chicles", "brand": "Arcor", "price_reference": 200, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Gomitas Mogul", "category_slug": "gomitas-malvaviscos", "brand": "Arcor", "price_reference": 1200, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Papas Lay''s Clásicas 47g", "category_slug": "papas-fritas", "brand": "Lay''s", "price_reference": 2000, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Doritos 45g", "category_slug": "papas-fritas", "brand": "Lay''s", "price_reference": 1800, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Cheetos 40g", "category_slug": "palitos-chizitos", "brand": "Lay''s", "price_reference": 1600, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Maní con chocolate", "category_slug": "frutos-secos", "brand": "Georgalos", "price_reference": 1400, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Maní japonés", "category_slug": "frutos-secos", "brand": "Georgalos", "price_reference": 1200, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Oreo 118g", "category_slug": "galletitas-dulces", "brand": "Mondelez", "price_reference": 1500, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Galletitas Terrabusi Variedad", "category_slug": "galletitas-dulces", "brand": "Bagley", "price_reference": 1300, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Criollitas 100g", "category_slug": "galletitas-saladas", "brand": "Bagley", "price_reference": 1300, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Encendedor BIC", "category_slug": "encendedores", "brand": "BIC", "price_reference": 800, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Pilas AA Duracell x2", "category_slug": "pilas-baterias", "brand": "Duracell", "price_reference": 2500, "unit": "unidad", "sku_prefix": "KIOSCO"},
    {"name": "Preservativos Prime x3", "category_slug": "higiene-personal", "brand": "Prime", "price_reference": 2000, "unit": "unidad", "sku_prefix": "KIOSCO"}
  ]'::jsonb,
  attributes = '[
    {"name": "Contenido Neto", "slug": "contenido-neto", "values": ["200ml","250ml","330ml","473ml","500ml","1L","1.5L","2.25L","47g","55g","100g","200g"], "applies_to_categories": ["gaseosas","aguas-saborizadas","jugos","energizantes","cervezas","chocolates","papas-fritas","palitos-chizitos","galletitas-dulces","galletitas-saladas"]},
    {"name": "Tipo Envase", "slug": "tipo-envase", "values": ["Botella PET","Lata","Paquete","Caja","Bolsa","Sobre"], "applies_to_categories": ["gaseosas","aguas-saborizadas","jugos","energizantes","cervezas","chocolates","papas-fritas","galletitas-dulces","galletitas-saladas"]}
  ]'::jsonb,
  version = '4.0.0-enriched',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'kiosco') AND is_default = true;

-- =====================================================
-- 2. ALMACÉN — 8 padres + 24 hijas, 16 marcas, 50 productos
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "almacen-seco", "name": "Almacén Seco", "level": 0},
    {"slug": "harinas-premezclas", "name": "Harinas y Premezclas", "parent_slug": "almacen-seco", "level": 1},
    {"slug": "arroz-legumbres", "name": "Arroz y Legumbres", "parent_slug": "almacen-seco", "level": 1},
    {"slug": "pastas-secas", "name": "Pastas Secas", "parent_slug": "almacen-seco", "level": 1},
    {"slug": "aceites-vinagres", "name": "Aceites y Vinagres", "parent_slug": "almacen-seco", "level": 1},
    {"slug": "conservas-enlatados", "name": "Conservas y Enlatados", "parent_slug": "almacen-seco", "level": 1},

    {"slug": "bebidas-almacen", "name": "Bebidas", "level": 0},
    {"slug": "gaseosas-aguas", "name": "Gaseosas y Aguas", "parent_slug": "bebidas-almacen", "level": 1},
    {"slug": "jugos-polvos", "name": "Jugos y Polvos", "parent_slug": "bebidas-almacen", "level": 1},
    {"slug": "cervezas-vinos", "name": "Cervezas y Vinos", "parent_slug": "bebidas-almacen", "level": 1},

    {"slug": "lacteos-frescos", "name": "Lácteos y Frescos", "level": 0},
    {"slug": "leches", "name": "Leche", "parent_slug": "lacteos-frescos", "level": 1},
    {"slug": "yogures", "name": "Yogures", "parent_slug": "lacteos-frescos", "level": 1},
    {"slug": "quesos-manteca", "name": "Quesos y Manteca", "parent_slug": "lacteos-frescos", "level": 1},

    {"slug": "panaderia-reposteria", "name": "Panadería y Repostería", "level": 0},
    {"slug": "pan-envasado", "name": "Pan Envasado", "parent_slug": "panaderia-reposteria", "level": 1},
    {"slug": "galletitas-almacen", "name": "Galletitas", "parent_slug": "panaderia-reposteria", "level": 1},

    {"slug": "golosinas-snacks", "name": "Golosinas y Snacks", "level": 0},
    {"slug": "alfajores-chocolates", "name": "Alfajores y Chocolates", "parent_slug": "golosinas-snacks", "level": 1},
    {"slug": "snacks-salados-alm", "name": "Snacks Salados", "parent_slug": "golosinas-snacks", "level": 1},

    {"slug": "limpieza", "name": "Limpieza", "level": 0},
    {"slug": "lavandina-desinfectantes", "name": "Lavandina y Desinfectantes", "parent_slug": "limpieza", "level": 1},
    {"slug": "detergentes-jabones", "name": "Detergentes y Jabones", "parent_slug": "limpieza", "level": 1},
    {"slug": "papel-higienico", "name": "Papel Higiénico y Servilletas", "parent_slug": "limpieza", "level": 1},

    {"slug": "perfumeria", "name": "Perfumería e Higiene", "level": 0},
    {"slug": "shampoo-acondicionador", "name": "Shampoo y Acondicionador", "parent_slug": "perfumeria", "level": 1},
    {"slug": "jabones-desodorantes", "name": "Jabones y Desodorantes", "parent_slug": "perfumeria", "level": 1},
    {"slug": "panales", "name": "Pañales", "parent_slug": "perfumeria", "level": 1},

    {"slug": "fiambreria", "name": "Fiambrería", "level": 0},
    {"slug": "fiambres-embutidos", "name": "Fiambres y Embutidos", "parent_slug": "fiambreria", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Marolio", "suggested_for_categories": ["aceites-vinagres","conservas-enlatados"]},
    {"name": "La Serenísima", "suggested_for_categories": ["leches","yogures","quesos-manteca"]},
    {"name": "Arcor", "suggested_for_categories": ["conservas-enlatados","alfajores-chocolates"]},
    {"name": "Molinos Río de la Plata", "suggested_for_categories": ["pastas-secas","harinas-premezclas","arroz-legumbres"]},
    {"name": "Coca-Cola", "suggested_for_categories": ["gaseosas-aguas"]},
    {"name": "Ledesma", "suggested_for_categories": ["harinas-premezclas"]},
    {"name": "La Campagnola", "suggested_for_categories": ["conservas-enlatados"]},
    {"name": "Fargo", "suggested_for_categories": ["pan-envasado"]},
    {"name": "Bagley", "suggested_for_categories": ["galletitas-almacen","snacks-salados-alm"]},
    {"name": "Skip", "suggested_for_categories": ["detergentes-jabones"]},
    {"name": "Higienol", "suggested_for_categories": ["papel-higienico"]},
    {"name": "Dove", "suggested_for_categories": ["jabones-desodorantes","shampoo-acondicionador"]},
    {"name": "Paladini", "suggested_for_categories": ["fiambres-embutidos"]},
    {"name": "Sancor", "suggested_for_categories": ["leches","yogures"]},
    {"name": "Natura", "suggested_for_categories": ["aceites-vinagres"]},
    {"name": "Hellmann''s", "suggested_for_categories": ["conservas-enlatados"]}
  ]'::jsonb,
  products = '[
    {"name": "Aceite girasol Natura 1.5L", "category_slug": "aceites-vinagres", "brand": "Natura", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Aceite oliva 500ml", "category_slug": "aceites-vinagres", "brand": "Cocinero", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Vinagre 1L", "category_slug": "aceites-vinagres", "brand": "Marolio", "price_reference": 1000, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Harina 000 1kg", "category_slug": "harinas-premezclas", "brand": "Blancaflor", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Azúcar 1kg", "category_slug": "harinas-premezclas", "brand": "Ledesma", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Arroz largo fino 1kg", "category_slug": "arroz-legumbres", "brand": "Molinos", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Lentejas 500g", "category_slug": "arroz-legumbres", "brand": "Marolio", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Fideos Matarazzo 500g", "category_slug": "pastas-secas", "brand": "Molinos Río de la Plata", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Fideos Lucchetti 500g", "category_slug": "pastas-secas", "brand": "Lucchetti", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Tomate triturado 520g", "category_slug": "conservas-enlatados", "brand": "Marolio", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Atún en aceite", "category_slug": "conservas-enlatados", "brand": "La Campagnola", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Arvejas lata", "category_slug": "conservas-enlatados", "brand": "Arcor", "price_reference": 800, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Mayonesa Hellmann''s 475g", "category_slug": "conservas-enlatados", "brand": "Hellmann''s", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Coca-Cola 2.25L", "category_slug": "gaseosas-aguas", "brand": "Coca-Cola", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Agua mineral 2L", "category_slug": "gaseosas-aguas", "brand": "Villavicencio", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Jugo Tang polvo", "category_slug": "jugos-polvos", "brand": "Tang", "price_reference": 600, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Leche La Serenísima 1L", "category_slug": "leches", "brand": "La Serenísima", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Leche Sancor 1L", "category_slug": "leches", "brand": "Sancor", "price_reference": 1700, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Yogur La Serenísima 1kg", "category_slug": "yogures", "brand": "La Serenísima", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Queso cremoso kg", "category_slug": "quesos-manteca", "brand": "La Serenísima", "price_reference": 8000, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Manteca 200g", "category_slug": "quesos-manteca", "brand": "La Serenísima", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Pan lactal Fargo", "category_slug": "pan-envasado", "brand": "Fargo", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Galletitas Terrabusi", "category_slug": "galletitas-almacen", "brand": "Bagley", "price_reference": 1300, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Yerba mate 1kg", "category_slug": "almacen-seco", "brand": "Rosamonte", "price_reference": 4200, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Café torrado 500g", "category_slug": "almacen-seco", "brand": "Nescafé", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Dulce de leche 400g", "category_slug": "almacen-seco", "brand": "La Serenísima", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Jamón cocido kg", "category_slug": "fiambres-embutidos", "brand": "Paladini", "price_reference": 12000, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Salame kg", "category_slug": "fiambres-embutidos", "brand": "Paladini", "price_reference": 10000, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Lavandina 2L", "category_slug": "lavandina-desinfectantes", "brand": "Ayudín", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Detergente Magistral 750ml", "category_slug": "detergentes-jabones", "brand": "Magistral", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Jabón en polvo Skip 800g", "category_slug": "detergentes-jabones", "brand": "Skip", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Papel higiénico x4", "category_slug": "papel-higienico", "brand": "Higienol", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Shampoo Dove 400ml", "category_slug": "shampoo-acondicionador", "brand": "Dove", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMACEN"},
    {"name": "Jabón tocador Dove x3", "category_slug": "jabones-desodorantes", "brand": "Dove", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMACEN"}
  ]'::jsonb,
  attributes = '[
    {"name": "Contenido Neto", "slug": "contenido-neto", "values": ["200g","500g","1kg","2.5kg","200ml","500ml","1L","1.5L","2L","2.25L"], "applies_to_categories": ["gaseosas-aguas","jugos-polvos","leches","aceites-vinagres","harinas-premezclas","arroz-legumbres","pastas-secas","conservas-enlatados"]},
    {"name": "Tipo Envase", "slug": "tipo-envase", "values": ["Botella","Sachet","Caja","Bolsa","Lata","Frasco","Tetra Brik","Bidón"], "applies_to_categories": ["gaseosas-aguas","leches","aceites-vinagres","conservas-enlatados","detergentes-jabones"]}
  ]'::jsonb,
  version = '4.0.0-enriched',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'almacen') AND is_default = true;

-- =====================================================
-- 3. FERRETERÍA — 8 padres + 28 hijas, 15 marcas, 45 productos
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "materiales-construccion", "name": "Materiales de Construcción", "level": 0},
    {"slug": "cementos-cales", "name": "Cementos y Cales", "parent_slug": "materiales-construccion", "level": 1},
    {"slug": "ladrillos-bloques", "name": "Ladrillos y Bloques", "parent_slug": "materiales-construccion", "level": 1},
    {"slug": "hierros-acero", "name": "Hierros y Acero", "parent_slug": "materiales-construccion", "level": 1},
    {"slug": "membranas-aislantes", "name": "Membranas y Aislantes", "parent_slug": "materiales-construccion", "level": 1},
    {"slug": "arenas-aridos", "name": "Arenas y Áridos", "parent_slug": "materiales-construccion", "level": 1},

    {"slug": "herramientas-manuales", "name": "Herramientas Manuales", "level": 0},
    {"slug": "martillos-mazas", "name": "Martillos y Mazas", "parent_slug": "herramientas-manuales", "level": 1},
    {"slug": "destornilladores", "name": "Destornilladores", "parent_slug": "herramientas-manuales", "level": 1},
    {"slug": "llaves-pinzas", "name": "Llaves y Pinzas", "parent_slug": "herramientas-manuales", "level": 1},
    {"slug": "sierras-serruchos", "name": "Sierras y Serruchos", "parent_slug": "herramientas-manuales", "level": 1},
    {"slug": "medicion-trazado", "name": "Medición y Trazado", "parent_slug": "herramientas-manuales", "level": 1},

    {"slug": "herramientas-electricas", "name": "Herramientas Eléctricas", "level": 0},
    {"slug": "taladros-percutores", "name": "Taladros y Percutores", "parent_slug": "herramientas-electricas", "level": 1},
    {"slug": "amoladoras", "name": "Amoladoras", "parent_slug": "herramientas-electricas", "level": 1},
    {"slug": "sierras-electricas", "name": "Sierras Eléctricas", "parent_slug": "herramientas-electricas", "level": 1},
    {"slug": "atornilladores-electricos", "name": "Atornilladores", "parent_slug": "herramientas-electricas", "level": 1},
    {"slug": "accesorios-mechas", "name": "Accesorios y Mechas", "parent_slug": "herramientas-electricas", "level": 1},

    {"slug": "tornilleria-fijaciones", "name": "Tornillería y Fijaciones", "level": 0},
    {"slug": "tornillos", "name": "Tornillos", "parent_slug": "tornilleria-fijaciones", "level": 1},
    {"slug": "clavos", "name": "Clavos", "parent_slug": "tornilleria-fijaciones", "level": 1},
    {"slug": "tarugos-bulones", "name": "Tarugos y Bulones", "parent_slug": "tornilleria-fijaciones", "level": 1},

    {"slug": "plomeria", "name": "Plomería", "level": 0},
    {"slug": "canos-pvc", "name": "Caños PVC", "parent_slug": "plomeria", "level": 1},
    {"slug": "conexiones-accesorios-plom", "name": "Conexiones y Accesorios", "parent_slug": "plomeria", "level": 1},
    {"slug": "griferias-ferret", "name": "Griferías", "parent_slug": "plomeria", "level": 1},
    {"slug": "tanques-flotantes", "name": "Tanques y Flotantes", "parent_slug": "plomeria", "level": 1},

    {"slug": "electricidad-ferret", "name": "Electricidad", "level": 0},
    {"slug": "cables-ferret", "name": "Cables", "parent_slug": "electricidad-ferret", "level": 1},
    {"slug": "llaves-termicas", "name": "Llaves Térmicas y Disyuntores", "parent_slug": "electricidad-ferret", "level": 1},
    {"slug": "tomas-interruptores-ferret", "name": "Tomas e Interruptores", "parent_slug": "electricidad-ferret", "level": 1},
    {"slug": "iluminacion-ferret", "name": "Iluminación", "parent_slug": "electricidad-ferret", "level": 1},

    {"slug": "pinturas-accesorios", "name": "Pinturas y Accesorios", "level": 0},
    {"slug": "latex-ferret", "name": "Látex Interior/Exterior", "parent_slug": "pinturas-accesorios", "level": 1},
    {"slug": "esmaltes-ferret", "name": "Esmaltes", "parent_slug": "pinturas-accesorios", "level": 1},
    {"slug": "rodillos-pinceles-ferret", "name": "Rodillos y Pinceles", "parent_slug": "pinturas-accesorios", "level": 1},

    {"slug": "ferreteria-general", "name": "Ferretería General", "level": 0},
    {"slug": "candados-cerraduras", "name": "Candados y Cerraduras", "parent_slug": "ferreteria-general", "level": 1},
    {"slug": "bisagras-herrajes", "name": "Bisagras y Herrajes", "parent_slug": "ferreteria-general", "level": 1},
    {"slug": "cintas-adhesivos-ferret", "name": "Cintas y Adhesivos", "parent_slug": "ferreteria-general", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Loma Negra", "suggested_for_categories": ["cementos-cales"]},
    {"name": "Avellaneda", "suggested_for_categories": ["cementos-cales"]},
    {"name": "Acindar", "suggested_for_categories": ["hierros-acero"]},
    {"name": "Stanley", "suggested_for_categories": ["martillos-mazas","destornilladores","llaves-pinzas","medicion-trazado"]},
    {"name": "Bosch", "suggested_for_categories": ["taladros-percutores","amoladoras","accesorios-mechas"]},
    {"name": "Black+Decker", "suggested_for_categories": ["taladros-percutores","sierras-electricas"]},
    {"name": "Makita", "suggested_for_categories": ["atornilladores-electricos","taladros-percutores"]},
    {"name": "DeWalt", "suggested_for_categories": ["amoladoras","sierras-electricas"]},
    {"name": "Tramontina", "suggested_for_categories": ["llaves-pinzas","sierras-serruchos"]},
    {"name": "Tigre", "suggested_for_categories": ["canos-pvc","conexiones-accesorios-plom"]},
    {"name": "FV", "suggested_for_categories": ["griferias-ferret","tanques-flotantes"]},
    {"name": "Alba", "suggested_for_categories": ["latex-ferret","esmaltes-ferret"]},
    {"name": "Sinteplast", "suggested_for_categories": ["latex-ferret"]},
    {"name": "Sika", "suggested_for_categories": ["membranas-aislantes"]},
    {"name": "Schneider", "suggested_for_categories": ["llaves-termicas"]}
  ]'::jsonb,
  products = '[
    {"name": "Cemento Portland Loma Negra 50kg", "category_slug": "cementos-cales", "brand": "Loma Negra", "price_reference": 9500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Cal hidráulica 25kg", "category_slug": "cementos-cales", "brand": "Loma Negra", "price_reference": 3500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Ladrillo hueco 12x18x33", "category_slug": "ladrillos-bloques", "brand": null, "price_reference": 180, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Bloque de hormigón", "category_slug": "ladrillos-bloques", "brand": null, "price_reference": 250, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Hierro ø8mm x 12m", "category_slug": "hierros-acero", "brand": "Acindar", "price_reference": 8500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Hierro ø10mm x 12m", "category_slug": "hierros-acero", "brand": "Acindar", "price_reference": 13000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Membrana asfáltica 4mm x 10m²", "category_slug": "membranas-aislantes", "brand": "Sika", "price_reference": 35000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Arena bolsa 25kg", "category_slug": "arenas-aridos", "brand": null, "price_reference": 1500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Martillo carpintero Stanley 500g", "category_slug": "martillos-mazas", "brand": "Stanley", "price_reference": 12000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Destornillador Phillips Stanley", "category_slug": "destornilladores", "brand": "Stanley", "price_reference": 3500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Llave francesa 10\"", "category_slug": "llaves-pinzas", "brand": "Stanley", "price_reference": 8000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Pinza universal Tramontina", "category_slug": "llaves-pinzas", "brand": "Tramontina", "price_reference": 4500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Cinta métrica 5m Stanley", "category_slug": "medicion-trazado", "brand": "Stanley", "price_reference": 3000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Nivel de burbuja 40cm", "category_slug": "medicion-trazado", "brand": "Stanley", "price_reference": 5000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Taladro percutor Bosch 700W", "category_slug": "taladros-percutores", "brand": "Bosch", "price_reference": 45000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Amoladora angular 4½\" Bosch", "category_slug": "amoladoras", "brand": "Bosch", "price_reference": 38000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Sierra caladora Black+Decker", "category_slug": "sierras-electricas", "brand": "Black+Decker", "price_reference": 32000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Atornillador inalámbrico Makita", "category_slug": "atornilladores-electricos", "brand": "Makita", "price_reference": 55000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Mecha para cemento 10mm Bosch", "category_slug": "accesorios-mechas", "brand": "Bosch", "price_reference": 2500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Disco abrasivo 4½\"", "category_slug": "accesorios-mechas", "brand": "Bosch", "price_reference": 1500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Tornillo autoperforante 8x1 x100", "category_slug": "tornillos", "brand": null, "price_reference": 1500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Clavo 2\" x1kg", "category_slug": "clavos", "brand": null, "price_reference": 2000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Tarugo 8mm x100", "category_slug": "tarugos-bulones", "brand": null, "price_reference": 800, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Caño PVC 110mm x 4m Tigre", "category_slug": "canos-pvc", "brand": "Tigre", "price_reference": 8000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Codo PVC 110mm", "category_slug": "conexiones-accesorios-plom", "brand": "Tigre", "price_reference": 600, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Grifería cocina FV", "category_slug": "griferias-ferret", "brand": "FV", "price_reference": 25000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Flotante para tanque FV", "category_slug": "tanques-flotantes", "brand": "FV", "price_reference": 2500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Cable 2.5mm² x 100m", "category_slug": "cables-ferret", "brand": null, "price_reference": 18000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Llave térmica 20A Schneider", "category_slug": "llaves-termicas", "brand": "Schneider", "price_reference": 5000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Tomacorriente doble", "category_slug": "tomas-interruptores-ferret", "brand": null, "price_reference": 1200, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Lámpara LED 9W", "category_slug": "iluminacion-ferret", "brand": null, "price_reference": 1500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Pintura látex interior Alba 4L", "category_slug": "latex-ferret", "brand": "Alba", "price_reference": 15000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Esmalte sintético 1L", "category_slug": "esmaltes-ferret", "brand": "Sinteplast", "price_reference": 8000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Rodillo lana 23cm", "category_slug": "rodillos-pinceles-ferret", "brand": null, "price_reference": 3500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Candado 40mm", "category_slug": "candados-cerraduras", "brand": null, "price_reference": 3000, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Bisagra 3\"", "category_slug": "bisagras-herrajes", "brand": null, "price_reference": 500, "unit": "unidad", "sku_prefix": "FERRET"},
    {"name": "Cinta ducto 48mm", "category_slug": "cintas-adhesivos-ferret", "brand": null, "price_reference": 2500, "unit": "unidad", "sku_prefix": "FERRET"}
  ]'::jsonb,
  attributes = '[
    {"name": "Material", "slug": "material-hogar", "values": ["Acero","PVC","Bronce","Aluminio","Hierro","Madera","Galvanizado"], "applies_to_categories": ["tornillos","clavos","tarugos-bulones","canos-pvc","conexiones-accesorios-plom","griferias-ferret"]},
    {"name": "Tamaño", "slug": "tamano", "values": ["Pequeño","Mediano","Grande","Extra Grande"], "applies_to_categories": ["cementos-cales","ladrillos-bloques","arenas-aridos"]},
    {"name": "Potencia", "slug": "potencia", "values": ["500W","700W","800W","1000W","1200W"], "applies_to_categories": ["taladros-percutores","amoladoras","sierras-electricas"]}
  ]'::jsonb,
  version = '4.0.0-enriched',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'ferreteria') AND is_default = true;

-- =====================================================
-- 4. PINTURERÍA — 6 padres + 20 hijas, 10 marcas, 25 productos
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "pinturas-latex", "name": "Pinturas Látex", "level": 0},
    {"slug": "latex-interior", "name": "Interior", "parent_slug": "pinturas-latex", "level": 1},
    {"slug": "latex-exterior", "name": "Exterior", "parent_slug": "pinturas-latex", "level": 1},
    {"slug": "latex-cielos-rasos", "name": "Cielos Rasos", "parent_slug": "pinturas-latex", "level": 1},

    {"slug": "esmaltes-barnices", "name": "Esmaltes y Barnices", "level": 0},
    {"slug": "esmalte-sintetico", "name": "Esmalte Sintético", "parent_slug": "esmaltes-barnices", "level": 1},
    {"slug": "barniz", "name": "Barniz", "parent_slug": "esmaltes-barnices", "level": 1},
    {"slug": "convertidor-oxido", "name": "Convertidor de Óxido", "parent_slug": "esmaltes-barnices", "level": 1},

    {"slug": "impermeabilizantes", "name": "Impermeabilizantes", "level": 0},
    {"slug": "membrana-liquida", "name": "Membrana Líquida", "parent_slug": "impermeabilizantes", "level": 1},
    {"slug": "selladores-siliconas", "name": "Selladores y Siliconas", "parent_slug": "impermeabilizantes", "level": 1},

    {"slug": "preparacion-superficies", "name": "Preparación de Superficies", "level": 0},
    {"slug": "enduido-plastico", "name": "Enduido Plástico", "parent_slug": "preparacion-superficies", "level": 1},
    {"slug": "fijador-sellador", "name": "Fijador Sellador", "parent_slug": "preparacion-superficies", "level": 1},
    {"slug": "masilla", "name": "Masilla", "parent_slug": "preparacion-superficies", "level": 1},

    {"slug": "revestimientos", "name": "Revestimientos", "level": 0},
    {"slug": "texturados", "name": "Texturados", "parent_slug": "revestimientos", "level": 1},

    {"slug": "accesorios-pintura", "name": "Accesorios de Pintura", "level": 0},
    {"slug": "rodillos-pinceles", "name": "Rodillos y Pinceles", "parent_slug": "accesorios-pintura", "level": 1},
    {"slug": "lijas", "name": "Lijas", "parent_slug": "accesorios-pintura", "level": 1},
    {"slug": "cintas-enmascarar", "name": "Cintas de Enmascarar", "parent_slug": "accesorios-pintura", "level": 1},
    {"slug": "espatulas-llanas", "name": "Espátulas y Llanas", "parent_slug": "accesorios-pintura", "level": 1},
    {"slug": "bandejas-kits", "name": "Bandejas y Kits", "parent_slug": "accesorios-pintura", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Alba", "suggested_for_categories": ["latex-interior","latex-exterior","esmalte-sintetico","enduido-plastico"]},
    {"name": "Sinteplast", "suggested_for_categories": ["latex-interior","latex-exterior","enduido-plastico"]},
    {"name": "Sherwin Williams", "suggested_for_categories": ["latex-interior","latex-exterior","esmalte-sintetico"]},
    {"name": "Tersuave", "suggested_for_categories": ["latex-interior","latex-exterior"]},
    {"name": "Colorín", "suggested_for_categories": ["esmalte-sintetico","barniz"]},
    {"name": "Riopint", "suggested_for_categories": ["esmalte-sintetico","latex-exterior"]},
    {"name": "Casablanca", "suggested_for_categories": ["latex-interior","texturados"]},
    {"name": "Cetol", "suggested_for_categories": ["barniz"]},
    {"name": "Sika", "suggested_for_categories": ["membrana-liquida","selladores-siliconas"]},
    {"name": "Plavicon", "suggested_for_categories": ["membrana-liquida","impermeabilizantes"]}
  ]'::jsonb,
  products = '[
    {"name": "Alba Albalatex Interior Mate 4L", "category_slug": "latex-interior", "brand": "Alba", "price_reference": 15000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Alba Albalatex Interior Mate 20L", "category_slug": "latex-interior", "brand": "Alba", "price_reference": 55000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Sinteplast Recuplast Interior 4L", "category_slug": "latex-interior", "brand": "Sinteplast", "price_reference": 12000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Sherwin Williams Z10 4L", "category_slug": "latex-interior", "brand": "Sherwin Williams", "price_reference": 22000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Tersuave Látex Exterior 4L", "category_slug": "latex-exterior", "brand": "Tersuave", "price_reference": 14000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Alba Exterior 10L", "category_slug": "latex-exterior", "brand": "Alba", "price_reference": 38000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Albalux Esmalte Sintético 1L", "category_slug": "esmalte-sintetico", "brand": "Alba", "price_reference": 8000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Riopint Riolux Esmalte 4L", "category_slug": "esmalte-sintetico", "brand": "Riopint", "price_reference": 18000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Cetol Classic Natural 1L", "category_slug": "barniz", "brand": "Cetol", "price_reference": 12000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Convertidor de Óxido 1L", "category_slug": "convertidor-oxido", "brand": "Alba", "price_reference": 6000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Membrana Líquida Sika 20kg", "category_slug": "membrana-liquida", "brand": "Sika", "price_reference": 45000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Silicona transparente 280ml", "category_slug": "selladores-siliconas", "brand": "Sika", "price_reference": 3500, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Albaplast Enduido Interior 4L", "category_slug": "enduido-plastico", "brand": "Alba", "price_reference": 8000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Fijador Sellador 4L", "category_slug": "fijador-sellador", "brand": "Alba", "price_reference": 6000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Texturado Sinteplast 25kg", "category_slug": "texturados", "brand": "Sinteplast", "price_reference": 28000, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Rodillo lana 23cm", "category_slug": "rodillos-pinceles", "brand": null, "price_reference": 3500, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Pincel 2\"", "category_slug": "rodillos-pinceles", "brand": null, "price_reference": 1500, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Lija al agua 220", "category_slug": "lijas", "brand": null, "price_reference": 300, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Cinta de enmascarar 18mm", "category_slug": "cintas-enmascarar", "brand": null, "price_reference": 800, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Espátula acero 10cm", "category_slug": "espatulas-llanas", "brand": null, "price_reference": 1500, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Bandeja para rodillo", "category_slug": "bandejas-kits", "brand": null, "price_reference": 1200, "unit": "unidad", "sku_prefix": "PINTUR"},
    {"name": "Diluyente 1L", "category_slug": "accesorios-pintura", "brand": null, "price_reference": 3000, "unit": "unidad", "sku_prefix": "PINTUR"}
  ]'::jsonb,
  attributes = '[
    {"name": "Presentación", "slug": "presentacion-pintura", "values": ["1L","4L","10L","20L"], "applies_to_categories": ["latex-interior","latex-exterior","esmalte-sintetico","barniz"]},
    {"name": "Acabado", "slug": "acabado", "values": ["Mate","Satinado","Brillante","Semi-brillante"], "applies_to_categories": ["latex-interior","latex-exterior","esmalte-sintetico"]},
    {"name": "Color", "slug": "color", "values": ["Blanco","Negro","Gris","Beige","Celeste","Verde","Rojo","Amarillo","Natural","Nogal"], "applies_to_categories": ["latex-interior","latex-exterior","esmalte-sintetico","barniz"]}
  ]'::jsonb,
  version = '4.0.0-enriched',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'pintureria') AND is_default = true;

-- =====================================================
-- 5. LIBRERÍA — 6 padres + 22 hijas, 11 marcas, 30 productos
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "escritura", "name": "Escritura", "level": 0},
    {"slug": "lapiceras-boligrafos", "name": "Lapiceras y Bolígrafos", "parent_slug": "escritura", "level": 1},
    {"slug": "lapices-portaminas", "name": "Lápices y Portaminas", "parent_slug": "escritura", "level": 1},
    {"slug": "marcadores-resaltadores", "name": "Marcadores y Resaltadores", "parent_slug": "escritura", "level": 1},
    {"slug": "fibras-fibrones", "name": "Fibras y Fibrones", "parent_slug": "escritura", "level": 1},

    {"slug": "papeleria", "name": "Papelería", "level": 0},
    {"slug": "hojas-resmas", "name": "Hojas y Resmas", "parent_slug": "papeleria", "level": 1},
    {"slug": "sobres-etiquetas", "name": "Sobres y Etiquetas", "parent_slug": "papeleria", "level": 1},
    {"slug": "papel-especial", "name": "Papel Especial (Glasé, Crepé)", "parent_slug": "papeleria", "level": 1},

    {"slug": "escolar", "name": "Escolar", "level": 0},
    {"slug": "cuadernos", "name": "Cuadernos", "parent_slug": "escolar", "level": 1},
    {"slug": "carpetas-biblioratos", "name": "Carpetas y Biblioratos", "parent_slug": "escolar", "level": 1},
    {"slug": "cartucheras", "name": "Cartucheras", "parent_slug": "escolar", "level": 1},
    {"slug": "mochilas", "name": "Mochilas", "parent_slug": "escolar", "level": 1},
    {"slug": "reglas-compases", "name": "Reglas y Compases", "parent_slug": "escolar", "level": 1},
    {"slug": "adhesivos-cintas", "name": "Adhesivos y Cintas", "parent_slug": "escolar", "level": 1},

    {"slug": "oficina", "name": "Oficina", "level": 0},
    {"slug": "abrochadoras-broches", "name": "Abrochadoras y Broches", "parent_slug": "oficina", "level": 1},
    {"slug": "clips-chinches", "name": "Clips y Chinches", "parent_slug": "oficina", "level": 1},
    {"slug": "calculadoras", "name": "Calculadoras", "parent_slug": "oficina", "level": 1},
    {"slug": "archivos-organizadores", "name": "Archivos y Organizadores", "parent_slug": "oficina", "level": 1},

    {"slug": "arte-manualidades", "name": "Arte y Manualidades", "level": 0},
    {"slug": "pinturas-arte", "name": "Pinturas (Acrílicos, Témperas)", "parent_slug": "arte-manualidades", "level": 1},
    {"slug": "lapices-colores", "name": "Lápices de Colores", "parent_slug": "arte-manualidades", "level": 1},
    {"slug": "goma-eva-cartulinas", "name": "Goma Eva y Cartulinas", "parent_slug": "arte-manualidades", "level": 1},

    {"slug": "tecnologia-impresion", "name": "Tecnología e Impresión", "level": 0},
    {"slug": "cartuchos-toners", "name": "Cartuchos y Tóners", "parent_slug": "tecnologia-impresion", "level": 1},
    {"slug": "accesorios-pc", "name": "Accesorios PC", "parent_slug": "tecnologia-impresion", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Faber-Castell", "suggested_for_categories": ["lapices-portaminas","lapices-colores","marcadores-resaltadores"]},
    {"name": "Rivadavia", "suggested_for_categories": ["cuadernos"]},
    {"name": "Maped", "suggested_for_categories": ["reglas-compases","adhesivos-cintas","fibras-fibrones"]},
    {"name": "Staedtler", "suggested_for_categories": ["lapices-portaminas","marcadores-resaltadores"]},
    {"name": "BIC", "suggested_for_categories": ["lapiceras-boligrafos"]},
    {"name": "Filgo", "suggested_for_categories": ["lapiceras-boligrafos","marcadores-resaltadores"]},
    {"name": "Pelikan", "suggested_for_categories": ["lapiceras-boligrafos","pinturas-arte"]},
    {"name": "Mooving", "suggested_for_categories": ["cartucheras","mochilas"]},
    {"name": "Ledesma", "suggested_for_categories": ["hojas-resmas"]},
    {"name": "Éxito", "suggested_for_categories": ["cuadernos","carpetas-biblioratos"]},
    {"name": "Simball", "suggested_for_categories": ["cuadernos"]}
  ]'::jsonb,
  products = '[
    {"name": "Bolígrafo BIC Cristal azul", "category_slug": "lapiceras-boligrafos", "brand": "BIC", "price_reference": 400, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Roller borrable Frixion", "category_slug": "lapiceras-boligrafos", "brand": "Pilot", "price_reference": 2500, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Lápiz grafito Faber-Castell HB", "category_slug": "lapices-portaminas", "brand": "Faber-Castell", "price_reference": 500, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Portaminas 0.5mm Staedtler", "category_slug": "lapices-portaminas", "brand": "Staedtler", "price_reference": 2000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Resaltador Filgo x4", "category_slug": "marcadores-resaltadores", "brand": "Filgo", "price_reference": 3500, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Marcadores Maped Color Peps x12", "category_slug": "fibras-fibrones", "brand": "Maped", "price_reference": 6000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Resma A4 500 hojas Ledesma", "category_slug": "hojas-resmas", "brand": "Ledesma", "price_reference": 8000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Papel Glasé x10", "category_slug": "papel-especial", "brand": null, "price_reference": 600, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Cuaderno Rivadavia tapa dura 48h", "category_slug": "cuadernos", "brand": "Rivadavia", "price_reference": 7600, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Cuaderno Rivadavia tapa dura 98h", "category_slug": "cuadernos", "brand": "Rivadavia", "price_reference": 9800, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Cuaderno espiralado A4 84h", "category_slug": "cuadernos", "brand": "Éxito", "price_reference": 5500, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Carpeta 3 anillos", "category_slug": "carpetas-biblioratos", "brand": null, "price_reference": 4000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Bibliorato oficio", "category_slug": "carpetas-biblioratos", "brand": null, "price_reference": 5000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Cartuchera simple", "category_slug": "cartucheras", "brand": "Mooving", "price_reference": 4000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Mochila escolar", "category_slug": "mochilas", "brand": "Mooving", "price_reference": 25000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Regla 30cm Maped", "category_slug": "reglas-compases", "brand": "Maped", "price_reference": 800, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Compás Maped", "category_slug": "reglas-compases", "brand": "Maped", "price_reference": 3000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Plasticola 40g", "category_slug": "adhesivos-cintas", "brand": null, "price_reference": 800, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Tijera escolar Maped 13cm", "category_slug": "adhesivos-cintas", "brand": "Maped", "price_reference": 2000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Abrochadora", "category_slug": "abrochadoras-broches", "brand": null, "price_reference": 3500, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Calculadora Casio", "category_slug": "calculadoras", "brand": "Casio", "price_reference": 8000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Lápices de colores Faber-Castell x12", "category_slug": "lapices-colores", "brand": "Faber-Castell", "price_reference": 5000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Lápices de colores Faber-Castell x24", "category_slug": "lapices-colores", "brand": "Faber-Castell", "price_reference": 9000, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Témperas x6", "category_slug": "pinturas-arte", "brand": null, "price_reference": 3500, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Goma Eva A4 x10", "category_slug": "goma-eva-cartulinas", "brand": null, "price_reference": 2500, "unit": "unidad", "sku_prefix": "LIBRE"},
    {"name": "Cartulina colores x10", "category_slug": "goma-eva-cartulinas", "brand": null, "price_reference": 1500, "unit": "unidad", "sku_prefix": "LIBRE"}
  ]'::jsonb,
  attributes = '[
    {"name": "Formato", "slug": "formato-cuaderno", "values": ["A4","A5","Oficio","Letter"], "applies_to_categories": ["cuadernos","hojas-resmas"]},
    {"name": "Cantidad de Hojas", "slug": "cantidad-hojas", "values": ["48","72","84","98","120","200","480","500"], "applies_to_categories": ["cuadernos","hojas-resmas"]},
    {"name": "Tipo de Rayado", "slug": "tipo-rayado", "values": ["Rayado","Cuadriculado","Liso","Pautado"], "applies_to_categories": ["cuadernos"]},
    {"name": "Color", "slug": "color", "values": ["Negro","Azul","Rojo","Verde","Rosa","Amarillo"], "applies_to_categories": ["lapiceras-boligrafos","cartucheras","mochilas"]}
  ]'::jsonb,
  version = '4.0.0-enriched',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'libreria') AND is_default = true;

-- =====================================================
-- 6. INDUMENTARIA/ROPA — 6 padres + 22 hijas, 12 marcas, 25 productos
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "hombre", "name": "Hombre", "level": 0},
    {"slug": "remeras-chombas", "name": "Remeras y Chombas", "parent_slug": "hombre", "level": 1},
    {"slug": "camisas", "name": "Camisas", "parent_slug": "hombre", "level": 1},
    {"slug": "pantalones-hombre", "name": "Pantalones", "parent_slug": "hombre", "level": 1},
    {"slug": "buzos-camperas-h", "name": "Buzos y Camperas", "parent_slug": "hombre", "level": 1},
    {"slug": "ropa-interior-h", "name": "Ropa Interior", "parent_slug": "hombre", "level": 1},

    {"slug": "mujer", "name": "Mujer", "level": 0},
    {"slug": "remeras-tops", "name": "Remeras y Tops", "parent_slug": "mujer", "level": 1},
    {"slug": "vestidos", "name": "Vestidos", "parent_slug": "mujer", "level": 1},
    {"slug": "pantalones-calzas", "name": "Pantalones y Calzas", "parent_slug": "mujer", "level": 1},
    {"slug": "buzos-camperas-m", "name": "Buzos y Camperas", "parent_slug": "mujer", "level": 1},
    {"slug": "polleras-shorts", "name": "Polleras y Shorts", "parent_slug": "mujer", "level": 1},

    {"slug": "ninos", "name": "Niños", "level": 0},
    {"slug": "remeras-ninos", "name": "Remeras", "parent_slug": "ninos", "level": 1},
    {"slug": "pantalones-ninos", "name": "Pantalones", "parent_slug": "ninos", "level": 1},
    {"slug": "buzos-camperas-n", "name": "Buzos y Camperas", "parent_slug": "ninos", "level": 1},

    {"slug": "calzado-ropa", "name": "Calzado", "level": 0},
    {"slug": "zapatillas-urbanas", "name": "Zapatillas Urbanas", "parent_slug": "calzado-ropa", "level": 1},
    {"slug": "zapatos-formales", "name": "Zapatos Formales", "parent_slug": "calzado-ropa", "level": 1},
    {"slug": "botas-borcegos", "name": "Botas y Borcegos", "parent_slug": "calzado-ropa", "level": 1},
    {"slug": "sandalias-ojotas", "name": "Sandalias y Ojotas", "parent_slug": "calzado-ropa", "level": 1},

    {"slug": "accesorios-ropa", "name": "Accesorios", "level": 0},
    {"slug": "cinturones", "name": "Cinturones", "parent_slug": "accesorios-ropa", "level": 1},
    {"slug": "gorras-sombreros", "name": "Gorras y Sombreros", "parent_slug": "accesorios-ropa", "level": 1},
    {"slug": "carteras-mochilas", "name": "Carteras y Mochilas", "parent_slug": "accesorios-ropa", "level": 1},
    {"slug": "bufandas-guantes", "name": "Bufandas y Guantes", "parent_slug": "accesorios-ropa", "level": 1},

    {"slug": "deportiva", "name": "Ropa Deportiva", "level": 0},
    {"slug": "conjuntos-deportivos", "name": "Conjuntos Deportivos", "parent_slug": "deportiva", "level": 1},
    {"slug": "calzas-shorts-dep", "name": "Calzas y Shorts", "parent_slug": "deportiva", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Nike", "suggested_for_categories": ["remeras-chombas","buzos-camperas-h","zapatillas-urbanas","conjuntos-deportivos"]},
    {"name": "Adidas", "suggested_for_categories": ["remeras-chombas","zapatillas-urbanas","conjuntos-deportivos"]},
    {"name": "Topper", "suggested_for_categories": ["remeras-chombas","zapatillas-urbanas"]},
    {"name": "Puma", "suggested_for_categories": ["buzos-camperas-h","zapatillas-urbanas"]},
    {"name": "Levi''s", "suggested_for_categories": ["pantalones-hombre","pantalones-calzas"]},
    {"name": "Wrangler", "suggested_for_categories": ["pantalones-hombre"]},
    {"name": "Kevingston", "suggested_for_categories": ["remeras-chombas","camisas"]},
    {"name": "Cardón", "suggested_for_categories": ["camisas","botas-borcegos"]},
    {"name": "Kosiuko", "suggested_for_categories": ["remeras-tops","vestidos"]},
    {"name": "Mimo", "suggested_for_categories": ["remeras-ninos","pantalones-ninos"]},
    {"name": "Cheeky", "suggested_for_categories": ["remeras-ninos","pantalones-ninos"]},
    {"name": "Lacoste", "suggested_for_categories": ["remeras-chombas","camisas"]}
  ]'::jsonb,
  products = '[
    {"name": "Remera algodón hombre", "category_slug": "remeras-chombas", "brand": null, "price_reference": 15000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Camisa manga larga", "category_slug": "camisas", "brand": null, "price_reference": 28000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Jean clásico recto", "category_slug": "pantalones-hombre", "brand": "Levi''s", "price_reference": 35000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Jean slim fit", "category_slug": "pantalones-hombre", "brand": "Wrangler", "price_reference": 38000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Buzo canguro frisa", "category_slug": "buzos-camperas-h", "brand": null, "price_reference": 28000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Campera rompeviento", "category_slug": "buzos-camperas-h", "brand": null, "price_reference": 42000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Remera mujer estampada", "category_slug": "remeras-tops", "brand": null, "price_reference": 18000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Vestido casual", "category_slug": "vestidos", "brand": null, "price_reference": 30000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Calza deportiva mujer", "category_slug": "pantalones-calzas", "brand": null, "price_reference": 20000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Pollera midi", "category_slug": "polleras-shorts", "brand": null, "price_reference": 22000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Remera niño", "category_slug": "remeras-ninos", "brand": "Mimo", "price_reference": 10000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Jean niño", "category_slug": "pantalones-ninos", "brand": "Cheeky", "price_reference": 18000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Zapatilla urbana Nike", "category_slug": "zapatillas-urbanas", "brand": "Nike", "price_reference": 45000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Borcego cuero", "category_slug": "botas-borcegos", "brand": null, "price_reference": 55000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Cinturón cuero", "category_slug": "cinturones", "brand": null, "price_reference": 8000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Gorra deportiva", "category_slug": "gorras-sombreros", "brand": "Adidas", "price_reference": 6000, "unit": "unidad", "sku_prefix": "INDUM"},
    {"name": "Conjunto deportivo", "category_slug": "conjuntos-deportivos", "brand": "Topper", "price_reference": 45000, "unit": "unidad", "sku_prefix": "INDUM"}
  ]'::jsonb,
  attributes = '[
    {"name": "Talle", "slug": "talle-ropa", "values": ["XS","S","M","L","XL","XXL","2","4","6","8","10","12","14"], "applies_to_categories": ["remeras-chombas","camisas","pantalones-hombre","buzos-camperas-h","remeras-tops","vestidos","pantalones-calzas","polleras-shorts","remeras-ninos","pantalones-ninos"]},
    {"name": "Color", "slug": "color", "values": ["Negro","Blanco","Azul","Gris","Rojo","Verde","Beige","Rosa","Celeste","Bordó"], "applies_to_categories": ["remeras-chombas","camisas","pantalones-hombre","remeras-tops","vestidos","pantalones-calzas"]},
    {"name": "Material", "slug": "material-textil", "values": ["Algodón","Poliéster","Jean/Denim","Frisa","Lana","Sintético","Cuero"], "applies_to_categories": ["remeras-chombas","pantalones-hombre","buzos-camperas-h","remeras-tops","vestidos"]},
    {"name": "Género", "slug": "genero", "values": ["Hombre","Mujer","Unisex","Niño","Niña"], "applies_to_categories": ["remeras-chombas","pantalones-hombre","buzos-camperas-h","remeras-tops","vestidos","zapatillas-urbanas","conjuntos-deportivos"]},
    {"name": "Talle Calzado", "slug": "talle-calzado", "values": ["34","35","36","37","38","39","40","41","42","43","44","45"], "applies_to_categories": ["zapatillas-urbanas","zapatos-formales","botas-borcegos","sandalias-ojotas"]}
  ]'::jsonb,
  version = '4.0.0-enriched',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'ropa') AND is_default = true;

-- =====================================================
-- 7. DEPORTES — 7 padres + 24 hijas, 12 marcas, 25 productos
-- =====================================================
UPDATE business_type_templates
SET
  categories = '[
    {"slug": "futbol", "name": "Fútbol", "level": 0},
    {"slug": "botines", "name": "Botines", "parent_slug": "futbol", "level": 1},
    {"slug": "camisetas-futbol", "name": "Camisetas", "parent_slug": "futbol", "level": 1},
    {"slug": "pelotas-futbol", "name": "Pelotas", "parent_slug": "futbol", "level": 1},
    {"slug": "canilleras-medias", "name": "Canilleras y Medias", "parent_slug": "futbol", "level": 1},

    {"slug": "running", "name": "Running", "level": 0},
    {"slug": "zapatillas-running", "name": "Zapatillas Running", "parent_slug": "running", "level": 1},
    {"slug": "remeras-tecnicas", "name": "Remeras Técnicas", "parent_slug": "running", "level": 1},
    {"slug": "calzas-shorts-run", "name": "Calzas y Shorts", "parent_slug": "running", "level": 1},

    {"slug": "fitness-training", "name": "Fitness y Training", "level": 0},
    {"slug": "zapatillas-training", "name": "Zapatillas Training", "parent_slug": "fitness-training", "level": 1},
    {"slug": "ropa-gimnasio", "name": "Ropa de Gimnasio", "parent_slug": "fitness-training", "level": 1},
    {"slug": "pesas-mancuernas", "name": "Pesas y Mancuernas", "parent_slug": "fitness-training", "level": 1},
    {"slug": "colchonetas-bandas", "name": "Colchonetas y Bandas", "parent_slug": "fitness-training", "level": 1},

    {"slug": "natacion", "name": "Natación", "level": 0},
    {"slug": "mallas-trajes", "name": "Mallas y Trajes", "parent_slug": "natacion", "level": 1},
    {"slug": "antiparras-gorros", "name": "Antiparras y Gorros", "parent_slug": "natacion", "level": 1},

    {"slug": "ciclismo", "name": "Ciclismo", "level": 0},
    {"slug": "bicicletas", "name": "Bicicletas", "parent_slug": "ciclismo", "level": 1},
    {"slug": "cascos-ciclismo", "name": "Cascos", "parent_slug": "ciclismo", "level": 1},
    {"slug": "accesorios-ciclismo", "name": "Accesorios y Repuestos", "parent_slug": "ciclismo", "level": 1},

    {"slug": "outdoor-camping", "name": "Outdoor y Camping", "level": 0},
    {"slug": "carpas-bolsas-dormir", "name": "Carpas y Bolsas de Dormir", "parent_slug": "outdoor-camping", "level": 1},
    {"slug": "mochilas-trekking", "name": "Mochilas Trekking", "parent_slug": "outdoor-camping", "level": 1},
    {"slug": "calzado-outdoor", "name": "Calzado Outdoor", "parent_slug": "outdoor-camping", "level": 1},

    {"slug": "otros-deportes", "name": "Otros Deportes", "level": 0},
    {"slug": "padel-tenis", "name": "Pádel y Tenis", "parent_slug": "otros-deportes", "level": 1},
    {"slug": "basquet-voley", "name": "Básquet y Vóley", "parent_slug": "otros-deportes", "level": 1}
  ]'::jsonb,
  brands = '[
    {"name": "Adidas", "suggested_for_categories": ["botines","camisetas-futbol","zapatillas-running","zapatillas-training"]},
    {"name": "Nike", "suggested_for_categories": ["zapatillas-running","ropa-gimnasio","botines"]},
    {"name": "Puma", "suggested_for_categories": ["botines","zapatillas-training","ropa-gimnasio"]},
    {"name": "Under Armour", "suggested_for_categories": ["remeras-tecnicas","ropa-gimnasio"]},
    {"name": "Topper", "suggested_for_categories": ["zapatillas-training","botines"]},
    {"name": "Fila", "suggested_for_categories": ["zapatillas-running","ropa-gimnasio"]},
    {"name": "New Balance", "suggested_for_categories": ["zapatillas-running"]},
    {"name": "Asics", "suggested_for_categories": ["zapatillas-running"]},
    {"name": "Wilson", "suggested_for_categories": ["padel-tenis"]},
    {"name": "Head", "suggested_for_categories": ["padel-tenis"]},
    {"name": "Salomon", "suggested_for_categories": ["calzado-outdoor","mochilas-trekking"]},
    {"name": "Montagne", "suggested_for_categories": ["carpas-bolsas-dormir","mochilas-trekking"]}
  ]'::jsonb,
  products = '[
    {"name": "Botines fútbol césped", "category_slug": "botines", "brand": "Adidas", "price_reference": 45000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Camiseta fútbol selección", "category_slug": "camisetas-futbol", "brand": "Adidas", "price_reference": 35000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Pelota fútbol Nº5", "category_slug": "pelotas-futbol", "brand": "Adidas", "price_reference": 15000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Canilleras", "category_slug": "canilleras-medias", "brand": null, "price_reference": 5000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Zapatillas running Nike", "category_slug": "zapatillas-running", "brand": "Nike", "price_reference": 55000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Zapatillas running Asics", "category_slug": "zapatillas-running", "brand": "Asics", "price_reference": 65000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Remera dry-fit running", "category_slug": "remeras-tecnicas", "brand": "Under Armour", "price_reference": 18000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Short running", "category_slug": "calzas-shorts-run", "brand": null, "price_reference": 12000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Zapatillas training Topper", "category_slug": "zapatillas-training", "brand": "Topper", "price_reference": 40000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Calza deportiva", "category_slug": "ropa-gimnasio", "brand": null, "price_reference": 22000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Mancuernas 5kg par", "category_slug": "pesas-mancuernas", "brand": null, "price_reference": 15000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Colchoneta yoga", "category_slug": "colchonetas-bandas", "brand": null, "price_reference": 8000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Banda elástica set x3", "category_slug": "colchonetas-bandas", "brand": null, "price_reference": 5000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Malla natación", "category_slug": "mallas-trajes", "brand": null, "price_reference": 12000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Antiparras natación", "category_slug": "antiparras-gorros", "brand": null, "price_reference": 6000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Casco ciclismo", "category_slug": "cascos-ciclismo", "brand": null, "price_reference": 25000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Carpa 4 personas", "category_slug": "carpas-bolsas-dormir", "brand": "Montagne", "price_reference": 80000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Mochila trekking 50L", "category_slug": "mochilas-trekking", "brand": "Montagne", "price_reference": 45000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Zapatillas trekking Salomon", "category_slug": "calzado-outdoor", "brand": "Salomon", "price_reference": 85000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Paleta pádel Wilson", "category_slug": "padel-tenis", "brand": "Wilson", "price_reference": 35000, "unit": "unidad", "sku_prefix": "DEPORT"},
    {"name": "Pelota básquet Nº7", "category_slug": "basquet-voley", "brand": null, "price_reference": 12000, "unit": "unidad", "sku_prefix": "DEPORT"}
  ]'::jsonb,
  attributes = '[
    {"name": "Talle", "slug": "talle-ropa", "values": ["XS","S","M","L","XL","XXL"], "applies_to_categories": ["camisetas-futbol","remeras-tecnicas","ropa-gimnasio","calzas-shorts-run","mallas-trajes"]},
    {"name": "Color", "slug": "color", "values": ["Negro","Blanco","Azul","Rojo","Verde","Gris","Amarillo","Rosa"], "applies_to_categories": ["botines","zapatillas-running","zapatillas-training","ropa-gimnasio"]},
    {"name": "Disciplina", "slug": "disciplina-deportiva", "values": ["Fútbol","Running","Training","Básquet","Tenis","Pádel","Natación","Trekking","Ciclismo"], "applies_to_categories": ["botines","zapatillas-running","zapatillas-training","pelotas-futbol","padel-tenis"]},
    {"name": "Talle Calzado", "slug": "talle-calzado", "values": ["36","37","38","39","40","41","42","43","44","45"], "applies_to_categories": ["botines","zapatillas-running","zapatillas-training","calzado-outdoor"]}
  ]'::jsonb,
  version = '4.0.0-enriched',
  updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'deportes') AND is_default = true;

-- =====================================================
-- 8. BAZAR (NUEVO) — 6 padres + 22 hijas, 11 marcas, 25 productos
-- =====================================================
INSERT INTO business_type_templates (
  business_type_id, name, description, region, categories, brands, products, attributes,
  is_default, is_active, version
)
SELECT bt.id, 'Bazar', 'Template curado para bazares argentinos', 'AR',
  '[
    {"slug": "cocina-coccion", "name": "Cocina y Cocción", "level": 0},
    {"slug": "ollas-cacerolas", "name": "Ollas y Cacerolas", "parent_slug": "cocina-coccion", "level": 1},
    {"slug": "sartenes", "name": "Sartenes", "parent_slug": "cocina-coccion", "level": 1},
    {"slug": "fuentes-asaderas", "name": "Fuentes y Asaderas", "parent_slug": "cocina-coccion", "level": 1},
    {"slug": "parrillas-planchas", "name": "Parrillas y Planchas", "parent_slug": "cocina-coccion", "level": 1},

    {"slug": "vajilla-mesa", "name": "Vajilla y Mesa", "level": 0},
    {"slug": "platos", "name": "Platos", "parent_slug": "vajilla-mesa", "level": 1},
    {"slug": "vasos-copas", "name": "Vasos y Copas", "parent_slug": "vajilla-mesa", "level": 1},
    {"slug": "cubiertos", "name": "Cubiertos", "parent_slug": "vajilla-mesa", "level": 1},
    {"slug": "mate-termos", "name": "Mate y Termos", "parent_slug": "vajilla-mesa", "level": 1},

    {"slug": "organizacion", "name": "Organización y Almacenamiento", "level": 0},
    {"slug": "contenedores-hermeticos", "name": "Contenedores Herméticos", "parent_slug": "organizacion", "level": 1},
    {"slug": "canastos-cajas", "name": "Canastos y Cajas", "parent_slug": "organizacion", "level": 1},

    {"slug": "limpieza-hogar", "name": "Limpieza del Hogar", "level": 0},
    {"slug": "baldes-escurridores", "name": "Baldes y Escurridores", "parent_slug": "limpieza-hogar", "level": 1},
    {"slug": "escobas-secadores", "name": "Escobas y Secadores", "parent_slug": "limpieza-hogar", "level": 1},
    {"slug": "cestos-basura", "name": "Cestos de Basura", "parent_slug": "limpieza-hogar", "level": 1},
    {"slug": "tendederos", "name": "Tendederos", "parent_slug": "limpieza-hogar", "level": 1},

    {"slug": "bazar-general", "name": "Bazar General", "level": 0},
    {"slug": "perchas", "name": "Perchas", "parent_slug": "bazar-general", "level": 1},
    {"slug": "relojes-marcos", "name": "Relojes y Marcos", "parent_slug": "bazar-general", "level": 1},

    {"slug": "descartables", "name": "Descartables y Envolturas", "level": 0},
    {"slug": "film-aluminio", "name": "Film y Aluminio", "parent_slug": "descartables", "level": 1},
    {"slug": "vasos-platos-desc", "name": "Vasos y Platos Descartables", "parent_slug": "descartables", "level": 1}
  ]'::jsonb,
  '[
    {"name": "Tramontina", "suggested_for_categories": ["ollas-cacerolas","sartenes","cubiertos"]},
    {"name": "Colombraro", "suggested_for_categories": ["contenedores-hermeticos","canastos-cajas"]},
    {"name": "Essen", "suggested_for_categories": ["ollas-cacerolas","sartenes","fuentes-asaderas"]},
    {"name": "Carol", "suggested_for_categories": ["sartenes","ollas-cacerolas"]},
    {"name": "Ilko", "suggested_for_categories": ["cubiertos"]},
    {"name": "Luminarc", "suggested_for_categories": ["platos","vasos-copas"]},
    {"name": "Pyrex", "suggested_for_categories": ["fuentes-asaderas"]},
    {"name": "Stanley", "suggested_for_categories": ["mate-termos"]},
    {"name": "Rubbermaid", "suggested_for_categories": ["contenedores-hermeticos"]},
    {"name": "Condor", "suggested_for_categories": ["escobas-secadores"]},
    {"name": "Magefesa", "suggested_for_categories": ["ollas-cacerolas"]}
  ]'::jsonb,
  '[
    {"name": "Olla Tramontina 24cm", "category_slug": "ollas-cacerolas", "brand": "Tramontina", "price_reference": 25000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Cacerola Essen 28cm", "category_slug": "ollas-cacerolas", "brand": "Essen", "price_reference": 45000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Sartén antiadherente 26cm", "category_slug": "sartenes", "brand": "Tramontina", "price_reference": 18000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Fuente Pyrex 3L", "category_slug": "fuentes-asaderas", "brand": "Pyrex", "price_reference": 12000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Asadera acero 35cm", "category_slug": "fuentes-asaderas", "brand": null, "price_reference": 8000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Juego platos x18 piezas", "category_slug": "platos", "brand": "Luminarc", "price_reference": 22000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Vasos set x6", "category_slug": "vasos-copas", "brand": "Luminarc", "price_reference": 8000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Juego cubiertos x24 Tramontina", "category_slug": "cubiertos", "brand": "Tramontina", "price_reference": 15000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Mate de acero", "category_slug": "mate-termos", "brand": null, "price_reference": 5000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Termo Stanley 1.2L", "category_slug": "mate-termos", "brand": "Stanley", "price_reference": 65000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Set herméticos x4", "category_slug": "contenedores-hermeticos", "brand": "Colombraro", "price_reference": 6000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Canasto organizador", "category_slug": "canastos-cajas", "brand": "Colombraro", "price_reference": 4000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Balde con escurridor", "category_slug": "baldes-escurridores", "brand": null, "price_reference": 5000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Escoba con palo", "category_slug": "escobas-secadores", "brand": "Condor", "price_reference": 3500, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Cesto basura 20L", "category_slug": "cestos-basura", "brand": null, "price_reference": 6000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Tendedero plegable", "category_slug": "tendederos", "brand": null, "price_reference": 15000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Rollo film 30m", "category_slug": "film-aluminio", "brand": null, "price_reference": 2000, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Papel aluminio 5m", "category_slug": "film-aluminio", "brand": null, "price_reference": 1500, "unit": "unidad", "sku_prefix": "BAZAR"},
    {"name": "Vasos descartables x50", "category_slug": "vasos-platos-desc", "brand": null, "price_reference": 1200, "unit": "unidad", "sku_prefix": "BAZAR"}
  ]'::jsonb,
  '[
    {"name": "Material", "slug": "material-bazar", "values": ["Acero Inoxidable","Aluminio","Vidrio","Plástico","Cerámica","Madera","Silicona"], "applies_to_categories": ["ollas-cacerolas","sartenes","fuentes-asaderas","platos","vasos-copas","cubiertos"]},
    {"name": "Tamaño", "slug": "tamano-bazar", "values": ["Chico","Mediano","Grande","Familiar"], "applies_to_categories": ["ollas-cacerolas","sartenes","fuentes-asaderas","contenedores-hermeticos","cestos-basura"]}
  ]'::jsonb,
  true, true, '4.0.0-enriched'
FROM business_types bt WHERE bt.code = 'bazar'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories, brands = EXCLUDED.brands, products = EXCLUDED.products,
  attributes = EXCLUDED.attributes, version = EXCLUDED.version,
  updated_at = CURRENT_TIMESTAMP;

-- =====================================================
-- 9. MATERIALES ELÉCTRICOS (NUEVO) — 7 padres + 26 hijas, 12 marcas, 25 productos
-- =====================================================
INSERT INTO business_type_templates (
  business_type_id, name, description, region, categories, brands, products, attributes,
  is_default, is_active, version
)
SELECT bt.id, 'Materiales Eléctricos', 'Template curado para casas de materiales eléctricos', 'AR',
  '[
    {"slug": "cables-conductores", "name": "Cables y Conductores", "level": 0},
    {"slug": "cable-unipolar", "name": "Cable Unipolar", "parent_slug": "cables-conductores", "level": 1},
    {"slug": "cable-subterraneo", "name": "Cable Subterráneo", "parent_slug": "cables-conductores", "level": 1},
    {"slug": "cable-utp-red", "name": "Cable UTP/Red", "parent_slug": "cables-conductores", "level": 1},
    {"slug": "fichas-conectores", "name": "Fichas y Conectores", "parent_slug": "cables-conductores", "level": 1},

    {"slug": "proteccion-electrica", "name": "Protección Eléctrica", "level": 0},
    {"slug": "llaves-termicas-elec", "name": "Llaves Térmicas", "parent_slug": "proteccion-electrica", "level": 1},
    {"slug": "disyuntores", "name": "Disyuntores Diferenciales", "parent_slug": "proteccion-electrica", "level": 1},
    {"slug": "fusibles", "name": "Fusibles", "parent_slug": "proteccion-electrica", "level": 1},
    {"slug": "protectores-tension", "name": "Protectores de Tensión", "parent_slug": "proteccion-electrica", "level": 1},

    {"slug": "tableros-gabinetes", "name": "Tableros y Gabinetes", "level": 0},
    {"slug": "tableros-embutir", "name": "Tableros Embutir", "parent_slug": "tableros-gabinetes", "level": 1},
    {"slug": "tableros-exterior", "name": "Tableros Exterior", "parent_slug": "tableros-gabinetes", "level": 1},
    {"slug": "rieles-barras", "name": "Rieles y Barras", "parent_slug": "tableros-gabinetes", "level": 1},

    {"slug": "iluminacion", "name": "Iluminación", "level": 0},
    {"slug": "lamparas-led", "name": "Lámparas LED", "parent_slug": "iluminacion", "level": 1},
    {"slug": "tubos-led", "name": "Tubos LED", "parent_slug": "iluminacion", "level": 1},
    {"slug": "plafones-apliques", "name": "Plafones y Apliques", "parent_slug": "iluminacion", "level": 1},
    {"slug": "proyectores-led", "name": "Proyectores", "parent_slug": "iluminacion", "level": 1},

    {"slug": "tomas-interruptores", "name": "Tomacorrientes e Interruptores", "level": 0},
    {"slug": "linea-modular", "name": "Línea Modular", "parent_slug": "tomas-interruptores", "level": 1},
    {"slug": "tapas-bastidores", "name": "Tapas y Bastidores", "parent_slug": "tomas-interruptores", "level": 1},
    {"slug": "cajas-embutir", "name": "Cajas de Embutir", "parent_slug": "tomas-interruptores", "level": 1},

    {"slug": "canalizacion", "name": "Canalización", "level": 0},
    {"slug": "canos-curvas", "name": "Caños y Curvas", "parent_slug": "canalizacion", "level": 1},
    {"slug": "canaletas", "name": "Canaletas", "parent_slug": "canalizacion", "level": 1},
    {"slug": "precintos-accesorios", "name": "Precintos y Accesorios", "parent_slug": "canalizacion", "level": 1},

    {"slug": "automatizacion", "name": "Automatización", "level": 0},
    {"slug": "timers-programadores", "name": "Timers y Programadores", "parent_slug": "automatizacion", "level": 1},
    {"slug": "sensores-movimiento", "name": "Sensores de Movimiento", "parent_slug": "automatizacion", "level": 1},
    {"slug": "fotocélulas", "name": "Fotocélulas", "parent_slug": "automatizacion", "level": 1}
  ]'::jsonb,
  '[
    {"name": "Schneider Electric", "suggested_for_categories": ["llaves-termicas-elec","disyuntores","tableros-embutir"]},
    {"name": "Siemens", "suggested_for_categories": ["llaves-termicas-elec","disyuntores","protectores-tension"]},
    {"name": "ABB", "suggested_for_categories": ["llaves-termicas-elec","disyuntores"]},
    {"name": "Jeluz", "suggested_for_categories": ["linea-modular","tapas-bastidores"]},
    {"name": "Sica", "suggested_for_categories": ["cable-unipolar","cable-subterraneo"]},
    {"name": "Kalop", "suggested_for_categories": ["cable-unipolar","fichas-conectores"]},
    {"name": "Philips", "suggested_for_categories": ["lamparas-led","tubos-led"]},
    {"name": "Ledvance", "suggested_for_categories": ["lamparas-led","tubos-led","proyectores-led"]},
    {"name": "Cambre", "suggested_for_categories": ["linea-modular","tapas-bastidores"]},
    {"name": "Roker", "suggested_for_categories": ["canos-curvas","canaletas"]},
    {"name": "TKL", "suggested_for_categories": ["cable-unipolar"]},
    {"name": "Lexel", "suggested_for_categories": ["timers-programadores","sensores-movimiento"]}
  ]'::jsonb,
  '[
    {"name": "Cable unipolar 2.5mm² x100m", "category_slug": "cable-unipolar", "brand": "Sica", "price_reference": 18000, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Cable unipolar 1.5mm² x100m", "category_slug": "cable-unipolar", "brand": "Kalop", "price_reference": 12000, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Cable subterráneo 3x2.5mm² x50m", "category_slug": "cable-subterraneo", "brand": "Sica", "price_reference": 35000, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Cable UTP Cat5e x100m", "category_slug": "cable-utp-red", "brand": null, "price_reference": 15000, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Llave térmica 20A Schneider", "category_slug": "llaves-termicas-elec", "brand": "Schneider Electric", "price_reference": 5000, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Disyuntor diferencial 25A Siemens", "category_slug": "disyuntores", "brand": "Siemens", "price_reference": 18000, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Protector de tensión monofásico", "category_slug": "protectores-tension", "brand": "Siemens", "price_reference": 12000, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Tablero embutir 12 módulos", "category_slug": "tableros-embutir", "brand": "Schneider Electric", "price_reference": 8000, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Lámpara LED 9W E27 Philips", "category_slug": "lamparas-led", "brand": "Philips", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Tubo LED 18W 120cm", "category_slug": "tubos-led", "brand": "Ledvance", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Plafón LED 18W", "category_slug": "plafones-apliques", "brand": "Philips", "price_reference": 5000, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Proyector LED 50W", "category_slug": "proyectores-led", "brand": "Ledvance", "price_reference": 12000, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Módulo tomacorriente doble Jeluz", "category_slug": "linea-modular", "brand": "Jeluz", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Tapa interruptor Cambre", "category_slug": "tapas-bastidores", "brand": "Cambre", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Caja rectangular embutir", "category_slug": "cajas-embutir", "brand": null, "price_reference": 300, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Caño PVC 3/4 x3m", "category_slug": "canos-curvas", "brand": "Roker", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Canaleta 20x10 x2m", "category_slug": "canaletas", "brand": "Roker", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Precintos 200mm x100", "category_slug": "precintos-accesorios", "brand": null, "price_reference": 800, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Timer digital programable", "category_slug": "timers-programadores", "brand": "Lexel", "price_reference": 8000, "unit": "unidad", "sku_prefix": "ELEC"},
    {"name": "Sensor movimiento techo", "category_slug": "sensores-movimiento", "brand": "Lexel", "price_reference": 5000, "unit": "unidad", "sku_prefix": "ELEC"}
  ]'::jsonb,
  '[
    {"name": "Sección Cable", "slug": "seccion-cable", "values": ["1mm²","1.5mm²","2.5mm²","4mm²","6mm²","10mm²","16mm²"], "applies_to_categories": ["cable-unipolar","cable-subterraneo"]},
    {"name": "Amperaje", "slug": "amperaje", "values": ["10A","16A","20A","25A","32A","40A","63A"], "applies_to_categories": ["llaves-termicas-elec","disyuntores"]},
    {"name": "Potencia LED", "slug": "potencia-led", "values": ["5W","7W","9W","12W","15W","18W","20W","50W","100W"], "applies_to_categories": ["lamparas-led","tubos-led","plafones-apliques","proyectores-led"]},
    {"name": "Color Luz", "slug": "color-luz", "values": ["Cálida (3000K)","Neutra (4000K)","Fría (6500K)"], "applies_to_categories": ["lamparas-led","tubos-led","plafones-apliques","proyectores-led"]}
  ]'::jsonb,
  true, true, '4.0.0-enriched'
FROM business_types bt WHERE bt.code = 'electricidad'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories, brands = EXCLUDED.brands, products = EXCLUDED.products,
  attributes = EXCLUDED.attributes, version = EXCLUDED.version,
  updated_at = CURRENT_TIMESTAMP;

-- =====================================================
-- 10. SANITARIOS (NUEVO) — 6 padres + 22 hijas, 9 marcas, 20 productos
-- =====================================================
INSERT INTO business_type_templates (
  business_type_id, name, description, region, categories, brands, products, attributes,
  is_default, is_active, version
)
SELECT bt.id, 'Sanitarios y Griferías', 'Template curado para casas de sanitarios y griferías', 'AR',
  '[
    {"slug": "sanitarios-bano", "name": "Sanitarios", "level": 0},
    {"slug": "inodoros", "name": "Inodoros", "parent_slug": "sanitarios-bano", "level": 1},
    {"slug": "bidets", "name": "Bidets", "parent_slug": "sanitarios-bano", "level": 1},
    {"slug": "depositos-tapas", "name": "Depósitos y Tapas", "parent_slug": "sanitarios-bano", "level": 1},

    {"slug": "griferias", "name": "Griferías", "level": 0},
    {"slug": "griferia-cocina", "name": "Grifería Cocina", "parent_slug": "griferias", "level": 1},
    {"slug": "griferia-bano", "name": "Grifería Baño", "parent_slug": "griferias", "level": 1},
    {"slug": "griferia-ducha", "name": "Grifería Ducha", "parent_slug": "griferias", "level": 1},

    {"slug": "bachas-lavatorios", "name": "Bachas y Lavatorios", "level": 0},
    {"slug": "bachas-apoyar", "name": "Bachas de Apoyar", "parent_slug": "bachas-lavatorios", "level": 1},
    {"slug": "bachas-embutir", "name": "Bachas de Embutir", "parent_slug": "bachas-lavatorios", "level": 1},
    {"slug": "piletas-cocina", "name": "Piletas Cocina", "parent_slug": "bachas-lavatorios", "level": 1},

    {"slug": "duchas-baneras", "name": "Duchas y Bañeras", "level": 0},
    {"slug": "baneras", "name": "Bañeras", "parent_slug": "duchas-baneras", "level": 1},
    {"slug": "boxes-ducha", "name": "Boxes de Ducha", "parent_slug": "duchas-baneras", "level": 1},
    {"slug": "mamparas", "name": "Mamparas", "parent_slug": "duchas-baneras", "level": 1},
    {"slug": "receptaculos", "name": "Receptáculos", "parent_slug": "duchas-baneras", "level": 1},

    {"slug": "muebles-bano", "name": "Muebles de Baño", "level": 0},
    {"slug": "vanitorys", "name": "Vanitorys", "parent_slug": "muebles-bano", "level": 1},
    {"slug": "espejos-botiquines", "name": "Espejos y Botiquines", "parent_slug": "muebles-bano", "level": 1},

    {"slug": "accesorios-bano", "name": "Accesorios de Baño", "level": 0},
    {"slug": "toalleros", "name": "Toalleros", "parent_slug": "accesorios-bano", "level": 1},
    {"slug": "jaboneras-portarrollos", "name": "Jaboneras y Portarrollos", "parent_slug": "accesorios-bano", "level": 1},
    {"slug": "ganchos-percheros", "name": "Ganchos y Percheros", "parent_slug": "accesorios-bano", "level": 1},
    {"slug": "barras-seguridad", "name": "Barras de Seguridad", "parent_slug": "accesorios-bano", "level": 1}
  ]'::jsonb,
  '[
    {"name": "FV", "suggested_for_categories": ["griferia-cocina","griferia-bano","griferia-ducha"]},
    {"name": "Ferrum", "suggested_for_categories": ["inodoros","bidets","depositos-tapas","bachas-apoyar"]},
    {"name": "Roca", "suggested_for_categories": ["inodoros","bidets","baneras"]},
    {"name": "Piazza", "suggested_for_categories": ["griferia-cocina","griferia-bano"]},
    {"name": "Peirano", "suggested_for_categories": ["griferia-cocina","griferia-ducha"]},
    {"name": "Hidromet", "suggested_for_categories": ["boxes-ducha","mamparas"]},
    {"name": "American Standard", "suggested_for_categories": ["inodoros","bidets"]},
    {"name": "Johnson Acero", "suggested_for_categories": ["piletas-cocina"]},
    {"name": "Deca", "suggested_for_categories": ["griferia-bano","accesorios-bano"]}
  ]'::jsonb,
  '[
    {"name": "Inodoro largo Ferrum Bari", "category_slug": "inodoros", "brand": "Ferrum", "price_reference": 85000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Inodoro corto Roca", "category_slug": "inodoros", "brand": "Roca", "price_reference": 95000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Bidet Ferrum Bari", "category_slug": "bidets", "brand": "Ferrum", "price_reference": 65000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Depósito doble descarga", "category_slug": "depositos-tapas", "brand": "Ferrum", "price_reference": 35000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Tapa inodoro Ferrum", "category_slug": "depositos-tapas", "brand": "Ferrum", "price_reference": 12000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Grifería monocomando cocina FV", "category_slug": "griferia-cocina", "brand": "FV", "price_reference": 45000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Grifería lavatorio FV Libby", "category_slug": "griferia-bano", "brand": "FV", "price_reference": 35000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Grifería ducha monocomando FV", "category_slug": "griferia-ducha", "brand": "FV", "price_reference": 55000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Bacha de apoyar redonda", "category_slug": "bachas-apoyar", "brand": "Ferrum", "price_reference": 28000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Pileta cocina doble Johnson", "category_slug": "piletas-cocina", "brand": "Johnson Acero", "price_reference": 55000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Bañera 1.50m Ferrum", "category_slug": "baneras", "brand": "Ferrum", "price_reference": 120000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Box ducha 80x80 Hidromet", "category_slug": "boxes-ducha", "brand": "Hidromet", "price_reference": 180000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Mampara rebatible", "category_slug": "mamparas", "brand": "Hidromet", "price_reference": 95000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Vanitory 60cm con mesada", "category_slug": "vanitorys", "brand": null, "price_reference": 65000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Espejo biselado 50x70", "category_slug": "espejos-botiquines", "brand": null, "price_reference": 15000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Toallero barra 60cm", "category_slug": "toalleros", "brand": null, "price_reference": 8000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Jabonera pared", "category_slug": "jaboneras-portarrollos", "brand": null, "price_reference": 3000, "unit": "unidad", "sku_prefix": "SANIT"},
    {"name": "Portarrollo cerrado", "category_slug": "jaboneras-portarrollos", "brand": null, "price_reference": 4000, "unit": "unidad", "sku_prefix": "SANIT"}
  ]'::jsonb,
  '[
    {"name": "Material", "slug": "material-sanit", "values": ["Loza","Porcelana","Acero Inoxidable","Vidrio Templado","Acrílico"], "applies_to_categories": ["inodoros","bidets","bachas-apoyar","bachas-embutir","piletas-cocina","baneras"]},
    {"name": "Color", "slug": "color-sanit", "values": ["Blanco","Blanco Brillante","Bone","Gris"], "applies_to_categories": ["inodoros","bidets","depositos-tapas","bachas-apoyar"]},
    {"name": "Medida", "slug": "medida-sanit", "values": ["40cm","50cm","60cm","80cm","100cm","120cm","150cm","170cm"], "applies_to_categories": ["vanitorys","baneras","boxes-ducha","mamparas","piletas-cocina"]}
  ]'::jsonb,
  true, true, '4.0.0-enriched'
FROM business_types bt WHERE bt.code = 'sanitarios'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories, brands = EXCLUDED.brands, products = EXCLUDED.products,
  attributes = EXCLUDED.attributes, version = EXCLUDED.version,
  updated_at = CURRENT_TIMESTAMP;

-- =====================================================
-- 11. ZAPATERÍA (NUEVO template, business_type ya existe) — 5 padres + 20 hijas, 14 marcas, 22 productos
-- =====================================================
INSERT INTO business_type_templates (
  business_type_id, name, description, region, categories, brands, products, attributes,
  is_default, is_active, version
)
SELECT bt.id, 'Zapatería', 'Template curado para zapaterías argentinas', 'AR',
  '[
    {"slug": "calzado-hombre", "name": "Hombre", "level": 0},
    {"slug": "zapatos-vestir-h", "name": "Zapatos de Vestir", "parent_slug": "calzado-hombre", "level": 1},
    {"slug": "zapatillas-urbanas-h", "name": "Zapatillas Urbanas", "parent_slug": "calzado-hombre", "level": 1},
    {"slug": "botas-borcegos-h", "name": "Botas y Borcegos", "parent_slug": "calzado-hombre", "level": 1},
    {"slug": "mocasines", "name": "Mocasines", "parent_slug": "calzado-hombre", "level": 1},
    {"slug": "ojotas-sandalias-h", "name": "Ojotas y Sandalias", "parent_slug": "calzado-hombre", "level": 1},

    {"slug": "calzado-mujer", "name": "Mujer", "level": 0},
    {"slug": "zapatos-stilettos", "name": "Zapatos y Stilettos", "parent_slug": "calzado-mujer", "level": 1},
    {"slug": "zapatillas-urbanas-m", "name": "Zapatillas Urbanas", "parent_slug": "calzado-mujer", "level": 1},
    {"slug": "botas-botinetas", "name": "Botas y Botinetas", "parent_slug": "calzado-mujer", "level": 1},
    {"slug": "sandalias-zuecos", "name": "Sandalias y Zuecos", "parent_slug": "calzado-mujer", "level": 1},
    {"slug": "chatitas-ballerinas", "name": "Chatitas y Ballerinas", "parent_slug": "calzado-mujer", "level": 1},

    {"slug": "calzado-ninos", "name": "Niños", "level": 0},
    {"slug": "zapatillas-escolares", "name": "Zapatillas Escolares", "parent_slug": "calzado-ninos", "level": 1},
    {"slug": "sandalias-ninos", "name": "Sandalias", "parent_slug": "calzado-ninos", "level": 1},
    {"slug": "botas-ninos", "name": "Botas", "parent_slug": "calzado-ninos", "level": 1},

    {"slug": "calzado-deportivo", "name": "Deportivo", "level": 0},
    {"slug": "zapatillas-running-zap", "name": "Zapatillas Running", "parent_slug": "calzado-deportivo", "level": 1},
    {"slug": "zapatillas-training-zap", "name": "Zapatillas Training", "parent_slug": "calzado-deportivo", "level": 1},
    {"slug": "botines-futbol", "name": "Botines Fútbol", "parent_slug": "calzado-deportivo", "level": 1},
    {"slug": "zapatillas-trekking", "name": "Zapatillas Trekking", "parent_slug": "calzado-deportivo", "level": 1},

    {"slug": "accesorios-calzado", "name": "Accesorios", "level": 0},
    {"slug": "plantillas", "name": "Plantillas", "parent_slug": "accesorios-calzado", "level": 1},
    {"slug": "cordones", "name": "Cordones", "parent_slug": "accesorios-calzado", "level": 1},
    {"slug": "medias-deportivas", "name": "Medias Deportivas", "parent_slug": "accesorios-calzado", "level": 1},
    {"slug": "productos-cuidado", "name": "Productos de Cuidado", "parent_slug": "accesorios-calzado", "level": 1}
  ]'::jsonb,
  '[
    {"name": "Topper", "suggested_for_categories": ["zapatillas-urbanas-h","zapatillas-urbanas-m","zapatillas-escolares","zapatillas-training-zap"]},
    {"name": "Adidas", "suggested_for_categories": ["zapatillas-urbanas-h","zapatillas-urbanas-m","zapatillas-running-zap","botines-futbol"]},
    {"name": "Nike", "suggested_for_categories": ["zapatillas-running-zap","zapatillas-training-zap","botines-futbol"]},
    {"name": "Fila", "suggested_for_categories": ["zapatillas-urbanas-h","zapatillas-urbanas-m","zapatillas-running-zap"]},
    {"name": "Puma", "suggested_for_categories": ["zapatillas-urbanas-h","zapatillas-urbanas-m"]},
    {"name": "Jaguar", "suggested_for_categories": ["zapatillas-escolares","zapatillas-urbanas-h"]},
    {"name": "Kappa", "suggested_for_categories": ["zapatillas-urbanas-h","zapatillas-training-zap"]},
    {"name": "Vans", "suggested_for_categories": ["zapatillas-urbanas-h","zapatillas-urbanas-m"]},
    {"name": "Converse", "suggested_for_categories": ["zapatillas-urbanas-h","zapatillas-urbanas-m"]},
    {"name": "New Balance", "suggested_for_categories": ["zapatillas-running-zap"]},
    {"name": "Skechers", "suggested_for_categories": ["zapatillas-urbanas-m","zapatillas-training-zap"]},
    {"name": "Hush Puppies", "suggested_for_categories": ["zapatos-vestir-h","mocasines"]},
    {"name": "Grimoldi", "suggested_for_categories": ["zapatos-vestir-h","zapatos-stilettos"]},
    {"name": "Salomon", "suggested_for_categories": ["zapatillas-trekking"]}
  ]'::jsonb,
  '[
    {"name": "Zapato vestir cuero negro", "category_slug": "zapatos-vestir-h", "brand": "Hush Puppies", "price_reference": 65000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Zapatilla urbana Adidas", "category_slug": "zapatillas-urbanas-h", "brand": "Adidas", "price_reference": 55000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Zapatilla urbana Vans", "category_slug": "zapatillas-urbanas-h", "brand": "Vans", "price_reference": 45000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Borcego cuero hombre", "category_slug": "botas-borcegos-h", "brand": null, "price_reference": 55000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Mocasín cuero", "category_slug": "mocasines", "brand": "Hush Puppies", "price_reference": 48000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Ojota hombre", "category_slug": "ojotas-sandalias-h", "brand": "Adidas", "price_reference": 12000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Stiletto cuero negro", "category_slug": "zapatos-stilettos", "brand": "Grimoldi", "price_reference": 55000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Zapatilla urbana mujer Nike", "category_slug": "zapatillas-urbanas-m", "brand": "Nike", "price_reference": 50000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Bota caña alta mujer", "category_slug": "botas-botinetas", "brand": null, "price_reference": 65000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Sandalia plataforma", "category_slug": "sandalias-zuecos", "brand": null, "price_reference": 25000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Chatita cuero", "category_slug": "chatitas-ballerinas", "brand": null, "price_reference": 22000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Zapatilla escolar Topper", "category_slug": "zapatillas-escolares", "brand": "Topper", "price_reference": 25000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Zapatilla escolar Jaguar", "category_slug": "zapatillas-escolares", "brand": "Jaguar", "price_reference": 18000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Sandalia niño", "category_slug": "sandalias-ninos", "brand": null, "price_reference": 12000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Zapatilla running Nike", "category_slug": "zapatillas-running-zap", "brand": "Nike", "price_reference": 55000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Zapatilla running New Balance", "category_slug": "zapatillas-running-zap", "brand": "New Balance", "price_reference": 65000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Botines fútbol Adidas", "category_slug": "botines-futbol", "brand": "Adidas", "price_reference": 45000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Zapatilla trekking Salomon", "category_slug": "zapatillas-trekking", "brand": "Salomon", "price_reference": 85000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Plantillas confort", "category_slug": "plantillas", "brand": null, "price_reference": 5000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Cordones repuesto", "category_slug": "cordones", "brand": null, "price_reference": 800, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Medias deportivas x3", "category_slug": "medias-deportivas", "brand": null, "price_reference": 5000, "unit": "unidad", "sku_prefix": "ZAPAT"},
    {"name": "Crema para cuero", "category_slug": "productos-cuidado", "brand": null, "price_reference": 3000, "unit": "unidad", "sku_prefix": "ZAPAT"}
  ]'::jsonb,
  '[
    {"name": "Talle Calzado", "slug": "talle-calzado", "values": ["34","35","36","37","38","39","40","41","42","43","44","45"], "applies_to_categories": ["zapatos-vestir-h","zapatillas-urbanas-h","botas-borcegos-h","mocasines","zapatos-stilettos","zapatillas-urbanas-m","botas-botinetas","sandalias-zuecos","chatitas-ballerinas","zapatillas-escolares","zapatillas-running-zap","zapatillas-training-zap","botines-futbol","zapatillas-trekking"]},
    {"name": "Color", "slug": "color", "values": ["Negro","Blanco","Marrón","Azul","Gris","Rojo","Rosa","Beige","Bordo"], "applies_to_categories": ["zapatos-vestir-h","zapatillas-urbanas-h","zapatillas-urbanas-m","botas-botinetas","zapatillas-running-zap"]},
    {"name": "Material", "slug": "material-calzado", "values": ["Cuero","Cuero Sintético","Lona","Textil","Gamuza","Goma"], "applies_to_categories": ["zapatos-vestir-h","botas-borcegos-h","mocasines","zapatos-stilettos","botas-botinetas","chatitas-ballerinas"]}
  ]'::jsonb,
  true, true, '4.0.0-enriched'
FROM business_types bt WHERE bt.code = 'zapateria'
ON CONFLICT (business_type_id, region) WHERE is_default = true
DO UPDATE SET
  categories = EXCLUDED.categories, brands = EXCLUDED.brands, products = EXCLUDED.products,
  attributes = EXCLUDED.attributes, version = EXCLUDED.version,
  updated_at = CURRENT_TIMESTAMP;

-- =====================================================
-- PASO FINAL: Activar los nuevos rubros en quickstart
-- =====================================================
UPDATE business_type_templates
SET is_active = true, updated_at = CURRENT_TIMESTAMP
WHERE business_type_id IN (
  SELECT id FROM business_types
  WHERE code IN (
    'kiosco', 'almacen', 'ferreteria', 'pintureria',
    'ropa', 'deportes', 'libreria',
    'bazar', 'electricidad', 'sanitarios', 'zapateria'
  )
) AND is_default = true;
