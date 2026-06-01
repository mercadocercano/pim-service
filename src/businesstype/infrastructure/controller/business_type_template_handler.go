package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"saas-mt-pim-service/src/businesstype/application/usecase"
	templateCriteria "saas-mt-pim-service/src/businesstype/infrastructure/criteria"
)

// BusinessTypeTemplateHandler maneja las requests HTTP para templates de business types
type BusinessTypeTemplateHandler struct {
	createUseCase    *usecase.CreateBusinessTypeTemplateUseCase
	updateUseCase    *usecase.UpdateBusinessTypeTemplateUseCase
	listUseCase      *usecase.ListBusinessTypeTemplatesUseCase
	getUseCase       *usecase.GetBusinessTypeTemplateUseCase
	deleteUseCase    *usecase.DeleteBusinessTypeTemplateUseCase
	analyticsUseCase *usecase.GetTemplateAnalyticsUseCase
	duplicateUseCase *usecase.DuplicateTemplateUseCase
	criteriaBuilder  *templateCriteria.BusinessTypeTemplateCriteriaBuilder
}

// NewBusinessTypeTemplateHandler crea una nueva instancia del handler
func NewBusinessTypeTemplateHandler(
	createUseCase *usecase.CreateBusinessTypeTemplateUseCase,
	updateUseCase *usecase.UpdateBusinessTypeTemplateUseCase,
	listUseCase *usecase.ListBusinessTypeTemplatesUseCase,
	getUseCase *usecase.GetBusinessTypeTemplateUseCase,
	deleteUseCase *usecase.DeleteBusinessTypeTemplateUseCase,
) *BusinessTypeTemplateHandler {
	return &BusinessTypeTemplateHandler{
		createUseCase:   createUseCase,
		updateUseCase:   updateUseCase,
		listUseCase:     listUseCase,
		getUseCase:      getUseCase,
		deleteUseCase:   deleteUseCase,
		criteriaBuilder: templateCriteria.NewBusinessTypeTemplateCriteriaBuilder(),
	}
}

// WithAnalyticsUseCase agrega el use case de analytics al handler
func (h *BusinessTypeTemplateHandler) WithAnalyticsUseCase(uc *usecase.GetTemplateAnalyticsUseCase) *BusinessTypeTemplateHandler {
	h.analyticsUseCase = uc
	return h
}

// WithDuplicateUseCase agrega el use case de duplicación al handler
func (h *BusinessTypeTemplateHandler) WithDuplicateUseCase(uc *usecase.DuplicateTemplateUseCase) *BusinessTypeTemplateHandler {
	h.duplicateUseCase = uc
	return h
}

// RegisterRoutes registra las rutas del handler
func (h *BusinessTypeTemplateHandler) RegisterRoutes(router *gin.RouterGroup) {
	templates := router.Group("/business-type-templates")
	{
		templates.POST("", h.CreateTemplate)
		templates.GET("", h.ListTemplates)
		templates.GET("/:id", h.GetTemplate)
		templates.PUT("/:id", h.UpdateTemplate)
		templates.DELETE("/:id", h.DeleteTemplate)
		templates.GET("/:id/analytics", h.GetTemplateAnalytics)
		templates.POST("/:id/duplicate", h.DuplicateTemplate)
	}
}

// CreateTemplate crea un nuevo template
func (h *BusinessTypeTemplateHandler) CreateTemplate(c *gin.Context) {
	var req usecase.CreateTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	if req.BusinessTypeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "business_type_id es requerido"})
		return
	}

	template, err := h.createUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"template": template})
}

// ListTemplates lista templates con filtros y paginación usando criteria
func (h *BusinessTypeTemplateHandler) ListTemplates(c *gin.Context) {
	validCriteria := h.criteriaBuilder.BuildValidated(c)

	result, err := h.listUseCase.Execute(c.Request.Context(), validCriteria)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetTemplate obtiene un template por ID
func (h *BusinessTypeTemplateHandler) GetTemplate(c *gin.Context) {
	id := c.Param("id")

	template, err := h.getUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"template": template})
}

// UpdateTemplate actualiza un template existente
func (h *BusinessTypeTemplateHandler) UpdateTemplate(c *gin.Context) {
	id := c.Param("id")

	var req usecase.UpdateTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	req.ID = id

	template, err := h.updateUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"template": template})
}

// DeleteTemplate elimina un template
func (h *BusinessTypeTemplateHandler) DeleteTemplate(c *gin.Context) {
	id := c.Param("id")

	err := h.deleteUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetTemplateAnalytics retorna las analíticas de uso de un template
func (h *BusinessTypeTemplateHandler) GetTemplateAnalytics(c *gin.Context) {
	if h.analyticsUseCase == nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "analytics no configurado"})
		return
	}

	id := c.Param("id")
	analytics, err := h.analyticsUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "template no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, analytics)
}

// DuplicateTemplate duplica un template existente
func (h *BusinessTypeTemplateHandler) DuplicateTemplate(c *gin.Context) {
	if h.duplicateUseCase == nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "duplicate no configurado"})
		return
	}

	id := c.Param("id")

	var body struct {
		NewName string `json:"new_name"`
	}
	// No forzamos binding; new_name es opcional
	_ = c.ShouldBindJSON(&body)

	req := usecase.DuplicateTemplateRequest{
		TemplateID: id,
		NewName:    body.NewName,
	}

	template, err := h.duplicateUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "template no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"template": template})
}
