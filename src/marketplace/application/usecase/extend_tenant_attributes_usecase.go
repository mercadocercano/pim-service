package usecase

import (
	"context"
	"fmt"
	"strings"

	"pim/src/marketplace/application/request"
	"pim/src/marketplace/application/response"
	"pim/src/marketplace/domain/entity"
	"pim/src/marketplace/domain/port"
)

// ExtendTenantAttributesUseCase maneja la extensión de atributos para un tenant
type ExtendTenantAttributesUseCase struct {
	categoryRepo        port.MarketplaceCategoryRepository
	customAttributeRepo port.TenantCustomAttributeRepository
}

// NewExtendTenantAttributesUseCase crea una nueva instancia del caso de uso
func NewExtendTenantAttributesUseCase(
	categoryRepo port.MarketplaceCategoryRepository,
	customAttributeRepo port.TenantCustomAttributeRepository,
) *ExtendTenantAttributesUseCase {
	return &ExtendTenantAttributesUseCase{
		categoryRepo:        categoryRepo,
		customAttributeRepo: customAttributeRepo,
	}
}

// Execute ejecuta el caso de uso de extensión de atributos tenant
func (uc *ExtendTenantAttributesUseCase) Execute(
	ctx context.Context,
	req *request.ExtendTenantAttributesRequest,
	tenantID string,
) (*response.TenantAttributeExtensionResponse, error) {
	// Debug: verificar que no tengamos nil pointers
	if uc == nil {
		return nil, fmt.Errorf("usecase is nil")
	}
	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}
	if uc.customAttributeRepo == nil {
		return nil, fmt.Errorf("customAttributeRepo is nil")
	}

	// Validar la petición
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// Verificar que la categoría marketplace existe (solo si tenemos repositorio)
	if uc.categoryRepo != nil {
		_, err := uc.categoryRepo.GetByID(ctx, req.MarketplaceCategoryID)
		if err != nil {
			return nil, fmt.Errorf("marketplace category not found: %w", err)
		}
	}

	// Obtener atributos existentes del tenant para esta categoría
	existingAttributes, err := uc.customAttributeRepo.GetByTenantAndCategory(
		ctx, tenantID, req.MarketplaceCategoryID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get existing attributes: %w", err)
	}

	// Crear mapa de atributos existentes por slug para verificar duplicados
	existingSlugMap := make(map[string]*entity.TenantCustomAttribute)
	for _, attr := range existingAttributes {
		existingSlugMap[attr.Slug] = attr
	}

	// Procesar cada atributo personalizado
	var createdAttributes []*entity.TenantCustomAttribute
	var customAttributeResponses []response.CustomAttributeResponse

	for _, customAttr := range req.CustomAttributes {
		// Generar slug a partir del nombre si no se proporciona
		slug := uc.generateSlug(customAttr.Name)

		// Verificar si ya existe un atributo con este slug
		exists, err := uc.customAttributeRepo.ExistsByTenantAndSlug(
			ctx, tenantID, slug, &req.MarketplaceCategoryID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to check attribute uniqueness: %w", err)
		}
		if exists {
			return nil, fmt.Errorf("attribute with slug '%s' already exists for this tenant and category", slug)
		}

		// Crear el atributo personalizado
		attribute, err := entity.NewTenantCustomAttribute(
			tenantID,
			&req.MarketplaceCategoryID,
			customAttr.Name,
			slug,
			customAttr.Type,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create custom attribute '%s': %w", customAttr.Name, err)
		}

		// Configurar propiedades adicionales
		attribute.SetFilterable(customAttr.IsFilterable)
		attribute.SetSearchable(customAttr.IsFilterable) // Si es filterable, también es buscable

		// Configurar reglas de validación para tipos select
		if customAttr.Type == "select" || customAttr.Type == "multi_select" {
			validationRules := map[string]interface{}{
				"options": customAttr.Options,
			}
			if customAttr.DefaultValue != nil {
				validationRules["default"] = *customAttr.DefaultValue
			}
			attribute.SetValidationRules(validationRules)
		}

		// Guardar el atributo
		if err := uc.customAttributeRepo.Save(ctx, attribute); err != nil {
			return nil, fmt.Errorf("failed to save custom attribute '%s': %w", customAttr.Name, err)
		}

		createdAttributes = append(createdAttributes, attribute)

		// Crear respuesta del atributo
		customAttributeResponses = append(customAttributeResponses, response.CustomAttributeResponse{
			ID:           attribute.ID,
			Name:         attribute.Name,
			Type:         attribute.Type,
			IsRequired:   customAttr.IsRequired,
			IsFilterable: attribute.IsFilterable,
			Options:      customAttr.Options,
			DefaultValue: customAttr.DefaultValue,
			CreatedAt:    attribute.CreatedAt,
			UpdatedAt:    attribute.UpdatedAt,
		})
	}

	// Crear respuesta consolidada
	return &response.TenantAttributeExtensionResponse{
		ID:                    fmt.Sprintf("%s-%s", tenantID, req.MarketplaceCategoryID), // ID compuesto
		TenantID:              tenantID,
		MarketplaceCategoryID: req.MarketplaceCategoryID,
		CustomAttributes:      customAttributeResponses,
		CreatedAt:             createdAttributes[0].CreatedAt, // Usar timestamp del primer atributo
		UpdatedAt:             createdAttributes[0].UpdatedAt,
	}, nil
}

// generateSlug genera un slug a partir del nombre del atributo
func (uc *ExtendTenantAttributesUseCase) generateSlug(name string) string {
	// Convertir a minúsculas y reemplazar espacios con guiones
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")

	// Remover caracteres especiales (mantener solo letras, números y guiones)
	var result strings.Builder
	for _, char := range slug {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			result.WriteRune(char)
		}
	}

	return result.String()
}
