package usecase

import (
	"context"
	"errors"

	"saas-mt-pim-service/src/attribute/domain/entity"
	"saas-mt-pim-service/src/attribute/domain/port"
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
	slug string,
	attributeType string,
	isFilterable bool,
	isSearchable bool,
	isRequiredForListing bool,
	validationRules map[string]interface{},
	sortOrder int,
) (*entity.MarketplaceAttribute, error) {
	// Validaciones básicas
	if id == "" {
		return nil, errors.New("id es requerido")
	}

	if name == "" {
		return nil, ErrInvalidAttributeName
	}

	if slug == "" {
		return nil, errors.New("slug es requerido")
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
	existingByName, err := uc.marketplaceAttributeRepo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if existingByName != nil && existingByName.ID != id {
		return nil, ErrMarketplaceAttributeExists
	}

	// Verificar que no existe otro atributo con el mismo slug
	existingBySlug, err := uc.marketplaceAttributeRepo.FindBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	if existingBySlug != nil && existingBySlug.ID != id {
		return nil, errors.New("ya existe un atributo con ese slug")
	}

	// Validar tipos permitidos según la constraint de la tabla
	validTypes := []string{"text", "number", "boolean", "select", "multi_select"}
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

	// Actualizar los campos
	attribute.Name = name
	attribute.Slug = slug
	attribute.Type = attributeType
	attribute.IsFilterable = isFilterable
	attribute.IsSearchable = isSearchable
	attribute.IsRequiredForListing = isRequiredForListing
	attribute.ValidationRules = validationRules
	attribute.SortOrder = sortOrder
	attribute.Update() // Actualiza el timestamp

	// Guardar cambios
	err = uc.marketplaceAttributeRepo.Update(ctx, attribute)
	if err != nil {
		return nil, err
	}

	return attribute, nil
}
