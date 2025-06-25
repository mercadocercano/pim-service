package usecase

import (
	"context"
	"fmt"

	"pim/src/attribute/application/request"
	"pim/src/attribute/application/response"
	"pim/src/attribute/domain/entity"
	"pim/src/attribute/domain/port"
)

// GetTenantCustomAttributesUseCase maneja la obtención de atributos personalizados de un tenant
type GetTenantCustomAttributesUseCase struct {
	customAttributeRepo port.TenantCustomAttributeRepository
}

// NewGetTenantCustomAttributesUseCase crea una nueva instancia del caso de uso
func NewGetTenantCustomAttributesUseCase(
	customAttributeRepo port.TenantCustomAttributeRepository,
) *GetTenantCustomAttributesUseCase {
	return &GetTenantCustomAttributesUseCase{
		customAttributeRepo: customAttributeRepo,
	}
}

// Execute ejecuta el caso de uso para obtener atributos personalizados
func (uc *GetTenantCustomAttributesUseCase) Execute(
	ctx context.Context,
	req *request.GetTenantCustomAttributesRequest,
) (*response.TenantCustomAttributesListResponse, error) {
	// Validar la petición
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var attributes []*entity.TenantCustomAttribute
	var err error

	// Obtener atributos según los filtros especificados
	if req.MarketplaceCategoryID != nil {
		// Filtrar por categoría específica
		attributes, err = uc.customAttributeRepo.GetByTenantAndCategory(
			ctx, req.TenantID, *req.MarketplaceCategoryID,
		)
	} else if req.IsFilterable != nil && *req.IsFilterable {
		// Obtener solo atributos filtrables
		attributes, err = uc.customAttributeRepo.GetFilterableByTenant(ctx, req.TenantID)
	} else if req.IsSearchable != nil && *req.IsSearchable {
		// Obtener solo atributos buscables
		attributes, err = uc.customAttributeRepo.GetSearchableByTenant(ctx, req.TenantID)
	} else {
		// Obtener todos los atributos del tenant
		attributes, err = uc.customAttributeRepo.GetByTenantID(ctx, req.TenantID)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get tenant custom attributes: %w", err)
	}

	// Filtrar por tipo si se especifica
	if req.AttributeType != nil {
		filteredAttributes := make([]*entity.TenantCustomAttribute, 0)
		for _, attr := range attributes {
			if attr.Type == *req.AttributeType {
				filteredAttributes = append(filteredAttributes, attr)
			}
		}
		attributes = filteredAttributes
	}

	// Convertir a respuesta
	attributeResponses := make([]response.CustomAttributeResponse, 0, len(attributes))
	for _, attr := range attributes {
		// Extraer opciones de las reglas de validación si existen
		var options []string
		if attr.ValidationRules != nil {
			if optionsInterface, exists := attr.ValidationRules["options"]; exists {
				if optionsSlice, ok := optionsInterface.([]interface{}); ok {
					for _, option := range optionsSlice {
						if optionStr, ok := option.(string); ok {
							options = append(options, optionStr)
						}
					}
				}
			}
		}

		// Extraer valor por defecto
		var defaultValue *string
		if attr.ValidationRules != nil {
			if defaultInterface, exists := attr.ValidationRules["default"]; exists {
				if defaultStr, ok := defaultInterface.(string); ok {
					defaultValue = &defaultStr
				}
			}
		}

		attributeResponses = append(attributeResponses, response.CustomAttributeResponse{
			ID:           attr.ID,
			Name:         attr.Name,
			Slug:         attr.Slug,
			Type:         attr.Type,
			IsRequired:   false, // TODO: Extraer de ValidationRules si existe
			IsFilterable: attr.IsFilterable,
			IsSearchable: attr.IsSearchable,
			Options:      options,
			DefaultValue: defaultValue,
			SortOrder:    attr.SortOrder,
			CreatedAt:    attr.CreatedAt,
			UpdatedAt:    attr.UpdatedAt,
		})
	}

	return &response.TenantCustomAttributesListResponse{
		TenantID:         req.TenantID,
		TotalCount:       len(attributeResponses),
		CustomAttributes: attributeResponses,
		FilteredBy: map[string]interface{}{
			"marketplace_category_id": req.MarketplaceCategoryID,
			"attribute_type":          req.AttributeType,
			"is_filterable":           req.IsFilterable,
			"is_searchable":           req.IsSearchable,
		},
	}, nil
}
