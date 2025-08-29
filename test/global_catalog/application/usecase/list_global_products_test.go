package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"saas-mt-pim-service/src/product/global_catalog/application/usecase"
	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

func TestListGlobalProducts_DefaultRequest_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewListGlobalProducts(mockRepo)

	request := usecase.ListGlobalProductsRequest{
		Offset: 0,
		Limit:  20,
	}

	// Crear productos mock
	products := createMockProducts(3)

	mockRepo.On("FindAll", 0, 20).Return(products, nil)
	mockRepo.On("CountTotal").Return(100, nil)
	mockRepo.On("CountArgentineProducts").Return(30, nil)
	mockRepo.On("CountByQualityScore", 70).Return(50, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Products, 3)
	assert.NotNil(t, result.Pagination)
	assert.NotNil(t, result.Summary)
	assert.Equal(t, 0, result.Pagination.Offset)
	assert.Equal(t, 20, result.Pagination.Limit)
	assert.Equal(t, 100, result.Pagination.Total)
	assert.Equal(t, 5, result.Pagination.TotalPages)
	assert.True(t, result.Pagination.HasNext)
	assert.False(t, result.Pagination.HasPrev)

	mockRepo.AssertExpectations(t)
}

func TestListGlobalProducts_WithFilters_BusinessType_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewListGlobalProducts(mockRepo)

	businessType := "retail"
	request := usecase.ListGlobalProductsRequest{
		BusinessType: &businessType,
		Limit:        20,
	}

	products := createMockProducts(2)

	mockRepo.On("FindByBusinessType", "retail", 20).Return(products, nil)
	mockRepo.On("CountTotal").Return(50, nil)
	mockRepo.On("CountArgentineProducts").Return(15, nil)
	mockRepo.On("CountByQualityScore", 70).Return(25, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Products, 2)
	assert.Equal(t, 50, result.Summary.TotalProducts)

	mockRepo.AssertExpectations(t)
}

func TestListGlobalProducts_WithFilters_Source_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewListGlobalProducts(mockRepo)

	source := "test_source"
	request := usecase.ListGlobalProductsRequest{
		Source: &source,
		Limit:  10,
	}

	products := createMockProducts(1)

	mockRepo.On("FindBySource", "test_source", 10).Return(products, nil)
	mockRepo.On("CountTotal").Return(25, nil)
	mockRepo.On("CountArgentineProducts").Return(10, nil)
	mockRepo.On("CountByQualityScore", 70).Return(15, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Products, 1)

	mockRepo.AssertExpectations(t)
}

func TestListGlobalProducts_WithQualityRange_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewListGlobalProducts(mockRepo)

	minQuality := 70
	maxQuality := 90
	request := usecase.ListGlobalProductsRequest{
		MinQuality: &minQuality,
		MaxQuality: &maxQuality,
		Limit:      15,
	}

	products := createMockProducts(2)

	mockRepo.On("FindByQualityScoreRange", 70, 90, 15).Return(products, nil)
	mockRepo.On("CountTotal").Return(80, nil)
	mockRepo.On("CountArgentineProducts").Return(20, nil)
	mockRepo.On("CountByQualityScore", 70).Return(40, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Products, 2)

	mockRepo.AssertExpectations(t)
}

func TestListGlobalProducts_SearchByName_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewListGlobalProducts(mockRepo)

	searchName := "iPhone"
	request := usecase.ListGlobalProductsRequest{
		SearchName: &searchName,
		Limit:      20,
	}

	products := createMockProducts(1)

	mockRepo.On("SearchByName", "iPhone", 20).Return(products, nil)
	mockRepo.On("CountTotal").Return(30, nil)
	mockRepo.On("CountArgentineProducts").Return(5, nil)
	mockRepo.On("CountByQualityScore", 70).Return(20, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Products, 1)

	mockRepo.AssertExpectations(t)
}

func TestListGlobalProducts_OnlyArgentine_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewListGlobalProducts(mockRepo)

	request := usecase.ListGlobalProductsRequest{
		OnlyArgentine: true,
		Offset:        0,
		Limit:         10,
	}

	products := createMockProducts(2)

	mockRepo.On("FindArgentineProducts", 0, 10).Return(products, nil)
	mockRepo.On("CountTotal").Return(50, nil)
	mockRepo.On("CountArgentineProducts").Return(25, nil)
	mockRepo.On("CountByQualityScore", 70).Return(30, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Products, 2)

	mockRepo.AssertExpectations(t)
}

func TestListGlobalProducts_OnlyVerified_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewListGlobalProducts(mockRepo)

	request := usecase.ListGlobalProductsRequest{
		OnlyVerified: true,
		Offset:       10,
		Limit:        20,
	}

	products := createMockProducts(1)

	mockRepo.On("FindVerified", 10, 20).Return(products, nil)
	mockRepo.On("CountTotal").Return(60, nil)
	mockRepo.On("CountArgentineProducts").Return(20, nil)
	mockRepo.On("CountByQualityScore", 70).Return(35, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Products, 1)
	assert.Equal(t, 10, result.Pagination.Offset)
	assert.True(t, result.Pagination.HasPrev)

	mockRepo.AssertExpectations(t)
}

func TestListGlobalProducts_NegativeOffset_UsesZero(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewListGlobalProducts(mockRepo)

	request := usecase.ListGlobalProductsRequest{
		Offset: -5,
		Limit:  20,
	}

	products := createMockProducts(1)

	mockRepo.On("FindAll", 0, 20).Return(products, nil) // Should use 0, not -5
	mockRepo.On("CountTotal").Return(20, nil)
	mockRepo.On("CountArgentineProducts").Return(5, nil)
	mockRepo.On("CountByQualityScore", 70).Return(10, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 0, result.Pagination.Offset)

	mockRepo.AssertExpectations(t)
}

func TestListGlobalProducts_InvalidLimit_UsesDefault(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewListGlobalProducts(mockRepo)

	request := usecase.ListGlobalProductsRequest{
		Offset: 0,
		Limit:  200, // Too high, should use default
	}

	products := createMockProducts(1)

	mockRepo.On("FindAll", 0, 20).Return(products, nil) // Should use 20, not 200
	mockRepo.On("CountTotal").Return(40, nil)
	mockRepo.On("CountArgentineProducts").Return(10, nil)
	mockRepo.On("CountByQualityScore", 70).Return(20, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 20, result.Pagination.Limit)

	mockRepo.AssertExpectations(t)
}

func TestListGlobalProducts_RepositoryError_ReturnsError(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewListGlobalProducts(mockRepo)

	request := usecase.ListGlobalProductsRequest{
		Offset: 0,
		Limit:  20,
	}

	mockRepo.On("FindAll", 0, 20).Return([]*entity.GlobalProduct(nil), errors.New("database error"))

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Error al listar productos")

	mockRepo.AssertExpectations(t)
}

func TestListGlobalProducts_EmptyResult_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewListGlobalProducts(mockRepo)

	request := usecase.ListGlobalProductsRequest{
		Offset: 0,
		Limit:  20,
	}

	emptyProducts := []*entity.GlobalProduct{}

	mockRepo.On("FindAll", 0, 20).Return(emptyProducts, nil)
	mockRepo.On("CountTotal").Return(0, nil)
	mockRepo.On("CountArgentineProducts").Return(0, nil)
	mockRepo.On("CountByQualityScore", 70).Return(0, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Products, 0)
	assert.Equal(t, 0, result.Summary.TotalProducts)
	assert.Equal(t, 0, result.Summary.AverageQuality)

	mockRepo.AssertExpectations(t)
}

// Helper function to create mock products
func createMockProducts(count int) []*entity.GlobalProduct {
	products := make([]*entity.GlobalProduct, count)
	validEANs := []string{"1234567890128", "9780201616224", "4006381333931", "3614272049017", "8711000036716"}

	for i := 0; i < count; i++ {
		eanCode := validEANs[i%len(validEANs)]
		ean, _ := value_object.NewEAN13(eanCode)
		source, _ := value_object.NewManualSource()
		product, _ := entity.NewGlobalProduct(
			ean,
			fmt.Sprintf("Test Product %d", i+1),
			stringPtr(fmt.Sprintf("Description %d", i+1)),
			stringPtr("Test Brand"),
			stringPtr("Test Category"),
			float64Ptr(100.0+float64(i*10)),
			stringPtr(fmt.Sprintf("https://test.com/image%d.jpg", i+1)),
			source,
		)
		products[i] = product
	}

	return products
}
