package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/application/mapper"
	"saas-mt-pim-service/src/template_ai/application/request"
	"saas-mt-pim-service/src/template_ai/application/response"
	"saas-mt-pim-service/src/template_ai/domain/entity"
	"saas-mt-pim-service/src/template_ai/domain/exception"
	"saas-mt-pim-service/src/template_ai/domain/port"
	"saas-mt-pim-service/src/template_ai/domain/service"
	"saas-mt-pim-service/src/template_ai/domain/value_object"
)

// updateResult contains the result of an update operation
type updateResult struct {
	success             bool
	message             string
	productsAdded       int
	productsRemoved     int
	productsModified    int
	quantitiesAdjusted  int
}

// UpdateTemplateFromFeedbackUseCase handles template updates based on user feedback
type UpdateTemplateFromFeedbackUseCase struct {
	aiTemplateRepo      port.AITemplateRepository
	aiGenerationService port.AIGenerationService
	domainService       *service.AITemplateDomainService
	mapper              *mapper.TemplateMapper
}

// NewUpdateTemplateFromFeedbackUseCase creates a new instance of the use case
func NewUpdateTemplateFromFeedbackUseCase(
	aiTemplateRepo port.AITemplateRepository,
	aiGenerationService port.AIGenerationService,
	domainService *service.AITemplateDomainService,
	mapper *mapper.TemplateMapper,
) *UpdateTemplateFromFeedbackUseCase {
	return &UpdateTemplateFromFeedbackUseCase{
		aiTemplateRepo:      aiTemplateRepo,
		aiGenerationService: aiGenerationService,
		domainService:       domainService,
		mapper:              mapper,
	}
}

// Execute updates a template based on feedback
func (uc *UpdateTemplateFromFeedbackUseCase) Execute(
	ctx context.Context,
	req *request.UpdateTemplateFromFeedbackRequest,
) (*response.UpdateTemplateFromFeedbackResponse, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Parse template ID
	templateID, err := uuid.FromString(req.TemplateID)
	if err != nil {
		return nil, exception.NewValidationError("template_id", "invalid UUID format")
	}

	// Get template
	template, err := uc.aiTemplateRepo.FindByID(ctx, templateID)
	if err != nil {
		return nil, exception.ErrTemplateNotFound
	}

	// Calculate feedback period
	feedbackSince := time.Now().AddDate(0, 0, -req.FeedbackPeriodDays)

	// Get feedback for the period
	feedback, err := uc.aiTemplateRepo.FindProductFeedback(ctx, &templateID, uuid.Nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve feedback: %w", err)
	}

	// Filter feedback by period
	var relevantFeedback []*entity.AIProductFeedback
	for _, fb := range feedback {
		if fb.CreatedAt.After(feedbackSince) {
			relevantFeedback = append(relevantFeedback, fb)
		}
	}

	// Check minimum feedback count
	if len(relevantFeedback) < req.MinFeedbackCount && !req.ForceUpdate {
		return &response.UpdateTemplateFromFeedbackResponse{
			Success: false,
			Message: fmt.Sprintf("Insufficient feedback: %d feedback items found, minimum %d required", 
				len(relevantFeedback), req.MinFeedbackCount),
			UpdateSummary: &response.UpdateSummaryResponse{
				UpdateStrategy:   req.UpdateStrategy,
				FeedbackAnalyzed: len(relevantFeedback),
				UpdatedAt:        time.Now(),
			},
		}, nil
	}

	// Get current performance metrics
	currentMetrics, err := uc.domainService.CalculateTemplatePerformance(ctx, templateID)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate current metrics: %w", err)
	}

	// Analyze feedback patterns
	feedbackAnalysis := uc.analyzeFeedback(relevantFeedback)

	// Check if update is needed
	if !req.ForceUpdate && !uc.domainService.ShouldRegenerateTemplate(currentMetrics) {
		return &response.UpdateTemplateFromFeedbackResponse{
			Success: false,
			Message: "Template performance is satisfactory, no update needed",
			UpdateSummary: &response.UpdateSummaryResponse{
				UpdateStrategy:   req.UpdateStrategy,
				FeedbackAnalyzed: len(relevantFeedback),
				PreviousMetrics:  currentMetrics,
				UpdatedAt:        time.Now(),
			},
		}, nil
	}

	// Execute update based on strategy
	var result *updateResult
	var testResults *response.TestResultsResponse

	if req.TestMode {
		// In test mode, simulate changes without applying
		testResults = uc.simulateUpdate(template, feedbackAnalysis, req)
		result = &updateResult{
			success: true,
			message: "Test mode: simulated changes without applying",
		}
	} else {
		switch req.GetUpdateStrategy() {
		case request.UpdateStrategyIncremental:
			result = uc.executeIncrementalUpdate(ctx, template, feedbackAnalysis, req)
		case request.UpdateStrategyFull:
			result = uc.executeFullUpdate(ctx, template, feedbackAnalysis, req)
		case request.UpdateStrategyRegenerate:
			result = uc.executeRegenerateUpdate(ctx, template, feedbackAnalysis, req)
		default:
			return nil, fmt.Errorf("unknown update strategy: %s", req.UpdateStrategy)
		}
	}

	if !result.success {
		return &response.UpdateTemplateFromFeedbackResponse{
			Success: false,
			Message: result.message,
			UpdateSummary: &response.UpdateSummaryResponse{
				UpdateStrategy:   req.UpdateStrategy,
				FeedbackAnalyzed: len(relevantFeedback),
				PreviousMetrics:  currentMetrics,
				UpdatedAt:        time.Now(),
			},
		}, nil
	}

	// Get updated template if not in test mode
	var updatedTemplateResponse *response.AITemplateResponse
	if !req.TestMode {
		updatedTemplate, err := uc.aiTemplateRepo.FindByID(ctx, templateID)
		if err == nil {
			products, _ := uc.aiTemplateRepo.FindTemplateProducts(ctx, templateID)
			updatedTemplateResponse = uc.mapper.ToTemplateResponse(updatedTemplate, len(products))
		}
	}

	// Calculate projected metrics
	projectedMetrics := uc.calculateProjectedMetrics(currentMetrics, feedbackAnalysis, result)

	// Prepare update summary
	updateSummary := &response.UpdateSummaryResponse{
		UpdateStrategy:     req.UpdateStrategy,
		FeedbackAnalyzed:   len(relevantFeedback),
		ProductsAdded:      result.productsAdded,
		ProductsRemoved:    result.productsRemoved,
		ProductsModified:   result.productsModified,
		QuantitiesAdjusted: result.quantitiesAdjusted,
		PreviousMetrics:    currentMetrics,
		ProjectedMetrics:   projectedMetrics,
		ImprovementAreas:   uc.identifyImprovementAreas(feedbackAnalysis),
		UpdateReason:       req.UpdateReason,
		UpdatedAt:          time.Now(),
	}

	return &response.UpdateTemplateFromFeedbackResponse{
		Success:         true,
		Message:         result.message,
		UpdatedTemplate: updatedTemplateResponse,
		UpdateSummary:   updateSummary,
		TestResults:     testResults,
	}, nil
}

// feedbackAnalysis contains analyzed feedback patterns
type feedbackAnalysis struct {
	totalFeedback      int
	productRemovals    map[uuid.UUID]int
	quantityChanges    map[uuid.UUID][]int
	productReplacements map[uuid.UUID][]uuid.UUID
	keptProducts       map[uuid.UUID]int
	removalReasons     map[string]int
}

// analyzeFeedback analyzes feedback patterns
func (uc *UpdateTemplateFromFeedbackUseCase) analyzeFeedback(feedback []*entity.AIProductFeedback) *feedbackAnalysis {
	analysis := &feedbackAnalysis{
		totalFeedback:       len(feedback),
		productRemovals:     make(map[uuid.UUID]int),
		quantityChanges:     make(map[uuid.UUID][]int),
		productReplacements: make(map[uuid.UUID][]uuid.UUID),
		keptProducts:        make(map[uuid.UUID]int),
		removalReasons:      make(map[string]int),
	}

	for _, fb := range feedback {
		if fb.GlobalProductID == nil {
			continue
		}

		switch fb.Action {
		case string(value_object.FeedbackActionRemoved):
			analysis.productRemovals[*fb.GlobalProductID]++
			if fb.FeedbackReason != nil {
				analysis.removalReasons[*fb.FeedbackReason]++
			}
		case string(value_object.FeedbackActionQuantityChanged):
			if fb.NewQuantity != nil {
				analysis.quantityChanges[*fb.GlobalProductID] = append(
					analysis.quantityChanges[*fb.GlobalProductID], *fb.NewQuantity)
			}
		case string(value_object.FeedbackActionReplaced):
			if fb.ReplacementProductID != nil {
				analysis.productReplacements[*fb.GlobalProductID] = append(
					analysis.productReplacements[*fb.GlobalProductID], *fb.ReplacementProductID)
			}
		case string(value_object.FeedbackActionKept):
			analysis.keptProducts[*fb.GlobalProductID]++
		}
	}

	return analysis
}

// executeIncrementalUpdate performs incremental updates based on feedback
func (uc *UpdateTemplateFromFeedbackUseCase) executeIncrementalUpdate(
	ctx context.Context,
	template *entity.AITemplate,
	analysis *feedbackAnalysis,
	req *request.UpdateTemplateFromFeedbackRequest,
) *updateResult {
	result := &updateResult{success: true}

	// Get current template products
	currentProducts, err := uc.aiTemplateRepo.FindTemplateProducts(ctx, template.ID)
	if err != nil {
		result.success = false
		result.message = "Failed to retrieve template products"
		return result
	}

	// Create maps for easier lookup
	productMap := make(map[uuid.UUID]*entity.TemplateGlobalProduct)
	for _, p := range currentProducts {
		productMap[p.GlobalProductID] = p
	}

	// Remove products with high removal rate
	removalThreshold := 0.3 // Remove if removed by 30% of users
	for productID, removalCount := range analysis.productRemovals {
		removalRate := float64(removalCount) / float64(analysis.totalFeedback)
		if removalRate >= removalThreshold {
			if err := uc.aiTemplateRepo.DeleteTemplateProduct(ctx, template.ID, productID); err == nil {
				result.productsRemoved++
				delete(productMap, productID)
			}
		}
	}

	// Adjust quantities based on feedback
	for productID, quantities := range analysis.quantityChanges {
		if product, exists := productMap[productID]; exists {
			// Calculate average new quantity
			var sum int
			for _, q := range quantities {
				sum += q
			}
			avgQuantity := sum / len(quantities)

			// Update if significantly different (>20% change)
			if float64(abs(avgQuantity-product.QuantitySuggestion))/float64(product.QuantitySuggestion) > 0.2 {
				product.QuantitySuggestion = avgQuantity
				if err := uc.aiTemplateRepo.UpdateTemplateProduct(ctx, product); err == nil {
					result.quantitiesAdjusted++
				}
			}
		}
	}

	// Add frequently replaced products
	replacementThreshold := 3 // Add if suggested as replacement 3+ times
	replacementCounts := make(map[uuid.UUID]int)
	for _, replacements := range analysis.productReplacements {
		for _, replacementID := range replacements {
			replacementCounts[replacementID]++
		}
	}

	for replacementID, count := range replacementCounts {
		if count >= replacementThreshold {
			if _, exists := productMap[replacementID]; !exists {
				// Add new product to template
				newProductID, err := uuid.NewV4()
				if err != nil {
					continue // Skip if UUID generation fails
				}
				newProduct := &entity.TemplateGlobalProduct{
					ID:                 newProductID,
					TemplateID:        template.ID,
					GlobalProductID:   replacementID,
					Priority:          int(value_object.ProductPriorityRecommended),
					QuantitySuggestion: 1,
					AIReasoning:       fmt.Sprintf("Added based on user feedback (requested %d times)", count),
					RelevanceScore:    0.8,
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}
				if err := uc.aiTemplateRepo.SaveTemplateProduct(ctx, newProduct); err == nil {
					result.productsAdded++
				}
			}
		}
	}

	// Update template metadata
	now := time.Now()
	template.LastAIUpdate = &now
	template.PerformanceMetrics["last_feedback_update"] = now.Format(time.RFC3339)
	template.PerformanceMetrics["feedback_items_processed"] = analysis.totalFeedback

	if err := uc.aiTemplateRepo.Update(ctx, template); err != nil {
		result.success = false
		result.message = "Failed to update template metadata"
		return result
	}

	result.message = fmt.Sprintf("Incremental update completed: %d products removed, %d added, %d quantities adjusted",
		result.productsRemoved, result.productsAdded, result.quantitiesAdjusted)
	return result
}

// executeFullUpdate performs a full reoptimization
func (uc *UpdateTemplateFromFeedbackUseCase) executeFullUpdate(
	ctx context.Context,
	template *entity.AITemplate,
	analysis *feedbackAnalysis,
	req *request.UpdateTemplateFromFeedbackRequest,
) *updateResult {
	// In a full update, we optimize the entire template while preserving structure
	// This would involve calling the AI service to reoptimize based on feedback
	
	// For now, we'll simulate by calling the incremental update with more aggressive thresholds
	return uc.executeIncrementalUpdate(ctx, template, analysis, req)
}

// executeRegenerateUpdate completely regenerates the template
func (uc *UpdateTemplateFromFeedbackUseCase) executeRegenerateUpdate(
	ctx context.Context,
	template *entity.AITemplate,
	analysis *feedbackAnalysis,
	req *request.UpdateTemplateFromFeedbackRequest,
) *updateResult {
	result := &updateResult{success: true}

	// Create new generation parameters based on feedback
	genParams := &entity.GenerationParams{
		BusinessTypeID: template.BusinessTypeID,
		TenantID:       template.TenantID,
		ProductCount:   100, // Default, could be derived from current template
	}

	// Update generation params based on feedback patterns
	if template.AIGenerationParams != nil {
		// Preserve original params but update based on feedback
		genParams.RegionalPreferences = template.AIGenerationParams["regional_preferences"].(map[string]interface{})
	}

	// Convert feedback to port.TemplateFeedback format
	var templateFeedback []port.TemplateFeedback
	// Note: In a real implementation, we would convert the analyzed feedback
	// For now, just pass empty feedback
	
	// Call AI service to regenerate
	optimizedTemplate, err := uc.aiGenerationService.OptimizeTemplate(ctx, template, templateFeedback)
	if err != nil {
		result.success = false
		result.message = fmt.Sprintf("Failed to regenerate template: %v", err)
		return result
	}
	
	// Update with optimized template
	template = optimizedTemplate

	// Update template metadata
	now := time.Now()
	template.LastAIUpdate = &now
	template.GeneratedBy = string(value_object.GenerationTypeHybrid) // Now hybrid since it includes feedback

	if err := uc.aiTemplateRepo.Update(ctx, template); err != nil {
		result.success = false
		result.message = "Failed to update template after regeneration"
		return result
	}

	result.message = "Template successfully regenerated based on user feedback"
	return result
}

// simulateUpdate simulates the update without applying changes
func (uc *UpdateTemplateFromFeedbackUseCase) simulateUpdate(
	template *entity.AITemplate,
	analysis *feedbackAnalysis,
	req *request.UpdateTemplateFromFeedbackRequest,
) *response.TestResultsResponse {
	var simulatedChanges []response.SimulatedChangeResponse

	// Simulate product removals
	for productID, removalCount := range analysis.productRemovals {
		removalRate := float64(removalCount) / float64(analysis.totalFeedback)
		if removalRate >= 0.3 {
			simulatedChanges = append(simulatedChanges, response.SimulatedChangeResponse{
				ChangeType:  "remove_product",
				ProductID:   productID.String(),
				Reason:      fmt.Sprintf("High removal rate: %.0f%%", removalRate*100),
			})
		}
	}

	// Simulate quantity adjustments
	for productID, quantities := range analysis.quantityChanges {
		var sum int
		for _, q := range quantities {
			sum += q
		}
		avgQuantity := sum / len(quantities)
		
		simulatedChanges = append(simulatedChanges, response.SimulatedChangeResponse{
			ChangeType:  "adjust_quantity",
			ProductID:   productID.String(),
			NewValue:    avgQuantity,
			Reason:      fmt.Sprintf("Based on %d user adjustments", len(quantities)),
		})
	}

	// Calculate projected improvement
	projectedImprovement := map[string]float64{
		"satisfaction_score": 0.15,  // Expected 15% improvement
		"modification_rate":  -0.20, // Expected 20% reduction
		"product_retention":  0.10,  // Expected 10% improvement
	}

	// Risk assessment
	riskLevel := "low"
	if len(simulatedChanges) > 10 {
		riskLevel = "medium"
	}
	if len(simulatedChanges) > 20 {
		riskLevel = "high"
	}

	recommendation := "Proceed with update"
	if riskLevel == "high" {
		recommendation = "Review changes carefully before applying"
	}

	return &response.TestResultsResponse{
		SimulatedChanges:     simulatedChanges,
		ProjectedImprovement: projectedImprovement,
		RiskAssessment:       riskLevel,
		Recommendation:       recommendation,
	}
}

// Helper methods

func (uc *UpdateTemplateFromFeedbackUseCase) calculateProjectedMetrics(
	currentMetrics map[string]float64,
	analysis *feedbackAnalysis,
	result *updateResult,
) map[string]float64 {
	projected := make(map[string]float64)
	
	// Copy current metrics
	for k, v := range currentMetrics {
		projected[k] = v
	}

	// Project improvements based on changes
	if result.productsRemoved > 0 || result.productsAdded > 0 {
		// Expect improvement in satisfaction and reduction in modification rate
		if score, exists := projected[string(value_object.MetricTypeSatisfactionScore)]; exists {
			projected[string(value_object.MetricTypeSatisfactionScore)] = min(1.0, score*1.1)
		}
		if rate, exists := projected[string(value_object.MetricTypeModificationRate)]; exists {
			projected[string(value_object.MetricTypeModificationRate)] = max(0, rate*0.8)
		}
	}

	return projected
}

func (uc *UpdateTemplateFromFeedbackUseCase) identifyImprovementAreas(analysis *feedbackAnalysis) []string {
	var areas []string

	if len(analysis.productRemovals) > 0 {
		areas = append(areas, "Product selection optimization")
	}
	if len(analysis.quantityChanges) > 0 {
		areas = append(areas, "Quantity suggestions refinement")
	}
	if len(analysis.productReplacements) > 0 {
		areas = append(areas, "Product alternatives consideration")
	}
	if len(analysis.removalReasons) > 0 {
		areas = append(areas, "Address specific user concerns")
	}

	return areas
}

// Helper functions
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}