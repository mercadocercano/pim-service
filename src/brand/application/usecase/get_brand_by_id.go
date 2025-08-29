package usecase

import (
	"context"

	"saas-mt-pim-service/src/brand/application/mapper"
	"saas-mt-pim-service/src/brand/application/response"
	"saas-mt-pim-service/src/brand/domain/exception"
	"saas-mt-pim-service/src/brand/domain/port"
)

// GetBrandByIDUseCase maneja la obtención de una marca por ID
type GetBrandByIDUseCase struct {
	brandRepo   port.BrandRepository
	brandMapper *mapper.BrandMapper
}

// NewGetBrandByIDUseCase crea una nueva instancia del caso de uso
func NewGetBrandByIDUseCase(
	brandRepo port.BrandRepository,
	brandMapper *mapper.BrandMapper,
) *GetBrandByIDUseCase {
	return &GetBrandByIDUseCase{
		brandRepo:   brandRepo,
		brandMapper: brandMapper,
	}
}

// Execute ejecuta el caso de uso de obtención de marca por ID
func (uc *GetBrandByIDUseCase) Execute(ctx context.Context, brandID, tenantID string) (*response.BrandResponse, error) {
	// Buscar la marca
	brand, err := uc.brandRepo.FindByID(ctx, brandID, tenantID)
	if err != nil {
		return nil, err
	}

	if brand == nil {
		return nil, exception.ErrBrandNotFound
	}

	// Convertir a response y retornar
	return uc.brandMapper.ToResponse(brand), nil
}
