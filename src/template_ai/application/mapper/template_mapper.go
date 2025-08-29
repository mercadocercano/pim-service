package mapper

import (
	"saas-mt-pim-service/src/template_ai/application/response"
	"saas-mt-pim-service/src/template_ai/domain/entity"
	"saas-mt-pim-service/src/template_ai/domain/value_object"
)

// TemplateMapper handles mapping between domain entities and DTOs
type TemplateMapper struct{}

// NewTemplateMapper creates a new instance of TemplateMapper
func NewTemplateMapper() *TemplateMapper {
	return &TemplateMapper{}
}

// ToTemplateResponse maps an AI template entity to response DTO
func (m *TemplateMapper) ToTemplateResponse(template *entity.AITemplate, productCount int) *response.AITemplateResponse {
	return &response.AITemplateResponse{
		ID:                    template.ID,
		BusinessTypeID:        template.BusinessTypeID,
		BusinessTypeName:      "", // This would be populated from business type service
		TenantID:              template.TenantID,
		Name:                  template.Name,
		Description:           template.Description,
		GeneratedBy:           template.GeneratedBy,
		ProductCount:          productCount,
		AIGenerationParams:    template.AIGenerationParams,
		PerformanceMetrics:    template.PerformanceMetrics,
		LastAIUpdate:          template.LastAIUpdate,
		ProductSelectionRules: template.ProductSelectionRules,
		CategoryDistribution:  template.CategoryDistribution,
		BrandPreferences:      template.BrandPreferences,
		RegionalAdaptations:   template.RegionalAdaptations,
		CreatedAt:             template.CreatedAt,
		UpdatedAt:             template.UpdatedAt,
	}
}

// ToTemplateProductResponse maps a template-product entity to response DTO
func (m *TemplateMapper) ToTemplateProductResponse(templateProduct *entity.TemplateGlobalProduct) *response.TemplateProductResponse {
	priorityLabel := value_object.ProductPriority(templateProduct.Priority).String()
	
	return &response.TemplateProductResponse{
		ID:                      templateProduct.ID,
		GlobalProductID:         templateProduct.GlobalProductID,
		Priority:                templateProduct.Priority,
		PriorityLabel:           priorityLabel,
		QuantitySuggestion:      templateProduct.QuantitySuggestion,
		AIReasoning:             templateProduct.AIReasoning,
		RelevanceScore:          templateProduct.RelevanceScore,
		CategoryMatchScore:      templateProduct.CategoryMatchScore,
		BrandMatchScore:         templateProduct.BrandMatchScore,
		RegionalPreferenceScore: templateProduct.RegionalPreferenceScore,
		CreatedAt:               templateProduct.CreatedAt,
		UpdatedAt:               templateProduct.UpdatedAt,
	}
}

// ToGenerationSummaryResponse creates a generation summary response
func (m *TemplateMapper) ToGenerationSummaryResponse(
	products []*entity.TemplateGlobalProduct,
	generationTimeMs int,
	aiModel string,
) *response.GenerationSummaryResponse {
	summary := &response.GenerationSummaryResponse{
		TotalProducts:       len(products),
		GenerationTimeMs:    generationTimeMs,
		AIModel:             aiModel,
		CategoryBreakdown:   make(map[string]int),
		BrandBreakdown:      make(map[string]int),
	}

	// Count products by priority
	for _, product := range products {
		switch product.Priority {
		case int(value_object.ProductPriorityEssential):
			summary.EssentialProducts++
		case int(value_object.ProductPriorityRecommended):
			summary.RecommendedProducts++
		case int(value_object.ProductPriorityOptional):
			summary.OptionalProducts++
		}
	}

	// Calculate optimization score based on distribution
	if len(products) > 0 {
		essentialRatio := float64(summary.EssentialProducts) / float64(len(products))
		recommendedRatio := float64(summary.RecommendedProducts) / float64(len(products))
		
		// Ideal distribution: 60% essential, 30% recommended, 10% optional
		essentialDiff := abs(essentialRatio - 0.6)
		recommendedDiff := abs(recommendedRatio - 0.3)
		
		summary.OptimizationScore = 1.0 - (essentialDiff + recommendedDiff)
		if summary.OptimizationScore < 0 {
			summary.OptimizationScore = 0
		}
	}

	return summary
}

// ToMetricValueResponse maps metric data to response
func (m *TemplateMapper) ToMetricValueResponse(value float64, target *float64) *response.MetricValueResponse {
	resp := &response.MetricValueResponse{
		Value:  value,
		Target: target,
	}

	// Determine status based on value and target
	if target != nil {
		if value >= *target {
			resp.Status = "good"
		} else if value >= *target*0.8 {
			resp.Status = "warning"
		} else {
			resp.Status = "critical"
		}
	} else {
		// Default status logic without target
		if value >= 0.8 {
			resp.Status = "good"
		} else if value >= 0.6 {
			resp.Status = "warning"
		} else {
			resp.Status = "critical"
		}
	}

	return resp
}

// ToRecommendationResponse creates a recommendation response
func (m *TemplateMapper) ToRecommendationResponse(
	recType, priority, title, description, action string,
	affectedMetrics []string,
) *response.RecommendationResponse {
	return &response.RecommendationResponse{
		Type:        recType,
		Priority:    priority,
		Title:       title,
		Description: description,
		Action:      action,
		Metrics:     affectedMetrics,
	}
}

// GetPerformanceRating calculates performance rating based on score
func (m *TemplateMapper) GetPerformanceRating(score float64) string {
	if score >= 0.9 {
		return "excellent"
	} else if score >= 0.75 {
		return "good"
	} else if score >= 0.6 {
		return "fair"
	} else if score >= 0.4 {
		return "poor"
	}
	return "critical"
}

// Helper function for absolute value
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}