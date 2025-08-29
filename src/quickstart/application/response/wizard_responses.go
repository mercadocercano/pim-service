package response

import (
	"time"

	"saas-mt-pim-service/src/quickstart/domain/entity"
)

// WizardStatusResponse representa la respuesta del estado del wizard
type WizardStatusResponse struct {
	WizardID       string    `json:"wizard_id"`
	TenantID       string    `json:"tenant_id"`
	BusinessTypeID string    `json:"business_type_id"`
	SetupData      string    `json:"setup_data"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// TemplateDataResponse representa la respuesta de datos del template
type TemplateDataResponse struct {
	BusinessTypeID string      `json:"business_type_id"`
	Section        string      `json:"section,omitempty"`
	Data           interface{} `json:"data"`
	Pagination     *Pagination `json:"pagination,omitempty"`
}

// Pagination representa los metadatos de paginación
type Pagination struct {
	Page        int  `json:"page"`
	PageSize    int  `json:"page_size"`
	TotalItems  int  `json:"total_items"`
	TotalPages  int  `json:"total_pages"`
	HasNext     bool `json:"has_next"`
	HasPrevious bool `json:"has_previous"`
}

// FromTenantQuickstartHistoryEntityToWizard convierte una entidad a respuesta de wizard
func FromTenantQuickstartHistoryEntityToWizard(entity *entity.TenantQuickstartHistory) *WizardStatusResponse {
	return &WizardStatusResponse{
		WizardID:       entity.ID,
		TenantID:       entity.TenantID,
		BusinessTypeID: entity.BusinessTypeID,
		SetupData:      string(entity.SetupData),
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
	}
}

// NewTemplateDataResponse crea una nueva respuesta de template data
func NewTemplateDataResponse(businessTypeID, section string, data interface{}, pagination *Pagination) *TemplateDataResponse {
	return &TemplateDataResponse{
		BusinessTypeID: businessTypeID,
		Section:        section,
		Data:           data,
		Pagination:     pagination,
	}
}