package controller

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"saas-mt-pim-service/src/product/global_catalog/application/usecase"
	"saas-mt-pim-service/src/product/global_catalog/domain/exception"
	"saas-mt-pim-service/src/product/global_catalog/infrastructure/controller/dto"
	"saas-mt-pim-service/src/product/global_catalog/infrastructure/criteria"
)

// GlobalCatalogController maneja las requests HTTP para el catálogo global
type GlobalCatalogController struct {
	createGlobalProduct           *usecase.CreateGlobalProduct
	searchByEAN                   *usecase.SearchByEAN
	listGlobalProducts            *usecase.ListGlobalProducts
	listGlobalProductsByCriteria  *usecase.ListGlobalProductsByCriteriaUseCase
	getGlobalProductByID          *usecase.GetGlobalProductByID
	updateGlobalProductByID       *usecase.UpdateGlobalProductByID
	deleteGlobalProduct           *usecase.DeleteGlobalProduct
	verifyGlobalProduct           *usecase.VerifyGlobalProduct
	unverifyGlobalProduct         *usecase.UnverifyGlobalProduct
	bulkImportGlobalProducts      *usecase.BulkImportGlobalProducts
	getBusinessTypeFacets         *usecase.GetBusinessTypeFacets
	listProductsNeedingEnrichment *usecase.ListProductsNeedingEnrichment
	getGlobalProductsByIDs        *usecase.GetGlobalProductsByIDs
	getDistinctBusinessTypes      *usecase.GetDistinctBusinessTypes
	criteriaBuilder               *criteria.GlobalProductCriteriaBuilder
}

// GlobalCatalogControllerDeps agrupa todas las dependencias del controlador
type GlobalCatalogControllerDeps struct {
	CreateGlobalProduct           *usecase.CreateGlobalProduct
	SearchByEAN                   *usecase.SearchByEAN
	ListGlobalProducts            *usecase.ListGlobalProducts
	ListGlobalProductsByCriteria  *usecase.ListGlobalProductsByCriteriaUseCase
	GetGlobalProductByID          *usecase.GetGlobalProductByID
	UpdateGlobalProductByID       *usecase.UpdateGlobalProductByID
	DeleteGlobalProduct           *usecase.DeleteGlobalProduct
	VerifyGlobalProduct           *usecase.VerifyGlobalProduct
	UnverifyGlobalProduct         *usecase.UnverifyGlobalProduct
	BulkImportGlobalProducts      *usecase.BulkImportGlobalProducts
	GetBusinessTypeFacets         *usecase.GetBusinessTypeFacets
	ListProductsNeedingEnrichment *usecase.ListProductsNeedingEnrichment
	GetGlobalProductsByIDs        *usecase.GetGlobalProductsByIDs
	GetDistinctBusinessTypes      *usecase.GetDistinctBusinessTypes
	CriteriaBuilder               *criteria.GlobalProductCriteriaBuilder
}

// NewGlobalCatalogController crea una nueva instancia del controlador
func NewGlobalCatalogController(
	createGlobalProduct *usecase.CreateGlobalProduct,
	searchByEAN *usecase.SearchByEAN,
	listGlobalProducts *usecase.ListGlobalProducts,
	listGlobalProductsByCriteria *usecase.ListGlobalProductsByCriteriaUseCase,
	getGlobalProductByID *usecase.GetGlobalProductByID,
	updateGlobalProductByID *usecase.UpdateGlobalProductByID,
	getBusinessTypeFacets *usecase.GetBusinessTypeFacets,
	listProductsNeedingEnrichment *usecase.ListProductsNeedingEnrichment,
	getGlobalProductsByIDs *usecase.GetGlobalProductsByIDs,
	getDistinctBusinessTypes *usecase.GetDistinctBusinessTypes,
	criteriaBuilder *criteria.GlobalProductCriteriaBuilder,
) *GlobalCatalogController {
	return &GlobalCatalogController{
		createGlobalProduct:           createGlobalProduct,
		searchByEAN:                   searchByEAN,
		listGlobalProducts:            listGlobalProducts,
		listGlobalProductsByCriteria:  listGlobalProductsByCriteria,
		getGlobalProductByID:          getGlobalProductByID,
		updateGlobalProductByID:       updateGlobalProductByID,
		getBusinessTypeFacets:         getBusinessTypeFacets,
		listProductsNeedingEnrichment: listProductsNeedingEnrichment,
		getGlobalProductsByIDs:        getGlobalProductsByIDs,
		getDistinctBusinessTypes:      getDistinctBusinessTypes,
		criteriaBuilder:               criteriaBuilder,
	}
}

// NewGlobalCatalogControllerWithDeps crea el controlador con las dependencias extendidas (incluyendo delete/verify/bulk)
func NewGlobalCatalogControllerWithDeps(deps GlobalCatalogControllerDeps) *GlobalCatalogController {
	return &GlobalCatalogController{
		createGlobalProduct:           deps.CreateGlobalProduct,
		searchByEAN:                   deps.SearchByEAN,
		listGlobalProducts:            deps.ListGlobalProducts,
		listGlobalProductsByCriteria:  deps.ListGlobalProductsByCriteria,
		getGlobalProductByID:          deps.GetGlobalProductByID,
		updateGlobalProductByID:       deps.UpdateGlobalProductByID,
		deleteGlobalProduct:           deps.DeleteGlobalProduct,
		verifyGlobalProduct:           deps.VerifyGlobalProduct,
		unverifyGlobalProduct:         deps.UnverifyGlobalProduct,
		bulkImportGlobalProducts:      deps.BulkImportGlobalProducts,
		getBusinessTypeFacets:         deps.GetBusinessTypeFacets,
		listProductsNeedingEnrichment: deps.ListProductsNeedingEnrichment,
		getGlobalProductsByIDs:        deps.GetGlobalProductsByIDs,
		getDistinctBusinessTypes:      deps.GetDistinctBusinessTypes,
		criteriaBuilder:               deps.CriteriaBuilder,
	}
}

// RegisterRoutes registra las rutas del API
func (gc *GlobalCatalogController) RegisterRoutes(router *gin.RouterGroup) {
	// Rutas públicas (sin autenticación)
	public := router.Group("/public/global-catalog")
	{
		public.GET("/health", gc.HealthCheck)
		public.GET("/search", gc.SearchByEANPublic)           // Búsqueda pública por EAN
		public.GET("/suggestions", gc.GetProductsSuggestions) // Sugerencias por tipo de negocio
		public.GET("/products/ean/:ean", gc.GetProductByEAN)  // Obtener producto por EAN
		public.GET("/facets", gc.GetBusinessTypeFacets)       // Marcas y categorías por business_type
	}

	// Rutas privadas (para administración y scraping)
	private := router.Group("/global-catalog")
	{
		private.POST("/products", gc.CreateProduct)                        // Crear producto (para scrapers)
		private.POST("/products/bulk-import", gc.BulkImportProducts)       // Importación masiva
		private.GET("/products", gc.ListProducts)                          // Listar productos con filtros
		private.GET("/products/search", gc.SearchByEAN)                    // Búsqueda avanzada
		private.GET("/products/:id", gc.GetProductByID)                    // Obtener producto por ID
		private.GET("/enrichment-queue", gc.ListProductsNeedingEnrichment) // Cola de scraping para webdata
		private.GET("/by-ids", gc.GetProductsByIDs)                        // On-demand enrichment por IDs
		private.PUT("/products/:id", gc.UpdateProductByID)                 // Actualizar producto por ID
		private.DELETE("/products/:id", gc.DeleteProductByID)              // Eliminar producto por ID
		private.PATCH("/products/:id/verify", gc.VerifyProduct)            // Verificar producto
		private.PATCH("/products/:id/unverify", gc.UnverifyProduct)        // Desverificar producto
	}
}

// CreateProduct crea un nuevo producto en el catálogo global
// POST /api/v1/global-catalog/products
func (gc *GlobalCatalogController) CreateProduct(c *gin.Context) {
	var request usecase.CreateGlobalProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "JSON inválido",
			"details": err.Error(),
		})
		return
	}

	response, err := gc.createGlobalProduct.Execute(request)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

// SearchByEAN busca un producto por su código EAN (endpoint privado)
// GET /api/v1/global-catalog/products/search?ean={ean}&only_active={bool}
func (gc *GlobalCatalogController) SearchByEAN(c *gin.Context) {
	ean := c.Query("ean")
	if ean == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parámetro 'ean' es obligatorio",
		})
		return
	}

	onlyActive := c.Query("only_active") == "true"

	request := usecase.SearchByEANRequest{
		EAN:        ean,
		OnlyActive: onlyActive,
	}

	response, err := gc.searchByEAN.Execute(request)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// SearchByEANPublic búsqueda pública por EAN (solo productos activos)
// GET /api/v1/public/global-catalog/search?ean={ean}
func (gc *GlobalCatalogController) SearchByEANPublic(c *gin.Context) {
	ean := c.Query("ean")
	if ean == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parámetro 'ean' es obligatorio",
		})
		return
	}

	request := usecase.SearchByEANRequest{
		EAN:        ean,
		OnlyActive: true, // Solo productos activos para API pública
	}

	response, err := gc.searchByEAN.Execute(request)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetProductByEAN busca un producto específico por EAN (público)
// GET /api/v1/public/global-catalog/products/ean/{ean}
func (gc *GlobalCatalogController) GetProductByEAN(c *gin.Context) {
	ean := c.Param("ean")
	if ean == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "EAN no especificado en la URL",
		})
		return
	}

	request := usecase.SearchByEANRequest{
		EAN:        ean,
		OnlyActive: true, // Solo productos activos para API pública
	}

	response, err := gc.searchByEAN.Execute(request)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListProducts lista productos del catálogo global con filtros
// GET /api/v1/global-catalog/products
func (gc *GlobalCatalogController) ListProducts(c *gin.Context) {
	// Construir criterios desde los query params
	searchCriteria := gc.criteriaBuilder.BuildValidated(c)

	// Ejecutar la búsqueda con criterios
	response, err := gc.listGlobalProductsByCriteria.Execute(c.Request.Context(), searchCriteria)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	// Convertir entidades a DTOs para evitar objetos vacíos en JSON
	dtoItems := make([]*dto.GlobalProductResponse, len(response.Items))
	for i, product := range response.Items {
		dtoItems[i] = dto.NewGlobalProductResponse(product)
	}

	// can_request: true cuando hay búsqueda activa pero 0 resultados
	hasSearchFilters := c.Query("search") != "" || c.Query("brand") != "" || c.Query("category") != "" || c.Query("description") != "" || c.Query("ean") != ""
	canRequest := hasSearchFilters && len(dtoItems) == 0

	finalResponse := struct {
		Items      []*dto.GlobalProductResponse `json:"items"`
		TotalCount int                          `json:"total_count"`
		Page       int                          `json:"page"`
		PageSize   int                          `json:"page_size"`
		TotalPages int                          `json:"total_pages"`
		CanRequest bool                         `json:"can_request"`
	}{
		Items:      dtoItems,
		TotalCount: response.TotalCount,
		Page:       response.Page,
		PageSize:   response.PageSize,
		TotalPages: response.TotalPages,
		CanRequest: canRequest,
	}

	c.JSON(http.StatusOK, finalResponse)
}

// GetProductsSuggestions obtiene productos sugeridos por tipo de negocio (público)
// GET /api/v1/public/global-catalog/suggestions?business_type={type}&limit={limit}
func (gc *GlobalCatalogController) GetProductsSuggestions(c *gin.Context) {
	businessType := c.Query("business_type")
	if businessType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parámetro 'business_type' es obligatorio",
		})
		return
	}

	limit := 10
	if limitStr := c.Query("limit"); limitStr != "" {
		if val, err := strconv.Atoi(limitStr); err == nil && val > 0 && val <= 50 {
			limit = val
		}
	}

	request := usecase.ListGlobalProductsRequest{
		BusinessType:    &businessType,
		Limit:           limit,
		OnlyActive:      true,
		OnlyHighQuality: true, // Solo productos de alta calidad para sugerencias
	}

	response, err := gc.listGlobalProducts.Execute(request)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetBusinessTypeFacets devuelve marcas y categorías únicas por business_type (público)
// GET /api/v1/public/global-catalog/facets?business_type={type}
func (gc *GlobalCatalogController) GetBusinessTypeFacets(c *gin.Context) {
	businessType := c.Query("business_type")
	if businessType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parámetro 'business_type' es obligatorio",
		})
		return
	}

	facets, err := gc.getBusinessTypeFacets.Execute(businessType)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, facets)
}

// handleUseCaseError maneja los errores de casos de uso y los mapea a respuestas HTTP
func (gc *GlobalCatalogController) handleUseCaseError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *exception.ValidationError:
		body := gin.H{"error": e.Message}
		if e.Cause != nil {
			body["details"] = e.Cause.Error()
		}
		c.JSON(http.StatusBadRequest, body)
	case *exception.ConflictError:
		c.JSON(http.StatusConflict, gin.H{
			"error": e.Message,
		})
	case *exception.GlobalProductNotFoundError:
		c.JSON(http.StatusNotFound, gin.H{
			"error": e.Error(),
		})
	case *exception.InternalError:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error interno del servidor",
			"details": e.Message,
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error interno del servidor",
			"details": err.Error(),
		})
	}
}

// HealthCheck endpoint para verificar el estado del servicio
// GET /api/v1/public/global-catalog/health
func (gc *GlobalCatalogController) HealthCheck(c *gin.Context) {
	response := gin.H{
		"status":    "healthy",
		"service":   "global-catalog",
		"version":   "1.0.0",
		"timestamp": "2024-12-18T12:00:00Z",
	}

	c.JSON(http.StatusOK, response)
}

// GetProductByID obtiene un producto por su ID
// GET /api/v1/global-catalog/products/{id}
func (gc *GlobalCatalogController) GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID del producto no especificado",
		})
		return
	}

	request := usecase.GetGlobalProductByIDRequest{
		ID: productID,
	}

	response, err := gc.getGlobalProductByID.Execute(c.Request.Context(), request)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	// Convertir la entidad a DTO para la respuesta JSON
	productDTO := dto.NewGlobalProductResponse(response.Product)

	c.JSON(http.StatusOK, gin.H{
		"product": productDTO,
	})
}

// UpdateProductByID actualiza un producto por su ID
// PUT /api/v1/global-catalog/products/{id}
func (gc *GlobalCatalogController) UpdateProductByID(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID del producto no especificado",
		})
		return
	}

	var request usecase.UpdateGlobalProductByIDRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "JSON inválido",
			"details": err.Error(),
		})
		return
	}

	// Establecer el ID del producto desde la URL
	request.ID = productID

	response, err := gc.updateGlobalProductByID.Execute(c.Request.Context(), request)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	// Parsear fechas desde string a time.Time
	createdAt, _ := time.Parse("2006-01-02T15:04:05Z07:00", response.CreatedAt)
	updatedAt, _ := time.Parse("2006-01-02T15:04:05Z07:00", response.UpdatedAt)

	// Convertir la respuesta a DTO para la respuesta JSON
	productDTO := dto.GlobalProductResponse{
		ID:           response.ID,
		EAN:          response.EAN,
		Name:         response.Name,
		Description:  response.Description,
		Brand:        response.Brand,
		Category:     response.Category,
		Price:        response.Price,
		ImageURL:     response.ImageURL,
		ImageURLs:    response.ImageURLs,
		QualityScore: response.QualityScore,
		Source:       response.Source,
		SourceURL:    response.SourceURL,
		IsVerified:   response.IsVerified,
		IsActive:     response.IsActive,
		BusinessType: response.BusinessType,
		Tags:         response.Tags,
		Metadata:     response.Metadata,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"product": productDTO,
	})
}

// ListProductsNeedingEnrichment devuelve productos con datos incompletos para la cola de scraping.
// GET /api/v1/global-catalog/enrichment-queue?business_type=ferreteria&limit=100&offset=0
func (gc *GlobalCatalogController) ListProductsNeedingEnrichment(c *gin.Context) {
	req := usecase.ListProductsNeedingEnrichmentRequest{}

	if bt := c.Query("business_type"); bt != "" {
		req.BusinessType = &bt
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if v, err := strconv.Atoi(limitStr); err == nil {
			req.Limit = v
		}
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		if v, err := strconv.Atoi(offsetStr); err == nil {
			req.Offset = v
		}
	}

	response, err := gc.listProductsNeedingEnrichment.Execute(req)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetProductsByIDs devuelve productos por lista de IDs para on-demand enrichment.
// GET /api/v1/global-catalog/by-ids?ids=uuid1,uuid2,...
func (gc *GlobalCatalogController) GetProductsByIDs(c *gin.Context) {
	idsParam := c.Query("ids")
	if idsParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "parámetro 'ids' es obligatorio",
		})
		return
	}

	ids := splitAndTrim(idsParam)
	if len(ids) > 100 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "se permiten máximo 100 IDs por request",
		})
		return
	}

	req := usecase.GetGlobalProductsByIDsRequest{IDs: ids}
	response, err := gc.getGlobalProductsByIDs.Execute(c.Request.Context(), req)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// splitAndTrim divide una cadena por coma y elimina espacios de cada elemento.
func splitAndTrim(s string) []string {
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		if trimmed := strings.TrimSpace(p); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// DeleteProductByID elimina un producto por su ID
// DELETE /api/v1/global-catalog/products/{id}
func (gc *GlobalCatalogController) DeleteProductByID(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto no especificado"})
		return
	}

	if gc.deleteGlobalProduct == nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Endpoint no implementado"})
		return
	}

	req := usecase.DeleteGlobalProductRequest{ID: productID}
	if err := gc.deleteGlobalProduct.Execute(c.Request.Context(), req); err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// VerifyProduct marca un producto como verificado
// PATCH /api/v1/global-catalog/products/{id}/verify
func (gc *GlobalCatalogController) VerifyProduct(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto no especificado"})
		return
	}

	if gc.verifyGlobalProduct == nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Endpoint no implementado"})
		return
	}

	req := usecase.VerifyGlobalProductRequest{ID: productID}
	if err := gc.verifyGlobalProduct.Execute(c.Request.Context(), req); err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto verificado"})
}

// UnverifyProduct marca un producto como no verificado
// PATCH /api/v1/global-catalog/products/{id}/unverify
func (gc *GlobalCatalogController) UnverifyProduct(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del producto no especificado"})
		return
	}

	if gc.unverifyGlobalProduct == nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Endpoint no implementado"})
		return
	}

	req := usecase.UnverifyGlobalProductRequest{ID: productID}
	if err := gc.unverifyGlobalProduct.Execute(c.Request.Context(), req); err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto desverificado"})
}

// BulkImportProducts importa productos en lote
// POST /api/v1/global-catalog/products/bulk-import
func (gc *GlobalCatalogController) BulkImportProducts(c *gin.Context) {
	if gc.bulkImportGlobalProducts == nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Endpoint no implementado"})
		return
	}

	var req usecase.BulkImportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "details": err.Error()})
		return
	}

	result, err := gc.bulkImportGlobalProducts.Execute(req)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetDistinctBusinessTypes devuelve todos los business types distintos del catálogo global.
// GET /api/v1/global-catalog/business-types
func (gc *GlobalCatalogController) GetDistinctBusinessTypes(c *gin.Context) {
	types, err := gc.getDistinctBusinessTypes.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"business_types": types})
}
