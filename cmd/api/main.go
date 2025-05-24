package main

import (
	"log"
	"net/http"
	"os"

	categoryConfig "pim/src/category/infrastructure/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	router := gin.Default()

	// Health check endpoint (público para verificación de servicios)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "up",
			"service": "pim",
		})
	})

	// API v1 grupo de rutas
	v1 := router.Group("/pim/api/v1")

	// Health check endpoint en API v1
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "up",
			"service": "pim",
			"version": "v1",
		})
	})

	// Configurar módulos
	categoryConfig.SetupCategoryModule(v1, db)

	// Iniciar el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
