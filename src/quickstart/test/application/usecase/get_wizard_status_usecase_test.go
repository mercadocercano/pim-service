package usecase

import (
	"context"
	"testing"

	"saas-mt-pim-service/src/quickstart/application/usecase"
	"saas-mt-pim-service/src/quickstart/domain/service"
	"saas-mt-pim-service/src/quickstart/test/domain/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetWizardStatusUseCase_Execute_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	historyOM := entity.NewTenantQuickstartHistoryObjectMother()
	expectedHistory := historyOM.WithTenantID(tenantID)

	mockTenantSetupRepo.On("GetLatestByTenantID", ctx, tenantID).Return(expectedHistory, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewGetWizardStatusUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tenantID, result.TenantID)
	assert.Equal(t, expectedHistory.ID, result.ID)

	mockTenantSetupRepo.AssertExpectations(t)
}

func TestGetWizardStatusUseCase_Execute_EmptyTenantID_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := ""

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewGetWizardStatusUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "tenant ID es requerido")
}

func TestGetWizardStatusUseCase_Execute_NoHistoryFound_ReturnsNil(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	mockTenantSetupRepo.On("GetLatestByTenantID", ctx, tenantID).Return(nil, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewGetWizardStatusUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID)

	// Assert
	assert.NoError(t, err)
	assert.Nil(t, result)

	mockTenantSetupRepo.AssertExpectations(t)
}

func TestGetWizardStatusUseCase_Execute_CompletedWizard_ReturnsCompletedStatus(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	historyOM := entity.NewTenantQuickstartHistoryObjectMother()
	expectedHistory := historyOM.Completed()
	expectedHistory.TenantID = tenantID

	mockTenantSetupRepo.On("GetLatestByTenantID", ctx, tenantID).Return(expectedHistory, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewGetWizardStatusUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tenantID, result.TenantID)
	assert.True(t, result.SetupCompleted)

	mockTenantSetupRepo.AssertExpectations(t)
}

func TestGetWizardStatusUseCase_Execute_InProgressWizard_ReturnsInProgressStatus(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	historyOM := entity.NewTenantQuickstartHistoryObjectMother()
	expectedHistory := historyOM.InProgress()
	expectedHistory.TenantID = tenantID

	mockTenantSetupRepo.On("GetLatestByTenantID", ctx, tenantID).Return(expectedHistory, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewGetWizardStatusUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tenantID, result.TenantID)
	assert.False(t, result.SetupCompleted)
	assert.True(t, result.IsPending())

	mockTenantSetupRepo.AssertExpectations(t)
}