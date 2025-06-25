package entity

import (
	"errors"
	"time"
)

// ErrInvalidTenantAttributeExtension representa errores relacionados con extensiones de atributos
var ErrInvalidTenantAttributeExtension = errors.New("extensión de atributo tenant inválida")

// TenantAttributeExtension representa la extensión de un atributo marketplace por un tenant específico
type TenantAttributeExtension struct {
	ID                     string
	TenantID               string
	MarketplaceAttributeID string
	CustomName             *string  // Nombre personalizado del atributo
	AdditionalValues       []string // Valores adicionales específicos del tenant
	IsActive               bool
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

// NewTenantAttributeExtension crea una nueva instancia de TenantAttributeExtension
func NewTenantAttributeExtension(tenantID, marketplaceAttributeID string, customName *string) (*TenantAttributeExtension, error) {
	if tenantID == "" {
		return nil, errors.New("el ID del tenant es obligatorio")
	}

	if marketplaceAttributeID == "" {
		return nil, errors.New("el ID del atributo marketplace es obligatorio")
	}

	now := time.Now()
	return &TenantAttributeExtension{
		TenantID:               tenantID,
		MarketplaceAttributeID: marketplaceAttributeID,
		CustomName:             customName,
		AdditionalValues:       []string{},
		IsActive:               true,
		CreatedAt:              now,
		UpdatedAt:              now,
	}, nil
}

// UpdateCustomName actualiza el nombre personalizado del atributo
func (tae *TenantAttributeExtension) UpdateCustomName(customName *string) {
	tae.CustomName = customName
	tae.UpdatedAt = time.Now()
}

// AddValue agrega un valor adicional específico del tenant
func (tae *TenantAttributeExtension) AddValue(value string) error {
	if value == "" {
		return errors.New("el valor adicional no puede estar vacío")
	}

	// Verificar que el valor no exista ya
	for _, existingValue := range tae.AdditionalValues {
		if existingValue == value {
			return errors.New("el valor ya existe en la extensión")
		}
	}

	tae.AdditionalValues = append(tae.AdditionalValues, value)
	tae.UpdatedAt = time.Now()
	return nil
}

// RemoveValue elimina un valor adicional específico del tenant
func (tae *TenantAttributeExtension) RemoveValue(value string) {
	for i, val := range tae.AdditionalValues {
		if val == value {
			tae.AdditionalValues = append(tae.AdditionalValues[:i], tae.AdditionalValues[i+1:]...)
			tae.UpdatedAt = time.Now()
			break
		}
	}
}

// UpdateValues reemplaza todos los valores adicionales
func (tae *TenantAttributeExtension) UpdateValues(values []string) {
	// Filtrar valores vacíos y duplicados
	uniqueValues := make([]string, 0, len(values))
	seen := make(map[string]bool)

	for _, value := range values {
		if value != "" && !seen[value] {
			uniqueValues = append(uniqueValues, value)
			seen[value] = true
		}
	}

	tae.AdditionalValues = uniqueValues
	tae.UpdatedAt = time.Now()
}

// Activate activa la extensión del atributo
func (tae *TenantAttributeExtension) Activate() {
	tae.IsActive = true
	tae.UpdatedAt = time.Now()
}

// Deactivate desactiva la extensión del atributo
func (tae *TenantAttributeExtension) Deactivate() {
	tae.IsActive = false
	tae.UpdatedAt = time.Now()
}

// HasCustomName verifica si la extensión tiene un nombre personalizado
func (tae *TenantAttributeExtension) HasCustomName() bool {
	return tae.CustomName != nil && *tae.CustomName != ""
}

// HasAdditionalValues verifica si la extensión tiene valores adicionales
func (tae *TenantAttributeExtension) HasAdditionalValues() bool {
	return len(tae.AdditionalValues) > 0
}

// GetDisplayName retorna el nombre a mostrar (custom o marketplace)
func (tae *TenantAttributeExtension) GetDisplayName(marketplaceAttributeName string) string {
	if tae.HasCustomName() {
		return *tae.CustomName
	}
	return marketplaceAttributeName
}

// GetAllValues combina valores marketplace con valores adicionales del tenant
func (tae *TenantAttributeExtension) GetAllValues(marketplaceValues []string) []string {
	// Combinar valores marketplace con valores adicionales, evitando duplicados
	seen := make(map[string]bool)
	allValues := make([]string, 0, len(marketplaceValues)+len(tae.AdditionalValues))

	// Agregar valores marketplace primero
	for _, value := range marketplaceValues {
		if !seen[value] {
			allValues = append(allValues, value)
			seen[value] = true
		}
	}

	// Agregar valores adicionales del tenant
	for _, value := range tae.AdditionalValues {
		if !seen[value] {
			allValues = append(allValues, value)
			seen[value] = true
		}
	}

	return allValues
}

// ValidateUniqueExtension verifica que no haya duplicados de extensión por tenant-atributo
func (tae *TenantAttributeExtension) ValidateUniqueExtension(existingExtensions []TenantAttributeExtension) error {
	for _, extension := range existingExtensions {
		if extension.ID != tae.ID &&
			extension.TenantID == tae.TenantID &&
			extension.MarketplaceAttributeID == tae.MarketplaceAttributeID {
			return errors.New("ya existe una extensión para este atributo marketplace en el tenant")
		}
	}
	return nil
}
