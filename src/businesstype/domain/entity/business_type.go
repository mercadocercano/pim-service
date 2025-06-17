package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// BusinessType representa un tipo de negocio disponible en el marketplace
type BusinessType struct {
	ID          string                 `json:"id"`
	Code        string                 `json:"code"`        // Código único del business type (ej: "retail", "restaurant")
	Name        string                 `json:"name"`        // Nombre descriptivo (ej: "Tienda Minorista")
	Description string                 `json:"description"` // Descripción detallada
	Icon        string                 `json:"icon"`        // Icono para UI
	Color       string                 `json:"color"`       // Color temático
	IsActive    bool                   `json:"is_active"`
	SortOrder   int                    `json:"sort_order"` // Orden de visualización
	Metadata    map[string]interface{} `json:"metadata"`   // Datos adicionales
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

// NewBusinessType crea una nueva instancia de BusinessType
func NewBusinessType(code, name, description string) (*BusinessType, error) {
	if code == "" {
		return nil, fmt.Errorf("code es requerido")
	}
	if name == "" {
		return nil, fmt.Errorf("name es requerido")
	}

	now := time.Now()
	return &BusinessType{
		ID:          uuid.New().String(),
		Code:        code,
		Name:        name,
		Description: description,
		IsActive:    true,
		SortOrder:   0,
		Metadata:    make(map[string]interface{}),
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// Update actualiza los campos de la entidad
func (bt *BusinessType) Update(name, description, icon, color string) {
	if name != "" {
		bt.Name = name
	}
	if description != "" {
		bt.Description = description
	}
	if icon != "" {
		bt.Icon = icon
	}
	if color != "" {
		bt.Color = color
	}
	bt.UpdatedAt = time.Now()
}

// Activate activa el business type
func (bt *BusinessType) Activate() {
	bt.IsActive = true
	bt.UpdatedAt = time.Now()
}

// Deactivate desactiva el business type
func (bt *BusinessType) Deactivate() {
	bt.IsActive = false
	bt.UpdatedAt = time.Now()
}

// SetSortOrder establece el orden de visualización
func (bt *BusinessType) SetSortOrder(order int) {
	bt.SortOrder = order
	bt.UpdatedAt = time.Now()
}

// AddMetadata agrega metadata al business type
func (bt *BusinessType) AddMetadata(key string, value interface{}) {
	if bt.Metadata == nil {
		bt.Metadata = make(map[string]interface{})
	}
	bt.Metadata[key] = value
	bt.UpdatedAt = time.Now()
}
