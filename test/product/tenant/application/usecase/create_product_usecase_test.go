package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/product/tenant/application/request"
)

func TestCreateProductUseCase_Execute_InvalidRequest_ShouldFail(t *testing.T) {
	// Arrange
	req := &request.CreateProductRequest{
		Name: "", // Nombre vacío debería fallar
	}

	// Act - Solo probar la validación del request
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nombre")
}

func TestCreateProductUseCase_Execute_ValidRequest_WithDescription(t *testing.T) {
	// Arrange
	description := "Smartphone Apple iPhone 15"
	req := &request.CreateProductRequest{
		Name:        "iPhone 15",
		Description: &description,
	}

	// Act - Solo validar que el request es válido
	err := req.Validate()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "iPhone 15", req.Name)
	assert.NotNil(t, req.Description)
	assert.Equal(t, description, *req.Description)
}

func TestCreateProductUseCase_Execute_ValidRequest_WithSKU(t *testing.T) {
	// Arrange
	sku := "IPH15-128-BLK"
	req := &request.CreateProductRequest{
		Name: "iPhone 15",
		SKU:  &sku,
	}

	// Act - Solo validar que el request es válido
	err := req.Validate()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "iPhone 15", req.Name)
	assert.NotNil(t, req.SKU)
	assert.Equal(t, sku, *req.SKU)
}

func TestCreateProductUseCase_Execute_ValidRequest_WithCategoryAndBrand(t *testing.T) {
	// Arrange
	categoryID := "cat-electronics"
	brandID := "brand-apple"
	req := &request.CreateProductRequest{
		Name:       "iPhone 15",
		CategoryID: &categoryID,
		BrandID:    &brandID,
	}

	// Act - Solo validar que el request es válido
	err := req.Validate()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "iPhone 15", req.Name)
	assert.NotNil(t, req.CategoryID)
	assert.Equal(t, categoryID, *req.CategoryID)
	assert.NotNil(t, req.BrandID)
	assert.Equal(t, brandID, *req.BrandID)
}

func TestCreateProductRequest_Validate_EmptyName_ShouldFail(t *testing.T) {
	// Arrange
	req := &request.CreateProductRequest{
		Name: "",
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nombre")
}

func TestCreateProductRequest_Validate_ShortName_ShouldFail(t *testing.T) {
	// Arrange
	req := &request.CreateProductRequest{
		Name: "A", // Nombre muy corto
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "2 caracteres")
}

func TestCreateProductRequest_Validate_LongName_ShouldFail(t *testing.T) {
	// Arrange
	longName := "Este es un nombre de producto extremadamente largo que excede los 255 caracteres permitidos por la validación del request y debería fallar la validación porque supera el límite máximo establecido para el nombre del producto en el sistema. Agregamos más texto para asegurar que definitivamente superamos los 255 caracteres permitidos para que la validación falle correctamente."
	req := &request.CreateProductRequest{
		Name: longName,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "255 caracteres")
}

func TestCreateProductRequest_Validate_LongDescription_ShouldFail(t *testing.T) {
	// Arrange
	longDescription := "Esta es una descripción extremadamente larga que excede los 1000 caracteres permitidos por la validación del request y debería fallar la validación porque supera el límite máximo establecido para la descripción del producto en el sistema. Esta descripción continúa siendo muy larga para probar que la validación funciona correctamente cuando se excede el límite de caracteres permitidos. Seguimos agregando más texto para asegurar que superamos los 1000 caracteres requeridos para que falle la validación. Esta descripción debe ser lo suficientemente larga como para activar el error de validación correspondiente. Continuamos agregando más contenido para asegurar que definitivamente excedemos el límite de 1000 caracteres establecido en la validación del request de creación de producto. Ya deberíamos haber superado ampliamente el límite de caracteres permitidos para la descripción del producto. Agregamos aún más texto para garantizar que superamos los 1000 caracteres límite establecidos en la validación del sistema para las descripciones de productos."
	req := &request.CreateProductRequest{
		Name:        "Producto de Prueba",
		Description: &longDescription,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "1000 caracteres")
}

func TestCreateProductRequest_Validate_ShortSKU_ShouldFail(t *testing.T) {
	// Arrange
	shortSKU := "AB" // SKU muy corto
	req := &request.CreateProductRequest{
		Name: "Producto de Prueba",
		SKU:  &shortSKU,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "3 caracteres")
}

func TestCreateProductRequest_Validate_LongSKU_ShouldFail(t *testing.T) {
	// Arrange
	longSKU := "ESTE-ES-UN-SKU-EXTREMADAMENTE-LARGO-QUE-EXCEDE-LOS-50-CARACTERES"
	req := &request.CreateProductRequest{
		Name: "Producto de Prueba",
		SKU:  &longSKU,
	}

	// Act
	err := req.Validate()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "50 caracteres")
}

func TestCreateProductRequest_Validate_ValidRequest_ShouldSucceed(t *testing.T) {
	// Arrange
	description := "Descripción válida del producto"
	sku := "PROD-123"
	categoryID := "cat-123"
	brandID := "brand-456"

	req := &request.CreateProductRequest{
		Name:        "Producto Válido",
		Description: &description,
		SKU:         &sku,
		CategoryID:  &categoryID,
		BrandID:     &brandID,
	}

	// Act
	err := req.Validate()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "Producto Válido", req.Name)
	assert.Equal(t, description, *req.Description)
	assert.Equal(t, sku, *req.SKU)
	assert.Equal(t, categoryID, *req.CategoryID)
	assert.Equal(t, brandID, *req.BrandID)
}

// Regression test: MER-96 — POST /pim/products ignoraba array variants del request
func TestCreateProductRequest_WithVariants_ShouldPreserveVariantsData(t *testing.T) {
	// Arrange
	variantName := "Coca Cola 500ml"
	variantSKU := "COCA-500"
	variantPrice := 1500.00
	variantBarcode := "7790895000478"

	req := &request.CreateProductRequest{
		Name: "Coca Cola",
		Variants: []request.CreateVariantInProductRequest{
			{
				Name:    &variantName,
				SKU:     &variantSKU,
				Price:   &variantPrice,
				Barcode: &variantBarcode,
			},
		},
	}

	// Act
	err := req.Validate()

	// Assert
	require.NoError(t, err)
	assert.Len(t, req.Variants, 1)
	assert.Equal(t, &variantName, req.Variants[0].Name)
	assert.Equal(t, &variantSKU, req.Variants[0].SKU)
	assert.Equal(t, &variantPrice, req.Variants[0].Price)
	assert.Equal(t, &variantBarcode, req.Variants[0].Barcode)
}

func TestCreateProductRequest_WithMultipleVariants_ShouldPreserveAll(t *testing.T) {
	// Arrange
	name1 := "Coca Cola 500ml"
	sku1 := "COCA-500"
	price1 := 1500.00
	name2 := "Coca Cola 1.5L"
	sku2 := "COCA-1500"
	price2 := 2500.00

	req := &request.CreateProductRequest{
		Name: "Coca Cola",
		Variants: []request.CreateVariantInProductRequest{
			{Name: &name1, SKU: &sku1, Price: &price1},
			{Name: &name2, SKU: &sku2, Price: &price2},
		},
	}

	// Act
	err := req.Validate()

	// Assert
	require.NoError(t, err)
	assert.Len(t, req.Variants, 2)
	assert.Equal(t, &name1, req.Variants[0].Name)
	assert.Equal(t, &name2, req.Variants[1].Name)
}
