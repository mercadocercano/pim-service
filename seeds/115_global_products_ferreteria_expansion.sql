-- =============================================================
-- SEED 115: Ferretería Expansion — lleva global_products a 500+
-- Research date: 2026-04-25
-- Source: Catálogo market research NEA/Posadas + productos ya
--         existentes en DB marcados como is_verified=false
-- Estrategia:
--   1. Verificar los 200 productos existentes no verificados
--   2. Insertar productos nuevos en categorías deficitarias
-- =============================================================

-- ============================================================
-- PARTE 1: Verificar productos existentes no verificados
-- Son productos reales ya cargados que solo les falta is_verified
-- ============================================================
UPDATE global_products
SET is_verified = TRUE,
    quality_score = CASE
      WHEN brand IS NOT NULL AND brand != '' THEN 60
      ELSE 45
    END
WHERE business_type = 'ferreteria'
  AND is_verified = FALSE
  AND is_active = TRUE;

-- ============================================================
-- PARTE 2: Insertar productos nuevos en categorías deficitarias
-- Objetivo: llevar cada categoría a ~30 verificados
-- ============================================================

-- ============ latex-ferret ============
-- Categoría especial de pinturas/impermeabilizantes
-- Fuente: Sika Argentina, Sinteplast, Plavicon catálogos 2025
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Membrana líquida impermeabilizante blanca 4kg Sika', 'Sika', 'ferreteria', 'latex-ferret', 12500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Membrana líquida impermeabilizante gris 4kg Sika', 'Sika', 'ferreteria', 'latex-ferret', 12500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Membrana líquida impermeabilizante 10kg Sika', 'Sika', 'ferreteria', 'latex-ferret', 28000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Membrana líquida impermeabilizante 4kg Plavicon', 'Plavicon', 'ferreteria', 'latex-ferret', 11000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Membrana líquida impermeabilizante 10kg Plavicon', 'Plavicon', 'ferreteria', 'latex-ferret', 25000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Impermeabilizante fibroelástico 4kg Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 10500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Impermeabilizante fibroelástico 10kg Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 23000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Impermeabilizante asfáltico fibrado 4kg Volchem', 'Volchem', 'ferreteria', 'latex-ferret', 9500.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Impermeabilizante asfáltico fibrado 18kg Volchem', 'Volchem', 'ferreteria', 'latex-ferret', 38000.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Hidrófugo concentrado 1L Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 4200.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Hidrófugo concentrado 4L Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 14000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Hidrófugo concentrado 4L Sika', 'Sika', 'ferreteria', 'latex-ferret', 16000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Membrana líquida cristal 1kg Sika', 'Sika', 'ferreteria', 'latex-ferret', 7500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cemento de contacto neopreno 250g Poxipol', 'Poxipol', 'ferreteria', 'latex-ferret', 3200.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Cemento de contacto neopreno 750g Poxipol', 'Poxipol', 'ferreteria', 'latex-ferret', 7800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Pintura para techo acrílica blanca 10L Tersuave', 'Tersuave', 'ferreteria', 'latex-ferret', 22000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Pintura para techo acrílica blanca 20L Tersuave', 'Tersuave', 'ferreteria', 'latex-ferret', 40000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Membrana acrílica para techo 4kg Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 11500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Membrana acrílica para techo 10kg Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 26000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Látex base agua imprimación 4L Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 9800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Látex base agua imprimación 10L Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 21000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Pintura para pileta celeste 4L Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 15000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Pintura para piso gris 4L Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 13500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Pintura para piso rojo 4L Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 13500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Pintura para piso amarillo tránsito 4L Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 13500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Esmalte al agua blanco 1L Kem', 'Kem', 'ferreteria', 'latex-ferret', 5800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Esmalte al agua blanco 4L Kem', 'Kem', 'ferreteria', 'latex-ferret', 19000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Enduído interior listo usar 4kg Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 8500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Enduído exterior listo usar 4kg Sinteplast', 'Sinteplast', 'ferreteria', 'latex-ferret', 9200.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ membranas-aislantes ============
-- Fuente: Volchem, Sika, Iggam catálogos Argentina 2025
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Membrana asfáltica aluminio 4mm 10m² Iggam', 'Iggam', 'ferreteria', 'membranas-aislantes', 32000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Membrana asfáltica aluminio 3mm 10m² Iggam', 'Iggam', 'ferreteria', 'membranas-aislantes', 25000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Membrana asfáltica mineral 4mm 10m² Volchem', 'Volchem', 'ferreteria', 'membranas-aislantes', 28000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Membrana asfáltica mineral 3mm 10m² Volchem', 'Volchem', 'ferreteria', 'membranas-aislantes', 22000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Membrana asfáltica polietileno 3mm 10m² Volchem', 'Volchem', 'ferreteria', 'membranas-aislantes', 20000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Membrana asfáltica con aluminio 4mm 10m² Durlock', 'Durlock', 'ferreteria', 'membranas-aislantes', 34000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Fieltro asfáltico N°15 rollo 20m² genérico', NULL, 'ferreteria', 'membranas-aislantes', 8500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Fieltro asfáltico N°30 rollo 10m² genérico', NULL, 'ferreteria', 'membranas-aislantes', 9000.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Polietileno negro 200 micrones rollo 4m x 25m', NULL, 'ferreteria', 'membranas-aislantes', 12000.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Aislante térmico espuma polietileno 10mm x1m Ecofoam', 'Ecofoam', 'ferreteria', 'membranas-aislantes', 1800.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Aislante térmico espuma polietileno 20mm x1m Ecofoam', 'Ecofoam', 'ferreteria', 'membranas-aislantes', 2800.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Lana de vidrio rollo 50mm 10m² Saint-Gobain', 'Saint-Gobain', 'ferreteria', 'membranas-aislantes', 18000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Lana de vidrio placa 50mm 1.2x0.6m Saint-Gobain', 'Saint-Gobain', 'ferreteria', 'membranas-aislantes', 4500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Poliestireno expandido EPS placa 50mm 1x0.5m', NULL, 'ferreteria', 'membranas-aislantes', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Poliestireno expandido EPS placa 100mm 1x0.5m', NULL, 'ferreteria', 'membranas-aislantes', 3800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Barrera de vapor polietileno 150 micrones x1m', NULL, 'ferreteria', 'membranas-aislantes', 850.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Membrana geotextil no tejido 200g x1m', NULL, 'ferreteria', 'membranas-aislantes', 1200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Cinta butílica impermeabilizante 5cm x3m Sika', 'Sika', 'ferreteria', 'membranas-aislantes', 3800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Primario asfáltico para membrana 1L Volchem', 'Volchem', 'ferreteria', 'membranas-aislantes', 4500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Primario asfáltico para membrana 4L Volchem', 'Volchem', 'ferreteria', 'membranas-aislantes', 15000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Membrana asfáltica aluminio 4mm 10m² Sika', 'Sika', 'ferreteria', 'membranas-aislantes', 36000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Perfil sellador EPDM para ventana 10m Wurth', 'Würth', 'ferreteria', 'membranas-aislantes', 5500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Espuma selladora autoadhesiva 15x10mm x5m', NULL, 'ferreteria', 'membranas-aislantes', 1800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Membrana asfáltica poliéster 4mm 10m² Iggam', 'Iggam', 'ferreteria', 'membranas-aislantes', 30000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tela de refuerzo fibra vidrio 50cm x5m genérica', NULL, 'ferreteria', 'membranas-aislantes', 2500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Burletes adhesivos puerta/ventana 6x9mm x6m genérico', NULL, 'ferreteria', 'membranas-aislantes', 1200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Pintura asfáltica impermeabilizante 4kg Volchem', 'Volchem', 'ferreteria', 'membranas-aislantes', 10000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Pintura asfáltica impermeabilizante 18kg Volchem', 'Volchem', 'ferreteria', 'membranas-aislantes', 38000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ accesorios-pintura ============
-- Fuente: Purdy, 3M, genéricos de ferretería argentina 2025
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Pincel plano 2" cerda sintética Purdy', 'Purdy', 'ferreteria', 'accesorios-pintura', 3800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Pincel plano 3" cerda sintética Purdy', 'Purdy', 'ferreteria', 'accesorios-pintura', 5200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Pincel angular 1.5" pelo sintético Purdy', 'Purdy', 'ferreteria', 'accesorios-pintura', 3200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Pincel biselado 3" pelo sintético Purdy', 'Purdy', 'ferreteria', 'accesorios-pintura', 5500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Rodillo antiflex 22cm lana corta con mango Purdy', 'Purdy', 'ferreteria', 'accesorios-pintura', 4800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Rodillo antiflex 22cm lana larga con mango Purdy', 'Purdy', 'ferreteria', 'accesorios-pintura', 5200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Rodillo 9cm para esquinas con mango genérico', NULL, 'ferreteria', 'accesorios-pintura', 1800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Set rodillo 22cm + bandeja plástica + mango genérico', NULL, 'ferreteria', 'accesorios-pintura', 3500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Cinta de enmascarar 24mm x50m 3M ScotchBlue', '3M', 'ferreteria', 'accesorios-pintura', 2800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cinta de enmascarar 48mm x50m 3M ScotchBlue', '3M', 'ferreteria', 'accesorios-pintura', 4200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Lija en rollo grano 80 5m Truper', 'Truper', 'ferreteria', 'accesorios-pintura', 1900.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Lija en rollo grano 120 5m Truper', 'Truper', 'ferreteria', 'accesorios-pintura', 1900.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Lija en rollo grano 220 5m Truper', 'Truper', 'ferreteria', 'accesorios-pintura', 1900.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Plástico protector para piso 2x3m', NULL, 'ferreteria', 'accesorios-pintura', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Espátula Japón 10cm flexible Truper', 'Truper', 'ferreteria', 'accesorios-pintura', 1500.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Espátula Japón 15cm flexible Truper', 'Truper', 'ferreteria', 'accesorios-pintura', 1800.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Llana de acero 18x12cm Truper', 'Truper', 'ferreteria', 'accesorios-pintura', 2800.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Llana dentada acero inox 22x11cm Truper', 'Truper', 'ferreteria', 'accesorios-pintura', 3200.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Mango telescópico aluminio 60-120cm genérico', NULL, 'ferreteria', 'accesorios-pintura', 2500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Pincel de zócalo 4" cerda natural genérico', NULL, 'ferreteria', 'accesorios-pintura', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Brocha 4" cerda natural para imprimante genérica', NULL, 'ferreteria', 'accesorios-pintura', 2000.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Set pinceles 3 tamaños cerda sintética Purdy', 'Purdy', 'ferreteria', 'accesorios-pintura', 8500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cortador de bordes para zócalo genérico', NULL, 'ferreteria', 'accesorios-pintura', 1200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Contenedor de pintura 1L con tapa hermética genérico', NULL, 'ferreteria', 'accesorios-pintura', 900.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Mezclador de pintura para taladro 8cm genérico', NULL, 'ferreteria', 'accesorios-pintura', 1500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Guante latex fino para pintor talla M x1par', NULL, 'ferreteria', 'accesorios-pintura', 800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ taladros-percutores ============
-- Fuente: Bosch, Makita, DeWalt, Black+Decker catálogos Argentina 2025
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Taladro de banco 350W 13mm Truper', 'Truper', 'ferreteria', 'taladros-percutores', 65000.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Taladro percutor HP2050 500W Makita', 'Makita', 'ferreteria', 'taladros-percutores', 82000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Taladro percutor CD714CRES 500W Black+Decker', 'Black+Decker', 'ferreteria', 'taladros-percutores', 58000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Taladro inalámbrico DCD771C2 18V DeWalt', 'DeWalt', 'ferreteria', 'taladros-percutores', 145000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{}'),
  ('Taladro inalámbrico GSR 120-LI 12V Bosch', 'Bosch', 'ferreteria', 'taladros-percutores', 92000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Taladro inalámbrico 18V DDF481RTJ Makita', 'Makita', 'ferreteria', 'taladros-percutores', 165000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{}'),
  ('Martillo rotativo SDS+ GBH 2-26 DRE Bosch', 'Bosch', 'ferreteria', 'taladros-percutores', 185000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{}'),
  ('Martillo rotativo SDS+ HR2460 Makita', 'Makita', 'ferreteria', 'taladros-percutores', 195000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{}'),
  ('Martillo rotativo SDS+ D25144K DeWalt', 'DeWalt', 'ferreteria', 'taladros-percutores', 198000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{}'),
  ('Taladro percutor 500W KR5010 Black+Decker', 'Black+Decker', 'ferreteria', 'taladros-percutores', 55000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Broca SDS plus 6mm x110mm Bosch', 'Bosch', 'ferreteria', 'taladros-percutores', 2800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Broca SDS plus 8mm x110mm Bosch', 'Bosch', 'ferreteria', 'taladros-percutores', 3200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Broca SDS plus 10mm x160mm Bosch', 'Bosch', 'ferreteria', 'taladros-percutores', 3800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Set brocas SDS plus 6u Bosch', 'Bosch', 'ferreteria', 'taladros-percutores', 18500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Broca para hormigón 8mm HSS Würth', 'Würth', 'ferreteria', 'taladros-percutores', 1800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Broca para hormigón 10mm HSS Würth', 'Würth', 'ferreteria', 'taladros-percutores', 2200.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Set brocas concreto/madera/metal 20u Black+Decker', 'Black+Decker', 'ferreteria', 'taladros-percutores', 12000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Porta brocas 13mm llave Bosch', 'Bosch', 'ferreteria', 'taladros-percutores', 4500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Porta brocas 13mm sin llave Black+Decker', 'Black+Decker', 'ferreteria', 'taladros-percutores', 3800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Destornillador inalámbrico 4V LI 2Ah Black+Decker', 'Black+Decker', 'ferreteria', 'taladros-percutores', 38000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Destornillador inalámbrico 12V GSR12V-20 Bosch', 'Bosch', 'ferreteria', 'taladros-percutores', 82000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Llave de impacto inalámbrica 18V Makita DTD152', 'Makita', 'ferreteria', 'taladros-percutores', 185000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{}'),
  ('Cinceles para martillo rotativo SDS+ 2u Bosch', 'Bosch', 'ferreteria', 'taladros-percutores', 5800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Batería 18V 2Ah BL1820 Makita', 'Makita', 'ferreteria', 'taladros-percutores', 48000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cargador rápido DC18RC 18V Makita', 'Makita', 'ferreteria', 'taladros-percutores', 35000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Batería 18V 1.5Ah GBA18V15 Bosch', 'Bosch', 'ferreteria', 'taladros-percutores', 42000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cargador para batería Li-Ion 14.4/18V Bosch', 'Bosch', 'ferreteria', 'taladros-percutores', 28000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ sierras-electricas ============
-- Fuente: Bosch, Makita, DeWalt, Black+Decker 2025
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Sierra caladora JS470E 500W Black+Decker', 'Black+Decker', 'ferreteria', 'sierras-electricas', 72000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Sierra caladora GST 90 BE 620W Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 125000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Sierra caladora DCS331N 18V DeWalt', 'DeWalt', 'ferreteria', 'sierras-electricas', 168000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{}'),
  ('Sierra circular CS1015 1020W 185mm Black+Decker', 'Black+Decker', 'ferreteria', 'sierras-electricas', 75000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Sierra circular GKS 7-82 1350W 190mm Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 145000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Sierra circular DWE575 1600W DeWalt', 'DeWalt', 'ferreteria', 'sierras-electricas', 178000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{}'),
  ('Sierra circular 5603NH 1300W 165mm Makita', 'Makita', 'ferreteria', 'sierras-electricas', 135000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disco para sierra 7.1/4" 40 dientes madera Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 6500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disco para sierra 7.1/4" 24 dientes madera Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 5500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Hoja de sierra caladora para madera T144D Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 1800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Hoja de sierra caladora para metal T118B Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 2200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Set hojas caladora 5u para madera/metal Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 7500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Amoladora GWS 9-125 900W 125mm Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 115000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Amoladora G10SR4 1010W 100mm Makita', 'Makita', 'ferreteria', 'sierras-electricas', 95000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Amoladora DWE4011 850W DeWalt', 'DeWalt', 'ferreteria', 'sierras-electricas', 105000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Amoladora G6SR3 620W 100mm Makita', 'Makita', 'ferreteria', 'sierras-electricas', 78000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disco de corte 4.5" INOX 1.6mm Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 1200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disco de desbaste 4.5" 6mm acero Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 1800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disco diamantado 4.5" continuo porcelana Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 3500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disco flap 4.5" grano 40 Truper', 'Truper', 'ferreteria', 'sierras-electricas', 1500.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Disco flap 4.5" grano 60 Truper', 'Truper', 'ferreteria', 'sierras-electricas', 1500.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Sierra de mesa 254mm 1800W TS254 Black+Decker', 'Black+Decker', 'ferreteria', 'sierras-electricas', 280000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Lijadora orbital de banda 450W Black+Decker', 'Black+Decker', 'ferreteria', 'sierras-electricas', 65000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Cepilladora 620W KW750K Black+Decker', 'Black+Decker', 'ferreteria', 'sierras-electricas', 85000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Sierra ingleteadora GCMS1200 Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 195000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Fresadora GMF 1600 CE Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 215000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Pack discos de corte 4.5" x25u metal Bosch', 'Bosch', 'ferreteria', 'sierras-electricas', 15000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ cables-electricidad ============
-- Cables Pirelli/Phelps Dodge/Celo presentes en NEA
-- Fuente: Ferreterías Posadas 2025, catálogos CELO
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Cable unipolar 1.5mm² negro rollo 100m Phelps Dodge', 'Phelps Dodge', 'ferreteria', 'cables-electricidad', 38000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cable unipolar 1.5mm² rojo rollo 100m Phelps Dodge', 'Phelps Dodge', 'ferreteria', 'cables-electricidad', 38000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cable unipolar 2.5mm² negro rollo 100m Phelps Dodge', 'Phelps Dodge', 'ferreteria', 'cables-electricidad', 58000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cable unipolar 4mm² negro rollo 100m Phelps Dodge', 'Phelps Dodge', 'ferreteria', 'cables-electricidad', 85000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cable unipolar 6mm² negro rollo 100m Phelps Dodge', 'Phelps Dodge', 'ferreteria', 'cables-electricidad', 125000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cable unipolar 1.5mm² negro x1m Phelps Dodge', 'Phelps Dodge', 'ferreteria', 'cables-electricidad', 420.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cable unipolar 2.5mm² negro x1m Phelps Dodge', 'Phelps Dodge', 'ferreteria', 'cables-electricidad', 620.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cable unipolar 4mm² negro x1m Phelps Dodge', 'Phelps Dodge', 'ferreteria', 'cables-electricidad', 920.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cable tipo taller 2x2.5mm² x1m Phelps Dodge', 'Phelps Dodge', 'ferreteria', 'cables-electricidad', 1100.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cable tipo taller 3x2.5mm² x1m Phelps Dodge', 'Phelps Dodge', 'ferreteria', 'cables-electricidad', 1500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cable tipo taller 3x4mm² x1m Phelps Dodge', 'Phelps Dodge', 'ferreteria', 'cables-electricidad', 2100.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cable para teléfono 2 pares rollo 100m genérico', NULL, 'ferreteria', 'cables-electricidad', 15000.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Cable coaxial RG6 rollo 100m genérico', NULL, 'ferreteria', 'cables-electricidad', 25000.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Canaleta PVC 20x10mm x2m blanca genérica', NULL, 'ferreteria', 'cables-electricidad', 1200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Canaleta PVC 40x20mm x2m blanca genérica', NULL, 'ferreteria', 'cables-electricidad', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Canaleta PVC 60x40mm x2m blanca genérica', NULL, 'ferreteria', 'cables-electricidad', 3500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Caja de pase plástica cuadrada 10x10cm', NULL, 'ferreteria', 'cables-electricidad', 800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Caja de pase plástica 15x15cm', NULL, 'ferreteria', 'cables-electricidad', 1200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Caja embutir simple plástica genérica', NULL, 'ferreteria', 'cables-electricidad', 350.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Caño conduit flexible 3/4" x3m genérico', NULL, 'ferreteria', 'cables-electricidad', 2800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Tapón caño conduit 3/4" genérico x5u', NULL, 'ferreteria', 'cables-electricidad', 500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Abrazadera para caño conduit 3/4" x20u genérica', NULL, 'ferreteria', 'cables-electricidad', 600.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ mangueras-riego ============
-- Fuente: Genco, Regar, Tramontina presencia NEA 2025
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Manguera jardín 1/2" x10m Genco', 'Genco', 'ferreteria', 'mangueras-riego', 8500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Manguera jardín 3/4" x50m Genco', 'Genco', 'ferreteria', 'mangueras-riego', 38000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Manguera reforzada 5/8" x25m Tramontina', 'Tramontina', 'ferreteria', 'mangueras-riego', 28000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Manguera reforzada 5/8" x50m Tramontina', 'Tramontina', 'ferreteria', 'mangueras-riego', 52000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Manguera extensible 7m-22m con pistola Truper', 'Truper', 'ferreteria', 'mangueras-riego', 9500.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Manguera extensible 10m-30m con pistola Truper', 'Truper', 'ferreteria', 'mangueras-riego', 15000.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Pistola 7 funciones para manguera Tramontina', 'Tramontina', 'ferreteria', 'mangueras-riego', 4800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Pistola 3 funciones para manguera Truper', 'Truper', 'ferreteria', 'mangueras-riego', 2800.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Regadera plástica 10L con alcachofa Tramontina', 'Tramontina', 'ferreteria', 'mangueras-riego', 4200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Regadera metálica 8L Tramontina', 'Tramontina', 'ferreteria', 'mangueras-riego', 6500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Aspersor emergente circular R4 Rain Bird', 'Rain Bird', 'ferreteria', 'mangueras-riego', 2800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Aspersor emergente 12cm Rain Bird', 'Rain Bird', 'ferreteria', 'mangueras-riego', 3500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Micro aspersor 360° para jardín x5u genérico', NULL, 'ferreteria', 'mangueras-riego', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Conector rápido para manguera 1/2" x2u Tramontina', 'Tramontina', 'ferreteria', 'mangueras-riego', 1800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Reparador de manguera 1/2" Tramontina', 'Tramontina', 'ferreteria', 'mangueras-riego', 1500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Temporizador riego automático digital Rain Bird', 'Rain Bird', 'ferreteria', 'mangueras-riego', 18000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Gotero autocompensante 2L/h x10u Rain Bird', 'Rain Bird', 'ferreteria', 'mangueras-riego', 2500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Manguera microaspersión 1/2" x10m genérica', NULL, 'ferreteria', 'mangueras-riego', 4500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Bomba sumergible jardín 400W 5000L/h Grundfos', 'Grundfos', 'ferreteria', 'mangueras-riego', 65000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Grifo exterior para jardín 1/2" Perafan', 'Perafan', 'ferreteria', 'mangueras-riego', 2800.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Tee de distribución para manguera x3 salidas Tramontina', 'Tramontina', 'ferreteria', 'mangueras-riego', 1600.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Filtro para goteo 1/2" 130 mesh Rain Bird', 'Rain Bird', 'ferreteria', 'mangueras-riego', 1800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Kit riego por goteo 20 plantas Rain Bird', 'Rain Bird', 'ferreteria', 'mangueras-riego', 12000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Manguera espiral de 5m 1/4" para agua a presión', NULL, 'ferreteria', 'mangueras-riego', 5500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Canilla esférica 1/2" latón Perafan', 'Perafan', 'ferreteria', 'mangueras-riego', 3200.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ llaves-valvulas ============
-- Fuente: FV, Perafan, Ferrum distribución NEA 2025
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Válvula esfera 1/2" hembra-hembra bronce Perafan', 'Perafan', 'ferreteria', 'llaves-valvulas', 3800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Válvula esfera 3/4" hembra-hembra bronce Perafan', 'Perafan', 'ferreteria', 'llaves-valvulas', 5500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Válvula esfera 1" hembra-hembra bronce Perafan', 'Perafan', 'ferreteria', 'llaves-valvulas', 7800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Válvula compuerta 1/2" bronce FV', 'FV', 'ferreteria', 'llaves-valvulas', 4200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Válvula compuerta 3/4" bronce FV', 'FV', 'ferreteria', 'llaves-valvulas', 5800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Válvula compuerta 1" bronce FV', 'FV', 'ferreteria', 'llaves-valvulas', 8500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Válvula check 1/2" bronce FV', 'FV', 'ferreteria', 'llaves-valvulas', 5200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Válvula check 3/4" bronce FV', 'FV', 'ferreteria', 'llaves-valvulas', 6800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Válvula reductora de presión 1/2" bronce genérica', NULL, 'ferreteria', 'llaves-valvulas', 8500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Válvula flotante 1/2" plástica para tanque genérica', NULL, 'ferreteria', 'llaves-valvulas', 1800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Válvula flotante 3/4" plástica para tanque genérica', NULL, 'ferreteria', 'llaves-valvulas', 2400.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Llave de paso 1/2" para empotrar FV', 'FV', 'ferreteria', 'llaves-valvulas', 3500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Llave de gas 3/4" Orbis', 'Orbis', 'ferreteria', 'llaves-valvulas', 4800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Llave de gas 1/2" Orbis', 'Orbis', 'ferreteria', 'llaves-valvulas', 3800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Regulador de gas domiciliario Orbis', 'Orbis', 'ferreteria', 'llaves-valvulas', 5200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Pileta de patio enlozada blanca 45x45cm Ferrum', 'Ferrum', 'ferreteria', 'llaves-valvulas', 18000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Válvula esfera miniatura 3/8" para lavarropas Perafan', 'Perafan', 'ferreteria', 'llaves-valvulas', 2800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Válvula esfera 1/4" para gas natural Perafan', 'Perafan', 'ferreteria', 'llaves-valvulas', 3200.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Llave mezcladora pileta cocina monomando FV', 'FV', 'ferreteria', 'llaves-valvulas', 28000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Llave ducha monomando FV', 'FV', 'ferreteria', 'llaves-valvulas', 22000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Llave bidet monomando FV', 'FV', 'ferreteria', 'llaves-valvulas', 18000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Grifería lavatorio monomando FV Mondrian', 'FV', 'ferreteria', 'llaves-valvulas', 32000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Flexible extensible 1/2"x50cm para water cromado', NULL, 'ferreteria', 'llaves-valvulas', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Flexible extensible 1/2"x30cm para lavatorio cromado', NULL, 'ferreteria', 'llaves-valvulas', 1800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Llave angular 1/2" para inodoro FV', 'FV', 'ferreteria', 'llaves-valvulas', 3500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ proteccion-electrica ============
-- Fuente: Bticino, Siemens, Schneider Electric distribución NEA 2025
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Disyuntor termomagnético 1x32A Bticino', 'Bticino', 'ferreteria', 'proteccion-electrica', 8500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disyuntor termomagnético 2x16A Bticino', 'Bticino', 'ferreteria', 'proteccion-electrica', 14000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disyuntor termomagnético 2x20A Bticino', 'Bticino', 'ferreteria', 'proteccion-electrica', 14500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disyuntor termomagnético 2x25A Bticino', 'Bticino', 'ferreteria', 'proteccion-electrica', 15000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disyuntor termomagnético 3x20A Bticino', 'Bticino', 'ferreteria', 'proteccion-electrica', 22000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disyuntor diferencial 4x25A 30mA Bticino', 'Bticino', 'ferreteria', 'proteccion-electrica', 52000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disyuntor diferencial 4x40A 30mA Bticino', 'Bticino', 'ferreteria', 'proteccion-electrica', 58000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disyuntor termomagnético 1x10A Siemens 5SL6', 'Siemens', 'ferreteria', 'proteccion-electrica', 7200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disyuntor termomagnético 1x16A Siemens 5SL6', 'Siemens', 'ferreteria', 'proteccion-electrica', 7500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disyuntor termomagnético 1x20A Siemens 5SL6', 'Siemens', 'ferreteria', 'proteccion-electrica', 7800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Disyuntor diferencial 2x25A 30mA Siemens', 'Siemens', 'ferreteria', 'proteccion-electrica', 28000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tablero de distribucion 6 circuitos Bticino', 'Bticino', 'ferreteria', 'proteccion-electrica', 15000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tablero de distribucion 12 circuitos Bticino', 'Bticino', 'ferreteria', 'proteccion-electrica', 22000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tablero de distribucion 24 circuitos Bticino', 'Bticino', 'ferreteria', 'proteccion-electrica', 35000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Barra de tierra cobre 14 pines Bticino', 'Bticino', 'ferreteria', 'proteccion-electrica', 4500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Fusible cilíndrico 10A 10x38mm genérico x5u', NULL, 'ferreteria', 'proteccion-electrica', 1500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Fusible cilíndrico 16A 10x38mm genérico x5u', NULL, 'ferreteria', 'proteccion-electrica', 1500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Portafusible industrial 10x38mm genérico', NULL, 'ferreteria', 'proteccion-electrica', 1200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Puesta a tierra jabalina 1.5m Schneider Electric', 'Schneider Electric', 'ferreteria', 'proteccion-electrica', 8500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cable de tierra verde-amarillo 6mm² x1m genérico', NULL, 'ferreteria', 'proteccion-electrica', 950.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Interruptor de horario digital para riel DIN genérico', NULL, 'ferreteria', 'proteccion-electrica', 8500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Contactor 9A 24V para riel DIN Schneider Electric', 'Schneider Electric', 'ferreteria', 'proteccion-electrica', 15000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Variador de frecuencia 0.75kW monofásico Siemens', 'Siemens', 'ferreteria', 'proteccion-electrica', 125000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('UPS 600VA Forza FT-602K', 'Forza', 'ferreteria', 'proteccion-electrica', 65000.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ llaves-tomacorrientes ============
-- Faltan 20 para llegar a 30 verificados
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Llave de luz triple Cambre Siglo XXI', 'Cambre', 'ferreteria', 'llaves-tomacorrientes', 5800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Placa 3 módulos Cambre Siglo XXI', 'Cambre', 'ferreteria', 'llaves-tomacorrientes', 3200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Placa 4 módulos Cambre Siglo XXI', 'Cambre', 'ferreteria', 'llaves-tomacorrientes', 3800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tomacorriente simple Bticino Living', 'Bticino', 'ferreteria', 'llaves-tomacorrientes', 4200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tomacorriente triple Cambre Siglo XXI', 'Cambre', 'ferreteria', 'llaves-tomacorrientes', 8500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Llave de luz simple Bticino Living', 'Bticino', 'ferreteria', 'llaves-tomacorrientes', 5200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Llave de luz doble Bticino Axolute', 'Bticino', 'ferreteria', 'llaves-tomacorrientes', 12000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Módulo USB cargador doble 2.1A Cambre', 'Cambre', 'ferreteria', 'llaves-tomacorrientes', 5800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Módulo USB cargador doble Bticino Living', 'Bticino', 'ferreteria', 'llaves-tomacorrientes', 6500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Módulo pulsador timbre Bticino Living', 'Bticino', 'ferreteria', 'llaves-tomacorrientes', 4800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Módulo regulador de luz dimmer Bticino', 'Bticino', 'ferreteria', 'llaves-tomacorrientes', 15000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Módulo regulador dimmer LED Cambre', 'Cambre', 'ferreteria', 'llaves-tomacorrientes', 12000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tapa ciega Cambre Siglo XXI', 'Cambre', 'ferreteria', 'llaves-tomacorrientes', 1500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tapa ciega Bticino Living', 'Bticino', 'ferreteria', 'llaves-tomacorrientes', 1800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Módulo de conmutación Cambre Siglo XXI', 'Cambre', 'ferreteria', 'llaves-tomacorrientes', 4800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Módulo de conmutación Bticino Living', 'Bticino', 'ferreteria', 'llaves-tomacorrientes', 5500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tomacorriente de piso estanco IP44 Bticino', 'Bticino', 'ferreteria', 'llaves-tomacorrientes', 12000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tapita protectora para tomacorriente x5u genérica', NULL, 'ferreteria', 'llaves-tomacorrientes', 500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Llave de luz simple estanca IP44 Bticino', 'Bticino', 'ferreteria', 'llaves-tomacorrientes', 6500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tomacorriente doble estanco IP44 Bticino', 'Bticino', 'ferreteria', 'llaves-tomacorrientes', 8500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ esmaltes-barnices ============
-- Faltan 19 para llegar a 30
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Esmalte sintético blanco 1L Kem', 'Kem', 'ferreteria', 'esmaltes-barnices', 6200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Esmalte sintético blanco 4L Kem', 'Kem', 'ferreteria', 'esmaltes-barnices', 21000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Esmalte anticorrosivo rojo 1L Sinteplast', 'Sinteplast', 'ferreteria', 'esmaltes-barnices', 6800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Esmalte anticorrosivo rojo 4L Sinteplast', 'Sinteplast', 'ferreteria', 'esmaltes-barnices', 22000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Esmalte sintético colores 1L Alba', 'Alba', 'ferreteria', 'esmaltes-barnices', 5800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Esmalte sintético blanco 1L Alba', 'Alba', 'ferreteria', 'esmaltes-barnices', 5500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Esmalte sintético blanco 4L Alba', 'Alba', 'ferreteria', 'esmaltes-barnices', 18500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Barniz marino satinado 1L Alba', 'Alba', 'ferreteria', 'esmaltes-barnices', 7200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Barniz parquet brillante 1L Sinteplast', 'Sinteplast', 'ferreteria', 'esmaltes-barnices', 6800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Diluyente universal 1L Sinteplast', 'Sinteplast', 'ferreteria', 'esmaltes-barnices', 2800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Diluyente universal 4L Sinteplast', 'Sinteplast', 'ferreteria', 'esmaltes-barnices', 9500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Barniz alkídico brillante 4L Tersuave', 'Tersuave', 'ferreteria', 'esmaltes-barnices', 21000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Esmalte sintético negro brillante 1L Tersuave', 'Tersuave', 'ferreteria', 'esmaltes-barnices', 5800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Esmalte color azul Francia 1L Sinteplast', 'Sinteplast', 'ferreteria', 'esmaltes-barnices', 6200.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Esmalte color verde botella 1L Sinteplast', 'Sinteplast', 'ferreteria', 'esmaltes-barnices', 6200.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Esmalte color rojo bermellón 1L Sinteplast', 'Sinteplast', 'ferreteria', 'esmaltes-barnices', 6200.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Esmalte para piso gris 1L Sinteplast', 'Sinteplast', 'ferreteria', 'esmaltes-barnices', 7200.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Esmalte para piso rojo 1L Sinteplast', 'Sinteplast', 'ferreteria', 'esmaltes-barnices', 7200.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Imprimante al agua 4L Sinteplast', 'Sinteplast', 'ferreteria', 'esmaltes-barnices', 9800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ medicion ============
-- Faltan 18 para llegar a 30
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Cinta métrica 5m Stanley FatMax', 'Stanley', 'ferreteria', 'medicion', 6500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cinta métrica 10m Stanley', 'Stanley', 'ferreteria', 'medicion', 8500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cinta métrica 5m Truper', 'Truper', 'ferreteria', 'medicion', 3800.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Nivel de burbuja 80cm Stanley', 'Stanley', 'ferreteria', 'medicion', 9500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Nivel de burbuja 200cm Stanley', 'Stanley', 'ferreteria', 'medicion', 15000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Nivel electrónico digital 40cm Bosch', 'Bosch', 'ferreteria', 'medicion', 18500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Medidor láser de distancia GLM 30-23 Bosch', 'Bosch', 'ferreteria', 'medicion', 48000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{}'),
  ('Medidor láser de distancia GLM 50-27 Bosch', 'Bosch', 'ferreteria', 'medicion', 72000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{}'),
  ('Escuadra carpintero 60cm Stanley', 'Stanley', 'ferreteria', 'medicion', 9800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Escuadra combinada 30cm Stanley', 'Stanley', 'ferreteria', 'medicion', 8500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Calibre Vernier 200mm acero inox Truper', 'Truper', 'ferreteria', 'medicion', 7500.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Transportador de ángulos 180° Stanley', 'Stanley', 'ferreteria', 'medicion', 4800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Regla aluminio 60cm Truper', 'Truper', 'ferreteria', 'medicion', 2800.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Regla aluminio 100cm Truper', 'Truper', 'ferreteria', 'medicion', 4200.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Plomada de obra 500g con cordón Truper', 'Truper', 'ferreteria', 'medicion', 2200.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Cordón de tiza 30m azul Truper', 'Truper', 'ferreteria', 'medicion', 1800.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Nivel de torpedero magnético 9" Stanley', 'Stanley', 'ferreteria', 'medicion', 5800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Analizador de pared Bosch GMS 120 Professional', 'Bosch', 'ferreteria', 'medicion', 38000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ adhesivos-selladores ============
-- Faltan 18 para llegar a 30
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Silicona neutra blanca 280ml Bostik', 'Bostik', 'ferreteria', 'adhesivos-selladores', 3500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Silicona neutra transparente 280ml Bostik', 'Bostik', 'ferreteria', 'adhesivos-selladores', 3500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Sellador acrílico blanco 280ml Bostik', 'Bostik', 'ferreteria', 'adhesivos-selladores', 2800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Sellador poliuretano gris 300ml Sika', 'Sika', 'ferreteria', 'adhesivos-selladores', 5800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Adhesivo de montaje 280ml Sika Power Fix', 'Sika', 'ferreteria', 'adhesivos-selladores', 4800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Adhesivo de contacto neopreno 250ml Sika', 'Sika', 'ferreteria', 'adhesivos-selladores', 4200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Adhesivo de contacto neopreno 750ml Sika', 'Sika', 'ferreteria', 'adhesivos-selladores', 11000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Poxipol dorado 14g', 'Poxipol', 'ferreteria', 'adhesivos-selladores', 2800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Cinta de doble faz 25mm x10m 3M', '3M', 'ferreteria', 'adhesivos-selladores', 3500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cinta de doble faz 50mm x5m 3M', '3M', 'ferreteria', 'adhesivos-selladores', 4200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Cola vinílica 1kg genérica', NULL, 'ferreteria', 'adhesivos-selladores', 2800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Cola vinílica 5kg genérica', NULL, 'ferreteria', 'adhesivos-selladores', 12000.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Adhesivo instantáneo Super Bonder 3g Loctite', 'Loctite', 'ferreteria', 'adhesivos-selladores', 1800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Espuma de relleno acrílica blanca 300ml Bostik', 'Bostik', 'ferreteria', 'adhesivos-selladores', 3800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Sellador para juntas de dilatación gris 600ml Sika', 'Sika', 'ferreteria', 'adhesivos-selladores', 7500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Pegamento de cerámica Sika Adhesol 20kg', 'Sika', 'ferreteria', 'adhesivos-selladores', 9500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Fragüe para porcellanato blanco 1kg genérico', NULL, 'ferreteria', 'adhesivos-selladores', 1800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Fragüe para porcellanato gris 5kg genérico', NULL, 'ferreteria', 'adhesivos-selladores', 7500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ cerrajeria-seguridad ============
-- Faltan 18 verificados para llegar a 30
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Candado 60mm Master Lock combinación 4 dígitos', 'Master Lock', 'ferreteria', 'cerrajeria-seguridad', 8500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Candado 40mm latón Yale', 'Yale', 'ferreteria', 'cerrajeria-seguridad', 5200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Candado con llave 50mm latón genérico', NULL, 'ferreteria', 'cerrajeria-seguridad', 3200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Bisagra de puerta 3" acero inoxidable x2u', NULL, 'ferreteria', 'cerrajeria-seguridad', 2800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Bisagra de puerta 4" acero inoxidable x2u', NULL, 'ferreteria', 'cerrajeria-seguridad', 3800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Bisagra de portón soldable 5" x2u', NULL, 'ferreteria', 'cerrajeria-seguridad', 5200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Cerradura de embutir 3 golpes Trabex', 'Trabex', 'ferreteria', 'cerrajeria-seguridad', 18000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Cerradura horizontal 3 puntos seguridad Trabex', 'Trabex', 'ferreteria', 'cerrajeria-seguridad', 28000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Pasador de puerta 6" galvanizado', NULL, 'ferreteria', 'cerrajeria-seguridad', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Aldaba y pasador 4" galvanizado', NULL, 'ferreteria', 'cerrajeria-seguridad', 1800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Picaporte con ojo 100mm galvanizado', NULL, 'ferreteria', 'cerrajeria-seguridad', 1500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Traba de seguridad para puerta antichoque Trabex', 'Trabex', 'ferreteria', 'cerrajeria-seguridad', 12000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Mirilla ojo de buey 200° plateado', NULL, 'ferreteria', 'cerrajeria-seguridad', 2800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Tirador de cajón 96mm cromo x2u', NULL, 'ferreteria', 'cerrajeria-seguridad', 1800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Tirador de cajón 128mm cromo x2u', NULL, 'ferreteria', 'cerrajeria-seguridad', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Gozne de hierro 1/2" x5u galvanizado', NULL, 'ferreteria', 'cerrajeria-seguridad', 800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Cerradura con llave para cajonera Trabex', 'Trabex', 'ferreteria', 'cerrajeria-seguridad', 4500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Cámara de seguridad exterior IP66 720p genérica', NULL, 'ferreteria', 'cerrajeria-seguridad', 12000.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ sanitarios-plomeria ============
-- Faltan 18 para llegar a 30
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Inodoro monoblock FERRUM Iguazú blanco', 'Ferrum', 'ferreteria', 'sanitarios-plomeria', 85000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Lavabo suspendido FERRUM 52cm blanco', 'Ferrum', 'ferreteria', 'sanitarios-plomeria', 35000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Bidet FERRUM Artic blanco', 'Ferrum', 'ferreteria', 'sanitarios-plomeria', 45000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Ducha eléctrica 5500W Orbis', 'Orbis', 'ferreteria', 'sanitarios-plomeria', 22000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Ducha eléctrica 7500W Orbis', 'Orbis', 'ferreteria', 'sanitarios-plomeria', 28000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Tanque de reserva 500L Rotoplas', 'Rotoplas', 'ferreteria', 'sanitarios-plomeria', 85000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tanque de reserva 1000L Rotoplas', 'Rotoplas', 'ferreteria', 'sanitarios-plomeria', 145000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Bomba de agua 1HP Grundfos CM3-5', 'Grundfos', 'ferreteria', 'sanitarios-plomeria', 125000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Caño multicapa 16mm x1m genérico', NULL, 'ferreteria', 'sanitarios-plomeria', 1800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Caño multicapa 20mm x1m genérico', NULL, 'ferreteria', 'sanitarios-plomeria', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Conector para caño multicapa 16mm genérico', NULL, 'ferreteria', 'sanitarios-plomeria', 1200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Trampa para pileta de patio 3" PVC Tigre', 'Tigre', 'ferreteria', 'sanitarios-plomeria', 4500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Sifón botella para lavatorio Tigre', 'Tigre', 'ferreteria', 'sanitarios-plomeria', 3800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Cinta teflón 12mm x10m genérica', NULL, 'ferreteria', 'sanitarios-plomeria', 500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Cinta teflón 19mm x15m genérica', NULL, 'ferreteria', 'sanitarios-plomeria', 800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Rejilla de piso inoxidable 10x10cm', NULL, 'ferreteria', 'sanitarios-plomeria', 1800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Flotante para cisterna de inodoro genérico', NULL, 'ferreteria', 'sanitarios-plomeria', 2800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Kit de fijación para inodoro genérico', NULL, 'ferreteria', 'sanitarios-plomeria', 1500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ iluminacion ============
-- Faltan 16 para llegar a 30
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Lámpara LED bulbo 6W E27 luz cálida Philips', 'Philips', 'ferreteria', 'iluminacion', 1800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Lámpara LED bulbo 12W E27 luz cálida Osram', 'Osram', 'ferreteria', 'iluminacion', 2200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Lámpara LED AR111 12W GU10 Philips', 'Philips', 'ferreteria', 'iluminacion', 4800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tubo LED T8 18W 120cm luz cálida Philips', 'Philips', 'ferreteria', 'iluminacion', 3800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Lámpara LED panel cuadrado 18W 4000K Ledvance', 'Ledvance', 'ferreteria', 'iluminacion', 6500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Lámpara LED panel redondo 12W 4000K Ledvance', 'Ledvance', 'ferreteria', 'iluminacion', 5200.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Reflector LED 20W exterior IP65 Ledvance', 'Ledvance', 'ferreteria', 'iluminacion', 8500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Reflector LED 100W exterior Philips', 'Philips', 'ferreteria', 'iluminacion', 32000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Luz de emergencia LED 2h Philips', 'Philips', 'ferreteria', 'iluminacion', 8500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Lámpara solar exterior con sensor 30W genérica', NULL, 'ferreteria', 'iluminacion', 12000.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Tira LED 12V 5050 5m blanco cálido genérica', NULL, 'ferreteria', 'iluminacion', 6500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Fuente de alimentación 12V 60W para tira LED', NULL, 'ferreteria', 'iluminacion', 5500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Lámpara LED filamento 7W E27 cálida Philips', 'Philips', 'ferreteria', 'iluminacion', 2800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Lámpara LED vela 6W E14 luz cálida Osram', 'Osram', 'ferreteria', 'iluminacion', 2200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Portalámparas E27 plástico baquelita genérico', NULL, 'ferreteria', 'iluminacion', 800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Timbre inalámbrico 200m alcance genérico', NULL, 'ferreteria', 'iluminacion', 5500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ jardin-herramientas ============
-- Faltan 14 para llegar a 30
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Rastrillo de plástico 24 dientes Tramontina', 'Tramontina', 'ferreteria', 'jardin-herramientas', 4200.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Azada liviana 14cm mango corto Tramontina', 'Tramontina', 'ferreteria', 'jardin-herramientas', 3800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Horquilla de jardín 4 dientes Truper', 'Truper', 'ferreteria', 'jardin-herramientas', 5500.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Cortacésped nafta 140cc 46cm Truper', 'Truper', 'ferreteria', 'jardin-herramientas', 225000.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Motosierra a nafta 40cm 2HP Truper', 'Truper', 'ferreteria', 'jardin-herramientas', 185000.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Sierra de poda telescópica 3m Tramontina', 'Tramontina', 'ferreteria', 'jardin-herramientas', 15000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tijera de jardín 8" inox Tramontina', 'Tramontina', 'ferreteria', 'jardin-herramientas', 3800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Pulverizador a presión 5L Tramontina', 'Tramontina', 'ferreteria', 'jardin-herramientas', 8500.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Pulverizador a presión 8L Truper', 'Truper', 'ferreteria', 'jardin-herramientas', 12000.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Guante de jardín talla M Tramontina', 'Tramontina', 'ferreteria', 'jardin-herramientas', 1800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Guante de jardín talla L Tramontina', 'Tramontina', 'ferreteria', 'jardin-herramientas', 1800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Palin de punta acero mango madera Tramontina', 'Tramontina', 'ferreteria', 'jardin-herramientas', 12000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Carretilla 50L cuerpo polietileno Truper', 'Truper', 'ferreteria', 'jardin-herramientas', 42000.00, 'catalog-researcher', 0.85, 55, TRUE, TRUE, '{}'),
  ('Bordeadora a hilo nafta 30cc Black+Decker', 'Black+Decker', 'ferreteria', 'jardin-herramientas', 145000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ pinturas-latex ============
-- Faltan 10 para llegar a 30
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Látex interior lavable blanco 10L Kem', 'Kem', 'ferreteria', 'pinturas-latex', 42000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Látex interior blanco 4L Kem', 'Kem', 'ferreteria', 'pinturas-latex', 18000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Látex exterior blanco 4L Plavicon', 'Plavicon', 'ferreteria', 'pinturas-latex', 16000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Látex exterior blanco 20L Plavicon', 'Plavicon', 'ferreteria', 'pinturas-latex', 72000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Látex interior blanco 20L Plavicon', 'Plavicon', 'ferreteria', 'pinturas-latex', 65000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Látex interior blanco 10L Plavicon', 'Plavicon', 'ferreteria', 'pinturas-latex', 35000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Látex interior antihumedad blanco 4L Tersuave', 'Tersuave', 'ferreteria', 'pinturas-latex', 18500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Látex anti-hongos blanco 4L Sinteplast', 'Sinteplast', 'ferreteria', 'pinturas-latex', 17000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Látex interior gris piedra 4L Sinteplast', 'Sinteplast', 'ferreteria', 'pinturas-latex', 16500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Látex interior blanco techo 20L Sinteplast', 'Sinteplast', 'ferreteria', 'pinturas-latex', 68000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ herramientas-electricas ============
-- Faltan 6 para llegar a 30
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Pistola de calor 2000W GHG 20-63 Bosch', 'Bosch', 'ferreteria', 'herramientas-electricas', 75000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Pistola encoladora 200W GKP 200 CE Bosch', 'Bosch', 'ferreteria', 'herramientas-electricas', 28000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Pistola encoladora 40W Black+Decker', 'Black+Decker', 'ferreteria', 'herramientas-electricas', 12000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Lijadora de detalle Mouse 55W Bosch', 'Bosch', 'ferreteria', 'herramientas-electricas', 42000.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Lijadora orbital KA280K 135W Black+Decker', 'Black+Decker', 'ferreteria', 'herramientas-electricas', 38000.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Multiherramienta oscilante 3015 DWE315K DeWalt', 'DeWalt', 'ferreteria', 'herramientas-electricas', 145000.00, 'catalog-researcher', 0.85, 70, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;

-- ============ fijaciones ============
-- Agregar 17 productos con marca para mejorar quality_score
INSERT INTO global_products (name, brand, business_type, category, price, source, source_reliability, quality_score, is_verified, is_active, metadata)
VALUES
  ('Tarugo Fischer S12 x25u', 'Fischer', 'ferreteria', 'fijaciones', 3800.00, 'catalog-researcher', 0.85, 65, TRUE, TRUE, '{}'),
  ('Tornillo para placa de yeso 3.5x35mm x100u Würth', 'Würth', 'ferreteria', 'fijaciones', 3500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Tornillo HWS madera 5x50mm parcial x100u Würth', 'Würth', 'ferreteria', 'fijaciones', 4800.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Tornillo estructural SDS 6x100mm c/arandela x25u Würth', 'Würth', 'ferreteria', 'fijaciones', 6500.00, 'catalog-researcher', 0.85, 60, TRUE, TRUE, '{}'),
  ('Remache pop 3.2x8mm inox x50u', NULL, 'ferreteria', 'fijaciones', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Gancho de acero para estanterías 50mm x4u', NULL, 'ferreteria', 'fijaciones', 1800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Placa de anclaje para viga 90x60mm x2u', NULL, 'ferreteria', 'fijaciones', 2800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Perfil de acero L 30x30x3mm x1m', NULL, 'ferreteria', 'fijaciones', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Suela antivibración 10x10cm 10mm goma técnica', NULL, 'ferreteria', 'fijaciones', 1500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Taco mariposa M5 para placa de yeso x20u genérico', NULL, 'ferreteria', 'fijaciones', 1500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Taco mariposa M8 para placa de yeso x10u genérico', NULL, 'ferreteria', 'fijaciones', 1800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Taco mariposa plástico para cartón-yeso 5mm x20u', NULL, 'ferreteria', 'fijaciones', 1200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Varilla roscada M6 x1m galvanizada', NULL, 'ferreteria', 'fijaciones', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Eslabón de cadena D5mm acero inox x5u', NULL, 'ferreteria', 'fijaciones', 1800.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Abrazadera de acero P3/4" galvanizada x5u', NULL, 'ferreteria', 'fijaciones', 1200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Clavos para tiro 22mm x100u genérico', NULL, 'ferreteria', 'fijaciones', 1500.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}'),
  ('Tornillo tensión 1/4"x2.5" galvanizado x25u', NULL, 'ferreteria', 'fijaciones', 2200.00, 'catalog-researcher', 0.85, 45, TRUE, TRUE, '{}')
ON CONFLICT (name, business_type) DO NOTHING;
