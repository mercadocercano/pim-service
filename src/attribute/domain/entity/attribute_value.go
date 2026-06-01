package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// AttributeValue representa un valor predefinido de un atributo tipo select/multi_select
type AttributeValue struct {
	ID          string    `json:"id"`
	AttributeID string    `json:"attribute_id"`
	Value       string    `json:"value"`
	Slug        string    `json:"slug"`
	SortOrder   int       `json:"sort_order"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

// NewAttributeValue crea una nueva instancia validada de AttributeValue
func NewAttributeValue(attributeID, value string, sortOrder int) (*AttributeValue, error) {
	if attributeID == "" {
		return nil, fmt.Errorf("attribute_id es requerido")
	}
	if strings.TrimSpace(value) == "" {
		return nil, fmt.Errorf("value es requerido")
	}

	slug := generateAttributeValueSlug(value)

	return &AttributeValue{
		ID:          uuid.New().String(),
		AttributeID: attributeID,
		Value:       value,
		Slug:        slug,
		SortOrder:   sortOrder,
		IsActive:    true,
		CreatedAt:   time.Now(),
	}, nil
}

// generateAttributeValueSlug genera un slug a partir del value
func generateAttributeValueSlug(value string) string {
	slug := strings.ToLower(strings.TrimSpace(value))
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "á", "a")
	slug = strings.ReplaceAll(slug, "é", "e")
	slug = strings.ReplaceAll(slug, "í", "i")
	slug = strings.ReplaceAll(slug, "ó", "o")
	slug = strings.ReplaceAll(slug, "ú", "u")
	slug = strings.ReplaceAll(slug, "ñ", "n")
	return slug
}
