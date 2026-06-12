package controller

import (
	httpresp "github.com/hornosg/go-shared/infrastructure/response"
	"net/http"

	"github.com/gin-gonic/gin"

	"saas-mt-pim-service/src/product/quickstart/application/usecase"
)

// QuickstartController maneja las peticiones HTTP para quickstart
type QuickstartController struct {
	createFromTemplateUseCase      *usecase.CreateFromTemplateUseCase
	importFromBusinessTypeUseCase  *usecase.ImportFromBusinessTypeUseCase
	importFromGlobalCatalogUseCase *usecase.ImportFromGlobalCatalogUseCase
	getQuickstartProgressUseCase   *usecase.GetQuickstartProgressUseCase
	backfillUseCase                *usecase.BackfillTenantImagesUseCase
}

// NewQuickstartController crea una nueva instancia del controller
func NewQuickstartController(
	createFromTemplateUseCase *usecase.CreateFromTemplateUseCase,
	importFromBusinessTypeUseCase *usecase.ImportFromBusinessTypeUseCase,
	importFromGlobalCatalogUseCase *usecase.ImportFromGlobalCatalogUseCase,
	getQuickstartProgressUseCase *usecase.GetQuickstartProgressUseCase,
	backfillUseCase *usecase.BackfillTenantImagesUseCase,
) *QuickstartController {
	return &QuickstartController{
		createFromTemplateUseCase:      createFromTemplateUseCase,
		importFromBusinessTypeUseCase:  importFromBusinessTypeUseCase,
		importFromGlobalCatalogUseCase: importFromGlobalCatalogUseCase,
		getQuickstartProgressUseCase:   getQuickstartProgressUseCase,
		backfillUseCase:                backfillUseCase,
	}
}

// CreateProductFromTemplate godoc
// @Summary Crear producto desde template del catálogo global
// @Description Crea un producto del tenant basado en un template del catálogo global con estados flexibles
// @Tags quickstart
// @Accept json
// @Produce json
// @Param product body usecase.CreateFromTemplateRequest true "Datos del template"
// @Success 201 {object} usecase.CreateFromTemplateResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /quickstart/products/from-template [post]
// @Security BearerAuth
func (ctrl *QuickstartController) CreateProductFromTemplate(c *gin.Context) {
	var req usecase.CreateFromTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresp.JSONWithDetails(c, http.StatusBadRequest, "Datos de entrada inválidos", err.Error())
		return
	}

	// Obtener tenant ID del header y agregarlo al request
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}
	req.TenantID = tenantID

	// Ejecutar caso de uso
	response, err := ctrl.createFromTemplateUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, response)
}

// ImportProductsFromBusinessType godoc
// @Summary Importar productos desde tipo de negocio
// @Description Importa múltiples productos desde el catálogo global basado en el tipo de negocio del tenant
// @Tags quickstart
// @Accept json
// @Produce json
// @Param import_request body usecase.ImportFromBusinessTypeRequest true "Datos de importación"
// @Success 201 {object} usecase.ImportFromBusinessTypeResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /quickstart/products/import-from-business-type [post]
// @Security BearerAuth
func (ctrl *QuickstartController) ImportProductsFromBusinessType(c *gin.Context) {
	var req usecase.ImportFromBusinessTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresp.JSONWithDetails(c, http.StatusBadRequest, "Datos de entrada inválidos", err.Error())
		return
	}

	// Obtener tenant ID del header y agregarlo al request
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}
	req.TenantID = tenantID

	// Ejecutar caso de uso
	response, err := ctrl.importFromBusinessTypeUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, response)
}

// ImportFromGlobalCatalog godoc
// @Summary Importar un producto del catálogo global al catálogo del tenant
// @Description Busca un producto en el catálogo global por ID y lo copia al catálogo del tenant con todos sus datos
// @Tags quickstart
// @Accept json
// @Produce json
// @Param import_request body usecase.ImportFromGlobalCatalogRequest true "ID del producto global"
// @Success 201 {object} usecase.ImportFromGlobalCatalogResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /quickstart/products/import-from-global [post]
// @Security BearerAuth
func (ctrl *QuickstartController) ImportFromGlobalCatalog(c *gin.Context) {
	var req usecase.ImportFromGlobalCatalogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresp.JSONWithDetails(c, http.StatusBadRequest, "Datos de entrada inválidos", err.Error())
		return
	}

	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}
	req.TenantID = tenantID

	response, err := ctrl.importFromGlobalCatalogUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetQuickstartProgress godoc
// @Summary Obtener progreso del quickstart
// @Description Retorna el progreso del proceso de quickstart para un tenant
// @Tags quickstart
// @Produce json
// @Success 200 {object} usecase.GetQuickstartProgressResponse
// @Failure 400 {object} map[string]interface{}
// @Router /quickstart/progress [get]
// @Security BearerAuth
func (ctrl *QuickstartController) GetQuickstartProgress(c *gin.Context) {
	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}

	// Crear request
	req := usecase.GetQuickstartProgressRequest{
		TenantID: tenantID,
	}

	// Ejecutar caso de uso
	response, err := ctrl.getQuickstartProgressUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

// BackfillTenantImages godoc
// @Summary Rellenar imágenes faltantes del tenant desde el catálogo global
// @Description Copia image_url desde global_products a los productos del tenant sin imagen, usando match por nombre y marca
// @Tags quickstart
// @Produce json
// @Success 200 {object} usecase.BackfillResult
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /quickstart/backfill-tenant-images [post]
// @Security BearerAuth
func (ctrl *QuickstartController) BackfillTenantImages(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}

	result, err := ctrl.backfillUseCase.Execute(c.Request.Context(), tenantID)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

// BackfillAllTenantImages godoc
// @Summary Rellenar imágenes para todos los tenants
// @Description Ejecuta el backfill de imágenes para todos los tenants que tienen productos registrados
// @Tags quickstart
// @Produce json
// @Success 200 {object} map[string]usecase.BackfillResult
// @Failure 500 {object} map[string]interface{}
// @Router /quickstart/backfill-all-tenant-images [post]
// @Security BearerAuth
func (ctrl *QuickstartController) BackfillAllTenantImages(c *gin.Context) {
	results, err := ctrl.backfillUseCase.ExecuteAll(c.Request.Context())
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}

// RegisterRoutes registra las rutas del controller
func (ctrl *QuickstartController) RegisterRoutes(router *gin.RouterGroup) {
	quickstart := router.Group("/quickstart")
	{
		// Productos desde templates
		products := quickstart.Group("/products")
		{
			products.POST("/from-template", ctrl.CreateProductFromTemplate)
			products.POST("/import-from-business-type", ctrl.ImportProductsFromBusinessType)
			products.POST("/import-from-global", ctrl.ImportFromGlobalCatalog)
		}

		// Progreso del quickstart
		quickstart.GET("/progress", ctrl.GetQuickstartProgress)

		// Backfill de imágenes
		quickstart.POST("/backfill-tenant-images", ctrl.BackfillTenantImages)
		quickstart.POST("/backfill-all-tenant-images", ctrl.BackfillAllTenantImages)
	}
}
