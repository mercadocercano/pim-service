package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"saas-mt-pim-service/src/category/application/request"
	"saas-mt-pim-service/src/category/infrastructure/controller"
)

func TestMarketplaceCategoryHandler_CreateMarketplaceCategory_Authorization(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Crear un handler con dependencias nil para probar solo la autorización
	handler := &controller.MarketplaceCategoryHandler{}

	t.Run("debería fallar sin rol de administrador", func(t *testing.T) {
		// Arrange
		reqBody := request.CreateMarketplaceCategoryRequest{
			Name: "Electrónicos",
			Slug: "electronicos",
		}
		jsonBody, _ := json.Marshal(reqBody)

		// Crear contexto HTTP sin rol de admin
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/marketplace/categories", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("X-User-Role", "tenant_user") // No es admin

		// Act
		handler.CreateMarketplaceCategory(c)

		// Assert
		assert.Equal(t, http.StatusForbidden, w.Code)

		var errorResponse map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		assert.Contains(t, errorResponse["error"], "Solo administradores")
	})

	t.Run("debería fallar con JSON inválido", func(t *testing.T) {
		// Arrange
		// JSON inválido
		invalidJSON := `{"name": "Test", "slug":}`

		// Crear contexto HTTP
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/marketplace/categories", bytes.NewBufferString(invalidJSON))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("X-User-Role", "marketplace_admin")

		// Act
		handler.CreateMarketplaceCategory(c)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var errorResponse map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		assert.Contains(t, errorResponse["error"], "Error en el formato")
	})
}

// TestMarketplaceCategoryHandler_GetTenantTaxonomy_Validation - TEMPORALMENTE COMENTADO
// El método GetTenantTaxonomy está comentado en el handler, este test se reactiva cuando se implemente.
// Ver marketplace_category_handler.go línea 319

func TestMarketplaceCategoryHandler_ValidateCategoryHierarchy_Authorization(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := &controller.MarketplaceCategoryHandler{}

	t.Run("debería fallar sin rol de administrador", func(t *testing.T) {
		// Arrange
		reqBody := map[string]interface{}{
			"category_id": "cat-123",
		}
		jsonBody, _ := json.Marshal(reqBody)

		// Crear contexto HTTP sin rol de admin
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/marketplace/categories/validate-hierarchy", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("X-User-Role", "tenant_user")

		// Act
		handler.ValidateCategoryHierarchy(c)

		// Assert
		assert.Equal(t, http.StatusForbidden, w.Code)

		var errorResponse map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		assert.Contains(t, errorResponse["error"], "Solo administradores")
	})

	// TODO: Arreglar este test - requiere implementar casos de uso
	// t.Run("debería pasar autorización con super_admin", func(t *testing.T) {
	// 	// Arrange
	// 	reqBody := map[string]interface{}{
	// 		"category_id": "cat-123",
	// 	}
	// 	jsonBody, _ := json.Marshal(reqBody)

	// 	// Crear contexto HTTP con rol super_admin
	// 	w := httptest.NewRecorder()
	// 	c, _ := gin.CreateTestContext(w)
	// 	c.Request = httptest.NewRequest("POST", "/marketplace/categories/validate-hierarchy", bytes.NewBuffer(jsonBody))
	// 	c.Request.Header.Set("Content-Type", "application/json")
	// 	c.Request.Header.Set("X-User-Role", "super_admin")

	// 	// Act
	// 	handler.ValidateCategoryHierarchy(c)

	// 	// Assert
	// 	assert.NotEqual(t, http.StatusForbidden, w.Code)
	// })
}

// TestMarketplaceCategoryHandler_SyncMarketplaceChanges_Authorization - TEMPORALMENTE COMENTADO
// El método SyncMarketplaceChanges está comentado en el handler, este test se reactiva cuando se implemente.
// Ver marketplace_category_handler.go línea 381
