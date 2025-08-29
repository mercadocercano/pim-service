package config

import (
	"database/sql"
	"time"

	"saas-mt-pim-service/src/schema_validation/application/usecase"
	"saas-mt-pim-service/src/schema_validation/domain/service"
	"saas-mt-pim-service/src/schema_validation/infrastructure/cache"
	"saas-mt-pim-service/src/schema_validation/infrastructure/controller"

	"github.com/gin-gonic/gin"
)

// SchemaValidationConfig contiene la configuración del módulo
type SchemaValidationConfig struct {
	db                         *sql.DB
	csvAnalyzer                *service.CSVAnalyzerService
	jsonAnalyzer               *service.JSONAnalyzerService
	validationCache            *cache.MemoryValidationCache
	validateCSVSchemaUseCase   *usecase.ValidateCSVSchemaUseCase
	validateJSONSchemaUseCase  *usecase.ValidateJSONSchemaUseCase
	schemaValidationController *controller.SchemaValidationController
}

// NewSchemaValidationConfig crea una nueva configuración del módulo
func NewSchemaValidationConfig(db *sql.DB) *SchemaValidationConfig {
	// Crear servicios
	csvAnalyzer := service.NewCSVAnalyzerService()
	jsonAnalyzer := service.NewJSONAnalyzerService()
	
	// Crear cache con limpieza cada 5 minutos
	validationCache := cache.NewMemoryValidationCache(5 * time.Minute)
	
	// Crear casos de uso
	validateCSVSchemaUseCase := usecase.NewValidateCSVSchemaUseCase(
		csvAnalyzer,
		validationCache,
	)
	
	validateJSONSchemaUseCase := usecase.NewValidateJSONSchemaUseCase(
		jsonAnalyzer,
		validationCache,
	)
	
	// Crear controller
	schemaValidationController := controller.NewSchemaValidationController(
		validateCSVSchemaUseCase,
		validateJSONSchemaUseCase,
	)
	
	return &SchemaValidationConfig{
		db:                         db,
		csvAnalyzer:                csvAnalyzer,
		jsonAnalyzer:               jsonAnalyzer,
		validationCache:            validationCache,
		validateCSVSchemaUseCase:   validateCSVSchemaUseCase,
		validateJSONSchemaUseCase:  validateJSONSchemaUseCase,
		schemaValidationController: schemaValidationController,
	}
}

// SetupSchemaValidationModule configura el módulo en el router
func SetupSchemaValidationModule(router *gin.RouterGroup, db *sql.DB) {
	config := NewSchemaValidationConfig(db)
	config.schemaValidationController.RegisterRoutes(router)
}

// GetController retorna el controller del módulo
func (c *SchemaValidationConfig) GetController() *controller.SchemaValidationController {
	return c.schemaValidationController
}