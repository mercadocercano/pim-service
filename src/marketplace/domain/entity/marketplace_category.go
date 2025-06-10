package entity

import (
	"errors"
	"time"
)

// ErrInvalidMarketplaceCategory representa errores relacionados con la validación de categorías marketplace
var ErrInvalidMarketplaceCategory = errors.New("categoría marketplace inválida")

// MarketplaceCategory representa una categoría global del marketplace
type MarketplaceCategory struct {
	ID          string
	Name        string
	Slug        string
	Description string
	ParentID    *string
	Level       int
	IsActive    bool
	SortOrder   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewMarketplaceCategory crea una nueva instancia de MarketplaceCategory con validaciones
func NewMarketplaceCategory(name, slug, description string, parentID *string) (*MarketplaceCategory, error) {
	if name == "" {
		return nil, errors.New("el nombre de la categoría marketplace es obligatorio")
	}

	if slug == "" {
		return nil, errors.New("el slug de la categoría marketplace es obligatorio")
	}

	now := time.Now()
	return &MarketplaceCategory{
		Name:        name,
		Slug:        slug,
		Description: description,
		ParentID:    parentID,
		Level:       0, // Se calculará automáticamente en el trigger
		IsActive:    true,
		SortOrder:   0,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// Update actualiza los campos de la categoría marketplace
func (mc *MarketplaceCategory) Update(name, slug, description string) error {
	if name == "" {
		return errors.New("el nombre de la categoría marketplace es obligatorio")
	}

	if slug == "" {
		return errors.New("el slug de la categoría marketplace es obligatorio")
	}

	mc.Name = name
	mc.Slug = slug
	mc.Description = description
	mc.UpdatedAt = time.Now()
	return nil
}

// Activate activa la categoría marketplace
func (mc *MarketplaceCategory) Activate() {
	mc.IsActive = true
	mc.UpdatedAt = time.Now()
}

// Deactivate desactiva la categoría marketplace
func (mc *MarketplaceCategory) Deactivate() {
	mc.IsActive = false
	mc.UpdatedAt = time.Now()
}

// UpdateSortOrder actualiza el orden de la categoría
func (mc *MarketplaceCategory) UpdateSortOrder(sortOrder int) {
	mc.SortOrder = sortOrder
	mc.UpdatedAt = time.Now()
}

// IsRoot verifica si la categoría es raíz (sin padre)
func (mc *MarketplaceCategory) IsRoot() bool {
	return mc.ParentID == nil
}

// ValidateLevel verifica que el nivel no exceda el máximo permitido
func (mc *MarketplaceCategory) ValidateLevel(maxLevel int) error {
	if mc.Level > maxLevel {
		return errors.New("la categoría marketplace excede el nivel máximo permitido")
	}
	return nil
}

// SetParent establece la categoría padre
func (mc *MarketplaceCategory) SetParent(parentID *string) {
	mc.ParentID = parentID
	mc.UpdatedAt = time.Now()
}
