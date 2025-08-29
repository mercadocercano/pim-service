package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"saas-mt-pim-service/src/template_ai/application/request"
	"saas-mt-pim-service/src/template_ai/application/usecase"
)

// AITemplateController handles HTTP requests for AI template operations
type AITemplateController struct {
	generateSmartTemplateUseCase       *usecase.GenerateSmartTemplateUseCase
	applyDynamicTemplateUseCase        *usecase.ApplyDynamicTemplateUseCase
	analyzeTemplatePerformanceUseCase  *usecase.AnalyzeTemplatePerformanceUseCase
	updateTemplateFromFeedbackUseCase  *usecase.UpdateTemplateFromFeedbackUseCase
}

// NewAITemplateController creates a new instance of the controller
func NewAITemplateController(
	generateSmartTemplateUseCase *usecase.GenerateSmartTemplateUseCase,
	applyDynamicTemplateUseCase *usecase.ApplyDynamicTemplateUseCase,
	analyzeTemplatePerformanceUseCase *usecase.AnalyzeTemplatePerformanceUseCase,
	updateTemplateFromFeedbackUseCase *usecase.UpdateTemplateFromFeedbackUseCase,
) *AITemplateController {
	return &AITemplateController{
		generateSmartTemplateUseCase:       generateSmartTemplateUseCase,
		applyDynamicTemplateUseCase:        applyDynamicTemplateUseCase,
		analyzeTemplatePerformanceUseCase:  analyzeTemplatePerformanceUseCase,
		updateTemplateFromFeedbackUseCase:  updateTemplateFromFeedbackUseCase,
	}
}

// RegisterRoutes registers all AI template routes
func (c *AITemplateController) RegisterRoutes(router *gin.RouterGroup) {
	templates := router.Group("/templates")
	{
		// AI-powered template generation
		templates.POST("/generate", c.GenerateSmartTemplate)
		
		// Apply template to tenant catalog
		templates.POST("/:id/apply", c.ApplyDynamicTemplate)
		
		// Performance analysis
		templates.GET("/:id/performance", c.GetTemplatePerformance)
		
		// Feedback-based updates
		templates.POST("/update-from-feedback", c.UpdateTemplateFromFeedback)
	}
}

// GenerateSmartTemplate handles AI-powered template generation
// @Summary Generate a smart template using AI
// @Description Generates an optimized product template based on business type and constraints
// @Tags AI Templates
// @Accept json
// @Produce json
// @Param X-Tenant-ID header string false "Tenant ID"
// @Param request body request.GenerateSmartTemplateRequest true "Template generation parameters"
// @Success 200 {object} response.GenerateSmartTemplateResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/templates/generate [post]
func (c *AITemplateController) GenerateSmartTemplate(ctx *gin.Context) {
	var req request.GenerateSmartTemplateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get tenant ID from header if not in request
	if req.TenantID == nil || *req.TenantID == "" {
		if tenantID := ctx.GetHeader("X-Tenant-ID"); tenantID != "" {
			req.TenantID = &tenantID
		}
	}

	response, err := c.generateSmartTemplateUseCase.Execute(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// ApplyDynamicTemplate handles template application to tenant catalog
// @Summary Apply a template to tenant catalog
// @Description Applies a template's products to a tenant's catalog with various modes
// @Tags AI Templates
// @Accept json
// @Produce json
// @Param X-Tenant-ID header string true "Tenant ID"
// @Param request body request.ApplyDynamicTemplateRequest true "Template application parameters"
// @Success 200 {object} response.ApplyDynamicTemplateResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/templates/{id}/apply [post]
func (c *AITemplateController) ApplyDynamicTemplate(ctx *gin.Context) {
	templateID := ctx.Param("id")
	if templateID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "template id is required"})
		return
	}

	var req request.ApplyDynamicTemplateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set template ID from URL param
	req.TemplateID = templateID

	// Override tenant ID from header if present
	if tenantID := ctx.GetHeader("X-Tenant-ID"); tenantID != "" {
		req.TenantID = tenantID
	}

	response, err := c.applyDynamicTemplateUseCase.Execute(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetTemplatePerformance handles template performance retrieval
// @Summary Get template performance metrics
// @Description Retrieves template performance metrics and analytics
// @Tags AI Templates
// @Accept json
// @Produce json
// @Param id path string true "Template ID"
// @Success 200 {object} response.TemplatePerformanceResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/templates/{id}/performance [get]
func (c *AITemplateController) GetTemplatePerformance(ctx *gin.Context) {
	templateID := ctx.Param("id")
	if templateID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "template id is required"})
		return
	}

	req := request.AnalyzeTemplatePerformanceRequest{
		TemplateID: templateID,
		// Add default date range if needed
		PeriodStart: nil, // Will use defaults in use case
		PeriodEnd:   nil,
	}

	response, err := c.analyzeTemplatePerformanceUseCase.Execute(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// UpdateTemplateFromFeedback handles template updates based on user feedback
// @Summary Update template based on feedback
// @Description Updates a template using collected user feedback and AI optimization
// @Tags AI Templates
// @Accept json
// @Produce json
// @Param request body request.UpdateTemplateFromFeedbackRequest true "Update parameters"
// @Success 200 {object} response.UpdateTemplateFromFeedbackResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/templates/update-from-feedback [post]
func (c *AITemplateController) UpdateTemplateFromFeedback(ctx *gin.Context) {
	var req request.UpdateTemplateFromFeedbackRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.updateTemplateFromFeedbackUseCase.Execute(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}