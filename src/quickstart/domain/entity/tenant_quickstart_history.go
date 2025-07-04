package entity

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// SetupStatus representa el estado de la configuración del quickstart
type SetupStatus string

const (
	SetupStatusPending   SetupStatus = "pending"
	SetupStatusCompleted SetupStatus = "completed"
	SetupStatusFailed    SetupStatus = "failed"
)

// TenantQuickstartHistory representa el historial de configuraciones de quickstart por tenant
type TenantQuickstartHistory struct {
	ID             string
	TenantID       string
	BusinessTypeID string
	TemplateID     *string
	SetupCompleted bool
	SetupData      json.RawMessage
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// NewTenantQuickstartHistory crea una nueva instancia de TenantQuickstartHistory
func NewTenantQuickstartHistory(tenantID, businessTypeID string, setupData string) (*TenantQuickstartHistory, error) {
	if tenantID == "" {
		return nil, errors.New("el tenant ID es obligatorio")
	}

	if businessTypeID == "" {
		return nil, errors.New("el business type ID es obligatorio")
	}

	if setupData == "" {
		setupData = "{}"
	}

	now := time.Now()
	return &TenantQuickstartHistory{
		ID:             uuid.New().String(),
		TenantID:       tenantID,
		BusinessTypeID: businessTypeID,
		SetupCompleted: false,
		SetupData:      json.RawMessage(setupData),
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

// MarkAsCompleted marca la configuración como completada
func (th *TenantQuickstartHistory) MarkAsCompleted() {
	th.SetupCompleted = true
	th.UpdatedAt = time.Now()
}

// UpdateSetupData actualiza los datos de configuración
func (th *TenantQuickstartHistory) UpdateSetupData(setupData string) {
	th.SetupData = json.RawMessage(setupData)
	th.UpdatedAt = time.Now()
}

// SetTemplateID establece el ID del template usado
func (th *TenantQuickstartHistory) SetTemplateID(templateID string) {
	th.TemplateID = &templateID
	th.UpdatedAt = time.Now()
}

// IsCompleted verifica si la configuración está completada
func (th *TenantQuickstartHistory) IsCompleted() bool {
	return th.SetupCompleted
}

// IsPending verifica si la configuración está pendiente
func (th *TenantQuickstartHistory) IsPending() bool {
	return !th.SetupCompleted
}
