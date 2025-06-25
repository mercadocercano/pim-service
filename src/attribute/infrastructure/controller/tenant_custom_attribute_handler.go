package controller

import (
	"net/http"

	"pim/src/attribute/application/request"
	"pim/src/attribute/application/usecase"

	"github.com/gin-gonic/gin"
)

// TenantCustomAttributeHandler maneja las peticiones HTTP para atributos custom de tenant
type TenantCustomAttributeHandler struct {
	extendTenantAttributesUseCase      *usecase.ExtendTenantAttributesUseCase
	getTenantCustomAttributesUseCase   *usecase.GetTenantCustomAttributesUseCase
	updateTenantCustomAttributeUseCase *usecase.UpdateTenantCustomAttributeUseCase
	deleteTenantCustomAttributeUseCase *usecase.DeleteTenantCustomAttributeUseCase
}

// NewTenantCustomAttributeHandler crea una nueva instancia del manejador
func NewTenantCustomAttributeHandler(
	extendTenantAttributesUseCase *usecase.ExtendTenantAttributesUseCase,
	getTenantCustomAttributesUseCase *usecase.GetTenantCustomAttributesUseCase,
	updateTenantCustomAttributeUseCase *usecase.UpdateTenantCustomAttributeUseCase,
	deleteTenantCustomAttributeUseCase *usecase.DeleteTenantCustomAttributeUseCase,
) *TenantCustomAttributeHandler {
	return &TenantCustomAttributeHandler{
		extendTenantAttributesUseCase:      extendTenantAttributesUseCase,
		getTenantCustomAttributesUseCase:   getTenantCustomAttributesUseCase,
		updateTenantCustomAttributeUseCase: updateTenantCustomAttributeUseCase,
		deleteTenantCustomAttributeUseCase: deleteTenantCustomAttributeUseCase,
	}
}

// RegisterRoutes registra las rutas del API para atributos custom de tenant
func (h *TenantCustomAttributeHandler) RegisterRoutes(router *gin.RouterGroup) {
	tenantAttributes := router.Group("/tenant/custom-attributes")
	{
		tenantAttributes.POST("", h.ExtendTenantAttributes)
		tenantAttributes.GET("", h.GetTenantCustomAttributes)
		tenantAttributes.PUT("/:attribute_id", h.UpdateTenantCustomAttribute)
		tenantAttributes.DELETE("/:attribute_id", h.DeleteTenantCustomAttribute)
	}
}

// ExtendTenantAttributes maneja la solicitud para extender atributos de un tenant
func (h *TenantCustomAttributeHandler) ExtendTenantAttributes(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	var req request.ExtendTenantAttributesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	extension, err := h.extendTenantAttributesUseCase.Execute(c.Request.Context(), &req, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, extension)
}

// GetTenantCustomAttributes maneja la solicitud para obtener atributos custom de un tenant
func (h *TenantCustomAttributeHandler) GetTenantCustomAttributes(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	// Crear request con parámetros de filtrado
	req := &request.GetTenantCustomAttributesRequest{
		TenantID:        tenantID,
		IncludeInactive: c.Query("include_inactive") == "true",
	}

	// Parámetros opcionales
	if categoryID := c.Query("marketplace_category_id"); categoryID != "" {
		req.MarketplaceCategoryID = &categoryID
	}

	if attrType := c.Query("attribute_type"); attrType != "" {
		req.AttributeType = &attrType
	}

	if filterable := c.Query("is_filterable"); filterable != "" {
		isFilterable := filterable == "true"
		req.IsFilterable = &isFilterable
	}

	if searchable := c.Query("is_searchable"); searchable != "" {
		isSearchable := searchable == "true"
		req.IsSearchable = &isSearchable
	}

	// Ejecutar caso de uso
	response, err := h.getTenantCustomAttributesUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateTenantCustomAttribute maneja la solicitud para actualizar un atributo custom
func (h *TenantCustomAttributeHandler) UpdateTenantCustomAttribute(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	attributeID := c.Param("attribute_id")
	if attributeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "attribute_id es obligatorio"})
		return
	}

	var req request.UpdateTenantCustomAttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	// Ejecutar caso de uso
	response, err := h.updateTenantCustomAttributeUseCase.Execute(c.Request.Context(), &req, tenantID, attributeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteTenantCustomAttribute maneja la solicitud para eliminar un atributo custom
func (h *TenantCustomAttributeHandler) DeleteTenantCustomAttribute(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	attributeID := c.Param("attribute_id")
	if attributeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "attribute_id es obligatorio"})
		return
	}

	// Ejecutar caso de uso
	err := h.deleteTenantCustomAttributeUseCase.Execute(c.Request.Context(), tenantID, attributeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
