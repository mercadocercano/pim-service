package usecase

import (
	"context"
	"testing"

	"pim/src/quickstart/application/usecase"
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

func (m *MockBusinessTypeRepository) Create(ctx context.Context, businessType interface{}) error {
	args := m.Called(ctx, businessType)
	return args.Error(0)
}

func (m *MockBusinessTypeRepository) Update(ctx context.Context, businessType interface{}) error {
	args := m.Called(ctx, businessType)
	return args.Error(0)
}

func (m *MockBusinessTypeRepository) FindByID(ctx context.Context, id string) (interface{}, error) {
	args := m.Called(ctx, id)
	return args.Get(0), args.Error(1)
}

func (m *MockBusinessTypeRepository) FindByCode(ctx context.Context, code string) (interface{}, error) {
	args := m.Called(ctx, code)
	return args.Get(0), args.Error(1)
}

func (m *MockBusinessTypeRepository) FindAll(ctx context.Context) (interface{}, error) {
	args := m.Called(ctx)
	return args.Get(0), args.Error(1)
}

func (m *MockBusinessTypeRepository) FindActive(ctx context.Context) (interface{}, error) {
	args := m.Called(ctx)
	return args.Get(0), args.Error(1)
}

func (m *MockBusinessTypeRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockBusinessTypeRepository) SearchByCriteria(ctx context.Context, crit interface{}) (interface{}, error) {
	args := m.Called(ctx, crit)
	return args.Get(0), args.Error(1)
}

func (m *MockBusinessTypeRepository) CountByCriteria(ctx context.Context, crit interface{}) (int, error) {
	args := m.Called(ctx, crit)
	return args.Int(0), args.Error(1)
}

// MockBusinessTypeTemplateRepository mock del repositorio de templates
type MockBusinessTypeTemplateRepository struct {
	mock.Mock
}

func (m *MockBusinessTypeTemplateRepository) Create(ctx context.Context, template interface{}) error {
	args := m.Called(ctx, template)
	return args.Error(0)
}

func (m *MockBusinessTypeTemplateRepository) Update(ctx context.Context, template interface{}) error {
	args := m.Called(ctx, template)
	return args.Error(0)
}

func (m *MockBusinessTypeTemplateRepository) FindByID(ctx context.Context, id string) (interface{}, error) {
	args := m.Called(ctx, id)
	return args.Get(0), args.Error(1)
}

func (m *MockBusinessTypeTemplateRepository) FindByBusinessTypeID(ctx context.Context, businessTypeID string) (interface{}, error) {
	args := m.Called(ctx, businessTypeID)
	return args.Get(0), args.Error(1)
}

func (m *MockBusinessTypeTemplateRepository) FindByBusinessTypeAndRegion(ctx context.Context, businessTypeID, region string) (interface{}, error) {
	args := m.Called(ctx, businessTypeID, region)
	return args.Get(0), args.Error(1)
}

func (m *MockBusinessTypeTemplateRepository) FindDefault(ctx context.Context, businessTypeID, region string) (interface{}, error) {
	args := m.Called(ctx, businessTypeID, region)
	return args.Get(0), args.Error(1)
}

func (m *MockBusinessTypeTemplateRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockBusinessTypeTemplateRepository) SearchByCriteria(ctx context.Context, crit interface{}) (interface{}, error) {
	args := m.Called(ctx, crit)
	return args.Get(0), args.Error(1)
}

func (m *MockBusinessTypeTemplateRepository) CountByCriteria(ctx context.Context, crit interface{}) (int, error) {
	args := m.Called(ctx, crit)
	return args.Int(0), args.Error(1)
}

// MockTenantSetupRepository mock del repositorio de tenant setup
type MockTenantSetupRepository struct {
	mock.Mock
}

func (m *MockTenantSetupRepository) Create(ctx context.Context, history interface{}) error {
	args := m.Called(ctx, history)
	return args.Error(0)
}

func (m *MockTenantSetupRepository) GetByTenantID(ctx context.Context, tenantID string) (interface{}, error) {
	args := m.Called(ctx, tenantID)
	return args.Get(0), args.Error(1)
}

func (m *MockTenantSetupRepository) GetByID(ctx context.Context, id string) (interface{}, error) {
	args := m.Called(ctx, id)
	return args.Get(0), args.Error(1)
}

func (m *MockTenantSetupRepository) Update(ctx context.Context, history interface{}) error {
	args := m.Called(ctx, history)
	return args.Error(0)
}

func (m *MockTenantSetupRepository) GetLatestByTenantID(ctx context.Context, tenantID string) (interface{}, error) {
	args := m.Called(ctx, tenantID)
	return args.Get(0), args.Error(1)
}

func TestStartWizardUseCase_Execute_Success(t *testing.T) {
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

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewStartWizardUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID, businessTypeID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tenantID, result.TenantID)
	assert.Equal(t, businessTypeID, result.BusinessTypeID)
	assert.False(t, result.SetupCompleted)

	mockBusinessTypeRepo.AssertExpectations(t)
	mockTemplateRepo.AssertExpectations(t)
	mockTenantSetupRepo.AssertExpectations(t)
}

func TestStartWizardUseCase_Execute_EmptyTenantID_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := ""
	businessTypeID := uuid.New().String()

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewStartWizardUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID, businessTypeID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "tenant ID es requerido")
}

func TestStartWizardUseCase_Execute_EmptyBusinessTypeID_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()
	businessTypeID := ""

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewStartWizardUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID, businessTypeID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "business type ID es requerido")
}

func TestStartWizardUseCase_Execute_BusinessTypeNotFound_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()
	businessTypeID := uuid.New().String()

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	mockBusinessTypeRepo.On("FindByID", ctx, businessTypeID).Return(nil, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewStartWizardUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID, businessTypeID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no encontrado")

	mockBusinessTypeRepo.AssertExpectations(t)
}

func TestStartWizardUseCase_Execute_NoTemplatesAvailable_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tenantID := uuid.New().String()
	businessTypeID := uuid.New().String()

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	businessType := struct {
		ID   string
		Code string
		Name string
	}{
		ID:   businessTypeID,
		Code: "polirubro",
		Name: "Polirubro",
	}

	emptyTemplates := []*interface{}{}

	mockBusinessTypeRepo.On("FindByID", ctx, businessTypeID).Return(&businessType, nil)
	mockTemplateRepo.On("FindByBusinessTypeID", ctx, businessTypeID).Return(emptyTemplates, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewStartWizardUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, tenantID, businessTypeID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no hay templates disponibles")

	mockBusinessTypeRepo.AssertExpectations(t)
	mockTemplateRepo.AssertExpectations(t)
}