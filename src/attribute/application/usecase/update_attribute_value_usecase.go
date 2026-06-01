package usecase

import (
	"context"
	"errors"

	"saas-mt-pim-service/src/attribute/domain/entity"
	"saas-mt-pim-service/src/attribute/domain/port"
)

// UpdateAttributeValueUseCase maneja la actualización de un valor de atributo
type UpdateAttributeValueUseCase struct {
	valueRepo port.AttributeValueRepository
}

// NewUpdateAttributeValueUseCase crea una nueva instancia del caso de uso
func NewUpdateAttributeValueUseCase(valueRepo port.AttributeValueRepository) *UpdateAttributeValueUseCase {
	return &UpdateAttributeValueUseCase{valueRepo: valueRepo}
}

// Execute actualiza el value y sort_order del valor indicado
func (uc *UpdateAttributeValueUseCase) Execute(ctx context.Context, attributeID, valueID, newValue string, sortOrder int) (*entity.AttributeValue, error) {
	if valueID == "" {
		return nil, errors.New("value_id es requerido")
	}
	if newValue == "" {
		return nil, errors.New("value es requerido")
	}

	existing, err := uc.valueRepo.FindByID(ctx, valueID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, ErrAttributeValueNotFound
	}
	if existing.AttributeID != attributeID {
		return nil, ErrAttributeValueNotFound
	}

	return uc.valueRepo.Update(ctx, valueID, newValue, sortOrder)
}
