package usecase_test

import (
	"context"
	"errors"
	"testing"

	"saas-mt-pim-service/src/template_ai/application/request"
	"saas-mt-pim-service/src/template_ai/application/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validGenerateRequest() *request.GenerateSmartTemplateRequest {
	businessTypeID := newUUID().String()
	return &request.GenerateSmartTemplateRequest{
		BusinessTypeID: businessTypeID,
		Name:           "Plantilla Kiosco Premium",
		Description:    "Plantilla optimizada para kioscos con productos premium y alta rotación",
		ProductCount:   50,
		BudgetRange: &request.BudgetRangeRequest{
			Min:      5000,
			Max:      50000,
			Currency: "ARS",
		},
	}
}

func TestGenerateSmartTemplate_HappyPath_ReturnsSuccess(t *testing.T) {
	template := newTemplate()
	repo := &mockAITemplateRepo{}
	aiService := &mockAIGenerationService{generateResult: template}
	domainSvc := &mockDomainService{}

	uc := usecase.NewGenerateSmartTemplateUseCase(repo, aiService, domainSvc, &mockMapper{})

	result, err := uc.Execute(context.Background(), validGenerateRequest())

	require.NoError(t, err)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Template)
	assert.Equal(t, "Plantilla Kiosco Premium", result.Template.Name)
	assert.Equal(t, 2, repo.saveHistoryCalls) // initial + completion update
	assert.Equal(t, 1, repo.saveCalls)
}

func TestGenerateSmartTemplate_InvalidRequest_ReturnsValidationError(t *testing.T) {
	uc := usecase.NewGenerateSmartTemplateUseCase(
		&mockAITemplateRepo{}, &mockAIGenerationService{}, &mockDomainService{}, &mockMapper{},
	)

	req := &request.GenerateSmartTemplateRequest{
		BusinessTypeID: "invalid-uuid",
		Name:           "Test",
		Description:    "Short",
		ProductCount:   5, // below minimum
	}

	_, err := uc.Execute(context.Background(), req)

	require.Error(t, err)
}

func TestGenerateSmartTemplate_AIGenerationFails_ReturnsError(t *testing.T) {
	repo := &mockAITemplateRepo{}
	aiService := &mockAIGenerationService{generateErr: errors.New("AI service unavailable")}

	uc := usecase.NewGenerateSmartTemplateUseCase(repo, aiService, &mockDomainService{}, &mockMapper{})

	_, err := uc.Execute(context.Background(), validGenerateRequest())

	require.Error(t, err)
	assert.Equal(t, 2, repo.saveHistoryCalls) // initial + error update
}

func TestGenerateSmartTemplate_DomainValidationFails_ReturnsError(t *testing.T) {
	template := newTemplate()
	repo := &mockAITemplateRepo{}
	aiService := &mockAIGenerationService{generateResult: template}
	domainSvc := &mockDomainService{validateCreationErr: errors.New("template already exists")}

	uc := usecase.NewGenerateSmartTemplateUseCase(repo, aiService, domainSvc, &mockMapper{})

	_, err := uc.Execute(context.Background(), validGenerateRequest())

	require.Error(t, err)
	assert.Contains(t, err.Error(), "template already exists")
	assert.Equal(t, 0, repo.saveCalls) // template NOT saved
}

func TestGenerateSmartTemplate_SaveFails_ReturnsError(t *testing.T) {
	template := newTemplate()
	repo := &mockAITemplateRepo{saveErr: errors.New("db connection lost")}
	aiService := &mockAIGenerationService{generateResult: template}

	uc := usecase.NewGenerateSmartTemplateUseCase(repo, aiService, &mockDomainService{}, &mockMapper{})

	_, err := uc.Execute(context.Background(), validGenerateRequest())

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to save template")
}
