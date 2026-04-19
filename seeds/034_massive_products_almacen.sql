-- =============================================================================
-- SEED 034: Catalogo masivo de almacen argentino (300 productos reales)
-- Actualiza business_type_templates.products para 'almacen'
-- Marcas reales argentinas: La Serenisima, Arcor, Marolio, Coca-Cola, Quilmes,
--   Molinos, Terrabusi, Bagley, Ala, Skip, Dove, Higienol, etc.
-- =============================================================================

UPDATE business_type_templates
SET products = '[
  {"name": "Aceite de girasol Cocinero 900ml", "category_slug": "aceites-vinagres", "brand": "Cocinero", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Aceite de girasol Natura 900ml", "category_slug": "aceites-vinagres", "brand": "Natura", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Aceite de girasol Cañuelas 1.5L", "category_slug": "aceites-vinagres", "brand": "Canuelas", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Aceite de oliva Cocinero 500ml", "category_slug": "aceites-vinagres", "brand": "Cocinero", "price_reference": 6500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Aceite mezcla Marolio 900ml", "category_slug": "aceites-vinagres", "brand": "Marolio", "price_reference": 2400, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Vinagre de alcohol Menoyo 1L", "category_slug": "aceites-vinagres", "brand": "Menoyo", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Vinagre de manzana Menoyo 500ml", "category_slug": "aceites-vinagres", "brand": "Menoyo", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Aceto balsamico Marolio 250ml", "category_slug": "aceites-vinagres", "brand": "Marolio", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Alfajor Havanna chocolate 65g", "category_slug": "alfajores-chocolates", "brand": "Havanna", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Alfajor Cachafaz dulce de leche 60g", "category_slug": "alfajores-chocolates", "brand": "Cachafaz", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Alfajor Guaymallen triple chocolate", "category_slug": "alfajores-chocolates", "brand": "Guaymallen", "price_reference": 600, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Alfajor Jorgito negro 55g", "category_slug": "alfajores-chocolates", "brand": "Jorgito", "price_reference": 800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Alfajor Terrabusi triple 70g", "category_slug": "alfajores-chocolates", "brand": "Terrabusi", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Chocolate Aguila taza 100g", "category_slug": "alfajores-chocolates", "brand": "Aguila", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Chocolate Milka leger 100g", "category_slug": "alfajores-chocolates", "brand": "Milka", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Bon o Bon clasico x6", "category_slug": "alfajores-chocolates", "brand": "Arcor", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Yerba mate Rosamonte 1kg", "category_slug": "almacen-seco", "brand": "Rosamonte", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Yerba mate Playadito 1kg", "category_slug": "almacen-seco", "brand": "Playadito", "price_reference": 4800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Yerba mate Taragui 1kg", "category_slug": "almacen-seco", "brand": "Taragui", "price_reference": 4200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Yerba mate Amanda 1kg", "category_slug": "almacen-seco", "brand": "Amanda", "price_reference": 4600, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Azucar Ledesma 1kg", "category_slug": "almacen-seco", "brand": "Ledesma", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Sal fina Celusal 500g", "category_slug": "almacen-seco", "brand": "Celusal", "price_reference": 700, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Sal gruesa Dos Anclas 1kg", "category_slug": "almacen-seco", "brand": "Dos Anclas", "price_reference": 800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Cafe molido La Virginia 250g", "category_slug": "almacen-seco", "brand": "La Virginia", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Cafe molido Cabrales 250g", "category_slug": "almacen-seco", "brand": "Cabrales", "price_reference": 5000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Te La Virginia saquitos x25", "category_slug": "almacen-seco", "brand": "La Virginia", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Arroz largo fino Gallo 1kg", "category_slug": "arroz-legumbres", "brand": "Gallo", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Arroz largo fino Molinos Ala 1kg", "category_slug": "arroz-legumbres", "brand": "Molinos Ala", "price_reference": 1900, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Arroz parboil Gallo 1kg", "category_slug": "arroz-legumbres", "brand": "Gallo", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Arroz integral Gallo 1kg", "category_slug": "arroz-legumbres", "brand": "Gallo", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Lentejas Marolio 500g", "category_slug": "arroz-legumbres", "brand": "Marolio", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Porotos pallares Marolio 500g", "category_slug": "arroz-legumbres", "brand": "Marolio", "price_reference": 1600, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Garbanzos Marolio 500g", "category_slug": "arroz-legumbres", "brand": "Marolio", "price_reference": 1700, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Porotos negros Marolio 500g", "category_slug": "arroz-legumbres", "brand": "Marolio", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Arvejas secas Marolio 500g", "category_slug": "arroz-legumbres", "brand": "Marolio", "price_reference": 1400, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Agua mineral Villa del Sur 1.5L", "category_slug": "bebidas-almacen", "brand": "Villa del Sur", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Agua mineral Eco de los Andes 1.5L", "category_slug": "bebidas-almacen", "brand": "Eco de los Andes", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Soda Ivess 2.25L", "category_slug": "bebidas-almacen", "brand": "Ivess", "price_reference": 1600, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Tonica Paso de los Toros 1.5L", "category_slug": "bebidas-almacen", "brand": "Paso de los Toros", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Pomelo Paso de los Toros 1.5L", "category_slug": "bebidas-almacen", "brand": "Paso de los Toros", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Cerveza Quilmes 1L retornable", "category_slug": "cervezas-vinos", "brand": "Quilmes", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Cerveza Stella Artois 1L", "category_slug": "cervezas-vinos", "brand": "Stella Artois", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Cerveza Brahma 1L", "category_slug": "cervezas-vinos", "brand": "Brahma", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Cerveza Quilmes lata 473ml", "category_slug": "cervezas-vinos", "brand": "Quilmes", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Cerveza Schneider lata 473ml", "category_slug": "cervezas-vinos", "brand": "Schneider", "price_reference": 1600, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Fernet Branca 750ml", "category_slug": "cervezas-vinos", "brand": "Fernet Branca", "price_reference": 12000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Fernet Branca 450ml", "category_slug": "cervezas-vinos", "brand": "Fernet Branca", "price_reference": 8000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Gancia 950ml", "category_slug": "cervezas-vinos", "brand": "Gancia", "price_reference": 6500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Vino tinto Estancia Mendoza 750ml", "category_slug": "cervezas-vinos", "brand": "Estancia Mendoza", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Vino tinto Termidor 1L", "category_slug": "cervezas-vinos", "brand": "Termidor", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Tomates enteros pelados La Campagnola 400g", "category_slug": "conservas-enlatados", "brand": "La Campagnola", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Pure de tomate La Campagnola 520g", "category_slug": "conservas-enlatados", "brand": "La Campagnola", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Salsa de tomate Arcor 520g", "category_slug": "conservas-enlatados", "brand": "Arcor", "price_reference": 1400, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Extracto de tomate La Campagnola 370g", "category_slug": "conservas-enlatados", "brand": "La Campagnola", "price_reference": 1600, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Arvejas Marolio 350g", "category_slug": "conservas-enlatados", "brand": "Marolio", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Choclo cremoso Marolio 350g", "category_slug": "conservas-enlatados", "brand": "Marolio", "price_reference": 1400, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jardinera Marolio 350g", "category_slug": "conservas-enlatados", "brand": "Marolio", "price_reference": 1300, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Atun en aceite La Campagnola 170g", "category_slug": "conservas-enlatados", "brand": "La Campagnola", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Caballa al natural Marolio 380g", "category_slug": "conservas-enlatados", "brand": "Marolio", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Duraznos en almibar Marolio 820g", "category_slug": "conservas-enlatados", "brand": "Marolio", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Aceitunas verdes Marolio 220g", "category_slug": "conservas-enlatados", "brand": "Marolio", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Detergente Magistral limon 750ml", "category_slug": "detergentes-jabones", "brand": "Magistral", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Detergente Ala regular 750ml", "category_slug": "detergentes-jabones", "brand": "Ala", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jabon en polvo Skip 800g", "category_slug": "detergentes-jabones", "brand": "Skip", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jabon en polvo Ala 800g", "category_slug": "detergentes-jabones", "brand": "Ala", "price_reference": 3800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jabon liquido Skip botella 800ml", "category_slug": "detergentes-jabones", "brand": "Skip", "price_reference": 5500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jabon en pan Federal blanco x3", "category_slug": "detergentes-jabones", "brand": "Federal", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Suavizante Comfort clasico 900ml", "category_slug": "detergentes-jabones", "brand": "Comfort", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Suavizante Vivere clasico 900ml", "category_slug": "detergentes-jabones", "brand": "Vivere", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Jamon cocido feteado Paladini 150g", "category_slug": "fiambreria", "brand": "Paladini", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Paleta feteada Paladini 150g", "category_slug": "fiambreria", "brand": "Paladini", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jamon crudo feteado Cagnoli 120g", "category_slug": "fiambreria", "brand": "Cagnoli", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Queso cremoso feteado La Serenisima 200g", "category_slug": "fiambreria", "brand": "La Serenisima", "price_reference": 4000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Mortadela feteada Paladini 150g", "category_slug": "fiambreria", "brand": "Paladini", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Salchichas tipo viena Paladini x6", "category_slug": "fiambres-embutidos", "brand": "Paladini", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Salchichas Vienissima x6", "category_slug": "fiambres-embutidos", "brand": "Vienissima", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Chorizo tipo cantimpalo Cagnoli 250g", "category_slug": "fiambres-embutidos", "brand": "Cagnoli", "price_reference": 4000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Lomito ahumado Paladini 200g", "category_slug": "fiambres-embutidos", "brand": "Paladini", "price_reference": 5500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Bondiola feteada Cagnoli 150g", "category_slug": "fiambres-embutidos", "brand": "Cagnoli", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Salame milan Cagnoli 150g", "category_slug": "fiambres-embutidos", "brand": "Cagnoli", "price_reference": 3800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Galletitas Traviata 303g", "category_slug": "galletitas-almacen", "brand": "Bagley", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Galletitas Criollitas 300g", "category_slug": "galletitas-almacen", "brand": "Bagley", "price_reference": 1600, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Galletitas Express 320g", "category_slug": "galletitas-almacen", "brand": "Bagley", "price_reference": 1700, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Galletitas Lincoln 153g", "category_slug": "galletitas-almacen", "brand": "Bagley", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Galletitas Rumba chocolate 112g", "category_slug": "galletitas-almacen", "brand": "Bagley", "price_reference": 1400, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Galletitas Chocolinas 170g", "category_slug": "galletitas-almacen", "brand": "Terrabusi", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Galletitas Terrabusi variedad 400g", "category_slug": "galletitas-almacen", "brand": "Terrabusi", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Galletitas Cerealitas clasicas 200g", "category_slug": "galletitas-almacen", "brand": "Cerealitas", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Galletitas Oreo 117g", "category_slug": "galletitas-almacen", "brand": "Terrabusi", "price_reference": 1600, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Galletitas de arroz Molinos 150g", "category_slug": "galletitas-almacen", "brand": "Molinos", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Coca-Cola 2.25L", "category_slug": "gaseosas-aguas", "brand": "Coca-Cola", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Coca-Cola 500ml", "category_slug": "gaseosas-aguas", "brand": "Coca-Cola", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Coca-Cola Zero 2.25L", "category_slug": "gaseosas-aguas", "brand": "Coca-Cola", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Pepsi 2.25L", "category_slug": "gaseosas-aguas", "brand": "Pepsi", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Sprite 2.25L", "category_slug": "gaseosas-aguas", "brand": "Sprite", "price_reference": 3300, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Fanta naranja 2.25L", "category_slug": "gaseosas-aguas", "brand": "Fanta", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Manaos cola 2.25L", "category_slug": "gaseosas-aguas", "brand": "Manaos", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Manaos uva 2.25L", "category_slug": "gaseosas-aguas", "brand": "Manaos", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Agua mineral Villavicencio 1.5L", "category_slug": "gaseosas-aguas", "brand": "Villavicencio", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Agua saborizada Levite pomelo 1.5L", "category_slug": "gaseosas-aguas", "brand": "Levite", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Tita 36g", "category_slug": "golosinas-snacks", "brand": "Terrabusi", "price_reference": 700, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Rhodesia 22g", "category_slug": "golosinas-snacks", "brand": "Terrabusi", "price_reference": 600, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Chicle Beldent x10", "category_slug": "golosinas-snacks", "brand": "Beldent", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Caramelos Butter Toffees x150g", "category_slug": "golosinas-snacks", "brand": "Arcor", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Chupetines Pico Dulce x24", "category_slug": "golosinas-snacks", "brand": "Arcor", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Turron Arcor mani 25g", "category_slug": "golosinas-snacks", "brand": "Arcor", "price_reference": 500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Gomitas Mogul frutas 150g", "category_slug": "golosinas-snacks", "brand": "Arcor", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Pastillas DRF menta 23g", "category_slug": "golosinas-snacks", "brand": "Arcor", "price_reference": 500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Harina 000 Blancaflor 1kg", "category_slug": "harinas-premezclas", "brand": "Blancaflor", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Harina 0000 Blancaflor 1kg", "category_slug": "harinas-premezclas", "brand": "Blancaflor", "price_reference": 1600, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Harina leudante Blancaflor 1kg", "category_slug": "harinas-premezclas", "brand": "Blancaflor", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Harina integral Pureza 1kg", "category_slug": "harinas-premezclas", "brand": "Pureza", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Premezcla para pizza Blancaflor 1kg", "category_slug": "harinas-premezclas", "brand": "Blancaflor", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Premezcla para torta Exquisita chocolate 540g", "category_slug": "harinas-premezclas", "brand": "Exquisita", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Premezcla para flan Exquisita 60g", "category_slug": "harinas-premezclas", "brand": "Exquisita", "price_reference": 900, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Fecula de maiz Maizena 500g", "category_slug": "harinas-premezclas", "brand": "Maizena", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Polenta Presto Pronta 500g", "category_slug": "harinas-premezclas", "brand": "Presto Pronta", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Levadura seca Levex 100g", "category_slug": "harinas-premezclas", "brand": "Levex", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Jabon de tocador Dove original 90g", "category_slug": "jabones-desodorantes", "brand": "Dove", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jabon de tocador Lux 125g", "category_slug": "jabones-desodorantes", "brand": "Lux", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jabon de tocador Rexona 125g", "category_slug": "jabones-desodorantes", "brand": "Rexona", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Desodorante Rexona aerosol 150ml", "category_slug": "jabones-desodorantes", "brand": "Rexona", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Desodorante Dove aerosol 150ml", "category_slug": "jabones-desodorantes", "brand": "Dove", "price_reference": 5000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Desodorante Axe aerosol 150ml", "category_slug": "jabones-desodorantes", "brand": "Axe", "price_reference": 4800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jabon liquido para manos Espadol 220ml", "category_slug": "jabones-desodorantes", "brand": "Espadol", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Alcohol en gel Espadol 250ml", "category_slug": "jabones-desodorantes", "brand": "Espadol", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Jugo Cepita naranja 1L", "category_slug": "jugos-polvos", "brand": "Cepita", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jugo Cepita multifruta 1L", "category_slug": "jugos-polvos", "brand": "Cepita", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jugo en polvo Tang naranja 25g", "category_slug": "jugos-polvos", "brand": "Tang", "price_reference": 500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jugo en polvo Tang naranja dulce 18g", "category_slug": "jugos-polvos", "brand": "Tang", "price_reference": 500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jugo en polvo Arcor manzana 20g", "category_slug": "jugos-polvos", "brand": "Arcor", "price_reference": 400, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Cacao en polvo Toddy 180g", "category_slug": "jugos-polvos", "brand": "Toddy", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Cacao en polvo Nesquik 360g", "category_slug": "jugos-polvos", "brand": "Nesquik", "price_reference": 4000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Leche entera La Serenisima sachet 1L", "category_slug": "lacteos-frescos", "brand": "La Serenisima", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Leche descremada La Serenisima sachet 1L", "category_slug": "lacteos-frescos", "brand": "La Serenisima", "price_reference": 1900, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Crema de leche La Serenisima 200ml", "category_slug": "lacteos-frescos", "brand": "La Serenisima", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Dulce de leche La Serenisima 400g", "category_slug": "lacteos-frescos", "brand": "La Serenisima", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Dulce de leche Ilolay 400g", "category_slug": "lacteos-frescos", "brand": "Ilolay", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Manteca La Serenisima 200g", "category_slug": "lacteos-frescos", "brand": "La Serenisima", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Ricota La Serenisima 250g", "category_slug": "lacteos-frescos", "brand": "La Serenisima", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Postre Serenito chocolate 100g", "category_slug": "lacteos-frescos", "brand": "La Serenisima", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Lavandina Ayudin original 1L", "category_slug": "lavandina-desinfectantes", "brand": "Ayudin", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Lavandina Ayudin concentrada 1L", "category_slug": "lavandina-desinfectantes", "brand": "Ayudin", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Lavandina Querubin 2L", "category_slug": "lavandina-desinfectantes", "brand": "Querubin", "price_reference": 1600, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Desinfectante Procenex original 900ml", "category_slug": "lavandina-desinfectantes", "brand": "Procenex", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Desinfectante Procenex lavanda 900ml", "category_slug": "lavandina-desinfectantes", "brand": "Procenex", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Leche entera larga vida La Serenisima 1L", "category_slug": "leches", "brand": "La Serenisima", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Leche descremada larga vida La Serenisima 1L", "category_slug": "leches", "brand": "La Serenisima", "price_reference": 2300, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Leche entera larga vida Sancor 1L", "category_slug": "leches", "brand": "Sancor", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Leche chocolatada Chocolatada Sancor 1L", "category_slug": "leches", "brand": "Sancor", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Leche en polvo entera La Serenisima 800g", "category_slug": "leches", "brand": "La Serenisima", "price_reference": 6500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Leche condensada La Serenisima 395g", "category_slug": "leches", "brand": "La Serenisima", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Limpiador Cif crema original 750ml", "category_slug": "limpieza", "brand": "Cif", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Limpiador Cif crema limon 750ml", "category_slug": "limpieza", "brand": "Cif", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Limpiador pisos Procenex lavanda 900ml", "category_slug": "limpieza", "brand": "Procenex", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Limpiador multiuso Mr Musculo 500ml", "category_slug": "limpieza", "brand": "Mr Musculo", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Limpiador vidrios Mr Musculo 500ml", "category_slug": "limpieza", "brand": "Mr Musculo", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Esponja Scotch-Brite multiuso", "category_slug": "limpieza", "brand": "Scotch-Brite", "price_reference": 800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Trapo de piso 48x60cm", "category_slug": "limpieza", "brand": null, "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Bolsa de residuos 45x60 x30 unidades", "category_slug": "limpieza", "brand": null, "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Bolsa de consorcio 80x110 x10 unidades", "category_slug": "limpieza", "brand": null, "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Prepizza La Saltenisima x2", "category_slug": "panaderia-reposteria", "brand": "La Saltenisima", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Tapas de empanada La Saltenisima x12", "category_slug": "panaderia-reposteria", "brand": "La Saltenisima", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Tapas de empanada para horno La Saltenisima x12", "category_slug": "panaderia-reposteria", "brand": "La Saltenisima", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Tapas de tarta La Saltenisima x2", "category_slug": "panaderia-reposteria", "brand": "La Saltenisima", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Dulce de membrillo La Campagnola 500g", "category_slug": "panaderia-reposteria", "brand": "La Campagnola", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Dulce de batata La Campagnola 500g", "category_slug": "panaderia-reposteria", "brand": "La Campagnola", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Mermelada de durazno Arcor 454g", "category_slug": "panaderia-reposteria", "brand": "Arcor", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Mermelada de frutilla Arcor 454g", "category_slug": "panaderia-reposteria", "brand": "Arcor", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Manteca de mani Aruba 350g", "category_slug": "panaderia-reposteria", "brand": "Aruba", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Miel de abejas pura 500g", "category_slug": "panaderia-reposteria", "brand": null, "price_reference": 4000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Panales Huggies Classic talle M x30", "category_slug": "panales", "brand": "Huggies", "price_reference": 12000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Panales Huggies Classic talle G x28", "category_slug": "panales", "brand": "Huggies", "price_reference": 12000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Panales Huggies Classic talle XG x24", "category_slug": "panales", "brand": "Huggies", "price_reference": 12000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Panales Pampers SuperSec talle M x30", "category_slug": "panales", "brand": "Pampers", "price_reference": 13000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Panales Pampers SuperSec talle G x26", "category_slug": "panales", "brand": "Pampers", "price_reference": 13000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Panales Babysec Ultra talle G x28", "category_slug": "panales", "brand": "Babysec", "price_reference": 9000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Toallitas humedas Huggies x48", "category_slug": "panales", "brand": "Huggies", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Pan lactal Bimbo 350g", "category_slug": "pan-envasado", "brand": "Bimbo", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Pan lactal Fargo 350g", "category_slug": "pan-envasado", "brand": "Fargo", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Pan lactal integral Bimbo 350g", "category_slug": "pan-envasado", "brand": "Bimbo", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Pan rallado Preferido 500g", "category_slug": "pan-envasado", "brand": "Preferido", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Tostadas Bimbo clasicas 200g", "category_slug": "pan-envasado", "brand": "Bimbo", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Grisines Marolio 100g", "category_slug": "pan-envasado", "brand": "Marolio", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Papel higienico Higienol 30m x4", "category_slug": "papel-higienico", "brand": "Higienol", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Papel higienico Higienol doble hoja 30m x4", "category_slug": "papel-higienico", "brand": "Higienol", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Papel higienico Elite doble hoja 30m x4", "category_slug": "papel-higienico", "brand": "Elite", "price_reference": 5000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Papel higienico Sussex doble hoja 30m x4", "category_slug": "papel-higienico", "brand": "Sussex", "price_reference": 4800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Rollo de cocina Sussex x3", "category_slug": "papel-higienico", "brand": "Sussex", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Servilletas Sussex x100", "category_slug": "papel-higienico", "brand": "Sussex", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Panuelos descartables Elite x10", "category_slug": "papel-higienico", "brand": "Elite", "price_reference": 800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Fideos spaghetti Matarazzo 500g", "category_slug": "pastas-secas", "brand": "Matarazzo", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Fideos tirabuzones Matarazzo 500g", "category_slug": "pastas-secas", "brand": "Matarazzo", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Fideos mostachol Matarazzo 500g", "category_slug": "pastas-secas", "brand": "Matarazzo", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Fideos tallarines Lucchetti 500g", "category_slug": "pastas-secas", "brand": "Lucchetti", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Fideos coditos Lucchetti 500g", "category_slug": "pastas-secas", "brand": "Lucchetti", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Fideos municiones Marolio 500g", "category_slug": "pastas-secas", "brand": "Marolio", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Fideos spaghetti Don Vicente 500g", "category_slug": "pastas-secas", "brand": "Don Vicente", "price_reference": 1400, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Fideos lasagna Matarazzo 500g", "category_slug": "pastas-secas", "brand": "Matarazzo", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Fideos cinta Favorita 500g", "category_slug": "pastas-secas", "brand": "Favorita", "price_reference": 1300, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Ravioles de verdura La Saltenisima x24", "category_slug": "pastas-secas", "brand": "La Saltenisima", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Crema de enjuague Pantene restauracion 400ml", "category_slug": "perfumeria", "brand": "Pantene", "price_reference": 5000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Crema corporal Hinds clasica 250ml", "category_slug": "perfumeria", "brand": "Hinds", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Protector solar Dermagloss FPS30 250ml", "category_slug": "perfumeria", "brand": "Dermagloss", "price_reference": 5500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Pasta dental Colgate Triple Accion 90g", "category_slug": "perfumeria", "brand": "Colgate", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Pasta dental Colgate Luminous White 90g", "category_slug": "perfumeria", "brand": "Colgate", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Cepillo de dientes Colgate Premier", "category_slug": "perfumeria", "brand": "Colgate", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Toallitas femeninas Always clasica x8", "category_slug": "perfumeria", "brand": "Always", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Tampones OB regular x20", "category_slug": "perfumeria", "brand": "OB", "price_reference": 4000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Algodon Estrella 75g", "category_slug": "perfumeria", "brand": "Estrella", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Queso cremoso La Serenisima horma 1kg", "category_slug": "quesos-manteca", "brand": "La Serenisima", "price_reference": 12000, "unit": "kg", "sku_prefix": "ALMAC"},
  {"name": "Queso sardo Sancor 1kg", "category_slug": "quesos-manteca", "brand": "Sancor", "price_reference": 15000, "unit": "kg", "sku_prefix": "ALMAC"},
  {"name": "Queso por salut La Serenisima 1kg", "category_slug": "quesos-manteca", "brand": "La Serenisima", "price_reference": 13000, "unit": "kg", "sku_prefix": "ALMAC"},
  {"name": "Queso rallado La Serenisima 150g", "category_slug": "quesos-manteca", "brand": "La Serenisima", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Queso untable Mendicrim 300g", "category_slug": "quesos-manteca", "brand": "Mendicrim", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Queso untable Casancrem 300g", "category_slug": "quesos-manteca", "brand": "Casancrem", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Manteca Sancor 200g", "category_slug": "quesos-manteca", "brand": "Sancor", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Queso crema Finlandia light 300g", "category_slug": "quesos-manteca", "brand": "Finlandia", "price_reference": 3800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Shampoo Head and Shoulders control caspa 375ml", "category_slug": "shampoo-acondicionador", "brand": "Head & Shoulders", "price_reference": 5500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Shampoo Pantene restauracion 400ml", "category_slug": "shampoo-acondicionador", "brand": "Pantene", "price_reference": 5000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Shampoo Dove reconstruccion 400ml", "category_slug": "shampoo-acondicionador", "brand": "Dove", "price_reference": 5000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Acondicionador Pantene restauracion 400ml", "category_slug": "shampoo-acondicionador", "brand": "Pantene", "price_reference": 5200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Acondicionador Dove reconstruccion 400ml", "category_slug": "shampoo-acondicionador", "brand": "Dove", "price_reference": 5200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Shampoo Sedal ceramidas 340ml", "category_slug": "shampoo-acondicionador", "brand": "Sedal", "price_reference": 4000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Papas fritas Lays clasicas 150g", "category_slug": "snacks-salados-alm", "brand": "Lays", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Papas fritas Lays corte americano 150g", "category_slug": "snacks-salados-alm", "brand": "Lays", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Papas fritas Pringles original 137g", "category_slug": "snacks-salados-alm", "brand": "Pringles", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Palitos salados Saladix queso 70g", "category_slug": "snacks-salados-alm", "brand": "Saladix", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Palitos salados Saladix jamon 70g", "category_slug": "snacks-salados-alm", "brand": "Saladix", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Twistos jamon 100g", "category_slug": "snacks-salados-alm", "brand": "Twistos", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Mani frito salado 500g", "category_slug": "snacks-salados-alm", "brand": null, "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Papas fritas tubo Krachitos 65g", "category_slug": "snacks-salados-alm", "brand": "Krachitos", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Chizitos Cheetos 80g", "category_slug": "snacks-salados-alm", "brand": "Cheetos", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Yogur entero La Serenisima frutilla 190g", "category_slug": "yogures", "brand": "La Serenisima", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Yogur entero La Serenisima vainilla 190g", "category_slug": "yogures", "brand": "La Serenisima", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Yogur Yogurisimo firme frutilla 190g", "category_slug": "yogures", "brand": "Yogurisimo", "price_reference": 1400, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Yogur Yogurisimo bebible durazno 1L", "category_slug": "yogures", "brand": "Yogurisimo", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Yogur Ser descremado natural 190g", "category_slug": "yogures", "brand": "Ser", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Yogur Sancor Vida vainilla 190g", "category_slug": "yogures", "brand": "Sancor", "price_reference": 1600, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Yogur Activia natural 150g", "category_slug": "yogures", "brand": "Activia", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Mayonesa Hellmanns 500g", "category_slug": "conservas-enlatados", "brand": "Hellmanns", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Ketchup Hellmanns 250g", "category_slug": "conservas-enlatados", "brand": "Hellmanns", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Mostaza Savora 250g", "category_slug": "conservas-enlatados", "brand": "Savora", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Caldo en cubos Knorr carne x12", "category_slug": "conservas-enlatados", "brand": "Knorr", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Caldo en cubos Knorr gallina x12", "category_slug": "conservas-enlatados", "brand": "Knorr", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Sopa crema Knorr zapallo 64g", "category_slug": "conservas-enlatados", "brand": "Knorr", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Yerba mate Union 1kg", "category_slug": "almacen-seco", "brand": "Union", "price_reference": 4400, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Woolite ropa fina 900ml", "category_slug": "detergentes-jabones", "brand": "Woolite", "price_reference": 5500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Quitamanchas Vanish 900ml", "category_slug": "detergentes-jabones", "brand": "Vanish", "price_reference": 5000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Arroz arborio Gallo 500g", "category_slug": "arroz-legumbres", "brand": "Gallo", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Aceite de girasol Cada Dia 900ml", "category_slug": "aceites-vinagres", "brand": "Cada Dia", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Aceite de oliva extra virgen Cocinero 250ml", "category_slug": "aceites-vinagres", "brand": "Cocinero", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Alfajor Capitanyo dulce de leche 40g", "category_slug": "alfajores-chocolates", "brand": "Capitanyo", "price_reference": 500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Alfajor Ser cereal 26g", "category_slug": "alfajores-chocolates", "brand": "Ser", "price_reference": 900, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Cafe instantaneo Nescafe clasico 100g", "category_slug": "almacen-seco", "brand": "Nescafe", "price_reference": 5500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Edulcorante Hileret x100 sobres", "category_slug": "almacen-seco", "brand": "Hileret", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Palmitos enteros La Campagnola 400g", "category_slug": "conservas-enlatados", "brand": "La Campagnola", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Queso pategras La Serenisima 1kg", "category_slug": "quesos-manteca", "brand": "La Serenisima", "price_reference": 14000, "unit": "kg", "sku_prefix": "ALMAC"},

  {"name": "Leche entera Sancor sachet 1L", "category_slug": "leches", "brand": "Sancor", "price_reference": 1700, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Lavandina Marolio 2L", "category_slug": "lavandina-desinfectantes", "brand": "Marolio", "price_reference": 1400, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Desinfectante Lysoform aerosol 360ml", "category_slug": "lavandina-desinfectantes", "brand": "Lysoform", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Limpiador inodoro Pato purific 500ml", "category_slug": "limpieza", "brand": "Pato Purific", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Desodorante de ambiente Glade aerosol 360ml", "category_slug": "limpieza", "brand": "Glade", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Jabon Odex antibacterial 90g", "category_slug": "jabones-desodorantes", "brand": "Odex", "price_reference": 1000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Jugo Baggio naranja 1L", "category_slug": "jugos-polvos", "brand": "Baggio", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Jugo Baggio manzana 1L", "category_slug": "jugos-polvos", "brand": "Baggio", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Pan de hamburguesa Bimbo x4", "category_slug": "pan-envasado", "brand": "Bimbo", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Pan de pancho Bimbo x6", "category_slug": "pan-envasado", "brand": "Bimbo", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Rollo de cocina Higienol x3", "category_slug": "papel-higienico", "brand": "Higienol", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Shampoo Suave miel y almendras 930ml", "category_slug": "shampoo-acondicionador", "brand": "Suave", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Yogur La Serenisima con cereales 157g", "category_slug": "yogures", "brand": "La Serenisima", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Yogur Yogurisimo bebible frutilla 1L", "category_slug": "yogures", "brand": "Yogurisimo", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Mayonesa Natura 500g", "category_slug": "conservas-enlatados", "brand": "Natura", "price_reference": 3800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Panales Pampers Premium talle M x36", "category_slug": "panales", "brand": "Pampers", "price_reference": 18000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Leche de almendras Silk 1L", "category_slug": "leches", "brand": "Silk", "price_reference": 4000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Rocklets confitados x40g", "category_slug": "golosinas-snacks", "brand": "Arcor", "price_reference": 1000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Fiambre de cerdo cocido Paladini 150g", "category_slug": "fiambres-embutidos", "brand": "Paladini", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Agua Jane 2L", "category_slug": "lavandina-desinfectantes", "brand": "Jane", "price_reference": 1800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Bebida deportiva Gatorade 500ml", "category_slug": "bebidas-almacen", "brand": "Gatorade", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Bebida energizante Speed 250ml", "category_slug": "bebidas-almacen", "brand": "Speed", "price_reference": 2500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Agua tonica Schweppes 1.5L", "category_slug": "bebidas-almacen", "brand": "Schweppes", "price_reference": 2800, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Queso Ilolay cremoso 1kg", "category_slug": "quesos-manteca", "brand": "Ilolay", "price_reference": 11500, "unit": "kg", "sku_prefix": "ALMAC"},

  {"name": "Dulce de leche Sancor 400g", "category_slug": "lacteos-frescos", "brand": "Sancor", "price_reference": 3000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Pan lactal Fargo integral 350g", "category_slug": "pan-envasado", "brand": "Fargo", "price_reference": 3200, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Papel higienico Voila triple hoja 30m x4", "category_slug": "papel-higienico", "brand": "Voila", "price_reference": 5500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Shampoo Tresemme keratina 750ml", "category_slug": "shampoo-acondicionador", "brand": "Tresemme", "price_reference": 5500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Repelente Off verde 200ml", "category_slug": "perfumeria", "brand": "Off", "price_reference": 4000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Leche Milkaut entera larga vida 1L", "category_slug": "leches", "brand": "Milkaut", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Yogur Danone natural 190g", "category_slug": "yogures", "brand": "Danone", "price_reference": 1500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Papas fritas Pringles sour cream 137g", "category_slug": "snacks-salados-alm", "brand": "Pringles", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Panales Nevax confort talle G x24", "category_slug": "panales", "brand": "Nevax", "price_reference": 8500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Jugo en polvo Tang pomelo rosado 18g", "category_slug": "jugos-polvos", "brand": "Tang", "price_reference": 500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Salchicha tipo frankfurt Paladini x12", "category_slug": "fiambres-embutidos", "brand": "Paladini", "price_reference": 4500, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Crema de leche Sancor 200ml", "category_slug": "lacteos-frescos", "brand": "Sancor", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Caramelos Flynn Paff tutti frutti x60", "category_slug": "golosinas-snacks", "brand": "Arcor", "price_reference": 2000, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Lavandina Ayudin floral 2L", "category_slug": "lavandina-desinfectantes", "brand": "Ayudin", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Pate de foie Cabalin 90g", "category_slug": "fiambreria", "brand": "Cabalin", "price_reference": 1200, "unit": "unidad", "sku_prefix": "ALMAC"},

  {"name": "Bebida isotonica Powerade mountain blast 500ml", "category_slug": "bebidas-almacen", "brand": "Powerade", "price_reference": 2200, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Shampoo Plusbelle clasico 1L", "category_slug": "shampoo-acondicionador", "brand": "Plusbelle", "price_reference": 3500, "unit": "unidad", "sku_prefix": "ALMAC"},
  {"name": "Acondicionador Sedal ceramidas 340ml", "category_slug": "shampoo-acondicionador", "brand": "Sedal", "price_reference": 4200, "unit": "unidad", "sku_prefix": "ALMAC"}
]'::jsonb,
updated_at = CURRENT_TIMESTAMP
WHERE business_type_id = (SELECT id FROM business_types WHERE code = 'almacen') AND is_default = true;
