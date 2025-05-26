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
	ID           string
	TenantID     string
	BusinessType string
	SetupData    json.RawMessage
	Status       SetupStatus
	ErrorMessage *string
	CreatedAt    time.Time
	CompletedAt  *time.Time
}

// NewTenantQuickstartHistory crea una nueva instancia de TenantQuickstartHistory
func NewTenantQuickstartHistory(tenantID, businessType string, setupData json.RawMessage) (*TenantQuickstartHistory, error) {
	if tenantID == "" {
		return nil, errors.New("el tenant ID es obligatorio")
	}

	if businessType == "" {
		return nil, errors.New("el tipo de negocio es obligatorio")
	}

	if len(setupData) == 0 {
		return nil, errors.New("los datos de configuración son obligatorios")
	}

	return &TenantQuickstartHistory{
		ID:           uuid.New().String(),
		TenantID:     tenantID,
		BusinessType: businessType,
		SetupData:    setupData,
		Status:       SetupStatusPending,
		CreatedAt:    time.Now(),
	}, nil
}

// MarkAsCompleted marca la configuración como completada
func (th *TenantQuickstartHistory) MarkAsCompleted() {
	th.Status = SetupStatusCompleted
	now := time.Now()
	th.CompletedAt = &now
	th.ErrorMessage = nil
}

// MarkAsFailed marca la configuración como fallida
func (th *TenantQuickstartHistory) MarkAsFailed(errorMessage string) {
	th.Status = SetupStatusFailed
	th.ErrorMessage = &errorMessage
}

// IsCompleted verifica si la configuración está completada
func (th *TenantQuickstartHistory) IsCompleted() bool {
	return th.Status == SetupStatusCompleted
}

// IsFailed verifica si la configuración falló
func (th *TenantQuickstartHistory) IsFailed() bool {
	return th.Status == SetupStatusFailed
}

// IsPending verifica si la configuración está pendiente
func (th *TenantQuickstartHistory) IsPending() bool {
	return th.Status == SetupStatusPending
}
