package loader

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"pim/src/quickstart/domain/entity"
	"pim/src/quickstart/domain/exception"

	"gopkg.in/yaml.v3"
)

// FileYamlDataLoader implementa YamlDataLoader cargando datos desde archivos YAML
type FileYamlDataLoader struct {
	dataPath string
	cache    map[string]interface{}
	mutex    sync.RWMutex
}

// NewFileYamlDataLoader crea una nueva instancia del cargador de archivos YAML
func NewFileYamlDataLoader(dataPath string) *FileYamlDataLoader {
	return &FileYamlDataLoader{
		dataPath: dataPath,
		cache:    make(map[string]interface{}),
	}
}

// LoadBusinessTypes carga todos los tipos de negocio desde business-types.yaml
func (loader *FileYamlDataLoader) LoadBusinessTypes(ctx context.Context) ([]*entity.BusinessType, error) {
	cacheKey := "business-types"

	// Verificar caché
	loader.mutex.RLock()
	if cached, exists := loader.cache[cacheKey]; exists {
		loader.mutex.RUnlock()
		return cached.([]*entity.BusinessType), nil
	}
	loader.mutex.RUnlock()

	// Cargar desde archivo
	filePath := filepath.Join(loader.dataPath, "business-types.yaml")
	data, err := loader.loadYamlFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error cargando tipos de negocio: %w", err)
	}

	// Parsear estructura específica
	var businessTypesData struct {
		BusinessTypes []struct {
			ID          string `yaml:"id"`
			Name        string `yaml:"name"`
			Description string `yaml:"description"`
			Icon        string `yaml:"icon"`
		} `yaml:"business_types"`
	}

	if err := yaml.Unmarshal(data, &businessTypesData); err != nil {
		return nil, fmt.Errorf("error parseando tipos de negocio: %w", err)
	}

	// Convertir a entidades
	businessTypes := make([]*entity.BusinessType, len(businessTypesData.BusinessTypes))
	for i, bt := range businessTypesData.BusinessTypes {
		businessType, err := entity.NewBusinessType(bt.ID, bt.Name, bt.Description, bt.Icon)
		if err != nil {
			return nil, fmt.Errorf("error creando tipo de negocio %s: %w", bt.ID, err)
		}
		businessTypes[i] = businessType
	}

	// Guardar en caché
	loader.mutex.Lock()
	loader.cache[cacheKey] = businessTypes
	loader.mutex.Unlock()

	return businessTypes, nil
}

// LoadCategoriesByBusinessType carga las categorías para un tipo de negocio específico
func (loader *FileYamlDataLoader) LoadCategoriesByBusinessType(ctx context.Context, businessType string) (interface{}, error) {
	cacheKey := fmt.Sprintf("categories-%s", businessType)

	// Verificar caché
	loader.mutex.RLock()
	if cached, exists := loader.cache[cacheKey]; exists {
		loader.mutex.RUnlock()
		return cached, nil
	}
	loader.mutex.RUnlock()

	// Cargar desde archivo
	filePath := filepath.Join(loader.dataPath, "categories", fmt.Sprintf("%s.yaml", businessType))
	data, err := loader.loadYamlFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error cargando categorías para %s: %w", businessType, err)
	}

	// Parsear como interfaz genérica para flexibilidad
	var categoriesData interface{}
	if err := yaml.Unmarshal(data, &categoriesData); err != nil {
		return nil, fmt.Errorf("error parseando categorías para %s: %w", businessType, err)
	}

	// Guardar en caché
	loader.mutex.Lock()
	loader.cache[cacheKey] = categoriesData
	loader.mutex.Unlock()

	return categoriesData, nil
}

// LoadAttributesByBusinessType carga los atributos para un tipo de negocio específico
func (loader *FileYamlDataLoader) LoadAttributesByBusinessType(ctx context.Context, businessType string) (interface{}, error) {
	cacheKey := fmt.Sprintf("attributes-%s", businessType)

	// Verificar caché
	loader.mutex.RLock()
	if cached, exists := loader.cache[cacheKey]; exists {
		loader.mutex.RUnlock()
		return cached, nil
	}
	loader.mutex.RUnlock()

	// Cargar desde archivo
	filePath := filepath.Join(loader.dataPath, "attributes", fmt.Sprintf("%s.yaml", businessType))
	data, err := loader.loadYamlFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error cargando atributos para %s: %w", businessType, err)
	}

	// Parsear como interfaz genérica
	var attributesData interface{}
	if err := yaml.Unmarshal(data, &attributesData); err != nil {
		return nil, fmt.Errorf("error parseando atributos para %s: %w", businessType, err)
	}

	// Guardar en caché
	loader.mutex.Lock()
	loader.cache[cacheKey] = attributesData
	loader.mutex.Unlock()

	return attributesData, nil
}

// LoadVariantsByBusinessType carga las variantes para un tipo de negocio específico
func (loader *FileYamlDataLoader) LoadVariantsByBusinessType(ctx context.Context, businessType string) (interface{}, error) {
	cacheKey := fmt.Sprintf("variants-%s", businessType)

	// Verificar caché
	loader.mutex.RLock()
	if cached, exists := loader.cache[cacheKey]; exists {
		loader.mutex.RUnlock()
		return cached, nil
	}
	loader.mutex.RUnlock()

	// Cargar desde archivo
	filePath := filepath.Join(loader.dataPath, "variants", fmt.Sprintf("%s.yaml", businessType))
	data, err := loader.loadYamlFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error cargando variantes para %s: %w", businessType, err)
	}

	// Parsear como interfaz genérica
	var variantsData interface{}
	if err := yaml.Unmarshal(data, &variantsData); err != nil {
		return nil, fmt.Errorf("error parseando variantes para %s: %w", businessType, err)
	}

	// Guardar en caché
	loader.mutex.Lock()
	loader.cache[cacheKey] = variantsData
	loader.mutex.Unlock()

	return variantsData, nil
}

// LoadProductsByBusinessType carga los productos para un tipo de negocio específico
func (loader *FileYamlDataLoader) LoadProductsByBusinessType(ctx context.Context, businessType string) (interface{}, error) {
	cacheKey := fmt.Sprintf("products-%s", businessType)

	// Verificar caché
	loader.mutex.RLock()
	if cached, exists := loader.cache[cacheKey]; exists {
		loader.mutex.RUnlock()
		return cached, nil
	}
	loader.mutex.RUnlock()

	// Cargar desde archivo
	filePath := filepath.Join(loader.dataPath, "products", fmt.Sprintf("%s.yaml", businessType))
	data, err := loader.loadYamlFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error cargando productos para %s: %w", businessType, err)
	}

	// Parsear como interfaz genérica
	var productsData interface{}
	if err := yaml.Unmarshal(data, &productsData); err != nil {
		return nil, fmt.Errorf("error parseando productos para %s: %w", businessType, err)
	}

	// Guardar en caché
	loader.mutex.Lock()
	loader.cache[cacheKey] = productsData
	loader.mutex.Unlock()

	return productsData, nil
}

// LoadBrandsByBusinessType carga las marcas para un tipo de negocio específico
func (loader *FileYamlDataLoader) LoadBrandsByBusinessType(ctx context.Context, businessType string) (interface{}, error) {
	cacheKey := fmt.Sprintf("brands-%s", businessType)

	// Verificar caché
	loader.mutex.RLock()
	if cached, exists := loader.cache[cacheKey]; exists {
		loader.mutex.RUnlock()
		return cached, nil
	}
	loader.mutex.RUnlock()

	// Cargar desde archivo
	filePath := filepath.Join(loader.dataPath, "brands", fmt.Sprintf("%s.yaml", businessType))
	data, err := loader.loadYamlFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error cargando marcas para %s: %w", businessType, err)
	}

	// Parsear como interfaz genérica
	var brandsData interface{}
	if err := yaml.Unmarshal(data, &brandsData); err != nil {
		return nil, fmt.Errorf("error parseando marcas para %s: %w", businessType, err)
	}

	// Guardar en caché
	loader.mutex.Lock()
	loader.cache[cacheKey] = brandsData
	loader.mutex.Unlock()

	return brandsData, nil
}

// loadYamlFile carga un archivo YAML y retorna los datos como bytes
func (loader *FileYamlDataLoader) loadYamlFile(filePath string) ([]byte, error) {
	// Verificar si el archivo existe
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, exception.ErrYamlFileNotFound
	}

	// Leer el archivo
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error leyendo archivo %s: %w", filePath, err)
	}

	return data, nil
}

// ClearCache limpia toda la caché
func (loader *FileYamlDataLoader) ClearCache() {
	loader.mutex.Lock()
	defer loader.mutex.Unlock()
	loader.cache = make(map[string]interface{})
}

// ClearCacheForBusinessType limpia la caché para un tipo de negocio específico
func (loader *FileYamlDataLoader) ClearCacheForBusinessType(businessType string) {
	loader.mutex.Lock()
	defer loader.mutex.Unlock()

	keys := []string{
		fmt.Sprintf("categories-%s", businessType),
		fmt.Sprintf("attributes-%s", businessType),
		fmt.Sprintf("variants-%s", businessType),
		fmt.Sprintf("products-%s", businessType),
	}

	for _, key := range keys {
		delete(loader.cache, key)
	}
}
