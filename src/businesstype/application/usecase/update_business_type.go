package usecase

import (
	"context"
	"fmt"
	"pim/src/businesstype/domain/entity"
	"pim/src/businesstype/domain/port"
	"time"
)

// UpdateBusinessTypeUseCase maneja la actualización de business types
type UpdateBusinessTypeUseCase struct {
	repository port.BusinessTypeRepository
}

// NewUpdateBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewUpdateBusinessTypeUseCase(repository port.BusinessTypeRepository) *UpdateBusinessTypeUseCase {
	return &UpdateBusinessTypeUseCase{
		repository: repository,
	}
}

// UpdateBusinessTypeRequest representa la solicitud para actualizar un business type
type UpdateBusinessTypeRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Icon        string                 `json:"icon"`
	Color       string                 `json:"color"`
	IsActive    *bool                  `json:"is_active"`
	SortOrder   *int                   `json:"sort_order"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// Execute ejecuta el caso de uso
func (uc *UpdateBusinessTypeUseCase) Execute(ctx context.Context, id string, req UpdateBusinessTypeRequest) (*entity.BusinessType, error) {
	// Buscar el business type existente
	businessType, err := uc.repository.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error buscando business type: %w", err)
	}
	if businessType == nil {
		return nil, fmt.Errorf("business type no encontrado")
	}

	// Actualizar campos si se proporcionan
	if req.Name != "" {
		businessType.Name = req.Name
	}
	if req.Description != "" {
		businessType.Description = req.Description
	}
	if req.Icon != "" {
		businessType.Icon = req.Icon
	}
	if req.Color != "" {
		businessType.Color = req.Color
	}
	if req.IsActive != nil {
		businessType.IsActive = *req.IsActive
	}
	if req.SortOrder != nil {
		businessType.SortOrder = *req.SortOrder
	}
	if req.Metadata != nil {
		businessType.Metadata = req.Metadata
	}

	// Actualizar timestamp
	businessType.UpdatedAt = time.Now()

	// Guardar cambios en repositorio
	if err := uc.repository.Update(ctx, businessType); err != nil {
		return nil, fmt.Errorf("error actualizando business type: %w", err)
	}

	return businessType, nil
}
