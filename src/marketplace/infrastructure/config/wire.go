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
	marketplaceGroup := router.Group("/marketplace")
	marketplaceGroup.Use(controller.CORSMiddleware())
	marketplaceGroup.Use(controller.MarketplaceAuthMiddleware())
	marketplaceGroup.Use(controller.TenantValidationMiddleware())
	marketplaceGroup.Use(controller.RequestValidationMiddleware())

	// Rutas que requieren permisos de administrador
	adminGroup := marketplaceGroup.Group("")
	adminGroup.Use(controller.AdminOnlyMiddleware())

	// Registrar rutas de administrador
	adminGroup.POST("/categories", marketplaceCategoryHandler.CreateMarketplaceCategory)
	adminGroup.POST("/categories/validate-hierarchy", marketplaceCategoryHandler.ValidateCategoryHierarchy)
	adminGroup.POST("/sync-changes", marketplaceCategoryHandler.SyncMarketplaceChanges)

	// Rutas para tenants (sin restricción de admin)
	marketplaceGroup.GET("/taxonomy", marketplaceCategoryHandler.GetTenantTaxonomy)

	// Registrar rutas de controladores tenant
	tenantCategoryMappingHandler.RegisterRoutes(marketplaceGroup)
	tenantCustomAttributeHandler.RegisterRoutes(marketplaceGroup)
}
