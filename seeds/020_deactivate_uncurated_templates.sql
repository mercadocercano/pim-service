-- Seeder 020: Desactivar templates de rubros sin datos curados
-- PROPÓSITO: Solo ofrecer en quickstart los rubros que tienen datos de calidad
-- RUBROS ACTIVOS: kiosco, almacen, ferreteria, ropa, deportes, libreria, pintureria
-- IDEMPOTENTE: Se puede re-ejecutar sin problemas

UPDATE business_type_templates
SET is_active = false, updated_at = CURRENT_TIMESTAMP
WHERE business_type_id IN (
  SELECT id FROM business_types
  WHERE code NOT IN (
    'kiosco',
    'almacen',
    'ferreteria',
    'pintureria',
    'ropa',
    'deportes',
    'libreria',
    'bazar',
    'electricidad',
    'sanitarios',
    'zapateria'
  )
);

UPDATE business_type_templates
SET is_active = true, updated_at = CURRENT_TIMESTAMP
WHERE business_type_id IN (
  SELECT id FROM business_types
  WHERE code IN (
    'kiosco',
    'almacen',
    'ferreteria',
    'pintureria',
    'ropa',
    'deportes',
    'libreria',
    'bazar',
    'electricidad',
    'sanitarios',
    'zapateria'
  )
);
