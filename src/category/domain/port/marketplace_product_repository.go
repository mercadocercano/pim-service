package port

import (
	"context"

	"saas-mt-pim-service/src/category/application/response"
)

// MarketplaceProductRepository define operaciones de lectura cross-tenant para el marketplace
type MarketplaceProductRepository interface {
	// FindProductsByStoreType busca productos de tenants con un business_type específico
	FindProductsByStoreType(ctx context.Context, storeTypeCode string, page, pageSize int) ([]*response.MarketplaceProductResponse, int, error)

	// FindAllProducts busca todos los productos activos cross-tenant con filtros opcionales
	FindAllProducts(ctx context.Context, search, businessType string, page, pageSize int) ([]*response.MarketplaceProductResponse, int, error)

	// FindProductByID busca un producto por ID (cross-tenant)
	FindProductByID(ctx context.Context, productID string) (*response.MarketplaceProductResponse, error)

	// FindProductsByTenantID busca productos de un tenant específico
	FindProductsByTenantID(ctx context.Context, tenantID string, page, pageSize int) ([]*response.MarketplaceProductResponse, int, error)

	// GetStoreTypesWithCounts lista business_types con conteo de tiendas y productos
	GetStoreTypesWithCounts(ctx context.Context) ([]*response.MarketplaceStoreTypeResponse, error)
}
