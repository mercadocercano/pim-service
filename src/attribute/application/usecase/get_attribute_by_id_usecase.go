package usecase

import (
	"context"
	"errors"

	"pim/src/attribute/domain/entity"
	"pim/src/attribute/domain/port"
)

var (
	ErrAttributeNotFound = errors.New("atributo no encontrado")
)

// GetAttributeByIDUseCase maneja la obtención de un atributo por ID
type GetAttributeByIDUseCase struct {
	attributeRepo port.AttributeRepository
}

// NewGetAttributeByIDUseCase crea una nueva instancia del caso de uso
func NewGetAttributeByIDUseCase(attributeRepo port.AttributeRepository) *GetAttributeByIDUseCase {
	return &GetAttributeByIDUseCase{
		attributeRepo: attributeRepo,
	}
}

// Execute ejecuta el caso de uso para obtener un atributo por ID
func (uc *GetAttributeByIDUseCase) Execute(
	ctx context.Context,
	id string,
	tenantID string,
) (*entity.Attribute, error) {
	// Validaciones básicas
	if id == "" {
		return nil, errors.New("id es requerido")
	}

	if tenantID == "" {
		return nil, errors.New("tenant_id es requerido")
	}

	// Obtener atributo por ID
	attribute, err := uc.attributeRepo.FindByID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	if attribute == nil {
		return nil, ErrAttributeNotFound
	}

	return attribute, nil
}
