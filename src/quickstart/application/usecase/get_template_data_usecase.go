package usecase

import (
	"context"

	"pim/src/quickstart/domain/service"
)

// GetTemplateDataUseCase obtiene los datos del template para el wizard
type GetTemplateDataUseCase struct {
	wizardService *service.QuickstartWizardService
}

// NewGetTemplateDataUseCase crea una nueva instancia del caso de uso
func NewGetTemplateDataUseCase(wizardService *service.QuickstartWizardService) *GetTemplateDataUseCase {
	return &GetTemplateDataUseCase{
		wizardService: wizardService,
	}
}

// Execute ejecuta el caso de uso para obtener datos del template
func (uc *GetTemplateDataUseCase) Execute(ctx context.Context, businessTypeID, section string) (interface{}, error) {
	return uc.wizardService.GetTemplateData(ctx, businessTypeID, section)
}