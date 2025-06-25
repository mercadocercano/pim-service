package response

import "time"

// TenantCategoryMappingResponse representa la respuesta de un mapeo de categoría tenant
type TenantCategoryMappingResponse struct {
	ID                      string    `json:"id"`
	TenantID                string    `json:"tenant_id"`
	MarketplaceCategoryID   string    `json:"marketplace_category_id"`
	MarketplaceCategoryName string    `json:"marketplace_category_name"`
	CustomName              string    `json:"custom_name"`
	IsActive                bool      `json:"is_active"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

// TenantCategoryMappingListResponse representa una lista de mapeos
type TenantCategoryMappingListResponse struct {
	Mappings   []*TenantCategoryMappingResponse `json:"mappings"`
	Total      int                              `json:"total"`
	Page       int                              `json:"page"`
	PageSize   int                              `json:"page_size"`
	TotalPages int                              `json:"total_pages"`
}
