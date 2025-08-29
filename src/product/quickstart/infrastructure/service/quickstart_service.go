package service

import (
	"context"

	"saas-mt-pim-service/src/product/quickstart/application/usecase"
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
	// Convertir templateData a CreateFromTemplateRequest
	// TODO: Implementar conversión adecuada según la estructura de templateData
	request := usecase.CreateFromTemplateRequest{
		TenantID: tenantID,
		// Agregar otros campos según templateData
	}

	_, err := s.createFromTemplateUseCase.Execute(ctx, request)
	return err
}
