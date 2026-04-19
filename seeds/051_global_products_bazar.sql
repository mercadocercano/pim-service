-- Seed 051: Bazar — 120 productos reales argentinos
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
-- COCINA (cocina)
-- ============================================================
('Olla acero inoxidable 24 cm', 'Tramontina', 'cocina', 45900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Olla acero inoxidable 20 cm', 'Tramontina', 'cocina', 38500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Olla a presión 6 litros', 'Magefesa', 'cocina', 89900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Cacerola antiadherente 24 cm', 'Tefal', 'cocina', 52900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Cacerola con tapa de vidrio 20 cm', 'Hudson', 'cocina', 28900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Sartén antiadherente 26 cm', 'Tefal', 'cocina', 42500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Sartén antiadherente 20 cm', 'Tramontina', 'cocina', 29900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Sartén de hierro 28 cm', 'Carol', 'cocina', 18500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Wok antiadherente 28 cm', 'Tramontina', 'cocina', 47900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Wok de acero inoxidable 30 cm', 'Hudson', 'cocina', 35900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Vaporera acero inoxidable 24 cm', 'Tramontina', 'cocina', 34900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Batería de cocina 7 piezas', 'Tramontina', 'cocina', 129900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),

-- ============================================================
-- CUBIERTOS (cubiertos)
-- ============================================================
('Juego de cubiertos 24 piezas', 'Tramontina', 'cubiertos', 42900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Juego de cubiertos 16 piezas', 'Brinox', 'cubiertos', 24900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Cuchillo de chef 8 pulgadas', 'Tramontina', 'cubiertos', 18900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Set de cuchillos 5 piezas con taco', 'Tramontina', 'cubiertos', 54900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Set de cuchillos para asado 6 piezas', 'Carol', 'cubiertos', 22500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Tenedor trinche asado', 'Tramontina', 'cubiertos', 8900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Cucharón de acero inoxidable', 'Tramontina', 'cubiertos', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Espumadera acero inoxidable', 'Brinox', 'cubiertos', 5900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Set asado 4 piezas con estuche', 'Tramontina', 'cubiertos', 67900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Set de cucharas de madera 3 piezas', 'Carol', 'cubiertos', 6500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Espátula de silicona', 'Tramontina', 'cubiertos', 5200.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),

-- ============================================================
-- VAJILLA (vajilla)
-- ============================================================
('Juego de platos 18 piezas', 'Luminarc', 'vajilla', 89900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Plato playo vidrio templado', 'Durax', 'vajilla', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Plato hondo vidrio templado', 'Durax', 'vajilla', 4200.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Plato de postre vidrio templado', 'Durax', 'vajilla', 3500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Taza de café con plato', 'Luminarc', 'vajilla', 6900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Tazón desayuno 500 ml', 'Luminarc', 'vajilla', 5900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Juego de vasos 6 unidades', 'Rigolleau', 'vajilla', 12900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Vaso trago largo 450 ml', 'Rigolleau', 'vajilla', 2900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Copa de vino cristal 350 ml', 'Rigolleau', 'vajilla', 4900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Juego de copas 6 unidades', 'Rigolleau', 'vajilla', 24900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Bowl vidrio templado 17 cm', 'Pyrex', 'vajilla', 8900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Set de bowls 3 piezas', 'Pyrex', 'vajilla', 29900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Fuente para horno rectangular', 'Pyrex', 'vajilla', 22900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),

-- ============================================================
-- ORGANIZACIÓN (organizacion)
-- ============================================================
('Caja organizadora con tapa 40 litros', 'Colombraro', 'organizacion', 18900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Caja organizadora con tapa 20 litros', 'Colombraro', 'organizacion', 12500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Canasto de mimbre grande', 'Drb', 'organizacion', 15900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Canasto organizador tela', 'Drb', 'organizacion', 9900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Perchero de pie metálico', 'Drb', 'organizacion', 34900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Perchero de pared 5 ganchos', 'Drb', 'organizacion', 12900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Ganchos adhesivos x6', 'Drb', 'organizacion', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Estante flotante 60 cm', 'Drb', 'organizacion', 14900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Zapatera 12 pares', 'Colombraro', 'organizacion', 22900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Organizador de cajón divisor', 'Colombraro', 'organizacion', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Porta bolsas de residuos', 'Colombraro', 'organizacion', 5900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),

-- ============================================================
-- DECORACIÓN (decoracion)
-- ============================================================
('Portarretratos madera 13x18 cm', 'Drb', 'decoracion', 7900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Portarretratos múltiple 6 fotos', 'Drb', 'decoracion', 18900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Vela aromática vaso vidrio', 'Drb', 'decoracion', 5900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Set de 3 velas cilíndricas', 'Drb', 'decoracion', 8900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Florero de vidrio 25 cm', 'Rigolleau', 'decoracion', 12900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Florero cerámica blanco', 'Drb', 'decoracion', 9900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Cuadro decorativo 40x60 cm', 'Drb', 'decoracion', 22900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Reloj de pared 30 cm', 'Drb', 'decoracion', 16900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Espejo redondo 50 cm', 'Drb', 'decoracion', 24900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Candelabro metálico 3 velas', 'Drb', 'decoracion', 14900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),

-- ============================================================
-- TEXTIL HOGAR (textil-hogar)
-- ============================================================
('Mantel antimanchas 150x200 cm', 'Carol', 'textil-hogar', 18900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Mantel PVC estampado 140x180 cm', 'Carol', 'textil-hogar', 12500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Repasador algodón pack x3', 'La Gauchita', 'textil-hogar', 7900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Repasador waffle pack x2', 'Carol', 'textil-hogar', 5900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Cortina de baño 180x180 cm', 'Carol', 'textil-hogar', 14900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Cortina blackout 150x220 cm', 'Drb', 'textil-hogar', 32900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Alfombra baño antideslizante', 'Carol', 'textil-hogar', 8900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Alfombra decorativa 120x170 cm', 'Drb', 'textil-hogar', 45900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Individual PVC pack x4', 'Carol', 'textil-hogar', 9900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Agarradera de cocina x2', 'La Gauchita', 'textil-hogar', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),

-- ============================================================
-- LIMPIEZA (limpieza)
-- ============================================================
('Balde con escurridor 12 litros', 'Vileda', 'limpieza', 22900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Balde plástico 10 litros', 'Colombraro', 'limpieza', 6900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Escoba de interior', 'Condor', 'limpieza', 8900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Escoba de exterior cerdas duras', 'Condor', 'limpieza', 9900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Trapo de piso algodón 48x60 cm', 'La Gauchita', 'limpieza', 3900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Trapo rejilla pack x3', 'La Gauchita', 'limpieza', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Secador de piso goma 40 cm', 'Condor', 'limpieza', 7900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Plumero microfibra extensible', 'Vileda', 'limpieza', 12900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Mopa microfibra con palo', 'Vileda', 'limpieza', 18900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Cepillo para inodoro con base', 'Colombraro', 'limpieza', 5900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Palita y escobillón set', 'Condor', 'limpieza', 6500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Guantes de limpieza par', 'Vileda', 'limpieza', 4900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),

-- ============================================================
-- ILUMINACIÓN (iluminacion)
-- ============================================================
('Lámpara de mesa pantalla tela', 'Drb', 'iluminacion', 24900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Lámpara colgante industrial', 'Drb', 'iluminacion', 32900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Velador de escritorio articulado', 'Drb', 'iluminacion', 28900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Velador de noche con pantalla', 'Drb', 'iluminacion', 19900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Guirnalda LED cálida 5 metros', 'Drb', 'iluminacion', 9900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Guirnalda LED bolitas 3 metros', 'Drb', 'iluminacion', 7900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Farol portavela metálico', 'Drb', 'iluminacion', 14900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Farol solar jardín pack x2', 'Drb', 'iluminacion', 12900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Aplique de pared bidireccional', 'Drb', 'iluminacion', 18900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Vela LED con timer', 'Drb', 'iluminacion', 5900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),

-- ============================================================
-- JARDÍN (jardin)
-- ============================================================
('Maceta plástica 25 cm', 'Colombraro', 'jardin', 5900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Maceta cerámica 20 cm', 'Drb', 'jardin', 9900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Maceta autoriego 30 cm', 'Colombraro', 'jardin', 14900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Regadera plástica 5 litros', 'Colombraro', 'jardin', 8900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Regadera metálica decorativa 2 litros', 'Drb', 'jardin', 12900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Set herramientas jardín 3 piezas', 'Tramontina', 'jardin', 18900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Pala de jardín mango corto', 'Tramontina', 'jardin', 7900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Tijera de podar', 'Tramontina', 'jardin', 14900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Manguera 15 metros con pistola', 'Colombraro', 'jardin', 29900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Soporte para maceta colgante', 'Drb', 'jardin', 6900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),

-- ============================================================
-- MESA (mesa)
-- ============================================================
('Panera de mimbre ovalada', 'Drb', 'mesa', 7900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Panera de tela y alambre', 'Carol', 'mesa', 5900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Ensaladera vidrio 23 cm', 'Pyrex', 'mesa', 14900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Ensaladera madera acacia', 'Tramontina', 'mesa', 24900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Tabla para picar bambú 38x28 cm', 'Tramontina', 'mesa', 16900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Tabla para asado madera 60x30 cm', 'Tramontina', 'mesa', 22900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Jarra de vidrio 1.5 litros', 'Rigolleau', 'mesa', 9900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Jarra térmica 1 litro', 'Silit', 'mesa', 18900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Termo Stanley clásico 1 litro', 'Stanley', 'mesa', 89900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Termo Stanley Adventure 750 ml', 'Stanley', 'mesa', 72900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Mate acero inoxidable Stanley', 'Stanley', 'mesa', 54900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Azucarera acero inoxidable', 'Tramontina', 'mesa', 8900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Salero y pimentero set', 'Tramontina', 'mesa', 7500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Aceitera vinagrera vidrio', 'Rigolleau', 'mesa', 6900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Bidet de baño portátil', 'Colombraro', 'limpieza', 4900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Escurreplatos de acero', 'Drb', 'cocina', 19900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Rallador 4 caras acero inoxidable', 'Tramontina', 'cocina', 8900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Porta rollo de cocina de pie', 'Drb', 'organizacion', 8900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Maceta colgante plástica 18 cm', 'Colombraro', 'jardin', 4500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb),
('Portavelas flotantes set x6', 'Drb', 'decoracion', 6900.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "set", "sku_prefix": "BAZ"}'::jsonb),
('Delantal de cocina algodón', 'La Gauchita', 'textil-hogar', 8500.00, 'seed', 0.5, 75, FALSE, TRUE, 'bazar', '{"unit": "unidad", "sku_prefix": "BAZ"}'::jsonb)
ON CONFLICT DO NOTHING;
