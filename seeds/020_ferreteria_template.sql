-- Seed 020: Template de Quickstart para Ferretería / Corralón
-- PROPÓSITO: Template predefinido para negocios de ferretería con categorías y atributos específicos
-- HITO 2: Quickstart Ferretería - Template mínimo con 6 categorías

-- =====================================================
-- PASO 1: Crear categorías de ferretería (marketplace_categories)
-- =====================================================

-- Categoría padre: Ferretería y Construcción
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order, is_active) VALUES
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f00', 'Ferretería y Construcción', 'ferreteria-construccion', 'Productos para ferretería, construcción y mejoras del hogar', NULL, 0, 100, true)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  slug = EXCLUDED.slug,
  description = EXCLUDED.description,
  is_active = EXCLUDED.is_active,
  updated_at = NOW();

-- 6 Categorías específicas (nivel 1)
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order, is_active) VALUES
-- 1. Tornillería
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'Tornillería', 'tornilleria', 'Tornillos, tuercas, arandelas y fijaciones', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f00', 1, 1, true),

-- 2. Herramientas manuales
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'Herramientas Manuales', 'herramientas-manuales', 'Martillos, destornilladores, llaves y herramientas de mano', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f00', 1, 2, true),

-- 3. Herramientas eléctricas
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'Herramientas Eléctricas', 'herramientas-electricas', 'Taladros, amoladoras, sierras eléctricas', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f00', 1, 3, true),

-- 4. Materiales de construcción
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', 'Materiales de Construcción', 'materiales-construccion', 'Cemento, cal, arena, ladrillos', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f00', 1, 4, true),

-- 5. Pinturas
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'Pinturas', 'pinturas', 'Pinturas, barnices, esmaltes y accesorios', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f00', 1, 5, true),

-- 6. Plomería / Sanitarios
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'Plomería y Sanitarios', 'plomeria-sanitarios', 'Caños, accesorios, sanitarios y griferías', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f00', 1, 6, true)

ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  slug = EXCLUDED.slug,
  description = EXCLUDED.description,
  parent_id = EXCLUDED.parent_id,
  level = EXCLUDED.level,
  sort_order = EXCLUDED.sort_order,
  is_active = EXCLUDED.is_active,
  updated_at = NOW();

-- =====================================================
-- PASO 2: Crear atributos marketplace (marketplace_attributes)
-- =====================================================

-- Atributos globales (aplicables a todos los productos)
INSERT INTO marketplace_attributes (id, name, slug, description, attribute_type, is_required, is_filterable, sort_order) VALUES
('fa1e8f2a-0000-0000-0000-000000000001', 'SKU', 'sku', 'Código único del producto', 'text', true, false, 1),
('fa1e8f2a-0000-0000-0000-000000000002', 'Marca', 'marca', 'Marca del producto', 'text', true, true, 2),
('fa1e8f2a-0000-0000-0000-000000000003', 'Unidad de Venta', 'unidad-venta', 'Unidad de medida para la venta', 'select', true, true, 3)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  slug = EXCLUDED.slug,
  description = EXCLUDED.description,
  attribute_type = EXCLUDED.attribute_type,
  is_required = EXCLUDED.is_required,
  is_filterable = EXCLUDED.is_filterable,
  updated_at = NOW();

-- Atributos específicos por categoría

-- Tornillería: diámetro, largo, material
INSERT INTO marketplace_attributes (id, name, slug, description, attribute_type, is_required, is_filterable, sort_order) VALUES
('fa1e8f2a-0001-0000-0000-000000000001', 'Diámetro', 'diametro', 'Diámetro del tornillo en mm', 'text', false, true, 10),
('fa1e8f2a-0001-0000-0000-000000000002', 'Largo', 'largo', 'Largo del tornillo (con unidad)', 'text', false, true, 11),
('fa1e8f2a-0001-0000-0000-000000000003', 'Material Tornillería', 'material-tornilleria', 'Material del tornillo/fijación', 'select', false, true, 12)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  slug = EXCLUDED.slug,
  description = EXCLUDED.description,
  updated_at = NOW();

-- Herramientas manuales: tipo, material
INSERT INTO marketplace_attributes (id, name, slug, description, attribute_type, is_required, is_filterable, sort_order) VALUES
('fa1e8f2a-0002-0000-0000-000000000001', 'Tipo Herramienta', 'tipo-herramienta', 'Tipo de herramienta manual', 'select', false, true, 20),
('fa1e8f2a-0002-0000-0000-000000000002', 'Material Herramienta', 'material-herramienta', 'Material de fabricación', 'select', false, true, 21)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  slug = EXCLUDED.slug,
  description = EXCLUDED.description,
  updated_at = NOW();

-- Herramientas eléctricas: potencia, voltaje
INSERT INTO marketplace_attributes (id, name, slug, description, attribute_type, is_required, is_filterable, sort_order) VALUES
('fa1e8f2a-0003-0000-0000-000000000001', 'Potencia', 'potencia', 'Potencia en watts (W)', 'text', false, true, 30),
('fa1e8f2a-0003-0000-0000-000000000002', 'Voltaje', 'voltaje', 'Voltaje de operación (V)', 'select', false, true, 31)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  slug = EXCLUDED.slug,
  description = EXCLUDED.description,
  updated_at = NOW();

-- Materiales construcción: presentación, peso
INSERT INTO marketplace_attributes (id, name, slug, description, attribute_type, is_required, is_filterable, sort_order) VALUES
('fa1e8f2a-0004-0000-0000-000000000001', 'Presentación', 'presentacion', 'Formato de presentación del producto', 'select', false, true, 40),
('fa1e8f2a-0004-0000-0000-000000000002', 'Peso', 'peso', 'Peso del producto (con unidad)', 'text', false, false, 41)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  slug = EXCLUDED.slug,
  description = EXCLUDED.description,
  updated_at = NOW();

-- Pinturas: color, presentación, rendimiento
INSERT INTO marketplace_attributes (id, name, slug, description, attribute_type, is_required, is_filterable, sort_order) VALUES
('fa1e8f2a-0005-0000-0000-000000000001', 'Color', 'color', 'Color del producto', 'text', false, true, 50),
('fa1e8f2a-0005-0000-0000-000000000002', 'Presentación Pintura', 'presentacion-pintura', 'Formato de presentación (lata, balde)', 'select', false, true, 51),
('fa1e8f2a-0005-0000-0000-000000000003', 'Rendimiento', 'rendimiento', 'Rendimiento por litro', 'text', false, false, 52)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  slug = EXCLUDED.slug,
  description = EXCLUDED.description,
  updated_at = NOW();

-- Plomería: diámetro, material
INSERT INTO marketplace_attributes (id, name, slug, description, attribute_type, is_required, is_filterable, sort_order) VALUES
('fa1e8f2a-0006-0000-0000-000000000001', 'Diámetro Plomería', 'diametro-plomeria', 'Diámetro de caño o accesorio', 'text', false, true, 60),
('fa1e8f2a-0006-0000-0000-000000000002', 'Material Plomería', 'material-plomeria', 'Material del producto', 'select', false, true, 61)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  slug = EXCLUDED.slug,
  description = EXCLUDED.description,
  updated_at = NOW();

-- =====================================================
-- PASO 3: Crear valores de atributos (marketplace_attribute_values)
-- =====================================================

-- Unidad de Venta
INSERT INTO marketplace_attribute_values (id, attribute_id, value, slug, sort_order) VALUES
('fv1-unidad-001', 'fa1e8f2a-0000-0000-0000-000000000003', 'Unidad', 'unidad', 1),
('fv1-unidad-002', 'fa1e8f2a-0000-0000-0000-000000000003', 'Kilogramo', 'kg', 2),
('fv1-unidad-003', 'fa1e8f2a-0000-0000-0000-000000000003', 'Bolsa', 'bolsa', 3),
('fv1-unidad-004', 'fa1e8f2a-0000-0000-0000-000000000003', 'Metro', 'metro', 4),
('fv1-unidad-005', 'fa1e8f2a-0000-0000-0000-000000000003', 'Litro', 'litro', 5),
('fv1-unidad-006', 'fa1e8f2a-0000-0000-0000-000000000003', 'Caja', 'caja', 6)
ON CONFLICT (id) DO UPDATE SET
  value = EXCLUDED.value,
  slug = EXCLUDED.slug,
  updated_at = NOW();

-- Material Tornillería
INSERT INTO marketplace_attribute_values (id, attribute_id, value, slug, sort_order) VALUES
('fv1-mat-torn-001', 'fa1e8f2a-0001-0000-0000-000000000003', 'Acero', 'acero', 1),
('fv1-mat-torn-002', 'fa1e8f2a-0001-0000-0000-000000000003', 'Acero Inoxidable', 'acero-inoxidable', 2),
('fv1-mat-torn-003', 'fa1e8f2a-0001-0000-0000-000000000003', 'Bronce', 'bronce', 3),
('fv1-mat-torn-004', 'fa1e8f2a-0001-0000-0000-000000000003', 'Aluminio', 'aluminio', 4)
ON CONFLICT (id) DO UPDATE SET
  value = EXCLUDED.value,
  slug = EXCLUDED.slug,
  updated_at = NOW();

-- Tipo Herramienta Manual
INSERT INTO marketplace_attribute_values (id, attribute_id, value, slug, sort_order) VALUES
('fv1-tipo-herr-001', 'fa1e8f2a-0002-0000-0000-000000000001', 'Martillo', 'martillo', 1),
('fv1-tipo-herr-002', 'fa1e8f2a-0002-0000-0000-000000000001', 'Destornillador', 'destornillador', 2),
('fv1-tipo-herr-003', 'fa1e8f2a-0002-0000-0000-000000000001', 'Llave', 'llave', 3),
('fv1-tipo-herr-004', 'fa1e8f2a-0002-0000-0000-000000000001', 'Pinza', 'pinza', 4),
('fv1-tipo-herr-005', 'fa1e8f2a-0002-0000-0000-000000000001', 'Sierra', 'sierra', 5)
ON CONFLICT (id) DO UPDATE SET
  value = EXCLUDED.value,
  slug = EXCLUDED.slug,
  updated_at = NOW();

-- Material Herramienta
INSERT INTO marketplace_attribute_values (id, attribute_id, value, slug, sort_order) VALUES
('fv1-mat-herr-001', 'fa1e8f2a-0002-0000-0000-000000000002', 'Acero', 'acero', 1),
('fv1-mat-herr-002', 'fa1e8f2a-0002-0000-0000-000000000002', 'Acero Forjado', 'acero-forjado', 2),
('fv1-mat-herr-003', 'fa1e8f2a-0002-0000-0000-000000000002', 'Madera y Acero', 'madera-acero', 3)
ON CONFLICT (id) DO UPDATE SET
  value = EXCLUDED.value,
  slug = EXCLUDED.slug,
  updated_at = NOW();

-- Voltaje
INSERT INTO marketplace_attribute_values (id, attribute_id, value, slug, sort_order) VALUES
('fv1-voltaje-001', 'fa1e8f2a-0003-0000-0000-000000000002', '220V', '220v', 1),
('fv1-voltaje-002', 'fa1e8f2a-0003-0000-0000-000000000002', '110V', '110v', 2),
('fv1-voltaje-003', 'fa1e8f2a-0003-0000-0000-000000000002', '12V', '12v', 3)
ON CONFLICT (id) DO UPDATE SET
  value = EXCLUDED.value,
  slug = EXCLUDED.slug,
  updated_at = NOW();

-- Presentación
INSERT INTO marketplace_attribute_values (id, attribute_id, value, slug, sort_order) VALUES
('fv1-pres-001', 'fa1e8f2a-0004-0000-0000-000000000001', 'Bolsa', 'bolsa', 1),
('fv1-pres-002', 'fa1e8f2a-0004-0000-0000-000000000001', 'Saco', 'saco', 2),
('fv1-pres-003', 'fa1e8f2a-0004-0000-0000-000000000001', 'Unidad', 'unidad', 3),
('fv1-pres-004', 'fa1e8f2a-0004-0000-0000-000000000001', 'Pallet', 'pallet', 4)
ON CONFLICT (id) DO UPDATE SET
  value = EXCLUDED.value,
  slug = EXCLUDED.slug,
  updated_at = NOW();

-- Presentación Pintura
INSERT INTO marketplace_attribute_values (id, attribute_id, value, slug, sort_order) VALUES
('fv1-pres-pint-001', 'fa1e8f2a-0005-0000-0000-000000000002', 'Lata 1L', 'lata-1l', 1),
('fv1-pres-pint-002', 'fa1e8f2a-0005-0000-0000-000000000002', 'Lata 4L', 'lata-4l', 2),
('fv1-pres-pint-003', 'fa1e8f2a-0005-0000-0000-000000000002', 'Balde 10L', 'balde-10l', 3),
('fv1-pres-pint-004', 'fa1e8f2a-0005-0000-0000-000000000002', 'Balde 20L', 'balde-20l', 4)
ON CONFLICT (id) DO UPDATE SET
  value = EXCLUDED.value,
  slug = EXCLUDED.slug,
  updated_at = NOW();

-- Material Plomería
INSERT INTO marketplace_attribute_values (id, attribute_id, value, slug, sort_order) VALUES
('fv1-mat-plom-001', 'fa1e8f2a-0006-0000-0000-000000000002', 'PVC', 'pvc', 1),
('fv1-mat-plom-002', 'fa1e8f2a-0006-0000-0000-000000000002', 'Cobre', 'cobre', 2),
('fv1-mat-plom-003', 'fa1e8f2a-0006-0000-0000-000000000002', 'Acero Galvanizado', 'acero-galvanizado', 3),
('fv1-mat-plom-004', 'fa1e8f2a-0006-0000-0000-000000000002', 'Bronce', 'bronce', 4)
ON CONFLICT (id) DO UPDATE SET
  value = EXCLUDED.value,
  slug = EXCLUDED.slug,
  updated_at = NOW();

-- =====================================================
-- PASO 4: Relacionar categorías con atributos (marketplace_category_attributes)
-- =====================================================

-- TORNILLERÍA
INSERT INTO marketplace_category_attributes (id, category_id, attribute_id, is_required, sort_order) VALUES
('fca-torn-001', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'fa1e8f2a-0000-0000-0000-000000000001', true, 1),  -- SKU
('fca-torn-002', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'fa1e8f2a-0000-0000-0000-000000000002', true, 2),  -- Marca
('fca-torn-003', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'fa1e8f2a-0000-0000-0000-000000000003', true, 3),  -- Unidad de Venta
('fca-torn-004', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'fa1e8f2a-0001-0000-0000-000000000001', false, 4), -- Diámetro
('fca-torn-005', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'fa1e8f2a-0001-0000-0000-000000000002', false, 5), -- Largo
('fca-torn-006', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'fa1e8f2a-0001-0000-0000-000000000003', false, 6)  -- Material
ON CONFLICT (id) DO UPDATE SET
  is_required = EXCLUDED.is_required,
  sort_order = EXCLUDED.sort_order,
  updated_at = NOW();

-- HERRAMIENTAS MANUALES
INSERT INTO marketplace_category_attributes (id, category_id, attribute_id, is_required, sort_order) VALUES
('fca-hman-001', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'fa1e8f2a-0000-0000-0000-000000000001', true, 1),  -- SKU
('fca-hman-002', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'fa1e8f2a-0000-0000-0000-000000000002', true, 2),  -- Marca
('fca-hman-003', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'fa1e8f2a-0000-0000-0000-000000000003', true, 3),  -- Unidad de Venta
('fca-hman-004', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'fa1e8f2a-0002-0000-0000-000000000001', false, 4), -- Tipo
('fca-hman-005', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'fa1e8f2a-0002-0000-0000-000000000002', false, 5)  -- Material
ON CONFLICT (id) DO UPDATE SET
  is_required = EXCLUDED.is_required,
  sort_order = EXCLUDED.sort_order,
  updated_at = NOW();

-- HERRAMIENTAS ELÉCTRICAS
INSERT INTO marketplace_category_attributes (id, category_id, attribute_id, is_required, sort_order) VALUES
('fca-helec-001', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'fa1e8f2a-0000-0000-0000-000000000001', true, 1),  -- SKU
('fca-helec-002', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'fa1e8f2a-0000-0000-0000-000000000002', true, 2),  -- Marca
('fca-helec-003', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'fa1e8f2a-0000-0000-0000-000000000003', true, 3),  -- Unidad de Venta
('fca-helec-004', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'fa1e8f2a-0003-0000-0000-000000000001', false, 4), -- Potencia
('fca-helec-005', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'fa1e8f2a-0003-0000-0000-000000000002', false, 5)  -- Voltaje
ON CONFLICT (id) DO UPDATE SET
  is_required = EXCLUDED.is_required,
  sort_order = EXCLUDED.sort_order,
  updated_at = NOW();

-- MATERIALES DE CONSTRUCCIÓN
INSERT INTO marketplace_category_attributes (id, category_id, attribute_id, is_required, sort_order) VALUES
('fca-mat-001', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', 'fa1e8f2a-0000-0000-0000-000000000001', true, 1),  -- SKU
('fca-mat-002', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', 'fa1e8f2a-0000-0000-0000-000000000002', true, 2),  -- Marca
('fca-mat-003', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', 'fa1e8f2a-0000-0000-0000-000000000003', true, 3),  -- Unidad de Venta
('fca-mat-004', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', 'fa1e8f2a-0004-0000-0000-000000000001', false, 4), -- Presentación
('fca-mat-005', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', 'fa1e8f2a-0004-0000-0000-000000000002', false, 5)  -- Peso
ON CONFLICT (id) DO UPDATE SET
  is_required = EXCLUDED.is_required,
  sort_order = EXCLUDED.sort_order,
  updated_at = NOW();

-- PINTURAS
INSERT INTO marketplace_category_attributes (id, category_id, attribute_id, is_required, sort_order) VALUES
('fca-pint-001', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'fa1e8f2a-0000-0000-0000-000000000001', true, 1),  -- SKU
('fca-pint-002', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'fa1e8f2a-0000-0000-0000-000000000002', true, 2),  -- Marca
('fca-pint-003', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'fa1e8f2a-0000-0000-0000-000000000003', true, 3),  -- Unidad de Venta
('fca-pint-004', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'fa1e8f2a-0005-0000-0000-000000000001', false, 4), -- Color
('fca-pint-005', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'fa1e8f2a-0005-0000-0000-000000000002', false, 5), -- Presentación
('fca-pint-006', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'fa1e8f2a-0005-0000-0000-000000000003', false, 6)  -- Rendimiento
ON CONFLICT (id) DO UPDATE SET
  is_required = EXCLUDED.is_required,
  sort_order = EXCLUDED.sort_order,
  updated_at = NOW();

-- PLOMERÍA Y SANITARIOS
INSERT INTO marketplace_category_attributes (id, category_id, attribute_id, is_required, sort_order) VALUES
('fca-plom-001', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'fa1e8f2a-0000-0000-0000-000000000001', true, 1),  -- SKU
('fca-plom-002', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'fa1e8f2a-0000-0000-0000-000000000002', true, 2),  -- Marca
('fca-plom-003', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'fa1e8f2a-0000-0000-0000-000000000003', true, 3),  -- Unidad de Venta
('fca-plom-004', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'fa1e8f2a-0006-0000-0000-000000000001', false, 4), -- Diámetro
('fca-plom-005', 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'fa1e8f2a-0006-0000-0000-000000000002', false, 5)  -- Material
ON CONFLICT (id) DO UPDATE SET
  is_required = EXCLUDED.is_required,
  sort_order = EXCLUDED.sort_order,
  updated_at = NOW();

-- =====================================================
-- VERIFICACIÓN FINAL
-- =====================================================

-- Resumen de lo creado
DO $$
DECLARE
  cat_count INT;
  attr_count INT;
  rel_count INT;
BEGIN
  SELECT COUNT(*) INTO cat_count FROM marketplace_categories 
    WHERE id LIKE 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f%';
  
  SELECT COUNT(*) INTO attr_count FROM marketplace_attributes 
    WHERE id LIKE 'fa1e8f2a-%';
  
  SELECT COUNT(*) INTO rel_count FROM marketplace_category_attributes 
    WHERE id LIKE 'fca-%';
  
  RAISE NOTICE '✅ Template Ferretería creado exitosamente:';
  RAISE NOTICE '   - Categorías: %', cat_count;
  RAISE NOTICE '   - Atributos: %', attr_count;
  RAISE NOTICE '   - Relaciones: %', rel_count;
END $$;

