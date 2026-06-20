package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/quickstart/application/usecase"
	"saas-mt-pim-service/src/quickstart/domain/port"
)

// MockListTemplatesRepository es un mock del repositorio de templates
type MockListTemplatesRepository struct {
	mock.Mock
}

func (m *MockListTemplatesRepository) LoadTemplatesFromBusinessTypeTemplates(ctx context.Context) ([]port.ListTemplate, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]port.ListTemplate), args.Error(1)
}

func (m *MockListTemplatesRepository) LoadTemplatesComputed(ctx context.Context) ([]port.ListTemplate, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]port.ListTemplate), args.Error(1)
}

func TestListTemplatesUseCase_Execute_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockListTemplatesRepository)
	uc := usecase.NewListTemplatesUseCase(mockRepo, false)

	templates := []port.ListTemplate{
		{
			ID:          "tpl-1",
			Name:        "Ferretería",
			Slug:        "ferreteria",
			Description: "Template para ferreterías",
			Categories:  []string{"herramientas", "electricidad"},
			IsActive:    true,
		},
		{
			ID:          "tpl-2",
			Name:        "Almacén",
			Slug:        "almacen",
			Description: "Template para almacenes",
			Categories:  []string{"bebidas", "alimentos"},
			IsActive:    true,
		},
	}
	mockRepo.On("LoadTemplatesFromBusinessTypeTemplates", mock.Anything).Return(templates, nil)

	// Act
	result, err := uc.Execute(context.Background())

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.Total)
	assert.Len(t, result.Templates, 2)
	assert.Equal(t, "Ferretería", result.Templates[0].Name)
	assert.Equal(t, "ferreteria", result.Templates[0].Slug)
	assert.True(t, result.Templates[0].IsActive)
	assert.Equal(t, []string{"herramientas", "electricidad"}, result.Templates[0].Categories)
	mockRepo.AssertExpectations(t)
}

func TestListTemplatesUseCase_Execute_EmptyResult(t *testing.T) {
	// Arrange
	mockRepo := new(MockListTemplatesRepository)
	uc := usecase.NewListTemplatesUseCase(mockRepo, false)

	mockRepo.On("LoadTemplatesFromBusinessTypeTemplates", mock.Anything).Return([]port.ListTemplate{}, nil)

	// Act
	result, err := uc.Execute(context.Background())

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, result.Total)
	assert.Empty(t, result.Templates)
	mockRepo.AssertExpectations(t)
}

func TestListTemplatesUseCase_Execute_RepositoryError_ShouldFail(t *testing.T) {
	// Arrange
	mockRepo := new(MockListTemplatesRepository)
	uc := usecase.NewListTemplatesUseCase(mockRepo, false)

	mockRepo.On("LoadTemplatesFromBusinessTypeTemplates", mock.Anything).Return(nil, errors.New("database error"))

	// Act
	result, err := uc.Execute(context.Background())

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "database error")
	mockRepo.AssertExpectations(t)
}

func TestListTemplatesUseCase_Execute_MapsAllFields(t *testing.T) {
	// Arrange
	mockRepo := new(MockListTemplatesRepository)
	uc := usecase.NewListTemplatesUseCase(mockRepo, false)

	templates := []port.ListTemplate{
		{
			ID:          "tpl-abc",
			Name:        "Kiosco",
			Slug:        "kiosco",
			Description: "Template kiosco argentino",
			Categories:  []string{"golosinas", "bebidas", "snacks"},
			IsActive:    false,
		},
	}
	mockRepo.On("LoadTemplatesFromBusinessTypeTemplates", mock.Anything).Return(templates, nil)

	// Act
	result, err := uc.Execute(context.Background())

	// Assert
	require.NoError(t, err)
	require.Len(t, result.Templates, 1)
	tpl := result.Templates[0]
	assert.Equal(t, "tpl-abc", tpl.ID)
	assert.Equal(t, "Kiosco", tpl.Name)
	assert.Equal(t, "kiosco", tpl.Slug)
	assert.Equal(t, "Template kiosco argentino", tpl.Description)
	assert.Equal(t, []string{"golosinas", "bebidas", "snacks"}, tpl.Categories)
	assert.False(t, tpl.IsActive)
}
