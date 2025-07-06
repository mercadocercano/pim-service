package usecase

import (
	"context"
	"pim/src/shared/domain/criteria"
	"pim/src/attribute/domain/entity"
	"pim/src/attribute/domain/port"
)

// ListMarketplaceAttributesByCriteriaUseCase lista atributos marketplace usando criterios
type ListMarketplaceAttributesByCriteriaUseCase struct {
	marketplaceAttrRepo port.MarketplaceAttributeRepository
}

// NewListMarketplaceAttributesByCriteriaUseCase crea una nueva instancia del caso de uso
func NewListMarketplaceAttributesByCriteriaUseCase(marketplaceAttrRepo port.MarketplaceAttributeRepository) *ListMarketplaceAttributesByCriteriaUseCase {
	return &ListMarketplaceAttributesByCriteriaUseCase{
		marketplaceAttrRepo: marketplaceAttrRepo,
	}
}

// Execute ejecuta el caso de uso
func (uc *ListMarketplaceAttributesByCriteriaUseCase) Execute(ctx context.Context, searchCriteria criteria.Criteria) (*criteria.ListResponse[entity.MarketplaceAttribute], error) {
	// Buscar atributos usando criterios
	attributes, err := uc.marketplaceAttrRepo.SearchByCriteria(ctx, searchCriteria)
	if err != nil {
		return nil, err
	}

	// Contar total de atributos con los mismos filtros (sin paginación)
	countCriteria := criteria.Criteria{
		Filters: searchCriteria.Filters,
		// No incluir Order ni Pagination para el conteo
	}
	total, err := uc.marketplaceAttrRepo.CountByCriteria(ctx, countCriteria)
	if err != nil {
		return nil, err
	}

	// Crear respuesta con información de paginación
	return criteria.NewListResponse(attributes, total, searchCriteria), nil
}