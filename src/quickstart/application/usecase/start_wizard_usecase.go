package usecase

import (
	"context"

	"saas-mt-pim-service/src/quickstart/domain/entity"
	"saas-mt-pim-service/src/quickstart/domain/service"
)

// StartWizardUseCase inicia el wizard de configuración para un tenant
type StartWizardUseCase struct {
	wizardService *service.QuickstartWizardService
}

// NewStartWizardUseCase crea una nueva instancia del caso de uso
func NewStartWizardUseCase(wizardService *service.QuickstartWizardService) *StartWizardUseCase {
	return &StartWizardUseCase{
		wizardService: wizardService,
	}
}

// Execute ejecuta el caso de uso para iniciar el wizard
func (uc *StartWizardUseCase) Execute(ctx context.Context, tenantID, businessTypeID string) (*entity.TenantQuickstartHistory, error) {
	return uc.wizardService.StartWizard(ctx, tenantID, businessTypeID)
}