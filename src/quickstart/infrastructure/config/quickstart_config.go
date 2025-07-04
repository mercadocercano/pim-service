package config

import (
	"database/sql"

	businessTypeRepository "pim/src/businesstype/infrastructure/persistence/repository"
	businessTypeUsecase "pim/src/businesstype/application/usecase"
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
	QuickstartWizardService            *service.QuickstartWizardService
	GetBusinessTypesUseCase            *usecase.GetBusinessTypesUseCase
	GetCategoriesByBusinessTypeUseCase *usecase.GetCategoriesByBusinessTypeUseCase
	GetAttributesByBusinessTypeUseCase *usecase.GetAttributesByBusinessTypeUseCase
	GetVariantsByBusinessTypeUseCase   *usecase.GetVariantsByBusinessTypeUseCase
	GetProductsByBusinessTypeUseCase   *usecase.GetProductsByBusinessTypeUseCase
	GetBrandsByBusinessTypeUseCase     *usecase.GetBrandsByBusinessTypeUseCase
	SetupTenantUseCase                 *usecase.SetupTenantUseCase
	QuickstartHandler                  *controller.QuickstartHandler
	WizardHandler                      *controller.WizardHandler
	SimpleWizardHandler                *controller.SimpleWizardHandler
}

// NewQuickstartModuleConfig crea una nueva configuración del módulo quickstart
func NewQuickstartModuleConfig(db *sql.DB, dataLoader port.YamlDataLoader, productService port.ProductService) *QuickstartModuleConfig {
	// Crear repositorios para wizard service
	businessTypeRepo := businessTypeRepository.NewBusinessTypePostgresRepository(db)
	// TODO: Agregar template repository cuando esté implementado
	// businessTypeTemplateRepo := businessTypeRepository.NewBusinessTypeTemplatePostgresRepository(db)
	// TODO: Agregar tenant setup repository cuando esté implementado
	// tenantSetupRepo := quickstartRepository.NewTenantQuickstartHistoryPostgresRepository(db)

	// Crear servicios de dominio
	quickstartDomainService := service.NewQuickstartService(dataLoader)
	// TODO: Crear wizard service cuando los repositorios estén listos
	// quickstartWizardService := service.NewQuickstartWizardService(businessTypeRepo, businessTypeTemplateRepo, tenantSetupRepo)
	
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

	// Crear handler simplificado del wizard
	listBusinessTypesUseCase := businessTypeUsecase.NewListBusinessTypesUseCase(businessTypeRepo)
	simpleWizardHandler := controller.NewSimpleWizardHandler(listBusinessTypesUseCase)

	// TODO: Crear wizard handler cuando los repositorios estén listos
	// wizardHandler := controller.NewWizardHandler(
	//     startWizardUseCase,
	//     getWizardStatusUseCase,
	//     updateWizardStepUseCase,
	//     getTemplateDataUseCase,
	// )

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
		SimpleWizardHandler:                simpleWizardHandler,
		// WizardHandler:                      wizardHandler,
	}
}

// GetQuickstartHandler devuelve el handler del quickstart
func (cfg *QuickstartModuleConfig) GetQuickstartHandler() *controller.QuickstartHandler {
	return cfg.QuickstartHandler
}

// GetWizardHandler devuelve el handler del wizard
func (cfg *QuickstartModuleConfig) GetWizardHandler() *controller.WizardHandler {
	return cfg.WizardHandler
}

// GetSimpleWizardHandler devuelve el handler simplificado del wizard
func (cfg *QuickstartModuleConfig) GetSimpleWizardHandler() *controller.SimpleWizardHandler {
	return cfg.SimpleWizardHandler
}

// NewYAMLDataLoader crea una nueva instancia del cargador de datos YAML
func NewYAMLDataLoader(dataPath string) port.YamlDataLoader {
	return loader.NewFileYamlDataLoader(dataPath)
}
