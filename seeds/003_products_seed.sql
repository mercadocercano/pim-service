-- Seeds para productos de prueba
-- Nota: Estos seeds asumen que ya existen categorías y marcas de los seeds anteriores

-- Productos de tecnología
INSERT INTO products (id, tenant_id, name, description, sku, category_id, category_name, brand_id, brand_name, status, created_at, updated_at) VALUES
-- Productos Apple
('550e8400-e29b-41d4-a716-446655440101', '550e8400-e29b-41d4-a716-446655440000', 'iPhone 15 Pro', 'Smartphone Apple con chip A17 Pro y cámara de 48MP', 'IPHONE-15-PRO-128GB', '550e8400-e29b-41d4-a716-446655440001', 'Smartphones', '550e8400-e29b-41d4-a716-446655440006', 'Apple', 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440102', '550e8400-e29b-41d4-a716-446655440000', 'iPhone 15 Pro Max', 'Smartphone Apple con chip A17 Pro y pantalla de 6.7 pulgadas', 'IPHONE-15-PRO-MAX-256GB', '550e8400-e29b-41d4-a716-446655440001', 'Smartphones', '550e8400-e29b-41d4-a716-446655440006', 'Apple', 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440103', '550e8400-e29b-41d4-a716-446655440000', 'MacBook Pro 14"', 'Laptop profesional con chip M3 Pro y pantalla Liquid Retina XDR', 'MACBOOK-PRO-14-M3-512GB', '550e8400-e29b-41d4-a716-446655440002', 'Laptops', '550e8400-e29b-41d4-a716-446655440006', 'Apple', 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440104', '550e8400-e29b-41d4-a716-446655440000', 'iPad Air', 'Tablet con chip M2 y pantalla Liquid Retina de 10.9 pulgadas', 'IPAD-AIR-M2-256GB', '550e8400-e29b-41d4-a716-446655440003', 'Tablets', '550e8400-e29b-41d4-a716-446655440006', 'Apple', 'active', NOW(), NOW()),

-- Productos Samsung
('550e8400-e29b-41d4-a716-446655440105', '550e8400-e29b-41d4-a716-446655440000', 'Galaxy S24 Ultra', 'Smartphone Samsung con S Pen y cámara de 200MP', 'GALAXY-S24-ULTRA-512GB', '550e8400-e29b-41d4-a716-446655440001', 'Smartphones', '550e8400-e29b-41d4-a716-446655440007', 'Samsung', 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440106', '550e8400-e29b-41d4-a716-446655440000', 'Galaxy Book3 Pro', 'Laptop ultradelgada con procesador Intel Core i7', 'GALAXY-BOOK3-PRO-1TB', '550e8400-e29b-41d4-a716-446655440002', 'Laptops', '550e8400-e29b-41d4-a716-446655440007', 'Samsung', 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440107', '550e8400-e29b-41d4-a716-446655440000', 'Galaxy Tab S9', 'Tablet Android con S Pen incluido y pantalla AMOLED', 'GALAXY-TAB-S9-256GB', '550e8400-e29b-41d4-a716-446655440003', 'Tablets', '550e8400-e29b-41d4-a716-446655440007', 'Samsung', 'active', NOW(), NOW()),

-- Productos Sony
('550e8400-e29b-41d4-a716-446655440108', '550e8400-e29b-41d4-a716-446655440000', 'PlayStation 5', 'Consola de videojuegos de nueva generación con SSD ultra rápido', 'PS5-STANDARD-825GB', '550e8400-e29b-41d4-a716-446655440004', 'Consolas', '550e8400-e29b-41d4-a716-446655440008', 'Sony', 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440109', '550e8400-e29b-41d4-a716-446655440000', 'WH-1000XM5', 'Audífonos inalámbricos con cancelación de ruido líder en la industria', 'WH-1000XM5-BLACK', '550e8400-e29b-41d4-a716-446655440005', 'Audio', '550e8400-e29b-41d4-a716-446655440008', 'Sony', 'active', NOW(), NOW()),

-- Productos sin categoría o marca específica
('550e8400-e29b-41d4-a716-446655440110', '550e8400-e29b-41d4-a716-446655440000', 'Monitor Gaming 27"', 'Monitor para gaming con resolución 4K y 144Hz', 'MONITOR-GAMING-27-4K', NULL, NULL, NULL, NULL, 'active', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440111', '550e8400-e29b-41d4-a716-446655440000', 'Teclado Mecánico RGB', 'Teclado mecánico para gaming con iluminación RGB personalizable', 'KEYBOARD-MECH-RGB', NULL, NULL, NULL, NULL, 'active', NOW(), NOW()),

-- Productos con diferentes estados
('550e8400-e29b-41d4-a716-446655440112', '550e8400-e29b-41d4-a716-446655440000', 'iPhone 14', 'Modelo anterior de iPhone, ahora descontinuado', 'IPHONE-14-128GB', '550e8400-e29b-41d4-a716-446655440001', 'Smartphones', '550e8400-e29b-41d4-a716-446655440006', 'Apple', 'discontinued', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440113', '550e8400-e29b-41d4-a716-446655440000', 'Galaxy S23', 'Modelo anterior de Galaxy S, temporalmente inactivo', 'GALAXY-S23-256GB', '550e8400-e29b-41d4-a716-446655440001', 'Smartphones', '550e8400-e29b-41d4-a716-446655440007', 'Samsung', 'inactive', NOW(), NOW()),

-- Productos para otro tenant (para probar aislamiento)
('550e8400-e29b-41d4-a716-446655440114', '550e8400-e29b-41d4-a716-446655440001', 'Producto Tenant 2', 'Producto que pertenece a otro tenant', 'PROD-TENANT2-001', NULL, NULL, NULL, NULL, 'active', NOW(), NOW());

-- Actualizar contadores de productos por categoría (esto sería manejado por triggers en un sistema real)
-- Por ahora solo insertamos los datos base 