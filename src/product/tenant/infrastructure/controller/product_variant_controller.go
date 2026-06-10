package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"saas-mt-pim-service/src/product/tenant/application/request"
	"saas-mt-pim-service/src/product/tenant/application/usecase"
)

// ProductVariantController maneja las peticiones HTTP para variantes de productos
type ProductVariantController struct {
	createProductVariantUseCase          *usecase.CreateProductVariantUseCase
	getProductVariantByIDUseCase         *usecase.GetProductVariantByIDUseCase
	getVariantBySKUUseCase               *usecase.GetVariantBySKUUseCase
	updateProductVariantUseCase          *usecase.UpdateProductVariantUseCase
	deleteProductVariantUseCase          *usecase.DeleteProductVariantUseCase
	listProductVariantsByCriteriaUseCase *usecase.ListProductVariantsByCriteriaUseCase
	getVariantsBySKUsUseCase             *usecase.GetVariantsBySKUsUseCase
}

// NewProductVariantController crea una nueva instancia del controller
func NewProductVariantController(
	createProductVariantUseCase *usecase.CreateProductVariantUseCase,
	getProductVariantByIDUseCase *usecase.GetProductVariantByIDUseCase,
	getVariantBySKUUseCase *usecase.GetVariantBySKUUseCase,
	updateProductVariantUseCase *usecase.UpdateProductVariantUseCase,
	deleteProductVariantUseCase *usecase.DeleteProductVariantUseCase,
	listProductVariantsByCriteriaUseCase *usecase.ListProductVariantsByCriteriaUseCase,
	getVariantsBySKUsUseCase *usecase.GetVariantsBySKUsUseCase,
) *ProductVariantController {
	return &ProductVariantController{
		createProductVariantUseCase:          createProductVariantUseCase,
		getProductVariantByIDUseCase:         getProductVariantByIDUseCase,
		getVariantBySKUUseCase:               getVariantBySKUUseCase,
		updateProductVariantUseCase:          updateProductVariantUseCase,
		deleteProductVariantUseCase:          deleteProductVariantUseCase,
		listProductVariantsByCriteriaUseCase: listProductVariantsByCriteriaUseCase,
		getVariantsBySKUsUseCase:             getVariantsBySKUsUseCase,
	}
}

// CreateProductVariant crea una nueva variante de producto
func (ctrl *ProductVariantController) CreateProductVariant(c *gin.Context) {
	var req request.CreateProductVariantRequest
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
	variant, err := ctrl.createProductVariantUseCase.Execute(c.Request.Context(), &req, tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, variant)
}

// GetProductVariant godoc
// @Summary Obtener una variante de producto por ID
// @Description Obtiene una variante específica de un producto por su ID
// @Tags product-variants
// @Accept json
// @Produce json
// @Param product_id path string true "ID del producto"
// @Param variant_id path string true "ID de la variante"
// @Success 200 {object} response.ProductVariantResponse
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{product_id}/variants/{variant_id} [get]
// @Security BearerAuth
func (ctrl *ProductVariantController) GetProductVariant(c *gin.Context) {
	productIDStr := c.Param("id") // Cambiado de product_id a id
	variantIDStr := c.Param("variant_id")

	if productIDStr == "" || variantIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto y de la variante son requeridos"})
		return
	}

	// Validar que los IDs sean UUIDs válidos
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto debe ser un UUID válido"})
		return
	}

	variantID, err := uuid.Parse(variantIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de la variante debe ser un UUID válido"})
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Ejecutar caso de uso
	variant, err := ctrl.getProductVariantByIDUseCase.Execute(c.Request.Context(), productID, variantID, tenantID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, variant)
}

// UpdateProductVariant godoc
// @Summary Actualizar una variante de producto
// @Description Actualiza una variante existente de un producto
// @Tags product-variants
// @Accept json
// @Produce json
// @Param product_id path string true "ID del producto"
// @Param variant_id path string true "ID de la variante"
// @Param variant body request.UpdateProductVariantRequest true "Datos actualizados de la variante"
// @Success 200 {object} response.ProductVariantResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{product_id}/variants/{variant_id} [put]
// @Security BearerAuth
func (ctrl *ProductVariantController) UpdateProductVariant(c *gin.Context) {
	productIDStr := c.Param("id") // Cambiado de product_id a id
	variantIDStr := c.Param("variant_id")

	if productIDStr == "" || variantIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto y de la variante son requeridos"})
		return
	}

	// Validar que los IDs sean UUIDs válidos
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto debe ser un UUID válido"})
		return
	}

	variantID, err := uuid.Parse(variantIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de la variante debe ser un UUID válido"})
		return
	}

	var req request.UpdateProductVariantRequest
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
	variant, err := ctrl.updateProductVariantUseCase.Execute(c.Request.Context(), productID, variantID, &req, tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, variant)
}

// DeleteProductVariant godoc
// @Summary Eliminar una variante de producto
// @Description Elimina una variante de producto (soft delete)
// @Tags product-variants
// @Accept json
// @Produce json
// @Param product_id path string true "ID del producto"
// @Param variant_id path string true "ID de la variante"
// @Success 204
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{product_id}/variants/{variant_id} [delete]
// @Security BearerAuth
func (ctrl *ProductVariantController) DeleteProductVariant(c *gin.Context) {
	productIDStr := c.Param("id") // Cambiado de product_id a id
	variantIDStr := c.Param("variant_id")

	if productIDStr == "" || variantIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto y de la variante son requeridos"})
		return
	}

	// Validar que los IDs sean UUIDs válidos
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto debe ser un UUID válido"})
		return
	}

	variantID, err := uuid.Parse(variantIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de la variante debe ser un UUID válido"})
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Ejecutar caso de uso
	if err := ctrl.deleteProductVariantUseCase.Execute(c.Request.Context(), productID, variantID, tenantID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListProductVariants godoc
// @Summary Listar variantes de un producto
// @Description Lista todas las variantes de un producto específico con filtros y paginación
// @Tags product-variants
// @Accept json
// @Produce json
// @Param product_id path string true "ID del producto"
// @Param page query int false "Número de página" default(1)
// @Param page_size query int false "Tamaño de página" default(10)
// @Param name query string false "Filtrar por nombre de variante"
// @Param sku query string false "Filtrar por SKU de variante"
// @Param status query string false "Filtrar por estado" Enums(active, inactive, discontinued)
// @Param is_default query bool false "Filtrar por variante por defecto"
// @Success 200 {object} response.ProductVariantListResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{product_id}/variants [get]
// @Security BearerAuth
func (ctrl *ProductVariantController) ListProductVariants(c *gin.Context) {
	productIDStr := c.Param("id") // Usar :id para consistencia con /products/:id
	if productIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto es requerido"})
		return
	}

	// Validar que el ID del producto sea un UUID válido
	_, err := uuid.Parse(productIDStr)
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

	// Agregar product_id a los parámetros de consulta para el filtro
	queryParams := c.Request.URL.Query()
	queryParams.Set("product_id", productIDStr)
	c.Request.URL.RawQuery = queryParams.Encode()

	// Ejecutar caso de uso
	variants, err := ctrl.listProductVariantsByCriteriaUseCase.Execute(c.Request.Context(), queryParams, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, variants)
}

// ListAllVariants godoc
// @Summary Listar todas las variantes
// @Description Lista todas las variantes de productos con filtros y paginación
// @Tags product-variants
// @Accept json
// @Produce json
// @Param page query int false "Número de página" default(1)
// @Param page_size query int false "Tamaño de página" default(10)
// @Param product_id query string false "Filtrar por ID de producto"
// @Param name query string false "Filtrar por nombre de variante"
// @Param sku query string false "Filtrar por SKU de variante"
// @Param status query string false "Filtrar por estado" Enums(active, inactive, discontinued)
// @Param is_default query bool false "Filtrar por variante por defecto"
// @Success 200 {object} response.ProductVariantListResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /variants [get]
// @Security BearerAuth
func (ctrl *ProductVariantController) ListAllVariants(c *gin.Context) {
	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Ejecutar caso de uso
	variants, err := ctrl.listProductVariantsByCriteriaUseCase.Execute(c.Request.Context(), c.Request.URL.Query(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, variants)
}

// RegisterRoutes registra las rutas del controlador
func (ctrl *ProductVariantController) RegisterRoutes(router *gin.RouterGroup) {
	// Rutas específicas para variantes de productos (standalone)
	router.POST("/product-variants", ctrl.CreateProductVariant)
	router.GET("/product-variants", ctrl.ListAllVariants)
	router.GET("/product-variants/:variant_id", ctrl.GetProductVariantByID)
	router.PUT("/product-variants/:variant_id", ctrl.UpdateProductVariantByID)
	router.DELETE("/product-variants/:variant_id", ctrl.DeleteProductVariantByID)

	// HITO A - Endpoint para buscar variante por SKU
	router.GET("/variants/by-sku/:sku", ctrl.GetVariantBySKU)

	// Batch lookup de variantes por SKUs (para orquestación BFF)
	router.POST("/variants/by-skus", ctrl.GetVariantsBySKUs)

	// Rutas anidadas bajo productos (usando :id en lugar de :product_id para evitar conflicto con /products/:id)
	// Estas se registran en el setup del módulo Product, DESPUÉS de las rutas de ProductController
}

// GetProductVariantByID obtiene una variante por su ID directamente
func (ctrl *ProductVariantController) GetProductVariantByID(c *gin.Context) {
	variantIDStr := c.Param("variant_id")
	if variantIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de la variante es requerido"})
		return
	}

	// Validar que el ID sea un UUID válido
	_, err := uuid.Parse(variantIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de la variante debe ser un UUID válido"})
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Para obtener por ID directo, necesitamos buscar en todas las variantes
	// Esto es una simplificación - en producción podrías querer un método más eficiente
	queryParams := c.Request.URL.Query()
	queryParams.Set("variant_id", variantIDStr)
	c.Request.URL.RawQuery = queryParams.Encode()

	variants, err := ctrl.listProductVariantsByCriteriaUseCase.Execute(c.Request.Context(), queryParams, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(variants.Variants) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Variante no encontrada"})
		return
	}

	c.JSON(http.StatusOK, variants.Variants[0])
}

// UpdateProductVariantByID actualiza una variante por su ID directamente
func (ctrl *ProductVariantController) UpdateProductVariantByID(c *gin.Context) {
	variantIDStr := c.Param("variant_id")
	if variantIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de la variante es requerido"})
		return
	}

	// Validar que el ID sea un UUID válido
	variantID, err := uuid.Parse(variantIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de la variante debe ser un UUID válido"})
		return
	}

	var req request.UpdateProductVariantRequest
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

	// Necesitamos obtener el product_id de la variante primero
	// Por simplicidad, vamos a requerir que se pase como query parameter
	productIDStr := c.Query("product_id")
	if productIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product_id query parameter es requerido"})
		return
	}

	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product_id debe ser un UUID válido"})
		return
	}

	// Ejecutar caso de uso
	variant, err := ctrl.updateProductVariantUseCase.Execute(c.Request.Context(), productID, variantID, &req, tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, variant)
}

// DeleteProductVariantByID elimina una variante por su ID directamente
func (ctrl *ProductVariantController) DeleteProductVariantByID(c *gin.Context) {
	variantIDStr := c.Param("variant_id")
	if variantIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de la variante es requerido"})
		return
	}

	// Validar que el ID sea un UUID válido
	variantID, err := uuid.Parse(variantIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de la variante debe ser un UUID válido"})
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Necesitamos obtener el product_id de la variante primero
	productIDStr := c.Query("product_id")
	if productIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product_id query parameter es requerido"})
		return
	}

	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product_id debe ser un UUID válido"})
		return
	}

	// Ejecutar caso de uso
	if err := ctrl.deleteProductVariantUseCase.Execute(c.Request.Context(), productID, variantID, tenantID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetVariantBySKU busca una variante por su SKU
// HITO A - Endpoint requerido por order-service para obtener snapshots
func (ctrl *ProductVariantController) GetVariantBySKU(c *gin.Context) {
	sku := c.Param("sku")
	if sku == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SKU es requerido"})
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Ejecutar caso de uso
	variant, err := ctrl.getVariantBySKUUseCase.Execute(c.Request.Context(), sku, tenantID)
	if err != nil {
		// Si no se encuentra, retornar 404
		c.JSON(http.StatusNotFound, gin.H{"error": "variant not found: " + sku})
		return
	}

	c.JSON(http.StatusOK, variant)
}

// GetVariantsBySKUs batch lookup of variants enriched with product and category info
func (ctrl *ProductVariantController) GetVariantsBySKUs(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	var req usecase.GetVariantsBySKUsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos", "details": err.Error()})
		return
	}

	result, err := ctrl.getVariantsBySKUsUseCase.Execute(c.Request.Context(), &req, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
