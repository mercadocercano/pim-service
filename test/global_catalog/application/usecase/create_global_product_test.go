package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pim/src/global_catalog/application/usecase"
	"pim/src/global_catalog/domain/entity"
	"pim/src/global_catalog/domain/value_object"
	"pim/src/shared/domain/criteria"
)

// MockGlobalProductRepository es un mock del repositorio
type MockGlobalProductRepository struct {
	mock.Mock
}

func (m *MockGlobalProductRepository) FindByEAN(ean string) (*entity.GlobalProduct, error) {
	args := m.Called(ean)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) Save(product *entity.GlobalProduct) (*entity.GlobalProduct, error) {
	args := m.Called(product)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) FindActiveByEAN(ean string) (*entity.GlobalProduct, error) {
	args := m.Called(ean)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) FindAll(offset, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) FindByBusinessType(businessType string, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(businessType, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) FindBySource(source string, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(source, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) FindByQualityScoreRange(min, max, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(min, max, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) SearchByName(name string, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(name, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) SearchByBrand(brand string, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(brand, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) SearchByCategory(category string, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(category, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) SearchByTags(tags []string, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(tags, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) FindArgentineProducts(offset, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) FindHighQualityProducts(offset, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) FindVerified(offset, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) FindActive(offset, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) CountTotal() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockGlobalProductRepository) CountArgentineProducts() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockGlobalProductRepository) CountByQualityScore(score int) (int, error) {
	args := m.Called(score)
	return args.Int(0), args.Error(1)
}

func (m *MockGlobalProductRepository) Update(product *entity.GlobalProduct) (*entity.GlobalProduct, error) {
	args := m.Called(product)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) FindByID(id string) (*entity.GlobalProduct, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockGlobalProductRepository) CountBySource(source string) (int, error) {
	args := m.Called(source)
	return args.Int(0), args.Error(1)
}

func (m *MockGlobalProductRepository) FindNeedingUpdate(maxAgeHours int, limit int) ([]*entity.GlobalProduct, error) {
	args := m.Called(maxAgeHours, limit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.GlobalProduct, error) {
	args := m.Called(ctx, crit)
	return args.Get(0).([]*entity.GlobalProduct), args.Error(1)
}

func (m *MockGlobalProductRepository) CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error) {
	args := m.Called(ctx, crit)
	return args.Int(0), args.Error(1)
}

func TestCreateGlobalProduct_ValidRequest_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewCreateGlobalProduct(mockRepo)

	request := usecase.CreateGlobalProductRequest{
		EAN:         "1234567890128", // EAN válido con checksum correcto
		Name:        "Test Product",
		Description: stringPtr("Test Description"),
		Brand:       stringPtr("Test Brand"),
		Source:      "manual", // Usar fuente válida
		SourceURL:   stringPtr("https://test.com"),
		Reliability: float64Ptr(0.8),
	}

	// Mock que no existe producto con este EAN
	mockRepo.On("FindByEAN", "1234567890128").Return(nil, nil)

	// Crear un producto mock para el Save
	ean, _ := value_object.NewEAN13("1234567890128")
	source, _ := value_object.NewManualSource() // Usar fuente manual para el test
	savedProduct, _ := entity.NewGlobalProduct(
		ean,
		"Test Product",
		stringPtr("Test Description"),
		stringPtr("Test Brand"),
		nil,
		nil,
		nil,
		source,
	)

	mockRepo.On("Save", mock.AnythingOfType("*entity.GlobalProduct")).Return(savedProduct, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "1234567890128", result.EAN)
	assert.Equal(t, "Test Product", result.Name)
	assert.Equal(t, "Test Description", *result.Description)
	assert.Equal(t, "Test Brand", *result.Brand)
	assert.Equal(t, "manual", result.Source)
	assert.False(t, result.IsVerified)
	assert.True(t, result.IsActive)

	mockRepo.AssertExpectations(t)
}

func TestCreateGlobalProduct_InvalidEAN_ReturnsError(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewCreateGlobalProduct(mockRepo)

	request := usecase.CreateGlobalProductRequest{
		EAN:    "invalid_ean",
		Name:   "Test Product",
		Source: "test_source",
	}

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "EAN inválido")
}

func TestCreateGlobalProduct_DuplicateEAN_ReturnsError(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewCreateGlobalProduct(mockRepo)

	request := usecase.CreateGlobalProductRequest{
		EAN:    "1234567890128",
		Name:   "Test Product",
		Source: "test_source",
	}

	// Mock que ya existe un producto con este EAN
	ean, _ := value_object.NewEAN13("1234567890128")
	source, _ := value_object.NewManualSource()
	existingProduct, _ := entity.NewGlobalProduct(
		ean,
		"Existing Product",
		nil,
		nil,
		nil,
		nil,
		nil,
		source,
	)

	mockRepo.On("FindByEAN", "1234567890128").Return(existingProduct, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Ya existe un producto con el EAN")

	mockRepo.AssertExpectations(t)
}

func TestCreateGlobalProduct_RepositoryError_ReturnsError(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewCreateGlobalProduct(mockRepo)

	request := usecase.CreateGlobalProductRequest{
		EAN:    "1234567890128",
		Name:   "Test Product",
		Source: "manual",
	}

	// Mock error en FindByEAN
	mockRepo.On("FindByEAN", "1234567890128").Return(nil, errors.New("database error"))

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Error al verificar EAN existente")

	mockRepo.AssertExpectations(t)
}

func TestCreateGlobalProduct_SaveError_ReturnsError(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewCreateGlobalProduct(mockRepo)

	request := usecase.CreateGlobalProductRequest{
		EAN:    "1234567890128",
		Name:   "Test Product",
		Source: "manual",
	}

	// Mock que no existe producto con este EAN
	mockRepo.On("FindByEAN", "1234567890128").Return(nil, nil)
	// Mock error en Save
	mockRepo.On("Save", mock.AnythingOfType("*entity.GlobalProduct")).Return(nil, errors.New("save error"))

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Error al guardar producto")

	mockRepo.AssertExpectations(t)
}

func TestCreateGlobalProduct_WithBusinessTypeAndTags_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockGlobalProductRepository)
	uc := usecase.NewCreateGlobalProduct(mockRepo)

	request := usecase.CreateGlobalProductRequest{
		EAN:          "1234567890128",
		Name:         "Test Product",
		Source:       "manual",
		BusinessType: stringPtr("retail"),
		Tags:         []string{"tag1", "tag2"},
		Metadata:     map[string]interface{}{"key1": "value1"},
	}

	// Mock que no existe producto con este EAN
	mockRepo.On("FindByEAN", "1234567890128").Return(nil, nil)

	// Crear un producto mock para el Save
	ean, _ := value_object.NewEAN13("1234567890128")
	source, _ := value_object.NewManualSource()
	savedProduct, _ := entity.NewGlobalProduct(
		ean,
		"Test Product",
		nil,
		nil,
		nil,
		nil,
		nil,
		source,
	)
	savedProduct.SetBusinessType("retail")
	savedProduct.AddTag("tag1")
	savedProduct.AddTag("tag2")
	savedProduct.SetMetadata("key1", "value1")

	mockRepo.On("Save", mock.AnythingOfType("*entity.GlobalProduct")).Return(savedProduct, nil)

	// Act
	result, err := uc.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "retail", *result.BusinessType)
	assert.Contains(t, result.Tags, "tag1")
	assert.Contains(t, result.Tags, "tag2")
	assert.Equal(t, "value1", result.Metadata["key1"])

	mockRepo.AssertExpectations(t)
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func float64Ptr(f float64) *float64 {
	return &f
}
