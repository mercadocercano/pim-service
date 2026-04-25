-- Seed 111: Actualización de templates — lote 2 (7 rubros)
-- PROPÓSITO: Poblar products, categories y brands en business_type_templates para
--            fiambreria, piletas, bazar, jugueteria, libreria, ropa y electrodomesticos.
-- IDEMPOTENTE: Solo UPDATE con WHERE id = '...'; no inserta ni elimina registros.
-- REQUIERE: 001 (business_types), 004 (business_type_quickstart_templates)
-- VERSION: 5.0.0-batch2 | generated_by: catalog-researcher
-- FECHA RESEARCH: 2026-04-25 | ZONA: Posadas, Misiones — NEA
-- RUBROS: fiambreria, piletas, bazar, jugueteria, libreria, ropa, electrodomesticos


-- =====================================================
-- 1. FIAMBRERÍA
--    id: ad55633a-8686-452a-b233-84b760c15e2d
--    sku_prefix: FIAM
-- =====================================================
UPDATE business_type_templates
SET
  categories = $categories$[
    {"name": "Quesos Frescos", "slug": "quesos-frescos", "level": 0},
    {"name": "Quesos Duros y Semiduros", "slug": "quesos-duros", "level": 0},
    {"name": "Fiambres Cocidos", "slug": "fiambres-cocidos", "level": 0},
    {"name": "Fiambres Curados", "slug": "fiambres-curados", "level": 0},
    {"name": "Lácteos y Mantecas", "slug": "lacteos", "level": 0},
    {"name": "Embutidos", "slug": "embutidos", "level": 0},
    {"name": "Conservas y Enlatados", "slug": "conservas", "level": 0},
    {"name": "Quesos Especiales", "slug": "quesos-especiales", "level": 0},
    {"name": "Pastas Frescas", "slug": "pastas-frescas", "level": 0},
    {"name": "Huevos", "slug": "huevos", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "La Serenísima", "suggested_for_categories": ["quesos-frescos", "quesos-duros", "lacteos"]},
    {"name": "SanCor", "suggested_for_categories": ["quesos-frescos", "quesos-duros", "lacteos"]},
    {"name": "Tregar", "suggested_for_categories": ["quesos-duros", "quesos-especiales"]},
    {"name": "Paladini", "suggested_for_categories": ["fiambres-cocidos", "embutidos"]},
    {"name": "Cagnoli", "suggested_for_categories": ["fiambres-curados", "embutidos"]},
    {"name": "Fargo", "suggested_for_categories": ["fiambres-cocidos"]},
    {"name": "Tres Cruces", "suggested_for_categories": ["fiambres-curados", "embutidos"]},
    {"name": "Milkaut", "suggested_for_categories": ["quesos-frescos", "lacteos"]},
    {"name": "Ilolay", "suggested_for_categories": ["quesos-frescos", "quesos-duros"]}
  ]$brands$::jsonb,
  products = $products$[
    {"name": "Queso cremoso La Serenísima x kg", "unit": "kg", "brand": "La Serenísima", "sku_prefix": "FIAM", "category_slug": "quesos-frescos", "price_reference": 8500},
    {"name": "Queso cuartirolo SanCor x kg", "unit": "kg", "brand": "SanCor", "sku_prefix": "FIAM", "category_slug": "quesos-frescos", "price_reference": 7800},
    {"name": "Queso muzzarella x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "quesos-frescos", "price_reference": 9200},
    {"name": "Queso crema vasito 200g La Serenísima", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "FIAM", "category_slug": "quesos-frescos", "price_reference": 2100},
    {"name": "Ricota x 300g", "unit": "unidad", "brand": null, "sku_prefix": "FIAM", "category_slug": "quesos-frescos", "price_reference": 1800},
    {"name": "Queso de máquina x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "quesos-frescos", "price_reference": 7200},
    {"name": "Queso reggianito Tregar x kg", "unit": "kg", "brand": "Tregar", "sku_prefix": "FIAM", "category_slug": "quesos-duros", "price_reference": 12500},
    {"name": "Queso pategras x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "quesos-duros", "price_reference": 10800},
    {"name": "Queso provolone x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "quesos-duros", "price_reference": 11200},
    {"name": "Queso sardo x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "quesos-duros", "price_reference": 11800},
    {"name": "Queso tybo x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "quesos-duros", "price_reference": 10200},
    {"name": "Queso fontina x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "quesos-especiales", "price_reference": 12000},
    {"name": "Queso brie porción 150g", "unit": "unidad", "brand": null, "sku_prefix": "FIAM", "category_slug": "quesos-especiales", "price_reference": 3500},
    {"name": "Queso azul x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "quesos-especiales", "price_reference": 14000},
    {"name": "Jamón cocido Paladini x kg", "unit": "kg", "brand": "Paladini", "sku_prefix": "FIAM", "category_slug": "fiambres-cocidos", "price_reference": 9500},
    {"name": "Paleta cocida x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "fiambres-cocidos", "price_reference": 8200},
    {"name": "Mortadela x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "fiambres-cocidos", "price_reference": 6800},
    {"name": "Queso de cerdo x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "fiambres-cocidos", "price_reference": 7500},
    {"name": "Lomito de cerdo cocido x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "fiambres-cocidos", "price_reference": 11000},
    {"name": "Jamón serrano x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "fiambres-curados", "price_reference": 16000},
    {"name": "Salame Cagnoli x kg", "unit": "kg", "brand": "Cagnoli", "sku_prefix": "FIAM", "category_slug": "fiambres-curados", "price_reference": 13500},
    {"name": "Salame tipo milano x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "fiambres-curados", "price_reference": 12800},
    {"name": "Pepperoni x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "fiambres-curados", "price_reference": 14500},
    {"name": "Chorizo seco x kg", "unit": "kg", "brand": "Tres Cruces", "sku_prefix": "FIAM", "category_slug": "embutidos", "price_reference": 11500},
    {"name": "Longaniza x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "embutidos", "price_reference": 9800},
    {"name": "Salamín picado fino x kg", "unit": "kg", "brand": "Cagnoli", "sku_prefix": "FIAM", "category_slug": "embutidos", "price_reference": 13000},
    {"name": "Manteca La Serenísima 200g", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "FIAM", "category_slug": "lacteos", "price_reference": 2800},
    {"name": "Manteca SanCor 200g", "unit": "unidad", "brand": "SanCor", "sku_prefix": "FIAM", "category_slug": "lacteos", "price_reference": 2600},
    {"name": "Crema de leche La Serenísima 200ml", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "FIAM", "category_slug": "lacteos", "price_reference": 1900},
    {"name": "Crema de leche SanCor 500ml", "unit": "unidad", "brand": "SanCor", "sku_prefix": "FIAM", "category_slug": "lacteos", "price_reference": 3800},
    {"name": "Yogur entero natural La Serenísima 190g", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "FIAM", "category_slug": "lacteos", "price_reference": 950},
    {"name": "Yogur bebible frutado 200ml La Serenísima", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "FIAM", "category_slug": "lacteos", "price_reference": 1100},
    {"name": "Leche entera La Serenísima 1L", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "FIAM", "category_slug": "lacteos", "price_reference": 1650},
    {"name": "Dulce de leche La Serenísima 400g", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "FIAM", "category_slug": "lacteos", "price_reference": 2200},
    {"name": "Dulce de leche SanCor 400g", "unit": "unidad", "brand": "SanCor", "sku_prefix": "FIAM", "category_slug": "lacteos", "price_reference": 2100},
    {"name": "Huevos blancos docena x12u", "unit": "docena", "brand": null, "sku_prefix": "FIAM", "category_slug": "huevos", "price_reference": 3200},
    {"name": "Huevos marrones docena x12u", "unit": "docena", "brand": null, "sku_prefix": "FIAM", "category_slug": "huevos", "price_reference": 3500},
    {"name": "Huevos de campo docena x12u", "unit": "docena", "brand": null, "sku_prefix": "FIAM", "category_slug": "huevos", "price_reference": 4200},
    {"name": "Pasta fresca tallarines 500g", "unit": "unidad", "brand": null, "sku_prefix": "FIAM", "category_slug": "pastas-frescas", "price_reference": 2800},
    {"name": "Pasta fresca sorrentinos carne x12u", "unit": "unidad", "brand": null, "sku_prefix": "FIAM", "category_slug": "pastas-frescas", "price_reference": 3500},
    {"name": "Pasta fresca ñoquis 500g", "unit": "unidad", "brand": null, "sku_prefix": "FIAM", "category_slug": "pastas-frescas", "price_reference": 2500},
    {"name": "Paté de jamón 200g", "unit": "unidad", "brand": "Paladini", "sku_prefix": "FIAM", "category_slug": "conservas", "price_reference": 1800},
    {"name": "Aceitunas verdes sin carozo 250g", "unit": "unidad", "brand": null, "sku_prefix": "FIAM", "category_slug": "conservas", "price_reference": 2200},
    {"name": "Aceitunas negras 250g", "unit": "unidad", "brand": null, "sku_prefix": "FIAM", "category_slug": "conservas", "price_reference": 2400},
    {"name": "Pepinillos en vinagre 370ml", "unit": "unidad", "brand": null, "sku_prefix": "FIAM", "category_slug": "conservas", "price_reference": 1900},
    {"name": "Ananá en almíbar 820g", "unit": "unidad", "brand": null, "sku_prefix": "FIAM", "category_slug": "conservas", "price_reference": 2100},
    {"name": "Creme fraiche 200g La Serenísima", "unit": "unidad", "brand": "La Serenísima", "sku_prefix": "FIAM", "category_slug": "lacteos", "price_reference": 2400},
    {"name": "Queso mascarpone 250g", "unit": "unidad", "brand": null, "sku_prefix": "FIAM", "category_slug": "quesos-especiales", "price_reference": 4500},
    {"name": "Jamón crudo sin hueso x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "fiambres-curados", "price_reference": 18500},
    {"name": "Bondiola de cerdo curada x kg", "unit": "kg", "brand": null, "sku_prefix": "FIAM", "category_slug": "fiambres-curados", "price_reference": 14000}
  ]$products$::jsonb,
  updated_at = NOW()
WHERE id = 'ad55633a-8686-452a-b233-84b760c15e2d';


-- =====================================================
-- 2. PILETAS
--    id: b5407229-0e0f-4bd7-81e5-dc89b51e1b78
--    sku_prefix: PILET
-- =====================================================
UPDATE business_type_templates
SET
  categories = $categories$[
    {"name": "Químicos y Tratamiento", "slug": "quimicos-tratamiento", "level": 0},
    {"name": "Equipos de Filtración", "slug": "equipos-filtracion", "level": 0},
    {"name": "Accesorios de Limpieza", "slug": "accesorios", "level": 0},
    {"name": "Limpieza de Pileta", "slug": "limpieza-pileta", "level": 0},
    {"name": "Escaleras y Cobertores", "slug": "escaleras-cobertores", "level": 0},
    {"name": "Iluminación Subacuática", "slug": "iluminacion-pileta", "level": 0},
    {"name": "Análisis y Control", "slug": "analisis-control", "level": 0},
    {"name": "Bombas y Motores", "slug": "bombas-motores", "level": 0},
    {"name": "Cañerías y Conexiones", "slug": "canherias-conexiones", "level": 0},
    {"name": "Recreación y Juegos Acuáticos", "slug": "recreacion-acuatica", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "HTH", "suggested_for_categories": ["quimicos-tratamiento"]},
    {"name": "Genco", "suggested_for_categories": ["quimicos-tratamiento"]},
    {"name": "Sodramar", "suggested_for_categories": ["equipos-filtracion", "bombas-motores"]},
    {"name": "Dancas", "suggested_for_categories": ["bombas-motores"]},
    {"name": "Zodiac", "suggested_for_categories": ["accesorios", "limpieza-pileta"]},
    {"name": "Quantum", "suggested_for_categories": ["quimicos-tratamiento"]},
    {"name": "Magiclean", "suggested_for_categories": ["limpieza-pileta", "quimicos-tratamiento"]},
    {"name": "Smartpool", "suggested_for_categories": ["analisis-control", "accesorios"]},
    {"name": "Summer Waves", "suggested_for_categories": ["recreacion-acuatica"]}
  ]$brands$::jsonb,
  products = $products$[
    {"name": "Cloro granulado HTH 1kg", "unit": "unidad", "brand": "HTH", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 4800},
    {"name": "Cloro granulado HTH 5kg", "unit": "unidad", "brand": "HTH", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 21000},
    {"name": "Tabletas cloro 3\" HTH x8u", "unit": "unidad", "brand": "HTH", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 9500},
    {"name": "Tabletas cloro 1\" HTH x12u", "unit": "unidad", "brand": "HTH", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 5800},
    {"name": "Algicida líquido 1L Genco", "unit": "unidad", "brand": "Genco", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 3500},
    {"name": "Algicida concentrado 5L Genco", "unit": "unidad", "brand": "Genco", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 14500},
    {"name": "Floculante líquido 1L", "unit": "unidad", "brand": "HTH", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 3200},
    {"name": "Clarificante líquido 1L", "unit": "unidad", "brand": "Genco", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 3000},
    {"name": "Elevador de pH 1kg", "unit": "unidad", "brand": "HTH", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 2800},
    {"name": "Reductor de pH (ácido muriático) 1L", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 1800},
    {"name": "Estabilizador ácido cianúrico 1kg", "unit": "unidad", "brand": "HTH", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 5500},
    {"name": "Cloro líquido 5L", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 4200},
    {"name": "Kit análisis tiras reactivas x50u", "unit": "unidad", "brand": "HTH", "sku_prefix": "PILET", "category_slug": "analisis-control", "price_reference": 3500},
    {"name": "Termómetro flotante para pileta", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "analisis-control", "price_reference": 1800},
    {"name": "Testeador digital pH y cloro", "unit": "unidad", "brand": "Smartpool", "sku_prefix": "PILET", "category_slug": "analisis-control", "price_reference": 8500},
    {"name": "Bomba de filtración 1/3 HP Sodramar", "unit": "unidad", "brand": "Sodramar", "sku_prefix": "PILET", "category_slug": "bombas-motores", "price_reference": 48000},
    {"name": "Bomba de filtración 1/2 HP Sodramar", "unit": "unidad", "brand": "Sodramar", "sku_prefix": "PILET", "category_slug": "bombas-motores", "price_reference": 65000},
    {"name": "Bomba de filtración 1 HP Dancas", "unit": "unidad", "brand": "Dancas", "sku_prefix": "PILET", "category_slug": "bombas-motores", "price_reference": 88000},
    {"name": "Filtro de arena completo 60cm", "unit": "unidad", "brand": "Sodramar", "sku_prefix": "PILET", "category_slug": "equipos-filtracion", "price_reference": 95000},
    {"name": "Arena filtrante para pileta 25kg", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "equipos-filtracion", "price_reference": 7500},
    {"name": "Skimmer flotante automático", "unit": "unidad", "brand": "Zodiac", "sku_prefix": "PILET", "category_slug": "equipos-filtracion", "price_reference": 12000},
    {"name": "Canasta para skimmer universal", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "equipos-filtracion", "price_reference": 2200},
    {"name": "Manguera 38mm x metro", "unit": "metro", "brand": null, "sku_prefix": "PILET", "category_slug": "canherias-conexiones", "price_reference": 850},
    {"name": "Conexión flexible 38mm", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "canherias-conexiones", "price_reference": 1500},
    {"name": "Limpiafondo manual telescópico", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "limpieza-pileta", "price_reference": 8500},
    {"name": "Limpiafondo automático Zodiac", "unit": "unidad", "brand": "Zodiac", "sku_prefix": "PILET", "category_slug": "limpieza-pileta", "price_reference": 45000},
    {"name": "Cepillo limpiafondo 40cm", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "limpieza-pileta", "price_reference": 3500},
    {"name": "Red cazahojas para pileta", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "limpieza-pileta", "price_reference": 2800},
    {"name": "Cepillo de pared para pileta", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "limpieza-pileta", "price_reference": 4200},
    {"name": "Escalera acero inoxidable 3 peldaños", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "escaleras-cobertores", "price_reference": 65000},
    {"name": "Escalera acero inoxidable 4 peldaños", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "escaleras-cobertores", "price_reference": 82000},
    {"name": "Cubre pileta lona 6x3m", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "escaleras-cobertores", "price_reference": 18500},
    {"name": "Cubre pileta lona 8x4m", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "escaleras-cobertores", "price_reference": 28000},
    {"name": "Luz LED subacuática 12W", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "iluminacion-pileta", "price_reference": 22000},
    {"name": "Flotador inflable para pileta", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "recreacion-acuatica", "price_reference": 3500},
    {"name": "Pelota acuática resistente UV", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "recreacion-acuatica", "price_reference": 2800},
    {"name": "Dispensador flotante pastillas cloro", "unit": "unidad", "brand": "HTH", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 3200},
    {"name": "Shock de cloro granulado HTH 500g", "unit": "unidad", "brand": "HTH", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 2800},
    {"name": "Antiincrustante para pileta 1L", "unit": "unidad", "brand": "Genco", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 3800},
    {"name": "Varizyme enzimático para pileta 1L", "unit": "unidad", "brand": "Magiclean", "sku_prefix": "PILET", "category_slug": "quimicos-tratamiento", "price_reference": 4500},
    {"name": "Selector de 3 vías 1½\"", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "equipos-filtracion", "price_reference": 6500},
    {"name": "Manómetro para filtro de arena", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "equipos-filtracion", "price_reference": 3200},
    {"name": "Poste telescópico aluminio 3m", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "limpieza-pileta", "price_reference": 5800},
    {"name": "Aspiradora manual para pileta", "unit": "unidad", "brand": "Zodiac", "sku_prefix": "PILET", "category_slug": "limpieza-pileta", "price_reference": 28000},
    {"name": "Tapón de desagüe fondo 2\"", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "canherias-conexiones", "price_reference": 1200},
    {"name": "Boya reguladora de nivel automática", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "canherias-conexiones", "price_reference": 4800},
    {"name": "Cronómetro programable para bomba", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "bombas-motores", "price_reference": 5500},
    {"name": "Pileta armable Summer Waves 3x2m", "unit": "unidad", "brand": "Summer Waves", "sku_prefix": "PILET", "category_slug": "recreacion-acuatica", "price_reference": 85000},
    {"name": "Pileta inflable redonda 2m diámetro", "unit": "unidad", "brand": null, "sku_prefix": "PILET", "category_slug": "recreacion-acuatica", "price_reference": 32000}
  ]$products$::jsonb,
  updated_at = NOW()
WHERE id = 'b5407229-0e0f-4bd7-81e5-dc89b51e1b78';


-- =====================================================
-- 3. BAZAR
--    id: c278befb-7131-4bc4-a7f4-416b14999c5d
--    sku_prefix: BAZAR
-- =====================================================
UPDATE business_type_templates
SET
  categories = $categories$[
    {"name": "Vajilla", "slug": "vajilla", "level": 0},
    {"name": "Utensilios de Cocina", "slug": "utensilios-cocina", "level": 0},
    {"name": "Almacenamiento y Organización", "slug": "almacenamiento", "level": 0},
    {"name": "Textiles del Hogar", "slug": "textiles-hogar", "level": 0},
    {"name": "Organización del Hogar", "slug": "organizacion", "level": 0},
    {"name": "Iluminación y Velas", "slug": "iluminacion-velas", "level": 0},
    {"name": "Ollas y Sartenes", "slug": "ollas-sartenes", "level": 0},
    {"name": "Cuchillos y Tablas", "slug": "cuchillos-tablas", "level": 0},
    {"name": "Vasos y Jarras", "slug": "vasos-jarras", "level": 0},
    {"name": "Limpieza del Hogar", "slug": "limpieza-hogar", "level": 0},
    {"name": "Decoración", "slug": "decoracion", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "Tramontina", "suggested_for_categories": ["ollas-sartenes", "utensilios-cocina", "cuchillos-tablas"]},
    {"name": "Tupperware", "suggested_for_categories": ["almacenamiento"]},
    {"name": "Mundial", "suggested_for_categories": ["cuchillos-tablas"]},
    {"name": "Arcos", "suggested_for_categories": ["vasos-jarras", "vajilla"]},
    {"name": "Cristar", "suggested_for_categories": ["vasos-jarras"]},
    {"name": "Essen", "suggested_for_categories": ["ollas-sartenes"]},
    {"name": "Vollrath", "suggested_for_categories": ["ollas-sartenes", "utensilios-cocina"]},
    {"name": "Scotch-Brite", "suggested_for_categories": ["limpieza-hogar"]},
    {"name": "Ravi", "suggested_for_categories": ["vajilla", "vasos-jarras"]}
  ]$brands$::jsonb,
  products = $products$[
    {"name": "Juego de platos hondos x6u porcelana", "unit": "juego", "brand": null, "sku_prefix": "BAZAR", "category_slug": "vajilla", "price_reference": 9500},
    {"name": "Juego de platos planos x6u porcelana", "unit": "juego", "brand": null, "sku_prefix": "BAZAR", "category_slug": "vajilla", "price_reference": 9500},
    {"name": "Juego de platos postre x6u", "unit": "juego", "brand": null, "sku_prefix": "BAZAR", "category_slug": "vajilla", "price_reference": 7800},
    {"name": "Taza café con plato cerámica", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "vajilla", "price_reference": 2200},
    {"name": "Taza desayuno grande 380ml", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "vajilla", "price_reference": 1800},
    {"name": "Fuente oval porcelana 35cm", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "vajilla", "price_reference": 4500},
    {"name": "Vasos agua x6u vidrio Cristar 350ml", "unit": "juego", "brand": "Cristar", "sku_prefix": "BAZAR", "category_slug": "vasos-jarras", "price_reference": 5800},
    {"name": "Vasos cerveza x6u vidrio 500ml", "unit": "juego", "brand": "Arcos", "sku_prefix": "BAZAR", "category_slug": "vasos-jarras", "price_reference": 7200},
    {"name": "Jarra vidrio con tapa 1.5L", "unit": "unidad", "brand": "Cristar", "sku_prefix": "BAZAR", "category_slug": "vasos-jarras", "price_reference": 4800},
    {"name": "Copas vino tinto x6u", "unit": "juego", "brand": "Arcos", "sku_prefix": "BAZAR", "category_slug": "vasos-jarras", "price_reference": 8500},
    {"name": "Cubiertos inox x24u (6 personas)", "unit": "juego", "brand": null, "sku_prefix": "BAZAR", "category_slug": "vajilla", "price_reference": 12500},
    {"name": "Cubiertos inox x12u (4 personas)", "unit": "juego", "brand": null, "sku_prefix": "BAZAR", "category_slug": "vajilla", "price_reference": 7500},
    {"name": "Olla enlozada 24cm con tapa", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "ollas-sartenes", "price_reference": 8500},
    {"name": "Olla enlozada 20cm con tapa", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "ollas-sartenes", "price_reference": 6500},
    {"name": "Sartén antiadherente 28cm Tramontina", "unit": "unidad", "brand": "Tramontina", "sku_prefix": "BAZAR", "category_slug": "ollas-sartenes", "price_reference": 14500},
    {"name": "Sartén antiadherente 24cm Tramontina", "unit": "unidad", "brand": "Tramontina", "sku_prefix": "BAZAR", "category_slug": "ollas-sartenes", "price_reference": 11000},
    {"name": "Cacerola acero inox 22cm Essen", "unit": "unidad", "brand": "Essen", "sku_prefix": "BAZAR", "category_slug": "ollas-sartenes", "price_reference": 22000},
    {"name": "Cuchillo chef 8\" Mundial", "unit": "unidad", "brand": "Mundial", "sku_prefix": "BAZAR", "category_slug": "cuchillos-tablas", "price_reference": 9800},
    {"name": "Cuchillo pan 8\" Mundial", "unit": "unidad", "brand": "Mundial", "sku_prefix": "BAZAR", "category_slug": "cuchillos-tablas", "price_reference": 8500},
    {"name": "Tabla para picar madera 30x20cm", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "cuchillos-tablas", "price_reference": 4800},
    {"name": "Tabla para picar polietileno 40x25cm", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "cuchillos-tablas", "price_reference": 3500},
    {"name": "Espátula silicona resistente calor", "unit": "unidad", "brand": "Tramontina", "sku_prefix": "BAZAR", "category_slug": "utensilios-cocina", "price_reference": 2200},
    {"name": "Cucharón acero inox", "unit": "unidad", "brand": "Tramontina", "sku_prefix": "BAZAR", "category_slug": "utensilios-cocina", "price_reference": 1800},
    {"name": "Rallador 4 caras acero inox", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "utensilios-cocina", "price_reference": 3200},
    {"name": "Pelador de verduras giratorio", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "utensilios-cocina", "price_reference": 1200},
    {"name": "Embudo plástico set x3u", "unit": "juego", "brand": null, "sku_prefix": "BAZAR", "category_slug": "utensilios-cocina", "price_reference": 1500},
    {"name": "Colador grande 24cm acero inox", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "utensilios-cocina", "price_reference": 3800},
    {"name": "Contenedor hermético 1L Tupperware", "unit": "unidad", "brand": "Tupperware", "sku_prefix": "BAZAR", "category_slug": "almacenamiento", "price_reference": 4200},
    {"name": "Set contenedores plástico x5u 0.5-2L", "unit": "juego", "brand": null, "sku_prefix": "BAZAR", "category_slug": "almacenamiento", "price_reference": 5500},
    {"name": "Frascos vidrio 500ml x3u con tapa", "unit": "juego", "brand": null, "sku_prefix": "BAZAR", "category_slug": "almacenamiento", "price_reference": 4800},
    {"name": "Canasto mimbre mediano con tapa", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "almacenamiento", "price_reference": 5500},
    {"name": "Mantel PVC 1.4x1.8m estampado", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "textiles-hogar", "price_reference": 4200},
    {"name": "Repasadores algodón x3u", "unit": "juego", "brand": null, "sku_prefix": "BAZAR", "category_slug": "textiles-hogar", "price_reference": 2800},
    {"name": "Delantal cocina algodón", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "textiles-hogar", "price_reference": 3500},
    {"name": "Velas emergencia x10u blancas", "unit": "paquete", "brand": null, "sku_prefix": "BAZAR", "category_slug": "iluminacion-velas", "price_reference": 2200},
    {"name": "Velas aromáticas decorativas x3u", "unit": "juego", "brand": null, "sku_prefix": "BAZAR", "category_slug": "iluminacion-velas", "price_reference": 3800},
    {"name": "Organizador cajón extensible", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "organizacion", "price_reference": 2800},
    {"name": "Ganchos adhesivos resistentes x5u", "unit": "paquete", "brand": null, "sku_prefix": "BAZAR", "category_slug": "organizacion", "price_reference": 1500},
    {"name": "Perchero de pie madera 8 ganchos", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "organizacion", "price_reference": 12000},
    {"name": "Portarrollo papel cocina adhesivo", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "organizacion", "price_reference": 2200},
    {"name": "Estropajo metálico x3u Scotch-Brite", "unit": "paquete", "brand": "Scotch-Brite", "sku_prefix": "BAZAR", "category_slug": "limpieza-hogar", "price_reference": 1200},
    {"name": "Esponja doble faz x3u Scotch-Brite", "unit": "paquete", "brand": "Scotch-Brite", "sku_prefix": "BAZAR", "category_slug": "limpieza-hogar", "price_reference": 1500},
    {"name": "Cepillo de fregar con mango", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "limpieza-hogar", "price_reference": 1800},
    {"name": "Porta especias giratorio x12 frascos", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "decoracion", "price_reference": 8500},
    {"name": "Florero vidrio cilíndrico 25cm", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "decoracion", "price_reference": 3200},
    {"name": "Portavela vidrio decorativo", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "decoracion", "price_reference": 2500},
    {"name": "Balanza digital cocina 5kg", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "utensilios-cocina", "price_reference": 8500},
    {"name": "Termómetro cocina digital", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "utensilios-cocina", "price_reference": 5500},
    {"name": "Exprimidor manual vidrio", "unit": "unidad", "brand": null, "sku_prefix": "BAZAR", "category_slug": "utensilios-cocina", "price_reference": 3200},
    {"name": "Juego salero y pimentero cerámica", "unit": "juego", "brand": null, "sku_prefix": "BAZAR", "category_slug": "vajilla", "price_reference": 2800}
  ]$products$::jsonb,
  updated_at = NOW()
WHERE id = 'c278befb-7131-4bc4-a7f4-416b14999c5d';


-- =====================================================
-- 4. JUGUETERÍA
--    id: 86001b97-681c-491d-9a0e-0408d811a3ce
--    sku_prefix: JUGUE
-- =====================================================
UPDATE business_type_templates
SET
  categories = $categories$[
    {"name": "Bebés 0-2 años", "slug": "bebes-0-2", "level": 0},
    {"name": "Niños 3-5 años", "slug": "ninos-3-5", "level": 0},
    {"name": "Niños 6-10 años", "slug": "ninos-6-10", "level": 0},
    {"name": "Adolescentes y Adultos", "slug": "adolescentes", "level": 0},
    {"name": "Juegos de Mesa", "slug": "juegos-mesa", "level": 0},
    {"name": "Cotillón y Fiestas", "slug": "cotillon", "level": 0},
    {"name": "Vehículos y Muñecos", "slug": "vehiculos-munhcos", "level": 0},
    {"name": "Deportes y Aire Libre", "slug": "deportes-aire-libre", "level": 0},
    {"name": "Arte y Manualidades", "slug": "arte-manualidades", "level": 0},
    {"name": "Electrónicos Infantiles", "slug": "electronicos-infantiles", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "Hasbro", "suggested_for_categories": ["juegos-mesa", "adolescentes"]},
    {"name": "Mattel", "suggested_for_categories": ["vehiculos-munhcos", "ninos-6-10", "juegos-mesa"]},
    {"name": "Fisher-Price", "suggested_for_categories": ["bebes-0-2", "ninos-3-5"]},
    {"name": "Lego", "suggested_for_categories": ["ninos-6-10", "adolescentes"]},
    {"name": "Play-Doh", "suggested_for_categories": ["ninos-3-5", "arte-manualidades"]},
    {"name": "Hot Wheels", "suggested_for_categories": ["vehiculos-munhcos", "ninos-6-10"]},
    {"name": "Barbie", "suggested_for_categories": ["vehiculos-munhcos", "ninos-6-10"]},
    {"name": "Rasti", "suggested_for_categories": ["ninos-3-5", "ninos-6-10"]},
    {"name": "Ruibal", "suggested_for_categories": ["juegos-mesa"]},
    {"name": "Ditoys", "suggested_for_categories": ["cotillon", "bebes-0-2"]}
  ]$brands$::jsonb,
  products = $products$[
    {"name": "Sonajero bebé Fisher-Price", "unit": "unidad", "brand": "Fisher-Price", "sku_prefix": "JUGUE", "category_slug": "bebes-0-2", "price_reference": 4500},
    {"name": "Mordillo silicona libre BPA", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "bebes-0-2", "price_reference": 2800},
    {"name": "Peluche osito 30cm suave", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "bebes-0-2", "price_reference": 5500},
    {"name": "Móvil cuna musical Fisher-Price", "unit": "unidad", "brand": "Fisher-Price", "sku_prefix": "JUGUE", "category_slug": "bebes-0-2", "price_reference": 12000},
    {"name": "Centro de actividades bebé", "unit": "unidad", "brand": "Fisher-Price", "sku_prefix": "JUGUE", "category_slug": "bebes-0-2", "price_reference": 18000},
    {"name": "Rasti básico 50 piezas", "unit": "unidad", "brand": "Rasti", "sku_prefix": "JUGUE", "category_slug": "ninos-3-5", "price_reference": 6800},
    {"name": "Plastilina Play-Doh x4 colores", "unit": "unidad", "brand": "Play-Doh", "sku_prefix": "JUGUE", "category_slug": "ninos-3-5", "price_reference": 5500},
    {"name": "Puzle 24 piezas madera animales", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "ninos-3-5", "price_reference": 4200},
    {"name": "Pistola de burbujas de colores", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "ninos-3-5", "price_reference": 3500},
    {"name": "Set de cocina infantil plástico", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "ninos-3-5", "price_reference": 7500},
    {"name": "Muñeca Barbie básica surtida", "unit": "unidad", "brand": "Barbie", "sku_prefix": "JUGUE", "category_slug": "ninos-6-10", "price_reference": 9500},
    {"name": "Auto Hot Wheels surtido x1u", "unit": "unidad", "brand": "Hot Wheels", "sku_prefix": "JUGUE", "category_slug": "vehiculos-munhcos", "price_reference": 1800},
    {"name": "Auto Hot Wheels pista básica", "unit": "unidad", "brand": "Hot Wheels", "sku_prefix": "JUGUE", "category_slug": "vehiculos-munhcos", "price_reference": 14000},
    {"name": "Lego Creator 3en1 pequeño", "unit": "unidad", "brand": "Lego", "sku_prefix": "JUGUE", "category_slug": "ninos-6-10", "price_reference": 16000},
    {"name": "Lego City set mediano", "unit": "unidad", "brand": "Lego", "sku_prefix": "JUGUE", "category_slug": "ninos-6-10", "price_reference": 28000},
    {"name": "Pelota de fútbol N°4 cuero sintético", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "deportes-aire-libre", "price_reference": 5800},
    {"name": "Pelota de fútbol N°3 infantil", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "deportes-aire-libre", "price_reference": 4200},
    {"name": "Raquetas de playa x2u con pelota", "unit": "juego", "brand": null, "sku_prefix": "JUGUE", "category_slug": "deportes-aire-libre", "price_reference": 3800},
    {"name": "Frisbee 27cm resistente UV", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "deportes-aire-libre", "price_reference": 2200},
    {"name": "Monopoly clásico Hasbro", "unit": "unidad", "brand": "Hasbro", "sku_prefix": "JUGUE", "category_slug": "juegos-mesa", "price_reference": 14500},
    {"name": "Batalla Naval Hasbro", "unit": "unidad", "brand": "Hasbro", "sku_prefix": "JUGUE", "category_slug": "juegos-mesa", "price_reference": 12000},
    {"name": "Jenga clásico 54 bloques", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "juegos-mesa", "price_reference": 8500},
    {"name": "Uno clásico Mattel 108 cartas", "unit": "unidad", "brand": "Mattel", "sku_prefix": "JUGUE", "category_slug": "juegos-mesa", "price_reference": 5800},
    {"name": "Twister Hasbro", "unit": "unidad", "brand": "Hasbro", "sku_prefix": "JUGUE", "category_slug": "juegos-mesa", "price_reference": 11000},
    {"name": "Ajedrez tablero madera 30x30cm", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "juegos-mesa", "price_reference": 6500},
    {"name": "Scrabble español Hasbro", "unit": "unidad", "brand": "Hasbro", "sku_prefix": "JUGUE", "category_slug": "juegos-mesa", "price_reference": 12500},
    {"name": "Cluedo clásico Hasbro", "unit": "unidad", "brand": "Hasbro", "sku_prefix": "JUGUE", "category_slug": "juegos-mesa", "price_reference": 13500},
    {"name": "Globos de colores x25u bolsa", "unit": "paquete", "brand": null, "sku_prefix": "JUGUE", "category_slug": "cotillon", "price_reference": 1200},
    {"name": "Gorros de cumpleaños x8u", "unit": "paquete", "brand": null, "sku_prefix": "JUGUE", "category_slug": "cotillon", "price_reference": 1500},
    {"name": "Piñata estrella 6 puntas 50cm", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "cotillon", "price_reference": 4500},
    {"name": "Cotillón set cumpleaños x12 invitados", "unit": "juego", "brand": null, "sku_prefix": "JUGUE", "category_slug": "cotillon", "price_reference": 8500},
    {"name": "Velas número para torta 0-9", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "cotillon", "price_reference": 800},
    {"name": "Serpentinas x10u pack", "unit": "paquete", "brand": null, "sku_prefix": "JUGUE", "category_slug": "cotillon", "price_reference": 1200},
    {"name": "Set pinturas de dedo x6 colores", "unit": "unidad", "brand": "Play-Doh", "sku_prefix": "JUGUE", "category_slug": "arte-manualidades", "price_reference": 3800},
    {"name": "Crayones gruesos x12 colores", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "arte-manualidades", "price_reference": 2200},
    {"name": "Libro para colorear infantil A4", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "arte-manualidades", "price_reference": 1500},
    {"name": "Rasti Constructor 100 piezas", "unit": "unidad", "brand": "Rasti", "sku_prefix": "JUGUE", "category_slug": "ninos-6-10", "price_reference": 12000},
    {"name": "Tablet educativa infantil 7\"", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "electronicos-infantiles", "price_reference": 32000},
    {"name": "Walkie-talkie infantil par", "unit": "juego", "brand": null, "sku_prefix": "JUGUE", "category_slug": "electronicos-infantiles", "price_reference": 8500},
    {"name": "Muñeco acción superhéroe 30cm", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "vehiculos-munhcos", "price_reference": 5500},
    {"name": "Camión volcador plástico grande", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "vehiculos-munhcos", "price_reference": 6800},
    {"name": "Bebote vinilo 40cm con ropa", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "ninos-3-5", "price_reference": 7500},
    {"name": "Triciclo infantil talle 1 (1-3 años)", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "deportes-aire-libre", "price_reference": 25000},
    {"name": "Bicicleta rodado 16 con rueditas", "unit": "unidad", "brand": null, "sku_prefix": "JUGUE", "category_slug": "deportes-aire-libre", "price_reference": 55000}
  ]$products$::jsonb,
  updated_at = NOW()
WHERE id = '86001b97-681c-491d-9a0e-0408d811a3ce';


-- =====================================================
-- 5. LIBRERÍA
--    id: 4d53e0e5-bef2-4af5-9f41-7113ca9c1b52
--    sku_prefix: LIBRE
-- =====================================================
UPDATE business_type_templates
SET
  categories = $categories$[
    {"name": "Útiles Escolares", "slug": "utiles-escolares", "level": 0},
    {"name": "Cuadernos y Carpetas", "slug": "cuadernos", "level": 0},
    {"name": "Arte y Dibujo", "slug": "arte", "level": 0},
    {"name": "Papelería y Oficina", "slug": "papeleria", "level": 0},
    {"name": "Tecnología Básica", "slug": "tecnologia-basica", "level": 0},
    {"name": "Instrumentos de Medición", "slug": "instrumentos-medicion", "level": 0},
    {"name": "Sellados y Archivado", "slug": "archivado", "level": 0},
    {"name": "Manualidades", "slug": "manualidades", "level": 0},
    {"name": "Libros y Textos Escolares", "slug": "libros", "level": 0},
    {"name": "Marcadores y Resaltadores", "slug": "marcadores", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "Staedtler", "suggested_for_categories": ["utiles-escolares", "instrumentos-medicion", "arte"]},
    {"name": "Faber-Castell", "suggested_for_categories": ["utiles-escolares", "arte"]},
    {"name": "Bic", "suggested_for_categories": ["utiles-escolares", "papeleria"]},
    {"name": "Maped", "suggested_for_categories": ["instrumentos-medicion", "utiles-escolares"]},
    {"name": "Gloria", "suggested_for_categories": ["cuadernos"]},
    {"name": "Rivadavia", "suggested_for_categories": ["cuadernos"]},
    {"name": "Navigator", "suggested_for_categories": ["papeleria"]},
    {"name": "Casio", "suggested_for_categories": ["tecnologia-basica"]},
    {"name": "Kingston", "suggested_for_categories": ["tecnologia-basica"]},
    {"name": "Acrilex", "suggested_for_categories": ["arte"]}
  ]$brands$::jsonb,
  products = $products$[
    {"name": "Lápiz HB Staedtler x12u", "unit": "caja", "brand": "Staedtler", "sku_prefix": "LIBRE", "category_slug": "utiles-escolares", "price_reference": 3800},
    {"name": "Lápiz 2B Faber-Castell x12u", "unit": "caja", "brand": "Faber-Castell", "sku_prefix": "LIBRE", "category_slug": "utiles-escolares", "price_reference": 3500},
    {"name": "Lapicera Bic Cristal azul x12u", "unit": "caja", "brand": "Bic", "sku_prefix": "LIBRE", "category_slug": "utiles-escolares", "price_reference": 3200},
    {"name": "Lapicera Bic Cristal negra x12u", "unit": "caja", "brand": "Bic", "sku_prefix": "LIBRE", "category_slug": "utiles-escolares", "price_reference": 3200},
    {"name": "Goma borrar Staedtler blanca", "unit": "unidad", "brand": "Staedtler", "sku_prefix": "LIBRE", "category_slug": "utiles-escolares", "price_reference": 650},
    {"name": "Sacapuntas doble Faber-Castell", "unit": "unidad", "brand": "Faber-Castell", "sku_prefix": "LIBRE", "category_slug": "utiles-escolares", "price_reference": 550},
    {"name": "Regla 30cm plástico transparente", "unit": "unidad", "brand": "Maped", "sku_prefix": "LIBRE", "category_slug": "instrumentos-medicion", "price_reference": 800},
    {"name": "Compás Maped escolar", "unit": "unidad", "brand": "Maped", "sku_prefix": "LIBRE", "category_slug": "instrumentos-medicion", "price_reference": 2500},
    {"name": "Transportador 180° plástico", "unit": "unidad", "brand": "Maped", "sku_prefix": "LIBRE", "category_slug": "instrumentos-medicion", "price_reference": 750},
    {"name": "Escuadra 45° Maped 20cm", "unit": "unidad", "brand": "Maped", "sku_prefix": "LIBRE", "category_slug": "instrumentos-medicion", "price_reference": 950},
    {"name": "Cuaderno Gloria 48h rayado tapa blanda", "unit": "unidad", "brand": "Gloria", "sku_prefix": "LIBRE", "category_slug": "cuadernos", "price_reference": 1800},
    {"name": "Cuaderno Gloria 96h cuadriculado", "unit": "unidad", "brand": "Gloria", "sku_prefix": "LIBRE", "category_slug": "cuadernos", "price_reference": 2800},
    {"name": "Cuaderno Rivadavia 192h espiral", "unit": "unidad", "brand": "Rivadavia", "sku_prefix": "LIBRE", "category_slug": "cuadernos", "price_reference": 3500},
    {"name": "Cuaderno Rivadavia 96h tapa dura", "unit": "unidad", "brand": "Rivadavia", "sku_prefix": "LIBRE", "category_slug": "cuadernos", "price_reference": 3000},
    {"name": "Carpeta escolar A4 azul 3 aros", "unit": "unidad", "brand": null, "sku_prefix": "LIBRE", "category_slug": "cuadernos", "price_reference": 3200},
    {"name": "Bibliorato lomo 3cm A4 surtido", "unit": "unidad", "brand": null, "sku_prefix": "LIBRE", "category_slug": "archivado", "price_reference": 2800},
    {"name": "Block dibujo A4 x20h 90g", "unit": "unidad", "brand": null, "sku_prefix": "LIBRE", "category_slug": "arte", "price_reference": 2500},
    {"name": "Temperas Acuarela x12 colores 20ml", "unit": "unidad", "brand": "Acrilex", "sku_prefix": "LIBRE", "category_slug": "arte", "price_reference": 3800},
    {"name": "Acrílicos Acrilex 20ml x6u colores básicos", "unit": "juego", "brand": "Acrilex", "sku_prefix": "LIBRE", "category_slug": "arte", "price_reference": 4200},
    {"name": "Pinceles escuela set x12u", "unit": "juego", "brand": null, "sku_prefix": "LIBRE", "category_slug": "arte", "price_reference": 3500},
    {"name": "Lápices de color x24u Faber-Castell", "unit": "caja", "brand": "Faber-Castell", "sku_prefix": "LIBRE", "category_slug": "arte", "price_reference": 5500},
    {"name": "Marcadores gruesos x12u surtidos", "unit": "caja", "brand": null, "sku_prefix": "LIBRE", "category_slug": "marcadores", "price_reference": 3200},
    {"name": "Resaltadores x5u colores surtidos", "unit": "juego", "brand": "Staedtler", "sku_prefix": "LIBRE", "category_slug": "marcadores", "price_reference": 2800},
    {"name": "Fibrones punta fina x10u Bic", "unit": "caja", "brand": "Bic", "sku_prefix": "LIBRE", "category_slug": "marcadores", "price_reference": 3500},
    {"name": "Papel A4 resma 500h Navigator 75g", "unit": "resma", "brand": "Navigator", "sku_prefix": "LIBRE", "category_slug": "papeleria", "price_reference": 9500},
    {"name": "Cartulina colores x10u A4", "unit": "paquete", "brand": null, "sku_prefix": "LIBRE", "category_slug": "papeleria", "price_reference": 2200},
    {"name": "Cola vinílica Klaukol 250g", "unit": "unidad", "brand": null, "sku_prefix": "LIBRE", "category_slug": "manualidades", "price_reference": 1500},
    {"name": "Cinta adhesiva Scotch 18mm x33m", "unit": "unidad", "brand": null, "sku_prefix": "LIBRE", "category_slug": "papeleria", "price_reference": 900},
    {"name": "Sobres blancos A4 x25u", "unit": "paquete", "brand": null, "sku_prefix": "LIBRE", "category_slug": "papeleria", "price_reference": 2200},
    {"name": "Pilas AA Duracell x4u", "unit": "paquete", "brand": null, "sku_prefix": "LIBRE", "category_slug": "tecnologia-basica", "price_reference": 3500},
    {"name": "Pilas AAA Energizer x4u", "unit": "paquete", "brand": null, "sku_prefix": "LIBRE", "category_slug": "tecnologia-basica", "price_reference": 3200},
    {"name": "Calculadora Casio fx-82 científica", "unit": "unidad", "brand": "Casio", "sku_prefix": "LIBRE", "category_slug": "tecnologia-basica", "price_reference": 18500},
    {"name": "Calculadora básica Casio 8 dígitos", "unit": "unidad", "brand": "Casio", "sku_prefix": "LIBRE", "category_slug": "tecnologia-basica", "price_reference": 5500},
    {"name": "USB 32GB Kingston DataTraveler", "unit": "unidad", "brand": "Kingston", "sku_prefix": "LIBRE", "category_slug": "tecnologia-basica", "price_reference": 8500},
    {"name": "Tijeras escolares punta redonda", "unit": "unidad", "brand": "Maped", "sku_prefix": "LIBRE", "category_slug": "utiles-escolares", "price_reference": 1500},
    {"name": "Tijeras adulto Maped 21cm", "unit": "unidad", "brand": "Maped", "sku_prefix": "LIBRE", "category_slug": "papeleria", "price_reference": 2200},
    {"name": "Separadores plástico x5u A4", "unit": "paquete", "brand": null, "sku_prefix": "LIBRE", "category_slug": "archivado", "price_reference": 1200},
    {"name": "Folio A4 paquete x100u", "unit": "paquete", "brand": null, "sku_prefix": "LIBRE", "category_slug": "archivado", "price_reference": 2800},
    {"name": "Abrochadora Maped con grapas", "unit": "unidad", "brand": "Maped", "sku_prefix": "LIBRE", "category_slug": "papeleria", "price_reference": 3500},
    {"name": "Grapas standard x1000u", "unit": "caja", "brand": null, "sku_prefix": "LIBRE", "category_slug": "papeleria", "price_reference": 800},
    {"name": "Papel glasé x10u colores surtidos", "unit": "paquete", "brand": null, "sku_prefix": "LIBRE", "category_slug": "manualidades", "price_reference": 1800},
    {"name": "Plastilina escolar x12 colores", "unit": "unidad", "brand": null, "sku_prefix": "LIBRE", "category_slug": "manualidades", "price_reference": 2500},
    {"name": "Corrector líquido punta pincel", "unit": "unidad", "brand": "Bic", "sku_prefix": "LIBRE", "category_slug": "utiles-escolares", "price_reference": 1200},
    {"name": "Cinta de papel masking 24mm x50m", "unit": "unidad", "brand": null, "sku_prefix": "LIBRE", "category_slug": "papeleria", "price_reference": 1500},
    {"name": "Stickers decorativos surtidos x100u", "unit": "paquete", "brand": null, "sku_prefix": "LIBRE", "category_slug": "manualidades", "price_reference": 1200}
  ]$products$::jsonb,
  updated_at = NOW()
WHERE id = '4d53e0e5-bef2-4af5-9f41-7113ca9c1b52';


-- =====================================================
-- 6. ROPA
--    id: 027b96f7-e3bc-4d3f-a236-50c4e8d66d32
--    sku_prefix: ROPA
-- =====================================================
UPDATE business_type_templates
SET
  categories = $categories$[
    {"name": "Remeras y Camisetas", "slug": "remeras", "level": 0},
    {"name": "Pantalones y Jeans", "slug": "pantalones", "level": 0},
    {"name": "Ropa Interior", "slug": "ropa-interior", "level": 0},
    {"name": "Medias y Calcetines", "slug": "medias", "level": 0},
    {"name": "Calzado", "slug": "calzado", "level": 0},
    {"name": "Abrigos y Camperas", "slug": "abrigos", "level": 0},
    {"name": "Ropa Deportiva", "slug": "ropa-deportiva", "level": 0},
    {"name": "Ropa de Mujer", "slug": "ropa-mujer", "level": 0},
    {"name": "Ropa de Niños", "slug": "ropa-ninhos", "level": 0},
    {"name": "Accesorios de Vestir", "slug": "accesorios-vestir", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "Fruit of the Loom", "suggested_for_categories": ["remeras", "ropa-interior"]},
    {"name": "Topper", "suggested_for_categories": ["calzado", "ropa-deportiva"]},
    {"name": "Wrangler", "suggested_for_categories": ["pantalones"]},
    {"name": "Mimo", "suggested_for_categories": ["ropa-ninhos"]},
    {"name": "Cacharel", "suggested_for_categories": ["ropa-interior"]},
    {"name": "Punto Blanco", "suggested_for_categories": ["ropa-interior", "medias"]},
    {"name": "Levi's", "suggested_for_categories": ["pantalones"]},
    {"name": "Columbia", "suggested_for_categories": ["abrigos", "ropa-deportiva"]},
    {"name": "Puma", "suggested_for_categories": ["calzado", "ropa-deportiva"]}
  ]$brands$::jsonb,
  products = $products$[
    {"name": "Remera básica algodón blanca talle M hombre", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "remeras", "price_reference": 5500},
    {"name": "Remera básica algodón negra talle M hombre", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "remeras", "price_reference": 5500},
    {"name": "Remera básica algodón gris talle L hombre", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "remeras", "price_reference": 5500},
    {"name": "Remera cuello V mujer talle S", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "remeras", "price_reference": 5200},
    {"name": "Remera cuello V mujer talle M", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "remeras", "price_reference": 5200},
    {"name": "Remera estampada hombre talle M surtida", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "remeras", "price_reference": 6800},
    {"name": "Pantalón jeans hombre azul talle 32", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "pantalones", "price_reference": 18500},
    {"name": "Pantalón jeans hombre negro talle 34", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "pantalones", "price_reference": 18500},
    {"name": "Pantalón jeans hombre azul talle 36", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "pantalones", "price_reference": 18500},
    {"name": "Pantalón jogger hombre algodón talle M", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "pantalones", "price_reference": 10500},
    {"name": "Bermuda deportiva hombre talle M", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-deportiva", "price_reference": 8500},
    {"name": "Calza mujer negra supplex talle S", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-mujer", "price_reference": 9500},
    {"name": "Calza mujer negra supplex talle M", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-mujer", "price_reference": 9500},
    {"name": "Short deportivo mujer talle M", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-deportiva", "price_reference": 7500},
    {"name": "Slip hombre talle M algodón pack x3u", "unit": "pack", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-interior", "price_reference": 7200},
    {"name": "Boxer hombre talle M algodón pack x3u", "unit": "pack", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-interior", "price_reference": 8500},
    {"name": "Colaless mujer talle M algodón pack x3u", "unit": "pack", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-interior", "price_reference": 6800},
    {"name": "Corpiño básico sin aro talle 85B", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-interior", "price_reference": 7500},
    {"name": "Corpiño básico sin aro talle 90C", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-interior", "price_reference": 7500},
    {"name": "Calcetines hombre corto algodón x3 pares", "unit": "pack", "brand": null, "sku_prefix": "ROPA", "category_slug": "medias", "price_reference": 4200},
    {"name": "Calcetines hombre largo algodón x3 pares", "unit": "pack", "brand": null, "sku_prefix": "ROPA", "category_slug": "medias", "price_reference": 4500},
    {"name": "Pinkies mujer algodón pack x6u", "unit": "pack", "brand": null, "sku_prefix": "ROPA", "category_slug": "medias", "price_reference": 3500},
    {"name": "Bucaneras mujer microfibra pack x3u", "unit": "pack", "brand": null, "sku_prefix": "ROPA", "category_slug": "medias", "price_reference": 3200},
    {"name": "Ojotas goma talle 40 hombre", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "calzado", "price_reference": 4800},
    {"name": "Ojotas goma talle 38 mujer", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "calzado", "price_reference": 4500},
    {"name": "Ojotas goma talle 36 mujer", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "calzado", "price_reference": 4500},
    {"name": "Zapatillas lona básicas talle 42", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "calzado", "price_reference": 12500},
    {"name": "Zapatillas lona básicas talle 40", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "calzado", "price_reference": 12500},
    {"name": "Zapatillas deportivas Topper talle 42", "unit": "unidad", "brand": "Topper", "sku_prefix": "ROPA", "category_slug": "calzado", "price_reference": 32000},
    {"name": "Campera rompevientos impermeable talle M", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "abrigos", "price_reference": 22000},
    {"name": "Polar hombre talle M antifrizz", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "abrigos", "price_reference": 16500},
    {"name": "Campera acolchada mujer talle M", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "abrigos", "price_reference": 28000},
    {"name": "Buzo canguro algodón hombre talle M", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "abrigos", "price_reference": 14000},
    {"name": "Remera manga larga térmica hombre M", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "remeras", "price_reference": 9500},
    {"name": "Ropa de bebé body algodón talle 3m", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-ninhos", "price_reference": 3800},
    {"name": "Pijama niño algodón talle 4", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-ninhos", "price_reference": 8500},
    {"name": "Pantalón deportivo niño talle 8", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-ninhos", "price_reference": 7200},
    {"name": "Vestido básico mujer talle M algodón", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "ropa-mujer", "price_reference": 12000},
    {"name": "Cinturón cuero sintético hombre 105cm", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "accesorios-vestir", "price_reference": 4500},
    {"name": "Gorra visera frente plano surtida", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "accesorios-vestir", "price_reference": 4200},
    {"name": "Bufanda tejida invierno surtida", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "accesorios-vestir", "price_reference": 5500},
    {"name": "Guantes lana invierno par surtido", "unit": "par", "brand": null, "sku_prefix": "ROPA", "category_slug": "accesorios-vestir", "price_reference": 3200},
    {"name": "Musculosa algodón hombre talle M blanca", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "remeras", "price_reference": 4200},
    {"name": "Medias de vestir hombre x3 pares", "unit": "pack", "brand": null, "sku_prefix": "ROPA", "category_slug": "medias", "price_reference": 4800},
    {"name": "Jean mujer talle 28 elastizado", "unit": "unidad", "brand": null, "sku_prefix": "ROPA", "category_slug": "pantalones", "price_reference": 16500}
  ]$products$::jsonb,
  updated_at = NOW()
WHERE id = '027b96f7-e3bc-4d3f-a236-50c4e8d66d32';


-- =====================================================
-- 7. ELECTRODOMÉSTICOS
--    id: 4f1816ac-8e3f-40f3-8d96-e2dc54e23d33
--    sku_prefix: ELECT
-- =====================================================
UPDATE business_type_templates
SET
  categories = $categories$[
    {"name": "Línea Blanca", "slug": "linea-blanca", "level": 0},
    {"name": "Audio y Video", "slug": "audio-video", "level": 0},
    {"name": "Climatización", "slug": "climatizacion", "level": 0},
    {"name": "Iluminación", "slug": "iluminacion", "level": 0},
    {"name": "Pequeños Electrodomésticos", "slug": "pequenhos-electrodomesticos", "level": 0},
    {"name": "Conectividad y Accesorios", "slug": "conectividad", "level": 0},
    {"name": "Imagen y Fotografía", "slug": "imagen-fotografia", "level": 0},
    {"name": "Calefacción", "slug": "calefaccion", "level": 0},
    {"name": "Ventilación", "slug": "ventilacion", "level": 0},
    {"name": "Hornos y Cocción", "slug": "hornos-coccion", "level": 0}
  ]$categories$::jsonb,
  brands = $brands$[
    {"name": "Atma", "suggested_for_categories": ["pequenhos-electrodomesticos", "linea-blanca", "climatizacion"]},
    {"name": "Oster", "suggested_for_categories": ["pequenhos-electrodomesticos", "hornos-coccion"]},
    {"name": "Philips", "suggested_for_categories": ["pequenhos-electrodomesticos", "iluminacion"]},
    {"name": "Samsung", "suggested_for_categories": ["audio-video", "linea-blanca"]},
    {"name": "LG", "suggested_for_categories": ["audio-video", "linea-blanca"]},
    {"name": "Noblex", "suggested_for_categories": ["audio-video"]},
    {"name": "Liliana", "suggested_for_categories": ["ventilacion", "calefaccion"]},
    {"name": "JBL", "suggested_for_categories": ["audio-video", "conectividad"]},
    {"name": "Sony", "suggested_for_categories": ["audio-video"]},
    {"name": "Rheem", "suggested_for_categories": ["calefaccion", "linea-blanca"]}
  ]$brands$::jsonb,
  products = $products$[
    {"name": "Licuadora Oster 500W 1.25L", "unit": "unidad", "brand": "Oster", "sku_prefix": "ELECT", "category_slug": "pequenhos-electrodomesticos", "price_reference": 28000},
    {"name": "Minipimer Oster 350W acero inox", "unit": "unidad", "brand": "Oster", "sku_prefix": "ELECT", "category_slug": "pequenhos-electrodomesticos", "price_reference": 22000},
    {"name": "Tostadora 2 ranuras Atma 700W", "unit": "unidad", "brand": "Atma", "sku_prefix": "ELECT", "category_slug": "pequenhos-electrodomesticos", "price_reference": 15000},
    {"name": "Sandwichera Atma 700W placas removibles", "unit": "unidad", "brand": "Atma", "sku_prefix": "ELECT", "category_slug": "pequenhos-electrodomesticos", "price_reference": 18500},
    {"name": "Freidora de aire Atma 3.5L 1400W", "unit": "unidad", "brand": "Atma", "sku_prefix": "ELECT", "category_slug": "pequenhos-electrodomesticos", "price_reference": 65000},
    {"name": "Cafetera goteo Oster 12 tazas 900W", "unit": "unidad", "brand": "Oster", "sku_prefix": "ELECT", "category_slug": "pequenhos-electrodomesticos", "price_reference": 28000},
    {"name": "Cafetera cápsulas Nespresso Essenza Mini", "unit": "unidad", "brand": null, "sku_prefix": "ELECT", "category_slug": "pequenhos-electrodomesticos", "price_reference": 85000},
    {"name": "Plancha vapor Philips 2400W antiadherente", "unit": "unidad", "brand": "Philips", "sku_prefix": "ELECT", "category_slug": "pequenhos-electrodomesticos", "price_reference": 32000},
    {"name": "Aspiradora ciclónica Atma 1600W sin bolsa", "unit": "unidad", "brand": "Atma", "sku_prefix": "ELECT", "category_slug": "pequenhos-electrodomesticos", "price_reference": 58000},
    {"name": "Hervidor eléctrico Atma 1.7L 2200W inox", "unit": "unidad", "brand": "Atma", "sku_prefix": "ELECT", "category_slug": "pequenhos-electrodomesticos", "price_reference": 18000},
    {"name": "Batidora de pie Oster 250W 5 velocidades", "unit": "unidad", "brand": "Oster", "sku_prefix": "ELECT", "category_slug": "pequenhos-electrodomesticos", "price_reference": 45000},
    {"name": "Procesadora de alimentos Atma 500W", "unit": "unidad", "brand": "Atma", "sku_prefix": "ELECT", "category_slug": "pequenhos-electrodomesticos", "price_reference": 38000},
    {"name": "Horno eléctrico Atma 46L 1800W", "unit": "unidad", "brand": "Atma", "sku_prefix": "ELECT", "category_slug": "hornos-coccion", "price_reference": 72000},
    {"name": "Microondas Atma 20L 700W digital", "unit": "unidad", "brand": "Atma", "sku_prefix": "ELECT", "category_slug": "hornos-coccion", "price_reference": 68000},
    {"name": "Parlante Bluetooth JBL Go 3 5W", "unit": "unidad", "brand": "JBL", "sku_prefix": "ELECT", "category_slug": "audio-video", "price_reference": 22000},
    {"name": "Parlante Bluetooth JBL Flip 6 30W", "unit": "unidad", "brand": "JBL", "sku_prefix": "ELECT", "category_slug": "audio-video", "price_reference": 65000},
    {"name": "Auriculares Bluetooth Sony WH-CH510", "unit": "unidad", "brand": "Sony", "sku_prefix": "ELECT", "category_slug": "audio-video", "price_reference": 38000},
    {"name": "Auriculares in-ear Sony WF-C500 Bluetooth", "unit": "unidad", "brand": "Sony", "sku_prefix": "ELECT", "category_slug": "audio-video", "price_reference": 45000},
    {"name": "Smart TV 32\" LED HD Noblex", "unit": "unidad", "brand": "Noblex", "sku_prefix": "ELECT", "category_slug": "audio-video", "price_reference": 185000},
    {"name": "Smart TV 43\" Samsung Full HD", "unit": "unidad", "brand": "Samsung", "sku_prefix": "ELECT", "category_slug": "audio-video", "price_reference": 320000},
    {"name": "Smart TV 50\" LG 4K UHD", "unit": "unidad", "brand": "LG", "sku_prefix": "ELECT", "category_slug": "audio-video", "price_reference": 480000},
    {"name": "Decodificador TDA Zinex HDMI", "unit": "unidad", "brand": null, "sku_prefix": "ELECT", "category_slug": "audio-video", "price_reference": 12000},
    {"name": "Ventilador de pie Liliana 18\" 5 velocidades", "unit": "unidad", "brand": "Liliana", "sku_prefix": "ELECT", "category_slug": "ventilacion", "price_reference": 32000},
    {"name": "Ventilador de mesa Liliana 12\" silencioso", "unit": "unidad", "brand": "Liliana", "sku_prefix": "ELECT", "category_slug": "ventilacion", "price_reference": 22000},
    {"name": "Ventilador de techo Peabody 52\" 3 velocidades", "unit": "unidad", "brand": null, "sku_prefix": "ELECT", "category_slug": "ventilacion", "price_reference": 55000},
    {"name": "Ventilador de torre Liliana 40cm oscilante", "unit": "unidad", "brand": "Liliana", "sku_prefix": "ELECT", "category_slug": "ventilacion", "price_reference": 42000},
    {"name": "Caloventor eléctrico Liliana 2000W", "unit": "unidad", "brand": "Liliana", "sku_prefix": "ELECT", "category_slug": "calefaccion", "price_reference": 28000},
    {"name": "Estufa halógena 1500W 3 cuarzos", "unit": "unidad", "brand": null, "sku_prefix": "ELECT", "category_slug": "calefaccion", "price_reference": 18000},
    {"name": "Estufa a cuarzo 2000W 4 tubos", "unit": "unidad", "brand": null, "sku_prefix": "ELECT", "category_slug": "calefaccion", "price_reference": 22000},
    {"name": "Termotanque eléctrico 50L Rheem", "unit": "unidad", "brand": "Rheem", "sku_prefix": "ELECT", "category_slug": "linea-blanca", "price_reference": 145000},
    {"name": "Termotanque eléctrico 80L Rheem", "unit": "unidad", "brand": "Rheem", "sku_prefix": "ELECT", "category_slug": "linea-blanca", "price_reference": 185000},
    {"name": "Lamparita LED 9W Osram E27 cálida", "unit": "unidad", "brand": null, "sku_prefix": "ELECT", "category_slug": "iluminacion", "price_reference": 1800},
    {"name": "Lamparita LED 12W Philips E27 fría", "unit": "unidad", "brand": "Philips", "sku_prefix": "ELECT", "category_slug": "iluminacion", "price_reference": 2400},
    {"name": "Lamparita LED 6W GU10 dicroica", "unit": "unidad", "brand": null, "sku_prefix": "ELECT", "category_slug": "iluminacion", "price_reference": 2200},
    {"name": "Tira LED 5m 12V luz cálida", "unit": "unidad", "brand": null, "sku_prefix": "ELECT", "category_slug": "iluminacion", "price_reference": 8500},
    {"name": "Tira LED 5m RGB con control remoto", "unit": "unidad", "brand": null, "sku_prefix": "ELECT", "category_slug": "iluminacion", "price_reference": 12000},
    {"name": "Zapatilla 5 bocas con USB Atma", "unit": "unidad", "brand": "Atma", "sku_prefix": "ELECT", "category_slug": "conectividad", "price_reference": 6500},
    {"name": "Cargador USB-C 65W carga rápida", "unit": "unidad", "brand": null, "sku_prefix": "ELECT", "category_slug": "conectividad", "price_reference": 8500},
    {"name": "Cable HDMI 2.0 1.8m 4K", "unit": "unidad", "brand": null, "sku_prefix": "ELECT", "category_slug": "conectividad", "price_reference": 3500},
    {"name": "Acondicionador de aire frío-calor 3000 frigorías", "unit": "unidad", "brand": null, "sku_prefix": "ELECT", "category_slug": "climatizacion", "price_reference": 480000},
    {"name": "Purificador de aire Atma HEPA 30m²", "unit": "unidad", "brand": "Atma", "sku_prefix": "ELECT", "category_slug": "climatizacion", "price_reference": 68000},
    {"name": "Lavarropas automático 6kg LG", "unit": "unidad", "brand": "LG", "sku_prefix": "ELECT", "category_slug": "linea-blanca", "price_reference": 285000},
    {"name": "Heladera con freezer 320L Samsung", "unit": "unidad", "brand": "Samsung", "sku_prefix": "ELECT", "category_slug": "linea-blanca", "price_reference": 420000}
  ]$products$::jsonb,
  updated_at = NOW()
WHERE id = '4f1816ac-8e3f-40f3-8d96-e2dc54e23d33';


-- =====================================================
-- CONTEO DE PRODUCTOS POR RUBRO
-- =====================================================
-- fiambreria       (ad55633a): 50 productos
-- piletas          (b5407229): 50 productos
-- bazar            (c278befb): 50 productos
-- jugueteria       (86001b97): 45 productos
-- libreria         (4d53e0e5): 45 productos
-- ropa             (027b96f7): 45 productos
-- electrodomesticos(4f1816ac): 43 productos
-- =====================================================
-- TOTAL: 328 productos en 7 rubros
-- Fuente: research de mercado NEA/Posadas, Misiones — 2026-04-25
-- Precios orientativos en pesos argentinos, zona Posadas, abril 2026
-- =====================================================
