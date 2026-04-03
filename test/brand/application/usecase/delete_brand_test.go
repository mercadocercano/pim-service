package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/brand/application/usecase"
	"saas-mt-pim-service/src/brand/domain/entity"
	"saas-mt-pim-service/src/brand/domain/service"
	testentity "saas-mt-pim-service/test/brand/domain/entity"
	"saas-mt-pim-service/test/brand/infrastructure/persistence/repository"
)

func TestDeleteBrandUseCase_Execute_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	uc := usecase.NewDeleteBrandUseCase(mockRepo, brandService)

	mother := testentity.Create()
	existingBrand := mother.Nike()
	existingBrand.TenantID = "tenant-123"
	mockRepo.SetupBrands([]*entity.Brand{existingBrand})

	// Act
	err := uc.Execute(ctx, existingBrand.ID, "tenant-123")

	// Assert
	require.NoError(t, err)
	// Verificar que la marca fue actualizada (soft delete via Update)
	assert.Equal(t, 1, mockRepo.GetCallCount("Update"))
}

func TestDeleteBrandUseCase_Execute_BrandNotFound_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	uc := usecase.NewDeleteBrandUseCase(mockRepo, brandService)

	// Act
	err := uc.Execute(ctx, "nonexistent-id", "tenant-123")

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no encontrada")
}

func TestDeleteBrandUseCase_Execute_AlreadyDeleted_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	uc := usecase.NewDeleteBrandUseCase(mockRepo, brandService)

	mother := testentity.Create()
	deletedBrand := mother.Deleted()
	deletedBrand.TenantID = "tenant-123"
	mockRepo.SetupBrands([]*entity.Brand{deletedBrand})

	// Act
	err := uc.Execute(ctx, deletedBrand.ID, "tenant-123")

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no se puede eliminar")
}

func TestDeleteBrandUseCase_Execute_RepositoryError_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	mockRepo.ShouldFailOn("FindByID")
	brandService := service.NewBrandDomainService(mockRepo)
	uc := usecase.NewDeleteBrandUseCase(mockRepo, brandService)

	// Act
	err := uc.Execute(ctx, "some-id", "tenant-123")

	// Assert
	assert.Error(t, err)
}

func TestDeleteBrandUseCase_Execute_WrongTenant_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	uc := usecase.NewDeleteBrandUseCase(mockRepo, brandService)

	mother := testentity.Create()
	existingBrand := mother.Nike()
	existingBrand.TenantID = "tenant-123"
	mockRepo.SetupBrands([]*entity.Brand{existingBrand})

	// Act - intentar eliminar con otro tenant
	err := uc.Execute(ctx, existingBrand.ID, "tenant-999")

	// Assert
	assert.Error(t, err)
}
