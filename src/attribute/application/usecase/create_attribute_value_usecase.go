package usecase

import (
	"context"
	"errors"

	"saas-mt-pim-service/src/attribute/domain/entity"
	"saas-mt-pim-service/src/attribute/domain/port"
)

// CreateAttributeValueUseCase maneja la creación de valores de atributo
type CreateAttributeValueUseCase struct {
	valueRepo port.AttributeValueRepository
	attrRepo  port.MarketplaceAttributeRepository
}

// NewCreateAttributeValueUseCase crea una nueva instancia del caso de uso
func NewCreateAttributeValueUseCase(
	valueRepo port.AttributeValueRepository,
	attrRepo port.MarketplaceAttributeRepository,
) *CreateAttributeValueUseCase {
	return &CreateAttributeValueUseCase{valueRepo: valueRepo, attrRepo: attrRepo}
}

// Execute crea un nuevo valor para un atributo existente
func (uc *CreateAttributeValueUseCase) Execute(ctx context.Context, attributeID, value string, sortOrder int) (*entity.AttributeValue, error) {
	if attributeID == "" {
		return nil, errors.New("attribute_id es requerido")
	}

	attr, err := uc.attrRepo.FindByID(ctx, attributeID)
	if err != nil {
		return nil, err
	}
	if attr == nil {
		return nil, ErrMarketplaceAttributeNotFound
	}

	av, err := entity.NewAttributeValue(attributeID, value, sortOrder)
	if err != nil {
		return nil, err
	}

	if err := uc.valueRepo.Create(ctx, av); err != nil {
		return nil, err
	}

	return av, nil
}
