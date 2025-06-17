package controller

import (
	"net/http"
	"pim/src/businesstype/application/usecase"

	"github.com/gin-gonic/gin"
)

// BusinessTypeHandler maneja las peticiones HTTP para business types
type BusinessTypeHandler struct {
	createUseCase *usecase.CreateBusinessTypeUseCase
	listUseCase   *usecase.ListBusinessTypesUseCase
	getUseCase    *usecase.GetBusinessTypeUseCase
	updateUseCase *usecase.UpdateBusinessTypeUseCase
}

// NewBusinessTypeHandler crea una nueva instancia del handler
func NewBusinessTypeHandler(
	createUseCase *usecase.CreateBusinessTypeUseCase,
	listUseCase *usecase.ListBusinessTypesUseCase,
	getUseCase *usecase.GetBusinessTypeUseCase,
	updateUseCase *usecase.UpdateBusinessTypeUseCase,
) *BusinessTypeHandler {
	return &BusinessTypeHandler{
		createUseCase: createUseCase,
		listUseCase:   listUseCase,
		getUseCase:    getUseCase,
		updateUseCase: updateUseCase,
	}
}

// RegisterRoutes registra las rutas del módulo
func (h *BusinessTypeHandler) RegisterRoutes(router *gin.RouterGroup) {
	businessTypes := router.Group("/business-types")
	{
		businessTypes.POST("", h.CreateBusinessType)
		businessTypes.GET("", h.ListBusinessTypes)
		businessTypes.GET("/:id", h.GetBusinessType)
		businessTypes.PUT("/:id", h.UpdateBusinessType)
		businessTypes.DELETE("/:id", h.DeleteBusinessType)
	}
}

// CreateBusinessType maneja la creación de un business type
func (h *BusinessTypeHandler) CreateBusinessType(c *gin.Context) {
	var req usecase.CreateBusinessTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Validar rol admin
	role := c.GetHeader("X-User-Role")
	if role != "marketplace_admin" && role != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden crear business types"})
		return
	}

	businessType, err := h.createUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando business type: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, businessType)
}

// ListBusinessTypes maneja el listado de business types
func (h *BusinessTypeHandler) ListBusinessTypes(c *gin.Context) {
	var req usecase.ListBusinessTypesRequest

	// Parsear query params
	if c.Query("only_active") == "true" {
		req.OnlyActive = true
	}

	result, err := h.listUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error listando business types: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetBusinessType obtiene un business type por ID
func (h *BusinessTypeHandler) GetBusinessType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID es requerido"})
		return
	}

	businessType, err := h.getUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "business type no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo business type: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, businessType)
}

// UpdateBusinessType actualiza un business type
func (h *BusinessTypeHandler) UpdateBusinessType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID es requerido"})
		return
	}

	// Validar rol admin
	role := c.GetHeader("X-User-Role")
	if role != "marketplace_admin" && role != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden actualizar business types"})
		return
	}

	var req usecase.UpdateBusinessTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	businessType, err := h.updateUseCase.Execute(c.Request.Context(), id, req)
	if err != nil {
		if err.Error() == "business type no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error actualizando business type: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, businessType)
}

// DeleteBusinessType elimina un business type
func (h *BusinessTypeHandler) DeleteBusinessType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID es requerido"})
		return
	}

	// Validar rol admin
	role := c.GetHeader("X-User-Role")
	if role != "marketplace_admin" && role != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden eliminar business types"})
		return
	}

	// TODO: Implementar caso de uso DeleteBusinessType
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Endpoint no implementado"})
}
