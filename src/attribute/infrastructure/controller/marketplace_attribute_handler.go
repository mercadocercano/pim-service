package controller

import (
	"net/http"

	"saas-mt-pim-service/src/attribute/application/request"
	"saas-mt-pim-service/src/attribute/application/response"
	"saas-mt-pim-service/src/attribute/application/usecase"
	"saas-mt-pim-service/src/attribute/infrastructure/criteria"

	"github.com/gin-gonic/gin"
)

// MarketplaceAttributeHandler maneja las peticiones HTTP para atributos marketplace
type MarketplaceAttributeHandler struct {
	createUseCase               *usecase.CreateMarketplaceAttributeUseCase
	listUseCase                 *usecase.ListMarketplaceAttributesUseCase
	listByCriteriaUseCase       *usecase.ListMarketplaceAttributesByCriteriaUseCase
	getByIDUseCase              *usecase.GetMarketplaceAttributeByIDUseCase
	updateUseCase               *usecase.UpdateMarketplaceAttributeUseCase
	deleteUseCase               *usecase.DeleteMarketplaceAttributeUseCase
	criteriaBuilder             *criteria.MarketplaceAttributeCriteriaBuilder
	listValuesUseCase           *usecase.ListAttributeValuesUseCase
	createValueUseCase          *usecase.CreateAttributeValueUseCase
	updateValueUseCase          *usecase.UpdateAttributeValueUseCase
	deleteValueUseCase          *usecase.DeleteAttributeValueUseCase
}

// NewMarketplaceAttributeHandler crea una nueva instancia del manejador (solo atributos)
func NewMarketplaceAttributeHandler(
	createUseCase *usecase.CreateMarketplaceAttributeUseCase,
	listUseCase *usecase.ListMarketplaceAttributesUseCase,
	listByCriteriaUseCase *usecase.ListMarketplaceAttributesByCriteriaUseCase,
	getByIDUseCase *usecase.GetMarketplaceAttributeByIDUseCase,
	updateUseCase *usecase.UpdateMarketplaceAttributeUseCase,
	deleteUseCase *usecase.DeleteMarketplaceAttributeUseCase,
	criteriaBuilder *criteria.MarketplaceAttributeCriteriaBuilder,
) *MarketplaceAttributeHandler {
	return &MarketplaceAttributeHandler{
		createUseCase:         createUseCase,
		listUseCase:           listUseCase,
		listByCriteriaUseCase: listByCriteriaUseCase,
		getByIDUseCase:        getByIDUseCase,
		updateUseCase:         updateUseCase,
		deleteUseCase:         deleteUseCase,
		criteriaBuilder:       criteriaBuilder,
	}
}

// WithValueUseCases inyecta los usecases de values en el handler
func (h *MarketplaceAttributeHandler) WithValueUseCases(
	list *usecase.ListAttributeValuesUseCase,
	create *usecase.CreateAttributeValueUseCase,
	update *usecase.UpdateAttributeValueUseCase,
	del *usecase.DeleteAttributeValueUseCase,
) *MarketplaceAttributeHandler {
	h.listValuesUseCase = list
	h.createValueUseCase = create
	h.updateValueUseCase = update
	h.deleteValueUseCase = del
	return h
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

		if h.listValuesUseCase != nil {
			marketplace.GET("/attributes/:id/values", h.ListValues)
			marketplace.POST("/attributes/:id/values", h.CreateValue)
			marketplace.PUT("/attributes/:id/values/:vid", h.UpdateValue)
			marketplace.DELETE("/attributes/:id/values/:vid", h.DeleteValue)
		}
	}
}

// isAdminRole verifica si el rol tiene permisos de administrador marketplace
func isAdminRole(role string) bool {
	return role == "marketplace_admin" || role == "super_admin"
}

// Create maneja la solicitud para crear un nuevo atributo marketplace
func (h *MarketplaceAttributeHandler) Create(c *gin.Context) {
	if !isAdminRole(c.GetHeader("X-User-Role")) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden crear atributos marketplace"})
		return
	}

	var req request.CreateMarketplaceAttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	validationRules := make(map[string]interface{})
	if len(req.AllowedValues) > 0 {
		validationRules["allowed_values"] = req.AllowedValues
	}

	attribute, err := h.createUseCase.Execute(
		c.Request.Context(),
		req.Name,
		"",
		req.Type,
		req.IsFilterable,
		req.IsSearchable,
		req.IsRequired,
		validationRules,
		0,
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
	if !isAdminRole(c.GetHeader("X-User-Role")) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden ver atributos marketplace"})
		return
	}

	searchCriteria := h.criteriaBuilder.BuildValidated(c)
	result, err := h.listByCriteriaUseCase.Execute(c.Request.Context(), searchCriteria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetByID maneja la solicitud para obtener un atributo marketplace por ID
func (h *MarketplaceAttributeHandler) GetByID(c *gin.Context) {
	if !isAdminRole(c.GetHeader("X-User-Role")) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden ver atributos marketplace"})
		return
	}

	attribute, err := h.getByIDUseCase.Execute(c.Request.Context(), c.Param("id"))
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
	if !isAdminRole(c.GetHeader("X-User-Role")) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden actualizar atributos marketplace"})
		return
	}

	var req request.UpdateMarketplaceAttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	validationRules := make(map[string]interface{})
	if len(req.AllowedValues) > 0 {
		validationRules["allowed_values"] = req.AllowedValues
	}

	attribute, err := h.updateUseCase.Execute(
		c.Request.Context(),
		c.Param("id"),
		req.Name,
		"", // slug se genera automáticamente desde el nombre
		req.Type,
		req.IsFilterable,
		req.IsSearchable,
		req.IsRequired,
		validationRules,
		0,
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
	if !isAdminRole(c.GetHeader("X-User-Role")) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden eliminar atributos marketplace"})
		return
	}

	err := h.deleteUseCase.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		if err == usecase.ErrMarketplaceAttributeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err == usecase.ErrAttributeInUse {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// ListValues lista los valores de un atributo
func (h *MarketplaceAttributeHandler) ListValues(c *gin.Context) {
	if !isAdminRole(c.GetHeader("X-User-Role")) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden ver valores de atributos"})
		return
	}

	values, err := h.listValuesUseCase.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": response.FromAttributeValueEntities(values)})
}

// CreateValue crea un nuevo valor para un atributo
func (h *MarketplaceAttributeHandler) CreateValue(c *gin.Context) {
	if !isAdminRole(c.GetHeader("X-User-Role")) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden crear valores de atributos"})
		return
	}

	var req request.CreateAttributeValueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	v, err := h.createValueUseCase.Execute(c.Request.Context(), c.Param("id"), req.Value, req.SortOrder)
	if err != nil {
		if err == usecase.ErrMarketplaceAttributeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.FromAttributeValueEntity(v))
}

// UpdateValue actualiza un valor existente de un atributo
func (h *MarketplaceAttributeHandler) UpdateValue(c *gin.Context) {
	if !isAdminRole(c.GetHeader("X-User-Role")) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden actualizar valores de atributos"})
		return
	}

	var req request.UpdateAttributeValueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	v, err := h.updateValueUseCase.Execute(c.Request.Context(), c.Param("id"), c.Param("vid"), req.Value, req.SortOrder)
	if err != nil {
		if err == usecase.ErrAttributeValueNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.FromAttributeValueEntity(v))
}

// DeleteValue elimina un valor de un atributo
func (h *MarketplaceAttributeHandler) DeleteValue(c *gin.Context) {
	if !isAdminRole(c.GetHeader("X-User-Role")) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden eliminar valores de atributos"})
		return
	}

	err := h.deleteValueUseCase.Execute(c.Request.Context(), c.Param("id"), c.Param("vid"))
	if err != nil {
		if err == usecase.ErrAttributeValueNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
