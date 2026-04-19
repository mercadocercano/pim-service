-- Solicitudes de productos no encontrados en el catálogo global.
-- Los tenants pueden pedir que se agregue un producto; el admin lo revisa y aprueba.

CREATE TABLE IF NOT EXISTS global_product_requests (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id       VARCHAR(255) NOT NULL,
    name            VARCHAR(500) NOT NULL,
    brand           VARCHAR(255),
    category        VARCHAR(255),
    description     TEXT,
    business_type   VARCHAR(100),
    status          VARCHAR(20) NOT NULL DEFAULT 'pending',
    admin_notes     TEXT,
    global_product_id UUID,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_request_status CHECK (status IN ('pending', 'approved', 'rejected', 'fulfilled'))
);

CREATE INDEX IF NOT EXISTS idx_gpr_status ON global_product_requests (status);
CREATE INDEX IF NOT EXISTS idx_gpr_tenant ON global_product_requests (tenant_id);
CREATE INDEX IF NOT EXISTS idx_gpr_created ON global_product_requests (created_at DESC);
