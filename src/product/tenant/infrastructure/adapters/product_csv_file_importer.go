package adapters

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"saas-mt-pim-service/src/product/tenant/domain/entity"
	"saas-mt-pim-service/src/product/tenant/domain/value_object"
	"saas-mt-pim-service/src/shared/domain/port"
	"saas-mt-pim-service/src/shared/infrastructure/adapters"

	"github.com/google/uuid"
)

// ProductCSVFileImporter implementa la importación de productos desde CSV
type ProductCSVFileImporter struct {
	*adapters.BaseCSVFileImporter[entity.Product]
	
	// categoryService para validar categorías
	categoryService CategoryValidator
	
	// brandService para validar marcas
	brandService BrandValidator
}

// CategoryValidator interfaz para validar categorías
type CategoryValidator interface {
	ValidateCategory(ctx context.Context, categoryID string, tenantID string) (bool, error)
}

// BrandValidator interfaz para validar marcas
type BrandValidator interface {
	ValidateBrand(ctx context.Context, brandID string, tenantID string) (bool, error)
}

// NewProductCSVFileImporter crea un nuevo importador de productos CSV
func NewProductCSVFileImporter(categoryService CategoryValidator, brandService BrandValidator) *ProductCSVFileImporter {
	requiredColumns := []string{"name", "sku", "price"}
	base := adapters.NewBaseCSVFileImporter[entity.Product](',', true, requiredColumns)
	
	return &ProductCSVFileImporter{
		BaseCSVFileImporter: base,
		categoryService:     categoryService,
		brandService:        brandService,
	}
}

// Import implementa la interfaz FileImporter
func (p *ProductCSVFileImporter) Import(ctx context.Context, reader io.Reader, tenantID string) (*port.ImportResult[entity.Product], error) {
	return p.BaseCSVFileImporter.Import(ctx, reader, tenantID, p)
}

// ParseRow implementa la interfaz RowParser para productos
func (p *ProductCSVFileImporter) ParseRow(row []string, headers []string, rowData map[string]string, tenantID string) (*entity.Product, []string) {
	errors := []string{}
	
	// Validar campos requeridos
	name := rowData["name"]
	if name == "" {
		errors = append(errors, "nombre es requerido")
	}
	
	sku := rowData["sku"]
	if sku == "" {
		errors = append(errors, "SKU es requerido")
	}
	
	// Validar y parsear precio
	var price float64
	if priceStr := rowData["price"]; priceStr != "" {
		parsedPrice, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			errors = append(errors, fmt.Sprintf("precio inválido: %s", priceStr))
		} else if parsedPrice < 0 {
			errors = append(errors, "el precio no puede ser negativo")
		} else {
			price = parsedPrice
		}
	} else {
		errors = append(errors, "precio es requerido")
	}
	
	// Si hay errores hasta aquí, retornar
	if len(errors) > 0 {
		return nil, errors
	}
	
	// Preparar descripción
	var description *string
	if desc := rowData["description"]; desc != "" {
		description = &desc
	}
	
	// Preparar SKU
	var productSKU *value_object.ProductSKU
	if sku != "" {
		skuObj, err := value_object.NewProductSKU(sku)
		if err != nil {
			errors = append(errors, fmt.Sprintf("SKU inválido: %v", err))
		} else {
			productSKU = skuObj
		}
	}
	
	// Preparar categoría si existe
	var categoryRef *value_object.CategoryReference
	if categoryID := rowData["category_id"]; categoryID != "" {
		// TODO: Implementar validación real con categoryService
		// Por ahora, solo validamos el formato UUID
		if _, err := uuid.Parse(categoryID); err != nil {
			errors = append(errors, fmt.Sprintf("category_id inválido: %s", categoryID))
		} else {
			categoryName := rowData["category_name"]
			if categoryName == "" {
				categoryName = "Sin categoría" // Valor por defecto
			}
			categoryRef, _ = value_object.NewCategoryReference(categoryID, categoryName)
		}
	}
	
	// Preparar marca si existe
	var brandRef *value_object.BrandReference
	if brandID := rowData["brand_id"]; brandID != "" {
		// TODO: Implementar validación real con brandService
		// Por ahora, solo validamos el formato UUID
		if _, err := uuid.Parse(brandID); err != nil {
			errors = append(errors, fmt.Sprintf("brand_id inválido: %s", brandID))
		} else {
			brandName := rowData["brand_name"]
			if brandName == "" {
				brandName = "Sin marca" // Valor por defecto
			}
			brandRef, _ = value_object.NewBrandReference(brandID, brandName)
		}
	}
	
	// Crear el producto usando el constructor
	product, err := entity.NewProduct(
		tenantID,
		name,
		description,
		productSKU,
		categoryRef,
		brandRef,
	)
	if err != nil {
		errors = append(errors, fmt.Sprintf("Error al crear producto: %v", err))
		return nil, errors
	}
	
	// Parsear atributos adicionales si existen
	// Nota: Los atributos adicionales del CSV se pueden usar para crear variantes con atributos
	// pero el producto en sí no tiene un campo de atributos genéricos
	attributes := p.parseAttributes(rowData)
	
	// Si tenemos un precio válido, crear una variante por defecto
	if price > 0 && len(errors) == 0 {
		// Convertir los atributos del CSV a VariantAttributes
		variantAttributes := make([]*value_object.VariantAttribute, 0)
		for key, value := range attributes {
			if strValue, ok := value.(string); ok {
				attr, err := value_object.NewVariantAttribute(key, strValue)
				if err == nil {
					variantAttributes = append(variantAttributes, attr)
				}
			}
		}
		
		// Crear la colección de atributos
		attrCollection, err := value_object.NewVariantAttributeCollection(variantAttributes)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Error al crear colección de atributos: %v", err))
		}
		
		// Crear la variante usando el constructor apropiado
		variant, err := entity.NewProductVariant(
			tenantID,
			product.ID(),
			"Variante principal",
			productSKU, // Usar el mismo SKU del producto
			true,       // Es la variante por defecto
			1,          // Primer orden
			attrCollection,
		)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Error al crear variante: %v", err))
		}
		
		// El producto necesita que se le agregue la variante usando AddVariant
		// Pero eso se debe hacer después de guardar el producto en el repositorio
		// Por ahora solo validamos que se puede crear
		_ = variant
		
		// Parsear stock si existe (se manejará en el servicio de stock)
		if stockStr := rowData["stock"]; stockStr != "" {
			stock, err := strconv.Atoi(stockStr)
			if err != nil {
				errors = append(errors, fmt.Sprintf("stock inválido: %s", stockStr))
			} else if stock < 0 {
				errors = append(errors, "el stock no puede ser negativo")
			}
			// El stock se manejará en el servicio de stock, no en la variante
		}
	}
	
	if len(errors) > 0 {
		return nil, errors
	}
	
	return product, nil
}

// parseAttributes extrae atributos adicionales del CSV
func (p *ProductCSVFileImporter) parseAttributes(rowData map[string]string) map[string]interface{} {
	// Columnas conocidas que no son atributos
	knownColumns := map[string]bool{
		"name":          true,
		"sku":           true,
		"description":   true,
		"price":         true,
		"stock":         true,
		"category_id":   true,
		"category_name": true,
		"brand_id":      true,
		"brand_name":    true,
	}
	
	attributes := make(map[string]interface{})
	
	// Cualquier columna que no sea conocida se considera un atributo
	for key, value := range rowData {
		lowerKey := strings.ToLower(strings.TrimSpace(key))
		if !knownColumns[lowerKey] && value != "" {
			attributes[key] = value
		}
	}
	
	return attributes
}