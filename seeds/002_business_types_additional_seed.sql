-- Seed: 002_business_types_additional_seed.sql
-- Purpose: Agregar tipos de negocio adicionales basados en investigación de mercado argentino
-- Date: 2025-01-27

-- Insertamos tipos de comercios adicionales identificados en la investigación
INSERT INTO business_types (code, name, description, icon, color, sort_order) VALUES

-- Servicios digitales y turismo
('agencia_viajes', 'Agencia de Viajes', 'Pasajes aéreos, paquetes turísticos, hoteles y servicios de turismo', 'plane', '#0EA5E9', 29),
('delivery', 'Delivery de Comidas', 'Servicio de entrega de comidas a domicilio, apps de delivery y restaurantes virtuales', 'bike', '#F97316', 30),
('servicios_digitales', 'Servicios Digitales', 'Software, aplicaciones, contenidos audiovisuales y servicios online', 'cloud', '#8B5CF6', 31),

-- Educación y capacitación
('centro_educativo', 'Centro Educativo', 'Institutos, academias, cursos, seminarios y capacitación profesional', 'graduation-cap', '#059669', 32),

-- Entretenimiento
('entretenimiento', 'Entretenimiento', 'Entradas a espectáculos, eventos, cines y actividades recreativas', 'ticket', '#DC2626', 33),

-- Servicios automotrices especializados
('lavadero', 'Lavadero de Autos', 'Lavado de vehículos, encerado, limpieza integral y servicios de detailing', 'car-wash', '#6B7280', 34),

-- Servicios profesionales
('servicios_profesionales', 'Servicios Profesionales', 'Consultoría, servicios contables, legales, técnicos y profesionales independientes', 'briefcase-business', '#1E40AF', 35);

-- Comentario
COMMENT ON TABLE business_types IS 'Tipos de comercios físicos y digitales argentinos para marketplace SaaS multitenant'; 