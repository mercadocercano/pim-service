package usecase

import (
	"context"
	"errors"

	"pim/src/attribute/domain/entity"
	"pim/src/attribute/domain/port"
)

var (
	ErrInvalidAttributeName = errors.New("nombre de atributo inválido")
	ErrAttributeExists      = errors.New("el atributo ya existe")
)

// CreateAttributeUseCase maneja la creación de atributos básicos
type CreateAttributeUseCase struct {
	attributeRepo port.AttributeRepository
}

// NewCreateAttributeUseCase crea una nueva instancia del caso de uso
func NewCreateAttributeUseCase(attributeRepo port.AttributeRepository) *CreateAttributeUseCase {
	return &CreateAttributeUseCase{
		attributeRepo: attributeRepo,
	}
}

// Execute ejecuta el caso de uso para crear un atributo
func (uc *CreateAttributeUseCase) Execute(
	ctx context.Context,
	tenantID string,
	name string,
) (*entity.Attribute, error) {
	// Validaciones básicas
	if name == "" {
		return nil, ErrInvalidAttributeName
	}

	if tenantID == "" {
		return nil, errors.New("tenant_id es requerido")
	}

	// Verificar que no existe un atributo con el mismo nombre para este tenant
	existingAttributes, err := uc.attributeRepo.FindByTenant(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	// Verificar duplicados
	for _, attr := range existingAttributes {
		if attr.Name == name {
			return nil, ErrAttributeExists
		}
	}

	// Crear la entidad atributo usando el constructor
	attribute, err := entity.NewAttribute(tenantID, name)
	if err != nil {
		return nil, err
	}

	// Guardar en el repositorio
	err = uc.attributeRepo.Create(ctx, attribute)
	if err != nil {
		return nil, err
	}

	return attribute, nil
}
