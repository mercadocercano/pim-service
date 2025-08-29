-- Crear tabla para trabajos de importación asíncronos
CREATE TABLE IF NOT EXISTS import_jobs (
    id UUID PRIMARY KEY,
    tenant_id VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_size_bytes BIGINT NOT NULL DEFAULT 0,
    total_records INTEGER NOT NULL DEFAULT 0,
    processed_records INTEGER NOT NULL DEFAULT 0,
    success_count INTEGER NOT NULL DEFAULT 0,
    failure_count INTEGER NOT NULL DEFAULT 0,
    progress DECIMAL(5,2) NOT NULL DEFAULT 0.0,
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255) NOT NULL,
    error_message TEXT,
    result_file_url TEXT,
    notification_sent BOOLEAN NOT NULL DEFAULT FALSE,
    webhook_url TEXT,
    email_notify VARCHAR(255),
    CONSTRAINT check_progress CHECK (progress >= 0 AND progress <= 100)
);

-- Índices para consultas eficientes
CREATE INDEX idx_import_jobs_tenant_id ON import_jobs(tenant_id);
CREATE INDEX idx_import_jobs_status ON import_jobs(status);
CREATE INDEX idx_import_jobs_created_at ON import_jobs(created_at);
CREATE INDEX idx_import_jobs_notification ON import_jobs(notification_sent, status) 
    WHERE notification_sent = FALSE AND status IN ('completed', 'failed', 'cancelled');
CREATE INDEX idx_import_jobs_pending ON import_jobs(status, created_at) 
    WHERE status = 'pending';

-- Trigger para actualizar updated_at
CREATE OR REPLACE FUNCTION update_import_jobs_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_import_jobs_updated_at_trigger
    BEFORE UPDATE ON import_jobs
    FOR EACH ROW
    EXECUTE FUNCTION update_import_jobs_updated_at();

-- Comentarios
COMMENT ON TABLE import_jobs IS 'Trabajos de importación asíncronos con notificaciones';
COMMENT ON COLUMN import_jobs.type IS 'Tipo de importación: csv_products, json_products, batch_create';
COMMENT ON COLUMN import_jobs.status IS 'Estado del trabajo: pending, processing, completed, failed, cancelled';
COMMENT ON COLUMN import_jobs.progress IS 'Progreso del trabajo de 0 a 100';
COMMENT ON COLUMN import_jobs.webhook_url IS 'URL para notificación webhook cuando termine';
COMMENT ON COLUMN import_jobs.email_notify IS 'Email para notificación cuando termine';