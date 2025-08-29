package config

import (
	"database/sql"

	"saas-mt-pim-service/src/template_ai/application/mapper"
	"saas-mt-pim-service/src/template_ai/application/usecase"
	"saas-mt-pim-service/src/template_ai/domain/port"
	"saas-mt-pim-service/src/template_ai/domain/service"
	"saas-mt-pim-service/src/template_ai/infrastructure/controller"
	"saas-mt-pim-service/src/template_ai/infrastructure/persistence/repository"
	aiservice "saas-mt-pim-service/src/template_ai/infrastructure/service"
)

// AITemplateConfig handles the configuration and dependency injection for AI template module
type AITemplateConfig struct {
	db                 *sql.DB
	controller         *controller.AITemplateController
	aiTemplateRepo     port.AITemplateRepository
	globalProductRepo  port.GlobalProductRepository
	aiGenerationService port.AIGenerationService
}

// NewAITemplateConfig creates a new configuration instance
func NewAITemplateConfig(db *sql.DB) *AITemplateConfig {
	config := &AITemplateConfig{db: db}
	config.initializeComponents()
	return config
}

// initializeComponents initializes all components with dependency injection
func (c *AITemplateConfig) initializeComponents() {
	// Initialize repositories
	c.aiTemplateRepo = repository.NewAITemplatePostgresRepository(c.db)
	c.globalProductRepo = repository.NewGlobalProductPostgresRepository(c.db)
	
	// Initialize external services
	c.aiGenerationService = aiservice.NewAIGenerationService()

	// Initialize domain service
	domainService := service.NewAITemplateDomainService(
		c.aiTemplateRepo,
		c.globalProductRepo,
		c.aiGenerationService,
	)

	// Initialize mapper
	templateMapper := mapper.NewTemplateMapper()

	// Initialize use cases
	generateSmartTemplateUseCase := usecase.NewGenerateSmartTemplateUseCase(
		c.aiTemplateRepo,
		c.aiGenerationService,
		domainService,
		templateMapper,
	)

	applyDynamicTemplateUseCase := usecase.NewApplyDynamicTemplateUseCase(
		c.aiTemplateRepo,
		c.globalProductRepo,
		domainService,
		templateMapper,
	)

	analyzeTemplatePerformanceUseCase := usecase.NewAnalyzeTemplatePerformanceUseCase(
		c.aiTemplateRepo,
		domainService,
		templateMapper,
	)

	updateTemplateFromFeedbackUseCase := usecase.NewUpdateTemplateFromFeedbackUseCase(
		c.aiTemplateRepo,
		c.aiGenerationService,
		domainService,
		templateMapper,
	)

	// Initialize controller
	c.controller = controller.NewAITemplateController(
		generateSmartTemplateUseCase,
		applyDynamicTemplateUseCase,
		analyzeTemplatePerformanceUseCase,
		updateTemplateFromFeedbackUseCase,
	)
}

// GetController returns the AI template controller
func (c *AITemplateConfig) GetController() *controller.AITemplateController {
	return c.controller
}

// GetAITemplateRepository returns the AI template repository
func (c *AITemplateConfig) GetAITemplateRepository() port.AITemplateRepository {
	return c.aiTemplateRepo
}

// GetGlobalProductRepository returns the global product repository
func (c *AITemplateConfig) GetGlobalProductRepository() port.GlobalProductRepository {
	return c.globalProductRepo
}

// GetAIGenerationService returns the AI generation service
func (c *AITemplateConfig) GetAIGenerationService() port.AIGenerationService {
	return c.aiGenerationService
}