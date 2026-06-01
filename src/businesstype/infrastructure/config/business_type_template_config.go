package config

import (
	"database/sql"

	"saas-mt-pim-service/src/businesstype/application/usecase"
	"saas-mt-pim-service/src/businesstype/infrastructure/controller"
	"saas-mt-pim-service/src/businesstype/infrastructure/persistence/repository"
	businessTypeRepository "saas-mt-pim-service/src/businesstype/infrastructure/persistence/repository"
)

// BusinessTypeTemplateModuleConfig contiene la configuración del módulo de templates
type BusinessTypeTemplateModuleConfig struct {
	DB                     *sql.DB
	TemplateRepository     *repository.BusinessTypeTemplatePostgresRepository
	BusinessTypeRepository *businessTypeRepository.BusinessTypePostgresRepository
	CreateUseCase          *usecase.CreateBusinessTypeTemplateUseCase
	UpdateUseCase          *usecase.UpdateBusinessTypeTemplateUseCase
	ListUseCase            *usecase.ListBusinessTypeTemplatesUseCase
	GetUseCase             *usecase.GetBusinessTypeTemplateUseCase
	DeleteUseCase          *usecase.DeleteBusinessTypeTemplateUseCase
	AnalyticsUseCase       *usecase.GetTemplateAnalyticsUseCase
	DuplicateUseCase       *usecase.DuplicateTemplateUseCase
	Handler                *controller.BusinessTypeTemplateHandler
}

// NewBusinessTypeTemplateModuleConfig crea una nueva configuración del módulo
func NewBusinessTypeTemplateModuleConfig(db *sql.DB) *BusinessTypeTemplateModuleConfig {
	templateRepo := repository.NewBusinessTypeTemplatePostgresRepository(db).(*repository.BusinessTypeTemplatePostgresRepository)
	businessTypeRepo := businessTypeRepository.NewBusinessTypePostgresRepository(db).(*businessTypeRepository.BusinessTypePostgresRepository)
	analyticsRepo := repository.NewTemplateAnalyticsPostgresRepository(db)

	createUseCase := usecase.NewCreateBusinessTypeTemplateUseCase(templateRepo, businessTypeRepo)
	updateUseCase := usecase.NewUpdateBusinessTypeTemplateUseCase(templateRepo)
	listUseCase := usecase.NewListBusinessTypeTemplatesUseCase(templateRepo)
	getUseCase := usecase.NewGetBusinessTypeTemplateUseCase(templateRepo)
	deleteUseCase := usecase.NewDeleteBusinessTypeTemplateUseCase(templateRepo)
	analyticsUseCase := usecase.NewGetTemplateAnalyticsUseCase(templateRepo, analyticsRepo)
	duplicateUseCase := usecase.NewDuplicateTemplateUseCase(templateRepo)

	handler := controller.NewBusinessTypeTemplateHandler(
		createUseCase,
		updateUseCase,
		listUseCase,
		getUseCase,
		deleteUseCase,
	).WithAnalyticsUseCase(analyticsUseCase).
		WithDuplicateUseCase(duplicateUseCase)

	return &BusinessTypeTemplateModuleConfig{
		DB:                     db,
		TemplateRepository:     templateRepo,
		BusinessTypeRepository: businessTypeRepo,
		CreateUseCase:          createUseCase,
		UpdateUseCase:          updateUseCase,
		ListUseCase:            listUseCase,
		GetUseCase:             getUseCase,
		DeleteUseCase:          deleteUseCase,
		AnalyticsUseCase:       analyticsUseCase,
		DuplicateUseCase:       duplicateUseCase,
		Handler:                handler,
	}
}

// GetHandler devuelve el handler del módulo
func (cfg *BusinessTypeTemplateModuleConfig) GetHandler() *controller.BusinessTypeTemplateHandler {
	return cfg.Handler
}
