package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"pim/src/category/application/usecase"
	"pim/src/category/domain/value_object"
	"pim/test/category/infrastructure/persistence/repository"
)

func TestCreateCategoryUseCase_Execute(t *testing.T) {
	// Arrange
	mockRepo := repository.NewMockCategoryRepository()
	createUseCase := usecase.NewCreateCategoryUseCase(mockRepo)
	ctx := context.Background()

	t.Run("debería crear una categoría con éxito", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		tenantID := "tenant-123"
		name := "Nueva Categoría"
		description := "Descripción de prueba"
		parentID := ""

		// Act
		category, err := createUseCase.Execute(ctx, tenantID, name, description, &parentID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, category)
		assert.Equal(t, tenantID, category.TenantID)
		assert.Equal(t, name, category.Name)
		assert.Equal(t, description, category.Description)
		assert.NotNil(t, category.ParentID)
		assert.Equal(t, "", *category.ParentID)
		assert.True(t, category.IsActive())
		assert.NotEmpty(t, category.ID)
		assert.Equal(t, 1, mockRepo.GetCallCount("Create"))

		// Verificar que la categoría está en el repositorio
		categories := mockRepo.GetCategories()
		assert.Len(t, categories, 1)
		assert.Equal(t, category.ID, categories[0].ID)
	})

	t.Run("debería fallar si el nombre es inválido", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		tenantID := "tenant-123"
		emptyName := ""
		description := "Descripción de prueba"
		parentID := ""

		// Act
		category, err := createUseCase.Execute(ctx, tenantID, emptyName, description, &parentID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, category)
		assert.Equal(t, value_object.ErrEmptyName, err)
		assert.Equal(t, 0, mockRepo.GetCallCount("Create")) // No incrementa porque falla antes
	})

	t.Run("debería fallar si el repositorio falla", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		mockRepo.ShouldFailOn("Create")
		tenantID := "tenant-123"
		name := "Nueva Categoría"
		description := "Descripción de prueba"
		parentID := ""

		// Act
		category, err := createUseCase.Execute(ctx, tenantID, name, description, &parentID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, category)
		assert.Equal(t, repository.ErrMockFailedOp, err)
		assert.Equal(t, 1, mockRepo.GetCallCount("Create"))
	})
}
