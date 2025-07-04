package usecase

import (
	"context"
	"fmt"

	"pim/src/businesstype/domain/entity"
	"pim/src/businesstype/domain/port"
)

// GetBusinessTypeTemplateUseCase maneja la obtención de un template específico
type GetBusinessTypeTemplateUseCase struct {
	templateRepo port.BusinessTypeTemplateRepository
}

// NewGetBusinessTypeTemplateUseCase crea una nueva instancia del caso de uso
func NewGetBusinessTypeTemplateUseCase(templateRepo port.BusinessTypeTemplateRepository) *GetBusinessTypeTemplateUseCase {
	return &GetBusinessTypeTemplateUseCase{
		templateRepo: templateRepo,
	}
}

// Execute ejecuta el caso de uso para obtener un template por ID
func (uc *GetBusinessTypeTemplateUseCase) Execute(ctx context.Context, id string) (*entity.BusinessTypeTemplate, error) {
	if id == "" {
		return nil, fmt.Errorf("template id is required")
	}

	template, err := uc.templateRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error finding template: %w", err)
	}

	if template == nil {
		return nil, fmt.Errorf("template with id %s not found", id)
	}

	return template, nil
}