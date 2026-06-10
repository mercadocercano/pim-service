package usecase

import (
	"context"
	"testing"

	cr "github.com/mercadocercano/criteria"
	businessTypeEntity "saas-mt-pim-service/src/businesstype/domain/entity"
	"saas-mt-pim-service/src/quickstart/application/usecase"
	quickstartEntity "saas-mt-pim-service/src/quickstart/domain/entity"
	"saas-mt-pim-service/src/quickstart/domain/service"
	"saas-mt-pim-service/src/quickstart/test/domain/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBusinessTypeRepository mock del repositorio de business types
type MockBusinessTypeRepository struct {
	mock.Mock
}

func (m *MockBusinessTypeRepository) Create(ctx context.Context, bt *businessTypeEntity.BusinessType) error {
	args := m.Called(ctx, bt)
	return args.Error(0)
}

func (m *MockBusinessTypeRepository) Update(ctx context.Context, bt *businessTypeEntity.BusinessType) error {
	args := m.Called(ctx, bt)
	return args.Error(0)
}

func (m *MockBusinessTypeRepository) FindByID(ctx context.Context, id string) (*businessTypeEntity.BusinessType, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*businessTypeEntity.BusinessType), args.Error(1)
}

func (m *MockBusinessTypeRepository) FindByCode(ctx context.Context, code string) (*businessTypeEntity.BusinessType, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*businessTypeEntity.BusinessType), args.Error(1)
}

func (m *MockBusinessTypeRepository) FindAll(ctx context.Context) ([]*businessTypeEntity.BusinessType, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*businessTypeEntity.BusinessType), args.Error(1)
}

func (m *MockBusinessTypeRepository) FindActive(ctx context.Context) ([]*businessTypeEntity.BusinessType, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*businessTypeEntity.BusinessType), args.Error(1)
}

func (m *MockBusinessTypeRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockBusinessTypeRepository) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*businessTypeEntity.BusinessType, error) {
	args := m.Called(ctx, crit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*businessTypeEntity.BusinessType), args.Error(1)
}

func (m *MockBusinessTypeRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	args := m.Called(ctx, crit)
	return args.Int(0), args.Error(1)
}

// MockBusinessTypeTemplateRepository mock del repositorio de templates
type MockBusinessTypeTemplateRepository struct {
	mock.Mock
}

func (m *MockBusinessTypeTemplateRepository) Create(ctx context.Context, template *businessTypeEntity.BusinessTypeTemplate) error {
	args := m.Called(ctx, template)
	return args.Error(0)
}

func (m *MockBusinessTypeTemplateRepository) Update(ctx context.Context, template *businessTypeEntity.BusinessTypeTemplate) error {
	args := m.Called(ctx, template)
	return args.Error(0)
}

func (m *MockBusinessTypeTemplateRepository) FindByID(ctx context.Context, id string) (*businessTypeEntity.BusinessTypeTemplate, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*businessTypeEntity.BusinessTypeTemplate), args.Error(1)
}

func (m *MockBusinessTypeTemplateRepository) FindByBusinessTypeID(ctx context.Context, businessTypeID string) ([]*businessTypeEntity.BusinessTypeTemplate, error) {
	args := m.Called(ctx, businessTypeID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*businessTypeEntity.BusinessTypeTemplate), args.Error(1)
}

func (m *MockBusinessTypeTemplateRepository) FindByBusinessTypeAndRegion(ctx context.Context, businessTypeID, region string) ([]*businessTypeEntity.BusinessTypeTemplate, error) {
	args := m.Called(ctx, businessTypeID, region)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*businessTypeEntity.BusinessTypeTemplate), args.Error(1)
}

func (m *MockBusinessTypeTemplateRepository) FindDefault(ctx context.Context, businessTypeID, region string) (*businessTypeEntity.BusinessTypeTemplate, error) {
	args := m.Called(ctx, businessTypeID, region)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*businessTypeEntity.BusinessTypeTemplate), args.Error(1)
}

func (m *MockBusinessTypeTemplateRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockBusinessTypeTemplateRepository) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*businessTypeEntity.BusinessTypeTemplate, error) {
	args := m.Called(ctx, crit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*businessTypeEntity.BusinessTypeTemplate), args.Error(1)
}

func (m *MockBusinessTypeTemplateRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	args := m.Called(ctx, crit)
	return args.Int(0), args.Error(1)
}

// MockTenantSetupRepository mock del repositorio de tenant setup
type MockTenantSetupRepository struct {
	mock.Mock
}

func (m *MockTenantSetupRepository) Create(ctx context.Context, history *quickstartEntity.TenantQuickstartHistory) error {
	args := m.Called(ctx, history)
	return args.Error(0)
}

func (m *MockTenantSetupRepository) GetByTenantID(ctx context.Context, tenantID string) ([]*quickstartEntity.TenantQuickstartHistory, error) {
	args := m.Called(ctx, tenantID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*quickstartEntity.TenantQuickstartHistory), args.Error(1)
}

func (m *MockTenantSetupRepository) GetByID(ctx context.Context, id string) (*quickstartEntity.TenantQuickstartHistory, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*quickstartEntity.TenantQuickstartHistory), args.Error(1)
}

func (m *MockTenantSetupRepository) Update(ctx context.Context, history *quickstartEntity.TenantQuickstartHistory) error {
	args := m.Called(ctx, history)
	return args.Error(0)
}

func (m *MockTenantSetupRepository) GetLatestByTenantID(ctx context.Context, tenantID string) (*quickstartEntity.TenantQuickstartHistory, error) {
	args := m.Called(ctx, tenantID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*quickstartEntity.TenantQuickstartHistory), args.Error(1)
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
	businessType := &businessTypeEntity.BusinessType{
		ID:   businessTypeID,
		Code: "polirubro",
		Name: "Polirubro",
	}

	templateOM := entity.NewBusinessTypeTemplateObjectMother()
	template := templateOM.WithBusinessTypeID(businessTypeID)
	templates := []*businessTypeEntity.BusinessTypeTemplate{template}

	mockBusinessTypeRepo.On("FindByID", ctx, businessTypeID).Return(businessType, nil)
	mockTemplateRepo.On("FindByBusinessTypeID", ctx, businessTypeID).Return(templates, nil)
	mockTenantSetupRepo.On("Create", ctx, mock.AnythingOfType("*entity.TenantQuickstartHistory")).Return(nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	uc := usecase.NewStartWizardUseCase(wizardService)

	// Act
	result, err := uc.Execute(ctx, tenantID, businessTypeID)

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
	uc := usecase.NewStartWizardUseCase(wizardService)

	// Act
	result, err := uc.Execute(ctx, tenantID, businessTypeID)

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
	uc := usecase.NewStartWizardUseCase(wizardService)

	// Act
	result, err := uc.Execute(ctx, tenantID, businessTypeID)

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
	uc := usecase.NewStartWizardUseCase(wizardService)

	// Act
	result, err := uc.Execute(ctx, tenantID, businessTypeID)

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

	businessType := &businessTypeEntity.BusinessType{
		ID:   businessTypeID,
		Code: "polirubro",
		Name: "Polirubro",
	}

	emptyTemplates := []*businessTypeEntity.BusinessTypeTemplate{}

	mockBusinessTypeRepo.On("FindByID", ctx, businessTypeID).Return(businessType, nil)
	mockTemplateRepo.On("FindByBusinessTypeID", ctx, businessTypeID).Return(emptyTemplates, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	uc := usecase.NewStartWizardUseCase(wizardService)

	// Act
	result, err := uc.Execute(ctx, tenantID, businessTypeID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no hay templates disponibles")

	mockBusinessTypeRepo.AssertExpectations(t)
	mockTemplateRepo.AssertExpectations(t)
}
