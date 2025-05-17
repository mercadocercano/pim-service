package config

import (
	"database/sql"

	"pim/src/category/application/usecase"
	"pim/src/category/domain/port"
	"pim/src/category/infrastructure/controller"
	"pim/src/category/infrastructure/persistence/repository"

	"github.com/gin-gonic/gin"
)

// SetupCategoryModule configura el módulo de categorías y sus dependencias
func SetupCategoryModule(router *gin.RouterGroup, db *sql.DB) {
	// Repositorio
	categoryRepo := repository.NewCategoryPostgresRepository(db)

	// Casos de uso
	createCategoryUC := usecase.NewCreateCategoryUseCase(categoryRepo)
	updateCategoryUC := usecase.NewUpdateCategoryUseCase(categoryRepo)
	changeCategoryStatusUC := usecase.NewChangeCategoryStatusUseCase(categoryRepo)
	moveCategoryUC := usecase.NewMoveCategoryUseCase(categoryRepo)
	getCategoriesUC := usecase.NewGetCategoriesUseCase(categoryRepo)
	getCategoryByIDUC := usecase.NewGetCategoryByIDUseCase(categoryRepo)

	// Controlador HTTP
	categoryHandler := controller.NewCategoryHandler(
		createCategoryUC,
		updateCategoryUC,
		changeCategoryStatusUC,
		moveCategoryUC,
		getCategoriesUC,
		getCategoryByIDUC,
	)

	categoryHandler.RegisterRoutes(router)
}

// InitializeCategoryModule inicializa el módulo de categorías y retorna el controlador
func InitializeCategoryModule(repo port.CategoryRepository) *controller.CategoryHandler {
	createUseCase := usecase.NewCreateCategoryUseCase(repo)
	updateUseCase := usecase.NewUpdateCategoryUseCase(repo)
	changeStatusUseCase := usecase.NewChangeCategoryStatusUseCase(repo)
	moveUseCase := usecase.NewMoveCategoryUseCase(repo)
	getCategoriesUseCase := usecase.NewGetCategoriesUseCase(repo)
	getCategoryByIDUseCase := usecase.NewGetCategoryByIDUseCase(repo)

	return controller.NewCategoryHandler(
		createUseCase,
		updateUseCase,
		changeStatusUseCase,
		moveUseCase,
		getCategoriesUseCase,
		getCategoryByIDUseCase,
	)
}
