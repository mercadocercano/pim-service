package port

import (
	"context"
	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
)

type ProductRequestRepository interface {
	Save(ctx context.Context, request *entity.ProductRequest) error
	FindByID(ctx context.Context, id string) (*entity.ProductRequest, error)
	Update(ctx context.Context, request *entity.ProductRequest) error
	FindPending(ctx context.Context, limit, offset int) ([]*entity.ProductRequest, error)
	CountPending(ctx context.Context) (int, error)
}
