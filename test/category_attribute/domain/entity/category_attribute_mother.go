package entity

import (
	"time"

	"pim/src/category_attribute/domain/entity"

	"github.com/google/uuid"
)

// CategoryAttributeMother implementa el patrón Object Mother para crear entities CategoryAttribute de prueba
type CategoryAttributeMother struct{}

// WithDefaults crea un atributo de categoría con valores por defecto
func (CategoryAttributeMother) WithDefaults() *entity.CategoryAttribute {
	now := time.Now()
	return &entity.CategoryAttribute{
		ID:            uuid.New().String(),
		TenantID:      "tenant-123",
		CategoryID:    "category-123",
		AttributeID:   "attribute-123",
		AllowedValues: []string{"Valor 1", "Valor 2", "Valor 3"},
		Status:        "active",
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

// WithID crea un atributo de categoría con un ID específico
func (ca CategoryAttributeMother) WithID(id string) *entity.CategoryAttribute {
	categoryAttribute := ca.WithDefaults()
	categoryAttribute.ID = id
	return categoryAttribute
}

// WithTenantID crea un atributo de categoría con un TenantID específico
func (ca CategoryAttributeMother) WithTenantID(tenantID string) *entity.CategoryAttribute {
	categoryAttribute := ca.WithDefaults()
	categoryAttribute.TenantID = tenantID
	return categoryAttribute
}

// WithCategoryID crea un atributo de categoría con un CategoryID específico
func (ca CategoryAttributeMother) WithCategoryID(categoryID string) *entity.CategoryAttribute {
	categoryAttribute := ca.WithDefaults()
	categoryAttribute.CategoryID = categoryID
	return categoryAttribute
}

// WithAttributeID crea un atributo de categoría con un AttributeID específico
func (ca CategoryAttributeMother) WithAttributeID(attributeID string) *entity.CategoryAttribute {
	categoryAttribute := ca.WithDefaults()
	categoryAttribute.AttributeID = attributeID
	return categoryAttribute
}

// WithAllowedValues crea un atributo de categoría con valores permitidos específicos
func (ca CategoryAttributeMother) WithAllowedValues(allowedValues []string) *entity.CategoryAttribute {
	categoryAttribute := ca.WithDefaults()
	categoryAttribute.AllowedValues = allowedValues
	return categoryAttribute
}

// WithoutAllowedValues crea un atributo de categoría sin valores permitidos
func (ca CategoryAttributeMother) WithoutAllowedValues() *entity.CategoryAttribute {
	categoryAttribute := ca.WithDefaults()
	categoryAttribute.AllowedValues = []string{}
	return categoryAttribute
}

// Inactive crea un atributo de categoría inactivo
func (ca CategoryAttributeMother) Inactive() *entity.CategoryAttribute {
	categoryAttribute := ca.WithDefaults()
	categoryAttribute.Status = "inactive"
	return categoryAttribute
}

// ColorCategoryAttribute crea un atributo de color para categoría
func (ca CategoryAttributeMother) ColorCategoryAttribute() *entity.CategoryAttribute {
	categoryAttribute := ca.WithDefaults()
	categoryAttribute.AttributeID = "color-attribute-id"
	categoryAttribute.AllowedValues = []string{"Rojo", "Azul", "Verde", "Negro", "Blanco"}
	return categoryAttribute
}

// SizeCategoryAttribute crea un atributo de talla para categoría
func (ca CategoryAttributeMother) SizeCategoryAttribute() *entity.CategoryAttribute {
	categoryAttribute := ca.WithDefaults()
	categoryAttribute.AttributeID = "size-attribute-id"
	categoryAttribute.AllowedValues = []string{"XS", "S", "M", "L", "XL", "XXL"}
	return categoryAttribute
}

// MaterialCategoryAttribute crea un atributo de material para categoría
func (ca CategoryAttributeMother) MaterialCategoryAttribute() *entity.CategoryAttribute {
	categoryAttribute := ca.WithDefaults()
	categoryAttribute.AttributeID = "material-attribute-id"
	categoryAttribute.AllowedValues = []string{"Algodón", "Poliéster", "Lana", "Seda", "Lino"}
	return categoryAttribute
}

// BrandCategoryAttribute crea un atributo de marca para categoría
func (ca CategoryAttributeMother) BrandCategoryAttribute() *entity.CategoryAttribute {
	categoryAttribute := ca.WithDefaults()
	categoryAttribute.AttributeID = "brand-attribute-id"
	categoryAttribute.AllowedValues = []string{"Nike", "Adidas", "Puma", "Reebok"}
	return categoryAttribute
}

// Complete crea un atributo de categoría con todos los parámetros especificados
func (CategoryAttributeMother) Complete(
	id, tenantID, categoryID, attributeID string,
	allowedValues []string,
	status string,
) *entity.CategoryAttribute {
	now := time.Now()
	return &entity.CategoryAttribute{
		ID:            id,
		TenantID:      tenantID,
		CategoryID:    categoryID,
		AttributeID:   attributeID,
		AllowedValues: allowedValues,
		Status:        status,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

// Create retorna una nueva instancia de CategoryAttributeMother
func Create() CategoryAttributeMother {
	return CategoryAttributeMother{}
}
