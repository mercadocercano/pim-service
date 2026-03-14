package controller

import (
	"net/http"
	"saas-mt-pim-service/src/businesstype/application/usecase"
	"saas-mt-pim-service/src/businesstype/domain/port"
	cr "github.com/mercadocercano/criteria"

	"github.com/gin-gonic/gin"
)

// BusinessTypeHandler maneja las peticiones HTTP para business types
type BusinessTypeHandler struct {
	createUseCase *usecase.CreateBusinessTypeUseCase
	listUseCase   *usecase.ListBusinessTypesUseCase
	getUseCase    *usecase.GetBusinessTypeUseCase
	updateUseCase *usecase.UpdateBusinessTypeUseCase
	repository    port.BusinessTypeRepository
}

// NewBusinessTypeHandler crea una nueva instancia del handler
func NewBusinessTypeHandler(
	createUseCase *usecase.CreateBusinessTypeUseCase,
	listUseCase *usecase.ListBusinessTypesUseCase,
	getUseCase *usecase.GetBusinessTypeUseCase,
	updateUseCase *usecase.UpdateBusinessTypeUseCase,
	repository port.BusinessTypeRepository,
) *BusinessTypeHandler {
	return &BusinessTypeHandler{
		createUseCase: createUseCase,
		listUseCase:   listUseCase,
		getUseCase:    getUseCase,
		updateUseCase: updateUseCase,
		repository:    repository,
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
	// Construir criterios usando el builder correcto
	criteriaBuilder := cr.NewCriteriaBuilder()

	// Poblar desde query parameters (paginación, ordenamiento básico)
	criteriaBuilder.FromURLValues(c.Request.URL.Query())

	// Filtro only_active
	if c.Query("only_active") == "true" {
		criteriaBuilder.AddEqualFilter("is_active", true)
	}

	// Filtros adicionales que pueden venir del frontend
	if search := c.Query("search"); search != "" {
		criteriaBuilder.AddLikeFilter("name", search)
	}

	if code := c.Query("code"); code != "" {
		criteriaBuilder.AddLikeFilter("code", code)
	}

	if isActive := c.Query("is_active"); isActive != "" {
		if isActive == "true" {
			criteriaBuilder.AddEqualFilter("is_active", true)
		} else if isActive == "false" {
			criteriaBuilder.AddEqualFilter("is_active", false)
		}
	}

	// Establecer valores por defecto si no se especifican
	if c.Query("sort_by") == "" {
		criteriaBuilder.SetOrder("sort_order", cr.OrderAsc)
	}

	// Construir criterios finales
	searchCriteria := criteriaBuilder.Build()

	// Buscar usando criterios
	businessTypes, err := h.repository.SearchByCriteria(c.Request.Context(), searchCriteria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error listando business types: " + err.Error()})
		return
	}

	// Contar total usando criterios (sin paginación)
	countCriteria := cr.NewCriteria(
		searchCriteria.Filters,
		[]cr.Order{},    // Sin ordenamiento para conteo
		cr.Pagination{}, // Sin paginación para conteo
	)

	total, err := h.repository.CountByCriteria(c.Request.Context(), countCriteria)
	if err != nil {
		total = len(businessTypes) // Fallback
	}

	// Respuesta con formato compatible con frontend
	response := map[string]interface{}{
		"items":       businessTypes,
		"total_count": total,
		"page":        searchCriteria.Pagination.Page,
		"page_size":   searchCriteria.Pagination.Limit,
		"total_pages": (total + searchCriteria.Pagination.Limit - 1) / searchCriteria.Pagination.Limit,
	}

	c.JSON(http.StatusOK, response)
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
