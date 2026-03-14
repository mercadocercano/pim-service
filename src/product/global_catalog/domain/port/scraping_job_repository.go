package port

import (
	"context"
	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
	cr "github.com/mercadocercano/criteria"
)

// ScrapingjobRepository define los métodos para persistir Scrapingjob
type ScrapingjobRepository interface {
	Create(ctx context.Context, scraping_job *entity.Scrapingjob) error
	Update(ctx context.Context, scraping_job *entity.Scrapingjob) error
	FindByID(ctx context.Context, id string, tenantID string) (*entity.Scrapingjob, error)
	FindByTenant(ctx context.Context, tenantID string) ([]*entity.Scrapingjob, error)
	Delete(ctx context.Context, id string, tenantID string) error
	SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.Scrapingjob, error)
	CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error)
}
