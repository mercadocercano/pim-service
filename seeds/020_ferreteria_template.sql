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
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, sort_order) VALUES
('fa1e8f2a-0000-0000-0000-000000000001', 'SKU', 'sku-ferreteria', 'text', false, 1),
('fa1e8f2a-0000-0000-0000-000000000002', 'Marca Ferretería', 'marca-ferreteria', 'text', true, 2),
('fa1e8f2a-0000-0000-0000-000000000003', 'Unidad de Venta', 'unidad-venta', 'select', true, 3)
ON CONFLICT (slug) DO UPDATE SET
  name = EXCLUDED.name,
  type = EXCLUDED.type,
  is_filterable = EXCLUDED.is_filterable,
  updated_at = NOW();

-- Atributos específicos por categoría

-- Tornillería: diámetro, largo, material
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, sort_order) VALUES
('fa1e8f2a-0001-0000-0000-000000000001', 'Diámetro', 'diametro-tornillo', 'text', true, 10),
('fa1e8f2a-0001-0000-0000-000000000002', 'Largo', 'largo-tornillo', 'text', true, 11),
('fa1e8f2a-0001-0000-0000-000000000003', 'Material Tornillería', 'material-tornilleria', 'select', true, 12)
ON CONFLICT (slug) DO UPDATE SET
  name = EXCLUDED.name,
  type = EXCLUDED.type,
  is_filterable = EXCLUDED.is_filterable,
  updated_at = NOW();

-- Herramientas manuales: tipo, material
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, sort_order) VALUES
('fa1e8f2a-0002-0000-0000-000000000001', 'Tipo Herramienta', 'tipo-herramienta', 'select', true, 20),
('fa1e8f2a-0002-0000-0000-000000000002', 'Material Herramienta', 'material-herramienta', 'select', true, 21)
ON CONFLICT (slug) DO UPDATE SET
  name = EXCLUDED.name,
  type = EXCLUDED.type,
  is_filterable = EXCLUDED.is_filterable,
  updated_at = NOW();

-- Herramientas eléctricas: potencia, voltaje
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, sort_order) VALUES
('fa1e8f2a-0003-0000-0000-000000000001', 'Potencia', 'potencia-herramienta', 'text', true, 30),
('fa1e8f2a-0003-0000-0000-000000000002', 'Voltaje', 'voltaje-herramienta', 'select', true, 31)
ON CONFLICT (slug) DO UPDATE SET
  name = EXCLUDED.name,
  type = EXCLUDED.type,
  is_filterable = EXCLUDED.is_filterable,
  updated_at = NOW();

-- Materiales construcción: presentación, peso
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, sort_order) VALUES
('fa1e8f2a-0004-0000-0000-000000000001', 'Presentación Material', 'presentacion-material', 'select', true, 40),
('fa1e8f2a-0004-0000-0000-000000000002', 'Peso', 'peso-material', 'text', false, 41)
ON CONFLICT (slug) DO UPDATE SET
  name = EXCLUDED.name,
  type = EXCLUDED.type,
  is_filterable = EXCLUDED.is_filterable,
  updated_at = NOW();

-- Pinturas: color, presentación, rendimiento
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, sort_order) VALUES
('fa1e8f2a-0005-0000-0000-000000000001', 'Color Pintura', 'color-pintura', 'text', true, 50),
('fa1e8f2a-0005-0000-0000-000000000002', 'Presentación Pintura', 'presentacion-pintura', 'select', true, 51),
('fa1e8f2a-0005-0000-0000-000000000003', 'Rendimiento Pintura', 'rendimiento-pintura', 'text', false, 52)
ON CONFLICT (slug) DO UPDATE SET
  name = EXCLUDED.name,
  type = EXCLUDED.type,
  is_filterable = EXCLUDED.is_filterable,
  updated_at = NOW();

-- Plomería: diámetro, material
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, sort_order) VALUES
('fa1e8f2a-0006-0000-0000-000000000001', 'Diámetro Plomería', 'diametro-plomeria', 'text', true, 60),
('fa1e8f2a-0006-0000-0000-000000000002', 'Material Plomería', 'material-plomeria', 'select', true, 61)
ON CONFLICT (slug) DO UPDATE SET
  name = EXCLUDED.name,
  type = EXCLUDED.type,
  is_filterable = EXCLUDED.is_filterable,
  updated_at = NOW();

-- =====================================================
-- PASO 3: Crear valores de atributos (marketplace_attribute_values)
-- =====================================================

-- Unidad de Venta
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('fa1e8f2a-0000-0000-0000-000000000003', 'Unidad', 'unidad-valor', 1),
('fa1e8f2a-0000-0000-0000-000000000003', 'Kilogramo', 'kg-valor', 2),
('fa1e8f2a-0000-0000-0000-000000000003', 'Bolsa', 'bolsa-valor', 3),
('fa1e8f2a-0000-0000-0000-000000000003', 'Metro', 'metro-valor', 4),
('fa1e8f2a-0000-0000-0000-000000000003', 'Litro', 'litro-valor', 5),
('fa1e8f2a-0000-0000-0000-000000000003', 'Caja', 'caja-valor', 6)
ON CONFLICT (attribute_id, slug) DO UPDATE SET
  value = EXCLUDED.value,
  sort_order = EXCLUDED.sort_order;

-- Material Tornillería
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('fa1e8f2a-0001-0000-0000-000000000003', 'Acero', 'acero-torn', 1),
('fa1e8f2a-0001-0000-0000-000000000003', 'Acero Inoxidable', 'acero-inoxidable-torn', 2),
('fa1e8f2a-0001-0000-0000-000000000003', 'Bronce', 'bronce-torn', 3),
('fa1e8f2a-0001-0000-0000-000000000003', 'Aluminio', 'aluminio-torn', 4)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Tipo Herramienta Manual
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('fa1e8f2a-0002-0000-0000-000000000001', 'Martillo', 'martillo-tipo', 1),
('fa1e8f2a-0002-0000-0000-000000000001', 'Destornillador', 'destornillador-tipo', 2),
('fa1e8f2a-0002-0000-0000-000000000001', 'Llave', 'llave-tipo', 3),
('fa1e8f2a-0002-0000-0000-000000000001', 'Pinza', 'pinza-tipo', 4),
('fa1e8f2a-0002-0000-0000-000000000001', 'Sierra', 'sierra-tipo', 5)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Material Herramienta
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('fa1e8f2a-0002-0000-0000-000000000002', 'Acero', 'acero-herr', 1),
('fa1e8f2a-0002-0000-0000-000000000002', 'Acero Forjado', 'acero-forjado-herr', 2),
('fa1e8f2a-0002-0000-0000-000000000002', 'Madera y Acero', 'madera-acero-herr', 3)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Voltaje
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('fa1e8f2a-0003-0000-0000-000000000002', '220V', '220v-volt', 1),
('fa1e8f2a-0003-0000-0000-000000000002', '110V', '110v-volt', 2),
('fa1e8f2a-0003-0000-0000-000000000002', '12V', '12v-volt', 3)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Presentación
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('fa1e8f2a-0004-0000-0000-000000000001', 'Bolsa', 'bolsa-pres', 1),
('fa1e8f2a-0004-0000-0000-000000000001', 'Saco', 'saco-pres', 2),
('fa1e8f2a-0004-0000-0000-000000000001', 'Unidad', 'unidad-pres', 3),
('fa1e8f2a-0004-0000-0000-000000000001', 'Pallet', 'pallet-pres', 4)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Presentación Pintura
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('fa1e8f2a-0005-0000-0000-000000000002', 'Lata 1L', 'lata-1l-pint', 1),
('fa1e8f2a-0005-0000-0000-000000000002', 'Lata 4L', 'lata-4l-pint', 2),
('fa1e8f2a-0005-0000-0000-000000000002', 'Balde 10L', 'balde-10l-pint', 3),
('fa1e8f2a-0005-0000-0000-000000000002', 'Balde 20L', 'balde-20l-pint', 4)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- Material Plomería
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('fa1e8f2a-0006-0000-0000-000000000002', 'PVC', 'pvc-plom', 1),
('fa1e8f2a-0006-0000-0000-000000000002', 'Cobre', 'cobre-plom', 2),
('fa1e8f2a-0006-0000-0000-000000000002', 'Acero Galvanizado', 'acero-galvanizado-plom', 3),
('fa1e8f2a-0006-0000-0000-000000000002', 'Bronce', 'bronce-plom', 4)
ON CONFLICT (attribute_id, slug) DO NOTHING;

-- =====================================================
-- PASO 4: Relacionar categorías con atributos (marketplace_category_attributes)
-- =====================================================

-- TORNILLERÍA
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'fa1e8f2a-0000-0000-0000-000000000001', true, 1),  -- SKU
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'fa1e8f2a-0000-0000-0000-000000000002', true, 2),  -- Marca
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'fa1e8f2a-0000-0000-0000-000000000003', true, 3),  -- Unidad de Venta
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'fa1e8f2a-0001-0000-0000-000000000001', false, 4), -- Diámetro
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'fa1e8f2a-0001-0000-0000-000000000002', false, 5), -- Largo
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f01', 'fa1e8f2a-0001-0000-0000-000000000003', false, 6)  -- Material
ON CONFLICT (category_id, attribute_id) DO UPDATE SET
  is_required = EXCLUDED.is_required,
  sort_order = EXCLUDED.sort_order;

-- HERRAMIENTAS MANUALES
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'fa1e8f2a-0000-0000-0000-000000000001', true, 1),  -- SKU
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'fa1e8f2a-0000-0000-0000-000000000002', true, 2),  -- Marca
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'fa1e8f2a-0000-0000-0000-000000000003', true, 3),  -- Unidad de Venta
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'fa1e8f2a-0002-0000-0000-000000000001', false, 4), -- Tipo
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f02', 'fa1e8f2a-0002-0000-0000-000000000002', false, 5)  -- Material
ON CONFLICT (category_id, attribute_id) DO UPDATE SET
  is_required = EXCLUDED.is_required,
  sort_order = EXCLUDED.sort_order;

-- HERRAMIENTAS ELÉCTRICAS
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'fa1e8f2a-0000-0000-0000-000000000001', true, 1),  -- SKU
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'fa1e8f2a-0000-0000-0000-000000000002', true, 2),  -- Marca
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'fa1e8f2a-0000-0000-0000-000000000003', true, 3),  -- Unidad de Venta
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'fa1e8f2a-0003-0000-0000-000000000001', false, 4), -- Potencia
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f03', 'fa1e8f2a-0003-0000-0000-000000000002', false, 5)  -- Voltaje
ON CONFLICT (category_id, attribute_id) DO UPDATE SET
  is_required = EXCLUDED.is_required,
  sort_order = EXCLUDED.sort_order;

-- MATERIALES DE CONSTRUCCIÓN
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', 'fa1e8f2a-0000-0000-0000-000000000001', true, 1),  -- SKU
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', 'fa1e8f2a-0000-0000-0000-000000000002', true, 2),  -- Marca
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', 'fa1e8f2a-0000-0000-0000-000000000003', true, 3),  -- Unidad de Venta
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', 'fa1e8f2a-0004-0000-0000-000000000001', false, 4), -- Presentación
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f04', 'fa1e8f2a-0004-0000-0000-000000000002', false, 5)  -- Peso
ON CONFLICT (category_id, attribute_id) DO UPDATE SET
  is_required = EXCLUDED.is_required,
  sort_order = EXCLUDED.sort_order;

-- PINTURAS
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'fa1e8f2a-0000-0000-0000-000000000001', true, 1),  -- SKU
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'fa1e8f2a-0000-0000-0000-000000000002', true, 2),  -- Marca
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'fa1e8f2a-0000-0000-0000-000000000003', true, 3),  -- Unidad de Venta
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'fa1e8f2a-0005-0000-0000-000000000001', false, 4), -- Color
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'fa1e8f2a-0005-0000-0000-000000000002', false, 5), -- Presentación
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f05', 'fa1e8f2a-0005-0000-0000-000000000003', false, 6)  -- Rendimiento
ON CONFLICT (category_id, attribute_id) DO UPDATE SET
  is_required = EXCLUDED.is_required,
  sort_order = EXCLUDED.sort_order;

-- PLOMERÍA Y SANITARIOS
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order) VALUES
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'fa1e8f2a-0000-0000-0000-000000000001', true, 1),  -- SKU
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'fa1e8f2a-0000-0000-0000-000000000002', true, 2),  -- Marca
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'fa1e8f2a-0000-0000-0000-000000000003', true, 3),  -- Unidad de Venta
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'fa1e8f2a-0006-0000-0000-000000000001', false, 4), -- Diámetro
('f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f06', 'fa1e8f2a-0006-0000-0000-000000000002', false, 5)  -- Material
ON CONFLICT (category_id, attribute_id) DO UPDATE SET
  is_required = EXCLUDED.is_required,
  sort_order = EXCLUDED.sort_order;

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
    WHERE id::text LIKE 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f%';
  
  SELECT COUNT(*) INTO attr_count FROM marketplace_attributes 
    WHERE id::text LIKE 'fa1e8f2a-%';
  
  SELECT COUNT(*) INTO rel_count FROM marketplace_category_attributes 
    WHERE category_id::text LIKE 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f%';
  
  RAISE NOTICE '✅ Template Ferretería creado exitosamente:';
  RAISE NOTICE '   - Categorías: %', cat_count;
  RAISE NOTICE '   - Atributos: %', attr_count;
  RAISE NOTICE '   - Relaciones: %', rel_count;
END $$;

