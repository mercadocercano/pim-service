package port

import (
	"context"
	"saas-mt-pim-service/src/attribute/domain/entity"
	"saas-mt-pim-service/src/shared/domain/criteria"
)

// AttributeRepository define los métodos para persistir Attribute
type AttributeRepository interface {
	Create(ctx context.Context, attribute *entity.Attribute) error
	Update(ctx context.Context, attribute *entity.Attribute) error
	FindByID(ctx context.Context, id string, tenantID string) (*entity.Attribute, error)
	FindByTenant(ctx context.Context, tenantID string) ([]*entity.Attribute, error)
	Delete(ctx context.Context, id string, tenantID string) error
	SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.Attribute, error)
	CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error)
}
