package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"saas-mt-pim-service/src/product/global_catalog/application/usecase"
	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

func TestSearchByEAN_ValidEAN_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewSearchByEAN(mockRepo)

	request := usecase.SearchByEANRequest{
		EAN:        "1234567890128",
		OnlyActive: false,
	}

	// Crear un producto mock
	ean, _ := value_object.NewEAN13("1234567890128")
	source, _ := value_object.NewManualSource()
	mockProduct, _ := entity.NewGlobalProduct(
		ean,
		"Test Product",
		stringPtr("Test Description"),
		stringPtr("Test Brand"),
		stringPtr("Test Category"),
		float64Ptr(100.0),
		stringPtr("https://test.com/image.jpg"),
		source,
	)

	mockRepo.On("FindByEAN", "1234567890128").Return(mockProduct, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "1234567890128", result.EAN)
	assert.Equal(t, "Test Product", result.Name)
	assert.Equal(t, "Test Description", *result.Description)
	assert.Equal(t, "Test Brand", *result.Brand)
	assert.Equal(t, "Test Category", *result.Category)
	assert.Equal(t, 100.0, *result.Price)
	assert.Equal(t, "https://test.com/image.jpg", *result.ImageURL)
	assert.Equal(t, "manual", result.Source)
	assert.False(t, result.IsVerified)
	assert.True(t, result.IsActive)

	mockRepo.AssertExpectations(t)
}

func TestSearchByEAN_OnlyActive_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewSearchByEAN(mockRepo)

	request := usecase.SearchByEANRequest{
		EAN:        "9780201616224",
		OnlyActive: true,
	}

	// Crear un producto mock
	ean, _ := value_object.NewEAN13("9780201616224")
	source, _ := value_object.NewManualSource()
	mockProduct, _ := entity.NewGlobalProduct(
		ean,
		"Test Product",
		nil,
		nil,
		nil,
		nil,
		nil,
		source,
	)

	mockRepo.On("FindActiveByEAN", "9780201616224").Return(mockProduct, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "9780201616224", result.EAN)
	assert.Equal(t, "Test Product", result.Name)
	assert.True(t, result.IsActive)

	mockRepo.AssertExpectations(t)
}

func TestSearchByEAN_InvalidEAN_ReturnsError(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewSearchByEAN(mockRepo)

	request := usecase.SearchByEANRequest{
		EAN:        "invalid_ean",
		OnlyActive: false,
	}

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "EAN inválido")
}

func TestSearchByEAN_ProductNotFound_ReturnsError(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewSearchByEAN(mockRepo)

	request := usecase.SearchByEANRequest{
		EAN:        "4006381333931",
		OnlyActive: false,
	}

	mockRepo.On("FindByEAN", "4006381333931").Return(nil, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "4006381333931") // Should contain the EAN in error message

	mockRepo.AssertExpectations(t)
}

func TestSearchByEAN_RepositoryError_ReturnsError(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewSearchByEAN(mockRepo)

	request := usecase.SearchByEANRequest{
		EAN:        "3614272049017",
		OnlyActive: false,
	}

	mockRepo.On("FindByEAN", "3614272049017").Return(nil, errors.New("database error"))

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Error al buscar producto por EAN")

	mockRepo.AssertExpectations(t)
}

func TestSearchByEAN_OnlyActiveRepositoryError_ReturnsError(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewSearchByEAN(mockRepo)

	request := usecase.SearchByEANRequest{
		EAN:        "8711000036716",
		OnlyActive: true,
	}

	mockRepo.On("FindActiveByEAN", "8711000036716").Return(nil, errors.New("database error"))

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Error al buscar producto por EAN")

	mockRepo.AssertExpectations(t)
}

func TestSearchByEAN_WithFullProductData_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewSearchByEAN(mockRepo)

	request := usecase.SearchByEANRequest{
		EAN:        "7891000100103",
		OnlyActive: false,
	}

	// Crear un producto mock con todos los datos
	ean, _ := value_object.NewEAN13("7891000100103")
	source, _ := value_object.NewScrapingSource("disco", "https://test.com/product", 0.9)
	mockProduct, _ := entity.NewGlobalProduct(
		ean,
		"Complete Test Product",
		stringPtr("Complete Test Description"),
		stringPtr("Test Brand"),
		stringPtr("Test Category"),
		float64Ptr(250.50),
		stringPtr("https://test.com/image.jpg"),
		source,
	)

	// Agregar datos adicionales
	mockProduct.SetBusinessType("retail")
	mockProduct.AddTag("electronics")
	mockProduct.AddTag("premium")
	mockProduct.SetMetadata("warranty", "2 years")
	mockProduct.SetMetadata("origin", "Argentina")

	mockRepo.On("FindByEAN", "7891000100103").Return(mockProduct, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "7891000100103", result.EAN)
	assert.Equal(t, "Complete Test Product", result.Name)
	assert.Equal(t, "Complete Test Description", *result.Description)
	assert.Equal(t, "Test Brand", *result.Brand)
	assert.Equal(t, "Test Category", *result.Category)
	assert.Equal(t, 250.50, *result.Price)
	assert.Equal(t, "https://test.com/image.jpg", *result.ImageURL)
	assert.Equal(t, "disco", result.Source)
	assert.Equal(t, "https://test.com/product", *result.SourceURL)
	assert.Equal(t, "retail", *result.BusinessType)
	assert.Contains(t, result.Tags, "electronics")
	assert.Contains(t, result.Tags, "premium")
	assert.Equal(t, "2 years", result.Metadata["warranty"])
	assert.Equal(t, "Argentina", result.Metadata["origin"])
	assert.True(t, result.IsArgentineProduct)
	assert.NotEmpty(t, result.QualityLevel)
	assert.NotEmpty(t, result.SourceDisplayName)
	assert.NotEmpty(t, result.CreatedAt)
	assert.NotEmpty(t, result.UpdatedAt)

	mockRepo.AssertExpectations(t)
}
