package entity

import (
	"fmt"
	"time"
	"github.com/google/uuid"
)

// Eanregistry representa la entidad ean_registry
type Eanregistry struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewEanregistry crea una nueva instancia de Eanregistry
func NewEanregistry(tenantID, name string) (*Eanregistry, error) {
	if tenantID == "" {
		return nil, fmt.Errorf("tenant_id es requerido")
	}
	if name == "" {
		return nil, fmt.Errorf("name es requerido")
	}
	
	now := time.Now()
	return &Eanregistry{
		ID:        uuid.New().String(),
		TenantID:  tenantID,
		Name:      name,
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Update actualiza los campos de la entidad
func (e *Eanregistry) Update() {
	e.UpdatedAt = time.Now()
}
