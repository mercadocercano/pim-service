package port

import (
	"context"
	"saas-mt-pim-service/src/businesstype/domain/entity"
	"saas-mt-pim-service/src/shared/domain/criteria"
)

// BusinessTypeTemplateRepository define los métodos para persistir BusinessTypeTemplate
type BusinessTypeTemplateRepository interface {
	Create(ctx context.Context, template *entity.BusinessTypeTemplate) error
	Update(ctx context.Context, template *entity.BusinessTypeTemplate) error
	FindByID(ctx context.Context, id string) (*entity.BusinessTypeTemplate, error)
	FindByBusinessTypeID(ctx context.Context, businessTypeID string) ([]*entity.BusinessTypeTemplate, error)
	FindByBusinessTypeAndRegion(ctx context.Context, businessTypeID, region string) ([]*entity.BusinessTypeTemplate, error)
	FindDefault(ctx context.Context, businessTypeID, region string) (*entity.BusinessTypeTemplate, error)
	Delete(ctx context.Context, id string) error
	SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.BusinessTypeTemplate, error)
	CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error)
}
