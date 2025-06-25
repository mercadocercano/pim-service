package usecase

import (
	"context"
	"errors"

	"pim/src/attribute/domain/entity"
	"pim/src/attribute/domain/port"
)

// UpdateAttributeUseCase maneja la actualización de atributos básicos
type UpdateAttributeUseCase struct {
	attributeRepo port.AttributeRepository
}

// NewUpdateAttributeUseCase crea una nueva instancia del caso de uso
func NewUpdateAttributeUseCase(attributeRepo port.AttributeRepository) *UpdateAttributeUseCase {
	return &UpdateAttributeUseCase{
		attributeRepo: attributeRepo,
	}
}

// Execute ejecuta el caso de uso para actualizar un atributo
func (uc *UpdateAttributeUseCase) Execute(
	ctx context.Context,
	id string,
	tenantID string,
	name string,
) (*entity.Attribute, error) {
	// Validaciones básicas
	if id == "" {
		return nil, errors.New("id es requerido")
	}

	if tenantID == "" {
		return nil, errors.New("tenant_id es requerido")
	}

	if name == "" {
		return nil, ErrInvalidAttributeName
	}

	// Obtener atributo existente
	attribute, err := uc.attributeRepo.FindByID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	if attribute == nil {
		return nil, ErrAttributeNotFound
	}

	// Verificar que no existe otro atributo con el mismo nombre
	existingAttributes, err := uc.attributeRepo.FindByTenant(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	for _, attr := range existingAttributes {
		if attr.Name == name && attr.ID != id {
			return nil, ErrAttributeExists
		}
	}

	// Actualizar los campos
	attribute.Name = name
	attribute.Update() // Actualiza el timestamp

	// Guardar cambios
	err = uc.attributeRepo.Update(ctx, attribute)
	if err != nil {
		return nil, err
	}

	return attribute, nil
}
