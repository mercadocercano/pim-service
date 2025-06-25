package usecase

import (
	"context"
	"errors"

	"pim/src/attribute/domain/port"
)

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
func (uc *DeleteMarketplaceAttributeUseCase) Execute(
	ctx context.Context,
	id string,
) error {
	// Validaciones básicas
	if id == "" {
		return errors.New("id es requerido")
	}

	// Verificar que el atributo existe
	attribute, err := uc.marketplaceAttributeRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if attribute == nil {
		return ErrMarketplaceAttributeNotFound
	}

	// Eliminar el atributo
	err = uc.marketplaceAttributeRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
