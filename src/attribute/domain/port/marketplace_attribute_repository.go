package port

import (
	"context"
	"saas-mt-pim-service/src/attribute/domain/entity"
	cr "github.com/mercadocercano/criteria"
)

// MarketplaceAttributeRepository define los métodos para persistir MarketplaceAttribute
type MarketplaceAttributeRepository interface {
	Create(ctx context.Context, attribute *entity.MarketplaceAttribute) error
	Update(ctx context.Context, attribute *entity.MarketplaceAttribute) error
	FindByID(ctx context.Context, id string) (*entity.MarketplaceAttribute, error)
	FindAll(ctx context.Context) ([]*entity.MarketplaceAttribute, error)
	FindByName(ctx context.Context, name string) (*entity.MarketplaceAttribute, error)
	FindBySlug(ctx context.Context, slug string) (*entity.MarketplaceAttribute, error)
	Delete(ctx context.Context, id string) error
	SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.MarketplaceAttribute, error)
	CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error)
}
