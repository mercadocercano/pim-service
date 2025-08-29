package port

import (
	"context"
	"time"
	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/domain/entity"
)

// GenerationRequest represents a request to generate a template
type GenerationRequest struct {
	BusinessTypeID   uuid.UUID
	TemplateName     string
	TargetSize       string
	Preferences      map[string]interface{}
	Constraints      map[string]interface{}
	OptimizationGoal string
	Region           string
	Season           string
	BudgetRange      string
}

// TemplateFeedback represents feedback on a template
type TemplateFeedback struct {
	ProductID   uuid.UUID
	Action      string
	Reason      string
	NewQuantity int
	Timestamp   time.Time
}

// PerformanceMetric represents a performance metric
type PerformanceMetric struct {
	TemplateID   uuid.UUID
	MetricType   string
	MetricValue  float64
	Metadata     map[string]interface{}
	PeriodStart  time.Time
	PeriodEnd    time.Time
}

// PerformanceAnalysis represents the analysis of performance metrics
type PerformanceAnalysis struct {
	Score           float64
	Insights        []string
	Recommendations []string
	Trends          map[string]interface{}
	Metadata        map[string]interface{}
}

// CommonModification represents a common modification pattern
type CommonModification struct {
	ProductID       uuid.UUID
	ProductName     string
	Action          string
	Frequency       int
	AvgQuantity     int
	AffectedTenants int
	Reason          string
}

// AIGenerationService defines the interface for AI template generation
type AIGenerationService interface {
	GenerateTemplate(ctx context.Context, request GenerationRequest) (*entity.AITemplate, error)
	ScoreProducts(ctx context.Context, products []*entity.GlobalProduct, params interface{}) ([]*entity.ScoredProduct, error)
	OptimizeTemplate(ctx context.Context, template *entity.AITemplate, feedback []TemplateFeedback) (*entity.AITemplate, error)
	AnalyzePerformance(ctx context.Context, metrics []PerformanceMetric) (*PerformanceAnalysis, error)
}