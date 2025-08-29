package usecase

import (
	"context"

	"saas-mt-pim-service/src/attribute/domain/entity"
	"saas-mt-pim-service/src/attribute/domain/port"
)

// ListMarketplaceAttributesUseCase maneja la obtención de atributos marketplace
type ListMarketplaceAttributesUseCase struct {
	marketplaceAttributeRepo port.MarketplaceAttributeRepository
}

// NewListMarketplaceAttributesUseCase crea una nueva instancia del caso de uso
func NewListMarketplaceAttributesUseCase(marketplaceAttributeRepo port.MarketplaceAttributeRepository) *ListMarketplaceAttributesUseCase {
	return &ListMarketplaceAttributesUseCase{
		marketplaceAttributeRepo: marketplaceAttributeRepo,
	}
}

// Execute ejecuta el caso de uso para listar atributos marketplace
func (uc *ListMarketplaceAttributesUseCase) Execute(
	ctx context.Context,
) ([]*entity.MarketplaceAttribute, error) {
	// Obtener todos los atributos marketplace
	attributes, err := uc.marketplaceAttributeRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return attributes, nil
}
