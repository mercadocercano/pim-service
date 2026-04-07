package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	apiConfig "saas-mt-pim-service/src/api/config"
	brandConfig "saas-mt-pim-service/src/brand/infrastructure/config"
	businesstypeUsecase "saas-mt-pim-service/src/businesstype/application/usecase"
	businesstypeController "saas-mt-pim-service/src/businesstype/infrastructure/controller"
	businesstypeRepository "saas-mt-pim-service/src/businesstype/infrastructure/persistence/repository"
	categoryConfig "saas-mt-pim-service/src/category/infrastructure/config"
	categoryAttributeConfig "saas-mt-pim-service/src/category_attribute/infrastructure/config"
	globalCatalogConfig "saas-mt-pim-service/src/product/global_catalog/infrastructure/config"
	productConfig "saas-mt-pim-service/src/product/tenant/infrastructure/config"
	quickstartConfig "saas-mt-pim-service/src/quickstart/infrastructure/config"
	sharedConfig "saas-mt-pim-service/src/shared/infrastructure/config"
	"saas-mt-pim-service/src/shared/infrastructure/database"

	// Brand imports
	brandController "saas-mt-pim-service/src/brand/infrastructure/controller"
	brandPersistence "saas-mt-pim-service/src/brand/infrastructure/persistence"
	brandRepository "saas-mt-pim-service/src/brand/infrastructure/persistence/repository"

	// Attribute imports
	attributeConfig "saas-mt-pim-service/src/attribute/infrastructure/config"

	// Business Type Template imports
	businessTypeTemplateConfig "saas-mt-pim-service/src/businesstype/infrastructure/config"

	// Category imports

	// Marketplace Categories imports
	categoryUsecase "saas-mt-pim-service/src/category/application/usecase"
	categoryController "saas-mt-pim-service/src/category/infrastructure/controller"
	categoryPersistence "saas-mt-pim-service/src/category/infrastructure/persistence"
	categoryRepository "saas-mt-pim-service/src/category/infrastructure/persistence/repository"

	// Batch imports
	batchPort "saas-mt-pim-service/src/batch/domain/port"
	batchUseCase "saas-mt-pim-service/src/batch/application/usecase"
	batchController "saas-mt-pim-service/src/batch/infrastructure/controller"

	// Product persistence
	persistence "saas-mt-pim-service/src/product/tenant/infrastructure/persistence"

	// Schema validation imports
	schemaValidationConfig "saas-mt-pim-service/src/schema_validation/infrastructure/config"

	// AI Template imports
	aiTemplateConfig "saas-mt-pim-service/src/template_ai/infrastructure/config"

	// Overview module imports
	overviewConfig "saas-mt-pim-service/src/overview/infrastructure/config"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Driver de PostgreSQL
	tenantmw "github.com/mercadocercano/middleware"
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
	router.MaxMultipartMemory = 50 << 20 // 50MB

	// Agregar middlewares básicos necesarios
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(tenantmw.TenantValidation(tenantmw.TenantValidationConfig{
		JWTSecret: os.Getenv("JWT_SECRET"),
		ExcludedRoutes: []string{
			"/api/v1/health",
			"/metrics",
			"/api/v1/marketplace-brands*",
			"/api/v1/marketplace/categories*",
			"/api/v1/marketplace-attributes*",
			"/api/v1/business-types*",
			"/api/v1/business-type-templates*",
			"/api/v1/global-catalog*",
			"/api/v1/public/global-catalog*",
		},
	}))

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

	// MongoDB: conexión condicional (solo si MONGO_HOST está configurado)
	if mongoHost := os.Getenv("MONGO_HOST"); mongoHost != "" {
		log.Printf("Inicializando conexión MongoDB (%s)...", mongoHost)
		mongoClient, err := database.NewMongoDBClient()
		if err != nil {
			log.Printf("⚠️ MongoDB no disponible: %v (continuando sin MongoDB)", err)
		} else {
			defer mongoClient.Close()
			if err := mongoClient.HealthCheck(nil); err != nil {
				log.Printf("⚠️ MongoDB health check falló: %v (continuando sin MongoDB)", err)
			} else {
				log.Println("✅ Conexión a MongoDB establecida")
			}
		}
	} else {
		log.Println("ℹ️ MONGO_HOST no configurado, MongoDB deshabilitado")
	}

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
	setupAttributeModule(v1, db)
	setupProductModule(v1, db)
	setupQuickstartModule(v1, db)
	setupBusinessTypeModule(v1, db)
	setupBusinessTypeTemplateModule(v1, db)
	setupGlobalCatalogModule(v1, db)
	setupMarketplaceCategoriesModule(v1, db)
	setupBatchModule(v1, db)
	setupSchemaValidationModule(v1, db)
	setupAITemplateModule(v1, db)
	overviewConfig.SetupOverviewModule(v1, db)

	// Aquí se agregarían más módulos:
	// - Ubicaciones de Stock

	// Iniciar el servidor
	port := getEnv("PORT", "8080")
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

	// Registrar rutas de Product Variants (standalone)
	productCfg.ProductVariantController.RegisterRoutes(router)
	
	// Registrar rutas anidadas de variantes bajo productos (DESPUÉS de las rutas de productos)
	productVariantCtrl := productCfg.ProductVariantController
	router.POST("/products/:id/variants", productVariantCtrl.CreateProductVariant)
	router.GET("/products/:id/variants", productVariantCtrl.ListProductVariants)
	router.GET("/products/:id/variants/:variant_id", productVariantCtrl.GetProductVariant)
	router.PUT("/products/:id/variants/:variant_id", productVariantCtrl.UpdateProductVariant)
	router.DELETE("/products/:id/variants/:variant_id", productVariantCtrl.DeleteProductVariant)

	// Registrar rutas del Quickstart
	productCfg.QuickstartController.RegisterRoutes(router)

	// HITO 2: Registrar rutas de importación bulk
	if productCfg.BulkImportController != nil {
		productCfg.BulkImportController.RegisterRoutes(router)
		log.Println("  POST   /api/v1/products/import (HITO 2)")
	}

	if productCfg.BulkUpdateController != nil {
		productCfg.BulkUpdateController.RegisterRoutes(router)
		log.Println("  PATCH  /api/v1/products/bulk-update")
	}

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

	// Crear configuración del módulo Quickstart
	quickstartCfg := quickstartConfig.NewQuickstartModuleConfig(db)

	// Obtener el handler principal de quickstart (incluye /templates)
	quickstartHandler := quickstartCfg.GetQuickstartHandler()
	quickstartHandler.RegisterRoutes(router)

	// Obtener el handler simplificado del wizard
	simpleWizardHandler := quickstartCfg.GetSimpleWizardHandler()

	// Registrar rutas del wizard simplificado
	simpleWizardHandler.RegisterRoutes(router)

	log.Println("Módulo Quickstart configurado exitosamente")
	log.Println("Rutas Quickstart disponibles:")
	log.Println("  GET    /api/v1/quickstart/templates")
	log.Println("  POST   /api/v1/quickstart/apply")
	log.Println("Rutas Wizard disponibles:")
	log.Println("  GET    /api/v1/wizard/status")
	log.Println("  POST   /api/v1/wizard/start")
	log.Println("  PUT    /api/v1/wizard/step")
	log.Println("  GET    /api/v1/wizard/template/:businessTypeId")
	log.Println("  GET    /api/v1/wizard/template/:businessTypeId/:section")
	log.Println("  DELETE /api/v1/wizard/reset (⚠️ TEMPORAL - BORRAR DESPUÉS)")
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
		businessTypeRepo,
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

// setupBusinessTypeTemplateModule configura el módulo BusinessTypeTemplate
func setupBusinessTypeTemplateModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo BusinessTypeTemplate...")

	// Crear configuración del módulo
	templateCfg := businessTypeTemplateConfig.NewBusinessTypeTemplateModuleConfig(db)

	// Obtener el handler
	templateHandler := templateCfg.GetHandler()

	// Registrar rutas
	templateHandler.RegisterRoutes(router)

	log.Println("Módulo BusinessTypeTemplate configurado exitosamente")
	log.Println("Rutas BusinessTypeTemplate disponibles:")
	log.Println("  POST   /api/v1/business-type-templates")
	log.Println("  GET    /api/v1/business-type-templates")
	log.Println("  GET    /api/v1/business-type-templates/:id")
	log.Println("  PUT    /api/v1/business-type-templates/:id")
	log.Println("  DELETE /api/v1/business-type-templates/:id")
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
		validateCategoryHierarchyUC,
		marketplaceCategoryRepo,
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

// setupAttributeModule configura el módulo Attribute
func setupAttributeModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo Attribute...")

	// Usar la función del config module
	attributeConfig.SetupAttributeModule(router, db)

	log.Println("Módulo Attribute configurado exitosamente")
	log.Println("Rutas Marketplace Attributes disponibles:")
	log.Println("  GET    /api/v1/marketplace/attributes")
	log.Println("  POST   /api/v1/marketplace/attributes")
	log.Println("  GET    /api/v1/marketplace/attributes/:id")
	log.Println("  PUT    /api/v1/marketplace/attributes/:id")
	log.Println("  DELETE /api/v1/marketplace/attributes/:id")
	log.Println("Rutas Tenant Attributes disponibles:")
	log.Println("  POST   /api/v1/attributes")
	log.Println("  GET    /api/v1/attributes")
	log.Println("  GET    /api/v1/attributes/:id")
	log.Println("  PUT    /api/v1/attributes/:id")
	log.Println("  DELETE /api/v1/attributes/:id")
}

// setupBatchModule configura el módulo Batch
func setupBatchModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo Batch...")

	// Importar dependencias necesarias
	// TODO: Mejorar importaciones cuando se refactorice
	categoryRepo := categoryRepository.NewCategoryPostgresRepository(db)
	brandRepo := brandPersistence.NewPostgresBrandRepository(db)
	productRepo := persistence.NewPostgresProductRepository(db)

	// Repository para mapeo de categorías
	categoryMappingRepo := categoryPersistence.NewTenantCategoryMappingPostgresRepository(db)

	// Crear caso de uso batch
	txBeginner := &batchPort.SQLDBTxBeginner{DB: db}
	batchUseCase := batchUseCase.NewBatchCreateUseCase(
		txBeginner,
		categoryRepo,
		brandRepo,
		productRepo,
		categoryMappingRepo,
	)

	// Crear controller
	batchController := batchController.NewBatchController(batchUseCase)

	// Registrar rutas
	batchController.RegisterRoutes(router)

	log.Println("Módulo Batch configurado exitosamente")
	log.Println("Rutas Batch disponibles:")
	log.Println("  POST   /api/v1/batch/create")
}

// setupSchemaValidationModule configura el módulo de validación de schema
func setupSchemaValidationModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo Schema Validation...")

	// Usar la función del config module
	schemaValidationConfig.SetupSchemaValidationModule(router, db)

	log.Println("Módulo Schema Validation configurado exitosamente")
	log.Println("Rutas Schema Validation disponibles:")
	log.Println("  POST   /api/v1/products/validate-schema")
	log.Println("  POST   /api/v1/products/apply-mapping")
	log.Println("  GET    /api/v1/products/csv-template")
}

// setupAITemplateModule configura el módulo de AI Templates
func setupAITemplateModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo AI Template...")

	// Crear configuración y obtener controller
	config := aiTemplateConfig.NewAITemplateConfig(db)
	controller := config.GetController()

	// Registrar rutas
	controller.RegisterRoutes(router)

	log.Println("Módulo AI Template configurado exitosamente")
	log.Println("Rutas AI Template disponibles:")
	log.Println("  POST   /api/v1/templates/generate")
	log.Println("  POST   /api/v1/templates/:id/apply")
	log.Println("  GET    /api/v1/templates/:id/performance")
	log.Println("  POST   /api/v1/templates/update-from-feedback")
}
