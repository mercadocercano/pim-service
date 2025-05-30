package controller

import (
	"net/http"

	"pim/src/category_attribute/application/request"
	"pim/src/category_attribute/application/response"
	"pim/src/category_attribute/application/usecase"
	categoryAttributeCriteria "pim/src/category_attribute/infrastructure/criteria"

	"github.com/gin-gonic/gin"
)

// CategoryAttributeHandler maneja las peticiones HTTP para atributos de categoría
type CategoryAttributeHandler struct {
	createUseCase                           *usecase.CreateCategoryAttributeUseCase
	updateUseCase                           *usecase.UpdateCategoryAttributeUseCase
	deleteUseCase                           *usecase.DeleteCategoryAttributeUseCase
	getUseCase                              *usecase.GetCategoryAttributesUseCase
	listCategoryAttributesByCriteriaUseCase *usecase.ListCategoryAttributesByCriteriaUseCase
	criteriaBuilder                         *categoryAttributeCriteria.CategoryAttributeCriteriaBuilder
}

// NewCategoryAttributeHandler crea una nueva instancia del manejador de atributos de categoría
func NewCategoryAttributeHandler(
	createUseCase *usecase.CreateCategoryAttributeUseCase,
	updateUseCase *usecase.UpdateCategoryAttributeUseCase,
	deleteUseCase *usecase.DeleteCategoryAttributeUseCase,
	getUseCase *usecase.GetCategoryAttributesUseCase,
	listCategoryAttributesByCriteriaUseCase *usecase.ListCategoryAttributesByCriteriaUseCase,
) *CategoryAttributeHandler {
	return &CategoryAttributeHandler{
		createUseCase:                           createUseCase,
		updateUseCase:                           updateUseCase,
		deleteUseCase:                           deleteUseCase,
		getUseCase:                              getUseCase,
		listCategoryAttributesByCriteriaUseCase: listCategoryAttributesByCriteriaUseCase,
		criteriaBuilder:                         categoryAttributeCriteria.NewCategoryAttributeCriteriaBuilder(),
	}
}

// RegisterRoutes registra las rutas del API para atributos de categoría
func (h *CategoryAttributeHandler) RegisterRoutes(router *gin.RouterGroup) {
	categoryAttributes := router.Group("/category-attributes")
	{
		categoryAttributes.GET("", h.ListWithCriteria)
		categoryAttributes.GET("/simple", h.List)
		categoryAttributes.POST("", h.Create)
		categoryAttributes.PUT("/:id", h.Update)
		categoryAttributes.DELETE("/:id", h.Delete)
	}
}

// ListWithCriteria maneja la solicitud para listar atributos de categoría con filtros y paginación
func (h *CategoryAttributeHandler) ListWithCriteria(c *gin.Context) {
	// Obtener el tenantID del header y agregarlo a los query parameters
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		tenantID = c.Query("tenant_id")
	}

	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el tenant_id es obligatorio"})
		return
	}

	// Agregar tenant_id a los query parameters para el filtrado
	query := c.Request.URL.Query()
	query.Set("tenant_id", tenantID)
	c.Request.URL.RawQuery = query.Encode()

	// Construir y validar criterios
	criteria := h.criteriaBuilder.BuildValidated(c)

	// Ejecutar el caso de uso
	result, err := h.listCategoryAttributesByCriteriaUseCase.Execute(c.Request.Context(), criteria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convertir las entidades a respuestas
	var responseItems []*response.CategoryAttributeResponse
	for _, categoryAttribute := range result.Items {
		responseItems = append(responseItems, response.FromEntity(categoryAttribute))
	}

	// Crear respuesta con paginación
	paginatedResponse := struct {
		Items      []*response.CategoryAttributeResponse `json:"items"`
		TotalCount int                                   `json:"total_count"`
		Page       int                                   `json:"page"`
		PageSize   int                                   `json:"page_size"`
		TotalPages int                                   `json:"total_pages"`
	}{
		Items:      responseItems,
		TotalCount: result.TotalCount,
		Page:       result.Page,
		PageSize:   result.PageSize,
		TotalPages: result.TotalPages,
	}

	c.JSON(http.StatusOK, paginatedResponse)
}

// List maneja la solicitud para obtener atributos de categoría (método simple)
func (h *CategoryAttributeHandler) List(c *gin.Context) {
	// Obtener el tenantID del header o query param
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		tenantID = c.Query("tenant_id")
	}

	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el tenant_id es obligatorio"})
		return
	}

	categoryID := c.Query("category_id")

	categoryAttributes, err := h.getUseCase.Execute(c.Request.Context(), tenantID, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response.FromEntityList(categoryAttributes),
	})
}

// Create maneja la solicitud para crear un nuevo atributo de categoría
// POST /api/v1/category-attributes
func (h *CategoryAttributeHandler) Create(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	var req request.CreateCategoryAttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryAttribute, err := h.createUseCase.Execute(
		c.Request.Context(),
		tenantID,
		req.CategoryID,
		req.AttributeID,
		req.AllowedValues,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.FromEntity(categoryAttribute))
}

// Update maneja la solicitud para actualizar un atributo de categoría
// PUT /api/v1/category-attributes/{id}
func (h *CategoryAttributeHandler) Update(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	id := c.Param("id")

	var req request.UpdateCategoryAttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.updateUseCase.Execute(c.Request.Context(), id, tenantID, req.AllowedValues)
	if err != nil {
		if err.Error() == "atributo de categoría no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Atributo de categoría actualizado exitosamente"})
}

// Delete maneja la solicitud para eliminar un atributo de categoría
// DELETE /api/v1/category-attributes/{id}
func (h *CategoryAttributeHandler) Delete(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	id := c.Param("id")

	err := h.deleteUseCase.Execute(c.Request.Context(), id, tenantID)
	if err != nil {
		if err.Error() == "atributo de categoría no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Atributo de categoría eliminado exitosamente"})
}
