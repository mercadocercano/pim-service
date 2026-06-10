package usecase

import (
	"context"
	businessTypePort "saas-mt-pim-service/src/businesstype/domain/port"
	"saas-mt-pim-service/src/quickstart/domain/entity"
)

// GetBusinessTypesUseCase obtiene todos los tipos de negocio
type GetBusinessTypesUseCase struct {
	businessTypeRepo businessTypePort.BusinessTypeRepository
}

// NewGetBusinessTypesUseCase crea una nueva instancia del caso de uso
func NewGetBusinessTypesUseCase(repo businessTypePort.BusinessTypeRepository) *GetBusinessTypesUseCase {
	return &GetBusinessTypesUseCase{
		businessTypeRepo: repo,
	}
}

// Execute ejecuta el caso de uso
func (uc *GetBusinessTypesUseCase) Execute(ctx context.Context) ([]*entity.BusinessType, error) {
	// Obtener tipos de negocio desde la base de datos usando FindAll
	businessTypes, err := uc.businessTypeRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// Convertir a entidades del dominio quickstart
	result := make([]*entity.BusinessType, 0, len(businessTypes))
	for _, bt := range businessTypes {
		if bt.IsActive {
			result = append(result, &entity.BusinessType{
				ID:          bt.ID,
				Name:        bt.Name,
				Description: bt.Description,
				Icon:        bt.Icon,
			})
		}
	}

	return result, nil
}
