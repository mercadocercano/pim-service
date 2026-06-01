package port

import (
	"context"

	"saas-mt-pim-service/src/attribute/domain/entity"
)

// AttributeValueRepository define los métodos para persistir AttributeValue
type AttributeValueRepository interface {
	Create(ctx context.Context, value *entity.AttributeValue) error
	Update(ctx context.Context, id, newValue string, sortOrder int) (*entity.AttributeValue, error)
	FindByID(ctx context.Context, id string) (*entity.AttributeValue, error)
	FindByAttributeID(ctx context.Context, attributeID string) ([]*entity.AttributeValue, error)
	Delete(ctx context.Context, id string) error
	IsAttributeInUse(ctx context.Context, attributeID string) (bool, error)
}
