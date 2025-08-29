package usecase

import (
	"context"

	"saas-mt-pim-service/src/category/domain/entity"
	"saas-mt-pim-service/src/category/domain/port"
	"saas-mt-pim-service/src/shared/domain/criteria"
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
