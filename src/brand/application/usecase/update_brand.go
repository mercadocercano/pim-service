package usecase

import (
	"context"

	"saas-mt-pim-service/src/brand/application/mapper"
	"saas-mt-pim-service/src/brand/application/request"
	"saas-mt-pim-service/src/brand/application/response"
	"saas-mt-pim-service/src/brand/domain/exception"
	"saas-mt-pim-service/src/brand/domain/port"
	"saas-mt-pim-service/src/brand/domain/service"
)

// UpdateBrandUseCase maneja la actualización de marcas
type UpdateBrandUseCase struct {
	brandRepo    port.BrandRepository
	brandService *service.BrandDomainService
	brandMapper  *mapper.BrandMapper
}

// NewUpdateBrandUseCase crea una nueva instancia del caso de uso
func NewUpdateBrandUseCase(
	brandRepo port.BrandRepository,
	brandService *service.BrandDomainService,
	brandMapper *mapper.BrandMapper,
) *UpdateBrandUseCase {
	return &UpdateBrandUseCase{
		brandRepo:    brandRepo,
		brandService: brandService,
		brandMapper:  brandMapper,
	}
}

// Execute ejecuta el caso de uso de actualización de marca
func (uc *UpdateBrandUseCase) Execute(ctx context.Context, brandID string, req *request.UpdateBrandRequest, tenantID string) (*response.BrandResponse, error) {
	// Validar el request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Buscar la marca existente
	brand, err := uc.brandRepo.FindByID(ctx, brandID, tenantID)
	if err != nil {
		return nil, err
	}

	if brand == nil {
		return nil, exception.ErrBrandNotFound
	}

	// Normalizar y validar los datos
	normalizedName, normalizedDescription, normalizedLogoURL, normalizedWebsite, err := uc.brandService.NormalizeBrandData(
		req.Name, req.Description, req.LogoURL, req.Website,
	)
	if err != nil {
		return nil, err
	}

	// Verificar que el nombre sea único (excluyendo la marca actual)
	if err := uc.brandService.ValidateUniqueName(ctx, normalizedName, tenantID, &brandID); err != nil {
		return nil, err
	}

	// Aplicar los cambios
	if err := uc.brandMapper.ApplyUpdateRequest(brand, &request.UpdateBrandRequest{
		Name:        normalizedName,
		Description: normalizedDescription,
		LogoURL:     normalizedLogoURL,
		Website:     normalizedWebsite,
		Color:       req.Color,
	}); err != nil {
		return nil, err
	}

	// Guardar los cambios
	if err := uc.brandRepo.Update(ctx, brand); err != nil {
		return nil, err
	}

	// Convertir a response y retornar
	return uc.brandMapper.ToResponse(brand), nil
}
