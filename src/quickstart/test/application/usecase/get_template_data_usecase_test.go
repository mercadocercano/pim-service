package usecase

import (
	"context"
	"testing"

	"saas-mt-pim-service/src/quickstart/application/usecase"
	"saas-mt-pim-service/src/quickstart/domain/service"
	businessTypeTestEntity "saas-mt-pim-service/src/quickstart/test/domain/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetTemplateDataUseCase_Execute_AllData_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	businessTypeID := uuid.New().String()
	section := "" // Empty means all data

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	templateOM := businessTypeTestEntity.NewBusinessTypeTemplateObjectMother()
	template := templateOM.WithBusinessTypeID(businessTypeID)
	templates := []*interface{}{&template}

	mockTemplateRepo.On("FindByBusinessTypeID", ctx, businessTypeID).Return(templates, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewGetTemplateDataUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, businessTypeID, section)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verificar que contiene las secciones esperadas
	resultMap, ok := result.(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, resultMap, "categories")
	assert.Contains(t, resultMap, "attributes")
	assert.Contains(t, resultMap, "products")
	assert.Contains(t, resultMap, "brands")
	assert.Contains(t, resultMap, "metadata")

	mockTemplateRepo.AssertExpectations(t)
}

func TestGetTemplateDataUseCase_Execute_CategoriesSection_Success(t *testing.T) {
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

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewGetTemplateDataUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, businessTypeID, section)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	mockTemplateRepo.AssertExpectations(t)
}

func TestGetTemplateDataUseCase_Execute_ProductsSection_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	businessTypeID := uuid.New().String()
	section := "products"

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	templateOM := businessTypeTestEntity.NewBusinessTypeTemplateObjectMother()
	template := templateOM.WithBusinessTypeID(businessTypeID)
	templates := []*interface{}{&template}

	mockTemplateRepo.On("FindByBusinessTypeID", ctx, businessTypeID).Return(templates, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewGetTemplateDataUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, businessTypeID, section)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	mockTemplateRepo.AssertExpectations(t)
}

func TestGetTemplateDataUseCase_Execute_EmptyBusinessTypeID_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	businessTypeID := ""
	section := "categories"

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewGetTemplateDataUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, businessTypeID, section)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "business type ID es requerido")
}

func TestGetTemplateDataUseCase_Execute_NoTemplatesFound_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	businessTypeID := uuid.New().String()
	section := "categories"

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	emptyTemplates := []*interface{}{}
	mockTemplateRepo.On("FindByBusinessTypeID", ctx, businessTypeID).Return(emptyTemplates, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewGetTemplateDataUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, businessTypeID, section)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no hay templates")

	mockTemplateRepo.AssertExpectations(t)
}

func TestGetTemplateDataUseCase_Execute_InvalidSection_ReturnsError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	businessTypeID := uuid.New().String()
	section := "invalid_section"

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	templateOM := businessTypeTestEntity.NewBusinessTypeTemplateObjectMother()
	template := templateOM.WithBusinessTypeID(businessTypeID)
	templates := []*interface{}{&template}

	mockTemplateRepo.On("FindByBusinessTypeID", ctx, businessTypeID).Return(templates, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewGetTemplateDataUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, businessTypeID, section)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "sección invalid_section no encontrada")

	mockTemplateRepo.AssertExpectations(t)
}

func TestGetTemplateDataUseCase_Execute_BrandsSection_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	businessTypeID := uuid.New().String()
	section := "brands"

	mockBusinessTypeRepo := new(MockBusinessTypeRepository)
	mockTemplateRepo := new(MockBusinessTypeTemplateRepository)
	mockTenantSetupRepo := new(MockTenantSetupRepository)

	templateOM := businessTypeTestEntity.NewBusinessTypeTemplateObjectMother()
	template := templateOM.WithBusinessTypeID(businessTypeID)
	templates := []*interface{}{&template}

	mockTemplateRepo.On("FindByBusinessTypeID", ctx, businessTypeID).Return(templates, nil)

	wizardService := service.NewQuickstartWizardService(
		mockBusinessTypeRepo,
		mockTemplateRepo,
		mockTenantSetupRepo,
	)
	useCase := usecase.NewGetTemplateDataUseCase(wizardService)

	// Act
	result, err := useCase.Execute(ctx, businessTypeID, section)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verificar que retorna un slice de brands
	brands, ok := result.([]string)
	assert.True(t, ok)
	assert.Greater(t, len(brands), 0)

	mockTemplateRepo.AssertExpectations(t)
}