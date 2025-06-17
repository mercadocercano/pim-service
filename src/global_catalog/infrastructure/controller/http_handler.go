package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"pim/src/global_catalog/application/usecase"
	"pim/src/global_catalog/domain/exception"
)

// GlobalCatalogController maneja las requests HTTP para el catálogo global
type GlobalCatalogController struct {
	createGlobalProduct *usecase.CreateGlobalProduct
	searchByEAN         *usecase.SearchByEAN
	listGlobalProducts  *usecase.ListGlobalProducts
}

// NewGlobalCatalogController crea una nueva instancia del controlador
func NewGlobalCatalogController(
	createGlobalProduct *usecase.CreateGlobalProduct,
	searchByEAN *usecase.SearchByEAN,
	listGlobalProducts *usecase.ListGlobalProducts,
) *GlobalCatalogController {
	return &GlobalCatalogController{
		createGlobalProduct: createGlobalProduct,
		searchByEAN:         searchByEAN,
		listGlobalProducts:  listGlobalProducts,
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
	}

	// Rutas privadas (para administración y scraping)
	private := router.Group("/global-catalog")
	{
		private.POST("/products", gc.CreateProduct)     // Crear producto (para scrapers)
		private.GET("/products", gc.ListProducts)       // Listar productos con filtros
		private.GET("/products/search", gc.SearchByEAN) // Búsqueda avanzada
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
	// Parámetros de paginación
	offset := 0
	limit := 20

	if offsetStr := c.Query("offset"); offsetStr != "" {
		if val, err := strconv.Atoi(offsetStr); err == nil && val >= 0 {
			offset = val
		}
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if val, err := strconv.Atoi(limitStr); err == nil && val > 0 && val <= 100 {
			limit = val
		}
	}

	// Filtros booleanos
	onlyActive := c.Query("only_active") == "true"
	onlyVerified := c.Query("only_verified") == "true"
	onlyArgentine := c.Query("only_argentine") == "true"
	onlyHighQuality := c.Query("only_high_quality") == "true"

	// Filtros de calidad
	var minQuality, maxQuality *int
	if minQualityStr := c.Query("min_quality"); minQualityStr != "" {
		if val, err := strconv.Atoi(minQualityStr); err == nil {
			minQuality = &val
		}
	}
	if maxQualityStr := c.Query("max_quality"); maxQualityStr != "" {
		if val, err := strconv.Atoi(maxQualityStr); err == nil {
			maxQuality = &val
		}
	}

	// Filtros de texto
	var businessType, source, searchName, searchBrand, searchCategory *string
	if bt := c.Query("business_type"); bt != "" {
		businessType = &bt
	}
	if s := c.Query("source"); s != "" {
		source = &s
	}
	if sn := c.Query("search_name"); sn != "" {
		searchName = &sn
	}
	if sb := c.Query("search_brand"); sb != "" {
		searchBrand = &sb
	}
	if sc := c.Query("search_category"); sc != "" {
		searchCategory = &sc
	}

	// Tags de búsqueda (separados por comas)
	var searchTags []string
	if tagsStr := c.Query("tags"); tagsStr != "" {
		searchTags = strings.Split(tagsStr, ",")
	}

	request := usecase.ListGlobalProductsRequest{
		Offset:          offset,
		Limit:           limit,
		BusinessType:    businessType,
		Source:          source,
		MinQuality:      minQuality,
		MaxQuality:      maxQuality,
		OnlyActive:      onlyActive,
		OnlyVerified:    onlyVerified,
		OnlyArgentine:   onlyArgentine,
		OnlyHighQuality: onlyHighQuality,
		SearchName:      searchName,
		SearchBrand:     searchBrand,
		SearchCategory:  searchCategory,
		SearchTags:      searchTags,
	}

	response, err := gc.listGlobalProducts.Execute(request)
	if err != nil {
		gc.handleUseCaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
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

// handleUseCaseError maneja los errores de casos de uso y los mapea a respuestas HTTP
func (gc *GlobalCatalogController) handleUseCaseError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *exception.ValidationError:
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   e.Message,
			"details": e.Cause.Error(),
		})
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
