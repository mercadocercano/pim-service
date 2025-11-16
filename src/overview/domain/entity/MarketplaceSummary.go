package entity

import (
	"fmt"
	"time"
	"github.com/google/uuid"
)

// MarketplaceSummary representa la entidad MarketplaceSummary
type MarketplaceSummary struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewMarketplaceSummary crea una nueva instancia de MarketplaceSummary
func NewMarketplaceSummary(tenantID, name string) (*MarketplaceSummary, error) {
	if tenantID == "" {
		return nil, fmt.Errorf("tenant_id es requerido")
	}
	if name == "" {
		return nil, fmt.Errorf("name es requerido")
	}
	
	now := time.Now()
	return &MarketplaceSummary{
		ID:        uuid.New().String(),
		TenantID:  tenantID,
		Name:      name,
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Update actualiza los campos de la entidad
func (e *MarketplaceSummary) Update() {
	e.UpdatedAt = time.Now()
}
