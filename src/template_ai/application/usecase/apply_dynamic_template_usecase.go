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
)

// ApplyDynamicTemplateUseCase handles applying templates to tenant catalogs
type ApplyDynamicTemplateUseCase struct {
	aiTemplateRepo    port.AITemplateRepository
	globalProductRepo port.GlobalProductRepository
	domainService     *service.AITemplateDomainService
	mapper            *mapper.TemplateMapper
	// In a real implementation, we would also have:
	// productService    port.ProductService
	// categoryService   port.CategoryService
	// brandService      port.BrandService
}

// NewApplyDynamicTemplateUseCase creates a new instance of the use case
func NewApplyDynamicTemplateUseCase(
	aiTemplateRepo port.AITemplateRepository,
	globalProductRepo port.GlobalProductRepository,
	domainService *service.AITemplateDomainService,
	mapper *mapper.TemplateMapper,
) *ApplyDynamicTemplateUseCase {
	return &ApplyDynamicTemplateUseCase{
		aiTemplateRepo:    aiTemplateRepo,
		globalProductRepo: globalProductRepo,
		domainService:     domainService,
		mapper:            mapper,
	}
}

// Execute applies a template to a tenant's catalog
func (uc *ApplyDynamicTemplateUseCase) Execute(
	ctx context.Context,
	req *request.ApplyDynamicTemplateRequest,
) (*response.ApplyDynamicTemplateResponse, error) {
	startTime := time.Now()

	// Validate request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Parse IDs
	templateID, err := uuid.FromString(req.TemplateID)
	if err != nil {
		return nil, exception.NewValidationError("template_id", "invalid UUID format")
	}

	tenantID, err := uuid.FromString(req.TenantID)
	if err != nil {
		return nil, exception.NewValidationError("tenant_id", "invalid UUID format")
	}

	// Get template
	template, err := uc.aiTemplateRepo.FindByID(ctx, templateID)
	if err != nil {
		return nil, exception.ErrTemplateNotFound
	}

	// Get template products
	templateProducts, err := uc.aiTemplateRepo.FindTemplateProducts(ctx, templateID)
	if err != nil {
		return nil, fmt.Errorf("failed to get template products: %w", err)
	}

	// Apply filters and overrides
	productsToApply := uc.filterAndOverrideProducts(templateProducts, req)

	// Get global product details
	productIDs := make([]uuid.UUID, 0, len(productsToApply))
	for _, p := range productsToApply {
		productIDs = append(productIDs, p.GlobalProductID)
	}

	globalProducts, err := uc.globalProductRepo.FindByIDs(ctx, productIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get global products: %w", err)
	}

	// Track application results
	applicationDetails := &response.ApplicationDetailsResponse{
		TemplateID:      templateID,
		TemplateName:    template.Name,
		TenantID:        tenantID,
		ApplyMode:       req.ApplyMode,
		StartedAt:       startTime,
		ProductsCreated: []string{},
		ProductsUpdated: []string{},
		ProductsSkipped: []string{},
	}

	var errors []response.ApplicationErrorResponse
	appliedCount := 0
	skippedCount := 0
	failedCount := 0

	// Apply each product based on mode
	for _, templateProduct := range productsToApply {
		globalProduct, exists := globalProducts[templateProduct.GlobalProductID]
		if !exists {
			errors = append(errors, response.ApplicationErrorResponse{
				ProductID: templateProduct.GlobalProductID.String(),
				Error:     "Global product not found",
				ErrorType: "not_found",
				Timestamp: time.Now(),
			})
			failedCount++
			continue
		}

		// Apply product based on mode
		switch req.GetApplyMode() {
		case request.ApplyModeFull:
			// In full mode, we would clear existing products first
			// Then create all template products
			if err := uc.createProductInTenant(ctx, tenantID, templateProduct, globalProduct, req); err != nil {
				errors = append(errors, response.ApplicationErrorResponse{
					ProductID:   templateProduct.GlobalProductID.String(),
					ProductName: uc.getProductName(globalProduct),
					Error:       err.Error(),
					ErrorType:   "creation_failed",
					Timestamp:   time.Now(),
				})
				failedCount++
			} else {
				applicationDetails.ProductsCreated = append(applicationDetails.ProductsCreated, templateProduct.GlobalProductID.String())
				appliedCount++
			}

		case request.ApplyModePartial:
			// In partial mode, only add products that don't exist
			if uc.productExistsInTenant(ctx, tenantID, templateProduct.GlobalProductID) {
				applicationDetails.ProductsSkipped = append(applicationDetails.ProductsSkipped, templateProduct.GlobalProductID.String())
				skippedCount++
			} else {
				if err := uc.createProductInTenant(ctx, tenantID, templateProduct, globalProduct, req); err != nil {
					errors = append(errors, response.ApplicationErrorResponse{
						ProductID:   templateProduct.GlobalProductID.String(),
						ProductName: uc.getProductName(globalProduct),
						Error:       err.Error(),
						ErrorType:   "creation_failed",
						Timestamp:   time.Now(),
					})
					failedCount++
				} else {
					applicationDetails.ProductsCreated = append(applicationDetails.ProductsCreated, templateProduct.GlobalProductID.String())
					appliedCount++
				}
			}

		case request.ApplyModeMerge:
			// In merge mode, create new products and update quantities for existing ones
			if uc.productExistsInTenant(ctx, tenantID, templateProduct.GlobalProductID) {
				if err := uc.updateProductQuantity(ctx, tenantID, templateProduct); err != nil {
					errors = append(errors, response.ApplicationErrorResponse{
						ProductID:   templateProduct.GlobalProductID.String(),
						ProductName: uc.getProductName(globalProduct),
						Error:       err.Error(),
						ErrorType:   "update_failed",
						Timestamp:   time.Now(),
					})
					failedCount++
				} else {
					applicationDetails.ProductsUpdated = append(applicationDetails.ProductsUpdated, templateProduct.GlobalProductID.String())
					appliedCount++
				}
			} else {
				if err := uc.createProductInTenant(ctx, tenantID, templateProduct, globalProduct, req); err != nil {
					errors = append(errors, response.ApplicationErrorResponse{
						ProductID:   templateProduct.GlobalProductID.String(),
						ProductName: uc.getProductName(globalProduct),
						Error:       err.Error(),
						ErrorType:   "creation_failed",
						Timestamp:   time.Now(),
					})
					failedCount++
				} else {
					applicationDetails.ProductsCreated = append(applicationDetails.ProductsCreated, templateProduct.GlobalProductID.String())
					appliedCount++
				}
			}
		}
	}

	// Save feedback for learning
	if len(req.ExcludeProducts) > 0 || len(req.OverrideQuantities) > 0 {
		uc.saveFeedback(ctx, tenantID, templateID, req)
	}

	// Complete application details
	applicationDetails.CompletedAt = time.Now()
	applicationDetails.DurationMs = int(time.Since(startTime).Milliseconds())

	// Prepare response
	resp := &response.ApplyDynamicTemplateResponse{
		Success:            appliedCount > 0,
		Message:            uc.generateApplicationMessage(appliedCount, skippedCount, failedCount),
		AppliedProducts:    appliedCount,
		SkippedProducts:    skippedCount,
		FailedProducts:     failedCount,
		ApplicationDetails: applicationDetails,
	}

	if len(errors) > 0 {
		resp.Errors = errors
	}

	return resp, nil
}

// filterAndOverrideProducts applies filters and overrides to template products
func (uc *ApplyDynamicTemplateUseCase) filterAndOverrideProducts(
	templateProducts []*entity.TemplateGlobalProduct,
	req *request.ApplyDynamicTemplateRequest,
) []*entity.TemplateGlobalProduct {
	// Create exclusion map
	excludeMap := make(map[string]bool)
	for _, id := range req.ExcludeProducts {
		excludeMap[id] = true
	}

	// Create manual selection map
	manualSelections := make(map[string]request.ProductSelection)
	for _, selection := range req.ProductSelections {
		manualSelections[selection.GlobalProductID] = selection
	}

	var result []*entity.TemplateGlobalProduct

	// Process template products
	for _, product := range templateProducts {
		productIDStr := product.GlobalProductID.String()

		// Skip excluded products
		if excludeMap[productIDStr] {
			continue
		}

		// Check for manual selection
		if selection, exists := manualSelections[productIDStr]; exists {
			// Override with manual selection
			product.QuantitySuggestion = selection.Quantity
			product.Priority = selection.Priority
			delete(manualSelections, productIDStr) // Remove from map
		} else if override, exists := req.OverrideQuantities[productIDStr]; exists {
			// Apply quantity override
			product.QuantitySuggestion = override
		}

		// Apply category adjustments if specified
		// This would require category information from global product
		// For now, we'll just add the product
		result = append(result, product)
	}

	// Add any remaining manual selections that weren't in template
	for _, selection := range manualSelections {
		globalProductID, _ := uuid.FromString(selection.GlobalProductID)
		templateProductID, err := uuid.NewV4()
		if err != nil {
			continue // Skip if UUID generation fails
		}
		result = append(result, &entity.TemplateGlobalProduct{
			ID:                 templateProductID,
			GlobalProductID:    globalProductID,
			Priority:           selection.Priority,
			QuantitySuggestion: selection.Quantity,
			AIReasoning:        "Manually selected by user",
			RelevanceScore:     1.0, // Maximum relevance for manual selections
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		})
	}

	return result
}

// Helper methods (these would typically call external services)

func (uc *ApplyDynamicTemplateUseCase) createProductInTenant(
	ctx context.Context,
	tenantID uuid.UUID,
	templateProduct *entity.TemplateGlobalProduct,
	globalProduct interface{},
	req *request.ApplyDynamicTemplateRequest,
) error {
	// In a real implementation, this would:
	// 1. Map global product to tenant product
	// 2. Create product via product service
	// 3. Create variants if requested
	// 4. Apply pricing if requested
	// For now, we'll simulate success
	return nil
}

func (uc *ApplyDynamicTemplateUseCase) productExistsInTenant(
	ctx context.Context,
	tenantID uuid.UUID,
	globalProductID uuid.UUID,
) bool {
	// In a real implementation, this would check via product service
	// For now, we'll return false
	return false
}

func (uc *ApplyDynamicTemplateUseCase) updateProductQuantity(
	ctx context.Context,
	tenantID uuid.UUID,
	templateProduct *entity.TemplateGlobalProduct,
) error {
	// In a real implementation, this would update via product service
	// For now, we'll simulate success
	return nil
}

func (uc *ApplyDynamicTemplateUseCase) getProductName(globalProduct interface{}) string {
	// In a real implementation, this would extract name from global product
	return "Product"
}

func (uc *ApplyDynamicTemplateUseCase) saveFeedback(
	ctx context.Context,
	tenantID uuid.UUID,
	templateID uuid.UUID,
	req *request.ApplyDynamicTemplateRequest,
) {
	// Save exclusion feedback
	for _, productID := range req.ExcludeProducts {
		globalProductID, _ := uuid.FromString(productID)
		feedbackID, err := uuid.NewV4()
		if err != nil {
			continue // Skip if UUID generation fails
		}
		feedback := &entity.AIProductFeedback{
			ID:              feedbackID,
			TenantID:        tenantID,
			TemplateID:      &templateID,
			GlobalProductID: &globalProductID,
			Action:          "removed",
			CreatedAt:       time.Now(),
		}
		uc.aiTemplateRepo.SaveProductFeedback(ctx, feedback)
	}

	// Save quantity override feedback
	for productID, newQuantity := range req.OverrideQuantities {
		globalProductID, _ := uuid.FromString(productID)
		feedbackID, err := uuid.NewV4()
		if err != nil {
			continue // Skip if UUID generation fails
		}
		feedback := &entity.AIProductFeedback{
			ID:              feedbackID,
			TenantID:        tenantID,
			TemplateID:      &templateID,
			GlobalProductID: &globalProductID,
			Action:          "quantity_changed",
			NewQuantity:     &newQuantity,
			CreatedAt:       time.Now(),
		}
		uc.aiTemplateRepo.SaveProductFeedback(ctx, feedback)
	}
}

func (uc *ApplyDynamicTemplateUseCase) generateApplicationMessage(applied, skipped, failed int) string {
	if failed == 0 && skipped == 0 {
		return fmt.Sprintf("Successfully applied %d products from template", applied)
	} else if failed == 0 {
		return fmt.Sprintf("Applied %d products, skipped %d existing products", applied, skipped)
	} else {
		return fmt.Sprintf("Applied %d products, skipped %d, failed %d", applied, skipped, failed)
	}
}