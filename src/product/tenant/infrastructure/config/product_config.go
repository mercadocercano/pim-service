package config

import (
	"database/sql"

	quickstartUseCase "saas-mt-pim-service/src/product/quickstart/application/usecase"
	quickstartCtrl "saas-mt-pim-service/src/product/quickstart/infrastructure/controller"
	quickstartService "saas-mt-pim-service/src/product/quickstart/infrastructure/service"
	"saas-mt-pim-service/src/product/tenant/application/mapper"
	"saas-mt-pim-service/src/product/tenant/application/usecase"
	tenantService "saas-mt-pim-service/src/product/tenant/domain/service"
	"saas-mt-pim-service/src/product/tenant/infrastructure/adapters"
	"saas-mt-pim-service/src/product/tenant/infrastructure/controller"
	"saas-mt-pim-service/src/product/tenant/infrastructure/criteria"
	"saas-mt-pim-service/src/product/tenant/infrastructure/persistence"
	categoryRepository "saas-mt-pim-service/src/category/infrastructure/persistence/repository"
	"saas-mt-pim-service/src/quickstart/domain/port"
)

// ProductConfig contiene todas las dependencias del módulo Product
type ProductConfig struct {
	// Repositorios
	ProductRepository persistence.PostgresProductRepository

	// Servicios de dominio
	ProductDomainService *tenantService.ProductDomainService
	ProductStatusService *tenantService.ProductStatusService

	// Use Cases - Productos
	CreateProductUseCase          *usecase.CreateProductUseCase
	GetProductByIDUseCase         *usecase.GetProductByIDUseCase
	UpdateProductUseCase          *usecase.UpdateProductUseCase
	UpdateProductStatusUseCase    *usecase.UpdateProductStatusUseCase
	DeleteProductUseCase          *usecase.DeleteProductUseCase
	ListProductsByCriteriaUseCase *usecase.ListProductsByCriteriaUseCase
	ImportProductsFromCSVUseCase  *usecase.ImportProductsFromCSVUseCase
	ImportProductsAsyncUseCase    *usecase.ImportProductsAsyncUseCase
	ValidateSKUsUseCase           *usecase.ValidateSKUsUseCase

	// Use Cases - Variantes
	CreateProductVariantUseCase          *usecase.CreateProductVariantUseCase
	GetProductVariantByIDUseCase         *usecase.GetProductVariantByIDUseCase
	UpdateProductVariantUseCase          *usecase.UpdateProductVariantUseCase
	DeleteProductVariantUseCase          *usecase.DeleteProductVariantUseCase
	ListProductVariantsByCriteriaUseCase *usecase.ListProductVariantsByCriteriaUseCase

	// Use Cases - Quickstart
	CreateFromTemplateUseCase     *quickstartUseCase.CreateFromTemplateUseCase
	ImportFromBusinessTypeUseCase *quickstartUseCase.ImportFromBusinessTypeUseCase
	GetQuickstartProgressUseCase  *quickstartUseCase.GetQuickstartProgressUseCase

	// Servicios - Quickstart
	QuickstartProductService port.ProductService

	// Controllers
	ProductController        *controller.ProductController
	ProductVariantController *controller.ProductVariantController
	QuickstartController     *quickstartCtrl.QuickstartController
	BulkImportController     *controller.BulkImportController  // HITO 2

	// Criteria Builders
	ProductCriteriaBuilder *criteria.ProductCriteriaBuilder
}

// NewProductConfig crea y configura todas las dependencias del módulo Product
func NewProductConfig(db *sql.DB) *ProductConfig {
	// Repositorios
	productRepo := persistence.NewPostgresProductRepository(db)

	// Servicios de dominio
	productDomainService := tenantService.NewProductDomainService(productRepo)
	productStatusService := tenantService.NewProductStatusService()

	// Mappers
	productMapper := mapper.NewProductMapper()
	variantMapper := mapper.NewProductVariantMapper()

	// Criteria Builders
	productCriteriaBuilder := criteria.NewProductCriteriaBuilder()

	// Use Cases - Productos
	createProductUseCase := usecase.NewCreateProductUseCase(
		productRepo,
		productDomainService,
		productMapper,
	)

	getProductByIDUseCase := usecase.NewGetProductByIDUseCase(
		productRepo,
		productMapper,
	)

	updateProductUseCase := usecase.NewUpdateProductUseCase(
		productRepo,
		productDomainService,
		productMapper,
	)

	deleteProductUseCase := usecase.NewDeleteProductUseCase(
		productRepo,
		productDomainService,
	)

	listProductsByCriteriaUseCase := usecase.NewListProductsByCriteriaUseCase(
		productRepo,
		productMapper,
	)

	updateProductStatusUseCase := usecase.NewUpdateProductStatusUseCase(
		productRepo,
		productStatusService,
	)

	// CSV File Importer
	// TODO: Pasar servicios reales de categoría y marca cuando estén disponibles
	productCSVFileImporter := adapters.NewProductCSVFileImporter(nil, nil)

	// Use Case de importación CSV
	importProductsFromCSVUseCase := usecase.NewImportProductsFromCSVUseCase(
		productRepo,
		productCSVFileImporter,
	)

	// Use Case de validación de SKUs
	validateSKUsUseCase := usecase.NewValidateSKUsUseCase(productRepo)

	// TODO: Agregar ImportProductsAsyncUseCase cuando las dependencias estén disponibles
	// importProductsAsyncUseCase := usecase.NewImportProductsAsyncUseCase(
	//     productRepo,
	//     fileImporter,
	//     importJobRepo,
	//     notificationSvc,
	//     fileStorage,
	// )

	// Use Cases - Variantes
	createProductVariantUseCase := usecase.NewCreateProductVariantUseCase(
		productRepo,
		variantMapper,
	)

	getProductVariantByIDUseCase := usecase.NewGetProductVariantByIDUseCase(
		productRepo,
		variantMapper,
	)

	updateProductVariantUseCase := usecase.NewUpdateProductVariantUseCase(
		productRepo,
		variantMapper,
	)

	deleteProductVariantUseCase := usecase.NewDeleteProductVariantUseCase(
		productRepo,
	)

	listProductVariantsByCriteriaUseCase := usecase.NewListProductVariantsByCriteriaUseCase(
		productRepo,
		variantMapper,
	)

	// HITO A - Use Case para buscar variante por SKU
	getVariantBySKUUseCase := usecase.NewGetVariantBySKUUseCase(productRepo, variantMapper)

	// Use Cases - Quickstart
	createFromTemplateUseCase := quickstartUseCase.NewCreateFromTemplateUseCase(productRepo)
	importFromBusinessTypeUseCase := quickstartUseCase.NewImportFromBusinessTypeUseCase(productRepo)
	getQuickstartProgressUseCase := quickstartUseCase.NewGetQuickstartProgressUseCase(productRepo)

	// Servicios - Quickstart
	quickstartProductService := quickstartService.NewQuickstartProductService(createFromTemplateUseCase)

	// Controllers
	productController := controller.NewProductController(
		createProductUseCase,
		getProductByIDUseCase,
		updateProductUseCase,
		updateProductStatusUseCase,
		deleteProductUseCase,
		listProductsByCriteriaUseCase,
		importProductsFromCSVUseCase,
		nil, // importProductsAsyncUseCase - TODO: implementar cuando las dependencias estén disponibles
		validateSKUsUseCase,
		productCriteriaBuilder,
	)

	productVariantController := controller.NewProductVariantController(
		createProductVariantUseCase,
		getProductVariantByIDUseCase,
		getVariantBySKUUseCase,
		updateProductVariantUseCase,
		deleteProductVariantUseCase,
		listProductVariantsByCriteriaUseCase,
	)

	quickstartController := quickstartCtrl.NewQuickstartController(
		createFromTemplateUseCase,
		importFromBusinessTypeUseCase,
		getQuickstartProgressUseCase,
	)

	// HITO 2: Configurar BulkImportController
	categoryRepo := categoryRepository.NewCategoryPostgresRepository(db)
	bulkImportUseCase := usecase.NewBulkImportProductsUseCase(productRepo, categoryRepo)
	bulkImportController := controller.NewBulkImportController(bulkImportUseCase)

	return &ProductConfig{
		ProductRepository:                    *productRepo.(*persistence.PostgresProductRepository),
		ProductDomainService:                 productDomainService,
		ProductStatusService:                 productStatusService,
		CreateProductUseCase:                 createProductUseCase,
		GetProductByIDUseCase:                getProductByIDUseCase,
		UpdateProductUseCase:                 updateProductUseCase,
		UpdateProductStatusUseCase:           updateProductStatusUseCase,
		DeleteProductUseCase:                 deleteProductUseCase,
		ListProductsByCriteriaUseCase:        listProductsByCriteriaUseCase,
		ImportProductsFromCSVUseCase:         importProductsFromCSVUseCase,
		ImportProductsAsyncUseCase:           nil, // TODO: implementar cuando las dependencias estén disponibles
		ValidateSKUsUseCase:                  validateSKUsUseCase,
		CreateProductVariantUseCase:          createProductVariantUseCase,
		GetProductVariantByIDUseCase:         getProductVariantByIDUseCase,
		UpdateProductVariantUseCase:          updateProductVariantUseCase,
		DeleteProductVariantUseCase:          deleteProductVariantUseCase,
		ListProductVariantsByCriteriaUseCase: listProductVariantsByCriteriaUseCase,
		CreateFromTemplateUseCase:            createFromTemplateUseCase,
		ImportFromBusinessTypeUseCase:        importFromBusinessTypeUseCase,
		GetQuickstartProgressUseCase:         getQuickstartProgressUseCase,
		QuickstartProductService:             quickstartProductService,
		ProductController:                    productController,
		ProductVariantController:             productVariantController,
		QuickstartController:                 quickstartController,
		BulkImportController:                 bulkImportController,
		ProductCriteriaBuilder:               productCriteriaBuilder,
	}
}
