package usecase

import (
	"context"
	"errors"

	"saas-mt-pim-service/src/attribute/domain/entity"
	"saas-mt-pim-service/src/attribute/domain/port"
)

// ListAttributesUseCase maneja la obtención de atributos básicos
type ListAttributesUseCase struct {
	attributeRepo port.AttributeRepository
}

// NewListAttributesUseCase crea una nueva instancia del caso de uso
func NewListAttributesUseCase(attributeRepo port.AttributeRepository) *ListAttributesUseCase {
	return &ListAttributesUseCase{
		attributeRepo: attributeRepo,
	}
}

// Execute ejecuta el caso de uso para listar atributos
func (uc *ListAttributesUseCase) Execute(
	ctx context.Context,
	tenantID string,
) ([]*entity.Attribute, error) {
	// Validaciones básicas
	if tenantID == "" {
		return nil, errors.New("tenant_id es requerido")
	}

	// Obtener atributos del tenant
	attributes, err := uc.attributeRepo.FindByTenant(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return attributes, nil
}
