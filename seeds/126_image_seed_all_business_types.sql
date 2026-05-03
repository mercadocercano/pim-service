-- Seed 126: Imágenes para todos los business_types con cobertura 0%
-- Fecha: 2026-05-02
-- IDEMPOTENTE: UPDATE WHERE image_url IS NULL
-- Fuentes: Wikimedia Commons (commons.wikimedia.org)
-- Zona: Argentina (NEA/Posadas)
-- ================================================================
-- Estado previo (medido 2026-05-02):
--   100%: almacen, corralon, electricidad, ferreteria, kiosco, panaderia
--     0%: bazar, carniceria, electrodomesticos, farmacia, fiambreria,
--         jugueteria, libreria, limpieza, peluqueria, perfumeria,
--         piletas, ropa, verduleria, veterinaria, vinoteca
-- Objetivo: llevar todos los 0% a >= 50% de cobertura
-- ================================================================


-- ===========================
-- === VERDULERIA (172 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Bananas_white_background_DS.jpg/400px-Bananas_white_background_DS.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%banana%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/15/Red_Apple.jpg/400px-Red_Apple.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND (name ILIKE '%manzana%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/90/Hapus_Mango.jpg/400px-Hapus_Mango.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%mango%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Oranges_and_orange_slices.jpg/400px-Oranges_and_orange_slices.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND (name ILIKE '%naranja%' OR name ILIKE '%mandarina%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/d/d2/Tomates.jpg/400px-Tomates.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%tomate%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a9/Potato_cultivar_Gunda.jpg/400px-Potato_cultivar_Gunda.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%papa%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a6/Pink_onion.jpg/400px-Pink_onion.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%cebolla%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Iceberg_lettuce.jpg/400px-Iceberg_lettuce.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%lechuga%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3b/Limones.jpg/400px-Limones.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND (name ILIKE '%limón%' OR name ILIKE '%limon%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/48/RedCapsicum.jpg/400px-RedCapsicum.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND (name ILIKE '%morrón%' OR name ILIKE '%morron%' OR name ILIKE '%pimiento%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/27/Zanahoria.jpg/400px-Zanahoria.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%zanahoria%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/0d/Calabaza_zapallo.jpg/400px-Calabaza_zapallo.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND (name ILIKE '%calabaza%' OR name ILIKE '%zapallo%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3e/Broccoli_and_cross_section_edit.jpg/400px-Broccoli_and_cross_section_edit.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%brócoli%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Zucchini_-_Edit.jpg/400px-Zucchini_-_Edit.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND (name ILIKE '%zapallito%' OR name ILIKE '%zucchini%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/28/Garlic.jpg/400px-Garlic.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%ajo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2e/Chard_leaves.jpg/400px-Chard_leaves.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%acelga%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/37/Popeye_spinach.jpg/400px-Popeye_spinach.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%espinaca%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/f/f7/Corn_on_the_cob.jpg/400px-Corn_on_the_cob.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%choclo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/13/Eggplant.jpg/400px-Eggplant.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%berenjena%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/44/Strawberries_white_bg.jpg/400px-Strawberries_white_bg.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%frutilla%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/Durazno.jpg/400px-Durazno.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%durazno%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e8/Peras_conferencia.jpg/400px-Peras_conferencia.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%pera%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Single_grape.png/400px-Single_grape.png',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%uva%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Kiwi-Fruit.jpg/400px-Kiwi-Fruit.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%kiwi%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/0a/Pineapple_and_cross_section.jpg/400px-Pineapple_and_cross_section.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%ananá%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Watermelon_seedless.jpg/400px-Watermelon_seedless.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%sandía%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6e/Batatas.jpg/400px-Batatas.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%batata%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/27/Chicken_Egg_Unlabeled.jpg/400px-Chicken_Egg_Unlabeled.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND name ILIKE '%huevo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e5/Almond_open_front.jpg/400px-Almond_open_front.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND (name ILIKE '%almendra%' OR name ILIKE '%cajú%' OR name ILIKE '%castaña%' OR name ILIKE '%nuez%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/01/Blueberries.jpg/400px-Blueberries.jpg',
    updated_at = NOW()
WHERE business_type = 'verduleria'
  AND (name ILIKE '%arándano%' OR name ILIKE '%frambuesa%')
  AND image_url IS NULL;


-- ===========================
-- === CARNICERIA (215 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/71/Roast_beef.jpg/400px-Roast_beef.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND (name ILIKE '%asado%' OR name ILIKE '%tira%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/1e/Pollo_entero_asado.jpg/400px-Pollo_entero_asado.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND name ILIKE '%pollo entero%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b4/Chicken_breast_with_rib.jpg/400px-Chicken_breast_with_rib.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND name ILIKE '%pechuga%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/40/Chorizo_bilbao.jpg/400px-Chorizo_bilbao.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND name ILIKE '%chorizo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/1b/Salami_aka.jpg/400px-Salami_aka.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND name ILIKE '%salame%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Hamburger_burger.jpg/400px-Hamburger_burger.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND name ILIKE '%hamburguesa%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6e/Sirloin_strip_steak.jpg/400px-Sirloin_strip_steak.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND (name ILIKE '%bife%' OR name ILIKE '%lomo%' OR name ILIKE '%nalga%' OR name ILIKE '%cuadril%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/41/Pork_ribs.jpg/400px-Pork_ribs.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND (name ILIKE '%cerdo%' OR name ILIKE '%chancho%' OR name ILIKE '%bondiola%' OR name ILIKE '%carré%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Milanesa_argentina.jpg/400px-Milanesa_argentina.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND name ILIKE '%milanesa%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/58/Sausages.jpg/400px-Sausages.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND (name ILIKE '%morcilla%' OR name ILIKE '%salchicha%' OR name ILIKE '%longaniza%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Chicken_thigh.jpg/400px-Chicken_thigh.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND (name ILIKE '%muslo%' OR name ILIKE '%pata%' OR name ILIKE '%ala%' OR name ILIKE '%suprema%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/1b/Jamon-serrano.jpg/400px-Jamon-serrano.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND (name ILIKE '%jamón%' OR name ILIKE '%jamon%' OR name ILIKE '%paleta%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/35/Panceta.jpg/400px-Panceta.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND name ILIKE '%panceta%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Matambre.jpg/400px-Matambre.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND name ILIKE '%matambre%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/0d/Vacio-corte-argentino.jpg/400px-Vacio-corte-argentino.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND (name ILIKE '%vacío%' OR name ILIKE '%vacio%' OR name ILIKE '%peceto%' OR name ILIKE '%roast%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/98/Liver_beef.jpg/400px-Liver_beef.jpg',
    updated_at = NOW()
WHERE business_type = 'carniceria'
  AND (name ILIKE '%hígado%' OR name ILIKE '%riñón%' OR name ILIKE '%molleja%' OR name ILIKE '%mondongo%')
  AND image_url IS NULL;


-- ===========================
-- === FARMACIA (90 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9b/Pampers_diapers.jpg/400px-Pampers_diapers.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%pañal%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/56/White_shark.jpg/400px-White_shark.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%preservativo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/1c/Sunscreen_products.jpg/400px-Sunscreen_products.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%protector solar%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/17/Toothbrushes-varying-sizes.jpg/400px-Toothbrushes-varying-sizes.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%cepillo de dientes%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/d/d8/Oral-B_Pro-500_toothbrush.jpg/400px-Oral-B_Pro-500_toothbrush.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%pasta dental%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Deodorant_spray.jpg/400px-Deodorant_spray.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%desodorante%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/08/Toallas_femeninas.jpg/400px-Toallas_femeninas.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND (name ILIKE '%toalla femenina%' OR name ILIKE '%toallas femeninas%' OR name ILIKE '%compresa%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/71/Tampons.jpg/400px-Tampons.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%tampón%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5c/Jabón_de_manos_Dove.jpg/400px-Jabón_de_manos_Dove.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND (name ILIKE '%jabón%' OR name ILIKE '%jabon%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/7b/Shampoo_bottle.jpg/400px-Shampoo_bottle.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%shampoo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/56/Body_lotion.jpg/400px-Body_lotion.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%crema hidratante%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3e/Band_aid_brand_adhesive_bandages.jpg/400px-Band_aid_brand_adhesive_bandages.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%curita%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Cotton_balls.jpg/400px-Cotton_balls.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%algodón%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/f/f3/Thermometer.jpg/400px-Thermometer.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%termómetro%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/be/Repelente_insectos.jpg/400px-Repelente_insectos.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%repelente%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/27/Baby_wipes.jpg/400px-Baby_wipes.jpg',
    updated_at = NOW()
WHERE business_type = 'farmacia'
  AND name ILIKE '%toallita%'
  AND image_url IS NULL;


-- ===========================
-- === FIAMBRERIA (148 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/31/Sliced_Ham.jpg/400px-Sliced_Ham.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND name ILIKE '%jamón cocido%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/f/f5/Jamon_serrano_01.jpg/400px-Jamon_serrano_01.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND name ILIKE '%jamón crudo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/1b/Salami_aka.jpg/400px-Salami_aka.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND name ILIKE '%salame%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c7/Chorizo_colorado.jpg/400px-Chorizo_colorado.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND name ILIKE '%chorizo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/8b/Queso-fresco.jpg/400px-Queso-fresco.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND (name ILIKE '%queso crema%' OR name ILIKE '%queso untable%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/Gouda_cheese.jpg/400px-Gouda_cheese.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND (name ILIKE '%queso gouda%' OR name ILIKE '%queso tybo%' OR name ILIKE '%queso dambo%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/7b/Mozzarella.jpg/400px-Mozzarella.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND name ILIKE '%queso mozzarella%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/0e/Gorgonzola.jpg/400px-Gorgonzola.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND (name ILIKE '%queso azul%' OR name ILIKE '%roquefort%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/67/Parmigiano_Reggiano.jpg/400px-Parmigiano_Reggiano.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND (name ILIKE '%parmesano%' OR name ILIKE '%reggianito%' OR name ILIKE '%sardo%' OR name ILIKE '%provolone%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Ricotta_cheese.jpg/400px-Ricotta_cheese.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND name ILIKE '%ricot%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/f/fc/Cream_cheese.jpg/400px-Cream_cheese.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND (name ILIKE '%queso brie%' OR name ILIKE '%queso camembert%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4d/Olive_oil.jpg/400px-Olive_oil.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND name ILIKE '%aceite de oliva%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/15/Aceitunas_verdes.jpg/400px-Aceitunas_verdes.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND name ILIKE '%aceituna%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5b/Mortadela.jpg/400px-Mortadela.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND name ILIKE '%mortadela%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/8a/Manteca-argentina.jpg/400px-Manteca-argentina.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND name ILIKE '%manteca%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/d/d5/Cream_250g.jpg/400px-Cream_250g.jpg',
    updated_at = NOW()
WHERE business_type = 'fiambreria'
  AND name ILIKE '%crema de leche%'
  AND image_url IS NULL;


-- ===========================
-- === LIMPIEZA (94 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/64/Bleach_jug.jpg/400px-Bleach_jug.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%lavandina%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3b/Laundry_detergent.jpg/400px-Laundry_detergent.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%jabón en polvo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/0c/Dish_soap.jpg/400px-Dish_soap.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%detergente%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/38/Cleaning_spray_bottle.jpg/400px-Cleaning_spray_bottle.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND (name ILIKE '%desinfectante%' OR name ILIKE '%desengrasante%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/8e/Esponja_cocina.jpg/400px-Esponja_cocina.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%esponja%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Rubbish_bin.jpg/400px-Rubbish_bin.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%bolsa de basura%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/1b/Broom.jpg/400px-Broom.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%escoba%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/36/Mop_bucket_and_mop_combination.jpg/400px-Mop_bucket_and_mop_combination.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%lampazo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5b/Glade_air_freshener.jpg/400px-Glade_air_freshener.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%desodorante de ambiente%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/30/Guantes_de_limpieza.jpg/400px-Guantes_de_limpieza.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%guantes%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/40/Raid_insecticide.jpg/400px-Raid_insecticide.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%insecticida%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/Vileda_mop.jpg/400px-Vileda_mop.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%mopa%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/7b/Toilet_bowl_cleaner.jpg/400px-Toilet_bowl_cleaner.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%inodoro%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Microfiber_cloth.jpg/400px-Microfiber_cloth.jpg',
    updated_at = NOW()
WHERE business_type = 'limpieza'
  AND name ILIKE '%franela%'
  AND image_url IS NULL;


-- ===========================
-- === BAZAR (174 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/Tramontina_pan.jpg/400px-Tramontina_pan.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND (name ILIKE '%sartén%' OR name ILIKE '%sarten%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Stainless_steel_pot.jpg/400px-Stainless_steel_pot.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND (name ILIKE '%olla%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Pressure_cooker.jpg/400px-Pressure_cooker.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%olla a presión%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/44/Chef_knife.jpg/400px-Chef_knife.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%cuchillo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/Cutlery_set.jpg/400px-Cutlery_set.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%cubiertos%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Juego_de_platos.jpg/400px-Juego_de_platos.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%juego de platos%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Stanley_thermos.jpg/400px-Stanley_thermos.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%termo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/74/Stanley_mate.jpg/400px-Stanley_mate.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%mate%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a0/Cutting_board_bamboo.jpg/400px-Cutting_board_bamboo.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%tabla%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e4/Glass_jug.jpg/400px-Glass_jug.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%jarra%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6e/Plastic_storage_containers.jpg/400px-Plastic_storage_containers.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%contenedor%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b1/Broom_and_dustpan.jpg/400px-Broom_and_dustpan.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%escoba%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/13/Floor_mop.jpg/400px-Floor_mop.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%mopa%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Glass_bowls.jpg/400px-Glass_bowls.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND (name ILIKE '%bowl%' OR name ILIKE '%ensaladera%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/8a/Plastic_bucket.jpg/400px-Plastic_bucket.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%balde%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c2/Drinking_glasses.jpg/400px-Drinking_glasses.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND (name ILIKE '%vaso%' OR name ILIKE '%copa%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/08/Taza-cafe-porcelana.jpg/400px-Taza-cafe-porcelana.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%taza%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/55/Plastic_planter.jpg/400px-Plastic_planter.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%maceta%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3a/LED_wall_clock.jpg/400px-LED_wall_clock.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND name ILIKE '%reloj%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c5/Candle.jpg/400px-Candle.jpg',
    updated_at = NOW()
WHERE business_type = 'bazar'
  AND (name ILIKE '%vela%')
  AND image_url IS NULL;


-- ===========================
-- === ELECTRODOMESTICOS (146 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/Refrigerator.jpg/400px-Refrigerator.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%heladera%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/70/Washing_machine.jpg/400px-Washing_machine.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%lavarropas%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/8d/Microwave_oven.jpg/400px-Microwave_oven.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%microondas%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/67/Smart_TV.jpg/400px-Smart_TV.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%smart tv%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c2/Air_conditioner_split.jpg/400px-Air_conditioner_split.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND (name ILIKE '%aire acondicionado%' OR name ILIKE '%split%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Electric_blender.jpg/400px-Electric_blender.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%licuadora%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Hair_dryer.jpg/400px-Hair_dryer.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%secador%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/30/Electric_iron.jpg/400px-Electric_iron.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%plancha%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Electric_kettle.jpg/400px-Electric_kettle.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%pava eléctrica%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e4/Coffee_maker.jpg/400px-Coffee_maker.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%cafetera%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/55/Vacuum_cleaner.jpg/400px-Vacuum_cleaner.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%aspiradora%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/95/Smartphone_Samsung.jpg/400px-Smartphone_Samsung.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%celular%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/43/Laptop_computer.jpg/400px-Laptop_computer.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%notebook%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/Gas_stove.jpg/400px-Gas_stove.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND (name ILIKE '%cocina a gas%' OR name ILIKE '%cocina eléctrica%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/7d/Air_fryer.jpg/400px-Air_fryer.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%freidora de aire%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3e/Electric_space_heater.jpg/400px-Electric_space_heater.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND (name ILIKE '%caloventor%' OR name ILIKE '%radiador%' OR name ILIKE '%estufa%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/52/Bluetooth_speaker.jpg/400px-Bluetooth_speaker.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%parlante%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/45/Bluetooth_headphones.jpg/400px-Bluetooth_headphones.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%auricular%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/f/f3/Tablet_computer.jpg/400px-Tablet_computer.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%tablet%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/8b/LED_bulb.jpg/400px-LED_bulb.jpg',
    updated_at = NOW()
WHERE business_type = 'electrodomesticos'
  AND name ILIKE '%lamparita%'
  AND image_url IS NULL;


-- ===========================
-- === JUGUETERIA (166 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/25/LEGO_logo.svg/400px-LEGO_logo.svg.png',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%lego%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3b/Barbie_doll.jpg/400px-Barbie_doll.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%barbie%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Hot_Wheels_cars.jpg/400px-Hot_Wheels_cars.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND (name ILIKE '%hot wheels%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/1b/Play-Doh_set.jpg/400px-Play-Doh_set.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%play-doh%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4d/Monopoly_board_game.jpg/400px-Monopoly_board_game.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%monopoly%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/d/d4/Nerf_gun.jpg/400px-Nerf_gun.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%nerf%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/Jigsaw_puzzle.jpg/400px-Jigsaw_puzzle.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND (name ILIKE '%rompecabezas%' OR name ILIKE '%puzzle%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Stuffed_teddy_bear.jpg/400px-Stuffed_teddy_bear.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%peluche%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Playmobil_figurines.jpg/400px-Playmobil_figurines.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%playmobil%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/0b/Rasti_Argentina.jpg/400px-Rasti_Argentina.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%rasti%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/39/Toy_bicycle.jpg/400px-Toy_bicycle.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%bicicleta%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/13/Coloring_crayons.jpg/400px-Coloring_crayons.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%crayones%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2a/Football_ball.jpg/400px-Football_ball.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%pelota de fútbol%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Jenga_game.jpg/400px-Jenga_game.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%jenga%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/f/fb/Scrabble_board.jpg/400px-Scrabble_board.jpg',
    updated_at = NOW()
WHERE business_type = 'jugueteria'
  AND name ILIKE '%scrabble%'
  AND image_url IS NULL;


-- ===========================
-- === LIBRERIA (177 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a2/BIC_Cristal.jpg/400px-BIC_Cristal.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND name ILIKE '%lapicera%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/37/Faber-Castell_pencils.jpg/400px-Faber-Castell_pencils.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND (name ILIKE '%lápiz%' OR name ILIKE '%lapiz%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6e/Colored_pencils.jpg/400px-Colored_pencils.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND name ILIKE '%lápices de colores%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/Notebook_spiral.jpg/400px-Notebook_spiral.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND name ILIKE '%cuaderno%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5c/Scotch_tape.jpg/400px-Scotch_tape.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND name ILIKE '%cinta%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a5/Plasticola_glue.jpg/400px-Plasticola_glue.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND (name ILIKE '%plasticola%' OR name ILIKE '%cola vinílica%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b0/A4_paper_ream.jpg/400px-A4_paper_ream.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND name ILIKE '%papel a4%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c0/Mochila_escolar.jpg/400px-Mochila_escolar.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND name ILIKE '%mochila%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Casio_scientific_calculator.jpg/400px-Casio_scientific_calculator.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND name ILIKE '%calculadora%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3e/Stapler.jpg/400px-Stapler.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND name ILIKE '%abrochadora%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a2/Folder_binder.jpg/400px-Folder_binder.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND (name ILIKE '%carpeta%' OR name ILIKE '%bibliorato%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/7c/Marker_pens.jpg/400px-Marker_pens.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND (name ILIKE '%marcador%' OR name ILIKE '%fibra%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/44/Eraser_white.jpg/400px-Eraser_white.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND name ILIKE '%goma de borrar%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/Watercolor_set.jpg/400px-Watercolor_set.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND name ILIKE '%acuarelas%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/0a/Duracell_AA_batteries.jpg/400px-Duracell_AA_batteries.jpg',
    updated_at = NOW()
WHERE business_type = 'libreria'
  AND name ILIKE '%pilas%'
  AND image_url IS NULL;


-- ===========================
-- === PERFUMERIA (183 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/64/Nivea_cream_tin.jpg/400px-Nivea_cream_tin.jpg',
    updated_at = NOW()
WHERE business_type = 'perfumeria'
  AND name ILIKE '%crema corporal%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/36/Perfume_bottles.jpg/400px-Perfume_bottles.jpg',
    updated_at = NOW()
WHERE business_type = 'perfumeria'
  AND (name ILIKE '%eau de parfum%' OR name ILIKE '%eau de toilette%' OR name ILIKE '%edt%' OR name ILIKE '%edp%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Nail_polish.jpg/400px-Nail_polish.jpg',
    updated_at = NOW()
WHERE business_type = 'perfumeria'
  AND name ILIKE '%esmalte%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a5/Lipstick.jpg/400px-Lipstick.jpg',
    updated_at = NOW()
WHERE business_type = 'perfumeria'
  AND (name ILIKE '%labial%' OR name ILIKE '%brillo labial%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/d/d4/Foundation_makeup.jpg/400px-Foundation_makeup.jpg',
    updated_at = NOW()
WHERE business_type = 'perfumeria'
  AND name ILIKE '%base%'
  AND name ILIKE '%maquillaje%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/56/Mascara_wand.jpg/400px-Mascara_wand.jpg',
    updated_at = NOW()
WHERE business_type = 'perfumeria'
  AND (name ILIKE '%delineador%' OR name ILIKE '%máscara%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4f/Shampoo_conditioner.jpg/400px-Shampoo_conditioner.jpg',
    updated_at = NOW()
WHERE business_type = 'perfumeria'
  AND (name ILIKE '%shampoo%' OR name ILIKE '%acondicionador%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/1c/Depilatory_cream.jpg/400px-Depilatory_cream.jpg',
    updated_at = NOW()
WHERE business_type = 'perfumeria'
  AND (name ILIKE '%crema depilat%' OR name ILIKE '%cera depilat%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/7f/Sunscreen_tube.jpg/400px-Sunscreen_tube.jpg',
    updated_at = NOW()
WHERE business_type = 'perfumeria'
  AND name ILIKE '%protector solar%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/56/Body_splash_perfume.jpg/400px-Body_splash_perfume.jpg',
    updated_at = NOW()
WHERE business_type = 'perfumeria'
  AND name ILIKE '%body splash%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9a/Hand_cream.jpg/400px-Hand_cream.jpg',
    updated_at = NOW()
WHERE business_type = 'perfumeria'
  AND name ILIKE '%crema de manos%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/35/Serum_facial.jpg/400px-Serum_facial.jpg',
    updated_at = NOW()
WHERE business_type = 'perfumeria'
  AND name ILIKE '%sérum%'
  AND image_url IS NULL;


-- ===========================
-- === PELUQUERIA (92 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6d/Hair_color_dye.jpg/400px-Hair_color_dye.jpg',
    updated_at = NOW()
WHERE business_type = 'peluqueria'
  AND name ILIKE '%tintura%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5c/Hair_developer.jpg/400px-Hair_developer.jpg',
    updated_at = NOW()
WHERE business_type = 'peluqueria'
  AND name ILIKE '%oxidante%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/7f/Hair_bleach_powder.jpg/400px-Hair_bleach_powder.jpg',
    updated_at = NOW()
WHERE business_type = 'peluqueria'
  AND name ILIKE '%decolorant%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Professional_hair_dryer.jpg/400px-Professional_hair_dryer.jpg',
    updated_at = NOW()
WHERE business_type = 'peluqueria'
  AND name ILIKE '%secador%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3e/Hair_straightener.jpg/400px-Hair_straightener.jpg',
    updated_at = NOW()
WHERE business_type = 'peluqueria'
  AND name ILIKE '%plancha alisadora%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Hair_curling_iron.jpg/400px-Hair_curling_iron.jpg',
    updated_at = NOW()
WHERE business_type = 'peluqueria'
  AND name ILIKE '%rizador%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Professional_hair_scissors.jpg/400px-Professional_hair_scissors.jpg',
    updated_at = NOW()
WHERE business_type = 'peluqueria'
  AND name ILIKE '%tijera%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b0/Hair_mask_treatment.jpg/400px-Hair_mask_treatment.jpg',
    updated_at = NOW()
WHERE business_type = 'peluqueria'
  AND name ILIKE '%mascarilla capilar%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e0/Hair_gel.jpg/400px-Hair_gel.jpg',
    updated_at = NOW()
WHERE business_type = 'peluqueria'
  AND name ILIKE '%gel fijador%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/ae/Hair_spray_can.jpg/400px-Hair_spray_can.jpg',
    updated_at = NOW()
WHERE business_type = 'peluqueria'
  AND name ILIKE '%laca capilar%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Professional_shampoo.jpg/400px-Professional_shampoo.jpg',
    updated_at = NOW()
WHERE business_type = 'peluqueria'
  AND name ILIKE '%shampoo profesional%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3b/Keratina_capilar.jpg/400px-Keratina_capilar.jpg',
    updated_at = NOW()
WHERE business_type = 'peluqueria'
  AND name ILIKE '%keratina%'
  AND image_url IS NULL;


-- ===========================
-- === PILETAS (116 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/0c/Pool_water_treatment.jpg/400px-Pool_water_treatment.jpg',
    updated_at = NOW()
WHERE business_type = 'piletas'
  AND name ILIKE '%cloro%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/94/Algaecide_pool.jpg/400px-Algaecide_pool.jpg',
    updated_at = NOW()
WHERE business_type = 'piletas'
  AND name ILIKE '%algicida%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/8b/Pool_pump.jpg/400px-Pool_pump.jpg',
    updated_at = NOW()
WHERE business_type = 'piletas'
  AND name ILIKE '%bomba%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c9/Pool_filter_sand.jpg/400px-Pool_filter_sand.jpg',
    updated_at = NOW()
WHERE business_type = 'piletas'
  AND (name ILIKE '%filtro%' OR name ILIKE '%arena%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/39/Pool_vacuum_head.jpg/400px-Pool_vacuum_head.jpg',
    updated_at = NOW()
WHERE business_type = 'piletas'
  AND name ILIKE '%aspiradora%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a2/Pool_net.jpg/400px-Pool_net.jpg',
    updated_at = NOW()
WHERE business_type = 'piletas'
  AND name ILIKE '%red%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/Bestway_inflatable_pool.jpg/400px-Bestway_inflatable_pool.jpg',
    updated_at = NOW()
WHERE business_type = 'piletas'
  AND (name ILIKE '%pileta%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/67/Ph_test_kit.jpg/400px-Ph_test_kit.jpg',
    updated_at = NOW()
WHERE business_type = 'piletas'
  AND (name ILIKE '%pH%' OR name ILIKE '%ph%' OR name ILIKE '%kit%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3e/Pool_cover.jpg/400px-Pool_cover.jpg',
    updated_at = NOW()
WHERE business_type = 'piletas'
  AND (name ILIKE '%cobert%' OR name ILIKE '%cubre%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b3/Pool_ladder.jpg/400px-Pool_ladder.jpg',
    updated_at = NOW()
WHERE business_type = 'piletas'
  AND name ILIKE '%escalera%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c2/Robot_pool_cleaner.jpg/400px-Robot_pool_cleaner.jpg',
    updated_at = NOW()
WHERE business_type = 'piletas'
  AND name ILIKE '%robot%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4f/Pool_brush.jpg/400px-Pool_brush.jpg',
    updated_at = NOW()
WHERE business_type = 'piletas'
  AND name ILIKE '%cepillo%'
  AND image_url IS NULL;


-- ===========================
-- === ROPA (147 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6b/Blue_jeans.jpg/400px-Blue_jeans.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND name ILIKE '%jean%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/33/Polar_fleece_pullover.jpg/400px-Polar_fleece_pullover.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND name ILIKE '%buzo polar%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Windbreaker_jacket.jpg/400px-Windbreaker_jacket.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND name ILIKE '%campera%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Bermuda_shorts.jpg/400px-Bermuda_shorts.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND name ILIKE '%bermuda%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/8a/Sneakers_shoes.jpg/400px-Sneakers_shoes.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND (name ILIKE '%zapatilla%' OR name ILIKE '%botita%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5c/Dress_shirt.jpg/400px-Dress_shirt.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND name ILIKE '%camisa%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6e/Leggings_yoga.jpg/400px-Leggings_yoga.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND name ILIKE '%calza%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/37/Hoodie_sweatshirt.jpg/400px-Hoodie_sweatshirt.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND (name ILIKE '%hoodie%' OR name ILIKE '%buzo canguro%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b1/Sports_socks.jpg/400px-Sports_socks.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND name ILIKE '%calcetín%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4d/Baseball_cap.jpg/400px-Baseball_cap.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND name ILIKE '%gorra%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6d/Jogger_pants.jpg/400px-Jogger_pants.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND name ILIKE '%jogger%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/School_uniform_coat.jpg/400px-School_uniform_coat.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND name ILIKE '%guardapolvo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/8b/Men_boxer_briefs.jpg/400px-Men_boxer_briefs.jpg',
    updated_at = NOW()
WHERE business_type = 'ropa'
  AND name ILIKE '%boxer%'
  AND image_url IS NULL;


-- ===========================
-- === VETERINARIA (120 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c6/Dog_food_kibble.jpg/400px-Dog_food_kibble.jpg',
    updated_at = NOW()
WHERE business_type = 'veterinaria'
  AND name ILIKE '%dog chow%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3e/Cat_food_bowl.jpg/400px-Cat_food_bowl.jpg',
    updated_at = NOW()
WHERE business_type = 'veterinaria'
  AND name ILIKE '%cat chow%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/67/Dog_collar_leash.jpg/400px-Dog_collar_leash.jpg',
    updated_at = NOW()
WHERE business_type = 'veterinaria'
  AND (name ILIKE '%collar%' OR name ILIKE '%correa%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Dog_house.jpg/400px-Dog_house.jpg',
    updated_at = NOW()
WHERE business_type = 'veterinaria'
  AND name ILIKE '%cucha%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5a/Cat_litter_box.jpg/400px-Cat_litter_box.jpg',
    updated_at = NOW()
WHERE business_type = 'veterinaria'
  AND name ILIKE '%arenero%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/8b/Cat_litter.jpg/400px-Cat_litter.jpg',
    updated_at = NOW()
WHERE business_type = 'veterinaria'
  AND name ILIKE '%arena sanitaria%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Pet_shampoo.jpg/400px-Pet_shampoo.jpg',
    updated_at = NOW()
WHERE business_type = 'veterinaria'
  AND name ILIKE '%shampoo%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9a/Bird_food.jpg/400px-Bird_food.jpg',
    updated_at = NOW()
WHERE business_type = 'veterinaria'
  AND (name ILIKE '%alimento aves%' OR name ILIKE '%alpiste%' OR name ILIKE '%mixtura%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3a/Fish_food_tetra.jpg/400px-Fish_food_tetra.jpg',
    updated_at = NOW()
WHERE business_type = 'veterinaria'
  AND name ILIKE '%alimento peces%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c5/Dog_clothes.jpg/400px-Dog_clothes.jpg',
    updated_at = NOW()
WHERE business_type = 'veterinaria'
  AND (name ILIKE '%abrigo%' OR name ILIKE '%camiseta%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/83/Pet_flea_treatment.jpg/400px-Pet_flea_treatment.jpg',
    updated_at = NOW()
WHERE business_type = 'veterinaria'
  AND (name ILIKE '%bravecto%' OR name ILIKE '%antipulgas%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/28/Pet_water_feeder.jpg/400px-Pet_water_feeder.jpg',
    updated_at = NOW()
WHERE business_type = 'veterinaria'
  AND (name ILIKE '%comedero%' OR name ILIKE '%bebedero%')
  AND image_url IS NULL;


-- ===========================
-- === VINOTECA (194 productos) ===
-- ===========================

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a6/Malbec_wine_glass.jpg/400px-Malbec_wine_glass.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND name ILIKE '%malbec%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/Sparkling_wine_glass.jpg/400px-Sparkling_wine_glass.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND (name ILIKE '%chandon%' OR name ILIKE '%extra brut%' OR name ILIKE '%espumante%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6b/Whiskey_bottle.jpg/400px-Whiskey_bottle.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND (name ILIKE '%whisky%' OR name ILIKE '%chivas%' OR name ILIKE '%ballantine%' OR name ILIKE '%jameson%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/0d/Absolut_Vodka.jpg/400px-Absolut_Vodka.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND (name ILIKE '%vodka%' OR name ILIKE '%absolut%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5f/Fernet_Branca.jpg/400px-Fernet_Branca.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND name ILIKE '%fernet%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/42/Campari_bottle.jpg/400px-Campari_bottle.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND (name ILIKE '%campari%' OR name ILIKE '%aperol%' OR name ILIKE '%cinzano%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/35/Bacardi_rum_bottle.jpg/400px-Bacardi_rum_bottle.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND (name ILIKE '%ron%' OR name ILIKE '%bacardi%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/Gin_bottle.jpg/400px-Gin_bottle.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND (name ILIKE '%gin%' OR name ILIKE '%beefeater%' OR name ILIKE '%bombay%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/71/Beer_bottle.jpg/400px-Beer_bottle.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND (name ILIKE '%budweiser%' OR name ILIKE '%brahma%' OR name ILIKE '%heineken%' OR name ILIKE '%beck%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5f/Antares_craft_beer.jpg/400px-Antares_craft_beer.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND name ILIKE '%antares%'
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3d/Wine_corkscrew.jpg/400px-Wine_corkscrew.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND (name ILIKE '%sacacorcho%' OR name ILIKE '%aerador%' OR name ILIKE '%decantador%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/8a/Wine_glass.jpg/400px-Wine_glass.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND (name ILIKE '%copa%' OR name ILIKE '%vaso%')
  AND image_url IS NULL;

UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/40/Catena_Zapata_Malbec.jpg/400px-Catena_Zapata_Malbec.jpg',
    updated_at = NOW()
WHERE business_type = 'vinoteca'
  AND name ILIKE '%catena%'
  AND image_url IS NULL;


-- ===========================
-- Verificación post-ejecución:
-- SELECT business_type, COUNT(*) total, COUNT(image_url) con_img,
--   ROUND(100.0*COUNT(image_url)/NULLIF(COUNT(*),0)) pct
-- FROM global_products
-- GROUP BY business_type ORDER BY pct DESC;
-- ===========================
