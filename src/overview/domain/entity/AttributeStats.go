package entity

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// AttributeStats representa la entidad AttributeStats
type AttributeStats struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewAttributeStats crea una nueva instancia de AttributeStats
func NewAttributeStats(tenantID, name string) (*AttributeStats, error) {
	if tenantID == "" {
		return nil, fmt.Errorf("tenant_id es requerido")
	}
	if name == "" {
		return nil, fmt.Errorf("name es requerido")
	}

	now := time.Now()
	return &AttributeStats{
		ID:        uuid.New().String(),
		TenantID:  tenantID,
		Name:      name,
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Update actualiza los campos de la entidad
func (e *AttributeStats) Update() {
	e.UpdatedAt = time.Now()
}
