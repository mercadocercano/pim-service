package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pim/src/marketplace/application/request"
	"pim/src/marketplace/application/usecase"
	"pim/src/marketplace/domain/entity"
	"pim/src/shared/domain/criteria"
)

// MockMarketplaceCategoryRepository es un mock del repositorio para tests de actualización
type MockMarketplaceCategoryRepositoryUpdate struct {
	mock.Mock
}

func (m *MockMarketplaceCategoryRepositoryUpdate) Save(ctx context.Context, category *entity.MarketplaceCategory) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockMarketplaceCategoryRepositoryUpdate) GetByID(ctx context.Context, id string) (*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryUpdate) GetBySlug(ctx context.Context, slug string) (*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, slug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryUpdate) GetByParentID(ctx context.Context, parentID *string) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, parentID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryUpdate) GetRootCategories(ctx context.Context) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryUpdate) GetTree(ctx context.Context) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryUpdate) FindByCriteria(ctx context.Context, criteria criteria.Criteria) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, criteria)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryUpdate) CountByCriteria(ctx context.Context, criteria criteria.Criteria) (int, error) {
	args := m.Called(ctx, criteria)
	return args.Int(0), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryUpdate) Update(ctx context.Context, category *entity.MarketplaceCategory) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockMarketplaceCategoryRepositoryUpdate) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockMarketplaceCategoryRepositoryUpdate) ExistsBySlug(ctx context.Context, slug string) (bool, error) {
	args := m.Called(ctx, slug)
	return args.Bool(0), args.Error(1)
}

func (m *MockMarketplaceCategoryRepositoryUpdate) GetCategoryPath(ctx context.Context, categoryID string) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, categoryID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func TestUpdateMarketplaceCategoryUseCase_Execute_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepositoryUpdate)
	useCase := usecase.NewUpdateMarketplaceCategoryUseCase(mockRepo)

	categoryID := "cat-123"
	name := "Updated Electronics"
	slug := "updated-electronics"
	description := "Updated electronic products category"
	sortOrder := 5
	req := &request.UpdateMarketplaceCategoryRequest{
		Name:        &name,
		Slug:        &slug,
		Description: &description,
		SortOrder:   &sortOrder,
	}

	ctx := context.Background()

	// Create existing category
	existingCategory := &entity.MarketplaceCategory{
		ID:          categoryID,
		Name:        "Electronics",
		Slug:        "electronics",
		Description: "Electronic products category",
		Level:       0,
		IsActive:    true,
		SortOrder:   1,
		CreatedAt:   time.Now().Add(-24 * time.Hour),
		UpdatedAt:   time.Now().Add(-1 * time.Hour),
	}

	// Mock expectations
	mockRepo.On("GetByID", ctx, categoryID).Return(existingCategory, nil)
	mockRepo.On("ExistsBySlug", ctx, "updated-electronics").Return(false, nil)
	mockRepo.On("Save", ctx, mock.AnythingOfType("*entity.MarketplaceCategory")).Return(nil)

	// Act
	response, err := useCase.Execute(ctx, categoryID, req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, categoryID, response.ID)
	assert.Equal(t, "Updated Electronics", response.Name)
	assert.Equal(t, "updated-electronics", response.Slug)
	assert.Equal(t, &description, response.Description)
	assert.Equal(t, 5, response.SortOrder)
	assert.NotZero(t, response.UpdatedAt)

	mockRepo.AssertExpectations(t)
}

func TestUpdateMarketplaceCategoryUseCase_Execute_CategoryNotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepositoryUpdate)
	useCase := usecase.NewUpdateMarketplaceCategoryUseCase(mockRepo)

	categoryID := "nonexistent-cat"
	name := "Updated Electronics"
	slug := "updated-electronics"
	req := &request.UpdateMarketplaceCategoryRequest{
		Name: &name,
		Slug: &slug,
	}

	ctx := context.Background()

	// Mock expectations
	mockRepo.On("GetByID", ctx, categoryID).Return(nil, errors.New("category not found"))

	// Act
	response, err := useCase.Execute(ctx, categoryID, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "category not found")

	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "ExistsBySlug")
	mockRepo.AssertNotCalled(t, "Save")
}

func TestUpdateMarketplaceCategoryUseCase_Execute_ValidationError(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepositoryUpdate)
	useCase := usecase.NewUpdateMarketplaceCategoryUseCase(mockRepo)

	categoryID := "cat-123"
	name := "" // Empty name should cause validation error
	slug := "updated-electronics"
	req := &request.UpdateMarketplaceCategoryRequest{
		Name: &name,
		Slug: &slug,
	}

	ctx := context.Background()

	// Act
	response, err := useCase.Execute(ctx, categoryID, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "name")

	mockRepo.AssertNotCalled(t, "GetByID")
	mockRepo.AssertNotCalled(t, "ExistsBySlug")
	mockRepo.AssertNotCalled(t, "Save")
}

func TestUpdateMarketplaceCategoryUseCase_Execute_DuplicateSlug(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepositoryUpdate)
	useCase := usecase.NewUpdateMarketplaceCategoryUseCase(mockRepo)

	categoryID := "cat-123"
	name := "Updated Electronics"
	slug := "existing-slug"
	req := &request.UpdateMarketplaceCategoryRequest{
		Name: &name,
		Slug: &slug,
	}

	ctx := context.Background()

	// Create existing category
	existingCategory := &entity.MarketplaceCategory{
		ID:   categoryID,
		Name: "Electronics",
		Slug: "electronics",
	}

	// Mock expectations
	mockRepo.On("GetByID", ctx, categoryID).Return(existingCategory, nil)
	mockRepo.On("ExistsBySlug", ctx, "existing-slug").Return(true, nil)

	// Act
	response, err := useCase.Execute(ctx, categoryID, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "already exists")

	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Save")
}

func TestUpdateMarketplaceCategoryUseCase_Execute_RepositoryError_Update(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepositoryUpdate)
	useCase := usecase.NewUpdateMarketplaceCategoryUseCase(mockRepo)

	categoryID := "cat-123"
	name := "Updated Electronics"
	slug := "updated-electronics"
	req := &request.UpdateMarketplaceCategoryRequest{
		Name: &name,
		Slug: &slug,
	}

	ctx := context.Background()

	// Create existing category
	existingCategory := &entity.MarketplaceCategory{
		ID:   categoryID,
		Name: "Electronics",
		Slug: "electronics",
	}

	// Mock expectations
	mockRepo.On("GetByID", ctx, categoryID).Return(existingCategory, nil)
	mockRepo.On("ExistsBySlug", ctx, "updated-electronics").Return(false, nil)
	mockRepo.On("Save", ctx, mock.AnythingOfType("*entity.MarketplaceCategory")).Return(errors.New("failed to update category"))

	// Act
	response, err := useCase.Execute(ctx, categoryID, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "failed to update category")

	mockRepo.AssertExpectations(t)
}
