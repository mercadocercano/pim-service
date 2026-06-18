package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"saas-mt-pim-service/src/product/global_catalog/application/usecase"
	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
	"saas-mt-pim-service/src/s2s/controller"
	s2sUsecase "saas-mt-pim-service/src/s2s/usecase"
)

func init() {
	gin.SetMode(gin.TestMode)
}

// --- Mock del repositorio que implementa domainport.ReclassifyRepository ---

type testReclassifyRepo struct {
	// applyCalled registra si el repo recibió una mutación (ApplyInTransaction).
	// Permite afirmar "cero mutación" en tests de rechazo de borde.
	applyCalled bool
}

func (t *testReclassifyRepo) CountCandidates(_ context.Context, _ value_object.ReclassifyScope) (int, error) {
	return 0, nil
}
func (t *testReclassifyRepo) FetchCandidates(_ context.Context, _ value_object.ReclassifyScope) ([]value_object.ReclassifyCandidate, error) {
	return nil, nil
}
func (t *testReclassifyRepo) ApplyInTransaction(_ context.Context, snapshotName string, _ []string, _ []value_object.ReclassifyUpdate) (string, int, int, error) {
	t.applyCalled = true
	return "global_products_bkp_" + snapshotName, 0, 0, nil
}
func (t *testReclassifyRepo) SaveAudit(_ context.Context, _ value_object.ReclassifyAuditRow) error {
	return nil
}

// buildTestUC construye un use case real con repo mock para tests del handler.
func buildTestUC() *usecase.ReclassifyBusinessTypesUseCase {
	return usecase.NewReclassifyBusinessTypesUseCase(&testReclassifyRepo{}, nil)
}

// buildTestUCWithRepo construye el use case sobre un repo provisto, para inspeccionarlo luego.
func buildTestUCWithRepo(repo *testReclassifyRepo) *usecase.ReclassifyBusinessTypesUseCase {
	return usecase.NewReclassifyBusinessTypesUseCase(repo, nil)
}

// setupTestRouter crea un router de prueba con el handler configurado.
func setupTestRouterWithUC(reclassifyUC *usecase.ReclassifyBusinessTypesUseCase) *gin.Engine {
	router := gin.New()
	v1 := router.Group("/api/v1")

	h := controller.NewInternalHandlerWithReclassify(
		(*s2sUsecase.RefreshTemplateProductsUseCase)(nil),
		(*s2sUsecase.GetTemplateStatusUseCase)(nil),
		reclassifyUC,
		nil,
	)
	h.RegisterRoutes(v1)
	return router
}

// makeReclassifyRequest envía POST al endpoint y devuelve el recorder.
func makeReclassifyRequest(router *gin.Engine, body interface{}, headers map[string]string) *httptest.ResponseRecorder {
	bodyBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/s2s/global-products/reclassify-business-types", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// T-032: body JSON malformado → 400.
func TestReclassifyHandler_InvalidBody_Returns400(t *testing.T) {
	router := setupTestRouterWithUC(buildTestUC())

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/s2s/global-products/reclassify-business-types", bytes.NewBufferString("{invalid json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("T-032: expected 400, got %d — body: %s", w.Code, w.Body.String())
	}
}

// T-033: dry_run=true sin X-Operator-Id → 200 (dry_run no requiere operador en L3).
func TestReclassifyHandler_DryRunNoOperator_Returns200(t *testing.T) {
	router := setupTestRouterWithUC(buildTestUC())

	body := map[string]interface{}{
		"dry_run": true,
		"scope":   map[string]interface{}{"source_prefix": "scraper", "max_rows": 100},
	}
	w := makeReclassifyRequest(router, body, nil)

	if w.Code != http.StatusOK {
		t.Errorf("T-033: expected 200, got %d — body: %s", w.Code, w.Body.String())
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("T-033: unmarshal response: %v", err)
	}
	data, ok := resp["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("T-033: expected 'data' key in response, got: %s", w.Body.String())
	}
	if data["mode"] != "dry_run" {
		t.Errorf("T-033: expected mode=dry_run, got %v", data["mode"])
	}
}

// T-034: scope con max_rows > cap (50001) → 422 (scope inválido: excede cap).
func TestReclassifyHandler_InvalidScope_Returns422(t *testing.T) {
	router := setupTestRouterWithUC(buildTestUC())

	// max_rows > ReclassifyMaxRowsCap → falla en construcción del scope → 422
	body := map[string]interface{}{
		"dry_run": true,
		"scope":   map[string]interface{}{"source_prefix": "scraper", "max_rows": value_object.ReclassifyMaxRowsCap + 1},
	}
	w := makeReclassifyRequest(router, body, nil)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("T-034: expected 422, got %d — body: %s", w.Code, w.Body.String())
	}
}

// T-035: dry_run=true con body válido → 200 con summary y mode=dry_run.
func TestReclassifyHandler_DryRunValid_Returns200WithSummary(t *testing.T) {
	router := setupTestRouterWithUC(buildTestUC())

	body := map[string]interface{}{
		"dry_run": true,
		"scope":   map[string]interface{}{"source_prefix": "scraper", "max_rows": 1000},
	}
	w := makeReclassifyRequest(router, body, map[string]string{"X-Operator-Id": "admin-123"})

	if w.Code != http.StatusOK {
		t.Errorf("T-035: expected 200, got %d — body: %s", w.Code, w.Body.String())
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("T-035: unmarshal response: %v", err)
	}
	data, ok := resp["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("T-035: expected 'data' key in response, got: %s", w.Body.String())
	}
	if data["mode"] != "dry_run" {
		t.Errorf("T-035: expected mode=dry_run, got %v", data["mode"])
	}
	// snapshot_ref debe ser null en dry_run
	if data["snapshot_ref"] != nil {
		t.Errorf("T-035: expected snapshot_ref=null in dry_run, got %v", data["snapshot_ref"])
	}
}

// T-036: apply exitoso → 200 con mode=applied.
func TestReclassifyHandler_ApplySucceeds_Returns200Applied(t *testing.T) {
	router := setupTestRouterWithUC(buildTestUC())

	body := map[string]interface{}{
		"dry_run": false,
		"confirm": true,
		"scope":   map[string]interface{}{"source_prefix": "scraper", "max_rows": 1000},
	}
	w := makeReclassifyRequest(router, body, map[string]string{"X-Operator-Id": "admin-123"})

	if w.Code != http.StatusOK {
		t.Errorf("T-036: expected 200, got %d — body: %s", w.Code, w.Body.String())
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("T-036: unmarshal response: %v", err)
	}
	data, ok := resp["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("T-036: expected 'data' key in response, got: %s", w.Body.String())
	}
	if data["mode"] != "applied" {
		t.Errorf("T-036: expected mode=applied, got %v", data["mode"])
	}
}

// T-037: apply (dry_run=false, confirm=true) SIN X-Operator-Id → 400 MISSING_OPERATOR_ID
// y cero mutación (el use case / repo no debe ejecutarse). Gemelo L4 de T-033.
func TestReclassifyHandler_ApplyNoOperator_Returns400AndNoMutation(t *testing.T) {
	repo := &testReclassifyRepo{}
	router := setupTestRouterWithUC(buildTestUCWithRepo(repo))

	body := map[string]interface{}{
		"dry_run": false,
		"confirm": true,
		"scope":   map[string]interface{}{"source_prefix": "scraper", "max_rows": 1000},
	}
	// Sin header X-Operator-Id.
	w := makeReclassifyRequest(router, body, nil)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("T-037: expected 400, got %d — body: %s", w.Code, w.Body.String())
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("T-037: unmarshal response: %v", err)
	}
	errObj, ok := resp["error"].(map[string]interface{})
	if !ok {
		t.Fatalf("T-037: expected 'error' key in response, got: %s", w.Body.String())
	}
	if errObj["code"] != "MISSING_OPERATOR_ID" {
		t.Errorf("T-037: expected code=MISSING_OPERATOR_ID, got %v", errObj["code"])
	}

	// Cero mutación: el rechazo es temprano, el repo nunca debe haber aplicado.
	if repo.applyCalled {
		t.Errorf("T-037: expected NO mutation (ApplyInTransaction must not be called) on apply without operator")
	}
}

// T-037b: apply con X-Operator-Id solo-whitespace → 400 MISSING_OPERATOR_ID (whitespace == vacío).
func TestReclassifyHandler_ApplyWhitespaceOperator_Returns400(t *testing.T) {
	repo := &testReclassifyRepo{}
	router := setupTestRouterWithUC(buildTestUCWithRepo(repo))

	body := map[string]interface{}{
		"dry_run": false,
		"confirm": true,
		"scope":   map[string]interface{}{"source_prefix": "scraper", "max_rows": 1000},
	}
	w := makeReclassifyRequest(router, body, map[string]string{"X-Operator-Id": "   "})

	if w.Code != http.StatusBadRequest {
		t.Fatalf("T-037b: expected 400, got %d — body: %s", w.Code, w.Body.String())
	}
	if repo.applyCalled {
		t.Errorf("T-037b: expected NO mutation on apply with whitespace-only operator")
	}
}
