package usecase

import (
	"context"

	"pim/src/quickstart/domain/entity"
	"pim/src/quickstart/domain/service"
)

// UpdateWizardStepUseCase actualiza el progreso del wizard
type UpdateWizardStepUseCase struct {
	wizardService *service.QuickstartWizardService
}

// NewUpdateWizardStepUseCase crea una nueva instancia del caso de uso
func NewUpdateWizardStepUseCase(wizardService *service.QuickstartWizardService) *UpdateWizardStepUseCase {
	return &UpdateWizardStepUseCase{
		wizardService: wizardService,
	}
}

// Execute ejecuta el caso de uso para actualizar un step del wizard
func (uc *UpdateWizardStepUseCase) Execute(ctx context.Context, tenantID, currentStep string, stepData map[string]interface{}) (*entity.TenantQuickstartHistory, error) {
	return uc.wizardService.UpdateWizardStep(ctx, tenantID, currentStep, stepData)
}