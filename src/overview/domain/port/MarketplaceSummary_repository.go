package port

import (
	"context"
	"saas-mt-pim-service/src/overview/domain/entity"
	cr "github.com/mercadocercano/criteria"
)

// MarketplaceSummaryRepository define los métodos para persistir MarketplaceSummary
type MarketplaceSummaryRepository interface {
	Create(ctx context.Context, MarketplaceSummary *entity.MarketplaceSummary) error
	Update(ctx context.Context, MarketplaceSummary *entity.MarketplaceSummary) error
	FindByID(ctx context.Context, id string, tenantID string) (*entity.MarketplaceSummary, error)
	FindByTenant(ctx context.Context, tenantID string) ([]*entity.MarketplaceSummary, error)
	Delete(ctx context.Context, id string, tenantID string) error
	SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.MarketplaceSummary, error)
	CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error)
}
