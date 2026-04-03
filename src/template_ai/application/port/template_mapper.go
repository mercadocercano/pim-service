package port

import (
	"saas-mt-pim-service/src/template_ai/application/response"
	"saas-mt-pim-service/src/template_ai/domain/entity"
)

// TemplateMapperPort define la interfaz del mapper de templates.
type TemplateMapperPort interface {
	ToTemplateResponse(template *entity.AITemplate, productCount int) *response.AITemplateResponse
	ToTemplateProductResponse(templateProduct *entity.TemplateGlobalProduct) *response.TemplateProductResponse
	ToGenerationSummaryResponse(products []*entity.TemplateGlobalProduct, generationTimeMs int, aiModel string) *response.GenerationSummaryResponse
	ToMetricValueResponse(value float64, target *float64) *response.MetricValueResponse
	ToRecommendationResponse(recType, priority, title, description, action string, affectedMetrics []string) *response.RecommendationResponse
	GetPerformanceRating(score float64) string
}
