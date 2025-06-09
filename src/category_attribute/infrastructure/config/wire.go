package config

import (
	"database/sql"

	"pim/src/category_attribute/application/usecase"
	"pim/src/category_attribute/domain/port"
	"pim/src/category_attribute/infrastructure/controller"
	"pim/src/category_attribute/infrastructure/persistence/repository"

	"log"

	"github.com/gin-gonic/gin"
)

// SetupCategoryAttributeModule configura el módulo de atributos de categoría y sus dependencias
func SetupCategoryAttributeModule(router *gin.RouterGroup, db *sql.DB) {
	log.Println("🚀🚀🚀 INICIANDO CONFIGURACIÓN DE CATEGORY ATTRIBUTE MODULE 🚀🚀🚀")

	// Repositorio
	categoryAttrRepo := repository.NewCategoryAttributePostgresRepository(db)

	// Casos de uso
	createCategoryAttrUC := usecase.NewCreateCategoryAttributeUseCase(categoryAttrRepo)
	updateCategoryAttrUC := usecase.NewUpdateCategoryAttributeUseCase(categoryAttrRepo)
	deleteCategoryAttrUC := usecase.NewDeleteCategoryAttributeUseCase(categoryAttrRepo)
	getCategoryAttrsUC := usecase.NewGetCategoryAttributesUseCase(categoryAttrRepo)
	getDetailedCategoryAttrsUC := usecase.NewGetDetailedCategoryAttributesUseCase(categoryAttrRepo)
	listCategoryAttrsByCriteriaUC := usecase.NewListCategoryAttributesByCriteriaUseCase(categoryAttrRepo)

	log.Println("✅ Casos de uso creados exitosamente")

	// Controlador HTTP
	categoryAttrHandler := controller.NewCategoryAttributeHandler(
		createCategoryAttrUC,
		updateCategoryAttrUC,
		deleteCategoryAttrUC,
		getCategoryAttrsUC,
		getDetailedCategoryAttrsUC,
		listCategoryAttrsByCriteriaUC,
	)

	log.Println("✅ Handler creado exitosamente")

	categoryAttrHandler.RegisterRoutes(router)

	log.Println("🎉🎉🎉 CATEGORY ATTRIBUTE MODULE CONFIGURADO COMPLETAMENTE 🎉🎉🎉")
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
