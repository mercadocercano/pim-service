package controller

import (
	httpresp "github.com/hornosg/go-shared/infrastructure/response"
	"net/http"

	"github.com/gin-gonic/gin"

	"saas-mt-pim-service/src/product/tenant/application/usecase"
)

// BulkImportController maneja las peticiones de importación bulk de productos
type BulkImportController struct {
	bulkImportProductsUseCase *usecase.BulkImportProductsUseCase
}

// NewBulkImportController crea una nueva instancia del controller
func NewBulkImportController(
	bulkImportProductsUseCase *usecase.BulkImportProductsUseCase,
) *BulkImportController {
	return &BulkImportController{
		bulkImportProductsUseCase: bulkImportProductsUseCase,
	}
}

// BulkImportProducts godoc
// @Summary Importación masiva de productos desde JSON
// @Description Crea múltiples productos en una sola petición desde datos JSON
// @Tags products
// @Accept json
// @Produce json
// @Param products body usecase.BulkImportProductsRequest true "Lista de productos a importar"
// @Success 201 {object} usecase.BulkImportProductsResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/import [post]
// @Security BearerAuth
func (ctrl *BulkImportController) BulkImportProducts(c *gin.Context) {
	// Obtener tenant ID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "X-Tenant-ID header es requerido")
		return
	}

	var req usecase.BulkImportProductsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresp.JSONWithDetails(c, http.StatusBadRequest, "Datos de entrada inválidos", err.Error())
		return
	}

	// Asignar tenant ID del header
	req.TenantID = tenantID

	// Ejecutar importación
	response, err := ctrl.bulkImportProductsUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Retornar respuesta con status apropiado
	statusCode := http.StatusCreated
	if !response.Success {
		statusCode = http.StatusPartialContent // 206 si algunos fallaron
	}

	c.JSON(statusCode, response)
}

// RegisterRoutes registra las rutas del controller
func (ctrl *BulkImportController) RegisterRoutes(router *gin.RouterGroup) {
	products := router.Group("/products")
	{
		products.POST("/import", ctrl.BulkImportProducts)
	}
}
