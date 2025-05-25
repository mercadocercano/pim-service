package usecase_test

import (
	"context"
	"testing"

	"pim/src/category/application/usecase"
	"pim/src/category/domain/entity"
	"pim/src/category/domain/value_object"
	testentity "pim/test/category/domain/entity"
	"pim/test/category/infrastructure/persistence/repository"

	"github.com/stretchr/testify/assert"
)

func TestUpdateCategoryUseCase_Execute(t *testing.T) {
	// Arrange
	mockRepo := repository.NewMockCategoryRepository()
	updateUseCase := usecase.NewUpdateCategoryUseCase(mockRepo)
	ctx := context.Background()
	categoryMother := testentity.CategoryMother{}
	tenantID := "tenant-123"

	t.Run("debería actualizar una categoría con éxito", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		existingCategory := categoryMother.WithName("Categoría Original")
		mockRepo.SetupCategories([]*entity.Category{categoryMother.WithID(existingCategory.ID)})

		newName := "Categoría Actualizada"
		newDescription := "Nueva descripción"
		var parentID *string
		if existingCategory.ParentID != nil {
			parentID = existingCategory.ParentID
		}

		// Act
		updatedCategory, err := updateUseCase.Execute(ctx, existingCategory.ID, tenantID, newName, newDescription, parentID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, updatedCategory)
		assert.Equal(t, newName, updatedCategory.Name)
		assert.Equal(t, newDescription, updatedCategory.Description)
		assert.Equal(t, existingCategory.ID, updatedCategory.ID)
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 1, mockRepo.GetCallCount("Update"))

		// Verificar que la categoría fue actualizada en el repositorio
		storedCategory, _ := mockRepo.FindByID(ctx, existingCategory.ID, tenantID)
		assert.Equal(t, newName, storedCategory.Name)
		assert.Equal(t, newDescription, storedCategory.Description)
	})

	t.Run("debería fallar si la categoría no existe", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		defaultCategory := categoryMother.WithDefaults()
		mockRepo.SetupCategories([]*entity.Category{defaultCategory})

		nonExistentID := "id-inexistente"
		newName := "Categoría Actualizada"
		newDescription := "Nueva descripción"

		// Act
		updatedCategory, err := updateUseCase.Execute(ctx, nonExistentID, tenantID, newName, newDescription, nil)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, updatedCategory)
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 0, mockRepo.GetCallCount("Update"))
	})

	t.Run("debería fallar si el nombre es inválido", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		existingCategory := categoryMother.WithDefaults()
		mockRepo.SetupCategories([]*entity.Category{existingCategory})

		invalidName := ""
		newDescription := "Nueva descripción"
		var parentID *string
		if existingCategory.ParentID != nil {
			parentID = existingCategory.ParentID
		}

		// Act
		updatedCategory, err := updateUseCase.Execute(ctx, existingCategory.ID, tenantID, invalidName, newDescription, parentID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, updatedCategory)
		assert.Equal(t, value_object.ErrEmptyName, err)
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 0, mockRepo.GetCallCount("Update")) // No se incrementa porque falla antes
	})

	t.Run("debería fallar si el repositorio de actualización falla", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		existingCategory := categoryMother.WithDefaults()
		mockRepo.SetupCategories([]*entity.Category{existingCategory})
		mockRepo.ShouldFailOn("Update")

		newName := "Categoría Actualizada"
		newDescription := "Nueva descripción"
		var parentID *string
		if existingCategory.ParentID != nil {
			parentID = existingCategory.ParentID
		}

		// Act
		updatedCategory, err := updateUseCase.Execute(ctx, existingCategory.ID, tenantID, newName, newDescription, parentID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, updatedCategory)
		assert.Equal(t, repository.ErrMockFailedOp, err)
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 1, mockRepo.GetCallCount("Update"))
	})
}
