package usecase_test

import (
	"context"
	"testing"

	"saas-mt-pim-service/src/category/application/usecase"
	"saas-mt-pim-service/src/category/domain/entity"
	testentity "saas-mt-pim-service/test/category/domain/entity"
	"saas-mt-pim-service/test/category/infrastructure/persistence/repository"

	"github.com/stretchr/testify/assert"
)

func TestChangeCategoryStatusUseCase_Activate(t *testing.T) {
	// Arrange
	mockRepo := repository.NewMockCategoryRepository()
	statusUseCase := usecase.NewChangeCategoryStatusUseCase(mockRepo)
	ctx := context.Background()
	categoryMother := testentity.CategoryMother{}
	tenantID := "tenant-123"

	t.Run("debería activar una categoría inactiva con éxito", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		inactiveCategory := categoryMother.Inactive()
		mockRepo.SetupCategories([]*entity.Category{inactiveCategory})

		// Act
		activatedCategory, err := statusUseCase.Activate(ctx, inactiveCategory.ID, tenantID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, activatedCategory)
		assert.True(t, activatedCategory.IsActive())
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 1, mockRepo.GetCallCount("Update"))

		// Verificar que la categoría fue actualizada en el repositorio
		storedCategory, _ := mockRepo.FindByID(ctx, inactiveCategory.ID, tenantID)
		assert.True(t, storedCategory.IsActive())
	})

	t.Run("debería activar una categoría que ya está activa", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		activeCategory := categoryMother.WithDefaults() // Ya está activa por defecto
		mockRepo.SetupCategories([]*entity.Category{activeCategory})

		// Act
		activatedCategory, err := statusUseCase.Activate(ctx, activeCategory.ID, tenantID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, activatedCategory)
		assert.True(t, activatedCategory.IsActive())
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 1, mockRepo.GetCallCount("Update")) // Siempre actualiza
	})

	t.Run("debería fallar si la categoría no existe", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		nonExistentID := "id-inexistente"

		// Act
		activatedCategory, err := statusUseCase.Activate(ctx, nonExistentID, tenantID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, activatedCategory)
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 0, mockRepo.GetCallCount("Update"))
	})

	t.Run("debería fallar si el repositorio de actualización falla", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		mockRepo.ShouldFailOn("Update")
		inactiveCategory := categoryMother.Inactive()
		mockRepo.SetupCategories([]*entity.Category{inactiveCategory})

		// Act
		activatedCategory, err := statusUseCase.Activate(ctx, inactiveCategory.ID, tenantID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, activatedCategory)
		assert.Equal(t, repository.ErrMockFailedOp, err)
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 1, mockRepo.GetCallCount("Update"))
	})
}

func TestChangeCategoryStatusUseCase_Deactivate(t *testing.T) {
	// Arrange
	mockRepo := repository.NewMockCategoryRepository()
	statusUseCase := usecase.NewChangeCategoryStatusUseCase(mockRepo)
	ctx := context.Background()
	categoryMother := testentity.CategoryMother{}
	tenantID := "tenant-123"

	t.Run("debería desactivar una categoría activa con éxito", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		activeCategory := categoryMother.WithDefaults() // Activa por defecto
		mockRepo.SetupCategories([]*entity.Category{activeCategory})

		// Act
		deactivatedCategory, err := statusUseCase.Deactivate(ctx, activeCategory.ID, tenantID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, deactivatedCategory)
		assert.False(t, deactivatedCategory.IsActive())
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 1, mockRepo.GetCallCount("Update"))

		// Verificar que la categoría fue actualizada en el repositorio
		storedCategory, _ := mockRepo.FindByID(ctx, activeCategory.ID, tenantID)
		assert.False(t, storedCategory.IsActive())
	})

	t.Run("debería desactivar una categoría que ya está inactiva", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		inactiveCategory := categoryMother.Inactive()
		mockRepo.SetupCategories([]*entity.Category{inactiveCategory})

		// Act
		deactivatedCategory, err := statusUseCase.Deactivate(ctx, inactiveCategory.ID, tenantID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, deactivatedCategory)
		assert.False(t, deactivatedCategory.IsActive())
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 1, mockRepo.GetCallCount("Update")) // Siempre actualiza
	})

	t.Run("debería fallar si la categoría no existe", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		nonExistentID := "id-inexistente"

		// Act
		deactivatedCategory, err := statusUseCase.Deactivate(ctx, nonExistentID, tenantID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, deactivatedCategory)
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 0, mockRepo.GetCallCount("Update"))
	})

	t.Run("debería fallar si el repositorio de actualización falla", func(t *testing.T) {
		// Arrange
		mockRepo.ResetFailures()
		mockRepo.ResetCallHistory()
		mockRepo.ShouldFailOn("Update")
		activeCategory := categoryMother.WithDefaults()
		mockRepo.SetupCategories([]*entity.Category{activeCategory})

		// Act
		deactivatedCategory, err := statusUseCase.Deactivate(ctx, activeCategory.ID, tenantID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, deactivatedCategory)
		assert.Equal(t, repository.ErrMockFailedOp, err)
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 1, mockRepo.GetCallCount("Update"))
	})
}
