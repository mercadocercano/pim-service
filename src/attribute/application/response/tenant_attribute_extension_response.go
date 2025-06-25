package response

import "time"

// TenantAttributeExtensionResponse representa la respuesta de una extensión de atributo tenant
type TenantAttributeExtensionResponse struct {
	ID                    string                    `json:"id"`
	TenantID              string                    `json:"tenant_id"`
	MarketplaceCategoryID string                    `json:"marketplace_category_id"`
	CustomAttributes      []CustomAttributeResponse `json:"custom_attributes"`
	CreatedAt             time.Time                 `json:"created_at"`
	UpdatedAt             time.Time                 `json:"updated_at"`
}

// CustomAttributeResponse representa un atributo personalizado en la respuesta
type CustomAttributeResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	Type         string    `json:"type"`
	IsRequired   bool      `json:"is_required"`
	IsFilterable bool      `json:"is_filterable"`
	IsSearchable bool      `json:"is_searchable"`
	Options      []string  `json:"options,omitempty"`
	DefaultValue *string   `json:"default_value,omitempty"`
	SortOrder    int       `json:"sort_order"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TenantAttributeExtensionListResponse representa una lista de extensiones de atributos
type TenantAttributeExtensionListResponse struct {
	Extensions []*TenantAttributeExtensionResponse `json:"extensions"`
	Total      int                                 `json:"total"`
	Page       int                                 `json:"page"`
	PageSize   int                                 `json:"page_size"`
	TotalPages int                                 `json:"total_pages"`
}
