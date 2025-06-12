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

func TestMarketplaceIntegration_WithMiddlewares(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Configurar router con middlewares
	router := gin.New()

	// Aplicar middlewares
	router.Use(controller.CORSMiddleware())
	router.Use(controller.MarketplaceAuthMiddleware())
	router.Use(controller.RequestValidationMiddleware())
	router.Use(controller.TenantValidationMiddleware())

	// Crear handlers con dependencias nil (solo para probar middlewares)
	categoryHandler := &controller.MarketplaceCategoryHandler{}
	mappingHandler := &controller.TenantCategoryMappingHandler{}

	// Registrar rutas manualmente para las pruebas
	api := router.Group("/api/v1")
	marketplace := api.Group("/marketplace")
	{
		// Rutas de administrador
		adminRoutes := marketplace.Group("")
		adminRoutes.Use(controller.AdminOnlyMiddleware())
		{
			adminRoutes.POST("/categories", categoryHandler.CreateMarketplaceCategory)
			adminRoutes.POST("/categories/validate-hierarchy", categoryHandler.ValidateCategoryHierarchy)
			adminRoutes.POST("/sync-changes", categoryHandler.SyncMarketplaceChanges)
		}

		// Rutas de tenant
		marketplace.GET("/taxonomy", categoryHandler.GetTenantTaxonomy)

		tenantRoutes := marketplace.Group("/tenant")
		{
			tenantRoutes.POST("/category-mappings", mappingHandler.MapTenantCategory)
			tenantRoutes.PUT("/category-mappings/:mapping_id", mappingHandler.UpdateTenantCategoryMapping)
			tenantRoutes.DELETE("/category-mappings/:mapping_id", mappingHandler.DeleteTenantCategoryMapping)
		}
	}

	t.Run("Middleware de autorización", func(t *testing.T) {
		t.Run("debería fallar sin X-User-Role", func(t *testing.T) {
			// Arrange
			reqBody := request.CreateMarketplaceCategoryRequest{
				Name: "Test Category",
				Slug: "test-category",
			}
			jsonBody, _ := json.Marshal(reqBody)

			// Act
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/marketplace/categories", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, http.StatusUnauthorized, w.Code)
		})

		t.Run("debería fallar con rol inválido", func(t *testing.T) {
			// Arrange
			reqBody := request.CreateMarketplaceCategoryRequest{
				Name: "Test Category",
				Slug: "test-category",
			}
			jsonBody, _ := json.Marshal(reqBody)

			// Act
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/marketplace/categories", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-User-Role", "invalid_role")
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, http.StatusForbidden, w.Code)
		})
	})

	t.Run("Middleware de administrador", func(t *testing.T) {
		t.Run("debería fallar con rol tenant_user", func(t *testing.T) {
			// Arrange
			reqBody := request.CreateMarketplaceCategoryRequest{
				Name: "Test Category",
				Slug: "test-category",
			}
			jsonBody, _ := json.Marshal(reqBody)

			// Act
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/marketplace/categories", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-User-Role", "tenant_user")
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, http.StatusForbidden, w.Code)
		})

		t.Run("debería pasar con rol marketplace_admin", func(t *testing.T) {
			// Arrange
			reqBody := request.CreateMarketplaceCategoryRequest{
				Name: "Test Category",
				Slug: "test-category",
			}
			jsonBody, _ := json.Marshal(reqBody)

			// Act
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/marketplace/categories", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-User-Role", "marketplace_admin")
			router.ServeHTTP(w, req)

			// Assert
			// No debería fallar por middleware (aunque falle por nil pointer en el handler)
			assert.NotEqual(t, http.StatusUnauthorized, w.Code)
			assert.NotEqual(t, http.StatusForbidden, w.Code)
		})
	})

	t.Run("Middleware de validación de tenant", func(t *testing.T) {
		t.Run("debería fallar sin X-Tenant-ID en rutas tenant", func(t *testing.T) {
			// Arrange
			reqBody := request.MapTenantCategoryRequest{
				CategoryID:            "cat-123",
				MarketplaceCategoryID: "mkt-cat-456",
				CustomName:            "Custom Name",
			}
			jsonBody, _ := json.Marshal(reqBody)

			// Act
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/marketplace/tenant/category-mappings", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-User-Role", "tenant_user")
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, http.StatusBadRequest, w.Code)

			var errorResponse map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
			assert.NoError(t, err)
			assert.Contains(t, errorResponse["error"], "X-Tenant-ID es obligatorio")
		})

		t.Run("debería fallar con UUID inválido", func(t *testing.T) {
			// Arrange
			reqBody := request.MapTenantCategoryRequest{
				CategoryID:            "cat-123",
				MarketplaceCategoryID: "mkt-cat-456",
				CustomName:            "Custom Name",
			}
			jsonBody, _ := json.Marshal(reqBody)

			// Act
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/marketplace/tenant/category-mappings", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-User-Role", "tenant_user")
			req.Header.Set("X-Tenant-ID", "invalid-uuid")
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, http.StatusBadRequest, w.Code)

			var errorResponse map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
			assert.NoError(t, err)
			assert.Contains(t, errorResponse["error"], "Formato de X-Tenant-ID no válido")
		})

		t.Run("debería pasar con UUID válido", func(t *testing.T) {
			// Arrange
			reqBody := request.MapTenantCategoryRequest{
				CategoryID:            "cat-123",
				MarketplaceCategoryID: "mkt-cat-456",
				CustomName:            "Custom Name",
			}
			jsonBody, _ := json.Marshal(reqBody)

			// Act
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/marketplace/tenant/category-mappings", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-User-Role", "tenant_user")
			req.Header.Set("X-Tenant-ID", "9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8")
			router.ServeHTTP(w, req)

			// Assert
			// No debería fallar por middleware (aunque falle por nil pointer en el handler)
			assert.NotEqual(t, http.StatusUnauthorized, w.Code)
			assert.NotEqual(t, http.StatusBadRequest, w.Code)
		})
	})

	t.Run("Middleware de validación de request", func(t *testing.T) {
		t.Run("debería fallar sin Content-Type para POST", func(t *testing.T) {
			// Arrange
			reqBody := request.CreateMarketplaceCategoryRequest{
				Name: "Test Category",
				Slug: "test-category",
			}
			jsonBody, _ := json.Marshal(reqBody)

			// Act
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/marketplace/categories", bytes.NewBuffer(jsonBody))
			req.Header.Set("X-User-Role", "marketplace_admin")
			// Sin Content-Type
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, http.StatusBadRequest, w.Code)

			var errorResponse map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
			assert.NoError(t, err)
			assert.Contains(t, errorResponse["error"], "Content-Type debe ser application/json")
		})

		t.Run("debería pasar para GET sin Content-Type", func(t *testing.T) {
			// Act
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/marketplace/taxonomy", nil)
			req.Header.Set("X-User-Role", "tenant_user")
			req.Header.Set("X-Tenant-ID", "9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8")
			router.ServeHTTP(w, req)

			// Assert
			// No debería fallar por middleware de validación de request
			assert.NotEqual(t, http.StatusBadRequest, w.Code)
		})
	})

	t.Run("CORS Middleware", func(t *testing.T) {
		t.Run("debería manejar OPTIONS request", func(t *testing.T) {
			// Act
			w := httptest.NewRecorder()
			req := httptest.NewRequest("OPTIONS", "/api/v1/marketplace/categories", nil)
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, http.StatusNoContent, w.Code)
			assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
		})

		t.Run("debería agregar headers CORS", func(t *testing.T) {
			// Act
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/marketplace/taxonomy", nil)
			req.Header.Set("X-User-Role", "tenant_user")
			req.Header.Set("X-Tenant-ID", "9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8")
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
			assert.Contains(t, w.Header().Get("Access-Control-Allow-Methods"), "GET")
			assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "X-Tenant-ID")
		})
	})
}
