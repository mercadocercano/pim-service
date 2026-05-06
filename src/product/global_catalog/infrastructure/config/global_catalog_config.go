package config

import (
	"database/sql"

	"saas-mt-pim-service/src/product/global_catalog/application/usecase"
	"saas-mt-pim-service/src/product/global_catalog/infrastructure/controller"
	"saas-mt-pim-service/src/product/global_catalog/infrastructure/criteria"
	"saas-mt-pim-service/src/product/global_catalog/infrastructure/persistence"
)

// GlobalCatalogConfig contiene la configuración del módulo global catalog
type GlobalCatalogConfig struct {
	DB                                      *sql.DB
	GlobalCatalogController                 *controller.GlobalCatalogController
	ProductRequestController                *controller.ProductRequestController
	globalProductRepository                 *persistence.PostgresGlobalProductRepository
	productRequestRepository                *persistence.PostgresProductRequestRepository
	createGlobalProductUseCase              *usecase.CreateGlobalProduct
	searchByEANUseCase                      *usecase.SearchByEAN
	listGlobalProductsUseCase               *usecase.ListGlobalProducts
	listGlobalProductsByCriteriaUseCase     *usecase.ListGlobalProductsByCriteriaUseCase
	getGlobalProductByIDUseCase             *usecase.GetGlobalProductByID
	updateGlobalProductByIDUseCase          *usecase.UpdateGlobalProductByID
	getBusinessTypeFacetsUseCase            *usecase.GetBusinessTypeFacets
	listProductsNeedingEnrichmentUseCase    *usecase.ListProductsNeedingEnrichment
	getGlobalProductsByIDsUseCase           *usecase.GetGlobalProductsByIDs
	createProductRequestUseCase             *usecase.CreateProductRequestUseCase
	listProductRequestsUseCase              *usecase.ListProductRequestsUseCase
	resolveProductRequestUseCase            *usecase.ResolveProductRequestUseCase
	criteriaBuilder                         *criteria.GlobalProductCriteriaBuilder
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
	c.productRequestRepository = persistence.NewPostgresProductRequestRepository(c.DB)
}

// initializeUseCases inicializa los casos de uso
func (c *GlobalCatalogConfig) initializeUseCases() {
	c.createGlobalProductUseCase = usecase.NewCreateGlobalProduct(c.globalProductRepository)
	c.searchByEANUseCase = usecase.NewSearchByEAN(c.globalProductRepository)
	c.listGlobalProductsUseCase = usecase.NewListGlobalProducts(c.globalProductRepository)
	c.listGlobalProductsByCriteriaUseCase = usecase.NewListGlobalProductsByCriteriaUseCase(c.globalProductRepository)
	c.getGlobalProductByIDUseCase = usecase.NewGetGlobalProductByID(c.globalProductRepository)
	c.updateGlobalProductByIDUseCase = usecase.NewUpdateGlobalProductByID(c.globalProductRepository)
	c.getBusinessTypeFacetsUseCase = usecase.NewGetBusinessTypeFacets(c.globalProductRepository)
	c.listProductsNeedingEnrichmentUseCase = usecase.NewListProductsNeedingEnrichment(c.globalProductRepository)
	c.getGlobalProductsByIDsUseCase = usecase.NewGetGlobalProductsByIDs(c.globalProductRepository)
	c.createProductRequestUseCase = usecase.NewCreateProductRequestUseCase(c.productRequestRepository)
	c.listProductRequestsUseCase = usecase.NewListProductRequestsUseCase(c.productRequestRepository)
	c.resolveProductRequestUseCase = usecase.NewResolveProductRequestUseCase(c.productRequestRepository)
}

// initializeControllers inicializa los controladores
func (c *GlobalCatalogConfig) initializeControllers() {
	c.criteriaBuilder = criteria.NewGlobalProductCriteriaBuilder()
	c.GlobalCatalogController = controller.NewGlobalCatalogController(
		c.createGlobalProductUseCase,
		c.searchByEANUseCase,
		c.listGlobalProductsUseCase,
		c.listGlobalProductsByCriteriaUseCase,
		c.getGlobalProductByIDUseCase,
		c.updateGlobalProductByIDUseCase,
		c.getBusinessTypeFacetsUseCase,
		c.listProductsNeedingEnrichmentUseCase,
		c.getGlobalProductsByIDsUseCase,
		c.criteriaBuilder,
	)
	c.ProductRequestController = controller.NewProductRequestController(
		c.createProductRequestUseCase,
		c.listProductRequestsUseCase,
		c.resolveProductRequestUseCase,
	)
}

// GetGlobalCatalogController retorna el controlador del catálogo global
func (c *GlobalCatalogConfig) GetGlobalCatalogController() *controller.GlobalCatalogController {
	return c.GlobalCatalogController
}
