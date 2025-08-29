package usecase

import (
	"context"
	"fmt"
	"saas-mt-pim-service/src/businesstype/domain/entity"
	"saas-mt-pim-service/src/businesstype/domain/port"
)

// GetBusinessTypeUseCase maneja la obtención de un business type por ID
type GetBusinessTypeUseCase struct {
	repository port.BusinessTypeRepository
}

// NewGetBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewGetBusinessTypeUseCase(repository port.BusinessTypeRepository) *GetBusinessTypeUseCase {
	return &GetBusinessTypeUseCase{
		repository: repository,
	}
}

// Execute ejecuta el caso de uso
func (uc *GetBusinessTypeUseCase) Execute(ctx context.Context, id string) (*entity.BusinessType, error) {
	if id == "" {
		return nil, fmt.Errorf("ID es requerido")
	}

	businessType, err := uc.repository.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error buscando business type: %w", err)
	}
	if businessType == nil {
		return nil, fmt.Errorf("business type no encontrado")
	}

	return businessType, nil
}
