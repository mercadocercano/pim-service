package value_object

import (
	"encoding/json"
	"fmt"
)

// SetupData representa los datos de configuración del quickstart
type SetupData struct {
	BusinessType       string   `json:"businessType"`
	SelectedCategories []string `json:"selectedCategories"`
	SelectedAttributes []string `json:"selectedAttributes"`
	SelectedVariants   []string `json:"selectedVariants"`
	SelectedProducts   []string `json:"selectedProducts"`
}

// NewSetupData crea una nueva instancia de SetupData con validaciones
func NewSetupData(businessType string, selectedCategories, selectedAttributes, selectedVariants, selectedProducts []string) (*SetupData, error) {
	if businessType == "" {
		return nil, fmt.Errorf("el tipo de negocio es obligatorio")
	}

	if len(selectedCategories) == 0 {
		return nil, fmt.Errorf("debe seleccionar al menos una categoría")
	}

	return &SetupData{
		BusinessType:       businessType,
		SelectedCategories: selectedCategories,
		SelectedAttributes: selectedAttributes,
		SelectedVariants:   selectedVariants,
		SelectedProducts:   selectedProducts,
	}, nil
}

// ToJSON convierte SetupData a JSON
func (sd *SetupData) ToJSON() (json.RawMessage, error) {
	data, err := json.Marshal(sd)
	if err != nil {
		return nil, fmt.Errorf("error al serializar SetupData: %w", err)
	}
	return data, nil
}

// FromJSON crea SetupData desde JSON
func FromJSON(data json.RawMessage) (*SetupData, error) {
	var setupData SetupData
	if err := json.Unmarshal(data, &setupData); err != nil {
		return nil, fmt.Errorf("error al deserializar SetupData: %w", err)
	}

	// Validar los datos deserializados
	if setupData.BusinessType == "" {
		return nil, fmt.Errorf("el tipo de negocio es obligatorio")
	}

	if len(setupData.SelectedCategories) == 0 {
		return nil, fmt.Errorf("debe seleccionar al menos una categoría")
	}

	return &setupData, nil
}

// HasCategories verifica si hay categorías seleccionadas
func (sd *SetupData) HasCategories() bool {
	return len(sd.SelectedCategories) > 0
}

// HasAttributes verifica si hay atributos seleccionados
func (sd *SetupData) HasAttributes() bool {
	return len(sd.SelectedAttributes) > 0
}

// HasVariants verifica si hay variantes seleccionadas
func (sd *SetupData) HasVariants() bool {
	return len(sd.SelectedVariants) > 0
}

// HasProducts verifica si hay productos seleccionados
func (sd *SetupData) HasProducts() bool {
	return len(sd.SelectedProducts) > 0
}
