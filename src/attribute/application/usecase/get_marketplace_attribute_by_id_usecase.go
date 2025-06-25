package usecase

import (
	"context"
	"errors"

	"pim/src/attribute/domain/entity"
	"pim/src/attribute/domain/port"
)

var (
	ErrMarketplaceAttributeNotFound = errors.New("atributo marketplace no encontrado")
)

// GetMarketplaceAttributeByIDUseCase maneja la obtención de un atributo marketplace por ID
type GetMarketplaceAttributeByIDUseCase struct {
	marketplaceAttributeRepo port.MarketplaceAttributeRepository
}

// NewGetMarketplaceAttributeByIDUseCase crea una nueva instancia del caso de uso
func NewGetMarketplaceAttributeByIDUseCase(marketplaceAttributeRepo port.MarketplaceAttributeRepository) *GetMarketplaceAttributeByIDUseCase {
	return &GetMarketplaceAttributeByIDUseCase{
		marketplaceAttributeRepo: marketplaceAttributeRepo,
	}
}

// Execute ejecuta el caso de uso para obtener un atributo marketplace por ID
func (uc *GetMarketplaceAttributeByIDUseCase) Execute(
	ctx context.Context,
	id string,
) (*entity.MarketplaceAttribute, error) {
	// Validaciones básicas
	if id == "" {
		return nil, errors.New("id es requerido")
	}

	// Obtener atributo por ID
	attribute, err := uc.marketplaceAttributeRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if attribute == nil {
		return nil, ErrMarketplaceAttributeNotFound
	}

	return attribute, nil
}
