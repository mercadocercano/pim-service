-- Migration 014: Create marketplace_brands table
-- PROPÓSITO: Crear catálogo global de marcas normalizadas para consistencia cross-tenant
-- BENEFICIO: Marcas normalizadas, enriquecidas y reutilizables entre tenants

CREATE TABLE marketplace_brands (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    normalized_name VARCHAR(255) NOT NULL, -- Nombre normalizado para matching
    description TEXT,
    logo_url VARCHAR(500),
    website VARCHAR(200),
    country_code VARCHAR(2) DEFAULT 'AR', -- Para marcas regionales
    category_tags TEXT[] DEFAULT '{}', -- Array de categorías asociadas
    verification_status VARCHAR(20) DEFAULT 'unverified', -- verified, unverified, disputed
    quality_score DECIMAL(3,2) DEFAULT 0.0, -- 0.0 - 1.0
    product_count INTEGER DEFAULT 0, -- Contador de productos que usan esta marca
    web_data JSONB DEFAULT '{}', -- Datos enriquecidos de APIs web
    aliases TEXT[] DEFAULT '{}', -- Variaciones del nombre de la marca
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    
    -- Constraints
    CONSTRAINT marketplace_brands_name_not_empty CHECK (LENGTH(TRIM(name)) > 0),
    CONSTRAINT marketplace_brands_slug_not_empty CHECK (LENGTH(TRIM(slug)) > 0),
    CONSTRAINT marketplace_brands_normalized_name_not_empty CHECK (LENGTH(TRIM(normalized_name)) > 0),
    CONSTRAINT marketplace_brands_quality_score_valid CHECK (quality_score >= 0.0 AND quality_score <= 1.0),
    CONSTRAINT marketplace_brands_verification_status_valid CHECK (verification_status IN ('verified', 'unverified', 'disputed', 'pending')),
    CONSTRAINT marketplace_brands_logo_url_format CHECK (logo_url IS NULL OR logo_url ~ '^https?://'),
    CONSTRAINT marketplace_brands_website_format CHECK (website IS NULL OR website ~ '^https?://')
);

-- Índices para performance
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_slug ON marketplace_brands(slug);
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_normalized_name ON marketplace_brands(normalized_name);
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_verification_status ON marketplace_brands(verification_status);
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_quality_score ON marketplace_brands(quality_score);
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_product_count ON marketplace_brands(product_count);
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_country ON marketplace_brands(country_code);
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_active ON marketplace_brands(is_active);
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_deleted_at ON marketplace_brands(deleted_at);

-- Índice para búsqueda de texto completo
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_search ON marketplace_brands USING gin(to_tsvector('spanish', name || ' ' || COALESCE(description, '')));

-- Índice para aliases array
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_aliases ON marketplace_brands USING gin(aliases);

-- Índice para category_tags array
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_category_tags ON marketplace_brands USING gin(category_tags);

-- Índice compuesto para listados ordenados
CREATE INDEX IF NOT EXISTS idx_marketplace_brands_listing ON marketplace_brands(is_active, quality_score DESC, product_count DESC);

-- Comentarios para documentación
COMMENT ON TABLE marketplace_brands IS 'Catálogo global de marcas normalizadas para consistencia cross-tenant';
COMMENT ON COLUMN marketplace_brands.name IS 'Nombre oficial de la marca';
COMMENT ON COLUMN marketplace_brands.slug IS 'Identificador único para URLs amigables';
COMMENT ON COLUMN marketplace_brands.normalized_name IS 'Nombre normalizado para matching y búsqueda';
COMMENT ON COLUMN marketplace_brands.verification_status IS 'Estado de verificación: verified, unverified, disputed, pending';
COMMENT ON COLUMN marketplace_brands.quality_score IS 'Puntuación de calidad basada en curación (0.0-1.0)';
COMMENT ON COLUMN marketplace_brands.product_count IS 'Contador de productos que usan esta marca';
COMMENT ON COLUMN marketplace_brands.web_data IS 'Datos enriquecidos de APIs web (Clearbit, Brandfetch, etc.)';
COMMENT ON COLUMN marketplace_brands.aliases IS 'Variaciones del nombre de la marca para matching';
COMMENT ON COLUMN marketplace_brands.category_tags IS 'Categorías asociadas a la marca';
COMMENT ON COLUMN marketplace_brands.deleted_at IS 'Timestamp de eliminación lógica, NULL si no está eliminado';

-- Trigger para actualizar updated_at automáticamente
CREATE OR REPLACE FUNCTION update_marketplace_brands_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_marketplace_brands_updated_at
    BEFORE UPDATE ON marketplace_brands
    FOR EACH ROW
    EXECUTE FUNCTION update_marketplace_brands_updated_at();

-- Función para generar slug automáticamente
CREATE OR REPLACE FUNCTION generate_marketplace_brand_slug(brand_name TEXT)
RETURNS TEXT AS $$
BEGIN
    RETURN LOWER(
        REGEXP_REPLACE(
            REGEXP_REPLACE(
                UNACCENT(brand_name), 
                '[^a-zA-Z0-9\s-]', '', 'g'
            ), 
            '\s+', '-', 'g'
        )
    );
END;
$$ language 'plpgsql';

-- Trigger para generar slug automáticamente si no se proporciona
CREATE OR REPLACE FUNCTION auto_generate_marketplace_brand_slug()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.slug IS NULL OR NEW.slug = '' THEN
        NEW.slug = generate_marketplace_brand_slug(NEW.name);
        
        -- Evitar duplicados agregando número secuencial
        WHILE EXISTS (SELECT 1 FROM marketplace_brands WHERE slug = NEW.slug AND id != COALESCE(NEW.id, '00000000-0000-0000-0000-000000000000'::uuid)) LOOP
            NEW.slug = generate_marketplace_brand_slug(NEW.name) || '-' || extract(epoch from now())::integer;
        END LOOP;
    END IF;
    
    -- Generar normalized_name si no se proporciona
    IF NEW.normalized_name IS NULL OR NEW.normalized_name = '' THEN
        NEW.normalized_name = UPPER(UNACCENT(TRIM(NEW.name)));
    END IF;
    
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_auto_generate_marketplace_brand_slug
    BEFORE INSERT OR UPDATE ON marketplace_brands
    FOR EACH ROW
    EXECUTE FUNCTION auto_generate_marketplace_brand_slug(); 