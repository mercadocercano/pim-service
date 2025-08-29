package entity

import (
	"time"

	"saas-mt-pim-service/src/brand/domain/exception"
	"saas-mt-pim-service/src/brand/domain/value_object"

	"github.com/google/uuid"
)

// Brand representa la entidad principal del dominio de marcas
type Brand struct {
	ID          string
	TenantID    string
	Name        string
	Description string
	LogoURL     *string
	Website     *string
	Status      value_object.BrandStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewBrand crea una nueva instancia de Brand con validaciones
func NewBrand(tenantID, name, description string, logoURL, website *string) (*Brand, error) {
	if tenantID == "" {
		return nil, exception.ErrBrandTenantRequired
	}

	if name == "" {
		return nil, exception.ErrBrandNameRequired
	}

	now := time.Now()
	return &Brand{
		ID:          uuid.New().String(),
		TenantID:    tenantID,
		Name:        name,
		Description: description,
		LogoURL:     logoURL,
		Website:     website,
		Status:      value_object.BrandStatusActive,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// Update actualiza los campos de la marca
func (b *Brand) Update(name, description string, logoURL, website *string) error {
	if name == "" {
		return exception.ErrBrandNameRequired
	}

	b.Name = name
	b.Description = description
	b.LogoURL = logoURL
	b.Website = website
	b.UpdatedAt = time.Now()
	return nil
}

// Activate establece la marca como activa
func (b *Brand) Activate() {
	b.Status = value_object.BrandStatusActive
	b.UpdatedAt = time.Now()
}

// Deactivate establece la marca como inactiva
func (b *Brand) Deactivate() {
	b.Status = value_object.BrandStatusInactive
	b.UpdatedAt = time.Now()
}

// Delete marca la marca como eliminada (soft delete)
func (b *Brand) Delete() error {
	if !b.CanBeDeleted() {
		return exception.ErrBrandCannotDelete
	}

	b.Status = value_object.BrandStatusDeleted
	b.UpdatedAt = time.Now()
	return nil
}

// IsActive verifica si la marca está activa
func (b *Brand) IsActive() bool {
	return b.Status.IsActive()
}

// IsDeleted verifica si la marca está eliminada
func (b *Brand) IsDeleted() bool {
	return b.Status.IsDeleted()
}

// CanBeDeleted verifica si la marca puede ser eliminada
func (b *Brand) CanBeDeleted() bool {
	// Por ahora permitimos eliminar cualquier marca
	// En el futuro podríamos verificar si tiene productos asociados
	return !b.IsDeleted()
}

// ToReference convierte la marca en una referencia para usar en otros contextos
func (b *Brand) ToReference() *value_object.BrandReference {
	ref, _ := value_object.NewBrandReference(b.ID, b.Name, b.Description)
	return ref
}

// HasLogo verifica si la marca tiene un logo configurado
func (b *Brand) HasLogo() bool {
	return b.LogoURL != nil && *b.LogoURL != ""
}

// HasWebsite verifica si la marca tiene un sitio web configurado
func (b *Brand) HasWebsite() bool {
	return b.Website != nil && *b.Website != ""
}
