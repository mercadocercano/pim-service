package controller

import (
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
	internal := router.Group("/internal")
	internal.POST("/refresh-template-products", h.RefreshTemplateProducts)
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
