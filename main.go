package main

import (
	"database/sql"
	"log"
	"os"

	apiConfig "pim/src/api/config"
	brandConfig "pim/src/brand/infrastructure/config"
	categoryConfig "pim/src/category/infrastructure/config"
	productConfig "pim/src/product/infrastructure/config"
	quickstartConfig "pim/src/quickstart/infrastructure/config"
	sharedConfig "pim/src/shared/infrastructure/config"

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

	// API v1 grupo de rutas
	v1 := router.Group("/api/v1")

	// Configurar el módulo API (health check y documentación)
	apiCfg := apiConfig.DefaultAPIConfig()
	apiCfg.DB = db
	apiCfg.Version = "1.0.0"
	apiConfig.SetupAPIModule(router, v1, apiCfg)

	// Configurar módulos
	categoryConfig.SetupCategoryModule(v1, db)
	setupBrandModule(v1, db)
	setupProductModule(v1, db)
	setupQuickstartModule(v1, db)

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
