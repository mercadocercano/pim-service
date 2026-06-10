package usecase

import (
	"context"

	cr "github.com/mercadocercano/criteria"
	"saas-mt-pim-service/src/brand/application/mapper"
	"saas-mt-pim-service/src/brand/application/response"
	"saas-mt-pim-service/src/brand/domain/port"
)

// ListBrandsByCriteriaUseCase maneja la obtención de marcas con filtros y paginación
type ListBrandsByCriteriaUseCase struct {
	brandRepo   port.BrandCriteriaRepository
	brandMapper *mapper.BrandMapper
}

// NewListBrandsByCriteriaUseCase crea una nueva instancia del caso de uso
func NewListBrandsByCriteriaUseCase(
	brandRepo port.BrandCriteriaRepository,
	brandMapper *mapper.BrandMapper,
) *ListBrandsByCriteriaUseCase {
	return &ListBrandsByCriteriaUseCase{
		brandRepo:   brandRepo,
		brandMapper: brandMapper,
	}
}

// Execute ejecuta el caso de uso de listado de marcas con criterios
func (uc *ListBrandsByCriteriaUseCase) Execute(ctx context.Context, searchCriteria cr.Criteria) (*response.BrandListResponse, error) {
	// Buscar marcas según los criterios
	brands, err := uc.brandRepo.SearchByCriteria(ctx, searchCriteria)
	if err != nil {
		return nil, err
	}

	// Obtener el total de marcas que coinciden con los filtros
	total, err := uc.brandRepo.CountByCriteria(ctx, searchCriteria)
	if err != nil {
		return nil, err
	}

	// Convertir a response y retornar
	return uc.brandMapper.ToResponseList(brands, total, searchCriteria), nil
}
