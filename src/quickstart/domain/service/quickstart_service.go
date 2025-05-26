package service

import (
	"context"
	"encoding/json"
	"fmt"

	"pim/src/quickstart/domain/entity"
	"pim/src/quickstart/domain/port"
)

// QuickstartService implementa la lógica de negocio para el quickstart
type QuickstartService struct {
	yamlLoader port.YamlDataLoader
}

// NewQuickstartService crea una nueva instancia del servicio de quickstart
func NewQuickstartService(yamlLoader port.YamlDataLoader) *QuickstartService {
	return &QuickstartService{
		yamlLoader: yamlLoader,
	}
}

// GetBusinessTypes obtiene todos los tipos de negocio disponibles
func (s *QuickstartService) GetBusinessTypes(ctx context.Context) ([]*entity.BusinessType, error) {
	return s.yamlLoader.LoadBusinessTypes(ctx)
}

// GetCategoriesByBusinessType obtiene las categorías para un tipo de negocio específico
func (s *QuickstartService) GetCategoriesByBusinessType(ctx context.Context, businessType string) (interface{}, error) {
	if businessType == "" {
		return nil, fmt.Errorf("el tipo de negocio es obligatorio")
	}

	return s.yamlLoader.LoadCategoriesByBusinessType(ctx, businessType)
}

// GetAttributesByBusinessType obtiene los atributos para un tipo de negocio específico
func (s *QuickstartService) GetAttributesByBusinessType(ctx context.Context, businessType string) (interface{}, error) {
	if businessType == "" {
		return nil, fmt.Errorf("el tipo de negocio es obligatorio")
	}

	return s.yamlLoader.LoadAttributesByBusinessType(ctx, businessType)
}

// GetVariantsByBusinessType obtiene las variantes para un tipo de negocio específico
func (s *QuickstartService) GetVariantsByBusinessType(ctx context.Context, businessType string) (interface{}, error) {
	if businessType == "" {
		return nil, fmt.Errorf("el tipo de negocio es obligatorio")
	}

	return s.yamlLoader.LoadVariantsByBusinessType(ctx, businessType)
}

// GetProductsByBusinessType obtiene los productos para un tipo de negocio específico
func (s *QuickstartService) GetProductsByBusinessType(ctx context.Context, businessType string) (interface{}, error) {
	if businessType == "" {
		return nil, fmt.Errorf("el tipo de negocio es obligatorio")
	}

	return s.yamlLoader.LoadProductsByBusinessType(ctx, businessType)
}

// GetBrandsByBusinessType obtiene las marcas para un tipo de negocio específico
func (s *QuickstartService) GetBrandsByBusinessType(ctx context.Context, businessType string) (interface{}, error) {
	if businessType == "" {
		return nil, fmt.Errorf("el tipo de negocio es obligatorio")
	}

	return s.yamlLoader.LoadBrandsByBusinessType(ctx, businessType)
}

// ValidateSetupData valida los datos de configuración del quickstart
func (s *QuickstartService) ValidateSetupData(setupData map[string]interface{}) error {
	businessType, ok := setupData["businessType"].(string)
	if !ok || businessType == "" {
		return fmt.Errorf("el tipo de negocio es obligatorio")
	}

	// Validar que las categorías seleccionadas existan
	if selectedCategories, ok := setupData["selectedCategories"].([]interface{}); ok {
		if len(selectedCategories) == 0 {
			return fmt.Errorf("debe seleccionar al menos una categoría")
		}
	}

	return nil
}

// PrepareSetupData prepara los datos de configuración para ser almacenados
func (s *QuickstartService) PrepareSetupData(setupData map[string]interface{}) (json.RawMessage, error) {
	if err := s.ValidateSetupData(setupData); err != nil {
		return nil, err
	}

	data, err := json.Marshal(setupData)
	if err != nil {
		return nil, fmt.Errorf("error al serializar los datos de configuración: %w", err)
	}

	return data, nil
}
