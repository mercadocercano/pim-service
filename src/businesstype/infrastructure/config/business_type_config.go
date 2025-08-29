package config

import (
	"database/sql"
	"saas-mt-pim-service/src/businesstype/application/usecase"
	"saas-mt-pim-service/src/businesstype/infrastructure/controller"
	"saas-mt-pim-service/src/businesstype/infrastructure/persistence/repository"

	"github.com/gin-gonic/gin"
)

// SetupBusinessTypeModule configura el módulo de business types
func SetupBusinessTypeModule(router *gin.RouterGroup, db *sql.DB) {
	// Crear repositorio
	businessTypeRepository := repository.NewBusinessTypePostgresRepository(db)

	// Crear casos de uso
	createUseCase := usecase.NewCreateBusinessTypeUseCase(businessTypeRepository)
	listUseCase := usecase.NewListBusinessTypesUseCase(businessTypeRepository)
	getUseCase := usecase.NewGetBusinessTypeUseCase(businessTypeRepository)
	updateUseCase := usecase.NewUpdateBusinessTypeUseCase(businessTypeRepository)

	// Crear handler
	businessTypeHandler := controller.NewBusinessTypeHandler(
		createUseCase,
		listUseCase,
		getUseCase,
		updateUseCase,
		businessTypeRepository,
	)

	// Registrar rutas
	businessTypeHandler.RegisterRoutes(router)
}
