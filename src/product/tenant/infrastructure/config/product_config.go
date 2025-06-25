package config

import (
	"database/sql"

	quickstartUseCase "pim/src/product/quickstart/application/usecase"
	quickstartCtrl "pim/src/product/quickstart/infrastructure/controller"
	quickstartService "pim/src/product/quickstart/infrastructure/service"
	"pim/src/product/tenant/application/mapper"
	"pim/src/product/tenant/application/usecase"
	tenantService "pim/src/product/tenant/domain/service"
	"pim/src/product/tenant/infrastructure/controller"
	"pim/src/product/tenant/infrastructure/criteria"
	"pim/src/product/tenant/infrastructure/persistence"
	"pim/src/quickstart/domain/port"
)

// ProductConfig contiene todas las dependencias del módulo Product
type ProductConfig struct {
	// Repositorios
	ProductRepository persistence.PostgresProductRepository

	// Servicios de dominio
	ProductDomainService *tenantService.ProductDomainService
	ProductStatusService *tenantService.ProductStatusService

	// Mappers
	ProductMapper        *mapper.ProductMapper
	ProductVariantMapper *mapper.ProductVariantMapper

	// Use Cases - Productos
	CreateProductUseCase          *usecase.CreateProductUseCase
	GetProductByIDUseCase         *usecase.GetProductByIDUseCase
	UpdateProductUseCase          *usecase.UpdateProductUseCase
	UpdateProductStatusUseCase    *usecase.UpdateProductStatusUseCase
	DeleteProductUseCase          *usecase.DeleteProductUseCase
	ListProductsByCriteriaUseCase *usecase.ListProductsByCriteriaUseCase

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
	productVariantMapper := mapper.NewProductVariantMapper()

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

	// Use Cases - Variantes
	createProductVariantUseCase := usecase.NewCreateProductVariantUseCase(
		productRepo,
		productVariantMapper,
	)

	getProductVariantByIDUseCase := usecase.NewGetProductVariantByIDUseCase(
		productRepo,
		productVariantMapper,
	)

	updateProductVariantUseCase := usecase.NewUpdateProductVariantUseCase(
		productRepo,
		productVariantMapper,
	)

	deleteProductVariantUseCase := usecase.NewDeleteProductVariantUseCase(
		productRepo,
	)

	listProductVariantsByCriteriaUseCase := usecase.NewListProductVariantsByCriteriaUseCase(
		productRepo,
		productVariantMapper,
	)

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
		productCriteriaBuilder,
	)

	productVariantController := controller.NewProductVariantController(
		createProductVariantUseCase,
		getProductVariantByIDUseCase,
		updateProductVariantUseCase,
		deleteProductVariantUseCase,
		listProductVariantsByCriteriaUseCase,
	)

	quickstartController := quickstartCtrl.NewQuickstartController(
		createFromTemplateUseCase,
		importFromBusinessTypeUseCase,
		getQuickstartProgressUseCase,
	)

	return &ProductConfig{
		ProductRepository:                    *productRepo.(*persistence.PostgresProductRepository),
		ProductDomainService:                 productDomainService,
		ProductStatusService:                 productStatusService,
		ProductMapper:                        productMapper,
		ProductVariantMapper:                 productVariantMapper,
		CreateProductUseCase:                 createProductUseCase,
		GetProductByIDUseCase:                getProductByIDUseCase,
		UpdateProductUseCase:                 updateProductUseCase,
		UpdateProductStatusUseCase:           updateProductStatusUseCase,
		DeleteProductUseCase:                 deleteProductUseCase,
		ListProductsByCriteriaUseCase:        listProductsByCriteriaUseCase,
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
		ProductCriteriaBuilder:               productCriteriaBuilder,
	}
}
