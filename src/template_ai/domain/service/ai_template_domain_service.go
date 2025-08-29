package service

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/domain/entity"
	"saas-mt-pim-service/src/template_ai/domain/exception"
	"saas-mt-pim-service/src/template_ai/domain/port"
	"saas-mt-pim-service/src/template_ai/domain/value_object"
)

// AITemplateDomainService handles business logic for AI templates
type AITemplateDomainService struct {
	aiTemplateRepo      port.AITemplateRepository
	globalProductRepo   port.GlobalProductRepository
	aiGenerationService port.AIGenerationService
}

// NewAITemplateDomainService creates a new instance of the domain service
func NewAITemplateDomainService(
	aiTemplateRepo port.AITemplateRepository,
	globalProductRepo port.GlobalProductRepository,
	aiGenerationService port.AIGenerationService,
) *AITemplateDomainService {
	return &AITemplateDomainService{
		aiTemplateRepo:      aiTemplateRepo,
		globalProductRepo:   globalProductRepo,
		aiGenerationService: aiGenerationService,
	}
}

// ValidateTemplateForCreation validates a template before creation
func (s *AITemplateDomainService) ValidateTemplateForCreation(ctx context.Context, template *entity.AITemplate) error {
	// Check if template already exists for the same business type and tenant
	existing, err := s.aiTemplateRepo.FindByBusinessTypeAndTenant(ctx, template.BusinessTypeID, template.TenantID)
	if err != nil {
		return fmt.Errorf("error checking existing templates: %w", err)
	}

	if len(existing) > 0 {
		for _, t := range existing {
			if t.Name == template.Name {
				return exception.ErrTemplateAlreadyExists
			}
		}
	}

	// Validate generation type
	genType := value_object.GenerationType(template.GeneratedBy)
	if err := genType.Valid(); err != nil {
		return exception.NewValidationError("generated_by", err.Error())
	}

	// Validate AI generation params if it's AI generated
	if genType == value_object.GenerationTypeAI || genType == value_object.GenerationTypeHybrid {
		if template.AIGenerationParams == nil || len(template.AIGenerationParams) == 0 {
			return exception.NewValidationError("ai_generation_params", "AI generation parameters are required for AI-generated templates")
		}
	}

	return nil
}

// ValidateProductForTemplate validates a product before adding to template
func (s *AITemplateDomainService) ValidateProductForTemplate(ctx context.Context, templateProduct *entity.TemplateGlobalProduct) error {
	// Validate priority
	priority := value_object.ProductPriority(templateProduct.Priority)
	if err := priority.Valid(); err != nil {
		return exception.NewValidationError("priority", err.Error())
	}

	// Validate scores
	if templateProduct.RelevanceScore < 0 || templateProduct.RelevanceScore > 1 {
		return exception.NewValidationError("relevance_score", "relevance score must be between 0 and 1")
	}

	if templateProduct.CategoryMatchScore != nil && (*templateProduct.CategoryMatchScore < 0 || *templateProduct.CategoryMatchScore > 1) {
		return exception.NewValidationError("category_match_score", "category match score must be between 0 and 1")
	}

	if templateProduct.BrandMatchScore != nil && (*templateProduct.BrandMatchScore < 0 || *templateProduct.BrandMatchScore > 1) {
		return exception.NewValidationError("brand_match_score", "brand match score must be between 0 and 1")
	}

	if templateProduct.RegionalPreferenceScore != nil && (*templateProduct.RegionalPreferenceScore < 0 || *templateProduct.RegionalPreferenceScore > 1) {
		return exception.NewValidationError("regional_preference_score", "regional preference score must be between 0 and 1")
	}

	// Validate quantity suggestion
	if templateProduct.QuantitySuggestion < 1 {
		return exception.NewValidationError("quantity_suggestion", "quantity suggestion must be at least 1")
	}

	// Check if product already exists in template
	existingProducts, err := s.aiTemplateRepo.FindTemplateProducts(ctx, templateProduct.TemplateID)
	if err != nil {
		return fmt.Errorf("error checking existing products: %w", err)
	}

	for _, p := range existingProducts {
		if p.GlobalProductID == templateProduct.GlobalProductID {
			return exception.ErrDuplicateProduct
		}
	}

	return nil
}

// CalculateTemplatePerformance calculates performance metrics for a template
func (s *AITemplateDomainService) CalculateTemplatePerformance(ctx context.Context, templateID uuid.UUID) (map[string]float64, error) {
	// Get aggregated metrics
	metrics, err := s.aiTemplateRepo.GetAggregatedMetrics(ctx, templateID)
	if err != nil {
		return nil, fmt.Errorf("error getting aggregated metrics: %w", err)
	}

	// Get feedback summary
	feedbackSummary, err := s.aiTemplateRepo.GetFeedbackSummary(ctx, templateID)
	if err != nil {
		return nil, fmt.Errorf("error getting feedback summary: %w", err)
	}

	// Calculate additional metrics based on feedback
	totalFeedback := 0
	for _, count := range feedbackSummary {
		totalFeedback += count
	}

	if totalFeedback > 0 {
		// Calculate modification rate
		modifications := feedbackSummary[string(value_object.FeedbackActionRemoved)] + 
			feedbackSummary[string(value_object.FeedbackActionQuantityChanged)] + 
			feedbackSummary[string(value_object.FeedbackActionReplaced)]
		
		modificationRate := float64(modifications) / float64(totalFeedback)
		metrics[string(value_object.MetricTypeModificationRate)] = modificationRate

		// Calculate product retention rate
		kept := feedbackSummary[string(value_object.FeedbackActionKept)]
		retentionRate := float64(kept) / float64(totalFeedback)
		metrics[string(value_object.MetricTypeProductRetention)] = retentionRate
	}

	return metrics, nil
}

// ShouldRegenerateTemplate determines if a template should be regenerated based on performance
func (s *AITemplateDomainService) ShouldRegenerateTemplate(metrics map[string]float64) bool {
	// Check modification rate threshold (if more than 30% of products are modified)
	if modRate, ok := metrics[string(value_object.MetricTypeModificationRate)]; ok && modRate > 0.3 {
		return true
	}

	// Check satisfaction score threshold (if less than 70%)
	if satScore, ok := metrics[string(value_object.MetricTypeSatisfactionScore)]; ok && satScore < 0.7 {
		return true
	}

	// Check product retention rate (if less than 60% of products are kept)
	if retRate, ok := metrics[string(value_object.MetricTypeProductRetention)]; ok && retRate < 0.6 {
		return true
	}

	return false
}

// ValidateFeedback validates product feedback
func (s *AITemplateDomainService) ValidateFeedback(feedback *entity.AIProductFeedback) error {
	// Validate action
	action := value_object.FeedbackAction(feedback.Action)
	if err := action.Valid(); err != nil {
		return exception.NewValidationError("action", err.Error())
	}

	// Validate action-specific fields
	switch action {
	case value_object.FeedbackActionQuantityChanged:
		if feedback.OriginalQuantity == nil || feedback.NewQuantity == nil {
			return exception.NewValidationError("quantity", "original and new quantity are required for quantity change feedback")
		}
		if *feedback.NewQuantity < 0 {
			return exception.NewValidationError("new_quantity", "new quantity cannot be negative")
		}
	case value_object.FeedbackActionReplaced:
		if feedback.ReplacementProductID == nil {
			return exception.NewValidationError("replacement_product_id", "replacement product ID is required for replacement feedback")
		}
	}

	return nil
}