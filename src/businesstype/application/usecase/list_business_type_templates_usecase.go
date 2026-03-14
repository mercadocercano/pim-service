package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/businesstype/domain/entity"
	"saas-mt-pim-service/src/businesstype/domain/port"
	cr "github.com/mercadocercano/criteria"
)

// ListBusinessTypeTemplatesUseCase maneja la listado de templates con criterios
type ListBusinessTypeTemplatesUseCase struct {
	templateRepo port.BusinessTypeTemplateRepository
}

// NewListBusinessTypeTemplatesUseCase crea una nueva instancia del caso de uso
func NewListBusinessTypeTemplatesUseCase(templateRepo port.BusinessTypeTemplateRepository) *ListBusinessTypeTemplatesUseCase {
	return &ListBusinessTypeTemplatesUseCase{
		templateRepo: templateRepo,
	}
}

// Execute ejecuta el caso de uso para listar templates usando criterios validados
func (uc *ListBusinessTypeTemplatesUseCase) Execute(ctx context.Context, searchCriteria cr.Criteria) (*cr.ListResponse[entity.BusinessTypeTemplate], error) {
	// Obtener templates usando criteria
	templates, err := uc.templateRepo.SearchByCriteria(ctx, searchCriteria)
	if err != nil {
		return nil, fmt.Errorf("error searching templates: %w", err)
	}

	// Obtener total count
	totalCount, err := uc.templateRepo.CountByCriteria(ctx, searchCriteria)
	if err != nil {
		return nil, fmt.Errorf("error counting templates: %w", err)
	}

	// Usar la respuesta genérica del dominio criteria
	return cr.NewListResponseFromCriteria(templates, totalCount, searchCriteria), nil
}