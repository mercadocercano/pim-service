package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"saas-mt-pim-service/src/template_ai/application/request"
	"saas-mt-pim-service/src/template_ai/application/usecase"
	"saas-mt-pim-service/src/template_ai/domain/entity"

	"github.com/gofrs/uuid/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validUpdateRequest(templateID string) *request.UpdateTemplateFromFeedbackRequest {
	return &request.UpdateTemplateFromFeedbackRequest{
		TemplateID:         templateID,
		FeedbackPeriodDays: 30,
		MinFeedbackCount:   5,
		UpdateStrategy:     "incremental",
		ForceUpdate:        true,
	}
}

func TestUpdateFromFeedback_InsufficientFeedback_ReturnsNotSuccess(t *testing.T) {
	template := newTemplate()
	repo := &mockAITemplateRepo{
		findByIDResult: template,
		feedbackResult: []*entity.AIProductFeedback{}, // empty
	}
	domainSvc := &mockDomainService{performanceMetrics: map[string]float64{}}

	uc := usecase.NewUpdateTemplateFromFeedbackUseCase(repo, &mockAIGenerationService{}, domainSvc, &mockMapper{})

	req := &request.UpdateTemplateFromFeedbackRequest{
		TemplateID:       template.ID.String(),
		UpdateStrategy:   "incremental",
		MinFeedbackCount: 10,
		ForceUpdate:      false, // don't force
	}

	result, err := uc.Execute(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.Success)
	assert.Contains(t, result.Message, "Insufficient feedback")
}

func TestUpdateFromFeedback_ForceUpdate_ProcessesDespiteLowFeedback(t *testing.T) {
	template := newTemplate()
	template.PerformanceMetrics = map[string]interface{}{}

	productID := newUUID()
	feedback := []*entity.AIProductFeedback{
		{
			ID:              newUUID(),
			TenantID:        newUUID(),
			TemplateID:      &template.ID,
			GlobalProductID: &productID,
			Action:          "removed",
			CreatedAt:       time.Now(),
		},
	}

	repo := &mockAITemplateRepo{
		findByIDResult:     template,
		feedbackResult:     feedback,
		findProductsResult: []*entity.TemplateGlobalProduct{},
	}
	domainSvc := &mockDomainService{
		performanceMetrics: map[string]float64{"satisfaction_score": 0.5},
		shouldRegenerate:   true,
	}

	uc := usecase.NewUpdateTemplateFromFeedbackUseCase(repo, &mockAIGenerationService{}, domainSvc, &mockMapper{})

	req := validUpdateRequest(template.ID.String())
	req.MinFeedbackCount = 100
	req.ForceUpdate = true

	result, err := uc.Execute(context.Background(), req)

	require.NoError(t, err)
	assert.True(t, result.Success)
	assert.NotNil(t, result.UpdateSummary)
}

func TestUpdateFromFeedback_TemplateNotFound_ReturnsError(t *testing.T) {
	repo := &mockAITemplateRepo{findByIDErr: errors.New("not found")}

	uc := usecase.NewUpdateTemplateFromFeedbackUseCase(repo, &mockAIGenerationService{}, &mockDomainService{}, &mockMapper{})

	req := validUpdateRequest(newUUID().String())

	_, err := uc.Execute(context.Background(), req)

	require.Error(t, err)
}

func TestUpdateFromFeedback_InvalidTemplateID_ReturnsError(t *testing.T) {
	uc := usecase.NewUpdateTemplateFromFeedbackUseCase(
		&mockAITemplateRepo{}, &mockAIGenerationService{}, &mockDomainService{}, &mockMapper{},
	)

	req := &request.UpdateTemplateFromFeedbackRequest{
		TemplateID:     "invalid-uuid",
		UpdateStrategy: "incremental",
	}

	_, err := uc.Execute(context.Background(), req)

	require.Error(t, err)
}

func TestUpdateFromFeedback_TestMode_SimulatesWithoutApplying(t *testing.T) {
	template := newTemplate()
	template.PerformanceMetrics = map[string]interface{}{}

	productID := newUUID()
	feedback := make([]*entity.AIProductFeedback, 0, 15)
	for i := 0; i < 15; i++ {
		feedback = append(feedback, &entity.AIProductFeedback{
			ID:              newUUID(),
			TenantID:        newUUID(),
			TemplateID:      &template.ID,
			GlobalProductID: &productID,
			Action:          "removed",
			CreatedAt:       time.Now(),
		})
	}

	repo := &mockAITemplateRepo{
		findByIDResult: template,
		feedbackResult: feedback,
	}
	domainSvc := &mockDomainService{
		performanceMetrics: map[string]float64{"satisfaction_score": 0.4},
		shouldRegenerate:   true,
	}

	uc := usecase.NewUpdateTemplateFromFeedbackUseCase(repo, &mockAIGenerationService{}, domainSvc, &mockMapper{})

	req := &request.UpdateTemplateFromFeedbackRequest{
		TemplateID:     template.ID.String(),
		UpdateStrategy: "incremental",
		ForceUpdate:    true,
		TestMode:       true,
	}

	result, err := uc.Execute(context.Background(), req)

	require.NoError(t, err)
	assert.True(t, result.Success)
	assert.NotNil(t, result.TestResults)
	assert.Equal(t, 0, repo.updateCalls) // no actual update in test mode
}

func TestUpdateFromFeedback_PerformanceSatisfactory_SkipsUpdate(t *testing.T) {
	template := newTemplate()
	repo := &mockAITemplateRepo{
		findByIDResult: template,
		feedbackResult: makeFeedback(template.ID, 20),
	}
	domainSvc := &mockDomainService{
		performanceMetrics: map[string]float64{"satisfaction_score": 0.95},
		shouldRegenerate:   false, // performance is good
	}

	uc := usecase.NewUpdateTemplateFromFeedbackUseCase(repo, &mockAIGenerationService{}, domainSvc, &mockMapper{})

	req := &request.UpdateTemplateFromFeedbackRequest{
		TemplateID:       template.ID.String(),
		UpdateStrategy:   "incremental",
		MinFeedbackCount: 5,
		ForceUpdate:      false,
	}

	result, err := uc.Execute(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.Success)
	assert.Contains(t, result.Message, "satisfactory")
}

func makeFeedback(templateID uuid.UUID, count int) []*entity.AIProductFeedback {
	feedback := make([]*entity.AIProductFeedback, 0, count)
	for i := 0; i < count; i++ {
		pid := newUUID()
		feedback = append(feedback, &entity.AIProductFeedback{
			ID:              newUUID(),
			TenantID:        newUUID(),
			TemplateID:      &templateID,
			GlobalProductID: &pid,
			Action:          "kept",
			CreatedAt:       time.Now(),
		})
	}
	return feedback
}
