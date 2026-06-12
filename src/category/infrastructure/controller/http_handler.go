package controller

import (
	"bytes"
	"fmt"
	httpresp "github.com/hornosg/go-shared/infrastructure/response"
	"io"
	"net/http"

	"saas-mt-pim-service/src/category/application/request"
	"saas-mt-pim-service/src/category/application/response"
	"saas-mt-pim-service/src/category/application/usecase"
	categoryCriteria "saas-mt-pim-service/src/category/infrastructure/criteria"

	"github.com/gin-gonic/gin"
)

// CategoryHandler maneja las peticiones HTTP para categorías
type CategoryHandler struct {
	createUseCase                   *usecase.CreateCategoryUseCase
	updateUseCase                   *usecase.UpdateCategoryUseCase
	changeCategoryStatus            *usecase.ChangeCategoryStatusUseCase
	moveCategoryUseCase             *usecase.MoveCategoryUseCase
	getCategoriesUseCase            *usecase.GetCategoriesUseCase
	getCategoryByIDUseCase          *usecase.GetCategoryByIDUseCase
	listCategoriesByCriteriaUseCase *usecase.ListCategoriesByCriteriaUseCase
	criteriaBuilder                 *categoryCriteria.CategoryCriteriaBuilder
}

// NewCategoryHandler crea una nueva instancia del manejador de categorías
func NewCategoryHandler(
	createUseCase *usecase.CreateCategoryUseCase,
	updateUseCase *usecase.UpdateCategoryUseCase,
	changeCategoryStatus *usecase.ChangeCategoryStatusUseCase,
	moveCategoryUseCase *usecase.MoveCategoryUseCase,
	getCategoriesUseCase *usecase.GetCategoriesUseCase,
	getCategoryByIDUseCase *usecase.GetCategoryByIDUseCase,
	listCategoriesByCriteriaUseCase *usecase.ListCategoriesByCriteriaUseCase,
) *CategoryHandler {
	return &CategoryHandler{
		createUseCase:                   createUseCase,
		updateUseCase:                   updateUseCase,
		changeCategoryStatus:            changeCategoryStatus,
		moveCategoryUseCase:             moveCategoryUseCase,
		getCategoriesUseCase:            getCategoriesUseCase,
		getCategoryByIDUseCase:          getCategoryByIDUseCase,
		listCategoriesByCriteriaUseCase: listCategoriesByCriteriaUseCase,
		criteriaBuilder:                 categoryCriteria.NewCategoryCriteriaBuilder(),
	}
}

// RegisterRoutes registra las rutas del API para categorías
func (h *CategoryHandler) RegisterRoutes(router *gin.RouterGroup) {
	categories := router.Group("/categories")
	{
		categories.POST("", h.Create)
		categories.GET("", h.ListWithCriteria)
		categories.GET("/simple", h.List)
		categories.GET("/tree", h.ListTree)
		categories.GET("/:id", h.GetByID)
		categories.PUT("/:id", h.Update)
		categories.PATCH("/:id/activate", h.Activate)
		categories.PATCH("/:id/deactivate", h.Deactivate)
		categories.PATCH("/:id/move", h.Move)
	}
}

// Create maneja la solicitud para crear una nueva categoría
func (h *CategoryHandler) Create(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "el header X-Tenant-ID es obligatorio")
		return
	}

	// Leer el cuerpo de la petición y depurar
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "Error leyendo el cuerpo de la petición: "+err.Error())
		return
	}

	// Restaurar el cuerpo para que pueda ser leído de nuevo durante el binding
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	// Imprimir el cuerpo para depuración
	fmt.Printf("Cuerpo recibido: %s\n", string(body))

	var req request.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "Error en el binding JSON: "+err.Error())
		return
	}

	fmt.Printf("Después del binding: name=%s, description=%s, parentID=%v\n",
		req.Name, req.Description, req.ParentID)

	category, err := h.createUseCase.Execute(c.Request.Context(), tenantID, req.Name, req.Description, req.ParentID)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, response.FromEntity(category))
}

// Update maneja la solicitud para actualizar una categoría
func (h *CategoryHandler) Update(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "el header X-Tenant-ID es obligatorio")
		return
	}

	id := c.Param("id")

	var req request.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	// Convertir el ParentID de string a *string
	var parentIDPtr *string
	if req.ParentID != "" {
		parentIDPtr = &req.ParentID
	}

	category, err := h.updateUseCase.Execute(c.Request.Context(), id, tenantID, req.Name, req.Description, parentIDPtr)
	if err == usecase.ErrCategoryNotFound {
		httpresp.JSON(c, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.FromEntity(category))
}

// Move maneja la solicitud para mover una categoría a un nuevo padre
func (h *CategoryHandler) Move(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "el header X-Tenant-ID es obligatorio")
		return
	}

	id := c.Param("id")

	var req request.MoveCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	// Convertir el ParentID de string a *string
	var parentIDPtr *string
	if req.ParentID != "" {
		parentIDPtr = &req.ParentID
	}

	category, err := h.moveCategoryUseCase.Execute(c.Request.Context(), id, tenantID, parentIDPtr)
	if err == usecase.ErrCategoryNotFound {
		httpresp.JSON(c, http.StatusNotFound, err.Error())
		return
	}
	if err == usecase.ErrInvalidMove {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.FromEntity(category))
}

// Activate maneja la solicitud para activar una categoría
func (h *CategoryHandler) Activate(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "el header X-Tenant-ID es obligatorio")
		return
	}

	id := c.Param("id")

	category, err := h.changeCategoryStatus.Activate(c.Request.Context(), id, tenantID)
	if err == usecase.ErrCategoryNotFound {
		httpresp.JSON(c, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.FromEntity(category))
}

// Deactivate maneja la solicitud para desactivar una categoría
func (h *CategoryHandler) Deactivate(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "el header X-Tenant-ID es obligatorio")
		return
	}

	id := c.Param("id")

	category, err := h.changeCategoryStatus.Deactivate(c.Request.Context(), id, tenantID)
	if err == usecase.ErrCategoryNotFound {
		httpresp.JSON(c, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.FromEntity(category))
}

// List maneja la solicitud para listar todas las categorías
func (h *CategoryHandler) List(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "el header X-Tenant-ID es obligatorio")
		return
	}

	categories, err := h.getCategoriesUseCase.Execute(c.Request.Context(), tenantID)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Convertir las entidades de categoría a respuestas
	var responseList []*response.CategoryResponse
	for _, category := range categories {
		responseList = append(responseList, response.FromEntity(category))
	}

	c.JSON(http.StatusOK, responseList)
}

// ListTree maneja la solicitud para listar todas las categorías en formato de árbol
func (h *CategoryHandler) ListTree(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "el header X-Tenant-ID es obligatorio")
		return
	}

	// Obtener el ID de categoría opcional de la query
	categoryID := c.Query("category_id")

	categories, err := h.getCategoriesUseCase.Execute(c.Request.Context(), tenantID)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Construir el árbol de categorías
	treeResponse := response.BuildCategoryTree(categories)

	// Si se especificó un ID de categoría, filtrar para devolver solo esa categoría y sus hijos
	if categoryID != "" {
		// Buscar la categoría específica en el árbol
		for _, category := range treeResponse {
			if category.ID == categoryID {
				c.JSON(http.StatusOK, []*response.CategoryTreeResponse{category})
				return
			}
			// Buscar en los hijos
			if found := findCategoryInTree(category.Children, categoryID); found != nil {
				c.JSON(http.StatusOK, []*response.CategoryTreeResponse{found})
				return
			}
		}
		httpresp.JSON(c, http.StatusNotFound, "categoría no encontrada")
		return
	}

	c.JSON(http.StatusOK, treeResponse)
}

// findCategoryInTree busca una categoría por ID en el árbol
func findCategoryInTree(categories []*response.CategoryTreeResponse, categoryID string) *response.CategoryTreeResponse {
	for _, category := range categories {
		if category.ID == categoryID {
			return category
		}
		if found := findCategoryInTree(category.Children, categoryID); found != nil {
			return found
		}
	}
	return nil
}

// GetByID maneja la solicitud para obtener una categoría por su ID
func (h *CategoryHandler) GetByID(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "el header X-Tenant-ID es obligatorio")
		return
	}

	id := c.Param("id")

	category, err := h.getCategoryByIDUseCase.Execute(c.Request.Context(), id, tenantID)
	if err == usecase.ErrCategoryNotFound {
		httpresp.JSON(c, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.FromEntity(category))
}

// ListWithCriteria maneja la solicitud para listar categorías con filtros y paginación
func (h *CategoryHandler) ListWithCriteria(c *gin.Context) {
	// Obtener el tenantID del header y agregarlo a los query parameters
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "el header X-Tenant-ID es obligatorio")
		return
	}

	// Agregar tenant_id a los query parameters para el filtrado
	query := c.Request.URL.Query()
	query.Set("tenant_id", tenantID)
	c.Request.URL.RawQuery = query.Encode()

	// Construir y validar criterios
	criteria := h.criteriaBuilder.BuildValidated(c)

	// Ejecutar el caso de uso
	result, err := h.listCategoriesByCriteriaUseCase.Execute(c.Request.Context(), criteria)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Convertir las entidades a respuestas
	var responseItems []*response.CategoryResponse
	for _, category := range result.Items {
		responseItems = append(responseItems, response.FromEntity(category))
	}

	// Crear respuesta con paginación
	paginatedResponse := struct {
		Items      []*response.CategoryResponse `json:"items"`
		TotalCount int                          `json:"total_count"`
		Page       int                          `json:"page"`
		PageSize   int                          `json:"page_size"`
		TotalPages int                          `json:"total_pages"`
	}{
		Items:      responseItems,
		TotalCount: result.TotalCount,
		Page:       result.Page,
		PageSize:   result.PageSize,
		TotalPages: result.TotalPages,
	}

	c.JSON(http.StatusOK, paginatedResponse)
}
