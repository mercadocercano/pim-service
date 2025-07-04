package usecase

import (
	"context"
	"fmt"

	"pim/src/businesstype/domain/port"
)

// DeleteBusinessTypeTemplateUseCase maneja la eliminación de templates
type DeleteBusinessTypeTemplateUseCase struct {
	templateRepo port.BusinessTypeTemplateRepository
}

// NewDeleteBusinessTypeTemplateUseCase crea una nueva instancia del caso de uso
func NewDeleteBusinessTypeTemplateUseCase(templateRepo port.BusinessTypeTemplateRepository) *DeleteBusinessTypeTemplateUseCase {
	return &DeleteBusinessTypeTemplateUseCase{
		templateRepo: templateRepo,
	}
}

// Execute ejecuta el caso de uso para eliminar un template
func (uc *DeleteBusinessTypeTemplateUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("template id is required")
	}

	// Verificar que el template existe
	template, err := uc.templateRepo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("error finding template: %w", err)
	}

	if template == nil {
		return fmt.Errorf("template with id %s not found", id)
	}

	// No permitir eliminar template default
	if template.IsDefault {
		return fmt.Errorf("cannot delete default template")
	}

	// Eliminar el template
	if err := uc.templateRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("error deleting template: %w", err)
	}

	return nil
}