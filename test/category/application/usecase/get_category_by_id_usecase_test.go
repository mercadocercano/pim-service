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

func TestGetCategoryByIDUseCase_Execute_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	getUseCase := usecase.NewGetCategoryByIDUseCase(mockRepo)
	categoryMother := testentity.Create()

	// Crear categoría usando Object Mother
	existingCategory := categoryMother.WithTenantID("tenant-123")
	existingCategory.Name = "Electrónicos"
	mockRepo.SetupCategories([]*entity.Category{existingCategory})

	// Act
	category, err := getUseCase.Execute(ctx, existingCategory.ID, existingCategory.TenantID)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, existingCategory.ID, category.ID)
	assert.Equal(t, existingCategory.TenantID, category.TenantID)
	assert.Equal(t, existingCategory.Name, category.Name)
	assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
}

func TestGetCategoryByIDUseCase_Execute_CategoryNotFound(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	getUseCase := usecase.NewGetCategoryByIDUseCase(mockRepo)

	tenantID := "tenant-123"
	nonExistentID := "550e8400-e29b-41d4-a716-446655440000" // UUID válido pero no existe

	// Act
	category, err := getUseCase.Execute(ctx, nonExistentID, tenantID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
}

func TestGetCategoryByIDUseCase_Execute_InvalidUUID_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	getUseCase := usecase.NewGetCategoryByIDUseCase(mockRepo)

	tenantID := "tenant-123"
	invalidID := "invalid-uuid"

	// Act
	category, err := getUseCase.Execute(ctx, invalidID, tenantID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Equal(t, usecase.ErrInvalidUUID, err)
	assert.Equal(t, 0, mockRepo.GetCallCount("FindByID")) // No debería llamar al repo
}

func TestGetCategoryByIDUseCase_Execute_EmptyID_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	getUseCase := usecase.NewGetCategoryByIDUseCase(mockRepo)

	tenantID := "tenant-123"
	emptyID := ""

	// Act
	category, err := getUseCase.Execute(ctx, emptyID, tenantID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Equal(t, usecase.ErrInvalidUUID, err)
	assert.Equal(t, 0, mockRepo.GetCallCount("FindByID"))
}

func TestGetCategoryByIDUseCase_Execute_RepositoryError_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	mockRepo.ShouldFailOn("FindByID")
	getUseCase := usecase.NewGetCategoryByIDUseCase(mockRepo)

	tenantID := "tenant-123"
	categoryID := "550e8400-e29b-41d4-a716-446655440000"

	// Act
	category, err := getUseCase.Execute(ctx, categoryID, tenantID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Equal(t, repository.ErrMockFailedOp, err)
	assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
}

func TestGetCategoryByIDUseCase_Execute_DifferentTenants_ShouldRespectTenancy(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	getUseCase := usecase.NewGetCategoryByIDUseCase(mockRepo)
	categoryMother := testentity.Create()

	// Crear categorías para diferentes tenants usando Object Mother
	category1 := categoryMother.WithTenantID("tenant-1")
	category1.Name = "Categoría Tenant 1"
	category2 := categoryMother.WithTenantID("tenant-2")
	category2.Name = "Categoría Tenant 2"

	mockRepo.SetupCategories([]*entity.Category{category1, category2})

	// Act - Intentar obtener categoría de tenant-1 usando tenant-2
	category, err := getUseCase.Execute(ctx, category1.ID, "tenant-2")

	// Assert - Debería fallar por tenancy
	assert.Error(t, err)
	assert.Nil(t, category)
}

func TestGetCategoryByIDUseCase_Execute_WithParentCategory(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	getUseCase := usecase.NewGetCategoryByIDUseCase(mockRepo)
	categoryMother := testentity.Create()

	// Crear categoría padre e hija usando Object Mother
	parentCategory := categoryMother.WithTenantID("tenant-123")
	parentCategory.Name = "Electrónicos"
	childCategory := categoryMother.WithParent(parentCategory.ID)
	childCategory.TenantID = "tenant-123"
	childCategory.Name = "Smartphones"

	mockRepo.SetupCategories([]*entity.Category{parentCategory, childCategory})

	// Act
	category, err := getUseCase.Execute(ctx, childCategory.ID, childCategory.TenantID)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, childCategory.ID, category.ID)
	assert.Equal(t, childCategory.Name, category.Name)
	assert.NotNil(t, category.ParentID)
	assert.Equal(t, parentCategory.ID, *category.ParentID)
}

func TestGetCategoryByIDUseCase_Execute_InactiveCategory_ShouldReturnCategory(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockCategoryRepository()
	getUseCase := usecase.NewGetCategoryByIDUseCase(mockRepo)
	categoryMother := testentity.Create()

	// Crear categoría inactiva usando Object Mother
	inactiveCategory := categoryMother.Inactive()
	inactiveCategory.TenantID = "tenant-123"
	inactiveCategory.Name = "Categoría Inactiva"
	mockRepo.SetupCategories([]*entity.Category{inactiveCategory})

	// Act
	category, err := getUseCase.Execute(ctx, inactiveCategory.ID, inactiveCategory.TenantID)

	// Assert - Debería devolver la categoría aunque esté inactiva
	require.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, inactiveCategory.ID, category.ID)
	assert.False(t, category.IsActive())
}
