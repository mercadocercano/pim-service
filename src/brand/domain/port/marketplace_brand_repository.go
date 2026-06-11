package port

import (
	"context"
	cr "github.com/hornosg/go-shared/criteria"
	"saas-mt-pim-service/src/brand/domain/entity"
)

// MarketplacebrandRepository define los métodos para persistir Marketplacebrand (tabla global)
type MarketplacebrandRepository interface {
	Create(ctx context.Context, marketplace_brand *entity.Marketplacebrand) error
	Update(ctx context.Context, marketplace_brand *entity.Marketplacebrand) error
	FindByID(ctx context.Context, id string) (*entity.Marketplacebrand, error)
	FindAll(ctx context.Context) ([]*entity.Marketplacebrand, error)
	Delete(ctx context.Context, id string) error
	SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.Marketplacebrand, error)
	CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error)
}
