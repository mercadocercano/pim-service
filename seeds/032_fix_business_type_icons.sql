-- Seed 032: Fix de iconos de business_types para lucide-react
-- IDEMPOTENTE: solo UPDATE

UPDATE business_types SET icon = 'droplets' WHERE code = 'lubricentro';
UPDATE business_types SET icon = 'trophy' WHERE code = 'deportes';
UPDATE business_types SET icon = 'cup-soda' WHERE code = 'bazar';
