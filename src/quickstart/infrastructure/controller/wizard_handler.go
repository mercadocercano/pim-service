package controller

import (
	"net/http"
	"strconv"

	"pim/src/quickstart/application/usecase"

	"github.com/gin-gonic/gin"
)

// WizardHandler maneja las peticiones HTTP para el wizard de configuración
type WizardHandler struct {
	startWizardUseCase     *usecase.StartWizardUseCase
	getWizardStatusUseCase *usecase.GetWizardStatusUseCase
	updateWizardStepUseCase *usecase.UpdateWizardStepUseCase
	getTemplateDataUseCase *usecase.GetTemplateDataUseCase
}

// NewWizardHandler crea una nueva instancia del manejador de wizard
func NewWizardHandler(
	startWizardUseCase *usecase.StartWizardUseCase,
	getWizardStatusUseCase *usecase.GetWizardStatusUseCase,
	updateWizardStepUseCase *usecase.UpdateWizardStepUseCase,
	getTemplateDataUseCase *usecase.GetTemplateDataUseCase,
) *WizardHandler {
	return &WizardHandler{
		startWizardUseCase:      startWizardUseCase,
		getWizardStatusUseCase:  getWizardStatusUseCase,
		updateWizardStepUseCase: updateWizardStepUseCase,
		getTemplateDataUseCase:  getTemplateDataUseCase,
	}
}

// RegisterRoutes registra las rutas del API para el wizard
func (h *WizardHandler) RegisterRoutes(router *gin.RouterGroup) {
	wizard := router.Group("/wizard")
	{
		// Wizard status y control
		wizard.GET("/status", h.GetWizardStatus)
		wizard.POST("/start", h.StartWizard)
		wizard.PUT("/step", h.UpdateWizardStep)
		
		// Template data para el wizard
		wizard.GET("/template/:businessTypeId", h.GetTemplateData)
		wizard.GET("/template/:businessTypeId/:section", h.GetTemplateSectionData)
	}
}

// StartWizardRequest representa la petición para iniciar el wizard
type StartWizardRequest struct {
	BusinessTypeID string `json:"business_type_id" binding:"required"`
}

// UpdateWizardStepRequest representa la petición para actualizar un step del wizard
type UpdateWizardStepRequest struct {
	CurrentStep string                 `json:"current_step" binding:"required"`
	StepData    map[string]interface{} `json:"step_data"`
}

// StartWizard inicia el wizard de configuración
func (h *WizardHandler) StartWizard(c *gin.Context) {
	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	var req StartWizardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de petición inválido: " + err.Error()})
		return
	}

	history, err := h.startWizardUseCase.Execute(c.Request.Context(), tenantID, req.BusinessTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"wizard_id":        history.ID,
		"tenant_id":        history.TenantID,
		"business_type_id": history.BusinessTypeID,
		"setup_data":       history.SetupData,
		"created_at":       history.CreatedAt,
	})
}

// GetWizardStatus obtiene el estado actual del wizard
func (h *WizardHandler) GetWizardStatus(c *gin.Context) {
	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	history, err := h.getWizardStatusUseCase.Execute(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if history == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No hay wizard iniciado para este tenant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"wizard_id":        history.ID,
		"tenant_id":        history.TenantID,
		"business_type_id": history.BusinessTypeID,
		"setup_data":       history.SetupData,
		"created_at":       history.CreatedAt,
		"updated_at":       history.UpdatedAt,
	})
}

// UpdateWizardStep actualiza el progreso del wizard
func (h *WizardHandler) UpdateWizardStep(c *gin.Context) {
	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	var req UpdateWizardStepRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de petición inválido: " + err.Error()})
		return
	}

	history, err := h.updateWizardStepUseCase.Execute(c.Request.Context(), tenantID, req.CurrentStep, req.StepData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"wizard_id":        history.ID,
		"tenant_id":        history.TenantID,
		"business_type_id": history.BusinessTypeID,
		"setup_data":       history.SetupData,
		"updated_at":       history.UpdatedAt,
	})
}

// GetTemplateData obtiene todos los datos del template
func (h *WizardHandler) GetTemplateData(c *gin.Context) {
	businessTypeID := c.Param("businessTypeId")
	if businessTypeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "business_type_id es requerido"})
		return
	}

	templateData, err := h.getTemplateDataUseCase.Execute(c.Request.Context(), businessTypeID, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"business_type_id": businessTypeID,
		"template_data":    templateData,
	})
}

// GetTemplateSectionData obtiene una sección específica del template
func (h *WizardHandler) GetTemplateSectionData(c *gin.Context) {
	businessTypeID := c.Param("businessTypeId")
	section := c.Param("section")
	
	if businessTypeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "business_type_id es requerido"})
		return
	}
	if section == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "section es requerido"})
		return
	}

	// Validar secciones permitidas
	allowedSections := []string{"categories", "brands", "products", "attributes", "variants"}
	isValidSection := false
	for _, allowedSection := range allowedSections {
		if section == allowedSection {
			isValidSection = true
			break
		}
	}
	if !isValidSection {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sección no válida. Permitidas: categories, brands, products, attributes, variants"})
		return
	}

	sectionData, err := h.getTemplateDataUseCase.Execute(c.Request.Context(), businessTypeID, section)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Agregar metadata de paginación si se solicita
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "50"))

	response := gin.H{
		"business_type_id": businessTypeID,
		"section":          section,
		"data":             sectionData,
	}

	// Si la respuesta es un array, agregar metadata de paginación
	if dataArray, ok := sectionData.([]interface{}); ok {
		total := len(dataArray)
		start := (page - 1) * pageSize
		end := start + pageSize
		if end > total {
			end = total
		}
		if start < total {
			response["data"] = dataArray[start:end]
		} else {
			response["data"] = []interface{}{}
		}
		
		response["pagination"] = gin.H{
			"page":         page,
			"page_size":    pageSize,
			"total_items":  total,
			"total_pages":  (total + pageSize - 1) / pageSize,
			"has_next":     end < total,
			"has_previous": page > 1,
		}
	}

	c.JSON(http.StatusOK, response)
}