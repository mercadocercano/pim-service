package port

import (
	"context"
	"saas-mt-pim-service/src/businesstype/domain/entity"
	cr "github.com/mercadocercano/criteria"
)

// BusinessTypeRepository define los métodos para persistir BusinessType
type BusinessTypeRepository interface {
	Create(ctx context.Context, businessType *entity.BusinessType) error
	Update(ctx context.Context, businessType *entity.BusinessType) error
	FindByID(ctx context.Context, id string) (*entity.BusinessType, error)
	FindByCode(ctx context.Context, code string) (*entity.BusinessType, error)
	FindAll(ctx context.Context) ([]*entity.BusinessType, error)
	FindActive(ctx context.Context) ([]*entity.BusinessType, error)
	Delete(ctx context.Context, id string) error
	SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.BusinessType, error)
	CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error)
}
