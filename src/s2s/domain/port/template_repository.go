package port

import (
	"context"
	"time"
)

// TemplateStatusRow contiene los datos de estado del template para un rubro.
type TemplateStatusRow struct {
	ComputedCount  int
	EditorialCount int
	LastRefresh    *time.Time
}

// TemplateRepository define las consultas de persistencia para el módulo s2s.
type TemplateRepository interface {
	// RefreshProductTemplates recalcula business_type_product_templates desde
	// global_products verificados. Es idempotente — seguro de llamar repetidamente.
	RefreshProductTemplates(ctx context.Context) (rowsAffected int64, err error)

	// GetTemplateStatus retorna el estado del template para el rubro con el slug dado.
	// Retorna (nil, nil) si el slug no existe.
	GetTemplateStatus(ctx context.Context, businessTypeSlug string) (*TemplateStatusRow, error)
}
