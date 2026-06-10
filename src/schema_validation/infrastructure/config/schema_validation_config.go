package config

import (
	"database/sql"
	"time"

	"saas-mt-pim-service/src/schema_validation/application/usecase"
	"saas-mt-pim-service/src/schema_validation/domain/service"
	"saas-mt-pim-service/src/schema_validation/infrastructure/adapter"
	"saas-mt-pim-service/src/schema_validation/infrastructure/cache"
	"saas-mt-pim-service/src/schema_validation/infrastructure/controller"

	"github.com/gin-gonic/gin"
)

// SchemaValidationConfig contiene la configuración del módulo
type SchemaValidationConfig struct {
	db                         *sql.DB
	csvAnalyzer                *service.CSVAnalyzerService
	jsonAnalyzer               *service.JSONAnalyzerService
	excelAnalyzer              *service.ExcelAnalyzerService
	validationCache            *cache.MemoryValidationCache
	validateCSVSchemaUseCase   *usecase.ValidateCSVSchemaUseCase
	validateJSONSchemaUseCase  *usecase.ValidateJSONSchemaUseCase
	validateExcelSchemaUseCase *usecase.ValidateExcelSchemaUseCase
	schemaValidationController *controller.SchemaValidationController
}

// NewSchemaValidationConfig crea una nueva configuración del módulo
func NewSchemaValidationConfig(db *sql.DB) *SchemaValidationConfig {
	csvAnalyzer := service.NewCSVAnalyzerService()
	jsonAnalyzer := service.NewJSONAnalyzerService()
	excelAnalyzer := service.NewExcelAnalyzerService(csvAnalyzer)
	categoryDeducer := service.NewCategoryDeductionService()
	categoryAdapter := adapter.NewPimCategoryAdapter(db)

	validationCache := cache.NewMemoryValidationCache(5 * time.Minute)

	validateCSVSchemaUseCase := usecase.NewValidateCSVSchemaUseCase(
		csvAnalyzer,
		categoryDeducer,
		categoryAdapter,
		validationCache,
	)

	validateJSONSchemaUseCase := usecase.NewValidateJSONSchemaUseCase(
		jsonAnalyzer,
		validationCache,
	)

	validateExcelSchemaUseCase := usecase.NewValidateExcelSchemaUseCase(
		excelAnalyzer,
		validateCSVSchemaUseCase,
		validationCache,
	)

	brandDeducer := service.NewBrandDeductionService()
	brandAdapter := adapter.NewPimBrandAdapter(db)

	importFromValidationUseCase := usecase.NewImportFromValidationUseCase(
		validationCache,
		brandDeducer,
		categoryAdapter.GetCategoryNames,
		brandAdapter.GetBrandNames,
	)

	schemaValidationController := controller.NewSchemaValidationController(
		validateCSVSchemaUseCase,
		validateJSONSchemaUseCase,
		validateExcelSchemaUseCase,
		importFromValidationUseCase,
	)

	return &SchemaValidationConfig{
		db:                         db,
		csvAnalyzer:                csvAnalyzer,
		jsonAnalyzer:               jsonAnalyzer,
		excelAnalyzer:              excelAnalyzer,
		validationCache:            validationCache,
		validateCSVSchemaUseCase:   validateCSVSchemaUseCase,
		validateJSONSchemaUseCase:  validateJSONSchemaUseCase,
		validateExcelSchemaUseCase: validateExcelSchemaUseCase,
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
