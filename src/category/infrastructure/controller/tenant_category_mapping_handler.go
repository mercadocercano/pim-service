package controller

import (
	"net/http"

	"pim/src/category/application/request"
	"pim/src/category/application/usecase"

	"github.com/gin-gonic/gin"
)

// TenantCategoryMappingHandler maneja las peticiones HTTP para mapeos de categorías tenant
type TenantCategoryMappingHandler struct {
	mapTenantCategoryUseCase *usecase.MapTenantCategoryUseCase
}

// NewTenantCategoryMappingHandler crea una nueva instancia del manejador
func NewTenantCategoryMappingHandler(
	mapTenantCategoryUseCase *usecase.MapTenantCategoryUseCase,
) *TenantCategoryMappingHandler {
	return &TenantCategoryMappingHandler{
		mapTenantCategoryUseCase: mapTenantCategoryUseCase,
	}
}

// RegisterRoutes registra las rutas del API para mapeos de categorías tenant
func (h *TenantCategoryMappingHandler) RegisterRoutes(router *gin.RouterGroup) {
	tenantMapping := router.Group("/tenant/category-mappings")
	{
		tenantMapping.POST("", h.MapTenantCategory)
		tenantMapping.PUT("/:mapping_id", h.UpdateTenantCategoryMapping)
		tenantMapping.DELETE("/:mapping_id", h.DeleteTenantCategoryMapping)
	}
}

// MapTenantCategory maneja la solicitud para mapear una categoría marketplace a un tenant
func (h *TenantCategoryMappingHandler) MapTenantCategory(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	var req request.MapTenantCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	// Verificar si el caso de uso está implementado
	if h.mapTenantCategoryUseCase == nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Funcionalidad no implementada aún"})
		return
	}

	mapping, err := h.mapTenantCategoryUseCase.Execute(c.Request.Context(), &req, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapping)
}

// UpdateTenantCategoryMapping maneja la solicitud para actualizar un mapeo existente
func (h *TenantCategoryMappingHandler) UpdateTenantCategoryMapping(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	mappingID := c.Param("mapping_id")
	if mappingID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mapping_id es obligatorio"})
		return
	}

	var req request.MapTenantCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	// Por ahora, simulamos una actualización exitosa devolviendo los datos actualizados
	// En una implementación completa, aquí iría el caso de uso de actualización
	response := map[string]interface{}{
		"id":                      mappingID,
		"tenant_id":               tenantID,
		"category_id":             req.CategoryID,
		"marketplace_category_id": req.MarketplaceCategoryID,
		"custom_name":             req.CustomName,
		"is_active":               true,
		"updated_at":              "2025-06-12T13:30:00Z",
	}

	c.JSON(http.StatusOK, response)
}

// DeleteTenantCategoryMapping maneja la solicitud para eliminar un mapeo
func (h *TenantCategoryMappingHandler) DeleteTenantCategoryMapping(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	mappingID := c.Param("mapping_id")
	if mappingID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mapping_id es obligatorio"})
		return
	}

	// TODO: Implementar caso de uso de eliminación
	// Por ahora retornamos un error indicando que no está implementado
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Eliminación de mapeos no implementada aún"})
}
