package controller_test

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/s2s/controller"
	"saas-mt-pim-service/src/s2s/usecase"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouter(db *sql.DB) *gin.Engine {
	refreshUC := usecase.NewRefreshTemplateProductsUseCase(db)
	templateUC := usecase.NewGetTemplateStatusUseCase(db)
	h := controller.NewInternalHandler(refreshUC, templateUC)

	r := gin.New()
	v1 := r.Group("/api/v1")
	h.RegisterRoutes(v1)
	return r
}

// TestTemplateStatusHandler_ComputedSource_Returns200 verifica que cuando el use case
// retorna source=computed, el handler responde 200 con data.source="computed".
func TestTemplateStatusHandler_ComputedSource_Returns200(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{"computed_count", "editorial_count", "last_refresh"}).
		AddRow(25, 20, sql.NullTime{Time: now, Valid: true})
	mock.ExpectQuery(`business_type_product_templates`).
		WithArgs("almacen").
		WillReturnRows(rows)

	r := setupRouter(db)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/s2s/business-types/almacen/template-status", nil)

	// Act
	r.ServeHTTP(w, req)

	// Assert
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
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"computed_count", "editorial_count", "last_refresh"}).
		AddRow(0, 15, sql.NullTime{Valid: false})
	mock.ExpectQuery(`business_type_product_templates`).
		WithArgs("relojeria").
		WillReturnRows(rows)

	r := setupRouter(db)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/s2s/business-types/relojeria/template-status", nil)

	// Act
	r.ServeHTTP(w, req)

	// Assert
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
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`business_type_product_templates`).
		WithArgs("no-existe").
		WillReturnError(sql.ErrNoRows)

	r := setupRouter(db)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/s2s/business-types/no-existe/template-status", nil)

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
	var body map[string]interface{}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
	errObj := body["error"].(map[string]interface{})
	assert.Equal(t, "NOT_FOUND", errObj["code"])
}

// TestTemplateStatusHandler_DBError_Returns500 verifica que un error de DB genérico
// retorna 500 con error.code="INTERNAL_ERROR".
func TestTemplateStatusHandler_DBError_Returns500(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`business_type_product_templates`).
		WithArgs("almacen").
		WillReturnError(assert.AnError)

	r := setupRouter(db)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/s2s/business-types/almacen/template-status", nil)

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var body map[string]interface{}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
	errObj := body["error"].(map[string]interface{})
	assert.Equal(t, "INTERNAL_ERROR", errObj["code"])
}
