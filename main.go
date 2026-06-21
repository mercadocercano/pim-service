package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	sharedmetrics "github.com/hornosg/go-shared/infrastructure/metrics"
	sharedmigrate "github.com/hornosg/go-shared/migrate"
	goshpostgres "github.com/hornosg/go-shared/infrastructure/postgres"
	apiConfig "saas-mt-pim-service/src/api/config"
	brandConfig "saas-mt-pim-service/src/brand/infrastructure/config"
	businesstypeUsecase "saas-mt-pim-service/src/businesstype/application/usecase"
	businesstypeController "saas-mt-pim-service/src/businesstype/infrastructure/controller"
	businesstypeRepository "saas-mt-pim-service/src/businesstype/infrastructure/persistence/repository"
	categoryConfig "saas-mt-pim-service/src/category/infrastructure/config"
	categoryAttributeConfig "saas-mt-pim-service/src/category_attribute/infrastructure/config"
	pimlogging "saas-mt-pim-service/src/pim/infrastructure/logging"
	pimport "saas-mt-pim-service/src/pim/domain/port"
	globalCatalogConfig "saas-mt-pim-service/src/product/global_catalog/infrastructure/config"
	backfillUseCase "saas-mt-pim-service/src/product/quickstart/application/usecase"
	productConfig "saas-mt-pim-service/src/product/tenant/infrastructure/config"
	quickstartConfig "saas-mt-pim-service/src/quickstart/infrastructure/config"
	sharedConfig "saas-mt-pim-service/src/shared/infrastructure/config"
	ratelimitsetup "saas-mt-pim-service/src/shared/infrastructure/ratelimit"

	// Brand imports
	brandController "saas-mt-pim-service/src/brand/infrastructure/controller"
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

	// Schema validation imports
	schemaValidationConfig "saas-mt-pim-service/src/schema_validation/infrastructure/config"

	// Overview module imports
	overviewConfig "saas-mt-pim-service/src/overview/infrastructure/config"

	// Internal S2S module
	s2sController "saas-mt-pim-service/src/s2s/controller"
	s2sPersistence "saas-mt-pim-service/src/s2s/infrastructure/persistence"
	s2sUsecase "saas-mt-pim-service/src/s2s/usecase"

	// Global Catalog — reclassify use case (E24/ADR-005)
	reclassifyUsecase "saas-mt-pim-service/src/product/global_catalog/application/usecase"
	reclassifyPersistence "saas-mt-pim-service/src/product/global_catalog/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	"github.com/hornosg/go-shared/infrastructure/env"
	tenantmw "github.com/hornosg/go-shared/infrastructure/middleware"
	_ "github.com/lib/pq" // Driver de PostgreSQL
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

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
			"/api/v1/marketplace/products*",
			"/api/v1/marketplace/store-types*",
			"/api/v1/global-catalog*",
			"/api/v1/public/global-catalog*",
			"/api/v1/internal*",
			// S2S: autenticado via API-Key en Kong, no requiere JWT del tenant
			"/api/v1/s2s*",
			// Admin: operan sobre múltiples tenants, no validan X-Tenant-ID contra token
			"/api/v1/quickstart/templates",
			"/api/v1/quickstart/backfill-tenant-images",
			"/api/v1/quickstart/backfill-all-tenant-images",
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
	dbHost := env.Get("DB_HOST", "localhost")
	dbPort := env.Get("DB_PORT", "5432")
	dbUser := env.Get("DB_USER", "postgres")
	dbPassword := env.Get("DB_PASSWORD", "postgres")
	dbName := env.Get("DB_NAME", "pim_db")

	// Conectar a la base de datos usando el helper compartido
	log.Printf("Intentando conectar a %s:%s/%s", dbHost, dbPort, dbName)
	db, err := goshpostgres.Connect(goshpostgres.Config{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		DBName:   dbName,
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()
	goshpostgres.StartPoolMonitor(context.Background(), db, goshpostgres.MonitorOptions{Service: "pim-service", DBName: dbName})
	log.Println("Conexión a la base de datos establecida con éxito")

	// Migraciones versionadas in-app (ADR-001) — fail-fast antes de servir tráfico.
	// Reemplaza el migrador casero scripts/migrate.sh (Job pim-migrate).
	if err := sharedmigrate.RunMigrations(db, MigrationsFS, dbName); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	// API v1 grupo de rutas
	v1 := router.Group("/api/v1")

	// Configurar el módulo API (health check y documentación)
	apiCfg := apiConfig.DefaultAPIConfig()
	apiCfg.DB = db
	apiCfg.Version = "1.0.0"
	apiConfig.SetupAPIModule(router, v1, apiCfg)

	// Logger canónico PIM — ADR-001. Un único adapter para toda la flota del servicio.
	pimLogger := pimlogging.NewPIMLogger("pim-service")

	// Configurar módulos
	categoryConfig.SetupCategoryModule(v1, db)
	categoryAttributeConfig.SetupCategoryAttributeModule(v1, db)
	setupBrandModule(v1, db)
	setupMarketplaceBrandModule(v1, db)
	setupAttributeModule(v1, db)
	productCfg := setupProductModuleWithLogger(v1, db, pimLogger)
	setupQuickstartModule(v1, db, productCfg.BackfillTenantImagesUseCase, pimLogger)
	setupBusinessTypeModule(v1, db)
	setupBusinessTypeTemplateModule(v1, db)
	setupGlobalCatalogModule(v1, db)
	setupMarketplaceCategoriesModule(v1, db)
	setupMarketplaceProductsModule(v1, db)
	setupSchemaValidationModule(v1, db)
	overviewConfig.SetupOverviewModule(v1, db)
	setupInternalModuleWithLogger(v1, db, pimLogger)

	// Aquí se agregarían más módulos:
	// - Ubicaciones de Stock

	// Iniciar el servidor
	port := env.Get("PORT", "8080")
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

// setupProductModule configura el módulo Product (sin logger canónico — compatibilidad).
func setupProductModule(router *gin.RouterGroup, db *sql.DB) *productConfig.ProductConfig {
	return setupProductModuleWithLogger(router, db, nil)
}

// setupProductModuleWithLogger configura el módulo Product con logger canónico ADR-001.
func setupProductModuleWithLogger(router *gin.RouterGroup, db *sql.DB, logger pimport.PIMEventLogger) *productConfig.ProductConfig {
	log.Println("Configurando módulo Product...")

	// Crear configuración del módulo Product
	productCfg := productConfig.NewProductConfigWithLogger(db, sharedmetrics.NewPrometheusRecorder(), logger)

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

	// HITO 2: Registrar rutas de importación bulk (con rate limiting por plan, ADR-003,
	// observe-only por default; corre después del TenantValidation global).
	if productCfg.BulkImportController != nil {
		productCfg.BulkImportController.RegisterRoutes(router, ratelimitsetup.NewBulkImportMiddleware())
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
	return productCfg
}

// setupQuickstartModule configura el módulo Quickstart
func setupQuickstartModule(router *gin.RouterGroup, db *sql.DB, backfillImages *backfillUseCase.BackfillTenantImagesUseCase, logger pimport.PIMEventLogger) {
	quickstartCfg := quickstartConfig.NewQuickstartModuleConfig(db, backfillImages, logger)
	quickstartCfg.GetQuickstartHandler().RegisterRoutes(router)
	quickstartCfg.GetSimpleWizardHandler().RegisterRoutes(router)
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
		private.GET("/enrichment-queue", globalCatalogController.ListProductsNeedingEnrichment)
		private.GET("/business-types", globalCatalogController.GetDistinctBusinessTypes)
		private.GET("/by-ids", globalCatalogController.GetProductsByIDs)
	}

	// Solicitudes de productos no encontrados
	globalCatalogCfg.ProductRequestController.RegisterRoutes(router)

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

// setupMarketplaceProductsModule configura el módulo de productos del marketplace (cross-tenant)
func setupMarketplaceProductsModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("Configurando módulo Marketplace Products...")

	// Crear repositorio
	marketplaceProductRepo := categoryPersistence.NewMarketplaceProductPostgresRepository(db)

	// Crear handler
	marketplaceProductHandler := categoryController.NewMarketplaceProductHandler(marketplaceProductRepo)

	// Registrar rutas
	marketplaceProductHandler.RegisterRoutes(router)

	log.Println("Módulo Marketplace Products configurado exitosamente")
	log.Println("Rutas Marketplace Products disponibles:")
	log.Println("  GET    /api/v1/marketplace/products")
	log.Println("  GET    /api/v1/marketplace/products/by-store-type/:code")
	log.Println("  GET    /api/v1/marketplace/store-types")
}

func setupInternalModule(router *gin.RouterGroup, db *sql.DB) {
	setupInternalModuleWithLogger(router, db, nil)
}

func setupInternalModuleWithLogger(router *gin.RouterGroup, db *sql.DB, logger pimport.PIMEventLogger) {
	// S2S template use cases (ya existentes)
	repo := s2sPersistence.NewPostgresTemplateRepository(db)
	refreshUC := s2sUsecase.NewRefreshTemplateProductsUseCaseWithLogger(repo, logger)
	templateUC := s2sUsecase.NewGetTemplateStatusUseCase(repo)

	// E24/ADR-005: use case de re-clasificación de business_type
	reclassifyRepo := reclassifyPersistence.NewPostgresReclassifyRepository(db)
	reclassifyUC := reclassifyUsecase.NewReclassifyBusinessTypesUseCase(reclassifyRepo, logger)

	// ADR-007: use case de normalización de category_slug (backfill S2S)
	normalizeRepo := reclassifyPersistence.NewPostgresNormalizeCategoryRepository(db)
	normalizeUC := reclassifyUsecase.NewNormalizeCategorySlugsUseCase(normalizeRepo, logger)

	handler := s2sController.
		NewInternalHandlerWithReclassify(refreshUC, templateUC, reclassifyUC, logger).
		WithNormalizeCategoryUseCase(normalizeUC)
	handler.RegisterRoutes(router)
}
