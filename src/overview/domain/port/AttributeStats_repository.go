package port

import (
	"context"
	"pim/src/overview/domain/entity"
	"pim/src/shared/domain/criteria"
)

// AttributeStatsRepository define los métodos para persistir AttributeStats
type AttributeStatsRepository interface {
	Create(ctx context.Context, AttributeStats *entity.AttributeStats) error
	Update(ctx context.Context, AttributeStats *entity.AttributeStats) error
	FindByID(ctx context.Context, id string, tenantID string) (*entity.AttributeStats, error)
	FindByTenant(ctx context.Context, tenantID string) ([]*entity.AttributeStats, error)
	Delete(ctx context.Context, id string, tenantID string) error
	SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.AttributeStats, error)
	CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error)
}
