package usecase_test

import (
	"context"
	"testing"

	"pim/src/category/application/usecase"
	"pim/src/category/domain/entity"
	testentity "pim/test/category/domain/entity"
	"pim/test/category/infrastructure/persistence/repository"

	"github.com/stretchr/testify/assert"
)

func TestMoveCategoryUseCase_Execute(t *testing.T) {
	// Arrange
	ctx := context.Background()
	categoryMother := testentity.CategoryMother{}
	tenantID := "tenant-123"

	t.Run("debería mover una categoría a un nuevo padre con éxito", func(t *testing.T) {
		// Arrange
		mockRepo := repository.NewMockCategoryRepository()
		moveUseCase := usecase.NewMoveCategoryUseCase(mockRepo)
		parentCategory := categoryMother.WithID("parent-id")
		childCategory := categoryMother.WithID("child-id")

		// Configurar el repositorio con las categorías
		mockRepo.SetupCategories([]*entity.Category{parentCategory, childCategory})

		// Act
		movedCategory, err := moveUseCase.Execute(ctx, childCategory.ID, tenantID, &parentCategory.ID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, movedCategory)
		assert.NotNil(t, movedCategory.ParentID)
		assert.Equal(t, parentCategory.ID, *movedCategory.ParentID)
		assert.Equal(t, 1, mockRepo.GetCallCount("Update"))

		// Verificar que la categoría fue actualizada en el repositorio
		storedCategory, _ := mockRepo.FindByID(ctx, childCategory.ID, tenantID)
		assert.NotNil(t, storedCategory.ParentID)
		assert.Equal(t, parentCategory.ID, *storedCategory.ParentID)
	})

	t.Run("debería quitar el padre de una categoría con éxito", func(t *testing.T) {
		// Arrange
		mockRepo := repository.NewMockCategoryRepository()
		moveUseCase := usecase.NewMoveCategoryUseCase(mockRepo)

		// Crear una categoría con padre
		parentID := "parent-id"
		childWithParent := categoryMother.WithParent(parentID)

		// Configurar el repositorio con la categoría
		mockRepo.SetupCategories([]*entity.Category{childWithParent})

		// Act - Mover a la raíz (sin padre)
		movedCategory, err := moveUseCase.Execute(ctx, childWithParent.ID, tenantID, nil)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, movedCategory)
		assert.Nil(t, movedCategory.ParentID)
		assert.Equal(t, 1, mockRepo.GetCallCount("Update"))

		// Verificar que la categoría fue actualizada en el repositorio
		storedCategory, _ := mockRepo.FindByID(ctx, childWithParent.ID, tenantID)
		assert.Nil(t, storedCategory.ParentID)
	})

	t.Run("debería fallar si la categoría no existe", func(t *testing.T) {
		// Arrange
		mockRepo := repository.NewMockCategoryRepository()
		moveUseCase := usecase.NewMoveCategoryUseCase(mockRepo)
		parentCategory := categoryMother.WithID("parent-id")
		mockRepo.SetupCategories([]*entity.Category{parentCategory})

		nonExistentID := "id-inexistente"
		parentIDStr := "parent-id"

		// Act
		movedCategory, err := moveUseCase.Execute(ctx, nonExistentID, tenantID, &parentIDStr)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, movedCategory)
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 0, mockRepo.GetCallCount("Update"))
	})

	t.Run("debería fallar si la categoría padre no existe", func(t *testing.T) {
		// Arrange
		mockRepo := repository.NewMockCategoryRepository()
		moveUseCase := usecase.NewMoveCategoryUseCase(mockRepo)
		childCategory := categoryMother.WithID("child-id")
		mockRepo.SetupCategories([]*entity.Category{childCategory})

		nonExistentParentID := "parent-inexistente"

		// Act
		movedCategory, err := moveUseCase.Execute(ctx, childCategory.ID, tenantID, &nonExistentParentID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, movedCategory)
		assert.Contains(t, err.Error(), "la categoría padre no existe")
		assert.Equal(t, 2, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 0, mockRepo.GetCallCount("Update"))
	})

	t.Run("debería fallar si una categoría intenta ser su propio padre", func(t *testing.T) {
		// Arrange
		mockRepo := repository.NewMockCategoryRepository()
		moveUseCase := usecase.NewMoveCategoryUseCase(mockRepo)
		category := categoryMother.WithID("category-id")
		mockRepo.SetupCategories([]*entity.Category{category})

		// Act
		movedCategory, err := moveUseCase.Execute(ctx, category.ID, tenantID, &category.ID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, movedCategory)
		assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
		assert.Equal(t, 0, mockRepo.GetCallCount("Update"))
	})

	t.Run("debería fallar si el repositorio de actualización falla", func(t *testing.T) {
		// Arrange
		mockRepo := repository.NewMockCategoryRepository()
		moveUseCase := usecase.NewMoveCategoryUseCase(mockRepo)
		mockRepo.ShouldFailOn("Update")
		parentCategory := categoryMother.WithID("parent-id")
		childCategory := categoryMother.WithID("child-id")
		mockRepo.SetupCategories([]*entity.Category{parentCategory, childCategory})

		// Act
		movedCategory, err := moveUseCase.Execute(ctx, childCategory.ID, tenantID, &parentCategory.ID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, movedCategory)
		assert.Equal(t, repository.ErrMockFailedOp, err)
		assert.Equal(t, 1, mockRepo.GetCallCount("Update"))
	})
}
