package response

// MarketplaceProductResponse representa un producto visible en el marketplace (cross-tenant)
type MarketplaceProductResponse struct {
	ID           string                          `json:"id"`
	TenantID     string                          `json:"tenant_id"`
	Name         string                          `json:"name"`
	Description  *string                         `json:"description"`
	CategoryName *string                         `json:"category_name"`
	BrandName    *string                         `json:"brand_name"`
	ImageURL     *string                         `json:"image_url"`
	StoreType    *MarketplaceStoreTypeInfo       `json:"store_type,omitempty"`
	Variant      *MarketplaceProductVariantInfo   `json:"variant,omitempty"`
}

// MarketplaceStoreTypeInfo info del tipo de comercio
type MarketplaceStoreTypeInfo struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
}

// MarketplaceProductVariantInfo info de la variante default del producto
type MarketplaceProductVariantInfo struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	SKU   *string `json:"sku"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// MarketplaceProductListResponse lista paginada de productos del marketplace
type MarketplaceProductListResponse struct {
	Products   []*MarketplaceProductResponse `json:"products"`
	Total      int                           `json:"total"`
	Page       int                           `json:"page"`
	PageSize   int                           `json:"page_size"`
	TotalPages int                           `json:"total_pages"`
}

// MarketplaceStoreTypeResponse representa un business_type con conteos para el marketplace
type MarketplaceStoreTypeResponse struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Icon         string `json:"icon"`
	Color        string `json:"color"`
	StoreCount   int    `json:"store_count"`
	ProductCount int    `json:"product_count"`
}
