package config

import (
	"database/sql"

	"pim/src/product/application/mapper"
	"pim/src/product/application/usecase"
	"pim/src/product/domain/service"
	"pim/src/product/infrastructure/controller"
	"pim/src/product/infrastructure/criteria"
	"pim/src/product/infrastructure/persistence"
)

// ProductConfig contiene todas las dependencias del módulo Product
type ProductConfig struct {
	// Repositorios
	ProductRepository persistence.PostgresProductRepository

	// Servicios de dominio
	ProductDomainService *service.ProductDomainService

	// Mappers
	ProductMapper        *mapper.ProductMapper
	ProductVariantMapper *mapper.ProductVariantMapper

	// Use Cases - Productos
	CreateProductUseCase          *usecase.CreateProductUseCase
	GetProductByIDUseCase         *usecase.GetProductByIDUseCase
	UpdateProductUseCase          *usecase.UpdateProductUseCase
	DeleteProductUseCase          *usecase.DeleteProductUseCase
	ListProductsByCriteriaUseCase *usecase.ListProductsByCriteriaUseCase

	// Use Cases - Variantes
	CreateProductVariantUseCase          *usecase.CreateProductVariantUseCase
	GetProductVariantByIDUseCase         *usecase.GetProductVariantByIDUseCase
	UpdateProductVariantUseCase          *usecase.UpdateProductVariantUseCase
	DeleteProductVariantUseCase          *usecase.DeleteProductVariantUseCase
	ListProductVariantsByCriteriaUseCase *usecase.ListProductVariantsByCriteriaUseCase

	// Controllers
	ProductController        *controller.ProductController
	ProductVariantController *controller.ProductVariantController

	// Criteria Builders
	ProductCriteriaBuilder *criteria.ProductCriteriaBuilder
}

// NewProductConfig crea y configura todas las dependencias del módulo Product
func NewProductConfig(db *sql.DB) *ProductConfig {
	// Repositorios
	productRepo := persistence.NewPostgresProductRepository(db)

	// Servicios de dominio
	productDomainService := service.NewProductDomainService(productRepo)

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

	// Controllers
	productController := controller.NewProductController(
		createProductUseCase,
		getProductByIDUseCase,
		updateProductUseCase,
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

	return &ProductConfig{
		ProductRepository:                    *productRepo.(*persistence.PostgresProductRepository),
		ProductDomainService:                 productDomainService,
		ProductMapper:                        productMapper,
		ProductVariantMapper:                 productVariantMapper,
		CreateProductUseCase:                 createProductUseCase,
		GetProductByIDUseCase:                getProductByIDUseCase,
		UpdateProductUseCase:                 updateProductUseCase,
		DeleteProductUseCase:                 deleteProductUseCase,
		ListProductsByCriteriaUseCase:        listProductsByCriteriaUseCase,
		CreateProductVariantUseCase:          createProductVariantUseCase,
		GetProductVariantByIDUseCase:         getProductVariantByIDUseCase,
		UpdateProductVariantUseCase:          updateProductVariantUseCase,
		DeleteProductVariantUseCase:          deleteProductVariantUseCase,
		ListProductVariantsByCriteriaUseCase: listProductVariantsByCriteriaUseCase,
		ProductController:                    productController,
		ProductVariantController:             productVariantController,
		ProductCriteriaBuilder:               productCriteriaBuilder,
	}
}
