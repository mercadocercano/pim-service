-- Migration: Add AI capabilities to template system
-- Description: Extends existing template system with AI generation and optimization features

-- 1. Extend business_type_templates with AI fields
ALTER TABLE business_type_templates
ADD COLUMN IF NOT EXISTS generated_by VARCHAR(50) DEFAULT 'manual', -- 'ai', 'manual', 'hybrid'
ADD COLUMN IF NOT EXISTS ai_generation_params JSONB,
ADD COLUMN IF NOT EXISTS performance_metrics JSONB,
ADD COLUMN IF NOT EXISTS last_ai_update TIMESTAMP,
ADD COLUMN IF NOT EXISTS product_selection_rules JSONB,
ADD COLUMN IF NOT EXISTS category_distribution JSONB,
ADD COLUMN IF NOT EXISTS brand_preferences JSONB,
ADD COLUMN IF NOT EXISTS regional_adaptations JSONB;

-- Add index for AI-generated templates
CREATE INDEX IF NOT EXISTS idx_business_type_templates_generated_by 
ON business_type_templates(generated_by);

-- 2. Create template_global_products to link templates with global catalog
CREATE TABLE IF NOT EXISTS template_global_products (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    template_id UUID NOT NULL REFERENCES business_type_templates(id) ON DELETE CASCADE,
    global_product_id UUID NOT NULL,
    priority INTEGER DEFAULT 2 CHECK (priority BETWEEN 1 AND 3), -- 1=essential, 2=recommended, 3=optional
    quantity_suggestion INTEGER DEFAULT 1,
    ai_reasoning TEXT,
    relevance_score DECIMAL(3,2) CHECK (relevance_score >= 0 AND relevance_score <= 1),
    category_match_score DECIMAL(3,2),
    brand_match_score DECIMAL(3,2),
    regional_preference_score DECIMAL(3,2),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    
    UNIQUE(template_id, global_product_id)
);

-- Indexes for template_global_products
CREATE INDEX idx_template_global_products_template_id ON template_global_products(template_id);
CREATE INDEX idx_template_global_products_priority ON template_global_products(priority);
CREATE INDEX idx_template_global_products_relevance ON template_global_products(relevance_score DESC);

-- 3. Create ai_template_generation_history for tracking AI generations
CREATE TABLE IF NOT EXISTS ai_template_generation_history (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    template_id UUID REFERENCES business_type_templates(id) ON DELETE SET NULL,
    tenant_id UUID,
    business_type_id UUID NOT NULL REFERENCES business_types(id),
    generation_params JSONB NOT NULL,
    ai_model VARCHAR(100),
    prompt_template TEXT,
    generated_content JSONB NOT NULL,
    generation_status VARCHAR(20) NOT NULL DEFAULT 'pending', -- pending, processing, completed, failed
    error_message TEXT,
    generation_time_ms INTEGER,
    applied_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Index for generation history
CREATE INDEX idx_ai_generation_history_tenant ON ai_template_generation_history(tenant_id);
CREATE INDEX idx_ai_generation_history_status ON ai_template_generation_history(generation_status);
CREATE INDEX idx_ai_generation_history_created ON ai_template_generation_history(created_at DESC);

-- 4. Create ai_template_performance_metrics for tracking template effectiveness
CREATE TABLE IF NOT EXISTS ai_template_performance_metrics (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    template_id UUID NOT NULL REFERENCES business_type_templates(id) ON DELETE CASCADE,
    metric_type VARCHAR(50) NOT NULL, -- usage_count, satisfaction_score, modification_rate, etc.
    metric_value DECIMAL(10,2) NOT NULL,
    metric_metadata JSONB,
    period_start TIMESTAMP NOT NULL,
    period_end TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    
    UNIQUE(template_id, metric_type, period_start, period_end)
);

-- Index for performance metrics
CREATE INDEX idx_ai_performance_metrics_template ON ai_template_performance_metrics(template_id);
CREATE INDEX idx_ai_performance_metrics_type ON ai_template_performance_metrics(metric_type);
CREATE INDEX idx_ai_performance_metrics_period ON ai_template_performance_metrics(period_start, period_end);

-- 5. Create ai_product_feedback for learning from user modifications
CREATE TABLE IF NOT EXISTS ai_product_feedback (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    tenant_id UUID NOT NULL,
    template_id UUID REFERENCES business_type_templates(id) ON DELETE SET NULL,
    global_product_id UUID,
    action VARCHAR(20) NOT NULL, -- 'kept', 'removed', 'quantity_changed', 'replaced'
    original_quantity INTEGER,
    new_quantity INTEGER,
    replacement_product_id UUID,
    feedback_reason TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Index for product feedback
CREATE INDEX idx_ai_product_feedback_tenant ON ai_product_feedback(tenant_id);
CREATE INDEX idx_ai_product_feedback_template ON ai_product_feedback(template_id);
CREATE INDEX idx_ai_product_feedback_action ON ai_product_feedback(action);

-- 6. Add trigger to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_template_global_products_updated_at 
BEFORE UPDATE ON template_global_products 
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- 7. Add comments for documentation
COMMENT ON TABLE template_global_products IS 'Links business type templates with global catalog products for AI-powered suggestions';
COMMENT ON COLUMN template_global_products.priority IS '1=essential products, 2=recommended products, 3=optional/opportunity products';
COMMENT ON COLUMN template_global_products.ai_reasoning IS 'AI-generated explanation for why this product was selected';

COMMENT ON TABLE ai_template_generation_history IS 'Tracks all AI-powered template generation attempts and results';
COMMENT ON COLUMN ai_template_generation_history.generation_params IS 'Input parameters used for AI generation (preferences, constraints, etc.)';

COMMENT ON TABLE ai_template_performance_metrics IS 'Tracks performance metrics for AI-generated templates to improve future generations';
COMMENT ON TABLE ai_product_feedback IS 'Captures user modifications to AI suggestions for continuous learning';

-- 8. Create default AI generation parameters for each business type
UPDATE business_type_templates
SET ai_generation_params = jsonb_build_object(
    'product_mix', jsonb_build_object(
        'essential', 0.6,
        'recommended', 0.3,
        'optional', 0.1
    ),
    'brand_distribution', jsonb_build_object(
        'premium', 0.35,
        'standard', 0.40,
        'generic', 0.25
    ),
    'category_priorities', CASE 
        WHEN business_type_id IN (SELECT id FROM business_types WHERE code = 'almacen_kiosco')
        THEN '["bebidas", "snacks", "golosinas", "cigarrillos", "limpieza"]'::jsonb
        WHEN business_type_id IN (SELECT id FROM business_types WHERE code = 'supermercado')
        THEN '["alimentos", "bebidas", "limpieza", "perfumeria", "frescos"]'::jsonb
        WHEN business_type_id IN (SELECT id FROM business_types WHERE code = 'farmacia')
        THEN '["medicamentos", "perfumeria", "cuidado_personal", "bebes", "suplementos"]'::jsonb
        ELSE '[]'::jsonb
    END,
    'regional_adaptations_enabled', true,
    'min_products', 50,
    'max_products', 150
)
WHERE ai_generation_params IS NULL;