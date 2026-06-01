package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/businesstype/domain/port"
)

// DeleteBusinessTypeUseCase maneja la eliminación de business types
type DeleteBusinessTypeUseCase struct {
	repository port.BusinessTypeRepository
}

// NewDeleteBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewDeleteBusinessTypeUseCase(repository port.BusinessTypeRepository) *DeleteBusinessTypeUseCase {
	return &DeleteBusinessTypeUseCase{repository: repository}
}

// Execute ejecuta la eliminación de un business type por ID
func (uc *DeleteBusinessTypeUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("id es requerido")
	}

	existing, err := uc.repository.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("error verificando business type: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("business type no encontrado")
	}

	if err := uc.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("error eliminando business type: %w", err)
	}

	return nil
}
