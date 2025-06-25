package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"pim/src/api/infrastructure/middleware"
)

func TestMarketplaceAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("debería fallar sin header X-User-Role", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		middlewareFunc := middleware.MarketplaceAuthMiddleware()

		// Act
		middlewareFunc(c)

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

		middlewareFunc := middleware.MarketplaceAuthMiddleware()

		// Act
		middlewareFunc(c)

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

		middlewareFunc := middleware.MarketplaceAuthMiddleware()

		// Act
		middlewareFunc(c)

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

				middlewareFunc := middleware.MarketplaceAuthMiddleware()

				// Act
				middlewareFunc(c)

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

		middlewareFunc := middleware.TenantValidationMiddleware()

		// Act
		middlewareFunc(c)

		// Assert
		assert.False(t, c.IsAborted())
	})

	t.Run("debería fallar sin X-Tenant-ID en rutas tenant", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/marketplace/tenant/categories", nil)

		middlewareFunc := middleware.TenantValidationMiddleware()

		// Act
		middlewareFunc(c)

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

		middlewareFunc := middleware.TenantValidationMiddleware()

		// Act
		middlewareFunc(c)

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

		middlewareFunc := middleware.TenantValidationMiddleware()

		// Act
		middlewareFunc(c)

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

		middlewareFunc := middleware.AdminOnlyMiddleware()

		// Act
		middlewareFunc(c)

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

		middlewareFunc := middleware.AdminOnlyMiddleware()

		// Act
		middlewareFunc(c)

		// Assert
		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.True(t, c.IsAborted())
	})

	t.Run("debería pasar con rol super_admin", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/test", nil)
		c.Request.Header.Set("X-User-Role", "super_admin")

		middlewareFunc := middleware.AdminOnlyMiddleware()

		// Act
		middlewareFunc(c)

		// Assert
		assert.False(t, c.IsAborted())
	})

	t.Run("debería pasar con rol marketplace_admin", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/test", nil)
		c.Request.Header.Set("X-User-Role", "marketplace_admin")

		middlewareFunc := middleware.AdminOnlyMiddleware()

		// Act
		middlewareFunc(c)

		// Assert
		assert.False(t, c.IsAborted())
	})
}

func TestRequestValidationMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("debería fallar sin Content-Type en POST", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/test", nil)

		middlewareFunc := middleware.RequestValidationMiddleware()

		// Act
		middlewareFunc(c)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.True(t, c.IsAborted())
	})

	t.Run("debería pasar con Content-Type correcto", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/test", nil)
		c.Request.Header.Set("Content-Type", "application/json")

		middlewareFunc := middleware.RequestValidationMiddleware()

		// Act
		middlewareFunc(c)

		// Assert
		assert.False(t, c.IsAborted())
	})

	t.Run("debería pasar sin validación en GET", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		middlewareFunc := middleware.RequestValidationMiddleware()

		// Act
		middlewareFunc(c)

		// Assert
		assert.False(t, c.IsAborted())
	})
}

func TestCORSMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("debería configurar headers CORS correctamente", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		middlewareFunc := middleware.CORSMiddleware()

		// Act
		middlewareFunc(c)

		// Assert
		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Contains(t, w.Header().Get("Access-Control-Allow-Methods"), "GET")
		assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "X-Tenant-ID")
	})

	t.Run("debería manejar OPTIONS correctamente", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("OPTIONS", "/test", nil)

		middlewareFunc := middleware.CORSMiddleware()

		// Act
		middlewareFunc(c)

		// Assert
		assert.Equal(t, http.StatusNoContent, w.Code)
		assert.True(t, c.IsAborted())
	})
}
