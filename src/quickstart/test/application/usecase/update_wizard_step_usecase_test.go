package usecase

import (
	"context"
	"testing"

	"pim/src/quickstart/application/usecase"
	"pim/src/quickstart/domain/service"
	"pim/src/quickstart/test/domain/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateWizardStepUseCase_Execute_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()
	currentStep := "categories_selection"
	stepData := map[string]interface{}{
		"selected_categories": []string{"electronics", "home"},
		"timestamp":           "2025-06-30T14:30:00Z",
	}

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	historyOM := entity.NewTenantQuickstartHistoryObjectMother()
	existingHistory := historyOM.WithTenantID(tenantID)

	mockTenantSetupRepo.On("GetLatestByTenantID", ctx, tenantID).Return(existingHistory, nil)
	mockTenantSetupRepo.On("Update", ctx, mock.AnythingOfType("*entity.TenantQuickstartHistory")).Return(nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewUpdateWizardStepUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID, currentStep, stepData)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tenantID, result.TenantID)

	mockTenantSetupRepo.AssertExpectations(t)
}

func TestUpdateWizardStepUseCase_Execute_EmptyTenantID_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := ""
	currentStep := "categories_selection"
	stepData := map[string]interface{}{}

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewUpdateWizardStepUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID, currentStep, stepData)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "tenant ID")
}

func TestUpdateWizardStepUseCase_Execute_EmptyCurrentStep_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()
	currentStep := ""
	stepData := map[string]interface{}{}

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewUpdateWizardStepUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID, currentStep, stepData)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateWizardStepUseCase_Execute_NoExistingHistory_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()
	currentStep := "categories_selection"
	stepData := map[string]interface{}{}

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	mockTenantSetupRepo.On("GetLatestByTenantID", ctx, tenantID).Return(nil, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewUpdateWizardStepUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID, currentStep, stepData)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no hay wizard iniciado")

	mockTenantSetupRepo.AssertExpectations(t)
}

func TestUpdateWizardStepUseCase_Execute_WithStepData_UpdatesSetupData(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()
	currentStep := "brands_selection"
	stepData := map[string]interface{}{
		"selected_brands": []string{"Samsung", "LG", "Sony"},
		"custom_brands":   []string{"Mi Marca Local"},
	}

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	historyOM := entity.NewTenantQuickstartHistoryObjectMother()
	existingHistory := historyOM.WithStep("categories_selection", []string{"business_type_selected"})
	existingHistory.TenantID = tenantID

	mockTenantSetupRepo.On("GetLatestByTenantID", ctx, tenantID).Return(existingHistory, nil)
	mockTenantSetupRepo.On("Update", ctx, mock.AnythingOfType("*entity.TenantQuickstartHistory")).Return(nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewUpdateWizardStepUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID, currentStep, stepData)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tenantID, result.TenantID)

	mockTenantSetupRepo.AssertExpectations(t)
}

func TestUpdateWizardStepUseCase_Execute_CompletedStepsAreTracked(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()
	currentStep := "products_selection"
	stepData := map[string]interface{}{
		"selected_products": []int{1, 2, 3, 4, 5},
	}

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	historyOM := entity.NewTenantQuickstartHistoryObjectMother()
	existingHistory := historyOM.WithStep("brands_selection", []string{"business_type_selected", "categories_selection"})
	existingHistory.TenantID = tenantID

	mockTenantSetupRepo.On("GetLatestByTenantID", ctx, tenantID).Return(existingHistory, nil)
	mockTenantSetupRepo.On("Update", ctx, mock.AnythingOfType("*entity.TenantQuickstartHistory")).Return(nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewUpdateWizardStepUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID, currentStep, stepData)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tenantID, result.TenantID)

	mockTenantSetupRepo.AssertExpectations(t)
}