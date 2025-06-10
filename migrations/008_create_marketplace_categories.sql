-- Migration 008: Create marketplace_categories table
-- PROPÓSITO: Crear estructura de navegación COMÚN para compradores
-- BENEFICIO: Seller no piensa en taxonomías complejas, elige entre pocas opciones claras

CREATE TABLE marketplace_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,        -- "Remeras y Tops"
    slug VARCHAR(255) NOT NULL UNIQUE, -- "fashion-tops" 
    description TEXT,                   -- "Remeras, musculosas, tops"
    parent_id UUID REFERENCES marketplace_categories(id),
    level INTEGER NOT NULL DEFAULT 0,  -- Máximo 3 niveles de profundidad
    is_active BOOLEAN DEFAULT TRUE,
    sort_order INTEGER DEFAULT 0,      -- Control de orden en navegación
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Constraints
    CONSTRAINT marketplace_categories_name_not_empty CHECK (LENGTH(TRIM(name)) > 0),
    CONSTRAINT marketplace_categories_slug_not_empty CHECK (LENGTH(TRIM(slug)) > 0),
    CONSTRAINT marketplace_categories_level_valid CHECK (level >= 0 AND level <= 3),
    CONSTRAINT marketplace_categories_no_self_reference CHECK (id != parent_id)
);

-- Índices para performance
CREATE INDEX IF NOT EXISTS idx_marketplace_categories_parent_id ON marketplace_categories(parent_id);
CREATE INDEX IF NOT EXISTS idx_marketplace_categories_slug ON marketplace_categories(slug);
CREATE INDEX IF NOT EXISTS idx_marketplace_categories_level ON marketplace_categories(level);
CREATE INDEX IF NOT EXISTS idx_marketplace_categories_active ON marketplace_categories(is_active);
CREATE INDEX IF NOT EXISTS idx_marketplace_categories_sort ON marketplace_categories(sort_order);
CREATE INDEX IF NOT EXISTS idx_marketplace_categories_name ON marketplace_categories(name);

-- Índice compuesto para navegación jerárquica
CREATE INDEX IF NOT EXISTS idx_marketplace_categories_hierarchy ON marketplace_categories(parent_id, sort_order, is_active);

-- Comentarios para documentación
COMMENT ON TABLE marketplace_categories IS 'Categorías globales del marketplace para navegación consistente cross-tenant';
COMMENT ON COLUMN marketplace_categories.name IS 'Nombre de la categoría mostrado al usuario';
COMMENT ON COLUMN marketplace_categories.slug IS 'Identificador único para URLs amigables';
COMMENT ON COLUMN marketplace_categories.level IS 'Profundidad en la jerarquía (0=raíz, máximo 3)';
COMMENT ON COLUMN marketplace_categories.parent_id IS 'Referencia a categoría padre, NULL para categorías raíz';
COMMENT ON COLUMN marketplace_categories.sort_order IS 'Orden de presentación dentro del mismo nivel';
COMMENT ON COLUMN marketplace_categories.is_active IS 'Indica si la categoría está disponible para nuevos productos';

-- Trigger para actualizar updated_at automáticamente
CREATE OR REPLACE FUNCTION update_marketplace_categories_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_marketplace_categories_updated_at
    BEFORE UPDATE ON marketplace_categories
    FOR EACH ROW
    EXECUTE FUNCTION update_marketplace_categories_updated_at();

-- Función para validar la profundidad de la jerarquía
CREATE OR REPLACE FUNCTION validate_marketplace_category_hierarchy()
RETURNS TRIGGER AS $$
DECLARE
    current_level INTEGER;
BEGIN
    -- Si es una categoría raíz, el nivel es 0
    IF NEW.parent_id IS NULL THEN
        NEW.level = 0;
        RETURN NEW;
    END IF;
    
    -- Calcular el nivel basado en el padre
    SELECT level + 1 INTO current_level
    FROM marketplace_categories
    WHERE id = NEW.parent_id;
    
    -- Validar que no exceda el máximo de 3 niveles
    IF current_level > 3 THEN
        RAISE EXCEPTION 'Marketplace categories cannot exceed 3 levels of depth';
    END IF;
    
    NEW.level = current_level;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_validate_marketplace_category_hierarchy
    BEFORE INSERT OR UPDATE ON marketplace_categories
    FOR EACH ROW
    EXECUTE FUNCTION validate_marketplace_category_hierarchy(); 