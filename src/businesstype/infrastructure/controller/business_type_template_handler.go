package controller

import (
	"net/http"

	"pim/src/businesstype/application/usecase"
	templateCriteria "pim/src/businesstype/infrastructure/criteria"

	"github.com/gin-gonic/gin"
)

// BusinessTypeTemplateHandler maneja las requests HTTP para templates de business types
type BusinessTypeTemplateHandler struct {
	createUseCase   *usecase.CreateBusinessTypeTemplateUseCase
	updateUseCase   *usecase.UpdateBusinessTypeTemplateUseCase
	listUseCase     *usecase.ListBusinessTypeTemplatesUseCase
	getUseCase      *usecase.GetBusinessTypeTemplateUseCase
	deleteUseCase   *usecase.DeleteBusinessTypeTemplateUseCase
	criteriaBuilder *templateCriteria.BusinessTypeTemplateCriteriaBuilder
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

// RegisterRoutes registra las rutas del handler
func (h *BusinessTypeTemplateHandler) RegisterRoutes(router *gin.RouterGroup) {
	templates := router.Group("/business-type-templates")
	{
		templates.POST("", h.CreateTemplate)
		templates.GET("", h.ListTemplates)
		templates.GET("/:id", h.GetTemplate)
		templates.PUT("/:id", h.UpdateTemplate)
		templates.DELETE("/:id", h.DeleteTemplate)
	}
}

// CreateTemplate crea un nuevo template
func (h *BusinessTypeTemplateHandler) CreateTemplate(c *gin.Context) {
	var req usecase.CreateTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	template, err := h.createUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"template": template,
	})
}

// ListTemplates lista templates con filtros y paginación usando criteria
func (h *BusinessTypeTemplateHandler) ListTemplates(c *gin.Context) {
	// Construir criterios validados desde la request
	validCriteria := h.criteriaBuilder.BuildValidated(c)

	// Ejecutar búsqueda usando criteria
	result, err := h.listUseCase.Execute(c.Request.Context(), validCriteria)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retornar respuesta estándar
	c.JSON(http.StatusOK, result)
}

// GetTemplate obtiene un template por ID
func (h *BusinessTypeTemplateHandler) GetTemplate(c *gin.Context) {
	id := c.Param("id")

	template, err := h.getUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"template": template,
	})
}

// UpdateTemplate actualiza un template existente
func (h *BusinessTypeTemplateHandler) UpdateTemplate(c *gin.Context) {
	id := c.Param("id")

	var req usecase.UpdateTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Asignar el ID del parámetro de ruta
	req.ID = id

	template, err := h.updateUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"template": template,
	})
}

// DeleteTemplate elimina un template
func (h *BusinessTypeTemplateHandler) DeleteTemplate(c *gin.Context) {
	id := c.Param("id")

	err := h.deleteUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Template deleted successfully",
	})
}

