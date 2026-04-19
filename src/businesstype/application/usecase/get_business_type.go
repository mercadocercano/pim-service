package usecase

import (
	"context"
	"fmt"
	"saas-mt-pim-service/src/businesstype/domain/entity"
	"saas-mt-pim-service/src/businesstype/domain/port"

	"github.com/google/uuid"
)

// GetBusinessTypeUseCase maneja la obtención de un business type por ID o código
type GetBusinessTypeUseCase struct {
	repository port.BusinessTypeRepository
}

// NewGetBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewGetBusinessTypeUseCase(repository port.BusinessTypeRepository) *GetBusinessTypeUseCase {
	return &GetBusinessTypeUseCase{
		repository: repository,
	}
}

// Execute busca por UUID o por código (ej: "almacen")
func (uc *GetBusinessTypeUseCase) Execute(ctx context.Context, idOrCode string) (*entity.BusinessType, error) {
	if idOrCode == "" {
		return nil, fmt.Errorf("ID o código es requerido")
	}

	// Si es UUID, buscar por ID; sino, buscar por código
	var businessType *entity.BusinessType
	var err error

	if _, parseErr := uuid.Parse(idOrCode); parseErr == nil {
		businessType, err = uc.repository.FindByID(ctx, idOrCode)
	} else {
		businessType, err = uc.repository.FindByCode(ctx, idOrCode)
	}

	if err != nil {
		return nil, fmt.Errorf("error buscando business type: %w", err)
	}
	if businessType == nil {
		return nil, fmt.Errorf("business type no encontrado")
	}

	return businessType, nil
}
