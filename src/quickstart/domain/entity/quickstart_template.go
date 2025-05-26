package entity

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// TemplateType representa los tipos de plantilla disponibles
type TemplateType string

const (
	TemplateTypeCategory  TemplateType = "category"
	TemplateTypeAttribute TemplateType = "attribute"
	TemplateTypeVariant   TemplateType = "variant"
	TemplateTypeProduct   TemplateType = "product"
)

// QuickstartTemplate representa una plantilla de configuración para el quickstart
type QuickstartTemplate struct {
	ID           string
	BusinessType string
	TemplateType TemplateType
	TemplateData json.RawMessage
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// NewQuickstartTemplate crea una nueva instancia de QuickstartTemplate
func NewQuickstartTemplate(businessType string, templateType TemplateType, templateData json.RawMessage) (*QuickstartTemplate, error) {
	if businessType == "" {
		return nil, errors.New("el tipo de negocio es obligatorio")
	}

	if templateType == "" {
		return nil, errors.New("el tipo de plantilla es obligatorio")
	}

	if len(templateData) == 0 {
		return nil, errors.New("los datos de la plantilla son obligatorios")
	}

	now := time.Now()
	return &QuickstartTemplate{
		ID:           uuid.New().String(),
		BusinessType: businessType,
		TemplateType: templateType,
		TemplateData: templateData,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}

// Update actualiza los datos de la plantilla
func (qt *QuickstartTemplate) Update(templateData json.RawMessage) error {
	if len(templateData) == 0 {
		return errors.New("los datos de la plantilla son obligatorios")
	}

	qt.TemplateData = templateData
	qt.UpdatedAt = time.Now()
	return nil
}

// IsValidTemplateType verifica si el tipo de plantilla es válido
func IsValidTemplateType(templateType string) bool {
	switch TemplateType(templateType) {
	case TemplateTypeCategory, TemplateTypeAttribute, TemplateTypeVariant, TemplateTypeProduct:
		return true
	default:
		return false
	}
}
