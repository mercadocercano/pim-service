package response

import (
	"encoding/json"
	"time"

	"pim/src/quickstart/domain/entity"
)

// SetupResponse representa la respuesta de la configuración del tenant
type SetupResponse struct {
	ID           string          `json:"id"`
	TenantID     string          `json:"tenantId"`
	BusinessType string          `json:"businessType"`
	SetupData    json.RawMessage `json:"setupData"`
	Status       string          `json:"status"`
	ErrorMessage *string         `json:"errorMessage,omitempty"`
	CreatedAt    time.Time       `json:"createdAt"`
	CompletedAt  *time.Time      `json:"completedAt,omitempty"`
}

// FromTenantQuickstartHistoryEntity convierte una entidad TenantQuickstartHistory a SetupResponse
func FromTenantQuickstartHistoryEntity(history *entity.TenantQuickstartHistory) *SetupResponse {
	return &SetupResponse{
		ID:           history.ID,
		TenantID:     history.TenantID,
		BusinessType: history.BusinessType,
		SetupData:    history.SetupData,
		Status:       string(history.Status),
		ErrorMessage: history.ErrorMessage,
		CreatedAt:    history.CreatedAt,
		CompletedAt:  history.CompletedAt,
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
