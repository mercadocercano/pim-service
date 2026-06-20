package usecase

import (
	"context"
	"errors"
	"testing"

	"saas-mt-pim-service/src/quickstart/application/usecase"
	"saas-mt-pim-service/src/quickstart/domain/port"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockListTemplatesRepository mock del repositorio de templates.
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

var (
	editorialTemplates = []port.ListTemplate{
		{ID: "t-1", Name: "Almacén", TotalProducts: 305, Brands: []port.ListTemplateBrand{{Name: "Marolio"}}},
	}
	computedTemplates = []port.ListTemplate{
		{ID: "t-1", Name: "Almacén", TotalProducts: 672, Brands: []port.ListTemplateBrand{{Name: "Arcor"}}},
	}
)

// Flag off → read-path editorial, nunca toca el computado.
func TestListTemplates_FlagOff_UsesEditorial(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockListTemplatesRepository)
	mockRepo.On("LoadTemplatesFromBusinessTypeTemplates", ctx).Return(editorialTemplates, nil)

	uc := usecase.NewListTemplatesUseCase(mockRepo, false)
	result, err := uc.Execute(ctx)

	assert.NoError(t, err)
	assert.Equal(t, 305, result.Templates[0].TotalProducts)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "LoadTemplatesComputed", ctx)
}

// Flag on → read-path computado.
func TestListTemplates_FlagOn_UsesComputed(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockListTemplatesRepository)
	mockRepo.On("LoadTemplatesComputed", ctx).Return(computedTemplates, nil)

	uc := usecase.NewListTemplatesUseCase(mockRepo, true)
	result, err := uc.Execute(ctx)

	assert.NoError(t, err)
	assert.Equal(t, 672, result.Templates[0].TotalProducts)
	assert.Equal(t, "Arcor", result.Templates[0].Brands[0].Name)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "LoadTemplatesFromBusinessTypeTemplates", ctx)
}

// Flag on pero el computado falla → fallback al editorial, sin romper el onboarding.
func TestListTemplates_FlagOn_ComputedFails_FallsBackToEditorial(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockListTemplatesRepository)
	mockRepo.On("LoadTemplatesComputed", ctx).Return(nil, errors.New("boom"))
	mockRepo.On("LoadTemplatesFromBusinessTypeTemplates", ctx).Return(editorialTemplates, nil)

	uc := usecase.NewListTemplatesUseCase(mockRepo, true)
	result, err := uc.Execute(ctx)

	assert.NoError(t, err)
	assert.Equal(t, 305, result.Templates[0].TotalProducts)
	mockRepo.AssertExpectations(t)
}
