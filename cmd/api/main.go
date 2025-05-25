package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	brandConfig "pim/src/brand/infrastructure/config"
	categoryConfig "pim/src/category/infrastructure/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Cargar variables de entorno (opcional)
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, using environment variables instead")
	}

	// Configurar la base de datos
	dbConfig := categoryConfig.NewDatabaseConfig()
	db, err := dbConfig.Connect()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Configurar el router
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

	// Health check endpoint (público para verificación de servicios)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "up",
			"service": "pim",
		})
	})

	// API v1 grupo de rutas
	v1 := router.Group("/api/v1")

	// Health check endpoint en API v1
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "up",
			"service": "pim",
			"version": "v1",
		})
	})

	// Configurar módulos
	setupCategoryModule(v1, db)
	setupBrandModule(v1, db)

	// Iniciar el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// setupCategoryModule configura las rutas del módulo de categorías
func setupCategoryModule(v1 *gin.RouterGroup, db *sql.DB) {
	categoryConfig.SetupCategoryModule(v1, db)
}

// setupBrandModule configura las rutas del módulo de marcas
func setupBrandModule(v1 *gin.RouterGroup, db *sql.DB) {
	brandConfig.SetupBrandModule(v1, db)
}
