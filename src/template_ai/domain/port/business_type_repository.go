package port

import (
	"context"
	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/domain/entity"
)

// BusinessTypeRepository defines the interface for business type persistence
type BusinessTypeRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.BusinessType, error)
	GetByCode(ctx context.Context, code string) (*entity.BusinessType, error)
	GetAll(ctx context.Context) ([]*entity.BusinessType, error)
}