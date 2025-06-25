package usecase

import (
	"context"
	"errors"

	"pim/src/attribute/domain/port"
)

// DeleteAttributeUseCase maneja la eliminación de atributos básicos
type DeleteAttributeUseCase struct {
	attributeRepo port.AttributeRepository
}

// NewDeleteAttributeUseCase crea una nueva instancia del caso de uso
func NewDeleteAttributeUseCase(attributeRepo port.AttributeRepository) *DeleteAttributeUseCase {
	return &DeleteAttributeUseCase{
		attributeRepo: attributeRepo,
	}
}

// Execute ejecuta el caso de uso para eliminar un atributo
func (uc *DeleteAttributeUseCase) Execute(
	ctx context.Context,
	id string,
	tenantID string,
) error {
	// Validaciones básicas
	if id == "" {
		return errors.New("id es requerido")
	}

	if tenantID == "" {
		return errors.New("tenant_id es requerido")
	}

	// Verificar que el atributo existe
	attribute, err := uc.attributeRepo.FindByID(ctx, id, tenantID)
	if err != nil {
		return err
	}

	if attribute == nil {
		return ErrAttributeNotFound
	}

	// Eliminar el atributo
	err = uc.attributeRepo.Delete(ctx, id, tenantID)
	if err != nil {
		return err
	}

	return nil
}
