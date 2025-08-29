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

	// Use Cases - Marketplace
	CreateMarketplaceAttributeUseCase  *usecase.CreateMarketplaceAttributeUseCase
	ListMarketplaceAttributesUseCase   *usecase.ListMarketplaceAttributesUseCase
	ListMarketplaceAttributesByCriteriaUseCase *usecase.ListMarketplaceAttributesByCriteriaUseCase
	GetMarketplaceAttributeByIDUseCase *usecase.GetMarketplaceAttributeByIDUseCase
	UpdateMarketplaceAttributeUseCase  *usecase.UpdateMarketplaceAttributeUseCase
	DeleteMarketplaceAttributeUseCase  *usecase.DeleteMarketplaceAttributeUseCase

	// Infrastructure
	MarketplaceAttributeCriteriaBuilder *criteria.MarketplaceAttributeCriteriaBuilder

	// Controllers
	MarketplaceAttributeHandler *controller.MarketplaceAttributeHandler
}

// NewAttributeModuleConfig crea e inicializa todas las dependencias del módulo Attribute
func NewAttributeModuleConfig(db *sql.DB) *AttributeModuleConfig {
	// Repositories
	marketplaceAttributeRepo := repository.NewMarketplaceAttributePostgresRepository(db)

	// Infrastructure
	marketplaceAttributeCriteriaBuilder := criteria.NewMarketplaceAttributeCriteriaBuilder()

	// Use Cases para Marketplace Attributes
	createMarketplaceAttributeUseCase := usecase.NewCreateMarketplaceAttributeUseCase(marketplaceAttributeRepo)
	listMarketplaceAttributesUseCase := usecase.NewListMarketplaceAttributesUseCase(marketplaceAttributeRepo)
	listMarketplaceAttributesByCriteriaUseCase := usecase.NewListMarketplaceAttributesByCriteriaUseCase(marketplaceAttributeRepo)
	getMarketplaceAttributeByIDUseCase := usecase.NewGetMarketplaceAttributeByIDUseCase(marketplaceAttributeRepo)
	updateMarketplaceAttributeUseCase := usecase.NewUpdateMarketplaceAttributeUseCase(marketplaceAttributeRepo)
	deleteMarketplaceAttributeUseCase := usecase.NewDeleteMarketplaceAttributeUseCase(marketplaceAttributeRepo)

	// Controllers
	marketplaceAttributeHandler := controller.NewMarketplaceAttributeHandler(
		createMarketplaceAttributeUseCase,
		listMarketplaceAttributesUseCase,
		listMarketplaceAttributesByCriteriaUseCase,
		getMarketplaceAttributeByIDUseCase,
		updateMarketplaceAttributeUseCase,
		deleteMarketplaceAttributeUseCase,
		marketplaceAttributeCriteriaBuilder,
	)

	return &AttributeModuleConfig{
		MarketplaceAttributeRepository:      marketplaceAttributeRepo,
		CreateMarketplaceAttributeUseCase:   createMarketplaceAttributeUseCase,
		ListMarketplaceAttributesUseCase:    listMarketplaceAttributesUseCase,
		ListMarketplaceAttributesByCriteriaUseCase: listMarketplaceAttributesByCriteriaUseCase,
		GetMarketplaceAttributeByIDUseCase:  getMarketplaceAttributeByIDUseCase,
		UpdateMarketplaceAttributeUseCase:   updateMarketplaceAttributeUseCase,
		DeleteMarketplaceAttributeUseCase:   deleteMarketplaceAttributeUseCase,
		MarketplaceAttributeCriteriaBuilder: marketplaceAttributeCriteriaBuilder,
		MarketplaceAttributeHandler:         marketplaceAttributeHandler,
	}
}

// GetMarketplaceAttributeHandler retorna el controller de Marketplace Attributes
func (c *AttributeModuleConfig) GetMarketplaceAttributeHandler() *controller.MarketplaceAttributeHandler {
	return c.MarketplaceAttributeHandler
}

// SetupAttributeModule configura el módulo de atributos y sus dependencias
func SetupAttributeModule(router *gin.RouterGroup, db *sql.DB) {
	// Crear configuración del módulo
	attributeConfig := NewAttributeModuleConfig(db)

	// Obtener el controller
	marketplaceAttributeHandler := attributeConfig.GetMarketplaceAttributeHandler()

	// Registrar las rutas de marketplace attributes
	marketplaceAttributeHandler.RegisterRoutes(router)
}
