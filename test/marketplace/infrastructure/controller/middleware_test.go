package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"pim/src/marketplace/infrastructure/controller"
)

func TestMarketplaceAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("debería fallar sin header X-User-Role", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		middleware := controller.MarketplaceAuthMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.True(t, c.IsAborted())
	})

	t.Run("debería fallar con rol inválido", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)
		c.Request.Header.Set("X-User-Role", "invalid_role")

		middleware := controller.MarketplaceAuthMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.True(t, c.IsAborted())
	})

	t.Run("debería pasar con rol válido", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)
		c.Request.Header.Set("X-User-Role", "marketplace_admin")

		middleware := controller.MarketplaceAuthMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code) // No se abortó
		assert.False(t, c.IsAborted())

		// Verificar que el rol se almacenó en el contexto
		userRole, exists := c.Get("user_role")
		assert.True(t, exists)
		assert.Equal(t, "marketplace_admin", userRole)
	})

	t.Run("debería aceptar todos los roles válidos", func(t *testing.T) {
		validRoles := []string{"super_admin", "marketplace_admin", "tenant_admin", "tenant_user"}

		for _, role := range validRoles {
			t.Run("rol_"+role, func(t *testing.T) {
				// Arrange
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest("GET", "/test", nil)
				c.Request.Header.Set("X-User-Role", role)

				middleware := controller.MarketplaceAuthMiddleware()

				// Act
				middleware(c)

				// Assert
				assert.False(t, c.IsAborted(), "El rol %s debería ser válido", role)
			})
		}
	})
}

func TestTenantValidationMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("debería pasar sin validación para rutas no-tenant", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/marketplace/categories", nil)

		middleware := controller.TenantValidationMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.False(t, c.IsAborted())
	})

	t.Run("debería fallar sin X-Tenant-ID en rutas tenant", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/marketplace/tenant/categories", nil)

		middleware := controller.TenantValidationMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.True(t, c.IsAborted())
	})

	t.Run("debería fallar con formato UUID inválido", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/marketplace/tenant/categories", nil)
		c.Request.Header.Set("X-Tenant-ID", "invalid-uuid")

		middleware := controller.TenantValidationMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.True(t, c.IsAborted())
	})

	t.Run("debería pasar con UUID válido", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/marketplace/tenant/categories", nil)
		c.Request.Header.Set("X-Tenant-ID", "9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8")

		middleware := controller.TenantValidationMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.False(t, c.IsAborted())

		// Verificar que el tenant ID se almacenó en el contexto
		tenantID, exists := c.Get("tenant_id")
		assert.True(t, exists)
		assert.Equal(t, "9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8", tenantID)
	})
}

func TestAdminOnlyMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("debería fallar con rol tenant_user", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/test", nil)
		c.Request.Header.Set("X-User-Role", "tenant_user")

		middleware := controller.AdminOnlyMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.True(t, c.IsAborted())
	})

	t.Run("debería fallar con rol tenant_admin", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/test", nil)
		c.Request.Header.Set("X-User-Role", "tenant_admin")

		middleware := controller.AdminOnlyMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.True(t, c.IsAborted())
	})

	t.Run("debería pasar con rol marketplace_admin", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/test", nil)
		c.Request.Header.Set("X-User-Role", "marketplace_admin")

		middleware := controller.AdminOnlyMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.False(t, c.IsAborted())
	})

	t.Run("debería pasar con rol super_admin", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/test", nil)
		c.Request.Header.Set("X-User-Role", "super_admin")

		middleware := controller.AdminOnlyMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.False(t, c.IsAborted())
	})
}

func TestRequestValidationMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("debería pasar para métodos GET", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		middleware := controller.RequestValidationMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.False(t, c.IsAborted())
	})

	t.Run("debería fallar sin Content-Type para POST", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/test", nil)

		middleware := controller.RequestValidationMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.True(t, c.IsAborted())
	})

	t.Run("debería pasar con Content-Type correcto para POST", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/test", nil)
		c.Request.Header.Set("Content-Type", "application/json")

		middleware := controller.RequestValidationMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.False(t, c.IsAborted())
	})
}

func TestCORSMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("debería agregar headers CORS", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		middleware := controller.CORSMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.False(t, c.IsAborted())
		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Contains(t, w.Header().Get("Access-Control-Allow-Methods"), "GET")
		assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "X-Tenant-ID")
	})

	t.Run("debería manejar OPTIONS request", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("OPTIONS", "/test", nil)

		middleware := controller.CORSMiddleware()

		// Act
		middleware(c)

		// Assert
		assert.True(t, c.IsAborted())
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
