package entity

import (
	"fmt"
	"time"
	"github.com/google/uuid"
)

// ProductStats representa la entidad ProductStats
type ProductStats struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewProductStats crea una nueva instancia de ProductStats
func NewProductStats(tenantID, name string) (*ProductStats, error) {
	if tenantID == "" {
		return nil, fmt.Errorf("tenant_id es requerido")
	}
	if name == "" {
		return nil, fmt.Errorf("name es requerido")
	}
	
	now := time.Now()
	return &ProductStats{
		ID:        uuid.New().String(),
		TenantID:  tenantID,
		Name:      name,
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Update actualiza los campos de la entidad
func (e *ProductStats) Update() {
	e.UpdatedAt = time.Now()
}
