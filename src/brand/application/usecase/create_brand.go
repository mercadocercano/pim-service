package usecase

import (
	"context"

	"saas-mt-pim-service/src/brand/application/mapper"
	"saas-mt-pim-service/src/brand/application/request"
	"saas-mt-pim-service/src/brand/application/response"
	"saas-mt-pim-service/src/brand/domain/port"
	"saas-mt-pim-service/src/brand/domain/service"
)

// CreateBrandUseCase maneja la creación de marcas
type CreateBrandUseCase struct {
	brandRepo    port.BrandRepository
	brandService *service.BrandDomainService
	brandMapper  *mapper.BrandMapper
}

// NewCreateBrandUseCase crea una nueva instancia del caso de uso
func NewCreateBrandUseCase(
	brandRepo port.BrandRepository,
	brandService *service.BrandDomainService,
	brandMapper *mapper.BrandMapper,
) *CreateBrandUseCase {
	return &CreateBrandUseCase{
		brandRepo:    brandRepo,
		brandService: brandService,
		brandMapper:  brandMapper,
	}
}

// Execute ejecuta el caso de uso de creación de marca
func (uc *CreateBrandUseCase) Execute(ctx context.Context, req *request.CreateBrandRequest, tenantID string) (*response.BrandResponse, error) {
	// Validar el request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Normalizar y validar los datos
	normalizedName, normalizedDescription, normalizedLogoURL, normalizedWebsite, err := uc.brandService.NormalizeBrandData(
		req.Name, req.Description, req.LogoURL, req.Website,
	)
	if err != nil {
		return nil, err
	}

	// Verificar que el nombre sea único
	if err := uc.brandService.ValidateUniqueName(ctx, normalizedName, tenantID, nil); err != nil {
		return nil, err
	}

	// Crear la entidad Brand
	brand, err := uc.brandMapper.FromCreateRequest(&request.CreateBrandRequest{
		Name:        normalizedName,
		Description: normalizedDescription,
		LogoURL:     normalizedLogoURL,
		Website:     normalizedWebsite,
	}, tenantID)
	if err != nil {
		return nil, err
	}

	// Guardar en el repositorio
	if err := uc.brandRepo.Create(ctx, brand); err != nil {
		return nil, err
	}

	// Convertir a response y retornar
	return uc.brandMapper.ToResponse(brand), nil
}
