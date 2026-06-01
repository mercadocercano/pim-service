package config

import (
	"database/sql"

	"saas-mt-pim-service/src/attribute/application/usecase"
	"saas-mt-pim-service/src/attribute/infrastructure/controller"
	"saas-mt-pim-service/src/attribute/infrastructure/criteria"
	"saas-mt-pim-service/src/attribute/infrastructure/persistence/repository"

	"github.com/gin-gonic/gin"
)

// AttributeModuleConfig contiene todas las dependencias del módulo Attribute
type AttributeModuleConfig struct {
	// Repositories
	MarketplaceAttributeRepository *repository.MarketplaceAttributePostgresRepository
	AttributeRepository            *repository.AttributePostgresRepository
	AttributeValueRepository       *repository.AttributeValuePostgresRepository

	// Use Cases - Marketplace
	CreateMarketplaceAttributeUseCase          *usecase.CreateMarketplaceAttributeUseCase
	ListMarketplaceAttributesUseCase           *usecase.ListMarketplaceAttributesUseCase
	ListMarketplaceAttributesByCriteriaUseCase *usecase.ListMarketplaceAttributesByCriteriaUseCase
	GetMarketplaceAttributeByIDUseCase         *usecase.GetMarketplaceAttributeByIDUseCase
	UpdateMarketplaceAttributeUseCase          *usecase.UpdateMarketplaceAttributeUseCase
	DeleteMarketplaceAttributeUseCase          *usecase.DeleteMarketplaceAttributeUseCase

	// Use Cases - Attribute Values
	ListAttributeValuesUseCase   *usecase.ListAttributeValuesUseCase
	CreateAttributeValueUseCase  *usecase.CreateAttributeValueUseCase
	UpdateAttributeValueUseCase  *usecase.UpdateAttributeValueUseCase
	DeleteAttributeValueUseCase  *usecase.DeleteAttributeValueUseCase

	// Use Cases - Tenant
	CreateAttributeUseCase  *usecase.CreateAttributeUseCase
	ListAttributesUseCase   *usecase.ListAttributesUseCase
	GetAttributeByIDUseCase *usecase.GetAttributeByIDUseCase
	UpdateAttributeUseCase  *usecase.UpdateAttributeUseCase
	DeleteAttributeUseCase  *usecase.DeleteAttributeUseCase

	// Infrastructure
	MarketplaceAttributeCriteriaBuilder *criteria.MarketplaceAttributeCriteriaBuilder

	// Controllers
	MarketplaceAttributeHandler *controller.MarketplaceAttributeHandler
	AttributeHandler            *controller.AttributeHandler
}

// NewAttributeModuleConfig crea e inicializa todas las dependencias del módulo Attribute
func NewAttributeModuleConfig(db *sql.DB) *AttributeModuleConfig {
	// Repositories
	marketplaceAttributeRepo := repository.NewMarketplaceAttributePostgresRepository(db)
	attributeRepo := repository.NewAttributePostgresRepository(db)
	attributeValueRepo := repository.NewAttributeValuePostgresRepository(db)

	// Infrastructure
	marketplaceAttributeCriteriaBuilder := criteria.NewMarketplaceAttributeCriteriaBuilder()

	// Use Cases para Marketplace Attributes
	createMarketplaceAttributeUseCase := usecase.NewCreateMarketplaceAttributeUseCase(marketplaceAttributeRepo)
	listMarketplaceAttributesUseCase := usecase.NewListMarketplaceAttributesUseCase(marketplaceAttributeRepo)
	listMarketplaceAttributesByCriteriaUseCase := usecase.NewListMarketplaceAttributesByCriteriaUseCase(marketplaceAttributeRepo)
	getMarketplaceAttributeByIDUseCase := usecase.NewGetMarketplaceAttributeByIDUseCase(marketplaceAttributeRepo)
	updateMarketplaceAttributeUseCase := usecase.NewUpdateMarketplaceAttributeUseCase(marketplaceAttributeRepo)
	deleteMarketplaceAttributeUseCase := usecase.NewDeleteMarketplaceAttributeUseCase(marketplaceAttributeRepo)

	// Use Cases para Attribute Values
	listAttributeValuesUseCase := usecase.NewListAttributeValuesUseCase(attributeValueRepo)
	createAttributeValueUseCase := usecase.NewCreateAttributeValueUseCase(attributeValueRepo, marketplaceAttributeRepo)
	updateAttributeValueUseCase := usecase.NewUpdateAttributeValueUseCase(attributeValueRepo)
	deleteAttributeValueUseCase := usecase.NewDeleteAttributeValueUseCase(attributeValueRepo)

	// Use Cases para Tenant Attributes
	createAttributeUseCase := usecase.NewCreateAttributeUseCase(attributeRepo)
	listAttributesUseCase := usecase.NewListAttributesUseCase(attributeRepo)
	getAttributeByIDUseCase := usecase.NewGetAttributeByIDUseCase(attributeRepo)
	updateAttributeUseCase := usecase.NewUpdateAttributeUseCase(attributeRepo)
	deleteAttributeUseCase := usecase.NewDeleteAttributeUseCase(attributeRepo)

	// Controllers
	marketplaceAttributeHandler := controller.NewMarketplaceAttributeHandler(
		createMarketplaceAttributeUseCase,
		listMarketplaceAttributesUseCase,
		listMarketplaceAttributesByCriteriaUseCase,
		getMarketplaceAttributeByIDUseCase,
		updateMarketplaceAttributeUseCase,
		deleteMarketplaceAttributeUseCase,
		marketplaceAttributeCriteriaBuilder,
	).WithValueUseCases(
		listAttributeValuesUseCase,
		createAttributeValueUseCase,
		updateAttributeValueUseCase,
		deleteAttributeValueUseCase,
	)

	attributeHandler := controller.NewAttributeHandler(
		createAttributeUseCase,
		listAttributesUseCase,
		getAttributeByIDUseCase,
		updateAttributeUseCase,
		deleteAttributeUseCase,
	)

	return &AttributeModuleConfig{
		MarketplaceAttributeRepository:             marketplaceAttributeRepo,
		AttributeRepository:                        attributeRepo,
		AttributeValueRepository:                   attributeValueRepo,
		CreateMarketplaceAttributeUseCase:          createMarketplaceAttributeUseCase,
		ListMarketplaceAttributesUseCase:           listMarketplaceAttributesUseCase,
		ListMarketplaceAttributesByCriteriaUseCase: listMarketplaceAttributesByCriteriaUseCase,
		GetMarketplaceAttributeByIDUseCase:         getMarketplaceAttributeByIDUseCase,
		UpdateMarketplaceAttributeUseCase:          updateMarketplaceAttributeUseCase,
		DeleteMarketplaceAttributeUseCase:          deleteMarketplaceAttributeUseCase,
		ListAttributeValuesUseCase:                 listAttributeValuesUseCase,
		CreateAttributeValueUseCase:                createAttributeValueUseCase,
		UpdateAttributeValueUseCase:                updateAttributeValueUseCase,
		DeleteAttributeValueUseCase:                deleteAttributeValueUseCase,
		CreateAttributeUseCase:                     createAttributeUseCase,
		ListAttributesUseCase:                      listAttributesUseCase,
		GetAttributeByIDUseCase:                    getAttributeByIDUseCase,
		UpdateAttributeUseCase:                     updateAttributeUseCase,
		DeleteAttributeUseCase:                     deleteAttributeUseCase,
		MarketplaceAttributeCriteriaBuilder:        marketplaceAttributeCriteriaBuilder,
		MarketplaceAttributeHandler:                marketplaceAttributeHandler,
		AttributeHandler:                           attributeHandler,
	}
}

// GetMarketplaceAttributeHandler retorna el controller de Marketplace Attributes
func (c *AttributeModuleConfig) GetMarketplaceAttributeHandler() *controller.MarketplaceAttributeHandler {
	return c.MarketplaceAttributeHandler
}

// GetAttributeHandler retorna el controller de Tenant Attributes
func (c *AttributeModuleConfig) GetAttributeHandler() *controller.AttributeHandler {
	return c.AttributeHandler
}

// SetupAttributeModule configura el módulo de atributos y sus dependencias
func SetupAttributeModule(router *gin.RouterGroup, db *sql.DB) {
	attributeConfig := NewAttributeModuleConfig(db)

	marketplaceAttributeHandler := attributeConfig.GetMarketplaceAttributeHandler()
	marketplaceAttributeHandler.RegisterRoutes(router)

	attributeHandler := attributeConfig.GetAttributeHandler()
	attributeHandler.RegisterRoutes(router)
}
