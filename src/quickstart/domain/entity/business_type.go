package entity

import (
	"errors"
	"time"
)

// ErrInvalidBusinessType representa errores relacionados con la validación de tipos de negocio
var ErrInvalidBusinessType = errors.New("tipo de negocio inválido")

// BusinessType representa un tipo de negocio disponible en el quickstart
type BusinessType struct {
	ID          string
	Name        string
	Description string
	Icon        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewBusinessType crea una nueva instancia de BusinessType con validaciones
func NewBusinessType(id, name, description, icon string) (*BusinessType, error) {
	if id == "" {
		return nil, errors.New("el ID del tipo de negocio es obligatorio")
	}

	if name == "" {
		return nil, errors.New("el nombre del tipo de negocio es obligatorio")
	}

	now := time.Now()
	return &BusinessType{
		ID:          id,
		Name:        name,
		Description: description,
		Icon:        icon,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// Update actualiza los campos del tipo de negocio
func (bt *BusinessType) Update(name, description, icon string) error {
	if name == "" {
		return errors.New("el nombre del tipo de negocio es obligatorio")
	}

	bt.Name = name
	bt.Description = description
	bt.Icon = icon
	bt.UpdatedAt = time.Now()
	return nil
}
