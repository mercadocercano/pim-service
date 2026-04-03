package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/brand/application/mapper"
	"saas-mt-pim-service/src/brand/application/request"
	"saas-mt-pim-service/src/brand/application/usecase"
	"saas-mt-pim-service/src/brand/domain/service"
	"saas-mt-pim-service/test/brand/infrastructure/persistence/repository"
)

func TestCreateBrandUseCase_Execute_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewCreateBrandUseCase(mockRepo, brandService, brandMapper)

	req := &request.CreateBrandRequest{
		Name:        "Nike",
		Description: "Marca deportiva internacional",
		LogoURL:     strPtr("https://example.com/nike.png"),
		Website:     strPtr("https://nike.com"),
	}

	// Act
	result, err := uc.Execute(ctx, req, "tenant-123")

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Nike", result.Name)
	assert.Equal(t, "Marca deportiva internacional", result.Description)
	assert.NotNil(t, result.LogoURL)
	assert.Equal(t, "https://example.com/nike.png", *result.LogoURL)
	assert.NotNil(t, result.Website)
	assert.Equal(t, "https://nike.com", *result.Website)
	assert.Equal(t, "active", result.Status)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, 1, mockRepo.GetCallCount("Create"))
}

func TestCreateBrandUseCase_Execute_WithoutOptionalFields(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewCreateBrandUseCase(mockRepo, brandService, brandMapper)

	req := &request.CreateBrandRequest{
		Name:        "Marca Simple",
		Description: "Sin logo ni web",
	}

	// Act
	result, err := uc.Execute(ctx, req, "tenant-123")

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Marca Simple", result.Name)
	assert.Nil(t, result.LogoURL)
	assert.Nil(t, result.Website)
}

func TestCreateBrandUseCase_Execute_DuplicateName_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewCreateBrandUseCase(mockRepo, brandService, brandMapper)

	// Crear primera marca
	firstReq := &request.CreateBrandRequest{
		Name:        "Nike",
		Description: "Primera",
	}
	_, err := uc.Execute(ctx, firstReq, "tenant-123")
	require.NoError(t, err)

	// Intentar crear con mismo nombre
	duplicateReq := &request.CreateBrandRequest{
		Name:        "Nike",
		Description: "Duplicada",
	}

	// Act
	result, err := uc.Execute(ctx, duplicateReq, "tenant-123")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "ya existe")
}

func TestCreateBrandUseCase_Execute_EmptyName_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewCreateBrandUseCase(mockRepo, brandService, brandMapper)

	req := &request.CreateBrandRequest{
		Name:        "",
		Description: "Sin nombre",
	}

	// Act
	result, err := uc.Execute(ctx, req, "tenant-123")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, 0, mockRepo.GetCallCount("Create"))
}

func TestCreateBrandUseCase_Execute_RepositoryError_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	mockRepo.ShouldFailOn("Create")
	brandService := service.NewBrandDomainService(mockRepo)
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewCreateBrandUseCase(mockRepo, brandService, brandMapper)

	req := &request.CreateBrandRequest{
		Name:        "Nike",
		Description: "Marca deportiva",
	}

	// Act
	result, err := uc.Execute(ctx, req, "tenant-123")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCreateBrandUseCase_Execute_SameNameDifferentTenant_ShouldSucceed(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandService := service.NewBrandDomainService(mockRepo)
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewCreateBrandUseCase(mockRepo, brandService, brandMapper)

	req1 := &request.CreateBrandRequest{
		Name:        "Nike",
		Description: "Tenant 1",
	}
	_, err := uc.Execute(ctx, req1, "tenant-1")
	require.NoError(t, err)

	req2 := &request.CreateBrandRequest{
		Name:        "Nike",
		Description: "Tenant 2",
	}

	// Act
	result, err := uc.Execute(ctx, req2, "tenant-2")

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Nike", result.Name)
}

func strPtr(s string) *string {
	return &s
}
