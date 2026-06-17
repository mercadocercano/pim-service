package config

import (
	"database/sql"

	"saas-mt-pim-service/src/category_attribute/application/usecase"
	"saas-mt-pim-service/src/category_attribute/domain/port"
	"saas-mt-pim-service/src/category_attribute/infrastructure/controller"
	"saas-mt-pim-service/src/category_attribute/infrastructure/persistence/repository"

	"github.com/gin-gonic/gin"
)

// SetupCategoryAttributeModule configura el módulo de atributos de categoría y sus dependencias
func SetupCategoryAttributeModule(router *gin.RouterGroup, db *sql.DB) {
	categoryAttrRepo := repository.NewCategoryAttributePostgresRepository(db)

	createCategoryAttrUC := usecase.NewCreateCategoryAttributeUseCase(categoryAttrRepo)
	updateCategoryAttrUC := usecase.NewUpdateCategoryAttributeUseCase(categoryAttrRepo)
	deleteCategoryAttrUC := usecase.NewDeleteCategoryAttributeUseCase(categoryAttrRepo)
	getCategoryAttrsUC := usecase.NewGetCategoryAttributesUseCase(categoryAttrRepo)
	getDetailedCategoryAttrsUC := usecase.NewGetDetailedCategoryAttributesUseCase(categoryAttrRepo)
	listCategoryAttrsByCriteriaUC := usecase.NewListCategoryAttributesByCriteriaUseCase(categoryAttrRepo)

	categoryAttrHandler := controller.NewCategoryAttributeHandler(
		createCategoryAttrUC,
		updateCategoryAttrUC,
		deleteCategoryAttrUC,
		getCategoryAttrsUC,
		getDetailedCategoryAttrsUC,
		listCategoryAttrsByCriteriaUC,
	)

	categoryAttrHandler.RegisterRoutes(router)
}

// InitializeCategoryAttributeModule inicializa el módulo de atributos de categoría y retorna el controlador
func InitializeCategoryAttributeModule(repo port.CategoryAttributeCriteriaRepository) *controller.CategoryAttributeHandler {
	createUseCase := usecase.NewCreateCategoryAttributeUseCase(repo)
	updateUseCase := usecase.NewUpdateCategoryAttributeUseCase(repo)
	deleteUseCase := usecase.NewDeleteCategoryAttributeUseCase(repo)
	getUseCase := usecase.NewGetCategoryAttributesUseCase(repo)
	getDetailedUseCase := usecase.NewGetDetailedCategoryAttributesUseCase(repo)
	listByCriteriaUseCase := usecase.NewListCategoryAttributesByCriteriaUseCase(repo)

	return controller.NewCategoryAttributeHandler(
		createUseCase,
		updateUseCase,
		deleteUseCase,
		getUseCase,
		getDetailedUseCase,
		listByCriteriaUseCase,
	)
}
