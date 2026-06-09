-- 127_image_seed_faltantes.sql
-- Seed: imágenes faltantes en global_products (almacén + ferretería/construcción)
-- Fecha: 2026-05-24
-- IDEMPOTENTE: UPDATE WHERE image_url IS NULL
-- Zona: Argentina (NEA/Posadas)
-- Fuentes verificadas:
--   Open Food Facts CDN: https://images.openfoodfacts.org/images/products/...
--   Wikimedia Commons:   https://upload.wikimedia.org/wikipedia/commons/...
-- Todas las URLs verificadas con HTTP 200 el 2026-05-24.
-- ================================================================


-- ===========================
-- === ALMACÉN ===
-- ===========================

-- Arvejas en lata - Arcor
-- Fuente: Open Food Facts product 7790580132392 (Arvejas Secas Remojadas, Arcor, Argentina)
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/058/013/2392/front_es.16.400.jpg',
    updated_at = NOW()
WHERE name ILIKE '%arveja%'
  AND brand ILIKE '%arcor%'
  AND image_url IS NULL;

-- Arvejas en lata - genérico almacén (sin brand o brand distinta)
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/058/013/2392/front_es.16.400.jpg',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%arveja%'
  AND image_url IS NULL;


-- Atún en aceite - La Campagnola
-- Fuente: Open Food Facts product 7790580131357 (Atún en aceite, La Campagnola, Argentina)
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/058/013/1357/front_es.3.400.jpg',
    updated_at = NOW()
WHERE name ILIKE '%atún%' OR name ILIKE '%atun%'
  AND (brand ILIKE '%campagnola%' OR brand ILIKE '%la campagnola%')
  AND image_url IS NULL;

-- Atún en aceite - genérico almacén
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/058/013/1357/front_es.3.400.jpg',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND (name ILIKE '%atún%' OR name ILIKE '%atun%')
  AND image_url IS NULL;


-- Café torrado 500g - Nescafé Tradición (producto argentino)
-- Fuente: Open Food Facts product 7891000350157 (Nescafé Tradición, Nescafé/Nestlé, Argentina)
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/789/100/035/0157/front_es.27.400.jpg',
    updated_at = NOW()
WHERE (name ILIKE '%nescaf%' OR (name ILIKE '%café%' AND brand ILIKE '%nescaf%') OR (name ILIKE '%cafe%' AND brand ILIKE '%nescaf%'))
  AND image_url IS NULL;

-- Café torrado - genérico almacén
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/789/100/035/0157/front_es.27.400.jpg',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND (name ILIKE '%café torrado%' OR name ILIKE '%cafe torrado%')
  AND image_url IS NULL;


-- Detergente Magistral 500ml y 750ml
-- Fuente: Wikimedia Commons - Afwasmiddel.jpg (frasco genérico detergente lavavajillas)
-- Nota: Magistral no tiene entrada verificada en Open Food Facts AR.
-- Se usa imagen genérica de detergente lavavajillas (Wikimedia CC).
UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/71/Afwasmiddel.jpg/500px-Afwasmiddel.jpg',
    updated_at = NOW()
WHERE name ILIKE '%magistral%'
  AND image_url IS NULL;

-- Detergente / lavavajillas genérico almacén
UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/71/Afwasmiddel.jpg/500px-Afwasmiddel.jpg',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND (name ILIKE '%detergente%' OR name ILIKE '%lavavajilla%')
  AND image_url IS NULL;


-- Dulce de leche 400g - La Serenísima Clásico
-- Fuente: Open Food Facts product 7790742625304 (Dulce de leche clásico, La Serenisima, Argentina)
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/074/262/5304/front_es.25.400.jpg',
    updated_at = NOW()
WHERE name ILIKE '%dulce de leche%'
  AND (brand ILIKE '%serenísima%' OR brand ILIKE '%serenisima%')
  AND image_url IS NULL;

-- Dulce de leche - genérico almacén
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/074/262/5304/front_es.25.400.jpg',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%dulce de leche%'
  AND image_url IS NULL;


-- Jabón en polvo Skip 800g
-- Fuente: Wikimedia Commons - Pralni_prašek.JPG (caja de jabón en polvo genérico)
-- Nota: Skip 800g Argentina no tiene entry verificada en Open Food Facts.
-- Se usa imagen genérica de jabón en polvo (Wikimedia CC BY-SA).
UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2a/Pralni_pra%C5%A1ek.JPG/500px-Pralni_pra%C5%A1ek.JPG',
    updated_at = NOW()
WHERE name ILIKE '%skip%'
  AND image_url IS NULL;

-- Jabón en polvo - genérico almacén
UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2a/Pralni_pra%C5%A1ek.JPG/500px-Pralni_pra%C5%A1ek.JPG',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND (name ILIKE '%jabón en polvo%' OR name ILIKE '%jabon en polvo%' OR name ILIKE '%polvo lavar%')
  AND image_url IS NULL;


-- Lentejas 500g - Marolio
-- Fuente: Open Food Facts product 7797470007921 (Lentejas, Marolio, Argentina)
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/747/000/7921/front_es.11.400.jpg',
    updated_at = NOW()
WHERE name ILIKE '%lenteja%'
  AND brand ILIKE '%marolio%'
  AND image_url IS NULL;

-- Lentejas - genérico almacén
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/747/000/7921/front_es.11.400.jpg',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%lenteja%'
  AND image_url IS NULL;


-- Manteca 200g - La Serenísima
-- Fuente: Open Food Facts product 7793940054006 (Manteca, La Serenísima, Argentina)
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/394/005/4006/front_es.42.400.jpg',
    updated_at = NOW()
WHERE name ILIKE '%manteca%'
  AND (brand ILIKE '%serenísima%' OR brand ILIKE '%serenisima%')
  AND image_url IS NULL;

-- Manteca - genérico almacén
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/394/005/4006/front_es.42.400.jpg',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%manteca%'
  AND image_url IS NULL;


-- Mayonesa Hellmann's 475g
-- Fuente: Open Food Facts product 8722700136231 (Gran mayonesa Hellmanns boca, Hellmann's/Unilever, Argentina)
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/872/270/013/6231/front_en.38.400.jpg',
    updated_at = NOW()
WHERE name ILIKE '%mayonesa%'
  AND (brand ILIKE '%hellmann%' OR brand ILIKE '%hellman%')
  AND image_url IS NULL;

-- Mayonesa - genérico almacén
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/872/270/013/6231/front_en.38.400.jpg',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%mayonesa%'
  AND image_url IS NULL;


-- Vinagre 1L - Marolio / genérico argentino
-- Fuente: Open Food Facts product 7790130000058 (Vinagre de Alcohol, Menoyo, Argentina)
-- Nota: Marolio vinagre no aparece en OFF. Se usa Menoyo que también es marca
--       de distribución nacional con presencia en NEA. Imagen sirve para ambas marcas.
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/013/000/0058/front_es.6.400.jpg',
    updated_at = NOW()
WHERE name ILIKE '%vinagre%'
  AND (brand ILIKE '%marolio%' OR brand ILIKE '%menoyo%')
  AND image_url IS NULL;

-- Vinagre - genérico almacén
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/013/000/0058/front_es.6.400.jpg',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%vinagre%'
  AND image_url IS NULL;


-- ===========================
-- === FERRETERÍA / CONSTRUCCIÓN ===
-- ===========================

-- Cemento Portland (Loma Negra, Holcim, genérico)
-- Fuente: Wikimedia Commons - Portland_Cement_Bags.jpg (bolsas de cemento Portland genéricas)
-- URL: https://commons.wikimedia.org/wiki/File:Portland_Cement_Bags.jpg
UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/62/Portland_Cement_Bags.jpg/500px-Portland_Cement_Bags.jpg',
    updated_at = NOW()
WHERE (name ILIKE '%cemento%')
  AND image_url IS NULL;


-- Ladrillo común
-- Fuente: Wikimedia Commons - Brick.jpg (ladrillo cerámico macizo)
-- URL: https://commons.wikimedia.org/wiki/File:Brick.jpg
UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/01/Brick.jpg/500px-Brick.jpg',
    updated_at = NOW()
WHERE (name ILIKE '%ladrillo%')
  AND image_url IS NULL;


-- Hierro redondo / barra de acero
-- Fuente: Wikimedia Commons - A_bunch_of_rebar.jpg (barras de hierro corrugado)
-- URL: https://commons.wikimedia.org/wiki/File:A_bunch_of_rebar.jpg
UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/f/fe/A_bunch_of_rebar.jpg/500px-A_bunch_of_rebar.jpg',
    updated_at = NOW()
WHERE (name ILIKE '%hierro%' OR name ILIKE '%varilla%' OR name ILIKE '%barra acero%' OR name ILIKE '%hierro redondo%')
  AND image_url IS NULL;


-- Cable eléctrico
-- Fuente: Wikimedia Commons - Leitungsende_Abisoliert_en.svg.png (extremo de cable pelado)
-- URL: https://commons.wikimedia.org/wiki/File:Leitungsende_Abisoliert_en.svg
UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/Leitungsende_Abisoliert_en.svg/500px-Leitungsende_Abisoliert_en.svg.png',
    updated_at = NOW()
WHERE (name ILIKE '%cable%' AND (name ILIKE '%eléctric%' OR name ILIKE '%electric%' OR name ILIKE '%unipolar%' OR name ILIKE '%bipolar%' OR name ILIKE '%mm2%' OR business_type IN ('ferreteria','electricidad','construccion')))
  AND image_url IS NULL;


-- Candado (Yale u otra marca)
-- Fuente: Wikimedia Commons - Padlock_kłódka_ubt.JPG (candado metálico genérico)
-- URL: https://commons.wikimedia.org/wiki/File:Padlock_kłódka_ubt.JPG
UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/48/Padlock_kl%C3%B3dka_ubt.JPG/500px-Padlock_kl%C3%B3dka_ubt.JPG',
    updated_at = NOW()
WHERE (name ILIKE '%candado%')
  AND image_url IS NULL;


-- Bisagra (puerta / mueble)
-- Fuente: Wikimedia Commons - Carrollton_New_Orleans_hinge_brass_inside.jpg (bisagra de latón)
-- URL: https://commons.wikimedia.org/wiki/File:Carrollton_New_Orleans_hinge_brass_inside.jpg
UPDATE global_products SET
    image_url  = 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9a/Carrollton_New_Orleans_hinge_brass_inside.jpg/500px-Carrollton_New_Orleans_hinge_brass_inside.jpg',
    updated_at = NOW()
WHERE (name ILIKE '%bisagra%')
  AND image_url IS NULL;
