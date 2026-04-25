-- =============================================================
-- SEED 117: Casa de Limpieza (business_type = 'limpieza')
-- Research date: 2026-04-25
-- Source: Catálogo market research NEA/Posadas 2026
--   - Distribuidoras mayoristas Posadas (Zona Norte, La Comercial)
--   - Precios de góndola relevados abril 2026
-- Zona: Posadas, Misiones (NEA)
-- Objetivo: >= 100 productos verificados
-- Categorías: limpieza, pequenos-limpieza, papel-higiene, cloro-quimicos
-- =============================================================

-- ============================================================
-- CATEGORÍA: limpieza — Productos de limpieza del hogar
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Lavandinas
('Lavandina concentrada 2.2L Ayudín', 'Ayudín', 'limpieza', 'limpieza', 1650.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Lavandina concentrada 1L Ayudín', 'Ayudín', 'limpieza', 'limpieza', 950.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Lavandina clásica 2L Ayudín', 'Ayudín', 'limpieza', 'limpieza', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Lavandina concentrada 2.2L Bianchi', 'Bianchi', 'limpieza', 'limpieza', 1500.00, 'catalog-researcher', 0.80, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Lavandina concentrada 1L Bianchi', 'Bianchi', 'limpieza', 'limpieza', 850.00, 'catalog-researcher', 0.80, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Lavandina concentrada 2L Cierto', 'Cierto', 'limpieza', 'limpieza', 1350.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Lavandina regular 5L bidon', NULL, 'limpieza', 'limpieza', 2200.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "bidon", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Desinfectantes y limpiadores de piso
('Limpiador de pisos pino 1L Poett', 'Poett', 'limpieza', 'limpieza', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Limpiador de pisos lavanda 1L Poett', 'Poett', 'limpieza', 'limpieza', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Limpiador de pisos pino 2L Poett', 'Poett', 'limpieza', 'limpieza', 2100.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Limpiador de pisos clásico 1L Cif', 'Cif', 'limpieza', 'limpieza', 1300.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Desinfectante multiuso lavanda 1L Lysoform', 'Lysoform', 'limpieza', 'limpieza', 1450.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Desinfectante multiuso pino 1L Lysoform', 'Lysoform', 'limpieza', 'limpieza', 1450.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Desinfectante pisos y superficies 1.8L Mr. Músculo', 'Mr. Músculo', 'limpieza', 'limpieza', 1700.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Limpiador desinfectante Baby 750ml Lysoform', 'Lysoform', 'limpieza', 'limpieza', 1100.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Desengrasantes
('Desengrasante cocina spray 500ml Mr. Músculo', 'Mr. Músculo', 'limpieza', 'limpieza', 1400.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Desengrasante cocina spray 500ml Cif', 'Cif', 'limpieza', 'limpieza', 1350.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Limpiavidrios spray 500ml Mr. Músculo', 'Mr. Músculo', 'limpieza', 'limpieza', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Limpiavidrios spray 500ml Cif', 'Cif', 'limpieza', 'limpieza', 1150.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Limpiador baño spray 500ml Cif crema', 'Cif', 'limpieza', 'limpieza', 1300.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Quitamanchas ropa líquido 1L Vanish', 'Vanish', 'limpieza', 'limpieza', 2500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Quitamanchas ropa en polvo 450g Vanish', 'Vanish', 'limpieza', 'limpieza', 1800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Detergentes
('Detergente lavavajillas limón 750ml Magistral', 'Magistral', 'limpieza', 'limpieza', 850.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Detergente lavavajillas limón 500ml Magistral', 'Magistral', 'limpieza', 'limpieza', 620.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Detergente lavavajillas 750ml Sunlight', 'Sunlight', 'limpieza', 'limpieza', 900.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Detergente lavavajillas 500ml Limpol', 'Limpol', 'limpieza', 'limpieza', 750.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Jabones en polvo y suavizantes
('Jabón en polvo 3kg Skip', 'Skip', 'limpieza', 'limpieza', 4500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "bolsa", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Jabón en polvo 1.5kg Skip', 'Skip', 'limpieza', 'limpieza', 2500.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "bolsa", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Jabón en polvo 3kg Ala', 'Ala', 'limpieza', 'limpieza', 3800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "bolsa", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Jabón en polvo 1.5kg Ala', 'Ala', 'limpieza', 'limpieza', 2100.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "bolsa", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Jabón en polvo 3kg Drive', 'Drive', 'limpieza', 'limpieza', 3500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "bolsa", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Suavizante ropa 900ml Comfort', 'Comfort', 'limpieza', 'limpieza', 1800.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Suavizante ropa 1.8L Comfort', 'Comfort', 'limpieza', 'limpieza', 3200.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Suavizante ropa 900ml Vivere', 'Vivere', 'limpieza', 'limpieza', 1600.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============================================================
-- CATEGORÍA: pequenos-limpieza — Utensilios y accesorios
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

-- Trapos, repasadores, franelas
('Trapo de piso absorbente 60x80cm Patito', 'Patito', 'limpieza', 'pequenos-limpieza', 850.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Trapo de piso algodón 70x80cm Vileda', 'Vileda', 'limpieza', 'pequenos-limpieza', 1200.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Repasador de cocina x3 unidades genérico', NULL, 'limpieza', 'pequenos-limpieza', 900.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "pack 3u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Franela multiuso amarilla 40x40cm Vileda', 'Vileda', 'limpieza', 'pequenos-limpieza', 750.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Franela multiuso pack x5 Patito', 'Patito', 'limpieza', 'pequenos-limpieza', 2500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "pack 5u", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Escobas, lampazo, cepillos
('Escoba plástica con mango Patito', 'Patito', 'limpieza', 'pequenos-limpieza', 1800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Escoba cerdas suaves con palo Vileda', 'Vileda', 'limpieza', 'pequenos-limpieza', 2200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Lampazo rectangular con palo Patito', 'Patito', 'limpieza', 'pequenos-limpieza', 2500.00, 'catalog-researcher', 0.85, 62, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Lampazo oval doble faz con palo Vileda', 'Vileda', 'limpieza', 'pequenos-limpieza', 3200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cepillo de fregar para platos', NULL, 'limpieza', 'pequenos-limpieza', 450.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cepillo de fregar con mango largo', NULL, 'limpieza', 'pequenos-limpieza', 700.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Balde plástico 10L con pico', NULL, 'limpieza', 'pequenos-limpieza', 1200.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Balde plástico 15L escurridor Vileda', 'Vileda', 'limpieza', 'pequenos-limpieza', 3500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Guantes de goma talle M Vileda', 'Vileda', 'limpieza', 'pequenos-limpieza', 900.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "par", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Guantes de goma talle G Vileda', 'Vileda', 'limpieza', 'pequenos-limpieza', 900.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "par", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Guantes de goma talle M Descarpack', 'Descarpack', 'limpieza', 'pequenos-limpieza', 650.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "par", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Esponjas y fibras
('Esponja verde/amarilla x3 unidades 3M', '3M', 'limpieza', 'pequenos-limpieza', 650.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "pack 3u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Esponja doble faz x2 Scotch-Brite', 'Scotch-Brite', 'limpieza', 'pequenos-limpieza', 700.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "pack 2u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Virulana lana acero x3 Vileda', 'Vileda', 'limpieza', 'pequenos-limpieza', 450.00, 'catalog-researcher', 0.85, 62, TRUE, TRUE, '{"unit": "pack 3u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Fibra de aluminio x3 genérico', NULL, 'limpieza', 'pequenos-limpieza', 350.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "pack 3u", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Bolsas de basura
('Bolsas de basura negras 60x90cm x30 unidades Polibag', 'Polibag', 'limpieza', 'pequenos-limpieza', 800.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "paquete 30u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Bolsas de basura negras 80x110cm x20 unidades Polibag', 'Polibag', 'limpieza', 'pequenos-limpieza', 900.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "paquete 20u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Bolsas de basura perfumadas 60x90cm x20 unidades', NULL, 'limpieza', 'pequenos-limpieza', 750.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "paquete 20u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Bolsas para heladera x100 unidades', NULL, 'limpieza', 'pequenos-limpieza', 500.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "paquete 100u", "zona": "Posadas", "fecha_precio": "2026-04"}'),

-- Plumeros e implementos varios
('Plumero de plástico largo 80cm', NULL, 'limpieza', 'pequenos-limpieza', 950.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Plumero clásico de avestruz sintético', NULL, 'limpieza', 'pequenos-limpieza', 1200.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Recogedor de basura con mango', NULL, 'limpieza', 'pequenos-limpieza', 600.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Portarollo papel higiénico plástico blanco', NULL, 'limpieza', 'pequenos-limpieza', 500.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============================================================
-- CATEGORÍA: papel-higiene — Papel de cocina, higiénico, servilletas
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

('Papel higiénico doble hoja x4 rollos Elite', 'Elite', 'limpieza', 'papel-higiene', 1200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "pack 4u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Papel higiénico doble hoja x8 rollos Elite', 'Elite', 'limpieza', 'papel-higiene', 2200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "pack 8u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Papel higiénico doble hoja x12 rollos Elite', 'Elite', 'limpieza', 'papel-higiene', 3200.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "pack 12u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Papel higiénico doble hoja x4 rollos Higienol', 'Higienol', 'limpieza', 'papel-higiene', 1100.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "pack 4u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Papel higiénico doble hoja x8 rollos Higienol', 'Higienol', 'limpieza', 'papel-higiene', 2000.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "pack 8u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Papel higiénico doble hoja x12 rollos Higienol', 'Higienol', 'limpieza', 'papel-higiene', 2900.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "pack 12u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Papel de cocina rollo doble hoja x2 Elite', 'Elite', 'limpieza', 'papel-higiene', 900.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "pack 2u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Papel de cocina rollo doble hoja x4 Elite', 'Elite', 'limpieza', 'papel-higiene', 1700.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "pack 4u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Papel de cocina rollo x2 Higienol', 'Higienol', 'limpieza', 'papel-higiene', 800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "pack 2u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Servilletas de papel blancas x100 Elite', 'Elite', 'limpieza', 'papel-higiene', 600.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "paquete 100u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Servilletas de papel blancas x50 Higienol', 'Higienol', 'limpieza', 'papel-higiene', 350.00, 'catalog-researcher', 0.85, 62, TRUE, TRUE, '{"unit": "paquete 50u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pañuelos descartables caja x60 Elite', 'Elite', 'limpieza', 'papel-higiene', 700.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "caja", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pañuelos descartables bolsillo x10 Elite', 'Elite', 'limpieza', 'papel-higiene', 200.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{"unit": "paquete 10u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Film estirable cocina 30cm x30m', NULL, 'limpieza', 'papel-higiene', 450.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "rollo", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Papel aluminio cocina 30cm x10m', NULL, 'limpieza', 'papel-higiene', 600.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "rollo", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============================================================
-- CATEGORÍA: cloro-quimicos — Hipoclorito, desinfectantes especiales
-- ============================================================
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES

('Hipoclorito de sodio 1L concentrado industrial', NULL, 'limpieza', 'cloro-quimicos', 750.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "litro", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Hipoclorito de sodio 5L bidon industrial', NULL, 'limpieza', 'cloro-quimicos', 3200.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "bidon 5L", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cloro en pastillas piscina 200g x10 unidades', NULL, 'limpieza', 'cloro-quimicos', 2800.00, 'catalog-researcher', 0.80, 55, TRUE, TRUE, '{"unit": "paquete 10u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Desodorante de ambiente aerosol 360ml Glade', 'Glade', 'limpieza', 'cloro-quimicos', 1500.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Desodorante de ambiente aerosol 360ml Breeze', 'Breeze', 'limpieza', 'cloro-quimicos', 1200.00, 'catalog-researcher', 0.80, 62, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Desodorante de ambiente eléctrico Glade recarga', 'Glade', 'limpieza', 'cloro-quimicos', 1800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "recarga", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pastilla de inodoro 50g Harpic', 'Harpic', 'limpieza', 'cloro-quimicos', 650.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Gel limpiador inodoro 750ml Harpic', 'Harpic', 'limpieza', 'cloro-quimicos', 1300.00, 'catalog-researcher', 0.85, 68, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Pastilla de inodoro tanque 50g', NULL, 'limpieza', 'cloro-quimicos', 400.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Destapa cañerías gel 500ml Destapol', 'Destapol', 'limpieza', 'cloro-quimicos', 1100.00, 'catalog-researcher', 0.80, 60, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Insecticida aerosol cucarachas y hormigas 360ml Raid', 'Raid', 'limpieza', 'cloro-quimicos', 1800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Insecticida aerosol mosquitos 360ml Raid', 'Raid', 'limpieza', 'cloro-quimicos', 1800.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Espiral antimosquitos x12 unidades', NULL, 'limpieza', 'cloro-quimicos', 900.00, 'catalog-researcher', 0.75, 55, TRUE, TRUE, '{"unit": "caja 12u", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Repelente eléctrico líquido recarga 45ml Raid', 'Raid', 'limpieza', 'cloro-quimicos', 1200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "recarga", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Naftalina bolas x100g', NULL, 'limpieza', 'cloro-quimicos', 350.00, 'catalog-researcher', 0.75, 50, TRUE, TRUE, '{"unit": "paquete 100g", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cera para piso líquida 900ml Johnson', 'Johnson', 'limpieza', 'cloro-quimicos', 1600.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}'),
('Cera para muebles spray 300ml Blem', 'Blem', 'limpieza', 'cloro-quimicos', 1400.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{"unit": "unidad", "zona": "Posadas", "fecha_precio": "2026-04"}')
ON CONFLICT (name, business_type) DO NOTHING;
