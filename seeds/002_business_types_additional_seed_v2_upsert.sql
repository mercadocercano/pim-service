-- Seed v2 (idempotente): Tipos de comercio adicionales (NO borra)
-- Reemplaza el comportamiento destructivo/solo-insert de seeds/002_business_types_additional_seed.sql
-- Upsert por (code).

INSERT INTO business_types (code, name, description, icon, color, sort_order, is_active)
VALUES
  -- Servicios digitales y turismo
  ('agencia_viajes', 'Agencia de Viajes', 'Pasajes aéreos, paquetes turísticos, hoteles y servicios de turismo', 'plane', '#0EA5E9', 29, true),
  ('delivery', 'Delivery de Comidas', 'Servicio de entrega de comidas a domicilio, apps de delivery y restaurantes virtuales', 'bike', '#F97316', 30, true),
  ('servicios_digitales', 'Servicios Digitales', 'Software, aplicaciones, contenidos audiovisuales y servicios online', 'cloud', '#8B5CF6', 31, true),

  -- Educación y capacitación
  ('centro_educativo', 'Centro Educativo', 'Institutos, academias, cursos, seminarios y capacitación profesional', 'graduation-cap', '#059669', 32, true),

  -- Entretenimiento
  ('entretenimiento', 'Entretenimiento', 'Entradas a espectáculos, eventos, cines y actividades recreativas', 'ticket', '#DC2626', 33, true),

  -- Servicios automotrices especializados
  ('lavadero', 'Lavadero de Autos', 'Lavado de vehículos, encerado, limpieza integral y servicios de detailing', 'car-wash', '#6B7280', 34, true),

  -- Servicios profesionales
  ('servicios_profesionales', 'Servicios Profesionales', 'Consultoría, servicios contables, legales, técnicos y profesionales independientes', 'briefcase-business', '#1E40AF', 35, true)
ON CONFLICT (code) DO UPDATE SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  icon = EXCLUDED.icon,
  color = EXCLUDED.color,
  sort_order = EXCLUDED.sort_order,
  is_active = EXCLUDED.is_active;


