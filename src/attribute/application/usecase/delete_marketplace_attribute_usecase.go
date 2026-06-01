package usecase

import (
	"context"
	"errors"

	"saas-mt-pim-service/src/attribute/domain/port"
)

// ErrAttributeInUse se retorna cuando se intenta eliminar un atributo en uso
var ErrAttributeInUse = errors.New("el atributo está en uso y no puede eliminarse")

// DeleteMarketplaceAttributeUseCase maneja la eliminación de atributos marketplace
type DeleteMarketplaceAttributeUseCase struct {
	marketplaceAttributeRepo port.MarketplaceAttributeRepository
}

// NewDeleteMarketplaceAttributeUseCase crea una nueva instancia del caso de uso
func NewDeleteMarketplaceAttributeUseCase(marketplaceAttributeRepo port.MarketplaceAttributeRepository) *DeleteMarketplaceAttributeUseCase {
	return &DeleteMarketplaceAttributeUseCase{
		marketplaceAttributeRepo: marketplaceAttributeRepo,
	}
}

// Execute ejecuta el caso de uso para eliminar un atributo marketplace
func (uc *DeleteMarketplaceAttributeUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id es requerido")
	}

	attribute, err := uc.marketplaceAttributeRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if attribute == nil {
		return ErrMarketplaceAttributeNotFound
	}

	inUse, err := uc.marketplaceAttributeRepo.IsInUse(ctx, id)
	if err != nil {
		return err
	}
	if inUse {
		return ErrAttributeInUse
	}

	return uc.marketplaceAttributeRepo.Delete(ctx, id)
}
