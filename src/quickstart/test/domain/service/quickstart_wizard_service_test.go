package service

import (
	"context"
	"encoding/json"
	"testing"

	"pim/src/quickstart/domain/service"
	"pim/src/quickstart/test/domain/entity"
	businessTypeTestEntity "pim/src/quickstart/test/domain/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBusinessTypeRepository mock del repositorio de business types
type MockBusinessTypeRepository struct {
	mock.Mock
}

func (m *MockBusinessTypeRepository) FindByID(ctx context.Context, id string) (interface{}, error) {
	args := m.Called(ctx, id)
	return args.Get(0), args.Error(1)
}

// MockBusinessTypeTemplateRepository mock del repositorio de templates
type MockBusinessTypeTemplateRepository struct {
	mock.Mock
}

func (m *MockBusinessTypeTemplateRepository) FindByBusinessTypeID(ctx context.Context, businessTypeID string) (interface{}, error) {
	args := m.Called(ctx, businessTypeID)
	return args.Get(0), args.Error(1)
}

// MockTenantSetupRepository mock del repositorio de tenant setup
type MockTenantSetupRepository struct {
	mock.Mock
}

func (m *MockTenantSetupRepository) Create(ctx context.Context, history interface{}) error {
	args := m.Called(ctx, history)
	return args.Error(0)
}

func (m *MockTenantSetupRepository) GetLatestByTenantID(ctx context.Context, tenantID string) (interface{}, error) {
	args := m.Called(ctx, tenantID)
	return args.Get(0), args.Error(1)
}

func (m *MockTenantSetupRepository) Update(ctx context.Context, history interface{}) error {
	args := m.Called(ctx, history)
	return args.Error(0)
}

func TestQuickstartWizardService_StartWizard_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()
	businessTypeID := uuid.New().String()

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	// Configurar mocks
	businessType := struct {
		ID   string
		Code string
		Name string
	}{
		ID:   businessTypeID,
		Code: "polirubro",
		Name: "Polirubro",
	}

	templateOM := businessTypeTestEntity.NewBusinessTypeTemplateObjectMother()
	template := templateOM.WithBusinessTypeID(businessTypeID)
	templates := []*interface{}{&template}

	mockBusinessTypeRepo.On("FindByID", ctx, businessTypeID).Return(&businessType, nil)
	mockTemplateRepo.On("FindByBusinessTypeID", ctx, businessTypeID).Return(templates, nil)
	mockTenantSetupRepo.On("Create", ctx, mock.AnythingOfType("*entity.TenantQuickstartHistory")).Return(nil)

	service := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)

	// Act
	result, err := service.StartWizard(ctx, tenantID, businessTypeID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tenantID, result.TenantID)
	assert.Equal(t, businessTypeID, result.BusinessTypeID)
	assert.False(t, result.SetupCompleted)

	// Verificar que el setup data contiene los campos esperados
	var setupData map[string]interface{}
	err = json.Unmarshal(result.SetupData, &setupData)
	assert.NoError(t, err)
	assert.Equal(t, businessTypeID, setupData["business_type_id"])
	assert.Equal(t, "business_type_selected", setupData["step"])
	assert.Contains(t, setupData["completed_steps"], "business_type_selected")

	mockBusinessTypeRepo.AssertExpectations(t)
	mockTemplateRepo.AssertExpectations(t)
	mockTenantSetupRepo.AssertExpectations(t)
}

func TestQuickstartWizardService_GetWizardStatus_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	historyOM := entity.NewTenantQuickstartHistoryObjectMother()
	expectedHistory := historyOM.WithTenantID(tenantID)

	mockTenantSetupRepo.On("GetLatestByTenantID", ctx, tenantID).Return(expectedHistory, nil)

	service := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)

	// Act
	result, err := service.GetWizardStatus(ctx, tenantID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tenantID, result.TenantID)

	mockTenantSetupRepo.AssertExpectations(t)
}

func TestQuickstartWizardService_UpdateWizardStep_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()
	currentStep := "categories_selection"
	stepData := map[string]interface{}{
		"selected_categories": []string{"electronics", "home"},
	}

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	historyOM := entity.NewTenantQuickstartHistoryObjectMother()
	existingHistory := historyOM.WithTenantID(tenantID)

	mockTenantSetupRepo.On("GetLatestByTenantID", ctx, tenantID).Return(existingHistory, nil)
	mockTenantSetupRepo.On("Update", ctx, mock.AnythingOfType("*entity.TenantQuickstartHistory")).Return(nil)

	service := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)

	// Act
	result, err := service.UpdateWizardStep(ctx, tenantID, currentStep, stepData)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tenantID, result.TenantID)

	// Verificar que el setup data se actualizó correctamente
	var updatedSetupData map[string]interface{}
	err = json.Unmarshal(result.SetupData, &updatedSetupData)
	assert.NoError(t, err)
	assert.Equal(t, currentStep, updatedSetupData["step"])

	mockTenantSetupRepo.AssertExpectations(t)
}

func TestQuickstartWizardService_CompleteWizard_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()
	finalSelections := map[string]interface{}{
		"final_categories": []string{"electronics", "home"},
		"final_brands":     []string{"Samsung", "LG"},
		"final_products":   []int{1, 2, 3},
	}

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	historyOM := entity.NewTenantQuickstartHistoryObjectMother()
	existingHistory := historyOM.InProgress()
	existingHistory.TenantID = tenantID

	mockTenantSetupRepo.On("GetLatestByTenantID", ctx, tenantID).Return(existingHistory, nil)
	mockTenantSetupRepo.On("Update", ctx, mock.AnythingOfType("*entity.TenantQuickstartHistory")).Return(nil)

	service := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)

	// Act
	result, err := service.CompleteWizard(ctx, tenantID, finalSelections)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tenantID, result.TenantID)

	// Verificar que el wizard se marcó como completado
	var completedSetupData map[string]interface{}
	err = json.Unmarshal(result.SetupData, &completedSetupData)
	assert.NoError(t, err)
	assert.Equal(t, "completed", completedSetupData["step"])
	assert.Equal(t, true, completedSetupData["completed"])
	assert.Equal(t, finalSelections, completedSetupData["final_selections"])

	mockTenantSetupRepo.AssertExpectations(t)
}

func TestQuickstartWizardService_GetTemplateData_AllSections_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	businessTypeID := uuid.New().String()
	section := ""

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	templateOM := businessTypeTestEntity.NewBusinessTypeTemplateObjectMother()
	template := templateOM.WithBusinessTypeID(businessTypeID)
	templates := []*interface{}{&template}

	mockTemplateRepo.On("FindByBusinessTypeID", ctx, businessTypeID).Return(templates, nil)

	service := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)

	// Act
	result, err := service.GetTemplateData(ctx, businessTypeID, section)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verificar estructura del template data
	templateData, ok := result.(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, templateData, "categories")
	assert.Contains(t, templateData, "attributes")
	assert.Contains(t, templateData, "products")
	assert.Contains(t, templateData, "brands")
	assert.Contains(t, templateData, "metadata")

	mockTemplateRepo.AssertExpectations(t)
}

func TestQuickstartWizardService_GetTemplateData_SpecificSection_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	businessTypeID := uuid.New().String()
	section := "categories"

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	templateOM := businessTypeTestEntity.NewBusinessTypeTemplateObjectMother()
	template := templateOM.WithBusinessTypeID(businessTypeID)
	templates := []*interface{}{&template}

	mockTemplateRepo.On("FindByBusinessTypeID", ctx, businessTypeID).Return(templates, nil)

	service := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)

	// Act
	result, err := service.GetTemplateData(ctx, businessTypeID, section)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	mockTemplateRepo.AssertExpectations(t)
}

func TestQuickstartWizardService_StartWizard_EmptyTenantID_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := ""
	businessTypeID := uuid.New().String()

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	service := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)

	// Act
	result, err := service.StartWizard(ctx, tenantID, businessTypeID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "tenant ID es requerido")
}

func TestQuickstartWizardService_StartWizard_BusinessTypeNotFound_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()
	businessTypeID := uuid.New().String()

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	mockBusinessTypeRepo.On("FindByID", ctx, businessTypeID).Return(nil, nil)

	service := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)

	// Act
	result, err := service.StartWizard(ctx, tenantID, businessTypeID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no encontrado")

	mockBusinessTypeRepo.AssertExpectations(t)
}