package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cr "github.com/mercadocercano/criteria"
	"saas-mt-pim-service/src/brand/application/mapper"
	"saas-mt-pim-service/src/brand/application/usecase"
	"saas-mt-pim-service/src/brand/domain/entity"
	testentity "saas-mt-pim-service/test/brand/domain/entity"
	"saas-mt-pim-service/test/brand/infrastructure/persistence/repository"
)

func TestListBrandsByCriteriaUseCase_Execute_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewListBrandsByCriteriaUseCase(mockRepo, brandMapper)

	mother := testentity.Create()
	brand1 := mother.Nike()
	brand2 := mother.Apple()
	mockRepo.SetupBrands([]*entity.Brand{brand1, brand2})

	criteria := cr.Criteria{
		Pagination: cr.Pagination{
			Page:     1,
			PageSize: 10,
		},
	}

	// Act
	result, err := uc.Execute(ctx, criteria)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Items, 2)
	assert.Equal(t, 1, result.Page)
	assert.Equal(t, 10, result.PageSize)
	assert.Equal(t, 1, mockRepo.GetCallCount("SearchByCriteria"))
	assert.Equal(t, 1, mockRepo.GetCallCount("CountByCriteria"))
}

func TestListBrandsByCriteriaUseCase_Execute_EmptyResult(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewListBrandsByCriteriaUseCase(mockRepo, brandMapper)

	criteria := cr.Criteria{
		Pagination: cr.Pagination{
			Page:     1,
			PageSize: 10,
		},
	}

	// Act
	result, err := uc.Execute(ctx, criteria)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, result.TotalCount)
	assert.Empty(t, result.Items)
}

func TestListBrandsByCriteriaUseCase_Execute_SearchError_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	mockRepo.ShouldFailOn("SearchByCriteria")
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewListBrandsByCriteriaUseCase(mockRepo, brandMapper)

	criteria := cr.Criteria{
		Pagination: cr.Pagination{
			Page:     1,
			PageSize: 10,
		},
	}

	// Act
	result, err := uc.Execute(ctx, criteria)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListBrandsByCriteriaUseCase_Execute_CountError_ShouldFail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	mockRepo.ShouldFailOn("CountByCriteria")
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewListBrandsByCriteriaUseCase(mockRepo, brandMapper)

	criteria := cr.Criteria{
		Pagination: cr.Pagination{
			Page:     1,
			PageSize: 10,
		},
	}

	// Act
	result, err := uc.Execute(ctx, criteria)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListBrandsByCriteriaUseCase_Execute_PaginationCalculation(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := repository.NewMockBrandRepository()
	brandMapper := mapper.NewBrandMapper()
	uc := usecase.NewListBrandsByCriteriaUseCase(mockRepo, brandMapper)

	mother := testentity.Create()
	brands := make([]*entity.Brand, 5)
	for i := 0; i < 5; i++ {
		brands[i] = mother.WithDefaults()
	}
	mockRepo.SetupBrands(brands)

	criteria := cr.Criteria{
		Pagination: cr.Pagination{
			Page:     1,
			PageSize: 2,
		},
	}

	// Act
	result, err := uc.Execute(ctx, criteria)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 5, result.TotalCount)
	assert.Equal(t, 3, result.TotalPages) // 5 items / 2 per page = 3 pages
}
