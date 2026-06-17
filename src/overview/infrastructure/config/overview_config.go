package config

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"saas-mt-pim-service/src/overview/application/usecase"
	"saas-mt-pim-service/src/overview/infrastructure/controller"

	// Repositorios
	attributeRepository "saas-mt-pim-service/src/attribute/infrastructure/persistence/repository"
	brandRepository "saas-mt-pim-service/src/brand/infrastructure/persistence/repository"
	categoryPersistence "saas-mt-pim-service/src/category/infrastructure/persistence"
	globalProductPersistence "saas-mt-pim-service/src/product/global_catalog/infrastructure/persistence"
)

// OverviewConfig representa la configuración del módulo overview
type OverviewConfig struct {
	db         *sql.DB
	controller *controller.OverviewHandler
}

// NewOverviewConfig crea una nueva configuración del módulo overview
func NewOverviewConfig(db *sql.DB) *OverviewConfig {
	// Instanciar repositorios reales
	marketplaceCategoryRepo := categoryPersistence.NewMarketplaceCategoryPostgresRepository(db)
	marketplaceBrandRepo := brandRepository.NewMarketplacebrandPostgresRepository(db)
	marketplaceAttributeRepo := attributeRepository.NewMarketplaceAttributePostgresRepository(db)
	globalProductRepo := globalProductPersistence.NewPostgresGlobalProductRepository(db)

	// Crear el caso de uso con repositorios reales
	getOverviewUseCase := usecase.NewGetMarketplaceOverviewUseCase(
		marketplaceCategoryRepo,
		marketplaceBrandRepo,
		marketplaceAttributeRepo,
		globalProductRepo,
	)

	// Crear controlador
	overviewController := controller.NewOverviewHandler(getOverviewUseCase)

	return &OverviewConfig{
		db:         db,
		controller: overviewController,
	}
}

// GetController devuelve el controlador del módulo
func (c *OverviewConfig) GetController() *controller.OverviewHandler {
	return c.controller
}

// SetupOverviewModule configura el módulo overview
func SetupOverviewModule(router *gin.RouterGroup, db *sql.DB) {
	config := NewOverviewConfig(db)
	controller := config.GetController()
	controller.RegisterRoutes(router)
}
