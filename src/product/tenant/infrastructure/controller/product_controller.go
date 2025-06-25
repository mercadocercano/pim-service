package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"pim/src/product/tenant/application/request"
	"pim/src/product/tenant/application/usecase"
	"pim/src/product/tenant/infrastructure/criteria"
)

// ProductController maneja las peticiones HTTP para productos
type ProductController struct {
	createProductUseCase          *usecase.CreateProductUseCase
	getProductByIDUseCase         *usecase.GetProductByIDUseCase
	updateProductUseCase          *usecase.UpdateProductUseCase
	updateProductStatusUseCase    *usecase.UpdateProductStatusUseCase
	deleteProductUseCase          *usecase.DeleteProductUseCase
	listProductsByCriteriaUseCase *usecase.ListProductsByCriteriaUseCase
	criteriaBuilder               *criteria.ProductCriteriaBuilder
}

// NewProductController crea una nueva instancia del controller
func NewProductController(
	createProductUseCase *usecase.CreateProductUseCase,
	getProductByIDUseCase *usecase.GetProductByIDUseCase,
	updateProductUseCase *usecase.UpdateProductUseCase,
	updateProductStatusUseCase *usecase.UpdateProductStatusUseCase,
	deleteProductUseCase *usecase.DeleteProductUseCase,
	listProductsByCriteriaUseCase *usecase.ListProductsByCriteriaUseCase,
	criteriaBuilder *criteria.ProductCriteriaBuilder,
) *ProductController {
	return &ProductController{
		createProductUseCase:          createProductUseCase,
		getProductByIDUseCase:         getProductByIDUseCase,
		updateProductUseCase:          updateProductUseCase,
		updateProductStatusUseCase:    updateProductStatusUseCase,
		deleteProductUseCase:          deleteProductUseCase,
		listProductsByCriteriaUseCase: listProductsByCriteriaUseCase,
		criteriaBuilder:               criteriaBuilder,
	}
}

// CreateProduct godoc
// @Summary Crear un nuevo producto
// @Description Crea un nuevo producto en el sistema
// @Tags products
// @Accept json
// @Produce json
// @Param product body request.CreateProductRequest true "Datos del producto"
// @Success 201 {object} response.ProductResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products [post]
// @Security BearerAuth
func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var req request.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos", "details": err.Error()})
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Ejecutar caso de uso
	product, err := ctrl.createProductUseCase.Execute(c.Request.Context(), &req, tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProduct godoc
// @Summary Obtener un producto por ID
// @Description Obtiene un producto específico por su ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "ID del producto"
// @Success 200 {object} response.ProductResponse
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{id} [get]
// @Security BearerAuth
func (ctrl *ProductController) GetProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	if productIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto es requerido"})
		return
	}

	// Validar que el ID sea un UUID válido
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto debe ser un UUID válido"})
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Ejecutar caso de uso
	product, err := ctrl.getProductByIDUseCase.Execute(c.Request.Context(), productID, tenantID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct godoc
// @Summary Actualizar un producto
// @Description Actualiza un producto existente
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "ID del producto"
// @Param product body request.UpdateProductRequest true "Datos actualizados del producto"
// @Success 200 {object} response.ProductResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{id} [put]
// @Security BearerAuth
func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	if productIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto es requerido"})
		return
	}

	// Validar que el ID sea un UUID válido
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto debe ser un UUID válido"})
		return
	}

	var req request.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos", "details": err.Error()})
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Ejecutar caso de uso
	product, err := ctrl.updateProductUseCase.Execute(c.Request.Context(), productID, &req, tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Eliminar un producto
// @Description Elimina un producto (soft delete)
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "ID del producto"
// @Success 204
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{id} [delete]
// @Security BearerAuth
func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	if productIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto es requerido"})
		return
	}

	// Validar que el ID sea un UUID válido
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto debe ser un UUID válido"})
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Ejecutar caso de uso
	if err := ctrl.deleteProductUseCase.Execute(c.Request.Context(), productID, tenantID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListProducts godoc
// @Summary Listar productos con filtros
// @Description Lista productos con filtros, ordenamiento y paginación
// @Tags products
// @Accept json
// @Produce json
// @Param page query int false "Número de página" default(1)
// @Param page_size query int false "Tamaño de página" default(10)
// @Param sort_by query string false "Campo de ordenamiento" default("created_at")
// @Param sort_dir query string false "Dirección de ordenamiento" Enums(asc, desc) default("desc")
// @Param name query string false "Filtrar por nombre (búsqueda parcial)"
// @Param description query string false "Filtrar por descripción (búsqueda parcial)"
// @Param sku query string false "Filtrar por SKU (búsqueda parcial)"
// @Param status query string false "Filtrar por estado" Enums(active, inactive, discontinued)
// @Param category_id query string false "Filtrar por ID de categoría"
// @Param brand_id query string false "Filtrar por ID de marca"
// @Param category_name query string false "Filtrar por nombre de categoría (búsqueda parcial)"
// @Param brand_name query string false "Filtrar por nombre de marca (búsqueda parcial)"
// @Param include_deleted query bool false "Incluir productos eliminados" default(false)
// @Param in_stock query bool false "Filtrar por disponibilidad en stock"
// @Param min_price query number false "Precio mínimo"
// @Param max_price query number false "Precio máximo"
// @Success 200 {object} response.ProductListResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products [get]
// @Security BearerAuth
func (ctrl *ProductController) ListProducts(c *gin.Context) {
	// Obtener el tenantID del header y agregarlo a los query parameters
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Agregar tenant_id a los query parameters para el filtrado
	query := c.Request.URL.Query()
	query.Set("tenant_id", tenantID)
	c.Request.URL.RawQuery = query.Encode()

	// Construir y validar criterios
	criteria := ctrl.criteriaBuilder.BuildValidated(c)

	// Ejecutar el caso de uso
	result, err := ctrl.listProductsByCriteriaUseCase.Execute(c.Request.Context(), criteria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// UpdateProductStatus godoc
// @Summary Actualizar estado del producto
// @Description Actualiza el estado de un producto con validaciones de transición
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "ID del producto"
// @Param status body map[string]string true "Nuevo estado" example({"status": "active"})
// @Success 200 {object} usecase.UpdateProductStatusResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /products/{id}/status [patch]
// @Security BearerAuth
func (ctrl *ProductController) UpdateProductStatus(c *gin.Context) {
	productIDStr := c.Param("id")
	if productIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto es requerido"})
		return
	}

	// Validar que el ID sea un UUID válido
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto debe ser un UUID válido"})
		return
	}

	var statusReq map[string]string
	if err := c.ShouldBindJSON(&statusReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos", "details": err.Error()})
		return
	}

	newStatus, exists := statusReq["status"]
	if !exists || newStatus == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El campo 'status' es requerido"})
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Crear request para el caso de uso
	request := usecase.UpdateProductStatusRequest{
		ProductID: productID,
		TenantID:  tenantID,
		NewStatus: newStatus,
	}

	// Ejecutar caso de uso
	response, err := ctrl.updateProductStatusUseCase.Execute(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAvailableStatusTransitions godoc
// @Summary Obtener transiciones de estado disponibles
// @Description Retorna las transiciones de estado disponibles para un producto
// @Tags products
// @Produce json
// @Param id path string true "ID del producto"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /products/{id}/status/transitions [get]
// @Security BearerAuth
func (ctrl *ProductController) GetAvailableStatusTransitions(c *gin.Context) {
	productIDStr := c.Param("id")
	if productIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto es requerido"})
		return
	}

	// Validar que el ID sea un UUID válido
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto debe ser un UUID válido"})
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Obtener transiciones disponibles
	transitions, err := ctrl.updateProductStatusUseCase.GetAvailableTransitions(c.Request.Context(), productID, tenantID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product_id":            productID,
		"available_transitions": transitions,
	})
}

// RegisterRoutes registra las rutas del controller
func (ctrl *ProductController) RegisterRoutes(router *gin.RouterGroup) {
	products := router.Group("/products")
	{
		products.POST("", ctrl.CreateProduct)
		products.GET("", ctrl.ListProducts)
		products.GET("/:id", ctrl.GetProduct)
		products.PUT("/:id", ctrl.UpdateProduct)
		products.PATCH("/:id/status", ctrl.UpdateProductStatus)
		products.GET("/:id/status/transitions", ctrl.GetAvailableStatusTransitions)
		products.DELETE("/:id", ctrl.DeleteProduct)
	}
}
