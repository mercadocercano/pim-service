package entity

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// Attribute representa la entidad attribute
type Attribute struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewAttribute crea una nueva instancia de Attribute
func NewAttribute(tenantID, name string) (*Attribute, error) {
	if tenantID == "" {
		return nil, fmt.Errorf("tenant_id es requerido")
	}
	if name == "" {
		return nil, fmt.Errorf("name es requerido")
	}

	now := time.Now()
	return &Attribute{
		ID:        uuid.New().String(),
		TenantID:  tenantID,
		Name:      name,
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Update actualiza los campos de la entidad
func (e *Attribute) Update() {
	e.UpdatedAt = time.Now()
}
