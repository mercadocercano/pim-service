package request

// BatchCreateRequest representa la solicitud para crear múltiples entidades
type BatchCreateRequest struct {
	Categories []CategoryBatchItem `json:"categories"`
	Brands     []BrandBatchItem    `json:"brands"`
	Products   []ProductBatchItem  `json:"products"`
}

// CategoryBatchItem representa una categoría para crear en batch
type CategoryBatchItem struct {
	Name        string              `json:"name" binding:"required"`
	Description string              `json:"description"`
	ParentID    *string             `json:"parent_id"`
	Mapping     *CategoryMappingData `json:"marketplace_mapping"`
}

// CategoryMappingData datos de mapeo con marketplace
type CategoryMappingData struct {
	MarketplaceCategoryID string `json:"marketplace_category_id" binding:"required"`
	CustomName            string `json:"custom_name"`
}

// BrandBatchItem representa una marca para crear en batch
type BrandBatchItem struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	LogoURL     string `json:"logo_url"`
	Website     string `json:"website"`
}

// ProductBatchItem representa un producto para crear en batch
type ProductBatchItem struct {
	Name        string                 `json:"name" binding:"required"`
	Description string                 `json:"description"`
	SKU         string                 `json:"sku" binding:"required"`
	Price       float64                `json:"price" binding:"required,min=0"`
	CategoryID  string                 `json:"category_id"`
	BrandID     string                 `json:"brand_id"`
	Stock       int                    `json:"stock"`
	Attributes  map[string]interface{} `json:"attributes"`
}