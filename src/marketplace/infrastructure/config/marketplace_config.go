package config

import (
	"database/sql"
	"fmt"
	"net/http"

	"pim/src/marketplace/application/request"
	"pim/src/marketplace/application/usecase"
	"pim/src/marketplace/infrastructure/controller"
	"pim/src/marketplace/infrastructure/persistence"
	"pim/src/shared/infrastructure/database"

	"github.com/gin-gonic/gin"
)

// SetupMarketplaceModule configura el módulo marketplace con MongoDB
func SetupMarketplaceModule(router *gin.RouterGroup, db *sql.DB, mongoClient *database.MongoDBClient) {
	fmt.Println("🔧 DEBUG: SetupMarketplaceModule iniciado - configurando CRUD completo")

	// Crear repositorios MongoDB
	fmt.Println("🔧 DEBUG: Creando repositorios MongoDB...")
	tenantCustomAttributeRepo := persistence.NewTenantCustomAttributeMongoRepository(mongoClient.Database)
	tenantCategoryMappingRepo := persistence.NewTenantCategoryMappingMongoRepository(mongoClient.Database)
	fmt.Println("🔧 DEBUG: Repositorios MongoDB creados exitosamente")

	// Para mantener compatibilidad, también creamos el repositorio PostgreSQL de categorías marketplace
	fmt.Println("🔧 DEBUG: Creando repositorio PostgreSQL de categorías...")
	marketplaceCategoryRepo := persistence.NewMarketplaceCategoryPostgresRepository(db)
	fmt.Println("🔧 DEBUG: Repositorio PostgreSQL creado exitosamente")

	// Crear casos de uso con repositorios MongoDB
	fmt.Println("🔧 DEBUG: Creando casos de uso...")
	extendTenantAttributesUC := usecase.NewExtendTenantAttributesUseCase(
		marketplaceCategoryRepo,
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

	mapTenantCategoryUC := usecase.NewMapTenantCategoryUseCase(
		marketplaceCategoryRepo,
		tenantCategoryMappingRepo,
	)

	getTenantTaxonomyUC := usecase.NewGetTenantTaxonomyUseCase(
		marketplaceCategoryRepo,
		tenantCategoryMappingRepo,
		tenantCustomAttributeRepo,
	)
	fmt.Println("🔧 DEBUG: Casos de uso creados exitosamente")

	// Crear controladores con casos de uso
	fmt.Println("🔧 DEBUG: Creando controladores...")
	tenantCustomAttributeHandler := controller.NewTenantCustomAttributeHandler(
		extendTenantAttributesUC,
		getTenantCustomAttributesUC,
		updateTenantCustomAttributeUC,
		deleteTenantCustomAttributeUC,
	)

	tenantCategoryMappingHandler := controller.NewTenantCategoryMappingHandler(
		mapTenantCategoryUC,
	)
	fmt.Println("🔧 DEBUG: Controladores creados exitosamente")

	// Configurar rutas marketplace
	fmt.Println("🔧 DEBUG: Configurando rutas marketplace...")
	marketplace := router.Group("/marketplace")
	{
		// Health check específico para marketplace con MongoDB
		marketplace.GET("/health", func(c *gin.Context) {
			// Verificar conexión MongoDB
			if err := mongoClient.HealthCheck(c.Request.Context()); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "error",
					"module": "marketplace",
					"error":  "MongoDB connection failed",
					"detail": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status":   "up",
				"module":   "marketplace",
				"database": "mongodb",
				"message":  "MongoDB connection successful - ready for marketplace features",
				"features": []string{
					"tenant_custom_attributes (MongoDB)",
					"tenant_category_mappings (MongoDB)",
				},
			})
		})

		// Endpoint de prueba para verificar repositorios MongoDB
		marketplace.GET("/test-mongo", func(c *gin.Context) {
			// Crear repositorios para probar
			attrRepo := persistence.NewTenantCustomAttributeMongoRepository(mongoClient.Database)
			mappingRepo := persistence.NewTenantCategoryMappingMongoRepository(mongoClient.Database)

			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "MongoDB repositories created successfully",
				"repositories": []string{
					"TenantCustomAttributeMongoRepository",
					"TenantCategoryMappingMongoRepository",
				},
				"ready_for":    "CRUD operations with ValidationRules as native JSON",
				"attr_repo":    attrRepo != nil,
				"mapping_repo": mappingRepo != nil,
			})
		})

		// Endpoint para obtener taxonomía del tenant (usando MongoDB)
		marketplace.GET("/taxonomy", func(c *gin.Context) {
			tenantID := c.GetHeader("X-Tenant-ID")
			if tenantID == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "X-Tenant-ID header is required",
				})
				return
			}

			// Crear request para el caso de uso
			req := &request.GetTenantTaxonomyRequest{
				TenantID:                tenantID,
				IncludeCustomAttributes: c.Query("include_custom_attributes") == "true",
				IncludeMarketplaceData:  c.Query("include_marketplace_data") == "true",
				IncludeInactive:         c.Query("include_inactive") == "true",
				Format:                  c.DefaultQuery("format", "tree"),
			}

			// Ejecutar caso de uso
			response, err := getTenantTaxonomyUC.Execute(c.Request.Context(), req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":  "Failed to get tenant taxonomy",
					"detail": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, response)
		})
	}

	// Registrar rutas de controladores tenant con MongoDB
	fmt.Println("🔧 DEBUG: Registrando rutas CRUD...")
	tenantGroup := marketplace.Group("/tenant")
	{
		// Rutas para atributos personalizados
		tenantGroup.POST("/custom-attributes", tenantCustomAttributeHandler.ExtendTenantAttributes)
		tenantGroup.GET("/custom-attributes", tenantCustomAttributeHandler.GetTenantCustomAttributes)
		tenantGroup.PUT("/custom-attributes/:attribute_id", tenantCustomAttributeHandler.UpdateTenantCustomAttribute)
		tenantGroup.DELETE("/custom-attributes/:attribute_id", tenantCustomAttributeHandler.DeleteTenantCustomAttribute)

		// Rutas para mapeos de categorías (usando RegisterRoutes)
		tenantCategoryMappingHandler.RegisterRoutes(tenantGroup)
	}

	// Log de rutas registradas
	fmt.Println("🔧 DEBUG: Rutas CRUD registradas exitosamente:")
	fmt.Println("  POST   /api/v1/marketplace/tenant/custom-attributes")
	fmt.Println("  GET    /api/v1/marketplace/tenant/custom-attributes")
	fmt.Println("  PUT    /api/v1/marketplace/tenant/custom-attributes/:attribute_id")
	fmt.Println("  DELETE /api/v1/marketplace/tenant/custom-attributes/:attribute_id")
	fmt.Println("🔧 DEBUG: SetupMarketplaceModule completado exitosamente")
}
