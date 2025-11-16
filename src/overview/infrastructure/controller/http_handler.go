package controller

import (
	"net/http"
	"strconv"
	"strings"

	"saas-mt-pim-service/src/overview/application/request"
	"saas-mt-pim-service/src/overview/application/usecase"

	"github.com/gin-gonic/gin"
)

// OverviewHandler maneja las peticiones HTTP para overview del marketplace
type OverviewHandler struct {
	getOverviewUseCase *usecase.GetMarketplaceOverviewUseCase
}

// NewOverviewHandler crea una nueva instancia del manejador
func NewOverviewHandler(getOverviewUseCase *usecase.GetMarketplaceOverviewUseCase) *OverviewHandler {
	return &OverviewHandler{
		getOverviewUseCase: getOverviewUseCase,
	}
}

// RegisterRoutes registra las rutas del API
func (h *OverviewHandler) RegisterRoutes(router *gin.RouterGroup) {
	marketplace := router.Group("/marketplace")
	{
		marketplace.GET("/overview", h.GetOverview)
		marketplace.GET("/overview/sections", h.GetAvailableSections)
	}
}

// GetOverview maneja la solicitud para obtener overview del marketplace
// GET /api/v1/marketplace/overview?sections=dashboard,taxonomy&include_stats=true&parallel=true
func (h *OverviewHandler) GetOverview(c *gin.Context) {
	// Obtener el tenantID del header (para marketplace admin es "global")
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		tenantID = "global" // Default para marketplace admin
	}

	// Parsear parámetros de query
	sectionsParam := c.Query("sections")
	if sectionsParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Parámetro 'sections' es requerido",
			"example": "/api/v1/marketplace/overview?sections=dashboard,taxonomy,brands",
		})
		return
	}

	// Crear request
	req := &request.GetMarketplaceOverviewRequest{
		Sections:      strings.Split(sectionsParam, ","),
		TenantID:      tenantID,
		IncludeStats:  c.Query("include_stats") == "true",
		Parallel:      c.Query("parallel") == "true",
		TimeRangeDays: 7,  // Default
		Limit:         10, // Default
	}

	// Parsear time_range_days
	if timeRangeStr := c.Query("time_range_days"); timeRangeStr != "" {
		if timeRange, err := strconv.Atoi(timeRangeStr); err == nil {
			req.TimeRangeDays = timeRange
		}
	}

	// Parsear limit
	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			req.Limit = limit
		}
	}

	// Ejecutar caso de uso
	resp, err := h.getOverviewUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Responder con éxito
	c.JSON(http.StatusOK, resp)
}

// GetAvailableSections retorna las secciones disponibles para el overview
// GET /api/v1/marketplace/overview/sections
func (h *OverviewHandler) GetAvailableSections(c *gin.Context) {
	sections := map[string]interface{}{
		"available_sections": []map[string]interface{}{
			{
				"name":        "dashboard",
				"description": "Métricas generales del marketplace",
				"fields":      []string{"total_categories", "total_brands", "total_attributes", "total_global_products", "active_tenants", "recent_activity"},
			},
			{
				"name":        "taxonomy",
				"description": "Estadísticas de categorías",
				"fields":      []string{"total_categories", "active_categories", "root_categories", "avg_depth", "most_used_categories"},
			},
			{
				"name":        "brands",
				"description": "Estadísticas de marcas",
				"fields":      []string{"total_brands", "verified_brands", "pending_brands", "top_brands", "recent_brands"},
			},
			{
				"name":        "global-catalog",
				"description": "Estadísticas del catálogo global",
				"fields":      []string{"total_products", "verified_products", "argentine_products", "avg_quality_score", "top_categories", "recent_products"},
			},
			{
				"name":        "attributes",
				"description": "Estadísticas de atributos",
				"fields":      []string{"total_attributes", "active_attributes", "required_attributes", "by_type", "most_used"},
			},
			{
				"name":        "curation",
				"description": "Estadísticas de curación de productos",
				"fields":      []string{"total_products", "pending_curation", "curated_products", "rejected_products", "avg_quality", "recent_activity"},
			},
		},
		"parameters": map[string]interface{}{
			"sections": map[string]interface{}{
				"type":        "string",
				"description": "Lista de secciones separadas por coma",
				"required":    true,
				"example":     "dashboard,taxonomy,brands",
			},
			"include_stats": map[string]interface{}{
				"type":        "boolean",
				"description": "Incluir estadísticas detalladas",
				"required":    false,
				"default":     false,
			},
			"parallel": map[string]interface{}{
				"type":        "boolean",
				"description": "Ejecutar consultas en paralelo",
				"required":    false,
				"default":     false,
			},
			"time_range_days": map[string]interface{}{
				"type":        "integer",
				"description": "Rango de días para métricas temporales",
				"required":    false,
				"default":     7,
				"min":         1,
				"max":         365,
			},
			"limit": map[string]interface{}{
				"type":        "integer",
				"description": "Límite para listas (top items, recent items)",
				"required":    false,
				"default":     10,
				"min":         1,
				"max":         100,
			},
		},
		"examples": []string{
			"/api/v1/marketplace/overview?sections=dashboard",
			"/api/v1/marketplace/overview?sections=dashboard,taxonomy,brands&include_stats=true",
			"/api/v1/marketplace/overview?sections=global-catalog,attributes&parallel=true&limit=20",
		},
	}

	c.JSON(http.StatusOK, sections)
}
