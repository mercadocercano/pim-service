package usecase

import (
	"context"

	cr "github.com/hornosg/go-shared/criteria"
	"saas-mt-pim-service/src/product/tenant/application/mapper"
	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/domain/port"
)

// ListProductsByCriteriaUseCase maneja la búsqueda de productos con criterios
type ListProductsByCriteriaUseCase struct {
	productRepo port.ProductCriteriaRepository
	mapper      *mapper.ProductMapper
}

// NewListProductsByCriteriaUseCase crea una nueva instancia del caso de uso
func NewListProductsByCriteriaUseCase(
	productRepo port.ProductCriteriaRepository,
	productMapper *mapper.ProductMapper,
) *ListProductsByCriteriaUseCase {
	return &ListProductsByCriteriaUseCase{
		productRepo: productRepo,
		mapper:      productMapper,
	}
}

// Execute ejecuta el caso de uso de búsqueda de productos por criterios
func (uc *ListProductsByCriteriaUseCase) Execute(
	ctx context.Context,
	searchCriteria cr.Criteria,
) (*response.ProductListResponse, error) {
	// Buscar productos
	products, err := uc.productRepo.SearchByCriteria(ctx, searchCriteria)
	if err != nil {
		return nil, err
	}

	// Contar total de productos
	total, err := uc.productRepo.CountByCriteria(ctx, searchCriteria)
	if err != nil {
		return nil, err
	}

	// Convertir a respuesta con paginación
	return uc.mapper.ToListResponse(
		products,
		searchCriteria.Pagination.Page,
		searchCriteria.Pagination.PageSize,
		total,
	), nil
}
