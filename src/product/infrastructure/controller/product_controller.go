package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"pim/src/product/application/request"
	"pim/src/product/application/usecase"
	"pim/src/product/infrastructure/criteria"
	domainCriteria "pim/src/shared/domain/criteria"
)

// ProductController maneja las peticiones HTTP para productos
type ProductController struct {
	createProductUseCase          *usecase.CreateProductUseCase
	getProductByIDUseCase         *usecase.GetProductByIDUseCase
	updateProductUseCase          *usecase.UpdateProductUseCase
	deleteProductUseCase          *usecase.DeleteProductUseCase
	listProductsByCriteriaUseCase *usecase.ListProductsByCriteriaUseCase
	criteriaBuilder               *criteria.ProductCriteriaBuilder
}

// NewProductController crea una nueva instancia del controller
func NewProductController(
	createProductUseCase *usecase.CreateProductUseCase,
	getProductByIDUseCase *usecase.GetProductByIDUseCase,
	updateProductUseCase *usecase.UpdateProductUseCase,
	deleteProductUseCase *usecase.DeleteProductUseCase,
	listProductsByCriteriaUseCase *usecase.ListProductsByCriteriaUseCase,
	criteriaBuilder *criteria.ProductCriteriaBuilder,
) *ProductController {
	return &ProductController{
		createProductUseCase:          createProductUseCase,
		getProductByIDUseCase:         getProductByIDUseCase,
		updateProductUseCase:          updateProductUseCase,
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
// @Success 200 {object} response.ProductListResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products [get]
// @Security BearerAuth
func (ctrl *ProductController) ListProducts(c *gin.Context) {
	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Construir criterios de búsqueda
	searchCriteria := ctrl.criteriaBuilder.BuildValidated(c)

	// Agregar filtro de tenant automáticamente
	tenantFilter := domainCriteria.Filter{
		Field:    "tenant_id",
		Operator: "=",
		Value:    tenantID,
	}
	searchCriteria.Filters.Add(tenantFilter)

	// Ejecutar caso de uso
	result, err := ctrl.listProductsByCriteriaUseCase.Execute(c.Request.Context(), searchCriteria)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// RegisterRoutes registra las rutas del controller
func (ctrl *ProductController) RegisterRoutes(router *gin.RouterGroup) {
	products := router.Group("/products")
	{
		products.POST("", ctrl.CreateProduct)
		products.GET("", ctrl.ListProducts)
		products.GET("/:id", ctrl.GetProduct)
		products.PUT("/:id", ctrl.UpdateProduct)
		products.DELETE("/:id", ctrl.DeleteProduct)
	}
}
