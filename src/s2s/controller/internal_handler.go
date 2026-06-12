package controller

import (
	"errors"
	httpresp "github.com/hornosg/go-shared/infrastructure/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"saas-mt-pim-service/src/s2s/usecase"
)

type InternalHandler struct {
	refreshUC  *usecase.RefreshTemplateProductsUseCase
	templateUC *usecase.GetTemplateStatusUseCase
}

func NewInternalHandler(refreshUC *usecase.RefreshTemplateProductsUseCase, templateUC *usecase.GetTemplateStatusUseCase) *InternalHandler {
	return &InternalHandler{refreshUC: refreshUC, templateUC: templateUC}
}

func (h *InternalHandler) RegisterRoutes(router *gin.RouterGroup) {
	// Ruta legacy — mantenida para no romper integraciones existentes
	internal := router.Group("/internal")
	internal.POST("/refresh-template-products", h.deprecatedRefreshTemplateProducts)

	// Rutas S2S — autenticadas via API-Key en Kong, sin JWT
	s2s := router.Group("/s2s")
	s2s.POST("/refresh-template-products", h.RefreshTemplateProducts)
	s2s.GET("/business-types/:slug/template-status", h.GetTemplateStatus)
}

// deprecatedRefreshTemplateProducts mantiene compatibilidad con la ruta /internal legacy.
func (h *InternalHandler) deprecatedRefreshTemplateProducts(c *gin.Context) {
	log.Println("[DEPRECATED] /internal/refresh-template-products - use /s2s/refresh-template-products")
	h.RefreshTemplateProducts(c)
}

func (h *InternalHandler) RefreshTemplateProducts(c *gin.Context) {
	result, err := h.refreshUC.Execute(c.Request.Context())
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "template products refreshed from global_products",
		"data":    result,
	})
}

func (h *InternalHandler) GetTemplateStatus(c *gin.Context) {
	slug := c.Param("slug")
	result, err := h.templateUC.Execute(c.Request.Context(), slug)
	if err != nil {
		if errors.Is(err, usecase.ErrBusinessTypeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "business type not found"}})
			return
		}
		log.Printf("[template-status] error slug=%s: %v", slug, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "internal server error"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
