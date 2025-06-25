package usecase

import (
	"context"
	"errors"

	"pim/src/attribute/domain/entity"
	"pim/src/attribute/domain/port"
)

// UpdateMarketplaceAttributeUseCase maneja la actualización de atributos marketplace
type UpdateMarketplaceAttributeUseCase struct {
	marketplaceAttributeRepo port.MarketplaceAttributeRepository
}

// NewUpdateMarketplaceAttributeUseCase crea una nueva instancia del caso de uso
func NewUpdateMarketplaceAttributeUseCase(marketplaceAttributeRepo port.MarketplaceAttributeRepository) *UpdateMarketplaceAttributeUseCase {
	return &UpdateMarketplaceAttributeUseCase{
		marketplaceAttributeRepo: marketplaceAttributeRepo,
	}
}

// Execute ejecuta el caso de uso para actualizar un atributo marketplace
func (uc *UpdateMarketplaceAttributeUseCase) Execute(
	ctx context.Context,
	id string,
	name string,
	attributeType string,
	description *string,
	isRequired bool,
	isSearchable bool,
	isFilterable bool,
	allowedValues []string,
) (*entity.MarketplaceAttribute, error) {
	// Validaciones básicas
	if id == "" {
		return nil, errors.New("id es requerido")
	}

	if name == "" {
		return nil, ErrInvalidAttributeName
	}

	if attributeType == "" {
		return nil, errors.New("tipo de atributo es requerido")
	}

	// Obtener atributo existente
	attribute, err := uc.marketplaceAttributeRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if attribute == nil {
		return nil, ErrMarketplaceAttributeNotFound
	}

	// Verificar que no existe otro atributo con el mismo nombre
	existingAttribute, err := uc.marketplaceAttributeRepo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if existingAttribute != nil && existingAttribute.ID != id {
		return nil, ErrMarketplaceAttributeExists
	}

	// Validar tipos permitidos
	validTypes := []string{"string", "number", "boolean", "date", "enum", "text"}
	isValidType := false
	for _, validType := range validTypes {
		if attributeType == validType {
			isValidType = true
			break
		}
	}
	if !isValidType {
		return nil, errors.New("tipo de atributo inválido")
	}

	// Validar que tipo enum tenga valores permitidos
	if attributeType == "enum" && len(allowedValues) == 0 {
		return nil, errors.New("atributos de tipo enum deben tener valores permitidos")
	}

	// Actualizar los campos
	attribute.Name = name
	attribute.Type = attributeType
	attribute.Description = description
	attribute.IsRequired = isRequired
	attribute.IsSearchable = isSearchable
	attribute.IsFilterable = isFilterable
	attribute.AllowedValues = allowedValues
	attribute.Update() // Actualiza el timestamp

	// Guardar cambios
	err = uc.marketplaceAttributeRepo.Update(ctx, attribute)
	if err != nil {
		return nil, err
	}

	return attribute, nil
}
