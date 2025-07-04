package usecase

import (
	"context"
	"errors"
	"strings"

	"pim/src/attribute/domain/entity"
	"pim/src/attribute/domain/port"
)

var (
	ErrMarketplaceAttributeExists = errors.New("el atributo marketplace ya existe")
)

// CreateMarketplaceAttributeUseCase maneja la creación de atributos marketplace
type CreateMarketplaceAttributeUseCase struct {
	marketplaceAttributeRepo port.MarketplaceAttributeRepository
}

// NewCreateMarketplaceAttributeUseCase crea una nueva instancia del caso de uso
func NewCreateMarketplaceAttributeUseCase(marketplaceAttributeRepo port.MarketplaceAttributeRepository) *CreateMarketplaceAttributeUseCase {
	return &CreateMarketplaceAttributeUseCase{
		marketplaceAttributeRepo: marketplaceAttributeRepo,
	}
}

// Execute ejecuta el caso de uso para crear un atributo marketplace
func (uc *CreateMarketplaceAttributeUseCase) Execute(
	ctx context.Context,
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
	if name == "" {
		return nil, ErrInvalidAttributeName
	}

	if attributeType == "" {
		return nil, errors.New("tipo de atributo es requerido")
	}

	// Generar slug si no se proporciona
	if slug == "" {
		slug = generateSlug(name)
	}

	// Verificar que no existe un atributo con el mismo nombre
	existingByName, err := uc.marketplaceAttributeRepo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if existingByName != nil {
		return nil, ErrMarketplaceAttributeExists
	}

	// Verificar que no existe un atributo con el mismo slug
	existingBySlug, err := uc.marketplaceAttributeRepo.FindBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	if existingBySlug != nil {
		return nil, errors.New("ya existe un atributo con ese slug")
	}

	// Crear la entidad atributo marketplace usando el constructor
	attribute, err := entity.NewMarketplaceAttribute(
		name,
		slug,
		attributeType,
		isFilterable,
		isSearchable,
		isRequiredForListing,
		validationRules,
		sortOrder,
	)
	if err != nil {
		return nil, err
	}

	// Guardar en el repositorio
	err = uc.marketplaceAttributeRepo.Create(ctx, attribute)
	if err != nil {
		return nil, err
	}

	return attribute, nil
}

// generateSlug genera un slug a partir de un nombre
func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "á", "a")
	slug = strings.ReplaceAll(slug, "é", "e")
	slug = strings.ReplaceAll(slug, "í", "i")
	slug = strings.ReplaceAll(slug, "ó", "o")
	slug = strings.ReplaceAll(slug, "ú", "u")
	slug = strings.ReplaceAll(slug, "ñ", "n")
	return slug
}
