package controller

import (
	"net/http"

	cr "github.com/hornosg/go-shared/criteria"
	"saas-mt-pim-service/src/attribute/application/request"
	"saas-mt-pim-service/src/attribute/application/response"
	"saas-mt-pim-service/src/attribute/application/usecase"

	"github.com/gin-gonic/gin"
)

// AttributeHandler maneja las peticiones HTTP para attributes
type AttributeHandler struct {
	createUseCase  *usecase.CreateAttributeUseCase
	listUseCase    *usecase.ListAttributesUseCase
	getByIDUseCase *usecase.GetAttributeByIDUseCase
	updateUseCase  *usecase.UpdateAttributeUseCase
	deleteUseCase  *usecase.DeleteAttributeUseCase
	criteriaHelper *cr.EntityCriteriaHelper
}

// NewAttributeHandler crea una nueva instancia del manejador
func NewAttributeHandler(
	createUseCase *usecase.CreateAttributeUseCase,
	listUseCase *usecase.ListAttributesUseCase,
	getByIDUseCase *usecase.GetAttributeByIDUseCase,
	updateUseCase *usecase.UpdateAttributeUseCase,
	deleteUseCase *usecase.DeleteAttributeUseCase,
) *AttributeHandler {
	return &AttributeHandler{
		createUseCase:  createUseCase,
		listUseCase:    listUseCase,
		getByIDUseCase: getByIDUseCase,
		updateUseCase:  updateUseCase,
		deleteUseCase:  deleteUseCase,
		criteriaHelper: cr.NewEntityCriteriaHelper(),
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

	var req request.CreateAttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	attribute, err := h.createUseCase.Execute(c.Request.Context(), tenantID, req.Name)
	if err != nil {
		if err == usecase.ErrInvalidAttributeName || err == usecase.ErrAttributeExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.FromEntity(attribute))
}

// List maneja la solicitud para listar attributes
func (h *AttributeHandler) List(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	attributes, err := h.listUseCase.Execute(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.FromEntities(attributes))
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

	attribute, err := h.getByIDUseCase.Execute(c.Request.Context(), id, tenantID)
	if err != nil {
		if err == usecase.ErrAttributeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.FromEntity(attribute))
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

	var req request.UpdateAttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	attribute, err := h.updateUseCase.Execute(c.Request.Context(), id, tenantID, req.Name)
	if err != nil {
		if err == usecase.ErrAttributeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err == usecase.ErrInvalidAttributeName || err == usecase.ErrAttributeExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.FromEntity(attribute))
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

	err := h.deleteUseCase.Execute(c.Request.Context(), id, tenantID)
	if err != nil {
		if err == usecase.ErrAttributeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
