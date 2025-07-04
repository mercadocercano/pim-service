package controller

import (
	"net/http"

	businessTypeUsecase "pim/src/businesstype/application/usecase"

	"github.com/gin-gonic/gin"
)

// SimpleWizardHandler maneja los endpoints del wizard de forma simplificada
// usando solo los business types desde la BD, sin template ni tenant setup
type SimpleWizardHandler struct {
	listBusinessTypesUseCase *businessTypeUsecase.ListBusinessTypesUseCase
}

// NewSimpleWizardHandler crea un nuevo handler simplificado del wizard
func NewSimpleWizardHandler(listBusinessTypesUseCase *businessTypeUsecase.ListBusinessTypesUseCase) *SimpleWizardHandler {
	return &SimpleWizardHandler{
		listBusinessTypesUseCase: listBusinessTypesUseCase,
	}
}

// RegisterRoutes registra las rutas del wizard simplificado
func (h *SimpleWizardHandler) RegisterRoutes(router *gin.RouterGroup) {
	wizard := router.Group("/wizard")
	{
		wizard.GET("/status", h.GetWizardStatus)
		wizard.POST("/start", h.StartWizard)
		wizard.PUT("/step", h.UpdateWizardStep)
		wizard.GET("/template/:businessTypeId", h.GetTemplateData)
		wizard.GET("/template/:businessTypeId/:section", h.GetTemplateSectionData)
	}
}

// GetWizardStatus obtiene el estado actual del wizard
func (h *SimpleWizardHandler) GetWizardStatus(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "X-Tenant-ID header is required",
		})
		return
	}

	// Respuesta simplificada para demostrar funcionalidad
	c.JSON(http.StatusOK, gin.H{
		"wizard_id":   "demo-wizard-id",
		"tenant_id":   tenantID,
		"setup_data":  gin.H{
			"step": "not_started",
			"completed_steps": []string{},
		},
		"setup_completed": false,
	})
}

// StartWizard inicia el proceso del wizard
func (h *SimpleWizardHandler) StartWizard(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "X-Tenant-ID header is required",
		})
		return
	}

	var req struct {
		BusinessTypeID string `json:"business_type_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Respuesta simplificada para demostrar funcionalidad
	c.JSON(http.StatusOK, gin.H{
		"wizard_id":        "demo-wizard-id",
		"tenant_id":        tenantID,
		"business_type_id": req.BusinessTypeID,
		"setup_data": gin.H{
			"step": "business_type_selected",
			"completed_steps": []string{"business_type_selected"},
			"business_type_id": req.BusinessTypeID,
		},
		"setup_completed": false,
	})
}

// UpdateWizardStep actualiza el paso actual del wizard
func (h *SimpleWizardHandler) UpdateWizardStep(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "X-Tenant-ID header is required",
		})
		return
	}

	var req struct {
		CurrentStep string                 `json:"current_step" binding:"required"`
		StepData    map[string]interface{} `json:"step_data" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Respuesta simplificada para demostrar funcionalidad
	c.JSON(http.StatusOK, gin.H{
		"wizard_id":   "demo-wizard-id",
		"tenant_id":   tenantID,
		"updated_at":  "2024-01-01T00:00:00Z",
		"setup_data":  gin.H{
			"step": req.CurrentStep,
			"completed_steps": []string{"business_type_selected", req.CurrentStep},
			"step_data": req.StepData,
		},
		"setup_completed": false,
	})
}

// GetTemplateData obtiene los datos del template completo
func (h *SimpleWizardHandler) GetTemplateData(c *gin.Context) {
	businessTypeID := c.Param("businessTypeId")
	if businessTypeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "business type ID is required",
		})
		return
	}

	// Respuesta simplificada con datos de ejemplo
	c.JSON(http.StatusOK, gin.H{
		"template_data": gin.H{
			"categories": []gin.H{
				{"id": 1, "name": "Electrónicos", "code": "electronics"},
				{"id": 2, "name": "Ropa", "code": "clothing"},
				{"id": 3, "name": "Hogar", "code": "home"},
			},
			"products": []gin.H{
				{"id": 1, "name": "Smartphone", "category": "electronics"},
				{"id": 2, "name": "Laptop", "category": "electronics"},
				{"id": 3, "name": "Camiseta", "category": "clothing"},
			},
			"brands": []string{"Samsung", "Apple", "LG", "Sony", "Nike", "Adidas"},
			"attributes": []gin.H{
				{"id": 1, "name": "Color", "type": "text"},
				{"id": 2, "name": "Tamaño", "type": "text"},
				{"id": 3, "name": "Material", "type": "text"},
			},
			"metadata": gin.H{
				"business_type_id": businessTypeID,
				"total_categories": 3,
				"total_products": 3,
				"total_brands": 6,
				"total_attributes": 3,
			},
		},
	})
}

// GetTemplateSectionData obtiene los datos de una sección específica del template
func (h *SimpleWizardHandler) GetTemplateSectionData(c *gin.Context) {
	businessTypeID := c.Param("businessTypeId")
	section := c.Param("section")
	
	if businessTypeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "business type ID is required",
		})
		return
	}

	// Parámetros de paginación
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	var data interface{}
	
	switch section {
	case "categories":
		data = []gin.H{
			{"id": 1, "name": "Electrónicos", "code": "electronics"},
			{"id": 2, "name": "Ropa", "code": "clothing"},
			{"id": 3, "name": "Hogar", "code": "home"},
		}
	case "products":
		data = []gin.H{
			{"id": 1, "name": "Smartphone", "category": "electronics"},
			{"id": 2, "name": "Laptop", "category": "electronics"},
			{"id": 3, "name": "Camiseta", "category": "clothing"},
		}
	case "brands":
		data = []string{"Samsung", "Apple", "LG", "Sony", "Nike", "Adidas"}
	case "attributes":
		data = []gin.H{
			{"id": 1, "name": "Color", "type": "text"},
			{"id": 2, "name": "Tamaño", "type": "text"},
			{"id": 3, "name": "Material", "type": "text"},
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid section. Valid sections are: categories, products, brands, attributes",
		})
		return
	}

	// Calcular longitud basado en el tipo de data
	var totalItems int
	switch d := data.(type) {
	case []gin.H:
		totalItems = len(d)
	case []string:
		totalItems = len(d)
	default:
		totalItems = 0
	}

	// Respuesta con paginación
	c.JSON(http.StatusOK, gin.H{
		"section": section,
		"data":    data,
		"pagination": gin.H{
			"page":        page,
			"page_size":   pageSize,
			"total_items": totalItems,
			"has_next":    false,
			"has_prev":    false,
		},
	})
}