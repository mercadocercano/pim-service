package usecase

import (
	"context"

	"pim/src/quickstart/domain/entity"
	"pim/src/quickstart/domain/service"
)

// GetWizardStatusUseCase obtiene el estado del wizard para un tenant
type GetWizardStatusUseCase struct {
	wizardService *service.QuickstartWizardService
}

// NewGetWizardStatusUseCase crea una nueva instancia del caso de uso
func NewGetWizardStatusUseCase(wizardService *service.QuickstartWizardService) *GetWizardStatusUseCase {
	return &GetWizardStatusUseCase{
		wizardService: wizardService,
	}
}

// Execute ejecuta el caso de uso para obtener el estado del wizard
func (uc *GetWizardStatusUseCase) Execute(ctx context.Context, tenantID string) (*entity.TenantQuickstartHistory, error) {
	return uc.wizardService.GetWizardStatus(ctx, tenantID)
}