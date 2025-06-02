package port

import (
	"context"
)

// CategoryService define las operaciones para crear categorías desde el quickstart
type CategoryService interface {
	CreateFromTemplate(ctx context.Context, tenantID string, templateData interface{}) error
}

// AttributeService define las operaciones para crear atributos desde el quickstart
type AttributeService interface {
	CreateFromTemplate(ctx context.Context, tenantID string, templateData interface{}) error
}

// CategoryAttributeService define las operaciones para crear relaciones categoría-atributo desde el quickstart
type CategoryAttributeService interface {
	CreateFromTemplate(ctx context.Context, tenantID string, templateData interface{}) error
}

// VariantService define las operaciones para crear variantes desde el quickstart
type VariantService interface {
	CreateFromTemplate(ctx context.Context, tenantID string, templateData interface{}) error
}

// ProductService define las operaciones para crear productos desde el quickstart
type ProductService interface {
	CreateFromTemplate(ctx context.Context, tenantID string, templateData interface{}) error
}
