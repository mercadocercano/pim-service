package controller

import (
	httpresp "github.com/hornosg/go-shared/infrastructure/response"
	"net/http"
	"strconv"

	"saas-mt-pim-service/src/category/application/response"
	"saas-mt-pim-service/src/category/domain/port"

	"github.com/gin-gonic/gin"
)

// MarketplaceProductHandler maneja las peticiones HTTP de productos del marketplace (cross-tenant)
type MarketplaceProductHandler struct {
	productRepo port.MarketplaceProductRepository
}

// NewMarketplaceProductHandler crea una nueva instancia del handler
func NewMarketplaceProductHandler(productRepo port.MarketplaceProductRepository) *MarketplaceProductHandler {
	return &MarketplaceProductHandler{productRepo: productRepo}
}

// RegisterRoutes registra las rutas de productos del marketplace
func (h *MarketplaceProductHandler) RegisterRoutes(router *gin.RouterGroup) {
	marketplace := router.Group("/marketplace")
	{
		marketplace.GET("/products", h.ListProducts)
		marketplace.GET("/products/:id", h.GetProduct)
		marketplace.GET("/products/by-store-type/:code", h.ListProductsByStoreType)
		marketplace.GET("/products/by-tenant/:tenant_id", h.ListProductsByTenant)
		marketplace.GET("/store-types", h.ListStoreTypes)
	}
}

// ListProducts lista todos los productos del marketplace con búsqueda y paginación
func (h *MarketplaceProductHandler) ListProducts(c *gin.Context) {
	page, pageSize := parsePagination(c)
	search := c.Query("search")
	businessType := c.Query("business_type")

	products, total, err := h.productRepo.FindAllProducts(c.Request.Context(), search, businessType, page, pageSize)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, "Error al obtener productos: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, &response.MarketplaceProductListResponse{
		Products:   products,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: calculateTotalPages(total, pageSize),
	})
}

// GetProduct obtiene un producto por ID (cross-tenant)
func (h *MarketplaceProductHandler) GetProduct(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "ID de producto requerido")
		return
	}

	product, err := h.productRepo.FindProductByID(c.Request.Context(), productID)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, "Error al obtener producto: "+err.Error())
		return
	}
	if product == nil {
		httpresp.JSON(c, http.StatusNotFound, "Producto no encontrado")
		return
	}

	c.JSON(http.StatusOK, product)
}

// ListProductsByStoreType lista productos filtrados por tipo de comercio (business_type)
func (h *MarketplaceProductHandler) ListProductsByStoreType(c *gin.Context) {
	storeTypeCode := c.Param("code")
	if storeTypeCode == "" {
		httpresp.JSON(c, http.StatusBadRequest, "código de tipo de comercio requerido")
		return
	}

	page, pageSize := parsePagination(c)

	products, total, err := h.productRepo.FindProductsByStoreType(c.Request.Context(), storeTypeCode, page, pageSize)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, "Error al obtener productos: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, &response.MarketplaceProductListResponse{
		Products:   products,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: calculateTotalPages(total, pageSize),
	})
}

// ListProductsByTenant lista productos de un tenant específico
func (h *MarketplaceProductHandler) ListProductsByTenant(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	if tenantID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "ID de tenant requerido")
		return
	}

	page, pageSize := parsePagination(c)

	products, total, err := h.productRepo.FindProductsByTenantID(c.Request.Context(), tenantID, page, pageSize)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, "Error al obtener productos: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, &response.MarketplaceProductListResponse{
		Products:   products,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: calculateTotalPages(total, pageSize),
	})
}

// ListStoreTypes lista los tipos de comercio con conteos de tiendas y productos
func (h *MarketplaceProductHandler) ListStoreTypes(c *gin.Context) {
	storeTypes, err := h.productRepo.GetStoreTypesWithCounts(c.Request.Context())
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, "Error al obtener tipos de comercio: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"store_types": storeTypes,
		"total":       len(storeTypes),
	})
}

func parsePagination(c *gin.Context) (int, int) {
	page := 1
	pageSize := 20

	if p := c.Query("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			page = val
		}
	}
	if ps := c.Query("page_size"); ps != "" {
		if val, err := strconv.Atoi(ps); err == nil && val > 0 && val <= 100 {
			pageSize = val
		}
	}

	return page, pageSize
}

func calculateTotalPages(total, pageSize int) int {
	if pageSize <= 0 {
		return 0
	}
	return (total + pageSize - 1) / pageSize
}
