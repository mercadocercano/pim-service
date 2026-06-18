package controller

import (
	"errors"
	httpresp "github.com/hornosg/go-shared/infrastructure/response"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"saas-mt-pim-service/src/pim/domain/port"
	globalcataloguc "saas-mt-pim-service/src/product/global_catalog/application/usecase"
	globalcatalogvo "saas-mt-pim-service/src/product/global_catalog/domain/value_object"
	"saas-mt-pim-service/src/s2s/usecase"
)

type InternalHandler struct {
	refreshUC      *usecase.RefreshTemplateProductsUseCase
	templateUC     *usecase.GetTemplateStatusUseCase
	reclassifyUC   *globalcataloguc.ReclassifyBusinessTypesUseCase
	logger         port.PIMEventLogger
}

func NewInternalHandler(
	refreshUC *usecase.RefreshTemplateProductsUseCase,
	templateUC *usecase.GetTemplateStatusUseCase,
	logger port.PIMEventLogger,
) *InternalHandler {
	return &InternalHandler{refreshUC: refreshUC, templateUC: templateUC, logger: logger}
}

// NewInternalHandlerWithReclassify crea el handler con el use case de re-clasificación.
// Usado por setupInternalModuleWithLogger en main.go (F-015).
func NewInternalHandlerWithReclassify(
	refreshUC *usecase.RefreshTemplateProductsUseCase,
	templateUC *usecase.GetTemplateStatusUseCase,
	reclassifyUC *globalcataloguc.ReclassifyBusinessTypesUseCase,
	logger port.PIMEventLogger,
) *InternalHandler {
	return &InternalHandler{
		refreshUC:    refreshUC,
		templateUC:   templateUC,
		reclassifyUC: reclassifyUC,
		logger:       logger,
	}
}

func (h *InternalHandler) RegisterRoutes(router *gin.RouterGroup) {
	// Ruta legacy — mantenida para no romper integraciones existentes
	internal := router.Group("/internal")
	internal.POST("/refresh-template-products", h.deprecatedRefreshTemplateProducts)

	// Rutas S2S — autenticadas via API-Key en Kong, sin JWT
	s2s := router.Group("/s2s")
	s2s.POST("/refresh-template-products", h.RefreshTemplateProducts)
	s2s.GET("/business-types/:slug/template-status", h.GetTemplateStatus)
	s2s.POST("/global-products/reclassify-business-types", h.ReclassifyBusinessTypes)
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
		h.logger.Log(port.PIMEvent{
			Event:  "pim.template_status_error",
			Reason: err.Error(),
			SKU:    slug,
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_ERROR", "message": "internal server error"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// reclassifyRequest es el body del endpoint POST /s2s/global-products/reclassify-business-types.
// Contrato ADR-005 §2.
type reclassifyRequest struct {
	DryRun  *bool                   `json:"dry_run"`
	Confirm bool                    `json:"confirm"`
	Scope   reclassifyScopeRequest  `json:"scope"`
}

type reclassifyScopeRequest struct {
	SourcePrefix string `json:"source_prefix"`
	MaxRows      int    `json:"max_rows"`
}

// ReclassifyBusinessTypes maneja POST /api/v1/s2s/global-products/reclassify-business-types.
//
// Flujo L4/PR-4:
//   - Parsea body; dry_run default true si se omite.
//   - Extrae X-Operator-Id del header.
//   - Si la operación es APPLY real (dry_run=false AND confirm=true) y X-Operator-Id viene
//     vacío (o solo whitespace) → 400 MISSING_OPERATOR_ID ANTES de invocar el use case
//     (cero mutación si falta operador). NO es 401/403: eso lo cubre key-auth de Kong;
//     acá es validación de request.
//   - En dry_run la simulación es read-only y NO requiere operador (T-033 sigue dando 200).
//   - Llama al use case y mapea errores (400, 422, 500).
//
// TEST-IDs: T-032 (body inválido→400), T-033 (dry_run sin operador→200), T-034 (422 max_rows),
//           T-035 (dry_run válido→200), T-036 (apply exitoso→200), T-037 (apply sin operador→400).
func (h *InternalHandler) ReclassifyBusinessTypes(c *gin.Context) {
	if h.reclassifyUC == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": gin.H{
			"code":    "USE_CASE_NOT_CONFIGURED",
			"message": "reclassify use case not configured",
		}})
		return
	}

	// Parsear body
	var req reclassifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{
			"code":    "INVALID_BODY",
			"message": err.Error(),
		}})
		return
	}

	// dry_run default true si se omite en el JSON
	dryRun := true
	if req.DryRun != nil {
		dryRun = *req.DryRun
	}

	// Extraer X-Operator-Id (mismo criterio de apply que el use case: !DryRun && Confirm).
	operatorID := strings.TrimSpace(c.GetHeader("X-Operator-Id"))
	isApply := !dryRun && req.Confirm

	// L4/PR-4: un apply real exige operador humano. Rechazar TEMPRANO, antes de invocar el
	// use case → cero mutación si falta el operador. dry_run (read-only) no lo requiere.
	if isApply && operatorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{
			"code":    "MISSING_OPERATOR_ID",
			"message": "X-Operator-Id header is required for an apply (dry_run=false, confirm=true)",
		}})
		return
	}

	// MaxRows default: si viene 0 usar el cap completo
	maxRows := req.Scope.MaxRows
	if maxRows <= 0 {
		maxRows = globalcatalogvo.ReclassifyMaxRowsCap
	}

	// Validar y construir scope
	scope, err := globalcatalogvo.NewReclassifyScope(req.Scope.SourcePrefix, maxRows)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": gin.H{
			"code":    "SCOPE_INVALID",
			"message": err.Error(),
		}})
		return
	}

	ucReq := globalcataloguc.ReclassifyRequest{
		DryRun:     dryRun,
		Confirm:    req.Confirm,
		Scope:      scope,
		OperatorID: operatorID,
	}

	result, err := h.reclassifyUC.Execute(c.Request.Context(), ucReq)
	if err != nil {
		var reclErr *globalcataloguc.ReclassifyError
		if errors.As(err, &reclErr) {
			c.JSON(reclErr.HTTPStatus, gin.H{"error": gin.H{
				"code":    reclErr.Code,
				"message": reclErr.Message,
			}})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{
			"code":    "INTERNAL_ERROR",
			"message": "internal server error",
		}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Ensure httpresp is used (imported from existing code)
var _ = httpresp.JSON
