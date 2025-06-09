package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pim/src/attribute/application/usecase"
	"pim/src/shared/domain/criteria"
	sharedCriteria "pim/src/shared/infrastructure/criteria"
)

// AttributeHandler maneja las peticiones HTTP para attributes
type AttributeHandler struct {
	// TODO: Agregar casos de uso cuando estén implementados
	criteriaHelper *sharedCriteria.EntityCriteriaHelper
}

// NewAttributeHandler crea una nueva instancia del manejador
func NewAttributeHandler() *AttributeHandler {
	return &AttributeHandler{
		criteriaHelper: sharedCriteria.NewEntityCriteriaHelper(),
	}
}

// RegisterRoutes registra las rutas del API
func (h *AttributeHandler) RegisterRoutes(router *gin.RouterGroup) {
	attributes := router.Group("/attributes")
	{
		attributes.POST("", h.Create)
		attributes.GET("", h.List)
		attributes.GET("/:id", h.GetByID)
		attributes.PUT("/:id", h.Update)
		attributes.DELETE("/:id", h.Delete)
	}
}

// Create maneja la solicitud para crear un nuevo attribute
func (h *AttributeHandler) Create(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	// TODO: Implementar binding de request y llamada al use case
	c.JSON(http.StatusNotImplemented, gin.H{"error": "no implementado - falta implementar casos de uso"})
}

// List maneja la solicitud para listar attributes
func (h *AttributeHandler) List(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	// TODO: Implementar listado con criterios
	c.JSON(http.StatusNotImplemented, gin.H{"error": "no implementado - falta implementar casos de uso"})
}

// GetByID maneja la solicitud para obtener un attribute por ID
func (h *AttributeHandler) GetByID(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	id := c.Param("id")
	// TODO: Implementar obtención por ID
	c.JSON(http.StatusNotImplemented, gin.H{"error": "no implementado - falta implementar casos de uso"})
}

// Update maneja la solicitud para actualizar un attribute
func (h *AttributeHandler) Update(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	id := c.Param("id")
	// TODO: Implementar actualización
	c.JSON(http.StatusNotImplemented, gin.H{"error": "no implementado - falta implementar casos de uso"})
}

// Delete maneja la solicitud para eliminar un attribute
func (h *AttributeHandler) Delete(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	id := c.Param("id")
	// TODO: Implementar eliminación
	c.JSON(http.StatusNotImplemented, gin.H{"error": "no implementado - falta implementar casos de uso"})
}
