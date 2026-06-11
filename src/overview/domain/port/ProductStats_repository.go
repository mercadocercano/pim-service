package port

import (
	"context"
	cr "github.com/hornosg/go-shared/criteria"
	"saas-mt-pim-service/src/overview/domain/entity"
)

// ProductStatsRepository define los métodos para persistir ProductStats
type ProductStatsRepository interface {
	Create(ctx context.Context, ProductStats *entity.ProductStats) error
	Update(ctx context.Context, ProductStats *entity.ProductStats) error
	FindByID(ctx context.Context, id string, tenantID string) (*entity.ProductStats, error)
	FindByTenant(ctx context.Context, tenantID string) ([]*entity.ProductStats, error)
	Delete(ctx context.Context, id string, tenantID string) error
	SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.ProductStats, error)
	CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error)
}
