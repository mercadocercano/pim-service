package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	cr "github.com/mercadocercano/criteria"
	"saas-mt-pim-service/src/category/application/usecase"
	"saas-mt-pim-service/src/category/domain/entity"
)

// MockMarketplaceCategoryRepositoryGetAll es un mock del repositorio para tests de obtener todas las categorías
type MockMarketplaceCategoryRepositoryGetAll struct {
	mock.Mock
}

func (m *MockMarketplaceCategoryRepositoryGetAll) Save(ctx context.Context, category *entity.MarketplaceCategory) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockMarketplaceCategoryRepositoryGetAll) GetByID(ctx context.Context, id string) (*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryGetAll) GetBySlug(ctx context.Context, slug string) (*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, slug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryGetAll) GetByParentID(ctx context.Context, parentID *string) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, parentID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryGetAll) GetRootCategories(ctx context.Context) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryGetAll) GetTree(ctx context.Context) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryGetAll) FindByCriteria(ctx context.Context, criteria cr.Criteria) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, criteria)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryGetAll) CountByCriteria(ctx context.Context, criteria cr.Criteria) (int, error) {
	args := m.Called(ctx, criteria)
	return args.Int(0), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryGetAll) Update(ctx context.Context, category *entity.MarketplaceCategory) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockMarketplaceCategoryRepositoryGetAll) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockMarketplaceCategoryRepositoryGetAll) ExistsBySlug(ctx context.Context, slug string) (bool, error) {
	args := m.Called(ctx, slug)
	return args.Bool(0), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryGetAll) GetCategoryPath(ctx context.Context, categoryID string) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, categoryID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func TestGetAllMarketplaceCategoriesUseCase_Execute_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepositoryGetAll)
	useCase := usecase.NewGetAllMarketplaceCategoriesUseCase(mockRepo)

	ctx := context.Background()

	// Create mock categories
	categories := []*entity.MarketplaceCategory{
		{
			ID:        "cat-1",
			Name:      "Electronics",
			Slug:      "electronics",
			Level:     0,
			IsActive:  true,
			SortOrder: 1,
		},
		{
			ID:        "cat-2",
			Name:      "Books",
			Slug:      "books",
			Level:     0,
			IsActive:  true,
			SortOrder: 2,
		},
	}

	// Mock expectations
	mockRepo.On("GetTree", ctx).Return(categories, nil)

	// Act
	response, err := useCase.Execute(ctx)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response, 2)
	assert.Equal(t, "Electronics", response[0].Name)
	assert.Equal(t, "Books", response[1].Name)

	mockRepo.AssertExpectations(t)
}

func TestGetAllMarketplaceCategoriesUseCase_Execute_EmptyResult(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepositoryGetAll)
	useCase := usecase.NewGetAllMarketplaceCategoriesUseCase(mockRepo)

	ctx := context.Background()

	// Mock expectations - no categories found
	mockRepo.On("GetTree", ctx).Return([]*entity.MarketplaceCategory{}, nil)

	// Act
	response, err := useCase.Execute(ctx)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response, 0)

	mockRepo.AssertExpectations(t)
}

func TestGetAllMarketplaceCategoriesUseCase_Execute_RepositoryError(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepositoryGetAll)
	useCase := usecase.NewGetAllMarketplaceCategoriesUseCase(mockRepo)

	ctx := context.Background()

	// Mock expectations - repository error
	mockRepo.On("GetTree", ctx).Return(nil, errors.New("database connection error"))

	// Act
	response, err := useCase.Execute(ctx)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "database connection error")

	mockRepo.AssertExpectations(t)
}
