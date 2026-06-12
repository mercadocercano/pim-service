package controller

import (
	httpresp "github.com/hornosg/go-shared/infrastructure/response"
	"net/http"

	"github.com/gin-gonic/gin"

	"saas-mt-pim-service/src/product/tenant/application/usecase"
)

type BulkUpdateController struct {
	bulkUpdateUseCase *usecase.BulkUpdateProductsUseCase
}

func NewBulkUpdateController(bulkUpdateUseCase *usecase.BulkUpdateProductsUseCase) *BulkUpdateController {
	return &BulkUpdateController{bulkUpdateUseCase: bulkUpdateUseCase}
}

func (ctrl *BulkUpdateController) BulkUpdateProducts(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header is required")
		return
	}

	var req usecase.BulkUpdateProductsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
		return
	}
	req.TenantID = tenantID

	result, err := ctrl.bulkUpdateUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, err.Error())
		return
	}

	status := http.StatusOK
	if !result.Success {
		status = http.StatusUnprocessableEntity
	}
	c.JSON(status, result)
}

func (ctrl *BulkUpdateController) RegisterRoutes(router *gin.RouterGroup) {
	router.PATCH("/products/bulk-update", ctrl.BulkUpdateProducts)
}
