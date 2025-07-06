package usecase

import (
	"context"
	"pim/src/shared/domain/criteria"
	"pim/src/product/global_catalog/domain/entity"
	"pim/src/product/global_catalog/domain/port"
)

// ListGlobalProductsByCriteriaUseCase lista productos globales usando criterios
type ListGlobalProductsByCriteriaUseCase struct {
	globalProductRepo port.GlobalProductRepository
}

// NewListGlobalProductsByCriteriaUseCase crea una nueva instancia del caso de uso
func NewListGlobalProductsByCriteriaUseCase(globalProductRepo port.GlobalProductRepository) *ListGlobalProductsByCriteriaUseCase {
	return &ListGlobalProductsByCriteriaUseCase{
		globalProductRepo: globalProductRepo,
	}
}

// Execute ejecuta el caso de uso
func (uc *ListGlobalProductsByCriteriaUseCase) Execute(ctx context.Context, searchCriteria criteria.Criteria) (*criteria.ListResponse[entity.GlobalProduct], error) {
	// Buscar productos usando criterios
	products, err := uc.globalProductRepo.SearchByCriteria(ctx, searchCriteria)
	if err != nil {
		return nil, err
	}

	// Contar total de productos con los mismos filtros (sin paginación)
	countCriteria := criteria.Criteria{
		Filters: searchCriteria.Filters,
		// No incluir Order ni Pagination para el conteo
	}
	total, err := uc.globalProductRepo.CountByCriteria(ctx, countCriteria)
	if err != nil {
		return nil, err
	}

	// Crear respuesta con información de paginación
	return criteria.NewListResponse(products, total, searchCriteria), nil
}