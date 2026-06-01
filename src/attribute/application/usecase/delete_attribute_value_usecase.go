package usecase

import (
	"context"
	"errors"

	"saas-mt-pim-service/src/attribute/domain/port"
)

// DeleteAttributeValueUseCase maneja la eliminación de un valor de atributo
type DeleteAttributeValueUseCase struct {
	valueRepo port.AttributeValueRepository
}

// NewDeleteAttributeValueUseCase crea una nueva instancia del caso de uso
func NewDeleteAttributeValueUseCase(valueRepo port.AttributeValueRepository) *DeleteAttributeValueUseCase {
	return &DeleteAttributeValueUseCase{valueRepo: valueRepo}
}

// Execute elimina el valor indicado validando que pertenece al atributo
func (uc *DeleteAttributeValueUseCase) Execute(ctx context.Context, attributeID, valueID string) error {
	if valueID == "" {
		return errors.New("value_id es requerido")
	}

	existing, err := uc.valueRepo.FindByID(ctx, valueID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrAttributeValueNotFound
	}
	if existing.AttributeID != attributeID {
		return ErrAttributeValueNotFound
	}

	return uc.valueRepo.Delete(ctx, valueID)
}
