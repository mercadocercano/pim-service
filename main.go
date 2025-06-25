package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	apiConfig "pim/src/api/config"
	brandConfig "pim/src/brand/infrastructure/config"
	businesstypeUsecase "pim/src/businesstype/application/usecase"
	businesstypeController "pim/src/businesstype/infrastructure/controller"
	businesstypeRepository "pim/src/businesstype/infrastructure/persistence/repository"
	categoryConfig "pim/src/category/infrastructure/config"
	categoryAttributeConfig "pim/src/category_attribute/infrastructure/config"
	globalCatalogConfig "pim/src/product/global_catalog/infrastructure/config"
	productConfig "pim/src/product/tenant/infrastructure/config"
	quickstartConfig "pim/src/quickstart/infrastructure/config"
	sharedConfig "pim/src/shared/infrastructure/config"
	"pim/src/shared/infrastructure/database"

	// Brand imports
	brandController "pim/src/brand/infrastructure/controller"
	brandRepository "pim/src/brand/infrastructure/persistence/repository"

	// Attribute imports

	// Category imports

	// Marketplace Categories imports
	categoryUsecase "pim/src/category/application/usecase"
	categoryController "pim/src/category/infrastructure/controller"
	categoryPersistence "pim/src/category/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Driver de PostgreSQL
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// getEnv obtiene una variable de entorno o devuelve un valor por defecto
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {
	fmt.Println("🚀🚀🚀 INICIO ABSOLUTO DEL MAIN.GO 🚀🚀🚀")
	log.Println("🚀 *** MAIN.GO DE LA RAÍZ EJECUTÁNDOSE - VERSIÓN NUEVA ***")
	// Configurar el router con Gin
	router := gin.New()

	// Agregar middlewares básicos necesarios
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Configurar Prometheus metrics si está habilitado
	prometheusEnabled := os.Getenv("PROMETHEUS_ENABLED")
	log.Printf("PROMETHEUS_ENABLED value: '%s'", prometheusEnabled)

	if prometheusEnabled == "true" {
		log.Println("Registering /metrics endpoint for PIM service")
		// Endpoint de métricas usando la librería oficial de Prometheus
		router.GET("/metrics", gin.WrapH(promhttp.Handler()))
		log.Println("/metrics endpoint registered successfully for PIM service")
	} else {
		log.Println("Prometheus metrics disabled for PIM service")
	}

	// Cargar plantillas HTML
	router.LoadHTMLGlob("templates/*")

	// Configurar middlewares compartidos
	sharedCfg := sharedConfig.DefaultSharedConfig()
	sharedConfig.SetupSharedMiddleware(router, sharedCfg)

	// Obtener configuración de la base de datos de variables de entorno
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "pim_db")

	// Crear string de conexión
	connStr := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	log.Printf("Intentando conectar a %s", connStr)

	// Conectar a la base de datos
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Comprobar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error al verificar la conexión a la base de datos: %v", err)
	}
	log.Println("Conexión a la base de datos establecida con éxito")

	// Configurar MongoDB para marketplace
	log.Println("Inicializando conexión MongoDB para marketplace...")
	mongoClient, err := database.NewMongoDBClient()
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}
	defer mongoClient.Close()

	// Verificar conexión MongoDB
	if err := mongoClient.HealthCheck(nil); err != nil {
		log.Fatalf("Error al verificar la conexión a MongoDB: %v", err)
	}
	log.Println("Conexión a MongoDB establecida con éxito")

	// API v1 grupo de rutas
	v1 := router.Group("/api/v1")

	// Configurar el módulo API (health check y documentación)
	apiCfg := apiConfig.DefaultAPIConfig()
	apiCfg.DB = db
	apiCfg.Version = "1.0.0"
	apiConfig.SetupAPIModule(router, v1, apiCfg)

	// Configurar módulos
	categoryConfig.SetupCategoryModule(v1, db)
	categoryAttributeConfig.SetupCategoryAttributeModule(v1, db)
	log.Println("Módulo CategoryAttribute configurado exitosamente")
	log.Println("Rutas CategoryAttribute disponibles:")
	log.Println("  GET    /api/v1/category-attributes (con filtros y paginación)")
	log.Println("  GET    /api/v1/category-attributes/simple (listado simple)")
	log.Println("  POST   /api/v1/category-attributes")
	log.Println("  PUT    /api/v1/category-attributes/:id")
	log.Println("  DELETE /api/v1/category-attributes/:id")
	setupBrandModule(v1, db)
	setupMarketplaceBrandModule(v1, db)
	setupProductModule(v1, db)
	setupQuickstartModule(v1, db)
	setupBusinessTypeModule(v1, db)
	setupGlobalCatalogModule(v1, db)
	setupMarketplaceCategoriesModule(v1, db)

	// Aquí se agregarían más módulos:
	// - Ubicaciones de Stock

	// Iniciar el servidor
	port := getEnv("PORT", "8090")
	log.Printf("Servidor iniciando en http://localhost:%s", port)
	router.Run(":" + port)
}

// setupBrandModule configura el módulo Brand
func setupBrandModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo Brand...")

	// Crear configuración del módulo Brand
	brandCfg := brandConfig.NewBrandModuleConfig(db)

	// Obtener el controller
	brandController := brandCfg.GetBrandController()

	// Configurar rutas de Brand
	brands := router.Group("/brands")
	{
		brands.POST("", brandController.CreateBrand)
		brands.GET("", brandController.ListBrands)
		brands.GET("/:id", brandController.GetBrand)
		brands.PUT("/:id", brandController.UpdateBrand)
		brands.DELETE("/:id", brandController.DeleteBrand)
	}

	log.Println("Módulo Brand configurado exitosamente")
	log.Println("Rutas Brand disponibles:")
	log.Println("  POST   /api/v1/brands")
	log.Println("  GET    /api/v1/brands")
	log.Println("  GET    /api/v1/brands/:id")
	log.Println("  PUT    /api/v1/brands/:id")
	log.Println("  DELETE /api/v1/brands/:id")
}

// setupProductModule configura el módulo Product
func setupProductModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo Product...")

	// Crear configuración del módulo Product
	productCfg := productConfig.NewProductConfig(db)

	// Registrar rutas del Product
	productCfg.ProductController.RegisterRoutes(router)

	// Registrar rutas de Product Variants
	productCfg.ProductVariantController.RegisterRoutes(router)

	// Registrar rutas del Quickstart
	productCfg.QuickstartController.RegisterRoutes(router)

	log.Println("Módulo Product configurado exitosamente")
	log.Println("Rutas Product disponibles:")
	log.Println("  POST   /api/v1/products")
	log.Println("  GET    /api/v1/products")
	log.Println("  GET    /api/v1/products/:id")
	log.Println("  PUT    /api/v1/products/:id")
	log.Println("  DELETE /api/v1/products/:id")
	log.Println("  PATCH  /api/v1/products/:id/status")
	log.Println("  GET    /api/v1/products/:id/status/transitions")
	log.Println("Rutas Product Variants disponibles:")
	log.Println("  POST   /api/v1/products/:product_id/variants")
	log.Println("  GET    /api/v1/products/:product_id/variants")
	log.Println("  GET    /api/v1/products/:product_id/variants/:variant_id")
	log.Println("  PUT    /api/v1/products/:product_id/variants/:variant_id")
	log.Println("  DELETE /api/v1/products/:product_id/variants/:variant_id")
	log.Println("  GET    /api/v1/variants")
	log.Println("Rutas Quickstart disponibles:")
	log.Println("  POST   /api/v1/quickstart/products/from-template")
	log.Println("  POST   /api/v1/quickstart/products/import-from-business-type")
	log.Println("  GET    /api/v1/quickstart/progress")
}

// setupQuickstartModule configura el módulo Quickstart
func setupQuickstartModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo Quickstart...")

	// Crear configuración del módulo Product para obtener el ProductService
	productCfg := productConfig.NewProductConfig(db)

	// Crear el loader de datos YAML
	dataLoader := quickstartConfig.NewYAMLDataLoader("src/quickstart/data")

	// Crear configuración del módulo Quickstart con ProductService
	quickstartCfg := quickstartConfig.NewQuickstartModuleConfig(db, dataLoader, productCfg.QuickstartProductService)

	// Obtener el handler
	quickstartHandler := quickstartCfg.GetQuickstartHandler()

	// Registrar rutas usando el método RegisterRoutes del handler
	quickstartHandler.RegisterRoutes(router)

	log.Println("Módulo Quickstart configurado exitosamente")
	log.Println("Rutas Quickstart disponibles:")
	log.Println("  GET    /api/v1/quickstart/business-types")
	log.Println("  GET    /api/v1/quickstart/categories/:businessType")
	log.Println("  GET    /api/v1/quickstart/attributes/:businessType")
	log.Println("  GET    /api/v1/quickstart/variants/:businessType")
	log.Println("  GET    /api/v1/quickstart/products/:businessType")
	log.Println("  GET    /api/v1/quickstart/brands/:businessType")
	log.Println("  POST   /api/v1/quickstart/setup")
}

// setupBusinessTypeModule configura el módulo BusinessType
func setupBusinessTypeModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo BusinessType...")

	// Crear repositorio
	businessTypeRepo := businesstypeRepository.NewBusinessTypePostgresRepository(db)

	// Crear casos de uso
	createBusinessTypeUC := businesstypeUsecase.NewCreateBusinessTypeUseCase(businessTypeRepo)
	listBusinessTypesUC := businesstypeUsecase.NewListBusinessTypesUseCase(businessTypeRepo)
	getBusinessTypeUC := businesstypeUsecase.NewGetBusinessTypeUseCase(businessTypeRepo)
	updateBusinessTypeUC := businesstypeUsecase.NewUpdateBusinessTypeUseCase(businessTypeRepo)

	// Crear handler
	businessTypeHandler := businesstypeController.NewBusinessTypeHandler(
		createBusinessTypeUC,
		listBusinessTypesUC,
		getBusinessTypeUC,
		updateBusinessTypeUC,
	)

	// Registrar rutas
	businessTypeHandler.RegisterRoutes(router)

	log.Println("Módulo BusinessType configurado exitosamente")
	log.Println("Rutas BusinessType disponibles:")
	log.Println("  POST   /api/v1/business-types")
	log.Println("  GET    /api/v1/business-types")
	log.Println("  GET    /api/v1/business-types/:id")
	log.Println("  PUT    /api/v1/business-types/:id")
	log.Println("  DELETE /api/v1/business-types/:id")
}

// setupGlobalCatalogModule configura el módulo Global Catalog
func setupGlobalCatalogModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo Global Catalog...")

	// Crear configuración del módulo Global Catalog
	globalCatalogCfg := globalCatalogConfig.NewGlobalCatalogConfig(db)

	// Obtener el controller
	globalCatalogController := globalCatalogCfg.GetGlobalCatalogController()

	// Configurar rutas públicas (sin autenticación)
	public := router.Group("/public/global-catalog")
	{
		public.GET("/health", globalCatalogController.HealthCheck)
		public.GET("/search", globalCatalogController.SearchByEANPublic)
		public.GET("/suggestions", globalCatalogController.GetProductsSuggestions)
		public.GET("/products/ean/:ean", globalCatalogController.GetProductByEAN)
	}

	// Configurar rutas privadas (para scrapers y admin)
	private := router.Group("/global-catalog")
	{
		private.POST("/products", globalCatalogController.CreateProduct)
		private.GET("/products", globalCatalogController.ListProducts)
		private.GET("/products/search", globalCatalogController.SearchByEAN)
		private.GET("/products/:id", globalCatalogController.GetProductByID)
		private.PUT("/products/:id", globalCatalogController.UpdateProductByID)
		private.DELETE("/products/:id", globalCatalogController.DeleteProductByID)
	}

	log.Println("Módulo Global Catalog configurado exitosamente")
	log.Println("Rutas Global Catalog públicas disponibles:")
	log.Println("  GET    /api/v1/public/global-catalog/health")
	log.Println("  GET    /api/v1/public/global-catalog/search?ean={ean}")
	log.Println("  GET    /api/v1/public/global-catalog/suggestions?business_type={type}")
	log.Println("  GET    /api/v1/public/global-catalog/products/ean/{ean}")
	log.Println("Rutas Global Catalog privadas disponibles:")
	log.Println("  POST   /api/v1/global-catalog/products")
	log.Println("  GET    /api/v1/global-catalog/products")
	log.Println("  GET    /api/v1/global-catalog/products/search?ean={ean}")
	log.Println("  GET    /api/v1/global-catalog/products/:id")
	log.Println("  PUT    /api/v1/global-catalog/products/:id")
	log.Println("  DELETE /api/v1/global-catalog/products/:id")
}

// setupMarketplaceCategoriesModule configura el módulo Marketplace Categories
func setupMarketplaceCategoriesModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo Marketplace Categories...")

	// Crear repositorio marketplace categories
	marketplaceCategoryRepo := categoryPersistence.NewMarketplaceCategoryPostgresRepository(db)

	// Crear casos de uso básicos (solo los que funcionan con un repositorio)
	createMarketplaceCategoryUC := categoryUsecase.NewCreateMarketplaceCategoryUseCase(marketplaceCategoryRepo)
	getAllMarketplaceCategoriesUC := categoryUsecase.NewGetAllMarketplaceCategoriesUseCase(marketplaceCategoryRepo)
	updateMarketplaceCategoryUC := categoryUsecase.NewUpdateMarketplaceCategoryUseCase(marketplaceCategoryRepo)
	validateCategoryHierarchyUC := categoryUsecase.NewValidateCategoryHierarchyUseCase(marketplaceCategoryRepo)

	// Crear handler con casos de uso básicos
	marketplaceCategoryHandler := categoryController.NewMarketplaceCategoryHandler(
		createMarketplaceCategoryUC,
		getAllMarketplaceCategoriesUC,
		updateMarketplaceCategoryUC,
		nil, // getTenantTaxonomyUC - requiere múltiples repos
		validateCategoryHierarchyUC,
		nil,                     // syncMarketplaceChangesUC - requiere múltiples repos
		marketplaceCategoryRepo, // Agregar el repositorio
	)

	// Registrar rutas
	marketplaceCategoryHandler.RegisterRoutes(router)

	log.Println("Módulo Marketplace Categories configurado exitosamente")
	log.Println("Rutas Marketplace Categories disponibles:")
	log.Println("  GET    /api/v1/marketplace/categories")
	log.Println("  POST   /api/v1/marketplace/categories")
	log.Println("  PUT    /api/v1/marketplace/categories/:id")
	log.Println("  POST   /api/v1/marketplace/categories/validate-hierarchy")
}

// setupMarketplaceBrandModule configura el módulo Marketplace Brand
func setupMarketplaceBrandModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo Marketplace Brand...")

	// Crear repositorio marketplace brands
	marketplaceBrandRepo := brandRepository.NewMarketplacebrandPostgresRepository(db)

	// Crear handler con repositorio
	marketplaceBrandHandler := brandController.NewMarketplaceBrandHandler(marketplaceBrandRepo)

	// Registrar rutas
	marketplaceBrandHandler.RegisterRoutes(router)

	log.Println("Módulo Marketplace Brand configurado exitosamente")
	log.Println("Rutas Marketplace Brand disponibles:")
	log.Println("  GET    /api/v1/marketplace-brands")
	log.Println("  POST   /api/v1/marketplace-brands")
	log.Println("  GET    /api/v1/marketplace-brands/:id")
	log.Println("  PUT    /api/v1/marketplace-brands/:id")
	log.Println("  DELETE /api/v1/marketplace-brands/:id")
	log.Println("  PUT    /api/v1/marketplace-brands/:id/verify")
	log.Println("  PUT    /api/v1/marketplace-brands/:id/unverify")
}
