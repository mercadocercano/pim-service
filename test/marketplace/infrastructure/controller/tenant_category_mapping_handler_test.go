package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"pim/src/marketplace/application/request"
	"pim/src/marketplace/infrastructure/controller"
)

func TestTenantCategoryMappingHandler_MapTenantCategory_Validation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := &controller.TenantCategoryMappingHandler{}

	t.Run("debería fallar sin X-Tenant-ID", func(t *testing.T) {
		// Arrange
		reqBody := request.MapTenantCategoryRequest{
			CategoryID:            "tenant-cat-123",
			MarketplaceCategoryID: "marketplace-cat-456",
			CustomName:            "Nombre personalizado",
		}
		jsonBody, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/marketplace/tenant/category-mappings", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		// Act
		handler.MapTenantCategory(c)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var errorResponse map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		assert.Contains(t, errorResponse["error"], "X-Tenant-ID es obligatorio")
	})

	t.Run("debería fallar con JSON inválido", func(t *testing.T) {
		// Arrange
		invalidJSON := `{"category_id": "test", "marketplace_category_id":}`

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/marketplace/tenant/category-mappings", bytes.NewBufferString(invalidJSON))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("X-Tenant-ID", "9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8")

		// Act
		handler.MapTenantCategory(c)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var errorResponse map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		assert.Contains(t, errorResponse["error"], "Error en el formato")
	})

	t.Run("debería pasar validaciones básicas", func(t *testing.T) {
		// Arrange
		reqBody := request.MapTenantCategoryRequest{
			CategoryID:            "tenant-cat-123",
			MarketplaceCategoryID: "marketplace-cat-456",
			CustomName:            "Nombre personalizado",
		}
		jsonBody, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/marketplace/tenant/category-mappings", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("X-Tenant-ID", "9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8")

		// Act
		handler.MapTenantCategory(c)

		// Assert
		// No debería fallar por validación básica (aunque falle por falta de caso de uso)
		assert.NotEqual(t, http.StatusBadRequest, w.Code)
	})
}

func TestTenantCategoryMappingHandler_UpdateTenantCategoryMapping_Validation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := &controller.TenantCategoryMappingHandler{}

	t.Run("debería fallar sin X-Tenant-ID", func(t *testing.T) {
		// Arrange
		reqBody := request.MapTenantCategoryRequest{
			CategoryID:            "tenant-cat-123",
			MarketplaceCategoryID: "marketplace-cat-456",
			CustomName:            "Nombre personalizado actualizado",
		}
		jsonBody, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("PUT", "/marketplace/tenant/category-mappings/mapping-123", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Params = []gin.Param{{Key: "mapping_id", Value: "mapping-123"}}

		// Act
		handler.UpdateTenantCategoryMapping(c)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var errorResponse map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		assert.Contains(t, errorResponse["error"], "X-Tenant-ID es obligatorio")
	})

	t.Run("debería manejar parámetro mapping_id", func(t *testing.T) {
		// Arrange
		reqBody := request.MapTenantCategoryRequest{
			CategoryID:            "tenant-cat-123",
			MarketplaceCategoryID: "marketplace-cat-456",
			CustomName:            "Nombre personalizado actualizado",
		}
		jsonBody, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("PUT", "/marketplace/tenant/category-mappings/mapping-123", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("X-Tenant-ID", "9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8")
		c.Params = []gin.Param{{Key: "mapping_id", Value: "mapping-123"}}

		// Act
		handler.UpdateTenantCategoryMapping(c)

		// Assert
		// Debería retornar NotImplemented ya que no está implementado
		assert.Equal(t, http.StatusNotImplemented, w.Code)
	})

	t.Run("debería fallar sin mapping_id", func(t *testing.T) {
		// Arrange
		reqBody := request.MapTenantCategoryRequest{
			CategoryID:            "tenant-cat-123",
			MarketplaceCategoryID: "marketplace-cat-456",
			CustomName:            "Nombre personalizado actualizado",
		}
		jsonBody, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("PUT", "/marketplace/tenant/category-mappings/", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("X-Tenant-ID", "9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8")
		// Sin parámetro mapping_id

		// Act
		handler.UpdateTenantCategoryMapping(c)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var errorResponse map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		assert.Contains(t, errorResponse["error"], "mapping_id es obligatorio")
	})
}

func TestTenantCategoryMappingHandler_DeleteTenantCategoryMapping_Validation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := &controller.TenantCategoryMappingHandler{}

	t.Run("debería fallar sin X-Tenant-ID", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("DELETE", "/marketplace/tenant/category-mappings/mapping-123", nil)
		c.Params = []gin.Param{{Key: "mapping_id", Value: "mapping-123"}}

		// Act
		handler.DeleteTenantCategoryMapping(c)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var errorResponse map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		assert.Contains(t, errorResponse["error"], "X-Tenant-ID es obligatorio")
	})

	t.Run("debería manejar parámetro mapping_id para DELETE", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("DELETE", "/marketplace/tenant/category-mappings/mapping-123", nil)
		c.Request.Header.Set("X-Tenant-ID", "9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8")
		c.Params = []gin.Param{{Key: "mapping_id", Value: "mapping-123"}}

		// Act
		handler.DeleteTenantCategoryMapping(c)

		// Assert
		// Debería retornar NotImplemented ya que no está implementado
		assert.Equal(t, http.StatusNotImplemented, w.Code)
	})

	t.Run("debería fallar sin mapping_id", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("DELETE", "/marketplace/tenant/category-mappings/", nil)
		c.Request.Header.Set("X-Tenant-ID", "9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8")
		// Sin parámetro mapping_id

		// Act
		handler.DeleteTenantCategoryMapping(c)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var errorResponse map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		assert.Contains(t, errorResponse["error"], "mapping_id es obligatorio")
	})
}

// Helper function para crear punteros a string
func stringPtr(s string) *string {
	return &s
}
