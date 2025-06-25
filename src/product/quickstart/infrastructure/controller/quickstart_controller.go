package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pim/src/product/quickstart/application/usecase"
)

// QuickstartController maneja las peticiones HTTP para quickstart
type QuickstartController struct {
	createFromTemplateUseCase     *usecase.CreateFromTemplateUseCase
	importFromBusinessTypeUseCase *usecase.ImportFromBusinessTypeUseCase
	getQuickstartProgressUseCase  *usecase.GetQuickstartProgressUseCase
}

// NewQuickstartController crea una nueva instancia del controller
func NewQuickstartController(
	createFromTemplateUseCase *usecase.CreateFromTemplateUseCase,
	importFromBusinessTypeUseCase *usecase.ImportFromBusinessTypeUseCase,
	getQuickstartProgressUseCase *usecase.GetQuickstartProgressUseCase,
) *QuickstartController {
	return &QuickstartController{
		createFromTemplateUseCase:     createFromTemplateUseCase,
		importFromBusinessTypeUseCase: importFromBusinessTypeUseCase,
		getQuickstartProgressUseCase:  getQuickstartProgressUseCase,
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos", "details": err.Error()})
		return
	}

	// Obtener tenant ID del header y agregarlo al request
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}
	req.TenantID = tenantID

	// Ejecutar caso de uso
	response, err := ctrl.createFromTemplateUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos", "details": err.Error()})
		return
	}

	// Obtener tenant ID del header y agregarlo al request
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}
	req.TenantID = tenantID

	// Ejecutar caso de uso
	response, err := ctrl.importFromBusinessTypeUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Crear request
	req := usecase.GetQuickstartProgressRequest{
		TenantID: tenantID,
	}

	// Ejecutar caso de uso
	response, err := ctrl.getQuickstartProgressUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
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
		}

		// Progreso del quickstart
		quickstart.GET("/progress", ctrl.GetQuickstartProgress)
	}
}
