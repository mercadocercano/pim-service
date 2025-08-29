package controller

import (
	"net/http"
	"saas-mt-pim-service/src/quickstart/application/usecase"
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
) *QuickstartHandler {
	return &QuickstartHandler{
		getBusinessTypesUseCase:            getBusinessTypesUseCase,
		getCategoriesByBusinessTypeUseCase: getCategoriesByBusinessTypeUseCase,
		getAttributesByBusinessTypeUseCase: getAttributesByBusinessTypeUseCase,
		getVariantsByBusinessTypeUseCase:   getVariantsByBusinessTypeUseCase,
		getProductsByBusinessTypeUseCase:   getProductsByBusinessTypeUseCase,
		getBrandsByBusinessTypeUseCase:     getBrandsByBusinessTypeUseCase,
		setupTenantUseCase:                 setupTenantUseCase,
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