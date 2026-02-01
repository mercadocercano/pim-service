package config

import (
	"database/sql"

	businessTypeRepository "saas-mt-pim-service/src/businesstype/infrastructure/persistence/repository"
	businessTypeUsecase "saas-mt-pim-service/src/businesstype/application/usecase"
	"saas-mt-pim-service/src/quickstart/application/usecase"
	"saas-mt-pim-service/src/quickstart/infrastructure/controller"
	quickstartRepository "saas-mt-pim-service/src/quickstart/infrastructure/persistence/repository"
)

// QuickstartModuleConfig contiene toda la configuración del módulo quickstart
type QuickstartModuleConfig struct {
	DB                                 *sql.DB
	GetBusinessTypesUseCase            *usecase.GetBusinessTypesUseCase
	GetCategoriesByBusinessTypeUseCase *usecase.GetCategoriesByBusinessTypeUseCase
	GetAttributesByBusinessTypeUseCase *usecase.GetAttributesByBusinessTypeUseCase
	GetVariantsByBusinessTypeUseCase   *usecase.GetVariantsByBusinessTypeUseCase
	GetProductsByBusinessTypeUseCase   *usecase.GetProductsByBusinessTypeUseCase
	GetBrandsByBusinessTypeUseCase     *usecase.GetBrandsByBusinessTypeUseCase
	SetupTenantUseCase                 *usecase.SetupTenantUseCase
	QuickstartHandler                  *controller.QuickstartHandler
	SimpleWizardHandler                *controller.SimpleWizardHandler
}

// NewQuickstartModuleConfig crea una nueva configuración del módulo quickstart
func NewQuickstartModuleConfig(db *sql.DB) *QuickstartModuleConfig {
	// Crear repositorio de business types
	businessTypeRepo := businessTypeRepository.NewBusinessTypePostgresRepository(db)
	
	// Crear casos de uso simplificados (sin YamlDataLoader)
	getBusinessTypesUseCase := usecase.NewGetBusinessTypesUseCase(businessTypeRepo)
	getCategoriesByBusinessTypeUseCase := usecase.NewGetCategoriesByBusinessTypeUseCase()
	getAttributesByBusinessTypeUseCase := usecase.NewGetAttributesByBusinessTypeUseCase()
	getVariantsByBusinessTypeUseCase := usecase.NewGetVariantsByBusinessTypeUseCase()
	getProductsByBusinessTypeUseCase := usecase.NewGetProductsByBusinessTypeUseCase()
	getBrandsByBusinessTypeUseCase := usecase.NewGetBrandsByBusinessTypeUseCase()
	setupTenantUseCase := usecase.NewSetupTenantUseCase()
	
	// HITO 2: Nuevos casos de uso para templates
	listTemplatesUseCase := usecase.NewListTemplatesUseCase(db)
	applyTemplateUseCase := usecase.NewApplyTemplateUseCase(db)
	
	// Crear handler principal de quickstart
	quickstartHandler := controller.NewQuickstartHandler(
		getBusinessTypesUseCase,
		getCategoriesByBusinessTypeUseCase,
		getAttributesByBusinessTypeUseCase,
		getVariantsByBusinessTypeUseCase,
		getProductsByBusinessTypeUseCase,
		getBrandsByBusinessTypeUseCase,
		setupTenantUseCase,
		listTemplatesUseCase,
		applyTemplateUseCase,
	)
	
	// Crear handler simplificado del wizard
	listBusinessTypesUseCase := businessTypeUsecase.NewListBusinessTypesUseCase(businessTypeRepo)
	historyRepo := quickstartRepository.NewQuickstartHistoryPostgresRepository(db)
	simpleWizardHandler := controller.NewSimpleWizardHandler(listBusinessTypesUseCase, db, historyRepo)

	return &QuickstartModuleConfig{
		DB:                                 db,
		GetBusinessTypesUseCase:            getBusinessTypesUseCase,
		GetCategoriesByBusinessTypeUseCase: getCategoriesByBusinessTypeUseCase,
		GetAttributesByBusinessTypeUseCase: getAttributesByBusinessTypeUseCase,
		GetVariantsByBusinessTypeUseCase:   getVariantsByBusinessTypeUseCase,
		GetProductsByBusinessTypeUseCase:   getProductsByBusinessTypeUseCase,
		GetBrandsByBusinessTypeUseCase:     getBrandsByBusinessTypeUseCase,
		SetupTenantUseCase:                 setupTenantUseCase,
		QuickstartHandler:                  quickstartHandler,
		SimpleWizardHandler:                simpleWizardHandler,
	}
}

// GetQuickstartHandler devuelve el handler principal de quickstart
func (cfg *QuickstartModuleConfig) GetQuickstartHandler() *controller.QuickstartHandler {
	return cfg.QuickstartHandler
}

// GetSimpleWizardHandler devuelve el handler simplificado del wizard
func (cfg *QuickstartModuleConfig) GetSimpleWizardHandler() *controller.SimpleWizardHandler {
	return cfg.SimpleWizardHandler
}