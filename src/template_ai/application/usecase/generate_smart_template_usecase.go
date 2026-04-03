package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid/v5"
	appPort "saas-mt-pim-service/src/template_ai/application/port"
	"saas-mt-pim-service/src/template_ai/application/request"
	"saas-mt-pim-service/src/template_ai/application/response"
	"saas-mt-pim-service/src/template_ai/domain/entity"
	"saas-mt-pim-service/src/template_ai/domain/exception"
	"saas-mt-pim-service/src/template_ai/domain/port"
	"saas-mt-pim-service/src/template_ai/domain/value_object"
)

// GenerateSmartTemplateUseCase handles AI-powered template generation
type GenerateSmartTemplateUseCase struct {
	aiTemplateRepo      port.AITemplateRepository
	aiGenerationService port.AIGenerationService
	domainService       port.AITemplateDomainServicePort
	mapper              appPort.TemplateMapperPort
}

// NewGenerateSmartTemplateUseCase creates a new instance of the use case
func NewGenerateSmartTemplateUseCase(
	aiTemplateRepo port.AITemplateRepository,
	aiGenerationService port.AIGenerationService,
	domainService port.AITemplateDomainServicePort,
	mapper appPort.TemplateMapperPort,
) *GenerateSmartTemplateUseCase {
	return &GenerateSmartTemplateUseCase{
		aiTemplateRepo:      aiTemplateRepo,
		aiGenerationService: aiGenerationService,
		domainService:       domainService,
		mapper:              mapper,
	}
}

// Execute generates a new smart template using AI
func (uc *GenerateSmartTemplateUseCase) Execute(
	ctx context.Context,
	req *request.GenerateSmartTemplateRequest,
) (*response.GenerateSmartTemplateResponse, error) {
	startTime := time.Now()

	// Validate request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Parse IDs
	businessTypeID, err := uuid.FromString(req.BusinessTypeID)
	if err != nil {
		return nil, exception.NewValidationError("business_type_id", "invalid UUID format")
	}

	var tenantID *uuid.UUID
	if req.TenantID != nil && *req.TenantID != "" {
		tid, err := uuid.FromString(*req.TenantID)
		if err != nil {
			return nil, exception.NewValidationError("tenant_id", "invalid UUID format")
		}
		tenantID = &tid
	}

	// Create generation parameters
	genParams := &entity.GenerationParams{
		BusinessTypeID:      businessTypeID,
		TenantID:            tenantID,
		ProductCount:        req.ProductCount,
		RegionalPreferences: req.RegionalPreferences,
		CategoryPriorities:  req.CategoryPriorities,
		BrandPreferences:    req.BrandPreferences,
		CustomConstraints:   req.CustomConstraints,
	}

	// Add budget range if provided
	if req.BudgetRange != nil {
		genParams.BudgetRange = &entity.BudgetRange{
			Min:      req.BudgetRange.Min,
			Max:      req.BudgetRange.Max,
			Currency: req.BudgetRange.Currency,
		}
	}

	// Create generation history entry
	historyID, err := uuid.NewV4()
	if err != nil {
		return nil, exception.NewInternalError("failed to generate history ID: " + err.Error())
	}
	generationHistory := &entity.AIGenerationHistory{
		ID:               historyID,
		TenantID:         tenantID,
		BusinessTypeID:   businessTypeID,
		GenerationParams: req.CustomConstraints, // Store full request params
		GenerationStatus: string(value_object.GenerationStatusProcessing),
		CreatedAt:        time.Now(),
	}

	// Save generation history
	if err := uc.aiTemplateRepo.SaveGenerationHistory(ctx, generationHistory); err != nil {
		return nil, fmt.Errorf("failed to save generation history: %w", err)
	}

	// Convert GenerationParams to GenerationRequest
	genRequest := port.GenerationRequest{
		BusinessTypeID:   genParams.BusinessTypeID,
		TemplateName:     req.Name,
		TargetSize:       "medium", // Default value
		Preferences:      req.RegionalPreferences,
		Constraints:      genParams.CustomConstraints,
		OptimizationGoal: "balanced",
		Region:           "",
		Season:           "",
		BudgetRange:      fmt.Sprintf("%.2f-%.2f", req.BudgetRange.Min, req.BudgetRange.Max),
	}

	// Generate template using AI service
	template, err := uc.aiGenerationService.GenerateTemplate(ctx, genRequest)
	if err != nil {
		// Update generation history with error
		generationHistory.GenerationStatus = string(value_object.GenerationStatusFailed)
		errorMsg := err.Error()
		generationHistory.ErrorMessage = &errorMsg
		uc.aiTemplateRepo.SaveGenerationHistory(ctx, generationHistory)
		
		return nil, exception.NewTemplateGenerationError("AI generation failed", err)
	}

	// Set template properties
	template.Name = req.Name
	template.Description = req.Description
	template.GeneratedBy = string(value_object.GenerationTypeAI)
	template.AIGenerationParams = map[string]interface{}{
		"product_count":        req.ProductCount,
		"budget_range":         req.BudgetRange,
		"regional_preferences": req.RegionalPreferences,
		"category_priorities":  req.CategoryPriorities,
		"brand_preferences":    req.BrandPreferences,
		"optimization_goals":   req.OptimizationGoals,
	}

	// Validate template before saving
	if err := uc.domainService.ValidateTemplateForCreation(ctx, template); err != nil {
		return nil, err
	}

	// Save template
	if err := uc.aiTemplateRepo.Save(ctx, template); err != nil {
		return nil, fmt.Errorf("failed to save template: %w", err)
	}

	// Generate product recommendations
	// TODO: This method doesn't exist in the interface yet
	// products, err := uc.aiGenerationService.GenerateProductRecommendations(ctx, template.ID, req.ProductCount)
	// if err != nil {
	// 	// Template is saved but product generation failed
	// 	// We don't delete the template, just return error
	// 	return nil, exception.NewTemplateGenerationError("Product recommendations generation failed", err)
	// }
	
	// For now, use empty products list
	var products []*entity.TemplateGlobalProduct

	// Validate and save each product recommendation
	var savedProducts []*entity.TemplateGlobalProduct
	for _, product := range products {
		product.TemplateID = template.ID
		
		// Validate product
		if err := uc.domainService.ValidateProductForTemplate(ctx, product); err != nil {
			// Skip invalid products but continue with others
			continue
		}

		// Save product
		if err := uc.aiTemplateRepo.SaveTemplateProduct(ctx, product); err != nil {
			// Log error but continue
			continue
		}

		savedProducts = append(savedProducts, product)
	}

	// Update generation history
	generationTimeMs := int(time.Since(startTime).Milliseconds())
	generationHistory.TemplateID = &template.ID
	generationHistory.GenerationStatus = string(value_object.GenerationStatusCompleted)
	generationHistory.GenerationTimeMs = &generationTimeMs
	generationHistory.GeneratedContent = map[string]interface{}{
		"template_id":    template.ID,
		"product_count":  len(savedProducts),
		"template_name":  template.Name,
	}
	now := time.Now()
	generationHistory.AppliedAt = &now

	if err := uc.aiTemplateRepo.SaveGenerationHistory(ctx, generationHistory); err != nil {
		// Log error but don't fail the operation
		fmt.Printf("Warning: failed to update generation history: %v\n", err)
	}

	// Convert to response
	templateResponse := uc.mapper.ToTemplateResponse(template, len(savedProducts))
	
	// Convert products to response
	var productResponses []response.TemplateProductResponse
	for _, product := range savedProducts {
		productResponses = append(productResponses, *uc.mapper.ToTemplateProductResponse(product))
	}

	// Create generation summary
	summary := uc.mapper.ToGenerationSummaryResponse(
		savedProducts,
		generationTimeMs,
		"gpt-4", // This would come from AI service
	)

	return &response.GenerateSmartTemplateResponse{
		Template:          templateResponse,
		Products:          productResponses,
		GenerationSummary: summary,
		Success:           true,
		Message:           fmt.Sprintf("Template generated successfully with %d products", len(savedProducts)),
	}, nil
}