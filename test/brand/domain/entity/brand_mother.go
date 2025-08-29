package entity

import (
	"time"

	"saas-mt-pim-service/src/brand/domain/entity"
	"saas-mt-pim-service/src/brand/domain/value_object"

	"github.com/google/uuid"
)

// BrandMother implementa el patrón Object Mother para crear entities Brand de prueba
type BrandMother struct{}

// stringPtr es una función helper para crear punteros a string
func stringPtr(s string) *string {
	return &s
}

// WithDefaults crea una marca con valores por defecto
func (BrandMother) WithDefaults() *entity.Brand {
	now := time.Now()
	status, _ := value_object.NewBrandStatus("active")

	return &entity.Brand{
		ID:          uuid.New().String(),
		TenantID:    "tenant-123", // TenantID por defecto para tests
		Name:        "Marca de prueba",
		Description: "Descripción de marca de prueba",
		LogoURL:     stringPtr("https://example.com/logo.png"),
		Website:     stringPtr("https://example.com"),
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// WithID crea una marca con un ID específico
func (b BrandMother) WithID(id string) *entity.Brand {
	brand := b.WithDefaults()
	brand.ID = id
	return brand
}

// WithTenantID crea una marca con un TenantID específico
func (b BrandMother) WithTenantID(tenantID string) *entity.Brand {
	brand := b.WithDefaults()
	brand.TenantID = tenantID
	return brand
}

// WithName crea una marca con un nombre específico
func (b BrandMother) WithName(name string) *entity.Brand {
	brand := b.WithDefaults()
	brand.Name = name
	return brand
}

// WithDescription crea una marca con una descripción específica
func (b BrandMother) WithDescription(description string) *entity.Brand {
	brand := b.WithDefaults()
	brand.Description = description
	return brand
}

// WithLogoURL crea una marca con una URL de logo específica
func (b BrandMother) WithLogoURL(logoURL string) *entity.Brand {
	brand := b.WithDefaults()
	brand.LogoURL = stringPtr(logoURL)
	return brand
}

// WithWebsite crea una marca con un sitio web específico
func (b BrandMother) WithWebsite(website string) *entity.Brand {
	brand := b.WithDefaults()
	brand.Website = stringPtr(website)
	return brand
}

// WithoutLogo crea una marca sin logo
func (b BrandMother) WithoutLogo() *entity.Brand {
	brand := b.WithDefaults()
	brand.LogoURL = nil
	return brand
}

// WithoutWebsite crea una marca sin sitio web
func (b BrandMother) WithoutWebsite() *entity.Brand {
	brand := b.WithDefaults()
	brand.Website = nil
	return brand
}

// Inactive crea una marca inactiva
func (b BrandMother) Inactive() *entity.Brand {
	brand := b.WithDefaults()
	status, _ := value_object.NewBrandStatus("inactive")
	brand.Status = status
	return brand
}

// Deleted crea una marca eliminada
func (b BrandMother) Deleted() *entity.Brand {
	brand := b.WithDefaults()
	status, _ := value_object.NewBrandStatus("deleted")
	brand.Status = status
	return brand
}

// Nike crea una marca Nike para tests
func (b BrandMother) Nike() *entity.Brand {
	brand := b.WithDefaults()
	brand.Name = "Nike"
	brand.Description = "Marca deportiva internacional líder en calzado y ropa deportiva"
	brand.LogoURL = stringPtr("https://logos.com/nike-logo.png")
	brand.Website = stringPtr("https://nike.com")
	return brand
}

// Apple crea una marca Apple para tests
func (b BrandMother) Apple() *entity.Brand {
	brand := b.WithDefaults()
	brand.Name = "Apple"
	brand.Description = "Empresa tecnológica estadounidense"
	brand.LogoURL = stringPtr("https://logos.com/apple-logo.png")
	brand.Website = stringPtr("https://apple.com")
	return brand
}

// Complete crea una marca con todos los parámetros especificados
func (BrandMother) Complete(id, tenantID, name, description, logoURL, website, status string) *entity.Brand {
	now := time.Now()
	brandStatus, _ := value_object.NewBrandStatus(status)

	var logoPtr, websitePtr *string
	if logoURL != "" {
		logoPtr = &logoURL
	}
	if website != "" {
		websitePtr = &website
	}

	return &entity.Brand{
		ID:          id,
		TenantID:    tenantID,
		Name:        name,
		Description: description,
		LogoURL:     logoPtr,
		Website:     websitePtr,
		Status:      brandStatus,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Create retorna una nueva instancia de BrandMother
func Create() BrandMother {
	return BrandMother{}
}
