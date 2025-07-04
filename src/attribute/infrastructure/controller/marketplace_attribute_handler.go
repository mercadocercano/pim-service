package controller

import (
	"net/http"

	"pim/src/attribute/application/request"
	"pim/src/attribute/application/response"
	"pim/src/attribute/application/usecase"

	"github.com/gin-gonic/gin"
)

// MarketplaceAttributeHandler maneja las peticiones HTTP para atributos marketplace
type MarketplaceAttributeHandler struct {
	createUseCase  *usecase.CreateMarketplaceAttributeUseCase
	listUseCase    *usecase.ListMarketplaceAttributesUseCase
	getByIDUseCase *usecase.GetMarketplaceAttributeByIDUseCase
	updateUseCase  *usecase.UpdateMarketplaceAttributeUseCase
	deleteUseCase  *usecase.DeleteMarketplaceAttributeUseCase
}

// NewMarketplaceAttributeHandler crea una nueva instancia del manejador
func NewMarketplaceAttributeHandler(
	createUseCase *usecase.CreateMarketplaceAttributeUseCase,
	listUseCase *usecase.ListMarketplaceAttributesUseCase,
	getByIDUseCase *usecase.GetMarketplaceAttributeByIDUseCase,
	updateUseCase *usecase.UpdateMarketplaceAttributeUseCase,
	deleteUseCase *usecase.DeleteMarketplaceAttributeUseCase,
) *MarketplaceAttributeHandler {
	return &MarketplaceAttributeHandler{
		createUseCase:  createUseCase,
		listUseCase:    listUseCase,
		getByIDUseCase: getByIDUseCase,
		updateUseCase:  updateUseCase,
		deleteUseCase:  deleteUseCase,
	}
}

// RegisterRoutes registra las rutas del API para atributos marketplace
func (h *MarketplaceAttributeHandler) RegisterRoutes(router *gin.RouterGroup) {
	marketplace := router.Group("/marketplace")
	{
		marketplace.GET("/attributes", h.List)
		marketplace.POST("/attributes", h.Create)
		marketplace.GET("/attributes/:id", h.GetByID)
		marketplace.PUT("/attributes/:id", h.Update)
		marketplace.DELETE("/attributes/:id", h.Delete)
	}
}

// Create maneja la solicitud para crear un nuevo atributo marketplace
func (h *MarketplaceAttributeHandler) Create(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden crear atributos marketplace"})
		return
	}

	var req request.CreateMarketplaceAttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	// Convertir allowed values a validation rules
	validationRules := make(map[string]interface{})
	if len(req.AllowedValues) > 0 {
		validationRules["allowed_values"] = req.AllowedValues
	}

	attribute, err := h.createUseCase.Execute(
		c.Request.Context(),
		req.Name,
		"", // slug se genera automáticamente
		req.Type,
		req.IsFilterable,
		req.IsSearchable,
		req.IsRequired,
		validationRules,
		0, // sortOrder por defecto
	)
	if err != nil {
		if err == usecase.ErrInvalidAttributeName || err == usecase.ErrMarketplaceAttributeExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.FromMarketplaceEntity(attribute))
}

// List maneja la solicitud para listar atributos marketplace
func (h *MarketplaceAttributeHandler) List(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden ver atributos marketplace"})
		return
	}

	attributes, err := h.listUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.FromMarketplaceEntities(attributes))
}

// GetByID maneja la solicitud para obtener un atributo marketplace por ID
func (h *MarketplaceAttributeHandler) GetByID(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden ver atributos marketplace"})
		return
	}

	id := c.Param("id")

	attribute, err := h.getByIDUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		if err == usecase.ErrMarketplaceAttributeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.FromMarketplaceEntity(attribute))
}

// Update maneja la solicitud para actualizar un atributo marketplace
func (h *MarketplaceAttributeHandler) Update(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "super_admin" && userRole != "marketplace_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden actualizar atributos marketplace"})
		return
	}

	id := c.Param("id")

	var req request.UpdateMarketplaceAttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	// Convertir allowed values a validation rules
	validationRules := make(map[string]interface{})
	if len(req.AllowedValues) > 0 {
		validationRules["allowed_values"] = req.AllowedValues
	}

	attribute, err := h.updateUseCase.Execute(
		c.Request.Context(),
		id,
		req.Name,
		"", // slug se genera automáticamente si no se proporciona
		req.Type,
		req.IsFilterable,
		req.IsSearchable,
		req.IsRequired,
		validationRules,
		0, // sortOrder por defecto
	)
	if err != nil {
		if err == usecase.ErrMarketplaceAttributeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err == usecase.ErrInvalidAttributeName || err == usecase.ErrMarketplaceAttributeExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.FromMarketplaceEntity(attribute))
}

// Delete maneja la solicitud para eliminar un atributo marketplace
func (h *MarketplaceAttributeHandler) Delete(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "super_admin" && userRole != "marketplace_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden eliminar atributos marketplace"})
		return
	}

	id := c.Param("id")

	err := h.deleteUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		if err == usecase.ErrMarketplaceAttributeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
