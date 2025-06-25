package usecase

import (
	"context"
	"fmt"

	"pim/src/attribute/application/request"
	"pim/src/attribute/application/response"
	"pim/src/attribute/domain/port"
)

// UpdateTenantCustomAttributeUseCase maneja la actualización de atributos personalizados de un tenant
type UpdateTenantCustomAttributeUseCase struct {
	customAttributeRepo port.TenantCustomAttributeRepository
}

// NewUpdateTenantCustomAttributeUseCase crea una nueva instancia del caso de uso
func NewUpdateTenantCustomAttributeUseCase(
	customAttributeRepo port.TenantCustomAttributeRepository,
) *UpdateTenantCustomAttributeUseCase {
	return &UpdateTenantCustomAttributeUseCase{
		customAttributeRepo: customAttributeRepo,
	}
}

// Execute ejecuta el caso de uso para actualizar un atributo personalizado
func (uc *UpdateTenantCustomAttributeUseCase) Execute(
	ctx context.Context,
	req *request.UpdateTenantCustomAttributeRequest,
	tenantID string,
	attributeID string,
) (*response.CustomAttributeResponse, error) {
	// Validar la petición
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// Obtener el atributo existente
	attribute, err := uc.customAttributeRepo.GetByID(ctx, attributeID)
	if err != nil {
		return nil, fmt.Errorf("attribute not found: %w", err)
	}

	// Verificar que el atributo pertenece al tenant
	if attribute.TenantID != tenantID {
		return nil, fmt.Errorf("attribute does not belong to tenant %s", tenantID)
	}

	// Actualizar campos modificables
	if req.Name != nil {
		attribute.Name = *req.Name
	}

	if req.IsFilterable != nil {
		attribute.SetFilterable(*req.IsFilterable)
	}

	if req.IsSearchable != nil {
		attribute.SetSearchable(*req.IsSearchable)
	}

	if req.SortOrder != nil {
		attribute.SortOrder = *req.SortOrder
	}

	// Actualizar reglas de validación si se proporcionan
	if req.ValidationRules != nil {
		attribute.SetValidationRules(req.ValidationRules)
	}

	// Guardar los cambios
	if err := uc.customAttributeRepo.Update(ctx, attribute); err != nil {
		return nil, fmt.Errorf("failed to update attribute: %w", err)
	}

	// Extraer opciones de las reglas de validación si existen
	var options []string
	if attribute.ValidationRules != nil {
		if optionsInterface, exists := attribute.ValidationRules["options"]; exists {
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
	if attribute.ValidationRules != nil {
		if defaultInterface, exists := attribute.ValidationRules["default"]; exists {
			if defaultStr, ok := defaultInterface.(string); ok {
				defaultValue = &defaultStr
			}
		}
	}

	// Convertir a respuesta
	return &response.CustomAttributeResponse{
		ID:           attribute.ID,
		Name:         attribute.Name,
		Slug:         attribute.Slug,
		Type:         attribute.Type,
		IsRequired:   false, // TODO: Extraer de ValidationRules si existe
		IsFilterable: attribute.IsFilterable,
		IsSearchable: attribute.IsSearchable,
		Options:      options,
		DefaultValue: defaultValue,
		SortOrder:    attribute.SortOrder,
		CreatedAt:    attribute.CreatedAt,
		UpdatedAt:    attribute.UpdatedAt,
	}, nil
}
