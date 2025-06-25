package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"pim/src/category/application/usecase"
	"pim/src/category/domain/entity"
	testentity "pim/test/category/domain/entity"
	"pim/test/category/infrastructure/persistence/repository"
)

func TestCreateCategoryUseCase_Execute_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	createUseCase := usecase.NewCreateCategoryUseCase(mockRepo)

	tenantID := "tenant-123"
	name := "Nueva Categoría"
	description := "Descripción de prueba"
	var parentID *string = nil

	// Act
	category, err := createUseCase.Execute(ctx, tenantID, name, description, parentID)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, tenantID, category.TenantID)
	assert.Equal(t, name, category.Name)
	assert.Equal(t, description, category.Description)
	assert.Nil(t, category.ParentID)
	assert.True(t, category.IsActive())
	assert.NotEmpty(t, category.ID)
	assert.Equal(t, 1, mockRepo.GetCallCount("Create"))

	// Verificar que la categoría está en el repositorio
	categories := mockRepo.GetCategories()
	assert.Len(t, categories, 1)
	assert.Equal(t, category.ID, categories[0].ID)
}

func TestCreateCategoryUseCase_Execute_WithParent(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	createUseCase := usecase.NewCreateCategoryUseCase(mockRepo)
	categoryMother := testentity.Create()

	// Crear categoría padre usando Object Mother
	parentCategory := categoryMother.WithTenantID("tenant-123")
	parentCategory.Name = "Categoría Padre"
	mockRepo.SetupCategories([]*entity.Category{parentCategory})

	tenantID := "tenant-123"
	name := "Categoría Hija"
	description := "Descripción de categoría hija"
	parentID := &parentCategory.ID

	// Act
	category, err := createUseCase.Execute(ctx, tenantID, name, description, parentID)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, tenantID, category.TenantID)
	assert.Equal(t, name, category.Name)
	assert.Equal(t, description, category.Description)
	assert.NotNil(t, category.ParentID)
	assert.Equal(t, parentCategory.ID, *category.ParentID)
	assert.True(t, category.IsActive())
	assert.Equal(t, 1, mockRepo.GetCallCount("Create"))
}

func TestCreateCategoryUseCase_Execute_EmptyName_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	createUseCase := usecase.NewCreateCategoryUseCase(mockRepo)

	tenantID := "tenant-123"
	emptyName := ""
	description := "Descripción de prueba"
	var parentID *string = nil

	// Act
	category, err := createUseCase.Execute(ctx, tenantID, emptyName, description, parentID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Contains(t, err.Error(), "nombre")
	assert.Equal(t, 0, mockRepo.GetCallCount("Create"))
}

func TestCreateCategoryUseCase_Execute_EmptyTenantID_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	createUseCase := usecase.NewCreateCategoryUseCase(mockRepo)

	emptyTenantID := ""
	name := "Categoría de Prueba"
	description := "Descripción de prueba"
	var parentID *string = nil

	// Act
	category, err := createUseCase.Execute(ctx, emptyTenantID, name, description, parentID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Contains(t, err.Error(), "tenant")
	assert.Equal(t, 0, mockRepo.GetCallCount("Create"))
}

func TestCreateCategoryUseCase_Execute_RepositoryError_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	mockRepo.ShouldFailOn("Create")
	createUseCase := usecase.NewCreateCategoryUseCase(mockRepo)

	tenantID := "tenant-123"
	name := "Nueva Categoría"
	description := "Descripción de prueba"
	var parentID *string = nil

	// Act
	category, err := createUseCase.Execute(ctx, tenantID, name, description, parentID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Equal(t, repository.ErrMockFailedOp, err)
	assert.Equal(t, 1, mockRepo.GetCallCount("Create"))
}

func TestCreateCategoryUseCase_Execute_MultipleCategories_ShouldSucceed(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	createUseCase := usecase.NewCreateCategoryUseCase(mockRepo)

	tenantID := "tenant-123"

	// Crear múltiples categorías
	testCases := []struct {
		name        string
		description string
	}{
		{"Electrónicos", "Productos electrónicos"},
		{"Ropa", "Ropa y accesorios"},
		{"Hogar", "Productos para el hogar"},
	}

	// Act & Assert
	for _, tc := range testCases {
		var parentID *string = nil
		category, err := createUseCase.Execute(ctx, tenantID, tc.name, tc.description, parentID)

		require.NoError(t, err)
		assert.NotNil(t, category)
		assert.Equal(t, tc.name, category.Name)
		assert.Equal(t, tc.description, category.Description)
		assert.Equal(t, tenantID, category.TenantID)
	}

	// Verificar que todas las categorías fueron creadas
	categories := mockRepo.GetCategories()
	assert.Len(t, categories, len(testCases))
	assert.Equal(t, len(testCases), mockRepo.GetCallCount("Create"))
}

func TestCreateCategoryUseCase_Execute_DifferentTenants_ShouldSucceed(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	createUseCase := usecase.NewCreateCategoryUseCase(mockRepo)

	// Datos de prueba para diferentes tenants
	testData := []struct {
		tenantID string
		name     string
	}{
		{"tenant-1", "Categoría Tenant 1"},
		{"tenant-2", "Categoría Tenant 2"},
		{"tenant-3", "Categoría Tenant 3"},
	}

	// Act & Assert
	for _, td := range testData {
		var parentID *string = nil
		category, err := createUseCase.Execute(ctx, td.tenantID, td.name, "Descripción", parentID)

		require.NoError(t, err)
		assert.NotNil(t, category)
		assert.Equal(t, td.tenantID, category.TenantID)
		assert.Equal(t, td.name, category.Name)
	}

	// Verificar que todas las categorías fueron creadas
	categories := mockRepo.GetCategories()
	assert.Len(t, categories, len(testData))
}

func TestCreateCategoryUseCase_Execute_InvalidNameLength_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	createUseCase := usecase.NewCreateCategoryUseCase(mockRepo)

	tenantID := "tenant-123"
	// Nombre muy largo (asumiendo que hay validación de longitud)
	longName := "Este es un nombre de categoría extremadamente largo que probablemente exceda los límites permitidos por la validación del objeto de valor Name y debería fallar"
	description := "Descripción de prueba"
	var parentID *string = nil

	// Act
	category, err := createUseCase.Execute(ctx, tenantID, longName, description, parentID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Equal(t, 0, mockRepo.GetCallCount("Create"))
}
