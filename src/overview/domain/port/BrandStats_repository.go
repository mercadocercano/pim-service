package port

import (
	"context"
	"pim/src/overview/domain/entity"
	"pim/src/shared/domain/criteria"
)

// BrandStatsRepository define los métodos para persistir BrandStats
type BrandStatsRepository interface {
	Create(ctx context.Context, BrandStats *entity.BrandStats) error
	Update(ctx context.Context, BrandStats *entity.BrandStats) error
	FindByID(ctx context.Context, id string, tenantID string) (*entity.BrandStats, error)
	FindByTenant(ctx context.Context, tenantID string) ([]*entity.BrandStats, error)
	Delete(ctx context.Context, id string, tenantID string) error
	SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.BrandStats, error)
	CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error)
}
