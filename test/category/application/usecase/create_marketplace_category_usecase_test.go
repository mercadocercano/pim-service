package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pim/src/marketplace/application/request"
	"pim/src/marketplace/application/usecase"
	"pim/src/marketplace/domain/entity"
	"pim/src/shared/domain/criteria"
)

// MockMarketplaceCategoryRepository es un mock del repositorio
type MockMarketplaceCategoryRepository struct {
	mock.Mock
}

func (m *MockMarketplaceCategoryRepository) Save(ctx context.Context, category *entity.MarketplaceCategory) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockMarketplaceCategoryRepository) GetByID(ctx context.Context, id string) (*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepository) GetBySlug(ctx context.Context, slug string) (*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, slug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepository) GetByParentID(ctx context.Context, parentID *string) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, parentID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepository) GetRootCategories(ctx context.Context) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepository) GetTree(ctx context.Context) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepository) FindByCriteria(ctx context.Context, criteria criteria.Criteria) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, criteria)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func (m *MockMarketplaceCategoryRepository) CountByCriteria(ctx context.Context, criteria criteria.Criteria) (int, error) {
	args := m.Called(ctx, criteria)
	return args.Int(0), args.Error(1)
}

func (m *MockMarketplaceCategoryRepository) Update(ctx context.Context, category *entity.MarketplaceCategory) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockMarketplaceCategoryRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockMarketplaceCategoryRepository) ExistsBySlug(ctx context.Context, slug string) (bool, error) {
	args := m.Called(ctx, slug)
	return args.Bool(0), args.Error(1)
}

func (m *MockMarketplaceCategoryRepository) GetCategoryPath(ctx context.Context, categoryID string) ([]*entity.MarketplaceCategory, error) {
	args := m.Called(ctx, categoryID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.MarketplaceCategory), args.Error(1)
}

func TestCreateMarketplaceCategoryUseCase_Execute_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepository)
	useCase := usecase.NewCreateMarketplaceCategoryUseCase(mockRepo)

	description := "Electronic products category"
	req := &request.CreateMarketplaceCategoryRequest{
		Name:        "Electronics",
		Slug:        "electronics",
		Description: &description,
		SortOrder:   1,
	}

	ctx := context.Background()

	// Mock expectations
	mockRepo.On("ExistsBySlug", ctx, "electronics").Return(false, nil)
	mockRepo.On("Save", ctx, mock.AnythingOfType("*entity.MarketplaceCategory")).Run(func(args mock.Arguments) {
		// Simular que el repositorio asigna un ID
		category := args.Get(1).(*entity.MarketplaceCategory)
		category.ID = "cat-123"
	}).Return(nil)

	// Act
	response, err := useCase.Execute(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.ID)
	assert.Equal(t, "Electronics", response.Name)
	assert.Equal(t, "electronics", response.Slug)
	assert.Equal(t, &description, response.Description)
	assert.Equal(t, 0, response.Level) // Root category
	assert.True(t, response.IsActive)
	assert.Equal(t, 1, response.SortOrder)
	assert.NotZero(t, response.CreatedAt)

	mockRepo.AssertExpectations(t)
}

func TestCreateMarketplaceCategoryUseCase_Execute_ValidationError_EmptyName(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepository)
	useCase := usecase.NewCreateMarketplaceCategoryUseCase(mockRepo)

	req := &request.CreateMarketplaceCategoryRequest{
		Name: "", // Empty name should cause validation error
		Slug: "electronics",
	}

	ctx := context.Background()

	// Act
	response, err := useCase.Execute(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "name")

	mockRepo.AssertNotCalled(t, "ExistsBySlug")
	mockRepo.AssertNotCalled(t, "Save")
}

func TestCreateMarketplaceCategoryUseCase_Execute_ValidationError_EmptySlug(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepository)
	useCase := usecase.NewCreateMarketplaceCategoryUseCase(mockRepo)

	req := &request.CreateMarketplaceCategoryRequest{
		Name: "Electronics",
		Slug: "", // Empty slug should cause validation error
	}

	ctx := context.Background()

	// Act
	response, err := useCase.Execute(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "slug")

	mockRepo.AssertNotCalled(t, "ExistsBySlug")
	mockRepo.AssertNotCalled(t, "Save")
}

func TestCreateMarketplaceCategoryUseCase_Execute_DuplicateSlug(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepository)
	useCase := usecase.NewCreateMarketplaceCategoryUseCase(mockRepo)

	req := &request.CreateMarketplaceCategoryRequest{
		Name: "Electronics",
		Slug: "electronics",
	}

	ctx := context.Background()

	// Mock expectations - slug already exists
	mockRepo.On("ExistsBySlug", ctx, "electronics").Return(true, nil)

	// Act
	response, err := useCase.Execute(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "already exists")

	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Save")
}

func TestCreateMarketplaceCategoryUseCase_Execute_RepositoryError_CheckExists(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepository)
	useCase := usecase.NewCreateMarketplaceCategoryUseCase(mockRepo)

	req := &request.CreateMarketplaceCategoryRequest{
		Name: "Electronics",
		Slug: "electronics",
	}

	ctx := context.Background()

	// Mock expectations - repository error when checking if exists
	mockRepo.On("ExistsBySlug", ctx, "electronics").Return(false, errors.New("database connection error"))

	// Act
	response, err := useCase.Execute(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "database connection error")

	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Save")
}

func TestCreateMarketplaceCategoryUseCase_Execute_RepositoryError_Save(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepository)
	useCase := usecase.NewCreateMarketplaceCategoryUseCase(mockRepo)

	req := &request.CreateMarketplaceCategoryRequest{
		Name: "Electronics",
		Slug: "electronics",
	}

	ctx := context.Background()

	// Mock expectations
	mockRepo.On("ExistsBySlug", ctx, "electronics").Return(false, nil)
	mockRepo.On("Save", ctx, mock.AnythingOfType("*entity.MarketplaceCategory")).Return(errors.New("failed to save category"))

	// Act
	response, err := useCase.Execute(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "failed to save category")

	mockRepo.AssertExpectations(t)
}

func TestCreateMarketplaceCategoryUseCase_Execute_WithParentCategory(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepository)
	useCase := usecase.NewCreateMarketplaceCategoryUseCase(mockRepo)

	parentID := "parent-123"
	req := &request.CreateMarketplaceCategoryRequest{
		Name:     "Smartphones",
		Slug:     "smartphones",
		ParentID: &parentID,
	}

	ctx := context.Background()

	// Create mock parent category
	parentCategory := &entity.MarketplaceCategory{
		ID:    "parent-123",
		Name:  "Electronics",
		Level: 0,
	}

	// Mock expectations
	mockRepo.On("ExistsBySlug", ctx, "smartphones").Return(false, nil)
	mockRepo.On("GetByID", ctx, "parent-123").Return(parentCategory, nil)
	mockRepo.On("Save", ctx, mock.AnythingOfType("*entity.MarketplaceCategory")).Return(nil)

	// Act
	response, err := useCase.Execute(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Smartphones", response.Name)
	assert.Equal(t, "smartphones", response.Slug)
	assert.Equal(t, &parentID, response.ParentID)
	assert.Equal(t, 1, response.Level) // Child category level

	mockRepo.AssertExpectations(t)
}

func TestCreateMarketplaceCategoryUseCase_Execute_MaxDepthExceeded(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepository)
	useCase := usecase.NewCreateMarketplaceCategoryUseCase(mockRepo)

	parentID := "parent-123"
	req := &request.CreateMarketplaceCategoryRequest{
		Name:     "Sub-subcategory",
		Slug:     "sub-subcategory",
		ParentID: &parentID,
	}

	ctx := context.Background()

	// Create mock parent category at max level
	parentCategory := &entity.MarketplaceCategory{
		ID:    "parent-123",
		Name:  "Subcategory",
		Level: 2, // Already at max level
	}

	// Mock expectations
	mockRepo.On("ExistsBySlug", ctx, "sub-subcategory").Return(false, nil)
	mockRepo.On("GetByID", ctx, "parent-123").Return(parentCategory, nil)

	// Act
	response, err := useCase.Execute(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "maximum category depth")

	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Save")
}

func TestCreateMarketplaceCategoryUseCase_Execute_ParentNotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockMarketplaceCategoryRepository)
	useCase := usecase.NewCreateMarketplaceCategoryUseCase(mockRepo)

	parentID := "nonexistent-parent"
	req := &request.CreateMarketplaceCategoryRequest{
		Name:     "Smartphones",
		Slug:     "smartphones",
		ParentID: &parentID,
	}

	ctx := context.Background()

	// Mock expectations
	mockRepo.On("ExistsBySlug", ctx, "smartphones").Return(false, nil)
	mockRepo.On("GetByID", ctx, "nonexistent-parent").Return(nil, errors.New("category not found"))

	// Act
	response, err := useCase.Execute(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "parent category not found")

git status	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Save")
}
