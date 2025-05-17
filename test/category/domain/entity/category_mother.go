package entity

import (
	"time"

	"pim/src/category/domain/entity"

	"github.com/google/uuid"
)

// CategoryMother implementa el patrón Object Mother para crear entities Category de prueba
type CategoryMother struct{}

// WithDefaults crea una categoría con valores por defecto
func (CategoryMother) WithDefaults() *entity.Category {
	now := time.Now()
	return &entity.Category{
		ID:          uuid.New().String(),
		Name:        "Categoría de prueba",
		Description: "Descripción de prueba",
		ParentID:    nil,
		Status:      "active",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// WithID crea una categoría con un ID específico
func (c CategoryMother) WithID(id string) *entity.Category {
	category := c.WithDefaults()
	category.ID = id
	return category
}

// WithName crea una categoría con un nombre específico
func (c CategoryMother) WithName(name string) *entity.Category {
	category := c.WithDefaults()
	category.Name = name
	return category
}

// WithParent crea una categoría con un padre específico
func (c CategoryMother) WithParent(parentID string) *entity.Category {
	category := c.WithDefaults()
	var parentIDPtr = parentID
	category.ParentID = &parentIDPtr
	return category
}

// Inactive crea una categoría inactiva
func (c CategoryMother) Inactive() *entity.Category {
	category := c.WithDefaults()
	category.Status = "inactive"
	return category
}

// Complete crea una categoría con todos los parámetros especificados
func (CategoryMother) Complete(id, name, description, parentID string, active bool) *entity.Category {
	now := time.Now()
	var parentIDPtr *string
	if parentID != "" {
		parentIDPtr = &parentID
	}

	status := "inactive"
	if active {
		status = "active"
	}

	return &entity.Category{
		ID:          id,
		Name:        name,
		Description: description,
		ParentID:    parentIDPtr,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Create retorna una nueva instancia de CategoryMother
func Create() CategoryMother {
	return CategoryMother{}
}
