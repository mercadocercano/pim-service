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

func TestAnalyzePerformance_HappyPath_ReturnsMetrics(t *testing.T) {
	template := newTemplate()
	repo := &mockAITemplateRepo{
		findByIDResult:    template,
		aggregatedMetrics: map[string]float64{"satisfaction_score": 0.85, "modification_rate": 0.15},
		feedbackSummary:   map[string]int{"kept": 80, "removed": 20},
	}
	domainSvc := &mockDomainService{
		performanceMetrics: map[string]float64{"satisfaction_score": 0.85, "modification_rate": 0.15},
	}

	uc := usecase.NewAnalyzeTemplatePerformanceUseCase(repo, domainSvc, &mockMapper{})

	req := &request.AnalyzeTemplatePerformanceRequest{
		TemplateID:  template.ID.String(),
		MetricTypes: []string{"satisfaction_score", "modification_rate"},
	}

	result, err := uc.Execute(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, template.ID, result.TemplateID)
	assert.Equal(t, template.Name, result.TemplateName)
	assert.NotEmpty(t, result.Metrics)
	assert.NotEmpty(t, result.PerformanceRating)
}

func TestAnalyzePerformance_TemplateNotFound_ReturnsError(t *testing.T) {
	repo := &mockAITemplateRepo{findByIDErr: errors.New("not found")}

	uc := usecase.NewAnalyzeTemplatePerformanceUseCase(repo, &mockDomainService{}, &mockMapper{})

	req := &request.AnalyzeTemplatePerformanceRequest{
		TemplateID: newUUID().String(),
	}

	_, err := uc.Execute(context.Background(), req)

	require.Error(t, err)
}

func TestAnalyzePerformance_InvalidTemplateID_ReturnsError(t *testing.T) {
	uc := usecase.NewAnalyzeTemplatePerformanceUseCase(
		&mockAITemplateRepo{}, &mockDomainService{}, &mockMapper{},
	)

	req := &request.AnalyzeTemplatePerformanceRequest{
		TemplateID: "invalid-uuid",
	}

	_, err := uc.Execute(context.Background(), req)

	require.Error(t, err)
}

func TestAnalyzePerformance_WithDefaultMetrics_UsesDefaults(t *testing.T) {
	template := newTemplate()
	repo := &mockAITemplateRepo{
		findByIDResult:    template,
		aggregatedMetrics: map[string]float64{},
		feedbackSummary:   map[string]int{},
	}
	domainSvc := &mockDomainService{
		performanceMetrics: map[string]float64{},
	}

	uc := usecase.NewAnalyzeTemplatePerformanceUseCase(repo, domainSvc, &mockMapper{})

	req := &request.AnalyzeTemplatePerformanceRequest{
		TemplateID: template.ID.String(),
		// No metric types specified — should use defaults
	}

	result, err := uc.Execute(context.Background(), req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.AnalysisPeriod)
}
