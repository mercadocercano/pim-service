package entity

import (
	"encoding/json"
	"time"

	"saas-mt-pim-service/src/quickstart/domain/entity"

	"github.com/google/uuid"
)

// TenantQuickstartHistoryObjectMother proporciona objetos de prueba para TenantQuickstartHistory
type TenantQuickstartHistoryObjectMother struct{}

// NewTenantQuickstartHistoryObjectMother crea una nueva instancia del Object Mother
func NewTenantQuickstartHistoryObjectMother() *TenantQuickstartHistoryObjectMother {
	return &TenantQuickstartHistoryObjectMother{}
}

// Default crea un TenantQuickstartHistory con valores por defecto para testing
func (om *TenantQuickstartHistoryObjectMother) Default() *entity.TenantQuickstartHistory {
	setupData := map[string]interface{}{
		"business_type_id": "550e8400-e29b-41d4-a716-446655440000",
		"step":             "business_type_selected",
		"completed_steps":  []string{"business_type_selected"},
		"total_steps":      7,
		"wizard_version":   "1.0",
	}
	setupDataJSON, _ := json.Marshal(setupData)

	return &entity.TenantQuickstartHistory{
		ID:             uuid.New().String(),
		TenantID:       "123e4567-e89b-12d3-a456-426614174000",
		BusinessTypeID: "550e8400-e29b-41d4-a716-446655440000",
		TemplateID:     nil,
		SetupCompleted: false,
		SetupData:      json.RawMessage(setupDataJSON),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

// WithTenantID crea un TenantQuickstartHistory con un tenant ID específico
func (om *TenantQuickstartHistoryObjectMother) WithTenantID(tenantID string) *entity.TenantQuickstartHistory {
	history := om.Default()
	history.TenantID = tenantID
	return history
}

// WithBusinessTypeID crea un TenantQuickstartHistory con un business type ID específico
func (om *TenantQuickstartHistoryObjectMother) WithBusinessTypeID(businessTypeID string) *entity.TenantQuickstartHistory {
	history := om.Default()
	history.BusinessTypeID = businessTypeID
	return history
}

// WithTemplateID crea un TenantQuickstartHistory con un template ID específico
func (om *TenantQuickstartHistoryObjectMother) WithTemplateID(templateID string) *entity.TenantQuickstartHistory {
	history := om.Default()
	history.TemplateID = &templateID
	return history
}

// Completed crea un TenantQuickstartHistory ya completado
func (om *TenantQuickstartHistoryObjectMother) Completed() *entity.TenantQuickstartHistory {
	history := om.Default()
	history.SetupCompleted = true
	
	setupData := map[string]interface{}{
		"business_type_id": history.BusinessTypeID,
		"step":             "completed",
		"completed":        true,
		"completed_steps":  []string{"business_type_selected", "categories_selection", "brands_selection", "products_selection", "final_review"},
		"total_steps":      7,
		"wizard_version":   "1.0",
		"completed_at":     time.Now().Format(time.RFC3339),
	}
	setupDataJSON, _ := json.Marshal(setupData)
	history.SetupData = json.RawMessage(setupDataJSON)
	
	return history
}

// WithStep crea un TenantQuickstartHistory en un step específico
func (om *TenantQuickstartHistoryObjectMother) WithStep(step string, completedSteps []string) *entity.TenantQuickstartHistory {
	history := om.Default()
	
	setupData := map[string]interface{}{
		"business_type_id": history.BusinessTypeID,
		"step":             step,
		"completed_steps":  completedSteps,
		"total_steps":      7,
		"wizard_version":   "1.0",
	}
	setupDataJSON, _ := json.Marshal(setupData)
	history.SetupData = json.RawMessage(setupDataJSON)
	
	return history
}

// WithSetupData crea un TenantQuickstartHistory con datos de setup específicos
func (om *TenantQuickstartHistoryObjectMother) WithSetupData(setupData map[string]interface{}) *entity.TenantQuickstartHistory {
	history := om.Default()
	
	setupDataJSON, _ := json.Marshal(setupData)
	history.SetupData = json.RawMessage(setupDataJSON)
	
	return history
}

// InProgress crea un TenantQuickstartHistory en progreso en el step de categorías
func (om *TenantQuickstartHistoryObjectMother) InProgress() *entity.TenantQuickstartHistory {
	return om.WithStep("categories_selection", []string{"business_type_selected"})
}

// Empty crea un TenantQuickstartHistory vacío para tests de validación
func (om *TenantQuickstartHistoryObjectMother) Empty() *entity.TenantQuickstartHistory {
	return &entity.TenantQuickstartHistory{}
}

// List crea una lista de TenantQuickstartHistory para testing
func (om *TenantQuickstartHistoryObjectMother) List(count int) []*entity.TenantQuickstartHistory {
	histories := make([]*entity.TenantQuickstartHistory, count)
	for i := 0; i < count; i++ {
		history := om.Default()
		history.TenantID = uuid.New().String()
		history.BusinessTypeID = uuid.New().String()
		histories[i] = history
	}
	return histories
}