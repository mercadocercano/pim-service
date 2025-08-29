package config

import (
	"database/sql"

	"saas-mt-pim-service/src/brand/application/mapper"
	"saas-mt-pim-service/src/brand/application/usecase"
	"saas-mt-pim-service/src/brand/domain/service"
	"saas-mt-pim-service/src/brand/infrastructure/controller"
	"saas-mt-pim-service/src/brand/infrastructure/criteria"
	"saas-mt-pim-service/src/brand/infrastructure/persistence"

	"github.com/gin-gonic/gin"
)

// BrandModuleConfig contiene todas las dependencias del módulo Brand
type BrandModuleConfig struct {
	// Repositories
	BrandRepository *persistence.PostgresBrandRepository

	// Domain Services
	BrandDomainService *service.BrandDomainService

	// Mappers
	BrandMapper *mapper.BrandMapper

	// Use Cases
	CreateBrandUseCase          *usecase.CreateBrandUseCase
	GetBrandByIDUseCase         *usecase.GetBrandByIDUseCase
	UpdateBrandUseCase          *usecase.UpdateBrandUseCase
	DeleteBrandUseCase          *usecase.DeleteBrandUseCase
	ListBrandsByCriteriaUseCase *usecase.ListBrandsByCriteriaUseCase

	// Infrastructure
	BrandCriteriaBuilder *criteria.BrandCriteriaBuilder
	BrandController      *controller.BrandController
}

// NewBrandModuleConfig crea e inicializa todas las dependencias del módulo Brand
func NewBrandModuleConfig(db *sql.DB) *BrandModuleConfig {
	// Repositories
	brandRepo := persistence.NewPostgresBrandRepository(db)

	// Domain Services
	brandDomainService := service.NewBrandDomainService(brandRepo)

	// Mappers
	brandMapper := mapper.NewBrandMapper()

	// Use Cases
	createBrandUseCase := usecase.NewCreateBrandUseCase(brandRepo, brandDomainService, brandMapper)
	getBrandByIDUseCase := usecase.NewGetBrandByIDUseCase(brandRepo, brandMapper)
	updateBrandUseCase := usecase.NewUpdateBrandUseCase(brandRepo, brandDomainService, brandMapper)
	deleteBrandUseCase := usecase.NewDeleteBrandUseCase(brandRepo, brandDomainService)
	listBrandsByCriteriaUseCase := usecase.NewListBrandsByCriteriaUseCase(brandRepo, brandMapper)

	// Infrastructure
	brandCriteriaBuilder := criteria.NewBrandCriteriaBuilder()
	brandController := controller.NewBrandController(
		createBrandUseCase,
		getBrandByIDUseCase,
		updateBrandUseCase,
		deleteBrandUseCase,
		listBrandsByCriteriaUseCase,
		brandCriteriaBuilder,
	)

	return &BrandModuleConfig{
		BrandRepository:             brandRepo,
		BrandDomainService:          brandDomainService,
		BrandMapper:                 brandMapper,
		CreateBrandUseCase:          createBrandUseCase,
		GetBrandByIDUseCase:         getBrandByIDUseCase,
		UpdateBrandUseCase:          updateBrandUseCase,
		DeleteBrandUseCase:          deleteBrandUseCase,
		ListBrandsByCriteriaUseCase: listBrandsByCriteriaUseCase,
		BrandCriteriaBuilder:        brandCriteriaBuilder,
		BrandController:             brandController,
	}
}

// GetBrandController retorna el controller de Brand
func (c *BrandModuleConfig) GetBrandController() *controller.BrandController {
	return c.BrandController
}

// GetBrandRepository retorna el repositorio de Brand
func (c *BrandModuleConfig) GetBrandRepository() *persistence.PostgresBrandRepository {
	return c.BrandRepository
}

// SetupBrandModule configura el módulo de marcas y sus dependencias
func SetupBrandModule(router *gin.RouterGroup, db *sql.DB) {
	// Crear configuración del módulo
	brandConfig := NewBrandModuleConfig(db)

	// Obtener el controller
	brandController := brandConfig.GetBrandController()

	// Registrar las rutas
	brandController.RegisterRoutes(router)
}
