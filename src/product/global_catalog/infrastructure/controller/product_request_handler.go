package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"saas-mt-pim-service/src/product/global_catalog/application/usecase"
)

// ProductRequestController maneja solicitudes de productos no encontrados
type ProductRequestController struct {
	createProductRequest  *usecase.CreateProductRequestUseCase
	listProductRequests   *usecase.ListProductRequestsUseCase
	resolveProductRequest *usecase.ResolveProductRequestUseCase
}

func NewProductRequestController(
	createProductRequest *usecase.CreateProductRequestUseCase,
	listProductRequests *usecase.ListProductRequestsUseCase,
	resolveProductRequest *usecase.ResolveProductRequestUseCase,
) *ProductRequestController {
	return &ProductRequestController{
		createProductRequest:  createProductRequest,
		listProductRequests:   listProductRequests,
		resolveProductRequest: resolveProductRequest,
	}
}

// RegisterRoutes registra las rutas de product requests
func (ctrl *ProductRequestController) RegisterRoutes(router *gin.RouterGroup) {
	// Tenant crea solicitud
	router.POST("/global-catalog/product-requests", ctrl.CreateRequest)

	// Admin gestiona solicitudes
	admin := router.Group("/global-catalog/product-requests")
	{
		admin.GET("", ctrl.ListPending)
		admin.GET("/metrics", ctrl.GetMetrics)
		admin.PUT("/:id/resolve", ctrl.Resolve)
	}
}

// CreateRequest — POST /api/v1/global-catalog/product-requests
func (ctrl *ProductRequestController) CreateRequest(c *gin.Context) {
	var req usecase.CreateProductRequestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}
	req.TenantID = tenantID

	resp, err := ctrl.createProductRequest.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// ListPending — GET /api/v1/global-catalog/product-requests?limit=50&offset=0
func (ctrl *ProductRequestController) ListPending(c *gin.Context) {
	limit := 50
	offset := 0
	if v := c.Query("limit"); v != "" {
		if val, err := strconv.Atoi(v); err == nil && val > 0 {
			limit = val
		}
	}
	if v := c.Query("offset"); v != "" {
		if val, err := strconv.Atoi(v); err == nil && val >= 0 {
			offset = val
		}
	}

	resp, err := ctrl.listProductRequests.Execute(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetMetrics — GET /api/v1/global-catalog/product-requests/metrics
func (ctrl *ProductRequestController) GetMetrics(c *gin.Context) {
	resp, err := ctrl.listProductRequests.Execute(c.Request.Context(), 0, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pending_requests": resp.PendingCount,
		"alert":            resp.PendingCount > 0,
		"alert_level":      alertLevel(resp.PendingCount),
	})
}

// Resolve — PUT /api/v1/global-catalog/product-requests/:id/resolve
func (ctrl *ProductRequestController) Resolve(c *gin.Context) {
	var req usecase.ResolveProductRequestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	req.RequestID = c.Param("id")

	resp, err := ctrl.resolveProductRequest.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func alertLevel(count int) string {
	switch {
	case count >= 20:
		return "critical"
	case count >= 10:
		return "warning"
	case count > 0:
		return "info"
	default:
		return "none"
	}
}
