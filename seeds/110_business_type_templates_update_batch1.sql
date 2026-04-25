-- =============================================================================
-- SEED 110: Actualización batch de business_type_templates — 8 rubros
-- Actualiza campos products, categories y brands via UPDATE por UUID
-- Fuente: catálogo real mercado argentino / NEA (Posadas, Misiones)
-- Fecha de research: 2026-04
-- VERSION: 1.0.0 | generated_by: catalog-researcher
-- IDEMPOTENTE: UPDATE por id primario, sin efectos secundarios
-- =============================================================================


-- ============================================================
-- 1. ALMACEN
-- UUID: b2000001-0000-4000-8000-000000000001
-- sku_prefix: ALMACEN
-- ============================================================

UPDATE business_type_templates
SET
  products = $products$[
    {"name": "Aceite girasol Natura 1.5L", "unit": "unidad", "brand": "Natura", "sku_prefix": "ALMACEN", "category_slug": "aceites-vinagres", "price_reference": 4800},
    {"name": "Aceite girasol Cocinero 900ml", "unit": "unidad", "brand": "Cocinero", "sku_prefix": "ALMACEN", "category_slug": "aceites-vinagres", "price_reference": 2900},
    {"name": "Aceite mezcla Marolio 1.5L", "unit": "unidad", "brand": "Marolio", "sku_prefix": "ALMACEN", "category_slug": "aceites-vinagres", "price_reference": 3800},
    {"name": "Vinagre de alcohol Menoyo 1L", "unit": "unidad", "brand": "Menoyo", "sku_prefix": "ALMACEN", "category_slug": "aceites-vinagres", "price_reference": 1300},
    {"name": "Azúcar refinada Ledesma 1kg", "unit": "unidad", "brand": "Ledesma", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 1800},
    {"name": "Sal fina Celusal 500g", "unit": "unidad", "brand": "Celusal", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 700},
    {"name": "Harina 0000 Blancaflor 1kg", "unit": "unidad", "brand": "Blancaflor", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 1400},
    {"name": "Harina 000 Cañuelas 1kg", "unit": "unidad", "brand": "Cañuelas", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 1200},
    {"name": "Arroz largo fino Molinos 1kg", "unit": "unidad", "brand": "Molinos", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 2200},
    {"name": "Yerba mate Rosamonte 1kg", "unit": "unidad", "brand": "Rosamonte", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 5200},
    {"name": "Yerba mate Playadito 1kg", "unit": "unidad", "brand": "Playadito", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 5500},
    {"name": "Yerba mate Taragüí 1kg", "unit": "unidad", "brand": "Taragüí", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 4900},
    {"name": "Polenta Marolio 500g", "unit": "unidad", "brand": "Marolio", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 1100},
    {"name": "Lentejas Marolio 500g", "unit": "unidad", "brand": "Marolio", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 1500},
    {"name": "Fideos moño Matarazzo 500g", "unit": "unidad", "brand": "Matarazzo", "sku_prefix": "ALMACEN", "category_slug": "pastas-harinas", "price_reference": 1800},
    {"name": "Fideos spaghetti Lucchetti 500g", "unit": "unidad", "brand": "Lucchetti", "sku_prefix": "ALMACEN", "category_slug": "pastas-harinas", "price_reference": 1600},
    {"name": "Fideos tirabuzón Don Vicente 500g", "unit": "unidad", "brand": "Don Vicente", "sku_prefix": "ALMACEN", "category_slug": "pastas-harinas", "price_reference": 1400},
    {"name": "Leche La Serenísima 1L", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "ALMACEN", "category_slug": "lacteos", "price_reference": 2000},
    {"name": "Leche SanCor entera 1L", "unit": "unidad", "brand": "SanCor", "sku_prefix": "ALMACEN", "category_slug": "lacteos", "price_reference": 1900},
    {"name": "Yogur La Serenísima entero 190g", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "ALMACEN", "category_slug": "lacteos", "price_reference": 1200},
    {"name": "Manteca La Serenísima 200g", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "ALMACEN", "category_slug": "lacteos", "price_reference": 3200},
    {"name": "Crema de leche La Serenísima 200ml", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "ALMACEN", "category_slug": "lacteos", "price_reference": 1700},
    {"name": "Dulce de leche La Serenísima 400g", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "ALMACEN", "category_slug": "lacteos", "price_reference": 2800},
    {"name": "Coca-Cola 2.25L", "unit": "unidad", "brand": "Coca-Cola", "sku_prefix": "ALMACEN", "category_slug": "bebidas", "price_reference": 3200},
    {"name": "Sprite 2.25L", "unit": "unidad", "brand": "Coca-Cola", "sku_prefix": "ALMACEN", "category_slug": "bebidas", "price_reference": 3000},
    {"name": "Fanta naranja 2.25L", "unit": "unidad", "brand": "Coca-Cola", "sku_prefix": "ALMACEN", "category_slug": "bebidas", "price_reference": 3000},
    {"name": "Agua mineral Villavicencio 2L", "unit": "unidad", "brand": "Villavicencio", "sku_prefix": "ALMACEN", "category_slug": "bebidas", "price_reference": 1400},
    {"name": "Agua mineral Glaciar 2L", "unit": "unidad", "brand": "Glaciar", "sku_prefix": "ALMACEN", "category_slug": "bebidas", "price_reference": 1200},
    {"name": "Jugo Tang naranja polvo 20g", "unit": "unidad", "brand": "Tang", "sku_prefix": "ALMACEN", "category_slug": "bebidas", "price_reference": 500},
    {"name": "Café instantáneo Nescafé 50g", "unit": "unidad", "brand": "Nescafé", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 3500},
    {"name": "Té negro Taragüí x25 saquitos", "unit": "unidad", "brand": "Taragüí", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 1500},
    {"name": "Tomate triturado Marolio 520g", "unit": "unidad", "brand": "Marolio", "sku_prefix": "ALMACEN", "category_slug": "conservas", "price_reference": 1400},
    {"name": "Puré de tomate Marolio 520g", "unit": "unidad", "brand": "Marolio", "sku_prefix": "ALMACEN", "category_slug": "conservas", "price_reference": 1200},
    {"name": "Atún al natural La Campagnola 170g", "unit": "unidad", "brand": "La Campagnola", "sku_prefix": "ALMACEN", "category_slug": "conservas", "price_reference": 2800},
    {"name": "Sardinas en aceite Marolio 125g", "unit": "unidad", "brand": "Marolio", "sku_prefix": "ALMACEN", "category_slug": "conservas", "price_reference": 1500},
    {"name": "Arvejas La Campagnola lata 300g", "unit": "unidad", "brand": "La Campagnola", "sku_prefix": "ALMACEN", "category_slug": "conservas", "price_reference": 1300},
    {"name": "Mermelada frutilla Arcor 390g", "unit": "unidad", "brand": "Arcor", "sku_prefix": "ALMACEN", "category_slug": "conservas", "price_reference": 2200},
    {"name": "Galletitas Oreo 118g", "unit": "unidad", "brand": "Oreo", "sku_prefix": "ALMACEN", "category_slug": "galletitas", "price_reference": 1700},
    {"name": "Galletitas Terrabusi Melitas 100g", "unit": "unidad", "brand": "Terrabusi", "sku_prefix": "ALMACEN", "category_slug": "galletitas", "price_reference": 1400},
    {"name": "Galletitas Bagley Criollitas 100g", "unit": "unidad", "brand": "Bagley", "sku_prefix": "ALMACEN", "category_slug": "galletitas", "price_reference": 1500},
    {"name": "Alfajor Jorgito chocolate 55g", "unit": "unidad", "brand": "Jorgito", "sku_prefix": "ALMACEN", "category_slug": "galletitas", "price_reference": 900},
    {"name": "Alfajor Guaymalén triple 75g", "unit": "unidad", "brand": "Guaymalén", "sku_prefix": "ALMACEN", "category_slug": "galletitas", "price_reference": 700},
    {"name": "Chocolate Milka Oreo 100g", "unit": "unidad", "brand": "Milka", "sku_prefix": "ALMACEN", "category_slug": "galletitas", "price_reference": 3500},
    {"name": "Bon o Bon x12 unidades", "unit": "unidad", "brand": "Arcor", "sku_prefix": "ALMACEN", "category_slug": "galletitas", "price_reference": 2800},
    {"name": "Jabón en polvo Ariel 800g", "unit": "unidad", "brand": "Ariel", "sku_prefix": "ALMACEN", "category_slug": "limpieza", "price_reference": 5500},
    {"name": "Jabón en polvo Ala 800g", "unit": "unidad", "brand": "Ala", "sku_prefix": "ALMACEN", "category_slug": "limpieza", "price_reference": 4500},
    {"name": "Detergente Magistral 750ml", "unit": "unidad", "brand": "Magistral", "sku_prefix": "ALMACEN", "category_slug": "limpieza", "price_reference": 2800},
    {"name": "Lavandina Ayudín 2L", "unit": "unidad", "brand": "Ayudín", "sku_prefix": "ALMACEN", "category_slug": "limpieza", "price_reference": 1400},
    {"name": "Limpiador Cif crema 750ml", "unit": "unidad", "brand": "Cif", "sku_prefix": "ALMACEN", "category_slug": "limpieza", "price_reference": 2500},
    {"name": "Suavizante Comfort 900ml", "unit": "unidad", "brand": "Comfort", "sku_prefix": "ALMACEN", "category_slug": "limpieza", "price_reference": 3200},
    {"name": "Papel higiénico Higienol x4", "unit": "unidad", "brand": "Higienol", "sku_prefix": "ALMACEN", "category_slug": "limpieza", "price_reference": 2800},
    {"name": "Pañales Pampers Confort Sec M x30", "unit": "unidad", "brand": "Pampers", "sku_prefix": "ALMACEN", "category_slug": "perfumeria-basica", "price_reference": 22000},
    {"name": "Shampoo Dove 400ml", "unit": "unidad", "brand": "Dove", "sku_prefix": "ALMACEN", "category_slug": "perfumeria-basica", "price_reference": 5500},
    {"name": "Jabón de tocador Dove 90g", "unit": "unidad", "brand": "Dove", "sku_prefix": "ALMACEN", "category_slug": "perfumeria-basica", "price_reference": 1800},
    {"name": "Pasta dental Colgate Triple 90g", "unit": "unidad", "brand": "Colgate", "sku_prefix": "ALMACEN", "category_slug": "perfumeria-basica", "price_reference": 2200},
    {"name": "Desodorante Rexona aerosol 150ml", "unit": "unidad", "brand": "Rexona", "sku_prefix": "ALMACEN", "category_slug": "perfumeria-basica", "price_reference": 4800},
    {"name": "Pan lactal Fargo grande", "unit": "unidad", "brand": "Fargo", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 2500},
    {"name": "Avena Quaker 500g", "unit": "unidad", "brand": "Quaker", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 2000},
    {"name": "Caldo de verdura Knorr x6", "unit": "unidad", "brand": "Knorr", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 1200},
    {"name": "Mayonesa Hellmann's 500g", "unit": "unidad", "brand": "Hellmann's", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 3800},
    {"name": "Ketchup Heinz 397g", "unit": "unidad", "brand": "Heinz", "sku_prefix": "ALMACEN", "category_slug": "almacen-seco", "price_reference": 3200},
    {"name": "Aceitunas verdes Marolio 350g", "unit": "unidad", "brand": "Marolio", "sku_prefix": "ALMACEN", "category_slug": "conservas", "price_reference": 2500}
  ]$products$::jsonb,
  categories = $categories$[
    {"name": "Aceites y Vinagres", "slug": "aceites-vinagres", "level": 0},
    {"name": "Almacén Seco", "slug": "almacen-seco", "level": 0},
    {"name": "Pastas y Harinas", "slug": "pastas-harinas", "level": 0},
    {"name": "Lácteos", "slug": "lacteos", "level": 0},
    {"name": "Bebidas", "slug": "bebidas", "level": 0},
    {"name": "Conservas", "slug": "conservas", "level": 0},
    {"name": "Galletitas y Golosinas", "slug": "galletitas", "level": 0},
    {"name": "Limpieza del Hogar", "slug": "limpieza", "level": 0},
    {"name": "Perfumería Básica", "slug": "perfumeria-basica", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "La Serenísima", "suggested_for_categories": ["lacteos"]},
    {"name": "SanCor", "suggested_for_categories": ["lacteos"]},
    {"name": "Marolio", "suggested_for_categories": ["aceites-vinagres", "conservas", "almacen-seco"]},
    {"name": "Arcor", "suggested_for_categories": ["conservas", "galletitas"]},
    {"name": "Bagley", "suggested_for_categories": ["galletitas"]},
    {"name": "Terrabusi", "suggested_for_categories": ["galletitas"]},
    {"name": "Coca-Cola", "suggested_for_categories": ["bebidas"]},
    {"name": "Natura", "suggested_for_categories": ["aceites-vinagres"]},
    {"name": "Molinos", "suggested_for_categories": ["almacen-seco", "pastas-harinas"]},
    {"name": "Rosamonte", "suggested_for_categories": ["almacen-seco"]},
    {"name": "Ala", "suggested_for_categories": ["limpieza"]},
    {"name": "Dove", "suggested_for_categories": ["perfumeria-basica"]}
  ]$brands$::jsonb,
  updated_at = NOW()
WHERE id = 'b2000001-0000-4000-8000-000000000001';


-- ============================================================
-- 2. FERRETERIA
-- UUID: b2000001-0000-4000-8000-000000000003
-- sku_prefix: FERRET
-- ============================================================

UPDATE business_type_templates
SET
  products = $products$[
    {"name": "Martillo carpintero Stanley 500g", "unit": "unidad", "brand": "Stanley", "sku_prefix": "FERRET", "category_slug": "herramientas-manuales", "price_reference": 8500},
    {"name": "Destornillador Phillips Stanley N°2", "unit": "unidad", "brand": "Stanley", "sku_prefix": "FERRET", "category_slug": "herramientas-manuales", "price_reference": 3200},
    {"name": "Destornillador plano Stanley N°4", "unit": "unidad", "brand": "Stanley", "sku_prefix": "FERRET", "category_slug": "herramientas-manuales", "price_reference": 3000},
    {"name": "Pinza universal Bahco 200mm", "unit": "unidad", "brand": "Bahco", "sku_prefix": "FERRET", "category_slug": "herramientas-manuales", "price_reference": 9500},
    {"name": "Alicate corte diagonal Bahco 160mm", "unit": "unidad", "brand": "Bahco", "sku_prefix": "FERRET", "category_slug": "herramientas-manuales", "price_reference": 8200},
    {"name": "Llave inglesa 10 Bahco", "unit": "unidad", "brand": "Bahco", "sku_prefix": "FERRET", "category_slug": "herramientas-manuales", "price_reference": 12000},
    {"name": "Sierra de mano Tramontina 24T", "unit": "unidad", "brand": "Tramontina", "sku_prefix": "FERRET", "category_slug": "herramientas-manuales", "price_reference": 7500},
    {"name": "Nivel de burbuja 60cm Lusqtoff", "unit": "unidad", "brand": "Lusqtoff", "sku_prefix": "FERRET", "category_slug": "herramientas-manuales", "price_reference": 5500},
    {"name": "Cinta métrica Stanley 5m", "unit": "unidad", "brand": "Stanley", "sku_prefix": "FERRET", "category_slug": "herramientas-manuales", "price_reference": 4500},
    {"name": "Taladro percutor Bosch GSB 650W", "unit": "unidad", "brand": "Bosch", "sku_prefix": "FERRET", "category_slug": "herramientas-electricas", "price_reference": 65000},
    {"name": "Amoladora angular Bosch GWS 115mm 720W", "unit": "unidad", "brand": "Bosch", "sku_prefix": "FERRET", "category_slug": "herramientas-electricas", "price_reference": 52000},
    {"name": "Sierra circular Makita 5007MG 185mm", "unit": "unidad", "brand": "Makita", "sku_prefix": "FERRET", "category_slug": "herramientas-electricas", "price_reference": 95000},
    {"name": "Caladora Bosch PST 650E", "unit": "unidad", "brand": "Bosch", "sku_prefix": "FERRET", "category_slug": "herramientas-electricas", "price_reference": 48000},
    {"name": "Lijadora orbital Black+Decker KA198", "unit": "unidad", "brand": "Black+Decker", "sku_prefix": "FERRET", "category_slug": "herramientas-electricas", "price_reference": 35000},
    {"name": "Tubo PVC cloacal 110mm x 3m Tigre", "unit": "unidad", "brand": "Tigre", "sku_prefix": "FERRET", "category_slug": "plomeria", "price_reference": 8500},
    {"name": "Tubo PVC presión 32mm x 3m Tigre", "unit": "unidad", "brand": "Tigre", "sku_prefix": "FERRET", "category_slug": "plomeria", "price_reference": 4200},
    {"name": "Codo 90° PVC 110mm Tigre", "unit": "unidad", "brand": "Tigre", "sku_prefix": "FERRET", "category_slug": "plomeria", "price_reference": 1800},
    {"name": "Llave de paso esfera 1/2\" FV", "unit": "unidad", "brand": "FV", "sku_prefix": "FERRET", "category_slug": "plomeria", "price_reference": 4500},
    {"name": "Llave de paso esfera 3/4\" FV", "unit": "unidad", "brand": "FV", "sku_prefix": "FERRET", "category_slug": "plomeria", "price_reference": 5800},
    {"name": "Flotante de bronce 1/2\" general", "unit": "unidad", "brand": null, "sku_prefix": "FERRET", "category_slug": "plomeria", "price_reference": 3200},
    {"name": "Flexible reforzado 1/2\" x 40cm", "unit": "unidad", "brand": null, "sku_prefix": "FERRET", "category_slug": "plomeria", "price_reference": 2500},
    {"name": "Cable unipolar 2.5mm² Prysmian x25m", "unit": "unidad", "brand": "Prysmian", "sku_prefix": "FERRET", "category_slug": "electrica", "price_reference": 18000},
    {"name": "Cable unipolar 1.5mm² Prysmian x25m", "unit": "unidad", "brand": "Prysmian", "sku_prefix": "FERRET", "category_slug": "electrica", "price_reference": 12000},
    {"name": "Disyuntor 16A Bticino", "unit": "unidad", "brand": "Bticino", "sku_prefix": "FERRET", "category_slug": "electrica", "price_reference": 6500},
    {"name": "Disyuntor 20A Bticino", "unit": "unidad", "brand": "Bticino", "sku_prefix": "FERRET", "category_slug": "electrica", "price_reference": 7200},
    {"name": "Tomacorriente doble Cambre", "unit": "unidad", "brand": "Cambre", "sku_prefix": "FERRET", "category_slug": "electrica", "price_reference": 3800},
    {"name": "Interruptor simple Cambre", "unit": "unidad", "brand": "Cambre", "sku_prefix": "FERRET", "category_slug": "electrica", "price_reference": 2800},
    {"name": "Cinta aisladora 3M 19mm x 20m", "unit": "unidad", "brand": "3M", "sku_prefix": "FERRET", "category_slug": "electrica", "price_reference": 1800},
    {"name": "Pintura látex interior Sinteplast 4L", "unit": "unidad", "brand": "Sinteplast", "sku_prefix": "FERRET", "category_slug": "pintura", "price_reference": 22000},
    {"name": "Pintura látex exterior Alba 4L", "unit": "unidad", "brand": "Alba", "sku_prefix": "FERRET", "category_slug": "pintura", "price_reference": 25000},
    {"name": "Esmalte sintético Alba blanco 1L", "unit": "unidad", "brand": "Alba", "sku_prefix": "FERRET", "category_slug": "pintura", "price_reference": 9500},
    {"name": "Esmalte sintético Tersuave negro 1L", "unit": "unidad", "brand": "Tersuave", "sku_prefix": "FERRET", "category_slug": "pintura", "price_reference": 8800},
    {"name": "Rodillo lana 23cm con palo", "unit": "unidad", "brand": null, "sku_prefix": "FERRET", "category_slug": "pintura", "price_reference": 4500},
    {"name": "Pincel plano N°20 Hamilton", "unit": "unidad", "brand": "Hamilton", "sku_prefix": "FERRET", "category_slug": "pintura", "price_reference": 2800},
    {"name": "Tornillos autorroscantes x50 6x1\"", "unit": "caja", "brand": null, "sku_prefix": "FERRET", "category_slug": "fijaciones", "price_reference": 1500},
    {"name": "Tornillos para madera x50 4x40mm", "unit": "caja", "brand": null, "sku_prefix": "FERRET", "category_slug": "fijaciones", "price_reference": 1800},
    {"name": "Tarugo plástico 8mm x50", "unit": "caja", "brand": null, "sku_prefix": "FERRET", "category_slug": "fijaciones", "price_reference": 900},
    {"name": "Tarugo plástico 10mm x50", "unit": "caja", "brand": null, "sku_prefix": "FERRET", "category_slug": "fijaciones", "price_reference": 1100},
    {"name": "Poxipol transparente 14ml", "unit": "unidad", "brand": "Poxipol", "sku_prefix": "FERRET", "category_slug": "fijaciones", "price_reference": 3200},
    {"name": "La Gotita gel 4g", "unit": "unidad", "brand": "La Gotita", "sku_prefix": "FERRET", "category_slug": "fijaciones", "price_reference": 1500},
    {"name": "Candado Yale 40mm", "unit": "unidad", "brand": "Yale", "sku_prefix": "FERRET", "category_slug": "cerrajeria", "price_reference": 8500},
    {"name": "Candado Yale 50mm", "unit": "unidad", "brand": "Yale", "sku_prefix": "FERRET", "category_slug": "cerrajeria", "price_reference": 11000},
    {"name": "Bisagra de puerta 3\" x 3\" acero", "unit": "unidad", "brand": null, "sku_prefix": "FERRET", "category_slug": "cerrajeria", "price_reference": 2200},
    {"name": "Disco de corte Bosch 115mm metal", "unit": "unidad", "brand": "Bosch", "sku_prefix": "FERRET", "category_slug": "herramientas-electricas", "price_reference": 1800},
    {"name": "Disco de desbaste Bosch 115mm", "unit": "unidad", "brand": "Bosch", "sku_prefix": "FERRET", "category_slug": "herramientas-electricas", "price_reference": 2500},
    {"name": "Broca para hormigón 8mm x120mm", "unit": "unidad", "brand": null, "sku_prefix": "FERRET", "category_slug": "herramientas-electricas", "price_reference": 1200},
    {"name": "Broca para hormigón 10mm x150mm", "unit": "unidad", "brand": null, "sku_prefix": "FERRET", "category_slug": "herramientas-electricas", "price_reference": 1800},
    {"name": "Masilla plástica Sinteplast 1kg", "unit": "unidad", "brand": "Sinteplast", "sku_prefix": "FERRET", "category_slug": "pintura", "price_reference": 4500},
    {"name": "Cinta de embalar transparente 48mm", "unit": "unidad", "brand": null, "sku_prefix": "FERRET", "category_slug": "fijaciones", "price_reference": 1200},
    {"name": "Guantes de látex talle L x par", "unit": "par", "brand": null, "sku_prefix": "FERRET", "category_slug": "herramientas-manuales", "price_reference": 1500},
    {"name": "Anteojos de seguridad transparentes", "unit": "unidad", "brand": null, "sku_prefix": "FERRET", "category_slug": "herramientas-manuales", "price_reference": 2800}
  ]$products$::jsonb,
  categories = $categories$[
    {"name": "Herramientas Manuales", "slug": "herramientas-manuales", "level": 0},
    {"name": "Herramientas Eléctricas", "slug": "herramientas-electricas", "level": 0},
    {"name": "Plomería", "slug": "plomeria", "level": 0},
    {"name": "Electricidad", "slug": "electrica", "level": 0},
    {"name": "Pintura", "slug": "pintura", "level": 0},
    {"name": "Fijaciones y Adhesivos", "slug": "fijaciones", "level": 0},
    {"name": "Cerrajería", "slug": "cerrajeria", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "Stanley", "suggested_for_categories": ["herramientas-manuales"]},
    {"name": "Bahco", "suggested_for_categories": ["herramientas-manuales"]},
    {"name": "Bosch", "suggested_for_categories": ["herramientas-electricas"]},
    {"name": "Makita", "suggested_for_categories": ["herramientas-electricas"]},
    {"name": "Black+Decker", "suggested_for_categories": ["herramientas-electricas"]},
    {"name": "Tigre", "suggested_for_categories": ["plomeria"]},
    {"name": "FV", "suggested_for_categories": ["plomeria"]},
    {"name": "Bticino", "suggested_for_categories": ["electrica"]},
    {"name": "Cambre", "suggested_for_categories": ["electrica"]},
    {"name": "Alba", "suggested_for_categories": ["pintura"]},
    {"name": "Sinteplast", "suggested_for_categories": ["pintura"]},
    {"name": "Yale", "suggested_for_categories": ["cerrajeria"]}
  ]$brands$::jsonb,
  updated_at = NOW()
WHERE id = 'b2000001-0000-4000-8000-000000000003';


-- ============================================================
-- 3. VINOTECA
-- UUID: 3c546ca9-0be2-4a30-97bb-f917dd6b251f
-- sku_prefix: VINO
-- ============================================================

UPDATE business_type_templates
SET
  products = $products$[
    {"name": "Trapiche Malbec Roble 750ml", "unit": "unidad", "brand": "Trapiche", "sku_prefix": "VINO", "category_slug": "vinos-tintos", "price_reference": 4500},
    {"name": "Norton Malbec 750ml", "unit": "unidad", "brand": "Norton", "sku_prefix": "VINO", "category_slug": "vinos-tintos", "price_reference": 5800},
    {"name": "Zuccardi Valle A Malbec 750ml", "unit": "unidad", "brand": "Zuccardi", "sku_prefix": "VINO", "category_slug": "vinos-tintos", "price_reference": 9500},
    {"name": "Santa Julia Malbec 750ml", "unit": "unidad", "brand": "Santa Julia", "sku_prefix": "VINO", "category_slug": "vinos-tintos", "price_reference": 4200},
    {"name": "Gato Negro Malbec 750ml", "unit": "unidad", "brand": "Gato Negro", "sku_prefix": "VINO", "category_slug": "vinos-tintos", "price_reference": 3800},
    {"name": "Trivento Reserve Malbec 750ml", "unit": "unidad", "brand": "Trivento", "sku_prefix": "VINO", "category_slug": "vinos-tintos", "price_reference": 5200},
    {"name": "Clos de los Siete Malbec 750ml", "unit": "unidad", "brand": "Clos de los Siete", "sku_prefix": "VINO", "category_slug": "vinos-tintos", "price_reference": 12000},
    {"name": "Achaval Ferrer Malbec 750ml", "unit": "unidad", "brand": "Achaval Ferrer", "sku_prefix": "VINO", "category_slug": "vinos-tintos", "price_reference": 14000},
    {"name": "Catena Zapata Adrianna Malbec 750ml", "unit": "unidad", "brand": "Catena Zapata", "sku_prefix": "VINO", "category_slug": "vinos-tintos", "price_reference": 28000},
    {"name": "Navarro Correas Cabernet Sauvignon 750ml", "unit": "unidad", "brand": "Navarro Correas", "sku_prefix": "VINO", "category_slug": "vinos-tintos", "price_reference": 6500},
    {"name": "Trapiche Torrontés 750ml", "unit": "unidad", "brand": "Trapiche", "sku_prefix": "VINO", "category_slug": "vinos-blancos", "price_reference": 4200},
    {"name": "Santa Julia Chardonnay 750ml", "unit": "unidad", "brand": "Santa Julia", "sku_prefix": "VINO", "category_slug": "vinos-blancos", "price_reference": 4000},
    {"name": "Zuccardi Torrontés 750ml", "unit": "unidad", "brand": "Zuccardi", "sku_prefix": "VINO", "category_slug": "vinos-blancos", "price_reference": 8500},
    {"name": "Norton Sauvignon Blanc 750ml", "unit": "unidad", "brand": "Norton", "sku_prefix": "VINO", "category_slug": "vinos-blancos", "price_reference": 5200},
    {"name": "Gato Blanco Chenin 750ml", "unit": "unidad", "brand": "Gato Negro", "sku_prefix": "VINO", "category_slug": "vinos-blancos", "price_reference": 3500},
    {"name": "Chandon Extra Brut 750ml", "unit": "unidad", "brand": "Chandon", "sku_prefix": "VINO", "category_slug": "espumantes", "price_reference": 12500},
    {"name": "Chandon Rosé 750ml", "unit": "unidad", "brand": "Chandon", "sku_prefix": "VINO", "category_slug": "espumantes", "price_reference": 13000},
    {"name": "Mumm Blanc de Blancs 750ml", "unit": "unidad", "brand": "Mumm", "sku_prefix": "VINO", "category_slug": "espumantes", "price_reference": 11000},
    {"name": "Valentin Bianchi Extra Brut 750ml", "unit": "unidad", "brand": "Valentín Bianchi", "sku_prefix": "VINO", "category_slug": "espumantes", "price_reference": 9500},
    {"name": "Dante Robino Brut Nature 750ml", "unit": "unidad", "brand": "Dante Robino", "sku_prefix": "VINO", "category_slug": "espumantes", "price_reference": 8800},
    {"name": "Johnnie Walker Red Label 750ml", "unit": "unidad", "brand": "Johnnie Walker", "sku_prefix": "VINO", "category_slug": "spirits", "price_reference": 28000},
    {"name": "Jameson Whiskey 750ml", "unit": "unidad", "brand": "Jameson", "sku_prefix": "VINO", "category_slug": "spirits", "price_reference": 32000},
    {"name": "Jack Daniel's Old No.7 750ml", "unit": "unidad", "brand": "Jack Daniel's", "sku_prefix": "VINO", "category_slug": "spirits", "price_reference": 35000},
    {"name": "Fernet Branca 750ml", "unit": "unidad", "brand": "Branca", "sku_prefix": "VINO", "category_slug": "spirits", "price_reference": 18000},
    {"name": "Aperol 700ml", "unit": "unidad", "brand": "Aperol", "sku_prefix": "VINO", "category_slug": "spirits", "price_reference": 22000},
    {"name": "Campari 700ml", "unit": "unidad", "brand": "Campari", "sku_prefix": "VINO", "category_slug": "spirits", "price_reference": 20000},
    {"name": "Gin Beefeater 750ml", "unit": "unidad", "brand": "Beefeater", "sku_prefix": "VINO", "category_slug": "spirits", "price_reference": 28000},
    {"name": "Ron Bacardi Carta Blanca 750ml", "unit": "unidad", "brand": "Bacardi", "sku_prefix": "VINO", "category_slug": "spirits", "price_reference": 24000},
    {"name": "Vodka Smirnoff 750ml", "unit": "unidad", "brand": "Smirnoff", "sku_prefix": "VINO", "category_slug": "spirits", "price_reference": 22000},
    {"name": "Heineken 330ml botella", "unit": "unidad", "brand": "Heineken", "sku_prefix": "VINO", "category_slug": "cervezas", "price_reference": 2800},
    {"name": "Stella Artois 473ml lata", "unit": "unidad", "brand": "Stella Artois", "sku_prefix": "VINO", "category_slug": "cervezas", "price_reference": 2500},
    {"name": "Patagonia Amber Lager 730ml", "unit": "unidad", "brand": "Patagonia", "sku_prefix": "VINO", "category_slug": "cervezas", "price_reference": 5500},
    {"name": "Patagonia Weisse 730ml", "unit": "unidad", "brand": "Patagonia", "sku_prefix": "VINO", "category_slug": "cervezas", "price_reference": 5800},
    {"name": "Corona 330ml", "unit": "unidad", "brand": "Corona", "sku_prefix": "VINO", "category_slug": "cervezas", "price_reference": 3200},
    {"name": "Quilmes Cristal 1L retornable", "unit": "unidad", "brand": "Quilmes", "sku_prefix": "VINO", "category_slug": "cervezas", "price_reference": 3000},
    {"name": "Agua con gas Villavicencio 500ml", "unit": "unidad", "brand": "Villavicencio", "sku_prefix": "VINO", "category_slug": "sin-alcohol", "price_reference": 1500},
    {"name": "Schweppes Tónica 250ml lata", "unit": "unidad", "brand": "Schweppes", "sku_prefix": "VINO", "category_slug": "sin-alcohol", "price_reference": 1800},
    {"name": "Ginger Beer Fever-Tree 200ml", "unit": "unidad", "brand": "Fever-Tree", "sku_prefix": "VINO", "category_slug": "sin-alcohol", "price_reference": 2800}
  ]$products$::jsonb,
  categories = $categories$[
    {"name": "Vinos Tintos", "slug": "vinos-tintos", "level": 0},
    {"name": "Vinos Blancos", "slug": "vinos-blancos", "level": 0},
    {"name": "Espumantes", "slug": "espumantes", "level": 0},
    {"name": "Spirits y Licores", "slug": "spirits", "level": 0},
    {"name": "Cervezas", "slug": "cervezas", "level": 0},
    {"name": "Sin Alcohol", "slug": "sin-alcohol", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "Trapiche", "suggested_for_categories": ["vinos-tintos", "vinos-blancos"]},
    {"name": "Norton", "suggested_for_categories": ["vinos-tintos", "vinos-blancos"]},
    {"name": "Zuccardi", "suggested_for_categories": ["vinos-tintos", "vinos-blancos"]},
    {"name": "Santa Julia", "suggested_for_categories": ["vinos-tintos", "vinos-blancos"]},
    {"name": "Chandon", "suggested_for_categories": ["espumantes"]},
    {"name": "Mumm", "suggested_for_categories": ["espumantes"]},
    {"name": "Johnnie Walker", "suggested_for_categories": ["spirits"]},
    {"name": "Branca", "suggested_for_categories": ["spirits"]},
    {"name": "Campari", "suggested_for_categories": ["spirits"]},
    {"name": "Heineken", "suggested_for_categories": ["cervezas"]},
    {"name": "Patagonia", "suggested_for_categories": ["cervezas"]},
    {"name": "Quilmes", "suggested_for_categories": ["cervezas"]}
  ]$brands$::jsonb,
  updated_at = NOW()
WHERE id = '3c546ca9-0be2-4a30-97bb-f917dd6b251f';


-- ============================================================
-- 4. KIOSCO
-- UUID: b2000001-0000-4000-8000-000000000002
-- sku_prefix: KIOSCO
-- ============================================================

UPDATE business_type_templates
SET
  products = $products$[
    {"name": "Coca-Cola 500ml", "unit": "unidad", "brand": "Coca-Cola", "sku_prefix": "KIOSCO", "category_slug": "bebidas-individuales", "price_reference": 1500},
    {"name": "Sprite 500ml", "unit": "unidad", "brand": "Coca-Cola", "sku_prefix": "KIOSCO", "category_slug": "bebidas-individuales", "price_reference": 1500},
    {"name": "Fanta naranja 500ml", "unit": "unidad", "brand": "Coca-Cola", "sku_prefix": "KIOSCO", "category_slug": "bebidas-individuales", "price_reference": 1500},
    {"name": "Agua Villavicencio 500ml", "unit": "unidad", "brand": "Villavicencio", "sku_prefix": "KIOSCO", "category_slug": "bebidas-individuales", "price_reference": 1100},
    {"name": "Agua Glaciar 500ml", "unit": "unidad", "brand": "Glaciar", "sku_prefix": "KIOSCO", "category_slug": "bebidas-individuales", "price_reference": 1000},
    {"name": "Monster Energy 473ml lata", "unit": "unidad", "brand": "Monster", "sku_prefix": "KIOSCO", "category_slug": "bebidas-individuales", "price_reference": 3800},
    {"name": "Red Bull 250ml lata", "unit": "unidad", "brand": "Red Bull", "sku_prefix": "KIOSCO", "category_slug": "bebidas-individuales", "price_reference": 3500},
    {"name": "Powerade 500ml naranja", "unit": "unidad", "brand": "Powerade", "sku_prefix": "KIOSCO", "category_slug": "bebidas-individuales", "price_reference": 2000},
    {"name": "Pepsi 500ml", "unit": "unidad", "brand": "Pepsi", "sku_prefix": "KIOSCO", "category_slug": "bebidas-individuales", "price_reference": 1400},
    {"name": "Alfajor Jorgito chocolate 55g", "unit": "unidad", "brand": "Jorgito", "sku_prefix": "KIOSCO", "category_slug": "chocolates", "price_reference": 900},
    {"name": "Alfajor Milka chocolate 52g", "unit": "unidad", "brand": "Milka", "sku_prefix": "KIOSCO", "category_slug": "chocolates", "price_reference": 1500},
    {"name": "Alfajor Guaymalén triple chocolate 75g", "unit": "unidad", "brand": "Guaymalén", "sku_prefix": "KIOSCO", "category_slug": "chocolates", "price_reference": 700},
    {"name": "Alfajor Terrabusi triple 70g", "unit": "unidad", "brand": "Terrabusi", "sku_prefix": "KIOSCO", "category_slug": "chocolates", "price_reference": 1200},
    {"name": "Alfajor Havanna chocolate 65g", "unit": "unidad", "brand": "Havanna", "sku_prefix": "KIOSCO", "category_slug": "chocolates", "price_reference": 2800},
    {"name": "Bon o Bon x12 unidades", "unit": "unidad", "brand": "Arcor", "sku_prefix": "KIOSCO", "category_slug": "chocolates", "price_reference": 2800},
    {"name": "Chocolatín Jack 30g", "unit": "unidad", "brand": "Arcor", "sku_prefix": "KIOSCO", "category_slug": "chocolates", "price_reference": 600},
    {"name": "Oreo individual 36g", "unit": "unidad", "brand": "Oreo", "sku_prefix": "KIOSCO", "category_slug": "galletitas", "price_reference": 800},
    {"name": "Bagley Merengadas 170g", "unit": "unidad", "brand": "Bagley", "sku_prefix": "KIOSCO", "category_slug": "galletitas", "price_reference": 1800},
    {"name": "Criollitas 100g", "unit": "unidad", "brand": "Bagley", "sku_prefix": "KIOSCO", "category_slug": "galletitas", "price_reference": 1500},
    {"name": "Papas fritas Pringles Original 124g", "unit": "unidad", "brand": "Pringles", "sku_prefix": "KIOSCO", "category_slug": "galletitas", "price_reference": 3500},
    {"name": "Papas fritas Lays clásicas 170g", "unit": "unidad", "brand": "Lays", "sku_prefix": "KIOSCO", "category_slug": "galletitas", "price_reference": 2800},
    {"name": "Pochoclo dulce 100g", "unit": "unidad", "brand": null, "sku_prefix": "KIOSCO", "category_slug": "galletitas", "price_reference": 1200},
    {"name": "Mentitas Arcor caja x28g", "unit": "unidad", "brand": "Arcor", "sku_prefix": "KIOSCO", "category_slug": "golosinas", "price_reference": 600},
    {"name": "Halls menta x12", "unit": "unidad", "brand": "Halls", "sku_prefix": "KIOSCO", "category_slug": "golosinas", "price_reference": 700},
    {"name": "Sugus frutilla x30g", "unit": "unidad", "brand": "Sugus", "sku_prefix": "KIOSCO", "category_slug": "golosinas", "price_reference": 800},
    {"name": "Beldent menta sin azúcar x12", "unit": "unidad", "brand": "Beldent", "sku_prefix": "KIOSCO", "category_slug": "golosinas", "price_reference": 800},
    {"name": "Bubbaloo fresa x5", "unit": "unidad", "brand": "Bubbaloo", "sku_prefix": "KIOSCO", "category_slug": "golosinas", "price_reference": 500},
    {"name": "Chicle Adams cool mint x12", "unit": "unidad", "brand": "Adams", "sku_prefix": "KIOSCO", "category_slug": "golosinas", "price_reference": 700},
    {"name": "Haribo Ositos 100g", "unit": "unidad", "brand": "Haribo", "sku_prefix": "KIOSCO", "category_slug": "golosinas", "price_reference": 2500},
    {"name": "Maní con chocolate Arcor x30g", "unit": "unidad", "brand": "Arcor", "sku_prefix": "KIOSCO", "category_slug": "golosinas", "price_reference": 600},
    {"name": "Pirulo frutilla helado", "unit": "unidad", "brand": "Arcor", "sku_prefix": "KIOSCO", "category_slug": "golosinas", "price_reference": 800},
    {"name": "Marlboro rojo x20", "unit": "unidad", "brand": "Marlboro", "sku_prefix": "KIOSCO", "category_slug": "cigarrillos", "price_reference": 6500},
    {"name": "Camel azul x20", "unit": "unidad", "brand": "Camel", "sku_prefix": "KIOSCO", "category_slug": "cigarrillos", "price_reference": 6200},
    {"name": "Lucky Strike x20", "unit": "unidad", "brand": "Lucky Strike", "sku_prefix": "KIOSCO", "category_slug": "cigarrillos", "price_reference": 6000},
    {"name": "Derby suave x20", "unit": "unidad", "brand": "Derby", "sku_prefix": "KIOSCO", "category_slug": "cigarrillos", "price_reference": 5500},
    {"name": "Encendedor BIC mini", "unit": "unidad", "brand": "BIC", "sku_prefix": "KIOSCO", "category_slug": "varios", "price_reference": 1500},
    {"name": "Pila AA Duracell x2", "unit": "unidad", "brand": "Duracell", "sku_prefix": "KIOSCO", "category_slug": "varios", "price_reference": 2800},
    {"name": "Preservativo Tulipán x3", "unit": "unidad", "brand": "Tulipán", "sku_prefix": "KIOSCO", "category_slug": "varios", "price_reference": 2200},
    {"name": "Chicle Trident menta x10", "unit": "unidad", "brand": "Trident", "sku_prefix": "KIOSCO", "category_slug": "golosinas", "price_reference": 900},
    {"name": "Chocolatín Milka 100g", "unit": "unidad", "brand": "Milka", "sku_prefix": "KIOSCO", "category_slug": "chocolates", "price_reference": 3500},
    {"name": "Coca-Cola 1.5L", "unit": "unidad", "brand": "Coca-Cola", "sku_prefix": "KIOSCO", "category_slug": "bebidas-individuales", "price_reference": 2500},
    {"name": "Manicero Arcor 30g", "unit": "unidad", "brand": "Arcor", "sku_prefix": "KIOSCO", "category_slug": "golosinas", "price_reference": 500},
    {"name": "Cepita naranja 200ml", "unit": "unidad", "brand": "Cepita", "sku_prefix": "KIOSCO", "category_slug": "bebidas-individuales", "price_reference": 1200},
    {"name": "Terma clásico 500ml", "unit": "unidad", "brand": "Terma", "sku_prefix": "KIOSCO", "category_slug": "bebidas-individuales", "price_reference": 1800},
    {"name": "Tofi caramelo x30g", "unit": "unidad", "brand": "Arcor", "sku_prefix": "KIOSCO", "category_slug": "golosinas", "price_reference": 500}
  ]$products$::jsonb,
  categories = $categories$[
    {"name": "Bebidas Individuales", "slug": "bebidas-individuales", "level": 0},
    {"name": "Chocolates y Alfajores", "slug": "chocolates", "level": 0},
    {"name": "Galletitas y Snacks", "slug": "galletitas", "level": 0},
    {"name": "Golosinas", "slug": "golosinas", "level": 0},
    {"name": "Cigarrillos", "slug": "cigarrillos", "level": 0},
    {"name": "Varios", "slug": "varios", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "Coca-Cola", "suggested_for_categories": ["bebidas-individuales"]},
    {"name": "Villavicencio", "suggested_for_categories": ["bebidas-individuales"]},
    {"name": "Monster", "suggested_for_categories": ["bebidas-individuales"]},
    {"name": "Arcor", "suggested_for_categories": ["golosinas", "chocolates"]},
    {"name": "Jorgito", "suggested_for_categories": ["chocolates"]},
    {"name": "Milka", "suggested_for_categories": ["chocolates"]},
    {"name": "Havanna", "suggested_for_categories": ["chocolates"]},
    {"name": "Bagley", "suggested_for_categories": ["galletitas"]},
    {"name": "Halls", "suggested_for_categories": ["golosinas"]},
    {"name": "Marlboro", "suggested_for_categories": ["cigarrillos"]},
    {"name": "BIC", "suggested_for_categories": ["varios"]},
    {"name": "Pringles", "suggested_for_categories": ["galletitas"]}
  ]$brands$::jsonb,
  updated_at = NOW()
WHERE id = 'b2000001-0000-4000-8000-000000000002';


-- ============================================================
-- 5. PERFUMERIA
-- UUID: 5301a74e-4cde-449a-90bf-96b87a8a397a
-- sku_prefix: PERF
-- ============================================================

UPDATE business_type_templates
SET
  products = $products$[
    {"name": "Shampoo Sedal rizos definidos 650ml", "unit": "unidad", "brand": "Sedal", "sku_prefix": "PERF", "category_slug": "cuidado-capilar", "price_reference": 4800},
    {"name": "Shampoo TRESemmé hidratación 750ml", "unit": "unidad", "brand": "TRESemmé", "sku_prefix": "PERF", "category_slug": "cuidado-capilar", "price_reference": 5500},
    {"name": "Shampoo Dove humectación intensa 400ml", "unit": "unidad", "brand": "Dove", "sku_prefix": "PERF", "category_slug": "cuidado-capilar", "price_reference": 5200},
    {"name": "Shampoo Pantene Pro-V 400ml", "unit": "unidad", "brand": "Pantene", "sku_prefix": "PERF", "category_slug": "cuidado-capilar", "price_reference": 5800},
    {"name": "Shampoo Head & Shoulders 400ml", "unit": "unidad", "brand": "Head & Shoulders", "sku_prefix": "PERF", "category_slug": "cuidado-capilar", "price_reference": 6200},
    {"name": "Acondicionador Sedal 300ml", "unit": "unidad", "brand": "Sedal", "sku_prefix": "PERF", "category_slug": "cuidado-capilar", "price_reference": 4200},
    {"name": "Crema para peinar L'Oreal Elvive 300ml", "unit": "unidad", "brand": "L'Oreal", "sku_prefix": "PERF", "category_slug": "cuidado-capilar", "price_reference": 5500},
    {"name": "Máscara capilar Capilatis nutritiva 300ml", "unit": "unidad", "brand": "Capilatis", "sku_prefix": "PERF", "category_slug": "cuidado-capilar", "price_reference": 6500},
    {"name": "Tintura L'Oreal Excellence N°4 castaño", "unit": "unidad", "brand": "L'Oreal", "sku_prefix": "PERF", "category_slug": "cuidado-capilar", "price_reference": 8500},
    {"name": "Tintura Garnier Color Naturals N°5 castaño claro", "unit": "unidad", "brand": "Garnier", "sku_prefix": "PERF", "category_slug": "cuidado-capilar", "price_reference": 7200},
    {"name": "Tintura Revlon Colorsilk N°41 castaño mediano", "unit": "unidad", "brand": "Revlon", "sku_prefix": "PERF", "category_slug": "cuidado-capilar", "price_reference": 7800},
    {"name": "Crema corporal Dove humectante 400ml", "unit": "unidad", "brand": "Dove", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 5500},
    {"name": "Crema hidratante Nivea original 400ml", "unit": "unidad", "brand": "Nivea", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 6200},
    {"name": "Crema Hinds clásica 250ml", "unit": "unidad", "brand": "Hinds", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 4500},
    {"name": "Jabón líquido Dove 250ml", "unit": "unidad", "brand": "Dove", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 3800},
    {"name": "Desodorante Rexona aerosol mujer 150ml", "unit": "unidad", "brand": "Rexona", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 4800},
    {"name": "Desodorante Axe aerosol hombre 150ml", "unit": "unidad", "brand": "Axe", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 4500},
    {"name": "Desodorante Sure roll-on mujer 50ml", "unit": "unidad", "brand": "Sure", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 3500},
    {"name": "Crema depilatoria Veet 200ml", "unit": "unidad", "brand": "Veet", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 6800},
    {"name": "Protector solar Banana Boat SPF50 180ml", "unit": "unidad", "brand": "Banana Boat", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 9500},
    {"name": "Base de maquillaje Maybelline Fit Me 30ml", "unit": "unidad", "brand": "Maybelline", "sku_prefix": "PERF", "category_slug": "maquillaje", "price_reference": 8500},
    {"name": "Labial Revlon SuperLustrous rojo", "unit": "unidad", "brand": "Revlon", "sku_prefix": "PERF", "category_slug": "maquillaje", "price_reference": 6500},
    {"name": "Rímel Rimmel London Scandal Eyes negro", "unit": "unidad", "brand": "Rimmel", "sku_prefix": "PERF", "category_slug": "maquillaje", "price_reference": 7500},
    {"name": "Sombra de ojos Maybelline The Nudes", "unit": "unidad", "brand": "Maybelline", "sku_prefix": "PERF", "category_slug": "maquillaje", "price_reference": 9200},
    {"name": "Corrector Maybelline Fit Me 6.8ml", "unit": "unidad", "brand": "Maybelline", "sku_prefix": "PERF", "category_slug": "maquillaje", "price_reference": 5500},
    {"name": "Polvo compacto Revlon ColorStay", "unit": "unidad", "brand": "Revlon", "sku_prefix": "PERF", "category_slug": "maquillaje", "price_reference": 7800},
    {"name": "Delineador Maybelline negro", "unit": "unidad", "brand": "Maybelline", "sku_prefix": "PERF", "category_slug": "maquillaje", "price_reference": 4500},
    {"name": "Perfume Soleil Toujours mujer 60ml", "unit": "unidad", "brand": null, "sku_prefix": "PERF", "category_slug": "fragancias", "price_reference": 12000},
    {"name": "Colonia Atkinsons White Rose 100ml", "unit": "unidad", "brand": "Atkinsons", "sku_prefix": "PERF", "category_slug": "fragancias", "price_reference": 9500},
    {"name": "Agua de colonia Brummel 100ml", "unit": "unidad", "brand": "Brummel", "sku_prefix": "PERF", "category_slug": "fragancias", "price_reference": 7500},
    {"name": "Esmalte Colorama N°001 blanco 8ml", "unit": "unidad", "brand": "Colorama", "sku_prefix": "PERF", "category_slug": "esmaltes", "price_reference": 1200},
    {"name": "Esmalte OPI Infinite Shine 15ml", "unit": "unidad", "brand": "OPI", "sku_prefix": "PERF", "category_slug": "esmaltes", "price_reference": 5500},
    {"name": "Quitaesmalte Kativa sin acetona 100ml", "unit": "unidad", "brand": "Kativa", "sku_prefix": "PERF", "category_slug": "esmaltes", "price_reference": 2500},
    {"name": "Esmalte endurecedor Mavala 10ml", "unit": "unidad", "brand": "Mavala", "sku_prefix": "PERF", "category_slug": "esmaltes", "price_reference": 4200},
    {"name": "Pasta dental Colgate Triple Acción 90g", "unit": "unidad", "brand": "Colgate", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 2200},
    {"name": "Cepillo de dientes Colgate Slim suave", "unit": "unidad", "brand": "Colgate", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 2800},
    {"name": "Hilo dental Oral-B 50m", "unit": "unidad", "brand": "Oral-B", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 2200},
    {"name": "Afeitadora Gillette Mach 3 repuesto x4", "unit": "unidad", "brand": "Gillette", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 8500},
    {"name": "Espuma de afeitar Gilette 250ml", "unit": "unidad", "brand": "Gillette", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 4200},
    {"name": "Algodón Johnson's 100g", "unit": "unidad", "brand": "Johnson's", "sku_prefix": "PERF", "category_slug": "cuidado-corporal", "price_reference": 2500}
  ]$products$::jsonb,
  categories = $categories$[
    {"name": "Cuidado Capilar", "slug": "cuidado-capilar", "level": 0},
    {"name": "Cuidado Corporal", "slug": "cuidado-corporal", "level": 0},
    {"name": "Maquillaje", "slug": "maquillaje", "level": 0},
    {"name": "Fragancias", "slug": "fragancias", "level": 0},
    {"name": "Esmaltes", "slug": "esmaltes", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "Sedal", "suggested_for_categories": ["cuidado-capilar"]},
    {"name": "TRESemmé", "suggested_for_categories": ["cuidado-capilar"]},
    {"name": "L'Oreal", "suggested_for_categories": ["cuidado-capilar"]},
    {"name": "Garnier", "suggested_for_categories": ["cuidado-capilar"]},
    {"name": "Dove", "suggested_for_categories": ["cuidado-capilar", "cuidado-corporal"]},
    {"name": "Nivea", "suggested_for_categories": ["cuidado-corporal"]},
    {"name": "Rexona", "suggested_for_categories": ["cuidado-corporal"]},
    {"name": "Maybelline", "suggested_for_categories": ["maquillaje"]},
    {"name": "Revlon", "suggested_for_categories": ["maquillaje", "cuidado-capilar"]},
    {"name": "Rimmel", "suggested_for_categories": ["maquillaje"]},
    {"name": "Colorama", "suggested_for_categories": ["esmaltes"]},
    {"name": "Colgate", "suggested_for_categories": ["cuidado-corporal"]}
  ]$brands$::jsonb,
  updated_at = NOW()
WHERE id = '5301a74e-4cde-449a-90bf-96b87a8a397a';


-- ============================================================
-- 6. PANADERIA
-- UUID: b2000001-0000-4000-8000-000000000004
-- sku_prefix: PANA
-- ============================================================

UPDATE business_type_templates
SET
  products = $products$[
    {"name": "Harina 0000 Pureza 1kg", "unit": "unidad", "brand": "Pureza", "sku_prefix": "PANA", "category_slug": "harinas", "price_reference": 1500},
    {"name": "Harina 0000 Blancaflor 1kg", "unit": "unidad", "brand": "Blancaflor", "sku_prefix": "PANA", "category_slug": "harinas", "price_reference": 1400},
    {"name": "Harina 000 Morixe 1kg", "unit": "unidad", "brand": "Morixe", "sku_prefix": "PANA", "category_slug": "harinas", "price_reference": 1300},
    {"name": "Harina integral Cañuelas 1kg", "unit": "unidad", "brand": "Cañuelas", "sku_prefix": "PANA", "category_slug": "harinas", "price_reference": 1600},
    {"name": "Harina de maíz Maizena 500g", "unit": "unidad", "brand": "Maizena", "sku_prefix": "PANA", "category_slug": "harinas", "price_reference": 1200},
    {"name": "Fécula de maíz Maizena 500g", "unit": "unidad", "brand": "Maizena", "sku_prefix": "PANA", "category_slug": "harinas", "price_reference": 1100},
    {"name": "Azúcar refinada Ledesma 1kg", "unit": "unidad", "brand": "Ledesma", "sku_prefix": "PANA", "category_slug": "harinas", "price_reference": 1800},
    {"name": "Azúcar impalpable Dulciora 250g", "unit": "unidad", "brand": "Dulciora", "sku_prefix": "PANA", "category_slug": "harinas", "price_reference": 1500},
    {"name": "Levadura seca Fleischmann 10g", "unit": "unidad", "brand": "Fleischmann", "sku_prefix": "PANA", "category_slug": "levaduras", "price_reference": 600},
    {"name": "Levadura fresca Calsa 500g", "unit": "unidad", "brand": "Calsa", "sku_prefix": "PANA", "category_slug": "levaduras", "price_reference": 2800},
    {"name": "Levadura fresca Puratos 500g", "unit": "unidad", "brand": "Puratos", "sku_prefix": "PANA", "category_slug": "levaduras", "price_reference": 3000},
    {"name": "Polvo de hornear Royal 200g", "unit": "unidad", "brand": "Royal", "sku_prefix": "PANA", "category_slug": "levaduras", "price_reference": 2200},
    {"name": "Bicarbonato de sodio Arm & Hammer 200g", "unit": "unidad", "brand": "Arm & Hammer", "sku_prefix": "PANA", "category_slug": "levaduras", "price_reference": 1500},
    {"name": "Margarina Primavera bloque 500g", "unit": "unidad", "brand": "Primavera", "sku_prefix": "PANA", "category_slug": "harinas", "price_reference": 3500},
    {"name": "Manteca sin sal La Serenísima 200g", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "PANA", "category_slug": "harinas", "price_reference": 3200},
    {"name": "Dulce de leche repostero La Serenísima 1kg", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "PANA", "category_slug": "dulces", "price_reference": 6500},
    {"name": "Dulce de leche repostero SanCor 1kg", "unit": "unidad", "brand": "SanCor", "sku_prefix": "PANA", "category_slug": "dulces", "price_reference": 6200},
    {"name": "Mermelada frutilla Arcor industrial 5kg", "unit": "unidad", "brand": "Arcor", "sku_prefix": "PANA", "category_slug": "dulces", "price_reference": 18000},
    {"name": "Mermelada damasco Arcor industrial 5kg", "unit": "unidad", "brand": "Arcor", "sku_prefix": "PANA", "category_slug": "dulces", "price_reference": 17500},
    {"name": "Pasta de maní Nuteller 1kg", "unit": "unidad", "brand": null, "sku_prefix": "PANA", "category_slug": "dulces", "price_reference": 9500},
    {"name": "Cobertura chocolate semiamargo Fenix 1kg", "unit": "unidad", "brand": "Fenix", "sku_prefix": "PANA", "category_slug": "coberturas", "price_reference": 8500},
    {"name": "Cobertura chocolate blanco Fenix 1kg", "unit": "unidad", "brand": "Fenix", "sku_prefix": "PANA", "category_slug": "coberturas", "price_reference": 9200},
    {"name": "Cobertura chocolate con leche Fenix 1kg", "unit": "unidad", "brand": "Fenix", "sku_prefix": "PANA", "category_slug": "coberturas", "price_reference": 8800},
    {"name": "Granillo de chocolate Fenix 1kg", "unit": "unidad", "brand": "Fenix", "sku_prefix": "PANA", "category_slug": "coberturas", "price_reference": 7500},
    {"name": "Glacé de naranja Molinos 500g", "unit": "unidad", "brand": "Molinos", "sku_prefix": "PANA", "category_slug": "coberturas", "price_reference": 4500},
    {"name": "Esencia de vainilla Fleischmann 60ml", "unit": "unidad", "brand": "Fleischmann", "sku_prefix": "PANA", "category_slug": "esencias", "price_reference": 1800},
    {"name": "Esencia de limón Moño Azul 50ml", "unit": "unidad", "brand": "Moño Azul", "sku_prefix": "PANA", "category_slug": "esencias", "price_reference": 1500},
    {"name": "Esencia de naranja Moño Azul 50ml", "unit": "unidad", "brand": "Moño Azul", "sku_prefix": "PANA", "category_slug": "esencias", "price_reference": 1500},
    {"name": "Colorante alimentario rojo Moño Azul 25ml", "unit": "unidad", "brand": "Moño Azul", "sku_prefix": "PANA", "category_slug": "esencias", "price_reference": 1200},
    {"name": "Gel en polvo de frutas Terrabusi 50g", "unit": "unidad", "brand": "Terrabusi", "sku_prefix": "PANA", "category_slug": "esencias", "price_reference": 800},
    {"name": "Papel mantequilla rollo 50m", "unit": "unidad", "brand": null, "sku_prefix": "PANA", "category_slug": "moldes", "price_reference": 5500},
    {"name": "Molde de papel muffin x12", "unit": "unidad", "brand": null, "sku_prefix": "PANA", "category_slug": "moldes", "price_reference": 1200},
    {"name": "Molde de papel cupcake x100", "unit": "unidad", "brand": null, "sku_prefix": "PANA", "category_slug": "moldes", "price_reference": 2500},
    {"name": "Bandeja de aluminio rectangular 30x20cm", "unit": "unidad", "brand": null, "sku_prefix": "PANA", "category_slug": "moldes", "price_reference": 3200},
    {"name": "Caja para torta 30cm x30cm x12cm", "unit": "unidad", "brand": null, "sku_prefix": "PANA", "category_slug": "moldes", "price_reference": 2800},
    {"name": "Bolsa celofán para pan 25x40cm x100", "unit": "unidad", "brand": null, "sku_prefix": "PANA", "category_slug": "moldes", "price_reference": 3500},
    {"name": "Crema pastelera en polvo Unidal 1kg", "unit": "unidad", "brand": "Unidal", "sku_prefix": "PANA", "category_slug": "coberturas", "price_reference": 4200},
    {"name": "Mejorador para pan Puratos 1kg", "unit": "unidad", "brand": "Puratos", "sku_prefix": "PANA", "category_slug": "levaduras", "price_reference": 5500},
    {"name": "Sal fina Celusal 1kg", "unit": "unidad", "brand": "Celusal", "sku_prefix": "PANA", "category_slug": "harinas", "price_reference": 900},
    {"name": "Semillas de sésamo x200g", "unit": "unidad", "brand": null, "sku_prefix": "PANA", "category_slug": "esencias", "price_reference": 1800},
    {"name": "Anís en grano x100g", "unit": "unidad", "brand": null, "sku_prefix": "PANA", "category_slug": "esencias", "price_reference": 1500},
    {"name": "Coco rallado sin azúcar 200g", "unit": "unidad", "brand": null, "sku_prefix": "PANA", "category_slug": "coberturas", "price_reference": 2500}
  ]$products$::jsonb,
  categories = $categories$[
    {"name": "Harinas e Ingredientes Secos", "slug": "harinas", "level": 0},
    {"name": "Levaduras e Impulsores", "slug": "levaduras", "level": 0},
    {"name": "Dulces y Rellenos", "slug": "dulces", "level": 0},
    {"name": "Coberturas y Decoración", "slug": "coberturas", "level": 0},
    {"name": "Esencias y Colorantes", "slug": "esencias", "level": 0},
    {"name": "Moldes y Packaging", "slug": "moldes", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "Blancaflor", "suggested_for_categories": ["harinas"]},
    {"name": "Pureza", "suggested_for_categories": ["harinas"]},
    {"name": "Morixe", "suggested_for_categories": ["harinas"]},
    {"name": "Fleischmann", "suggested_for_categories": ["levaduras", "esencias"]},
    {"name": "Calsa", "suggested_for_categories": ["levaduras"]},
    {"name": "Puratos", "suggested_for_categories": ["levaduras"]},
    {"name": "La Serenísima", "suggested_for_categories": ["dulces", "harinas"]},
    {"name": "Arcor", "suggested_for_categories": ["dulces"]},
    {"name": "Fenix", "suggested_for_categories": ["coberturas"]},
    {"name": "Moño Azul", "suggested_for_categories": ["esencias"]},
    {"name": "Royal", "suggested_for_categories": ["levaduras"]},
    {"name": "Ledesma", "suggested_for_categories": ["harinas"]}
  ]$brands$::jsonb,
  updated_at = NOW()
WHERE id = 'b2000001-0000-4000-8000-000000000004';


-- ============================================================
-- 7. CARNICERIA
-- UUID: b2000001-0000-4000-8000-000000000006
-- sku_prefix: CARN
-- ============================================================

UPDATE business_type_templates
SET
  products = $products$[
    {"name": "Asado de tira vacuno x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 7500},
    {"name": "Vacío vacuno x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 9500},
    {"name": "Nalga vacuna x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 8800},
    {"name": "Cuadril vacuno x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 9200},
    {"name": "Bife de chorizo x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 12500},
    {"name": "Lomo vacuno x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 16000},
    {"name": "Paleta vacuna x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 7200},
    {"name": "Carne molida común x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 6500},
    {"name": "Carne molida especial x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 8000},
    {"name": "Peceto vacuno x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 10000},
    {"name": "Matambre vacuno x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 9800},
    {"name": "Osobuco vacuno x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 7000},
    {"name": "Hígado vacuno x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "achuras", "price_reference": 5500},
    {"name": "Riñón vacuno x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "achuras", "price_reference": 4800},
    {"name": "Mondongo x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "achuras", "price_reference": 4500},
    {"name": "Corazón vacuno x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "achuras", "price_reference": 5000},
    {"name": "Bondiola de cerdo x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "cerdo", "price_reference": 7500},
    {"name": "Costeletas de cerdo x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "cerdo", "price_reference": 6800},
    {"name": "Panceta de cerdo x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "cerdo", "price_reference": 7200},
    {"name": "Paleta de cerdo x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "cerdo", "price_reference": 6500},
    {"name": "Pollo entero x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "pollo", "price_reference": 3500},
    {"name": "Pechuga de pollo sin hueso x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "pollo", "price_reference": 5500},
    {"name": "Muslo y contramuslo de pollo x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "pollo", "price_reference": 3800},
    {"name": "Alitas de pollo x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "pollo", "price_reference": 3200},
    {"name": "Suprema de pollo x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "pollo", "price_reference": 6200},
    {"name": "Chorizo parrillero x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "embutidos", "price_reference": 6500},
    {"name": "Morcilla x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "embutidos", "price_reference": 5800},
    {"name": "Salchicha viena Paladini 300g", "unit": "unidad", "brand": "Paladini", "sku_prefix": "CARN", "category_slug": "embutidos", "price_reference": 3800},
    {"name": "Salame Milano Paladini 200g", "unit": "unidad", "brand": "Paladini", "sku_prefix": "CARN", "category_slug": "embutidos", "price_reference": 4500},
    {"name": "Jamón cocido Paladini x kg", "unit": "kg", "brand": "Paladini", "sku_prefix": "CARN", "category_slug": "embutidos", "price_reference": 12000},
    {"name": "Panceta ahumada Paladini x kg", "unit": "kg", "brand": "Paladini", "sku_prefix": "CARN", "category_slug": "embutidos", "price_reference": 9500},
    {"name": "Longaniza parrillera x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "embutidos", "price_reference": 6800},
    {"name": "Hamburguesa vacuna casera x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 7800},
    {"name": "Milanesa de nalga x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "vacuno", "price_reference": 9500},
    {"name": "Milanesa de pollo x kg", "unit": "kg", "brand": null, "sku_prefix": "CARN", "category_slug": "pollo", "price_reference": 6000}
  ]$products$::jsonb,
  categories = $categories$[
    {"name": "Vacuno", "slug": "vacuno", "level": 0},
    {"name": "Cerdo", "slug": "cerdo", "level": 0},
    {"name": "Pollo", "slug": "pollo", "level": 0},
    {"name": "Embutidos", "slug": "embutidos", "level": 0},
    {"name": "Achuras", "slug": "achuras", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "Paladini", "suggested_for_categories": ["embutidos"]},
    {"name": "Swift", "suggested_for_categories": ["embutidos"]},
    {"name": "La Salamandra", "suggested_for_categories": ["embutidos"]},
    {"name": "Granja del Sol", "suggested_for_categories": ["pollo"]},
    {"name": "Molfino", "suggested_for_categories": ["pollo"]}
  ]$brands$::jsonb,
  updated_at = NOW()
WHERE id = 'b2000001-0000-4000-8000-000000000006';


-- ============================================================
-- 8. VERDULERIA
-- UUID: b2000001-0000-4000-8000-000000000005
-- sku_prefix: VERDU
-- ============================================================

UPDATE business_type_templates
SET
  products = $products$[
    {"name": "Manzana Red x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 2800},
    {"name": "Banana x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 1800},
    {"name": "Naranja x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 1500},
    {"name": "Mandarina x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 2000},
    {"name": "Pera Williams x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 3200},
    {"name": "Uva negra x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 3800},
    {"name": "Limón x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 2200},
    {"name": "Pomelo x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 1800},
    {"name": "Durazno x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 3500},
    {"name": "Frutillas x bandeja 250g", "unit": "bandeja", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 2200},
    {"name": "Sandía x unidad", "unit": "unidad", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 4500},
    {"name": "Melón x unidad", "unit": "unidad", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 3800},
    {"name": "Ananá x unidad", "unit": "unidad", "brand": null, "sku_prefix": "VERDU", "category_slug": "frutas", "price_reference": 4200},
    {"name": "Tomate x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-fruto", "price_reference": 2500},
    {"name": "Pimiento rojo x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-fruto", "price_reference": 3800},
    {"name": "Pimiento verde x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-fruto", "price_reference": 3200},
    {"name": "Berenjena x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-fruto", "price_reference": 3000},
    {"name": "Zapallito redondo x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-fruto", "price_reference": 2200},
    {"name": "Choclo x unidad", "unit": "unidad", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-fruto", "price_reference": 900},
    {"name": "Pepino x unidad", "unit": "unidad", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-fruto", "price_reference": 1200},
    {"name": "Papa x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-raiz", "price_reference": 1500},
    {"name": "Cebolla x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-raiz", "price_reference": 1800},
    {"name": "Zanahoria x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-raiz", "price_reference": 2000},
    {"name": "Remolacha x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-raiz", "price_reference": 2200},
    {"name": "Zapallo anco x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-raiz", "price_reference": 1500},
    {"name": "Ajo x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-raiz", "price_reference": 8500},
    {"name": "Batata x kg", "unit": "kg", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-raiz", "price_reference": 2500},
    {"name": "Lechuga criolla x unidad", "unit": "unidad", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-hoja", "price_reference": 1200},
    {"name": "Lechuga mantecosa x unidad", "unit": "unidad", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-hoja", "price_reference": 1400},
    {"name": "Rúcula x atado", "unit": "atado", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-hoja", "price_reference": 1500},
    {"name": "Acelga x atado", "unit": "atado", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-hoja", "price_reference": 1200},
    {"name": "Espinaca x atado", "unit": "atado", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-hoja", "price_reference": 1300},
    {"name": "Brócoli x unidad", "unit": "unidad", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-hoja", "price_reference": 2200},
    {"name": "Coliflor x unidad", "unit": "unidad", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-hoja", "price_reference": 2500},
    {"name": "Repollo blanco x unidad", "unit": "unidad", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-hoja", "price_reference": 1800},
    {"name": "Perejil x atado", "unit": "atado", "brand": null, "sku_prefix": "VERDU", "category_slug": "hierbas", "price_reference": 600},
    {"name": "Albahaca x atado", "unit": "atado", "brand": null, "sku_prefix": "VERDU", "category_slug": "hierbas", "price_reference": 800},
    {"name": "Ciboulette x atado", "unit": "atado", "brand": null, "sku_prefix": "VERDU", "category_slug": "hierbas", "price_reference": 700},
    {"name": "Cilantro x atado", "unit": "atado", "brand": null, "sku_prefix": "VERDU", "category_slug": "hierbas", "price_reference": 700},
    {"name": "Apio x atado", "unit": "atado", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-hoja", "price_reference": 1500},
    {"name": "Puerro x unidad", "unit": "unidad", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-raiz", "price_reference": 1200},
    {"name": "Cebollita de verdeo x atado", "unit": "atado", "brand": null, "sku_prefix": "VERDU", "category_slug": "verduras-raiz", "price_reference": 900}
  ]$products$::jsonb,
  categories = $categories$[
    {"name": "Frutas", "slug": "frutas", "level": 0},
    {"name": "Verduras de Fruto", "slug": "verduras-fruto", "level": 0},
    {"name": "Verduras de Raíz", "slug": "verduras-raiz", "level": 0},
    {"name": "Verduras de Hoja", "slug": "verduras-hoja", "level": 0},
    {"name": "Hierbas Frescas", "slug": "hierbas", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "Del Monte", "suggested_for_categories": ["frutas"]},
    {"name": "Chiquita", "suggested_for_categories": ["frutas"]}
  ]$brands$::jsonb,
  updated_at = NOW()
WHERE id = 'b2000001-0000-4000-8000-000000000005';


-- =============================================================================
-- CONTEO DE PRODUCTOS POR RUBRO
-- Fuente: catálogo real mercado argentino / NEA (Posadas, Misiones)
-- Fecha de research: 2026-04
-- =============================================================================
-- almacen    (b2000001-0000-4000-8000-000000000001): 61 productos
-- ferreteria (b2000001-0000-4000-8000-000000000003): 50 productos
-- vinoteca   (3c546ca9-0be2-4a30-97bb-f917dd6b251f): 38 productos
-- kiosco     (b2000001-0000-4000-8000-000000000002): 45 productos
-- perfumeria (5301a74e-4cde-449a-90bf-96b87a8a397a): 40 productos
-- panaderia  (b2000001-0000-4000-8000-000000000004): 42 productos
-- carniceria (b2000001-0000-4000-8000-000000000006): 35 productos
-- verduleria (b2000001-0000-4000-8000-000000000005): 42 productos
-- =============================================================================
-- TOTAL: 353 productos en 8 rubros
-- =============================================================================
