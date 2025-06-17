package config

import (
	"database/sql"

	"pim/src/global_catalog/application/usecase"
	"pim/src/global_catalog/infrastructure/controller"
	"pim/src/global_catalog/infrastructure/persistence"
)

// GlobalCatalogConfig contiene la configuración del módulo global catalog
type GlobalCatalogConfig struct {
	DB                         *sql.DB
	GlobalCatalogController    *controller.GlobalCatalogController
	globalProductRepository    *persistence.PostgresGlobalProductRepository
	createGlobalProductUseCase *usecase.CreateGlobalProduct
	searchByEANUseCase         *usecase.SearchByEAN
	listGlobalProductsUseCase  *usecase.ListGlobalProducts
}

// NewGlobalCatalogConfig crea una nueva configuración del módulo
func NewGlobalCatalogConfig(db *sql.DB) *GlobalCatalogConfig {
	config := &GlobalCatalogConfig{
		DB: db,
	}

	// Inicializar dependencias
	config.initializeRepositories()
	config.initializeUseCases()
	config.initializeControllers()

	return config
}

// initializeRepositories inicializa los repositorios
func (c *GlobalCatalogConfig) initializeRepositories() {
	c.globalProductRepository = persistence.NewPostgresGlobalProductRepository(c.DB).(*persistence.PostgresGlobalProductRepository)
}

// initializeUseCases inicializa los casos de uso
func (c *GlobalCatalogConfig) initializeUseCases() {
	c.createGlobalProductUseCase = usecase.NewCreateGlobalProduct(c.globalProductRepository)
	c.searchByEANUseCase = usecase.NewSearchByEAN(c.globalProductRepository)
	c.listGlobalProductsUseCase = usecase.NewListGlobalProducts(c.globalProductRepository)
}

// initializeControllers inicializa los controladores
func (c *GlobalCatalogConfig) initializeControllers() {
	c.GlobalCatalogController = controller.NewGlobalCatalogController(
		c.createGlobalProductUseCase,
		c.searchByEANUseCase,
		c.listGlobalProductsUseCase,
	)
}

// GetGlobalCatalogController retorna el controlador del catálogo global
func (c *GlobalCatalogConfig) GetGlobalCatalogController() *controller.GlobalCatalogController {
	return c.GlobalCatalogController
}
