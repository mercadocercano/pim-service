package usecase

import (
	"context"
	"testing"

	"saas-mt-pim-service/src/quickstart/application/usecase"
	"saas-mt-pim-service/src/quickstart/domain/port"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockGetCategoriesByBusinessTypeRepository mock del repositorio
type MockGetCategoriesByBusinessTypeRepository struct {
	mock.Mock
}

func (m *MockGetCategoriesByBusinessTypeRepository) GetCategoriesByBusinessType(ctx context.Context, businessTypeSlug string) ([]port.CategoryByBusinessType, error) {
	args := m.Called(ctx, businessTypeSlug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]port.CategoryByBusinessType), args.Error(1)
}

func TestGetCategoriesByBusinessTypeUseCase_Execute_ReturnsCategoriesFromRepo(t *testing.T) {
	// Arrange
	ctx := context.Background()
	slug := "almacen"
	expected := []port.CategoryByBusinessType{
		{ID: "cat-1", Name: "Bebidas", Slug: "bebidas"},
		{ID: "cat-2", Name: "Lácteos", Slug: "lacteos"},
	}

	mockRepo := new(MockGetCategoriesByBusinessTypeRepository)
	mockRepo.On("GetCategoriesByBusinessType", ctx, slug).Return(expected, nil)

	uc := usecase.NewGetCategoriesByBusinessTypeUseCase(mockRepo)

	// Act
	result, err := uc.Execute(ctx, slug)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetCategoriesByBusinessTypeUseCase_Execute_BusinessTypeNotFound_ReturnsEmptySlice(t *testing.T) {
	// Arrange
	ctx := context.Background()
	slug := "tipo-inexistente"
	empty := []port.CategoryByBusinessType{}

	mockRepo := new(MockGetCategoriesByBusinessTypeRepository)
	mockRepo.On("GetCategoriesByBusinessType", ctx, slug).Return(empty, nil)

	uc := usecase.NewGetCategoriesByBusinessTypeUseCase(mockRepo)

	// Act
	result, err := uc.Execute(ctx, slug)

	// Assert
	assert.NoError(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}
