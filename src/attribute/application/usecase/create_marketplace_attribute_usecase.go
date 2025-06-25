package usecase

import (
	"context"
	"errors"

	"pim/src/attribute/domain/entity"
	"pim/src/attribute/domain/port"
)

var (
	ErrMarketplaceAttributeExists = errors.New("el atributo marketplace ya existe")
)

// CreateMarketplaceAttributeUseCase maneja la creación de atributos marketplace
type CreateMarketplaceAttributeUseCase struct {
	marketplaceAttributeRepo port.MarketplaceAttributeRepository
}

// NewCreateMarketplaceAttributeUseCase crea una nueva instancia del caso de uso
func NewCreateMarketplaceAttributeUseCase(marketplaceAttributeRepo port.MarketplaceAttributeRepository) *CreateMarketplaceAttributeUseCase {
	return &CreateMarketplaceAttributeUseCase{
		marketplaceAttributeRepo: marketplaceAttributeRepo,
	}
}

// Execute ejecuta el caso de uso para crear un atributo marketplace
func (uc *CreateMarketplaceAttributeUseCase) Execute(
	ctx context.Context,
	name string,
	attributeType string,
	description *string,
	isRequired bool,
	isSearchable bool,
	isFilterable bool,
	allowedValues []string,
) (*entity.MarketplaceAttribute, error) {
	// Validaciones básicas
	if name == "" {
		return nil, ErrInvalidAttributeName
	}

	if attributeType == "" {
		return nil, errors.New("tipo de atributo es requerido")
	}

	// Verificar que no existe un atributo con el mismo nombre
	existingAttribute, err := uc.marketplaceAttributeRepo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if existingAttribute != nil {
		return nil, ErrMarketplaceAttributeExists
	}

	// Crear la entidad atributo marketplace usando el constructor
	attribute, err := entity.NewMarketplaceAttribute(
		name,
		attributeType,
		description,
		isRequired,
		isSearchable,
		isFilterable,
		allowedValues,
	)
	if err != nil {
		return nil, err
	}

	// Guardar en el repositorio
	err = uc.marketplaceAttributeRepo.Create(ctx, attribute)
	if err != nil {
		return nil, err
	}

	return attribute, nil
}
