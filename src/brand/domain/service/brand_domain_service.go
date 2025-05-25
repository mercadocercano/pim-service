package service

import (
	"context"
	"strings"

	"pim/src/brand/domain/exception"
	"pim/src/brand/domain/port"
)

// BrandDomainService contiene la lógica de negocio compleja del dominio Brand
type BrandDomainService struct {
	brandRepo port.BrandRepository
}

// NewBrandDomainService crea una nueva instancia del servicio de dominio
func NewBrandDomainService(brandRepo port.BrandRepository) *BrandDomainService {
	return &BrandDomainService{
		brandRepo: brandRepo,
	}
}

// ValidateUniqueName verifica que el nombre de la marca sea único en el tenant
func (s *BrandDomainService) ValidateUniqueName(ctx context.Context, name, tenantID string, excludeID *string) error {
	// Normalizar el nombre para la comparación
	normalizedName := s.normalizeName(name)

	exists, err := s.brandRepo.ExistsByName(ctx, normalizedName, tenantID, excludeID)
	if err != nil {
		return err
	}

	if exists {
		return exception.ErrBrandAlreadyExists
	}

	return nil
}

// CanDeleteBrand verifica si una marca puede ser eliminada
func (s *BrandDomainService) CanDeleteBrand(ctx context.Context, brandID, tenantID string) error {
	brand, err := s.brandRepo.FindByID(ctx, brandID, tenantID)
	if err != nil {
		return err
	}

	if brand == nil {
		return exception.ErrBrandNotFound
	}

	if !brand.CanBeDeleted() {
		return exception.ErrBrandCannotDelete
	}

	// Aquí podríamos agregar validaciones adicionales como:
	// - Verificar si tiene productos asociados
	// - Verificar si está siendo usada en otros contextos

	return nil
}

// NormalizeBrandData normaliza y valida los datos de una marca
func (s *BrandDomainService) NormalizeBrandData(name, description string, logoURL, website *string) (string, string, *string, *string, error) {
	// Normalizar nombre
	normalizedName := s.normalizeName(name)
	if normalizedName == "" {
		return "", "", nil, nil, exception.ErrBrandNameRequired
	}

	// Normalizar descripción
	normalizedDescription := strings.TrimSpace(description)

	// Validar y normalizar URL del logo
	var normalizedLogoURL *string
	if logoURL != nil && *logoURL != "" {
		trimmedURL := strings.TrimSpace(*logoURL)
		if s.isValidURL(trimmedURL) {
			normalizedLogoURL = &trimmedURL
		}
	}

	// Validar y normalizar website
	var normalizedWebsite *string
	if website != nil && *website != "" {
		trimmedWebsite := strings.TrimSpace(*website)
		if s.isValidURL(trimmedWebsite) {
			normalizedWebsite = &trimmedWebsite
		}
	}

	return normalizedName, normalizedDescription, normalizedLogoURL, normalizedWebsite, nil
}

// normalizeName normaliza el nombre de la marca
func (s *BrandDomainService) normalizeName(name string) string {
	return strings.TrimSpace(name)
}

// isValidURL valida si una URL tiene un formato básico válido
func (s *BrandDomainService) isValidURL(url string) bool {
	// Validación básica de URL
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}
