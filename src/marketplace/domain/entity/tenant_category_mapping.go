package entity

import (
	"errors"
	"time"
)

// ErrInvalidTenantCategoryMapping representa errores relacionados con el mapeo de categorías
var ErrInvalidTenantCategoryMapping = errors.New("mapeo de categoría tenant inválido")

// TenantCategoryMapping representa el mapeo entre una categoría tenant y una categoría marketplace
type TenantCategoryMapping struct {
	ID                    string
	TenantID              string
	CategoryID            string  // ID de la categoría tenant existente
	MarketplaceCategoryID string  // ID de la categoría marketplace
	CustomName            *string // Nombre personalizado que ve el tenant
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

// NewTenantCategoryMapping crea una nueva instancia de TenantCategoryMapping con validaciones
func NewTenantCategoryMapping(tenantID, categoryID, marketplaceCategoryID string, customName *string) (*TenantCategoryMapping, error) {
	if tenantID == "" {
		return nil, errors.New("el ID del tenant es obligatorio")
	}

	if categoryID == "" {
		return nil, errors.New("el ID de la categoría tenant es obligatorio")
	}

	if marketplaceCategoryID == "" {
		return nil, errors.New("el ID de la categoría marketplace es obligatorio")
	}

	now := time.Now()
	return &TenantCategoryMapping{
		TenantID:              tenantID,
		CategoryID:            categoryID,
		MarketplaceCategoryID: marketplaceCategoryID,
		CustomName:            customName,
		CreatedAt:             now,
		UpdatedAt:             now,
	}, nil
}

// UpdateCustomName actualiza el nombre personalizado de la categoría
func (tcm *TenantCategoryMapping) UpdateCustomName(customName *string) {
	tcm.CustomName = customName
	tcm.UpdatedAt = time.Now()
}

// UpdateMarketplaceCategory cambia la categoría marketplace mapeada
func (tcm *TenantCategoryMapping) UpdateMarketplaceCategory(marketplaceCategoryID string) error {
	if marketplaceCategoryID == "" {
		return errors.New("el ID de la categoría marketplace es obligatorio")
	}

	tcm.MarketplaceCategoryID = marketplaceCategoryID
	tcm.UpdatedAt = time.Now()
	return nil
}

// HasCustomName verifica si el mapeo tiene un nombre personalizado
func (tcm *TenantCategoryMapping) HasCustomName() bool {
	return tcm.CustomName != nil && *tcm.CustomName != ""
}

// GetDisplayName retorna el nombre a mostrar (custom o marketplace)
func (tcm *TenantCategoryMapping) GetDisplayName(marketplaceCategoryName string) string {
	if tcm.HasCustomName() {
		return *tcm.CustomName
	}
	return marketplaceCategoryName
}

// ValidateUniqueMapping verifica que no haya duplicados por tenant
func (tcm *TenantCategoryMapping) ValidateUniqueMapping(existingMappings []TenantCategoryMapping) error {
	for _, mapping := range existingMappings {
		if mapping.ID != tcm.ID && mapping.TenantID == tcm.TenantID {
			// Un tenant no puede mapear la misma categoría tenant a múltiples categorías marketplace
			if mapping.CategoryID == tcm.CategoryID {
				return errors.New("la categoría tenant ya está mapeada a otra categoría marketplace")
			}
			// Un tenant no puede mapear múltiples categorías tenant a la misma categoría marketplace
			if mapping.MarketplaceCategoryID == tcm.MarketplaceCategoryID {
				return errors.New("la categoría marketplace ya está mapeada desde otra categoría tenant")
			}
		}
	}
	return nil
}
