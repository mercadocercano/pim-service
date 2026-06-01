package usecase

import (
	"context"
	"errors"

	"saas-mt-pim-service/src/attribute/domain/entity"
	"saas-mt-pim-service/src/attribute/domain/port"
)

var ErrAttributeValueNotFound = errors.New("valor de atributo no encontrado")

// ListAttributeValuesUseCase lista los valores de un atributo dado
type ListAttributeValuesUseCase struct {
	valueRepo port.AttributeValueRepository
}

// NewListAttributeValuesUseCase crea una nueva instancia del caso de uso
func NewListAttributeValuesUseCase(valueRepo port.AttributeValueRepository) *ListAttributeValuesUseCase {
	return &ListAttributeValuesUseCase{valueRepo: valueRepo}
}

// Execute retorna los valores del atributo indicado
func (uc *ListAttributeValuesUseCase) Execute(ctx context.Context, attributeID string) ([]*entity.AttributeValue, error) {
	if attributeID == "" {
		return nil, errors.New("attribute_id es requerido")
	}
	return uc.valueRepo.FindByAttributeID(ctx, attributeID)
}
