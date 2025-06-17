-- Migration: 014_create_business_types_tables.sql
-- Purpose: Create tables for business types and their quickstart templates
-- Date: 2025-06-13

-- Create business_types table
CREATE TABLE IF NOT EXISTS business_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    icon VARCHAR(255),
    color VARCHAR(50),
    is_active BOOLEAN DEFAULT true,
    sort_order INTEGER DEFAULT 0,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for business_types
CREATE INDEX idx_business_types_code ON business_types(code);
CREATE INDEX idx_business_types_is_active ON business_types(is_active);
CREATE INDEX idx_business_types_sort_order ON business_types(sort_order);

-- Create business_type_templates table
CREATE TABLE IF NOT EXISTS business_type_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_type_id UUID NOT NULL REFERENCES business_types(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    version VARCHAR(20) DEFAULT '1.0.0',
    region VARCHAR(10) DEFAULT 'GLOBAL',
    categories JSONB DEFAULT '[]',
    attributes JSONB DEFAULT '[]',
    products JSONB DEFAULT '[]',
    brands JSONB DEFAULT '[]',
    is_active BOOLEAN DEFAULT true,
    is_default BOOLEAN DEFAULT false,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for business_type_templates
CREATE INDEX idx_business_type_templates_business_type_id ON business_type_templates(business_type_id);
CREATE INDEX idx_business_type_templates_region ON business_type_templates(region);
CREATE INDEX idx_business_type_templates_is_active ON business_type_templates(is_active);
CREATE INDEX idx_business_type_templates_is_default ON business_type_templates(is_default);

-- Create partial unique index instead of constraint
CREATE UNIQUE INDEX idx_unique_default_per_business_type ON business_type_templates(business_type_id, region) WHERE is_default = true;

-- Create tenant_business_type_setup table (tracks which business type a tenant selected during onboarding)
CREATE TABLE IF NOT EXISTS tenant_business_type_setup (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    business_type_id UUID NOT NULL REFERENCES business_types(id),
    template_id UUID REFERENCES business_type_templates(id),
    setup_completed BOOLEAN DEFAULT false,
    setup_data JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_tenant_business_type UNIQUE (tenant_id)
);

-- Create indexes for tenant_business_type_setup
CREATE INDEX idx_tenant_business_type_setup_tenant_id ON tenant_business_type_setup(tenant_id);
CREATE INDEX idx_tenant_business_type_setup_business_type_id ON tenant_business_type_setup(business_type_id);

-- Trigger to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Apply trigger to business_types
CREATE TRIGGER update_business_types_updated_at 
    BEFORE UPDATE ON business_types 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Apply trigger to business_type_templates
CREATE TRIGGER update_business_type_templates_updated_at 
    BEFORE UPDATE ON business_type_templates 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Apply trigger to tenant_business_type_setup
CREATE TRIGGER update_tenant_business_type_setup_updated_at 
    BEFORE UPDATE ON tenant_business_type_setup 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Insert initial business types
INSERT INTO business_types (code, name, description, icon, color, sort_order) VALUES
('retail', 'Tienda Minorista', 'Venta de productos al por menor', 'shopping-bag', '#4F46E5', 1),
('restaurant', 'Restaurante', 'Servicio de comida y bebida', 'utensils', '#EF4444', 2),
('services', 'Servicios', 'Prestación de servicios profesionales', 'briefcase', '#10B981', 3),
('wholesale', 'Mayorista', 'Venta al por mayor', 'truck', '#F59E0B', 4),
('manufacturing', 'Manufactura', 'Producción y fabricación', 'factory', '#6366F1', 5),
('healthcare', 'Salud', 'Servicios médicos y de salud', 'heart', '#EC4899', 6),
('education', 'Educación', 'Servicios educativos', 'graduation-cap', '#8B5CF6', 7),
('technology', 'Tecnología', 'Productos y servicios tecnológicos', 'laptop', '#3B82F6', 8);

-- Comments
COMMENT ON TABLE business_types IS 'Tipos de negocio disponibles en el marketplace';
COMMENT ON TABLE business_type_templates IS 'Plantillas de configuración quickstart para cada tipo de negocio';
COMMENT ON TABLE tenant_business_type_setup IS 'Registro de qué tipo de negocio seleccionó cada tenant'; 