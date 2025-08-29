package response

import (
	"encoding/json"
	"time"

	"saas-mt-pim-service/src/quickstart/domain/entity"
)

// SetupResponse representa la respuesta de la configuración del tenant
type SetupResponse struct {
	ID             string          `json:"id"`
	TenantID       string          `json:"tenant_id"`
	BusinessTypeID string          `json:"business_type_id"`
	TemplateID     *string         `json:"template_id"`
	SetupCompleted bool            `json:"setup_completed"`
	SetupData      json.RawMessage `json:"setup_data"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

// FromTenantQuickstartHistoryEntity convierte una entidad TenantQuickstartHistory a SetupResponse
func FromTenantQuickstartHistoryEntity(history *entity.TenantQuickstartHistory) *SetupResponse {
	return &SetupResponse{
		ID:             history.ID,
		TenantID:       history.TenantID,
		BusinessTypeID: history.BusinessTypeID,
		TemplateID:     history.TemplateID,
		SetupCompleted: history.SetupCompleted,
		SetupData:      history.SetupData,
		CreatedAt:      history.CreatedAt,
		UpdatedAt:      history.UpdatedAt,
	}
}

// FromTenantQuickstartHistoryEntities convierte una lista de entidades a SetupResponse
func FromTenantQuickstartHistoryEntities(histories []*entity.TenantQuickstartHistory) []*SetupResponse {
	responses := make([]*SetupResponse, len(histories))
	for i, history := range histories {
		responses[i] = FromTenantQuickstartHistoryEntity(history)
	}
	return responses
}
