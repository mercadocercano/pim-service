package response

import (
	"time"

	"pim/src/quickstart/domain/entity"
)

// BusinessTypeResponse representa la respuesta para un tipo de negocio
type BusinessTypeResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// FromEntity convierte una entidad BusinessType a BusinessTypeResponse
func FromBusinessTypeEntity(bt *entity.BusinessType) *BusinessTypeResponse {
	return &BusinessTypeResponse{
		ID:          bt.ID,
		Name:        bt.Name,
		Description: bt.Description,
		Icon:        bt.Icon,
		CreatedAt:   bt.CreatedAt,
		UpdatedAt:   bt.UpdatedAt,
	}
}

// FromBusinessTypeEntities convierte una lista de entidades BusinessType a BusinessTypeResponse
func FromBusinessTypeEntities(businessTypes []*entity.BusinessType) []*BusinessTypeResponse {
	responses := make([]*BusinessTypeResponse, len(businessTypes))
	for i, bt := range businessTypes {
		responses[i] = FromBusinessTypeEntity(bt)
	}
	return responses
}
