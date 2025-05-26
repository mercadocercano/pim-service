package controller

import (
	"net/http"

	"pim/src/quickstart/application/request"
	"pim/src/quickstart/application/response"
	"pim/src/quickstart/application/usecase"
	"pim/src/quickstart/domain/exception"

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

// NewQuickstartHandler crea una nueva instancia del manejador de quickstart
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

// RegisterRoutes registra las rutas del API para quickstart
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

// GetBusinessTypes maneja la solicitud para obtener todos los tipos de negocio
func (h *QuickstartHandler) GetBusinessTypes(c *gin.Context) {
	businessTypes, err := h.getBusinessTypesUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.FromBusinessTypeEntities(businessTypes))
}

// GetCategoriesByBusinessType maneja la solicitud para obtener categorías por tipo de negocio
func (h *QuickstartHandler) GetCategoriesByBusinessType(c *gin.Context) {
	businessType := c.Param("businessType")
	if businessType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el tipo de negocio es obligatorio"})
		return
	}

	categories, err := h.getCategoriesByBusinessTypeUseCase.Execute(c.Request.Context(), businessType)
	if err != nil {
		if err == exception.ErrBusinessTypeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetAttributesByBusinessType maneja la solicitud para obtener atributos por tipo de negocio
func (h *QuickstartHandler) GetAttributesByBusinessType(c *gin.Context) {
	businessType := c.Param("businessType")
	if businessType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el tipo de negocio es obligatorio"})
		return
	}

	attributes, err := h.getAttributesByBusinessTypeUseCase.Execute(c.Request.Context(), businessType)
	if err != nil {
		if err == exception.ErrBusinessTypeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attributes)
}

// GetVariantsByBusinessType maneja la solicitud para obtener variantes por tipo de negocio
func (h *QuickstartHandler) GetVariantsByBusinessType(c *gin.Context) {
	businessType := c.Param("businessType")
	if businessType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el tipo de negocio es obligatorio"})
		return
	}

	variants, err := h.getVariantsByBusinessTypeUseCase.Execute(c.Request.Context(), businessType)
	if err != nil {
		if err == exception.ErrBusinessTypeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, variants)
}

// GetProductsByBusinessType maneja la solicitud para obtener productos por tipo de negocio
func (h *QuickstartHandler) GetProductsByBusinessType(c *gin.Context) {
	businessType := c.Param("businessType")
	if businessType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el tipo de negocio es obligatorio"})
		return
	}

	products, err := h.getProductsByBusinessTypeUseCase.Execute(c.Request.Context(), businessType)
	if err != nil {
		if err == exception.ErrBusinessTypeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetBrandsByBusinessType maneja la solicitud para obtener marcas por tipo de negocio
func (h *QuickstartHandler) GetBrandsByBusinessType(c *gin.Context) {
	businessType := c.Param("businessType")
	if businessType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el tipo de negocio es obligatorio"})
		return
	}

	brands, err := h.getBrandsByBusinessTypeUseCase.Execute(c.Request.Context(), businessType)
	if err != nil {
		if err == exception.ErrBusinessTypeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brands)
}

// SetupTenant maneja la solicitud para configurar un tenant con quickstart
func (h *QuickstartHandler) SetupTenant(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	var req request.SetupTenantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	history, err := h.setupTenantUseCase.Execute(c.Request.Context(), tenantID, req.ToMap())
	if err != nil {
		if err == exception.ErrQuickstartAlreadyCompleted {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if err == exception.ErrQuickstartInProgress {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if err == exception.ErrInvalidSetupData {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.FromTenantQuickstartHistoryEntity(history))
}
