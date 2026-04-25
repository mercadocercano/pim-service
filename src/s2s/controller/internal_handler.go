package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"saas-mt-pim-service/src/s2s/usecase"
)

type InternalHandler struct {
	refreshUC *usecase.RefreshTemplateProductsUseCase
}

func NewInternalHandler(refreshUC *usecase.RefreshTemplateProductsUseCase) *InternalHandler {
	return &InternalHandler{refreshUC: refreshUC}
}

func (h *InternalHandler) RegisterRoutes(router *gin.RouterGroup) {
	// Ruta legacy — mantenida para no romper integraciones existentes
	internal := router.Group("/internal")
	internal.POST("/refresh-template-products", h.deprecatedRefreshTemplateProducts)

	// Ruta S2S — autenticada via API-Key en Kong, sin JWT
	s2s := router.Group("/s2s")
	s2s.POST("/refresh-template-products", h.RefreshTemplateProducts)
}

// deprecatedRefreshTemplateProducts mantiene compatibilidad con la ruta /internal legacy.
func (h *InternalHandler) deprecatedRefreshTemplateProducts(c *gin.Context) {
	log.Println("[DEPRECATED] /internal/refresh-template-products - use /s2s/refresh-template-products")
	h.RefreshTemplateProducts(c)
}

func (h *InternalHandler) RefreshTemplateProducts(c *gin.Context) {
	result, err := h.refreshUC.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "template products refreshed from global_products",
		"data":    result,
	})
}
