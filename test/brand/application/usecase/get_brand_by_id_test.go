package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/brand/application/mapper"
	"saas-mt-pim-service/src/brand/application/usecase"
	"saas-mt-pim-service/src/brand/domain/entity"
	testentity "saas-mt-pim-service/test/brand/domain/entity"
	"saas-mt-pim-service/test/brand/infrastructure/persistence/repository"
)

func TestGetBrandByIDUseCase_Execute_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewGetBrandByIDUseCase(mockRepo, brandMapper)

	mother := testentity.Create()
	existingBrand := mother.Nike()
	existingBrand.TenantID = "tenant-123"
	mockRepo.SetupBrands([]*entity.Brand{existingBrand})

	// Act
	result, err := uc.Execute(ctx, existingBrand.ID, "tenant-123")

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, existingBrand.ID, result.ID)
	assert.Equal(t, "Nike", result.Name)
	assert.Equal(t, "active", result.Status)
	assert.Equal(t, 1, mockRepo.GetCallCount("FindByID"))
}

func TestGetBrandByIDUseCase_Execute_NotFound_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewGetBrandByIDUseCase(mockRepo, brandMapper)

	// Act
	result, err := uc.Execute(ctx, "nonexistent-id", "tenant-123")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no encontrada")
}

func TestGetBrandByIDUseCase_Execute_WrongTenant_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewGetBrandByIDUseCase(mockRepo, brandMapper)

	mother := testentity.Create()
	existingBrand := mother.Nike()
	existingBrand.TenantID = "tenant-123"
	mockRepo.SetupBrands([]*entity.Brand{existingBrand})

	// Act - intentar acceder con otro tenant
	result, err := uc.Execute(ctx, existingBrand.ID, "tenant-999")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetBrandByIDUseCase_Execute_RepositoryError_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	mockRepo.ShouldFailOn("FindByID")
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewGetBrandByIDUseCase(mockRepo, brandMapper)

	// Act
	result, err := uc.Execute(ctx, "some-id", "tenant-123")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
}
