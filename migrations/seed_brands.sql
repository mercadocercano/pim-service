-- Seed data for brands table
-- Description: Datos de prueba para marcas populares
-- Date: 2024-01-15

-- Marcas de moda y deportes
INSERT INTO brands (id, tenant_id, name, description, logo_url, website, status) VALUES
('550e8400-e29b-41d4-a716-446655440001', '123e4567-e89b-12d3-a456-426614174000', 'Nike', 'Marca deportiva internacional líder en calzado y ropa deportiva', 'https://logos.com/nike-logo.png', 'https://nike.com', 'active'),
('550e8400-e29b-41d4-a716-446655440002', '123e4567-e89b-12d3-a456-426614174000', 'Adidas', 'Marca alemana de artículos deportivos', 'https://logos.com/adidas-logo.png', 'https://adidas.com', 'active'),
('550e8400-e29b-41d4-a716-446655440003', '123e4567-e89b-12d3-a456-426614174000', 'Zara', 'Cadena española de moda rápida', 'https://logos.com/zara-logo.png', 'https://zara.com', 'active'),
('550e8400-e29b-41d4-a716-446655440004', '123e4567-e89b-12d3-a456-426614174000', 'H&M', 'Marca sueca de moda accesible', 'https://logos.com/hm-logo.png', 'https://hm.com', 'active'),
('550e8400-e29b-41d4-a716-446655440005', '123e4567-e89b-12d3-a456-426614174000', 'Puma', 'Marca alemana de artículos deportivos', 'https://logos.com/puma-logo.png', 'https://puma.com', 'active');

-- Marcas de electrónicos
INSERT INTO brands (id, tenant_id, name, description, logo_url, website, status) VALUES
('550e8400-e29b-41d4-a716-446655440006', '123e4567-e89b-12d3-a456-426614174000', 'Apple', 'Empresa tecnológica estadounidense', 'https://logos.com/apple-logo.png', 'https://apple.com', 'active'),
('550e8400-e29b-41d4-a716-446655440007', '123e4567-e89b-12d3-a456-426614174000', 'Samsung', 'Conglomerado surcoreano de electrónicos', 'https://logos.com/samsung-logo.png', 'https://samsung.com', 'active'),
('550e8400-e29b-41d4-a716-446655440008', '123e4567-e89b-12d3-a456-426614174000', 'Sony', 'Empresa japonesa de electrónicos y entretenimiento', 'https://logos.com/sony-logo.png', 'https://sony.com', 'active'),
('550e8400-e29b-41d4-a716-446655440009', '123e4567-e89b-12d3-a456-426614174000', 'LG', 'Empresa surcoreana de electrónicos', 'https://logos.com/lg-logo.png', 'https://lg.com', 'active'),
('550e8400-e29b-41d4-a716-446655440010', '123e4567-e89b-12d3-a456-426614174000', 'Xiaomi', 'Empresa china de electrónicos', 'https://logos.com/xiaomi-logo.png', 'https://xiaomi.com', 'active');

-- Marcas de alimentación
INSERT INTO brands (id, tenant_id, name, description, logo_url, website, status) VALUES
('550e8400-e29b-41d4-a716-446655440011', '123e4567-e89b-12d3-a456-426614174000', 'Coca-Cola', 'Empresa de bebidas estadounidense', 'https://logos.com/cocacola-logo.png', 'https://coca-cola.com', 'active'),
('550e8400-e29b-41d4-a716-446655440012', '123e4567-e89b-12d3-a456-426614174000', 'Nestlé', 'Empresa suiza de alimentos y bebidas', 'https://logos.com/nestle-logo.png', 'https://nestle.com', 'active'),
('550e8400-e29b-41d4-a716-446655440013', '123e4567-e89b-12d3-a456-426614174000', 'Unilever', 'Empresa anglo-holandesa de bienes de consumo', 'https://logos.com/unilever-logo.png', 'https://unilever.com', 'active'),
('550e8400-e29b-41d4-a716-446655440014', '123e4567-e89b-12d3-a456-426614174000', 'Danone', 'Empresa francesa de productos lácteos', 'https://logos.com/danone-logo.png', 'https://danone.com', 'active'),
('550e8400-e29b-41d4-a716-446655440015', '123e4567-e89b-12d3-a456-426614174000', 'Kelloggs', 'Empresa estadounidense de cereales', 'https://logos.com/kelloggs-logo.png', 'https://kelloggs.com', 'active');

-- Marcas genéricas/sin marca
INSERT INTO brands (id, tenant_id, name, description, status) VALUES
('550e8400-e29b-41d4-a716-446655440016', '123e4567-e89b-12d3-a456-426614174000', 'Genérica', 'Marca genérica para productos sin marca específica', 'active'),
('550e8400-e29b-41d4-a716-446655440017', '123e4567-e89b-12d3-a456-426614174000', 'Marca Propia', 'Marca propia del comercio', 'active');

-- Marcas para otro tenant (para testing)
INSERT INTO brands (id, tenant_id, name, description, status) VALUES
('550e8400-e29b-41d4-a716-446655440018', '456e7890-e89b-12d3-a456-426614174001', 'Nike', 'Marca deportiva para otro tenant', 'active'),
('550e8400-e29b-41d4-a716-446655440019', '456e7890-e89b-12d3-a456-426614174001', 'Local Brand', 'Marca local del segundo tenant', 'active'); 