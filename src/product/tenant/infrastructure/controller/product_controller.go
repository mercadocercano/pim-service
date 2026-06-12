package controller

import (
	httpresp "github.com/hornosg/go-shared/infrastructure/response"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	sharedport "github.com/hornosg/go-shared/domain/port"
	"saas-mt-pim-service/src/product/tenant/application/request"
	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/application/usecase"
	tenantport "saas-mt-pim-service/src/product/tenant/domain/port"
	"saas-mt-pim-service/src/product/tenant/infrastructure/criteria"
)

// ProductController maneja las peticiones HTTP para productos
type ProductController struct {
	createProductUseCase          *usecase.CreateProductUseCase
	getProductByIDUseCase         *usecase.GetProductByIDUseCase
	updateProductUseCase          *usecase.UpdateProductUseCase
	updateProductStatusUseCase    *usecase.UpdateProductStatusUseCase
	deleteProductUseCase          *usecase.DeleteProductUseCase
	listProductsByCriteriaUseCase *usecase.ListProductsByCriteriaUseCase
	importProductsFromCSVUseCase  *usecase.ImportProductsFromCSVUseCase
	importProductsAsyncUseCase    *usecase.ImportProductsAsyncUseCase
	validateSKUsUseCase           *usecase.ValidateSKUsUseCase
	criteriaBuilder               *criteria.ProductCriteriaBuilder
	metrics                       sharedport.MetricsRecorder
}

// NewProductController crea una nueva instancia del controller
func NewProductController(
	createProductUseCase *usecase.CreateProductUseCase,
	getProductByIDUseCase *usecase.GetProductByIDUseCase,
	updateProductUseCase *usecase.UpdateProductUseCase,
	updateProductStatusUseCase *usecase.UpdateProductStatusUseCase,
	deleteProductUseCase *usecase.DeleteProductUseCase,
	listProductsByCriteriaUseCase *usecase.ListProductsByCriteriaUseCase,
	importProductsFromCSVUseCase *usecase.ImportProductsFromCSVUseCase,
	importProductsAsyncUseCase *usecase.ImportProductsAsyncUseCase,
	validateSKUsUseCase *usecase.ValidateSKUsUseCase,
	criteriaBuilder *criteria.ProductCriteriaBuilder,
	metricsRecorder sharedport.MetricsRecorder,
) *ProductController {
	return &ProductController{
		createProductUseCase:          createProductUseCase,
		getProductByIDUseCase:         getProductByIDUseCase,
		updateProductUseCase:          updateProductUseCase,
		updateProductStatusUseCase:    updateProductStatusUseCase,
		deleteProductUseCase:          deleteProductUseCase,
		listProductsByCriteriaUseCase: listProductsByCriteriaUseCase,
		importProductsFromCSVUseCase:  importProductsFromCSVUseCase,
		importProductsAsyncUseCase:    importProductsAsyncUseCase,
		validateSKUsUseCase:           validateSKUsUseCase,
		criteriaBuilder:               criteriaBuilder,
		metrics:                       metricsRecorder,
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
		httpresp.JSONWithDetails(c, http.StatusBadRequest, "Datos de entrada inválidos", err.Error())
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}

	// Ejecutar caso de uso
	product, err := ctrl.createProductUseCase.Execute(c.Request.Context(), &req, tenantID)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
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
		httpresp.JSON(c, http.StatusBadRequest, "ID del producto es requerido")
		return
	}

	// Validar que el ID sea un UUID válido
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "ID del producto debe ser un UUID válido")
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}

	// Ejecutar caso de uso
	product, err := ctrl.getProductByIDUseCase.Execute(c.Request.Context(), productID, tenantID)
	if err != nil {
		httpresp.JSON(c, http.StatusNotFound, err.Error())
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
		httpresp.JSON(c, http.StatusBadRequest, "ID del producto es requerido")
		return
	}

	// Validar que el ID sea un UUID válido
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "ID del producto debe ser un UUID válido")
		return
	}

	var req request.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresp.JSONWithDetails(c, http.StatusBadRequest, "Datos de entrada inválidos", err.Error())
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}

	// Ejecutar caso de uso
	product, err := ctrl.updateProductUseCase.Execute(c.Request.Context(), productID, &req, tenantID)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
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
		httpresp.JSON(c, http.StatusBadRequest, "ID del producto es requerido")
		return
	}

	// Validar que el ID sea un UUID válido
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "ID del producto debe ser un UUID válido")
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}

	// Ejecutar caso de uso
	if err := ctrl.deleteProductUseCase.Execute(c.Request.Context(), productID, tenantID); err != nil {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
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
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
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
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
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
		httpresp.JSON(c, http.StatusBadRequest, "ID del producto es requerido")
		return
	}

	// Validar que el ID sea un UUID válido
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "ID del producto debe ser un UUID válido")
		return
	}

	var statusReq map[string]string
	if err := c.ShouldBindJSON(&statusReq); err != nil {
		httpresp.JSONWithDetails(c, http.StatusBadRequest, "Datos de entrada inválidos", err.Error())
		return
	}

	newStatus, exists := statusReq["status"]
	if !exists || newStatus == "" {
		httpresp.JSON(c, http.StatusBadRequest, "El campo 'status' es requerido")
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
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
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
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
		httpresp.JSON(c, http.StatusBadRequest, "ID del producto es requerido")
		return
	}

	// Validar que el ID sea un UUID válido
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "ID del producto debe ser un UUID válido")
		return
	}

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}

	// Obtener transiciones disponibles
	transitions, err := ctrl.updateProductStatusUseCase.GetAvailableTransitions(c.Request.Context(), productID, tenantID)
	if err != nil {
		httpresp.JSON(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product_id":            productID,
		"available_transitions": transitions,
	})
}

// ImportProductsCSV godoc
// @Summary Importar productos desde archivo CSV
// @Description Importa múltiples productos desde un archivo CSV
// @Tags products
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Archivo CSV con productos"
// @Param create_variants formData bool false "Crear variantes por defecto"
// @Success 200 {object} response.ImportProductsCSVResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/import-csv [post]
// @Security BearerAuth
func (ctrl *ProductController) ImportProductsCSV(c *gin.Context) {
	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}

	// Obtener archivo del form
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		httpresp.JSONWithDetails(c, http.StatusBadRequest, "No se pudo obtener el archivo", err.Error())
		return
	}
	defer file.Close()

	// Validar extensión del archivo
	if !strings.HasSuffix(strings.ToLower(header.Filename), ".csv") {
		httpresp.JSON(c, http.StatusBadRequest, "El archivo debe ser CSV")
		return
	}

	ctrl.metrics.Record(sharedport.MetricEvent{
		Name:  tenantport.MetricImportFileSize,
		Kind:  sharedport.MetricKindHistogram,
		Unit:  sharedport.MetricUnitBytes,
		Value: float64(header.Size),
		Labels: map[string]string{
			"tenant_id": tenantID,
			"type":      "csv_products",
		},
	})

	// Ejecutar caso de uso
	result, err := ctrl.importProductsFromCSVUseCase.Execute(c.Request.Context(), file, tenantID)
	if err != nil {
		httpresp.JSONWithDetails(c, http.StatusInternalServerError, "Error al procesar archivo", err.Error())
		return
	}

	// Crear respuesta
	response := response.NewImportProductsCSVResponse(
		result.ImportResult,
		result.SavedProducts,
		result.ProcessingErrors,
	)

	c.JSON(http.StatusOK, response)
}

// ValidateSKUs godoc
// @Summary Validar SKUs existentes
// @Description Verifica qué SKUs ya existen en el sistema del tenant
// @Tags products
// @Accept json
// @Produce json
// @Param body body request.ValidateSKUsRequest true "Lista de SKUs a validar"
// @Success 200 {object} response.ValidateSKUsResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/validate-skus [post]
// @Security BearerAuth
func (ctrl *ProductController) ValidateSKUs(c *gin.Context) {
	startTime := time.Now()

	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}

	// Parsear request
	var req request.ValidateSKUsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresp.JSONWithDetails(c, http.StatusBadRequest, "Datos de entrada inválidos", err.Error())
		return
	}

	// Ejecutar caso de uso
	result, err := ctrl.validateSKUsUseCase.Execute(c.Request.Context(), &req, tenantID)
	if err != nil {
		httpresp.JSONWithDetails(c, http.StatusInternalServerError, "Error al validar SKUs", err.Error())
		return
	}

	duration := time.Since(startTime).Seconds()
	ctrl.metrics.Record(sharedport.MetricEvent{
		Name:   tenantport.MetricSKUValidation,
		Kind:   sharedport.MetricKindHistogram,
		Unit:   sharedport.MetricUnitSeconds,
		Value:  duration,
		Labels: map[string]string{"tenant_id": tenantID},
	})
	ctrl.metrics.Record(sharedport.MetricEvent{
		Name:   tenantport.MetricSKUBatchSize,
		Kind:   sharedport.MetricKindHistogram,
		Value:  float64(len(req.SKUs)),
		Labels: map[string]string{"tenant_id": tenantID},
	})

	c.JSON(http.StatusOK, result)
}

// ImportProductsCSVAsync godoc
// @Summary Importar productos desde archivo CSV de forma asíncrona
// @Description Inicia una importación asíncrona de productos desde un archivo CSV
// @Tags products
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Archivo CSV con productos"
// @Param webhook_url formData string false "URL para notificación webhook"
// @Param email_notify formData string false "Email para notificación"
// @Success 202 {object} usecase.ImportAsyncResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/import-csv/async [post]
// @Security BearerAuth
func (ctrl *ProductController) ImportProductsCSVAsync(c *gin.Context) {
	if ctrl.importProductsAsyncUseCase == nil {
		httpresp.JSON(c, http.StatusNotImplemented, "async import not available")
		return
	}
	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}

	// Obtener user ID del contexto o header
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		userID = "system"
	}

	// Obtener archivo del form
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		httpresp.JSONWithDetails(c, http.StatusBadRequest, "No se pudo obtener el archivo", err.Error())
		return
	}
	defer file.Close()

	// Validar extensión del archivo
	if !strings.HasSuffix(strings.ToLower(header.Filename), ".csv") {
		httpresp.JSON(c, http.StatusBadRequest, "El archivo debe ser CSV")
		return
	}

	// Obtener parámetros opcionales
	var webhookURL *string
	if webhook := c.PostForm("webhook_url"); webhook != "" {
		webhookURL = &webhook
	}

	var emailNotify *string
	if email := c.PostForm("email_notify"); email != "" {
		emailNotify = &email
	}

	// Crear request
	req := &usecase.ImportAsyncRequest{
		FileName:    header.Filename,
		FileReader:  file,
		FileSize:    header.Size,
		TenantID:    tenantID,
		UserID:      userID,
		WebhookURL:  webhookURL,
		EmailNotify: emailNotify,
	}

	// Ejecutar caso de uso
	result, err := ctrl.importProductsAsyncUseCase.StartImport(c.Request.Context(), req)
	if err != nil {
		httpresp.JSONWithDetails(c, http.StatusInternalServerError, "Error al iniciar importación", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, result)
}

// GetImportJobStatus godoc
// @Summary Obtener estado de trabajo de importación
// @Description Obtiene el estado actual de un trabajo de importación
// @Tags import-jobs
// @Accept json
// @Produce json
// @Param id path string true "ID del trabajo"
// @Success 200 {object} entity.ImportJob
// @Failure 404 {object} map[string]interface{}
// @Router /import-jobs/{id} [get]
// @Security BearerAuth
func (ctrl *ProductController) GetImportJobStatus(c *gin.Context) {
	jobIDStr := c.Param("id")
	if jobIDStr == "" {
		httpresp.JSON(c, http.StatusBadRequest, "ID del trabajo es requerido")
		return
	}

	jobID, err := uuid.Parse(jobIDStr)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "ID del trabajo debe ser un UUID válido")
		return
	}

	job, err := ctrl.importProductsAsyncUseCase.GetImportJobStatus(c.Request.Context(), jobID)
	if err != nil {
		httpresp.JSON(c, http.StatusNotFound, "Trabajo no encontrado")
		return
	}

	c.JSON(http.StatusOK, job)
}

// CancelImportJob godoc
// @Summary Cancelar trabajo de importación
// @Description Cancela un trabajo de importación en progreso
// @Tags import-jobs
// @Accept json
// @Produce json
// @Param id path string true "ID del trabajo"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /import-jobs/{id}/cancel [post]
// @Security BearerAuth
func (ctrl *ProductController) CancelImportJob(c *gin.Context) {
	jobIDStr := c.Param("id")
	if jobIDStr == "" {
		httpresp.JSON(c, http.StatusBadRequest, "ID del trabajo es requerido")
		return
	}

	jobID, err := uuid.Parse(jobIDStr)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "ID del trabajo debe ser un UUID válido")
		return
	}

	if err := ctrl.importProductsAsyncUseCase.CancelImportJob(c.Request.Context(), jobID); err != nil {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Trabajo cancelado exitosamente"})
}

// RegisterRoutes registra las rutas del controller
func (ctrl *ProductController) RegisterRoutes(router *gin.RouterGroup) {
	products := router.Group("/products")
	{
		// Rutas sin parámetros primero
		products.POST("", ctrl.CreateProduct)
		products.GET("", ctrl.ListProducts)
		products.POST("/import-csv", ctrl.ImportProductsCSV)
		products.POST("/import-csv/async", ctrl.ImportProductsCSVAsync)
		products.POST("/validate-skus", ctrl.ValidateSKUs)

		// Luego rutas con /:id
		products.GET("/:id", ctrl.GetProduct)
		products.PUT("/:id", ctrl.UpdateProduct)
		products.PATCH("/:id/status", ctrl.UpdateProductStatus)
		products.GET("/:id/status/transitions", ctrl.GetAvailableStatusTransitions)
		products.DELETE("/:id", ctrl.DeleteProduct)
	}

	// Rutas de trabajos de importación
	importJobs := router.Group("/import-jobs")
	{
		importJobs.GET("/:id", ctrl.GetImportJobStatus)
		importJobs.POST("/:id/cancel", ctrl.CancelImportJob)
	}
}
