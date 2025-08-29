package service

import (
	"context"

	"saas-mt-pim-service/src/product/tenant/application/usecase"
	"saas-mt-pim-service/src/quickstart/domain/port"
)

// QuickstartProductService implementa la interfaz ProductService del quickstart
type QuickstartProductService struct {
	createFromTemplateUseCase *usecase.CreateFromTemplateUseCase
}

// NewQuickstartProductService crea una nueva instancia del servicio
func NewQuickstartProductService(
	createFromTemplateUseCase *usecase.CreateFromTemplateUseCase,
) port.ProductService {
	return &QuickstartProductService{
		createFromTemplateUseCase: createFromTemplateUseCase,
	}
}

// CreateFromTemplate implementa la interfaz ProductService
func (s *QuickstartProductService) CreateFromTemplate(ctx context.Context, tenantID string, templateData interface{}) error {
	return s.createFromTemplateUseCase.Execute(ctx, tenantID, templateData)
}
