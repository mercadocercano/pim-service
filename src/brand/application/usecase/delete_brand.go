package usecase

import (
	"context"

	"pim/src/brand/domain/exception"
	"pim/src/brand/domain/port"
	"pim/src/brand/domain/service"
)

// DeleteBrandUseCase maneja la eliminación de marcas
type DeleteBrandUseCase struct {
	brandRepo    port.BrandRepository
	brandService *service.BrandDomainService
}

// NewDeleteBrandUseCase crea una nueva instancia del caso de uso
func NewDeleteBrandUseCase(
	brandRepo port.BrandRepository,
	brandService *service.BrandDomainService,
) *DeleteBrandUseCase {
	return &DeleteBrandUseCase{
		brandRepo:    brandRepo,
		brandService: brandService,
	}
}

// Execute ejecuta el caso de uso de eliminación de marca
func (uc *DeleteBrandUseCase) Execute(ctx context.Context, brandID, tenantID string) error {
	// Verificar que la marca puede ser eliminada
	if err := uc.brandService.CanDeleteBrand(ctx, brandID, tenantID); err != nil {
		return err
	}

	// Buscar la marca
	brand, err := uc.brandRepo.FindByID(ctx, brandID, tenantID)
	if err != nil {
		return err
	}

	if brand == nil {
		return exception.ErrBrandNotFound
	}

	// Marcar como eliminada (soft delete)
	if err := brand.Delete(); err != nil {
		return err
	}

	// Guardar los cambios
	return uc.brandRepo.Update(ctx, brand)
}
