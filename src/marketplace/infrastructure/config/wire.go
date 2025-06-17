package config

import (
	"database/sql"

	"pim/src/marketplace/application/usecase"
	"pim/src/marketplace/infrastructure/controller"
	"pim/src/marketplace/infrastructure/persistence"

	"github.com/gin-gonic/gin"
)

// SetupMarketplaceModulePostgres configura el módulo marketplace con PostgreSQL (legacy)
func SetupMarketplaceModulePostgres(router *gin.RouterGroup, db *sql.DB) {
	// Repositorios
	marketplaceCategoryRepo := persistence.NewMarketplaceCategoryPostgresRepository(db)
	tenantCategoryMappingRepo := persistence.NewTenantCategoryMappingPostgresRepository(db)
	tenantCustomAttributeRepo := persistence.NewTenantCustomAttributePostgresRepository(db)

	// Casos de uso
	createMarketplaceCategoryUC := usecase.NewCreateMarketplaceCategoryUseCase(marketplaceCategoryRepo)
	getAllMarketplaceCategoriesUC := usecase.NewGetAllMarketplaceCategoriesUseCase(marketplaceCategoryRepo)
	updateMarketplaceCategoryUC := usecase.NewUpdateMarketplaceCategoryUseCase(marketplaceCategoryRepo)
	getTenantTaxonomyUC := usecase.NewGetTenantTaxonomyUseCase(
		marketplaceCategoryRepo,
		tenantCategoryMappingRepo,
		tenantCustomAttributeRepo,
	)
	validateCategoryHierarchyUC := usecase.NewValidateCategoryHierarchyUseCase(marketplaceCategoryRepo)
	syncMarketplaceChangesUC := usecase.NewSyncMarketplaceChangesUseCase(
		marketplaceCategoryRepo,
		tenantCategoryMappingRepo,
		tenantCustomAttributeRepo,
	)
	mapTenantCategoryUC := usecase.NewMapTenantCategoryUseCase(
		marketplaceCategoryRepo,
		tenantCategoryMappingRepo,
	)
	extendTenantAttributesUC := usecase.NewExtendTenantAttributesUseCase(
		marketplaceCategoryRepo,
		tenantCustomAttributeRepo,
	)

	// Controladores HTTP
	marketplaceCategoryHandler := controller.NewMarketplaceCategoryHandler(
		createMarketplaceCategoryUC,
		getAllMarketplaceCategoriesUC,
		updateMarketplaceCategoryUC,
		getTenantTaxonomyUC,
		validateCategoryHierarchyUC,
		syncMarketplaceChangesUC,
	)

	tenantCategoryMappingHandler := controller.NewTenantCategoryMappingHandler(
		mapTenantCategoryUC,
	)

	// Crear casos de uso adicionales para CRUD completo
	getTenantCustomAttributesUC := usecase.NewGetTenantCustomAttributesUseCase(
		tenantCustomAttributeRepo,
	)

	updateTenantCustomAttributeUC := usecase.NewUpdateTenantCustomAttributeUseCase(
		tenantCustomAttributeRepo,
	)

	deleteTenantCustomAttributeUC := usecase.NewDeleteTenantCustomAttributeUseCase(
		tenantCustomAttributeRepo,
	)

	tenantCustomAttributeHandler := controller.NewTenantCustomAttributeHandler(
		extendTenantAttributesUC,
		getTenantCustomAttributesUC,
		updateTenantCustomAttributeUC,
		deleteTenantCustomAttributeUC,
	)

	// Aplicar middlewares
	marketplaceGroup := router.Group("")
	marketplaceGroup.Use(controller.CORSMiddleware())
	marketplaceGroup.Use(controller.MarketplaceAuthMiddleware())
	marketplaceGroup.Use(controller.TenantValidationMiddleware())
	marketplaceGroup.Use(controller.RequestValidationMiddleware())

	// Registrar rutas usando el método RegisterRoutes del controlador
	marketplaceCategoryHandler.RegisterRoutes(marketplaceGroup)
	tenantCategoryMappingHandler.RegisterRoutes(marketplaceGroup)
	tenantCustomAttributeHandler.RegisterRoutes(marketplaceGroup)
}
