-- Seed 053: Electrodomésticos — 100 productos reales argentinos
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
-- HELADERAS (10 productos)
-- ============================================================
('Heladera No Frost 394L', 'Samsung', 'heladeras', 1250000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Heladera No Frost 382L', 'Whirlpool', 'heladeras', 1180000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Heladera No Frost 454L Inverter', 'LG', 'heladeras', 1450000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Heladera Cycle Defrost 334L', 'Electrolux', 'heladeras', 780000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Heladera No Frost 345L', 'Drean', 'heladeras', 850000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Heladera con Freezer 295L', 'Philco', 'heladeras', 720000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Heladera Under Counter 120L', 'Midea', 'heladeras', 420000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Heladera No Frost 363L', 'Siam', 'heladeras', 890000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Heladera No Frost Side by Side 541L', 'Samsung', 'heladeras', 2100000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Heladera No Frost 350L', 'BGH', 'heladeras', 830000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),

-- ============================================================
-- LAVARROPAS (10 productos)
-- ============================================================
('Lavarropas Automático 8kg 1200RPM', 'Drean', 'lavarropas', 750000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Lavarropas Automático 9kg Inverter', 'Samsung', 'lavarropas', 980000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Lavarropas Automático 8.5kg', 'Whirlpool', 'lavarropas', 820000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Lavarropas Automático 7kg', 'Electrolux', 'lavarropas', 680000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Lavarropas Automático 10kg', 'LG', 'lavarropas', 1100000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Lavarropas Semiautomático 7kg', 'Drean', 'lavarropas', 380000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Lavarropas Automático 6kg', 'Philco', 'lavarropas', 550000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Secarropas Centrífugo 5.5kg', 'Drean', 'lavarropas', 280000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Lavarropas Automático 8kg', 'Aurora', 'lavarropas', 620000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Lavarropa Semiautomático 5kg', 'Philco', 'lavarropas', 320000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),

-- ============================================================
-- COCINAS (10 productos)
-- ============================================================
('Cocina a Gas 56cm 4 Hornallas', 'Longvie', 'cocinas', 650000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Cocina a Gas 56cm Multigas', 'Orbis', 'cocinas', 580000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Cocina a Gas 56cm Acero Inoxidable', 'Siam', 'cocinas', 720000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Cocina a Gas 76cm 5 Hornallas', 'Longvie', 'cocinas', 950000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Anafe a Gas 4 Hornallas Empotrable', 'Orbis', 'cocinas', 350000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Anafe Eléctrico Vitrocerámico 4 Zonas', 'Whirlpool', 'cocinas', 580000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Horno Empotrable a Gas', 'Longvie', 'cocinas', 620000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Horno Empotrable Eléctrico 60L', 'Electrolux', 'cocinas', 750000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Cocina a Gas 56cm 4 Hornallas', 'Aurora', 'cocinas', 520000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Cocina Eléctrica Vitrocerámica 60cm', 'Electrolux', 'cocinas', 1100000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),

-- ============================================================
-- AIRES ACONDICIONADOS (10 productos)
-- ============================================================
('Aire Acondicionado Split 3200W Frío/Calor', 'Samsung', 'aires-acondicionados', 980000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Aire Acondicionado Split 2600W Inverter', 'LG', 'aires-acondicionados', 1150000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Aire Acondicionado Split 5200W Frío/Calor', 'Midea', 'aires-acondicionados', 1250000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Aire Acondicionado Split 3500W Inverter', 'BGH', 'aires-acondicionados', 1050000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Aire Acondicionado Split 2300W Frío Solo', 'Philco', 'aires-acondicionados', 720000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Aire Acondicionado Portátil 3500W', 'Atma', 'aires-acondicionados', 680000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Aire Acondicionado Split 2600W Frío/Calor', 'Noblex', 'aires-acondicionados', 780000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Aire Acondicionado Split 4500W Inverter', 'Samsung', 'aires-acondicionados', 1450000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Aire Acondicionado Ventana 3000W', 'BGH', 'aires-acondicionados', 550000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Aire Acondicionado Split 3200W Frío/Calor', 'Hisense', 'aires-acondicionados', 850000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),

-- ============================================================
-- TV Y AUDIO (10 productos)
-- ============================================================
('Smart TV LED 43" Full HD', 'Noblex', 'tv-audio', 520000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Smart TV LED 50" 4K UHD', 'Samsung', 'tv-audio', 850000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Smart TV LED 55" 4K UHD', 'LG', 'tv-audio', 980000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Smart TV LED 32" HD', 'Philco', 'tv-audio', 320000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Smart TV LED 50" 4K', 'TCL', 'tv-audio', 680000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Smart TV LED 55" 4K UHD', 'Hisense', 'tv-audio', 750000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Smart TV LED 65" 4K UHD', 'Samsung', 'tv-audio', 1550000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Soundbar 2.1 Bluetooth 300W', 'Samsung', 'tv-audio', 380000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Soundbar 2.0 Bluetooth 120W', 'Philips', 'tv-audio', 250000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Parlante Bluetooth Portátil 20W', 'Sony', 'tv-audio', 180000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),

-- ============================================================
-- PEQUEÑOS ELECTRODOMÉSTICOS - COCINA (12 productos)
-- ============================================================
('Microondas 20L Digital', 'BGH', 'pequenos-cocina', 280000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Microondas 28L con Grill', 'Samsung', 'pequenos-cocina', 420000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Microondas 23L Digital', 'Philco', 'pequenos-cocina', 250000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Licuadora de Pie 1.5L 600W', 'Philips', 'pequenos-cocina', 120000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Licuadora de Pie 2L 800W', 'Atma', 'pequenos-cocina', 85000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Batidora de Mano 500W', 'Philips', 'pequenos-cocina', 95000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Batidora Planetaria 4L 600W', 'Liliana', 'pequenos-cocina', 250000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Tostadora 2 Ranuras Acero Inox', 'Peabody', 'pequenos-cocina', 65000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Cafetera de Filtro 12 Tazas', 'Atma', 'pequenos-cocina', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Cafetera Express 15 Bar', 'Peabody', 'pequenos-cocina', 320000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Pava Eléctrica 1.7L Acero Inox', 'Philips', 'pequenos-cocina', 85000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Pava Eléctrica 1.7L', 'Liliana', 'pequenos-cocina', 48000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),

-- ============================================================
-- PEQUEÑOS ELECTRODOMÉSTICOS - LIMPIEZA (8 productos)
-- ============================================================
('Aspiradora Robot Wi-Fi', 'Samsung', 'pequenos-limpieza', 650000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Aspiradora Trineo 1600W', 'Philips', 'pequenos-limpieza', 280000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Aspiradora Vertical Inalámbrica', 'Midea', 'pequenos-limpieza', 350000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Aspiradora Ciclónica Sin Bolsa 1800W', 'Atma', 'pequenos-limpieza', 180000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Plancha a Vapor 2400W', 'Philips', 'pequenos-limpieza', 95000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Plancha a Vapor 2200W Cerámica', 'Liliana', 'pequenos-limpieza', 65000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Vaporizador Vertical 1500W', 'Philips', 'pequenos-limpieza', 120000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Plancha a Vapor 2000W', 'Ranser', 'pequenos-limpieza', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),

-- ============================================================
-- CALEFACCIÓN (10 productos)
-- ============================================================
('Estufa a Gas Infrarroja 3800 kcal', 'Eskabe', 'calefaccion', 280000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Estufa a Gas Tiro Balanceado 5000 kcal', 'Eskabe', 'calefaccion', 450000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Estufa a Gas Tiro Balanceado 3000 kcal', 'Orbis', 'calefaccion', 350000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Estufa a Gas Catalítica 3000 kcal', 'Longvie', 'calefaccion', 320000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Panel Calefactor 1000W', 'Liliana', 'calefaccion', 85000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Panel Calefactor 1400W con Termostato', 'Peabody', 'calefaccion', 120000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Caloventor 2000W', 'Atma', 'calefaccion', 55000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Caloventor 2000W Split', 'Liliana', 'calefaccion', 65000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Radiador Eléctrico 7 Elementos', 'Peabody', 'calefaccion', 95000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Radiador Eléctrico 9 Elementos', 'Ranser', 'calefaccion', 110000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),

-- ============================================================
-- VENTILACIÓN (10 productos)
-- ============================================================
('Ventilador de Pie 20" 3 Velocidades', 'Liliana', 'ventilacion', 75000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Ventilador de Pie 20" Turbo', 'Atma', 'ventilacion', 85000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Ventilador de Techo 3 Palas con Luz', 'Liliana', 'ventilacion', 150000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Ventilador Turbo 20" Reclinable', 'Peabody', 'ventilacion', 95000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Ventilador de Pie 16" Silencioso', 'Philips', 'ventilacion', 110000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Ventilador Turbo 16" Piso/Pared', 'Ranser', 'ventilacion', 65000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Ventilador de Techo 5 Palas', 'Peabody', 'ventilacion', 180000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Purificador de Aire con Filtro HEPA', 'Philips', 'ventilacion', 350000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Ventilador Industrial 20" Metal', 'Liliana', 'ventilacion', 120000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Climatizador Evaporativo Portátil', 'Midea', 'ventilacion', 280000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),

-- ============================================================
-- TECNOLOGÍA (10 productos)
-- ============================================================
('Notebook 15.6" Intel i5 8GB 256GB SSD', 'Samsung', 'tecnologia', 1200000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Notebook 14" Intel i3 8GB 512GB SSD', 'Noblex', 'tecnologia', 780000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Notebook 15.6" AMD Ryzen 5 16GB 512GB SSD', 'LG', 'tecnologia', 1450000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Tablet 10" 4GB 64GB Wi-Fi', 'Samsung', 'tecnologia', 450000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Tablet 10.1" 3GB 32GB', 'Philco', 'tecnologia', 280000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Celular 6.5" 128GB 4G', 'Samsung', 'tecnologia', 450000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Celular 6.7" 256GB 5G', 'Samsung', 'tecnologia', 1200000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Celular 6.1" 128GB 4G', 'TCL', 'tecnologia', 280000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Impresora Multifunción Wi-Fi', 'Epson', 'tecnologia', 350000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb),
('Impresora Láser Monocromática Wi-Fi', 'Samsung', 'tecnologia', 420000.00, 'seed', 0.5, 75, FALSE, TRUE, 'electrodomesticos', '{"unit": "unidad", "sku_prefix": "ELEC"}'::jsonb)

ON CONFLICT DO NOTHING;
