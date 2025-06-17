package usecase

import (
	"context"
	"fmt"
	"pim/src/businesstype/domain/entity"
	"pim/src/businesstype/domain/port"
)

// CreateBusinessTypeUseCase maneja la creación de business types
type CreateBusinessTypeUseCase struct {
	repository port.BusinessTypeRepository
}

// NewCreateBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewCreateBusinessTypeUseCase(repository port.BusinessTypeRepository) *CreateBusinessTypeUseCase {
	return &CreateBusinessTypeUseCase{
		repository: repository,
	}
}

// CreateBusinessTypeRequest representa la solicitud para crear un business type
type CreateBusinessTypeRequest struct {
	Code        string                 `json:"code" binding:"required"`
	Name        string                 `json:"name" binding:"required"`
	Description string                 `json:"description"`
	Icon        string                 `json:"icon"`
	Color       string                 `json:"color"`
	SortOrder   int                    `json:"sort_order"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// Execute ejecuta el caso de uso
func (uc *CreateBusinessTypeUseCase) Execute(ctx context.Context, req CreateBusinessTypeRequest) (*entity.BusinessType, error) {
	// Validar que el código no exista
	existing, err := uc.repository.FindByCode(ctx, req.Code)
	if err != nil {
		return nil, fmt.Errorf("error verificando código existente: %w", err)
	}
	if existing != nil {
		return nil, fmt.Errorf("ya existe un business type con el código: %s", req.Code)
	}

	// Crear nueva entidad
	businessType, err := entity.NewBusinessType(req.Code, req.Name, req.Description)
	if err != nil {
		return nil, fmt.Errorf("error creando business type: %w", err)
	}

	// Configurar campos opcionales
	if req.Icon != "" {
		businessType.Icon = req.Icon
	}
	if req.Color != "" {
		businessType.Color = req.Color
	}
	if req.SortOrder > 0 {
		businessType.SortOrder = req.SortOrder
	}
	if req.Metadata != nil {
		businessType.Metadata = req.Metadata
	}

	// Guardar en repositorio
	if err := uc.repository.Create(ctx, businessType); err != nil {
		return nil, fmt.Errorf("error guardando business type: %w", err)
	}

	return businessType, nil
}
