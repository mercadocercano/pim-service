package config

import (
	"database/sql"

	globalCatalogPort "saas-mt-pim-service/src/product/global_catalog/domain/port"
	"saas-mt-pim-service/src/product/quickstart/application/usecase"
	"saas-mt-pim-service/src/product/quickstart/infrastructure/controller"
	"saas-mt-pim-service/src/product/tenant/domain/port"
)

// QuickstartConfig contiene todas las dependencias del módulo quickstart
type QuickstartConfig struct {
	// Use Cases
	CreateFromTemplateUseCase     *usecase.CreateFromTemplateUseCase
	ImportFromBusinessTypeUseCase *usecase.ImportFromBusinessTypeUseCase
	GetQuickstartProgressUseCase  *usecase.GetQuickstartProgressUseCase

	// Controllers
	QuickstartController *controller.QuickstartController
}

// NewQuickstartConfig crea una nueva configuración del módulo quickstart
func NewQuickstartConfig(
	db *sql.DB,
	productRepo port.ProductRepository,
	globalCatalogRepo globalCatalogPort.GlobalProductRepository,
) *QuickstartConfig {
	// Use Cases
	createFromTemplateUseCase := usecase.NewCreateFromTemplateUseCase(productRepo)
	importFromBusinessTypeUseCase := usecase.NewImportFromBusinessTypeUseCase(productRepo, globalCatalogRepo)
	getQuickstartProgressUseCase := usecase.NewGetQuickstartProgressUseCase(productRepo)

	// Controllers
	quickstartController := controller.NewQuickstartController(
		createFromTemplateUseCase,
		importFromBusinessTypeUseCase,
		getQuickstartProgressUseCase,
	)

	return &QuickstartConfig{
		CreateFromTemplateUseCase:     createFromTemplateUseCase,
		ImportFromBusinessTypeUseCase: importFromBusinessTypeUseCase,
		GetQuickstartProgressUseCase:  getQuickstartProgressUseCase,
		QuickstartController:          quickstartController,
	}
}

// RegisterRoutes registra las rutas del módulo quickstart
func (config *QuickstartConfig) RegisterRoutes(router interface{}) {
	// Se implementará cuando se integre con el router principal
}
