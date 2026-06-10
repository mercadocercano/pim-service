package controller_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/s2s/controller"
	s2sport "saas-mt-pim-service/src/s2s/domain/port"
	"saas-mt-pim-service/src/s2s/usecase"
)

func init() {
	gin.SetMode(gin.TestMode)
}

// mockTemplateRepo implementa port.TemplateRepository para tests del handler.
type mockTemplateRepo struct {
	getStatusFn func(ctx context.Context, slug string) (*s2sport.TemplateStatusRow, error)
	refreshFn   func(ctx context.Context) (int64, error)
}

func (m *mockTemplateRepo) GetTemplateStatus(ctx context.Context, slug string) (*s2sport.TemplateStatusRow, error) {
	return m.getStatusFn(ctx, slug)
}

func (m *mockTemplateRepo) RefreshProductTemplates(ctx context.Context) (int64, error) {
	return m.refreshFn(ctx)
}

func setupRouter(repo s2sport.TemplateRepository) *gin.Engine {
	refreshUC := usecase.NewRefreshTemplateProductsUseCase(repo)
	templateUC := usecase.NewGetTemplateStatusUseCase(repo)
	h := controller.NewInternalHandler(refreshUC, templateUC)

	r := gin.New()
	v1 := r.Group("/api/v1")
	h.RegisterRoutes(v1)
	return r
}

// TestTemplateStatusHandler_ComputedSource_Returns200 verifica que cuando el use case
// retorna source=computed, el handler responde 200 con data.source="computed".
func TestTemplateStatusHandler_ComputedSource_Returns200(t *testing.T) {
	now := time.Now().UTC()
	repo := &mockTemplateRepo{
		getStatusFn: func(_ context.Context, _ string) (*s2sport.TemplateStatusRow, error) {
			return &s2sport.TemplateStatusRow{ComputedCount: 25, EditorialCount: 20, LastRefresh: &now}, nil
		},
	}

	r := setupRouter(repo)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/s2s/business-types/almacen/template-status", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var body map[string]interface{}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
	data := body["data"].(map[string]interface{})
	assert.Equal(t, "computed", data["source"])
	assert.Equal(t, float64(25), data["computed_count"])
}

// TestTemplateStatusHandler_EditorialSource_Returns200 verifica que cuando computed_count=0
// el handler responde 200 con data.source="editorial".
func TestTemplateStatusHandler_EditorialSource_Returns200(t *testing.T) {
	repo := &mockTemplateRepo{
		getStatusFn: func(_ context.Context, _ string) (*s2sport.TemplateStatusRow, error) {
			return &s2sport.TemplateStatusRow{ComputedCount: 0, EditorialCount: 15, LastRefresh: nil}, nil
		},
	}

	r := setupRouter(repo)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/s2s/business-types/relojeria/template-status", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var body map[string]interface{}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
	data := body["data"].(map[string]interface{})
	assert.Equal(t, "editorial", data["source"])
	assert.Nil(t, data["last_refresh"])
}

// TestTemplateStatusHandler_UnknownSlug_Returns404 verifica que un slug inexistente
// retorna 404 con error.code="NOT_FOUND".
func TestTemplateStatusHandler_UnknownSlug_Returns404(t *testing.T) {
	repo := &mockTemplateRepo{
		getStatusFn: func(_ context.Context, _ string) (*s2sport.TemplateStatusRow, error) {
			return nil, nil // nil,nil = no encontrado
		},
	}

	r := setupRouter(repo)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/s2s/business-types/no-existe/template-status", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var body map[string]interface{}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
	errObj := body["error"].(map[string]interface{})
	assert.Equal(t, "NOT_FOUND", errObj["code"])
}

// TestTemplateStatusHandler_DBError_Returns500 verifica que un error de repo genérico
// retorna 500 con error.code="INTERNAL_ERROR".
func TestTemplateStatusHandler_DBError_Returns500(t *testing.T) {
	repo := &mockTemplateRepo{
		getStatusFn: func(_ context.Context, _ string) (*s2sport.TemplateStatusRow, error) {
			return nil, errors.New("connection refused")
		},
	}

	r := setupRouter(repo)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/s2s/business-types/almacen/template-status", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var body map[string]interface{}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
	errObj := body["error"].(map[string]interface{})
	assert.Equal(t, "INTERNAL_ERROR", errObj["code"])
}
