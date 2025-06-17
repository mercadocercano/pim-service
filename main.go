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
	globalCatalogConfig "pim/src/global_catalog/infrastructure/config"
	"pim/src/marketplace/application/usecase"
	marketplaceConfig "pim/src/marketplace/infrastructure/config"
	"pim/src/marketplace/infrastructure/controller"
	"pim/src/marketplace/infrastructure/persistence"
	productConfig "pim/src/product/infrastructure/config"
	quickstartConfig "pim/src/quickstart/infrastructure/config"
	sharedConfig "pim/src/shared/infrastructure/config"
	"pim/src/shared/infrastructure/database"

	// Marketplace imports

	// Database imports

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
	setupProductModule(v1, db)
	setupQuickstartModule(v1, db)
	setupBusinessTypeModule(v1, db)
	setupGlobalCatalogModule(v1, db)
	// Usar el archivo de configuración del marketplace
	marketplaceConfig.SetupMarketplaceModule(v1, db, mongoClient)

	// Aquí se agregarían más módulos:
	// - Ubicaciones de Stock

	// Iniciar el servidor
	log.Println("Servidor iniciando en http://localhost:8080")
	router.Run(":8080")
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

	log.Println("Módulo Product configurado exitosamente")
	log.Println("Rutas Product disponibles:")
	log.Println("  POST   /api/v1/products")
	log.Println("  GET    /api/v1/products")
	log.Println("  GET    /api/v1/products/:id")
	log.Println("  PUT    /api/v1/products/:id")
	log.Println("  DELETE /api/v1/products/:id")
	log.Println("Rutas Product Variants disponibles:")
	log.Println("  POST   /api/v1/products/:product_id/variants")
	log.Println("  GET    /api/v1/products/:product_id/variants")
	log.Println("  GET    /api/v1/products/:product_id/variants/:variant_id")
	log.Println("  PUT    /api/v1/products/:product_id/variants/:variant_id")
	log.Println("  DELETE /api/v1/products/:product_id/variants/:variant_id")
	log.Println("  GET    /api/v1/variants")
}

// setupQuickstartModule configura el módulo Quickstart
func setupQuickstartModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo Quickstart...")

	// Crear el loader de datos YAML
	dataLoader := quickstartConfig.NewYAMLDataLoader("src/quickstart/data")

	// Crear configuración del módulo Quickstart
	quickstartCfg := quickstartConfig.NewQuickstartModuleConfig(db, dataLoader)

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

// setupMarketplaceModuleWithMongoDB configura el módulo Marketplace con MongoDB (COMPLETO)
func setupMarketplaceModuleWithMongoDB(router *gin.RouterGroup, db *sql.DB, mongoClient *database.MongoDBClient) {
	log.Println("Configurando módulo Marketplace con MongoDB...")

	// Crear repositorios MongoDB
	log.Println("🔧 Creando repositorios MongoDB...")
	tenantCustomAttributeRepo := persistence.NewTenantCustomAttributeMongoRepository(mongoClient.Database)
	tenantCategoryMappingRepo := persistence.NewTenantCategoryMappingMongoRepository(mongoClient.Database)

	// Crear repositorios PostgreSQL (para categorías marketplace)
	log.Println("🔧 Creando repositorios PostgreSQL...")
	marketplaceCategoryRepo := persistence.NewMarketplaceCategoryPostgresRepository(db)

	// Crear casos de uso para atributos personalizados
	log.Println("🔧 Creando casos de uso para atributos personalizados...")
	extendTenantAttributesUC := usecase.NewExtendTenantAttributesUseCase(
		marketplaceCategoryRepo, // Usar el repositorio correcto en lugar de nil
		tenantCustomAttributeRepo,
	)

	getTenantCustomAttributesUC := usecase.NewGetTenantCustomAttributesUseCase(
		tenantCustomAttributeRepo,
	)

	updateTenantCustomAttributeUC := usecase.NewUpdateTenantCustomAttributeUseCase(
		tenantCustomAttributeRepo,
	)

	deleteTenantCustomAttributeUC := usecase.NewDeleteTenantCustomAttributeUseCase(
		tenantCustomAttributeRepo,
	)

	// Crear casos de uso para categorías marketplace
	log.Println("🔧 Creando casos de uso para categorías marketplace...")
	createMarketplaceCategoryUC := usecase.NewCreateMarketplaceCategoryUseCase(marketplaceCategoryRepo)
	getAllMarketplaceCategoriesUC := usecase.NewGetAllMarketplaceCategoriesUseCase(marketplaceCategoryRepo)
	updateMarketplaceCategoryUC := usecase.NewUpdateMarketplaceCategoryUseCase(marketplaceCategoryRepo)
	getTenantTaxonomyUC := usecase.NewGetTenantTaxonomyUseCase(marketplaceCategoryRepo, tenantCategoryMappingRepo, tenantCustomAttributeRepo)
	validateCategoryHierarchyUC := usecase.NewValidateCategoryHierarchyUseCase(marketplaceCategoryRepo)
	syncMarketplaceChangesUC := usecase.NewSyncMarketplaceChangesUseCase(marketplaceCategoryRepo, tenantCategoryMappingRepo, tenantCustomAttributeRepo)

	// Crear caso de uso para mapeo de categorías
	log.Println("🔧 Creando caso de uso para mapeo de categorías...")
	mapTenantCategoryUC := usecase.NewMapTenantCategoryUseCase(marketplaceCategoryRepo, tenantCategoryMappingRepo)

	// Crear controladores
	log.Println("🔧 Creando controladores...")
	tenantCustomAttributeHandler := controller.NewTenantCustomAttributeHandler(
		extendTenantAttributesUC,
		getTenantCustomAttributesUC,
		updateTenantCustomAttributeUC,
		deleteTenantCustomAttributeUC,
	)

	marketplaceCategoryHandler := controller.NewMarketplaceCategoryHandler(
		createMarketplaceCategoryUC,
		getAllMarketplaceCategoriesUC,
		updateMarketplaceCategoryUC,
		getTenantTaxonomyUC,
		validateCategoryHierarchyUC,
		syncMarketplaceChangesUC,
	)

	tenantCategoryMappingHandler := controller.NewTenantCategoryMappingHandler(
		mapTenantCategoryUC,
	)

	log.Println("🔧 Controladores creados exitosamente")

	// Configurar rutas marketplace
	log.Println("🔧 Configurando rutas marketplace...")
	log.Println("🔧 LÍNEA 340 - ANTES DE CREAR GRUPO")
	log.Println("🔧 Creando grupo marketplace...")
	marketplace := router.Group("/marketplace")
	log.Println("🔧 Grupo marketplace creado exitosamente")
	{
		log.Println("🔧 Configurando endpoints básicos...")

		// Registrar rutas de categorías marketplace directamente
		log.Println("🔧 Registrando ruta GET /categories...")
		marketplace.GET("/categories", marketplaceCategoryHandler.GetAllMarketplaceCategories)
		log.Println("🔧 Registrando ruta POST /categories...")
		marketplace.POST("/categories", marketplaceCategoryHandler.CreateMarketplaceCategory)
		log.Println("🔧 Registrando ruta PUT /categories/:id...")
		log.Println("🔧 Antes de crear función PUT...")
		// Endpoint temporal de prueba
		putHandler := func(c *gin.Context) {
			log.Printf("🔧 PUT endpoint llamado con ID: %s", c.Param("id"))
			c.JSON(200, gin.H{"message": "PUT endpoint funcionando", "id": c.Param("id")})
		}
		log.Println("🔧 Función PUT creada, registrando ruta...")
		marketplace.PUT("/categories/:id", putHandler)
		log.Println("🔧 Ruta PUT registrada exitosamente")
		log.Println("🔧 Registrando ruta POST /categories/validate-hierarchy...")
		marketplace.POST("/categories/validate-hierarchy", marketplaceCategoryHandler.ValidateCategoryHierarchy)
		log.Println("🔧 Registrando ruta POST /sync-changes...")
		marketplace.POST("/sync-changes", marketplaceCategoryHandler.SyncMarketplaceChanges)
		log.Println("🔧 Registrando ruta GET /taxonomy...")
		marketplace.GET("/taxonomy", marketplaceCategoryHandler.GetTenantTaxonomy)

		// Rutas para tenants (atributos personalizados y mapeos)
		tenantGroup := marketplace.Group("/tenant")
		{
			// Atributos personalizados
			tenantGroup.POST("/custom-attributes", tenantCustomAttributeHandler.ExtendTenantAttributes)
			tenantGroup.GET("/custom-attributes", tenantCustomAttributeHandler.GetTenantCustomAttributes)
			tenantGroup.PUT("/custom-attributes/:attribute_id", tenantCustomAttributeHandler.UpdateTenantCustomAttribute)
			tenantGroup.DELETE("/custom-attributes/:attribute_id", tenantCustomAttributeHandler.DeleteTenantCustomAttribute)

			// Mapeos de categorías
			tenantGroup.POST("/category-mappings", tenantCategoryMappingHandler.MapTenantCategory)
			tenantGroup.PUT("/category-mappings/:mapping_id", tenantCategoryMappingHandler.UpdateTenantCategoryMapping)
			tenantGroup.DELETE("/category-mappings/:mapping_id", tenantCategoryMappingHandler.DeleteTenantCategoryMapping)
		}
	}

	log.Println("Módulo Marketplace configurado exitosamente con MongoDB")
	log.Println("Rutas Marketplace disponibles:")
	log.Println("  GET    /api/v1/marketplace/health")
	log.Println("  GET    /api/v1/marketplace/test-mongo")
	log.Println("  GET    /api/v1/marketplace/categories (admin)")
	log.Println("  POST   /api/v1/marketplace/categories (admin)")
	log.Println("  PUT    /api/v1/marketplace/categories/:id (admin)")
	log.Println("  POST   /api/v1/marketplace/categories/validate-hierarchy (admin)")
	log.Println("  POST   /api/v1/marketplace/sync-changes (admin)")
	log.Println("  GET    /api/v1/marketplace/taxonomy (tenant)")
	log.Println("  POST   /api/v1/marketplace/tenant/custom-attributes")
	log.Println("  GET    /api/v1/marketplace/tenant/custom-attributes")
	log.Println("  PUT    /api/v1/marketplace/tenant/custom-attributes/:attribute_id")
	log.Println("  DELETE /api/v1/marketplace/tenant/custom-attributes/:attribute_id")
	log.Println("  POST   /api/v1/marketplace/tenant/category-mappings")
	log.Println("  PUT    /api/v1/marketplace/tenant/category-mappings/:mapping_id")
	log.Println("  DELETE /api/v1/marketplace/tenant/category-mappings/:mapping_id")
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
}
