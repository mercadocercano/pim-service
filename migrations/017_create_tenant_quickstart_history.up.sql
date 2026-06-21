-- Migración: 017_create_tenant_quickstart_history.sql
-- Descripción: Crear tabla para almacenar el historial y progreso del wizard quickstart por tenant

-- Crear tabla tenant_quickstart_history
CREATE TABLE IF NOT EXISTS tenant_quickstart_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    business_type_id UUID NOT NULL,
    template_id UUID,
    setup_completed BOOLEAN DEFAULT FALSE,
    setup_data JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign keys
    CONSTRAINT fk_business_type FOREIGN KEY (business_type_id) 
        REFERENCES business_types(id) ON DELETE RESTRICT,
    CONSTRAINT fk_template FOREIGN KEY (template_id) 
        REFERENCES business_type_templates(id) ON DELETE SET NULL,
    
    -- Índices
    CONSTRAINT unique_tenant_active_wizard UNIQUE (tenant_id, setup_completed)
);

-- Crear índices para mejorar rendimiento
CREATE INDEX idx_tenant_quickstart_history_tenant_id ON tenant_quickstart_history(tenant_id);
CREATE INDEX idx_tenant_quickstart_history_business_type_id ON tenant_quickstart_history(business_type_id);
CREATE INDEX idx_tenant_quickstart_history_setup_completed ON tenant_quickstart_history(setup_completed);
CREATE INDEX idx_tenant_quickstart_history_created_at ON tenant_quickstart_history(created_at);

-- Crear función para actualizar updated_at automáticamente
CREATE OR REPLACE FUNCTION update_tenant_quickstart_history_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Crear trigger para actualizar updated_at
CREATE TRIGGER update_tenant_quickstart_history_updated_at_trigger
    BEFORE UPDATE ON tenant_quickstart_history
    FOR EACH ROW
    EXECUTE FUNCTION update_tenant_quickstart_history_updated_at();

-- Comentarios en la tabla y columnas
COMMENT ON TABLE tenant_quickstart_history IS 'Historial y progreso del wizard quickstart por tenant';
COMMENT ON COLUMN tenant_quickstart_history.id IS 'Identificador único del registro';
COMMENT ON COLUMN tenant_quickstart_history.tenant_id IS 'ID del tenant que ejecuta el wizard';
COMMENT ON COLUMN tenant_quickstart_history.business_type_id IS 'Tipo de negocio seleccionado';
COMMENT ON COLUMN tenant_quickstart_history.template_id IS 'Template utilizado (opcional)';
COMMENT ON COLUMN tenant_quickstart_history.setup_completed IS 'Indica si el wizard fue completado';
COMMENT ON COLUMN tenant_quickstart_history.setup_data IS 'Datos del progreso y selecciones del wizard en formato JSON';
COMMENT ON COLUMN tenant_quickstart_history.created_at IS 'Fecha de creación del registro';
COMMENT ON COLUMN tenant_quickstart_history.updated_at IS 'Fecha de última actualización';