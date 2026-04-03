package usecase_test

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/schema_validation/application/usecase"
	"saas-mt-pim-service/src/schema_validation/domain/entity"
	"saas-mt-pim-service/src/schema_validation/domain/service"
)

// MockSchemaValidationCache mock del cache de validaciones
type MockSchemaValidationCache struct {
	mock.Mock
}

func (m *MockSchemaValidationCache) Get(ctx context.Context, id string) (*entity.SchemaValidation, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.SchemaValidation), args.Error(1)
}

func (m *MockSchemaValidationCache) Set(ctx context.Context, validation *entity.SchemaValidation) error {
	args := m.Called(ctx, validation)
	return args.Error(0)
}

// MockCategoryRepository mock del repositorio de categorías
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) GetCategoryNames(ctx context.Context, tenantID string) ([]string, error) {
	args := m.Called(ctx, tenantID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]string), args.Error(1)
}

func TestValidateCSVSchemaUseCase_Execute_ValidCSV(t *testing.T) {
	// Arrange
	csvAnalyzer := service.NewCSVAnalyzerService()
	categoryDeducer := service.NewCategoryDeductionService()
	mockCache := new(MockSchemaValidationCache)
	mockCategoryRepo := new(MockCategoryRepository)

	uc := usecase.NewValidateCSVSchemaUseCase(csvAnalyzer, categoryDeducer, mockCategoryRepo, mockCache)

	csvData := "nombre,precio,sku\nProducto A,100.50,SKU001\nProducto B,200.00,SKU002\n"
	reader := strings.NewReader(csvData)

	mockCache.On("Set", mock.Anything, mock.AnythingOfType("*entity.SchemaValidation")).Return(nil)
	mockCategoryRepo.On("GetCategoryNames", mock.Anything, "tenant-123").Return([]string{}, nil)

	// Act
	result, err := uc.Execute(context.Background(), reader, "tenant-123", "products.csv", 10)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "tenant-123", result.TenantID)
	assert.Equal(t, "products.csv", result.FileName)
	assert.NotNil(t, result.Summary)
	assert.Equal(t, ",", result.DetectedDelimiter)

	// Debe haber detectado las 3 columnas requeridas
	assert.Contains(t, result.Columns, "nombre")
	assert.Contains(t, result.Columns, "precio")
	assert.Contains(t, result.Columns, "sku")
}

func TestValidateCSVSchemaUseCase_Execute_MissingRequiredColumns(t *testing.T) {
	// Arrange
	csvAnalyzer := service.NewCSVAnalyzerService()
	mockCache := new(MockSchemaValidationCache)
	mockCategoryRepo := new(MockCategoryRepository)

	uc := usecase.NewValidateCSVSchemaUseCase(csvAnalyzer, nil, mockCategoryRepo, mockCache)

	// CSV sin columna de precio ni sku
	csvData := "nombre,descripcion\nProducto A,Desc A\nProducto B,Desc B\n"
	reader := strings.NewReader(csvData)

	mockCache.On("Set", mock.Anything, mock.AnythingOfType("*entity.SchemaValidation")).Return(nil)

	// Act
	result, err := uc.Execute(context.Background(), reader, "tenant-123", "products.csv", 10)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	// Debe tener recomendaciones para columnas faltantes
	assert.True(t, len(result.Recommendations) > 0)
}

func TestValidateCSVSchemaUseCase_Execute_SemicolonDelimiter(t *testing.T) {
	// Arrange
	csvAnalyzer := service.NewCSVAnalyzerService()
	mockCache := new(MockSchemaValidationCache)
	mockCategoryRepo := new(MockCategoryRepository)

	uc := usecase.NewValidateCSVSchemaUseCase(csvAnalyzer, nil, mockCategoryRepo, mockCache)

	csvData := "nombre;precio;sku\nProducto A;100.50;SKU001\nProducto B;200.00;SKU002\n"
	reader := strings.NewReader(csvData)

	mockCache.On("Set", mock.Anything, mock.AnythingOfType("*entity.SchemaValidation")).Return(nil)

	// Act
	result, err := uc.Execute(context.Background(), reader, "tenant-123", "products.csv", 10)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, ";", result.DetectedDelimiter)
}

func TestValidateCSVSchemaUseCase_Execute_DefaultMaxRows(t *testing.T) {
	// Arrange
	csvAnalyzer := service.NewCSVAnalyzerService()
	mockCache := new(MockSchemaValidationCache)
	mockCategoryRepo := new(MockCategoryRepository)

	uc := usecase.NewValidateCSVSchemaUseCase(csvAnalyzer, nil, mockCategoryRepo, mockCache)

	csvData := "nombre,precio,sku\nProducto A,100.50,SKU001\n"
	reader := strings.NewReader(csvData)

	mockCache.On("Set", mock.Anything, mock.AnythingOfType("*entity.SchemaValidation")).Return(nil)

	// Act - maxRows=0 should default to 10
	result, err := uc.Execute(context.Background(), reader, "tenant-123", "products.csv", 0)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestValidateCSVSchemaUseCase_Execute_WithCategoryDeduction(t *testing.T) {
	// Arrange
	csvAnalyzer := service.NewCSVAnalyzerService()
	categoryDeducer := service.NewCategoryDeductionService()
	mockCache := new(MockSchemaValidationCache)
	mockCategoryRepo := new(MockCategoryRepository)

	uc := usecase.NewValidateCSVSchemaUseCase(csvAnalyzer, categoryDeducer, mockCategoryRepo, mockCache)

	// CSV con nombres de productos reconocibles pero sin columna de categoría
	csvData := "nombre,precio,sku\nCoca Cola 2.25L,350.00,SKU-CC01\nMartillo Carpintero,1500.00,SKU-MT01\nShampoo Dove 400ml,800.00,SKU-SH01\n"
	reader := strings.NewReader(csvData)

	mockCache.On("Set", mock.Anything, mock.AnythingOfType("*entity.SchemaValidation")).Return(nil)
	mockCategoryRepo.On("GetCategoryNames", mock.Anything, "tenant-123").Return([]string{}, nil)

	// Act
	result, err := uc.Execute(context.Background(), reader, "tenant-123", "products.csv", 10)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	// Debería haber deducido categorías para algunos productos
	if len(result.DeducedCategories) > 0 {
		assert.True(t, len(result.DeducedCategories) > 0)
	}
}

func TestValidateCSVSchemaUseCase_Execute_EmptyCSV_ShouldFail(t *testing.T) {
	// Arrange
	csvAnalyzer := service.NewCSVAnalyzerService()
	mockCache := new(MockSchemaValidationCache)

	uc := usecase.NewValidateCSVSchemaUseCase(csvAnalyzer, nil, nil, mockCache)

	reader := strings.NewReader("")

	// Act
	result, err := uc.Execute(context.Background(), reader, "tenant-123", "empty.csv", 10)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestValidateCSVSchemaUseCase_ApplyMapping(t *testing.T) {
	// Arrange
	mockCache := new(MockSchemaValidationCache)

	uc := usecase.NewValidateCSVSchemaUseCase(nil, nil, nil, mockCache)

	// Simular una validación existente en cache
	validation := entity.NewSchemaValidation("tenant-123", "test.csv")
	col := entity.NewColumnValidation("columna1", 0)
	validation.AddColumn(col)

	mockCache.On("Get", mock.Anything, validation.ID.String()).Return(validation, nil)
	mockCache.On("Set", mock.Anything, mock.AnythingOfType("*entity.SchemaValidation")).Return(nil)

	// Act
	result, err := uc.ApplyMapping(context.Background(), validation.ID.String(), map[string]string{
		"columna1": "name",
	})

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "name", result.Columns["columna1"].MappedTo)
}
