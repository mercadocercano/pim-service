package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/brand/application/mapper"
	"saas-mt-pim-service/src/brand/application/request"
	"saas-mt-pim-service/src/brand/application/usecase"
	"saas-mt-pim-service/src/brand/domain/entity"
	"saas-mt-pim-service/src/brand/domain/service"
	testentity "saas-mt-pim-service/test/brand/domain/entity"
	"saas-mt-pim-service/test/brand/infrastructure/persistence/repository"
)

func TestUpdateBrandUseCase_Execute_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewUpdateBrandUseCase(mockRepo, brandService, brandMapper)

	mother := testentity.Create()
	existingBrand := mother.Nike()
	existingBrand.TenantID = "tenant-123"
	mockRepo.SetupBrands([]*entity.Brand{existingBrand})

	req := &request.UpdateBrandRequest{
		Name:        "Nike Updated",
		Description: "Descripción actualizada",
		LogoURL:     strPtr("https://example.com/new-logo.png"),
		Website:     strPtr("https://nike.com/new"),
	}

	// Act
	result, err := uc.Execute(ctx, existingBrand.ID, req, "tenant-123")

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Nike Updated", result.Name)
	assert.Equal(t, "Descripción actualizada", result.Description)
	assert.Equal(t, "https://example.com/new-logo.png", *result.LogoURL)
	assert.Equal(t, "https://nike.com/new", *result.Website)
	assert.Equal(t, 1, mockRepo.GetCallCount("Update"))
}

func TestUpdateBrandUseCase_Execute_BrandNotFound_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewUpdateBrandUseCase(mockRepo, brandService, brandMapper)

	req := &request.UpdateBrandRequest{
		Name:        "Updated",
		Description: "Test",
	}

	// Act
	result, err := uc.Execute(ctx, "nonexistent-id", req, "tenant-123")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no encontrada")
}

func TestUpdateBrandUseCase_Execute_EmptyName_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewUpdateBrandUseCase(mockRepo, brandService, brandMapper)

	mother := testentity.Create()
	existingBrand := mother.WithDefaults()
	existingBrand.TenantID = "tenant-123"
	mockRepo.SetupBrands([]*entity.Brand{existingBrand})

	req := &request.UpdateBrandRequest{
		Name:        "",
		Description: "Sin nombre",
	}

	// Act
	result, err := uc.Execute(ctx, existingBrand.ID, req, "tenant-123")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateBrandUseCase_Execute_DuplicateName_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewUpdateBrandUseCase(mockRepo, brandService, brandMapper)

	mother := testentity.Create()
	brand1 := mother.Nike()
	brand1.TenantID = "tenant-123"
	brand2 := mother.Apple()
	brand2.TenantID = "tenant-123"
	mockRepo.SetupBrands([]*entity.Brand{brand1, brand2})

	// Intentar renombrar brand2 al nombre de brand1
	req := &request.UpdateBrandRequest{
		Name:        "Nike",
		Description: "Intentar duplicar nombre",
	}

	// Act
	result, err := uc.Execute(ctx, brand2.ID, req, "tenant-123")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "ya existe")
}

func TestUpdateBrandUseCase_Execute_RepositoryError_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	mockRepo.ShouldFailOn("Update")
	brandService := service.NewBrandDomainService(mockRepo)
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewUpdateBrandUseCase(mockRepo, brandService, brandMapper)

	mother := testentity.Create()
	existingBrand := mother.WithDefaults()
	existingBrand.TenantID = "tenant-123"
	mockRepo.SetupBrands([]*entity.Brand{existingBrand})

	req := &request.UpdateBrandRequest{
		Name:        "Updated",
		Description: "Test",
	}

	// Act
	result, err := uc.Execute(ctx, existingBrand.ID, req, "tenant-123")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
}
