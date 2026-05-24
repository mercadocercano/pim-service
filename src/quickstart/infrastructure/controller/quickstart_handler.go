package controller

import (
	"context"
	"log"
	"net/http"
	"saas-mt-pim-service/src/quickstart/application/usecase"
	backfillUseCase "saas-mt-pim-service/src/product/quickstart/application/usecase"
	"github.com/gin-gonic/gin"
)

// QuickstartHandler maneja las peticiones HTTP para el módulo quickstart
type QuickstartHandler struct {
	getBusinessTypesUseCase            *usecase.GetBusinessTypesUseCase
	getCategoriesByBusinessTypeUseCase *usecase.GetCategoriesByBusinessTypeUseCase
	getAttributesByBusinessTypeUseCase *usecase.GetAttributesByBusinessTypeUseCase
	getVariantsByBusinessTypeUseCase   *usecase.GetVariantsByBusinessTypeUseCase
	getProductsByBusinessTypeUseCase   *usecase.GetProductsByBusinessTypeUseCase
	getBrandsByBusinessTypeUseCase     *usecase.GetBrandsByBusinessTypeUseCase
	setupTenantUseCase                 *usecase.SetupTenantUseCase
	listTemplatesUseCase               *usecase.ListTemplatesUseCase
	applyTemplateUseCase               *usecase.ApplyTemplateUseCase
	backfillImagesUseCase              *backfillUseCase.BackfillTenantImagesUseCase
}

// NewQuickstartHandler crea una nueva instancia del handler
func NewQuickstartHandler(
	getBusinessTypesUseCase *usecase.GetBusinessTypesUseCase,
	getCategoriesByBusinessTypeUseCase *usecase.GetCategoriesByBusinessTypeUseCase,
	getAttributesByBusinessTypeUseCase *usecase.GetAttributesByBusinessTypeUseCase,
	getVariantsByBusinessTypeUseCase *usecase.GetVariantsByBusinessTypeUseCase,
	getProductsByBusinessTypeUseCase *usecase.GetProductsByBusinessTypeUseCase,
	getBrandsByBusinessTypeUseCase *usecase.GetBrandsByBusinessTypeUseCase,
	setupTenantUseCase *usecase.SetupTenantUseCase,
	listTemplatesUseCase *usecase.ListTemplatesUseCase,
	applyTemplateUseCase *usecase.ApplyTemplateUseCase,
	backfillImagesUseCase *backfillUseCase.BackfillTenantImagesUseCase,
) *QuickstartHandler {
	return &QuickstartHandler{
		getBusinessTypesUseCase:            getBusinessTypesUseCase,
		getCategoriesByBusinessTypeUseCase: getCategoriesByBusinessTypeUseCase,
		getAttributesByBusinessTypeUseCase: getAttributesByBusinessTypeUseCase,
		getVariantsByBusinessTypeUseCase:   getVariantsByBusinessTypeUseCase,
		getProductsByBusinessTypeUseCase:   getProductsByBusinessTypeUseCase,
		getBrandsByBusinessTypeUseCase:     getBrandsByBusinessTypeUseCase,
		setupTenantUseCase:                 setupTenantUseCase,
		listTemplatesUseCase:               listTemplatesUseCase,
		applyTemplateUseCase:               applyTemplateUseCase,
		backfillImagesUseCase:              backfillImagesUseCase,
	}
}

// RegisterRoutes registra las rutas del quickstart
func (h *QuickstartHandler) RegisterRoutes(router *gin.RouterGroup) {
	quickstart := router.Group("/quickstart")
	{
		quickstart.GET("/business-types", h.GetBusinessTypes)
		quickstart.GET("/categories/:businessType", h.GetCategoriesByBusinessType)
		quickstart.GET("/attributes/:businessType", h.GetAttributesByBusinessType)
		quickstart.GET("/variants/:businessType", h.GetVariantsByBusinessType)
		quickstart.GET("/products/:businessType", h.GetProductsByBusinessType)
		quickstart.GET("/brands/:businessType", h.GetBrandsByBusinessType)
		quickstart.POST("/setup", h.SetupTenant)
		
		// HITO 2: Nuevos endpoints para templates
		quickstart.GET("/templates", h.ListTemplates)
		quickstart.POST("/apply", h.ApplyTemplate) // Body: {"template_id": "xxx"}
		quickstart.POST("/templates/:id/apply", h.ApplyTemplateByID) // Path param: /templates/xxx/apply
	}
}

// GetBusinessTypes obtiene todos los tipos de negocio
func (h *QuickstartHandler) GetBusinessTypes(c *gin.Context) {
	businessTypes, err := h.getBusinessTypesUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, businessTypes)
}

// GetCategoriesByBusinessType obtiene categorías por tipo de negocio
func (h *QuickstartHandler) GetCategoriesByBusinessType(c *gin.Context) {
	businessType := c.Param("businessType")
	categories, err := h.getCategoriesByBusinessTypeUseCase.Execute(c.Request.Context(), businessType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GetAttributesByBusinessType obtiene atributos por tipo de negocio
func (h *QuickstartHandler) GetAttributesByBusinessType(c *gin.Context) {
	businessType := c.Param("businessType")
	attributes, err := h.getAttributesByBusinessTypeUseCase.Execute(c.Request.Context(), businessType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, attributes)
}

// GetVariantsByBusinessType obtiene variantes por tipo de negocio
func (h *QuickstartHandler) GetVariantsByBusinessType(c *gin.Context) {
	businessType := c.Param("businessType")
	variants, err := h.getVariantsByBusinessTypeUseCase.Execute(c.Request.Context(), businessType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, variants)
}

// GetProductsByBusinessType obtiene productos por tipo de negocio
func (h *QuickstartHandler) GetProductsByBusinessType(c *gin.Context) {
	businessType := c.Param("businessType")
	products, err := h.getProductsByBusinessTypeUseCase.Execute(c.Request.Context(), businessType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetBrandsByBusinessType obtiene marcas por tipo de negocio
func (h *QuickstartHandler) GetBrandsByBusinessType(c *gin.Context) {
	businessType := c.Param("businessType")
	brands, err := h.getBrandsByBusinessTypeUseCase.Execute(c.Request.Context(), businessType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

// SetupTenant configura un tenant con quickstart
func (h *QuickstartHandler) SetupTenant(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header is required"})
		return
	}

	var setupData map[string]interface{}
	if err := c.ShouldBindJSON(&setupData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	history, err := h.setupTenantUseCase.Execute(c.Request.Context(), tenantID, setupData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, history)
}

// ListTemplates obtiene la lista de templates disponibles
func (h *QuickstartHandler) ListTemplates(c *gin.Context) {
	templates, err := h.listTemplatesUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, templates)
}

// ApplyTemplate aplica un template de quickstart al tenant (template_id en body)
func (h *QuickstartHandler) ApplyTemplate(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header is required"})
		return
	}

	var req struct {
		TemplateID string `json:"template_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	applyReq := usecase.ApplyTemplateRequest{
		TemplateID: req.TemplateID,
		TenantID:   tenantID,
	}

	response, err := h.applyTemplateUseCase.Execute(c.Request.Context(), applyReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.triggerBackfill(tenantID)

	c.JSON(http.StatusCreated, response)
}

// ApplyTemplateByID aplica un template de quickstart al tenant (template_id en URL)
func (h *QuickstartHandler) ApplyTemplateByID(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header is required"})
		return
	}

	templateID := c.Param("id")
	if templateID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "template_id is required in path"})
		return
	}

	applyReq := usecase.ApplyTemplateRequest{
		TemplateID: templateID,
		TenantID:   tenantID,
	}

	response, err := h.applyTemplateUseCase.Execute(c.Request.Context(), applyReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.triggerBackfill(tenantID)

	c.JSON(http.StatusCreated, response)
}

// triggerBackfill lanza el backfill de imágenes en background, sin bloquear la respuesta.
func (h *QuickstartHandler) triggerBackfill(tenantID string) {
	if h.backfillImagesUseCase == nil {
		return
	}
	go func() {
		result, err := h.backfillImagesUseCase.Execute(context.Background(), tenantID)
		if err != nil {
			log.Printf("[quickstart] backfill de imágenes falló para tenant %s: %v", tenantID, err)
			return
		}
		log.Printf("[quickstart] backfill completado tenant=%s updated=%d skipped=%d errors=%d",
			tenantID, result.Updated, result.Skipped, result.Errors)
	}()
}