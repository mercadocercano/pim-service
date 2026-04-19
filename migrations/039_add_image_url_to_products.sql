-- Agregar campo image_url a productos del tenant
ALTER TABLE products ADD COLUMN IF NOT EXISTS image_url TEXT;

COMMENT ON COLUMN products.image_url IS 'URL de imagen principal del producto';
