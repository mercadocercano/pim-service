package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/product/tenant/domain/entity"
	"saas-mt-pim-service/src/product/tenant/domain/port"
	"saas-mt-pim-service/src/product/tenant/domain/value_object"
)

// CreateFromTemplateUseCase implementa la creación de productos desde templates del global catalog
type CreateFromTemplateUseCase struct {
	productRepo port.ProductRepository
}

// NewCreateFromTemplateUseCase crea una nueva instancia del caso de uso
func NewCreateFromTemplateUseCase(productRepo port.ProductRepository) *CreateFromTemplateUseCase {
	return &CreateFromTemplateUseCase{
		productRepo: productRepo,
	}
}

// Execute ejecuta la creación de productos desde templates del global catalog
func (uc *CreateFromTemplateUseCase) Execute(ctx context.Context, tenantID string, templateData interface{}) error {
	// Convertir templateData a slice de productos del global catalog
	globalProducts, ok := templateData.([]map[string]interface{})
	if !ok {
		return fmt.Errorf("templateData debe ser un slice de productos del global catalog")
	}

	// Procesar cada producto template
	for _, globalProductData := range globalProducts {
		if err := uc.createProductFromTemplate(ctx, tenantID, globalProductData); err != nil {
			return fmt.Errorf("error creando producto desde template: %w", err)
		}
	}

	return nil
}

// createProductFromTemplate crea un producto del tenant basado en un template del global catalog
func (uc *CreateFromTemplateUseCase) createProductFromTemplate(ctx context.Context, tenantID string, templateData map[string]interface{}) error {
	// Extraer datos del template
	name, _ := templateData["name"].(string)
	description, _ := templateData["description"].(string)
	ean, _ := templateData["ean"].(string)
	brandName, _ := templateData["brand"].(string)
	categoryName, _ := templateData["category"].(string)

	if name == "" {
		return fmt.Errorf("el nombre del producto es obligatorio en el template")
	}

	// Crear SKU único basado en EAN y tenant
	var productSKU *value_object.ProductSKU
	if ean != "" {
		skuValue := fmt.Sprintf("%s-%s", tenantID, ean)
		sku, err := value_object.NewProductSKU(skuValue)
		if err != nil {
			return fmt.Errorf("error creando SKU: %w", err)
		}
		productSKU = sku
	}

	// Crear referencia de categoría (temporal con nombre, sin ID por ahora)
	var categoryRef *value_object.CategoryReference
	if categoryName != "" {
		// TODO: En el futuro, buscar la categoría por nombre y obtener su ID
		// Por ahora creamos una referencia temporal
		ref, err := value_object.NewCategoryReference("temp-category-id", categoryName)
		if err != nil {
			return fmt.Errorf("error creando referencia de categoría: %w", err)
		}
		categoryRef = ref
	}

	// Crear referencia de marca (temporal con nombre, sin ID por ahora)
	var brandRef *value_object.BrandReference
	if brandName != "" {
		// TODO: En el futuro, buscar la marca por nombre y obtener su ID
		// Por ahora creamos una referencia temporal
		ref, err := value_object.NewBrandReference("temp-brand-id", brandName)
		if err != nil {
			return fmt.Errorf("error creando referencia de marca: %w", err)
		}
		brandRef = ref
	}

	// Preparar descripción como puntero
	var descPtr *string
	if description != "" {
		descPtr = &description
	}

	// Crear la entidad producto
	product, err := entity.NewProduct(
		tenantID,
		name,
		descPtr,
		productSKU,
		categoryRef,
		brandRef,
	)
	if err != nil {
		return fmt.Errorf("error creando entidad producto: %w", err)
	}

	// TODO: Agregar referencia al global catalog cuando se implemente en la entidad
	// if globalCatalogID != "" {
	//     product.SetGlobalCatalogReference(globalCatalogID)
	// }

	// Guardar en el repositorio
	if err := uc.productRepo.Save(ctx, product); err != nil {
		return fmt.Errorf("error guardando producto: %w", err)
	}

	return nil
}
