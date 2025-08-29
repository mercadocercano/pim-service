package controller

import (
	"net/http"
	"time"

	"saas-mt-pim-service/src/batch/application/request"
	"saas-mt-pim-service/src/batch/application/usecase"
	"saas-mt-pim-service/src/shared/infrastructure/metrics"

	"github.com/gin-gonic/gin"
)

// BatchController maneja las operaciones batch
type BatchController struct {
	batchCreateUseCase *usecase.BatchCreateUseCase
}

// NewBatchController crea una nueva instancia del controlador
func NewBatchController(batchCreateUseCase *usecase.BatchCreateUseCase) *BatchController {
	return &BatchController{
		batchCreateUseCase: batchCreateUseCase,
	}
}

// RegisterRoutes registra las rutas del controlador batch
func (c *BatchController) RegisterRoutes(router *gin.RouterGroup) {
	batch := router.Group("/batch")
	{
		batch.POST("/create", c.BatchCreate)
	}
}

// BatchCreate maneja la creación masiva de entidades
// @Summary Crear múltiples entidades en una transacción
// @Description Crea categorías, marcas y productos de forma transaccional
// @Tags batch
// @Accept json
// @Produce json
// @Param X-Tenant-ID header string true "ID del tenant"
// @Param body body request.BatchCreateRequest true "Datos de las entidades a crear"
// @Success 200 {object} response.BatchCreateResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /batch/create [post]
// @Security BearerAuth
func (c *BatchController) BatchCreate(ctx *gin.Context) {
	startTime := time.Now()
	
	// Obtener tenant ID
	tenantID := ctx.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Parsear request
	var req request.BatchCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos", "details": err.Error()})
		return
	}

	// Registrar tamaño de batch por tipo de entidad
	if len(req.Categories) > 0 {
		metrics.RecordBatchOperation(tenantID, "categories", len(req.Categories))
	}
	if len(req.Brands) > 0 {
		metrics.RecordBatchOperation(tenantID, "brands", len(req.Brands))
	}
	if len(req.Products) > 0 {
		metrics.RecordBatchOperation(tenantID, "products", len(req.Products))
	}

	// Ejecutar caso de uso
	result, err := c.batchCreateUseCase.Execute(ctx.Request.Context(), &req, tenantID)
	if err != nil {
		metrics.ImportOperationsTotal.WithLabelValues(tenantID, "batch_create", "failure").Inc()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Determinar código de respuesta
	statusCode := http.StatusOK
	totalCreated := len(result.Created.Categories) + len(result.Created.Brands) + len(result.Created.Products)
	totalFailed := len(result.Errors)
	
	if totalCreated == 0 {
		statusCode = http.StatusBadRequest
	}

	// Registrar métricas
	duration := time.Since(startTime).Seconds()
	success := totalFailed == 0 && err == nil
	totalProcessed := len(req.Categories) + len(req.Brands) + len(req.Products)
	
	metrics.RecordImportMetrics(
		tenantID,
		"batch_create",
		success,
		totalProcessed,
		totalCreated,
		totalFailed,
		duration,
	)

	ctx.JSON(statusCode, result)
}