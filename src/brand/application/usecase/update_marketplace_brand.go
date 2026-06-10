package usecase

import (
	"context"
	"fmt"

	cr "github.com/mercadocercano/criteria"
	"saas-mt-pim-service/src/brand/application/request"
	"saas-mt-pim-service/src/brand/domain/port"
)

// UpdateMarketplaceBrandUseCase maneja la lógica de negocio para actualizar una marca marketplace
type UpdateMarketplaceBrandUseCase struct {
	repository port.MarketplacebrandRepository
}

// NewUpdateMarketplaceBrandUseCase crea una nueva instancia del caso de uso
func NewUpdateMarketplaceBrandUseCase(repository port.MarketplacebrandRepository) *UpdateMarketplaceBrandUseCase {
	return &UpdateMarketplaceBrandUseCase{
		repository: repository,
	}
}

// Execute ejecuta el caso de uso de actualización de marca marketplace
func (uc *UpdateMarketplaceBrandUseCase) Execute(ctx context.Context, req *request.UpdateMarketplaceBrandRequest) error {
	// Validar que la marca existe
	existingBrand, err := uc.repository.FindByID(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("error buscando marca: %w", err)
	}

	if existingBrand == nil {
		return fmt.Errorf("marca con ID %s no encontrada", req.ID)
	}

	// Validar que el nombre no esté duplicado (si cambió)
	if existingBrand.Name != req.Name {
		if err := uc.validateUniqueName(ctx, req.Name, req.ID); err != nil {
			return err
		}
	}

	// Actualizar los campos de la entidad
	existingBrand.UpdateFields(
		req.Name,
		req.Description,
		req.LogoURL,
		req.Website,
		req.Aliases,
		req.CategoryTags,
		req.Sources,
		req.QualityScore,
		req.IsActive,
	)

	// Persistir los cambios
	if err := uc.repository.Update(ctx, existingBrand); err != nil {
		return fmt.Errorf("error actualizando marca: %w", err)
	}

	return nil
}

// validateUniqueName valida que el nombre de la marca sea único globalmente
func (uc *UpdateMarketplaceBrandUseCase) validateUniqueName(ctx context.Context, name, excludeID string) error {
	// Buscar marcas con el mismo nombre
	criteriaBuilder := cr.NewCriteriaBuilder()
	criteriaBuilder.AddEqualFilter("name", name)

	existingBrands, err := uc.repository.SearchByCriteria(ctx, criteriaBuilder.Build())
	if err != nil {
		return fmt.Errorf("error validando nombre único: %w", err)
	}

	// Verificar si hay alguna marca con el mismo nombre (excluyendo la actual)
	for _, brand := range existingBrands {
		if brand.ID != excludeID {
			return fmt.Errorf("ya existe una marca con el nombre '%s'", name)
		}
	}

	return nil
}
