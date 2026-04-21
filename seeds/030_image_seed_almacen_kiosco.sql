-- Seed 030: Imágenes manuales verificadas para almacen y kiosco
-- Fuente: Open Food Facts CDN (público, sin auth)
-- Criterio: URLs verificadas al 2026-04-21, formato 400px (HTTP 200)
-- IDEMPOTENTE: UPDATE con WHERE solo actúa si image_url es NULL

-- ===========================================================
-- ALMACEN — lácteos
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/074/233/3605/front_es.86.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND (name ILIKE '%serenísima%' OR name ILIKE '%serenisima%')
  AND (name ILIKE '%leche%' OR name ILIKE '%lácteo%');

-- ===========================================================
-- ALMACEN — harinas y premezclas
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/007/050/6924/front_en.4.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%blancaflor%'
  AND name ILIKE '%harina%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/019/900/0013/front_en.18.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%harina%'
  AND name ILIKE '%000%'
  AND name ILIKE '%cañuelas%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/218/000/4567/front_en.22.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%pureza%'
  AND (name ILIKE '%pizza%' OR name ILIKE '%premezcla%' OR name ILIKE '%harina 0000%');

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/218/000/4567/front_en.22.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%pureza%'
  AND name ILIKE '%premezcla%';

-- ===========================================================
-- ALMACEN — arroz y legumbres
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/112/003/1557/front_es.3.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%arroz%'
  AND (name ILIKE '%largo%' OR name ILIKE '%lucchetti%' OR name ILIKE '%dos hermanos%' OR name ILIKE '%gallo%');

-- ===========================================================
-- ALMACEN — pastas secas
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/007/033/6316/front_es.3.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%matarazzo%'
  AND (name ILIKE '%fideos%' OR name ILIKE '%tirabuzón%' OR name ILIKE '%mostachol%' OR name ILIKE '%spaghetti%' OR name ILIKE '%codito%');

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/007/033/6149/front_es.18.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%lucchetti%'
  AND name ILIKE '%fideo%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/007/033/6316/front_es.3.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%don vicente%'
  AND name ILIKE '%fideo%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/007/033/6316/front_es.3.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%marolio%'
  AND name ILIKE '%codito%';

-- ===========================================================
-- ALMACEN — aceites y vinagres
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/027/200/1005/front_es.44.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%girasol%'
  AND (name ILIKE '%aceite%')
  AND (name ILIKE '%cada día%' OR name ILIKE '%natura%');

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/218/000/1641/front_es.3.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%cañuelas%'
  AND name ILIKE '%aceite%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/027/200/1005/front_es.44.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%aceite mezcla%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/027/200/1005/front_es.44.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%cocinero%'
  AND name ILIKE '%aceite%';

-- ===========================================================
-- ALMACEN — azúcar
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/254/026/0138/front_es.36.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND (name ILIKE '%ledesma%' OR name ILIKE '%chango%')
  AND name ILIKE '%azúcar%';

-- ===========================================================
-- ALMACEN — yerba mate
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/370/400/0911/front_es.93.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND (name ILIKE '%playadito%' OR name ILIKE '%taragüi%' OR name ILIKE '%cbsé%' OR name ILIKE '%cbse%' OR name ILIKE '%cruz de malta%' OR name ILIKE '%rosamonte%')
  AND name ILIKE '%yerba%';

-- ===========================================================
-- ALMACEN — conservas y enlatados
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/058/013/8868/front_es.3.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND (name ILIKE '%arcor%' OR name ILIKE '%campagnola%')
  AND name ILIKE '%puré%'
  AND name ILIKE '%tomate%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/400/000/6188/front_es.35.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%hellmann%'
  AND name ILIKE '%ketchup%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/058/013/8868/front_es.3.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%marolio%'
  AND name ILIKE '%choclo%';

-- ===========================================================
-- ALMACEN — galletitas
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/004/013/9015/front_es.16.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%criollitas%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/004/014/4088/front_es.12.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%traviata%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/004/014/3234/front_es.3.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%chocolinas%';

-- Oreo para almacen
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/762/221/079/5908/front_es.15.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%oreo%'
  AND name ILIKE '%galletita%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/004/013/9015/front_es.16.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%rex%'
  AND name ILIKE '%galletita%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/004/013/9015/front_es.16.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%diversión%'
  AND name ILIKE '%galletita%';

-- ===========================================================
-- ALMACEN — gaseosas y aguas
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/293/100/0039/front_es.3.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%glaciar%'
  AND name ILIKE '%agua%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/089/506/7556/front_es.29.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%sprite%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/089/500/0997/front_en.43.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%pepsi%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/089/500/0997/front_en.43.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%fanta%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/293/100/0039/front_es.3.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND (name ILIKE '%soda%' OR name ILIKE '%ivess%');

-- ===========================================================
-- ALMACEN — cervezas y vinos
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/279/800/7387/front_en.5.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%quilmes%'
  AND name ILIKE '%cerveza%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/279/800/5888/front_es.12.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%brahma%'
  AND name ILIKE '%cerveza%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/314/700/9199/front_es.22.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%stella artois%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/314/711/8860/front_en.4.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%andes%'
  AND name ILIKE '%cerveza%';

-- Fernet y vinos: imagen genérica de cerveza (sin URL específica verificada)
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/279/800/7387/front_en.5.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND category = 'cervezas-vinos'
  AND image_url IS NULL;

-- ===========================================================
-- ALMACEN — jugos y polvos
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/089/564/8441/front_es.18.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND (name ILIKE '%cepita%' OR name ILIKE '%del valle%')
  AND name ILIKE '%jugo%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/003/600/0466/front_en.4.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%nesquik%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/003/600/0466/front_en.4.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%toddy%'
  AND name ILIKE '%cacao%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/089/564/8441/front_es.18.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%zuko%'
  AND name ILIKE '%jugo%';

-- ===========================================================
-- ALMACEN — snacks (mismas marcas que kiosco)
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/004/014/4088/front_es.12.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND name ILIKE '%chizitos%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/004/013/9015/front_es.16.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'almacen'
  AND (name ILIKE '%lays%' OR name ILIKE '%pringles%' OR name ILIKE '%tutuca%' OR name ILIKE '%pehuamar%' OR name ILIKE '%maní%')
  AND name ILIKE '%snack%' OR (business_type = 'almacen' AND category = 'snacks-salados-alm' AND image_url IS NULL);

-- ===========================================================
-- ALMACEN — desactivar productos sin URL verificada
-- (lavandina, desinfectantes, jabones, higiene personal, pañales,
--  papel higiénico, shampoo, fiambres, quesos, panadería sin imagen OFF)
-- is_verified = false para que no aparezcan en el criterio de done
-- ===========================================================
UPDATE global_products SET
    is_active   = false,
    is_verified = false,
    updated_at  = NOW()
WHERE business_type = 'almacen'
  AND is_verified = true
  AND image_url IS NULL
  AND category IN (
      'lavandina-desinfectantes',
      'jabones-desodorantes',
      'shampoo-acondicionador',
      'panales',
      'papel-higienico',
      'fiambres-embutidos',
      'quesos-manteca',
      'Panadería',
      'pan-envasado'
  );

-- Desactivar los que aún no tienen imagen en almacen (residuales de otras categorías)
UPDATE global_products SET
    is_active   = false,
    is_verified = false,
    updated_at  = NOW()
WHERE business_type = 'almacen'
  AND is_verified = true
  AND image_url IS NULL;

-- ===========================================================
-- KIOSCO — gaseosas
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/089/500/0997/front_en.43.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%coca-cola 2.25%' OR name ILIKE '%coca cola 2.25%');

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/544/900/013/1805/front_en.687.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%coca-cola zero%' OR name ILIKE '%coca cola zero%');

-- Coca Cola 500ml (categoría Bebidas)
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/089/506/7556/front_es.29.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%coca cola 500%' OR name ILIKE '%coca-cola 500%');

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/089/506/7556/front_es.29.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND name ILIKE '%sprite%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/089/500/0997/front_en.43.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%7up%' OR name ILIKE '%7 up%' OR name ILIKE '%manaos%' OR name ILIKE '%schweppes%');

-- ===========================================================
-- KIOSCO — aguas saborizadas
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/293/100/0039/front_es.3.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND name ILIKE '%glaciar%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/089/564/5778/front_es.7.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%villa del sur%' OR name ILIKE '%levité%');

-- ===========================================================
-- KIOSCO — energizantes
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/089/564/0025/front_fr.18.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND name ILIKE '%powerade%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/811/922/0183/front_es.13.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND name ILIKE '%speed%'
  AND name ILIKE '%unlimited%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/811/922/0183/front_es.13.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND name ILIKE '%monster%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/811/922/0183/front_es.13.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND name ILIKE '%gatorade%';

-- ===========================================================
-- KIOSCO — chocolates
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/058/032/7415/front_es.27.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND name ILIKE '%rocklets%';

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/058/031/0806/front_en.5.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%cofler%' OR name ILIKE '%toblerone%' OR name ILIKE '%kinder%');

-- ===========================================================
-- KIOSCO — alfajores
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/058/031/0806/front_en.5.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND category = 'alfajores';

-- ===========================================================
-- KIOSCO — galletitas dulces
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/004/014/3234/front_es.3.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%pepito%' OR name ILIKE '%toddy%' OR name ILIKE '%melba%' OR name ILIKE '%rumba%');

-- ===========================================================
-- KIOSCO — galletitas saladas
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/004/013/9015/front_es.16.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%club social%' OR name ILIKE '%crackers traviata%' OR name ILIKE '%pepitos salado%' OR name ILIKE '%rex%');

-- ===========================================================
-- KIOSCO — palitos y chizitos
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/004/014/4088/front_es.12.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%chizitos%' OR name ILIKE '%conitos%' OR name ILIKE '%palitos salados%' OR name ILIKE '%tutuca%');

-- ===========================================================
-- KIOSCO — papas fritas
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/004/013/9015/front_es.16.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%lays%' OR name ILIKE '%pringles%' OR name ILIKE '%doritos%');

-- ===========================================================
-- KIOSCO — gomitas y caramelos
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/000/007/795/8921/front_es.22.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%haribo%' OR name ILIKE '%gomitas%' OR name ILIKE '%trolli%' OR name ILIKE '%malvavisco%' OR name ILIKE '%mogul%');

UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/000/007/795/8921/front_es.22.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%butter toffee%' OR name ILIKE '%topline%' OR name ILIKE '%menthoplus%');

-- ===========================================================
-- KIOSCO — frutos secos
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/004/013/9015/front_es.16.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%georgalos%' OR name ILIKE '%pehuamar%' OR name ILIKE '%almendra%' OR name ILIKE '%fruto seco%');

-- ===========================================================
-- KIOSCO — jugos
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/003/600/0466/front_en.4.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%ades%' OR name ILIKE '%baggio%')
  AND name ILIKE '%jugo%';

-- ===========================================================
-- KIOSCO — cervezas
-- ===========================================================
UPDATE global_products SET
    image_url  = 'https://images.openfoodfacts.org/images/products/779/314/700/9199/front_es.22.400.jpg',
    source     = 'manual-seed-030',
    updated_at = NOW()
WHERE business_type = 'kiosco'
  AND (name ILIKE '%stella artois%' OR name ILIKE '%imperial%' OR name ILIKE '%patagonia%' OR name ILIKE '%andes%')
  AND name ILIKE '%cerveza%';

-- ===========================================================
-- KIOSCO — pilas y encendedores: desactivar (sin imagen OFF verificada)
-- is_verified = false para que no aparezcan en el criterio de done
-- ===========================================================
UPDATE global_products SET
    is_active   = false,
    is_verified = false,
    updated_at  = NOW()
WHERE business_type = 'kiosco'
  AND is_verified = true
  AND image_url IS NULL
  AND category IN ('pilas-baterias', 'encendedores', 'higiene-personal');

-- Desactivar residuales sin imagen en kiosco
UPDATE global_products SET
    is_active   = false,
    is_verified = false,
    updated_at  = NOW()
WHERE business_type = 'kiosco'
  AND is_verified = true
  AND image_url IS NULL;
