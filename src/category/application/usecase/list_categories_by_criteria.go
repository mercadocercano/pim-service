package usecase

import (
	"context"

	"pim/src/category/domain/entity"
	"pim/src/category/domain/port"
	"pim/src/shared/domain/criteria"
)

// ListCategoriesByCriteriaUseCase maneja el listado de categorías con filtros y paginación
type ListCategoriesByCriteriaUseCase struct {
	repository port.CategoryCriteriaRepository
}

// NewListCategoriesByCriteriaUseCase crea una nueva instancia del caso de uso
func NewListCategoriesByCriteriaUseCase(repository port.CategoryCriteriaRepository) *ListCategoriesByCriteriaUseCase {
	return &ListCategoriesByCriteriaUseCase{
		repository: repository,
	}
}

// Execute ejecuta el caso de uso para listar categorías con criteria
func (uc *ListCategoriesByCriteriaUseCase) Execute(ctx context.Context, crit criteria.Criteria) (*criteria.ListResponse[entity.Category], error) {
	return uc.repository.ListByCriteria(ctx, crit)
}
