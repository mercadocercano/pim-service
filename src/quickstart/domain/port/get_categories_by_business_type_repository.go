package port

import "context"

// CategoryByBusinessType representa una categoría asociada a un tipo de negocio
type CategoryByBusinessType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// GetCategoriesByBusinessTypeRepository define las operaciones para obtener categorías por tipo de negocio
type GetCategoriesByBusinessTypeRepository interface {
	GetCategoriesByBusinessType(ctx context.Context, businessTypeSlug string) ([]CategoryByBusinessType, error)
}
