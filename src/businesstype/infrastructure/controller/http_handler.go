package controller

import (
	"net/http"

	cr "github.com/mercadocercano/criteria"
	"github.com/gin-gonic/gin"

	"saas-mt-pim-service/src/businesstype/application/usecase"
	"saas-mt-pim-service/src/businesstype/domain/port"
)

// BusinessTypeHandler maneja las peticiones HTTP para business types
type BusinessTypeHandler struct {
	createUseCase     *usecase.CreateBusinessTypeUseCase
	listUseCase       *usecase.ListBusinessTypesUseCase
	getUseCase        *usecase.GetBusinessTypeUseCase
	updateUseCase     *usecase.UpdateBusinessTypeUseCase
	deleteUseCase     *usecase.DeleteBusinessTypeUseCase
	activateUseCase   *usecase.ActivateBusinessTypeUseCase
	deactivateUseCase *usecase.DeactivateBusinessTypeUseCase
	repository        port.BusinessTypeRepository
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
		createUseCase:     createUseCase,
		listUseCase:       listUseCase,
		getUseCase:        getUseCase,
		updateUseCase:     updateUseCase,
		deleteUseCase:     usecase.NewDeleteBusinessTypeUseCase(repository),
		activateUseCase:   usecase.NewActivateBusinessTypeUseCase(repository),
		deactivateUseCase: usecase.NewDeactivateBusinessTypeUseCase(repository),
		repository:        repository,
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
		businessTypes.PATCH("/:id/activate", h.ActivateBusinessType)
		businessTypes.PATCH("/:id/deactivate", h.DeactivateBusinessType)
	}
}

// CreateBusinessType maneja la creación de un business type
func (h *BusinessTypeHandler) CreateBusinessType(c *gin.Context) {
	var req usecase.CreateBusinessTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden crear business types"})
		return
	}

	businessType, err := h.createUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Error creando business type: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, businessType)
}

// ListBusinessTypes maneja el listado de business types
func (h *BusinessTypeHandler) ListBusinessTypes(c *gin.Context) {
	criteriaBuilder := cr.NewCriteriaBuilder()
	criteriaBuilder.FromURLValues(c.Request.URL.Query())

	if c.Query("only_active") == "true" {
		criteriaBuilder.AddEqualFilter("is_active", true)
	}

	if search := c.Query("search"); search != "" {
		criteriaBuilder.AddLikeFilter("name", search)
	}

	if isActive := c.Query("is_active"); isActive != "" {
		if isActive == "true" {
			criteriaBuilder.AddEqualFilter("is_active", true)
		} else if isActive == "false" {
			criteriaBuilder.AddEqualFilter("is_active", false)
		}
	}

	if c.Query("sort_by") == "" {
		criteriaBuilder.SetOrder("sort_order", cr.OrderAsc)
	}

	searchCriteria := criteriaBuilder.Build()

	businessTypes, err := h.repository.SearchByCriteria(c.Request.Context(), searchCriteria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error listando business types: " + err.Error()})
		return
	}

	countCriteria := cr.NewCriteria(searchCriteria.Filters, []cr.Order{}, cr.Pagination{})
	total, err := h.repository.CountByCriteria(c.Request.Context(), countCriteria)
	if err != nil {
		total = len(businessTypes)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"items":       businessTypes,
		"total_count": total,
		"page":        searchCriteria.Pagination.Page,
		"page_size":   searchCriteria.Pagination.Limit,
		"total_pages": calcTotalPages(total, searchCriteria.Pagination.Limit),
	})
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

	if !isAdmin(c) {
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

	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden eliminar business types"})
		return
	}

	if err := h.deleteUseCase.Execute(c.Request.Context(), id); err != nil {
		if err.Error() == "business type no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error eliminando business type: " + err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// ActivateBusinessType activa un business type
func (h *BusinessTypeHandler) ActivateBusinessType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID es requerido"})
		return
	}

	bt, err := h.activateUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "business type no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error activando business type: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, bt)
}

// DeactivateBusinessType desactiva un business type
func (h *BusinessTypeHandler) DeactivateBusinessType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID es requerido"})
		return
	}

	bt, err := h.deactivateUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "business type no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error desactivando business type: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, bt)
}

// isAdmin verifica que el header X-User-Role sea admin
func isAdmin(c *gin.Context) bool {
	role := c.GetHeader("X-User-Role")
	return role == "marketplace_admin" || role == "super_admin"
}

// calcTotalPages calcula el total de páginas evitando división por cero
func calcTotalPages(total, pageSize int) int {
	if pageSize <= 0 {
		return 1
	}
	return (total + pageSize - 1) / pageSize
}
