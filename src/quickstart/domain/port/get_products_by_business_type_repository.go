package port

import "context"

// GetProductsByBusinessTypeRepository define el contrato para obtener productos sugeridos de un template
type GetProductsByBusinessTypeRepository interface {
	GetProductsByBusinessType(ctx context.Context, businessTypeSlug string) ([]TemplateProduct, error)
}
