package entity

import (
	"fmt"
	"time"
	"github.com/google/uuid"
)

// Scrapingjob representa la entidad scraping_job
type Scrapingjob struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewScrapingjob crea una nueva instancia de Scrapingjob
func NewScrapingjob(tenantID, name string) (*Scrapingjob, error) {
	if tenantID == "" {
		return nil, fmt.Errorf("tenant_id es requerido")
	}
	if name == "" {
		return nil, fmt.Errorf("name es requerido")
	}
	
	now := time.Now()
	return &Scrapingjob{
		ID:        uuid.New().String(),
		TenantID:  tenantID,
		Name:      name,
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Update actualiza los campos de la entidad
func (e *Scrapingjob) Update() {
	e.UpdatedAt = time.Now()
}
