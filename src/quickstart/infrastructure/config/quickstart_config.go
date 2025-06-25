package config

import (
	"database/sql"

	"pim/src/quickstart/application/usecase"
	"pim/src/quickstart/domain/port"
	"pim/src/quickstart/domain/service"
	"pim/src/quickstart/infrastructure/controller"
	"pim/src/quickstart/infrastructure/loader"
	infrastructureService "pim/src/quickstart/infrastructure/service"
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
func NewQuickstartModuleConfig(db *sql.DB, dataLoader port.YamlDataLoader, productService port.ProductService) *QuickstartModuleConfig {
	// Crear servicios de dominio
	quickstartDomainService := service.NewQuickstartService(dataLoader)
	categoryService := infrastructureService.NewCategoryService(db)
	attributeService := infrastructureService.NewAttributeService(db)
	categoryAttributeService := infrastructureService.NewCategoryAttributeService(db)

	// Crear casos de uso
	getBusinessTypesUseCase := usecase.NewGetBusinessTypesUseCase(quickstartDomainService)
	getCategoriesByBusinessTypeUseCase := usecase.NewGetCategoriesByBusinessTypeUseCase(quickstartDomainService)
	getAttributesByBusinessTypeUseCase := usecase.NewGetAttributesByBusinessTypeUseCase(quickstartDomainService)
	getVariantsByBusinessTypeUseCase := usecase.NewGetVariantsByBusinessTypeUseCase(quickstartDomainService)
	getProductsByBusinessTypeUseCase := usecase.NewGetProductsByBusinessTypeUseCase(quickstartDomainService)
	getBrandsByBusinessTypeUseCase := usecase.NewGetBrandsByBusinessTypeUseCase(quickstartDomainService)

	// Para el setup tenant, usar servicios reales donde sea posible
	setupTenantUseCase := usecase.NewSetupTenantUseCase(
		quickstartDomainService,
		nil,                      // historyRepo - temporal nil (necesita implementación)
		categoryService,          // categoryService - ahora implementado
		attributeService,         // attributeService - ahora implementado
		categoryAttributeService, // categoryAttributeService - ahora implementado
		nil,                      // variantService - temporal nil (necesita implementación)
		productService,           // productService - ahora implementado
	)

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
		QuickstartService:                  quickstartDomainService,
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
