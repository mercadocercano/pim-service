package port

import (
	"context"
	cr "github.com/mercadocercano/criteria"
	"saas-mt-pim-service/src/overview/domain/entity"
)

// CategoryStatsRepository define los métodos para persistir CategoryStats
type CategoryStatsRepository interface {
	Create(ctx context.Context, CategoryStats *entity.CategoryStats) error
	Update(ctx context.Context, CategoryStats *entity.CategoryStats) error
	FindByID(ctx context.Context, id string, tenantID string) (*entity.CategoryStats, error)
	FindByTenant(ctx context.Context, tenantID string) ([]*entity.CategoryStats, error)
	Delete(ctx context.Context, id string, tenantID string) error
	SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.CategoryStats, error)
	CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error)
}
