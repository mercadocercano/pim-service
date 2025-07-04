package config

import (
	"database/sql"

	"pim/src/businesstype/application/usecase"
	"pim/src/businesstype/infrastructure/controller"
	"pim/src/businesstype/infrastructure/persistence/repository"
	businessTypeRepository "pim/src/businesstype/infrastructure/persistence/repository"
)

// BusinessTypeTemplateModuleConfig contiene la configuración del módulo de templates
type BusinessTypeTemplateModuleConfig struct {
	DB                      *sql.DB
	TemplateRepository      *repository.BusinessTypeTemplatePostgresRepository
	BusinessTypeRepository *businessTypeRepository.BusinessTypePostgresRepository
	CreateUseCase           *usecase.CreateBusinessTypeTemplateUseCase
	UpdateUseCase           *usecase.UpdateBusinessTypeTemplateUseCase
	ListUseCase             *usecase.ListBusinessTypeTemplatesUseCase
	GetUseCase              *usecase.GetBusinessTypeTemplateUseCase
	DeleteUseCase           *usecase.DeleteBusinessTypeTemplateUseCase
	Handler                 *controller.BusinessTypeTemplateHandler
}

// NewBusinessTypeTemplateModuleConfig crea una nueva configuración del módulo
func NewBusinessTypeTemplateModuleConfig(db *sql.DB) *BusinessTypeTemplateModuleConfig {
	// Crear repositorios
	templateRepo := repository.NewBusinessTypeTemplatePostgresRepository(db).(*repository.BusinessTypeTemplatePostgresRepository)
	businessTypeRepo := businessTypeRepository.NewBusinessTypePostgresRepository(db).(*businessTypeRepository.BusinessTypePostgresRepository)

	// Crear casos de uso
	createUseCase := usecase.NewCreateBusinessTypeTemplateUseCase(templateRepo, businessTypeRepo)
	updateUseCase := usecase.NewUpdateBusinessTypeTemplateUseCase(templateRepo)
	listUseCase := usecase.NewListBusinessTypeTemplatesUseCase(templateRepo)
	getUseCase := usecase.NewGetBusinessTypeTemplateUseCase(templateRepo)
	deleteUseCase := usecase.NewDeleteBusinessTypeTemplateUseCase(templateRepo)

	// Crear handler
	handler := controller.NewBusinessTypeTemplateHandler(
		createUseCase,
		updateUseCase,
		listUseCase,
		getUseCase,
		deleteUseCase,
	)

	return &BusinessTypeTemplateModuleConfig{
		DB:                      db,
		TemplateRepository:      templateRepo,
		BusinessTypeRepository: businessTypeRepo,
		CreateUseCase:           createUseCase,
		UpdateUseCase:           updateUseCase,
		ListUseCase:             listUseCase,
		GetUseCase:              getUseCase,
		DeleteUseCase:           deleteUseCase,
		Handler:                 handler,
	}
}

// GetHandler devuelve el handler del módulo
func (cfg *BusinessTypeTemplateModuleConfig) GetHandler() *controller.BusinessTypeTemplateHandler {
	return cfg.Handler
}