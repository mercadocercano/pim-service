package usecase

import (
	"context"
	"fmt"
	"pim/src/businesstype/domain/entity"
	"pim/src/businesstype/domain/port"
)

// ListBusinessTypesUseCase maneja el listado de business types
type ListBusinessTypesUseCase struct {
	repository port.BusinessTypeRepository
}

// NewListBusinessTypesUseCase crea una nueva instancia del caso de uso
func NewListBusinessTypesUseCase(repository port.BusinessTypeRepository) *ListBusinessTypesUseCase {
	return &ListBusinessTypesUseCase{
		repository: repository,
	}
}

// ListBusinessTypesRequest representa los filtros para listar business types
type ListBusinessTypesRequest struct {
	OnlyActive bool `json:"only_active"`
}

// ListBusinessTypesResponse representa la respuesta del listado
type ListBusinessTypesResponse struct {
	BusinessTypes []*entity.BusinessType `json:"business_types"`
	Total         int                    `json:"total"`
}

// Execute ejecuta el caso de uso
func (uc *ListBusinessTypesUseCase) Execute(ctx context.Context, req ListBusinessTypesRequest) (*ListBusinessTypesResponse, error) {
	var businessTypes []*entity.BusinessType
	var err error

	if req.OnlyActive {
		businessTypes, err = uc.repository.FindActive(ctx)
	} else {
		businessTypes, err = uc.repository.FindAll(ctx)
	}

	if err != nil {
		return nil, fmt.Errorf("error listando business types: %w", err)
	}

	return &ListBusinessTypesResponse{
		BusinessTypes: businessTypes,
		Total:         len(businessTypes),
	}, nil
}
