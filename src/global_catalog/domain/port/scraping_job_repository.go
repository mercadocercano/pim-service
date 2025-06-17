package port

import (
	"context"
	"pim/src/global_catalog/domain/entity"
	"pim/src/shared/domain/criteria"
)

// ScrapingjobRepository define los métodos para persistir Scrapingjob
type ScrapingjobRepository interface {
	Create(ctx context.Context, scraping_job *entity.Scrapingjob) error
	Update(ctx context.Context, scraping_job *entity.Scrapingjob) error
	FindByID(ctx context.Context, id string, tenantID string) (*entity.Scrapingjob, error)
	FindByTenant(ctx context.Context, tenantID string) ([]*entity.Scrapingjob, error)
	Delete(ctx context.Context, id string, tenantID string) error
	SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.Scrapingjob, error)
	CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error)
}
