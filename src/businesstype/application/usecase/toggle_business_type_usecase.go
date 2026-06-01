package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/businesstype/domain/entity"
	"saas-mt-pim-service/src/businesstype/domain/port"
)

// ActivateBusinessTypeUseCase activa un business type
type ActivateBusinessTypeUseCase struct {
	repository port.BusinessTypeRepository
}

// NewActivateBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewActivateBusinessTypeUseCase(repository port.BusinessTypeRepository) *ActivateBusinessTypeUseCase {
	return &ActivateBusinessTypeUseCase{repository: repository}
}

// Execute activa un business type por ID
func (uc *ActivateBusinessTypeUseCase) Execute(ctx context.Context, id string) (*entity.BusinessType, error) {
	return toggleStatus(ctx, id, true, uc.repository)
}

// DeactivateBusinessTypeUseCase desactiva un business type
type DeactivateBusinessTypeUseCase struct {
	repository port.BusinessTypeRepository
}

// NewDeactivateBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewDeactivateBusinessTypeUseCase(repository port.BusinessTypeRepository) *DeactivateBusinessTypeUseCase {
	return &DeactivateBusinessTypeUseCase{repository: repository}
}

// Execute desactiva un business type por ID
func (uc *DeactivateBusinessTypeUseCase) Execute(ctx context.Context, id string) (*entity.BusinessType, error) {
	return toggleStatus(ctx, id, false, uc.repository)
}

// toggleStatus es la lógica compartida para activar/desactivar
func toggleStatus(ctx context.Context, id string, activate bool, repo port.BusinessTypeRepository) (*entity.BusinessType, error) {
	if id == "" {
		return nil, fmt.Errorf("id es requerido")
	}

	bt, err := repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error buscando business type: %w", err)
	}
	if bt == nil {
		return nil, fmt.Errorf("business type no encontrado")
	}

	if activate {
		bt.Activate()
	} else {
		bt.Deactivate()
	}

	if err := repo.Update(ctx, bt); err != nil {
		return nil, fmt.Errorf("error actualizando business type: %w", err)
	}

	return bt, nil
}
