package entity

import (
	"time"

	"saas-mt-pim-service/src/attribute/domain/entity"

	"github.com/google/uuid"
)

// AttributeMother implementa el patrón Object Mother para crear entities Attribute de prueba
type AttributeMother struct{}

// WithDefaults crea un atributo con valores por defecto
func (AttributeMother) WithDefaults() *entity.Attribute {
	now := time.Now()
	return &entity.Attribute{
		ID:        uuid.New().String(),
		TenantID:  "tenant-123",
		Name:      "Atributo de prueba",
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// WithID crea un atributo con un ID específico
func (a AttributeMother) WithID(id string) *entity.Attribute {
	attribute := a.WithDefaults()
	attribute.ID = id
	return attribute
}

// WithTenantID crea un atributo con un TenantID específico
func (a AttributeMother) WithTenantID(tenantID string) *entity.Attribute {
	attribute := a.WithDefaults()
	attribute.TenantID = tenantID
	return attribute
}

// WithName crea un atributo con un nombre específico
func (a AttributeMother) WithName(name string) *entity.Attribute {
	attribute := a.WithDefaults()
	attribute.Name = name
	return attribute
}

// Inactive crea un atributo inactivo
func (a AttributeMother) Inactive() *entity.Attribute {
	attribute := a.WithDefaults()
	attribute.Active = false
	return attribute
}

// ColorAttribute crea un atributo de color para tests
func (a AttributeMother) ColorAttribute() *entity.Attribute {
	attribute := a.WithDefaults()
	attribute.Name = "Color"
	return attribute
}

// SizeAttribute crea un atributo de talla para tests
func (a AttributeMother) SizeAttribute() *entity.Attribute {
	attribute := a.WithDefaults()
	attribute.Name = "Talla"
	return attribute
}

// MaterialAttribute crea un atributo de material para tests
func (a AttributeMother) MaterialAttribute() *entity.Attribute {
	attribute := a.WithDefaults()
	attribute.Name = "Material"
	return attribute
}

// Complete crea un atributo con todos los parámetros especificados
func (AttributeMother) Complete(id, tenantID, name string, active bool) *entity.Attribute {
	now := time.Now()
	return &entity.Attribute{
		ID:        id,
		TenantID:  tenantID,
		Name:      name,
		Active:    active,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Create retorna una nueva instancia de AttributeMother
func Create() AttributeMother {
	return AttributeMother{}
}
