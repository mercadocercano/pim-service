package usecase_test

import (
	"context"
	"time"

	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/application/response"
	"saas-mt-pim-service/src/template_ai/domain/entity"
	"saas-mt-pim-service/src/template_ai/domain/port"
)

// --- AITemplateRepository Mock ---

type mockAITemplateRepo struct {
	saveErr                error
	updateErr              error
	findByIDResult         *entity.AITemplate
	findByIDErr            error
	findByBusinessResult   []*entity.AITemplate
	findByBusinessErr      error
	saveHistoryErr         error
	findProductsResult     []*entity.TemplateGlobalProduct
	findProductsErr        error
	saveProductErr         error
	updateProductErr       error
	deleteProductErr       error
	findMetricsResult      []*entity.AIPerformanceMetric
	findMetricsErr         error
	aggregatedMetrics      map[string]float64
	aggregatedMetricsErr   error
	feedbackResult         []*entity.AIProductFeedback
	feedbackErr            error
	feedbackSummary        map[string]int
	feedbackSummaryErr     error

	saveCalls          int
	updateCalls        int
	saveHistoryCalls   int
	saveProductCalls   int
	saveFeedbackCalls  int
}

func (m *mockAITemplateRepo) Save(_ context.Context, _ *entity.AITemplate) error {
	m.saveCalls++
	return m.saveErr
}

func (m *mockAITemplateRepo) Update(_ context.Context, _ *entity.AITemplate) error {
	m.updateCalls++
	return m.updateErr
}

func (m *mockAITemplateRepo) FindByID(_ context.Context, _ uuid.UUID) (*entity.AITemplate, error) {
	return m.findByIDResult, m.findByIDErr
}

func (m *mockAITemplateRepo) FindByBusinessTypeAndTenant(_ context.Context, _ uuid.UUID, _ *uuid.UUID) ([]*entity.AITemplate, error) {
	return m.findByBusinessResult, m.findByBusinessErr
}

func (m *mockAITemplateRepo) Delete(_ context.Context, _ uuid.UUID) error { return nil }

func (m *mockAITemplateRepo) SaveTemplateProduct(_ context.Context, _ *entity.TemplateGlobalProduct) error {
	m.saveProductCalls++
	return m.saveProductErr
}

func (m *mockAITemplateRepo) UpdateTemplateProduct(_ context.Context, _ *entity.TemplateGlobalProduct) error {
	return m.updateProductErr
}

func (m *mockAITemplateRepo) FindTemplateProducts(_ context.Context, _ uuid.UUID) ([]*entity.TemplateGlobalProduct, error) {
	return m.findProductsResult, m.findProductsErr
}

func (m *mockAITemplateRepo) DeleteTemplateProduct(_ context.Context, _, _ uuid.UUID) error {
	return m.deleteProductErr
}

func (m *mockAITemplateRepo) SaveGenerationHistory(_ context.Context, _ *entity.AIGenerationHistory) error {
	m.saveHistoryCalls++
	return m.saveHistoryErr
}

func (m *mockAITemplateRepo) FindGenerationHistory(_ context.Context, _ *uuid.UUID, _ *uuid.UUID, _ int) ([]*entity.AIGenerationHistory, error) {
	return nil, nil
}

func (m *mockAITemplateRepo) SavePerformanceMetric(_ context.Context, _ *entity.AIPerformanceMetric) error {
	return nil
}

func (m *mockAITemplateRepo) FindPerformanceMetrics(_ context.Context, _ uuid.UUID, _ string) ([]*entity.AIPerformanceMetric, error) {
	return m.findMetricsResult, m.findMetricsErr
}

func (m *mockAITemplateRepo) GetAggregatedMetrics(_ context.Context, _ uuid.UUID) (map[string]float64, error) {
	return m.aggregatedMetrics, m.aggregatedMetricsErr
}

func (m *mockAITemplateRepo) SaveProductFeedback(_ context.Context, _ *entity.AIProductFeedback) error {
	m.saveFeedbackCalls++
	return nil
}

func (m *mockAITemplateRepo) FindProductFeedback(_ context.Context, _ *uuid.UUID, _ uuid.UUID) ([]*entity.AIProductFeedback, error) {
	return m.feedbackResult, m.feedbackErr
}

func (m *mockAITemplateRepo) GetFeedbackSummary(_ context.Context, _ uuid.UUID) (map[string]int, error) {
	return m.feedbackSummary, m.feedbackSummaryErr
}

// --- AIGenerationService Mock ---

type mockAIGenerationService struct {
	generateResult    *entity.AITemplate
	generateErr       error
	scoreResult       []*entity.ScoredProduct
	scoreErr          error
	optimizeResult    *entity.AITemplate
	optimizeErr       error
	analyzeResult     *port.PerformanceAnalysis
	analyzeErr        error
}

func (m *mockAIGenerationService) GenerateTemplate(_ context.Context, _ port.GenerationRequest) (*entity.AITemplate, error) {
	return m.generateResult, m.generateErr
}

func (m *mockAIGenerationService) ScoreProducts(_ context.Context, _ []*entity.GlobalProduct, _ interface{}) ([]*entity.ScoredProduct, error) {
	return m.scoreResult, m.scoreErr
}

func (m *mockAIGenerationService) OptimizeTemplate(_ context.Context, _ *entity.AITemplate, _ []port.TemplateFeedback) (*entity.AITemplate, error) {
	return m.optimizeResult, m.optimizeErr
}

func (m *mockAIGenerationService) AnalyzePerformance(_ context.Context, _ []port.PerformanceMetric) (*port.PerformanceAnalysis, error) {
	return m.analyzeResult, m.analyzeErr
}

// --- GlobalProductRepository Mock ---

type mockGlobalProductRepo struct {
	findByIDsResult map[uuid.UUID]interface{}
	findByIDsErr    error
}

func (m *mockGlobalProductRepo) FindByIDs(_ context.Context, _ []uuid.UUID) (map[uuid.UUID]interface{}, error) {
	return m.findByIDsResult, m.findByIDsErr
}

func (m *mockGlobalProductRepo) SearchByCategory(_ context.Context, _ string, _ int) ([]interface{}, error) {
	return nil, nil
}

func (m *mockGlobalProductRepo) SearchByBrand(_ context.Context, _ string, _ int) ([]interface{}, error) {
	return nil, nil
}

func (m *mockGlobalProductRepo) GetFilteredProducts(_ context.Context, _ map[string]interface{}) ([]*entity.GlobalProduct, error) {
	return nil, nil
}

// --- AITemplateDomainServicePort Mock ---

type mockDomainService struct {
	validateCreationErr    error
	validateProductErr     error
	performanceMetrics     map[string]float64
	performanceErr         error
	shouldRegenerate       bool
	validateFeedbackErr    error
}

func (m *mockDomainService) ValidateTemplateForCreation(_ context.Context, _ *entity.AITemplate) error {
	return m.validateCreationErr
}

func (m *mockDomainService) ValidateProductForTemplate(_ context.Context, _ *entity.TemplateGlobalProduct) error {
	return m.validateProductErr
}

func (m *mockDomainService) CalculateTemplatePerformance(_ context.Context, _ uuid.UUID) (map[string]float64, error) {
	return m.performanceMetrics, m.performanceErr
}

func (m *mockDomainService) ShouldRegenerateTemplate(_ map[string]float64) bool {
	return m.shouldRegenerate
}

func (m *mockDomainService) ValidateFeedback(_ *entity.AIProductFeedback) error {
	return m.validateFeedbackErr
}

// --- TemplateMapperPort Mock ---

type mockMapper struct{}

func (m *mockMapper) ToTemplateResponse(template *entity.AITemplate, productCount int) *response.AITemplateResponse {
	return &response.AITemplateResponse{
		ID:          template.ID,
		Name:        template.Name,
		Description: template.Description,
		GeneratedBy: template.GeneratedBy,
		ProductCount: productCount,
	}
}

func (m *mockMapper) ToTemplateProductResponse(tp *entity.TemplateGlobalProduct) *response.TemplateProductResponse {
	return &response.TemplateProductResponse{
		ID:              tp.ID,
		GlobalProductID: tp.GlobalProductID,
		Priority:        tp.Priority,
	}
}

func (m *mockMapper) ToGenerationSummaryResponse(_ []*entity.TemplateGlobalProduct, timeMs int, model string) *response.GenerationSummaryResponse {
	return &response.GenerationSummaryResponse{
		GenerationTimeMs: timeMs,
		AIModel:          model,
		CategoryBreakdown: make(map[string]int),
		BrandBreakdown:    make(map[string]int),
	}
}

func (m *mockMapper) ToMetricValueResponse(value float64, target *float64) *response.MetricValueResponse {
	status := "good"
	if value < 0.6 {
		status = "critical"
	} else if value < 0.8 {
		status = "warning"
	}
	return &response.MetricValueResponse{Value: value, Target: target, Status: status}
}

func (m *mockMapper) ToRecommendationResponse(recType, priority, title, desc, action string, metrics []string) *response.RecommendationResponse {
	return &response.RecommendationResponse{
		Type: recType, Priority: priority, Title: title, Description: desc, Action: action, Metrics: metrics,
	}
}

func (m *mockMapper) GetPerformanceRating(score float64) string {
	if score >= 0.9 {
		return "excellent"
	} else if score >= 0.75 {
		return "good"
	}
	return "fair"
}

// --- Helpers ---

func newUUID() uuid.UUID {
	id, _ := uuid.NewV4()
	return id
}

func newTemplate() *entity.AITemplate {
	return &entity.AITemplate{
		ID:                 newUUID(),
		BusinessTypeID:     newUUID(),
		Name:               "Template Test",
		Description:        "Template de prueba para tests",
		GeneratedBy:        "ai",
		AIGenerationParams: make(map[string]interface{}),
		PerformanceMetrics: make(map[string]interface{}),
		Version:            1,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
}
