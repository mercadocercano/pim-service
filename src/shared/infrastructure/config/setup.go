package config

import (
	apimiddleware "saas-mt-pim-service/src/api/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	sharedconfig "github.com/mercadocercano/go-shared/infrastructure/config"
	sharedmiddleware "github.com/mercadocercano/go-shared/infrastructure/middleware"
)

// SharedConfig is an alias for go-shared's SharedConfig, extended with pim defaults.
type SharedConfig = sharedconfig.SharedConfig

// DefaultSharedConfig returns pim-service-specific defaults.
func DefaultSharedConfig() SharedConfig {
	return SharedConfig{
		EnableGzip:            true,
		AlwaysTryDecompress:   true,
		ForceGzipCompression:  false,
		ForceGzipCheckSupport: true,
		ForceGzipPaths:        []string{"/pim/api/v1/products"},
		GzipExcludedPaths:     []string{"/health", "/metrics", "/api-docs", "/pim/api/v1/categories"},
	}
}

// SetupSharedMiddleware configures CORS (pim-specific) and the shared gzip/decompress middleware.
func SetupSharedMiddleware(router *gin.Engine, cfg SharedConfig) {
	router.Use(apimiddleware.CORSMiddleware())

	if cfg.AlwaysTryDecompress {
		router.Use(sharedmiddleware.GzipReader())
	}

	if cfg.EnableGzip {
		router.Use(sharedmiddleware.GzipMiddleware(sharedmiddleware.GzipOptions{
			ExcludedPaths: cfg.GzipExcludedPaths,
		}))

		if cfg.ForceGzipCompression && len(cfg.ForceGzipPaths) > 0 {
			opts := sharedmiddleware.ForceGzipOptions{CheckClientSupport: cfg.ForceGzipCheckSupport}
			for _, path := range cfg.ForceGzipPaths {
				router.Group(path).Use(sharedmiddleware.ForceGzipMiddleware(opts))
			}
		}
	}
}
