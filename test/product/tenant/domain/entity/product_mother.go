package entity

import (
	"time"

	"saas-mt-pim-service/src/product/tenant/domain/entity"
	"saas-mt-pim-service/src/product/tenant/domain/value_object"

	"github.com/google/uuid"
)

// ProductMother implementa el patrón Object Mother para crear entities Product de prueba
type ProductMother struct{}

// WithDefaults crea un producto con valores por defecto
func (ProductMother) WithDefaults() (*entity.Product, error) {
	// Crear value objects necesarios
	sku, _ := value_object.NewProductSKU("PROD-001")
	categoryRef, _ := value_object.NewCategoryReference(
		uuid.New().String(),
		"Categoría de Prueba",
	)
	brandRef, _ := value_object.NewBrandReference(
		uuid.New().String(),
		"Marca de Prueba",
	)

	description := "Descripción del producto de prueba"

	return entity.NewProduct(
		"tenant-123",
		"Producto de Prueba",
		&description,
		sku,
		categoryRef,
		brandRef,
	)
}

// WithTenantID crea un producto con un TenantID específico
func (p ProductMother) WithTenantID(tenantID string) (*entity.Product, error) {
	sku, _ := value_object.NewProductSKU("PROD-001")
	categoryRef, _ := value_object.NewCategoryReference(
		uuid.New().String(),
		"Categoría de Prueba",
	)
	brandRef, _ := value_object.NewBrandReference(
		uuid.New().String(),
		"Marca de Prueba",
	)

	description := "Descripción del producto de prueba"

	return entity.NewProduct(
		tenantID,
		"Producto de Prueba",
		&description,
		sku,
		categoryRef,
		brandRef,
	)
}

// WithName crea un producto con un nombre específico
func (p ProductMother) WithName(name string) (*entity.Product, error) {
	sku, _ := value_object.NewProductSKU("PROD-001")
	categoryRef, _ := value_object.NewCategoryReference(
		uuid.New().String(),
		"Categoría de Prueba",
	)
	brandRef, _ := value_object.NewBrandReference(
		uuid.New().String(),
		"Marca de Prueba",
	)

	description := "Descripción del producto de prueba"

	return entity.NewProduct(
		"tenant-123",
		name,
		&description,
		sku,
		categoryRef,
		brandRef,
	)
}

// WithSKU crea un producto con un SKU específico
func (p ProductMother) WithSKU(skuValue string) (*entity.Product, error) {
	sku, err := value_object.NewProductSKU(skuValue)
	if err != nil {
		return nil, err
	}

	categoryRef, _ := value_object.NewCategoryReference(
		uuid.New().String(),
		"Categoría de Prueba",
	)
	brandRef, _ := value_object.NewBrandReference(
		uuid.New().String(),
		"Marca de Prueba",
	)

	description := "Descripción del producto de prueba"

	return entity.NewProduct(
		"tenant-123",
		"Producto de Prueba",
		&description,
		sku,
		categoryRef,
		brandRef,
	)
}

// ActiveProduct crea un producto en estado activo
func (p ProductMother) ActiveProduct() (*entity.Product, error) {
	sku, _ := value_object.NewProductSKU("PROD-001")
	categoryRef, _ := value_object.NewCategoryReference(
		uuid.New().String(),
		"Categoría de Prueba",
	)
	brandRef, _ := value_object.NewBrandReference(
		uuid.New().String(),
		"Marca de Prueba",
	)

	description := "Descripción del producto de prueba"
	now := time.Now()

	return entity.NewProductFromRepository(
		uuid.New(),
		"tenant-123",
		"Producto de Prueba",
		&description,
		nil,
		sku,
		categoryRef,
		brandRef,
		value_object.ProductStatusActive(),
		now,
		now,
	)
}

// ElectronicsProduct crea un producto de electrónicos
func (p ProductMother) ElectronicsProduct() (*entity.Product, error) {
	sku, _ := value_object.NewProductSKU("ELEC-001")
	categoryRef, _ := value_object.NewCategoryReference(
		uuid.New().String(),
		"Electrónicos",
	)
	brandRef, _ := value_object.NewBrandReference(
		uuid.New().String(),
		"Samsung",
	)

	description := "Producto electrónico de alta calidad"

	return entity.NewProduct(
		"tenant-123",
		"Smartphone Samsung Galaxy",
		&description,
		sku,
		categoryRef,
		brandRef,
	)
}

// Create retorna una nueva instancia de ProductMother
func Create() ProductMother {
	return ProductMother{}
}
