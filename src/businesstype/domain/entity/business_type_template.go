package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// BusinessTypeTemplate representa una plantilla de configuración quickstart para un business type
type BusinessTypeTemplate struct {
	ID             string                 `json:"id"`
	BusinessTypeID string                 `json:"business_type_id"` // Referencia al business type
	Name           string                 `json:"name"`             // Nombre de la plantilla
	Description    string                 `json:"description"`      // Descripción de qué incluye
	Version        string                 `json:"version"`          // Versión de la plantilla
	Region         string                 `json:"region"`           // Región o país (ej: "AR", "MX", "GLOBAL")
	Categories     []CategoryTemplate     `json:"categories"`       // Categorías incluidas
	Attributes     []AttributeTemplate    `json:"attributes"`       // Atributos incluidos
	Products       []ProductTemplate      `json:"products"`         // Productos de ejemplo
	Brands         []BrandTemplate        `json:"brands"`           // Marcas sugeridas
	IsActive       bool                   `json:"is_active"`
	IsDefault      bool                   `json:"is_default"` // Si es la plantilla por defecto
	Metadata       map[string]interface{} `json:"metadata"`   // Configuración adicional
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
}

// BrandTemplate representa una marca sugerida en la plantilla
type BrandTemplate struct {
	Name                   string   `json:"name"`
	SuggestedForCategories []string `json:"suggested_for_categories,omitempty"`
}

// CategoryTemplate representa una categoría en la plantilla
type CategoryTemplate struct {
	ID          string `json:"id"`                    // ID de la categoría del marketplace
	Name        string `json:"name"`                  // Nombre legible de la categoría
	Slug        string `json:"slug"`                  // Slug de la categoría
	Description string `json:"description,omitempty"` // Descripción opcional
	ParentID    string `json:"parent_id,omitempty"`   // ID de la categoría padre
	Level       int    `json:"level"`                 // Nivel de jerarquía (0 = root)
}

// AttributeTemplate representa un atributo en la plantilla
type AttributeTemplate struct {
	ID           string   `json:"id"`                      // ID del atributo del marketplace
	Code         string   `json:"code"`                    // Código único del atributo
	Name         string   `json:"name"`                    // Nombre del atributo
	Type         string   `json:"type"`                    // Tipo: text, number, select, etc
	IsRequired   bool     `json:"is_required"`             // Si es obligatorio
	DefaultValue string   `json:"default_value,omitempty"` // Valor por defecto
	Options      []string `json:"options,omitempty"`       // Opciones para tipo select
}

// ProductTemplate representa un producto de ejemplo en la plantilla
type ProductTemplate struct {
	ID           string                 `json:"id,omitempty"`            // ID del producto (opcional para templates)
	Name         string                 `json:"name"`                    // Nombre del producto
	Description  string                 `json:"description,omitempty"`   // Descripción
	CategoryID   string                 `json:"category_id"`             // ID de la categoría
	CategoryName string                 `json:"category_name,omitempty"` // Nombre de la categoría (para display)
	BrandID      string                 `json:"brand_id,omitempty"`      // ID de la marca
	BrandName    string                 `json:"brand_name,omitempty"`    // Nombre de la marca
	SKU          string                 `json:"sku"`                     // SKU de ejemplo
	Price        float64                `json:"price"`                   // Precio sugerido
	Attributes   map[string]interface{} `json:"attributes,omitempty"`    // Valores de atributos
}

// NewBusinessTypeTemplate crea una nueva instancia de BusinessTypeTemplate
func NewBusinessTypeTemplate(businessTypeID, name, description, version, region string) (*BusinessTypeTemplate, error) {
	if businessTypeID == "" {
		return nil, fmt.Errorf("business_type_id es requerido")
	}
	if name == "" {
		return nil, fmt.Errorf("name es requerido")
	}
	if version == "" {
		version = "1.0.0"
	}
	if region == "" {
		region = "GLOBAL"
	}

	now := time.Now()
	return &BusinessTypeTemplate{
		ID:             uuid.New().String(),
		BusinessTypeID: businessTypeID,
		Name:           name,
		Description:    description,
		Version:        version,
		Region:         region,
		Categories:     []CategoryTemplate{},
		Attributes:     []AttributeTemplate{},
		Products:       []ProductTemplate{},
		Brands:         []BrandTemplate{},
		IsActive:       true,
		IsDefault:      false,
		Metadata:       make(map[string]interface{}),
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

// AddCategory agrega una categoría a la plantilla
func (bt *BusinessTypeTemplate) AddCategory(category CategoryTemplate) {
	bt.Categories = append(bt.Categories, category)
	bt.UpdatedAt = time.Now()
}

// AddAttribute agrega un atributo a la plantilla
func (bt *BusinessTypeTemplate) AddAttribute(attribute AttributeTemplate) {
	bt.Attributes = append(bt.Attributes, attribute)
	bt.UpdatedAt = time.Now()
}

// AddProduct agrega un producto de ejemplo a la plantilla
func (bt *BusinessTypeTemplate) AddProduct(product ProductTemplate) {
	bt.Products = append(bt.Products, product)
	bt.UpdatedAt = time.Now()
}

// SetAsDefault marca esta plantilla como la por defecto
func (bt *BusinessTypeTemplate) SetAsDefault() {
	bt.IsDefault = true
	bt.UpdatedAt = time.Now()
}

// Activate activa la plantilla
func (bt *BusinessTypeTemplate) Activate() {
	bt.IsActive = true
	bt.UpdatedAt = time.Now()
}

// Deactivate desactiva la plantilla
func (bt *BusinessTypeTemplate) Deactivate() {
	bt.IsActive = false
	bt.UpdatedAt = time.Now()
}
