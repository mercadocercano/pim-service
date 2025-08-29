package entity

import (
	"fmt"
	"time"

	businessTypeEntity "saas-mt-pim-service/src/businesstype/domain/entity"

	"github.com/google/uuid"
)

// BusinessTypeTemplateObjectMother proporciona objetos de prueba para BusinessTypeTemplate
type BusinessTypeTemplateObjectMother struct{}

// NewBusinessTypeTemplateObjectMother crea una nueva instancia del Object Mother
func NewBusinessTypeTemplateObjectMother() *BusinessTypeTemplateObjectMother {
	return &BusinessTypeTemplateObjectMother{}
}

// Default crea un BusinessTypeTemplate con valores por defecto para testing
func (om *BusinessTypeTemplateObjectMother) Default() *businessTypeEntity.BusinessTypeTemplate {
	now := time.Now()
	return &businessTypeEntity.BusinessTypeTemplate{
		ID:             uuid.New().String(),
		BusinessTypeID: "550e8400-e29b-41d4-a716-446655440000",
		Name:           "Template Polirubro Argentina",
		Description:    "Template por defecto para negocios polirubro en Argentina",
		Version:        "1.0.0",
		Region:         "AR",
		Categories: []businessTypeEntity.CategoryTemplate{
			{
				ID:          "cat-1",
				Name:        "Electrónicos",
				Slug:        "electronicos",
				Description: "Productos electrónicos y tecnología",
				ParentID:    "",
				Level:       0,
			},
			{
				ID:          "cat-2",
				Name:        "Hogar",
				Slug:        "hogar",
				Description: "Artículos para el hogar",
				ParentID:    "",
				Level:       0,
			},
		},
		Attributes: []businessTypeEntity.AttributeTemplate{
			{
				ID:           "attr-1",
				Code:         "marca",
				Name:         "Marca",
				Type:         "select",
				IsRequired:   true,
				DefaultValue: "",
				Options:      []string{"Samsung", "LG", "Sony", "Philips"},
			},
			{
				ID:           "attr-2",
				Code:         "color",
				Name:         "Color",
				Type:         "select",
				IsRequired:   false,
				DefaultValue: "blanco",
				Options:      []string{"blanco", "negro", "plateado", "dorado"},
			},
		},
		Products: []businessTypeEntity.ProductTemplate{
			{
				ID:           "prod-1",
				Name:        "Televisor LED 32\"",
				Description: "Televisor LED de 32 pulgadas",
				CategoryID:   "cat-1",
				CategoryName: "Electrónicos",
				BrandID:      "brand-1",
				BrandName:    "Samsung",
				SKU:         "TV-SAM-32-001",
				Price:       45000.0,
				Attributes: map[string]interface{}{
					"marca":    "Samsung",
					"modelo":   "UN32T4300",
					"garantia": "12 meses",
				},
			},
			{
				ID:           "prod-2",
				Name:        "Silla de Oficina",
				Description: "Silla ergonómica para oficina",
				CategoryID:   "cat-2",
				CategoryName: "Hogar",
				BrandID:      "brand-2",
				BrandName:    "Genérica",
				SKU:         "SILLA-OFF-001",
				Price:       25000.0,
				Attributes: map[string]interface{}{
					"material": "cuero sintético",
					"color":    "negro",
					"tamaño":   "mediano",
				},
			},
		},
		Brands:    []string{"Samsung", "LG", "Sony", "Philips", "Genérica"},
		IsActive:  true,
		IsDefault: true,
		Metadata: map[string]interface{}{
			"created_by": "system",
			"region":     "AR",
			"industry":   "retail",
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// WithBusinessTypeID crea un BusinessTypeTemplate con un business type ID específico
func (om *BusinessTypeTemplateObjectMother) WithBusinessTypeID(businessTypeID string) *businessTypeEntity.BusinessTypeTemplate {
	template := om.Default()
	template.BusinessTypeID = businessTypeID
	return template
}

// WithRegion crea un BusinessTypeTemplate con una región específica
func (om *BusinessTypeTemplateObjectMother) WithRegion(region string) *businessTypeEntity.BusinessTypeTemplate {
	template := om.Default()
	template.Region = region
	return template
}

// Inactive crea un BusinessTypeTemplate inactivo
func (om *BusinessTypeTemplateObjectMother) Inactive() *businessTypeEntity.BusinessTypeTemplate {
	template := om.Default()
	template.IsActive = false
	template.IsDefault = false
	return template
}

// WithName crea un BusinessTypeTemplate con un nombre específico
func (om *BusinessTypeTemplateObjectMother) WithName(name string) *businessTypeEntity.BusinessTypeTemplate {
	template := om.Default()
	template.Name = name
	return template
}

// SimpleTemplate crea un BusinessTypeTemplate simple con pocas categorías
func (om *BusinessTypeTemplateObjectMother) SimpleTemplate() *businessTypeEntity.BusinessTypeTemplate {
	template := om.Default()
	template.Name = "Template Simple"
	template.Description = "Template simple para testing"
	template.Categories = []businessTypeEntity.CategoryTemplate{
		{
			Name:        "General",
			Description: "Productos generales",
			ParentName:  "",
			SortOrder:   1,
			Attributes:  []string{"nombre", "precio"},
		},
	}
	template.Attributes = []businessTypeEntity.AttributeTemplate{
		{
			ID:           "attr-nom-1",
			Code:         "nombre",
			Name:         "Nombre",
			Type:         "text",
			IsRequired:   true,
			DefaultValue: "",
			Options:      []string{},
		},
	}
	template.Products = []businessTypeEntity.ProductTemplate{
		{
			ID:           "prod-test-1",
			Name:        "Producto de Prueba",
			Description: "Producto para testing",
			CategoryID:   "cat-gen-1",
			CategoryName: "General",
			BrandID:      "brand-test-1",
			BrandName:    "Test",
			SKU:         "TEST-001",
			Price:       100.0,
			Attributes: map[string]interface{}{
				"nombre": "Producto de Prueba",
			},
		},
	}
	template.Brands = []string{"Test"}
	return template
}

// WithoutCategories crea un BusinessTypeTemplate sin categorías
func (om *BusinessTypeTemplateObjectMother) WithoutCategories() *businessTypeEntity.BusinessTypeTemplate {
	template := om.Default()
	template.Categories = []businessTypeEntity.CategoryTemplate{}
	return template
}

// WithoutProducts crea un BusinessTypeTemplate sin productos
func (om *BusinessTypeTemplateObjectMother) WithoutProducts() *businessTypeEntity.BusinessTypeTemplate {
	template := om.Default()
	template.Products = []businessTypeEntity.ProductTemplate{}
	return template
}

// List crea una lista de BusinessTypeTemplate para testing
func (om *BusinessTypeTemplateObjectMother) List(count int) []*businessTypeEntity.BusinessTypeTemplate {
	templates := make([]*businessTypeEntity.BusinessTypeTemplate, count)
	for i := 0; i < count; i++ {
		template := om.Default()
		template.ID = uuid.New().String()
		template.BusinessTypeID = uuid.New().String()
		template.Name = fmt.Sprintf("Template %d", i+1)
		templates[i] = template
	}
	return templates
}

// Empty crea un BusinessTypeTemplate vacío para tests de validación
func (om *BusinessTypeTemplateObjectMother) Empty() *businessTypeEntity.BusinessTypeTemplate {
	return &businessTypeEntity.BusinessTypeTemplate{}
}