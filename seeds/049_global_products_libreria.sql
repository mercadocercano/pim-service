-- Seed 049: Librería — 120 productos reales argentinos (10 categorías)
-- Generado: 2026-04-18
-- Fuente: global_products (v2.0)
-- ON CONFLICT DO NOTHING: idempotente.

INSERT INTO global_products (
    name, brand, category, price,
    source, source_reliability, quality_score,
    is_verified, is_active, business_type, metadata
)
VALUES
-- ============================================================
-- CUADERNOS
-- ============================================================
('Cuaderno tapa dura Rivadavia 98 hojas rayado', 'Rivadavia', 'cuadernos', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cuaderno tapa dura Rivadavia 98 hojas cuadriculado', 'Rivadavia', 'cuadernos', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cuaderno tapa dura Rivadavia 48 hojas rayado', 'Rivadavia', 'cuadernos', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cuaderno tapa blanda Éxito 48 hojas rayado', 'Éxito', 'cuadernos', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cuaderno tapa blanda Éxito 84 hojas cuadriculado', 'Éxito', 'cuadernos', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cuaderno espiral A4 Ledesma 80 hojas rayado', 'Ledesma', 'cuadernos', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cuaderno espiral A4 Ledesma 120 hojas cuadriculado', 'Ledesma', 'cuadernos', 6200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cuaderno universitario Congreso 84 hojas rayado', 'Congreso', 'cuadernos', 3000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cuaderno universitario Gloria 82 hojas cuadriculado', 'Gloria', 'cuadernos', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cuaderno escolar Rivadavia 50 hojas liso', 'Rivadavia', 'cuadernos', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Repuesto hojas A4 Rivadavia 480 hojas rayado', 'Rivadavia', 'cuadernos', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cuaderno espiral Trabi A4 120 hojas rayado', 'Trabi', 'cuadernos', 5000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),

-- ============================================================
-- CARPETAS
-- ============================================================
('Carpeta N3 Rivadavia con ganchos', 'Rivadavia', 'carpetas', 6500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Carpeta N3 Éxito escolar', 'Éxito', 'carpetas', 4800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Carpeta oficio con ganchos Avios', 'Avios', 'carpetas', 5200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Carpeta A4 con elástico Ledesma', 'Ledesma', 'carpetas', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Bibliorato A4 lomo ancho Avios', 'Avios', 'carpetas', 8500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Bibliorato oficio lomo angosto Util-Of', 'Util-Of', 'carpetas', 7200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Carpeta colgante oficio Nepaco x25', 'Nepaco', 'carpetas', 12000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Carpeta con cierre A4 Filgo', 'Filgo', 'carpetas', 4000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),

-- ============================================================
-- ESCRITURA
-- ============================================================
('Lapicera BIC Cristal azul', 'BIC', 'escritura', 800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Lapicera BIC Cristal negra', 'BIC', 'escritura', 800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Lapicera BIC Cristal roja', 'BIC', 'escritura', 800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Lapicera Filgo Stick 026 azul', 'Filgo', 'escritura', 600.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Lapicera Simball Trimax azul', 'Simball', 'escritura', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Lápiz negro Faber Castell HB N2', 'Faber Castell', 'escritura', 500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Lápiz negro Staedtler Noris HB', 'Staedtler', 'escritura', 700.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Lápices de colores Faber Castell x12', 'Faber Castell', 'escritura', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Lápices de colores Faber Castell x24', 'Faber Castell', 'escritura', 8500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Lápices de colores Lyra Groove x12', 'Lyra', 'escritura', 6000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Marcador permanente Edding 400 negro', 'Edding', 'escritura', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Marcador permanente Edding 400 rojo', 'Edding', 'escritura', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Marcador p/ pizarra Filgo x4 colores', 'Filgo', 'escritura', 4000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Resaltador Filgo Flúo amarillo', 'Filgo', 'escritura', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Resaltador Stabilo Boss amarillo', 'Stabilo', 'escritura', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Resaltador Faber Castell Textliner pastel x6', 'Faber Castell', 'escritura', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Fibras Simball x10 colores', 'Simball', 'escritura', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Fibras Filgo punta fina x12', 'Filgo', 'escritura', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Microfibras Staedtler Triplus x10', 'Staedtler', 'escritura', 15000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),

-- ============================================================
-- UTILES ESCOLARES
-- ============================================================
('Regla 20cm Maped Twist n Flex', 'Maped', 'utiles-escolares', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Regla 30cm Pizzini acrílica', 'Pizzini', 'utiles-escolares', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Compás escolar Maped Stop System', 'Maped', 'utiles-escolares', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Compás Pizzini metálico escolar', 'Pizzini', 'utiles-escolares', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Tijera escolar Maped Vivo 13cm', 'Maped', 'utiles-escolares', 3000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Tijera Pizzini 17cm punta roma', 'Pizzini', 'utiles-escolares', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Goma de borrar Pelikan WS30', 'Pelikan', 'utiles-escolares', 600.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Goma de borrar Staedtler Mars Plastic', 'Staedtler', 'utiles-escolares', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Goma de borrar Faber Castell Dust-Free', 'Faber Castell', 'utiles-escolares', 900.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Sacapuntas Maped Croc Croc', 'Maped', 'utiles-escolares', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Sacapuntas metálico Faber Castell simple', 'Faber Castell', 'utiles-escolares', 800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Plasticola blanca 40g', 'Plasticola', 'utiles-escolares', 1000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Plasticola color x6 40g', 'Plasticola', 'utiles-escolares', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Set de geometría Pizzini 4 piezas', 'Pizzini', 'utiles-escolares', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Transportador Maped 180° 12cm', 'Maped', 'utiles-escolares', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),

-- ============================================================
-- PAPELERIA
-- ============================================================
('Resma A4 Ledesma autor 75g 500 hojas', 'Ledesma', 'papeleria', 12000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "resma", "sku_prefix": "LIB"}'::jsonb),
('Resma A4 Ledesma autor 80g 500 hojas', 'Ledesma', 'papeleria', 14000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "resma", "sku_prefix": "LIB"}'::jsonb),
('Resma oficio Ledesma 75g 500 hojas', 'Ledesma', 'papeleria', 13500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "resma", "sku_prefix": "LIB"}'::jsonb),
('Resma A4 Boreal 75g 500 hojas', 'Boreal', 'papeleria', 10500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "resma", "sku_prefix": "LIB"}'::jsonb),
('Papel glasé x10 hojas colores surtidos', 'Ledesma', 'papeleria', 800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cartulina escolar 50x70 blanca', 'Ledesma', 'papeleria', 500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cartulina escolar 50x70 colores surtidos', 'Ledesma', 'papeleria', 500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Papel afiche 74x110 colores surtidos', 'Ledesma', 'papeleria', 600.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Papel crepé 50x200cm colores surtidos', 'Ledesma', 'papeleria', 700.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Papel barrilete 50x70 colores surtidos', 'Ledesma', 'papeleria', 300.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Papel madera 70x100', 'Ledesma', 'papeleria', 400.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),

-- ============================================================
-- ARTE
-- ============================================================
('Acuarelas Pelikan x12 colores', 'Pelikan', 'arte', 6500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Acuarelas Alba x12 colores', 'Alba', 'arte', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Témperas Alba x6 colores 8ml', 'Alba', 'arte', 5000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Témperas Alba x10 colores 8ml', 'Alba', 'arte', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Pinceles Proarte set escolar x6', 'Proarte', 'arte', 4000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Pincel Proarte N8 pelo de cerda', 'Proarte', 'arte', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Crayones Pelikan Jumbo x12', 'Pelikan', 'arte', 4000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Crayones de cera Faber Castell x12', 'Faber Castell', 'arte', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Pasteles al óleo Simball x12', 'Simball', 'arte', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Plastilina alba x6 colores', 'Alba', 'arte', 3200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Masa para modelar Alba x6 colores', 'Alba', 'arte', 3800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Block de dibujo Rivadavia N5', 'Rivadavia', 'arte', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),

-- ============================================================
-- MOCHILAS
-- ============================================================
('Mochila escolar Samsonite 18" con bolsillos', 'Samsonite', 'mochilas', 85000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Mochila urbana Totto mediana', 'Totto', 'mochilas', 65000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Mochila escolar Xtrem Impact 20"', 'Xtrem', 'mochilas', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Mochila con carrito Topper escolar', 'Topper', 'mochilas', 75000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Mochila con carrito Xtrem Trolley', 'Xtrem', 'mochilas', 80000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Lonchera térmica Totto infantil', 'Totto', 'mochilas', 25000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Mochila urbana Samsonite laptop 15.6"', 'Samsonite', 'mochilas', 95000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cartuchera doble Totto estampada', 'Totto', 'mochilas', 18000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),

-- ============================================================
-- ADHESIVOS
-- ============================================================
('Cinta scotch 3M transparente 18mm x 30m', '3M', 'adhesivos', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cinta scotch 3M Magic 19mm x 33m', '3M', 'adhesivos', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cinta de papel 18mm x 50m', '3M', 'adhesivos', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cinta de embalar 48mm x 50m', '3M', 'adhesivos', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Voligoma 30ml', 'Voligoma', 'adhesivos', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Voligoma barra 21g', 'Voligoma', 'adhesivos', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('La Gotita adhesivo instantáneo 3g', 'La Gotita', 'adhesivos', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('UHU barra adhesiva 21g', 'UHU', 'adhesivos', 2800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('UHU barra adhesiva 40g', 'UHU', 'adhesivos', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Pritt barra adhesiva 22g', 'Pritt', 'adhesivos', 3000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),

-- ============================================================
-- ORGANIZACION
-- ============================================================
('Archivador revistero cartón Util-Of', 'Util-Of', 'organizacion', 2500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Sobre manila oficio x10', 'Ledesma', 'organizacion', 2000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Sobre blanco carta x10', 'Ledesma', 'organizacion', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Folios A4 x100 40 micrones', 'Util-Of', 'organizacion', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Folios oficio x100 cristal', 'Util-Of', 'organizacion', 5000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Separadores A4 x10 colores cartulina', 'Ledesma', 'organizacion', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Separadores A4 x5 plásticos', 'Util-Of', 'organizacion', 2200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Clips metálicos N2 x100', 'Util-Of', 'organizacion', 800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Clips mariposa N5 x12', 'Util-Of', 'organizacion', 1200.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Abrochadora Maped Essentials 26/6', 'Maped', 'organizacion', 5500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Broches Maped 26/6 x1000', 'Maped', 'organizacion', 1500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Sacabroches Maped', 'Maped', 'organizacion', 1800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Perforadora Maped Essentials 2 agujeros', 'Maped', 'organizacion', 6000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Bandejón acrílico triple apilable', 'Util-Of', 'organizacion', 8500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),

-- ============================================================
-- JUEGOS DIDACTICOS
-- ============================================================
('Rompecabezas 100 piezas infantil Ruibal', 'Ruibal', 'juegos-didacticos', 8500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Rompecabezas 500 piezas paisaje Ruibal', 'Ruibal', 'juegos-didacticos', 15000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Rompecabezas 1000 piezas Antex', 'Antex', 'juegos-didacticos', 22000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Libro infantil Cuentos Clásicos ilustrado', 'Sigmar', 'juegos-didacticos', 6500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Libro para colorear infantil A4 64 págs', 'Sigmar', 'juegos-didacticos', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Juego de mesa Carrera de Mente', 'Ruibal', 'juegos-didacticos', 28000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Juego de mesa Preguntados', 'Toyco', 'juegos-didacticos', 25000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Juego didáctico Abecedario magnético', 'Ruibal', 'juegos-didacticos', 12000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Memotest animales infantil', 'Ruibal', 'juegos-didacticos', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Juego de mesa Scrabble Junior', 'Mattel', 'juegos-didacticos', 32000.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb),
('Cuaderno tapa dura Gloria 98 hojas rayado', 'Gloria', 'cuadernos', 4800.00, 'seed', 0.5, 75, FALSE, TRUE, 'libreria', '{"unit": "unidad", "sku_prefix": "LIB"}'::jsonb)

ON CONFLICT DO NOTHING;
