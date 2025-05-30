package usecase

import (
	"context"

	"pim/src/category_attribute/domain/entity"
	"pim/src/category_attribute/domain/port"
	"pim/src/shared/domain/criteria"
)

// ListCategoryAttributesByCriteriaUseCase maneja el listado de atributos de categoría con filtros y paginación
type ListCategoryAttributesByCriteriaUseCase struct {
	repository port.CategoryAttributeCriteriaRepository
}

// NewListCategoryAttributesByCriteriaUseCase crea una nueva instancia del caso de uso
func NewListCategoryAttributesByCriteriaUseCase(repository port.CategoryAttributeCriteriaRepository) *ListCategoryAttributesByCriteriaUseCase {
	return &ListCategoryAttributesByCriteriaUseCase{
		repository: repository,
	}
}

// Execute ejecuta el caso de uso para listar atributos de categoría con criteria
func (uc *ListCategoryAttributesByCriteriaUseCase) Execute(ctx context.Context, crit criteria.Criteria) (*criteria.ListResponse[entity.CategoryAttribute], error) {
	return uc.repository.ListByCriteria(ctx, crit)
}
