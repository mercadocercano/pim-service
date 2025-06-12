package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// MarketplaceAuthMiddleware valida la autorización para endpoints del marketplace
func MarketplaceAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validar que el header X-User-Role esté presente
		userRole := c.GetHeader("X-User-Role")
		if userRole == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Header X-User-Role es obligatorio"})
			c.Abort()
			return
		}

		// Validar roles válidos
		validRoles := map[string]bool{
			"super_admin":       true,
			"marketplace_admin": true,
			"tenant_admin":      true,
			"tenant_user":       true,
		}

		if !validRoles[userRole] {
			c.JSON(http.StatusForbidden, gin.H{"error": "Rol de usuario no válido"})
			c.Abort()
			return
		}

		// Almacenar el rol en el contexto para uso posterior
		c.Set("user_role", userRole)
		c.Next()
	}
}

// TenantValidationMiddleware valida que el tenant ID esté presente y sea válido
func TenantValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Solo aplicar validación a rutas que requieren tenant
		if strings.Contains(c.Request.URL.Path, "/tenant/") {
			tenantID := c.GetHeader("X-Tenant-ID")
			if tenantID == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Header X-Tenant-ID es obligatorio para operaciones de tenant"})
				c.Abort()
				return
			}

			// Validar formato UUID básico (36 caracteres con guiones)
			if len(tenantID) != 36 || !isValidUUIDFormat(tenantID) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de X-Tenant-ID no válido"})
				c.Abort()
				return
			}

			// Almacenar el tenant ID en el contexto
			c.Set("tenant_id", tenantID)
		}

		c.Next()
	}
}

// AdminOnlyMiddleware restringe el acceso solo a administradores
func AdminOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetHeader("X-User-Role")
		if userRole != "super_admin" && userRole != "marketplace_admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso restringido a administradores"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// RequestValidationMiddleware valida el formato de las peticiones JSON
func RequestValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Solo validar para métodos que envían datos
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			contentType := c.GetHeader("Content-Type")
			if !strings.Contains(contentType, "application/json") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type debe ser application/json"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

// CORSMiddleware maneja las políticas CORS para el marketplace
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Tenant-ID, X-User-Role")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// isValidUUIDFormat valida el formato básico de un UUID
func isValidUUIDFormat(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	// Verificar posiciones de los guiones: 8-4-4-4-12
	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	// Verificar que el resto sean caracteres hexadecimales
	for i, char := range uuid {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue // Saltar guiones
		}
		if !isHexChar(char) {
			return false
		}
	}

	return true
}

// isHexChar verifica si un carácter es hexadecimal válido
func isHexChar(char rune) bool {
	return (char >= '0' && char <= '9') ||
		(char >= 'a' && char <= 'f') ||
		(char >= 'A' && char <= 'F')
}
