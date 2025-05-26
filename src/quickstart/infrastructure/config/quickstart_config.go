package config

import (
	"database/sql"

	"pim/src/quickstart/application/usecase"
	"pim/src/quickstart/domain/port"
	"pim/src/quickstart/domain/service"
	"pim/src/quickstart/infrastructure/controller"
	"pim/src/quickstart/infrastructure/loader"
)

// QuickstartModuleConfig contiene toda la configuración del módulo quickstart
type QuickstartModuleConfig struct {
	DB                                 *sql.DB
	DataLoader                         port.YamlDataLoader
	QuickstartService                  *service.QuickstartService
	GetBusinessTypesUseCase            *usecase.GetBusinessTypesUseCase
	GetCategoriesByBusinessTypeUseCase *usecase.GetCategoriesByBusinessTypeUseCase
	GetAttributesByBusinessTypeUseCase *usecase.GetAttributesByBusinessTypeUseCase
	GetVariantsByBusinessTypeUseCase   *usecase.GetVariantsByBusinessTypeUseCase
	GetProductsByBusinessTypeUseCase   *usecase.GetProductsByBusinessTypeUseCase
	GetBrandsByBusinessTypeUseCase     *usecase.GetBrandsByBusinessTypeUseCase
	SetupTenantUseCase                 *usecase.SetupTenantUseCase
	QuickstartHandler                  *controller.QuickstartHandler
}

// NewQuickstartModuleConfig crea una nueva configuración del módulo quickstart
func NewQuickstartModuleConfig(db *sql.DB, dataLoader port.YamlDataLoader) *QuickstartModuleConfig {
	// Crear servicio de dominio
	quickstartService := service.NewQuickstartService(dataLoader)

	// Crear casos de uso
	getBusinessTypesUseCase := usecase.NewGetBusinessTypesUseCase(quickstartService)
	getCategoriesByBusinessTypeUseCase := usecase.NewGetCategoriesByBusinessTypeUseCase(quickstartService)
	getAttributesByBusinessTypeUseCase := usecase.NewGetAttributesByBusinessTypeUseCase(quickstartService)
	getVariantsByBusinessTypeUseCase := usecase.NewGetVariantsByBusinessTypeUseCase(quickstartService)
	getProductsByBusinessTypeUseCase := usecase.NewGetProductsByBusinessTypeUseCase(quickstartService)
	getBrandsByBusinessTypeUseCase := usecase.NewGetBrandsByBusinessTypeUseCase(quickstartService)

	// Para el setup tenant, necesitamos implementaciones mock por ahora
	// TODO: Implementar repositorios y servicios reales
	setupTenantUseCase := &usecase.SetupTenantUseCase{}

	// Crear handler
	quickstartHandler := controller.NewQuickstartHandler(
		getBusinessTypesUseCase,
		getCategoriesByBusinessTypeUseCase,
		getAttributesByBusinessTypeUseCase,
		getVariantsByBusinessTypeUseCase,
		getProductsByBusinessTypeUseCase,
		getBrandsByBusinessTypeUseCase,
		setupTenantUseCase,
	)

	return &QuickstartModuleConfig{
		DB:                                 db,
		DataLoader:                         dataLoader,
		QuickstartService:                  quickstartService,
		GetBusinessTypesUseCase:            getBusinessTypesUseCase,
		GetCategoriesByBusinessTypeUseCase: getCategoriesByBusinessTypeUseCase,
		GetAttributesByBusinessTypeUseCase: getAttributesByBusinessTypeUseCase,
		GetVariantsByBusinessTypeUseCase:   getVariantsByBusinessTypeUseCase,
		GetProductsByBusinessTypeUseCase:   getProductsByBusinessTypeUseCase,
		GetBrandsByBusinessTypeUseCase:     getBrandsByBusinessTypeUseCase,
		SetupTenantUseCase:                 setupTenantUseCase,
		QuickstartHandler:                  quickstartHandler,
	}
}

// GetQuickstartHandler devuelve el handler del quickstart
func (cfg *QuickstartModuleConfig) GetQuickstartHandler() *controller.QuickstartHandler {
	return cfg.QuickstartHandler
}

// NewYAMLDataLoader crea una nueva instancia del cargador de datos YAML
func NewYAMLDataLoader(dataPath string) port.YamlDataLoader {
	return loader.NewFileYamlDataLoader(dataPath)
}
